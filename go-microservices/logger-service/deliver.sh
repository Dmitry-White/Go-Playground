#!/bin/bash

ACCOUNT=$1
NAME=$2
VERSION=$3
TAG=$ACCOUNT/$NAME:$VERSION

echo "Tag: $TAG"

echo "Logging in Docker Hub..."
docker login
echo "Logging in Successful!"

echo "Building Docker image..."
docker build -t $TAG .
echo "Building Successful!"

echo "Pushing Docker image..."
docker push $TAG
echo "Pushing Successful!"
