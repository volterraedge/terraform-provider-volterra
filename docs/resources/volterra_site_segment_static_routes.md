---

page_title: "Volterra: site_segment_static_routes"

description: "The site_segment_static_routes allows CRUD of Site Segment Static Routes resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------------------

Resource volterra_site_segment_static_routes
============================================

The Site Segment Static Routes allows CRUD of Site Segment Static Routes resource on Volterra SaaS

~> **Note:** Please refer to [Site Segment Static Routes API docs](https://docs.cloud.f5.com/docs-v2/api/site-segment-static-routes) to learn more

Example Usage
-------------

```hcl
resource "volterra_site_segment_static_routes" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  site {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }

  virtual_network {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
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

`nameserver_v4` - (Optional) Optional IPv4 DNS server to be used for name resolution (`String`).

`nameserver_v6` - (Optional) Optional IPv6 DNS server to be used for name resolution (`String`).

`secondary_nameserver_v4` - (Optional) Optional Secondary IPv4 DNS server to be used for name resolution (`String`).

`secondary_nameserver_v6` - (Optional) Optional Secondary IPv6 DNS server to be used for name resolution (`String`).

`site` - (Required) Reference to the site. See [ref](#ref) below for details.

`static_routes` - (Optional) List of static routes on the Site. See [Static Routes ](#static-routes) below for details.

`static_v6_routes` - (Optional) List of IPV6 static routes on the Site. See [Static V6 Routes ](#static-v6-routes) below for details.

`virtual_network` - (Required) Reference to the virtual network. See [ref](#ref) below for details.

### Static Routes

List of static routes on the Site.

`attrs` - (Optional) List of route attributes associated with the static route (`List of Strings`).

`labels` - (Optional) Add Labels for this Static Route, these labels can be used in network policy (`String`).

`nexthop` - (Optional) Nexthop for the route. See [Static Routes Nexthop ](#static-routes-nexthop) below for details.

`subnets` - (Required) List of route prefixes. See [Static Routes Subnets ](#static-routes-subnets) below for details.

### Static V6 Routes

List of IPV6 static routes on the Site.

`attrs` - (Optional) List of route attributes associated with the static route (`List of Strings`).

`labels` - (Optional) Add Labels for this Static Route, these labels can be used in network policy (`String`).

`nexthop` - (Optional) Nexthop for the route. See [Static V6 Routes Nexthop ](#static-v6-routes-nexthop) below for details.

`subnets` - (Required) List of route prefixes. See [Static V6 Routes Subnets ](#static-v6-routes-subnets) below for details.

### Nexthop Nexthop Address

Nexthop address when type is "Use-Configured".

###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Static Routes Nexthop

Nexthop for the route.

`interface` - (Optional) Nexthop is network interface when type is "Network-Interface". See [ref](#ref) below for details.

`nexthop_address` - (Optional) Nexthop address when type is "Use-Configured". See [Nexthop Nexthop Address ](#nexthop-nexthop-address) below for details.

`type` - (Optional) Identifies the type of next-hop (`String`).

### Static Routes Subnets

List of route prefixes.

###### One of the arguments from this list "ipv4, ipv6" must be set

`ipv4` - (Optional) IPv4 Subnet Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Subnet Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Static V6 Routes Nexthop

Nexthop for the route.

`interface` - (Optional) Nexthop is network interface when type is "Network-Interface". See [ref](#ref) below for details.

`nexthop_address` - (Optional) Nexthop address when type is "Use-Configured". See [Nexthop Nexthop Address ](#nexthop-nexthop-address) below for details.

`type` - (Optional) Identifies the type of next-hop (`String`).

### Static V6 Routes Subnets

List of route prefixes.

###### One of the arguments from this list "ipv4, ipv6" must be set

`ipv4` - (Optional) IPv4 Subnet Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Subnet Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Ver Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ver Ipv4

IPv4 Subnet Address.

`plen` - (Optional) Prefix-length of the IPv4 subnet. Must be <= 32 (`Int`).

`prefix` - (Optional) Prefix part of the IPv4 subnet in string form with dot-decimal notation (`String`).

### Ver Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

### Ver Ipv6

IPv6 Subnet Address.

`plen` - (Optional) Prefix length of the IPv6 subnet. Must be <= 128 (`Int`).

`prefix` - (Optional) e.g. "2001:db8::2::" (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured site_segment_static_routes.
