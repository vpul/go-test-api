# Makefile

# Variables
GO := go
BINARY_NAME := go-test-api

# Targets
all: build

build:
	$(GO) build -o ./bin/$(BINARY_NAME) ./cmd/api

run:
	$(GO) run ./cmd/api

start:
	./bin/$(BINARY_NAME)

clean:
	rm -f ./bin/$(BINARY_NAME)

.PHONY: all build run clean
