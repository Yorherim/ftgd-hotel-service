package api

import "github.com/gofiber/fiber/v2"

type ResponseData struct {
	Code  int    `json:"code"`
	Data  any    `json:"data"`
	Error string `json:"error"`
}

func response(c *fiber.Ctx, response ResponseData) error {
	return c.Status(400).JSON(fiber.Map{
		"code":  response.Code,
		"error": response.Error,
		"data":  response.Data,
	})
}
