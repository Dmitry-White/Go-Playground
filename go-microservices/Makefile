.PHONY: clean, start, stop, prepare, summon_local, summon_remote, dismiss, infra_swarm_up, infra_swarm_down

clean:
	docker system prune -af

start:
	docker-compose up -d

stop:
	docker-compose down

prepare:
	./deliver.sh

summon_local:
	./deploy.sh development

summon_remote:
	./deploy.sh production

dismiss:
	./remove.sh

infra_swarm_up:
	./infrastructure.sh DmitryPC swarm up

infra_swarm_down:
	./infrastructure.sh DmitryPC swarm down
