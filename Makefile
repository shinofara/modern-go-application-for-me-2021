oapi-generate:
	oapi-codegen -generate types -package oapi openapi/openapi.yaml > http/openapi/types.gen.go
	oapi-codegen -generate chi-server -package oapi openapi/openapi.yaml > http/openapi/server.gen.go
