package igrpc

import (
	"context"
	"fmt"
	"net"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/ruslanSorokin/lock-manager-api/gen/grpc/go"
	"github.com/ruslanSorokin/lock-manager/internal/infra/handler/igrpc/lock"
	"github.com/ruslanSorokin/lock-manager/internal/infra/handler/igrpc/unlock"
	"github.com/ruslanSorokin/lock-manager/internal/service"
)

type LockHandlerI interface {
	Lock(ctx context.Context, req *pb.LockReq) (*pb.LockRes, error)
	Unlock(ctx context.Context, req *pb.UnlockReq) (*pb.UnlockRes, error)
}

type LockHandler struct {
	pb.UnimplementedLockManagerServiceServer

	srv *grpc.Server

	log logr.Logger
	svc service.LockServiceI
	cfg Config

	lock   lock.Handler
	unlock unlock.Handler
}

func New(
	srv *grpc.Server,
	log logr.Logger,
	svc service.LockServiceI,
	port string, withReflection bool,
) LockHandler {
	h := LockHandler{
		srv:    srv,
		log:    log,
		svc:    svc,
		cfg:    Config{Port: port},
		lock:   lock.New(log, svc),
		unlock: unlock.New(log, svc),
	}
	pb.RegisterLockManagerServiceServer(h.srv, h)

	if withReflection {
		log.Info("grpc reflection enabled")
		reflection.Register(srv)
	}

	return h
}

func NewFromConfig(
	srv *grpc.Server,
	log logr.Logger,
	svc service.LockServiceI,
	cfg Config,
) LockHandler {
	return New(srv, log, svc, cfg.Port, cfg.Reflection)
}

func (h LockHandler) Start() error {
	addr := fmt.Sprintf(":%s", h.cfg.Port)

	lst, err := net.Listen("tcp", addr)
	if err != nil {
		h.log.Error(err, fmt.Sprintf("unable to listen on %s", addr))
		return err
	}

	h.log.Info(fmt.Sprintf("listen on %s", addr))
	return h.srv.Serve(lst)
}

func (h LockHandler) GracefulStop() {
	h.srv.GracefulStop()
}

func (h LockHandler) Stop() {
	h.srv.Stop()
}

func (h LockHandler) Lock(
	ctx context.Context,
	req *pb.LockReq,
) (*pb.LockRes, error) {
	return h.lock(ctx, req)
}

func (h LockHandler) Unlock(
	ctx context.Context,
	req *pb.UnlockReq,
) (*pb.UnlockRes, error) {
	return h.unlock(ctx, req)
}
