package config

import (
	"bytes"
	"os"

	"github.com/zikwall/app_metrica/pkg/kfk"
	"gopkg.in/yaml.v3"

	"github.com/zikwall/app_metrica/pkg/click"
)

type Server struct {
	Bugsnag            Bugsnag `yaml:"bugsnag"`
	MaxMindDatabaseDir string  `yaml:"max_mind_database_dir"`

	Clickhouse  *click.Opt     `yaml:"clickhouse"`
	KafkaReader *kfk.ReaderOpt `yaml:"kafka_reader"`
	KafkaWriter *kfk.WriterOpt `yaml:"kafka_writer"`

	Prefork bool `yaml:"prefork"`
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
