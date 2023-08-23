#!/bin/bash

set -e

echo "Setting up Docker's apt repository..."
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get -y install \
    ca-certificates \
    curl \
    gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

# Add the repository to Apt sources:
echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" |
    sudo tee /etc/apt/sources.list.d/docker.list >/dev/null
sudo apt-get update
echo "Set up Successul!"

echo "Installing the Docker packages..."
sudo apt-get -y install \
    docker-ce \
    docker-ce-cli \
    containerd.io \
    docker-buildx-plugin \
    docker-compose-plugin
echo "Installation Successul!"

echo "Preparing project folders..."
mkdir \
    ./reverse-proxy/dal/data/data \
    ./reverse-proxy/dal/data/config \
    ./authentication-service/dal/data \
    ./logger-service/dal/data \
    ./listener-service/dal/data 
echo "Preparation Successul!"
