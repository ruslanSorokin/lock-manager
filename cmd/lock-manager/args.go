package main

import "flag"

type args struct {
	Config string
}

func parseArgs() *args {
	var config string

	flag.StringVar(&config, "config", "default.development.yaml", "config file")

	flag.Parse()
	return &args{
		Config: config,
	}
}
