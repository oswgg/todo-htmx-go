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

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/base.html", "web/templates/index.html"))
	list, err := h.service.List()
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.ExecuteTemplate(w, "base.html", list)
	if err != nil {
		fmt.Println(err)
	}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var name string = r.PostFormValue("Name")
	taskToCreate := models.Task{
		Name:      name,
		CreatedAt: time.Now(),
		Completed: false,
	}

	// Call the service to create the task
	newTask, err := h.service.Create(&taskToCreate)
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	err = tmpl.ExecuteTemplate(w, "task", newTask)
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
