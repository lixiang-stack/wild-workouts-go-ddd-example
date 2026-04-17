package main

//go:generate oapi-codegen -generate types -o openapi_types.gen.go -package main ../../api/openapi/trainings.yml
//go:generate oapi-codegen -generate chi-server -o openapi_api.gen.go -package main ../../api/openapi/trainings.yml
