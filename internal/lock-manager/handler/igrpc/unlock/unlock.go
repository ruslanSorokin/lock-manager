package unlock

import (
	"context"

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

type (
	errToCodeFunc func(error) codes.Code
	errToMsgFunc  func(error) string
)

func newErrToCodeMapper() errToCodeFunc {
	errToCodeMapper := map[error]codes.Code{
		nil:                          codes.OK,
		service.ErrInvalidResourceID: codes.InvalidArgument,
		service.ErrInvalidToken:      codes.InvalidArgument,
		provider.ErrWrongToken:       codes.InvalidArgument,
		provider.ErrLockNotFound:     codes.NotFound,
	}

	return func(e error) codes.Code {
		if code, ok := errToCodeMapper[e]; ok {
			return code
		}
		return codes.Internal
	}
}

func newErrToMsgMapper() errToMsgFunc {
	errToCodeMapper := map[error]string{
		nil:                          "OK",
		service.ErrInvalidResourceID: "INVALID_RESOURCE_ID",
		service.ErrInvalidToken:      "INVALID_TOKEN",
		provider.ErrWrongToken:       "TOKEN_DOES_NOT_FIT",
		provider.ErrLockNotFound:     "RESOURCE_NOT_FOUND",
	}

	return func(e error) string {
		if code, ok := errToCodeMapper[e]; ok {
			return code
		}
		return "INTERNAL_ERROR"
	}
}

func newPBRes() *pb.UnlockRes {
	return &pb.UnlockRes{}
}

func New(log logr.Logger, svc service.LockServiceI) Handler {
	logicalErrs := []error{
		service.ErrInvalidResourceID,
		service.ErrInvalidToken,
		provider.ErrWrongToken,
		provider.ErrLockNotFound,
	}

	errToCode := newErrToCodeMapper()
	errToMsg := newErrToMsgMapper()

	return func(
		ctx context.Context,
		req *pb.UnlockReq,
	) (*pb.UnlockRes, error) {
		rID := req.GetResourceId()
		tkn := req.GetToken()

		err := svc.Unlock(ctx, rID, tkn)

		code := errToCode(err)
		msg := errToMsg(err)

		switch {
		case err == nil:
			return newPBRes(), nil

		case slices.Contains(logicalErrs, err):
			log.Error(err, "bad attempt to unlock resource",
				LogKeyResourceID, rID,
				LogKeyToken, tkn,
				LogKeyGRPCCode, code)
			return newPBRes(), status.Error(code, msg)

		default:
			log.Error(err, "internal error during attempt to unlock resource",
				LogKeyResourceID, rID,
				LogKeyToken, tkn,
				LogKeyGRPCCode, code)
			return newPBRes(), status.Error(code, msg)
		}
	}
}
