package user

import (
	"context"
	userDomain "lazy/domain/user"
)

type Handler struct {
	userRepo userDomain.Repository
}

func NewHandler(userRepo userDomain.Repository) Handler {
	return Handler{userRepo}
}

func (h *Handler) Execute(ctx context.Context, query Query) (Result, error) {
	id := userDomain.IdFrom(query.Id)
	user, err := h.userRepo.FindById(ctx, id)
	if err != nil {
		return Result{}, err
	}
	if user == nil {
		return Result{}, userDomain.ErrUserDoesNotExist
	}

	return *user, nil
}
