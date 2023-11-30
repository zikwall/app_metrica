package kfk

import (
	"github.com/segmentio/kafka-go"
)

type WriterWrapper struct {
	w *kafka.Writer
}

func (c *WriterWrapper) Writer() *kafka.Writer {
	return c.w
}

func (c *WriterWrapper) Drop() error {
	return c.w.Close()
}

func (c *WriterWrapper) DropMsg() string {
	return "close kafka writer"
}

func NewWriterWrapper(w *kafka.Writer) *WriterWrapper {
	return &WriterWrapper{w: w}
}
