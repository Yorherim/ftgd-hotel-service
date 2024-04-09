package user

import (
	"context"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller"

	"github.com/Yorherim/ftgd-hotel-service/internal/domain/entity"
)

type UserService interface {
	GetUserByID(context.Context, string) (*entity.User, error)
	GetUsers(context.Context) ([]*entity.User, error)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) controller.Handler {
	return &UserHandler{
		userService: userService,
	}
}
