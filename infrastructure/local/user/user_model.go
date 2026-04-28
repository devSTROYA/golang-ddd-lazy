package local

import (
	"time"
)

type LocalUser struct {
	Id        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
