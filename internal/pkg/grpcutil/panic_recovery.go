package grpcutil

import (
	"runtime/debug"

	"github.com/go-logr/logr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecoveryMetricI interface {
	Inc()
}

func NewPanicRecoveryHandler(l logr.Logger, m RecoveryMetricI) func(any) error {
	return func(p any) error {
		m.Inc()

		l.Info("msg", "recovered from panic",
			"panic", p, "stack", debug.Stack())

		return status.Errorf(codes.Internal, "%s", p)
	}
}
