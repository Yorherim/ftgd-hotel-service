package controller

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Register(appGroup fiber.Router)
}
