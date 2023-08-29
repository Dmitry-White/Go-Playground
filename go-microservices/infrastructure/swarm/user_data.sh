#!/bin/bash

METADATA_URL=http://169.254.169.254/latest/meta-data

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
echo "Set up Successful!"

echo "Installing the Docker packages..."
sudo apt-get -y install \
    docker-ce \
    docker-ce-cli \
    containerd.io \
    docker-buildx-plugin \
    docker-compose-plugin
echo "Installation Successful!"

echo "Preparing project folders..."
sudo mkdir -p \
    ${tf_home_path}/reverse-proxy/dal/data/data \
    ${tf_home_path}/reverse-proxy/dal/data/config \
    ${tf_home_path}/authentication-service/dal/data \
    ${tf_home_path}/logger-service/dal/data \
    ${tf_home_path}/listener-service/dal/data
echo "Preparation Successful!"

echo "Configuring instance..."
sudo usermod -aG docker ${tf_node_user}
nodeName=$(curl -H "X-aws-ec2-metadata-token: $TOKEN" -v $METADATA_URL/tags/instance/Name)
sudo hostnamectl set-hostname $nodeName
echo "Configuration Successful!"
