package postgres

import (
	"database/sql"
	"log"
	"testing"
	"todo-app/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var t1 = &model.Todo{
	ID:   uuid.New().String(),
	Todo: "Test Database",
	Done: false,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFindTodos(t *testing.T) {
	db, mock := NewMock()
	repo := &todoRepository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, todo, done FROM todos"

	rows := sqlmock.NewRows([]string{"id", "todo", "done"}).AddRow(t1.ID, t1.Todo, t1.Done)

	mock.ExpectQuery(query).WillReturnRows(rows)

	todos, err := repo.Find()
	assert.NotEmpty(t, todos)
	assert.NoError(t, err)
	assert.Len(t, todos, 1)
}

func TestCreateTodo(t *testing.T) {
	db, mock := NewMock()
	repo := &todoRepository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO todos \\(id,todo,done\\) VALUES \\(\\$1, \\$2, \\$3\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(t1.ID, t1.Todo, t1.Done).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Create(t1)
	assert.NoError(t, err)

}
