package config

import (
	"bytes"
	"os"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/zikwall/app_metrica/pkg/click"
	"github.com/zikwall/app_metrica/pkg/kfk"
)

type Server struct {
	Bugsnag            Bugsnag `yaml:"bugsnag"`
	MaxMindDatabaseDir string  `yaml:"max_mind_database_dir"`

	Clickhouse  *click.Opt     `yaml:"clickhouse"`
	KafkaReader *kfk.ReaderOpt `yaml:"kafka_reader"`
	KafkaWriter *kfk.WriterOpt `yaml:"kafka_writer"`

	Prefork  bool     `yaml:"prefork"`
	Internal Internal `yaml:"internal"`
}

type Internal struct {
	ProducerPerInstanceSize  int           `yaml:"handler_proc_size"`
	ConsumerQueueHandlerSize int           `yaml:"consumer_queue_handler_size"`
	ConsumerPerInstanceSize  int           `yaml:"consumer_per_instance_size"`
	BufSize                  uint          `yaml:"buf_size"`
	BufFlushInterval         uint          `yaml:"buf_flush_interval"`
	ChWriteTimeout           time.Duration `yaml:"ch_write_timeout"`
	MetricTable              string        `yaml:"metric_table"`
	Debug                    bool          `yaml:"debug"`
	WithGeo                  bool          `yaml:"with_geo"`
}

type Bugsnag struct {
	APIKey       string `yaml:"api_key"`
	ReleaseStage string `yaml:"release_stage"`
}

func (b *Bugsnag) Maybe() bool {
	return b.APIKey != "" && b.ReleaseStage != ""
}

type Config struct {
	Server `yaml:"server"`
}

func New(filepath string) (*Config, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	d := yaml.NewDecoder(bytes.NewReader(content))
	if err = d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}
