.PHONY: summon, dismiss, inside

summon:
	./deploy.sh $(AWS_PROFILE) up

dismiss:
	./deploy.sh $(AWS_PROFILE) down

inside:
	ssh -i "~/.ssh/id_rsa" ubuntu@$(NODE_IP)
