package iprom

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ruslanSorokin/lock-manager/internal/metric"
)

type Metrics struct {
	reg prometheus.Registerer

	lockedTotal   prometheus.Counter
	unlockedTotal prometheus.Counter
}

var _ metric.ServiceMetricI = (*Metrics)(nil)

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
	}

	m.reg.MustRegister(m.lockedTotal, m.unlockedTotal)
	return m
}

func (m Metrics) IncLockedTotal() {
	m.lockedTotal.Inc()
}

func (m Metrics) IncUnlockedTotal() {
	m.unlockedTotal.Inc()
}
