package main

import (
	"flag"
	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
	"log"
	"mygo/config"
	"mygo/ent"
	"mygo/http/handler"
	oapi "mygo/http/openapi"
	"mygo/infrastructure/database"
	"mygo/interfaces"

	"net/http"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config yaml path")
	flag.Parse()

	provides := []interface{}{
		interfaces.NewDummyMailer,
		handler.NewHandler,
		database.Open,
		config.DB,
	}

	container := dig.New()
	if err := container.Provide(func() (*config.Config, error) {
		return config.New(configPath)
	}); err != nil {
		panic(err)
	}

	for _, p := range provides {
		if err := container.Provide(p); err != nil {
			panic(err)
		}
	}

	if err := container.Invoke(Server); err != nil {
		panic(err)
	}
}

func Server(mux handler.Handler, db *ent.Client) error {
	defer func() {
		db.Close()
		log.Println("DB Close")
	}()

	r := chi.NewRouter()
	oapi.HandlerFromMux(&mux, r)
	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Println("run: " + s.Addr)

	return s.ListenAndServe()
}
