package user

import (
	"context"
	uow "lazy/common"
	todoDomain "lazy/domain/todo"
	"lazy/infrastructure/config"
)

type Handler struct {
	todoRepo todoDomain.Repository
	uow      uow.UnitOfWork
	env      config.Env
}

func NewHandler(todoRepo todoDomain.Repository, uow uow.UnitOfWork, env config.Env) Handler {
	return Handler{todoRepo, uow, env}
}

func (h *Handler) Execute(ctx context.Context, cmd Command) (Result, error) {
	todoId := todoDomain.IdFrom(cmd.TodoId)
	todo, err := h.todoRepo.FindById(ctx, todoId)
	if err != nil {
		return Result{}, err
	}

	err = todo.Complete()
	if err != nil {
		return Result{}, err
	}

	err = h.uow.WithTransaction(ctx, func(ctx context.Context) error {
		if err := h.todoRepo.Save(ctx, *todo); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return Result{}, err
	}

	return Result{}, nil
}
