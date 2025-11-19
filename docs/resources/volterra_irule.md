---

page_title: "Volterra: irule"

description: "The irule allows CRUD of Irule resource on Volterra SaaS"
-----------------------------------------------------------------------

Resource volterra_irule
=======================

The Irule allows CRUD of Irule resource on Volterra SaaS

~> **Note:** Please refer to [Irule API docs](https://docs.cloud.f5.com/docs-v2/api/irule) to learn more

Example Usage
-------------

```hcl
resource "volterra_irule" "example" {
  name        = "acmecorp-web"
  namespace   = "staging"
  description = ["description"]
  irule       = ["when HTTP_REQUEST {if { !([HTTP::host] contains \".example.f5.com\") } {reject}}"]
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

`description` - (Required) Specify Description for iRule (`String`).

`irule` - (Required) irule content (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured irule.
