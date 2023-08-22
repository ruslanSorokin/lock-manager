package grpcutil

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

const (
	debugVerbosity = 4
	infoVerbosity  = 2
	warnVerbosity  = 1
	errorVerbosity = 0
)

func NewInterceptorLogger(log logr.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(_ context.Context, lvl grpclog.Level, msg string, args ...any) {
		l := log.WithValues(args...)
		switch lvl {
		case grpclog.LevelDebug:
			l.V(debugVerbosity).Info(msg)
		case grpclog.LevelInfo:
			l.V(infoVerbosity).Info(msg)
		case grpclog.LevelWarn:
			l.V(warnVerbosity).Info(msg)
		case grpclog.LevelError:
			l.V(errorVerbosity).Info(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
