package user

import (
	"regexp"
	"strings"
)

type Email struct {
	value string
}

func (e Email) Value() string {
	return e.value
}

var emailRegex = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,}$`)

func NewEmail(value string) (Email, error) {
	sanitized := strings.ToLower(strings.TrimSpace(value))
	if !emailRegex.MatchString(sanitized) {
		return Email{}, ErrInvalidEmailFormat
	}
	return Email{value: sanitized}, nil
}

func EmailFrom(value string) Email {
	return Email{value: value}
}

func (e Email) Equals(other Email) bool {
	return e.value == other.value
}
