package app

import (
	"context"
	"net/http"

	"github.com/go-logr/logr"
	promgrpc "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc"
	iservice "github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	redisconn "github.com/ruslanSorokin/lock-manager/internal/pkg/conn/redis"
	apputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app"
	httputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/http"
	promutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/prom"
)

type (
	repo struct {
		redisConn *redisconn.Conn
	}
	srv struct {
		grpcServer  *grpc.Server
		grpcHandler igrpc.LockHandlerI
	}
	mtr struct {
		promReg     *prometheus.Registry
		promGRPC    *promgrpc.ServerMetrics
		httpSrv     *http.Server
		mux         *http.ServeMux
		httpHandler *httputil.Handler
		app         apputil.MetricI
	}
)

type App struct {
	cfg *Config
	log logr.Logger

	storage repo
	server  srv
	metric  mtr
	service iservice.LockServiceI

	environment apputil.Env
	version     apputil.Ver
}

func New(
	cfg *Config,
	log logr.Logger,
	redisConn *redisconn.Conn,
	promReg *prometheus.Registry,
	promGRPCMetric *promgrpc.ServerMetrics,
	mux *http.ServeMux,
	httpSrv *http.Server,
	grpcSrv *grpc.Server,
	svc iservice.LockServiceI,
	grpcHandler igrpc.LockHandlerI,
	httpMtrHandler *httputil.Handler,
	app apputil.MetricI,
	env apputil.Env,
	ver apputil.Ver,
) *App {
	return &App{
		cfg:     cfg,
		log:     log,
		storage: repo{redisConn: redisConn},
		server: srv{
			grpcServer:  grpcSrv,
			grpcHandler: grpcHandler,
		},
		metric: mtr{
			promReg:     promReg,
			promGRPC:    promGRPCMetric,
			httpSrv:     httpSrv,
			mux:         mux,
			httpHandler: httpMtrHandler,
			app:         app,
		},
		service:     svc,
		environment: env,
		version:     ver,
	}
}

func (a App) prepare() {
	mux := a.metric.httpHandler.Mux()
	promutil.Register(mux, a.metric.promReg)

	grpcSrv := a.server.grpcHandler.Server()
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

	rg.Add(a.server.grpcHandler.Start,
		func(_ error) {
			a.server.grpcHandler.GracefulStop()
		})

	rg.Add(a.metric.httpHandler.Start,
		func(err error) {
			if err := a.metric.httpHandler.GracefulStop(); err != nil {
				a.log.Error(err, "http metric server shutdown error")
			}
		})

	return rg.Run()
}
