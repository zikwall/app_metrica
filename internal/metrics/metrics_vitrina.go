package metrics

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type MetricsVitrina struct {
	rm *prometheus.CounterVec
	qm *prometheus.CounterVec

	gatewayErr *prometheus.CounterVec
	queueErr   *prometheus.CounterVec

	rmDur *prometheus.HistogramVec
	qmDur *prometheus.HistogramVec
}

func (m *MetricsVitrina) IncRequests(isBatch bool) {
	m.rm.WithLabelValues(strconv.FormatBool(isBatch)).Inc()
}

func (m *MetricsVitrina) IncQueue(size int, byTicker bool, err error) {
	m.qm.WithLabelValues(strconv.Itoa(size), strconv.FormatBool(byTicker), strconv.FormatBool(err != nil)).Inc()
}

func (m *MetricsVitrina) RequestsDuration(start time.Time, isBatch bool) {
	m.rmDur.WithLabelValues(strconv.FormatBool(isBatch)).Observe(time.Since(start).Seconds())
}

func (m *MetricsVitrina) QueueDuration(start time.Time, size int, err error) {
	m.qmDur.WithLabelValues(strconv.Itoa(size), strconv.FormatBool(err != nil)).Observe(time.Since(start).Seconds())
}

func (m *MetricsVitrina) IncConsumerError(err error) {
	m.queueErr.WithLabelValues(err.Error()).Inc()
}

func (m *MetricsVitrina) IncGatewayError(err error) {
	m.gatewayErr.WithLabelValues(err.Error()).Inc()
}

func NewVitrina() *MetricsVitrina {
	m := &MetricsVitrina{
		rm: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "mediavitrina",
			Subsystem: "gateway",
			Name:      "requests",
		},
			[]string{"is_batch"},
		),
		qm: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "mediavitrina",
			Subsystem: "gateway",
			Name:      "queue",
		},
			[]string{"size", "by_ticker", "err"},
		),

		queueErr: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "mediavitrina",
			Subsystem: "consumer",
			Name:      "errors",
		},
			[]string{"err"},
		),

		gatewayErr: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "mediavitrina",
			Subsystem: "gateway",
			Name:      "errors",
		},
			[]string{"err"},
		),

		rmDur: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "mediavitrina",
				Subsystem: "gateway",
				Name:      "requests_duration",
				Buckets:   prometheus.LinearBuckets(0.005, 0.05, 10),
			},
			[]string{"is_batch"},
		),
		qmDur: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "mediavitrina",
				Subsystem: "gateway",
				Name:      "queue_duration",
				Buckets:   prometheus.LinearBuckets(0.005, 0.05, 10),
			},
			[]string{"size", "err"},
		),
	}

	prometheus.MustRegister(m.rm, m.qm, m.queueErr, m.gatewayErr, m.rmDur, m.qmDur)

	return m
}
