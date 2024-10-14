---

page_title: "Volterra: virtual_network"
description: "The virtual_network allows CRUD of Virtual Network resource on Volterra SaaS"

---

Resource volterra_virtual_network
=================================

The Virtual Network allows CRUD of Virtual Network resource on Volterra SaaS

~> **Note:** Please refer to [Virtual Network API docs](https://docs.cloud.f5.com/docs-v2/api/virtual-network) to learn more

Example Usage
-------------

```hcl
resource "volterra_virtual_network" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "global_network legacy_type site_local_inside_network site_local_network srv6_network" must be set

  site_local_inside_network = true
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

###### One of the arguments from this list "global_network, legacy_type, site_local_inside_network, site_local_network, srv6_network" must be set

`global_network` - (Optional) Global network can extend to multiple sites. (`Bool`).

`legacy_type` - (Optional) Type of virtual network (`String`).

`site_local_inside_network` - (Optional) Site local Inside network, also known as inside network (`Bool`).

`site_local_network` - (Optional) Site local network, also known as outside network (`Bool`).

`srv6_network` - (Optional) Configure a per site srv6 network. See [Network Choice Srv6 Network ](#network-choice-srv6-network) below for details.(Deprecated)

`static_routes` - (Optional) List of static routes on the virtual network. See [Static Routes ](#static-routes) below for details.

`static_v6_routes` - (Optional) List of static IPv6 routes on the virtual network. See [Static V6 Routes ](#static-v6-routes) below for details.

### Static Routes

List of static routes on the virtual network.

`attrs` - (Optional) List of attributes that control forwarding, dynamic routing and control plane (host) reachability (`List of Strings`).

`ip_prefixes` - (Required) List of route prefixes that have common next hop and attributes (`String`).

###### One of the arguments from this list "default_gateway, interface, ip_address" must be set

`default_gateway` - (Optional) Traffic matching the ip prefixes is sent to the default gateway (`Bool`).

`interface` - (Optional) Traffic matching the ip prefixes is sent to this interface. See [ref](#ref) below for details.

`ip_address` - (Optional) Traffic matching the ip prefixes is sent to this IP Address (`String`).

### Static V6 Routes

List of static IPv6 routes on the virtual network.

`attrs` - (Optional) List of attributes that control forwarding, dynamic routing and control plane (host) reachability (`List of Strings`).

`ip_prefixes` - (Required) List of IPv6 route prefixes that have common next hop and attributes (`String`).

###### One of the arguments from this list "default_gateway, interface, ip_address" must be set

`default_gateway` - (Optional) Traffic matching the ip prefixes is sent to the default gateway (`Bool`).

`interface` - (Optional) Traffic matching the ip prefixes is sent to this interface. See [ref](#ref) below for details.

`ip_address` - (Optional) Traffic matching the ip prefixes is sent to this IP Address (`String`).

### Default Vip Choice Fleet Vip

Configure per site anycast vip allocator.

`vip_allocator` - (Required) Anycast VIP address allocator reference. See [ref](#ref) below for details.

### Default Vip Choice Interface Ip Vip

Default VIP is interface ip of respective node.

### Namespace Choice No Namespace Network

Namespace network is not connected to this network.

### Namespace Choice Srv6 Network Ns Params

Name of namespace whose network is connected.

`namespace` - (Required) Name of namespace that is connected to srv6 Network (`String`).

### Network Choice Srv6 Network

Configure a per site srv6 network.

`access_network_rtargets` - (Optional) Import Route Targets for connectivity to Access Networks.. See [Srv6 Network Access Network Rtargets ](#srv6-network-access-network-rtargets) below for details.

###### One of the arguments from this list "anycast_vip, fleet_vip, interface_ip_vip" must be set

`anycast_vip` - (Optional) Configure anycast VIP (`String`).

`fleet_vip` - (Optional) Configure per site anycast vip allocator. See [Default Vip Choice Fleet Vip ](#default-vip-choice-fleet-vip) below for details.(Deprecated)

`interface_ip_vip` - (Optional) Default VIP is interface ip of respective node (`Bool`).

`enterprise_network_rtargets` - (Optional) Import Route Targets for connectivity to Enterprise Networks.. See [Srv6 Network Enterprise Network Rtargets ](#srv6-network-enterprise-network-rtargets) below for details.

`export_rtargets` - (Optional) Export Route Targets for advertised routes.. See [Srv6 Network Export Rtargets ](#srv6-network-export-rtargets) below for details.

`fleets` - (Optional) The set of sites where this virtual network is to be instantiated.. See [ref](#ref) below for details.

`internet_rtargets` - (Optional) Import Route Targets for connectivity to the Internet.. See [Srv6 Network Internet Rtargets ](#srv6-network-internet-rtargets) below for details.

###### One of the arguments from this list "no_namespace_network, srv6_network_ns_params" must be set

`no_namespace_network` - (Optional) Namespace network is not connected to this network (`Bool`).

`srv6_network_ns_params` - (Optional) Name of namespace whose network is connected. See [Namespace Choice Srv6 Network Ns Params ](#namespace-choice-srv6-network-ns-params) below for details.

`remote_sid_stats_plen` - (Optional) Number of most significant bits of remote SIDs to use for maintaining per-SID counters. (`Int`).

`slice` - (Required) The srv6_network_slice to which this network belongs.. See [ref](#ref) below for details.

###### One of the arguments from this list "fleet_snat_pool, interface_ip_snat_pool, site_snat_pool" must be set

`fleet_snat_pool` - (Optional) Configure address allocator for SNAT pool for a Fleet. See [Snat Pool Choice Fleet Snat Pool ](#snat-pool-choice-fleet-snat-pool) below for details.(Deprecated)

`interface_ip_snat_pool` - (Optional) SNAT pool is interface ip of respective node (`Bool`).

`site_snat_pool` - (Optional) Configure per node SNAT pool for a site. See [Snat Pool Choice Site Snat Pool ](#snat-pool-choice-site-snat-pool) below for details.

### Next Hop Choice Default Gateway

Traffic matching the ip prefixes is sent to the default gateway.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rtarget Choice Asn2byte Rtarget

Two-Octet AS Specific Route Target..

`as_number` - (Required) Two-Octet AS Number. (`Int`).

`value` - (Required) A 4-byte value that is unique in the scope of the ASN. (`Int`).

### Rtarget Choice Asn4byte Rtarget

Four-Octet AS Specific Route Target..

`as_number` - (Required) Four-Octet AS Number. (`Int`).

`value` - (Required) A 2-byte value that is unique in the scope of the ASN. (`Int`).

### Rtarget Choice Ipv4 Addr Rtarget

IPv4 Address Specific Route Target..

`address` - (Required) IPv4 Address (`String`).

`value` - (Required) A 2-byte value that is unique in the scope of the IPv4 address. (`Int`).

### Site Snat Pool Node Snat Pool

Per node ip v4 prefix for snat pool.

`ipv4_prefixes` - (Optional) List of IPv4 prefixes used as SNAT pool (`String`).

`ipv6_prefixes` - (Optional) List of IPv6 prefixes used as SNAT pool (`String`).

### Snat Pool Choice Fleet Snat Pool

Configure address allocator for SNAT pool for a Fleet.

`snat_pool_allocator` - (Required) SNAT pool address allocator reference. See [ref](#ref) below for details.

### Snat Pool Choice Interface Ip Snat Pool

SNAT pool is interface ip of respective node.

### Snat Pool Choice Site Snat Pool

Configure per node SNAT pool for a site.

`node_snat_pool` - (Required) Per node ip v4 prefix for snat pool. See [Site Snat Pool Node Snat Pool ](#site-snat-pool-node-snat-pool) below for details.

### Srv6 Network Access Network Rtargets

Import Route Targets for connectivity to Access Networks..

###### One of the arguments from this list "asn2byte_rtarget, asn4byte_rtarget, ipv4_addr_rtarget" must be set

`asn2byte_rtarget` - (Optional) Two-Octet AS Specific Route Target.. See [Rtarget Choice Asn2byte Rtarget ](#rtarget-choice-asn2byte-rtarget) below for details.

`asn4byte_rtarget` - (Optional) Four-Octet AS Specific Route Target.. See [Rtarget Choice Asn4byte Rtarget ](#rtarget-choice-asn4byte-rtarget) below for details.

`ipv4_addr_rtarget` - (Optional) IPv4 Address Specific Route Target.. See [Rtarget Choice Ipv4 Addr Rtarget ](#rtarget-choice-ipv4-addr-rtarget) below for details.

### Srv6 Network Enterprise Network Rtargets

Import Route Targets for connectivity to Enterprise Networks..

###### One of the arguments from this list "asn2byte_rtarget, asn4byte_rtarget, ipv4_addr_rtarget" must be set

`asn2byte_rtarget` - (Optional) Two-Octet AS Specific Route Target.. See [Rtarget Choice Asn2byte Rtarget ](#rtarget-choice-asn2byte-rtarget) below for details.

`asn4byte_rtarget` - (Optional) Four-Octet AS Specific Route Target.. See [Rtarget Choice Asn4byte Rtarget ](#rtarget-choice-asn4byte-rtarget) below for details.

`ipv4_addr_rtarget` - (Optional) IPv4 Address Specific Route Target.. See [Rtarget Choice Ipv4 Addr Rtarget ](#rtarget-choice-ipv4-addr-rtarget) below for details.

### Srv6 Network Export Rtargets

Export Route Targets for advertised routes..

###### One of the arguments from this list "asn2byte_rtarget, asn4byte_rtarget, ipv4_addr_rtarget" must be set

`asn2byte_rtarget` - (Optional) Two-Octet AS Specific Route Target.. See [Rtarget Choice Asn2byte Rtarget ](#rtarget-choice-asn2byte-rtarget) below for details.

`asn4byte_rtarget` - (Optional) Four-Octet AS Specific Route Target.. See [Rtarget Choice Asn4byte Rtarget ](#rtarget-choice-asn4byte-rtarget) below for details.

`ipv4_addr_rtarget` - (Optional) IPv4 Address Specific Route Target.. See [Rtarget Choice Ipv4 Addr Rtarget ](#rtarget-choice-ipv4-addr-rtarget) below for details.

### Srv6 Network Internet Rtargets

Import Route Targets for connectivity to the Internet..

###### One of the arguments from this list "asn2byte_rtarget, asn4byte_rtarget, ipv4_addr_rtarget" must be set

`asn2byte_rtarget` - (Optional) Two-Octet AS Specific Route Target.. See [Rtarget Choice Asn2byte Rtarget ](#rtarget-choice-asn2byte-rtarget) below for details.

`asn4byte_rtarget` - (Optional) Four-Octet AS Specific Route Target.. See [Rtarget Choice Asn4byte Rtarget ](#rtarget-choice-asn4byte-rtarget) below for details.

`ipv4_addr_rtarget` - (Optional) IPv4 Address Specific Route Target.. See [Rtarget Choice Ipv4 Addr Rtarget ](#rtarget-choice-ipv4-addr-rtarget) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured virtual_network.
