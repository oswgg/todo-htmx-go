package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oswgg/todo-htmx/internal/handlers"
	"github.com/oswgg/todo-htmx/internal/repositories"
	"github.com/oswgg/todo-htmx/internal/service"
	"net/http"
)

var connection *sql.DB

func Init() {
	var err error
	dsn := "root:rootOGG040520.dev@tcp(localhost:3306)/go_test?parseTime=true"
	connection, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	Init()
	taskRepo := repositories.NewMariadbTaskRepository(connection)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	root := http.NewServeMux()

	root.HandleFunc("POST /task", taskHandler.Create)
	root.HandleFunc("PUT /task/{id}/toggle", taskHandler.ToggleTask)
	root.HandleFunc("GET /task/update/{id}", taskHandler.UpdateView)
	root.HandleFunc("POST /task/update/{id}", taskHandler.Update)
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
