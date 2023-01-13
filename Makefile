run: ### Run docker-compose
	docker-compose up --build -d webapp
.PHONY: run

run-postgres: ### Run docker-compose only with postgres
	docker-compose up -d postgres
.PHONY: run-postgres

run-dev: ### Run docker-compose
	make run-postgres
	go run ./cmd/app/main.go
.PHONY: run

integration-tests:
	docker-compose -f docker-compose.integration.yaml up --build postgres webapp
.PHONY: integration-tests

stop: ### Down docker-compose
	docker-compose down
.PHONY: stop

init-swag: ### Init swag
	swag init -g internal/endpoints/http/router.go
.PHONY: init-swag