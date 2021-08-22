package main

import (
	"context"
	"flag"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
	"mygo/config"
	"mygo/ent"
	"mygo/http/handler"
	oapi "mygo/http/oapi"
	"mygo/infrastructure/database"
	"mygo/infrastructure/mailer"
	"mygo/repository"
	"mygo/usecase"
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
		repository.NewRepository,
		usecase.NewUseCase,
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

	// Provideに登録した依存を解決させて、Serverを実行
	if err := container.Invoke(Server); err != nil {
		panic(err)
	}
}

func Server(ctx context.Context, p struct{
	dig.In

	// Server起動時に利用する引数は可変の可能性がある為、struct化してDIにて注入
	Mux handler.Handler
	DB *ent.Client
}) error {
	defer func() {
		p.DB.Close()
		log.Println("DB Close")
	}()

	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("creating stdout exporter: %v", err)
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exporter),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("hogehoge"),
			attribute.String("environment", "development"),
			attribute.Int64("ID", 123),
		)),
	)
	otel.SetTracerProvider(tp)

	r := chi.NewRouter()
	oapi.HandlerFromMux(&p.Mux, r)
	srv := &http.Server{
		Handler: otelhttp.NewHandler(r, "server",
			otelhttp.WithMessageEvents(otelhttp.ReadEvents, otelhttp.WriteEvents),
		),
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
