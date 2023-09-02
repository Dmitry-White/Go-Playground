#!/bin/bash

AWS_PROFILE=$1
MODE=$2
MANIFESTS_FOLDER="./cluster"

set -e

echo "Running Kubernetes Deploy..."

cd $MANIFESTS_FOLDER

services=$(ls -d */ | cut -f1 -d'/')
for value in $services; do
    echo "--------- $value ---------"
    if [ $MODE == "up" ]; then
        kubectl apply -f $value
    else
        kubectl delete -f $value
    fi
    echo "--------------------------"
done

cd ..

echo "Kubernetes Deploy Successful!"
