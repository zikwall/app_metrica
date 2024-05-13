package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/urfave/cli/v2"

	"github.com/zikwall/app_metrica/pkg/log"
)

type Payload struct {
	ApplicationID         int    `json:"application_id"`
	IosIfa                string `json:"ios_ifa"`
	IosIfv                string `json:"ios_ifv"`
	AndroidID             string `json:"android_id"`
	GoogleAid             string `json:"google_aid"`
	ProfileID             string `json:"profile_id"`
	OsName                string `json:"os_name"`
	OsVersion             string `json:"os_version"`
	DeviceManufacturer    string `json:"device_manufacturer"`
	DeviceModel           string `json:"device_model"`
	DeviceType            string `json:"device_type"`
	DeviceLocale          string `json:"device_locale"`
	AppVersionName        string `json:"app_version_name"`
	AppPackageName        string `json:"app_package_name"`
	EventName             string `json:"event_name"`
	EventJSON             string `json:"event_json"`
	EventDatetime         string `json:"event_datetime"`
	EventTimestamp        int    `json:"event_timestamp"`
	EventReceiveDatetime  string `json:"event_receive_datetime"`
	EventReceiveTimestamp string `json:"event_receive_timestamp"`
	ConnectionType        string `json:"connection_type"`
	OperatorName          string `json:"operator_name"`
	Mcc                   string `json:"mcc"`
	Mnc                   string `json:"mnc"`
	CountryIsoCode        string `json:"country_iso_code"`
	City                  string `json:"city"`
	AppmetricaDeviceID    string `json:"appmetrica_device_id"`
	InstallationID        string `json:"installation_id"`
	SessionID             string `json:"session_id"`
	XLHDAgent             string `json:"xlhd_agent"`
	UserAgent             string `json:"user_agent"`
	HardwareOrGUI         string `json:"hardware_or_gui"`
}

