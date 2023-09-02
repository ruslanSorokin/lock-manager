//go:build wireinject
// +build wireinject

package app

import (
	"net/http"

	"github.com/go-logr/logr"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/wire"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc"
	ipromsvc "github.com/ruslanSorokin/lock-manager/internal/lock-manager/imetric/iprom"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider/repository/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/apputil"
	ipromapp "github.com/ruslanSorokin/lock-manager/internal/pkg/apputil/iprom"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/grpcutil"
	ipromgrpc "github.com/ruslanSorokin/lock-manager/internal/pkg/grpcutil/iprom"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/promutil"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/redisconn"
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
		redisconn.WireProvideConn,

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
