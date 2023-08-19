package promutil

import "time"

type Config struct {
	Port        string        `yaml:"port"`
	ReadTimeOut time.Duration `yaml:"readTimeOut"`
}
