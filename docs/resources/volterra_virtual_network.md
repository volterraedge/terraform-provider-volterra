---

page_title: "Volterra: virtual_network"

description: "The virtual_network allows CRUD of Virtual Network resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

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

  // One of the arguments from this list "global_network legacy_type site_local_inside_network site_local_network" must be set

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

###### One of the arguments from this list "global_network, legacy_type, site_local_inside_network, site_local_network" must be set

`global_network` - (Optional) Global network can extend to multiple sites. (`Bool`).

`legacy_type` - (Optional) Type of virtual network (`String`).

`site_local_inside_network` - (Optional) Site local Inside network, also known as inside network (`Bool`).

`site_local_network` - (Optional) Site local network, also known as outside network (`Bool`).

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

### Next Hop Choice Default Gateway

Traffic matching the ip prefixes is sent to the default gateway.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured virtual_network.
