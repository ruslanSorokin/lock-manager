package httputil

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-logr/logr"
)

type HandlerI interface {
	Start() error
	GracefulStop() error
	Stop() error

	Server() *http.Server
	Mux() *http.ServeMux
}

type Handler struct {
	log logr.Logger

	srv *http.Server
	mux *http.ServeMux

	cfg *Config
}

var _ HandlerI = (*Handler)(nil)

func NewHandler(
	log logr.Logger,
	srv *http.Server,
	mux *http.ServeMux,
	port string,
	readTO time.Duration,
) *Handler {
	return &Handler{
		log: log,
		srv: srv,
		mux: mux,
		cfg: &Config{Port: port, ReadTimeOut: readTO},
	}
}

func NewHandlerFromConfig(
	log logr.Logger,
	srv *http.Server,
	mux *http.ServeMux,
	cfg *Config,
) *Handler {
	return &Handler{
		log: log,
		srv: srv,
		mux: mux,
		cfg: cfg,
	}
}

func (h Handler) Start() error {
	addr := fmt.Sprintf(":%s", h.cfg.Port)

	h.srv.Addr = addr
	h.srv.Handler = h.mux

	h.log.Info(fmt.Sprintf("http server is up on %s", addr))
	return h.srv.ListenAndServe()
}

func (h Handler) Stop() error {
	return h.srv.Shutdown(context.TODO())
}

func (h Handler) GracefulStop() error {
	return h.srv.Close()
}

func (h Handler) Server() *http.Server {
	return h.srv
}

func (h Handler) Mux() *http.ServeMux {
	return h.mux
}
