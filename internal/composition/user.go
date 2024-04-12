package composition

import (
	user2 "github.com/Yorherim/ftgd-hotel-service/internal/adapters/db/mongodb/user"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller/http/v1/user"
	v1 "github.com/Yorherim/ftgd-hotel-service/internal/controller/http/v1/user"
	"github.com/Yorherim/ftgd-hotel-service/internal/domain/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type UserComposition struct {
	Store   service.UserStore
	Service user.UserService
	Handler controller.Handler
}

func NewUserComposition(
	client *mongo.Client,
	validate *validator.Validate,
	logger *zap.SugaredLogger,
	apiV1 *fiber.Router) *UserComposition {

	logger.Info("user composition init")
	logger.Warn("user composition warn")
	logger.Errorf("user composition err")

	userStorage := user2.NewMongoUserStore(client, logger)
	userService := service.NewUserService(userStorage, logger)
	userHandler := v1.NewUserHandler(userService, validate, logger)

	userComposition := UserComposition{
		Store:   userStorage,
		Service: userService,
		Handler: userHandler,
	}

	userComposition.Handler.Register(*apiV1)

	return &userComposition
}