func (p *Payload) Bytes() ([]byte, error) {
	b, err := json.Marshal(*p)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func randomPayload(rnd *randomMachine) *Payload {
	return &Payload{
		ApplicationID:         rnd.rndAppID(),
		IosIfa:                rnd.rndIosIfa(),
		IosIfv:                rnd.rndIosIfv(),
		AndroidID:             rnd.rndAndroidID(),
		GoogleAid:             rnd.rndGoogleAid(),
		ProfileID:             rnd.rndProfileID(),
		OsName:                rnd.rndOsName(),
		OsVersion:             rnd.rndOsVersion(),
		DeviceManufacturer:    rnd.rndDeviceManufacturer(),
		DeviceModel:           rnd.rndDeviceModel(),
		DeviceType:            rnd.rndDeviceType(),
		DeviceLocale:          rnd.rndDeviceLocale(),
		AppVersionName:        rnd.rndAppVersion(),
		AppPackageName:        rnd.rndPackage(),
		EventName:             rnd.rndEventName(),
		EventJSON:             rnd.rndEventJSON(),
		EventDatetime:         rnd.rndEventDatetime(),
		EventTimestamp:        rnd.rndEventTimestamp(),
		EventReceiveDatetime:  rnd.rndEventReceiveDatetime(),
		EventReceiveTimestamp: rnd.rndEventReceiveTimestamp(),
		ConnectionType:        rnd.rndConnectionType(),
		OperatorName:          rnd.rndOperatorName(),
		Mcc:                   rnd.rndMcc(),
		Mnc:                   rnd.rndMnc(),
		CountryIsoCode:        rnd.rndCountry(),
		City:                  rnd.rndCity(),
		AppmetricaDeviceID:    rnd.rndAppmetricaDeviceID(),
	}
}

func main() {
	application := cli.App{
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "addresses",
				Usage:    "URL address",
				Required: false,
				Value:    cli.NewStringSlice("http://localhost:3000/internal/api/v1/event"),
			},
			&cli.IntFlag{
				Name:     "count",
				Usage:    "is the total number of requests to make.",
				Required: false,
				Value:    1,
			},
			&cli.IntFlag{
				Name:     "parallelism",
				Usage:    "is concurrency level, the number of concurrent workers to run.",
				Required: false,
				Value:    1,
			},
			&cli.Float64Flag{
				Name:     "qps",
				Usage:    "Qps is the rate limit in queries per second.",
				Required: false,
				Value:    100,
			},
			&cli.DurationFlag{
				Name:     "duration",
				Usage:    "Duration of application to send requests. When duration is reached,\napplication stops and exits. If duration is specified, is ignored.\n",
				Required: false,
				Value:    time.Minute,
			},
			&cli.IntFlag{
				Name:     "timeout",
				Usage:    "Timeout for each request in seconds. Default is 20, use 0 for infinite.",
				Required: false,
				Value:    20,
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
	appContext, cancel := context.WithTimeout(ctx.Context, ctx.Duration("duration"))
	defer func() {
		cancel()
		log.Info("app context is canceled, AppMetrica Mocker is down!")
	}()

	urls := make([]*url.URL, 0, len(ctx.StringSlice("addresses")))
	for _, address := range ctx.StringSlice("addresses") {
		u, err := url.Parse(address)
		if err != nil {
			return err
		}

		urls = append(urls, u)
	}

	w := &worker{
		URLs:    urls,
		N:       ctx.Int("count"),
		C:       ctx.Int("parallelism"),
		Timeout: ctx.Int("timeout"),
		QPS:     ctx.Float64("qps"),
		start:   time.Now(),
		randomMachine: &randomMachine{
			countries:          generateStrings(1000, gofakeit.Country),
			cities:             generateStrings(1000, gofakeit.City),
			appsIDs:            generateIntegers(1000, 1, 1000),
			installationIDs:    generateStrings(1000, gofakeit.UUID),
			sessionIDs:         generateStrings(1000, gofakeit.UUID),
			appVersions:        generateStrings(1000, gofakeit.AppVersion),
			deviceModels:       generateStrings(1000, gofakeit.Letter),
			deviceTypes:        generateStrings(1000, gofakeit.Letter),
			deviceLocales:      generateStrings(1000, gofakeit.Language),
			deviceManufactures: generateStrings(1000, gofakeit.Company),
			osNames:            generateStrings(1000, gofakeit.Letter),
			osVersions:         generateStrings(1000, gofakeit.Letter),
			eventNames:         generateStrings(1000, gofakeit.Letter, 10),
			packages:           generateStrings(1000, gofakeit.Letter, 10),
			appMetricaIDs:      generateStrings(1000, gofakeit.UUID),
		},
	}

	w.work(appContext)
	return nil
}

type worker struct {
	URLs []*url.URL

	// N is the total number of requests to make.
	N int

	// C is the concurrency level, the number of concurrent workers to run.
	C int

	// Timeout in seconds.
	Timeout int

	Duration time.Duration

	// Qps is the rate limit in queries per second.
	QPS float64

	start time.Time

	randomMachine *randomMachine

	completedRequestsCount int64
	completedWithErrors    int64
	next                   uint32
	finished               uint64

	ctx    context.Context
	cancel context.CancelFunc
}

func (w *worker) Next() *url.URL {
	return w.URLs[(int(atomic.AddUint32(&w.next, 1))-1)%len(w.URLs)]
}

func (w *worker) runWorker(ctx context.Context, client *http.Client, n int) {
	log.Infof("run worker with capacity: %d", n)

	var throttle <-chan time.Time
	if w.QPS > 0 {
		throttle = time.Tick(time.Duration(1e6/(w.QPS)) * time.Microsecond)
	}

	for i := 0; i < n; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			if w.QPS > 0 {
				<-throttle
			}

			if err := w.httpRequestAsync(ctx, client, &Payload{
				ApplicationID:         w.randomMachine.rndAppID(),
				IosIfa:                w.randomMachine.rndIosIfa(),
				IosIfv:                w.randomMachine.rndIosIfv(),
				AndroidID:             w.randomMachine.rndAndroidID(),
				GoogleAid:             w.randomMachine.rndGoogleAid(),
				ProfileID:             w.randomMachine.rndProfileID(),
				OsName:                w.randomMachine.rndOsName(),
				OsVersion:             w.randomMachine.rndOsVersion(),
				DeviceManufacturer:    w.randomMachine.rndDeviceManufacturer(),
				DeviceModel:           w.randomMachine.rndDeviceModel(),
				DeviceType:            w.randomMachine.rndDeviceType(),
				DeviceLocale:          w.randomMachine.rndDeviceLocale(),
				AppVersionName:        w.randomMachine.rndAppVersion(),
				AppPackageName:        w.randomMachine.rndAppPackageName(),
				EventName:             w.randomMachine.rndEventName(),
				EventJSON:             w.randomMachine.rndEventJSON(),
				EventDatetime:         w.randomMachine.rndEventDatetime(),
				EventTimestamp:        w.randomMachine.rndEventTimestamp(),
				EventReceiveDatetime:  w.randomMachine.rndEventReceiveDatetime(),
				EventReceiveTimestamp: w.randomMachine.rndEventReceiveTimestamp(),
				ConnectionType:        "WiFi",
				OperatorName:          "Some Operator",
				Mcc:                   "123",
				Mnc:                   "456",
				CountryIsoCode:        w.randomMachine.rndCountry(),
				City:                  w.randomMachine.rndCity(),
				AppmetricaDeviceID:    w.randomMachine.rndAppmetricaDeviceID(),
				InstallationID:        w.randomMachine.rndInstallationID(),
				SessionID:             w.randomMachine.rndSessionID(),
				XLHDAgent:             "lhd header",
				UserAgent:             "user agent",
				HardwareOrGUI:         "idk",
			}); err != nil {
				log.Warningf("http request: %s", err)
				atomic.AddInt64(&w.completedWithErrors, 1)
			}

			atomic.AddInt64(&w.completedRequestsCount, 1)
		}
	}

	if atomic.AddUint64(&w.finished, 1) == uint64(w.C) {
		w.cancel()
	}
}

