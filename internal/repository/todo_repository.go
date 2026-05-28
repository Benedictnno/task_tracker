package repository

import (
	"context"
	"todo-api/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepository struct {
	DB *pgxpool.Pool
}

func NewTodoRepository(db *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) GetTodos(ctx context.Context) ([]models.Todo, error) {
	rows, err := r.DB.Query(ctx, "SELECT id, title, completed, created_at FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]models.Todo, 0)
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodoRepository) GetTodoByID(ctx context.Context, id int) (*models.Todo, error) {
	var todo models.Todo
	err := r.DB.QueryRow(ctx, "SELECT id, title, completed, created_at FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) CreateTodo(ctx context.Context, title string) (*models.Todo, error) {
	var todo models.Todo
	err := r.DB.QueryRow(ctx, "INSERT INTO todos (title) VALUES ($1) RETURNING id, title, completed, created_at", title).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) UpdateTodo(ctx context.Context, id int, title string, completed bool) (*models.Todo, error) {
	var todo models.Todo
	err := r.DB.QueryRow(ctx, "UPDATE todos SET title = $1, completed = $2 WHERE id = $3 RETURNING id, title, completed, created_at", title, completed, id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, id int) error {
	result, err := r.DB.Exec(ctx, "DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *TodoRepository) Close() {
	r.DB.Close()
}
