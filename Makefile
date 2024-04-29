.SILENT:
.DEFAULT_GOAL: fast-run

include ./.env

export ENABLE_DEBUG_LOGS
export ENABLE_HTTP_SERVER_DEBUG_MODE
export HTTP_SERVER_LISTEN_IP
export HTTP_SERVER_LISTEN_PORT

.PHONY: fast-run
fast-run:
	go run ./cmd/server

.PHONY: build
build:
	go build -o ./build/server ./cmd/server/

.PHONY: run
run:
	./build/server

.PHONY: swag
swag:
	swag init -g ./cmd/server/main.go -o ./docs
