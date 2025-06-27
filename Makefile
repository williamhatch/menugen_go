.PHONY: dev build test clean install deps

# Default target
dev: deps
	buffalo dev

# Install dependencies
deps:
	go mod tidy
	yarn install

# Build the application
build:
	buffalo build

# Run tests
test:
	buffalo test

# Clean generated files
clean:
	rm -rf bin/
	rm -rf public/assets/
	rm -rf uploads/*

# Install tools
install:
	go install github.com/gobuffalo/cli/cmd/buffalo@latest

# Start production server
start:
	./bin/menugen

# Build for production
build-prod: deps
	yarn build
	buffalo build

# Help
help:
	@echo "Available targets:"
	@echo "  dev        - Start development server with hot reload"
	@echo "  build      - Build the application"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean generated files"
	@echo "  deps       - Install dependencies"
	@echo "  install    - Install Buffalo CLI"
	@echo "  start      - Start production server"
	@echo "  help       - Show this help message" 