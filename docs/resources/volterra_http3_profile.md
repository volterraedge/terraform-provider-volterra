---

page_title: "Volterra: http3_profile"

description: "The http3_profile allows CRUD of Http3 Profile resource on Volterra SaaS"
---------------------------------------------------------------------------------------

Resource volterra_http3_profile
===============================

The Http3 Profile allows CRUD of Http3 Profile resource on Volterra SaaS

~> **Note:** Please refer to [Http3 Profile API docs](https://docs.cloud.f5.com/docs-v2/api/http3-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_http3_profile" "example" {
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

`app_service` - (Optional) The application service to which the object belongs. (`String`).

`command` - (Optional) x-displayName: "Command" (`String`).

`header_table_size` - (Optional) Specifies what table size will be used for the compression of headers. (`Int`).

`sub_path` - (Optional) x-displayName: "Sub Path" (`String`).

`user_spec` - (Optional) User specified properties. (`List of String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured http3_profile.
