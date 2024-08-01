











---
page_title: "Volterra: secret_policy_rule"
description: "The secret_policy_rule allows CRUD of Secret Policy Rule  resource on Volterra SaaS"
---
# Resource volterra_secret_policy_rule

The Secret Policy Rule  allows CRUD of Secret Policy Rule  resource on Volterra SaaS

~> **Note:** Please refer to [Secret Policy Rule  API docs](https://docs.cloud.f5.com/docs-v2/api/secret-policy-rule) to learn more

## Example Usage

```hcl
resource "volterra_secret_policy_rule" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
  action    = ["action"]

  // One of the arguments from this list "client_name_matcher client_name client_selector" must be set

  client_name = "ver.re01.int.ves.io"
}

```

## Argument Reference

### Metadata Argument Reference
`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).


`description` - (Optional) Human readable description for the object (`String`).


`disable` - (Optional) A value of true will administratively disable the object (`Bool`).


`labels` - (Optional) by selector expression (`String`).


`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).


`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).



### Spec Argument Reference
`action` - (Required) Action to be enforced if all the predicates evaluates to true. (`String`).




`client_name` - (Optional) This predicate evaluates to true if client name matches the configured name (`String`).


`client_name_matcher` - (Optional) The predicate evaluates to true if any of the client's actual names match any of the exact values or regular expressions in the client name matcher.. See [Client Choice Client Name Matcher ](#client-choice-client-name-matcher) below for details.
		





`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.
		






`label_matcher` - (Optional) The values of any other labels do not matter.. See [Label Matcher ](#label-matcher) below for details.(Deprecated)




### Label Matcher 

 The values of any other labels do not matter..

`keys` - (Optional) The list of label key names that have to match (`String`).



### Client Choice Client Name Matcher 

 The predicate evaluates to true if any of the client's actual names match any of the exact values or regular expressions in the client name matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).



### Client Choice Client Selector 

 The predicate evaluates to true if the expressions in the label selector are true for the client labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).



## Attribute Reference

* `id` - This is the id of the configured secret_policy_rule.

