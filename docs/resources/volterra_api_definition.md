---

page_title: "Volterra: api_definition"

description: "The api_definition allows CRUD of Api Definition resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_api_definition
================================

The Api Definition allows CRUD of Api Definition resource on Volterra SaaS

~> **Note:** Please refer to [Api Definition API docs](https://docs.cloud.f5.com/docs/api/views-api-definition) to learn more

Example Usage
-------------

```hcl
resource "volterra_api_definition" "example" {
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

`api_inventory_exclusion_list` - (Optional) note: hidden until full release of the API Inventory feature. See [Api Inventory Exclusion List ](#api-inventory-exclusion-list) below for details.

`api_inventory_inclusion_list` - (Optional) note: hidden until full release of the API Inventory feature. See [Api Inventory Inclusion List ](#api-inventory-inclusion-list) below for details.

`non_api_endpoints` - (Optional) note: hidden until full release of the API Inventory feature. See [Non Api Endpoints ](#non-api-endpoints) below for details.

`swagger_specs` - (Optional) Notice file versions. If OpenAPI file is updated, need to select a new version here to redefine the API. (`List of String`).

### Api Inventory Exclusion List

note: hidden until full release of the API Inventory feature.

`method` - (Required) Method to match the input request API method against. (`String`).

`path` - (Required) The path should comply with RFC 3986 and may have parameters according to OpenAPI specification (`String`).

### Api Inventory Inclusion List

note: hidden until full release of the API Inventory feature.

`method` - (Required) Method to match the input request API method against. (`String`).

`path` - (Required) The path should comply with RFC 3986 and may have parameters according to OpenAPI specification (`String`).

### Non Api Endpoints

note: hidden until full release of the API Inventory feature.

`method` - (Required) Method to match the input request API method against. (`String`).

`path` - (Required) The path should comply with RFC 3986 and may have parameters according to OpenAPI specification (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured api_definition.
