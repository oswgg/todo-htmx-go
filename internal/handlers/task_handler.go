package handlers

import (
	"fmt"
	"github.com/oswgg/todo-htmx/internal/models"
	"github.com/oswgg/todo-htmx/internal/service"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a Task
	var name string = r.PostFormValue("Name")
	newTask := models.Task{
		ID:        1,
		Name:      name,
		CreatedAt: time.Now(),
		Completed: false,
	}

	// Call the service to create the task
	list, err := h.service.Create(&newTask)
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	err = tmpl.ExecuteTemplate(w, "task", list[len(list)-1])
	if err != nil {
		fmt.Println(err)
	}

}

func (h *TaskHandler) ToggleTask(w http.ResponseWriter, r *http.Request) {
	var id int
	var err error
	var updatedTask *models.Task

	id, err = strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Println(err)
	}

	updatedTask, err = h.service.Toggle(int64(id))

	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	err = tmpl.ExecuteTemplate(w, "task", updatedTask)
	if err != nil {
		fmt.Println(err)
	}
}
