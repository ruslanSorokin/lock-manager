package unlock

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

type Handler func(context.Context, *pb.UnlockReq) (*pb.UnlockRes, error)

func respCode(e error) codes.Code {
	switch {
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

func pbErr(e error) pb.UnlockRes_Error {
	switch {
	case errors.Is(e, service.ErrInvalidResourceID):
		return pb.UnlockRes_ERR_INVALID_RESOURCE_ID

	case errors.Is(e, service.ErrInvalidToken):
		return pb.UnlockRes_ERR_INVALID_TOKEN

	case errors.Is(e, provider.ErrLockNotFound):
		return pb.UnlockRes_ERR_LOCK_NOT_FOUND

	case errors.Is(e, provider.ErrWrongToken):
		return pb.UnlockRes_ERR_WRONG_TOKEN

	default:
		return pb.UnlockRes_ERR_UNSPECIFIED
	}
}

func newPbRes(errs []pb.UnlockRes_Error) *pb.UnlockRes {
	return &pb.UnlockRes{Errors: errs}
}

func response(
	code codes.Code,
	msg string,
	errs []pb.UnlockRes_Error,
) (*pb.UnlockRes, error) {
	var err error
	if code != codes.OK {
		err = status.Error(code, msg)
	}
	return newPbRes(errs), err
}

func New(log logr.Logger, svc service.LockServiceI) Handler {
	possibleErrs := []error{
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
		if err == nil {
			return response(codes.OK, "", nil)
		}

		code := respCode(err)
		if code == codes.Internal {
			log.Error(err, "internal error during attempt to unlock resource",
				LogKeyResourceID, rID,
				LogKeyToken, tkn,
				LogKeyGRPCCode, code)
			return response(code, msgInternalError, nil)
		}

		// TODO: use slices.Contains to return internal error right here
		// It requires 1.20+ to use slices.Contains on error type, so we have to
		// wait until 1.21 release to make version 1.20 our HEAD~1 version.
		protoErrs := make([]pb.UnlockRes_Error, 0)
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
			return response(code, msgInternalError, nil)
		}
		msg := protoErrs[0].String()
		if len(protoErrs) > 1 {
			msg = "too many errors"
		}

		log.Error(err, "bad attempt to unlock resource",
			LogKeyResourceID, rID,
			LogKeyToken, tkn,
			LogKeyGRPCCode, code)
		return response(code, msg, protoErrs)
	}
}
