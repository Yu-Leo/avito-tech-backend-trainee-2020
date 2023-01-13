run: ### Run docker-compose
	docker-compose up --build -d webapp
.PHONY: run

run-postgres: ### Run docker-compose only with postgres
	docker-compose up postgres
.PHONY: run-postgres

stop: ### Down docker-compose
	docker-compose down
.PHONY: stop

init-swag: ### Init swag
	swag init -g internal/endpoints/http/router.go
.PHONY: init-swag