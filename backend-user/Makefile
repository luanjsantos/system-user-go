include .env

MIGRATE=migrate -path db/migrations -database $(DATABASE_URL)

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-force:
	$(MIGRATE) force 1

migrate-new:
	@read -p "Nome da migration: " name; \
	migrate create -ext sql -dir db/migrations -seq $$name

run:
	go run cmd/main.go

dev:
	air

swagger:
	swag init -g cmd/main.go
	swag fmt

swagger-ui:
	@echo "Acesse: http://localhost:8080/swagger/index.html"
