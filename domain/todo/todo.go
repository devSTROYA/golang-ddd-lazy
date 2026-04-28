package todo

import (
	"time"

	"lazy/domain/user"
)

type Todo struct {
	id          Id
	title       Title
	description *string
	userId      user.Id
	createdAt   time.Time
	completedAt *time.Time
}

type NewTodoProps struct {
	Title       Title
	Description *string
	UserId      user.Id
}

type TodoFromProps struct {
	Id          Id
	Title       Title
	Description *string
	UserId      user.Id
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func (t *Todo) Id() Id {
	return t.id
}

func (t *Todo) Title() Title {
	return t.title
}

func (t *Todo) Description() *string {
	return t.description
}

func (t *Todo) UserId() user.Id {
	return t.userId
}

func (t *Todo) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Todo) CompletedAt() *time.Time {
	return t.completedAt
}

func (t *Todo) Complete() error {
	if t.completedAt != nil {
		return ErrTodoAlreadyCompleted
	}

	now := time.Now()
	t.completedAt = &now

	return nil
}

func New(props NewTodoProps) Todo {
	return Todo{
		id:          NewId(),
		title:       props.Title,
		description: props.Description,
		userId:      props.UserId,
		createdAt:   time.Now(),
		completedAt: nil,
	}
}

func From(props TodoFromProps) Todo {
	return Todo{
		id:          props.Id,
		title:       props.Title,
		description: props.Description,
		userId:      props.UserId,
		createdAt:   props.CreatedAt,
		completedAt: props.CompletedAt,
	}
}
