package eventbus

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zikwall/app_metrica/internal/infrastructure/kafka"
	"github.com/zikwall/app_metrica/pkg/buffer"
	"github.com/zikwall/app_metrica/pkg/log"
	"github.com/zikwall/app_metrica/pkg/xerror"
)

func (e *EventBus) listen(ctx context.Context, number int) {
	log.Infof("run producer proc with circular buffer(64): %d", number)
	defer log.Infof("stop producer proc with circular buffer(64): %d", number)

	writer := kafka.New(e.writerOpt)

	defer func() {
		if err := writer.Close(); err != nil {
			log.Warningf("failed to drop Kafka producer: %s", err)
		}
	}()

	circular := buffer.NewCircularBuffer(e.opt.CircularBufferSize)

	// flushBuffer is a closure that is responsible for sending the filled circular buffer data to Kafka.
	flushBuffer := func() error {
		var (
			err error
			m   [][]byte
		)

		defer func() {
			e.metrics.IncQueue(len(m), false, err)
		}()

		if e.opt.Debug {
			log.Infof("buffer #%d is not empty, flush..", number)
		}

		// Add code here for sending the buffer data to Kafka
		// Return any potential error that occurs during the process
		m, err = circular.DequeueBatch(circular.Size())
		if err != nil {
			return fmt.Errorf("dequeue batch from buffer: %w", err)
		}

		err = writer.WriteBytesMessages(ctx, m)
		if err != nil {
			return fmt.Errorf("handle write to kafka: %w", err)
		}

		return nil
	}

	// handleVoidBuffer is a closure that is called when the circular buffer is empty.
	handleVoidBuffer := func() error {
		if !circular.IsEmpty() {
			var (
				err error
				m   [][]byte
			)

			defer func() {
				e.metrics.IncQueue(len(m), true, err)
			}()

			if e.opt.Debug {
				log.Infof("buffer #%d is not empty, flush by ticker..", number)
			}

			m, err = circular.DequeueBatch(circular.Size())
			if err != nil {
				return fmt.Errorf("ticker handle dequeue batch: %w", err)
			}

			err = writer.WriteBytesMessages(ctx, m)
			if err != nil {
				return fmt.Errorf("ticker handle write to kafka: %w", err)
			}
		}
		return nil
	}

	defer func() {
		if err := handleVoidBuffer(); err != nil {
			log.Warningf("producer proc defer: %s", err)
		}
	}()

	handleEvent := func(event Event) error {
		var (
			err   error
			bytes []byte
			start = time.Now()
		)

		defer func() {
			e.metrics.QueueDuration(start, 1, err)
		}()

		bytes, err = e.byteHandler(event)
		if err != nil {
			return fmt.Errorf("handle: %w", err)
		}

		err = circular.Enqueue(bytes)
		if err != nil {
			return fmt.Errorf("handle write to buffer: %w", err)
		}

		if circular.IsFull() {
			if err = flushBuffer(); err != nil {
				return fmt.Errorf("throughFilling: %w", err)
			}
		}

		return nil
	}

	handleEvents := func(event Event) error {
		var (
			err   error
			bytes [][]byte
			start = time.Now()
		)

		defer func() {
			e.metrics.QueueDuration(start, len(bytes), err)
		}()

		bytes, err = e.bytesHandler(event)
		if err != nil {
			return fmt.Errorf("handle: %w", err)
		}

		for i := range bytes {
			err = circular.Enqueue(bytes[i])
			if err != nil {
				return fmt.Errorf("handle write to buffer: %w", err)
			}

			if circular.IsFull() {
				if err = flushBuffer(); err != nil {
					return fmt.Errorf("throughFilling: %w", err)
				}
			}
		}
		return nil
	}

	flushTicker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ctx.Done():
			flushTicker.Stop()
			return
		case ev := <-e.events:
			switch ev.Evt {
			case EventTypeInline:
				if err := handleEvent(ev); err != nil {
					e.handleErrorMetric(err)
				}
			case EventTypeBatch:
				if err := handleEvents(ev); err != nil {
					e.handleErrorMetric(err)
				}
			}
		case <-flushTicker.C:
			if err := handleVoidBuffer(); err != nil {
				e.handleErrorMetric(err)
			}
		}
	}
}

func (e *EventBus) handleErrorMetric(err error) {
	log.Warningf("producer proc: %s", err)

	var ep *xerror.ErrPacket
	if errors.As(err, &ep) {
		// without data fields
		e.metrics.IncGatewayError(ep.Unwrap())
		return
	}
	e.metrics.IncGatewayError(err)
}
