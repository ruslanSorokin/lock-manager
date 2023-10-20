package fiberutil

import (
	"net/http"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Middleware func(*fiber.Ctx) error

func WireProvideDefaultMiddleware() []Middleware {
	return []Middleware{
		recover.New(recover.Config{EnableStackTrace: true}),
	}
}

func WireProvideDefaultApp(
	c *Config,
	mw []Middleware,
	mtrMux *http.ServeMux,
) *fiber.App {
	fc := c.toFiberConfig()

	fc.JSONDecoder = json.Unmarshal
	fc.JSONEncoder = json.Marshal

	a := fiber.New(fc)

	Register(mtrMux)

	for _, m := range mw {
		a.Use((func(*fiber.Ctx) error)(m))
	}
	return a
}
