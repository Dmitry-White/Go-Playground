variable "region" {
  description = "Region to deploy to"
  default     = "eu-central-1"
}

variable "instance_type" {
  description = "Type of EC2 instance to provision"
  default     = "t3.nano"
}

variable "ssh_key_path" {
  type        = string
  description = "SSH Key Path"
  default     = "~/.ssh/id_rsa.pub"
}

variable "profile" {
  description = "AWS Profile to use"
}

locals {
  sg_name = "Swarm node security group"
  ssh_key = "swarm-node-key"
  configuration = [
    {
      name = "node-1"
    },
    {
      name = "node-2"
    }
  ]
}

data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}
