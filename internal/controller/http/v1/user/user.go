package user

import (
	"context"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller"
	"github.com/go-playground/validator/v10"

	"github.com/Yorherim/ftgd-hotel-service/internal/domain/entity"
)

type UserService interface {
	GetUserByID(context.Context, string) (*entity.User, error)
	GetUsers(context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, dto CreateUserDTO) (*entity.User, error)
}

type UserHandler struct {
	userService UserService
	validate    *validator.Validate
}

func NewUserHandler(userService UserService, validate *validator.Validate) controller.Handler {
	return &UserHandler{
		userService: userService,
		validate:    validate,
	}
}
