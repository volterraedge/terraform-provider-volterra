---

page_title: "Volterra: ip_prefix_set"

description: "The ip_prefix_set allows CRUD of Ip Prefix Set resource on Volterra SaaS"
---------------------------------------------------------------------------------------

Resource volterra_ip_prefix_set
===============================

The Ip Prefix Set allows CRUD of Ip Prefix Set resource on Volterra SaaS

~> **Note:** Please refer to [Ip Prefix Set API docs](https://volterra.io/docs/api/ip-prefix-set) to learn more

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

`ipv6_prefix` - (Optional) An unordered list of IPv6 prefixes. (`List of String`).

`prefix` - (Optional) An unordered list of IPv4 prefixes. (`List of String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured ip_prefix_set.
