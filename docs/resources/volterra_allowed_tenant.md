---

page_title: "Volterra: allowed_tenant"

description: "The allowed_tenant allows CRUD of Allowed Tenant resource on Volterra SaaS"

---

Resource volterra_allowed_tenant
================================

The Allowed Tenant allows CRUD of Allowed Tenant resource on Volterra SaaS

~> **Note:** Please refer to [Allowed Tenant API docs](https://docs.cloud.f5.com/docs/api/allowed-tenant) to learn more

Example Usage
-------------

```hcl
resource "volterra_allowed_tenant" "example" {
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

`allowed_groups` - (Optional) User access from original tenant into an allowed tenant will be associated to underlying roles in this user_group.. See [ref](#ref) below for details.

`tenant_id` - (Optional) NOTE: this is the name of the tenant configuration obj. not UID. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured allowed_tenant.
