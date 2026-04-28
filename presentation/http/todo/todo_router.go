package todo

import (
	"github.com/labstack/echo/v4"
)

func NewRouter(group *echo.Group, todoController Controller) {
	group.GET("", todoController.GetUserTodos)
	group.POST("", todoController.AddTodo)
	group.PATCH("/:todoId", todoController.CompleteTodo)
}
