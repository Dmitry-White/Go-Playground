#!/bin/bash

SERVICE_NAME=$1
SERVICE_COUNT=$2

echo "Scaling $SERVICE_NAME Docker Service to $SERVICE_COUNT..."
docker service scale $SERVICE_NAME=$SERVICE_COUNT
echo "Scaling Successul!"
