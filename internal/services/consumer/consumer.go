package consumer

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/mailru/easyjson"
	"github.com/oschwald/geoip2-golang"
	"github.com/segmentio/kafka-go"

	"github.com/zikwall/app_metrica/config"
	"github.com/zikwall/app_metrica/pkg/domain/event"
	"github.com/zikwall/app_metrica/pkg/log"
)

type EventRepository interface {
	AddEvent(
		ctx context.Context,
		event *event.EventExtended,
	) error
}

type Metrics interface {
	IncConsumerError(err error)
}

type Consumer struct {
	city            *geoip2.Reader
	asn             *geoip2.Reader
	eventRepository EventRepository

	wg  *sync.WaitGroup
	opt *config.Config

	metrics Metrics

	queue chan kafka.Message
}

func (c *Consumer) Run(ctx context.Context) {
	for partition := 0; partition < c.opt.Internal.ConsumerPerInstanceSize; partition++ {
		c.wg.Add(1)
		go func(p int) {
			defer c.wg.Done()

			c.proc(ctx, p)
		}(partition)
	}

	for i := 1; i <= c.opt.Server.Internal.ConsumerQueueHandlerSize; i++ {
		c.wg.Add(1)
		go func(number int) {
			defer c.wg.Done()

			c.handle(ctx, number)
		}(i)
	}

	c.wg.Wait()
	close(c.queue)
}

func (c *Consumer) proc(ctx context.Context, partition int) {
	log.Infof("run consumer for partition: %d", partition)
	defer log.Infof("stop consumer for partition: %d", partition)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     c.opt.KafkaReader.Brokers,
		GroupID:     c.opt.KafkaReader.GroupID,
		GroupTopics: c.opt.KafkaReader.GroupTopics,
		Topic:       c.opt.KafkaReader.Topic,
		// Partition:        partition,
		QueueCapacity:    c.opt.KafkaReader.QueueCapacity,
		MinBytes:         c.opt.KafkaReader.MinBytes,
		MaxBytes:         c.opt.KafkaReader.MaxBytes,
		MaxWait:          c.opt.KafkaReader.MaxWait,
		ReadBatchTimeout: c.opt.KafkaReader.ReadBatchTimeout,
		ReadLagInterval:  c.opt.KafkaReader.ReadLagInterval,
	})
	defer func() {
		if err := reader.Close(); err != nil {
			log.Warningf("close kafka reader: %s", err)
		}
	}()

	var (
		err error
		m   kafka.Message
	)

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		m, err = reader.ReadMessage(ctx)
		if err != nil {
			continue
		}

		if c.opt.Internal.Debug {
			fmt.Printf("cumsumer#(%d) message at topic/partition/offset %v/%v/%v: %s = %s\n",
				partition, m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value),
			)
		}

		c.queue <- m
	}
}

func (c *Consumer) handle(ctx context.Context, number int) {
	log.Infof("run consumer message handler: %d", number)
	defer log.Infof("stop consumer message handler: %d", number)

	for {
		select {
		case <-ctx.Done():
			return
		case m := <-c.queue:
			now := time.Now()

			evt := &event.EventExtended{}
			if err := easyjson.Unmarshal(m.Value, evt); err != nil {
				c.metrics.IncConsumerError(err)

				log.Warningf("consumer proc: unmarshal json %s", err)
				continue
			}

			evt.FromQueueDatetime = now

			if c.opt.Internal.WithGeo && evt.IP != "" {
				nIP := net.ParseIP(evt.IP)
				if nIP != nil {
					c.enrichEventLocation(evt, nIP)
				}
			}

			if err := c.eventRepository.AddEvent(ctx, evt); err != nil {
				log.Warningf("consumer proc: add event: %s", err)
			}
		}
	}
}

func (c *Consumer) enrichEventLocation(record *event.EventExtended, nIP net.IP) {
	var (
		err  error
		city *geoip2.City
		as   *geoip2.ASN
	)

	if city, err = c.city.City(nIP); err == nil {
		record.CountryIsoCode = city.Country.IsoCode

		cityFound := false
		if name, ok := city.City.Names["en"]; ok {
			record.City = name
			cityFound = true
		}

		if !cityFound {
			if name, ok := city.City.Names["EN"]; ok {
				record.City = name
				cityFound = true
			}
		}

		if !cityFound && len(city.City.Names) > 0 {
			// peek first city name, unordered
			for k := range city.City.Names {
				record.City = city.City.Names[k]
				break
			}
		}

		if len(city.Subdivisions) > 0 {
			record.Region = city.Subdivisions[0].IsoCode
		}
	}

	if as, err = c.asn.ASN(nIP); err == nil {
		record.AS = uint32(as.AutonomousSystemNumber)
		record.ORG = as.AutonomousSystemOrganization
	}
}

func New(
	eventRepository EventRepository,
	city, asn *geoip2.Reader,
	opt *config.Config,
	metrics Metrics,
) *Consumer {
	return &Consumer{
		eventRepository: eventRepository,
		city:            city,
		asn:             asn,
		wg:              &sync.WaitGroup{},
		metrics:         metrics,
		opt:             opt,
		queue:           make(chan kafka.Message, 10000),
	}
}
