package config

import (
	"errors"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller"
	"github.com/gofiber/fiber/v2"
)

var FiberConfig = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		if err != nil {
			return ctx.Status(code).JSON(controller.ResponseType{
				Error: e.Message,
				Data:  nil,
				Code:  e.Code,
			})
		}

		return nil
	},
}
