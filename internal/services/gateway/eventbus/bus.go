package eventbus

import (
	"context"
	"sync"
	"time"

	"github.com/zikwall/app_metrica/config"
	"github.com/zikwall/app_metrica/internal/metrics"
)

type Metrics interface {
	IncRequests(isBatch bool)
	IncQueue(size int, byTicker bool, err error)
	RequestsDuration(start time.Time, isBatch bool)
	QueueDuration(start time.Time, size int, err error)
	IncConsumerError(err error)
	IncGatewayError(err error)
}

type ByteHandler func(event Event) ([]byte, error)
type BytesHandler func(event Event) ([][]byte, error)
type DebugHandler func(ev Event) error

type EventBus struct {
	wg      *sync.WaitGroup
	opt     *config.Config
	metrics Metrics

	events chan Event

	name            string
	byteHandler     ByteHandler
	bytesHandler    BytesHandler
	messageDebugger DebugHandler

	Done chan struct{}
}

func NewEventBus(
	opt *config.Config,
	metrics *metrics.Metrics,
	name string,
	byteHandler ByteHandler,
	bytesHandler BytesHandler,
	messageDebugger DebugHandler,
) *EventBus {
	bus := &EventBus{
		wg:              &sync.WaitGroup{},
		opt:             opt,
		metrics:         metrics,
		events:          make(chan Event, opt.Internal.ProducerPerInstanceSize+10000),
		name:            name,
		byteHandler:     byteHandler,
		bytesHandler:    bytesHandler,
		messageDebugger: messageDebugger,
		Done:            make(chan struct{}, 1),
	}
	return bus
}

func (e *EventBus) SendEvent(event Event) {
	e.events <- event
}

func (e *EventBus) DebugMessage(event Event) error {
	return e.messageDebugger(event)
}

func (e *EventBus) Run(ctx context.Context) {
	for i := 1; i <= e.opt.Internal.ProducerPerInstanceSize; i++ {
		e.wg.Add(1)
		go func(number int) {
			defer e.wg.Done()

			e.listen(ctx, number)
		}(i)
	}

	e.wg.Wait()

	e.Done <- struct{}{}
}
