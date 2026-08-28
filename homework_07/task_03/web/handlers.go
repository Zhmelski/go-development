package web

import (
	"embed"
	"html/template"
	"net/http"
	"strconv"
	"task_02/todoapp"
)

//go:embed templates/*
var templateFS embed.FS

// 'Handler' processes web requests
type Handler struct {
	app  *todoapp.TodoApp
	tmpl *template.Template
}

// 'NewHandler' creates a new web handler
func NewHandler(app *todoapp.TodoApp) (*Handler, error) {
	// Download templates with formatting features
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"formatDate": func(t interface{}) string {
			if t == nil {
				return "-"
			}
			// Date formatting
			return ""
		},
	}).ParseFS(templateFS, "templates/*.html")

	if err != nil {
		return nil, err
	}

	return &Handler{
		app:  app,
		tmpl: tmpl,
	}, nil
}

// 'IndexHandler' displays the main page with a list of tasks
func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tasks := h.app.GetAllTasks()

	data := struct {
		Tasks []interface{}
	}{
		Tasks: make([]interface{}, len(tasks)),
	}

	for i, task := range tasks {
		completedAt := ""
		if task.CompletedAt != nil {
			completedAt = task.CompletedAt.Format("02.01.2006 15:04")
		}

		data.Tasks[i] = map[string]interface{}{
			"ID":          task.ID,
			"Description": task.Description,
			"Completed":   task.Completed,
			"CreatedAt":   task.CreatedAt.Format("02.01.2006 15:04"),
			"CompletedAt": completedAt,
		}
	}

	if err := h.tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// 'AddTaskHandler' handles adding a new task
func (h *Handler) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	description := r.FormValue("description")
	if description == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := h.app.AddTask(description); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// 'CompleteTaskHandler' marks the task as completed
func (h *Handler) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID задачи", http.StatusBadRequest)
		return
	}

	if err := h.app.CompleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// 'DeleteTaskHandler' deletes a task
func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID задачи", http.StatusBadRequest)
		return
	}

	if err := h.app.DeleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
