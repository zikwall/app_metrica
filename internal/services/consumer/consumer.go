package consumer

import (
	"context"
	"fmt"
	"sync"

	"github.com/segmentio/kafka-go"
	"github.com/zikwall/app_metrica/pkg/log"
)

type Consumer struct {
	reader *kafka.Reader

	wg *sync.WaitGroup

	poolSize int
	debug    bool
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
			break
		}
		if c.debug {
			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n",
				m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value),
			)
		}
	}
}

func New(reader *kafka.Reader, poolSize int, debug bool) *Consumer {
	return &Consumer{
		reader:   reader,
		wg:       &sync.WaitGroup{},
		poolSize: poolSize,
		debug:    debug,
	}
}
