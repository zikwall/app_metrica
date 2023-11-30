package click

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Wrapper struct {
	conn          driver.Conn
	executionTime time.Duration
}

func (c *Wrapper) Conn() driver.Conn {
	return c.conn
}

func (c *Wrapper) Drop() error {
	return c.conn.Close()
}

func (c *Wrapper) DropMsg() string {
	return "close clickhouse connection"
}

func NewClickhouseWrapper(conn driver.Conn, timeout int) *Wrapper {
	return &Wrapper{conn: conn, executionTime: time.Duration(timeout) * time.Second}
}

func NewClickhouse(ctx context.Context, timeout int, options *clickhouse.Options) (*Wrapper, error) {
	conn, err := clickhouse.Open(options)
	if err != nil {
		return nil, err
	}
	ctx = clickhouse.Context(ctx, clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
	}))
	if err = conn.Ping(ctx); err != nil {
		var e *clickhouse.Exception
		if errors.As(err, &e) {
			fmt.Printf("catch exception [%d] %s \n%s\n", e.Code, e.Message, e.StackTrace)
		}
		return nil, err
	}
	return NewClickhouseWrapper(conn, timeout), nil
}
