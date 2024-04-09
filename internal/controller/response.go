package controller

import "github.com/gofiber/fiber/v2"

type ResponseType struct {
	Code  int    `json:"code"`
	Data  any    `json:"data"`
	Error string `json:"error"`
}

func Response(c *fiber.Ctx, r ResponseType) error {
	return c.Status(400).JSON(fiber.Map{
		"code":  r.Code,
		"error": r.Error,
		"data":  r.Data,
	})
}
