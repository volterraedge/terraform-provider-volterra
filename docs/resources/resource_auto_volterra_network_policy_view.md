---

page_title: "Volterra: network_policy_view"

description: "The network_policy_view allows CRUD of Network Policy View resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------

Resource volterra_network_policy_view
=====================================

The Network Policy View allows CRUD of Network Policy View resource on Volterra SaaS

~> **Note:** Please refer to [Network Policy View API docs](https://volterra.io/docs/api/views-network-policy-view) to learn more

Example Usage
-------------

```hcl
resource "volterra_network_policy_view" "example" {
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

`egress_rules` - (Optional) Ordered list of rules applied to connections from policy endpoints.. See [Egress Rules ](#egress-rules) below for details.

`endpoint` - (Optional) Policy is for set of endpoints defined, rules are applied to connections to or from these endpoints.. See [Endpoint ](#endpoint) below for details.

`ingress_rules` - (Optional) Ordered list of rules applied to connections to policy endpoints.. See [Ingress Rules ](#ingress-rules) below for details.

### Egress Rules

Ordered list of rules applied to connections from policy endpoints..

`action` - (Optional) Action to be taken at rule match. Currently supported actions are Allow and Deny (`String`).

`keys` - (Optional) can talk to "db" in site "abc" and can not talk to "db" in site "xyz" (`String`).

`any` - (Optional) Any Endpoint that matches 0/0 ip prefix (bool).

`inside_endpoints` - (Optional) All ip prefixes that are reachable via inside interfaces are chosen as Endpoints (bool).

`ip_prefix_set` - (Optional) Reference to object which represents list of IP prefixes that will be referred as remote endpoint. See [Ip Prefix Set ](#ip-prefix-set) below for details.

`label_selector` - (Optional) local end point is set of prefixes determined by label selector expression. See [Label Selector ](#label-selector) below for details.

`namespace` - (Optional) All ip prefixes that are of a namespace are chosen as Endpoints (`String`).

`outside_endpoints` - (Optional) All ip prefixes that are reachable via outside interfaces are chosen as Endpoints (bool).

`prefix_list` - (Optional) For Ingress rules: To these endpoints from remote endpoints these ip prefixes are destination IPs.. See [Prefix List ](#prefix-list) below for details.

`rule_description` - (Optional) Human readable description for the rule (`String`).

`rule_name` - (Required) Rule Name that will be used to query metrics for this rule. (`String`).

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

`keys` - (Optional) can talk to "db" in site "abc" and can not talk to "db" in site "xyz" (`String`).

`any` - (Optional) Any Endpoint that matches 0/0 ip prefix (bool).

`inside_endpoints` - (Optional) All ip prefixes that are reachable via inside interfaces are chosen as Endpoints (bool).

`ip_prefix_set` - (Optional) Reference to object which represents list of IP prefixes that will be referred as remote endpoint. See [Ip Prefix Set ](#ip-prefix-set) below for details.

`label_selector` - (Optional) local end point is set of prefixes determined by label selector expression. See [Label Selector ](#label-selector) below for details.

`namespace` - (Optional) All ip prefixes that are of a namespace are chosen as Endpoints (`String`).

`outside_endpoints` - (Optional) All ip prefixes that are reachable via outside interfaces are chosen as Endpoints (bool).

`prefix_list` - (Optional) For Ingress rules: To these endpoints from remote endpoints these ip prefixes are destination IPs.. See [Prefix List ](#prefix-list) below for details.

`rule_description` - (Optional) Human readable description for the rule (`String`).

`rule_name` - (Required) Rule Name that will be used to query metrics for this rule. (`String`).

`all_tcp_traffic` - (Optional) Select all TCP traffic to match (bool).

`all_traffic` - (Optional) Select all traffic to match (bool).

`all_udp_traffic` - (Optional) Select all UDP traffic to match (bool).

`applications` - (Optional) Select Application traffic to match. See [Applications ](#applications) below for details.

`protocol_port_range` - (Optional) Select specific protocol and port ranges traffic to match. See [Protocol Port Range ](#protocol-port-range) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_policy_view.
