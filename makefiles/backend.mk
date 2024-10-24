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
else
	GO_FLAGS += -gcflags="all=-N -l"
endif

# Flags for the dlv command
DLV_FLAGS=--log

# Include env vars in local.env
ifneq (,$(wildcard local.env))
    include local.env
    export $(shell sed 's/=.*//' local.env)
endif

###
### LOCALLY RUN THE PROJECT
###
# Run the main binary
run-main:
	@echo "Running $(BINARY_NAME_MAIN)..."
	@dlv debug $(SRC_MAIN) $(DLV_FLAGS)

# Run the main binary
run-redirector:
	@echo "Running $(BINARY_NAME_REDIRECTOR)..."
	@dlv debug $(SRC_REDIRECTOR) $(DLV_FLAGS)

###
### RUN DOCKER WITH DEV ENV
###
build-docker-main: apply-env
	@docker build -f Dockerfile.main -t shrtnr_main:$(IMAGE_TAG) .

build-docker-redir: apply-env
	@docker build -f Dockerfile.redirector -t shrtnr_redir:$(IMAGE_TAG) .

deploy-docker-dev: build-docker-main build-docker-redir
	@$(shell export IMAGE_TAG=$(IMAGE_TAG))
	@docker stack deploy -c docker-compose.dev.yml url-shrtnr -d

stop-docker-dev:
	@docker service rm url-shrtnr_shortener-db \
	url-shrtnr_shortener-main-app url-shrtnr_shortener-redir-app

###
### RUN DOCKER WITH PROD ENV
###
build-docker-prod: clean apply-env
	@docker build -f Dockerfile.prod.main -t rbennum2329/shrtnr_main:$(IMAGE_TAG) \
	--progress=plain .
	@docker push rbennum2329/shrtnr_main:$(IMAGE_TAG)
	@docker build -f Dockerfile.prod.redirector -t rbennum2329/shrtnr_redir:$(IMAGE_TAG) \
	--progress=plain .
	@docker push rbennum2329/shrtnr_redir:$(IMAGE_TAG)

deploy-docker-prod: apply-env
	@$(shell export IMAGE_TAG=$(IMAGE_TAG))
	@docker stack deploy -c docker-compose.prod.yml url-shrtnr -d

###
### MISC
###
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