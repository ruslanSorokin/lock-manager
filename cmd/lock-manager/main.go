package main

import (
	"context"
	"errors"
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

	rg := &run.Group{}

	ctx, cancel := context.WithCancel(context.TODO())

	rg.Add(func() error { return a.Run(ctx) },
		func(err error) { cancel() })

	rg.Add(run.SignalHandler(context.TODO(),
		syscall.SIGINT, syscall.SIGTERM))

	if err := rg.Run(); !errors.As(err, &run.SignalError{}) {
		return err
	}
	return nil
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
