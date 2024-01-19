---

page_title: "Volterra: enhanced_firewall_policy"

description: "The enhanced_firewall_policy allows CRUD of Enhanced Firewall Policy resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------------------

Resource volterra_enhanced_firewall_policy
==========================================

The Enhanced Firewall Policy allows CRUD of Enhanced Firewall Policy resource on Volterra SaaS

~> **Note:** Please refer to [Enhanced Firewall Policy API docs](https://docs.cloud.f5.com/docs/api/enhanced-firewall-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_enhanced_firewall_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "allow_all allowed_sources allowed_destinations deny_all denied_sources denied_destinations rule_list" must be set

  rule_list {
    rules {
      // One of the arguments from this list "deny allow insert_service" must be set
      deny = true

      advanced_action {
        action = "action"
      }

      // One of the arguments from this list "inside_destinations all_slo_vips all_sli_vips destination_aws_subnet_ids destination_label_selector destination_aws_vpc_ids all_destinations destination_prefix_list destination_ip_prefix_set outside_destinations destination_namespace" must be set
      destination_namespace = "destination_namespace"

      label_matcher {
        keys = ["['environment', 'location', 'deployment']"]
      }

      metadata {
        description = "Virtual Host for acmecorp website"
        disable     = true
        name        = "acmecorp-web"
      }

      // One of the arguments from this list "all_sources source_prefix_list outside_sources source_aws_subnet_ids source_ip_prefix_set inside_sources source_namespace source_label_selector source_aws_vpc_ids" must be set
      all_sources = true

      // One of the arguments from this list "all_traffic all_tcp_traffic all_udp_traffic applications protocol_port_range" must be set
      all_traffic = true
    }
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

`allow_all` - (Optional) Allow all connections from any source to any destination (bool).

`allowed_destinations` - (Optional) Allow all connections to list of destinations from any source. See [Allowed Destinations ](#allowed-destinations) below for details.

`allowed_sources` - (Optional) Allow all connections from list of sources to any destination. See [Allowed Sources ](#allowed-sources) below for details.

`denied_destinations` - (Optional) Deny all connections to list of destinations from any source. See [Denied Destinations ](#denied-destinations) below for details.

`denied_sources` - (Optional) Deny all connections from list of sources to any destination. See [Denied Sources ](#denied-sources) below for details.

`deny_all` - (Optional) Deny all connections from any source to any destination (bool).

`rule_list` - (Optional) Custom Enhanced Firewall Policy Rule Selection. See [Rule List ](#rule-list) below for details.

`segment_policy` - (Optional) Skip the configuration or set option as Any to ignore corresponding segment match. See [Segment Policy ](#segment-policy) below for details.

### Advanced Action

Log any connection matching the rule.

`action` - (Optional) Enable or disable logging. (`String`).

### All Destinations

Any address that matches 0/0 ip prefix.

### All Sli Vips

Destination is virtual-ip of all loadbalancer on site-local-inside network.

### All Slo Vips

Destination is virtual-ip of all loadbalancer on site-local-outside network.

### All Sources

Any address that matches 0/0 ip prefix.

### All Tcp Traffic

Select all TCP traffic to match.

### All Traffic

Select all traffic to match.

### All Udp Traffic

Select all UDP traffic to match.

### Allow

Allow any connection matching the rule.

### Allowed Destinations

Allow all connections to list of destinations from any source.

`ipv6_prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

### Allowed Sources

Allow all connections from list of sources to any destination.

`ipv6_prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

### Applications

Select Application traffic to match.

`applications` - (Optional) Application protocols like HTTP, SNMP (`List of Strings`).

### Denied Destinations

Deny all connections to list of destinations from any source.

`ipv6_prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

### Denied Sources

Deny all connections from list of sources to any destination.

`ipv6_prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

### Deny

Deny any connection matching the rule.

### Destination Aws Subnet Ids

Destination is any address in list of AWS Subnets.

`subnet_id` - (Required) List of Subnet Identifiers in AWS (`String`).

### Destination Aws Vpc Ids

Destination is any address in list of AWS VPCs.

`vpc_id` - (Required) List of VPC Identifiers in AWS (`String`).

### Destination Ip Prefix Set

Addresses that match one of the prefix in the ip-prefix-set.

`ref` - (Optional) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Destination Label Selector

These labels can be cloud tags defined on the vpc resource, labels on the global network or others..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Destination Prefix List

Addresses that match one of the prefix in the list.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Dst Any

Traffic is not matched against any segment.

### Dst Segments

Traffic is matched against destination segment in selected segments.

`segments` - (Required) Select list of segments. See [ref](#ref) below for details.

### Insert Service

Send selected traffic to NFV Service of type Palo Alto Networks VM-Series Firewall for inspection.

`nfv_service` - (Required) Select External Service, to which the traffic should be forwarded to. Forwarding to Palo Alto Networks external service is supported.. See [ref](#ref) below for details.

### Inside Destinations

All addresses reachable in site-local inside interfaces.

### Inside Sources

All addresses reachable in site-local inside interfaces.

### Intra Segment

Traffic is matched for source and destination on the same segment.

### Label Matcher

not specified here, just the label keys. This facilitates reuse of policies across multiple dimensions such as deployment, environment, and location..

`keys` - (Optional) The list of label key names that have to match (`String`).

### Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Outside Destinations

All addresses reachable in site-local outside interfaces.

### Outside Sources

All addresses reachable in site-local outside interfaces.

### Protocol Port Range

Select specific protocol and port ranges traffic to match.

`port_ranges` - (Optional) List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192 (`String`).

`protocol` - (Optional) Values are tcp, udp, and icmp (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rule List

Custom Enhanced Firewall Policy Rule Selection.

`rules` - (Required) Ordered List of Enhanced Firewall Policy Rules. See [Rules ](#rules) below for details.

### Rules

Ordered List of Enhanced Firewall Policy Rules.

`allow` - (Optional) Allow any connection matching the rule (bool).

`deny` - (Optional) Deny any connection matching the rule (bool).

`insert_service` - (Optional) Send selected traffic to NFV Service of type Palo Alto Networks VM-Series Firewall for inspection. See [Insert Service ](#insert-service) below for details.

`advanced_action` - (Optional) Log any connection matching the rule. See [Advanced Action ](#advanced-action) below for details.

`all_destinations` - (Optional) Any address that matches 0/0 ip prefix (bool).

`all_sli_vips` - (Optional) Destination is virtual-ip of all loadbalancer on site-local-inside network (bool).

`all_slo_vips` - (Optional) Destination is virtual-ip of all loadbalancer on site-local-outside network (bool).

`destination_aws_subnet_ids` - (Optional) Destination is any address in list of AWS Subnets. See [Destination Aws Subnet Ids ](#destination-aws-subnet-ids) below for details.

`destination_aws_vpc_ids` - (Optional) Destination is any address in list of AWS VPCs. See [Destination Aws Vpc Ids ](#destination-aws-vpc-ids) below for details.

`destination_ip_prefix_set` - (Optional) Addresses that match one of the prefix in the ip-prefix-set. See [Destination Ip Prefix Set ](#destination-ip-prefix-set) below for details.

`destination_label_selector` - (Optional) These labels can be cloud tags defined on the vpc resource, labels on the global network or others.. See [Destination Label Selector ](#destination-label-selector) below for details.

`destination_namespace` - (Optional) All addresses in a namespace (`String`).

`destination_prefix_list` - (Optional) Addresses that match one of the prefix in the list. See [Destination Prefix List ](#destination-prefix-list) below for details.

`inside_destinations` - (Optional) All addresses reachable in site-local inside interfaces (bool).

`outside_destinations` - (Optional) All addresses reachable in site-local outside interfaces (bool).

`label_matcher` - (Optional) not specified here, just the label keys. This facilitates reuse of policies across multiple dimensions such as deployment, environment, and location.. See [Label Matcher ](#label-matcher) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`all_sources` - (Optional) Any address that matches 0/0 ip prefix (bool).

`inside_sources` - (Optional) All addresses reachable in site-local inside interfaces (bool).

`outside_sources` - (Optional) All addresses reachable in site-local outside interfaces (bool).

`source_aws_subnet_ids` - (Optional) Source is any address in list of AWS Subnets. See [Source Aws Subnet Ids ](#source-aws-subnet-ids) below for details.

`source_aws_vpc_ids` - (Optional) Source is any address in list of AWS VPCs. See [Source Aws Vpc Ids ](#source-aws-vpc-ids) below for details.

`source_ip_prefix_set` - (Optional) Addresses that match one of the prefix in the ip-prefix-set. See [Source Ip Prefix Set ](#source-ip-prefix-set) below for details.

`source_label_selector` - (Optional) These labels can be cloud tags defined on the vpc resource, labels on the global network or others.. See [Source Label Selector ](#source-label-selector) below for details.

`source_namespace` - (Optional) All addresses in a namespace (`String`).

`source_prefix_list` - (Optional) Addresses that match one of the prefix in the list. See [Source Prefix List ](#source-prefix-list) below for details.

`all_tcp_traffic` - (Optional) Select all TCP traffic to match (bool).

`all_traffic` - (Optional) Select all traffic to match (bool).

`all_udp_traffic` - (Optional) Select all UDP traffic to match (bool).

`applications` - (Optional) Select Application traffic to match. See [Applications ](#applications) below for details.

`protocol_port_range` - (Optional) Select specific protocol and port ranges traffic to match. See [Protocol Port Range ](#protocol-port-range) below for details.

### Segment Policy

Skip the configuration or set option as Any to ignore corresponding segment match.

`dst_any` - (Optional) Traffic is not matched against any segment (bool).

`dst_segments` - (Optional) Traffic is matched against destination segment in selected segments. See [Dst Segments ](#dst-segments) below for details.

`intra_segment` - (Optional) Traffic is matched for source and destination on the same segment (bool).

`src_any` - (Optional) Traffic is not matched against any segment (bool).

`src_segments` - (Optional) Source traffic is matched against selected segments. See [Src Segments ](#src-segments) below for details.

### Source Aws Subnet Ids

Source is any address in list of AWS Subnets.

`subnet_id` - (Required) List of Subnet Identifiers in AWS (`String`).

### Source Aws Vpc Ids

Source is any address in list of AWS VPCs.

`vpc_id` - (Required) List of VPC Identifiers in AWS (`String`).

### Source Ip Prefix Set

Addresses that match one of the prefix in the ip-prefix-set.

`ref` - (Optional) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Source Label Selector

These labels can be cloud tags defined on the vpc resource, labels on the global network or others..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Source Prefix List

Addresses that match one of the prefix in the list.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Src Any

Traffic is not matched against any segment.

### Src Segments

Source traffic is matched against selected segments.

`segments` - (Required) Select list of segments. See [ref](#ref) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured enhanced_firewall_policy.
