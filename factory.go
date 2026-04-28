package main

import (
	"lazy/common/platform-echo/guards"
	"lazy/presentation"
	todoPresentation "lazy/presentation/http/todo"
	userPresentation "lazy/presentation/http/user"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Controller struct {
	User userPresentation.Controller
	Todo todoPresentation.Controller
}

func NewApp(controller Controller) *echo.Echo {
	app := echo.New()
	app.Debug = true
	app.HideBanner = true
	app.Use(
		middleware.RequestLogger(),
		middleware.Recover(),
		middleware.RequestIDWithConfig(middleware.RequestIDConfig{
			Generator: func() string {
				return uuid.NewString()
			},
		}),
	)

	app.HTTPErrorHandler = presentation.HTTPFilter
	app.GET("/", RootHandler)

	authGroup := app.Group("/auth")
	userPresentation.NewRouter(authGroup, controller.User)

	todosGroup := app.Group("/todos")
	todosGroup.Use(guards.AuthGuard())
	todoPresentation.NewRouter(todosGroup, controller.Todo)

	return app
}

func RootHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello World!",
	})
}
