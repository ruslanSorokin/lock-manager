package promutil

import "time"

type Config struct {
	Port        string        `yaml:"port"        env:"PORT" env-required:"true"`
	ReadTimeOut time.Duration `yaml:"readTimeout"                                env-default:"5s"`
}
