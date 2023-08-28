#!/bin/bash

AWS_PROFILE=$1
MODE=$2
PLAN_PATH=tfplan

set -e

echo "Running Terraform Init..."
terraform init
echo "Init Successul!"

if [ $MODE == "up" ]; then
    echo "Running Terraform Plan..."
    terraform plan \
        -var="profile=$1" \
        -out=$PLAN_PATH
    echo "Plan Successul!"

    echo "Running Terraform Apply..."
    terraform apply \
        $PLAN_PATH
    echo "Apply Successul!"
else
    echo "Running Terraform Plan Destroy..."
    terraform plan \
        -destroy \
        -var="profile=$1" \
        -out=$PLAN_PATH
    echo "Plan Destroy Successul!"

    echo "Running Terraform Apply Destroy..."
    terraform apply \
        -destroy \
        $PLAN_PATH
    echo "Apply Destroy Successul!"
fi
