package user

import (
	"context"
	uow "lazy/common"
	todoDomain "lazy/domain/todo"
	userDomain "lazy/domain/user"
)

type Handler struct {
	todoRepo todoDomain.Repository
	uow      uow.UnitOfWork
}

func NewHandler(todoRepo todoDomain.Repository, uow uow.UnitOfWork) Handler {
	return Handler{todoRepo, uow}
}

func (h *Handler) Execute(ctx context.Context, cmd Command) (Result, error) {
	title, err := todoDomain.NewTitle(cmd.Title)
	if err != nil {
		return Result{}, err
	}

	todo := todoDomain.New(todoDomain.NewTodoProps{
		Title:       title,
		Description: cmd.Description,
		UserId:      userDomain.IdFrom(cmd.UserId),
	})

	err = h.uow.WithTransaction(ctx, func(ctx context.Context) error {
		if err := h.todoRepo.Save(ctx, todo); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return Result{}, err
	}

	return Result{}, nil
}
