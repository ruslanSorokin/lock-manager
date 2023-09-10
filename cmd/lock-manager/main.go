package main

import (
	"context"
	"errors"
	"os"
	"syscall"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/oklog/run"
	"github.com/rs/zerolog"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/app"
	apputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app"
	cfgutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/config"
)

const appName = "lock-manager"

func newLogger() logr.Logger {
	zl := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log := zerologr.New(&zl)
	return log
}

func start(appName apputil.Name, file cfgutil.File) error {
	log := newLogger()

	cfg := &app.Config{}
	cfgutil.MustLoad(cfg, appName, file)

	a, cleanup, err := app.Wire(log, cfg)
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
	args := parseArgs()

	a := apputil.Name(appName)
	c := cfgutil.File(args.Config)

	if err := start(a, c); err != nil {
		panic(err)
	}
}
