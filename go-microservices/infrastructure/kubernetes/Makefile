.PHONY: clean, init, start, stop, see, summon, dismiss

clean:
	minikube delete

init:
	minikube start --nodes=2 

start:
	minikube start

stop:
	minikube stop

see:
	minikube dashboard

summon:
	./deploy.sh nil up

dismiss:
	./deploy.sh nil down
