package main

import (
	"context"
	"errors"
	"flag"
	"os"
	"syscall"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/oklog/run"
	"github.com/rs/zerolog"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/app"
	apputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app"
	configutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/config"
)

const appName = "lock-manager"

func newLogger() logr.Logger {
	zl := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log := zerologr.New(&zl)
	return log
}

func start(n apputil.Name, e apputil.Env, c configutil.File) error {
	log := newLogger()

	appEnvelope := configutil.NewConfig[Config](n)
	appEnvelope.MustLoad(c)
	cfg := appEnvelope.AppConfig()

	log.Info("application environment", "env", e)
	log.Info("application version", "version", cfg.App.Version)

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

func parseArgs() *arguments {
	var env string
	var configFile string

	flag.StringVar(&env, "env", "dev", "app environment")
	flag.StringVar(&configFile, "config", "default.development.yaml", "app config file")

	flag.Parse()
	return &arguments{
		Env:        env,
		ConfigFile: configFile,
	}
}

func main() {
	args := parseArgs()

	n := apputil.Name(appName)
	e := apputil.MustParseEnv(args.Env)
	c := configutil.File(args.ConfigFile)
	if err := start(n, e, c); err != nil {
		panic(err)
	}
}
