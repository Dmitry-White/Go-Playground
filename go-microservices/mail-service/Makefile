.PHONY: build, run, inside, kill

build:
	docker build -t mail-service .

run:
	docker run --name mail-service -d -p 8084:80 mail-service

inside:
	docker exec -it mail-service sh

kill:
	docker container stop mail-service && docker container rm mail-service
