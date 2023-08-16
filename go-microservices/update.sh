#!/bin/bash

SERVICE_NAME=$1
IMAGE_TAG=$2

echo "Updating $SERVICE_NAME Docker Service to use $IMAGE_TAG..."
docker service update --image $IMAGE_TAG $SERVICE_NAME
echo "Update Successul!"
