package handler

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func (h Handler) Todo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "add todo")
	return
}