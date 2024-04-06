package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/Yorherim/ftgd-hotel-service/api/handlers"
	"github.com/Yorherim/ftgd-hotel-service/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	dbUri = "mongodb://localhost:27017"
)

func main() {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if err != nil {
				return ctx.Status(code).JSON(api.ResponseData{
					Error: e.Message,
					Data:  nil,
					Code:  e.Code,
				})
			}

			return nil
		},
	})

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	apiV1 := app.Group("api/v1")
	{
		apiV1.Get("/users", userHandler.HandleGetUsers)
		apiV1.Get("/users/:id", userHandler.HandleGetUserByID)
	}

	listenAddr := flag.String(
		"listenAddr",
		":5000",
		"The listen address of the API server")
	flag.Parse()

	if err := app.Listen(*listenAddr); err != nil {
		_ = fmt.Errorf("err listen server: %s", err)
	}
}
