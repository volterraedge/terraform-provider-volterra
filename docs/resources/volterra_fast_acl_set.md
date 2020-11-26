---

page_title: "Volterra: fast_acl_set"

description: "The fast_acl_set allows CRUD of Fast Acl Set resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_fast_acl_set
==============================

The Fast Acl Set allows CRUD of Fast Acl Set resource on Volterra SaaS

~> **Note:** Please refer to [Fast Acl Set API docs](https://volterra.io/docs/api/fast-acl-set) to learn more

Example Usage
-------------

```hcl
resource "volterra_fast_acl_set" "example" {
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

`acl_list` - (Optional) Ordered list of references to the fast ACL that make up this fast ACL set.. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured fast_acl_set.
