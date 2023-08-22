package lock

import (
	"context"
	"errors"

	"github.com/go-logr/logr"
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

const (
	msgInternalError = "internal error"
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
		return "invalid resource id"

	case errors.Is(e, provider.ErrLockAlreadyExists):
		return "resource already locked"

	default:
		return "internal error"
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
	return func(
		ctx context.Context,
		req *pb.LockReq,
	) (*pb.LockRes, error) {
		rID := req.GetResourceId()
		tkn, err := svc.Lock(ctx, rID)
		code := errToCode(err)
		switch code {
		case codes.OK:
			return response(err, &tkn)

		case codes.Internal:
			log.Error(err, "internal error during attempt to lock resource",
				LogKeyResourceID, rID,
				LogKeyToken, tkn,
				LogKeyGRPCCode, code)

		default:
			log.Error(err, "bad attempt to lock resource",
				LogKeyResourceID, rID,
				LogKeyToken, tkn,
				LogKeyGRPCCode, code)

		}

		return response(err, nil)
	}
}