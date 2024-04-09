package user

import (
	"errors"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (h *UserHandler) Register(appGroup fiber.Router) {
	users := appGroup.Group("users")
	{
		users.Get("/", h.HandleGetUsers)
		users.Get("/:id", h.HandleGetUserByID)
		users.Post("/create", h.HandleCreateUser)
	}
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

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var dto CreateUserDTO

	if err := c.BodyParser(&dto); err != nil {
		err = fiber.NewError(fiber.StatusBadRequest, "error parse body")
		return err
	}

	if err := h.validate.Struct(dto); err != nil {
		var errs []string
		var validatorErrs validator.ValidationErrors
		errors.As(err, &validatorErrs)

		for _, e := range validatorErrs {
			errs = append(errs, e.Error())
		}

		errorMessage := strings.Join(errs, ",\n")

		err = fiber.NewError(fiber.StatusBadRequest, errorMessage)
		return err
	}

	user, err := h.userService.CreateUser(c.Context(), dto)
	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "error create user")
		return err
	}

	return controller.Response(c, controller.ResponseType{
		Code: 200,
		Data: user,
	})
}
