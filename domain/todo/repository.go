package todo

import (
	"context"
	"lazy/domain/user"
)

type Repository interface {
	Save(ctx context.Context, todo Todo) error
	FindById(ctx context.Context, id Id) (*Todo, error)
	FindAllByUserId(ctx context.Context, userId user.Id) ([]Todo, error)
}
