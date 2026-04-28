package user

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, user User) error
	FindById(ctx context.Context, id Id) (*User, error)
	Exists(ctx context.Context, email Email) (bool, error)
	FindAll(ctx context.Context) ([]User, error)
}
