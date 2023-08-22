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
	cfg *Config

	log logr.Logger
	reg *prometheus.Registry
	mux *http.ServeMux
	srv *http.Server
}

func NewHandler(
	log logr.Logger,
	reg *prometheus.Registry,
	mux *http.ServeMux,
	port string,
	readTO time.Duration,
) *Handler {
	return &Handler{
		cfg: &Config{Port: port, ReadTimeOut: readTO},

		log: log,
		reg: reg,
		mux: mux,
		srv: &http.Server{
			ReadTimeout: readTO,
		},
	}
}

func NewHandlerFromConfig(
	log logr.Logger,
	reg *prometheus.Registry,
	mux *http.ServeMux,
	cfg *Config,
) *Handler {
	return NewHandler(log, reg, mux, cfg.Port, cfg.ReadTimeOut)
}

func (h Handler) Start() error {
	addr := fmt.Sprintf(":%s", h.cfg.Port)
	h.srv.Addr = addr
	h.srv.Handler = h.mux
	promHandler := promhttp.HandlerFor(h.reg, promhttp.HandlerOpts{})

	h.mux.Handle(pattern, promHandler)

	h.log.Info(fmt.Sprintf("metrics are up on %s", addr))
	return h.srv.ListenAndServe()
}

func (h Handler) GracefulStop() error {
	return h.srv.Shutdown(context.TODO())
}

func (h Handler) Stop() error {
	return h.srv.Close()
}
