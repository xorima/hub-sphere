.PHONY: fmt build lint test setup

setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	pre-commit install
	mkdir bin

fmt:
	go fmt

build:
	go mod tidy
	go build -o bin/hub-sphere

lint:
	golangci-lint run

test:
	go test ./... -race

all: fmt build lint test
