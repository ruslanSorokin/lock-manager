//go:build wireinject
// +build wireinject

package grpcutil

import "github.com/google/wire"

//nolint:gochecknoglobals // Wire Set
var (
	InterceptorLoggerSet       = wire.NewSet(NewInterceptorLogger)
	PanicRecoveryHandlerSet    = wire.NewSet(NewPanicRecoveryHandler)
	ProcessingTimeHistogramSet = wire.NewSet(NewProcessingTimeHistogram)
)
