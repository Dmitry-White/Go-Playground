#!/bin/bash

ACCOUNT=$1
NAME=$2:1.0.0
TAG=$ACCOUNT/$NAME

echo "Tag: $TAG"

echo "Logging in Docker Hub..."
docker login
echo "Logging in Successul!"

echo "Building Docker image..."
docker build -t $TAG .
echo "Building Successul!"

echo "Pushing Docker image..."
docker push $TAG
echo "Pushing Successul!"