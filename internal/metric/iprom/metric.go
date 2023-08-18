package iprom

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ruslanSorokin/lock-manager/internal/metric"
)

const (
	verLabel = "version"
	envLabel = "environment"
)

type Metrics struct {
	reg prometheus.Registerer

	lockedTotal   prometheus.Counter
	unlockedTotal prometheus.Counter

	Version     *prometheus.GaugeVec
	Environment *prometheus.GaugeVec
}

var _ metric.MetricI = (*Metrics)(nil)

func New(r prometheus.Registerer) Metrics {
	m := Metrics{
		reg: r,

		lockedTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "locked_total",
			Help: "Number of locked resources in total.",
		}),
		unlockedTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "unlocked_total",
			Help: "Number of unlocked resources in total.",
		}),

		Environment: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "environment_type",
			Help: "Type of environment of currently loaded config.",
		}, []string{envLabel}),
		Version: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "application_version",
			Help: "Version of currently running application.",
		}, []string{verLabel}),
	}

	m.reg.MustRegister(m.Environment, m.Version)
	return m
}

func (m Metrics) SetAppVersion(ver string) {
	m.Version.With(prometheus.Labels{verLabel: ver}).Set(1)
}

func (m Metrics) SetAppEnvironment(env string) {
	m.Environment.With(prometheus.Labels{envLabel: env}).Set(1)
}

func (m Metrics) IncLockedTotal() {
	m.lockedTotal.Inc()
}

func (m Metrics) IncUnlockedTotal() {
	m.unlockedTotal.Inc()
}
