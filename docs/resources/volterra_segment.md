---

page_title: "Volterra: segment"

description: "The segment allows CRUD of Segment resource on Volterra SaaS"
---------------------------------------------------------------------------

Resource volterra_segment
=========================

The Segment allows CRUD of Segment resource on Volterra SaaS

~> **Note:** Please refer to [Segment API docs](https://docs.cloud.f5.com/docs-v2/api/segment) to learn more

Example Usage
-------------

```hcl
resource "volterra_segment" "example" {
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

###### One of the arguments from this list "disable, enable" can be set

`disable` - (Optional) x-displayName: "Deny traffic from this segment to Internet" (`Bool`).

`enable` - (Optional) x-displayName: "Allow traffic from this segment to Internet" (`Bool`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured segment.
