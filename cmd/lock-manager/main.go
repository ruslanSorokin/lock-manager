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
	"github.com/ruslanSorokin/lock-manager/internal/metric/iprom"
	"github.com/ruslanSorokin/lock-manager/internal/service"
	"github.com/ruslanSorokin/lock-manager/pkg/grpcutil"
	"github.com/ruslanSorokin/lock-manager/pkg/promutil"
)

func start(c config.Type) error {
	zl := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log := zerologr.New(&zl)

	cfg := config.MustLoad[Config](log, c)

	dbRedis, err := iredis.NewClientFromConfig(
		log, cfg.Repository.Redis)
	if err != nil {
		return err
	}
	defer iredis.Close(dbRedis)

	lockRepo := iredis.NewLockStorage(log, dbRedis)

	promReg := prometheus.NewRegistry()

	mtrHandler := promutil.NewFromConfig(
		log,
		promReg,
		http.NewServeMux(),
		cfg.Observability.Pull.Metric)

	svc := service.NewLockServiceFromConfig(
		log, lockRepo, iprom.New(promReg), cfg.Service)

	grpcProcessingTimeHistogram := grpcutil.NewProcessingTimeHistogram(promReg)

	grpcSrv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(grpcProcessingTimeHistogram.UnaryServerInterceptor()),
		grpc.ChainStreamInterceptor(grpcProcessingTimeHistogram.StreamServerInterceptor()))

	grpcProcessingTimeHistogram.InitializeMetrics(grpcSrv)

	grpcHandler := igrpc.NewFromConfig(
		grpcSrv, log, svc, cfg.Handler.GRPC)

	g := &run.Group{}

	g.Add(grpcHandler.Start,
		func(err error) { grpcHandler.GracefulStop() })

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
