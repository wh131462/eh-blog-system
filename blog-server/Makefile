.PHONY: run build docker

run:
    go run cmd/server/main.go

build:
    GOOS=linux GOARCH=amd64 go build -o bin/server ./cmd/server

docker:
    docker build -t blog-server:latest -f deployments/Dockerfile .

migrate:
    go run migrations/main.go

test:
    go test -v ./...