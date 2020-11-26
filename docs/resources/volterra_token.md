---

page_title: "Volterra: volterra_token"

description: "The volterra_token runs action api on view sites"
---------------------------------------------------------------

Resource volterra_token
=======================

====================================

volterra_token: Token object is used to manage site admission. User must generate token before provisioning and pass this token to site during it's registration. Invalid tokens are refused and site with invalid token won't be able to join Volterra network. Single token can be used to register multiple sites. If tokens already exists in the tenant, then this resource will use that instead of creating another token.

~> **Note:** Please refer to [Volterra Token API docs](https://volterra.io/docs/api/token) to learn more

Example Usage
-------------

---

```hcl
resource "volterra_token" "example" {
  name       = "token1"
  namespace  = "system"
}

```

### Argument Reference

---

`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).

`description` - (Optional) Human readable description for the object (`String`).

`disable` - (Optional) A value of true will administratively disable the object (`Bool`).

`labels` - (Optional) by selector expression (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

`namespace` - (Required) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).

Attribute Reference
-------------------

---

-	`created_by_tf` - Flag to indicate if token was created by terraform create or it used existing one

-	`id` - Token ID

-	`tenant_name` - Tenant name configured in volterra
