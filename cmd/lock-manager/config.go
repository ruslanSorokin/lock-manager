package main

import "github.com/ruslanSorokin/lock-manager/internal/lock-manager/infra/repository/iredis"

type Config struct {
	redis iredis.Config
}
