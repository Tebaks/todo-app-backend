package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-app/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type TestTodoService struct {
	Todos []*model.Todo
}

func (t TestTodoService) GetTodos() ([]*model.Todo, error) {
	if t.Todos == nil {
		return []*model.Todo{}, nil
	}
	return t.Todos, nil
}

func (t TestTodoService) AddTodo(todo *model.Todo) (*model.Todo, error) {
	todo.ID = "1"
	return todo, nil
}

func TestAddTodos(t *testing.T) {
	todoPostBody := map[string]interface{}{
		"todo": "Test Add Todo",
	}
	body, _ := json.Marshal(todoPostBody)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))

	t.Run("Return todo with id", func(t *testing.T) {
		testService := TestTodoService{}
		handler := NewTodoHandler(testService)
		w := httptest.NewRecorder()
		handler.AddTodo(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "{\"id\":\"1\",\"todo\":\"Test Add Todo\",\"done\":false}\n", w.Body.String())
	})
}
func TestGetTodos(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	t.Run("Return empty data when no data in database", func(t *testing.T) {
		testService := TestTodoService{}
		handler := NewTodoHandler(testService)
		w := httptest.NewRecorder()
		handler.GetTodos(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "[]\n", w.Body.String())
	})

	t.Run("Return todos when data in database", func(t *testing.T) {
		testService := TestTodoService{Todos: []*model.Todo{
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
		}}
		handler := NewTodoHandler(testService)
		w := httptest.NewRecorder()
		handler.GetTodos(w, req)
		assert.Equal(t, 200, w.Code)

		var todos []*model.Todo
		err := json.Unmarshal(w.Body.Bytes(), &todos)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(todos))
	})

}
