---

page_title: "Volterra: endpoint"

description: "The endpoint allows CRUD of Endpoint resource on Volterra SaaS"
-----------------------------------------------------------------------------

Resource volterra_endpoint
==========================

The Endpoint allows CRUD of Endpoint resource on Volterra SaaS

~> **Note:** Please refer to [Endpoint API docs](https://docs.cloud.f5.com/docs-v2/api/endpoint) to learn more

Example Usage
-------------

```hcl
resource "volterra_endpoint" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
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

###### One of the arguments from this list "dns_name, dns_name_advanced, ip, service_info" can be set

`dns_name` - (Optional) Endpoint's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

`dns_name_advanced` - (Optional) Specifies name and TTL used for DNS resolution.. See [Endpoint Address Dns Name Advanced ](#endpoint-address-dns-name-advanced) below for details.

`ip` - (Optional) Endpoint is reachable at the given ipv4/ipv6 address (`String`).

`service_info` - (Optional) In case of Consul, tags on the service is matched against service_selector. See [Endpoint Address Service Info ](#endpoint-address-service-info) below for details.

`health_check_port` - (Optional) Setting this with a non-zero value allows an endpoint to have different health check port. (`Int`).

`port` - (Optional) Endpoint service is available on this port (`Int`).

`protocol` - (Optional) Both TCP and UDP protocols are supported (`String`).

`snat_pool` - (Optional) Configured SNAT Pool. See [Snat Pool ](#snat-pool) below for details.

`where` - (Optional) This endpoint is present in site, virtual_site or virtual_network selected by following field.. See [Where ](#where) below for details.

### Snat Pool

Configured SNAT Pool.

###### One of the arguments from this list "no_snat_pool, snat_pool" can be set

`no_snat_pool` - (Optional) No configured SNAT Pool to reach Origin Server (`Bool`).

`snat_pool` - (Optional) Configure SNAT Pool to reach Origin Server. See [Snat Pool Choice Snat Pool ](#snat-pool-choice-snat-pool) below for details.

### Where

This endpoint is present in site, virtual_site or virtual_network selected by following field..

###### One of the arguments from this list "site, virtual_network, virtual_site" must be set

`site` - (Optional) Direct reference to site object. See [Ref Or Selector Site ](#ref-or-selector-site) below for details.

`virtual_network` - (Optional) Direct reference to virtual network object. See [Ref Or Selector Virtual Network ](#ref-or-selector-virtual-network) below for details.

`virtual_site` - (Optional) Direct reference to virtual site object. See [Ref Or Selector Virtual Site ](#ref-or-selector-virtual-site) below for details.

### Endpoint Address Dns Name Advanced

Specifies name and TTL used for DNS resolution..

`name` - (Optional) Endpoint's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

###### One of the arguments from this list "refresh_interval" can be set

`refresh_interval` - (Optional) Interval for DNS refresh in seconds. (`Int`).

### Endpoint Address Service Info

```
 In case of Consul, tags on the service is matched against service_selector.
```

`discovery_type` - (Required) Specifies whether the discovery is from Kubernetes or Consul cluster (`String`).

###### One of the arguments from this list "service_name, service_selector" can be set

`service_name` - (Optional) discovery objects of the site. (`String`).

`service_selector` - (Optional) discovery has to happen. This implicit label is added to service_selector. See [Service Info Service Selector ](#service-info-service-selector) below for details.

### Internet Vip Choice Disable Internet Vip

Do not enable advertise on external internet vip..

### Internet Vip Choice Enable Internet Vip

Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site..

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Ref Or Selector Site

Direct reference to site object.

###### One of the arguments from this list "disable_internet_vip, enable_internet_vip" must be set

`disable_internet_vip` - (Optional) Do not enable advertise on external internet vip. (`Bool`).

`enable_internet_vip` - (Optional) Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site. (`Bool`).

`network_type` - (Optional) The type of network on the referred site (`String`).

`ref` - (Required) A site direct reference. See [ref](#ref) below for details.

### Ref Or Selector Virtual Network

Direct reference to virtual network object.

`ref` - (Required) A virtual network direct reference. See [ref](#ref) below for details.

### Ref Or Selector Virtual Site

Direct reference to virtual site object.

###### One of the arguments from this list "disable_internet_vip, enable_internet_vip" must be set

`disable_internet_vip` - (Optional) Do not enable advertise on external internet vip. (`Bool`).

`enable_internet_vip` - (Optional) Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site. (`Bool`).

`network_type` - (Optional) The type of network on the referred virtual_site (`String`).

`ref` - (Required) A virtual_site direct reference. See [ref](#ref) below for details.

### Service Info Service Selector

discovery has to happen. This implicit label is added to service_selector.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Snat Pool Choice No Snat Pool

No configured SNAT Pool to reach Origin Server.

### Snat Pool Choice Snat Pool

Configure SNAT Pool to reach Origin Server.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured endpoint.
