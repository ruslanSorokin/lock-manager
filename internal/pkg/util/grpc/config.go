package grpcutil

import "time"

type (
	ping struct {
		After   time.Duration `yaml:"after"   env:"AFTER"   env-default:"10m"`
		Timeout time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"15s"`
	}
	conn struct {
		MaxIdle time.Duration `yaml:"maxIdle" env-default:"5m"`
		MaxAge  time.Duration `yaml:"maxAge"  env-default:"5m"`
	}
	Config struct {
		Port           string `yaml:"port"           env:"PORT"            env-required:"true"`
		WithReflection bool   `yaml:"withReflection" env:"WITH_REFLECTION"                     env-default:"false"`

		Ping ping `yaml:"ping" env-prefix:"PING_"`

		Conn conn `yaml:"conn" env-prefix:"CONN_"`
	}
)
