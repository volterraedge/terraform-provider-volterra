---

page_title: "Volterra: network_policy_rule"

description: "The network_policy_rule allows CRUD of Network Policy Rule resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------

Resource volterra_network_policy_rule
=====================================

The Network Policy Rule allows CRUD of Network Policy Rule resource on Volterra SaaS

~> **Note:** Please refer to [Network Policy Rule API docs](https://docs.cloud.f5.com/docs/api/network-policy-rule) to learn more

Example Usage
-------------

```hcl
resource "volterra_network_policy_rule" "example" {
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

`action` - (Optional) Action to be taken at rule match. Currently supported actions are Allow and Deny (`String`).

`advanced_action` - (Optional) Enable or disable logging.. See [Advanced Action ](#advanced-action) below for details.

`label_matcher` - (Optional) List of label keys to be matched in prefix_selector configured in remote_endpoint. See [Label Matcher ](#label-matcher) below for details.

`ports` - (Optional) List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192 (`List of String`).

`protocol` - (Optional) Values are tcp, udp, and icmp (`String`).

`ip_prefix_set` - (Optional) Reference to object which represents list of IP prefixes that will be referred as remote endpoint. See [Ip Prefix Set ](#ip-prefix-set) below for details.

`prefix` - (Optional) these IP prefixes are destination. See [Prefix ](#prefix) below for details.

`prefix_selector` - (Optional) Only first expression is selected even though LabelSelectorType can provide multiple. See [Prefix Selector ](#prefix-selector) below for details.

### Advanced Action

Enable or disable logging..

`action` - (Optional) Enable or disable logging. (`String`).

### Ip Prefix Set

Reference to object which represents list of IP prefixes that will be referred as remote endpoint.

`ref` - (Optional) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Label Matcher

List of label keys to be matched in prefix_selector configured in remote_endpoint.

`keys` - (Optional) The list of label key names that have to match (`String`).

### Prefix

these IP prefixes are destination.

`ipv6_prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

### Prefix Selector

Only first expression is selected even though LabelSelectorType can provide multiple.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_policy_rule.
