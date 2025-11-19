---

page_title: "Volterra: mobile_base_config"

description: "The mobile_base_config allows CRUD of Mobile Base Config resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------

Resource volterra_mobile_base_config
====================================

The Mobile Base Config allows CRUD of Mobile Base Config resource on Volterra SaaS

~> **Note:** Please refer to [Mobile Base Config API docs](https://docs.cloud.f5.com/docs-v2/api/mobile-base-config) to learn more

Example Usage
-------------

```hcl
resource "volterra_mobile_base_config" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
  fqdn      = ["mobile.example.com"]
  os        = ["os"]
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

`fqdn` - (Required) Fully qualified domain name (FQDN) (`String`).

`os` - (Required) Specify Mobile SDK Operating System to use (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured mobile_base_config.
