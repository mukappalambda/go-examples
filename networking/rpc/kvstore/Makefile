GO ?= go

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

clean: ## clean binaries
	@echo "$(WHALE) $@"
	@rm -rf ./bin

build: ## build the cmd
	@echo "$(WHALE) $@"
	@$(GO) build -o bin/kvstore-server ./cmd/kvstore-server
	@$(GO) build -o bin/kvstore-cli ./cmd/kvstore-cli
