package app

import (
	"net/http"

	"github.com/go-logr/logr"
	promgrpc "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc"
	iservice "github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/apputil"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/promutil"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/redisconn"
	"google.golang.org/grpc"
)

type (
	repository struct {
		redisConn *redisconn.Conn
	}
	server struct {
		grpcServer  *grpc.Server
		grpcHandler igrpc.LockHandlerI
	}
	metric struct {
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

	repository repository
	server     server
	metric     metric
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
		repository: repository{redisConn: redisConn},
		server: server{
			grpcServer:  grpcSrv,
			grpcHandler: grpcHandler,
		},
		metric: metric{
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

func (a App) Configure() {
	a.metric.promGRPC.InitializeMetrics(a.server.grpcServer)

	a.metric.app.SetVersion(a.version)
	a.metric.app.SetEnvironment(a.environment)
}

func (a App) RunGRPC() error {
	return a.server.grpcHandler.Start()
}

func (a App) GracefulStopGRPC() {
	a.server.grpcHandler.GracefulStop()
}

func (a App) RunMetric() error {
	return a.metric.httpHandler.Start()
}

func (a App) GracefulStopMetric() {
	if err := a.metric.httpHandler.GracefulStop(); err != nil {
		a.log.Error(err, "grpc shutdown error")
	}
}
