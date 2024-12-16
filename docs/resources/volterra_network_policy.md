---

page_title: "Volterra: network_policy"

description: "The network_policy allows CRUD of Network Policy resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_network_policy
================================

The Network Policy allows CRUD of Network Policy resource on Volterra SaaS

~> **Note:** Please refer to [Network Policy API docs](https://docs.cloud.f5.com/docs-v2/api/network-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_network_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  endpoint {
    // One of the arguments from this list "any inside_endpoints interface label_selector namespace outside_endpoints prefix_list" must be set

    outside_endpoints = true
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

### Endpoint

Policy is for set of endpoints defined, rules are applied to connections to or from these endpoints..

###### One of the arguments from this list "any, inside_endpoints, interface, label_selector, namespace, outside_endpoints, prefix_list" must be set

`any` - (Optional) Any Endpoint that matches 0/0 ip prefix (`Bool`).

`inside_endpoints` - (Optional) All ip prefixes that are reachable via inside interfaces are chosen as Endpoints (`Bool`).

`interface` - (Optional) All ip prefixes that are reachable via an interfaces are chosen as Endpoints. See [ref](#ref) below for details.(Deprecated)

`label_selector` - (Optional) local end point is set of prefixes determined by label selector expression. See [Endpoint Choice Label Selector ](#endpoint-choice-label-selector) below for details.

`namespace` - (Optional) All ip prefixes that are of a namespace are chosen as Endpoints (`String`).(Deprecated)

`outside_endpoints` - (Optional) All ip prefixes that are reachable via outside interfaces are chosen as Endpoints (`Bool`).

`prefix_list` - (Optional) For Ingress rules: To this endpoints from remote endpoints these ip prefixes are destination ip.. See [Endpoint Choice Prefix List ](#endpoint-choice-prefix-list) below for details.

### Rules

Network Policy Rules.

`egress_rules` - (Optional) Ordered list of rules applied to connections from policy endpoints.. See [Rules Egress Rules ](#rules-egress-rules) below for details.

`ingress_rules` - (Optional) Ordered list of rules applied to connections to policy endpoints.. See [Rules Ingress Rules ](#rules-ingress-rules) below for details.

### Egress Rules Adv Action

Enable or disable logging..

`action` - (Optional) Enable or disable logging. (`String`).

### Egress Rules Label Matcher

not specified here, just the label keys. This facilitates reuse of policies across multiple dimensions such as deployment, environment, and location..

`keys` - (Optional) The list of label key names that have to match (`String`).

### Egress Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Endpoint Choice Any

Any Endpoint that matches 0/0 ip prefix.

### Endpoint Choice Inside Endpoints

All ip prefixes that are reachable via inside interfaces are chosen as Endpoints.

### Endpoint Choice Label Selector

local end point is set of prefixes determined by label selector expression.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Endpoint Choice Outside Endpoints

All ip prefixes that are reachable via outside interfaces are chosen as Endpoints.

### Endpoint Choice Prefix List

For Ingress rules: To this endpoints from remote endpoints these ip prefixes are destination ip..

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Ingress Rules Adv Action

Enable or disable logging..

`action` - (Optional) Enable or disable logging. (`String`).

### Ingress Rules Label Matcher

not specified here, just the label keys. This facilitates reuse of policies across multiple dimensions such as deployment, environment, and location..

`keys` - (Optional) The list of label key names that have to match (`String`).

### Ingress Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Other Endpoint Any

Any Endpoint that matches 0/0 ip prefix.

### Other Endpoint Inside Endpoints

All ip prefixes that are reachable via inside interfaces are chosen as Endpoints.

### Other Endpoint Ip Prefix Set

Reference to object which represents list of IP prefixes that will be referred as remote endpoint.

`ref` - (Optional) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Other Endpoint Label Selector

local end point is set of prefixes determined by label selector expression.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Other Endpoint Outside Endpoints

All ip prefixes that are reachable via outside interfaces are chosen as Endpoints.

### Other Endpoint Prefix List

For Ingress rules: To these endpoints from remote endpoints these ip prefixes are destination IPs..

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rules Egress Rules

Ordered list of rules applied to connections from policy endpoints..

`action` - (Optional) Action to be taken at rule match. Currently supported actions are Allow and Deny (`String`).

`adv_action` - (Optional) Enable or disable logging.. See [Egress Rules Adv Action ](#egress-rules-adv-action) below for details.

`keys` - (Optional) can talk to "db" in site "abc" and can not talk to "db" in site "xyz" (`String`).(Deprecated)

`label_matcher` - (Optional) not specified here, just the label keys. This facilitates reuse of policies across multiple dimensions such as deployment, environment, and location.. See [Egress Rules Label Matcher ](#egress-rules-label-matcher) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Egress Rules Metadata ](#egress-rules-metadata) below for details.

###### One of the arguments from this list "any, inside_endpoints, ip_prefix_set, label_selector, namespace, outside_endpoints, prefix_list" can be set

`any` - (Optional) Any Endpoint that matches 0/0 ip prefix (`Bool`).

`inside_endpoints` - (Optional) All ip prefixes that are reachable via inside interfaces are chosen as Endpoints (`Bool`).

`ip_prefix_set` - (Optional) Reference to object which represents list of IP prefixes that will be referred as remote endpoint. See [Other Endpoint Ip Prefix Set ](#other-endpoint-ip-prefix-set) below for details.

`label_selector` - (Optional) local end point is set of prefixes determined by label selector expression. See [Other Endpoint Label Selector ](#other-endpoint-label-selector) below for details.

`namespace` - (Optional) All ip prefixes that are of a namespace are chosen as Endpoints (`String`).(Deprecated)

`outside_endpoints` - (Optional) All ip prefixes that are reachable via outside interfaces are chosen as Endpoints (`Bool`).

`prefix_list` - (Optional) For Ingress rules: To these endpoints from remote endpoints these ip prefixes are destination IPs.. See [Other Endpoint Prefix List ](#other-endpoint-prefix-list) below for details.

`rule_description` - (Optional) Human readable description for the rule (`String`).(Deprecated)

`rule_name` - (Optional) Rule Name that will be used to query metrics for this rule. (`String`).(Deprecated)

###### One of the arguments from this list "all_tcp_traffic, all_traffic, all_udp_traffic, applications, protocol_port_range" can be set

`all_tcp_traffic` - (Optional) Select all TCP traffic to match (`Bool`).

`all_traffic` - (Optional) Select all traffic to match (`Bool`).

`all_udp_traffic` - (Optional) Select all UDP traffic to match (`Bool`).

`applications` - (Optional) Select Application traffic to match. See [Traffic Choice Applications ](#traffic-choice-applications) below for details.

`protocol_port_range` - (Optional) Select specific protocol and port ranges traffic to match. See [Traffic Choice Protocol Port Range ](#traffic-choice-protocol-port-range) below for details.

### Rules Ingress Rules

Ordered list of rules applied to connections to policy endpoints..

`action` - (Optional) Action to be taken at rule match. Currently supported actions are Allow and Deny (`String`).

`adv_action` - (Optional) Enable or disable logging.. See [Ingress Rules Adv Action ](#ingress-rules-adv-action) below for details.

`keys` - (Optional) can talk to "db" in site "abc" and can not talk to "db" in site "xyz" (`String`).(Deprecated)

`label_matcher` - (Optional) not specified here, just the label keys. This facilitates reuse of policies across multiple dimensions such as deployment, environment, and location.. See [Ingress Rules Label Matcher ](#ingress-rules-label-matcher) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Ingress Rules Metadata ](#ingress-rules-metadata) below for details.

###### One of the arguments from this list "any, inside_endpoints, ip_prefix_set, label_selector, namespace, outside_endpoints, prefix_list" can be set

`any` - (Optional) Any Endpoint that matches 0/0 ip prefix (`Bool`).

`inside_endpoints` - (Optional) All ip prefixes that are reachable via inside interfaces are chosen as Endpoints (`Bool`).

`ip_prefix_set` - (Optional) Reference to object which represents list of IP prefixes that will be referred as remote endpoint. See [Other Endpoint Ip Prefix Set ](#other-endpoint-ip-prefix-set) below for details.

`label_selector` - (Optional) local end point is set of prefixes determined by label selector expression. See [Other Endpoint Label Selector ](#other-endpoint-label-selector) below for details.

`namespace` - (Optional) All ip prefixes that are of a namespace are chosen as Endpoints (`String`).(Deprecated)

`outside_endpoints` - (Optional) All ip prefixes that are reachable via outside interfaces are chosen as Endpoints (`Bool`).

`prefix_list` - (Optional) For Ingress rules: To these endpoints from remote endpoints these ip prefixes are destination IPs.. See [Other Endpoint Prefix List ](#other-endpoint-prefix-list) below for details.

`rule_description` - (Optional) Human readable description for the rule (`String`).(Deprecated)

`rule_name` - (Optional) Rule Name that will be used to query metrics for this rule. (`String`).(Deprecated)

###### One of the arguments from this list "all_tcp_traffic, all_traffic, all_udp_traffic, applications, protocol_port_range" can be set

`all_tcp_traffic` - (Optional) Select all TCP traffic to match (`Bool`).

`all_traffic` - (Optional) Select all traffic to match (`Bool`).

`all_udp_traffic` - (Optional) Select all UDP traffic to match (`Bool`).

`applications` - (Optional) Select Application traffic to match. See [Traffic Choice Applications ](#traffic-choice-applications) below for details.

`protocol_port_range` - (Optional) Select specific protocol and port ranges traffic to match. See [Traffic Choice Protocol Port Range ](#traffic-choice-protocol-port-range) below for details.

### Traffic Choice All Tcp Traffic

Select all TCP traffic to match.

### Traffic Choice All Traffic

Select all traffic to match.

### Traffic Choice All Udp Traffic

Select all UDP traffic to match.

### Traffic Choice Applications

Select Application traffic to match.

`applications` - (Optional) Application protocols like HTTP, SNMP (`List of Strings`).

### Traffic Choice Protocol Port Range

Select specific protocol and port ranges traffic to match.

`port_ranges` - (Optional) List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192 (`String`).

`protocol` - (Optional) Values are tcp, udp, and icmp (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_policy.
