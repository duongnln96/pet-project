DB_URL=postgres://postgres:postgres@localhost:5432/blog_realworld_db?sslmode=disable

migration_up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migration_down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc_gen:
	sqlc generate

buf_gen:
	rm -rf ./gen/go/* && \
	buf generate && \
	go mod tidy
.PHONY: buf_gen

clean:
	go clean

wire:
	cd internal/user/adapter/http_server && wire && cd - && \
	cd internal/auth/adapter/grpc_server && wire && cd -
.PHONY: wire

up_infras_env:
	docker compose -f docker-compose-infras.yaml up -d
.PHONY: up_core_env

down_infras_env:
	docker compose -f docker-compose-infras.yaml down
.PHONY: down_core_env

up_svc_env:
	docker compose -f docker-compose-service.yaml up -d
.PHONY: up_core_env

down_svc_env:
	docker compose -f docker-compose-service.yaml down
.PHONY: down_core_env