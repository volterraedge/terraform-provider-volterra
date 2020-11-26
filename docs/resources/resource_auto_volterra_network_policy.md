---

page_title: "Volterra: network_policy"

description: "The network_policy allows CRUD of Network Policy resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_network_policy
================================

The Network Policy allows CRUD of Network Policy resource on Volterra SaaS

~> **Note:** Please refer to [Network Policy API docs](https://volterra.io/docs/api/network-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_network_policy" "example" {
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

`egress_rules` - (Optional) If egress rules are not specified or is empty list, then policy will assume default deny and to ANY destination from local endpoint. See [ref](#ref) below for details.

`ingress_rules` - (Optional) If ingress rules are not specified or is empty list, then policy will assume default deny and from ANY destination to local endpoint. See [ref](#ref) below for details.

`prefix` - (Optional) For ingress rules these ip prefixes are destination. See [Prefix ](#prefix) below for details.

`prefix_selector` - (Optional) Multiple repeated expressions are not supported, only first expression is evaluated. See [Prefix Selector ](#prefix-selector) below for details.

### Prefix

For ingress rules these ip prefixes are destination.

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

### Prefix Selector

Multiple repeated expressions are not supported, only first expression is evaluated.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_policy.
