.PHONY: run build clean setup

# Default target
all: build

# Get local IP address
IP := $(shell ifconfig | grep "inet " | grep -v 127.0.0.1 | awk '{print $$2}' | head -n 1)
PORT := 8080
URL := http://$(IP):$(PORT)

# Setup development environment
setup:
	@echo "Setting up development environment..."
	npm install
	npm run build:css

# Build the application
build: setup
	@echo "Building application..."
	go build -o ./tmp/main ./cmd/server

# Run the application with hot reload
run: setup
	@echo "Starting application with hot reload..."
	@mkdir -p tmp
	@echo "\n\033[1;32mApplication is running at:\033[0m"
	@echo "\033[1;34m$(URL)\033[0m"
	@echo "\n\033[1;33mPress Ctrl+C to stop\033[0m\n"
	air

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf tmp
	@rm -f air_errors.log
	@rm -rf node_modules
	@rm -f package-lock.json
	@rm -f assets/css/output.css

# Run the application without hot reload (for production)
run-prod: build
	@echo "Starting application..."
	@echo "\n\033[1;32mApplication is running at:\033[0m"
	@echo "\033[1;34m$(URL)\033[0m"
	@echo "\n\033[1;33mPress Ctrl+C to stop\033[0m\n"
	./tmp/main

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go install github.com/air-verse/air@latest
	npm install 