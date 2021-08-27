package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.opentelemetry.io/otel/propagation"

	"github.com/shinofara/modern-go-application-for-me-2021/openapi"

	"github.com/shinofara/modern-go-application-for-me-2021/config"
	"github.com/shinofara/modern-go-application-for-me-2021/ent"
	"github.com/shinofara/modern-go-application-for-me-2021/http/handler"
	"github.com/shinofara/modern-go-application-for-me-2021/infrastructure/database"
	"github.com/shinofara/modern-go-application-for-me-2021/infrastructure/logger"
	"github.com/shinofara/modern-go-application-for-me-2021/infrastructure/mailer"
	"github.com/shinofara/modern-go-application-for-me-2021/infrastructure/trace"
	"github.com/shinofara/modern-go-application-for-me-2021/repository"
	"github.com/shinofara/modern-go-application-for-me-2021/usecase"

	"github.com/rs/zerolog/log"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"net/http"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config yaml path")
	flag.Parse()

	cfg, err := config.New(configPath)
	if err != nil {
		panic(err)
	}

	l := logger.NewLogger(cfg.Logger)
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
		cfg.Clone,
		trace.NewExporter,
	}

	container := dig.New()
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
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(p.TraceExporter),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("example"),
			attribute.String("environment", "development"),
		)),
	)
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Err(err).Msg("")
		}
	}()

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	r := chi.NewRouter()

	swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Printf("failed to get swagger spec: %v\n", err)
	}
	swagger.Servers = nil

	r.Use(middleware.OapiRequestValidator(swagger))
	rr := openapi.HandlerFromMux(&p.Mux, r)

	srv := &http.Server{
		Handler: otelhttp.NewHandler(rr, "",
			otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
				return r.Method + ": " + r.URL.Path
			}),
			otelhttp.WithMessageEvents(
				otelhttp.ReadEvents,
				otelhttp.WriteEvents,
			),
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
