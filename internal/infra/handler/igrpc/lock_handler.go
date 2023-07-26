package igrpc

import (
	"context"
	"fmt"
	"net"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"

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

func NewLockHandler(srv *grpc.Server, log logr.Logger, svc service.LockServiceI, ip, port string) LockHandler {
	return LockHandler{
		srv:    srv,
		log:    log,
		svc:    svc,
		cfg:    Config{IP: ip, Port: port},
		lock:   lock.New(log, svc),
		unlock: unlock.New(log, svc),
	}
}

func NewLockHandlerFromConfig(srv *grpc.Server, log logr.Logger, svc service.LockServiceI, cfg Config) LockHandler {
	return LockHandler{
		srv:    srv,
		log:    log,
		svc:    svc,
		cfg:    cfg,
		lock:   lock.New(log, svc),
		unlock: unlock.New(log, svc),
	}
}

func (h LockHandler) Register() {
	pb.RegisterLockManagerServiceServer(h.srv, h)
}

func (h LockHandler) Start() error {
	addr := fmt.Sprintf("%s:%s", h.cfg.IP, h.cfg.Port)

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
