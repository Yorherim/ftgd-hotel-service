package user

import (
	"github.com/Yorherim/ftgd-hotel-service/internal/domain/entity"
)

type getUserByID200Example struct {
	Code  int         `json:"code" example:"200"`
	Data  entity.User `json:"data"`
	Error string      `json:"error" example:""`
}

type getUserByID400Example struct {
	Code  int    `json:"code" example:"400"`
	Data  string `json:"data" example:"null" extensions:"x-nullable"`
	Error string `json:"error" example:"invalid id"`
}

type getUserByID404Example struct {
	Code  int    `json:"code" example:"404"`
	Data  string `json:"data" example:"null" extensions:"x-nullable"`
	Error string `json:"error" example:"user not found"`
}