func (w *worker) work(ctx context.Context) {
	const maxIdleConn = 500

	w.ctx, w.cancel = context.WithCancel(ctx)

	var wg sync.WaitGroup
	wg.Add(w.C)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         w.URLs[0].Host,
		},
		MaxIdleConnsPerHost: min(w.C, maxIdleConn),
	}
	client := &http.Client{Transport: tr, Timeout: time.Duration(w.Timeout) * time.Second}

	for i := 0; i < w.C; i++ {
		go func() {
			w.runWorker(w.ctx, client, w.N/w.C)
			wg.Done()
		}()
	}

	wg.Add(1)
	go func() {
		ticker := time.NewTicker(time.Second)
		defer func() {
			wg.Done()
			ticker.Stop()
		}()
		for {
			select {
			case <-w.ctx.Done():
				return
			case <-ticker.C:
				log.Infof("[completed requests]: %d", atomic.LoadInt64(&w.completedRequestsCount))
				log.Infof("[completed requests (errors)]: %d", atomic.LoadInt64(&w.completedWithErrors))
			}
		}
	}()

	wg.Wait()
	log.Info("all requests completed, wait some time...")
	<-time.After(time.Second * 5)
}

func (w *worker) httpRequestAsync(ctx context.Context, client *http.Client, payload *Payload) error {
	b, err := payload.Bytes()
	if err != nil {
		return err
	}

	u := w.Next()
	address := fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, "/internal/api/v1/event")
	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, address, bytes.NewReader(b),
	)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("undefined HTTP code: %d", res.StatusCode)
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
