.PHONY: build, run

build:
	docker build -t authentication-service .

run:
	docker run --name authentication-service -d -p 8082:80 authentication-service