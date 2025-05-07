GO ?= go

APP_NAME=$$(basename $$(pwd))

GO_LDFLAGS=-ldflags '-s -w'

help: ## Show help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

clean: ## Clean binaries
	@echo "$(WHALE) $@"
	@${GO} clean
	@rm -f ${APP_NAME}

build: clean ## Build the binary
	@echo "$(WHALE) $@"
	@go build -o ${APP_NAME} ${GO_LDFLAGS} .

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
	@./${APP_NAME}
