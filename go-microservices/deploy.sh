#!/bin/bash

COMPOSE_FILE=swarm.yaml
STACK_NAME=go-microservices

echo "Initialising Docker Swarm..."
docker swarm init
echo "Initialisation Successul!"

echo "Deploying Docker Stack..."
docker stack deploy -c $COMPOSE_FILE $STACK_NAME
echo "Deployment Successul!"
