package main

import (
	"context"
	"net"
	"os"

	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
	"github.com/zikwall/app_metrica/internal/services/gateway"

	"github.com/zikwall/app_metrica/config"
	"github.com/zikwall/app_metrica/internal/appmetrica"
	"github.com/zikwall/app_metrica/pkg/fiberext"
	"github.com/zikwall/app_metrica/pkg/log"
	"github.com/zikwall/app_metrica/pkg/signal"
)

func main() {
	application := cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config-file",
				Required: true,
				Usage:    "YAML config filepath",
				EnvVars:  []string{"APPMETRICA_CONFIG_FILE"},
				FilePath: "/srv/app_metrica_secrets/config_file",
			},
			&cli.StringFlag{
				Name:     "bind-address",
				Usage:    "IP и порт сервера, например: 0.0.0.0:3000",
				Required: false,
				Value:    "0.0.0.0:3000",
				EnvVars:  []string{"APPMETRICA_HTTP_BIND_ADDRESS"},
			},
			&cli.StringFlag{
				Name:     "bind-socket",
				Usage:    "Путь к Unix сокет файлу",
				Required: false,
				Value:    "/tmp/gateway.sock",
				EnvVars:  []string{"APPMETRICA_BIND_SOCKET"},
			},
			&cli.IntFlag{
				Name:     "listener",
				Usage:    "Unix socket or TCP",
				Required: false,
				Value:    1,
				EnvVars:  []string{"APPMETRICA_HTTP_LISTENER"},
			},
		},
		Action: Main,
		After: func(c *cli.Context) error {
			log.Info("stopped")
			return nil
		},
	}

	if err := application.Run(os.Args); err != nil {
		log.Error(err)
	}
}

// nolint:funlen // it's ok
func Main(ctx *cli.Context) error {
	appContext, cancel := context.WithCancel(ctx.Context)
	defer func() {
		cancel()
		log.Info("app context is canceled, AppMetrica is down!")
	}()

	cfg, err := config.New(ctx.String("config-file"))
	if err != nil {
		return err
	}

	metrica, err := appmetrica.New(appContext, &appmetrica.Options{
		Click:              cfg.Clickhouse,
		KafkaWriter:        cfg.KafkaWriter,
		MaxMindDatabaseDir: cfg.MaxMindDatabaseDir,
	})
	if err != nil {
		return err
	}

	if cfg.Bugsnag.Maybe() {
		bugsnag.Configure(bugsnag.Configuration{
			// Your Bugsnag project API key, required unless set as environment
			// variable $BUGSNAG_API_KEY
			APIKey: cfg.Bugsnag.APIKey,
			// The development stage of your application build, like "alpha" or
			// "production"
			ReleaseStage: cfg.Bugsnag.ReleaseStage,
			// The import paths for the Go packages containing your source files
			ProjectPackages: []string{"main", "github.com/zikwall/app_metrica"},
			// more configuration options
		})

		log.EnableBugsnagReporter()
	}

	defer func() {
		metrica.Shutdown(func(err error) {
			log.Warning(err)
		})
		metrica.Stacktrace()
	}()

	await, stop := signal.Notifier(func() {
		log.Info("received a system signal to shutdown AppMetrica, start shutdown process..")
	})

	app := fiber.New(fiber.Config{
		Prefork:      cfg.Prefork,
		ErrorHandler: fiberext.ErrorHandler,
	})

	handler := gateway.NewHandler(metrica.KafkaWriter.Writer(), cfg.Internal.HandlerProcSize)
	handler.MountRoutes(app)

	go func() {
		handler.Run(metrica.Context())
	}()

	go func() {
		var ln net.Listener
		if ln, err = signal.Listener(
			metrica.Context(),
			ctx.Int("listener"),
			ctx.String("bind-socket"),
			ctx.String("bind-address"),
		); err != nil {
			stop(err)
			return
		}
		if err = app.Listener(ln); err != nil {
			stop(err)
		}
	}()

	log.Info("AppMetrica: statistic and analytic system")
	return await()
}
