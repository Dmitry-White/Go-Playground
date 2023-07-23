.PHONY: build, run

build:
	docker build -t frontend-service .

run:
	docker run --name frontend-service -d -p 8080:80 frontend-service

inside:
	docker exec -it frontend-service sh

kill:
	docker container stop frontend-service && docker container rm frontend-service
