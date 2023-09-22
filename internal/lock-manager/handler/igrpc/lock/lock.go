package lock

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

type Handler func(context.Context, *pb.LockReq) (*pb.LockRes, error)

type (
	errToCodeFunc func(error) codes.Code
	errToMsgFunc  func(error) string
)

func newErrToCodeMapper() errToCodeFunc {
	errToCodeMapper := map[error]codes.Code{
		nil:                           codes.OK,
		service.ErrInvalidResourceID:  codes.InvalidArgument,
		provider.ErrLockAlreadyExists: codes.AlreadyExists,
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
		nil:                           "OK",
		service.ErrInvalidResourceID:  "INVALID_RESOURCE_ID",
		provider.ErrLockAlreadyExists: "RESOURCE_ALREADY_LOCKED",
	}

	return func(e error) string {
		if code, ok := errToCodeMapper[e]; ok {
			return code
		}
		return "INTERNAL_ERROR"
	}
}

func newPBRes(tkn *string) *pb.LockRes {
	return &pb.LockRes{
		Token: tkn,
	}
}

func New(
	log logr.Logger,
	svc service.LockServiceI,
) Handler {
	logicalErrs := []error{
		service.ErrInvalidResourceID,
		provider.ErrLockAlreadyExists,
	}
	errToCode := newErrToCodeMapper()
	errToMsg := newErrToMsgMapper()

	return func(
		ctx context.Context,
		req *pb.LockReq,
	) (*pb.LockRes, error) {
		rID := req.GetResourceId()
		tkn, err := svc.Lock(ctx, rID)

		code := errToCode(err)
		msg := errToMsg(err)

		switch {
		case err == nil:
			return newPBRes(&tkn), nil

		case slices.Contains(logicalErrs, err):
			log.Error(err, "bad attempt to lock resource",
				LogKeyResourceID, rID,
				LogKeyToken, tkn,
				LogKeyGRPCCode, code)
			return newPBRes(nil), status.Error(code, msg)

		default:
			log.Error(err, "internal error during attempt to lock resource",
				LogKeyResourceID, rID,
				LogKeyToken, tkn,
				LogKeyGRPCCode, code)
			return newPBRes(nil), status.Error(code, msg)
		}
	}
}
