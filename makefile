# Makefile

# Variables
GO := go
BINARY_NAME := go-test-api

# Targets
all: build

build:
	$(GO) build -o $(BINARY_NAME) ./cmd/api

run:
	$(GO) run ./cmd/api

start:
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

.PHONY: all build run clean
