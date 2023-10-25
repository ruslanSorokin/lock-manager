package lock

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/ifiber/shared"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/ierror"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/ilog"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
)

type (
	Handler func(*fiber.Ctx) error
)

func New(
	log logr.Logger,
	svc service.LockServiceI,
) Handler {
	const internalErrLogMsg = "internal error during attempt to lock resource"
	const badAttemptLogMsg = "bad attempt to lock resource"

	return func(
		c *fiber.Ctx,
	) error {
		rID := c.Params(shared.PathParamNameResourceID)
		rID = utils.CopyString(rID)
		tkn, err := svc.Lock(c.Context(), rID)

		if err == nil {
			return c.Status(http.StatusCreated).JSON(fiber.Map{
				"token": tkn,
				"_links": fiber.Map{
					"unlock": fiber.Map{
						"href": fmt.Sprintf("%s%s?token=%s",
							c.BaseURL(), c.OriginalURL(), tkn),
						"type": "delete",
					},
				},
			})
		}

		var t interface {
			error
			ierror.HTTPConvertible
			ierror.EnumConvertible
		}
		logMsg := internalErrLogMsg
		code := http.StatusInternalServerError
		apiStCode := shared.APIStCodeInternalError

		if errors.As(err, &t) {
			logMsg = badAttemptLogMsg
			code = t.ToHTTP()
			apiStCode = t.ToEnum()
		}

		log.Error(err, logMsg,
			ilog.TagResourceID, rID,
			ilog.TagToken, tkn,
			ilog.TagHTTPStCode, code)

		return c.Status(code).JSON(fiber.Map{
			"error": apiStCode,
		})
	}
}
