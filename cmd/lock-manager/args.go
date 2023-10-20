package main

import (
	"flag"
	"os"
)

type args struct {
	Config string
}

func parseArgs() *args {
	var cfg string
	var a args

	flag.StringVar(&cfg, "config", "default.testing.yaml", "config file")
	flag.Parse()
	a.Config = cfg

	if c, ok := os.LookupEnv("LOCK_MANAGER_CONFIG_FILE"); ok {
		a.Config = c
	}

	return &a
}
