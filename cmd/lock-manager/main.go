package main

import (
	"context"
	"flag"
	"os"
	"syscall"

	"github.com/go-logr/zerologr"
	"github.com/oklog/run"
	"github.com/rs/zerolog"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/app"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/apputil"
)

func start(e apputil.Env) error {
	zl := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log := zerologr.New(&zl)

	log.Info("environment", "env", e)

	cfg := apputil.MustLoad[Config](e)

	a, cleanup, err := app.Wire(e, log, cfg.ToAppConfig())
	if err != nil {
		return err
	}
	defer cleanup()

	a.Configure()

	rg := &run.Group{}

	rg.Add(a.RunGRPC,
		func(_ error) { a.GracefulStopGRPC() })

	rg.Add(a.RunMetric,
		func(_ error) { a.GracefulStopMetric() })

	rg.Add(run.SignalHandler(context.TODO(),
		syscall.SIGINT, syscall.SIGTERM))

	return rg.Run()
}

func main() {
	var envFlag string
	flag.StringVar(&envFlag, "env", "dev", "app environment")
	flag.Parse()

	env := apputil.MustParseEnv(envFlag)
	if err := start(env); err != nil {
		panic(err)
	}
}
