.PHONY: build
build:
	go build -v ./cmd

.PHONY: test
test:
	go test -v -race -timeout 30s ./task1
.DEFAULT_GOAL := build