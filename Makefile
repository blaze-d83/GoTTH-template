# Variables
BINARY_NAME=blog
SRC_DIR=cmd
MAIN=$(SRC_DIR)/main.go
BIN_DIR=bin
ENV_FILE=.env

# Default target to build the application
build: 
	@echo "Building the Go application..."
	@if not exist $(BIN_DIR) (mkdir $(BIN_DIR))
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN)

# Run the application
run: build
	@echo "Running the Go application..."
	./$(BIN_DIR)/$(BINARY_NAME)

# Tidy up the dependencies
tidy:
	@echo "Tidying up Go modules..."
	go mod tidy

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download

# Clean the directory
clean:
	@echo "Cleaning up..."
	@if exist $(BIN_DIR) (rmdir /S /Q $(BIN_DIR))

# Format the Go code
format:
	@echo "Formatting Go code..."
	go fmt ./...

# Check for issues with the code
lint:
	@echo "Linting Go code..."
	golangci-lint run

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Help menu
help:
	@echo "Makefile commands:"
	@echo "  build       - Build the Go application"
	@echo "  run         - Run the Go application"
	@echo "  tidy        - Tidy up Go modules"
	@echo "  deps        - Install dependencies"
	@echo "  clean       - Clean up generated files"
	@echo "  format      - Format the Go code"
	@echo "  lint        - Lint the Go code"
	@echo "  test        - Run tests"
	@echo "  help        - Display this help menu"

# Default target (help menu)
.DEFAULT_GOAL := help



