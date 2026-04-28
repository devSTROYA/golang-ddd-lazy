package user

import (
	"context"
	uow "lazy/common"
	userDomain "lazy/domain/user"
	"lazy/infrastructure/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Handler struct {
	userRepo userDomain.Repository
	uow      uow.UnitOfWork
	env      config.Env
}

func NewHandler(userRepo userDomain.Repository, uow uow.UnitOfWork, env config.Env) Handler {
	return Handler{userRepo, uow, env}
}

func (h *Handler) Execute(ctx context.Context, cmd Command) (Result, error) {
	name, err := userDomain.NewName(cmd.Name)
	if err != nil {
		return Result{}, err
	}

	email, err := userDomain.NewEmail(cmd.Email)
	if err != nil {
		return Result{}, err
	}

	password, err := userDomain.NewPassword(cmd.Password)
	if err != nil {
		return Result{}, err
	}

	existingUser, err := h.userRepo.Exists(ctx, email)
	if err != nil {
		return Result{}, err
	}
	if existingUser {
		return Result{}, userDomain.ErrEmailAlreadyInUse
	}

	user := userDomain.New(userDomain.NewUserProps{
		Name:     name,
		Email:    email,
		Password: password,
	})

	err = h.uow.WithTransaction(ctx, func(ctx context.Context) error {
		if err := h.userRepo.Save(ctx, user); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return Result{}, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   user.Id().Value(),
		ID:        uuid.NewString(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(h.env.JwtExpirationTime) * time.Hour)),
	})
	accessToken, err := jwtToken.SignedString([]byte(h.env.JwtSecret))
	if err != nil {
		return Result{}, err
	}

	return Result{accessToken}, nil
}
