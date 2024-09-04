package main

import (
	"context"
	"net"
	"os"

	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/urfave/cli/v2"

	_ "github.com/zikwall/app_metrica/docs"

	"github.com/zikwall/app_metrica/config"
	"github.com/zikwall/app_metrica/internal/appmetrica"
	"github.com/zikwall/app_metrica/internal/metrics"
	"github.com/zikwall/app_metrica/internal/services/gateway"
	"github.com/zikwall/app_metrica/internal/services/gateway/eventbus"
	"github.com/zikwall/app_metrica/internal/services/gateway/handlers/mediavitrina"
	"github.com/zikwall/app_metrica/internal/services/gateway/handlers/own"
	"github.com/zikwall/app_metrica/pkg/fiberext"
	"github.com/zikwall/app_metrica/pkg/log"
	"github.com/zikwall/app_metrica/pkg/prometheus"
	"github.com/zikwall/app_metrica/pkg/signal"
)

// @title           OwnMetrics aka AppMetrica gateway service
// @version         1.0
// @description     Under construct
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  a.kapitonov@limehd.tv

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      lm.limehd.tv
// @BasePath  /
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
		Click:       cfg.Clickhouse,
		KafkaWriter: cfg.KafkaWriter,
		MaxMind:     cfg.MaxMind,
		Internal:    cfg.Internal,
	})
	if err != nil {
		return err
	}

	if !cfg.Bugsnag.IsEmpty() {
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

	metro := metrics.New()
	metroVitrina := metrics.NewVitrina()

	ownBus := eventbus.NewEventBus(
		cfg.Internal,
		metro,
		cfg.KafkaWriter,
		"own_metrics",
		own.HandleBytes,
		own.HandleMultipleBytes,
		own.DebugMessage,
	)

	mediaBus := eventbus.NewEventBus(
		cfg.Internal,
		metro,
		cfg.KafkaWriterMediaVitrina,
		"media_vitrina",
		mediavitrina.HandleBytes,
		mediavitrina.HandleMultipleBytes,
		mediavitrina.DebugMessage,
	)

	defer func() {
		metrica.Shutdown(func(err error) {
			log.Warning(err)
		})

		<-ownBus.Done
		<-mediaBus.Done

		metrica.Stacktrace()
	}()

	await, stop := signal.Notifier(func() {
		log.Info("received a system signal to shutdown AppMetrica, start shutdown process..")
	})

	// run event bus
	go func() {
		ownBus.Run(metrica.Context())
	}()

	// run event bus
	go func() {
		mediaBus.Run(metrica.Context())
	}()

	// run HTTP server for receive events
	go func() {
		app := fiber.New(fiber.Config{
			Prefork:      cfg.Prefork,
			ErrorHandler: fiberext.ErrorHandler,
		})
		app.Use(cors.New())

		app.Get("/swagger/*", swagger.HandlerDefault)
		app.Get("/metrics", prometheus.FastHTTPAdapter())

		handler := gateway.NewHandler(ownBus, mediaBus, metro, metroVitrina)
		handler.MountRoutes(app)

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
