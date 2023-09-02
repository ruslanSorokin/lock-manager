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
	"github.com/ruslanSorokin/lock-manager/internal/pkg/apputil"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/promutil"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/redisconn"
)

type (
	repositoryComponent struct {
		redisConn *redisconn.Conn
	}
	serverComponent struct {
		grpcServer  *grpc.Server
		grpcHandler igrpc.LockHandlerI
	}
	metricComponent struct {
		promReg     *prometheus.Registry
		promGRPC    *promgrpc.ServerMetrics
		httpSrv     *http.Server
		mux         *http.ServeMux
		httpHandler *promutil.Handler
		app         apputil.MetricI
	}
)

type App struct {
	cfg *Config
	log logr.Logger

	repository repositoryComponent
	server     serverComponent
	metric     metricComponent
	service    iservice.LockServiceI

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
	httpMtrHandler *promutil.Handler,
	app apputil.MetricI,
	env apputil.Env,
	ver apputil.Ver,
) *App {
	return &App{
		cfg:        cfg,
		log:        log,
		repository: repositoryComponent{redisConn: redisConn},
		server: serverComponent{
			grpcServer:  grpcSrv,
			grpcHandler: grpcHandler,
		},
		metric: metricComponent{
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
	a.metric.promGRPC.InitializeMetrics(a.server.grpcServer)

	a.metric.app.SetVersion(a.version)
	a.metric.app.SetEnvironment(a.environment)
}

func (a App) Run(ctx context.Context) error {
	a.prepare()
	rg := run.Group{}

	rg.Add(a.server.grpcHandler.Start,
		func(_ error) {
			a.server.grpcHandler.GracefulStop()
		})

	rg.Add(a.metric.httpHandler.Start,
		func(_ error) {
			if err := a.metric.httpHandler.GracefulStop(); err != nil {
				a.log.Error(err, "http metric server shutdown error")
			}
		})

	rg.Add(run.SignalHandler(ctx))

	return rg.Run()
}
