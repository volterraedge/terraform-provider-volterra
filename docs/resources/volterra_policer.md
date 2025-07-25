---

page_title: "Volterra: policer"

description: "The policer allows CRUD of Policer resource on Volterra SaaS"
---------------------------------------------------------------------------

Resource volterra_policer
=========================

The Policer allows CRUD of Policer resource on Volterra SaaS

~> **Note:** Please refer to [Policer API docs](https://docs.cloud.f5.com/docs-v2/api/policer) to learn more

Example Usage
-------------

```hcl
resource "volterra_policer" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  burst_size = ["10000"]

  committed_information_rate = ["10000"]
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

`burst_size` - (Required) e.g. 10000 pps burst (`Int`).

`committed_information_rate` - (Required) e.g. 10000 pps (`Int`).

`policer_mode` - (Optional) be created with mode as "Shared" (`String`).

`policer_type` - (Optional) policer (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured policer.
