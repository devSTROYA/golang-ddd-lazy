package todo

import "github.com/google/uuid"

type Id struct {
	value string
}

func NewId() Id {
	return Id{value: uuid.NewString()}
}

func IdFrom(value string) Id {
	return Id{value: value}
}

func (id Id) Value() string {
	return id.value
}

func (id Id) Equals(other Id) bool {
	return id.value == other.value
}
