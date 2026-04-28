package todo

import (
	"errors"
)

var (
	ErrTitleTooShort        = errors.New("TODO_TITLE_TOO_SHORT")
	ErrTodoDoesNotExist     = errors.New("TODO_DOES_NOT_EXIST")
	ErrTodoAlreadyCompleted = errors.New("TODO_ALREADY_COMPLETED")
)
