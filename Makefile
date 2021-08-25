oapi-generate:
	oapi-codegen -generate types -package oapi openapi.yaml > http/oapi/types.gen.go
	oapi-codegen -generate chi-server -package oapi openapi.yaml > http/oapi/server.gen.go

run: migrate
	docker compose up app

model:
	go generate ./ent

migrate:
	docker compose run --rm migration --dryrun=false --config environment/development/config.yml

migrate-dryrun:
	docker compose run --rm migration

mockgen:
	go install github.com/golang/mock/mockgen@v1.6.0
	go run github.com/sanposhiho/gomockhandler@latest -f mockgen

init:
	go get ./...
	$(MAKE) model
	$(MAKE) oapi-generate
	$(MAKE) migrate
	$(MAKE) mockgen
