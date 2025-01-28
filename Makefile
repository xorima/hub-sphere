.PHONY: fmt build lint test setup

GOBIN ?= $$(go env GOPATH)/bin

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
	make check-coverage


install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

check-coverage:  install-go-test-coverage
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./... -race
	${GOBIN}/go-test-coverage --config=./.testcoverage.yml

all: fmt build lint test
