package handler

import (
	"encoding/json"
	"net/http"
	"todo-app/model"
	"todo-app/service"
)

type todoHandler struct {
	service service.TodoService
}

func NewTodoHandler(service service.TodoService) *todoHandler {
	return &todoHandler{
		service: service,
	}
}

func (t *todoHandler) GetTodos(rw http.ResponseWriter, r *http.Request) {
	todos, err := t.service.GetTodos()
	if err != nil {
		http.Error(rw, "unable to get todos", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(todos)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
		return
	}
}

func (t *todoHandler) AddTodo(rw http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	e := json.NewDecoder(r.Body)
	err := e.Decode(&todo)
	if err != nil {
		http.Error(rw, "Error reading comment", http.StatusBadRequest)
		return
	}

	newTodo, err := t.service.AddTodo(&todo)
	if err != nil {
		http.Error(rw, "unable to add todo", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(newTodo)
}
