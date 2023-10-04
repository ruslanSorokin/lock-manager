package unlock

import (
	"context"
	"errors"

	"github.com/go-logr/logr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ruslanSorokin/lock-manager-api/gen/grpc/go"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc/shared"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
)

type Handler func(context.Context, *pb.UnlockReq) (*pb.UnlockRes, error)

func newPBRes() *pb.UnlockRes {
	return &pb.UnlockRes{}
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
			GRPCStatusCode() codes.Code
			APIStatusCode() string
		}
		logMsg := internalErrLogMsg
		code := codes.Internal
		apiStCode := shared.APIStCodeInternalError

		if errors.As(err, &t) {
			logMsg = badAttemptLogMsg
			code = t.GRPCStatusCode()
			apiStCode = t.APIStatusCode()
		}

		log.Error(err, logMsg,
			shared.LogTagResourceID, rID,
			shared.LogTagToken, tkn,
			shared.LogTagGRPCCode, code)

		return newPBRes(), status.Error(code, apiStCode)
	}
}
