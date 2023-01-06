run: ### Run docker-compose
	docker-compose up --build -d webapp
.PHONY: run

down: ### Down docker-compose
	docker-compose down
.PHONY: down

init-swag-v1: ### Init swag
	swag init -g internal/controller/http/v1/router.go
.PHONY: init-swag-v1