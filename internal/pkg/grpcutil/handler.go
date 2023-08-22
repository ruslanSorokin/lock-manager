package grpcutil

import (
	"fmt"
	"net"

	"github.com/go-logr/logr"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type HandlerI interface {
	Start() error
	Server() *grpc.Server
	GracefulStop()
	Stop()
}

type Handler struct {
	cfg Config

	srv *grpc.Server
	log logr.Logger
}

var _ HandlerI = (*Handler)(nil)

//nolint:gochecknoglobals // Wire Sets
var (
	HandlerSet           = wire.NewSet(NewHandler)
	HandlerFromConfigSet = wire.NewSet(NewHandlerFromConfig)
)

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
	cfg *Config,
) *Handler {
	return NewHandler(srv, log, cfg.Port, cfg.WithReflection)
}

func (h Handler) Start() error {
	h.log.Info("grpc reflection", "enabled", h.cfg.WithReflection)
	if h.cfg.WithReflection {
		reflection.Register(h.srv)
	}

	addr := fmt.Sprintf(":%s", h.cfg.Port)

	lst, err := net.Listen("tcp", addr)
	if err != nil {
		h.log.Error(err, fmt.Sprintf("unable to listen on %s", addr))
		return err
	}

	h.log.Info(fmt.Sprintf("grpc server is up on %s", addr))
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
