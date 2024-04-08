package eventbus

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mailru/easyjson"
	"github.com/segmentio/kafka-go"

	"github.com/zikwall/app_metrica/pkg/buffer"
	"github.com/zikwall/app_metrica/pkg/domain"
	"github.com/zikwall/app_metrica/pkg/log"
	"github.com/zikwall/app_metrica/pkg/x"
)

func (e *EventBus) newKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:            kafka.TCP(e.opt.KafkaWriter.Brokers...),
		Topic:           e.opt.KafkaWriter.Topic,
		Balancer:        &kafka.LeastBytes{},
		MaxAttempts:     e.opt.KafkaWriter.MaxAttempts,
		WriteBackoffMin: e.opt.KafkaWriter.WriteBackoffMin,
		WriteBackoffMax: e.opt.KafkaWriter.WriteBackoffMax,
		BatchSize:       e.opt.KafkaWriter.BatchSize,
		BatchBytes:      e.opt.KafkaWriter.BatchBytes,
		BatchTimeout:    e.opt.KafkaWriter.BatchTimeout,
		ReadTimeout:     e.opt.KafkaWriter.ReadTimeout,
		WriteTimeout:    e.opt.KafkaWriter.WriteTimeout,
	}
}

func (e *EventBus) DebugMessage(ev Event) error {
	var (
		err   error
		event = &domain.Event{}
	)
	if err = json.Unmarshal(ev.data, event); err != nil {
		return err
	}
	return nil
}

func (e *EventBus) listen(ctx context.Context, number int) {
	log.Infof("run producer proc with circular buffer(64): %d", number)
	defer log.Infof("stop producer proc with circular buffer(64): %d", number)

	writer := e.newKafkaWriter()

	defer func() {
		if err := writer.Close(); err != nil {
			log.Warningf("failed to drop Kafka producer: %s", err)
		}
	}()

	circular := buffer.NewCircularBuffer(64)

	// flushBuffer is a closure that is responsible for sending the filled circular buffer data to Kafka.
	flushBuffer := func() error {
		if e.opt.Internal.Debug {
			log.Infof("buffer #%d is not empty, flush..", number)
		}

		// Add code here for sending the buffer data to Kafka
		// Return any potential error that occurs during the process
		m, err := circular.DequeueBatch(circular.Size())
		if err != nil {
			return fmt.Errorf("dequeue batch from buffer: %w", err)
		}

		err = writer.WriteMessages(ctx, assembleMessages(m)...)
		if err != nil {
			return fmt.Errorf("handle write to kafka: %w", err)
		}
		return nil
	}

	// handleVoidBuffer is a closure that is called when the circular buffer is empty.
	handleVoidBuffer := func() error {
		if !circular.IsEmpty() {
			if e.opt.Internal.Debug {
				log.Infof("buffer #%d is not empty, flush by ticker..", number)
			}

			m, err := circular.DequeueBatch(circular.Size())
			if err != nil {
				return fmt.Errorf("ticker handle dequeue batch: %w", err)
			}

			err = writer.WriteMessages(ctx, assembleMessages(m)...)

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
		bytes, err := handleBytes(event)
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
		bytes, err := handleMultipleBytes(event)
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
		case event := <-e.events:
			switch event.evt {
			case EventTypeInline:
				if err := handleEvent(event); err != nil {
					log.Warningf("producer proc: %s", err)
				}
			case EventTypeBatch:
				if err := handleEvents(event); err != nil {
					log.Warningf("producer proc: %s", err)
				}
			}
		case <-flushTicker.C:
			if err := handleVoidBuffer(); err != nil {
				log.Warningf("producer proc: %s", err)
			}
		}
	}
}

func handleBytes(e Event) ([]byte, error) {
	var (
		err   error
		event = &domain.Event{}
	)
	if err = easyjson.Unmarshal(e.data, event); err != nil {
		return nil, fmt.Errorf("unmarshal json object: %w, %s", err, string(e.data))
	}

	exEvent := domain.ExtendEvent(event, time.Now(), e.t)
	exEvent.IP = e.ip

	var bytes []byte
	if bytes, err = easyjson.Marshal(exEvent); err != nil {
		return nil, err
	}
	return bytes, nil
}

func handleMultipleBytes(e Event) ([][]byte, error) {
	var (
		err    error
		events domain.Events
	)
	if err = easyjson.Unmarshal(e.data, &events); err != nil {
		return nil, err
	}

	bytes := make([][]byte, len(events))
	for i := range events {
		exEvent := domain.ExtendEvent(events[i], time.Now(), e.t)
		exEvent.IP = e.ip

		var extBytes []byte
		if extBytes, err = easyjson.Marshal(exEvent); err != nil {
			return nil, err
		}

		bytes[i] = extBytes
	}

	return bytes, nil
}

func assembleMessages(m [][]byte) []kafka.Message {
	return x.Map[[]byte, kafka.Message](m, func(i []byte, _ int) kafka.Message {
		return kafka.Message{Value: i}
	})
}
