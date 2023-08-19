package grpcutil

import (
	iprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

func NewProcessingTimeHistogram(r prometheus.Registerer) *iprom.ServerMetrics {
	m := iprom.NewServerMetrics(
		iprom.WithServerHandlingTimeHistogram(
			iprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	r.MustRegister(m)

	return m
}
