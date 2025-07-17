# Project settings
APP_NAME=boilerplate-simple
CMD_DIR=cmd/server
PORT=8080

# Load .env variables
include .env
export

.PHONY: run build tidy fmt test migrate help

## Run the application
run:
	@echo "🚀 Running $(APP_NAME)..."
	go run $(CMD_DIR)/main.go

## Build the application
build:
	@echo "📦 Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) $(CMD_DIR)/main.go

## Tidy up Go modules
tidy:
	go mod tidy

## Format code
fmt:
	go fmt ./...

## Run tests
test:
	go test ./... -v

migrate:
	@echo "📦 Running database migration via CLI args..."
	docker exec -it boilerplate-simple-app go run cmd/server/main.go migrate

docker-up:
	@echo "Building Docker image for development..."
	docker compose \
	-f deploy/docker/docker-compose.yml \
	--env-file .env \
	--project-name boilerplate-simple \
	up --build

docker-down:
	@echo "Stopping Docker containers..."
	docker compose \
	-f deploy/docker/docker-compose.yml \
	--env-file .env \
	--project-name boilerplate-simple \
	down -v

## Show help
help:
	@echo "Available commands:"
	@echo "  run       - Run the application"
	@echo "  build     - Build the binary"
	@echo "  tidy      - Clean go.mod"
	@echo "  fmt       - Format code"
	@echo "  test      - Run tests"
	@echo "  migrate   - Run database migrations"
