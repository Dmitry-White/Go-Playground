#!/bin/bash

ACCOUNT=$1
NAME=$2
VERSION=$3
TAG=$ACCOUNT/$NAME:$VERSION

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
