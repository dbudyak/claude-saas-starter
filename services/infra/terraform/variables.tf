variable "hcloud_token" {
  description = "Hetzner Cloud API token. Get from console.hetzner.cloud → Security → API Tokens."
  type        = string
  sensitive   = true
}

variable "project_name" {
  description = "Short project name used for resource naming (e.g. 'myapp')."
  type        = string
}

variable "server_type" {
  description = "Hetzner server type. cx22 (2 vCPU, 4GB) is enough for most early-stage apps."
  type        = string
  default     = "cx22"
}

variable "location" {
  description = "Data centre location. Options: nbg1 (Nuremberg), fsn1 (Falkenstein), hel1 (Helsinki)."
  type        = string
  default     = "hel1"
}

variable "ssh_public_key_path" {
  description = "Path to the SSH public key for server access."
  type        = string
  default     = "~/.ssh/id_ed25519.pub"
}
