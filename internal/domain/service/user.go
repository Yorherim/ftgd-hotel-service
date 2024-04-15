package service

import (
	"context"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller/http/v1/user"
	"github.com/Yorherim/ftgd-hotel-service/internal/domain/entity"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
)

type UserStore interface {
	GetUserByID(ctx context.Context, Id string) (*entity.User, error)
	GetUsers(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, dto user.CreateUserDTO) (*entity.User, error)
}

type Service struct {
	store  UserStore
	logger *zap.SugaredLogger
}

func NewUserService(storage UserStore, logger *zap.SugaredLogger) *Service {
	return &Service{
		store:  storage,
		logger: logger,
	}
}

func (s *Service) GetUserByID(ctx context.Context, Id string) (*entity.User, error) {
	return s.store.GetUserByID(ctx, Id)

}

func (s *Service) GetUsers(ctx context.Context) ([]*entity.User, error) {
	return s.store.GetUsers(ctx)
}

func (s *Service) CreateUser(ctx context.Context, dto user.CreateUserDTO) (*entity.User, error) {
	s.logger.Info("service CreateUser start")

	encPw, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcryptCost)
	if err != nil {
		s.logger.Errorf("service CreateUser error GenerateFromPassword: %s", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "server error")
	}

	dto.Password = string(encPw)

	return s.store.CreateUser(ctx, dto)
}
