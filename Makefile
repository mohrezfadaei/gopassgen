BINARY_NAME=passgen
SRC_DIR=./cmd
MAIN_SRC=$(SRC_DIR)/main.go
BUILD_DIR=./build/bin
BINARY_SRC=$(BUILD_DIR)/$(BINARY_NAME)

all: build

build:
		@echo "Building the binary for the current OS and architecture..."
		go build -o $(BINARY_SRC) $(MAIN_SRC)
		chmod +x $(BINARY_SRC)

build-windows:
		@echo "Building the binary for Windows..."
		GOOS=windows GOARCH=amd64 go build -o $(BINARY_SRC)-windows-amd64.exe $(MAIN_SRC)
		GOOS=windows GOARCH=386 go build -o $(BINARY_SRC)-windows-386.exe $(MAIN_SRC)

build-macos:
		@echo "Building the binary for macOS..."
		GOOS=darwin GOARCH=amd64 go build -o $(BINARY_SRC)-darwin-amd64 $(MAIN_SRC)
		GOOS=darwin GOARCH=arm64 go build -o $(BINARY_SRC)-darwin-arm64 $(MAIN_SRC)

build-linux:
		@echo "Build the binary for Linux..."
		GOOS=linux GOARCH=amd64 go build -o $(BINARY_SRC)-linux-amd64 $(MAIN_SRC)
		GOOS=linux GOARCH=386 go build -o $(BINARY_SRC)-linux-386 $(MAIN_SRC)
		GOOS=linux GOARCH=arm64 go build -o $(BINARY_SRC)-linux-arm64 $(MAIN_SRC)
		GOOS=linux GOARCH=arm go build -o $(BINARY_SRC)-linux-arm $(MAIN_SRC)

build-all: build-windows build-darwin build-linux

run: build
		@echo "Running the binary..."
		./$(BINARY_SRC)

clean:
		@echo "Cleaning up..."
		@if [ -f $(BINARY_SRC) ]; then rm $(BINARY_SRC); fi

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
