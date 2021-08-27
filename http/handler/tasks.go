package handler

import (
	"fmt"
	"net/http"

	"github.com/shinofara/modern-go-application-for-me-2021/openapi"

	"github.com/goccy/go-json"
)

// GetMyTasks Get: /my_tasks
func (h *Handler) GetMyTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := h.db.User.Get(ctx, 16)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tasks, err := user.QueryAssignTasks().All(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	ts := new(openapi.Tasks)
	ts.Data = []openapi.Task{}
	for _, t := range tasks {
		ts.Data = append(ts.Data, openapi.Task{
			Title:           t.Title,
			CreatedDateTime: t.CreatedAt,
		})
	}
	//h.response().setStatus(http.StatusOK).json(w, tr)
	h.response().json(w, ts)
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
