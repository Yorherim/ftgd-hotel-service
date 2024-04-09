package composition

import (
	user2 "github.com/Yorherim/ftgd-hotel-service/internal/adapters/db/mongodb/user"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller/http/v1/user"
	v1 "github.com/Yorherim/ftgd-hotel-service/internal/controller/http/v1/user"
	"github.com/Yorherim/ftgd-hotel-service/internal/domain/service"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserComposition struct {
	Store   service.UserStore
	Service user.UserService
	Handler controller.Handler
}

func NewUserComposition(client *mongo.Client, validate *validator.Validate) *UserComposition {
	userStorage := user2.NewMongoUserStore(client)
	userService := service.NewUserService(userStorage)
	userHandler := v1.NewUserHandler(userService, validate)

	return &UserComposition{
		Store:   userStorage,
		Service: userService,
		Handler: userHandler,
	}
}
