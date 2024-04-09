package eventbus

import (
	"context"
	"sync"
	"time"

	"github.com/zikwall/app_metrica/config"
	"github.com/zikwall/app_metrica/internal/metrics"
)

type EventType uint8

const (
	EventTypeInline EventType = iota
	EventTypeBatch
)

type Event struct {
	data []byte
	ip   string
	t    time.Time

	evt EventType
}

func NewEvent(data []byte, ip string, received time.Time, evt EventType) Event {
	return Event{
		data: data,
		ip:   ip,
		t:    received,
		evt:  evt,
	}
}

type EventBus struct {
	wg      *sync.WaitGroup
	opt     *config.Config
	metrics *metrics.Metrics

	events chan Event

	Done chan struct{}
}

func NewEventBus(opt *config.Config, metrics *metrics.Metrics) *EventBus {
	bus := &EventBus{
		wg:      &sync.WaitGroup{},
		opt:     opt,
		metrics: metrics,
		events:  make(chan Event, opt.Internal.ProducerPerInstanceSize+10000),
		Done:    make(chan struct{}, 1),
	}
	return bus
}

func (e *EventBus) SendEvent(event Event) {
	e.events <- event
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
