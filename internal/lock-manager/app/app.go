package app

import (
	"context"
	"net/http"

	"github.com/go-logr/logr"
	promgrpc "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/ifiber"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc"
	iservice "github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	redisconn "github.com/ruslanSorokin/lock-manager/internal/pkg/conn/redis"
	apputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app"
	httputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/http"
	promutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/prom"
)

type (
	lockStorageComp struct {
		redis *redisconn.Conn
	}
	grpcLockHandlerComp struct {
		server      *grpc.Server
		lockHandler igrpc.LockHandlerI
	}
	fiberLockHandlerComp struct {
		lockHandler *ifiber.LockHandler
	}
	lockHandlerComp struct {
		grpc grpcLockHandlerComp
		http fiberLockHandlerComp
	}
	metricComp struct {
		handler  httputil.HandlerI
		server   *http.Server
		mux      *http.ServeMux
		promReg  *prometheus.Registry
		promGRPC *promgrpc.ServerMetrics
		app      apputil.MetricI
	}
)

type App struct {
	cfg *Config
	log logr.Logger

	storage lockStorageComp
	handler lockHandlerComp
	metric  metricComp
	service iservice.LockServiceI

	environment apputil.Env
	version     apputil.Ver
}

func New(
	cfg *Config,
	log logr.Logger,
	service iservice.LockServiceI,
	redis *redisconn.Conn,
	grpcSrv *grpc.Server,
	grpcLockHandler igrpc.LockHandlerI,
	fiberLockHandler *ifiber.LockHandler,
	httpMetricServer *http.Server,
	httpMetricMux *http.ServeMux,
	httpMetricHandler httputil.HandlerI,
	appMetric apputil.MetricI,
	promReg *prometheus.Registry,
	promGRPCMetrics *promgrpc.ServerMetrics,
	environment apputil.Env,
	version apputil.Ver,
) *App {
	return &App{
		cfg:     cfg,
		log:     log,
		storage: lockStorageComp{redis: redis},
		handler: lockHandlerComp{
			grpc: grpcLockHandlerComp{server: grpcSrv, lockHandler: grpcLockHandler},
			http: fiberLockHandlerComp{lockHandler: fiberLockHandler},
		},
		metric: metricComp{
			server:   httpMetricServer,
			mux:      httpMetricMux,
			handler:  httpMetricHandler,
			promReg:  promReg,
			promGRPC: promGRPCMetrics,
			app:      appMetric,
		},
		service:     service,
		environment: environment,
		version:     version,
	}
}

func (a App) prepare() {
	mux := a.metric.handler.Mux()
	promutil.Register(mux, a.metric.promReg)

	grpcSrv := a.handler.grpc.lockHandler.Server()
	a.metric.promGRPC.InitializeMetrics(grpcSrv)

	a.log.Info("application environment", "env", a.environment.String())
	a.metric.app.SetVersion(a.version)
	a.log.Info("application version", "version", a.version.String())
	a.metric.app.SetEnvironment(a.environment)
}

func (a App) Run(ctx context.Context) error {
	a.prepare()

	rg := run.Group{}
	rg.Add(run.ContextHandler(ctx))

	grpcSrv := a.handler.grpc.lockHandler
	rg.Add(grpcSrv.Start,
		func(_ error) { grpcSrv.GracefulStop() })

	httpSrv := a.handler.http.lockHandler
	rg.Add(httpSrv.Start,
		func(_ error) {
			if err := httpSrv.GracefulStop(); err != nil {
				a.log.Error(
					err,
					"http server graceful shutdown error, shutting down forcefully",
				)
				httpSrv.Stop()
			}
		})

	metricSrv := a.metric.handler
	rg.Add(metricSrv.Start,
		func(_ error) {
			if err := metricSrv.GracefulStop(); err != nil {
				a.log.Error(
					err,
					"http metric server graceful shutdown error, shutting down forcefully",
				)
				metricSrv.Stop()
			}
		})

	return rg.Run()
}
