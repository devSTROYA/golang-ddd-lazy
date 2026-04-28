package todo

import todoDomain "lazy/domain/todo"

func ToDto(domain todoDomain.Todo) TodoDto {
	return TodoDto{
		Id:          domain.Id().Value(),
		Title:       domain.Title().Value(),
		Description: domain.Description(),
		CreatedAt:   domain.CreatedAt().Unix(),
		IsCompleted: domain.CompletedAt() != nil,
	}
}
