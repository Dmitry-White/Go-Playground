.PHONY: build, run, inside, kill

build:
	docker build -t broker-service .

run:
	docker run --name broker-service -d -p 8081:80 broker-service

inside:
	docker exec -it broker-service sh

kill:
	docker container stop broker-service && docker container rm broker-service
