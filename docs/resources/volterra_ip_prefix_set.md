---

page_title: "Volterra: ip_prefix_set"

description: "The ip_prefix_set allows CRUD of Ip Prefix Set resource on Volterra SaaS"
---------------------------------------------------------------------------------------

Resource volterra_ip_prefix_set
===============================

The Ip Prefix Set allows CRUD of Ip Prefix Set resource on Volterra SaaS

~> **Note:** Please refer to [Ip Prefix Set API docs](https://docs.cloud.f5.com/docs-v2/api/ip-prefix-set) to learn more

Example Usage
-------------

```hcl
resource "volterra_ip_prefix_set" "example" {
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

`ipv4_prefixes` - (Optional) list of IPv4 prefixes with description.. See [Ipv4 Prefixes ](#ipv4-prefixes) below for details.

`ipv6_prefixes` - (Optional) list of IPv6 prefixes with description.. See [Ipv6 Prefixes ](#ipv6-prefixes) below for details.

### Ipv4 Prefixes

list of IPv4 prefixes with description..

`description` - (Optional) x-example: "blocked ip" (`String`).

`ipv4_prefix` - (Required) x-example: "192.0.2.146/22" (`String`).

### Ipv6 Prefixes

list of IPv6 prefixes with description..

`description` - (Optional) x-example: "blocked ip" (`String`).

`ipv6_prefix` - (Required) x-example: "2001:db8:3c4d:15::/64" (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured ip_prefix_set.
