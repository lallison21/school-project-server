.PHONY: build up down
build:
	go build -v ./cmd/to-do

up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down

.DEFAULT_GOAL := build
