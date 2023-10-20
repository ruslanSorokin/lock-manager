package httputil

import "time"

type Config struct {
	Port              string        `yaml:"port"              env:"PORT" env-required:"true"`
	ReadTimeOut       time.Duration `yaml:"readTimeout"                                      env-default:"5s"`
	ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"                                env-default:"5s"`
	IdleTimeout       time.Duration `yaml:"idleTimeout"                                      env-default:"5s"`
	WriteTimeout      time.Duration `yaml:"writeTimeout"                                     env-default:"5s"`
}
