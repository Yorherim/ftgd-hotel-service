package service

import (
	"context"
	"github.com/Yorherim/ftgd-hotel-service/internal/domain/entity"
)

type UserStore interface {
	GetUserByID(ctx context.Context, Id string) (*entity.User, error)
	GetUsers(ctx context.Context) ([]*entity.User, error)
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
