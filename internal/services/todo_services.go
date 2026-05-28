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

func (s *TodoService) CreateTodo(ctx context.Context, title string) (models.Todo, error) {
	if title == "" {
		return models.Todo{}, errors.New("title cannot be empty")
	}

	todo, err := s.repo.CreateTodo(ctx, title)
	if err != nil {
		return models.Todo{}, err
	}

	return *todo, nil
}

func (s *TodoService) GetTodos(ctx context.Context) ([]models.Todo, error) {
	return s.repo.GetTodos(ctx)
}

func (s *TodoService) GetTodoByID(ctx context.Context, id int) (models.Todo, error) {
	todo, err := s.repo.GetTodoByID(ctx, id)
	if err != nil {
		return models.Todo{}, err
	}

	return *todo, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, id int, title string, completed bool) (models.Todo, error) {
	if title == "" {
		return models.Todo{}, errors.New("title cannot be empty")
	}

	todo, err := s.repo.UpdateTodo(ctx, id, title, completed)
	if err != nil {
		return models.Todo{}, err
	}

	return *todo, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, id int) (bool, error) {
	err := s.repo.DeleteTodo(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *TodoService) Close() {
	s.repo.Close()
}
