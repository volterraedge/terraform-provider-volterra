---

page_title: "Volterra: infraprotect_asn_prefix"

description: "The infraprotect_asn_prefix allows CRUD of Infraprotect Asn Prefix resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------------

Resource volterra_infraprotect_asn_prefix
=========================================

The Infraprotect Asn Prefix allows CRUD of Infraprotect Asn Prefix resource on Volterra SaaS

~> **Note:** Please refer to [Infraprotect Asn Prefix API docs](https://docs.cloud.f5.com/docs-v2/api/infraprotect-asn-prefix) to learn more

Example Usage
-------------

```hcl
resource "volterra_infraprotect_asn_prefix" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  asn {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }

  prefix = ["10.10.10.0/24"]
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

`asn` - (Required) ASN. See [ref](#ref) below for details.

`prefix` - (Required) Prefix (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured infraprotect_asn_prefix.
