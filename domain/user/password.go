package user

import (
	"golang.org/x/crypto/bcrypt"
)

const MIN_PASSWORD_LENGTH = 8

type Password struct {
	value string
}

func (pw Password) Value() string {
	return pw.value
}

func NewPassword(value string) (Password, error) {
	if len(value) < MIN_PASSWORD_LENGTH {
		return Password{}, ErrPasswordTooShort
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(value), 10)
	if err != nil {
		return Password{}, err
	}
	return Password{value: string(bytes)}, nil
}

func PasswordFrom(value string) Password {
	return Password{value: value}
}

func (pw Password) Equals(other Password) bool {
	return pw.value == other.value
}

func (pw Password) Compare(other Password) bool {
	return bcrypt.CompareHashAndPassword([]byte(pw.value), []byte(other.value)) == nil
}
