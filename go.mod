module mygo

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/XSAM/otelsql v0.5.0
	github.com/cosmtrek/air v1.27.3
	github.com/go-chi/chi/v5 v5.0.3
	github.com/go-sql-driver/mysql v1.6.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.22.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.22.0
	go.opentelemetry.io/otel v1.0.0-RC2
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0-RC2 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.0-RC2
	go.opentelemetry.io/otel/sdk v1.0.0-RC2
	go.opentelemetry.io/otel/trace v1.0.0-RC2
	go.uber.org/dig v1.12.0
	gopkg.in/yaml.v2 v2.4.0

)
