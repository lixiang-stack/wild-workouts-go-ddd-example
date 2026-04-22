include .env

.PHONY: help
help: ## Show available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

.PHONY: generate
generate: ## Generate OpenAPI Go code (types + chi-server)
	cd internal/trainer && go generate ./...
	cd internal/trainings && go generate ./...
	cd internal/users && go generate ./...

.PHONY: proto
proto: ## Generate protobuf/gRPC Go code (requires protoc)
	cd internal/common && go generate ./...

.PHONY: openapi_js
openapi_js: ## Generate JavaScript API clients (requires Docker)
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.0 generate \
        -i /local/api/openapi/trainings.yml \
        -g javascript \
        -o /local/web/src/repositories/clients/trainings

	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.0 generate \
		-i /local/api/openapi/trainer.yml \
		-g javascript \
		-o /local/web/src/repositories/clients/trainer

	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.0 generate \
		-i /local/api/openapi/users.yml \
		-g javascript \
		-o /local/web/src/repositories/clients/users

.PHONY: clean
clean: ## Remove generated OpenAPI Go code
	rm -f internal/trainer/openapi_*.gen.go
	rm -f internal/trainings/openapi_*.gen.go
	rm -f internal/users/openapi_*.gen.go

# .PHONY: clean-all
# clean-all: clean ## Remove all generated code (OpenAPI + protobuf + JS clients)
# 	rm -f internal/common/genproto/trainer/*.pb.go
# 	rm -f internal/common/genproto/users/*.pb.go
# 	rm -rf web/src/repositories/clients/trainer
# 	rm -rf web/src/repositories/clients/trainings
# 	rm -rf web/src/repositories/clients/users

.PHONY: lint
lint: ## Run linter on all services
	@./scripts/lint.sh trainer
	@./scripts/lint.sh trainings
	@./scripts/lint.sh users


.PHONY: mycli
mycli:
	mycli -u ${MYSQL_USER} -p ${MYSQL_PASSWORD} ${MYSQL_DATABASE}