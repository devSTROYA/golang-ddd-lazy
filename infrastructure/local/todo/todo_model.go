package local

import "time"

type LocalTodo struct {
	Id          string
	Title       string
	Description *string
	UserId      string
	CreatedAt   time.Time
	CompletedAt *time.Time
}
