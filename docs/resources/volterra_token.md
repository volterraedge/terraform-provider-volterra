---

page_title: "Volterra: token"

description: "The token allows CRUD of Token resource on Volterra SaaS"

---

Resource volterra_token
=======================

The Token allows CRUD of Token resource on Volterra SaaS

~> **Note:** Please refer to [Token API docs](https://docs.cloud.f5.com/docs/api/token) to learn more

Example Usage
-------------

```hcl
resource "volterra_token" "example" {
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

`type` - (Optional) Token type has two values, 0 or 1 (0 refers to a NORMAL token and 1 JWT token). For Secure Mesh Site v2 UI, JWT is used, other workflow sticks to normal.Single token can be used to register multiple sites. (`int`\)

`site_name` - (Optional) Only used for JWT token, to add cluster_name to JWT token. (`String`\)

### Spec Argument Reference

Attribute Reference
-------------------

-	`id` - This is the id of the configured token.
