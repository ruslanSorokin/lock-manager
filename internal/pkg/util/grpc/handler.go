package grpcutil

import (
	"fmt"
	"net"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type HandlerI interface {
	Start() error
	GracefulStop()
	Stop()

	Server() *grpc.Server
}

type Handler struct {
	cfg *Config

	srv *grpc.Server
	log logr.Logger
}

var _ HandlerI = (*Handler)(nil)

func NewHandler(
	srv *grpc.Server,
	log logr.Logger,
	cfg *Config,
) *Handler {
	return &Handler{
		cfg: cfg,
		srv: srv,
		log: log,
	}
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

func (h Handler) Server() *grpc.Server { return h.srv }

func (h Handler) GracefulStop() { h.srv.GracefulStop() }

func (h Handler) Stop() { h.srv.Stop() }
