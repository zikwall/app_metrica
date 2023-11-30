package fiberext

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/zikwall/app_metrica/pkg/log"
	"github.com/zikwall/app_metrica/pkg/xerror"
)

const defaultErrorMessage = "Internal Server Error"

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if err != nil {
		code := fiber.StatusInternalServerError
		value := defaultErrorMessage

		var e *xerror.HTTPError
		var w *xerror.WrapError
		var f *fiber.Error

		if errors.As(err, &f) {
			code = f.Code
			value = f.Message
		} else if errors.As(err, &e) {
			code = e.Code()
			value = e.Message()
		}

		if errors.As(err, &w) {
			var pub *xerror.PublicError
			var pri *xerror.PrivateError
			if errors.As(err, &pub) {
				value = fmt.Sprintf("%s: %v", w.Context, pub.Error())
			} else if errors.As(err, &pri) {
				log.Warningf("error handler: %s", err)
			}
			log.ViaNotify(ctx.Context(), w, map[string]interface{}{
				"_context": w.Context,
			})
		}

		return ctx.Status(code).JSON(fiber.Map{
			"status":  code,
			"message": value,
		})
	}

	return nil
}
