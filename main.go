package main

import (
	"github.com/go-chi/chi/v5"
	"mygo/http/handler"
	oapi "mygo/http/openapi"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	mux := &handler.Handler{}
	oapi.HandlerFromMux(mux, r)
	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
