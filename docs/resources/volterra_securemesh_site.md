---

page_title: "Volterra: securemesh_site"

description: "The securemesh_site allows CRUD of Securemesh Site resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_securemesh_site
=================================

The Securemesh Site allows CRUD of Securemesh Site resource on Volterra SaaS

~> **Note:** Please refer to [Securemesh Site API docs](https://docs.cloud.f5.com/docs-v2/api/views-securemesh-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_securemesh_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "blocked_services default_blocked_services" must be set

  default_blocked_services = true

  // One of the arguments from this list "bond_device_list no_bond_devices" must be set

  no_bond_devices = true

  // One of the arguments from this list "log_receiver logs_streaming_disabled" must be set

  logs_streaming_disabled = true
  master_node_configuration {
    name = "master-0"

    public_ip = "192.168.0.156"
  }

  // One of the arguments from this list "custom_network_config default_network_config" must be set

  default_network_config = true
  volterra_certified_hw = ["isv-8000-series-voltmesh"]
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

`address` - (Optional) Site's geographical address that can be used determine its latitude and longitude. (`String`).

###### One of the arguments from this list "blocked_services, default_blocked_services" must be set

`blocked_services` - (Optional) Use custom blocked services configuration. See [Blocked Services Choice Blocked Services ](#blocked-services-choice-blocked-services) below for details.

`default_blocked_services` - (Optional) Use default behavior which allows SSH (port 22), HTTPS (port 65500) and ICMP node access in blocked services (`Bool`).

###### One of the arguments from this list "bond_device_list, no_bond_devices" must be set

`bond_device_list` - (Optional) Configure Bond Devices for this Secure Mesh site. See [Bond Choice Bond Device List ](#bond-choice-bond-device-list) below for details.

`no_bond_devices` - (Optional) No Bond Devices configured for this Secure Mesh site (`Bool`).

`coordinates` - (Optional) Coordinates of the site, longitude and latitude. See [Coordinates ](#coordinates) below for details.

`kubernetes_upgrade_drain` - (Optional) Enable Kubernetes Drain during OS or SW upgrade. See [Kubernetes Upgrade Drain ](#kubernetes-upgrade-drain) below for details.

###### One of the arguments from this list "log_receiver, logs_streaming_disabled" must be set

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (`Bool`).

`master_node_configuration` - (Required) Configuration of master nodes. See [Master Node Configuration ](#master-node-configuration) below for details.

###### One of the arguments from this list "custom_network_config, default_network_config" must be set

`custom_network_config` - (Optional) Use custom networking configuration. See [Network Cfg Choice Custom Network Config ](#network-cfg-choice-custom-network-config) below for details.

`default_network_config` - (Optional) Use default networking configuration based on certified hardware. (`Bool`).

`offline_survivability_mode` - (Optional) Enable/Disable offline survivability mode. See [Offline Survivability Mode ](#offline-survivability-mode) below for details.

`os` - (Optional) Operating System Details. See [Os ](#os) below for details.

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode to optimize for L3 or L7 networking. See [Performance Enhancement Mode ](#performance-enhancement-mode) below for details.

`sw` - (Optional) F5XC Software Details. See [Sw ](#sw) below for details.

`volterra_certified_hw` - (Required) Name for generic server certified hardware to form this Secure Mesh site. (`String`).

`worker_nodes` - (Optional) Names of worker nodes (`List of String`).

### Coordinates

Coordinates of the site, longitude and latitude.

`latitude` - (Optional) Latitude of the site location (`Float`).

`longitude` - (Optional) longitude of site location (`Float`).

### Kubernetes Upgrade Drain

Enable Kubernetes Drain during OS or SW upgrade.

###### One of the arguments from this list "disable_upgrade_drain, enable_upgrade_drain" must be set

`disable_upgrade_drain` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_upgrade_drain` - (Optional) x-displayName: "Enable". See [Kubernetes Upgrade Drain Enable Choice Enable Upgrade Drain ](#kubernetes-upgrade-drain-enable-choice-enable-upgrade-drain) below for details.

### Master Node Configuration

Configuration of master nodes.

`name` - (Required) Names of master node (`String`).

`public_ip` - (Optional) via Site Mesh Group (`String`).

### Offline Survivability Mode

Enable/Disable offline survivability mode.

###### One of the arguments from this list "enable_offline_survivability_mode, no_offline_survivability_mode" must be set

`enable_offline_survivability_mode` - (Optional) x-displayName: "Enabled" (`Bool`).

`no_offline_survivability_mode` - (Optional) x-displayName: "Disabled" (`Bool`).

### Os

Operating System Details.

###### One of the arguments from this list "default_os_version, operating_system_version" must be set

`default_os_version` - (Optional) Will assign latest available OS version (`Bool`).

`operating_system_version` - (Optional) Specify a OS version to be used e.g. 9.2024.6. (`String`).

### Performance Enhancement Mode

Performance Enhancement Mode to optimize for L3 or L7 networking.

###### One of the arguments from this list "perf_mode_l3_enhanced, perf_mode_l7_enhanced" must be set

`perf_mode_l3_enhanced` - (Optional) Site optimized for L3 traffic processing. See [Perf Mode Choice Perf Mode L3 Enhanced ](#perf-mode-choice-perf-mode-l3-enhanced) below for details.

`perf_mode_l7_enhanced` - (Optional) Site optimized for L7 traffic processing (`Bool`).

### Sw

F5XC Software Details.

###### One of the arguments from this list "default_sw_version, volterra_software_version" must be set

`default_sw_version` - (Optional) Will assign latest available F5XC Software Version (`Bool`).

`volterra_software_version` - (Optional) Specify a F5XC Software Version to be used e.g. crt-20210329-1002. (`String`).

### Address Choice Dhcp Client

Interface gets it's IP address from external DHCP server.

### Address Choice Dhcp Server

DHCP Server is configured for this interface. IP for this Interface will be derived from the DHCP Server configuration..

`dhcp_networks` - (Required) List of networks from which DHCP Server can allocate IPv4 Addresses. See [Dhcp Server Dhcp Networks ](#dhcp-server-dhcp-networks) below for details.

`fixed_ip_map` - (Optional) Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client. (`String`).

###### One of the arguments from this list "automatic_from_end, automatic_from_start, interface_ip_map" must be set

`automatic_from_end` - (Optional) Assign automatically from end of the first network in the DHCP Network list (`Bool`).

`automatic_from_start` - (Optional) Assign automatically from start of the first network in the DHCP Network list (`Bool`).

`interface_ip_map` - (Optional) Statically configure a IPv4 address for every node. See [Interfaces Addressing Choice Interface Ip Map ](#interfaces-addressing-choice-interface-ip-map) below for details.

### Address Choice Stateful

works along with Router Advertisement' Managed flag.

`dhcp_networks` - (Required) List of networks from which DHCP server can allocate ip addresses. See [Stateful Dhcp Networks ](#stateful-dhcp-networks) below for details.

`fixed_ip_map` - (Optional) Assign fixed IPv6 addresses based on the MAC Address of the DHCP Client. (`String`).

###### One of the arguments from this list "automatic_from_end, automatic_from_start, interface_ip_map" must be set

`automatic_from_end` - (Optional) Assign automatically from End of the first network in the list (`Bool`).

`automatic_from_start` - (Optional) Assign automatically from start of the first network in the list (`Bool`).

`interface_ip_map` - (Optional) Configured address for every node. See [Interfaces Addressing Choice Interface Ip Map ](#interfaces-addressing-choice-interface-ip-map) below for details.

### Address Choice Static Ip

Interface IP is configured statically.

###### One of the arguments from this list "cluster_static_ip, node_static_ip" must be set

`cluster_static_ip` - (Optional) Static IP configuration for a specific node. See [Network Prefix Choice Cluster Static Ip ](#network-prefix-choice-cluster-static-ip) below for details.

`node_static_ip` - (Optional) Static IP configuration for the Node. See [Network Prefix Choice Node Static Ip ](#network-prefix-choice-node-static-ip) below for details.

### Autoconfig Choice Host

auto configuration routers. This is similar to a DHCP Client..

### Autoconfig Choice Router

System behaves like auto config Router and provides auto config parameters. This is similar to a DHCP Server..

###### One of the arguments from this list "network_prefix, stateful" must be set

`network_prefix` - (Optional) Allowed only /64 prefix length as per RFC 4862 (`String`).

`stateful` - (Optional) works along with Router Advertisement' Managed flag. See [Address Choice Stateful ](#address-choice-stateful) below for details.

`dns_config` - (Optional) Dns information that needs to added in the RouterAdvetisement. See [Router Dns Config ](#router-dns-config) below for details.

### Blocked Services Blocked Sevice

x-displayName: "Disable Node Local Services".

###### One of the arguments from this list "dns, ssh, web_user_interface" can be set

`dns` - (Optional) Matches DNS port 53 (`Bool`).

`ssh` - (Optional) x-displayName: "SSH" (`Bool`).

`web_user_interface` - (Optional) x-displayName: "Web UI" (`Bool`).

`network_type` - (Optional) Site Local VRF on which this service will be disabled (`String`).

### Blocked Services Choice Blocked Services

Use custom blocked services configuration.

`blocked_sevice` - (Optional) x-displayName: "Disable Node Local Services". See [Blocked Services Blocked Sevice ](#blocked-services-blocked-sevice) below for details.

### Blocked Services Value Type Choice Dns

Matches DNS port 53.

### Blocked Services Value Type Choice Ssh

x-displayName: "SSH".

### Blocked Services Value Type Choice Web User Interface

x-displayName: "Web UI".

### Bond Choice Bond Device List

Configure Bond Devices for this Secure Mesh site.

`bond_devices` - (Required) List of bond devices. See [Bond Device List Bond Devices ](#bond-device-list-bond-devices) below for details.

### Bond Device List Bond Devices

List of bond devices.

`devices` - (Required) Ethernet devices that will make up this bond (`String`).

###### One of the arguments from this list "active_backup, lacp" must be set

`active_backup` - (Optional) Configure active/backup based bond device (`Bool`).

`lacp` - (Optional) Configure LACP (802.3ad) based bond device. See [Lacp Choice Lacp ](#lacp-choice-lacp) below for details.

`link_polling_interval` - (Required) Link polling interval in milliseconds (`Int`).

`link_up_delay` - (Required) Milliseconds wait before link is declared up (`Int`).

`name` - (Required) Name for the Bond. Ex 'bond0' (`String`).

### Cluster Static Ip Interface Ip Map

Map of Node to Static ip configuration value, Key:Node, Value:IP Address.

`default_gw` - (Optional) IP address of the default gateway. (`String`).

`ip_address` - (Required) IP address of the interface and prefix length (`String`).

### Connection Choice Sli To Global Dr

Site local inside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Connection Choice Slo To Global Dr

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Dc Cluster Group Choice No Dc Cluster Group

This site is not a member of dc cluster group.

### Dc Cluster Group Connectivity Interface Choice Dc Cluster Group Connectivity Interface Disabled

Do not use this interface to connect to DC cluster group peers. .

### Dc Cluster Group Connectivity Interface Choice Dc Cluster Group Connectivity Interface Enabled

Use this interface to connect to DC cluster group peers..

### Dhcp Networks Pools

List of non overlapping ip address ranges..

`end_ip` - (Optional) In case of address allocator, offset is derived based on network prefix. (`String`).

`start_ip` - (Optional) 2001::1 with prefix length of 64, start offset is 5 (`String`).

### Dhcp Networks Pools

List of non overlapping ip address ranges..

`end_ip` - (Optional) 10.1.1.200 with prefix length of 24, end offset is 0.0.0.200 (`String`).

`start_ip` - (Optional) 10.1.1.5 with prefix length of 24, start offset is 0.0.0.5 (`String`).

### Dhcp Server Dhcp Networks

List of networks from which DHCP Server can allocate IPv4 Addresses.

###### One of the arguments from this list "dns_address, same_as_dgw" must be set

`dns_address` - (Optional) Enter a IPv4 address from the network prefix to be used as the DNS server. (`String`).

`same_as_dgw` - (Optional) DNS server address is same as default gateway address (`Bool`).

###### One of the arguments from this list "dgw_address, first_address, last_address" must be set

`dgw_address` - (Optional) Enter a IPv4 address from the network prefix to be used as the default gateway. (`String`).

`first_address` - (Optional) First usable address from the network prefix is chosen as default gateway (`Bool`).

`last_address` - (Optional) Last usable address from the network prefix is chosen as default gateway (`Bool`).

###### One of the arguments from this list "network_prefix" must be set

`network_prefix` - (Optional) Set the network prefix for the site. ex: 10.1.1.0/24 (`String`).

`pool_settings` - (Required) Controls how DHCP pools are handled (`String`).

`pools` - (Optional) List of non overlapping ip address ranges.. See [Dhcp Networks Pools ](#dhcp-networks-pools) below for details.

### Dns Choice Configured List

Configured address outside network range - external dns server.

`dns_list` - (Required) List of IPV6 Addresses acting as Dns servers (`String`).

### Dns Choice Local Dns

Choose the address from the network prefix range as dns server.

###### One of the arguments from this list "configured_address, first_address, last_address" must be set

`configured_address` - (Optional) Configured address from the network prefix is chosen as dns server (`String`).

`first_address` - (Optional) First usable address from the network prefix is chosen as dns server (`Bool`).

`last_address` - (Optional) Last usable address from the network prefix is chosen as dns server (`Bool`).

### Dns Choice Same As Dgw

DNS server address is same as default gateway address.

### Forward Proxy Choice Active Forward Proxy Policies

Enable Forward Proxy for this site and manage policies.

`forward_proxy_policies` - (Required) Ordered List of Forward Proxy Policies active. See [ref](#ref) below for details.

### Forward Proxy Choice Forward Proxy Allow All

Enable Forward Proxy for this site and allow all requests..

### Forward Proxy Choice No Forward Proxy

Disable Forward Proxy for this site.

### Gateway Choice First Address

First usable address from the network prefix is chosen as default gateway.

### Gateway Choice Last Address

Last usable address from the network prefix is chosen as default gateway.

### Global Network Choice Global Network List

List of global network connections.

`global_network_connections` - (Required) Global network connections. See [Global Network List Global Network Connections ](#global-network-list-global-network-connections) below for details.

### Global Network Choice No Global Network

No global network to connect.

### Global Network List Global Network Connections

Global network connections.

###### One of the arguments from this list "sli_to_global_dr, slo_to_global_dr" must be set

`sli_to_global_dr` - (Optional) Site local inside is connected directly to a given global network. See [Connection Choice Sli To Global Dr ](#connection-choice-sli-to-global-dr) below for details.

`slo_to_global_dr` - (Optional) Site local outside is connected directly to a given global network. See [Connection Choice Slo To Global Dr ](#connection-choice-slo-to-global-dr) below for details.

### Interface Choice Dedicated Interface

Networking configuration for dedicated interface is configured locally on site e.g. (outside/inside)Ethernet.

`device` - (Required) Name of the device for which interface is configured. Use wwan0 for 4G/LTE. (`String`).

###### One of the arguments from this list "monitor, monitor_disabled" can be set

`monitor` - (Optional) Link Quality Monitoring parameters. Choosing the option will enable link quality monitoring.. See [Monitoring Choice Monitor ](#monitoring-choice-monitor) below for details.

`monitor_disabled` - (Optional) Link quality monitoring disabled on the interface. (`Bool`).

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

###### One of the arguments from this list "cluster, node" must be set

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site. (`Bool`).

`node` - (Optional) Configuration will apply to a device on the given node of the site. (`String`).

###### One of the arguments from this list "is_primary, not_primary" must be set

`is_primary` - (Optional) This interface is primary (`Bool`).

`not_primary` - (Optional) This interface is not primary (`Bool`).

`priority` - (Optional) Greater the value, higher the priority (`Int`).

### Interface Choice Dedicated Management Interface

Fallback management interfaces can be made into dedicated management interface.

`device` - (Required) Name of the device for which interface is configured (`String`).

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

###### One of the arguments from this list "cluster, node" must be set

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site. (`Bool`).

`node` - (Optional) Configuration will apply to a device on the given node of the site. (`String`).

### Interface Choice Default Interface Config

Interface configuration is done based on certified hardware for this site.

### Interface Choice Ethernet Interface

Ethernet interface configuration..

###### One of the arguments from this list "dhcp_client, dhcp_server, static_ip" must be set

`dhcp_client` - (Optional) Interface gets it's IP address from external DHCP server (`Bool`).

`dhcp_server` - (Optional) DHCP Server is configured for this interface. IP for this Interface will be derived from the DHCP Server configuration.. See [Address Choice Dhcp Server ](#address-choice-dhcp-server) below for details.

`static_ip` - (Optional) Interface IP is configured statically. See [Address Choice Static Ip ](#address-choice-static-ip) below for details.

`device` - (Required) Interface configuration for the ethernet device (`String`).

###### One of the arguments from this list "ipv6_auto_config, no_ipv6_address, static_ipv6_address" can be set

`ipv6_auto_config` - (Optional) Configuration corresponding to IPV6 auto configuration. See [Ipv6 Address Choice Ipv6 Auto Config ](#ipv6-address-choice-ipv6-auto-config) below for details.

`no_ipv6_address` - (Optional) Interface does not have an IPv6 Address. (`Bool`).

`static_ipv6_address` - (Optional) Interface IP is configured statically. See [Ipv6 Address Choice Static Ipv6 Address ](#ipv6-address-choice-static-ipv6-address) below for details.

###### One of the arguments from this list "monitor, monitor_disabled" can be set

`monitor` - (Optional) Link Quality Monitoring parameters. Choosing the option will enable link quality monitoring.. See [Monitoring Choice Monitor ](#monitoring-choice-monitor) below for details.

`monitor_disabled` - (Optional) Link quality monitoring disabled on the interface. (`Bool`).

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

###### One of the arguments from this list "segment_network, site_local_inside_network, site_local_network, storage_network" must be set

`segment_network` - (Optional) x-displayName: "Segment". See [ref](#ref) below for details.

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (`Bool`).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (`Bool`).

`storage_network` - (Optional) Interface belongs to site local network inside (`Bool`).

###### One of the arguments from this list "cluster, node" must be set

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site. (`Bool`).

`node` - (Optional) Configuration will apply to a device on the given node. (`String`).

###### One of the arguments from this list "is_primary, not_primary" must be set

`is_primary` - (Optional) This interface is primary (`Bool`).

`not_primary` - (Optional) This interface is not primary (`Bool`).

`priority` - (Optional) Greater the value, higher the priority (`Int`).

###### One of the arguments from this list "untagged, vlan_id" must be set

`untagged` - (Optional) Configure a untagged ethernet interface (`Bool`).

`vlan_id` - (Optional) Configure a VLAN tagged ethernet interface (`Int`).

### Interface Choice Interface List

Add all interfaces belonging to this site.

`interfaces` - (Required) Configure network interfaces for this Secure Mesh site. See [Interface List Interfaces ](#interface-list-interfaces) below for details.

### Interface List Interfaces

Configure network interfaces for this Secure Mesh site.

###### One of the arguments from this list "dc_cluster_group_connectivity_interface_disabled, dc_cluster_group_connectivity_interface_enabled" must be set

`dc_cluster_group_connectivity_interface_disabled` - (Optional) Do not use this interface to connect to DC cluster group peers. (`Bool`).

`dc_cluster_group_connectivity_interface_enabled` - (Optional) Use this interface to connect to DC cluster group peers. (`Bool`).

`description` - (Optional) Description for this Interface (`String`).

###### One of the arguments from this list "dedicated_interface, dedicated_management_interface, ethernet_interface" must be set

`dedicated_interface` - (Optional) Networking configuration for dedicated interface is configured locally on site e.g. (outside/inside)Ethernet. See [Interface Choice Dedicated Interface ](#interface-choice-dedicated-interface) below for details.

`dedicated_management_interface` - (Optional) Fallback management interfaces can be made into dedicated management interface. See [Interface Choice Dedicated Management Interface ](#interface-choice-dedicated-management-interface) below for details.

`ethernet_interface` - (Optional) Ethernet interface configuration.. See [Interface Choice Ethernet Interface ](#interface-choice-ethernet-interface) below for details.

`labels` - (Optional) Add Labels for this Interface, these labels can be used in firewall policy (`String`).

### Interfaces Addressing Choice Automatic From End

Assign automatically from end of the first network in the DHCP Network list.

### Interfaces Addressing Choice Automatic From Start

Assign automatically from start of the first network in the DHCP Network list.

### Interfaces Addressing Choice Interface Ip Map

Statically configure a IPv4 address for every node.

`interface_ip_map` - (Optional) Specify static IPv4 addresses per site:node. (`String`).

### Interfaces Addressing Choice Interface Ip Map

Configured address for every node.

`interface_ip_map` - (Optional) Map of Site:Node to IPV6 address. (`String`).

### Ipv6 Address Choice Ipv6 Auto Config

Configuration corresponding to IPV6 auto configuration.

###### One of the arguments from this list "host, router" must be set

`host` - (Optional) auto configuration routers. This is similar to a DHCP Client. (`Bool`).

`router` - (Optional) System behaves like auto config Router and provides auto config parameters. This is similar to a DHCP Server.. See [Autoconfig Choice Router ](#autoconfig-choice-router) below for details.

### Ipv6 Address Choice No Ipv6 Address

Interface does not have an IPv6 Address..

### Ipv6 Address Choice Static Ipv6 Address

Interface IP is configured statically.

###### One of the arguments from this list "cluster_static_ip, node_static_ip" must be set

`cluster_static_ip` - (Optional) Static IP configuration for a specific node. See [Network Prefix Choice Cluster Static Ip ](#network-prefix-choice-cluster-static-ip) below for details.

`node_static_ip` - (Optional) Static IP configuration for the Node. See [Network Prefix Choice Node Static Ip ](#network-prefix-choice-node-static-ip) below for details.

### Kubernetes Upgrade Drain Enable Choice Disable Upgrade Drain

x-displayName: "Disable".

### Kubernetes Upgrade Drain Enable Choice Enable Upgrade Drain

x-displayName: "Enable".

###### One of the arguments from this list "drain_max_unavailable_node_count" must be set

`drain_max_unavailable_node_count` - (Optional) x-example: "1" (`Int`).

`drain_node_timeout` - (Required) (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is recommended to use the default value). (`Int`).

### Lacp Choice Active Backup

Configure active/backup based bond device.

### Lacp Choice Lacp

Configure LACP (802.3ad) based bond device.

`rate` - (Optional) Interval in seconds to transmit LACP packets (`Int`).

### Local Dns Choice First Address

First usable address from the network prefix is chosen as dns server.

### Local Dns Choice Last Address

Last usable address from the network prefix is chosen as dns server.

### Monitoring Choice Monitor

Link Quality Monitoring parameters. Choosing the option will enable link quality monitoring..

### Monitoring Choice Monitor Disabled

Link quality monitoring disabled on the interface..

### Network Cfg Choice Custom Network Config

Use custom networking configuration.

###### One of the arguments from this list "active_forward_proxy_policies, forward_proxy_allow_all, no_forward_proxy" must be set

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Forward Proxy Choice Active Forward Proxy Policies ](#forward-proxy-choice-active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (`Bool`).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (`Bool`).

###### One of the arguments from this list "global_network_list, no_global_network" must be set

`global_network_list` - (Optional) List of global network connections. See [Global Network Choice Global Network List ](#global-network-choice-global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (`Bool`).

###### One of the arguments from this list "default_interface_config, interface_list" must be set

`default_interface_config` - (Optional) Interface configuration is done based on certified hardware for this site (`Bool`).

`interface_list` - (Optional) Add all interfaces belonging to this site. See [Interface Choice Interface List ](#interface-choice-interface-list) below for details.

###### One of the arguments from this list "active_enhanced_firewall_policies, active_network_policies, no_network_policy" must be set

`active_enhanced_firewall_policies` - (Optional) with an additional option for service insertion.. See [Network Policy Choice Active Enhanced Firewall Policies ](#network-policy-choice-active-enhanced-firewall-policies) below for details.

`active_network_policies` - (Optional) Firewall Policies active for this site.. See [Network Policy Choice Active Network Policies ](#network-policy-choice-active-network-policies) below for details.

`no_network_policy` - (Optional) Firewall Policy is disabled for this site. (`Bool`).

###### One of the arguments from this list "sm_connection_public_ip, sm_connection_pvt_ip" must be set

`sm_connection_public_ip` - (Optional) which are part of the site mesh group (`Bool`).

`sm_connection_pvt_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (`Bool`).

###### One of the arguments from this list "default_sli_config, sli_config" can be set

`default_sli_config` - (Optional) Use default configuration for site local network (`Bool`).

`sli_config` - (Optional) Configuration for site local inside network. See [Sli Choice Sli Config ](#sli-choice-sli-config) below for details.

###### One of the arguments from this list "default_config, slo_config" must be set

`default_config` - (Optional) Use default configuration for site local network (`Bool`).

`slo_config` - (Optional) Configuration for site local network. See [Slo Choice Slo Config ](#slo-choice-slo-config) below for details.

`tunnel_dead_timeout` - (Optional) When not set (== 0), a default value of 10000 msec will be used. (`Int`).

`vip_vrrp_mode` - (Optional) When Outside VIP / Inside VIP are configured, it is recommended to turn on vrrp and also configure BGP. (`String`).

### Network Choice Site Local Inside Network

Interface belongs to site local network inside.

### Network Choice Site Local Network

Interface belongs to site local network (outside).

### Network Choice Storage Network

Interface belongs to site local network inside.

### Network Policy Choice Active Enhanced Firewall Policies

with an additional option for service insertion..

`enhanced_firewall_policies` - (Required) Ordered List of Enhanced Firewall Policies active. See [ref](#ref) below for details.

### Network Policy Choice Active Network Policies

Firewall Policies active for this site..

`network_policies` - (Required) Ordered List of Firewall Policies active for this network firewall. See [ref](#ref) below for details.

### Network Policy Choice No Network Policy

Firewall Policy is disabled for this site..

### Network Prefix Choice Cluster Static Ip

Static IP configuration for a specific node.

`interface_ip_map` - (Optional) Map of Node to Static ip configuration value, Key:Node, Value:IP Address. See [Cluster Static Ip Interface Ip Map ](#cluster-static-ip-interface-ip-map) below for details.

### Network Prefix Choice Node Static Ip

Static IP configuration for the Node.

`default_gw` - (Optional) IP address of the default gateway. (`String`).

`ip_address` - (Required) IP address of the interface and prefix length (`String`).

### Next Hop Choice Default Gateway

Traffic matching the ip prefixes is sent to the default gateway.

### Node Choice Cluster

Configuration will apply to given device on all nodes of the site..

### Offline Survivability Mode Choice Enable Offline Survivability Mode

x-displayName: "Enabled".

### Offline Survivability Mode Choice No Offline Survivability Mode

x-displayName: "Disabled".

### Operating System Version Choice Default Os Version

Will assign latest available OS version.

### Perf Mode Choice Jumbo

x-displayName: "Enabled".

### Perf Mode Choice No Jumbo

x-displayName: "Disabled".

### Perf Mode Choice Perf Mode L3 Enhanced

Site optimized for L3 traffic processing.

###### One of the arguments from this list "jumbo, no_jumbo" must be set

`jumbo` - (Optional) x-displayName: "Enabled" (`Bool`).

`no_jumbo` - (Optional) x-displayName: "Disabled" (`Bool`).

### Perf Mode Choice Perf Mode L7 Enhanced

Site optimized for L7 traffic processing.

### Primary Choice Is Primary

This interface is primary.

### Primary Choice Not Primary

This interface is not primary.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Router Dns Config

Dns information that needs to added in the RouterAdvetisement.

###### One of the arguments from this list "configured_list, local_dns" must be set

`configured_list` - (Optional) Configured address outside network range - external dns server. See [Dns Choice Configured List ](#dns-choice-configured-list) below for details.

`local_dns` - (Optional) Choose the address from the network prefix range as dns server. See [Dns Choice Local Dns ](#dns-choice-local-dns) below for details.

### Site Mesh Group Choice Sm Connection Public Ip

which are part of the site mesh group.

### Site Mesh Group Choice Sm Connection Pvt Ip

creating ipsec between two sites which are part of the site mesh group.

### Sli Choice Default Sli Config

Use default configuration for site local network.

### Sli Choice Sli Config

Configuration for site local inside network.

###### One of the arguments from this list "dc_cluster_group, no_dc_cluster_group" must be set

`dc_cluster_group` - (Optional) This site is member of dc cluster group via network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (`Bool`).

`labels` - (Optional) Add Labels for this network, these labels can be used in firewall policy (`String`).

`nameserver` - (Optional) Optional DNS V4 server IP to be used for name resolution (`String`).

`nameserver_v6` - (Optional) Optional DNS V6 server IP to be used for name resolution (`String`).

###### One of the arguments from this list "no_static_routes, static_routes" must be set

`no_static_routes` - (Optional) Static Routes disabled for site local network. (`Bool`).

`static_routes` - (Optional) Manage static routes for site local network.. See [Static Route Choice Static Routes ](#static-route-choice-static-routes) below for details.

###### One of the arguments from this list "no_v6_static_routes, static_v6_routes" must be set

`no_v6_static_routes` - (Optional) Static IPv6 Routes disabled for site local network. (`Bool`).

`static_v6_routes` - (Optional) Manage IPv6 static routes for site local network.. See [Static V6 Route Choice Static V6 Routes ](#static-v6-route-choice-static-v6-routes) below for details.

`vip` - (Optional) Optional common virtual V4 IP across all nodes to be used as automatic VIP. (`String`).

`vip_v6` - (Optional) Optional common virtual V6 IP across all nodes to be used as automatic VIP. (`String`).

### Slo Choice Default Config

Use default configuration for site local network.

### Slo Choice Slo Config

Configuration for site local network.

###### One of the arguments from this list "dc_cluster_group, no_dc_cluster_group" must be set

`dc_cluster_group` - (Optional) This site is member of dc cluster group via network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (`Bool`).

`labels` - (Optional) Add Labels for this network, these labels can be used in firewall policy (`String`).

`nameserver` - (Optional) Optional DNS V4 server IP to be used for name resolution (`String`).

`nameserver_v6` - (Optional) Optional DNS V6 server IP to be used for name resolution (`String`).

###### One of the arguments from this list "no_static_routes, static_routes" must be set

`no_static_routes` - (Optional) Static Routes disabled for site local network. (`Bool`).

`static_routes` - (Optional) Manage static routes for site local network.. See [Static Route Choice Static Routes ](#static-route-choice-static-routes) below for details.

###### One of the arguments from this list "no_v6_static_routes, static_v6_routes" must be set

`no_v6_static_routes` - (Optional) Static IPv6 Routes disabled for site local network. (`Bool`).

`static_v6_routes` - (Optional) Manage IPv6 static routes for site local network.. See [Static V6 Route Choice Static V6 Routes ](#static-v6-route-choice-static-v6-routes) below for details.

`vip` - (Optional) Optional common virtual V4 IP across all nodes to be used as automatic VIP. (`String`).

`vip_v6` - (Optional) Optional common virtual V6 IP across all nodes to be used as automatic VIP. (`String`).

### Stateful Dhcp Networks

List of networks from which DHCP server can allocate ip addresses.

###### One of the arguments from this list "network_prefix" must be set

`network_prefix` - (Optional) Network Prefix to be used for IPV6 address auto configuration (`String`).

`pool_settings` - (Required) Controls how DHCPV6 pools are handled (`String`).

`pools` - (Optional) List of non overlapping ip address ranges.. See [Dhcp Networks Pools ](#dhcp-networks-pools) below for details.

### Static Route Choice No Static Routes

Static Routes disabled for site local network..

### Static Route Choice Static Routes

Manage static routes for site local network..

`static_routes` - (Required) List of static routes. See [Static Routes Static Routes ](#static-routes-static-routes) below for details.

### Static Routes Static Routes

List of static routes.

`attrs` - (Optional) List of attributes that control forwarding, dynamic routing and control plane (host) reachability (`List of Strings`).

`ip_prefixes` - (Required) List of route prefixes that have common next hop and attributes (`String`).

###### One of the arguments from this list "default_gateway, interface, ip_address" must be set

`default_gateway` - (Optional) Traffic matching the ip prefixes is sent to the default gateway (`Bool`).

`interface` - (Optional) Traffic matching the ip prefixes is sent to this interface. See [ref](#ref) below for details.

`ip_address` - (Optional) Traffic matching the ip prefixes is sent to this IP Address (`String`).

### Static V6 Route Choice No V6 Static Routes

Static IPv6 Routes disabled for site local network..

### Static V6 Route Choice Static V6 Routes

Manage IPv6 static routes for site local network..

`static_routes` - (Required) List of IPv6 static routes. See [Static V6 Routes Static Routes ](#static-v6-routes-static-routes) below for details.

### Static V6 Routes Static Routes

List of IPv6 static routes.

`attrs` - (Optional) List of attributes that control forwarding, dynamic routing and control plane (host) reachability (`List of Strings`).

`ip_prefixes` - (Required) List of IPv6 route prefixes that have common next hop and attributes (`String`).

###### One of the arguments from this list "default_gateway, interface, ip_address" must be set

`default_gateway` - (Optional) Traffic matching the ip prefixes is sent to the default gateway (`Bool`).

`interface` - (Optional) Traffic matching the ip prefixes is sent to this interface. See [ref](#ref) below for details.

`ip_address` - (Optional) Traffic matching the ip prefixes is sent to this IP Address (`String`).

### Vlan Choice Untagged

Configure a untagged ethernet interface.

### Volterra Sw Version Choice Default Sw Version

Will assign latest available F5XC Software Version.

Attribute Reference
-------------------

-	`id` - This is the id of the configured securemesh_site.
