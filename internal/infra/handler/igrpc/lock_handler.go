package igrpc

import (
	"context"

	"github.com/go-logr/logr"

	pb "github.com/ruslanSorokin/lock-manager-api/gen/grpc/go"
	"github.com/ruslanSorokin/lock-manager/internal/infra/handler/igrpc/lock"
	"github.com/ruslanSorokin/lock-manager/internal/infra/handler/igrpc/unlock"
	"github.com/ruslanSorokin/lock-manager/internal/service"
	"github.com/ruslanSorokin/lock-manager/pkg/grpcutil"
)

type LockHandlerI interface {
	Lock(ctx context.Context, req *pb.LockReq) (*pb.LockRes, error)
	Unlock(ctx context.Context, req *pb.UnlockReq) (*pb.UnlockRes, error)
}

type LockHandler struct {
	pb.UnimplementedLockManagerServiceServer
	*grpcutil.Handler

	log logr.Logger
	svc service.LockServiceI

	lock   lock.Handler
	unlock unlock.Handler
}

func New(
	h *grpcutil.Handler,
	log logr.Logger,
	svc service.LockServiceI,
) LockHandler {
	lh := LockHandler{
		Handler: h,
		log:     log,
		svc:     svc,
		lock:    lock.New(log, svc),
		unlock:  unlock.New(log, svc),
	}

	pb.RegisterLockManagerServiceServer(h.Server(), lh)

	return lh
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
