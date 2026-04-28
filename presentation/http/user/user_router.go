package user

import (
	"lazy/common/platform-echo/guards"

	"github.com/labstack/echo/v4"
)

func NewRouter(group *echo.Group, userController Controller) {
	group.POST("/register", userController.RegisterUser)
	group.GET("/info", userController.GetCurrentUser, guards.AuthGuard())
}
