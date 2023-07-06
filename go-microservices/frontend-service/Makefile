.PHONY: build, run

build:
	docker build -t frontend-service .

run:
	docker run --name frontend-service -d -p 8081:80 frontend-service