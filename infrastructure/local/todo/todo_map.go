package local

import (
	"lazy/domain/todo"
	"lazy/domain/user"
)

func ToDomain(persistence LocalTodo) todo.Todo {
	return todo.From(todo.TodoFromProps{
		Id:          todo.IdFrom(persistence.Id),
		Title:       todo.TitleFrom(persistence.Title),
		Description: persistence.Description,
		UserId:      user.IdFrom(persistence.UserId),
		CreatedAt:   persistence.CreatedAt,
		CompletedAt: persistence.CompletedAt,
	})
}

func ToPersistence(domain todo.Todo) LocalTodo {
	return LocalTodo{
		Id:          domain.Id().Value(),
		Title:       domain.Title().Value(),
		Description: domain.Description(),
		UserId:      domain.UserId().Value(),
		CreatedAt:   domain.CreatedAt(),
		CompletedAt: domain.CompletedAt(),
	}
}
