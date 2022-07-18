run-app:
	docker-compose up -d

stop-app:
	docker-compose down -v

clear-di: ## Clear project DI
	go mod tidy

test: ## Run go test ./...
	@go test -count=1 -v ./...

lint: ## Run golangci-lint
	@golangci-lint run

create-migration: ## Create new migration files
	MIGRATION_NAME=$(shell bash -c 'read -p "Migration name: " mname; echo $$mname') && \
	DT=$$(date '+%Y%m%d%H%M%S') && \
	touch "internal/infrastructure/repository/postgres/migrations/$${DT}_$${MIGRATION_NAME}.up.sql" && \
	touch "internal/infrastructure/repository/postgres/migrations/$${DT}_$${MIGRATION_NAME}.down.sql"