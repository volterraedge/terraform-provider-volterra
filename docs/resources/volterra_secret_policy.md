











---
page_title: "Volterra: secret_policy"
description: "The secret_policy allows CRUD of Secret Policy  resource on Volterra SaaS"
---
# Resource volterra_secret_policy

The Secret Policy  allows CRUD of Secret Policy  resource on Volterra SaaS

~> **Note:** Please refer to [Secret Policy  API docs](https://docs.cloud.f5.com/docs-v2/api/secret-policy) to learn more

## Example Usage

```hcl
resource "volterra_secret_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
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
`algo` - (Optional) - DENY_OVERRIDES Rules with a DENY action are evaluated prior to rules with an ALLOW action (`String`).(Deprecated)



`allow_f5xc` - (Optional) if allow_f5xc is set to true, it allows relevant F5XC infrastructure services to decrypt the secret encrypted using this policy. (`Bool`).



`decrypt_cache_timeout` - (Optional) Value for this parameter is a string ending in the suffix "s" (indicating seconds), suffix "m" (indicating minutes) or suffix "h" (indicating hours) (`String`).




`legacy_rule_list` - (Optional) x-displayName: "Legacy Rule List". See [Rule Choice Legacy Rule List ](#rule-choice-legacy-rule-list) below for details.(Deprecated)
		




`rule_list` - (Optional) x-displayName: "Custom Rule List". See [Rule Choice Rule List ](#rule-choice-rule-list) below for details.
		


		


		





		






		







		





		






`rules` - (Optional) The order of evaluation of the rules depends on the rule combining algorithm.. See [ref](#ref) below for details.(Deprecated)


### Client Choice Client Name Matcher 

 The predicate evaluates to true if any of the client's actual names match any of the exact values or regular expressions in the client name matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).



### Client Choice Client Selector 

 The predicate evaluates to true if the expressions in the label selector are true for the client labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).



### Ref 


Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).



### Rule Choice Legacy Rule List 

 x-displayName: "Legacy Rule List".

`rules` - (Optional) The order of evaluation of the rules depends on the rule combining algorithm.. See [ref](#ref) below for details.



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

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).



### Rules Spec 

 Specification for the rule including match predicates and actions..

`action` - (Required) Action to be enforced if all the predicates evaluates to true. (`String`).



###### One of the arguments from this list "client_name, client_selector, client_name_matcher" must be set

`client_name` - (Optional) This predicate evaluates to true if client name matches the configured name (`String`).


`client_name_matcher` - (Optional) The predicate evaluates to true if any of the client's actual names match any of the exact values or regular expressions in the client name matcher.. See [Client Choice Client Name Matcher ](#client-choice-client-name-matcher) below for details.


`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.


`label_matcher` - (Optional) The values of any other labels do not matter.. See [Spec Label Matcher ](#spec-label-matcher) below for details.(Deprecated)



### Spec Label Matcher 

 The values of any other labels do not matter..

`keys` - (Optional) The list of label key names that have to match (`String`).



## Attribute Reference

* `id` - This is the id of the configured secret_policy.

