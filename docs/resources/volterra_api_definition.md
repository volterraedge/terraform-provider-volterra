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

  swagger_specs = ["https://my.tenant.domain/api/object_store/namespaces/my-ns/stored_objects/swagger/file-name/v1-22-01-12"]
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

`swagger_specs` - (Required) Notice file versions. If swagger file is updated, need to select a new version here to redefine the API. (`List of String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured api_definition.
