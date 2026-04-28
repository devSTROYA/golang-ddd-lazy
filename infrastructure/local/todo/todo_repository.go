package local

import (
	"context"
	"lazy/domain/todo"
	"lazy/domain/user"
)

type todoRepository struct {
	todos []LocalTodo
}

func NewTodoRepository() todo.Repository {
	todos := make([]LocalTodo, 0)
	return &todoRepository{todos}
}

func (repo *todoRepository) FindAllByUserId(ctx context.Context, userId user.Id) ([]todo.Todo, error) {
	todos := make([]todo.Todo, 0)
	for _, todo := range repo.todos {
		if todo.UserId == userId.Value() {
			todos = append(todos, ToDomain(todo))
		}
	}

	return todos, nil
}

func (repo *todoRepository) FindById(ctx context.Context, id todo.Id) (*todo.Todo, error) {
	for _, todo := range repo.todos {
		if todo.Id == id.Value() {
			selectedTodo := ToDomain(todo)
			return &selectedTodo, nil
		}
	}

	return nil, todo.ErrTodoDoesNotExist
}

func (repo *todoRepository) Save(ctx context.Context, todo todo.Todo) error {
	snapshot := ToPersistence(todo)
	for i, existingTodo := range repo.todos {
		if existingTodo.Id == todo.Id().Value() {
			repo.todos[i] = snapshot
			return nil
		}
	}

	repo.todos = append(repo.todos, snapshot)
	return nil
}
