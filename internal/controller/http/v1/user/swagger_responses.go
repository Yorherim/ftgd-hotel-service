package user

import (
	"github.com/Yorherim/ftgd-hotel-service/internal/domain/entity"
)

type getUserByID200Example struct {
	Code   int         `json:"code" example:"200"`
	Data   entity.User `json:"data"`
	Errors []string    `json:"errors"`
}

type getUserByID400Example struct {
	Code   int      `json:"code" example:"400"`
	Data   string   `json:"data" example:"null" extensions:"x-nullable"`
	Errors []string `json:"errors" example:"invalid id"`
}

type getUserByID404Example struct {
	Code   int      `json:"code" example:"404"`
	Data   string   `json:"data" example:"null" extensions:"x-nullable"`
	Errors []string `json:"errors" example:"user not found"`
}

type getUsers200Example struct {
	Code   int           `json:"code" example:"200"`
	Data   []entity.User `json:"data"`
	Errors []string      `json:"errors"`
}

type createUser400Example struct {
	Code   int      `json:"code" example:"400"`
	Data   string   `json:"data" example:"null" extensions:"x-nullable"`
	Errors []string `json:"errors" example:"error parse body"`
}

type serverError500Example struct {
	Code   int      `json:"code" example:"500"`
	Data   string   `json:"data" example:"null" extensions:"x-nullable"`
	Errors []string `json:"errors" example:"server error"`
}
