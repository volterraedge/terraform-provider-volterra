











---
page_title: "Volterra: enhanced_firewall_policy"
description: "The enhanced_firewall_policy allows CRUD of Enhanced Firewall Policy  resource on Volterra SaaS"
---
# Resource volterra_enhanced_firewall_policy

The Enhanced Firewall Policy  allows CRUD of Enhanced Firewall Policy  resource on Volterra SaaS

~> **Note:** Please refer to [Enhanced Firewall Policy  API docs](https://docs.cloud.f5.com/docs-v2/api/enhanced-firewall-policy) to learn more

## Example Usage

```hcl
resource "volterra_enhanced_firewall_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "rule_list allow_all allowed_sources allowed_destinations deny_all denied_sources denied_destinations" must be set

  allowed_destinations {
    ipv6_prefix = ["[2001:db8::1::/112, 2001::db8::2::/112]"]

    prefix = ["[192.168.1.0/24, 192.168.2.0/24]\""]
  }
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


`allow_all` - (Optional) Allow all connections from any source to any destination (`Bool`).


`allowed_destinations` - (Optional) Allow all connections to list of destinations from any source. See [Rule Choice Allowed Destinations ](#rule-choice-allowed-destinations) below for details.
		





`allowed_sources` - (Optional) Allow all connections from list of sources to any destination. See [Rule Choice Allowed Sources ](#rule-choice-allowed-sources) below for details.
		





`denied_destinations` - (Optional) Deny all connections to list of destinations from any source. See [Rule Choice Denied Destinations ](#rule-choice-denied-destinations) below for details.
		





`denied_sources` - (Optional) Deny all connections from list of sources to any destination. See [Rule Choice Denied Sources ](#rule-choice-denied-sources) below for details.
		





`deny_all` - (Optional) Deny all connections from any source to any destination (`Bool`).


`rule_list` - (Optional) Custom Enhanced Firewall Policy Rule Selection. See [Rule Choice Rule List ](#rule-choice-rule-list) below for details.
		


		




		




		




		





		





		




		




		




		





		





		





		






		






		




		




		



		







		




		




		




		





		





		





		






		








		




		




		




		





		









`segment_policy` - (Optional) Skip the configuration or set option as Any to ignore corresponding segment match. See [Segment Policy ](#segment-policy) below for details.




		




		





		






		




		






### Segment Policy 

 Skip the configuration or set option as Any to ignore corresponding segment match.




###### One of the arguments from this list "dst_any, intra_segment, dst_segments" can be set

`dst_any` - (Optional) Traffic is not matched against any segment (`Bool`).


`dst_segments` - (Optional) Traffic is matched against destination segment in selected segments. See [Dst Segment Choice Dst Segments ](#dst-segment-choice-dst-segments) below for details.


`intra_segment` - (Optional) Traffic is matched for source and destination on the same segment (`Bool`).





###### One of the arguments from this list "src_any, src_segments" can be set

`src_any` - (Optional) Traffic is not matched against any segment (`Bool`).


`src_segments` - (Optional) Source traffic is matched against selected segments. See [Src Segment Choice Src Segments ](#src-segment-choice-src-segments) below for details.




### Action Choice Allow 

 Allow any connection matching the rule.



### Action Choice Deny 

 Deny any connection matching the rule.



### Action Choice Insert Service 

 Send selected traffic to NFV Service of type Palo Alto Networks VM-Series Firewall for inspection.

`nfv_service` - (Required) Select External Service, to which the traffic should be forwarded to. Forwarding to Palo Alto Networks external service is supported.. See [ref](#ref) below for details.



### Destination Choice All Destinations 

 Any address that matches 0/0 ip prefix.



### Destination Choice All Sli Vips 

 Destination is virtual-ip of all loadbalancer on site-local-inside network.



### Destination Choice All Slo Vips 

 Destination is virtual-ip of all loadbalancer on site-local-outside network.



### Destination Choice Destination Aws Subnet Ids 

 Destination is any address in list of AWS Subnets.

`subnet_id` - (Required) List of Subnet Identifiers in AWS (`String`).



### Destination Choice Destination Aws Vpc Ids 

 Destination is any address in list of AWS VPCs.

`vpc_id` - (Required) List of VPC Identifiers in AWS (`String`).



### Destination Choice Destination Ip Prefix Set 

 Addresses that match one of the prefix in the ip-prefix-set.

`ref` - (Optional) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.



### Destination Choice Destination Label Selector 

 These labels can be cloud tags defined on the vpc resource, labels on the global network or others..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).



### Destination Choice Destination Prefix List 

 Addresses that match one of the prefix in the list.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).



### Destination Choice Inside Destinations 

 All addresses reachable in site-local inside interfaces.



### Destination Choice Outside Destinations 

 All addresses reachable in site-local outside interfaces.



### Dst Segment Choice Dst Any 

 Traffic is not matched against any segment.



### Dst Segment Choice Dst Segments 

 Traffic is matched against destination segment in selected segments.

`segments` - (Required) Select list of segments. See [ref](#ref) below for details.



### Dst Segment Choice Intra Segment 

 Traffic is matched for source and destination on the same segment.



### Ref 


Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).



### Rule Choice Allowed Destinations 

 Allow all connections to list of destinations from any source.

`ipv6_prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).



