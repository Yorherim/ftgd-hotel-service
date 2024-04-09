package composition

import (
	"github.com/Yorherim/ftgd-hotel-service/internal/adapters/db/mongodb"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller/http/v1/user"
	v1 "github.com/Yorherim/ftgd-hotel-service/internal/controller/http/v1/user"
	"github.com/Yorherim/ftgd-hotel-service/internal/domain/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserComposition struct {
	Store   service.UserStore
	Service user.UserService
	Handler controller.Handler
}

func NewUserComposition(client *mongo.Client) *UserComposition {
	userStorage := mongodb.NewMongoUserStore(client)
	userService := service.NewUserService(userStorage)
	userHandler := v1.NewUserHandler(userService)

	return &UserComposition{
		Store:   userStorage,
		Service: userService,
		Handler: userHandler,
	}
}
