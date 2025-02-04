package appmetrica

import (
	"context"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	clickhousebuffer "github.com/zikwall/clickhouse-buffer/v4"
	"github.com/zikwall/clickhouse-buffer/v4/src/buffer/cxmem"
	"github.com/zikwall/clickhouse-buffer/v4/src/cx"
	"github.com/zikwall/clickhouse-buffer/v4/src/db/cxnative"

	"github.com/zikwall/app_metrica/config"
	"github.com/zikwall/app_metrica/internal/infrastructure/repositories/clickhouse/event"
	"github.com/zikwall/app_metrica/internal/infrastructure/repositories/clickhouse/mediavitrina"
	"github.com/zikwall/app_metrica/pkg/click"
	"github.com/zikwall/app_metrica/pkg/drop"
	"github.com/zikwall/app_metrica/pkg/geolocation"
	"github.com/zikwall/app_metrica/pkg/kfk"
	"github.com/zikwall/app_metrica/pkg/xerror"
)

type Options struct {
	Click *click.Opt

	KafkaReader             *kfk.ReaderOpt
	KafkaReaderMediaVitrina *kfk.ReaderOpt

	KafkaWriter             *kfk.WriterOpt
	KafkaWriterMediaVitrina *kfk.WriterOpt

	MaxMind  config.MaxMind
	Internal config.Internal
}

type AppMetrica struct {
	*drop.Impl

	ReaderCity    *geolocation.Wrapper
	ReaderASN     *geolocation.Wrapper
	Clickhouse    *click.Wrapper
	Writer        clickhousebuffer.Writer
	MediaWriter   clickhousebuffer.Writer
	bufferWrapper *click.BufferWrapper
	clientWrapper *click.ClientWrapper
}

func New(ctx context.Context, opt *Options) (*AppMetrica, error) {
	var err error

	metrica := &AppMetrica{
		Impl: drop.NewContext(ctx),
	}

	if !opt.MaxMind.IsEmpty() && opt.Internal.WithGeo {
		metrica.ReaderCity, err = geolocation.Reader(opt.MaxMind.CityPath)
		if err != nil {
			return nil, err
		}
		metrica.AddDropper(metrica.ReaderCity)
		metrica.ReaderASN, err = geolocation.Reader(opt.MaxMind.ASNPath)
		if err != nil {
			return nil, err
		}
		metrica.AddDropper(metrica.ReaderASN)
	} else if opt.Internal.WithGeo {
		return nil, xerror.ErrMaxMindWithGeo
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
			cx.NewView(opt.Internal.MetricTable, event.Columns()),
			cxmem.NewBuffer(client.Options().BatchSize()),
		)

		metrica.MediaWriter = client.Writer(
			clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
				"max_execution_time": opt.Internal.ChWriteTimeout.Seconds(),
			})),
			cx.NewView(opt.Internal.MetricMediaVitrinaTable, mediavitrina.Columns()),
			cxmem.NewBuffer(client.Options().BatchSize()),
		)
	}

	return metrica, nil
}
