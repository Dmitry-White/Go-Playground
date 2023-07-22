.PHONY: build, run

build:
	docker build -t broker-service .

run:
	docker run --name broker-service -d -p 8081:80 broker-service

inside:
	docker exec -it broker-service sh
