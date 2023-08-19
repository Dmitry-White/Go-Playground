.PHONY: clean, start, stop, prepare, summon, dismiss

clean:
	docker system prune -af

start:
	docker-compose up -d

stop:
	docker-compose down

prepare:
	./deliver.sh

summon:
	./deploy.sh

dismiss:
	./remove.sh
