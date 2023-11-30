package kfk

import (
	"github.com/segmentio/kafka-go"
)

type ReaderWrapper struct {
	r *kafka.Reader
}

func (c *ReaderWrapper) Drop() error {
	return c.r.Close()
}

func (c *ReaderWrapper) DropMsg() string {
	return "close kafka reader"
}

func NewReaderWrapper(r *kafka.Reader) *ReaderWrapper {
	return &ReaderWrapper{r: r}
}
