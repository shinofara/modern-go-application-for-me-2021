package handler

import (
	"encoding/json"
	"fmt"
	"log"
	oapi "mygo/http/openapi"
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

	u := h.DB.User.Create()
	u.SetEmail(*p.Email).SetName(*p.Name).SetPassword(*p.Password)
	if _, err := u.Save(ctx); err != nil {
		log.Println(err.Error())
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, p)
}
