package user

import userDomain "lazy/domain/user"

func ToDto(domain userDomain.User) UserDto {
	return UserDto{
		Id:        domain.Id().Value(),
		Name:      domain.Name().Value(),
		Email:     domain.Email().Value(),
		CreatedAt: domain.CreatedAt().Unix(),
	}
}
