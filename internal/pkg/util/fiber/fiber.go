package fiberutil

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type HandlerI interface {
	Start() error
	GracefulStop() error
	Stop()

	Router() fiber.Router
}

type Handler struct {
	log logr.Logger
	cfg *Config

	srv *fiber.App
}

func NewHandlerFromConfig(log logr.Logger, c *Config) *Handler {
	fc := c.toFiberConfig()

	fc.JSONDecoder = json.Unmarshal
	fc.JSONEncoder = json.Marshal

	return &Handler{
		srv: fiber.New(fc),
		log: log,
		cfg: c,
	}
}

func (h Handler) Start() error {
	addr := fmt.Sprintf(":%s", h.cfg.Port)
	h.log.Info(fmt.Sprintf("http-fiber server is up on %s", addr))
	return h.srv.Listen(addr)
}

func (h Handler) GracefulStop() error {
	return h.srv.Shutdown()
}

func (h Handler) Stop() {
	ctx, cancel := context.WithTimeout(context.TODO(), 0)
	defer cancel()
	_ = h.srv.ShutdownWithContext(ctx)
}

func (h Handler) Router() fiber.Router {
	return h.srv
}
