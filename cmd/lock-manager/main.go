package main

import (
	"os"

	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/infra/repository/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/config"
)

const (
	appName = "lock-manager"
)

func run() error {
	zl := zerolog.New(os.Stdout)
	zl = zl.With().Timestamp().Logger()
	log := zerologr.New(&zl)

	cfg := config.MustLoad[Config](log, appName, config.Local)

	_, err := iredis.NewClientFromConfig(
		log, cfg.redis,
	)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
