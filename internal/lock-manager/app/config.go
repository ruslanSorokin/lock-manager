package app

import (
	redisconn "github.com/ruslanSorokin/lock-manager/internal/pkg/conn/redis"
	apputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app"
	grpcutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/grpc"
	httputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/http"
)

type (
	storage struct {
		Redis redisconn.Config `yaml:"redis" env-prefix:"REDIS_"`
	}
	handler struct {
		GRPC grpcutil.Config `yaml:"grpc" env-prefix:"GRPC_"`
	}
	metric struct {
		Pull httputil.Config `yaml:"pull" env-prefix:"PULL_"`
	}

	app struct {
		Environment string `yaml:"environment" env:"ENV"`
		Version     string `yaml:"version"     env:"VERSION"`
	}

	Config struct {
		Storage storage `yaml:"storage"     env-prefix:"STORAGE_"`
		Handler handler `yaml:"handler"     env-prefix:"HANDLER_"`
		Metric  metric  `yaml:"metric"      env-prefix:"METRIC_"`
		App     app     `yaml:"application" env-prefix:"APP_"`
	}
)

func toWireConfig(c *Config) (*wireConfig, error) {
	env, err := apputil.ParseEnv(c.App.Environment)
	if err != nil {
		return nil, err
	}
	return &wireConfig{
		Redis:       &c.Storage.Redis,
		GRPC:        &c.Handler.GRPC,
		Pull:        &c.Metric.Pull,
		Version:     apputil.Ver(c.App.Version),
		Environment: env,
	}, nil
}

type wireConfig struct {
	Redis       *redisconn.Config
	GRPC        *grpcutil.Config
	Pull        *httputil.Config
	Version     apputil.Ver
	Environment apputil.Env
}
