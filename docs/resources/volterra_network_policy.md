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

  endpoint {
    // One of the arguments from this list "interface namespace label_selector prefix_list any outside_endpoints inside_endpoints" must be set
    namespace = "namespace"
  }
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

`endpoint` - (Required) Policy is for set of endpoints defined, rules are applied to connections to or from these endpoints.. See [Endpoint ](#endpoint) below for details.

`rules` - (Optional) Network Policy Rules. See [Rules ](#rules) below for details.

### Adv Action

Enable or disable logging..

`action` - (Optional) Enable or disable logging. (`String`).

### All Tcp Traffic

Select all TCP traffic to match.

### All Traffic

Select all traffic to match.

### All Udp Traffic

Select all UDP traffic to match.

### Any

Any Endpoint that matches 0/0 ip prefix.

### Applications

Select Application traffic to match.

`applications` - (Optional) Application protocols like HTTP, SNMP (`List of Strings`).

### Egress Rules

Ordered list of rules applied to connections from policy endpoints..

`action` - (Optional) Action to be taken at rule match. Currently supported actions are Allow and Deny (`String`).

`adv_action` - (Optional) Enable or disable logging.. See [Adv Action ](#adv-action) below for details.

`keys` - (Optional) can talk to "db" in site "abc" and can not talk to "db" in site "xyz" (`String`).

`label_matcher` - (Optional) not specified here, just the label keys. This facilitates reuse of policies across multiple dimensions such as deployment, environment, and location.. See [Label Matcher ](#label-matcher) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`any` - (Optional) Any Endpoint that matches 0/0 ip prefix (bool).

`inside_endpoints` - (Optional) All ip prefixes that are reachable via inside interfaces are chosen as Endpoints (bool).

`ip_prefix_set` - (Optional) Reference to object which represents list of IP prefixes that will be referred as remote endpoint. See [Ip Prefix Set ](#ip-prefix-set) below for details.

`label_selector` - (Optional) local end point is set of prefixes determined by label selector expression. See [Label Selector ](#label-selector) below for details.

`namespace` - (Optional) All ip prefixes that are of a namespace are chosen as Endpoints (`String`).

`outside_endpoints` - (Optional) All ip prefixes that are reachable via outside interfaces are chosen as Endpoints (bool).

`prefix_list` - (Optional) For Ingress rules: To these endpoints from remote endpoints these ip prefixes are destination IPs.. See [Prefix List ](#prefix-list) below for details.

`rule_description` - (Optional) Human readable description for the rule (`String`).

`rule_name` - (Optional) Rule Name that will be used to query metrics for this rule. (`String`).

`all_tcp_traffic` - (Optional) Select all TCP traffic to match (bool).

`all_traffic` - (Optional) Select all traffic to match (bool).

`all_udp_traffic` - (Optional) Select all UDP traffic to match (bool).

`applications` - (Optional) Select Application traffic to match. See [Applications ](#applications) below for details.

`protocol_port_range` - (Optional) Select specific protocol and port ranges traffic to match. See [Protocol Port Range ](#protocol-port-range) below for details.

### Endpoint

Policy is for set of endpoints defined, rules are applied to connections to or from these endpoints..

`any` - (Optional) Any Endpoint that matches 0/0 ip prefix (bool).

`inside_endpoints` - (Optional) All ip prefixes that are reachable via inside interfaces are chosen as Endpoints (bool).

`interface` - (Optional) All ip prefixes that are reachable via an interfaces are chosen as Endpoints. See [ref](#ref) below for details.

`label_selector` - (Optional) local end point is set of prefixes determined by label selector expression. See [Label Selector ](#label-selector) below for details.

`namespace` - (Optional) All ip prefixes that are of a namespace are chosen as Endpoints (`String`).

`outside_endpoints` - (Optional) All ip prefixes that are reachable via outside interfaces are chosen as Endpoints (bool).

`prefix_list` - (Optional) For Ingress rules: To this endpoints from remote endpoints these ip prefixes are destination ip.. See [Prefix List ](#prefix-list) below for details.

### Ingress Rules

Ordered list of rules applied to connections to policy endpoints..

`action` - (Optional) Action to be taken at rule match. Currently supported actions are Allow and Deny (`String`).

`adv_action` - (Optional) Enable or disable logging.. See [Adv Action ](#adv-action) below for details.

`keys` - (Optional) can talk to "db" in site "abc" and can not talk to "db" in site "xyz" (`String`).

`label_matcher` - (Optional) not specified here, just the label keys. This facilitates reuse of policies across multiple dimensions such as deployment, environment, and location.. See [Label Matcher ](#label-matcher) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`any` - (Optional) Any Endpoint that matches 0/0 ip prefix (bool).

`inside_endpoints` - (Optional) All ip prefixes that are reachable via inside interfaces are chosen as Endpoints (bool).

`ip_prefix_set` - (Optional) Reference to object which represents list of IP prefixes that will be referred as remote endpoint. See [Ip Prefix Set ](#ip-prefix-set) below for details.

`label_selector` - (Optional) local end point is set of prefixes determined by label selector expression. See [Label Selector ](#label-selector) below for details.

`namespace` - (Optional) All ip prefixes that are of a namespace are chosen as Endpoints (`String`).

`outside_endpoints` - (Optional) All ip prefixes that are reachable via outside interfaces are chosen as Endpoints (bool).

`prefix_list` - (Optional) For Ingress rules: To these endpoints from remote endpoints these ip prefixes are destination IPs.. See [Prefix List ](#prefix-list) below for details.

`rule_description` - (Optional) Human readable description for the rule (`String`).

`rule_name` - (Optional) Rule Name that will be used to query metrics for this rule. (`String`).

`all_tcp_traffic` - (Optional) Select all TCP traffic to match (bool).

`all_traffic` - (Optional) Select all traffic to match (bool).

`all_udp_traffic` - (Optional) Select all UDP traffic to match (bool).

`applications` - (Optional) Select Application traffic to match. See [Applications ](#applications) below for details.

`protocol_port_range` - (Optional) Select specific protocol and port ranges traffic to match. See [Protocol Port Range ](#protocol-port-range) below for details.

### Inside Endpoints

All ip prefixes that are reachable via inside interfaces are chosen as Endpoints.

### Ip Prefix Set

Reference to object which represents list of IP prefixes that will be referred as remote endpoint.

`ref` - (Optional) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Label Matcher

not specified here, just the label keys. This facilitates reuse of policies across multiple dimensions such as deployment, environment, and location..

`keys` - (Optional) The list of label key names that have to match (`String`).

### Label Selector

local end point is set of prefixes determined by label selector expression.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Outside Endpoints

All ip prefixes that are reachable via outside interfaces are chosen as Endpoints.

### Prefix List

For Ingress rules: To this endpoints from remote endpoints these ip prefixes are destination ip..

`prefixes` - (Required) List of IPv4 prefixes that represent an endpoint (`String`).

### Protocol Port Range

Select specific protocol and port ranges traffic to match.

`port_ranges` - (Optional) List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192 (`String`).

`protocol` - (Optional) Values are tcp, udp, and icmp (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rules

Network Policy Rules.

`egress_rules` - (Optional) Ordered list of rules applied to connections from policy endpoints.. See [Egress Rules ](#egress-rules) below for details.

`ingress_rules` - (Optional) Ordered list of rules applied to connections to policy endpoints.. See [Ingress Rules ](#ingress-rules) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_policy.
