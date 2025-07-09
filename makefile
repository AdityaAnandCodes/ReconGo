BINARY_NAME = recongo
BUILD_DIR = builds

.PHONY: all build clean run install

## Build for current OS
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

## Build for Linux
build-linux:
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux main.go

## Build for macOS (ARM)
build-mac:
	mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-mac main.go

## Clean builds
clean:
	rm -rf $(BUILD_DIR)

## Run locally
run:
	go run main.go

## Install globally
install:
	go install

## Help
help:
	@echo "Makefile options:"
	@echo "  make build        Build for current OS"
	@echo "  make build-linux  Build for Linux"
	@echo "  make build-mac    Build for macOS"
	@echo "  make clean        Remove builds"
	@echo "  make run          Run with go run"
	@echo "  make install      Install globally"
