#!/bin/bash

AWS_PROFILE=$1
MODE=$2

cd ./infrastructure/swarm
set -e

echo "Running Terraform Init..."
terraform init
echo "Init Successul!"

if [ $MODE == "up" ]; then
    echo "Running Terraform Plan..."
    terraform plan -var="profile=$1"
    echo "Plan Successul!"

    echo "Running Terraform Apply..."
    terraform apply -var="profile=$1"
    echo "Apply Successul!"
else
    echo "Running Terraform Plan Destroy..."
    terraform plan -destroy -var="profile=$1"
    echo "Plan Destroy Successul!"

    echo "Running Terraform Apply Destroy..."
    terraform apply -destroy -var="profile=$1"
    echo "Apply Destroy Successul!"
fi

cd ../../
