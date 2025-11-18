---

page_title: "Volterra: nat_policy"

description: "The nat_policy allows CRUD of Nat Policy resource on Volterra SaaS"
---------------------------------------------------------------------------------

Resource volterra_nat_policy
============================

The Nat Policy allows CRUD of Nat Policy resource on Volterra SaaS

~> **Note:** Please refer to [Nat Policy API docs](https://docs.cloud.f5.com/docs-v2/api/nat-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_nat_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "site" must be set

  site {
    refs {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }
  rules {
    action {
      // One of the arguments from this list "dynamic virtual_cidr" can be set

      virtual_cidr = "virtual_cidr"
    }

    criteria {
      destination_cidr = ["1.1.1.0/24 or 2001::10/64"]

      destination_port {
        // One of the arguments from this list "no_port_match port port_ranges" can be set

        port = "6443"
      }

      // One of the arguments from this list "segment site_local_inside_network site_local_network virtual_network" can be set

      virtual_network {
        refs {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
      }
      protocol = "ALL"

      // One of the arguments from this list "any icmp tcp udp" can be set

      any = true
      source_cidr = ["1.1.1.0/24 or 2001:10/64"]
      source_port {
        // One of the arguments from this list "no_port_match port port_ranges" can be set

        port = "6443"
      }
    }

    // One of the arguments from this list "disable enable" must be set

    enable = true
    name = "NAT to Internet"

    // One of the arguments from this list "cloud_connect network_interface node_interface segment virtual_network" must be set

    network_interface {
      refs {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }
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

###### One of the arguments from this list "site" must be set

`site` - (Optional) Site reference. See [Applies To Choice Site ](#applies-to-choice-site) below for details.

`rules` - (Required) List of rules to apply under the NAT Policy. Rule that matches first would be applied. See [Rules ](#rules) below for details.

### Rules

List of rules to apply under the NAT Policy. Rule that matches first would be applied.

`action` - (Required) Action to apply if rule is applied. See [Rules Action ](#rules-action) below for details.

`criteria` - (Optional) Criteria to match on the packet to apply the Action. See [Rules Criteria ](#rules-criteria) below for details.

###### One of the arguments from this list "disable, enable" must be set

`disable` - (Optional) Disables the rule (`Bool`).

`enable` - (Optional) Enables the rule (`Bool`).

`name` - (Required) Name of the Rule (`String`).

###### One of the arguments from this list "cloud_connect, network_interface, node_interface, segment, virtual_network" must be set

`cloud_connect` - (Optional) NAT rule is applied to packet coming from cloud connect. See [Scope Choice Cloud Connect ](#scope-choice-cloud-connect) below for details.

`network_interface` - (Optional) NAT rule is applied to packet coming from interface. See [Scope Choice Network Interface ](#scope-choice-network-interface) below for details.

`node_interface` - (Optional) NAT rule is applied to packet coming from one or more interfaces of nodes. See [Scope Choice Node Interface ](#scope-choice-node-interface) below for details.(Deprecated)

`segment` - (Optional) NAT rule is applied to packet in the segment. See [Scope Choice Segment ](#scope-choice-segment) below for details.

`virtual_network` - (Optional) NAT rule is applied to packet in the virtual network. See [Scope Choice Virtual Network ](#scope-choice-virtual-network) below for details.

### Applies To Choice Site

Site reference.

`refs` - (Required) Reference to Site Object. See [ref](#ref) below for details.

### Criteria Destination Port

Destination port of the packet to match.

###### One of the arguments from this list "no_port_match, port, port_ranges" can be set

`no_port_match` - (Optional) Disable matching of ports (`Bool`).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Criteria Source Port

Source port of the packet to match.

###### One of the arguments from this list "no_port_match, port, port_ranges" can be set

`no_port_match` - (Optional) Disable matching of ports (`Bool`).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Enable Choice Disable

Disables the rule.

### Enable Choice Enable

Enables the rule.

### Network Choice Segment

When there is no segment connector, this field need not be specified. When there is segment connector configured and if packet is destined to destination segment, destination segment needs to be configured here.

`refs` - (Required) Reference to Segment Object. See [ref](#ref) below for details.

`virtual_networks` - (Optional) Internally used to resolve segment ref to networks.. See [ref](#ref) below for details.(Deprecated)

### Network Choice Site Local Inside Network

x-displayName: "Site Local Inside Virtual Network".

### Network Choice Site Local Network

x-displayName: "Site Local Outside Virtual Network".

### Network Choice Virtual Network

When there is no network connector, this field need not be specified. When there is network connector configured and if packet is destined to destination network, destination network needs to be configured here.

`refs` - (Optional) Reference to virtual network. See [ref](#ref) below for details.

### Node Interface List

On a multinode site, this list holds the nodes and corresponding networking_interface.

`interface` - (Optional) Interface reference on this node. See [ref](#ref) below for details.

`node` - (Optional) Node name on this site (`String`).

### Pool Choice Elastic Ips

Use Cloud Elastic IP Object as SNAT Pool that is routable in Internet in Cloud Environments.

`refs` - (Required) Reference to one or more cloud elastic ip objects. See [ref](#ref) below for details.

### Pool Choice Pools

Use the configured Pool as SNAT Pool.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Port Match No Port Match

Disable matching of ports.

### Protocol Choice Any

Match Any Protocol.

### Protocol Choice Icmp

Match ICMP Protocol.

### Protocol Choice Tcp

Match TCP Protocol.

`destination_port` - (Optional) Destination port of the packet to match. See [Tcp Destination Port ](#tcp-destination-port) below for details.

`source_port` - (Optional) Source port of the packet to match. See [Tcp Source Port ](#tcp-source-port) below for details.

### Protocol Choice Udp

Match UDP Protocol.

`destination_port` - (Optional) Destination port of the packet to match. See [Udp Destination Port ](#udp-destination-port) below for details.

`source_port` - (Optional) Source port of the packet to match. See [Udp Source Port ](#udp-source-port) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rules Action

Action to apply if rule is applied.

###### One of the arguments from this list "dynamic, virtual_cidr" can be set

`dynamic` - (Optional) Note that for a 3-node cluster the IPv4 prefix list should provide a minimum of 3 addresses to ensure each node gets an IP address from the pool.. See [Source Nat Choice Dynamic ](#source-nat-choice-dynamic) below for details.

`virtual_cidr` - (Optional) The range of the real CIDR and virtual CIDRs should be the same (e.g. if the real CIDR has the CIDR 10.10.10.0/24, the virtual CIDR has 100.100.100.0/24. (`String`).

### Rules Criteria

Criteria to match on the packet to apply the Action.

`destination_cidr` - (Optional) Destination IP of the packet to match (`String`).

`destination_port` - (Optional) Destination port of the packet to match. See [Criteria Destination Port ](#criteria-destination-port) below for details.

###### One of the arguments from this list "segment, site_local_inside_network, site_local_network, virtual_network" can be set

`segment` - (Optional) When there is no segment connector, this field need not be specified. When there is segment connector configured and if packet is destined to destination segment, destination segment needs to be configured here. See [Network Choice Segment ](#network-choice-segment) below for details.

`site_local_inside_network` - (Optional) x-displayName: "Site Local Inside Virtual Network" (`Bool`).(Deprecated)

`site_local_network` - (Optional) x-displayName: "Site Local Outside Virtual Network" (`Bool`).(Deprecated)

`virtual_network` - (Optional) When there is no network connector, this field need not be specified. When there is network connector configured and if packet is destined to destination network, destination network needs to be configured here. See [Network Choice Virtual Network ](#network-choice-virtual-network) below for details.

`protocol` - (Optional) Protocol of the packet to match (`String`).

###### One of the arguments from this list "any, icmp, tcp, udp" can be set

`any` - (Optional) Match Any Protocol (`Bool`).(Deprecated)

`icmp` - (Optional) Match ICMP Protocol (`Bool`).(Deprecated)

`tcp` - (Optional) Match TCP Protocol. See [Protocol Choice Tcp ](#protocol-choice-tcp) below for details.(Deprecated)

`udp` - (Optional) Match UDP Protocol. See [Protocol Choice Udp ](#protocol-choice-udp) below for details.(Deprecated)

`source_cidr` - (Optional) Source IP of the packet to match (`String`).

`source_port` - (Optional) Source port of the packet to match. See [Criteria Source Port ](#criteria-source-port) below for details.

### Scope Choice Cloud Connect

NAT rule is applied to packet coming from cloud connect.

`refs` - (Required) Reference to Cloud Connect Object. See [ref](#ref) below for details.

### Scope Choice Network Interface

NAT rule is applied to packet coming from interface.

`refs` - (Required) Reference to Network Interface Object. See [ref](#ref) below for details.

### Scope Choice Node Interface

NAT rule is applied to packet coming from one or more interfaces of nodes.

`list` - (Optional) On a multinode site, this list holds the nodes and corresponding networking_interface. See [Node Interface List ](#node-interface-list) below for details.

### Scope Choice Segment

NAT rule is applied to packet in the segment.

`refs` - (Required) Reference to Segment Object. See [ref](#ref) below for details.

`virtual_networks` - (Optional) Internally used to resolve segment ref to networks.. See [ref](#ref) below for details.(Deprecated)

### Scope Choice Virtual Network

NAT rule is applied to packet in the virtual network.

`refs` - (Optional) Reference to virtual network. See [ref](#ref) below for details.

### Source Nat Choice Dynamic

Note that for a 3-node cluster the IPv4 prefix list should provide a minimum of 3 addresses to ensure each node gets an IP address from the pool..

###### One of the arguments from this list "elastic_ips, pools" must be set

`elastic_ips` - (Optional) Use Cloud Elastic IP Object as SNAT Pool that is routable in Internet in Cloud Environments. See [Pool Choice Elastic Ips ](#pool-choice-elastic-ips) below for details.

`pools` - (Optional) Use the configured Pool as SNAT Pool. See [Pool Choice Pools ](#pool-choice-pools) below for details.

### Tcp Destination Port

Destination port of the packet to match.

###### One of the arguments from this list "no_port_match, port, port_ranges" can be set

`no_port_match` - (Optional) Disable matching of ports (`Bool`).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Tcp Source Port

Source port of the packet to match.

###### One of the arguments from this list "no_port_match, port, port_ranges" can be set

`no_port_match` - (Optional) Disable matching of ports (`Bool`).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Udp Destination Port

Destination port of the packet to match.

###### One of the arguments from this list "no_port_match, port, port_ranges" can be set

`no_port_match` - (Optional) Disable matching of ports (`Bool`).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Udp Source Port

Source port of the packet to match.

###### One of the arguments from this list "no_port_match, port, port_ranges" can be set

`no_port_match` - (Optional) Disable matching of ports (`Bool`).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured nat_policy.
