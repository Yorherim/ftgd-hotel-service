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

// HandleGetUserByID
//
//	@Summary		get user by ID
//	@Description	get user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	getUserByID200Example
//	@Failure		400	{object}	getUserByID400Example{data=nil}
//	@Failure		404	{object}	getUserByID404Example{data=nil}
//	@Router			/users/{id} [get]
func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)

	user, err := h.userService.GetUserByID(c.Context(), id)
	if err != nil {
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
		return fiber.NewError(fiber.StatusInternalServerError, "error get users")
	}

	return controller.Response(c, controller.ResponseType{
		Code: 200,
		Data: users,
	})
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var dto CreateUserDTO

	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "error parse body")
	}

	if err := h.validate.Struct(dto); err != nil {
		var errs []string
		var validatorErrs validator.ValidationErrors
		errors.As(err, &validatorErrs)

		for _, e := range validatorErrs {
			errs = append(errs, e.Error())
		}

		errorMessage := strings.Join(errs, ",\n")

		return fiber.NewError(fiber.StatusBadRequest, errorMessage)
	}

	user, err := h.userService.CreateUser(c.Context(), dto)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "error create user")
	}

	return controller.Response(c, controller.ResponseType{
		Code: 200,
		Data: user,
	})
}
