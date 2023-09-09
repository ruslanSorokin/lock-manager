package main

import (
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/app"
	redisconn "github.com/ruslanSorokin/lock-manager/internal/pkg/conn/redis"
	apputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app"
	grpcutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/grpc"
	promutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/prom"
)

type Config struct {
	Repository struct {
		Redis *redisconn.Config `yaml:"redis"`
	} `yaml:"repository"`
	Handler struct {
		GRPC *grpcutil.Config `yaml:"grpc"`
	} `yaml:"handler"`
	Observability struct {
		Pull struct {
			Metric *promutil.Config `yaml:"metric"`
		} `yaml:"pull"`
	} `yaml:"observability"`
	App struct {
		Environment string `yaml:"environment"`
		Version     string `yaml:"version"`
	} `yaml:"application"`
}

func (c Config) ToAppConfig() *app.Config {
	return &app.Config{
		Redis:      c.Repository.Redis,
		GRPC:       c.Handler.GRPC,
		HTTPMetric: c.Observability.Pull.Metric,
		Ver:        apputil.Ver(c.App.Version),
	}
}
