package main

import (
	"os"

	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"

	"github.com/ruslanSorokin/lock-manager/internal/config"
	"github.com/ruslanSorokin/lock-manager/internal/infra/provider/repository/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/service"
)

func run() error {
	zl := zerolog.New(os.Stdout)
	zl = zl.With().Timestamp().Logger()
	log := zerologr.New(&zl)

	cfg := config.MustLoad[Config](log, config.Local)

	dbRedis, err := iredis.NewClientFromConfig(
		log, cfg.repository.redis,
	)
	if err != nil {
		return err
	}

	lp := iredis.NewLockStorage(log, dbRedis)

	_ = service.NewLockService(
		log, lp,
	)

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
