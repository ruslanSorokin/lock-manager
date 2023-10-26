package unlock

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
	"github.com/ruslanSorokin/lock-manager/internal/pkg/ierror"
)

type Handler func(context.Context, *pb.UnlockReq) (*pb.UnlockRes, error)

func newPBRes() *pb.UnlockRes {
	return &pb.UnlockRes{Response: nil}
}

func New(log logr.Logger, svc service.LockServiceI) Handler {
	const internalErrLogMsg = "internal error during attempt to unlock resource"
	const badAttemptLogMsg = "bad attempt to unlock resource"
	return func(
		ctx context.Context,
		req *pb.UnlockReq,
	) (*pb.UnlockRes, error) {
		rID := req.GetResourceId()
		tkn := req.GetToken()

		err := svc.Unlock(ctx, rID, tkn)
		if err == nil {
			return newPBRes(), nil
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

		return newPBRes(), status.Error(code, apiStCode)
	}
}
