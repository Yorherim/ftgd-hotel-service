package user

import (
	"github.com/Yorherim/ftgd-hotel-service/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) Register(appGroup fiber.Router) {
	users := appGroup.Group("users")
	{
		users.Get("/", h.HandleGetUsers)
		users.Get("/:id", h.HandleGetUserByID)
	}
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)

	user, err := h.userService.GetUserByID(c.Context(), id)
	if err != nil {
		err = fiber.NewError(fiber.StatusBadRequest, "invalid id")
		return err
	}

	return controller.Response(c, controller.ResponseType{
		Code: 200,
		Data: user,
	})
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetUsers(c.Context())
	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "error get users")
		return err
	}

	return controller.Response(c, controller.ResponseType{
		Code: 200,
		Data: users,
	})
}
