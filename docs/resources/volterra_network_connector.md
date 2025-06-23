---

page_title: "Volterra: network_connector"

description: "The network_connector allows CRUD of Network Connector resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_network_connector
===================================

The Network Connector allows CRUD of Network Connector resource on Volterra SaaS

~> **Note:** Please refer to [Network Connector API docs](https://docs.cloud.f5.com/docs-v2/api/network-connector) to learn more

Example Usage
-------------

```hcl
resource "volterra_network_connector" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "sli_to_global_dr sli_to_slo_snat slo_to_global_dr" must be set

  sli_to_global_dr {
    global_vn {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
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

###### One of the arguments from this list "sli_to_global_dr, sli_to_slo_snat, slo_to_global_dr" must be set

`sli_to_global_dr` - (Optional) Site local inside is connected directly to a given global network. See [Connector Choice Sli To Global Dr ](#connector-choice-sli-to-global-dr) below for details.

`sli_to_slo_snat` - (Optional) Site local inside is connected to site local outside, using SNAT. See [Connector Choice Sli To Slo Snat ](#connector-choice-sli-to-slo-snat) below for details.

`slo_to_global_dr` - (Optional) Site local outside is connected directly to a given global network. See [Connector Choice Slo To Global Dr ](#connector-choice-slo-to-global-dr) below for details.

###### One of the arguments from this list "disable_forward_proxy, enable_forward_proxy" must be set

`disable_forward_proxy` - (Optional) Forward Proxy is disabled for this connector (`Bool`).

`enable_forward_proxy` - (Optional) Forward Proxy is enabled for this connector. See [Forward Proxy Choice Enable Forward Proxy ](#forward-proxy-choice-enable-forward-proxy) below for details.

### Connector Choice Sli To Global Dr

Site local inside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Connector Choice Sli To Slo Snat

Site local inside is connected to site local outside, using SNAT.

###### One of the arguments from this list "interface_ip" must be set

`interface_ip` - (Optional) Interface IP of the outgoing interface will be used for SNAT (`Bool`).

###### One of the arguments from this list "default_gw_snat" must be set

`default_gw_snat` - (Optional) Default route in inside network to SNATed network (`Bool`).

### Connector Choice Slo To Global Dr

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Forward Proxy Choice Enable Forward Proxy

Forward Proxy is enabled for this connector.

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2000 (2 seconds) (`Int`).

`max_connect_attempts` - (Optional) Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1. (`Int`).

`white_listed_ports` - (Optional) Example "tmate" server port (`Int`).

`white_listed_prefixes` - (Optional) Example "tmate" server ip (`String`).

### Pool Choice Interface Ip

Interface IP of the outgoing interface will be used for SNAT.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Routing Choice Default Gw Snat

Default route in inside network to SNATed network.

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_connector.
