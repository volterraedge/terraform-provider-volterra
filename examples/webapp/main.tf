provider "volterra" {
  api_p12_file = "${var.api_p12_file}"
  api_cert     = "${var.api_p12_file != "" ? "" : var.api_cert}"
  api_key      = "${var.api_p12_file != "" ? "" : var.api_key}"
  api_ca_cert  = "${var.api_ca_cert}"
  url          = "${var.api_url}"
}

resource "volterra_namespace" "volterra_ns" {
  name = "webapp-${var.name}-ns"
}

resource "volterra_api_credential" "api_cred" {
  count                 = "${var.diable_api_cred}"
  name                  = "api198712"
  api_credential_type   = "KUBE_CONFIG"
  virtual_k8s_namespace = "hipster"
  virtual_k8s_name      = "default"
}

resource "volterra_endpoint" "volterra_ep" {
  name      = "webapp-${var.name}-ep"
  namespace = "${volterra_namespace.volterra_ns.name}"

  endpoint_address {
    dns_name = "${var.webapp}"
  }

  port = 80

  where {
    ref_or_selector {
      site {
        ref {
          name      = "${var.site_name}"
          namespace = "system"
          tenant    = "ves-io"
        }

        network_type = "VIRTUAL_NETWORK_SITE_LOCAL"
      }
    }
  }

  depends_on = ["volterra_namespace.volterra_ns"]
}

resource "volterra_virtual_host" "volterra_test" {
  name      = "webapp-${var.name}-vh"
  namespace = "${volterra_namespace.volterra_ns.name}"

  advertise_policies {
    namespace = "${volterra_namespace.volterra_ns.name}"
    tenant    = "${volterra_namespace.volterra_ns.tenant_name}"
    name      = "${volterra_advertise_policy.volterra_ap.name}"
  }

  domains = ["${volterra_namespace.volterra_ns.tenant_name}.${var.webapp}.helloworld.app"]
  proxy   = "HTTP_PROXY"

  routes {
    namespace = "${volterra_namespace.volterra_ns.name}"
    tenant    = "${volterra_namespace.volterra_ns.tenant_name}"
    name      = "${volterra_route.volterra_rt.name}"
  }

  depends_on = ["volterra_namespace.volterra_ns", "volterra_advertise_policy.volterra_ap", "volterra_route.volterra_rt"]
}

data "volterra_virtual_host_dns_info" "dns_info" {
  name       = "webapp-${var.name}-vh"
  namespace  = "${volterra_namespace.volterra_ns.name}"
  depends_on = ["volterra_virtual_host.volterra_test"]
}

resource "volterra_advertise_policy" "volterra_ap" {
  name       = "webapp-${var.name}-ap"
  namespace  = "${volterra_namespace.volterra_ns.name}"
  depends_on = ["volterra_route.volterra_rt"]

  where {
    ref_or_selector {
      virtual_network {
        ref {
          name      = "public"
          namespace = "shared"
          tenant    = "ves-io"
        }
      }
    }
  }

  port = 80
}

resource "volterra_cluster" "volterra_cl" {
  name      = "webapp-${var.name}-cl"
  namespace = "${volterra_namespace.volterra_ns.name}"

  endpoints {
    namespace = "${volterra_namespace.volterra_ns.name}"
    tenant    = "${volterra_namespace.volterra_ns.tenant_name}"
    name      = "${volterra_endpoint.volterra_ep.name}"
  }

  depends_on = ["volterra_endpoint.volterra_ep"]
}

resource "volterra_route" "volterra_rt" {
  name      = "webapp-${var.name}-rt"
  namespace = "${volterra_namespace.volterra_ns.name}"

  routes {
    match {
      path {
        path_match {
          regex = "(.*?)"
        }
      }
    }

    route_action {
      route_destination {
        destinations {
          cluster {
            namespace = "${volterra_namespace.volterra_ns.name}"
            tenant    = "${volterra_namespace.volterra_ns.tenant_name}"
            name      = "${volterra_cluster.volterra_cl.name}"
          }
        }
      }
    }
  }

  depends_on = ["volterra_cluster.volterra_cl"]
}
