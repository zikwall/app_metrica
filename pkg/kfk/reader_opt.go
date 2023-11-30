package kfk

import "time"

type ReaderOpt struct {
	// The list of broker addresses used to connect to the kafka cluster.
	Brokers []string `yaml:"brokers"`

	// GroupID holds the optional consumer group id.  If GroupID is specified, then
	// Partition should NOT be specified e.g. 0
	GroupID string `yaml:"group_id"`

	// GroupTopics allows specifying multiple topics, but can only be used in
	// combination with GroupID, as it is a consumer-group feature. As such, if
	// GroupID is set, then either Topic or GroupTopics must be defined.
	GroupTopics []string `yaml:"group_topics"`

	// The topic to read messages from.
	Topic string `yaml:"topic"`

	// Partition to read messages from.  Either Partition or GroupID may
	// be assigned, but not both
	Partition int `yaml:"partition"`

	// The capacity of the internal message queue, defaults to 100 if none is
	// set.
	QueueCapacity int

	// MinBytes indicates to the broker the minimum batch size that the consumer
	// will accept. Setting a high minimum when consuming from a low-volume topic
	// may result in delayed delivery when the broker does not have enough data to
	// satisfy the defined minimum.
	//
	// Default: 1
	MinBytes int

	// MaxBytes indicates to the broker the maximum batch size that the consumer
	// will accept. The broker will truncate a message to satisfy this maximum, so
	// choose a value that is high enough for your largest message size.
	//
	// Default: 1MB
	MaxBytes int

	// Maximum amount of time to wait for new data to come when fetching batches
	// of messages from kafka.
	//
	// Default: 10s
	MaxWait time.Duration

	// ReadBatchTimeout amount of time to wait to fetch message from kafka messages batch.
	//
	// Default: 10s
	ReadBatchTimeout time.Duration

	// ReadLagInterval sets the frequency at which the reader lag is updated.
	// Setting this field to a negative value disables lag reporting.
	ReadLagInterval time.Duration
}
