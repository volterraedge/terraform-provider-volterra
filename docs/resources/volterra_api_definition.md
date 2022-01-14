---

page_title: "Volterra: api_definition"

description: "The api_definition allows CRUD of Api Definition resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_api_definition
================================

The Api Definition allows CRUD of Api Definition resource on Volterra SaaS

~> **Note:** Please refer to [Api Definition API docs](https://volterra.io/docs/api/views-api-definition) to learn more

Example Usage
-------------

```hcl
resource "volterra_api_definition" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  swagger_specs = ["swagger_specs"]
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

`swagger_specs` - (Required) Swagger Specs for this API Definition. (`List of String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured api_definition.
