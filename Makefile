BINARY := pkgview
BUILD_DIR := build
SRC := ./main.go

.PHONY: all build run install fmt tidy clean

all: build

# Build the binary into build/ directory
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY) $(SRC)

# Run the built binary
run: build
	$(BUILD_DIR)/$(BINARY)

# Install the binary to /usr/bin with proper permissions
install: build
	sudo install -Dm755 $(BUILD_DIR)/$(BINARY) /usr/bin/$(BINARY)

# Format Go code
fmt:
	go fmt ./...

# Tidy go.mod and go.sum
tidy:
	go mod tidy

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)
