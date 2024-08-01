











---
page_title: "Volterra: bgp_asn_set"
description: "The bgp_asn_set allows CRUD of Bgp Asn Set  resource on Volterra SaaS"
---
# Resource volterra_bgp_asn_set

The Bgp Asn Set  allows CRUD of Bgp Asn Set  resource on Volterra SaaS

~> **Note:** Please refer to [Bgp Asn Set  API docs](https://docs.cloud.f5.com/docs-v2/api/bgp-asn-set) to learn more

## Example Usage

```hcl
resource "volterra_bgp_asn_set" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  as_numbers = ["[713, 7932, 847325, 4683, 15269, 1000001]"]
}

```

## Argument Reference

### Metadata Argument Reference
`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).


`description` - (Optional) Human readable description for the object (`String`).


`disable` - (Optional) A value of true will administratively disable the object (`Bool`).


`labels` - (Optional) by selector expression (`String`).


`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).


`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).



### Spec Argument Reference

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create whitelists or blacklists for use in network policy or service policy. (`List of Int`).



## Attribute Reference

* `id` - This is the id of the configured bgp_asn_set.

