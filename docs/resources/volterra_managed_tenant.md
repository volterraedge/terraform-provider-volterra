---

page_title: "Volterra: managed_tenant"

description: "The managed_tenant allows CRUD of Managed Tenant resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_managed_tenant
================================

The Managed Tenant allows CRUD of Managed Tenant resource on Volterra SaaS

~> **Note:** Please refer to [Managed Tenant API docs](https://volterra.io/docs/api/managed-tenant) to learn more

Example Usage
-------------

```hcl
resource "volterra_managed_tenant" "example" {
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

`groups` - (Optional) List of local user group association to user groups in the managed tenant specified in the tenant_choice.. See [Groups ](#groups) below for details.

`all_tenants` - (Optional) Internal option only. (bool).

`tenant_id` - (Optional) NOTE: this is the name of the tenant configuration obj. not UID. (`String`).

`tenant_regex` - (Optional) Internal option only. (`String`).

### Groups

List of local user group association to user groups in the managed tenant specified in the tenant_choice..

`group` - (Optional) User should be member of this group to gain access into managed tenant.. See [ref](#ref) below for details.

`managed_tenant_groups` - (Optional) access managed tenant, underlying roles from managed tenant will be applied to user. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured managed_tenant.
