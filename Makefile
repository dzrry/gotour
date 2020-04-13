.PHONY: build
build:
	go build -v ./cmd

.PHONY: test1
test1:
	go test -v -race -timeout 30s ./task1

.PHONY: test4
test4:
	go test -v -race -timeout 30s ./task4

.DEFAULT_GOAL := build