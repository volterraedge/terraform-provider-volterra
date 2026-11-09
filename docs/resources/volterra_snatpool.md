---

page_title: "Volterra: snatpool"

description: "The snatpool allows CRUD of Snatpool resource on Volterra SaaS"
-----------------------------------------------------------------------------

Resource volterra_snatpool
==========================

The Snatpool allows CRUD of Snatpool resource on Volterra SaaS

~> **Note:** Please refer to [Snatpool API docs](https://docs.cloud.f5.com/docs-v2/api/snatpool) to learn more

Example Usage
-------------

```hcl
resource "volterra_snatpool" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
  members   = ["members"]
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

`full_path` - (Optional) Specifies the canonical, system-wide unique path combining partition, optional sub_path, and name (`String`).

`members` - (Required) Specifies a translation address (IPv4 or IPv6) to add to or delete from a SNAT pool. (`List of String`).

`sub_path` - (Optional) Optional folder under the partition where the object resides (`String`).

`traffic_acceleration_status` - (Optional) Defines the available traffic-acceleration allocation modes for the SNAT pool. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured snatpool.
