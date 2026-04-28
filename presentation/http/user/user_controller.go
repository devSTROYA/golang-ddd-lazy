package user

import (
	registerUser "lazy/application/user/commands/register_user"
	getCurrentUser "lazy/application/user/queries/get_current_user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	registerUserHandler   registerUser.Handler
	getCurrentUserHandler getCurrentUser.Handler
}

func NewController(registerUserHandler registerUser.Handler, getCurrentUserHandler getCurrentUser.Handler) Controller {
	return Controller{registerUserHandler, getCurrentUserHandler}
}

func (c *Controller) RegisterUser(ctx echo.Context) error {
	var req RegisterUserRequest

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "INVALID_REQUEST")
	}

	result, err := c.registerUserHandler.Execute(ctx.Request().Context(), registerUser.Command{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return err
	}

	response := RegisterUserResponse{
		AccessToken: result.AccessToken,
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (c *Controller) GetCurrentUser(ctx echo.Context) error {
	userId, ok := ctx.Get("USER_ID").(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "USER_ID_NOT_FOUND")
	}

	result, err := c.getCurrentUserHandler.Execute(ctx.Request().Context(), getCurrentUser.Query{
		Id: userId,
	})
	if err != nil {
		return err
	}

	user := ToDto(result)

	response := GetCurrentUserResponse{
		Data: UserDto{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
	}

	return ctx.JSON(http.StatusOK, response)
}
