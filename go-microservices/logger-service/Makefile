.PHONY: build, run, in

build:
	docker build -t logger-service .

run:
	docker run --name logger-service -d -p 8083:80 logger-service

inside:
	docker exec -it logger-service sh

kill:
	docker container stop logger-service && docker container rm logger-service
