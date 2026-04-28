package local

import "lazy/domain/user"

func ToDomain(persistence LocalUser) user.User {
	domain := user.From(user.UserFromProps{
		Id:        user.IdFrom(persistence.Id),
		Name:      user.NameFrom(persistence.Name),
		Email:     user.EmailFrom(persistence.Email),
		Password:  user.PasswordFrom(persistence.Password),
		CreatedAt: persistence.CreatedAt,
		UpdatedAt: persistence.UpdatedAt,
	})

	return domain
}

func ToPersistence(domain user.User) LocalUser {
	return LocalUser{
		Id:        domain.Id().Value(),
		Name:      domain.Name().Value(),
		Email:     domain.Email().Value(),
		Password:  domain.Password().Value(),
		CreatedAt: domain.CreatedAt(),
		UpdatedAt: domain.UpdatedAt(),
	}
}
