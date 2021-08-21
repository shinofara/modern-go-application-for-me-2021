package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"mygo/ent/auth"
	oapi "mygo/http/openapi"
	"mygo/usecase"
	"net/http"
)

type SignupRequest struct {
	Email    string
	Password string
	Name     string
}

func (h *Handler) PostSignup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var p oapi.Signup
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := usecase.Signup(ctx, &usecase.ServiceLocator{DB: h.DB, Mailer: h.Mailer}, &p); err != nil {
		log.Println(err)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, p)
}

func (h *Handler) PostSignin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var p oapi.Signin
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a, err := h.DB.Auth.Query().Where(auth.Email(p.Email), auth.Password(p.Password)).Only(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	u, err := a.QueryUser().Only(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprint(w, u)
}
