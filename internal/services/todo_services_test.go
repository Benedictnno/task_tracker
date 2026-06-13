package service

import (
	"context"
	"testing"
	"todo-api/internal/models"
)

type MockTodoRepo struct {
	TodoRepository // Embed interface to satisfy it without implementing all methods if needed, though implementing them is cleaner.
}

func (m *MockTodoRepo) CreateTodo(
	ctx context.Context,
	title string,
	userID int,
) (models.Todo, error) {

	return models.Todo{
		ID:     1,
		Title:  title,
		UserID: userID,
	}, nil
}

// Dummy implementations for the rest of the interface
func (m *MockTodoRepo) GetTodos(ctx context.Context, userID int) ([]models.Todo, error) {
	return nil, nil
}
func (m *MockTodoRepo) GetTodoByID(ctx context.Context, id int, userID int) (*models.Todo, error) {
	return nil, nil
}
func (m *MockTodoRepo) UpdateTodo(ctx context.Context, id int, title string, completed bool, userID int) (*models.Todo, error) {
	return nil, nil
}
func (m *MockTodoRepo) DeleteTodo(ctx context.Context, id int, userID int) error { return nil }
func (m *MockTodoRepo) Close()                                                   {}

func TestCreateTodo_EmptyTitle(
	t *testing.T,
) {

	repo := &MockTodoRepo{}

	service := NewTodoService(repo)

	_, err := service.CreateTodo(
		context.Background(),
		"",
		1,
	)

	if err == nil {
		t.Error("expected error")
	}
}
