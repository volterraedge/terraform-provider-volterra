---

page_title: "Volterra: http_router"

description: "The http_router allows CRUD of Http Router resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_http_router
=============================

The Http Router allows CRUD of Http Router resource on Volterra SaaS

~> **Note:** Please refer to [Http Router API docs](https://docs.cloud.f5.com/docs-v2/api/http-router) to learn more

Example Usage
-------------

```hcl
resource "volterra_http_router" "example" {
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

`app_service` - (Optional) Specifies the Application Service that owns and manages this object (`String`).

`full_path` - (Optional) Specifies the canonical, system-wide unique path combining partition, optional sub_path, and name (`String`).

`sub_path` - (Optional) Optional folder under the partition where the object resides (`String`).

`user_spec` - (Optional) User specified properties. (`List of String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured http_router.
