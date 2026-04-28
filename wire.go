// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	addTodoCommand "lazy/application/todo/commands/add_todo"
	completeTodoCommand "lazy/application/todo/commands/complete_todo"
	getUserTodosQuery "lazy/application/todo/queries/get_todos_user"
	registerUserCommand "lazy/application/user/commands/register_user"
	getCurrentUserQuery "lazy/application/user/queries/get_current_user"
	"lazy/infrastructure/config"
	localUow "lazy/infrastructure/local"
	localTodo "lazy/infrastructure/local/todo"
	localUser "lazy/infrastructure/local/user"
	todoPresentation "lazy/presentation/http/todo"
	userPresentation "lazy/presentation/http/user"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func InitializeApp() *echo.Echo {
	wire.Build(
		config.NewEnv,
		localUser.NewUserRepository,
		localTodo.NewTodoRepository,
		localUow.NewUnitOfWork,

		registerUserCommand.NewHandler,
		getCurrentUserQuery.NewHandler,
		userPresentation.NewController,

		getUserTodosQuery.NewHandler,
		todoPresentation.NewController,
		addTodoCommand.NewHandler,
		completeTodoCommand.NewHandler,
		wire.Struct(new(Controller), "*"),
		NewApp,
	)

	return echo.New()
}
