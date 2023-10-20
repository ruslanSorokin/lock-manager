package promutil

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const pattern = "/metrics"

func Register(mux *http.ServeMux, reg *prometheus.Registry) {
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	mux.Handle(pattern, promHandler)
}
