package controller

import "github.com/gofiber/fiber/v2"

type ResponseType struct {
	Code   int   `json:"code"`
	Data   any   `json:"data"`
	Errors []any `json:"errors"`
}

func Response(c *fiber.Ctx, r ResponseType) error {
	return c.Status(r.Code).JSON(fiber.Map{
		"code":  r.Code,
		"error": r.Errors,
		"data":  r.Data,
	})
}
