.SILENT:
.DEFAULT_GOAL: run-fast

include ./.env

export ENABLE_DEBUG_LOGS
export ENABLE_HTTP_SERVER_DEBUG_MODE
export HTTP_SERVER_LISTEN_IP
export HTTP_SERVER_LISTEN_PORT

export DB_PASSWORD
export DB_HOSTNAME
export DB_MIGRATION_CONTAINER_PORT
export DB_USERNAME
export DB_DRIVER
export DB_DATABASE
export DB_SSLMODE

.PHONY: build
build:
	go build -o ./build/server ./cmd/server/

.PHONY: run
run:
	./build/server

.PHONY: run-fast
run-fast:
	go run ./cmd/server

.PHONY: migrate
migrate:
	./wait-for-postgres.sh \
	${DB_PASSWORD} \
	${DB_HOSTNAME} \
	${DB_PORT} \
	${DB_USERNAME} \
	migrate \
	-path ./schema \
	-database ${DB_DRIVER}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOSTNAME}:${DB_PORT}/${DB_DATABASE}?sslmode=${DB_SSLMODE} \
	up

.PHONY: swag
swag:
	swag init -g ./cmd/server/main.go -o ./docs

.PHONY: migrate-docker
migrate-docker: run-mdc migrate-docker-up stop-mdc

.PHONY: run-mdc
run-mdc:
	docker run \
	--rm \
	--name ${DB_MIGRATION_CONTAINER_NAME} \
	-d \
	-e PGDATA=/data \
	-e POSTGRES_USER=${DB_USERNAME} \
	-e POSTGRES_PASSWORD=${DB_PASSWORD} \
	-e POSTGRES_DB=${DB_DATABASE} \
	-v ${DB_LOCAL_DIR}:/data \
	-p ${DB_MIGRATION_CONTAINER_PORT}:5432 \
	postgres:${POSTGRES_VER}-alpine${ALPINE_VER}

.PHONY: migrate-docker-up
migrate-docker-up:
	./wait-for-postgres.sh \
	${DB_PASSWORD} \
	${DB_HOSTNAME} \
	${DB_MIGRATION_CONTAINER_PORT} \
	${DB_USERNAME} \
	migrate \
	-path ./schema \
	-database ${DB_DRIVER}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOSTNAME}:${DB_MIGRATION_CONTAINER_PORT}/${DB_DATABASE}?sslmode=${DB_SSLMODE} \
	up

.PHONY: stop-mdc
stop-mdc:
	docker stop ${DB_MIGRATION_CONTAINER_NAME}
