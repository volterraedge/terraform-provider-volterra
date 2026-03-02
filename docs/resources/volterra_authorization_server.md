---

page_title: "Volterra: authorization_server"

description: "The authorization_server allows CRUD of Authorization Server resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------

Resource volterra_authorization_server
======================================

The Authorization Server allows CRUD of Authorization Server resource on Volterra SaaS

~> **Note:** Please refer to [Authorization Server API docs](https://docs.cloud.f5.com/docs-v2/api/authorization-server) to learn more

Example Usage
-------------

```hcl
resource "volterra_authorization_server" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
  jwks_uri  = ["jwks_uri"]
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

`jwks_uri` - (Required) x-textBlockPosition: "inside" (`String`).

Attribute Reference
-------------------

*   `id` - This is the id of the configured authorization_server.
