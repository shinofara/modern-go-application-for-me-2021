package main

import (
	"context"
	"flag"
	"log"
	"mygo/config"
	"mygo/ent"
	"mygo/http/handler"
	oapi "mygo/http/openapi"
	"mygo/infrastructure/database"
	"mygo/infrastructure/mailer"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"

	"net/http"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config yaml path")
	flag.Parse()

	provides := []interface{}{
		context.Background,
		mailer.NewDummyMailer,
		handler.NewHandler,
		database.NewClient,
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

func Server(ctx context.Context, mux handler.Handler, db *ent.Client) error {
	defer func() {
		db.Close()
		log.Println("DB Close")
	}()

	r := chi.NewRouter()
	oapi.HandlerFromMux(&mux, r)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	go func() {
		log.Println("run: " + srv.Addr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Fatalln("Server closed with error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}
