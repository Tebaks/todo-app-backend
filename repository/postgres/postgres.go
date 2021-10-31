package postgres

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"time"
	"todo-app/model"
	repo "todo-app/repository"

	_ "github.com/lib/pq"
)

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository() (repo.TodoRepository, error) {
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_MINUTES"))

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	Migrate(db)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set database connection settings
	db.SetMaxOpenConns(maxConn)                           // the default is 0, means unlimited
	db.SetMaxIdleConns(maxIdleConn)                       // the default is 2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) //0, connections are reused forever

	return &todoRepository{db}, nil
}

func (r *todoRepository) Close() {
	r.db.Close()
}

func (r *todoRepository) Find() ([]*model.Todo, error) {
	todos := make([]*model.Todo, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT id, todo, done FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		todo := new(model.Todo)
		err = rows.Scan(
			&todo.ID,
			&todo.Todo,
			&todo.Done,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)

	}

	return todos, nil
}

func (r *todoRepository) Create(todo *model.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO todos (id,todo,done) VALUES ($1, $2, $3)"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, todo.ID, todo.Todo, false)

	return err
}
