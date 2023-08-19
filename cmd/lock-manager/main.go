package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"syscall"

	"github.com/go-logr/zerologr"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/ruslanSorokin/lock-manager/internal/config"
	"github.com/ruslanSorokin/lock-manager/internal/infra/handler/igrpc"
	"github.com/ruslanSorokin/lock-manager/internal/infra/provider/repository/iredis"
	ipromsvc "github.com/ruslanSorokin/lock-manager/internal/metric/iprom"
	"github.com/ruslanSorokin/lock-manager/internal/service"
	"github.com/ruslanSorokin/lock-manager/pkg/grpcutil"
	"github.com/ruslanSorokin/lock-manager/pkg/promutil"
	"github.com/ruslanSorokin/lock-manager/pkg/redisutil"
)

func start(c config.Type) error {
	zl := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log := zerologr.New(&zl)

	cfg := config.MustLoad[Config](log, c)

	dbRedis, err := redisutil.NewClientFromConfig(
		log, cfg.Repository.Redis)
	if err != nil {
		return err
	}
	defer redisutil.Close(dbRedis)

	lockRepo := iredis.NewLockStorage(log, dbRedis)

	promReg := prometheus.NewRegistry()

	mtrHandler := promutil.NewFromConfig(
		log,
		promReg,
		http.NewServeMux(),
		cfg.Observability.Pull.Metric)

	svc := service.NewLockServiceFromConfig(
		log, lockRepo, ipromsvc.New(promReg), cfg.Service)

	grpcProcessingTimeHistogram := grpcutil.NewProcessingTimeHistogram(promReg)

	grpcSrv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(grpcProcessingTimeHistogram.UnaryServerInterceptor()),
		grpc.ChainStreamInterceptor(grpcProcessingTimeHistogram.StreamServerInterceptor()))

	grpcProcessingTimeHistogram.InitializeMetrics(grpcSrv)

	grpcHandler := grpcutil.NewHandlerFromConfig(grpcSrv, log, cfg.Handler.GRPC)

	lockGRPCHandler := igrpc.New(grpcHandler, log, svc)

	g := &run.Group{}

	g.Add(lockGRPCHandler.Start,
		func(err error) { lockGRPCHandler.GracefulStop() })

	g.Add(mtrHandler.Start,
		func(err error) { panic(mtrHandler.GracefulStop()) })

	g.Add(run.SignalHandler(context.TODO(),
		syscall.SIGINT, syscall.SIGTERM))

	return g.Run()
}

func main() {
	cfgFlag := flag.String("config", "dev", "startup config")
	flag.Parse()
	cfg := config.Type(*cfgFlag)
	if err := start(cfg); err != nil {
		panic(err)
	}
}
