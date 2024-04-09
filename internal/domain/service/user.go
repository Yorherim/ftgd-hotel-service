package service

import (
	"context"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller/http/v1/user"
	"github.com/Yorherim/ftgd-hotel-service/internal/domain/entity"
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
	store UserStore
}

func NewUserService(storage UserStore) *Service {
	return &Service{store: storage}
}

func (s *Service) GetUserByID(ctx context.Context, Id string) (*entity.User, error) {
	return s.store.GetUserByID(ctx, Id)

}

func (s *Service) GetUsers(ctx context.Context) ([]*entity.User, error) {
	return s.store.GetUsers(ctx)
}

func (s *Service) CreateUser(ctx context.Context, dto user.CreateUserDTO) (*entity.User, error) {
	encPw, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	dto.Password = string(encPw)

	return s.store.CreateUser(ctx, dto)
}
