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
    asn = "64512"

    bgp_router_id {
      // One of the arguments from this list "ipv4 ipv6" must be set

      ipv4 {
        addr = "192.168.1.1"
      }
    }

    bgp_router_id_key  = "value"
    bgp_router_id_type = "bgp_router_id_type"

    // One of the arguments from this list "local_address from_site ip_address" must be set
    ip_address = "ip_address"
  }

  peers {
    metadata {
      description = "Virtual Host for acmecorp website"
      disable     = true
      name        = "acmecorp-web"
    }

    target_service = "value"

    // One of the arguments from this list "external internal" must be set

    internal {
      // One of the arguments from this list "address from_site dns_name" must be set
      from_site = true

      family_inet6vpn {
        // One of the arguments from this list "disable enable" must be set
        enable = true
      }

      family_inetvpn {
        // One of the arguments from this list "enable disable" must be set

        enable {
          // One of the arguments from this list "enable disable" must be set
          enable = true
        }
      }

      family_rtarget {
        // One of the arguments from this list "enable disable" must be set
        enable = true
      }

      family_uuidvpn {
        // One of the arguments from this list "enable disable" must be set
        enable = true
      }

      // One of the arguments from this list "disable_mtls enable_mtls" must be set
      disable_mtls = true
      port         = "179"
    }
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

`peers` - (Required) List of peers. See [Peers ](#peers) below for details.

`where` - (Required) Site or virtual site where this BGP configuration should be applied.. See [Where ](#where) below for details.

### Bgp Parameters

BGP parameters for local site.

`asn` - (Required) Autonomous System Number (`Int`).

`bgp_router_id` - (Optional) If Router ID Type is set to "From IP Address", this is used as Router ID. Else, this is ignored.. See [Bgp Router Id ](#bgp-router-id) below for details.

`bgp_router_id_key` - (Optional) from site template parameters map in site object. Else, this is ignored. (`String`).

`bgp_router_id_type` - (Optional) Decides how BGP router id is derived (`String`).

`from_site` - (Optional) Use the Router ID field from the site object. (bool).

`ip_address` - (Optional) Use the configured IPv4 Address as Router ID. (`String`).

`local_address` - (Optional) Use an interface address of the site as the Router ID. (bool).

### Bgp Router Id

If Router ID Type is set to "From IP Address", this is used as Router ID. Else, this is ignored..

`ipv4` - (Optional) IPv4 Address. See [Ipv4 ](#ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ipv6 ](#ipv6) below for details.

### Default Gateway

Use the default gateway address..

### Disable

Disable the IPv4 Unicast family..

### Disable Mtls

Disable MTLS.

### Enable

Enable the IPv4 Unicast family..

### Enable Mtls

Enable MTLS.

### External

External BGP peer..

`address` - (Optional) Specify peer address. (`String`).

`default_gateway` - (Optional) Use the default gateway address. (bool).

`from_site` - (Optional) Use the address specified in the site object. (bool).

`subnet_begin_offset` - (Optional) Calculate peer address using offset from the beginning of the subnet. (`Int`).

`subnet_end_offset` - (Optional) Calculate peer address using offset from the end of the subnet. (`Int`).

`asn` - (Required) Autonomous System Number for BGP peer (`Int`).

`family_inet` - (Optional) Parameters for IPv4 Unicast family.. See [Family Inet ](#family-inet) below for details.

`inside_interfaces` - (Optional) All interfaces in the site local inside network. (bool).

`interface` - (Optional) Specify interface.. See [ref](#ref) below for details.

`interface_list` - (Optional) List of network interfaces.. See [Interface List ](#interface-list) below for details.

`outside_interfaces` - (Optional) All interfaces in the site local outside network. (bool).

`port` - (Optional) Peer TCP port number. (`Int`).

### Family Inet

Parameters for IPv4 Unicast family..

`disable` - (Optional) Disable the IPv4 Unicast family. (bool).

`enable` - (Optional) Enable the IPv4 Unicast family. (bool).

### Family Inet6vpn

Parameters for IPv6 VPN Unicast family..

`disable` - (Optional) Disable the IPv6 Unicast family. (bool).

`enable` - (Optional) Enable the IPv6 Unicast family. (bool).

### Family Inetvpn

Parameters for IPv4 VPN Unicast family..

`disable` - (Optional) Disable the IPv4 Unicast family. (bool).

`enable` - (Optional) Enable the IPv4 Unicast family.. See [Enable ](#enable) below for details.

### Family Rtarget

Parameters for Route Target family..

`disable` - (Optional) Disable the Route Target family. (bool).

`enable` - (Optional) Enable the Route Target family. (bool).

### Family Uuidvpn

Parameters for UUID VPN Unicast family..

`disable` - (Optional) Disable the UUID Unicast family. (bool).

`enable` - (Optional) Enable the UUID Unicast family. (bool).

### From Site

Use the Router ID field from the site object..

### Inside Interfaces

All interfaces in the site local inside network..

### Interface List

List of network interfaces..

`interfaces` - (Required) List of network interfaces.. See [ref](#ref) below for details.

### Internal

External BGP peer..

`address` - (Optional) Specify peer address. (`String`).

`dns_name` - (Optional) Use the addresse by resolving the given DNS name. (`String`).

`from_site` - (Optional) Use the address specified in the site object. (bool).

`family_inet6vpn` - (Optional) Parameters for IPv6 VPN Unicast family.. See [Family Inet6vpn ](#family-inet6vpn) below for details.

`family_inetvpn` - (Optional) Parameters for IPv4 VPN Unicast family.. See [Family Inetvpn ](#family-inetvpn) below for details.

`family_rtarget` - (Optional) Parameters for Route Target family.. See [Family Rtarget ](#family-rtarget) below for details.

`family_uuidvpn` - (Optional) Parameters for UUID VPN Unicast family.. See [Family Uuidvpn ](#family-uuidvpn) below for details.

`disable_mtls` - (Optional) Disable MTLS (bool).

`enable_mtls` - (Optional) Enable MTLS (bool).

`port` - (Optional) Local Peer TCP Port Number. (`Int`).

### Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

### Local Address

Use an interface address of the site as the Router ID..

### Metadata

Common attributes for the peer including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Outside Interfaces

All interfaces in the site local outside network..

### Peers

List of peers.

`metadata` - (Required) Common attributes for the peer including name and description.. See [Metadata ](#metadata) below for details.

`target_service` - (Optional) Specify whether this peer should be configured in "phobos" or "frr". (`String`).

`external` - (Optional) External BGP peer.. See [External ](#external) below for details.

`internal` - (Optional) External BGP peer.. See [Internal ](#internal) below for details.

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

Site or virtual site where this BGP configuration should be applied..

`site` - (Optional) Direct reference to site object. See [Site ](#site) below for details.

`virtual_site` - (Optional) Direct reference to virtual site object. See [Virtual Site ](#virtual-site) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured bgp.
