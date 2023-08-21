resource "aws_key_pair" "swarm_node_key" {
  key_name   = local.ssh_key
  public_key = file(var.ssh_key_path)
}

resource "aws_instance" "swarm_node" {
  for_each = { for index, node in local.configuration : node.name => node }

  ami                    = data.aws_ami.ubuntu.id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.swarm_node_key.key_name
  vpc_security_group_ids = [aws_security_group.swarm_node_sg.id]

  tags = {
    Name = each.value.name
  }
}
