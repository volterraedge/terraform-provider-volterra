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

`dns_name_advance` - (Optional) Specifies name and TTL used for DNS resolution.. See [Dns Name Advance ](#dns-name-advance) below for details.

`ip` - (Optional) Endpoint is reachable at the given ip address (`String`).

`service_info` - (Optional) In case of Consul, tags on the service is matched against service_selector. See [Service Info ](#service-info) below for details.

`port` - (Optional) Endpoint service is available on this port (`Int`).

`protocol` - (Optional) Both TCP and UDP protocols are supported (`String`).

`where` - (Optional) This endpoint is present in site, virtual_site or virtual_network selected by following field.. See [Where ](#where) below for details.

### Dns Name Advance

Specifies name and TTL used for DNS resolution..

`name` - (Optional) Endpoint's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

`refresh_interval` - (Optional) Interval for DNS refresh in seconds. (`Int`).

`strict_ttl` - (Optional) Use TTL value returned by DNS Server during DNS resolution as DNS refresh interval (bool).

### Service Info

```
 In case of Consul, tags on the service is matched against service_selector.
```

`discovery_type` - (Required) Specifies whether the discovery is from Kubernetes or Consul cluster (`String`).

`service_name` - (Optional) If namespace is not specified, then discovery is done in "default" namespace. (`String`).

`service_selector` - (Optional) Labels of the service to discover. See [Service Selector ](#service-selector) below for details.

### Where

This endpoint is present in site, virtual_site or virtual_network selected by following field..

`site` - (Optional) Direct reference to site object. See [Site ](#site) below for details.

`virtual_network` - (Optional) Direct reference to virtual network object. See [Virtual Network ](#virtual-network) below for details.

`virtual_site` - (Optional) Direct reference to virtual site object. See [Virtual Site ](#virtual-site) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured endpoint.
