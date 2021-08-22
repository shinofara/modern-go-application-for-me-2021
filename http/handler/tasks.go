package handler

import (
	"fmt"
	"net/http"
)

// GetMyTasks Get: /my_tasks
func (h *Handler) GetMyTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tasks, err := h.DB.Task.Query().All(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprint(w, tasks)
}