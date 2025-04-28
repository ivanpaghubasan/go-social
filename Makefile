include .envrc
MIGRATIONS_PATH=./cmd/migrate/migrations

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down 1

.PHONY: migrate-reset
migrate-reset:
	@echo "Resetting database..."
	# Drop all tables (postgres version)
	@psql $(DB_ADDR) -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;" || true
	# Force clean if necessary
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) drop -f
	# Reapply all migrations
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go
