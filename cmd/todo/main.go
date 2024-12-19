package main

import (
	"fmt"
	"github.com/oswgg/todo-htmx/internal/handlers"
	"github.com/oswgg/todo-htmx/internal/repositories"
	"github.com/oswgg/todo-htmx/internal/service"
	"net/http"
)

func main() {
	taskRepo := repositories.NewMockTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	root := http.NewServeMux()

	root.HandleFunc("POST /task", taskHandler.Create)
	root.HandleFunc("PUT /task/{id}/toggle", taskHandler.ToggleTask)
	root.HandleFunc("GET /", taskHandler.List)

	server := &http.Server{
		Addr:    ":8080",
		Handler: root,
	}
	fmt.Printf("Listening on %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
