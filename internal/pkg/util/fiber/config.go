package fiberutil

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Port        string `yaml:"port"        env:"PORT"        env-required:"true"`
	Prefork     bool   `yaml:"prefork"     env:"PREFORK"                         env-default:"false"`
	Concurrency int    `yaml:"concurrency" env:"CONCURRENCY"                     env-default:"262144"`

	DisableKeepAlive bool `yaml:"disableKeepAlive" env:"DISABLE_KEEP_ALIVE" env-default:"false"`

	ReadTimeout  time.Duration `yaml:"readTimeout"  env:"READ_TIMEOUT"  env-default:"-1"`
	WriteTimeout time.Duration `yaml:"writeTimeout" env:"WRITE_TIMEOUT" env-default:"-1"`
	IdleTimeout  time.Duration `yaml:"idleTimeout"  env:"IDLE_TIMEOUT"  env-default:"-1"`
}

func (c *Config) toFiberConfig() fiber.Config {
	return fiber.Config{
		Prefork:          c.Prefork,
		Concurrency:      c.Concurrency,
		DisableKeepalive: c.DisableKeepAlive,
		ReadTimeout:      c.ReadTimeout,
		WriteTimeout:     c.WriteTimeout,
		IdleTimeout:      c.IdleTimeout,
	}
}
