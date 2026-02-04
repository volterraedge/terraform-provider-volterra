terraform {
  required_providers {
    volterra = {
      source  = "localhost/volterraedge/volterra"
      version = "0.1.1"
      #source  = "volterraedge/volterra"
      #version = "= 0.11.44"
    }
  }
}

variable "f5xc_api_url" {
  type = string
}

variable "f5xc_api_p12_file" {
  type = string
}

provider "volterra" {
  api_p12_file = var.f5xc_api_p12_file
  url          = var.f5xc_api_url
}

resource "volterra_dns_lb_health_check" "system_auto-icmp-jolt" {
  icmp_health_check = true
  name              = "auto-icmp-jolt"
  namespace         = "system"
}

resource "volterra_dns_lb_health_check" "system_auto-tcp-rbpq" {
  tcp_health_check {
    health_check_port           = 7001
    health_check_secondary_port = 0
    receive                     = ""
    send                        = ""
  }
  name      = "auto-tcp-rbpq"
  namespace = "system"
}

resource "volterra_app_firewall" "system_auto-test-waf-xghr" {
  allow_all_response_codes   = true
  default_anonymization      = true
  use_default_blocking_page  = true
  default_bot_setting        = true
  default_detection_settings = true
  blocking                   = true
  name                       = "auto-test-waf-xghr"
  namespace                  = "ns-azure-ce-perf-0"
}

resource "volterra_origin_pool" "system_ny8-nyc-origin-poo-01" {
  endpoint_selection     = "DISTRIBUTED"
  loadbalancer_algorithm = "ROUND_ROBIN"

  origin_servers {
    k8s_service {
      outside_network = true
      protocol        = "PROTOCOL_TCP"
      service_name    = "productpage.system"
      site_locator {
        site {
          name      = "ny8-nyc"
          namespace = "system"
          tenant    = "ves-io"
        }
      }
    }
  }

  port      = 9080
  no_tls    = true
  name      = "ny8-nyc-origin-pool-01"
  namespace = "ns-azure-ce-perf-0"
}

resource "volterra_origin_pool" "system_ny8-nyc-origin-poo-02" {
  endpoint_selection     = "DISTRIBUTED"
  loadbalancer_algorithm = "ROUND_ROBIN"

  origin_servers {
    k8s_service {
      outside_network = true
      protocol        = "PROTOCOL_TCP"
      service_name    = "productpage.system"
      site_locator {
        site {
          name      = "ny8-nyc"
          namespace = "system"
          tenant    = "ves-io"
        }
      }
    }
  }

  port      = 9080
  no_tls    = true
  name      = "ny8-nyc-origin-pool-02"
  namespace = "ns-azure-ce-perf-0"
}

resource "volterra_origin_pool" "system_ny8-nyc-origin-poo-03" {
  endpoint_selection     = "DISTRIBUTED"
  loadbalancer_algorithm = "ROUND_ROBIN"

  origin_servers {
    k8s_service {
      outside_network = true
      protocol        = "PROTOCOL_TCP"
      service_name    = "productpage.system"
      site_locator {
        site {
          name      = "ny8-nyc"
          namespace = "system"
          tenant    = "ves-io"
        }
      }
    }
  }

  port      = 9080
  no_tls    = true
  name      = "ny8-nyc-origin-pool-03"
  namespace = "ns-azure-ce-perf-0"
}

resource "volterra_origin_pool" "system_ny8-nyc-origin-poo-04" {
  endpoint_selection     = "DISTRIBUTED"
  loadbalancer_algorithm = "ROUND_ROBIN"


  origin_servers {
    k8s_service {
      outside_network = true
      protocol        = "PROTOCOL_TCP"
      service_name    = "productpage.system"
      site_locator {
        site {
          name      = "ny8-nyc"
          namespace = "system"
          tenant    = "ves-io"
        }
      }
    }
  }

  port      = 9080
  no_tls    = true
  name      = "ny8-nyc-origin-pool-04"
  namespace = "ns-azure-ce-perf-0"
}

