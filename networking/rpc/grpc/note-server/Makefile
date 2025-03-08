GO ?= go

clean: ## clean binaries
	@echo "$(WHALE) $@"
	@rm -rf ./bin

protoc: ## gen proto files
	@echo "$(WHALE) $@"
	@protoc --proto_path=. --proto_path=$$HOME/protobuf/src --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./api/services/notes/notes.proto

build: ## build the cmd
	@echo "$(WHALE) $@"
	@$(GO) build -o bin/noted ./cmd/noted
	@$(GO) build -o bin/note-cli ./cmd/note-cli

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort
