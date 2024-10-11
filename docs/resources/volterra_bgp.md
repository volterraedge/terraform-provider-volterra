---

page_title: "Volterra: bgp"
description: "The bgp allows CRUD of Bgp resource on Volterra SaaS"

---

Resource volterra_bgp
=====================

The Bgp allows CRUD of Bgp resource on Volterra SaaS

~> **Note:** Please refer to [Bgp API docs](https://docs.cloud.f5.com/docs-v2/api/bgp) to learn more

Example Usage
-------------

```hcl
resource "volterra_bgp" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  bgp_parameters {
    asn = "64512"

    bgp_router_id {
      // One of the arguments from this list "ipv4 ipv6" can be set

      ipv4 {
        addr = "192.168.1.1"
      }
    }

    bgp_router_id_key = "value"

    bgp_router_id_type = "bgp_router_id_type"

    // One of the arguments from this list "from_site ip_address local_address" must be set

    local_address = true
  }

  peers {
    metadata {
      description = "Virtual Host for acmecorp website"

      disable = true

      name = "acmecorp-web"
    }

    // One of the arguments from this list "passive_mode_disabled passive_mode_enabled" must be set

    passive_mode_disabled = true
    target_service = "value"

    // One of the arguments from this list "external internal" must be set

    external {
      // One of the arguments from this list "address default_gateway disable from_site subnet_begin_offset subnet_end_offset" must be set

      address = "address"

      // One of the arguments from this list "address_ipv6 default_gateway_v6 disable_v6 from_site_v6 subnet_begin_offset_v6 subnet_end_offset_v6" must be set

      default_gateway_v6 = true
      asn = "64512"

      // One of the arguments from this list "md5_auth_key no_authentication" can be set

      no_authentication = true
      family_inet {
        // One of the arguments from this list "disable enable" must be set

        enable = true
      }
      family_inet_v6 {
        // One of the arguments from this list "disable enable" must be set

        disable = true
      }

      // One of the arguments from this list "inside_interfaces interface interface_list outside_interfaces" must be set

      interface {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }
      port = "179"
    }
  }

  where {
    // One of the arguments from this list "site virtual_site" must be set

    site {
      // One of the arguments from this list "disable_internet_vip enable_internet_vip" must be set

      disable_internet_vip = true

      network_type = "network_type"

      ref {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }

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

`bgp_parameters` - (Required) BGP parameters for local site. See [Bgp Parameters ](#bgp-parameters) below for details.

`peers` - (Required) List of peers. See [Peers ](#peers) below for details.

`where` - (Required) Site or virtual site where this BGP configuration should be applied.. See [Where ](#where) below for details.

### Bgp Parameters

BGP parameters for local site.

`asn` - (Required) Autonomous System Number (`Int`).

`bgp_router_id` - (Optional) If Router ID Type is set to "From IP Address", this is used as Router ID. Else, this is ignored.. See [Bgp Parameters Bgp Router Id ](#bgp-parameters-bgp-router-id) below for details.(Deprecated)

`bgp_router_id_key` - (Optional) from site template parameters map in site object. Else, this is ignored. (`String`).(Deprecated)

`bgp_router_id_type` - (Optional) Decides how BGP router id is derived (`String`).(Deprecated)

###### One of the arguments from this list "from_site, ip_address, local_address" must be set

`from_site` - (Optional) Use the Router ID field from the site object. (`Bool`).

`ip_address` - (Optional) Use the configured IPv4 Address as Router ID. (`String`).

`local_address` - (Optional) Use an interface address of the site as the Router ID. (`Bool`).

### Peers

List of peers.

`metadata` - (Required) Common attributes for the peer including name and description.. See [Peers Metadata ](#peers-metadata) below for details.

###### One of the arguments from this list "passive_mode_disabled, passive_mode_enabled" must be set

`passive_mode_disabled` - (Optional) x-displayName: "Disabled" (`Bool`).

`passive_mode_enabled` - (Optional) x-displayName: "Enabled" (`Bool`).

`target_service` - (Optional) Specify whether this peer should be configured in "phobos" or "frr". (`String`).(Deprecated)

###### One of the arguments from this list "external, internal" must be set

`external` - (Optional) External BGP peer.. See [Type Choice External ](#type-choice-external) below for details.

`internal` - (Optional) Internal BGP peer.. See [Type Choice Internal ](#type-choice-internal) below for details.(Deprecated)

### Where

Site or virtual site where this BGP configuration should be applied..

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Direct reference to site object. See [Ref Or Selector Site ](#ref-or-selector-site) below for details.

`virtual_site` - (Optional) Direct reference to virtual site object. See [Ref Or Selector Virtual Site ](#ref-or-selector-virtual-site) below for details.

### Address Choice Default Gateway

Use the default gateway address..

### Address Choice Disable

No Peer Ipv4 Address..

### Address Choice From Site

Use the address specified in the site object..

### Address Choice V6 Default Gateway V6

Use the default gateway address..

### Address Choice V6 Disable V6

No Peer IPv6 Address..

### Address Choice V6 From Site V6

Use the address specified in the site object..

### Auth Choice No Authentication

No Authentication of BGP session.

### Bgp Parameters Bgp Router Id

If Router ID Type is set to "From IP Address", this is used as Router ID. Else, this is ignored..

###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Enable Choice Disable

Disable IPv4 family Route Exchange..

### Enable Choice Enable

Enable IPv4 family Route Exchange..

### Enable Choice Enable

Enable the IPv4 Unicast family..

###### One of the arguments from this list "disable, enable" must be set

`disable` - (Optional) Disable the IPv4 Unicast family. (`Bool`).(Deprecated)

`enable` - (Optional) Enable the IPv4 Unicast family. (`Bool`).

### External Family Inet

Enable/Disable Ipv4 family of routes exchange with peer.

###### One of the arguments from this list "disable, enable" must be set

`disable` - (Optional) Disable IPv4 family Route Exchange. (`Bool`).

`enable` - (Optional) Enable IPv4 family Route Exchange. (`Bool`).

### External Family Inet V6

Enable/Disable IPv6 family of routes exchange with peer.

###### One of the arguments from this list "disable, enable" must be set

`disable` - (Optional) Disable IPv6 family Route Exchange. (`Bool`).

`enable` - (Optional) Enable IPv6 family Route Exchange. (`Bool`).

### Interface Choice Inside Interfaces

All interfaces in the site local inside network..

### Interface Choice Interface List

List of network interfaces..

`interfaces` - (Required) List of network interfaces.. See [ref](#ref) below for details.

### Interface Choice Outside Interfaces

All interfaces in the site local outside network..

### Internal Family Inet6vpn

Parameters for IPv6 VPN Unicast family..

###### One of the arguments from this list "disable, enable" must be set

`disable` - (Optional) Disable the IPv6 Unicast family. (`Bool`).(Deprecated)

`enable` - (Optional) Enable the IPv6 Unicast family. (`Bool`).

### Internal Family Inetvpn

Parameters for IPv4 VPN Unicast family..

###### One of the arguments from this list "disable, enable" must be set

`disable` - (Optional) Disable the IPv4 Unicast family. (`Bool`).(Deprecated)

`enable` - (Optional) Enable the IPv4 Unicast family.. See [Enable Choice Enable ](#enable-choice-enable) below for details.

### Internal Family Rtarget

Parameters for Route Target family..

###### One of the arguments from this list "disable, enable" must be set

`disable` - (Optional) Disable the Route Target family. (`Bool`).

`enable` - (Optional) Enable the Route Target family. (`Bool`).

### Internal Family Uuidvpn

Parameters for UUID VPN Unicast family..

###### One of the arguments from this list "disable, enable" must be set

`disable` - (Optional) Disable the UUID Unicast family. (`Bool`).(Deprecated)

`enable` - (Optional) Enable the UUID Unicast family. (`Bool`).

### Internet Vip Choice Disable Internet Vip

Do not enable advertise on external internet vip..

### Internet Vip Choice Enable Internet Vip

Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site..

### Mtls Choice Disable Mtls

Disable mTLS.

### Mtls Choice Enable Mtls

Enable mTLS.

### Passive Choice Passive Mode Disabled

x-displayName: "Disabled".

### Passive Choice Passive Mode Enabled

x-displayName: "Enabled".

### Peers Metadata

Common attributes for the peer including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Ref Or Selector Site

Direct reference to site object.

###### One of the arguments from this list "disable_internet_vip, enable_internet_vip" must be set

`disable_internet_vip` - (Optional) Do not enable advertise on external internet vip. (`Bool`).

`enable_internet_vip` - (Optional) Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site. (`Bool`).

`network_type` - (Optional) The type of network on the referred site (`String`).

`ref` - (Required) A site direct reference. See [ref](#ref) below for details.

`refs` - (Optional) Reference to virtual network. See [ref](#ref) below for details.(Deprecated)

### Ref Or Selector Virtual Site

Direct reference to virtual site object.

###### One of the arguments from this list "disable_internet_vip, enable_internet_vip" must be set

`disable_internet_vip` - (Optional) Do not enable advertise on external internet vip. (`Bool`).

`enable_internet_vip` - (Optional) Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site. (`Bool`).

`network_type` - (Optional) The type of network on the referred virtual_site (`String`).

`ref` - (Required) A virtual_site direct reference. See [ref](#ref) below for details.

`refs` - (Optional) Reference to virtual network. See [ref](#ref) below for details.(Deprecated)

### Router Id Choice From Site

Use the Router ID field from the site object..

### Router Id Choice Local Address

Use an interface address of the site as the Router ID..

### Sr Choice Disable

Disable the IPv4 Unicast family..

### Sr Choice Enable

Enable the IPv4 Unicast family..

### Type Choice External

External BGP peer..

###### One of the arguments from this list "address, default_gateway, disable, from_site, subnet_begin_offset, subnet_end_offset" must be set

`address` - (Optional) Specify IPV4 peer address. (`String`).

`default_gateway` - (Optional) Use the default gateway address. (`Bool`).

`disable` - (Optional) No Peer Ipv4 Address. (`Bool`).

`from_site` - (Optional) Use the address specified in the site object. (`Bool`).

`subnet_begin_offset` - (Optional) Calculate peer address using offset from the beginning of the subnet. (`Int`).

`subnet_end_offset` - (Optional) Calculate peer address using offset from the end of the subnet. (`Int`).

###### One of the arguments from this list "address_ipv6, default_gateway_v6, disable_v6, from_site_v6, subnet_begin_offset_v6, subnet_end_offset_v6" must be set

`address_ipv6` - (Optional) Specify peer IPv6 address. (`String`).

`default_gateway_v6` - (Optional) Use the default gateway address. (`Bool`).

`disable_v6` - (Optional) No Peer IPv6 Address. (`Bool`).

`from_site_v6` - (Optional) Use the address specified in the site object. (`Bool`).

`subnet_begin_offset_v6` - (Optional) Calculate peer address using offset from the beginning of the subnet. (`Int`).

`subnet_end_offset_v6` - (Optional) Calculate peer address using offset from the end of the subnet. (`Int`).

`asn` - (Required) Autonomous System Number for BGP peer (`Int`).

###### One of the arguments from this list "md5_auth_key, no_authentication" can be set

`md5_auth_key` - (Optional) MD5 key for protecting BGP Sessions (RFC 2385) (`String`).

`no_authentication` - (Optional) No Authentication of BGP session (`Bool`).

`family_inet` - (Optional) Enable/Disable Ipv4 family of routes exchange with peer. See [External Family Inet ](#external-family-inet) below for details.

`family_inet_v6` - (Optional) Enable/Disable IPv6 family of routes exchange with peer. See [External Family Inet V6 ](#external-family-inet-v6) below for details.

###### One of the arguments from this list "inside_interfaces, interface, interface_list, outside_interfaces" must be set

`inside_interfaces` - (Optional) All interfaces in the site local inside network. (`Bool`).(Deprecated)

`interface` - (Optional) Specify interface.. See [ref](#ref) below for details.

`interface_list` - (Optional) List of network interfaces.. See [Interface Choice Interface List ](#interface-choice-interface-list) below for details.

`outside_interfaces` - (Optional) All interfaces in the site local outside network. (`Bool`).(Deprecated)

`port` - (Optional) Peer TCP port number. (`Int`).

### Type Choice Internal

Internal BGP peer..

###### One of the arguments from this list "address, dns_name, from_site" must be set

`address` - (Optional) Specify peer address. (`String`).

`dns_name` - (Optional) Use the addresse by resolving the given DNS name. (`String`).(Deprecated)

`from_site` - (Optional) Use the address specified in the site object. (`Bool`).

`family_inet6vpn` - (Optional) Parameters for IPv6 VPN Unicast family.. See [Internal Family Inet6vpn ](#internal-family-inet6vpn) below for details.

`family_inetvpn` - (Optional) Parameters for IPv4 VPN Unicast family.. See [Internal Family Inetvpn ](#internal-family-inetvpn) below for details.

`family_rtarget` - (Optional) Parameters for Route Target family.. See [Internal Family Rtarget ](#internal-family-rtarget) below for details.

`family_uuidvpn` - (Optional) Parameters for UUID VPN Unicast family.. See [Internal Family Uuidvpn ](#internal-family-uuidvpn) below for details.

###### One of the arguments from this list "disable_mtls, enable_mtls" can be set

`disable_mtls` - (Optional) Disable mTLS (`Bool`).(Deprecated)

`enable_mtls` - (Optional) Enable mTLS (`Bool`).(Deprecated)

`port` - (Optional) Local Peer TCP Port Number. (`Int`).

### Ver Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ver Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured bgp.
