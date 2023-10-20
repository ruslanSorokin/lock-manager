package igrpc

import (
	"context"

	"github.com/go-logr/logr"

	pb "github.com/ruslanSorokin/lock-manager-api/gen/grpc/go"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc/lock"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc/unlock"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	grpcutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/grpc"
)

type LockHandlerI interface {
	grpcutil.HandlerI

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

var _ grpcutil.HandlerI = (*LockHandler)(nil)

func NewLockHandler(
	h *grpcutil.Handler,
	log logr.Logger,
	svc service.LockServiceI,
) LockHandler {
	lh := LockHandler{
		UnimplementedLockManagerServiceServer: pb.UnimplementedLockManagerServiceServer{},
		Handler:                               h,
		log:                                   log,
		svc:                                   svc,
		lock:                                  lock.New(log, svc),
		unlock:                                unlock.New(log, svc),
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
