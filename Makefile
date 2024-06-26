run: ## Run project on host machine
	go run cmd/main.go

build: ## Building the project
	go build cmd/main.go

test: ## Run all unit tests in the project
	go test -v ./...

test-cover: ## Run all unit tests in the project with test coverage
	go test -v ./... -covermode=count -coverprofile=coverage.out

html-cover: test-cover
	go tool cover -html="coverage.out"

## Documentation
SWAGGER_PARENT_INDEX_YAML_PATH := docs/swagger/index.yaml
SWAGGER_YAML_PATH              := docs/swagger/generated-swagger.yaml
POSTMAN_COLLECTION_PATH        := docs/swagger/generated-postman.json

swagger-postman-deps: # Installs required dependencies for swagger and postman
	@echo "Installing swagger and postman dependencies"
	npm i -g swagger-cli@4.0.4
	npm i -g openapi-to-postmanv2@4.21.0

generate-swagger-yaml:
	@swagger-cli bundle -t yaml $(SWAGGER_PARENT_INDEX_YAML_PATH) --outfile $(SWAGGER_YAML_PATH)

generate-postman-collection: generate-swagger-yaml
	@openapi2postmanv2 -s $(SWAGGER_YAML_PATH) -o $(POSTMAN_COLLECTION_PATH) -p
