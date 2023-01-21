# Documentation

init-swag: ### Init OpenAPI Specification files
	swag init -g internal/endpoints/rest/router.go
.PHONY: init-swag


# Developer mode

dev-up-all: ### Run Postgres and webapp services in Docker containers
	docker-compose -f docker-compose.dev.yaml up --build -d postgres webapp
.PHONY: dev-up-all

dev-up-postgres: ### Run only Postgres in Docker container
	docker-compose -f docker-compose.dev.yaml up -d postgres
.PHONY: dev-up-postgres

dev-init-db: ### Init database in Postgres Docker container
	docker-compose -f docker-compose.dev.yaml up -d postgres
	sleep 2
	docker-compose -f docker-compose.dev.yaml up --build init-db
	docker-compose -f docker-compose.dev.yaml down
.PHONY: dev-init-db

dev-down: ### Stop and delete all running containers
	docker-compose -f docker-compose.dev.yaml down
.PHONY: dev-down


# End-to-end tests

end-to-end-tests-up:
	docker-compose -f docker-compose.e2e.yaml up -d postgres
	sleep 2
	docker-compose -f docker-compose.e2e.yaml up --build init-db
	sleep 1
	docker-compose -f docker-compose.e2e.yaml up -d webapp
	sleep 1
	docker-compose -f docker-compose.e2e.yaml up --build end-to-end-tests
	sleep 2
	docker-compose -f docker-compose.e2e.yaml down
.PHONY: end-to-end-tests-up

end-to-end-tests-forced-down:
	docker-compose -f docker-compose.e2e.yaml down
.PHONY: end-to-end-tests-forced-down
