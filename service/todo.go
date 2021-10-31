package service

import (
	"todo-app/model"
	"todo-app/repository"

	"github.com/google/uuid"
)

type TodoService interface {
	GetTodos() ([]*model.Todo, error)
	AddTodo(*model.Todo) (*model.Todo, error)
}

type todoService struct {
	repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{
		repo,
	}
}

func (t *todoService) GetTodos() ([]*model.Todo, error) {
	todos, err := t.Find()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoService) AddTodo(todo *model.Todo) (*model.Todo, error) {
	todo.ID = uuid.New().String()
	err := t.Create(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
