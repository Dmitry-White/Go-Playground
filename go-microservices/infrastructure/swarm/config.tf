data "template_file" "user_data" {
  template = local.user_data_file_content
  vars = {
    tf_node_user = local.node_user
    tf_home_path = local.home_path
  }
}

data "cloudinit_config" "swarm_node_config" {
  gzip          = false
  base64_encode = false

  part {
    content_type = "text/cloud-config"
    filename     = local.swarm_file
    content      = local.swarm_file_content
  }

  part {
    content_type = "text/x-shellscript"
    filename     = local.user_data_file
    content      = data.template_file.user_data.rendered
  }
}
