package promutil

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	readTO time.Duration,
) *Handler {
	return &Handler{
		cfg: Config{Port: port, ReadTimeOut: readTO},

		log: log,
		reg: reg,
		mux: mux,
		srv: &http.Server{
			ReadTimeout: readTO,
		},
	}
}

func NewFromConfig(
	log logr.Logger,
	reg *prometheus.Registry,
	mux *http.ServeMux,
	cfg *Config,
) *Handler {
	return New(log, reg, mux, cfg.Port, cfg.ReadTimeOut)
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
