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

resource "volterra_fleet" "fleet" {
  name        = "test-fleet"
  namespace   = "system"
  fleet_label = "test-fleet"

  inside_virtual_network {
    name = "test1"
    namespace = "system"
    tenant    = data.volterra_namespace.system.tenant_name
  }

  outside_virtual_network {
    name = "test2"
    namespace = "system"
    tenant    = data.volterra_namespace.system.tenant_name
  }
  network_connectors {
    name      = "test3"
    namespace = "system"
    tenant    = data.volterra_namespace.system.tenant_name
  }
}
