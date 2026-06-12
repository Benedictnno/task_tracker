package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/internal/middleware"
	service "todo-api/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok || userID <= 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	todos, err := h.service.GetTodos(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := r.Context().Value(middleware.UserIDKey).(int)

	if input.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	todo, err := h.service.CreateTodo(
		r.Context(),
		input.Title,
		userID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	userID := r.Context().Value(middleware.UserIDKey).(int)
	todo, err := h.service.GetTodoByID(r.Context(), id, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var input struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := r.Context().Value(middleware.UserIDKey).(int)

	todo, err := h.service.UpdateTodo(r.Context(), id, input.Title, input.Completed, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	userID := r.Context().Value(middleware.UserIDKey).(int)	
	_, err = h.service.DeleteTodo(r.Context(), id, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted successfully"})
}
