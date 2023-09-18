---

page_title: "Volterra: app_type"

description: "The app_type allows CRUD of App Type resource on Volterra SaaS"
-----------------------------------------------------------------------------

Resource volterra_app_type
==========================

The App Type allows CRUD of App Type resource on Volterra SaaS

~> **Note:** Please refer to [App Type API docs](https://docs.cloud.f5.com/docs/api/app-type) to learn more

Example Usage
-------------

```hcl
resource "volterra_app_type" "example" {
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

`business_logic_markup_setting` - (Optional) Setting specifying how Business Logic Markup will be performed. See [Business Logic Markup Setting ](#business-logic-markup-setting) below for details.

`features` - (Optional) List of various AI/ML features enabled. See [Features ](#features) below for details.

### Business Logic Markup Setting

Setting specifying how Business Logic Markup will be performed.

`disable` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (bool).

`enable` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (bool).

### Disable

Disable learning API patterns from traffic with redirect response codes 3xx.

### Enable

Enable learning API patterns from traffic with redirect response codes 3xx.

### Features

List of various AI/ML features enabled.

`type` - (Optional) Feature type to be enabled (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured app_type.
