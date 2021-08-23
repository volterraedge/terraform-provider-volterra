---

page_title: "Volterra: secret_policy"

description: "The secret_policy allows CRUD of Secret Policy resource on Volterra SaaS"
---------------------------------------------------------------------------------------

Resource volterra_secret_policy
===============================

The Secret Policy allows CRUD of Secret Policy resource on Volterra SaaS

~> **Note:** Please refer to [Secret Policy API docs](https://volterra.io/docs/api/secret-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_secret_policy" "example" {
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

`algo` - (Optional) - DENY_OVERRIDES Rules with a DENY action are evaluated prior to rules with an ALLOW action (`String`).

`allow_volterra` - (Optional) if allow_volterra is set to true, it allows relevant volterra infrastructure services to decrypt the secret encrypted using this policy. (`Bool`).

`decrypt_cache_timeout` - (Optional) Value for this parameter is a string ending in the suffix "s" (indicating seconds), suffix "m" (indicating minutes) or suffix "h" (indicating hours) (`String`).

`rules` - (Optional) The order of evaluation of the rules depends on the rule combining algorithm.. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured secret_policy.
