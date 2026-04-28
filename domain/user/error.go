package user

import (
	"errors"
)

var (
	ErrEmailAlreadyInUse  = errors.New("EMAIL_ALREADY_IN_USE")
	ErrInvalidEmailFormat = errors.New("INVALID_EMAIL_FORMAT")
	ErrNameTooShort       = errors.New("USER_NAME_TOO_SHORT")
	ErrUserDoesNotExist   = errors.New("USER_DOES_NOT_EXIST")
	ErrPasswordTooShort   = errors.New("PASSWORD_TOO_SHORT")
)
