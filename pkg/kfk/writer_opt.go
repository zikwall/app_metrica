package kfk

import "time"

type WriterOpt struct {
	// The list of broker addresses used to connect to the kafka cluster.
	Brokers []string `yaml:"brokers"`

	// Topic is the name of the topic that the writer will produce messages to.
	//
	// Setting this field or not is a mutually exclusive option. If you set Topic
	// here, you must not set Topic for any produced Message. Otherwise, if you	do
	// not set Topic, every Message must have Topic specified.
	Topic string

	// Limit on how many attempts will be made to deliver a message.
	//
	// The default is to try at most 10 times.
	MaxAttempts int

	// WriteBackoffMin optionally sets the smallest amount of time the writer waits before
	// it attempts to write a batch of messages
	//
	// Default: 100ms
	WriteBackoffMin time.Duration

	// WriteBackoffMax optionally sets the maximum amount of time the writer waits before
	// it attempts to write a batch of messages
	//
	// Default: 1s
	WriteBackoffMax time.Duration

	// Limit on how many messages will be buffered before being sent to a
	// partition.
	//
	// The default is to use a target batch size of 100 messages.
	BatchSize int

	// Limit the maximum size of a request in bytes before being sent to
	// a partition.
	//
	// The default is to use a kafka default value of 1048576.
	BatchBytes int64

	// Time limit on how often incomplete message batches will be flushed to
	// kafka.
	//
	// The default is to flush at least every second.
	BatchTimeout time.Duration

	// Timeout for read operations performed by the Writer.
	//
	// Defaults to 10 seconds.
	ReadTimeout time.Duration

	// Timeout for write operation performed by the Writer.
	//
	// Defaults to 10 seconds.
	WriteTimeout time.Duration

	// Setting this flag to true causes the WriteMessages method to never block.
	// It also means that errors are ignored since the caller will not receive
	// the returned value. Use this only if you don't care about guarantees of
	// whether the messages were written to kafka.
	//
	// Defaults to false.
	Async bool
}
