---

page_title: "Volterra: endpoint"

description: "The endpoint allows CRUD of Endpoint resource on Volterra SaaS"
-----------------------------------------------------------------------------

Resource volterra_endpoint
==========================

The Endpoint allows CRUD of Endpoint resource on Volterra SaaS

~> **Note:** Please refer to [Endpoint API docs](https://volterra.io/docs/api/endpoint) to learn more

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

`dns_name` - (Optional) Endpoint's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

`dns_name_advanced` - (Optional) Specifies name and TTL used for DNS resolution.. See [Dns Name Advanced ](#dns-name-advanced) below for details.

`ip` - (Optional) Endpoint is reachable at the given ip address (`String`).

`service_info` - (Optional) In case of Consul, tags on the service is matched against service_selector. See [Service Info ](#service-info) below for details.

`health_check_port` - (Optional) Setting this with a non-zero value allows an endpoint to have different health check port. (`Int`).

`port` - (Optional) Endpoint service is available on this port (`Int`).

`protocol` - (Optional) Both TCP and UDP protocols are supported (`String`).

`where` - (Optional) This endpoint is present in site, virtual_site or virtual_network selected by following field.. See [Where ](#where) below for details.

### Disable Internet Vip

Do not enable advertise on external internet vip..

### Dns Name Advanced

Specifies name and TTL used for DNS resolution..

`name` - (Optional) Endpoint's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

`refresh_interval` - (Optional) Interval for DNS refresh in seconds. (`Int`).

`strict_ttl` - (Optional) Use TTL value returned by DNS Server during DNS resolution as DNS refresh interval (bool).

### Enable Internet Vip

Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site..

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Service Info

```
 In case of Consul, tags on the service is matched against service_selector.
```

`discovery_type` - (Required) Specifies whether the discovery is from Kubernetes or Consul cluster (`String`).

`service_name` - (Optional) discovery objects of the site. (`String`).

`service_selector` - (Optional) discovery has to happen. This implicit label is added to service_selector. See [Service Selector ](#service-selector) below for details.

### Service Selector

discovery has to happen. This implicit label is added to service_selector.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Site

Direct reference to site object.

`disable_internet_vip` - (Optional) Do not enable advertise on external internet vip. (bool).

`enable_internet_vip` - (Optional) Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site. (bool).

`network_type` - (Optional) The type of network on the referred site (`String`).

`ref` - (Optional) A site direct reference. See [ref](#ref) below for details.

### Strict Ttl

Use TTL value returned by DNS Server during DNS resolution as DNS refresh interval.

### Virtual Network

Direct reference to virtual network object.

`ref` - (Optional) A virtual network direct reference. See [ref](#ref) below for details.

### Virtual Site

Direct reference to virtual site object.

`disable_internet_vip` - (Optional) Do not enable advertise on external internet vip. (bool).

`enable_internet_vip` - (Optional) Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site. (bool).

`network_type` - (Optional) The type of network on the referred virtual_site (`String`).

`ref` - (Optional) A virtual_site direct reference. See [ref](#ref) below for details.

### Where

This endpoint is present in site, virtual_site or virtual_network selected by following field..

`site` - (Optional) Direct reference to site object. See [Site ](#site) below for details.

`virtual_network` - (Optional) Direct reference to virtual network object. See [Virtual Network ](#virtual-network) below for details.

`virtual_site` - (Optional) Direct reference to virtual site object. See [Virtual Site ](#virtual-site) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured endpoint.
