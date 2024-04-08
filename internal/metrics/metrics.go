package metrics

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	rm *prometheus.CounterVec
	qm *prometheus.CounterVec

	rmDur *prometheus.HistogramVec
	qmDur *prometheus.HistogramVec
}

func (m *Metrics) IncRequests(isBatch bool) {
	m.rm.WithLabelValues(strconv.FormatBool(isBatch)).Inc()
}

func (m *Metrics) IncQueue(size int, byTicker bool, err error) {
	m.qm.WithLabelValues(strconv.Itoa(size), strconv.FormatBool(byTicker), strconv.FormatBool(err != nil)).Inc()
}

func (m *Metrics) RequestsDuration(start time.Time, isBatch bool) {
	m.rmDur.WithLabelValues(strconv.FormatBool(isBatch)).Observe(time.Since(start).Seconds())
}

func (m *Metrics) QueueDuration(start time.Time, size int, err error) {
	m.qmDur.WithLabelValues(strconv.Itoa(size), strconv.FormatBool(err != nil)).Observe(time.Since(start).Seconds())
}

func New() *Metrics {
	m := &Metrics{
		rm: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "own_metrics",
			Subsystem: "gateway",
			Name:      "requests",
		},
			[]string{"is_batch"},
		),
		qm: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "own_metrics",
			Subsystem: "gateway",
			Name:      "queue",
		},
			[]string{"size", "by_ticker", "err"},
		),

		rmDur: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "own_metrics",
				Subsystem: "gateway",
				Name:      "requests_duration",
				Buckets:   prometheus.LinearBuckets(0.005, 0.05, 10),
			},
			[]string{"is_batch"},
		),
		qmDur: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "own_metrics",
				Subsystem: "gateway",
				Name:      "queue_duration",
				Buckets:   prometheus.LinearBuckets(0.005, 0.05, 10),
			},
			[]string{"size", "err"},
		),
	}

	prometheus.MustRegister(m.rm, m.qm, m.rmDur, m.qmDur)

	return m
}
