.PHONY: build, run

build:
	docker build -t authentication-service .

run:
	docker run --name authentication-service -d -p 8082:80 authentication-service

inside:
	docker exec -it authentication-service sh

migrate:
	go run cmd/main.go
