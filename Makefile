run: ### Run docker-compose
	docker-compose up --build -d webapp
.PHONY: run

run-postgres: ### Run docker-compose only with postgres
	docker-compose up --build -d postgres
.PHONY: run-postgres

stop: ### Down docker-compose
	docker-compose down
.PHONY: stop

init-swag-v1: ### Init swag
	swag init -g internal/controller/http/v1/router.go
.PHONY: init-swag-v1