package ifiber

import (
	"fmt"

	"github.com/go-logr/logr"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/ifiber/lock"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/ifiber/shared"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/ifiber/unlock"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	fiberutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/fiber"
)

type LockHandler struct {
	fiberutil.HandlerI

	log logr.Logger
	svc service.LockServiceI

	lock   lock.Handler
	unlock unlock.Handler
}

var _ fiberutil.HandlerI = (*LockHandler)(nil)

func NewLockHandler(
	log logr.Logger,
	svc service.LockServiceI,
	handler fiberutil.HandlerI,
) *LockHandler {
	h := &LockHandler{
		HandlerI: handler,
		log:      log,
		svc:      svc,
		lock:     lock.New(log, svc),
		unlock:   unlock.New(log, svc),
	}

	h.registerRoutes()

	return h
}

func (h LockHandler) registerRoutes() {
	r := h.HandlerI.Router()

	rg := r.Group("/api/v1/locks")

	rIDParam := shared.PathParamNameResourceID
	rg.Post(fmt.Sprintf("/:%s", rIDParam), h.lock)
	rg.Delete(fmt.Sprintf("/:%s", rIDParam), h.unlock)
}
