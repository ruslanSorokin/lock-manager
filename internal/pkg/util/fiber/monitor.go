package fiberutil

import (
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

const pattern = "/metrics/http-fiber"

func Register(mux *http.ServeMux) {
	fh := monitor.New(
		monitor.Config{Title: "http handler metrics", APIOnly: false},
	)
	h := adaptor.FiberHandler(fh)
	mux.Handle(pattern, h)
}
