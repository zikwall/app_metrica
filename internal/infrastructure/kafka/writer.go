package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"

	"github.com/zikwall/app_metrica/pkg/kfk"
	"github.com/zikwall/app_metrica/pkg/x"
)

type Writer struct {
	writer *kafka.Writer
}

func New(opt *kfk.WriterOpt) *Writer {
	return &Writer{
		writer: &kafka.Writer{
			Addr:            kafka.TCP(opt.Brokers...),
			Topic:           opt.Topic,
			Balancer:        &kafka.LeastBytes{},
			MaxAttempts:     opt.MaxAttempts,
			WriteBackoffMin: opt.WriteBackoffMin,
			WriteBackoffMax: opt.WriteBackoffMax,
			BatchSize:       opt.BatchSize,
			BatchBytes:      opt.BatchBytes,
			BatchTimeout:    opt.BatchTimeout,
			ReadTimeout:     opt.ReadTimeout,
			WriteTimeout:    opt.WriteTimeout,
		},
	}
}

func (w *Writer) WriteBytesMessages(ctx context.Context, messages [][]byte) error {
	err := w.writer.WriteMessages(ctx, assembleMessages(messages)...)
	if err != nil {
		return fmt.Errorf("handle write to kafka: %w", err)
	}

	return nil
}

func (w *Writer) Close() error {
	return w.writer.Close()
}

func assembleMessages(m [][]byte) []kafka.Message {
	return x.Map[[]byte, kafka.Message](m, func(i []byte, _ int) kafka.Message {
		return kafka.Message{Value: i}
	})
}
