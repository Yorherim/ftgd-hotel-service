package api

import (
	"context"
	"github.com/Yorherim/ftgd-hotel-service/db"
	"github.com/Yorherim/ftgd-hotel-service/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)

	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		err = fiber.NewError(fiber.StatusBadRequest, "invalid id")
		return err
	}

	return response(c, ResponseData{
		Code: 200,
		Data: user,
	})
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	return response(c, ResponseData{
		Code: 200,
		Data: []types.User{
			{FirstName: "Vasya", LastName: "222"},
			{FirstName: "James", LastName: "Bond"},
		},
	})
}
