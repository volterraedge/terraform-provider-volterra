---

page_title: "Volterra: uztna_application_tag"

description: "The uztna_application_tag allows CRUD of Uztna Application Tag resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------------

Resource volterra_uztna_application_tag
=======================================

The Uztna Application Tag allows CRUD of Uztna Application Tag resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Application Tag API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-application-tag) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_application_tag" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
  tag       = ["APPGRP-MKTG"]
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

`tag` - (Required) This field is used to create a tag that will be attached to the application. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_application_tag.
