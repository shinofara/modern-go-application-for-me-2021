package trace

import (
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

type Config struct {
	Jaeger *Jaeger
}

type Jaeger struct {
	URL string
}

func NewExporter(cfg *Config) (tracesdk.SpanExporter, error) {
	switch {
	case cfg.Jaeger != nil:
		return JaegerExporter(cfg.Jaeger)
	}

	return StdoutExporter()
}

func StdoutExporter() (tracesdk.SpanExporter, error){
	return stdouttrace.New(stdouttrace.WithPrettyPrint())
}

func JaegerExporter(cfg *Jaeger) (tracesdk.SpanExporter, error){
	url := cfg.URL
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
}