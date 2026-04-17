include .env
export

DB_URL ?= $(DATABASE_URL)
MIGRATE = migrate -path ./migrations -database "$(DB_URL)"

.PHONY: run build migrate-up migrate-down migrate-create docker-up docker-down

run:
	go run ./cmd/api

build:
	go build -o bin/api ./cmd/api

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

migrate-create:
	$(MIGRATE) create -ext sql -dir ./migrations -seq $(name)

docker-up:
	docker compose up -d

docker-down:
	docker compose down

tidy:
	go mod tidy
