.PHONY: build up down
build:
	go build -o ./bin/school-project-server ./cmd/school-project-server

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down

.DEFAULT_GOAL := build