### Rule Choice Allowed Sources 

 Allow all connections from list of sources to any destination.

`ipv6_prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).



### Rule Choice Denied Destinations 

 Deny all connections to list of destinations from any source.

`ipv6_prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).



### Rule Choice Denied Sources 

 Deny all connections from list of sources to any destination.

`ipv6_prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).



### Rule Choice Rule List 

 Custom Enhanced Firewall Policy Rule Selection.

`rules` - (Required) Ordered List of Enhanced Firewall Policy Rules. See [Rule List Rules ](#rule-list-rules) below for details.



### Rule List Rules 

 Ordered List of Enhanced Firewall Policy Rules.



###### One of the arguments from this list "deny, allow, insert_service" must be set

`allow` - (Optional) Allow any connection matching the rule (`Bool`).


`deny` - (Optional) Deny any connection matching the rule (`Bool`).


`insert_service` - (Optional) Send selected traffic to NFV Service of type Palo Alto Networks VM-Series Firewall for inspection. See [Action Choice Insert Service ](#action-choice-insert-service) below for details.


`advanced_action` - (Optional) Log any connection matching the rule. See [Rules Advanced Action ](#rules-advanced-action) below for details.



###### One of the arguments from this list "inside_destinations, destination_label_selector, all_slo_vips, all_sli_vips, destination_aws_vpc_ids, all_destinations, destination_ip_prefix_set, destination_namespace, destination_aws_subnet_ids, destination_prefix_list, outside_destinations" must be set

`all_destinations` - (Optional) Any address that matches 0/0 ip prefix (`Bool`).


`all_sli_vips` - (Optional) Destination is virtual-ip of all loadbalancer on site-local-inside network (`Bool`).


`all_slo_vips` - (Optional) Destination is virtual-ip of all loadbalancer on site-local-outside network (`Bool`).


`destination_aws_subnet_ids` - (Optional) Destination is any address in list of AWS Subnets. See [Destination Choice Destination Aws Subnet Ids ](#destination-choice-destination-aws-subnet-ids) below for details.


`destination_aws_vpc_ids` - (Optional) Destination is any address in list of AWS VPCs. See [Destination Choice Destination Aws Vpc Ids ](#destination-choice-destination-aws-vpc-ids) below for details.


`destination_ip_prefix_set` - (Optional) Addresses that match one of the prefix in the ip-prefix-set. See [Destination Choice Destination Ip Prefix Set ](#destination-choice-destination-ip-prefix-set) below for details.


`destination_label_selector` - (Optional) These labels can be cloud tags defined on the vpc resource, labels on the global network or others.. See [Destination Choice Destination Label Selector ](#destination-choice-destination-label-selector) below for details.


`destination_namespace` - (Optional) All addresses in a namespace (`String`).(Deprecated)


`destination_prefix_list` - (Optional) Addresses that match one of the prefix in the list. See [Destination Choice Destination Prefix List ](#destination-choice-destination-prefix-list) below for details.


`inside_destinations` - (Optional) All addresses reachable in site-local inside interfaces (`Bool`).


`outside_destinations` - (Optional) All addresses reachable in site-local outside interfaces (`Bool`).


`label_matcher` - (Optional) not specified here, just the label keys. This facilitates reuse of policies across multiple dimensions such as deployment, environment, and location.. See [Rules Label Matcher ](#rules-label-matcher) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Rules Metadata ](#rules-metadata) below for details.



###### One of the arguments from this list "source_aws_subnet_ids, all_sources, source_ip_prefix_set, source_namespace, source_label_selector, source_aws_vpc_ids, source_prefix_list, inside_sources, outside_sources" must be set

`all_sources` - (Optional) Any address that matches 0/0 ip prefix (`Bool`).


`inside_sources` - (Optional) All addresses reachable in site-local inside interfaces (`Bool`).


`outside_sources` - (Optional) All addresses reachable in site-local outside interfaces (`Bool`).


`source_aws_subnet_ids` - (Optional) Source is any address in list of AWS Subnets. See [Source Choice Source Aws Subnet Ids ](#source-choice-source-aws-subnet-ids) below for details.


`source_aws_vpc_ids` - (Optional) Source is any address in list of AWS VPCs. See [Source Choice Source Aws Vpc Ids ](#source-choice-source-aws-vpc-ids) below for details.


`source_ip_prefix_set` - (Optional) Addresses that match one of the prefix in the ip-prefix-set. See [Source Choice Source Ip Prefix Set ](#source-choice-source-ip-prefix-set) below for details.


`source_label_selector` - (Optional) These labels can be cloud tags defined on the vpc resource, labels on the global network or others.. See [Source Choice Source Label Selector ](#source-choice-source-label-selector) below for details.


`source_namespace` - (Optional) All addresses in a namespace (`String`).(Deprecated)


`source_prefix_list` - (Optional) list contains sublist of both v4 and v6 prfix list. See [Source Choice Source Prefix List ](#source-choice-source-prefix-list) below for details.




###### One of the arguments from this list "applications, protocol_port_range, all_traffic, all_tcp_traffic, all_udp_traffic" must be set

`all_tcp_traffic` - (Optional) Select all TCP traffic to match (`Bool`).


`all_traffic` - (Optional) Select all traffic to match (`Bool`).


`all_udp_traffic` - (Optional) Select all UDP traffic to match (`Bool`).


`applications` - (Optional) Select Application traffic to match. See [Traffic Choice Applications ](#traffic-choice-applications) below for details.


`protocol_port_range` - (Optional) Select specific protocol and port ranges traffic to match. See [Traffic Choice Protocol Port Range ](#traffic-choice-protocol-port-range) below for details.




### Rules Advanced Action 

 Log any connection matching the rule.

`action` - (Optional) Enable or disable logging. (`String`).



### Rules Label Matcher 

 not specified here, just the label keys. This facilitates reuse of policies across multiple dimensions such as deployment, environment, and location..

`keys` - (Optional) The list of label key names that have to match (`String`).



### Rules Metadata 

 Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).



### Source Choice All Sources 

 Any address that matches 0/0 ip prefix.



### Source Choice Inside Sources 

 All addresses reachable in site-local inside interfaces.



### Source Choice Outside Sources 

 All addresses reachable in site-local outside interfaces.



### Source Choice Source Aws Subnet Ids 

 Source is any address in list of AWS Subnets.

`subnet_id` - (Required) List of Subnet Identifiers in AWS (`String`).



### Source Choice Source Aws Vpc Ids 

 Source is any address in list of AWS VPCs.

`vpc_id` - (Required) List of VPC Identifiers in AWS (`String`).



### Source Choice Source Ip Prefix Set 

 Addresses that match one of the prefix in the ip-prefix-set.

`ref` - (Optional) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.



### Source Choice Source Label Selector 

 These labels can be cloud tags defined on the vpc resource, labels on the global network or others..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).



### Source Choice Source Prefix List 

 list contains sublist of both v4 and v6 prfix list.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).



### Src Segment Choice Src Any 

 Traffic is not matched against any segment.



### Src Segment Choice Src Segments 

 Source traffic is matched against selected segments.

`segments` - (Required) Select list of segments. See [ref](#ref) below for details.



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



## Attribute Reference

* `id` - This is the id of the configured enhanced_firewall_policy.

