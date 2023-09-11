package main

import (
	"flag"
	"os"
)

type args struct {
	Config string
}

func parseArgs() *args {
	var config string
	args := &args{}

	flag.StringVar(&config, "config", "default.testing.yaml", "config file")
	flag.Parse()
	args.Config = config

	config, isSet := os.LookupEnv("LOCK_MANAGER_CONFIG_FILE")
	if isSet {
		args.Config = config
	}

	return args
}
