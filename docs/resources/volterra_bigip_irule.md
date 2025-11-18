---

page_title: "Volterra: bigip_irule"

description: "The bigip_irule allows CRUD of Bigip Irule resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_bigip_irule
=============================

The Bigip Irule allows CRUD of Bigip Irule resource on Volterra SaaS

~> **Note:** Please refer to [Bigip Irule API docs](https://docs.cloud.f5.com/docs-v2/api/bigip-irule) to learn more

Example Usage
-------------

```hcl
resource "volterra_bigip_irule" "example" {
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

`code` - (Optional) iRule code content, this content will be base64 encoded for preserving formating (`String`).

`irule_name` - (Optional) iRule name (`String`).

`source` - (Optional) iRule generation/updation source (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured bigip_irule.
