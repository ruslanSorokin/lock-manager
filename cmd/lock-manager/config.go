package main

import (
	"github.com/ruslanSorokin/lock-manager/internal/infra/handler/igrpc"
	"github.com/ruslanSorokin/lock-manager/internal/infra/provider/repository/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/service"
	"github.com/ruslanSorokin/lock-manager/pkg/promutil"
)

type Config struct {
	Repository struct {
		Redis iredis.Config `yaml:"redis"`
	} `yaml:"repository"`
	Handler struct {
		GRPC igrpc.Config `yaml:"grpc"`
	} `yaml:"handler"`
	Service       service.Config `yaml:"service"`
	Observability struct {
		Pull struct {
			Metric promutil.Config `yaml:"metric"`
		} `yaml:"pull"`
	} `yaml:"observability"`
	App struct {
		Environment string `yaml:"environment"`
		Version     string `yaml:"version"`
	} `yaml:"application"`
}
