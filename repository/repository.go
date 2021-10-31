package repository

import "todo-app/model"

type TodoRepository interface {
	Close()
	Find() ([]*model.Todo, error)
	Create(user *model.Todo) error
}
