GO ?= go

APP_NAME='gin-demo-server'

GO_LDFLAGS=-ldflags '-s -w'

help: ## Show help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

clean: ## Clean binaries
	@echo "$(WHALE) $@"
	@rm -rf bin

build: clean ## Build the binary
	@echo "$(WHALE) $@"
	@go build -o bin/${APP_NAME} ${GO_LDFLAGS} cmd/gin-demo-server/main.go

build-image: clean ## Build the docker image
	@echo "$(WHALE) $@"
	@docker build -t ${APP_NAME}:latest .

up-container: build-image ## Start up the docker container
	@echo "$(WHALE) $@"
	@docker run -dt --name ${APP_NAME} -p 8080:8080 ${APP_NAME}:latest

down-container: ## Tear down the docker container
	@echo "$(WHALE) $@"
	@docker rm -f ${APP_NAME}

run: build ## Run the binary
	@echo "$(WHALE) $@"
	@./bin/${APP_NAME}
