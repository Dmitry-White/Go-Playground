#!/bin/bash

ACCOUNT=dmitrywhite
NAME=reverse-proxy-production
VERSION=1.0.0
TAG=$ACCOUNT/$NAME:$VERSION
DOCKER_FILE=production.dockerfile

echo "Tag: $TAG"

echo "Logging in Docker Hub..."
docker login
echo "Logging in Successful!"

echo "Building Docker image..."
docker build -t $TAG -f $DOCKER_FILE .
echo "Building Successful!"

echo "Pushing Docker image..."
docker push $TAG
echo "Pushing Successful!"
