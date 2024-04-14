package user

import (
	"context"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller"
	"github.com/Yorherim/ftgd-hotel-service/pkg/utils"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

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
	logger      *zap.SugaredLogger
	utils       *utils.Utils
}

func NewUserHandler(
	userService UserService,
	validate *validator.Validate,
	logger *zap.SugaredLogger,
	utils *utils.Utils) controller.Handler {

	return &UserHandler{
		userService: userService,
		validate:    validate,
		logger:      logger,
		utils:       utils,
	}
}
