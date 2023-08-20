resource "aws_instance" "swarm_node" {
  for_each = { for index, node in local.configuration : node.name => node }

  ami           = data.aws_ami.ubuntu.id
  instance_type = var.instance_type

  tags = {
    Name = each.value.name
  }
}
