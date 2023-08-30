#!/bin/bash

ENV=$1
COMPOSE_FILE=swarm.$1.yaml
STACK_NAME=go-microservices

echo "Initialising Docker Swarm..."
docker swarm init
echo "Initialisation Successful!"

echo "Deploying Docker Stack..."
docker stack deploy -c $COMPOSE_FILE $STACK_NAME
echo "Deployment Successful!"
