include .env
export

.PHONY: run build migrate-up migrate-down migrate-reset migrate-status docker-up docker-down tidy

run:
	go run ./cmd/api

build:
	go build -o bin/api ./cmd/api

migrate-up:
	go run ./cmd/migrate -command=up

migrate-down:
	go run ./cmd/migrate -command=down

migrate-reset:
	go run ./cmd/migrate -command=reset

migrate-status:
	go run ./cmd/migrate -command=status

docker-up:
	docker compose up -d

docker-down:
	docker compose down

tidy:
	go mod tidy
