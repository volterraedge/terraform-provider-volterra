---

page_title: "Volterra: ike1"

description: "The ike1 allows CRUD of Ike1 resource on Volterra SaaS"
---------------------------------------------------------------------

Resource volterra_ike1
======================

The Ike1 allows CRUD of Ike1 resource on Volterra SaaS

~> **Note:** Please refer to [Ike1 API docs](https://docs.cloud.f5.com/docs-v2/api/ike1) to learn more

Example Usage
-------------

```hcl
resource "volterra_ike1" "example" {
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

Attribute Reference
-------------------

-	`id` - This is the id of the configured ike1.
