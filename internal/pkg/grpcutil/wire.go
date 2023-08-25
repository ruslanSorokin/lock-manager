//go:build wireinject
// +build wireinject

package grpcutil

import "github.com/google/wire"

//nolint:gochecknoglobals // Wire Sets
var (
	WireHandlerSet           = wire.NewSet(NewHandler)
	WireHandlerFromConfigSet = wire.NewSet(NewHandlerFromConfig)

	WireInterceptorLoggerSet       = wire.NewSet(NewInterceptorLogger)
	WirePanicRecoveryHandlerSet    = wire.NewSet(NewPanicRecoveryHandler)
	WireProcessingTimeHistogramSet = wire.NewSet(NewProcessingTimeHistogram)
)
