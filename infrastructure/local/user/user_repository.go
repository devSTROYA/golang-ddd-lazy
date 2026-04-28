package local

import (
	"context"
	"lazy/domain/user"
)

type userRepository struct {
	users []LocalUser
}

func NewUserRepository() user.Repository {
	users := make([]LocalUser, 0)
	return &userRepository{users}
}

func (repo *userRepository) FindAll(ctx context.Context) ([]user.User, error) {
	users := make([]user.User, 0)
	for _, user := range repo.users {
		users = append(users, ToDomain(user))
	}

	return users, nil
}

func (repo *userRepository) FindById(ctx context.Context, id user.Id) (*user.User, error) {
	for _, user := range repo.users {
		if user.Id == id.Value() {
			selectedUser := ToDomain(user)
			return &selectedUser, nil
		}
	}

	return nil, user.ErrUserDoesNotExist
}

func (repo *userRepository) Exists(ctx context.Context, email user.Email) (bool, error) {
	for _, user := range repo.users {
		if user.Email == email.Value() {
			return true, nil
		}
	}

	return false, nil
}

func (repo *userRepository) Save(ctx context.Context, user user.User) error {
	snapshot := ToPersistence(user)
	for i, existingUser := range repo.users {
		if existingUser.Id == user.Id().Value() {
			repo.users[i] = snapshot

			return nil
		}
	}

	repo.users = append(repo.users, snapshot)

	return nil
}
