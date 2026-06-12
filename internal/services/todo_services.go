package service

import (
	"context"
	"errors"

	"todo-api/internal/models"
	"todo-api/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(ctx context.Context, title string, userID int) (models.Todo, error) {
	if title == "" {
		return models.Todo{}, errors.New("title cannot be empty")
	}

	todo, err := s.repo.CreateTodo(ctx, title, userID)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) GetTodos(ctx context.Context, userID int) ([]models.Todo, error) {
	return s.repo.GetTodos(ctx, userID)
}

func (s *TodoService) GetTodoByID(ctx context.Context, id int, userID int) (models.Todo, error) {
	todo, err := s.repo.GetTodoByID(ctx, id, userID)
	if err != nil {
		return models.Todo{}, err
	}

	return *todo, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, id int, title string, completed bool, userID int) (models.Todo, error) {
	if title == "" {
		return models.Todo{}, errors.New("title cannot be empty")
	}

	todo, err := s.repo.UpdateTodo(ctx, id, title, completed, userID)
	if err != nil {
		return models.Todo{}, err
	}

	return *todo, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, id int, userID int) (bool, error) {
	err := s.repo.DeleteTodo(ctx, id, userID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *TodoService) Close() {
	s.repo.Close()
}
