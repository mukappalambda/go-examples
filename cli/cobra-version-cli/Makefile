GO ?= go
GIT_VERSION := $(shell git --version | awk '{print $$3}')
GIT_COMMIT  := $(shell git rev-parse HEAD)
BUILD_DATE  := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
USERNAME := mukappalambda
REPO_NAME := go-examples
PROJECT_NAME := cobra-version-cli
REPO_URL := "https://github.com/${USERNAME}/${REPO_NAME}"
LDFLAGS := -X '${PROJECT_NAME}/cmd.GitVersion=$(GIT_VERSION)' \
	-X '${PROJECT_NAME}/cmd.GitCommit=$(GIT_COMMIT)' \
	-X '${PROJECT_NAME}/cmd.BuildDate=$(BUILD_DATE)' \
	-X '${PROJECT_NAME}/cmd.RepoUrl=$(REPO_URL)'

SHELL := /bin/bash

help: ## Show help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

build: ## Build the binary
	@echo "$(WHALE) $@"
	@${GO} build -ldflags "$(LDFLAGS) -s -w" -o ./bin/cobra-version-cli main.go
	@echo "✓ Built the binary under \"./bin\""
