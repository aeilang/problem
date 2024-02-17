.PHONY: migrate-init
migrate-init:
	migrate create -ext sql -dir ./internal/repo/migration -seq init

.PHONY: migrate-up
migrate-up:
	migrate -path ./internal/repo/migration -database postgres://lang:password@localhost:5432/problem?sslmode=disable up

.PHONY: migrate-down
migrate-down:
	migrate -path ./internal/repo/migration -database postgres://lang:password@localhost:5432/problem?sslmode=disable down

migrate-drop:
	migrate -path ./internal/repo/migration -database postgres://lang:password@localhost:5432/problem?sslmode=disable drop -f	