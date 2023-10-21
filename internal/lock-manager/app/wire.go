//go:build wireinject
// +build wireinject

package app

import (
	"net/http"

	"github.com/go-logr/logr"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/wire"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/ifiber"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc"
	ipromsvc "github.com/ruslanSorokin/lock-manager/internal/lock-manager/metric/iprom"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider/storage/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	redisconn "github.com/ruslanSorokin/lock-manager/internal/pkg/conn/redis"
	ipromapp "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app/iprom"
	fiberutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/fiber"
	grpcutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/grpc"
	ipromgrpc "github.com/ruslanSorokin/lock-manager/internal/pkg/util/grpc/iprom"
	httputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/http"
	promutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/prom"
)

func Wire(logr.Logger, *Config) (*App, func(), error) {
	panic(wire.Build(
		toWireConfig,
		wire.FieldsOf(new(*wireConfig),
			"Redis",
			"GRPC",
			"HTTP",
			"Pull",
			"Environment",
			"Version",
		),

		New,
		provideValidator,
		redisconn.WireProvide,
		http.NewServeMux,
		httputil.WireProvideServer,
		grpcutil.WireProvideDefaultServerOpts,
		grpcutil.WireProvideServer,
		fiberutil.WireProvideDefaultMiddleware,
		fiberutil.WireProvideDefaultApp,
		grpcutil.NewInterceptorLogger,
		grpcutil.NewPanicRecoveryHandler,
		grpcutil.NewProcessingTimeHistogram,
		ifiber.NewLockHandler,

		grpcutil.WireHandlerSet,
		httputil.WireHandlerSet,
		igrpc.WireLockHandlerSet,
		fiberutil.WireHandlerSet,
		ipromgrpc.WireRecoveryMetricSet,
		ipromsvc.WireSet,
		ipromapp.WireMetricSet,
		promutil.WireRegistrySet,
		service.WireSet,
		iredis.WireLockStorageSet,
	))
}

func provideValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}
