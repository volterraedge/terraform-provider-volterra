---

page_title: "Volterra: app_type"

description: "The app_type allows CRUD of App Type resource on Volterra SaaS"
-----------------------------------------------------------------------------

Resource volterra_app_type
==========================

The App Type allows CRUD of App Type resource on Volterra SaaS

~> **Note:** Please refer to [App Type API docs](https://docs.cloud.f5.com/docs-v2/api/app-type) to learn more

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

`business_logic_markup_setting` - (Optional) Setting specifying how API Discovery will be performed. See [Business Logic Markup Setting ](#business-logic-markup-setting) below for details.

`features` - (Optional) List of various AI/ML features enabled. See [Features ](#features) below for details.

### Business Logic Markup Setting

Setting specifying how API Discovery will be performed.

`discovered_api_settings` - (Optional) x-displayName: "Discovered API Settings". See [Business Logic Markup Setting Discovered Api Settings ](#business-logic-markup-setting-discovered-api-settings) below for details.

###### One of the arguments from this list "disable, enable" can be set

`disable` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`enable` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

### Features

List of various AI/ML features enabled.

`type` - (Required) Feature type to be enabled (`String`).

### Business Logic Markup Setting Discovered Api Settings

x-displayName: "Discovered API Settings".

`purge_duration_for_inactive_discovered_apis` - (Optional) Inactive discovered API will be deleted after configured duration. (`Int`).

### Learn From Redirect Traffic Disable

Disable learning API patterns from traffic with redirect response codes 3xx.

### Learn From Redirect Traffic Enable

Enable learning API patterns from traffic with redirect response codes 3xx.

Attribute Reference
-------------------

-	`id` - This is the id of the configured app_type.
