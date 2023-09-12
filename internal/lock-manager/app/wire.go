//go:build wireinject
// +build wireinject

package app

import (
	"net/http"

	"github.com/go-logr/logr"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/wire"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc"
	ipromsvc "github.com/ruslanSorokin/lock-manager/internal/lock-manager/metric/iprom"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider/storage/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	redisconn "github.com/ruslanSorokin/lock-manager/internal/pkg/conn/redis"
	ipromapp "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app/iprom"
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
			"Pull",
			"Environment",
			"Version",
		),
		New,
		validator.New,
		redisconn.WireProvide,

		ipromapp.WireMetricSet,

		promutil.WireRegistrySet,

		httputil.WireHandlerFromConfigSet,

		grpcutil.WireInterceptorLoggerSet,
		grpcutil.WirePanicRecoveryHandlerSet,
		grpcutil.WireProcessingTimeHistogramSet,
		grpcutil.WireHandlerFromConfigSet,

		ipromgrpc.WireRecoveryMetricSet,

		http.NewServeMux,

		provideHTTPServer,

		grpcutil.WireProvideInterceptors,
		grpcutil.WireProvideServer,

		service.WireSet,
		ipromsvc.WireSet,

		iredis.WireLockStorageSet,

		igrpc.WireLockHandlerSet,
	))
}

func provideHTTPServer() *http.Server {
	return &http.Server{}
}
