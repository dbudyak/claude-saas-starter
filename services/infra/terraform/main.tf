# Terraform starter for provisioning a server.
# This example uses Hetzner Cloud (cheap, simple, European).
# Swap the provider block for DigitalOcean, UpCloud, AWS EC2, etc.
#
# To use:
#   cd services/infra/terraform
#   cp terraform.tfvars.example terraform.tfvars
#   terraform init
#   terraform apply

terraform {
  required_version = ">= 1.5"

  required_providers {
    hcloud = {
      source  = "hetznercloud/hcloud"
      version = "~> 1.45"
    }
  }
}

provider "hcloud" {
  token = var.hcloud_token
}

resource "hcloud_ssh_key" "deploy" {
  name       = "${var.project_name}-deploy"
  public_key = file(var.ssh_public_key_path)
}

resource "hcloud_server" "app" {
  name        = var.project_name
  server_type = var.server_type
  image       = "ubuntu-24.04"
  location    = var.location
  ssh_keys    = [hcloud_ssh_key.deploy.id]

  user_data = templatefile("${path.module}/cloud-init.yml", {
    project_name = var.project_name
  })

  labels = {
    project = var.project_name
  }
}

resource "hcloud_firewall" "app" {
  name = "${var.project_name}-firewall"

  rule {
    direction = "in"
    protocol  = "tcp"
    port      = "22"
    source_ips = ["0.0.0.0/0", "::/0"]
  }

  rule {
    direction = "in"
    protocol  = "tcp"
    port      = "80"
    source_ips = ["0.0.0.0/0", "::/0"]
  }

  rule {
    direction = "in"
    protocol  = "tcp"
    port      = "443"
    source_ips = ["0.0.0.0/0", "::/0"]
  }
}

resource "hcloud_firewall_attachment" "app" {
  firewall_id = hcloud_firewall.app.id
  server_ids  = [hcloud_server.app.id]
}

output "server_ip" {
  value       = hcloud_server.app.ipv4_address
  description = "Set this as DEPLOY_HOST in GitHub secrets and point your DNS here."
}
