#!/bin/bash

AWS_PROFILE=$1
TOOL=$2
MODE=$3

cd ./infrastructure/$TOOL
set -e

echo "Deploying Infrastructure..."
./deploy.sh $AWS_PROFILE $MODE
echo "Deploy Successul!"

cd ../../
