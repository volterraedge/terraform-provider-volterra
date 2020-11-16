---

page_title: "Volterra: service_policy"

description: "The service_policy allows CRUD of Service Policy resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_service_policy
================================

The Service Policy allows CRUD of Service Policy resource on Volterra SaaS

~> **Note:** Please refer to [Service Policy API docs](https://volterra.io/docs/api/service-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_service_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
  algo      = ["algo"]

  // One of the arguments from this list "any_server server_name server_selector server_name_matcher" must be set
  server_name = "database.production.customer.volterra.us"
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

`algo` - (Required) - DENY_OVERRIDES Rules with a DENY action are evaluated prior to rules with an ALLOW action (`String`).

`port_matcher` - (Optional) The list of port ranges to which the destination port should belong. In case of an HTTP Connect, the port is extracted from the desired destination.. See [Port Matcher ](#port-matcher) below for details.

`rules` - (Optional) The order of evaluation of the rules depends on the rule combining algorithm.. See [ref](#ref) below for details.

`any_server` - (Optional) Any Server (bool).

`server_name` - (Optional) The predicate evaluates to true if any of the actual names is the same as the expected server name. (`String`).

`server_name_matcher` - (Optional) The predicate evaluates to true if any of the server's actual names match any of the exact values or regular expressions in the server name matcher.. See [Server Name Matcher ](#server-name-matcher) below for details.

`server_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the server labels.. See [Server Selector ](#server-selector) below for details.

### Port Matcher

The list of port ranges to which the destination port should belong. In case of an HTTP Connect, the port is extracted from the desired destination..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`ports` - (Required) to be part of the range. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Server Name Matcher

The predicate evaluates to true if any of the server's actual names match any of the exact values or regular expressions in the server name matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Server Selector

The predicate evaluates to true if the expressions in the label selector are true for the server labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured service_policy.
