---

page_title: "Volterra: secret_policy"

description: "The secret_policy allows CRUD of Secret Policy resource on Volterra SaaS"
---------------------------------------------------------------------------------------

Resource volterra_secret_policy
===============================

The Secret Policy allows CRUD of Secret Policy resource on Volterra SaaS

~> **Note:** Please refer to [Secret Policy API docs](https://docs.cloud.f5.com/docs-v2/api/secret-policy) to learn more

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

`allow_f5xc` - (Optional) if allow_f5xc is set to true, it allows relevant F5XC infrastructure services to decrypt the secret encrypted using this policy. (`Bool`).

`decrypt_cache_timeout` - (Optional) Value for this parameter is a string ending in the suffix "s" (indicating seconds), suffix "m" (indicating minutes) or suffix "h" (indicating hours) (`String`).

###### One of the arguments from this list "rule_list" can be set

`rule_list` - (Optional) x-displayName: "Custom Rule List". See [Rule Choice Rule List ](#rule-choice-rule-list) below for details.

### Client Choice Client Name Matcher

The predicate evaluates to true if any of the client's actual names match any of the exact values or regular expressions in the client name matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Client Choice Client Selector

The predicate evaluates to true if the expressions in the label selector are true for the client labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Rule Choice Rule List

x-displayName: "Custom Rule List".

`rules` - (Optional) Rules are evaluated from top to bottom in the list.. See [Rule List Rules ](#rule-list-rules) below for details.

### Rule List Rules

Rules are evaluated from top to bottom in the list..

`metadata` - (Required) Common attributes for the rule including name and description.. See [Rules Metadata ](#rules-metadata) below for details.

`spec` - (Required) Specification for the rule including match predicates and actions.. See [Rules Spec ](#rules-spec) below for details.

### Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Rules Spec

Specification for the rule including match predicates and actions..

`action` - (Required) Action to be enforced if all the predicates evaluates to true. (`String`).

###### One of the arguments from this list "client_name, client_name_matcher, client_selector" must be set

`client_name` - (Optional) This predicate evaluates to true if client name matches the configured name (`String`).

`client_name_matcher` - (Optional) The predicate evaluates to true if any of the client's actual names match any of the exact values or regular expressions in the client name matcher.. See [Client Choice Client Name Matcher ](#client-choice-client-name-matcher) below for details.

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured secret_policy.
