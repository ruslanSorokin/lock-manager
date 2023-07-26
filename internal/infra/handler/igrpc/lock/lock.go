package lock

import (
	"context"
	"errors"

	"github.com/go-logr/logr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ruslanSorokin/lock-manager-api/gen/grpc/go"
	"github.com/ruslanSorokin/lock-manager/internal/infra/provider"
	"github.com/ruslanSorokin/lock-manager/internal/service"
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

func respCode(e error) codes.Code {
	switch {
	case errors.Is(e, service.ErrInvalidResourceID):
		return codes.InvalidArgument

	case errors.Is(e, provider.ErrLockAlreadyExists):
		return codes.AlreadyExists

	default:
		return codes.Internal
	}
}

func pbErr(e error) pb.LockRes_Error {
	switch {
	case errors.Is(e, service.ErrInvalidResourceID):
		return pb.LockRes_ERR_INVALID_RESOURCE_ID

	case errors.Is(e, provider.ErrLockAlreadyExists):
		return pb.LockRes_ERR_LOCK_ALREADY_EXISTS

	default:
		return pb.LockRes_ERR_UNSPECIFIED
	}
}

func newResp(errs []pb.LockRes_Error, tkn *string) *pb.LockRes {
	return &pb.LockRes{
		Errors: errs,
		Token:  tkn,
	}
}

func response(
	code codes.Code,
	msg string,
	errs []pb.LockRes_Error,
	tkn *string,
) (*pb.LockRes, error) {
	var err error
	if code != codes.OK {
		err = status.Error(code, msg)
	}
	return newResp(errs, tkn), err
}

func New(
	log logr.Logger,
	svc service.LockServiceI,
) Handler {
	possibleErrs := []error{
		service.ErrInvalidResourceID,
		provider.ErrLockAlreadyExists,
	}

	return func(
		ctx context.Context,
		req *pb.LockReq,
	) (*pb.LockRes, error) {
		rID := req.GetResourceId()
		tkn, err := svc.Lock(ctx, rID)
		if err == nil {
			return response(codes.OK, "", nil, &tkn)
		}

		code := respCode(err)
		if code == codes.Internal {
			log.Error(err, "internal error during attempt to lock resource",
				LogKeyResourceID, rID,
				LogKeyToken, tkn,
				LogKeyGRPCCode, code)
			return response(code, msgInternalError, nil, nil)
		}

		// TODO: use slices.Contains to return internal error right here
		// It requires 1.20+ to use slices.Contains on error type, so we have to
		// wait until 1.21 release to make version 1.20 our HEAD~1 version.
		protoErrs := make([]pb.LockRes_Error, 0)
		for _, e := range possibleErrs {
			if errors.Is(err, e) {
				// Unwrap joined errors
				protoErrs = append(protoErrs, pbErr(e))
			}
		}

		if len(protoErrs) == 0 {
			log.Error(err, "bad assertion",
				"condition", "if err != nil && code != internal",
				"expected", "len(proto_errs) > 0")
			return response(code, msgInternalError, nil, nil)
		}
		msg := protoErrs[0].String()
		if len(protoErrs) > 1 {
			msg = "too many errors"
		}

		log.Error(err, "bad attempt to lock resource",
			LogKeyResourceID, rID,
			LogKeyToken, tkn,
			LogKeyGRPCCode, code)
		return response(code, msg, protoErrs, nil)
	}
}
