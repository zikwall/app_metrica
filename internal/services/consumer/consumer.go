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
	clickhousebuffer "github.com/zikwall/clickhouse-buffer/v4"

	"github.com/zikwall/app_metrica/pkg/domain"
	"github.com/zikwall/app_metrica/pkg/log"
)

type Consumer struct {
	reader *kafka.Reader
	city   *geoip2.Reader
	asn    *geoip2.Reader
	writer clickhousebuffer.Writer

	wg *sync.WaitGroup

	poolSize          int
	debug             bool
	enrichWithGeoData bool
}

func (c *Consumer) Run(ctx context.Context) {
	for i := 1; i <= c.poolSize; i++ {
		c.wg.Add(1)
		go func(number int) {
			defer c.wg.Done()

			c.proc(ctx, number)
		}(i)
	}
	c.wg.Wait()
}

func (c *Consumer) proc(ctx context.Context, number int) {
	log.Infof("run consumer proc: %d", number)
	defer log.Infof("stop consumer proc: %d", number)

	var (
		err error
		m   kafka.Message
	)

	for {
		m, err = c.reader.ReadMessage(ctx)
		if err != nil {
			continue
		}

		now := time.Now()
		if c.debug {
			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n",
				m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value),
			)
		}

		event := &domain.EventExtended{}
		if err = easyjson.Unmarshal(m.Value, event); err != nil {
			log.Warningf("consumer proc: unmarshal json %s", err)
			continue
		}

		record := domain.RecordFromEvent(event)
		record.FromQueueDatetime = now
		record.FromQueueTimestamp = now.Unix()

		if c.enrichWithGeoData && event.IP != "" {
			nIP := net.ParseIP(event.IP)
			if nIP != nil {
				c.enrichRecordLocation(record, nIP)
			}
		}

		c.writer.TryWriteRow(record)
	}
}

func (c *Consumer) enrichRecordLocation(record *domain.Record, nIP net.IP) {
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
	reader *kafka.Reader,
	writer clickhousebuffer.Writer,
	city, asn *geoip2.Reader,
	poolSize int,
	withGeo, debug bool,
) *Consumer {
	return &Consumer{
		reader:            reader,
		writer:            writer,
		city:              city,
		asn:               asn,
		wg:                &sync.WaitGroup{},
		poolSize:          poolSize,
		enrichWithGeoData: withGeo,
		debug:             debug,
	}
}
