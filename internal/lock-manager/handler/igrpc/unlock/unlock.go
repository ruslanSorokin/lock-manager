package unlock

import (
	"context"
	"errors"

	"github.com/go-logr/logr"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ruslanSorokin/lock-manager-api/gen/grpc/go"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
)

const (
	LogKeyResourceID = "resource_id"
	LogKeyToken      = "token"
	LogKeyGRPCCode   = "grpc_code"
)

type Handler func(context.Context, *pb.UnlockReq) (*pb.UnlockRes, error)

func errToCode(e error) codes.Code {
	switch {
	case errors.Is(e, nil):
		return codes.OK

	case errors.Is(e, service.ErrInvalidResourceID):
		return codes.InvalidArgument

	case errors.Is(e, service.ErrInvalidToken):
		return codes.InvalidArgument

	case errors.Is(e, provider.ErrWrongToken):
		return codes.InvalidArgument

	case errors.Is(e, provider.ErrLockNotFound):
		return codes.NotFound

	default:
		return codes.Internal
	}
}

func errToMsg(e error) string {
	switch {
	case errors.Is(e, nil):
		return "resource has been unlocked"

	case errors.Is(e, service.ErrInvalidResourceID):
		return "invalid resource id"

	case errors.Is(e, service.ErrInvalidToken):
		return "invalid token"

	case errors.Is(e, provider.ErrWrongToken):
		return "token doesn't fit"

	case errors.Is(e, provider.ErrLockNotFound):
		return "resource not found"

	default:
		return "internal error"
	}
}

func newPBRes() *pb.UnlockRes {
	return &pb.UnlockRes{}
}

func response(
	e error,
) (*pb.UnlockRes, error) {
	code := errToCode(e)
	msg := errToMsg(e)
	return newPBRes(), status.Error(code, msg)
}

func New(log logr.Logger, svc service.LockServiceI) Handler {
	logicalErrs := []error{
		service.ErrInvalidResourceID,
		service.ErrInvalidToken,
		provider.ErrWrongToken,
		provider.ErrLockNotFound,
	}

	return func(
		ctx context.Context,
		req *pb.UnlockReq,
	) (*pb.UnlockRes, error) {
		rID := req.GetResourceId()
		tkn := req.GetToken()

		err := svc.Unlock(ctx, rID, tkn)
		switch {
		case err == nil:
			return response(err)

		case slices.Contains(logicalErrs, err):
			log.Error(err, "bad attempt to unlock resource",
				LogKeyResourceID, rID,
				LogKeyToken, tkn,
				LogKeyGRPCCode, errToCode(err))
			return response(err)

		default:
			log.Error(err, "internal error during attempt to unlock resource",
				LogKeyResourceID, rID,
				LogKeyToken, tkn,
				LogKeyGRPCCode, errToCode(err))
			return response(err)
		}
	}
}
