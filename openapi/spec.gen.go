// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RV0U/bPhD+V9D9fm9LE1OEBH5jE5qQtgkN3hBCXnI0ZklsfJd1VZX/fbKTNCkpA6nd",
	"U1z7vrvv++y7riE1pTUVVkwg10BpjqUKyxu9qHT1HZ9rJPYb+FuVtsCwLJUuQHbfCKwiWhqXgRyWTQTW",
	"GYuONdIItAZeWQQJxE5XixC3gU8OmwgcPtfaYQbyblrwPuoR5scTpuzTeea1fR/zSpUeHD5762iTHUTg",
	"S0K7dN4q+vlCXupQMWYPmWJ8YB20zcVczMTx7OT4dn4iT8/l6fkHcS6FgAhYswd234nUHelG5f6a+dG4",
	"UjFI8NhZwEZTY7r6b7ky5dFDX/OFXhiTKVYg7/Z3aO8M9xOXW25r0Ixl2Pjf4SNI+C8ZmjPpOjMJl95s",
	"VCvn1GriV8g4tcaH6erRhKs1Fas0tEfXA5edW00EGVLqtGVtqtFBBIVOsSIcgS6sSnM8msdeae18a+XM",
	"lmSSLJfLWIXj2LhF0mEp+XL16fLbzeVsHos457IYvYO+1NHF9RVE8AsdtRSOYxELH2gsVspqkHAStnyL",
	"cB5cS8rVA/dXv8AgzdusvIwr39Ofkb+u2tfh/SJrPCEfNxeiNwWrgFTWFjoN2OSJPIl+Or7nhqg1e9vH",
	"3CzQ7zcRJBSmq09lDe2gem2I2wkM7dUi8UeTrQ7Gcnu872BL4+rtw2JXY7PbuW1sHzFWW9u31db2H6od",
	"/hJeUbupvqfazSN8XezwCA+vtZ0QU4me1gEENhEQOt+aYZ4OPS+TpDCpKnJDLM/EmYDGjyC1aAO3U9aE",
	"bvj7Db+a++ZPAAAA//8Ojh68jwgAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

