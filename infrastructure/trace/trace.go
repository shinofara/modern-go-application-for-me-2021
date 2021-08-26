package trace

import (
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
)

type Config struct {
	Jaeger *Jaeger
	GCP *GCP
}

type Jaeger struct {
	URL string
}

type GCP struct {
	ProjectID string `yaml:"project_id"`
}

func NewExporter(cfg *Config) (tracesdk.SpanExporter, error) {
	switch {
	case cfg.Jaeger != nil:
		return JaegerExporter(cfg.Jaeger)
	case cfg.GCP != nil:
		return texporter.New(texporter.WithProjectID(cfg.GCP.ProjectID))
	}

	return StdoutExporter()
}

func StdoutExporter() (tracesdk.SpanExporter, error) {
	return stdouttrace.New(stdouttrace.WithPrettyPrint())
}

func JaegerExporter(cfg *Jaeger) (tracesdk.SpanExporter, error) {
	url := cfg.URL
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
}
