package iprom

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/apputil"
)

const (
	verLabel = "version"
	envLabel = "environment"
)

type Metric struct {
	reg prometheus.Registerer

	Version     *prometheus.GaugeVec
	Environment *prometheus.GaugeVec
}

var _ apputil.MetricI = (*Metric)(nil)

func New(r prometheus.Registerer) *Metric {
	return &Metric{
		reg: r,

		Environment: promauto.With(r).NewGaugeVec(prometheus.GaugeOpts{
			Name: "environment_type",
			Help: "Type of environment of currently loaded config.",
		}, []string{envLabel}),
		Version: promauto.With(r).NewGaugeVec(prometheus.GaugeOpts{
			Name: "application_version",
			Help: "Version of currently running application.",
		}, []string{verLabel}),
	}
}

func (m Metric) SetVersion(ver apputil.Ver) {
	m.Version.With(prometheus.Labels{verLabel: string(ver)}).Set(1)
}

func (m Metric) SetEnvironment(env apputil.Env) {
	m.Environment.With(prometheus.Labels{envLabel: string(env)}).Set(1)
}
