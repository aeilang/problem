.PHONY: migrate-init
migrate-init:
	migrate create -ext sql -dir ./internal/repo/migration -seq init

.PHONY: migrate-up
migrate-up:
	migrate -path ./internal/repo/migration -database postgres://lang:password@localhost:5432/problem?sslmode=disable up

.PHONY: migrate-down
migrate-down:
	migrate -path ./internal/repo/migration -database postgres://lang:password@localhost:5432/problem?sslmode=disable down

.PHONY: migrate-drop
migrate-drop:
	migrate -path ./internal/repo/migration -database postgres://lang:password@localhost:5432/problem?sslmode=disable drop -f	

.PHONY: set-postgres
set-postgres:
	docker run -d \
 	 -e POSTGRES_USER=lang \
 	 -e POSTGRES_PASSWORD=password \
	 -e POSTGRES_DB=problem \
  	 --name my-postgres \
	 -p 5432:5432 \
	 -d \
  	 postgres:alpine

