output "swarm_node_ip" {
  description = "Public IP address of the Swarm node"
  value       = [for node in aws_instance.swarm_node : node.public_ip]
}
