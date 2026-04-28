package todo

import (
	"context"
	todoDomain "lazy/domain/todo"
	userDomain "lazy/domain/user"
)

type Handler struct {
	todoRepo todoDomain.Repository
	userRepo userDomain.Repository
}

func NewHandler(todoRepo todoDomain.Repository, userRepo userDomain.Repository) Handler {
	return Handler{todoRepo, userRepo}
}

func (h *Handler) Execute(ctx context.Context, query Query) (Result, error) {
	userId := userDomain.IdFrom(query.UserId)
	user, err := h.userRepo.FindById(ctx, userId)
	if err != nil {
		return Result{}, err
	}
	println(user)
	if user == nil {
		return Result{}, userDomain.ErrUserDoesNotExist
	}

	todos, err := h.todoRepo.FindAllByUserId(ctx, userId)
	if err != nil {
		return Result{}, err
	}

	return Result{Todos: todos}, nil
}
