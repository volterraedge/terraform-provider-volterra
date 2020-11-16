---

page_title: "Volterra: network_interface"

description: "The network_interface allows CRUD of Network Interface resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_network_interface
===================================

The Network Interface allows CRUD of Network Interface resource on Volterra SaaS

~> **Note:** Please refer to [Network Interface API docs](https://volterra.io/docs/api/network-interface) to learn more

Example Usage
-------------

```hcl
resource "volterra_network_interface" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "dedicated_management_interface dedicated_interface ethernet_interface tunnel_interface legacy_interface" must be set

  tunnel_interface {
    mtu = "1450"

    // One of the arguments from this list "site_local_network site_local_inside_network inside_network" must be set
    site_local_network = true

    // One of the arguments from this list "cluster node" must be set
    cluster  = true
    priority = "42"

    static_ip {
      // One of the arguments from this list "fleet_static_ip node_static_ip cluster_static_ip" must be set

      node_static_ip {
        default_gw = "192.168.20.1"
        dns_server = "192.168.20.1"
        ip_address = "192.168.20.1/24"
      }
    }

    tunnel {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
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

`dedicated_interface` - (Optional) Networking configuration for dedicated interface is configured locally on site e.g. (outside/inside)Ethernet, WLAN, or LTE/4G. . See [Dedicated Interface ](#dedicated-interface) below for details.

`dedicated_management_interface` - (Optional) In dc cluster sites fallback management interfaces can be made into dedicated management interface. See [Dedicated Management Interface ](#dedicated-management-interface) below for details.

`ethernet_interface` - (Optional) Ethernet interface configuration.. See [Ethernet Interface ](#ethernet-interface) below for details.

`legacy_interface` - (Optional) Old method of interface configuration. See [Legacy Interface ](#legacy-interface) below for details.

`tunnel_interface` - (Optional) Tunnel interface, Ipsec tunnels to other networking devices.. See [Tunnel Interface ](#tunnel-interface) below for details.

### DNS Server

Configures how DNS server is derived for the subnet in static addresses.

`dns_mode` - (Required) Mode of obtaining DNS server (`String`).

`dns_server` - (Optional) Address of DNS server when mode is "use-configured". See [Dns Server ](#dns-server) below for details.

### Dedicated Interface

Networking configuration for dedicated interface is configured locally on site e.g. (outside/inside)Ethernet, WLAN, or LTE/4G. .

`device` - (Required) Name of the device for which interface is configured. Use wwan0 for 4G/LTE. (`String`).

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site. (bool).

`node` - (Optional) Configuration will apply to a device on the given node of the site. (`String`).

`is_primary` - (Optional) This interface is primary (bool).

`not_primary` - (Optional) This interface is not primary (bool).

`priority` - (Optional) Greater the value, higher the priority (`Int`).

### Dedicated Management Interface

In dc cluster sites fallback management interfaces can be made into dedicated management interface.

`device` - (Required) Name of the device for which interface is configured (`String`).

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site. (bool).

`node` - (Optional) Configuration will apply to a device on the given node of the site. (`String`).

### Default Gateway

Configures how default gateway is derived for the subnet static addresses.

`default_gateway_address` - (Optional) Address of default gateway when mode is "use-configured". See [Default Gateway Address ](#default-gateway-address) below for details.

`default_gateway_mode` - (Required) Mode of obtaining default gateway (`String`).

### Default Gateway Address

Address of default gateway when mode is "use-configured".

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Dns Server

Address of DNS server when mode is "use-configured".

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ethernet Interface

Ethernet interface configuration..

`dhcp_client` - (Optional) Interface gets it IP address from external DHCP server (bool).

`dhcp_server` - (Optional) DHCP Server is configured for this interface, Interface IP from DHCP server configuration.. See [Dhcp Server ](#dhcp-server) below for details.

`static_ip` - (Optional) Interface IP is configured statically. See [Static Ip ](#static-ip) below for details.

`device` - (Required) Interface configuration for the ethernet device (`String`).

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

`inside_network` - (Optional) Interface belongs to user configured inside network. See [ref](#ref) below for details.

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (bool).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (bool).

`storage_network` - (Optional) Interface belongs to site local network inside (bool).

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site. (bool).

`node` - (Optional) Configuration will apply to a device on the given node. (`String`).

`is_primary` - (Optional) This interface is primary (bool).

`not_primary` - (Optional) This interface is not primary (bool).

`priority` - (Optional) Greater the value, higher the priority (`Int`).

`untagged` - (Optional) Configure a untagged ethernet interface (bool).

`vlan_id` - (Optional) Configure a VLAN tagged ethernet interface (`Int`).

### Legacy Interface

Old method of interface configuration.

`DHCP_server` - (Required) Behave as DHCP server for subnet configured in static addresses. (`String`).

`DNS_server` - (Optional) Configures how DNS server is derived for the subnet in static addresses. See [DNS Server ](#DNS-server) below for details.

`address_allocator` - (Optional) allocate a subnet for the interface and an address from the subnet is set on the interface.. See [ref](#ref) below for details.

`default_gateway` - (Optional) Configures how default gateway is derived for the subnet static addresses. See [Default Gateway ](#default-gateway) below for details.

`device_name` - (Required) Name of the physical network interface device which this network interface represents. (`String`).

`dhcp_address` - (Required) Enable DHCP based address assignment for this interface. (`String`).

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

`priority` - (Optional) Greater the value, higher the priority (`Int`).

`static_addresses` - (Optional) If DHCP server is enabled, configures the subnet to be used for IP allocation.. See [Static Addresses ](#static-addresses) below for details.

`tunnel` - (Optional) When interface is created as TUNNEL type, then reference to tunnel is specified here. See [Tunnel ](#tunnel) below for details.

`type` - (Required) Specifies the type of interface (ethernet, vlan, lacp etc) (`String`).

`virtual_network` - (Optional) This is optional and can contain at most one entry. See [ref](#ref) below for details.

`vlan_tag` - (Optional) when vlan_tagging is enabled, value must be between 1 - 4094 (`Int`).

`vlan_tagging` - (Optional) Must be enabled for VLAN interfaces (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Static Addresses

If DHCP server is enabled, configures the subnet to be used for IP allocation..

`plen` - (Optional) Prefix-length of the IPv4 subnet. Must be <= 32 (`Int`).

`prefix` - (Optional) Prefix part of the IPv4 subnet in string form with dot-decimal notation (`String`).

### Static Ip

Interface IP is configured statically.

`cluster_static_ip` - (Optional) Static IP configuration for a specific node. See [Cluster Static Ip ](#cluster-static-ip) below for details.

`fleet_static_ip` - (Optional) Static IP configuration for the fleet. See [Fleet Static Ip ](#fleet-static-ip) below for details.

`node_static_ip` - (Optional) Static IP configuration for the Node. See [Node Static Ip ](#node-static-ip) below for details.

### Tunnel

When interface is created as TUNNEL type, then reference to tunnel is specified here.

`tunnel` - (Optional) Tunnel which is attached to this interface. See [ref](#ref) below for details.

### Tunnel Interface

Tunnel interface, Ipsec tunnels to other networking devices..

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

`inside_network` - (Optional) Interface belongs to user configured inside network. See [ref](#ref) below for details.

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (bool).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (bool).

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site (bool).

`node` - (Optional) Configuration will apply to a given device on the given node. (`String`).

`priority` - (Optional) Greater the value, higher the priority (`Int`).

`static_ip` - (Optional) Interface IP is configured statically. See [Static Ip ](#static-ip) below for details.

`tunnel` - (Optional) Tunnel Configuration for this Interface. See [ref](#ref) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_interface.
