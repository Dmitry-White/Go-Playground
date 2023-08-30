#!/bin/bash

STACK_NAME=go-microservices

echo "Removing Docker Stack..."
docker stack rm $STACK_NAME
echo "Removal Successful!"

echo "Leaving Docker Swarm..."
docker swarm leave --force
echo "Leave Successful!"
