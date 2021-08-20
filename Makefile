oapi-generate:
	oapi-codegen -generate types -package oapi openapi/openapi.yaml > http/openapi/types.gen.go
	oapi-codegen -generate chi-server -package oapi openapi/openapi.yaml > http/openapi/server.gen.go

run:
	go run github.com/cosmtrek/air -c .air.toml

init:
	go install github.com/facebook/ent/cmd/entc@latest

model:
	go generate ./ent

migrate:
	docker compose run --rm migration --dryrun=false

migrate-dryrun:
	docker compose run --rm migration