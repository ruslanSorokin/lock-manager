package promutil

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const pattern = "/metrics"

type Handler struct {
	cfg Config

	log logr.Logger
	reg *prometheus.Registry
	mux *http.ServeMux
	srv *http.Server
}

func New(
	log logr.Logger,
	reg *prometheus.Registry,
	mux *http.ServeMux,
	port string,
) *Handler {
	return &Handler{
		cfg: Config{Port: port},

		log: log,
		reg: reg,
		mux: mux,
		srv: &http.Server{},
	}
}

func (m Handler) Start() error {
	m.srv.Addr = fmt.Sprintf(":%s", m.cfg.Port)
	m.srv.Handler = m.mux
	promHandler := promhttp.HandlerFor(m.reg, promhttp.HandlerOpts{})

	m.mux.Handle(pattern, promHandler)
	return m.srv.ListenAndServe()
}

func (m Handler) GracefulStop() error {
	return m.srv.Shutdown(context.TODO())
}

func (m Handler) Stop() error {
	return m.srv.Close()
}
