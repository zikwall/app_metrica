package eventbus

import (
	"context"
	"fmt"
	"sync"

	kafkaLib "github.com/segmentio/kafka-go"

	"github.com/zikwall/app_metrica/config"
	"github.com/zikwall/app_metrica/internal/infrastructure/kafka"
	"github.com/zikwall/app_metrica/pkg/kfk"
	"github.com/zikwall/app_metrica/pkg/log"
)

type Handler interface {
	Handle(ctx context.Context, message kafkaLib.Message) error
}

type Consumer struct {
	wg        *sync.WaitGroup
	opt       config.Internal
	readerOpt *kfk.ReaderOpt

	queue chan kafkaLib.Message

	handler Handler
}

func (c *Consumer) Run(ctx context.Context) {
	usePartition := false
	if len(c.readerOpt.Partitions) > 0 {
		usePartition = true
	}

	for _, partition := range c.readerOpt.Partitions {
		c.wg.Add(1)

		go func(p int) {
			defer c.wg.Done()

			c.proc(ctx, p, usePartition)
		}(partition)
	}

	for i := 1; i <= c.opt.ConsumerQueueHandlerSize; i++ {
		c.wg.Add(1)
		go func(number int) {
			defer c.wg.Done()

			c.handle(ctx, number)
		}(i)
	}

	c.wg.Wait()
	close(c.queue)
}

func (c *Consumer) proc(ctx context.Context, partition int, usePartition bool) {
	log.Infof("run consumer for partition: %d", partition)
	defer log.Infof("stop consumer for partition: %d", partition)

	reader := kafka.NewReader(c.readerOpt, partition, usePartition)

	defer func() {
		if err := reader.Close(); err != nil {
			log.Warningf("close kafka reader: %s", err)
		}
	}()

	var (
		err error
		m   kafkaLib.Message
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

		if c.opt.Debug {
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
			if err := c.handler.Handle(ctx, m); err != nil {
				log.Warning(err)
			}
		}
	}
}

func New(
	opt config.Internal,
	readerOpt *kfk.ReaderOpt,
	handler Handler,
) *Consumer {
	return &Consumer{
		wg:        &sync.WaitGroup{},
		opt:       opt,
		readerOpt: readerOpt,
		queue:     make(chan kafkaLib.Message, 10000),
		handler:   handler,
	}
}
