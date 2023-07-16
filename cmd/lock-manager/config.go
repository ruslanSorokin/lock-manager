package main

import (
	"github.com/ruslanSorokin/lock-manager/internal/infra/provider/repository/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/service"
)

type Config struct {
	Repository struct {
		Redis iredis.Config `yaml:"redis"`
	} `yaml:"repository"`
	Service service.Config `yaml:"service"`
}
