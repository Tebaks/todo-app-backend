package service

import (
	"testing"
	"todo-app/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type TestTodoRepository struct {
	Todos []*model.Todo
}

func (t TestTodoRepository) Close() {

}

func (t TestTodoRepository) Create(todo *model.Todo) error {
	return nil
}

func (t TestTodoRepository) Find() ([]*model.Todo, error) {
	return t.Todos, nil
}

func TestGetTodos(t *testing.T) {
	t.Run("Return empty data when no data in database", func(t *testing.T) {
		testRepo := TestTodoRepository{}
		service := NewTodoService(testRepo)

		todos, err := service.GetTodos()

		assert.Nil(t, err)
		assert.Equal(t, 0, len(todos))
	})
	t.Run("Return empty data when data in database", func(t *testing.T) {

		testRepo := TestTodoRepository{Todos: []*model.Todo{
			{
				ID:   uuid.New().String(),
				Todo: "Learn Test",
				Done: false,
			},
			{
				ID:   uuid.New().String(),
				Todo: "Learn Test Service",
				Done: true,
			},
		},
		}

		service := NewTodoService(testRepo)
		todos, err := service.GetTodos()

		assert.Nil(t, err)
		assert.Equal(t, 2, len(todos))
	})

}

func TestAddTodo(t *testing.T) {
	testRepo := TestTodoRepository{}
	service := NewTodoService(testRepo)

	todo := &model.Todo{
		Todo: "Learn Test",
		Done: false,
	}
	newTodo, err := service.AddTodo(todo)

	assert.Nil(t, err)
	assert.Equal(t, newTodo.Todo, todo.Todo)
	assert.NotNil(t, newTodo.ID)
}
