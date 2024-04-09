package main

import (
	"flag"
	"fmt"
	"github.com/Yorherim/ftgd-hotel-service/internal/composition"
	"github.com/Yorherim/ftgd-hotel-service/internal/config"
	mongodb2 "github.com/Yorherim/ftgd-hotel-service/pkg/client/mongodb"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(config.FiberConfig)

	client := mongodb2.InitClient()
	validate := validator.New()

	apiV1 := app.Group("api/v1")

	userComposition := composition.NewUserComposition(client, validate)
	userComposition.Handler.Register(apiV1)

	listenAddr := flag.String(
		"listenAddr",
		":5000",
		"The listen address of the API server")
	flag.Parse()

	if err := app.Listen(*listenAddr); err != nil {
		_ = fmt.Errorf("err listen server: %s", err)
	}
}
