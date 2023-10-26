package unlock

import (
	"errors"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/ifiber/shared"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/ilog"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/ierror"
)

type (
	Handler func(*fiber.Ctx) error
)

func New(log logr.Logger, svc service.LockServiceI) Handler {
	const internalErrLogMsg = "internal error during attempt to unlock resource"
	const badAttemptLogMsg = "bad attempt to unlock resource"

	return func(
		c *fiber.Ctx,
	) error {
		rID := c.Params(shared.PathParamNameResourceID)
		rID = utils.CopyString(rID)
		tkn := c.Query(shared.QueryParamNameToken)
		tkn = utils.CopyString(tkn)

		err := svc.Unlock(c.Context(), rID, tkn)
		if err == nil {
			return c.SendStatus(http.StatusOK)
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
