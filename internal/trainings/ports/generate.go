package ports

//go:generate oapi-codegen -generate types -o openapi_types.gen.go -package ports ../../../api/openapi/trainings.yml
//go:generate oapi-codegen -generate chi-server -o openapi_api.gen.go -package ports ../../../api/openapi/trainings.yml
