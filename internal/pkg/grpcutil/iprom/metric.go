package iprom

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/grpcutil"
)

type RecoveryMetric struct {
	mtr prometheus.Counter
}

var _ grpcutil.RecoveryMetricI = (*RecoveryMetric)(nil)

func NewRecoveryMetric(r prometheus.Registerer) *RecoveryMetric {
	return &RecoveryMetric{
		mtr: promauto.With(r).NewCounter(prometheus.CounterOpts{
			Name: "grpc_panic_recovered_total",
			Help: "Total number of gRPC requests recovered from internal panic.",
		}),
	}
}

func (m RecoveryMetric) Inc() {
	m.mtr.Inc()
}
