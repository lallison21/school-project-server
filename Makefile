.PHONY: build
build:
	go build -v ./cmd/to-do

.DEFAULT_GOAL := build
