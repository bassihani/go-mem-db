# Variables
BINARY_NAME=go-memdb
GO_FILES=$(shell find . -name "*.go")

.PHONY: all build run test clean help


all: build ## all: Default target to build the binary

build: ## build: Compile the project into a binary in the bin/ folder
	@echo "Building the binary..."
	@go build -o bin/$(BINARY_NAME) ./cmd/server/main.go

run: build ## run: Build and immediately run the application
	@echo "Running the application..."
	@./bin/$(BINARY_NAME)

test: ## test: Run all unit tests
	@echo "Running Unit test....."
	@go test -v ./internal/db/ ./pkg/utils
test_race: ## test: Run all unit tests in the project with the race detector
	@echo "Running tests with race detector..."
	@go test -v -race ./...

clean: ## clean: Remove the compiled binary and build artifacts
	@echo "Cleaning up..."
	@rm -rf bin/
	@go clean -cache -modcache -testcache

help: ## help: Display this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'