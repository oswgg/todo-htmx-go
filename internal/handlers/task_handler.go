package handlers

import (
	"fmt"
	"github.com/oswgg/todo-htmx/internal/models"
	"github.com/oswgg/todo-htmx/internal/service"
	"net/http"
	"strconv"
	"strings"
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

func (h *TaskHandler) UpdateView(w http.ResponseWriter, r *http.Request) {
	var task *models.Task
	var err error
	var id int

	id, err = strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/task/update/"))
	if err != nil {
		fmt.Println(err)
	}

	task, err = h.service.FindByID(int64(id))

	tmpl := template.Must(template.ParseFiles("web/templates/base.html", "web/templates/update-task.html"))
	err = tmpl.ExecuteTemplate(w, "base.html", task)
	if err != nil {
		fmt.Println(err)
	}
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	var task *models.Task
	var err error
	var id int

	id, err = strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/task/update/"))
	if err != nil {
		fmt.Println(err)
	}

	name := r.PostFormValue("Name")

	task, err = h.service.FindByID(int64(id))
	task.Name = name

	_, err = h.service.Update(task)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("HX-Location", "/")
	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) ToggleTask(w http.ResponseWriter, r *http.Request) {
}
