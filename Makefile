# Variables
BINARY_NAME_MAIN=main_binary
BINARY_NAME_REDIRECTOR=redirector_binary
SRC_MAIN=./cmd/app
SRC_REDIRECTOR=./cmd/redirector
BUILD_DIR=./build
GOFILES=$(shell find . -type f -name '*.go')

# Default environment (can be overridden by setting ENV=prod, etc.)
ENV=dev

# Flags for the Go command
GO_FLAGS=
ifeq ($(ENV),prod)
	GO_FLAGS += -ldflags="-s -w"
endif

# Tasks
# .PHONY: all build clean run test

# Build the main binary
build-main:
	@echo "Building $(BINARY_NAME_MAIN)..."
	@go build $(GO_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME_MAIN) ./$(SRC_MAIN)
	@chmod 755 $(BUILD_DIR)/$(BINARY_NAME_MAIN)

# Run the main binary
run-main: build-main
	@echo "Running $(BINARY_NAME_MAIN)..."
	@./$(BUILD_DIR)/$(BINARY_NAME_MAIN)

# Build the redirector binary
build-redirector:
	@echo "Building $(BINARY_NAME_REDIRECTOR)..."
	@go build $(GO_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME_REDIRECTOR) ./$(SRC_REDIRECTOR)
	@chmod 755 $(BUILD_DIR)/$(BINARY_NAME_REDIRECTOR)

# Run the main binary
run-redirector: build-redirector
	@echo "Running $(BINARY_NAME_REDIRECTOR)..."
	@./$(BUILD_DIR)/$(BINARY_NAME_REDIRECTOR)

run-all: run-main run-redirector

apply-env:
	@if [ -z "$(IMAGE_TAG)" ]; then \
		echo "ERROR: IMAGE_TAG is not set."; \
		exit 1; \
	fi
	@if [ -z "$(PORT_DB)" ]; then \
		echo "ERROR: PORT_DB is not set."; \
		exit 1; \
	fi
	@if [ -z "$(PORT_REDIR)" ]; then \
		echo "ERROR: PORT_REDIR is not set."; \
		exit 1; \
	fi
	@if [ -z "$(PORT_MAIN)" ]; then \
		echo "ERROR: PORT_MAIN is not set."; \
		exit 1; \
	fi

build-docker-main: apply-env
	@docker build -f Dockerfile.main -t shrtnr_main:$(IMAGE_TAG) .

build-docker-redir: apply-env
	@docker build -f Dockerfile.redirector -t shrtnr_redir:$(IMAGE_TAG) .

deploy-docker-dev: build-docker-main build-docker-redir
	@$(shell export IMAGE_TAG=$(IMAGE_TAG))
	@docker stack deploy -c docker-compose.dev.yml url-shrtnr -d

# Start debugging using Delve
debug-main: build-main
	@echo "Start debugging $(BINARY_NAME_MAIN)"
	@dlv exec ./$(BUILD_DIR)/$(BINARY_NAME_MAIN)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)/*

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Watch files and re-build on changes (requires entr or similar tool)
# watch:
# 	@find . -name "*.go" | entr -c make run

# Advanced features
deps:
	@echo "Getting dependencies..."
	@go mod tidy

vet:
	@echo "Running go vet..."
	@go vet ./...

fmt:
	@echo "Formatting code..."
	@go fmt ./...

lint: vet fmt
	@echo "Running linter..."
	@golangci-lint run

# Run in different environments
# dev: ENV=dev
# dev: run

# prod: ENV=prod
# prod: run
