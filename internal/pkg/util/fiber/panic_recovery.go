package fiberutil

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewPanicRecoveryHandler(cfg recover.Config) func(*fiber.Ctx) error {
	return recover.New(cfg)
}
