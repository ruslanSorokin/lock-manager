package main

import (
	"flag"
	"os"

	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/ruslanSorokin/lock-manager/internal/config"
	"github.com/ruslanSorokin/lock-manager/internal/infra/handler/igrpc"
	"github.com/ruslanSorokin/lock-manager/internal/infra/provider/repository/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/service"
)

func run(c config.Type) error {
	zl := zerolog.New(os.Stdout)
	zl = zl.With().Timestamp().Logger()
	log := zerologr.New(&zl)

	cfg := config.MustLoad[Config](log, c)

	dbRedis, err := iredis.NewClientFromConfig(
		log, cfg.Repository.Redis,
	)
	if err != nil {
		return err
	}

	lp := iredis.NewLockStorage(log, dbRedis)

	svc := service.NewLockServiceFromConfig(
		log, lp, cfg.Service,
	)

	grpcSrv := grpc.NewServer()
	grpcHandler := igrpc.NewLockHandlerFromConfig(
		grpcSrv, log, svc, cfg.Handler.GRPC,
	)

	grpcHandler.Register()
	err = grpcHandler.Start()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	cfgFlag := flag.String("config", "dev", "startup config")
	flag.Parse()
	cfg := config.Type(*cfgFlag)
	if err := run(cfg); err != nil {
		panic(err)
	}
}
