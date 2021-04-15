---

page_title: "Volterra: virtual_network"

description: "The virtual_network allows CRUD of Virtual Network resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_virtual_network
=================================

The Virtual Network allows CRUD of Virtual Network resource on Volterra SaaS

~> **Note:** Please refer to [Virtual Network API docs](https://volterra.io/docs/api/virtual-network) to learn more

Example Usage
-------------

```hcl
resource "volterra_virtual_network" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "srv6_network global_network site_local_network site_local_inside_network legacy_type" must be set
  global_network = true
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

`global_network` - (Optional) Global network can extend to multiple sites. (bool).

`legacy_type` - (Optional) Type of virtual network (`String`).

`site_local_inside_network` - (Optional) Site local Inside network, also known as inside network (bool).

`site_local_network` - (Optional) Site local network, also known as outside network (bool).

`srv6_network` - (Optional) Configure a per site srv6 network. See [Srv6 Network ](#srv6-network) below for details.

`static_routes` - (Optional) List of static routes on the virtual network. See [Static Routes ](#static-routes) below for details.

### Access Network Rtargets

Import Route Targets for connectivity to Access Networks..

`asn2byte_rtarget` - (Optional) Two-Octet AS Specific Route Target.. See [Asn2byte Rtarget ](#asn2byte-rtarget) below for details.

`asn4byte_rtarget` - (Optional) Four-Octet AS Specific Route Target.. See [Asn4byte Rtarget ](#asn4byte-rtarget) below for details.

`ipv4_addr_rtarget` - (Optional) IPv4 Address Specific Route Target.. See [Ipv4 Addr Rtarget ](#ipv4-addr-rtarget) below for details.

### Asn2byte Rtarget

Two-Octet AS Specific Route Target..

`as_number` - (Required) Two-Octet AS Number. (`Int`).

`value` - (Required) A 4-byte value that is unique in the scope of the ASN. (`Int`).

### Asn4byte Rtarget

Four-Octet AS Specific Route Target..

`as_number` - (Required) Four-Octet AS Number. (`Int`).

`value` - (Required) A 2-byte value that is unique in the scope of the ASN. (`Int`).

### Default Gateway

Traffic matching the ip prefixes is sent to default gateway .

### Enterprise Network Rtargets

Import Route Targets for connectivity to Enterprise Networks..

`asn2byte_rtarget` - (Optional) Two-Octet AS Specific Route Target.. See [Asn2byte Rtarget ](#asn2byte-rtarget) below for details.

`asn4byte_rtarget` - (Optional) Four-Octet AS Specific Route Target.. See [Asn4byte Rtarget ](#asn4byte-rtarget) below for details.

`ipv4_addr_rtarget` - (Optional) IPv4 Address Specific Route Target.. See [Ipv4 Addr Rtarget ](#ipv4-addr-rtarget) below for details.

### Export Rtargets

Export Route Targets for advertised routes..

`asn2byte_rtarget` - (Optional) Two-Octet AS Specific Route Target.. See [Asn2byte Rtarget ](#asn2byte-rtarget) below for details.

`asn4byte_rtarget` - (Optional) Four-Octet AS Specific Route Target.. See [Asn4byte Rtarget ](#asn4byte-rtarget) below for details.

`ipv4_addr_rtarget` - (Optional) IPv4 Address Specific Route Target.. See [Ipv4 Addr Rtarget ](#ipv4-addr-rtarget) below for details.

### Fleet Snat Pool

Configure address allocator for SNAT pool for a Fleet.

`snat_pool_allocator` - (Required) SNAT pool address allocator reference. See [ref](#ref) below for details.

### Fleet Vip

Configure per site anycast vip allocator.

`vip_allocator` - (Required) Anycast VIP address allocator reference. See [ref](#ref) below for details.

### Interface Ip Snat Pool

SNAT pool is interface ip of respective node.

### Interface Ip Vip

Default VIP is interface ip of respective node.

### Internet Rtargets

Import Route Targets for connectivity to the Internet..

`asn2byte_rtarget` - (Optional) Two-Octet AS Specific Route Target.. See [Asn2byte Rtarget ](#asn2byte-rtarget) below for details.

`asn4byte_rtarget` - (Optional) Four-Octet AS Specific Route Target.. See [Asn4byte Rtarget ](#asn4byte-rtarget) below for details.

`ipv4_addr_rtarget` - (Optional) IPv4 Address Specific Route Target.. See [Ipv4 Addr Rtarget ](#ipv4-addr-rtarget) below for details.

### Ipv4 Addr Rtarget

IPv4 Address Specific Route Target..

`address` - (Required) IPv4 Address (`String`).

`value` - (Required) A 4-byte value that is unique in the scope of the IPv4 address. (`Int`).

### No Namespace Network

Namespace network is not connected to this network.

### Node Snat Pool

Per node ip v4 prefix for snat pool.

`ipv4_prefixes` - (Required) List of IPv4 prefixes used as SNAT pool (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Site Snat Pool

Configure per node SNAT pool for a site.

`node_snat_pool` - (Required) Per node ip v4 prefix for snat pool. See [Node Snat Pool ](#node-snat-pool) below for details.

### Srv6 Network

Configure a per site srv6 network.

`access_network_rtargets` - (Optional) Import Route Targets for connectivity to Access Networks.. See [Access Network Rtargets ](#access-network-rtargets) below for details.

`anycast_vip` - (Optional) Configure anycast VIP (`String`).

`fleet_vip` - (Optional) Configure per site anycast vip allocator. See [Fleet Vip ](#fleet-vip) below for details.

`interface_ip_vip` - (Optional) Default VIP is interface ip of respective node (bool).

`enterprise_network_rtargets` - (Optional) Import Route Targets for connectivity to Enterprise Networks.. See [Enterprise Network Rtargets ](#enterprise-network-rtargets) below for details.

`export_rtargets` - (Optional) Export Route Targets for advertised routes.. See [Export Rtargets ](#export-rtargets) below for details.

`fleets` - (Optional) The set of sites where this virtual network is to be instantiated.. See [ref](#ref) below for details.

`internet_rtargets` - (Optional) Import Route Targets for connectivity to the Internet.. See [Internet Rtargets ](#internet-rtargets) below for details.

`no_namespace_network` - (Optional) Namespace network is not connected to this network (bool).

`srv6_network_ns_params` - (Optional) Name of namespace whose network is connected. See [Srv6 Network Ns Params ](#srv6-network-ns-params) below for details.

`slice` - (Required) The srv6_network_slice to which this network belongs.. See [ref](#ref) below for details.

`fleet_snat_pool` - (Optional) Configure address allocator for SNAT pool for a Fleet. See [Fleet Snat Pool ](#fleet-snat-pool) below for details.

`interface_ip_snat_pool` - (Optional) SNAT pool is interface ip of respective node (bool).

`site_snat_pool` - (Optional) Configure per node SNAT pool for a site. See [Site Snat Pool ](#site-snat-pool) below for details.

### Srv6 Network Ns Params

Name of namespace whose network is connected.

`namespace` - (Required) Name of namespace that is connected to srv6 Network (`String`).

### Static Routes

List of static routes on the virtual network.

`attrs` - (Optional) List of attributes that control forwarding, dynamic routing and control plane(host) reachability (`List of Strings`).

`ip_prefixes` - (Required) List of route prefixes that have common next hop and attributes (`String`).

`default_gateway` - (Optional) Traffic matching the ip prefixes is sent to default gateway (bool).

`interface` - (Optional) Traffic matching the ip prefixes is sent to the interface. See [ref](#ref) below for details.

`ip_address` - (Optional) Traffic matching the ip prefixes is sent to IP Address (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured virtual_network.
