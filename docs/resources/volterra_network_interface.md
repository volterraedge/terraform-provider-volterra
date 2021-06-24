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

  // One of the arguments from this list "dedicated_interface ethernet_interface tunnel_interface legacy_interface dedicated_management_interface" must be set

  dedicated_interface {
    device = "eth0"

    // One of the arguments from this list "monitor_disabled monitor" must be set
    monitor_disabled = true
    mtu              = "1450"

    // One of the arguments from this list "node cluster" must be set
    cluster = true

    // One of the arguments from this list "not_primary is_primary" must be set
    not_primary = true
    priority    = "42"
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

### Automatic From End

Assign automatically from End of the first network in the list.

### Automatic From Start

Assign automatically from start of the first network in the list.

### Cluster

Configuration will apply to given device on all nodes of the site..

### Cluster Static Ip

Static IP configuration for a specific node.

`interface_ip_map` - (Optional) Map of Node to Static ip configuration value, Key:Node, Value:IP Address. See [Interface Ip Map ](#interface-ip-map) below for details.

### Dedicated Interface

Networking configuration for dedicated interface is configured locally on site e.g. (outside/inside)Ethernet, WLAN, or LTE/4G. .

`device` - (Required) Name of the device for which interface is configured. Use wwan0 for 4G/LTE. (`String`).

`monitor` - (Optional) Link Quality Monitoring parameters. Choosing the option will enable link quality monitoring.. See [Monitor ](#monitor) below for details.

`monitor_disabled` - (Optional) Link quality monitoring disabled on the interface. (bool).

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

### Dhcp Client

Interface gets it IP address from external DHCP server.

### Dhcp Networks

List of networks from which DHCP server can allocate ip addresses.

`dns_address` - (Optional) Configured address is chosen as DNS server address in DHCP response. (`String`).

`same_as_dgw` - (Optional) DNS server address is same as default gateway address (bool).

`dgw_address` - (Optional) Configured address from the network prefix is chosen as default gateway. (`String`).

`first_address` - (Optional) First usable address from the network prefix is chosen as default gateway (bool).

`last_address` - (Optional) Last usable address from the network prefix is chosen as default gateway (bool).

`network_prefix` - (Optional) Network Prefix for a single site. (`String`).

`network_prefix_allocator` - (Optional) Prefix length from address allocator scheme is used to calculate offsets. See [ref](#ref) below for details.

`pool_settings` - (Required) Controls how DHCP pools are handled (`String`).

`pools` - (Optional) List of non overlapping ip address ranges.. See [Pools ](#pools) below for details.

### Dhcp Server

DHCP Server is configured for this interface, Interface IP from DHCP server configuration..

`dhcp_networks` - (Required) List of networks from which DHCP server can allocate ip addresses. See [Dhcp Networks ](#dhcp-networks) below for details.

`dhcp_option82_tag` - (Optional) Optional tag that can be given to this configuration (`String`).

`fixed_ip_map` - (Optional) Fixed MAC address to ip assignments, Key: Mac address, Value: IP Address (`String`).

`automatic_from_end` - (Optional) Assign automatically from End of the first network in the list (bool).

`automatic_from_start` - (Optional) Assign automatically from start of the first network in the list (bool).

`interface_ip_map` - (Optional) Configured address for every node. See [Interface Ip Map ](#interface-ip-map) below for details.

### Dns Server

Address of DNS server when mode is "use-configured".

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ethernet Interface

Ethernet interface configuration..

`dhcp_client` - (Optional) Interface gets it IP address from external DHCP server (bool).

`dhcp_server` - (Optional) DHCP Server is configured for this interface, Interface IP from DHCP server configuration.. See [Dhcp Server ](#dhcp-server) below for details.

`static_ip` - (Optional) Interface IP is configured statically. See [Static Ip ](#static-ip) below for details.

`device` - (Required) Interface configuration for the ethernet device (`String`).

`no_ipv6_address` - (Optional) Interface does not have an IPv6 Address. (bool).

`static_ipv6_address` - (Optional) Interface IP is configured statically. See [Static Ipv6 Address ](#static-ipv6-address) below for details.

`monitor` - (Optional) Link Quality Monitoring parameters. Choosing the option will enable link quality monitoring.. See [Monitor ](#monitor) below for details.

`monitor_disabled` - (Optional) Link quality monitoring disabled on the interface. (bool).

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

`inside_network` - (Optional) Interface belongs to user configured inside network. See [ref](#ref) below for details.

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (bool).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (bool).

`srv6_network` - (Optional) Interface belongs to per site srv6 network. See [ref](#ref) below for details.

`storage_network` - (Optional) Interface belongs to site local network inside (bool).

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site. (bool).

`node` - (Optional) Configuration will apply to a device on the given node. (`String`).

`is_primary` - (Optional) This interface is primary (bool).

`not_primary` - (Optional) This interface is not primary (bool).

`priority` - (Optional) Greater the value, higher the priority (`Int`).

`untagged` - (Optional) Configure a untagged ethernet interface (bool).

`vlan_id` - (Optional) Configure a VLAN tagged ethernet interface (`Int`).

### First Address

First usable address from the network prefix is chosen as default gateway.

### Fleet Static Ip

Static IP configuration for the fleet.

`default_gw` - (Optional) IP address offset of the default gateway, prefix len is used to calculate offset (`String`).

`dns_server` - (Optional) IP address offset of the DNS server, prefix len is used to calculate offset (`String`).

`network_prefix_allocator` - (Optional) Static IP configuration for the fleet. See [ref](#ref) below for details.

### Interface Ip Map

Configured address for every node.

`interface_ip_map` - (Optional) Map of Site:Node to IP address. (`String`).

### Is Primary

This interface is primary.

### Last Address

Last usable address from the network prefix is chosen as default gateway.

### Legacy Interface

Old method of interface configuration.

`DHCP_server` - (Required) Behave as DHCP server for subnet configured in static addresses. (`String`).

`DNS_server` - (Optional) Configures how DNS server is derived for the subnet in static addresses. See [DNS Server ](#DNS-server) below for details.

`address_allocator` - (Optional) allocate a subnet for the interface and an address from the subnet is set on the interface.. See [ref](#ref) below for details.

`default_gateway` - (Optional) Configures how default gateway is derived for the subnet static addresses. See [Default Gateway ](#default-gateway) below for details.

`device_name` - (Required) Name of the physical network interface device which this network interface represents. (`String`).

`dhcp_address` - (Required) Enable DHCP based address assignment for this interface. (`String`).

`monitor` - (Optional) Link Quality Monitoring parameters. Choosing the option will enable link quality monitoring.. See [Monitor ](#monitor) below for details.

`monitor_disabled` - (Optional) Link quality monitoring disabled on the interface. (bool).

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

`priority` - (Optional) Greater the value, higher the priority (`Int`).

`static_addresses` - (Optional) If DHCP server is enabled, configures the subnet to be used for IP allocation.. See [Static Addresses ](#static-addresses) below for details.

`tunnel` - (Optional) When interface is created as TUNNEL type, then reference to tunnel is specified here. See [Tunnel ](#tunnel) below for details.

`type` - (Required) Specifies the type of interface (ethernet, vlan, lacp etc) (`String`).

`virtual_network` - (Optional) This is optional and can contain at most one entry. See [ref](#ref) below for details.

`vlan_tag` - (Optional) when vlan_tagging is enabled, value must be between 1 - 4094 (`Int`).

`vlan_tagging` - (Optional) Must be enabled for VLAN interfaces (`String`).

### Monitor

Link Quality Monitoring parameters. Choosing the option will enable link quality monitoring..

### Monitor Disabled

Link quality monitoring disabled on the interface..

### No Ipv6 Address

Interface does not have an IPv6 Address..

### Node Static Ip

Static IP configuration for the Node.

`default_gw` - (Optional) IP address of the default gateway. (`String`).

`dns_server` - (Optional) IP address of the DNS server (`String`).

`ip_address` - (Required) IP address of the interface and prefix length (`String`).

### Not Primary

This interface is not primary.

### Pools

List of non overlapping ip address ranges..

`end_ip` - (Optional) 10.1.1.200 with prefix length of 24, end offset is 0.0.0.200 (`String`).

`exclude` - (Optional) If exclude is true, IP addresses are not assigned from this range. (`Bool`).

`start_ip` - (Optional) 10.1.1.5 with prefix length of 24, start offset is 0.0.0.5 (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Same As Dgw

DNS server address is same as default gateway address.

### Site Local Inside Network

Interface belongs to site local network inside.

### Site Local Network

Interface belongs to site local network (outside).

### Static Addresses

If DHCP server is enabled, configures the subnet to be used for IP allocation..

`plen` - (Optional) Prefix-length of the IPv4 subnet. Must be <= 32 (`Int`).

`prefix` - (Optional) Prefix part of the IPv4 subnet in string form with dot-decimal notation (`String`).

### Static Ip

Interface IP is configured statically.

`cluster_static_ip` - (Optional) Static IP configuration for a specific node. See [Cluster Static Ip ](#cluster-static-ip) below for details.

`fleet_static_ip` - (Optional) Static IP configuration for the fleet. See [Fleet Static Ip ](#fleet-static-ip) below for details.

`node_static_ip` - (Optional) Static IP configuration for the Node. See [Node Static Ip ](#node-static-ip) below for details.

### Static Ipv6 Address

Interface IP is configured statically.

`cluster_static_ip` - (Optional) Static IP configuration for a specific node. See [Cluster Static Ip ](#cluster-static-ip) below for details.

`fleet_static_ip` - (Optional) Static IP configuration for the fleet. See [Fleet Static Ip ](#fleet-static-ip) below for details.

`node_static_ip` - (Optional) Static IP configuration for the Node. See [Node Static Ip ](#node-static-ip) below for details.

### Storage Network

Interface belongs to site local network inside.

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

### Untagged

Configure a untagged ethernet interface.

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_interface.
