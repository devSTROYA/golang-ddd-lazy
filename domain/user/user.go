package user

import (
	"time"
)

type User struct {
	id        Id
	name      Name
	email     Email
	password  Password
	createdAt time.Time
	updatedAt *time.Time
}

type NewUserProps struct {
	Name     Name
	Email    Email
	Password Password
}

type UserFromProps struct {
	Id        Id
	Name      Name
	Email     Email
	Password  Password
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (u User) Id() Id {
	return u.id
}

func (u User) Name() Name {
	return u.name
}

func (u User) Email() Email {
	return u.email
}

func (u User) Password() Password {
	return u.password
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) UpdatedAt() *time.Time {
	return u.updatedAt
}

func New(props NewUserProps) User {
	return User{
		id:        NewId(),
		name:      props.Name,
		email:     props.Email,
		password:  props.Password,
		createdAt: time.Now(),
		updatedAt: nil,
	}
}

func From(props UserFromProps) User {
	return User{
		id:        props.Id,
		name:      props.Name,
		email:     props.Email,
		password:  props.Password,
		createdAt: props.CreatedAt,
		updatedAt: props.UpdatedAt,
	}
}
