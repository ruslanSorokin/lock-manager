package lock

import (
	"context"
	"errors"

	"github.com/go-logr/logr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ruslanSorokin/lock-manager-api/gen/proto/go"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc/shared"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/ilog"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	"github.com/ruslanSorokin/lock-manager/pkg/ierror"
)

type Handler func(context.Context, *pb.LockReq) (*pb.LockRes, error)

func newPBRes(tkn *string) *pb.LockRes {
	return &pb.LockRes{
		Token: tkn,
	}
}

func New(
	log logr.Logger,
	svc service.LockServiceI,
) Handler {
	const internalErrLogMsg = "internal error during attempt to lock resource"
	const badAttemptLogMsg = "bad attempt to lock resource"
	return func(
		ctx context.Context,
		req *pb.LockReq,
	) (*pb.LockRes, error) {
		rID := req.GetResourceId()
		tkn, err := svc.Lock(ctx, rID)
		if err == nil {
			return newPBRes(&tkn), nil
		}

		var t interface {
			error
			ierror.GRPCConvertible
			ierror.EnumConvertible
		}
		logMsg := internalErrLogMsg
		code := codes.Internal
		apiStCode := shared.APIStCodeInternalError

		if errors.As(err, &t) {
			logMsg = badAttemptLogMsg
			code = t.ToGRPC()
			apiStCode = t.ToEnum()
		}

		log.Error(err, logMsg,
			ilog.TagResourceID, rID,
			ilog.TagToken, tkn,
			ilog.TagGRPCStCode, code)

		return newPBRes(nil), status.Error(code, apiStCode)
	}
}
