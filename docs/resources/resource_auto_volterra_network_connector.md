---

page_title: "Volterra: network_connector"

description: "The network_connector allows CRUD of Network Connector resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_network_connector
===================================

The Network Connector allows CRUD of Network Connector resource on Volterra SaaS

~> **Note:** Please refer to [Network Connector API docs](https://volterra.io/docs/api/network-connector) to learn more

Example Usage
-------------

```hcl
resource "volterra_network_connector" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "sli_to_slo_snat sli_to_slo_dr sli_to_global_dr sli_to_global_snat slo_to_global_dr slo_to_global_snat" must be set

  slo_to_global_snat {
    global_vn {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }

    snat_config {
      // One of the arguments from this list "interface_ip snat_pool snat_pool_allocator" must be set
      snat_pool = "snat_pool"

      // One of the arguments from this list "default_gw_snat dynamic_routing" must be set
      default_gw_snat = true
    }
  }
  // One of the arguments from this list "disable_forward_proxy enable_forward_proxy" must be set
  disable_forward_proxy = true
}

```

Argument Reference
------------------

### Metadata Argument Reference

`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).

`description` - (Optional) Human readable description for the object (`String`).

`disable` - (Optional) A value of true will administratively disable the object (`Bool`).

`labels` - (Optional) by selector expression (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).

### Spec Argument Reference

`sli_to_global_dr` - (Optional) Site local inside is connected directly to a given global network. See [Sli To Global Dr ](#sli-to-global-dr) below for details.

`sli_to_global_snat` - (Optional) Site local inside is connected to a given global network, using SNAT. See [Sli To Global Snat ](#sli-to-global-snat) below for details.

`sli_to_slo_dr` - (Optional) Site local inside is connected directly to site local outside (bool).

`sli_to_slo_snat` - (Optional) Site local inside is connected to site local outside, using SNAT. See [Sli To Slo Snat ](#sli-to-slo-snat) below for details.

`slo_to_global_dr` - (Optional) Site local outside is connected directly to a given global network. See [Slo To Global Dr ](#slo-to-global-dr) below for details.

`slo_to_global_snat` - (Optional) Site local outside is connected directly to a given global network. See [Slo To Global Snat ](#slo-to-global-snat) below for details.

`disable_forward_proxy` - (Optional) Forward Proxy is disabled for this connector (bool).

`enable_forward_proxy` - (Optional) Forward Proxy is enabled for this connector. See [Enable Forward Proxy ](#enable-forward-proxy) below for details.

### Enable Forward Proxy

Forward Proxy is enabled for this connector.

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2 seconds (`Int`).

`max_connect_attempts` - (Optional) Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1. (`Int`).

`white_listed_ports` - (Optional) Example "tmate" server port (`Int`).

`white_listed_prefixes` - (Optional) Example "tmate" server ip (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Sli To Global Dr

Site local inside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Sli To Global Snat

Site local inside is connected to a given global network, using SNAT.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

`snat_config` - (Optional) SNAT configuration to connect to global network. See [Snat Config ](#snat-config) below for details.

### Sli To Slo Snat

Site local inside is connected to site local outside, using SNAT.

`interface_ip` - (Optional) Interface IP of the outgoing interface will be used for SNAT (bool).

`snat_pool` - (Optional) IP from the ip pool prefix will be used for SNAT (`String`).

`snat_pool_allocator` - (Optional) IP from the ip pool prefix will be used for SNAT. See [ref](#ref) below for details.

`default_gw_snat` - (Optional) Default route in inside network to SNATed network (bool).

`dynamic_routing` - (Optional) Routes are imported in inside network from SNATed network (bool).

### Slo To Global Dr

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Slo To Global Snat

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

`snat_config` - (Optional) SNAT configuration to connect to global network. See [Snat Config ](#snat-config) below for details.

### Snat Config

SNAT configuration to connect to global network.

`interface_ip` - (Optional) Interface IP of the outgoing interface will be used for SNAT (bool).

`snat_pool` - (Optional) IP from the ip pool prefix will be used for SNAT (`String`).

`snat_pool_allocator` - (Optional) IP from the ip pool prefix will be used for SNAT. See [ref](#ref) below for details.

`default_gw_snat` - (Optional) Default route in inside network to SNATed network (bool).

`dynamic_routing` - (Optional) Routes are imported in inside network from SNATed network (bool).

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_connector.
