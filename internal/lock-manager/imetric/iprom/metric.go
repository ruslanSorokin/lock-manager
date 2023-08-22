package iprom

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/imetric"
)

type Metric struct {
	reg prometheus.Registerer

	lockedTotal   prometheus.Counter
	unlockedTotal prometheus.Counter
}

var _ imetric.ServiceMetricI = (*Metric)(nil)

func New(r prometheus.Registerer) *Metric {
	m := &Metric{
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

func (m Metric) IncLockedTotal() {
	m.lockedTotal.Inc()
}

func (m Metric) IncUnlockedTotal() {
	m.unlockedTotal.Inc()
}
