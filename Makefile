run: ### Run docker-compose
	docker-compose -f docker-compose.dev.yaml up --build -d webapp
.PHONY: run

init-postgres:
	docker-compose -f docker-compose.dev.yaml up -d postgres
	sleep 5
	docker-compose -f docker-compose.dev.yaml up --build init-db
	sleep 5
	docker-compose -f docker-compose.dev.yaml down
.PHONY: init-postgres

run-postgres: ### Run docker-compose only with postgres
	docker-compose -f docker-compose.dev.yaml up -d postgres
.PHONY: run-postgres

run-dev: ### Run docker-compose
	make run-postgres
	go run ./cmd/app/main.go
.PHONY: run

integration-tests:
	docker-compose -f docker-compose.integration.yaml up --build integration
.PHONY: integration-tests

stop: ### Down docker-compose
	docker-compose -f docker-compose.dev.yaml down
.PHONY: stop

init-swag: ### Init swag
	swag init -g internal/endpoints/http/router.go
.PHONY: init-swag