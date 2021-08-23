package main

import (
	"context"
	"flag"
	"mygo/config"
	"mygo/ent"
	"mygo/http/handler"
	oapi "mygo/http/oapi"
	"mygo/infrastructure/database"
	"mygo/infrastructure/logger"
	"mygo/infrastructure/mailer"
	"mygo/infrastructure/trace"
	"mygo/repository"
	"mygo/usecase"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"

	"net/http"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config yaml path")
	flag.Parse()

	l := logger.NewLogger("development")
	sh, err := logger.NewSentryHook()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	_ = l
	_ = sh
	//実際利用する差異はこのあたりを整える
	//l.AddHook(sh)

	provides := []interface{}{
		context.Background,
		mailer.NewDummyMailer,
		handler.NewHandler,
		repository.NewRepository,
		usecase.NewUseCase,
		database.NewClient,
		config.DB,
		config.Trace,
		trace.NewExporter,
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

func Server(ctx context.Context, p struct {
	dig.In

	// Server起動時に利用する引数は可変の可能性がある為、struct化してDIにて注入
	Mux           handler.Handler
	DB            *ent.Client
	TraceExporter tracesdk.SpanExporter
}) error {
	defer func() {
		p.DB.Close()
		log.Debug().Msg("DB Close")
	}()

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(p.TraceExporter),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("example"),
			attribute.String("environment", "development"),
		)),
	)
	otel.SetTracerProvider(tp)

	r := chi.NewRouter()
	r.Use(otelmux.Middleware("example", otelmux.WithTracerProvider(tp)))
	oapi.HandlerFromMux(&p.Mux, r)

	srv := &http.Server{
		Handler: otelhttp.NewHandler(r, "server",
			otelhttp.WithMessageEvents(otelhttp.ReadEvents, otelhttp.WriteEvents),
		),
		Addr: "0.0.0.0:8080",
	}

	go func() {
		log.Debug().Msgf("run: " + srv.Addr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Fatal().Msgf("Server closed with error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}
