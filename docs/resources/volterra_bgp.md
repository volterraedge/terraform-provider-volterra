---

page_title: "Volterra: bgp"

description: "The bgp allows CRUD of Bgp resource on Volterra SaaS"
-------------------------------------------------------------------

Resource volterra_bgp
=====================

The Bgp allows CRUD of Bgp resource on Volterra SaaS

~> **Note:** Please refer to [Bgp API docs](https://volterra.io/docs/api/bgp) to learn more

Example Usage
-------------

```hcl
resource "volterra_bgp" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  bgp_parameters {
    asn = "asn"

    bgp_router_id {
      // One of the arguments from this list "ipv4 ipv6" must be set

      ipv4 {
        addr = "192.168.1.1"
      }
    }

    bgp_router_id_key  = "bgp_router_id_key"
    bgp_router_id_type = "bgp_router_id_type"
  }

  bgp_peers {
    asn = "asn"

    bgp_peer_address {
      // One of the arguments from this list "ipv4 ipv6" must be set

      ipv4 {
        addr = "192.168.1.1"
      }
    }

    bgp_peer_address_key   = "bgp_peer_address_key"
    bgp_peer_address_type  = "bgp_peer_address_type"
    bgp_peer_subnet_offset = "bgp_peer_subnet_offset"
    port                   = "port"
  }

  network_interface {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }

  where {
    // One of the arguments from this list "site virtual_site" must be set

    site {
      network_type = "network_type"

      ref {
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

`bgp_parameters` - (Required) BGP parameters for local site. See [Bgp Parameters ](#bgp-parameters) below for details.

`bgp_peers` - (Required) BGP parameters for peer. See [Bgp Peers ](#bgp-peers) below for details.

`network_interface` - (Required) List of interfaces to which the BGP configuration is applied. See [ref](#ref) below for details.

`where` - (Required) Set of sites in which this configuration is valid and must be present. See [Where ](#where) below for details.

### Bgp Parameters

BGP parameters for local site.

`asn` - (Required) Autonomous System Number for current site (`Int`).

`bgp_router_id` - (Optional) If Router ID Type is set to "From IP Address", this is used as Router ID. Else, this is ignored.. See [Bgp Router Id ](#bgp-router-id) below for details.

`bgp_router_id_key` - (Optional) from site template parameters map in site object. Else, this is ignored. (`String`).

`bgp_router_id_type` - (Required) Decides how BGP router id is derived (`String`).

### Bgp Peer Address

If Peer Address Type is set to "From IP Address", this is used as Peer Address. Else, this is ignored..

`ipv4` - (Optional) IPv4 Address. See [Ipv4 ](#ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ipv6 ](#ipv6) below for details.

### Bgp Peers

BGP parameters for peer.

`asn` - (Required) Autonomous System Number for BGP peer (`Int`).

`bgp_peer_address` - (Optional) If Peer Address Type is set to "From IP Address", this is used as Peer Address. Else, this is ignored.. See [Bgp Peer Address ](#bgp-peer-address) below for details.

`bgp_peer_address_key` - (Optional) from site template parameters map in site object. Else, this is ignored. (`String`).

`bgp_peer_address_type` - (Required) Decides how bgp peer address is derived for this peer (`String`).

`bgp_peer_subnet_offset` - (Optional) set to 5 with Peer Address Type set to "Offset from end of Subnet", peer address of 1.2.3.250 is used. (`Int`).

`port` - (Optional) Peer's port number, defaults to port 179 when not set (`Int`).

### Bgp Router Id

If Router ID Type is set to "From IP Address", this is used as Router ID. Else, this is ignored..

`ipv4` - (Optional) IPv4 Address. See [Ipv4 ](#ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ipv6 ](#ipv6) below for details.

### Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Site

Direct reference to site object.

`network_type` - (Optional) The type of network on the referred site (`String`).

`ref` - (Optional) A site direct reference. See [ref](#ref) below for details.

### Virtual Site

Direct reference to virtual site object.

`network_type` - (Optional) The type of network on the referred virtual_site (`String`).

`ref` - (Optional) A virtual_site direct reference. See [ref](#ref) below for details.

### Where

Set of sites in which this configuration is valid and must be present.

`site` - (Optional) Direct reference to site object. See [Site ](#site) below for details.

`virtual_site` - (Optional) Direct reference to virtual site object. See [Virtual Site ](#virtual-site) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured bgp.
