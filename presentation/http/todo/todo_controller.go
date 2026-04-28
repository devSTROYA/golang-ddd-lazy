package todo

import (
	addTodo "lazy/application/todo/commands/add_todo"
	completeTodo "lazy/application/todo/commands/complete_todo"
	getUserTodos "lazy/application/todo/queries/get_todos_user"
	"lazy/common/utils/array"
	todoDomain "lazy/domain/todo"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	getUserTodosHandler getUserTodos.Handler
	addTodoHandler      addTodo.Handler
	completeTodoHandler completeTodo.Handler
}

func NewController(getUserTodosHandler getUserTodos.Handler, addTodoHandler addTodo.Handler, completeTodoHandler completeTodo.Handler) Controller {
	return Controller{getUserTodosHandler, addTodoHandler, completeTodoHandler}
}

func (c *Controller) GetUserTodos(ctx echo.Context) error {
	userId, ok := ctx.Get("USER_ID").(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "USER_ID_NOT_FOUND")
	}

	result, err := c.getUserTodosHandler.Execute(ctx.Request().Context(), getUserTodos.Query{
		UserId: userId,
	})
	if err != nil {
		return err
	}

	response := GetUserTodosResponse{
		Data: array.Map(result.Todos, func(todo todoDomain.Todo) TodoDto {
			return ToDto(todo)
		}),
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) AddTodo(ctx echo.Context) error {
	userId, ok := ctx.Get("USER_ID").(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "USER_ID_NOT_FOUND")
	}

	var req AddTodoRequest

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "INVALID_REQUEST")
	}

	_, err := c.addTodoHandler.Execute(ctx.Request().Context(), addTodo.Command{
		UserId:      userId,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}

func (c *Controller) CompleteTodo(ctx echo.Context) error {
	todoId := ctx.Param("todoId")
	request := CompleteTodoRequest{
		TodoId: todoId,
	}

	_, err := c.completeTodoHandler.Execute(ctx.Request().Context(), completeTodo.Command{
		TodoId: request.TodoId,
	})
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
