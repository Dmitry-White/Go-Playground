.PHONY: build, run, inside, kill, production

build:
	docker build -t reverse-proxy .

run:
	docker run --name reverse-proxy -d --expose 80 --expose 443 reverse-proxy

inside:
	docker exec -it reverse-proxy sh

kill:
	docker container stop reverse-proxy && docker container rm reverse-proxy

production:
	docker build -t reverse-proxy-production -f production.dockerfile .