resource "volterra_http_loadbalancer" "system_productpage" {
  add_location                    = false
  advertise_on_public_default_vip = true
  no_challenge                    = true
  domains                         = ["vk8s-bookinfo-ves-io-all-res-cesslui"]
  round_robin                     = true

  http {
    dns_volterra_managed = false
    port                 = 80
  }

  /*
  caching_policy {
    custom_cache_rule {
      cdn_cache_rules {
        name = "ck-test-01"
        namespace = "system"
        tenant    = ""
      }
    }
    default_cache_action {
      cache_ttl_default = "7h"
    }
  }
   */

  disable_caching = true
  multi_lb_app       = true
  disable_rate_limit = true

  routes {
    simple_route {
      advanced_options {
        waf_exclusion_policy {
          name      = "ck-test-01"
          namespace = "ns-azure-ce-perf-0"
          tenant    = "playground-urxgwtyy"
        }
        common_buffering          = true
        do_not_retract_cluster    = true
        disable_location_add      = false
        common_hash_policy        = true
        disable_mirroring         = true
        priority                  = "DEFAULT"
        default_retry_policy      = true
        prefix_rewrite            = "/"
        disable_spdy              = true
        timeout                   = 0
        inherited_waf             = true
        disable_web_socket_config = true
      }
      disable_host_rewrite = true
      http_method          = "ANY"

      origin_pools {
        pool {
          name      = volterra_origin_pool.system_ny8-nyc-origin-poo-04.name
          namespace = "ns-azure-ce-perf-0"
          tenant    = ""
        }
        priority = 0
        weight   = 1
      }
      path {
        path = "/ny8-nyc/"
      }

      uuid = ""
    }
  }
  routes {
    simple_route {
      advanced_options {
        waf_exclusion_policy {
          name      = "ck-test-01"
          namespace = "ns-azure-ce-perf-0"
          tenant    = ""
        }
        common_buffering          = true
        do_not_retract_cluster    = true
        disable_location_add      = false
        common_hash_policy        = true
        disable_mirroring         = true
        priority                  = "DEFAULT"
        default_retry_policy      = true
        prefix_rewrite            = "/"
        disable_spdy              = true
        timeout                   = 0
        inherited_waf             = true
        disable_web_socket_config = true
      }
      disable_host_rewrite = true
      http_method          = "ANY"

      origin_pools {
        pool {
          name      = volterra_origin_pool.system_ny8-nyc-origin-poo-01.name
          namespace = "ns-azure-ce-perf-0"
          tenant    = "playground-urxgwtyy"
        }
        priority = 0
        weight   = 1
      }
      path {
        path = "/sv10-sjc/"
      }

      uuid = ""
    }
  }

  routes {
    simple_route {
      advanced_options {
        waf_exclusion_policy {
          name      = "ck-test-01"
          namespace = "ns-azure-ce-perf-0"
          tenant    = ""
        }
        common_buffering          = true
        do_not_retract_cluster    = true
        disable_location_add      = false
        common_hash_policy        = true
        disable_mirroring         = true
        priority                  = "DEFAULT"
        default_retry_policy      = true
        prefix_rewrite            = "/"
        disable_spdy              = true
        timeout                   = 0
        inherited_waf             = true
        disable_web_socket_config = true
      }
      disable_host_rewrite = true
      http_method          = "ANY"

      origin_pools {
        pool {
          name      = "ny8-nyc-origin-pool-02"
          namespace = "ns-azure-ce-perf-0"
          tenant    = ""
        }
        priority = 0
        weight   = 1
      }
      path {
        path = "/pa4-par/"
      }

      uuid = ""
    }
  }
  routes {
    simple_route {
      advanced_options {
        waf_exclusion_policy {
          name      = "ck-test-01"
          namespace = "ns-azure-ce-perf-0"
          tenant    = ""
        }
        common_buffering          = true
        do_not_retract_cluster    = true
        disable_location_add      = false
        common_hash_policy        = true
        disable_mirroring         = true
        priority                  = "DEFAULT"
        default_retry_policy      = true
        prefix_rewrite            = "/"
        disable_spdy              = true
        timeout                   = 0
        inherited_waf             = true
        disable_web_socket_config = true
      }

      disable_host_rewrite = true
      http_method          = "ANY"

      origin_pools {
        pool {
          name      = volterra_origin_pool.system_ny8-nyc-origin-poo-02.name
          namespace = "ns-azure-ce-perf-0"
          tenant    = ""
        }
        priority = 0
        weight   = 1
      }

      path {
        path = "/auto-aws-ce01-new/"
      }

      uuid = ""
    }
  }
  routes {
    simple_route {
      advanced_options {
        waf_exclusion_policy {
          name      = "ck-test-01"
          namespace = "ns-azure-ce-perf-0"
          tenant    = ""
        }
        common_buffering          = true
        do_not_retract_cluster    = true
        disable_location_add      = false
        common_hash_policy        = true
        disable_mirroring         = true
        priority                  = "DEFAULT"
        default_retry_policy      = true
        prefix_rewrite            = "/"
        disable_spdy              = true
        timeout                   = 0
        inherited_waf             = true
        disable_web_socket_config = true
      }
      disable_host_rewrite = true
      http_method          = "ANY"

      origin_pools {
        pool {
          name      = volterra_origin_pool.system_ny8-nyc-origin-poo-03.name
          namespace = "ns-azure-ce-perf-0"
          tenant    = ""
        }
        priority = 0
        weight   = 1
      }
      path {
        path = "/auto-az-new-test/"
      }

      uuid = ""
    }
  }

  /*
  l7_ddos_protection {
    mitigation_block = true
    default_rps_threshold = true
    clientside_action_none = true
    ddos_policy_none = true
  }
   */

  service_policies_from_namespace = true
  user_id_client_ip               = true
  disable_waf                     = true
  name                            = "productpage"
  namespace                       = "ns-azure-ce-perf-0"
}

resource "volterra_origin_pool" "aj-test-ns_auto-auto-lb-op-hdct" {
  endpoint_selection = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"

  origin_servers {
    public_ip {
      ip = "34.168.129.164"
    }
  }

  port   = 10101
  no_tls = true
  name      = "auto-auto-lb-op-hdct"
  namespace = "default"
}

/*
resource "volterra_udp_loadbalancer" "aj-test-ns_auto-auto-udp-lb-wjkq" {
  advertise_on_public_default_vip  = true
  retract_cluster                  = true
  dns_volterra_managed             = false
  domains                          = ["example.com"]
  enable_per_packet_load_balancing = false
  hash_policy_choice_round_robin   = true
  idle_timeout                     = 0
  udp                              = true

  origin_pools_weights {
    pool {
      name      = volterra_origin_pool.aj-test-ns_auto-auto-lb-op-hdct.name
      namespace = "default"
      tenant    = ""
    }
    priority = 0
    weight   = 0
  }

  listen_port = 37088
  name        = "auto-auto-udp-lb-wjkq"
  namespace   = "default"
}
 */