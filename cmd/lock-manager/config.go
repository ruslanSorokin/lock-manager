package main

import (
	"github.com/ruslanSorokin/lock-manager/internal/infra/provider/repository/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/service"
)

type Config struct {
	repository struct{ redis iredis.Config }
	service    service.Config
}
