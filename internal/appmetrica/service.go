package appmetrica

import (
	"context"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/segmentio/kafka-go"
	"github.com/zikwall/app_metrica/pkg/domain"
	clickhousebuffer "github.com/zikwall/clickhouse-buffer/v4"
	"github.com/zikwall/clickhouse-buffer/v4/src/buffer/cxmem"
	"github.com/zikwall/clickhouse-buffer/v4/src/cx"
	"github.com/zikwall/clickhouse-buffer/v4/src/db/cxnative"

	"github.com/zikwall/app_metrica/config"
	"github.com/zikwall/app_metrica/pkg/click"
	"github.com/zikwall/app_metrica/pkg/drop"
	"github.com/zikwall/app_metrica/pkg/geolocation"
	"github.com/zikwall/app_metrica/pkg/kfk"
)

type Options struct {
	Click       *click.Opt
	KafkaReader *kfk.ReaderOpt
	KafkaWriter *kfk.WriterOpt

	MaxMindDatabaseDir string

	Internal config.Internal
}

type AppMetrica struct {
	*drop.Impl

	ReaderCity    *geolocation.Wrapper
	Clickhouse    *click.Wrapper
	KafkaReader   *kfk.ReaderWrapper
	KafkaWriter   *kfk.WriterWrapper
	Writer        clickhousebuffer.Writer
	bufferWrapper *click.BufferWrapper
	clientWrapper *click.ClientWrapper
}

func New(ctx context.Context, opt *Options) (*AppMetrica, error) {
	var err error

	metrica := &AppMetrica{
		Impl: drop.NewContext(ctx),
	}

	if opt.MaxMindDatabaseDir != "" {
		metrica.ReaderCity, err = geolocation.Reader(opt.MaxMindDatabaseDir)
		if err != nil {
			return nil, err
		}
		metrica.AddDropper(metrica.ReaderCity)
	}

	if opt.Click != nil {
		metrica.Clickhouse, err = click.NewClickhouse(ctx, opt.Click.MaxExecutionTime, &clickhouse.Options{
			Addr: opt.Click.Hosts,
			Auth: clickhouse.Auth{
				Database: opt.Click.Database,
				Username: opt.Click.Username,
				Password: opt.Click.Password,
			},
			Settings: clickhouse.Settings{
				"max_execution_time": opt.Click.MaxExecutionTime,
			},
			DialTimeout: 5 * time.Second,
			Compression: &clickhouse.Compression{
				Method: clickhouse.CompressionLZ4,
			},
			MaxOpenConns:    opt.Click.MaxOpenConns,
			MaxIdleConns:    opt.Click.MaxIdleConns,
			ConnMaxLifetime: opt.Click.MaxConnMaxLifetime,
			Debug:           opt.Internal.Debug,
		})
		if err != nil {
			return nil, err
		}
		// metrica.AddDropper(metrica.Clickhouse)

		ch := cxnative.NewClickhouseWithConn(metrica.Clickhouse.Conn(), &cx.RuntimeOptions{
			WriteTimeout: opt.Internal.ChWriteTimeout,
		})
		client := clickhousebuffer.NewClientWithOptions(ctx, ch, clickhousebuffer.NewOptions(
			clickhousebuffer.WithFlushInterval(opt.Internal.BufFlushInterval),
			clickhousebuffer.WithBatchSize(opt.Internal.BufSize+1),
			clickhousebuffer.WithDebugMode(opt.Internal.Debug),
			clickhousebuffer.WithRetry(true),
		))

		metrica.bufferWrapper = click.NewBufferWrapper(ch)
		metrica.clientWrapper = click.NewClientWrapper(client)

		metrica.AddDroppers(metrica.bufferWrapper, metrica.clientWrapper)

		metrica.Writer = client.Writer(
			clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
				"max_execution_time": opt.Internal.ChWriteTimeout.Seconds(),
			})),
			cx.NewView(opt.Internal.MetricTable, domain.Columns()),
			cxmem.NewBuffer(client.Options().BatchSize()),
		)
	}

	if opt.KafkaReader != nil {
		metrica.KafkaReader = kfk.NewReaderWrapper(kafka.NewReader(kafka.ReaderConfig{
			Brokers:          opt.KafkaReader.Brokers,
			GroupID:          opt.KafkaReader.GroupID,
			GroupTopics:      opt.KafkaReader.GroupTopics,
			Topic:            opt.KafkaReader.Topic,
			Partition:        opt.KafkaReader.Partition,
			QueueCapacity:    opt.KafkaReader.QueueCapacity,
			MinBytes:         opt.KafkaReader.MinBytes,
			MaxBytes:         opt.KafkaReader.MaxBytes,
			MaxWait:          opt.KafkaReader.MaxWait,
			ReadBatchTimeout: opt.KafkaReader.ReadBatchTimeout,
			ReadLagInterval:  opt.KafkaReader.ReadLagInterval,
		}))
		metrica.AddDropper(metrica.KafkaReader)
	}

	if opt.KafkaWriter != nil {
		metrica.KafkaWriter = kfk.NewWriterWrapper(&kafka.Writer{
			Addr:            kafka.TCP(opt.KafkaWriter.Brokers...),
			Topic:           opt.KafkaWriter.Topic,
			Balancer:        &kafka.LeastBytes{},
			MaxAttempts:     opt.KafkaWriter.MaxAttempts,
			WriteBackoffMin: opt.KafkaWriter.WriteBackoffMin,
			WriteBackoffMax: opt.KafkaWriter.WriteBackoffMax,
			BatchSize:       opt.KafkaWriter.BatchSize,
			BatchBytes:      opt.KafkaWriter.BatchBytes,
			BatchTimeout:    opt.KafkaWriter.BatchTimeout,
			ReadTimeout:     opt.KafkaWriter.ReadTimeout,
			WriteTimeout:    opt.KafkaWriter.WriteTimeout,
		})
		metrica.AddDropper(metrica.KafkaWriter)
	}

	return metrica, nil
}
