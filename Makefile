include .env

SOURCE=cmd/main.go
MIGRATIONS=./migrations
DOCKER=./

all: docker-up

dbshell:
	@psql -U ${DB_USER} -d ${DB_NAME}

migrations-up:
	@goose -dir=${MIGRATIONS} postgres ${DB_URL} up

migrations-down:
	@goose -dir=${MIGRATIONS} postgres ${DB_URL} down 

local:
	@go run $(SOURCE)

docker-up: docker-build
	@docker-compose up

docker-build:
	@docker build -t project $(DOCKER)

tidy:
	@go mod tidy
