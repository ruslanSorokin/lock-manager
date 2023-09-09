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
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider/repository/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	redisconn "github.com/ruslanSorokin/lock-manager/internal/pkg/conn/redis"
	apputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app"
	ipromapp "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app/iprom"
	grpcutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/grpc"
	ipromgrpc "github.com/ruslanSorokin/lock-manager/internal/pkg/util/grpc/iprom"
	promutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/prom"
)

func Wire(apputil.Env, logr.Logger, *Config) (*App, func(), error) {
	panic(wire.Build(
		wire.FieldsOf(new(*Config),
			"Redis",
			"GRPC",
			"HTTPMetric",
			"Ver",
		),
		New,

		validator.New,
		redisconn.WireProvide,

		ipromapp.WireMetricSet,

		promutil.WireHandlerFromConfigSet,
		promutil.WireRegistrySet,

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
