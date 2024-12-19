package main

import (
	"fmt"
	"github.com/oswgg/todo-htmx/internal/handlers"
	"github.com/oswgg/todo-htmx/internal/models"
	"github.com/oswgg/todo-htmx/internal/repositories"
	"github.com/oswgg/todo-htmx/internal/service"
	"net/http"
	"text/template"
)

func main() {
	taskRepo := repositories.NewMockTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	root := http.NewServeMux()

	root.HandleFunc("POST /task", taskHandler.Create)
	root.HandleFunc("PUT /task/{id}/toggle", taskHandler.ToggleTask)
	root.HandleFunc("GET /", homeHandler)

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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/base.html", "web/templates/index.html"))
	err := tmpl.ExecuteTemplate(w, "base.html", []models.Task{})
	if err != nil {
		fmt.Println(err)
	}
}
