package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"

	"github.com/zikwall/app_metrica/pkg/kfk"
)

type Reader struct {
	reader *kafka.Reader
}

func (r *Reader) ReadMessage(ctx context.Context) (kafka.Message, error) {
	return r.reader.ReadMessage(ctx)
}

func (r *Reader) Close() error {
	return r.reader.Close()
}

func NewReader(opt *kfk.ReaderOpt, partition int, usePartition bool) *Reader {
	groupID := opt.GroupID

	if usePartition {
		groupID = ""
	} else {
		partition = 0
	}

	return &Reader{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:          opt.Brokers,
			GroupID:          groupID,
			GroupTopics:      opt.GroupTopics,
			Topic:            opt.Topic,
			Partition:        partition,
			QueueCapacity:    opt.QueueCapacity,
			MinBytes:         opt.MinBytes,
			MaxBytes:         opt.MaxBytes,
			MaxWait:          opt.MaxWait,
			ReadBatchTimeout: opt.ReadBatchTimeout,
			ReadLagInterval:  opt.ReadLagInterval,
		}),
	}
}
