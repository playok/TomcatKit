.PHONY: build run clean test

APP_NAME=tomcatkit
BUILD_DIR=bin

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/tomcatkit

run: build
	./$(BUILD_DIR)/$(APP_NAME)

clean:
	rm -rf $(BUILD_DIR)
	go clean

test:
	go test -v ./...

deps:
	go mod tidy

.DEFAULT_GOAL := build
