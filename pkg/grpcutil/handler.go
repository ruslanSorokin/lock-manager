package grpcutil

import (
	"fmt"
	"net"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Handler struct {
	cfg Config

	srv *grpc.Server
	log logr.Logger
}

func NewHandler(
	srv *grpc.Server,
	log logr.Logger,
	port string, withReflection bool,
) *Handler {
	return &Handler{
		cfg: Config{Port: port, WithReflection: withReflection},
		srv: srv,
		log: log,
	}
}

func NewHandlerFromConfig(
	srv *grpc.Server,
	log logr.Logger,
	cfg Config,
) *Handler {
	return NewHandler(srv, log, cfg.Port, cfg.WithReflection)
}

func (h Handler) preStart() {
	if h.cfg.WithReflection {
		h.log.Info("grpc_reflection", "enabled")
		reflection.Register(h.srv)
	} else {
		h.log.Info("grpc_reflection", "disabled")
	}
}

func (h Handler) Start() error {
	h.preStart()
	addr := fmt.Sprintf(":%s", h.cfg.Port)

	lst, err := net.Listen("tcp", addr)
	if err != nil {
		h.log.Error(err, fmt.Sprintf("unable to listen on %s", addr))
		return err
	}

	h.log.Info(fmt.Sprintf("listen on %s", addr))
	return h.srv.Serve(lst)
}

func (h Handler) Server() *grpc.Server {
	return h.srv
}

func (h Handler) GracefulStop() {
	h.srv.GracefulStop()
}

func (h Handler) Stop() {
	h.srv.Stop()
}
