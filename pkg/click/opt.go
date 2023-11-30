package click

import "time"

type Opt struct {
	Hosts              []string      `yaml:"hosts"`
	Username           string        `yaml:"username"`
	Password           string        `yaml:"password"`
	Database           string        `yaml:"database"`
	MaxExecutionTime   int           `yaml:"max_execution_time"`
	MaxOpenConns       int           `yaml:"max_open_conns"`
	MaxIdleConns       int           `yaml:"max_idle_conns"`
	MaxConnMaxLifetime time.Duration `yaml:"max_conn_max_lifetime"`
}
