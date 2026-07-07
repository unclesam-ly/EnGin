.PHONY: run generate build

run:
	go run cmd/server/*.go server

generate:
	go generate ./internal/ent

build:
	go build -o bin/server cmd/server/*.go
