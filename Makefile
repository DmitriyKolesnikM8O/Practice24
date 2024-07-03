.PHONY: build
build:
	go build -o build/ -v ./cmd/main.go

.DEFAULT_GOAL := build
