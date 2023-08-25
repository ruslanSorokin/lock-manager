package lock

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

type Handler func(context.Context, *pb.LockReq) (*pb.LockRes, error)

func errToCode(e error) codes.Code {
	switch {
	case errors.Is(e, nil):
		return codes.OK

	case errors.Is(e, service.ErrInvalidResourceID):
		return codes.InvalidArgument

	case errors.Is(e, provider.ErrLockAlreadyExists):
		return codes.AlreadyExists

	default:
		return codes.Internal
	}
}

func errToMsg(e error) string {
	switch {
	case errors.Is(e, service.ErrInvalidResourceID):
		return "INVALID_RESOURCE_ID"

	case errors.Is(e, provider.ErrLockAlreadyExists):
		return "RESOURCE_ALREADY_LOCKED"

	default:
		return "INTERNAL_ERROR"
	}
}

func newPBRes(tkn *string) *pb.LockRes {
	return &pb.LockRes{
		Token: tkn,
	}
}

func response(
	e error,
	tkn *string,
) (*pb.LockRes, error) {
	code := errToCode(e)
	msg := errToMsg(e)
	return newPBRes(tkn), status.Error(code, msg)
}

func New(
	log logr.Logger,
	svc service.LockServiceI,
) Handler {
	logicalErrs := []error{
		service.ErrInvalidResourceID,
		provider.ErrLockAlreadyExists,
	}

	return func(
		ctx context.Context,
		req *pb.LockReq,
	) (*pb.LockRes, error) {
		rID := req.GetResourceId()
		tkn, err := svc.Lock(ctx, rID)
		code := errToCode(err)
		switch {
		case err == nil:
			return response(err, &tkn)

		case slices.Contains(logicalErrs, err):
			log.Error(err, "bad attempt to lock resource",
				LogKeyResourceID, rID,
				LogKeyToken, tkn,
				LogKeyGRPCCode, code)
			return response(err, nil)

		default:
			log.Error(err, "internal error during attempt to lock resource",
				LogKeyResourceID, rID,
				LogKeyToken, tkn,
				LogKeyGRPCCode, code)
			return response(err, nil)
		}
	}
}
