.PHONY: build, run, inside, kill, production

summon:
	./deploy.sh <AWS_PROFILE> up

dismiss:
	./deploy.sh <AWS_PROFILE> down

inside:
	ssh -i "~/.ssh/id_rsa" ubuntu@<IP>
