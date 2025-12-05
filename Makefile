.PHONY: help build run dev test clean migrate seed docker-up docker-down

# Default target
help:
	@echo "Available commands:"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make dev          - Run with hot reload (requires air)"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make migrate      - Run database migrations"
	@echo "  make seed         - Seed the database"
	@echo "  make deps         - Download dependencies"
	@echo "  make swagger      - Generate swagger docs"

# Build the application
build:
	@echo "Building..."
	go build -o bin/server main.go

# Run the application
run: build
	@echo "Running..."
	./bin/server

# Run with hot reload (requires air: go install github.com/cosmtrek/air@latest)
dev:
	@echo "Running with hot reload..."
	air

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# Run database migrations
migrate:
	@echo "Running migrations..."
	@mysql -u${DB_USER} -p${DB_PASSWORD} -h${DB_HOST} ${DB_NAME} < migrations/001_init_schema.sql

# Generate swagger documentation (requires swag: go install github.com/swaggo/swag/cmd/swag@latest)
swagger:
	@echo "Generating swagger docs..."
	swag init

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	golangci-lint run

# Create .env from example
env:
	@if [ ! -f .env ]; then cp .env.example .env; echo ".env file created"; else echo ".env file already exists"; fi

# Create uploads directory
uploads:
	@mkdir -p uploads/images uploads/documents uploads/general
	@echo "Upload directories created"

# Initialize project
init: env deps uploads
	@echo "Project initialized successfully"

# Build for Linux
build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build -o bin/server-linux main.go

# Build for Windows
build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -o bin/server.exe main.go

# Build for macOS
build-mac:
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 go build -o bin/server-mac main.go
