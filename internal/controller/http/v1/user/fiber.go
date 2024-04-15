package user

import (
	"errors"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller"
	"github.com/Yorherim/ftgd-hotel-service/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"unicode"
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
	h.logger.Info("HandleGetUserByID start")

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

// HandleGetUsers
//
//	@Summary		get all users
//	@Description	get all users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	getUsers200Example
//	@Failure		500	{object}	serverError500Example{data=nil}
//	@Router			/users [get]
func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	h.logger.Info("HandleGetUsers start")

	users, err := h.userService.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return controller.Response(c, controller.ResponseType{
		Code: 200,
		Data: users,
	})
}

// HandleCreateUser
//
//	@Summary		create user
//	@Description	create user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateUserDTO	true	"create user fields"
//	@Success		201		{object}	getUserByID200Example
//	@Failure		400		{object}	getUserByID400Example{data=nil}
//	@Failure		500		{object}	serverError500Example{data=nil}
//	@Router			/users/create [post]
func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	h.logger.Info("HandleCreateUser start")

	var dto CreateUserDTO

	if err := c.BodyParser(&dto); err != nil {
		h.logger.Errorf("HandleCreateUser error BodyParser: %s", err)
		return fiber.NewError(fiber.StatusBadRequest, "error parse body")
	}

	// validate fields
	if err := h.validate.Struct(dto); err != nil {
		var errs []any
		var validatorErrs validator.ValidationErrors

		errors.As(err, &validatorErrs)

		for _, e := range validatorErrs {
			customErr := h.utils.ValidatorGetCustomMsg(e)

			r := []rune(e.Field())
			r[0] = unicode.ToLower(r[0])

			errs = append(errs, utils.ValidateError{
				Param:   string(r),
				Message: customErr,
			})
		}

		h.logger.Errorf("HandleCreateUser error validate: %s", errs)
		return controller.Response(c, controller.ResponseType{
			Code:   400,
			Data:   nil,
			Errors: errs,
		})
	}

	user, err := h.userService.CreateUser(c.Context(), dto)
	if err != nil {
		return err
	}

	return controller.Response(c, controller.ResponseType{
		Code: 201,
		Data: user,
	})
}
