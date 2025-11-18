---

page_title: "Volterra: infraprotect_asn"

description: "The infraprotect_asn allows CRUD of Infraprotect Asn resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_infraprotect_asn
==================================

The Infraprotect Asn allows CRUD of Infraprotect Asn resource on Volterra SaaS

~> **Note:** Please refer to [Infraprotect Asn API docs](https://docs.cloud.f5.com/docs-v2/api/infraprotect-asn) to learn more

Example Usage
-------------

```hcl
resource "volterra_infraprotect_asn" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
  asn       = ["12345"]

  // One of the arguments from this list "bgp_session_disabled bgp_session_enabled" must be set

  bgp_session_enabled = true
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

`asn` - (Required) 2-byte or 4-byte Autonomous System Number (ASN) (`Int`).

###### One of the arguments from this list "bgp_session_disabled, bgp_session_enabled" must be set

`bgp_session_disabled` - (Optional) BGP Sessions Not Enabled (`Bool`).

`bgp_session_enabled` - (Optional) BGP Sessions Enabled (`Bool`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured infraprotect_asn.
