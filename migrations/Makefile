include .env

migrate-up:
	migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable" -path ./ -verbose up

migrate-down:
	migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable" -path ./ -verbose down

.PHONY: migrate-up migrate-down
