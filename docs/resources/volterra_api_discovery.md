---

page_title: "Volterra: api_discovery"

description: "The api_discovery allows CRUD of Api Discovery resource on Volterra SaaS"
---------------------------------------------------------------------------------------

Resource volterra_api_discovery
===============================

The Api Discovery allows CRUD of Api Discovery resource on Volterra SaaS

~> **Note:** Please refer to [Api Discovery API docs](https://docs.cloud.f5.com/docs-v2/api/api-discovery) to learn more

Example Usage
-------------

```hcl
resource "volterra_api_discovery" "example" {
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

`custom_auth_types` - (Optional) Select your custom authentication types to be detected in the API discovery. See [Custom Auth Types ](#custom-auth-types) below for details.

### Custom Auth Types

Select your custom authentication types to be detected in the API discovery.

`parameter_name` - (Required) The authentication parameter name. (`String`).

`parameter_type` - (Required) x-displayName: "Parameter Type" (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured api_discovery.
