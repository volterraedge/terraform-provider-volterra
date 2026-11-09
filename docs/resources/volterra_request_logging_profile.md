---

page_title: "Volterra: request_logging_profile"

description: "The request_logging_profile allows CRUD of Request Logging Profile resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------------

Resource volterra_request_logging_profile
=========================================

The Request Logging Profile allows CRUD of Request Logging Profile resource on Volterra SaaS

~> **Note:** Please refer to [Request Logging Profile API docs](https://docs.cloud.f5.com/docs-v2/api/request-logging-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_request_logging_profile" "example" {
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

Attribute Reference
-------------------

-	`id` - This is the id of the configured request_logging_profile.
