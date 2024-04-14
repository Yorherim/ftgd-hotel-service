package main

import (
	"flag"
	_ "github.com/Yorherim/ftgd-hotel-service/docs"
	"github.com/Yorherim/ftgd-hotel-service/internal/composition"
	"github.com/Yorherim/ftgd-hotel-service/internal/config"
	mongodb2 "github.com/Yorherim/ftgd-hotel-service/pkg/client/mongodb"
	logger2 "github.com/Yorherim/ftgd-hotel-service/pkg/logger"
	utils2 "github.com/Yorherim/ftgd-hotel-service/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

//	@title						hotel reservation
//	@version					1.0
//	@description				hotel service
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				Igor
//	@contact.email				Yorherim@gmail.com
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@host						localhost:5000
//	@BasePath					/api/v1
//	@securityDefinitions.basic	BasicAuth
func main() {
	app := fiber.New(config.FiberConfig)

	client := mongodb2.InitClient()
	validate := validator.New()
	logger := logger2.NewZapSugarLogger()
	utils := utils2.NewUtils()

	app.Get("/swagger/*", swagger.New(config.SwaggerConfig))
	apiV1 := app.Group("api/v1")

	composition.NewUserComposition(client, validate, logger, &apiV1, utils)

	listenAddr := flag.String(
		"listenAddr",
		":5000",
		"The listen address of the API server")
	flag.Parse()

	if err := app.Listen(*listenAddr); err != nil {
		logger.Errorf("err listen server: %s", err)
	}
}
