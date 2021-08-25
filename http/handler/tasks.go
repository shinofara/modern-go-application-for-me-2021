package handler

import (
	"fmt"
	"github.com/shinofara/example-go-2021/openapi"
	"net/http"

	"github.com/goccy/go-json"
)

// GetMyTasks Get: /my_tasks
func (h *Handler) GetMyTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := h.db.User.Get(ctx, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tasks, err := user.QueryAssignTasks().All(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var tr []openapi.Task
	for _, t := range tasks {
		tr = append(tr, openapi.Task{
			Title: t.Title,
		})
	}

	fmt.Fprint(w, tr)
}

func (h *Handler) PostTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var p openapi.Task

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.useCase.CreateTask(ctx, &p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "ok")
}
