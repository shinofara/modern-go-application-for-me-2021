package handler

import (
	"encoding/json"
	"fmt"
	"mygo/http/oapi"
	"net/http"
)

// GetMyTasks Get: /my_tasks
func (h *Handler) GetMyTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := h.DB.User.Get(ctx, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tasks, err := user.QueryAssignTasks().All(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprint(w, tasks)
}

func (h *Handler) PostTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var p oapi.Task
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.DB.User.Get(ctx, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tc := h.DB.Task.Create()

	if _, err := tc.SetTitle(p.Title).
		SetCreator(user).
		SetAssign(user).Save(ctx);
	err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "ok")
}