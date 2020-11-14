provider "volterra" {
  api_p12_file = var.api_p12_file
  api_cert     = var.api_p12_file != "" ? "" : var.api_cert
  api_key      = var.api_p12_file != "" ? "" : var.api_key
  api_ca_cert  = var.api_ca_cert
  url          = var.api_url
}

data "volterra_namespace" "system" {
  name = "system"
}

resource "volterra_network_firewall" "ps" {
  name      = "${var.name}-firewall"
  namespace = "system"
  disable_forward_proxy_policy = true
  disable_network_policy = true
  disable_fast_acl = true
}
