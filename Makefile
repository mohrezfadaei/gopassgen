BINARY_NAME=passgen
SRC_DIR=./cmd
MAIN_SRC=$(SRC_DIR)/main.go
BUILD_DIR=./build/bin
BINARY_SRC=$(BUILD_DIR)/$(BINARY_NAME)

all: build

build:
		@echo "Building the binary..."
		go build -o $(BINARY_SRC) $(MAIN_SRC)
		chmod +x $(BINARY_SRC)

run: build
		@echo "Running the binary..."
		./$(BINARY_SRC)

clean:
		@echo "Cleaning up..."
		@if [ -f $BINARY_SRC ]; then rm $(BINARY_SRC); fi

deps:
		@echo "Installing dependencies..."
		go mod tidy

fmt:
		@echo "Formatting the code..."
		go fmt ./...

lint:
		@echo "Linting the code..."
		golangci-lint run

test:
		@echo "Running tests..."
		go test -v ./...

help:
		@echo "Makefile for $(BINARY_NAME)"
		@echo ""
		@echo "Usage:"
		@echo "  make [target]"
		@echo ""
		@echo "Targets:"
		@echo "  all		- Build the binary (default)"
		@echo "  build		- Build the binary"
		@echo "  run		- Run the application"
		@echo "  clean		- Clean the binary"
		@echo "  deps		- Install dependencies"
		@echo "  fmt		- Format the code"
		@echo "  lint		- Lint the code"
		@echo "  help		- Display help message"

.PHONY: all build run clean deps fmt line test help
