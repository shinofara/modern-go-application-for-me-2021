oapi-merge:
	@docker run --rm -v "${PWD}/openapi/src:/w" -w /w openapitools/openapi-generator-cli generate \
	-g openapi-yaml -i openapi.yaml -o generated

oapi-generate:
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -generate types -package openapi openapi/src/generated/openapi/openapi.yaml > openapi/types.gen.go;
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -generate chi-server -package openapi openapi/src/generated/openapi/openapi.yaml > openapi/server.gen.go
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -generate spec -package openapi openapi/src/generated/openapi/openapi.yaml > openapi/spec.gen.go

run: migrate
	docker compose up app

model:
	go generate ./ent

migrate:
	docker compose run --rm migration

migrate-dryrun:
	docker compose run --rm migration

mockgen:
	go install github.com/golang/mock/mockgen@v1.6.0
	go run github.com/sanposhiho/gomockhandler@latest -f mockgen

init:
	docker compose up -d db
	$(MAKE) oapi-generate
	go get ./...
	$(MAKE) model
	$(MAKE) migrate
	$(MAKE) mockgen
	go test ./...
