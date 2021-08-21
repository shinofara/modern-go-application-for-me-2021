package handler

import (
	"encoding/json"
	"fmt"
	"log"
	oapi "mygo/http/openapi"
	"mygo/usecase"
	"net/http"
)

type SignupRequest struct {
	Email string
	Password string
	Name string
}

func (h *Handler) 	PostSignup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var p oapi.Signup
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := usecase.Signup(ctx, h.DB, &p); err != nil {
		log.Println(err)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, p)
}
