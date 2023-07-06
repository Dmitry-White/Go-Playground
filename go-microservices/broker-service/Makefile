.PHONY: build, run

build:
	docker build -t broker-service .

run:
	docker run --name broker-service -d -p 8080:80 broker-service