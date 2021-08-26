package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shinofara/example-go-2021/openapi"

	"github.com/goccy/go-json"
	"github.com/shinofara/example-go-2021/ent/auth"
)

func (h *Handler) PostSignup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var p openapi.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.useCase.Signup(ctx, &p); err != nil {
		log.Println(err)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, p)
}

func (h *Handler) PostSignin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var p openapi.SigninRequest
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a, err := h.db.Auth.Query().Where(auth.Email(p.Email), auth.Password(p.Password)).Only(ctx)
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
