.SILENT:
.DEFAULT_GOAL := run-fast

include ./.env

export DB_DATABASE
export DB_DRIVER
export DB_HOSTNAME
export DB_PASSWORD
export DB_PORT
export DB_SSLMODE
export DB_USERNAME

export DB_LOCAL_DIR

export HTTP_SERVER_LISTEN_IP_ADDRESS
export HTTP_SERVER_LISTEN_PORT

export POSTGRES_VER
export ALPINE_VER

export SERVER_APP_NAME

export GIN_MODE

export PASSWORD_HASH_SALT

.PHONY: build-local
build-local:
	go build -o ./build/${SERVER_APP_NAME} ./cmd/${SERVER_APP_NAME}

.PHONY: run-local
run-local:
	./build/${SERVER_APP_NAME}

.PHONY: run-fast
run-fast:
	go run ./cmd/${SERVER_APP_NAME}

.PHONY: db
db:
	docker run \
	-d \
	-e PGDATA=/data \
	-e POSTGRES_USER=${DB_USERNAME} \
	-e POSTGRES_PASSWORD=${DB_PASSWORD} \
	-e POSTGRES_DB=${DB_DATABASE} \
	-v ${DB_LOCAL_DIR}:/data \
	-p ${DB_PORT}:${DB_PORT} \
	--rm \
	--name pastebin-db \
	postgres:${POSTGRES_VER}-alpine${ALPINE_VER}

.PHONY: db-stop
db-stop:
	docker stop pastebin-db

.PHONY: swag
swag:
	swag init -g ./cmd/${SERVER_APP_NAME}/main.go -o ./docs
