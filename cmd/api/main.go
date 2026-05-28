package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"todo-api/internal/db"
	"todo-api/internal/handlers"
	"todo-api/internal/middleware"
	"todo-api/internal/repository"
	 "todo-api/internal/services"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPool, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbPool.Close()

	todoRepo := repository.NewTodoRepository(dbPool)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	r := chi.NewRouter()
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.Timeout(5 * time.Second))
	r.Use(middleware.Logger)
	
	r.Get("/todos", todoHandler.GetTodos)
	r.Post("/todos", todoHandler.CreateTodo)
	r.Get("/todos/{id}", todoHandler.GetTodoByID)
	r.Put("/todos/{id}", todoHandler.UpdateTodo)
	r.Delete("/todos/{id}", todoHandler.DeleteTodo)

	
	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}
	go func() {
		log.Printf("Server is running on port %s\n", os.Getenv("PORT"))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", os.Getenv("PORT"), err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed:%+v", err)
	}
	log.Println("Server gracefully stopped")
}
