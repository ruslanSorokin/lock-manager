package main

import "github.com/ruslanSorokin/lock-manager/internal/lock-manager/repository/iredis"

type Config struct {
	redis iredis.Config
}
