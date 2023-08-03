.PHONY: build, run, inside, kill

build:
	docker build -t listener-service .

run:
	docker run --name listener-service -d -p 8085:80 listener-service

inside:
	docker exec -it listener-service sh

kill:
	docker container stop listener-service && docker container rm listener-service
