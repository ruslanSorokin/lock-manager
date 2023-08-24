//go:build wireinject
// +build wireinject

package app

import (
	"net/http"

	"github.com/go-logr/logr"
	"github.com/google/wire"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"

	"github.com/ruslanSorokin/lock-manager/internal/pkg/apputil"
	ipromapp "github.com/ruslanSorokin/lock-manager/internal/pkg/apputil/iprom"

	promgrpc "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc"
	ipromsvc "github.com/ruslanSorokin/lock-manager/internal/lock-manager/imetric/iprom"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider/repository/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
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
			"LockService",
			"HTTPMetric",
			"Ver",
		),

		New,

		redisconn.WireProvide,

		ipromapp.MetricSet,

		promutil.HandlerFromConfigSet,
		promutil.RegistrySet,

		ipromgrpc.RecoveryMetricSet,

		grpcutil.InterceptorLoggerSet,
		grpcutil.PanicRecoveryHandlerSet,
		grpcutil.ProcessingTimeHistogramSet,

		http.NewServeMux,

		provideHTTPServer,
		provideGRPCServer,

		grpcutil.HandlerFromConfigSet,

		service.FromConfigSet,
		iredis.LockStorageSet,
		ipromsvc.Set,
		igrpc.LockHandlerSet,
	))
}

func provideHTTPServer() *http.Server {
	return &http.Server{}
}

func provideGRPCServer(
	log logging.Logger,
	metric *promgrpc.ServerMetrics,
	recoveryHandler func(any) error,
) *grpc.Server {
	unaryInters := []grpc.UnaryServerInterceptor{
		metric.UnaryServerInterceptor(),
		logging.UnaryServerInterceptor(log),
		recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(recoveryHandler)),
	}
	streamInters := []grpc.StreamServerInterceptor{
		metric.StreamServerInterceptor(),
		logging.StreamServerInterceptor(log),
		recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(recoveryHandler)),
	}

	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(unaryInters...),
		grpc.ChainStreamInterceptor(streamInters...))
}
