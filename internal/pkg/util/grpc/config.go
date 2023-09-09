package grpcutil

import "time"

type Config struct {
	Port           string `yaml:"port"`
	WithReflection bool   `yaml:"withReflection"`
	Ping           struct {
		After   time.Duration `yaml:"after"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"keepAlive"`
	Conn struct {
		MaxIdle time.Duration `yaml:"maxIdle"`
		MaxAge  time.Duration `yaml:"maxAge"`
	} `yaml:"conn"`
}
