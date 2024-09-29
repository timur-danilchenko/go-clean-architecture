SOURCE=cmd/main.go
DOCKER=./

all: docker-up

local:
	@go run $(SOURCE)

docker-up: docker-build
	@docker-compose up

docker-build:
	@docker build -t project $(DOCKER)

tidy:
	@go mod tidy