---

page_title: "Volterra: securemesh_site_v2"
description: "The securemesh_site_v2 allows CRUD of Securemesh Site V2 resource on Volterra SaaS"

---

Resource volterra_securemesh_site_v2
====================================

The Securemesh Site V2 allows CRUD of Securemesh Site V2 resource on Volterra SaaS

~> **Note:** Please refer to [Securemesh Site V2 API docs](https://docs.cloud.f5.com/docs-v2/api/views-securemesh-site-v2) to learn more

Example Usage
-------------

```hcl
resource "volterra_securemesh_site_v2" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "block_all_services blocked_services" must be set

  block_all_services = true

  // One of the arguments from this list "log_receiver logs_streaming_disabled" must be set

  logs_streaming_disabled = true

  // One of the arguments from this list "aws azure baremetal gcp kvm oci rseries vmware" must be set

  azure {
    // One of the arguments from this list "not_managed" can be set

    not_managed {
      node_list {
        hostname = "Control"

        interface_list {
          // One of the arguments from this list "dhcp_client dhcp_server no_ipv4_address static_ip" must be set

          static_ip {
            default_gw = "192.168.20.1"

            dns_server = "192.168.20.1"

            ip_address = "192.168.20.1/24"
          }

          description = "value"

          // One of the arguments from this list "bond_interface ethernet_interface vlan_interface" must be set

          bond_interface {
            devices = ["eth0"]

            // One of the arguments from this list "active_backup lacp" must be set

            lacp {
              rate = "30"
            }
            link_polling_interval = "1000"
            link_up_delay = "200"
            name = "bond0"
          }

          // One of the arguments from this list "ipv6_auto_config no_ipv6_address static_ipv6_address" can be set

          no_ipv6_address = true
          is_management = true
          is_primary = true
          labels = {
            "key1" = "value1"
          }

          // One of the arguments from this list "monitor monitor_disabled" can be set

          monitor_disabled = true
          mtu = "1450"
          name = "value"
          network_option {
            // One of the arguments from this list "segment_network site_local_inside_network site_local_network" can be set

            site_local_network = true
          }
          priority = "42"

          // One of the arguments from this list "site_to_site_connectivity_interface_disabled site_to_site_connectivity_interface_enabled" can be set

          site_to_site_connectivity_interface_disabled = true
        }

        public_ip = "1.1.1.1"

        type = "Control"
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

###### One of the arguments from this list "block_all_services, blocked_services" must be set

`block_all_services` - (Optional) Enable WebUI, SSH and DNS on all nodes in this site. (`Bool`).

`blocked_services` - (Optional) It is recommended to disable node local services after the nodes register or after configuration/deugging is complete.. See [Blocked Services Choice Blocked Services ](#blocked-services-choice-blocked-services) below for details.

###### One of the arguments from this list "active_forward_proxy_policies, no_forward_proxy" can be set

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site. Traffic will be processed in the order that Forward Proxy Policies are added.. See [Forward Proxy Choice Active Forward Proxy Policies ](#forward-proxy-choice-active-forward-proxy-policies) below for details.

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site. (`Bool`).

`load_balancing` - (Optional) This section contains settings on the site that relate to Load Balancing functionality.. See [Load Balancing ](#load-balancing) below for details.

`local_vrf` - (Optional) The Site Local Inside (SLI) local VRF is used to connect LAN side workloads to this site. SLI local VRF is optional.. See [Local Vrf ](#local-vrf) below for details.

###### One of the arguments from this list "log_receiver, logs_streaming_disabled" must be set

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) x-displayName: "Disable Logs Streaming" (`Bool`).

###### One of the arguments from this list "active_enhanced_firewall_policies, no_network_policy" can be set

`active_enhanced_firewall_policies` - (Optional) Enable Network Firewall for this site. Traffic will be processed in the order that Network Firewall Policies are added.. See [Network Policy Choice Active Enhanced Firewall Policies ](#network-policy-choice-active-enhanced-firewall-policies) below for details.

`no_network_policy` - (Optional) Disable Network Firewall for this site. (`Bool`).

###### One of the arguments from this list "disable_ha, enable_ha" can be set

`disable_ha` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_ha` - (Optional) x-displayName: "Enable" (`Bool`).

`offline_survivability_mode` - (Optional) When the mode is toggled, services will restart and traffic disruption will be seen.. See [Offline Survivability Mode ](#offline-survivability-mode) below for details.

`performance_enhancement_mode` - (Optional) Optimize the site for L3 or L7 traffic processing. By default, the site is optimized for L7 traffic processing.. See [Performance Enhancement Mode ](#performance-enhancement-mode) below for details.

###### One of the arguments from this list "aws, azure, baremetal, gcp, kvm, oci, rseries, vmware" must be set

`aws` - (Optional) x-displayName: "AWS". See [Provider Choice Aws ](#provider-choice-aws) below for details.

`azure` - (Optional) x-displayName: "Azure". See [Provider Choice Azure ](#provider-choice-azure) below for details.

`baremetal` - (Optional) x-displayName: "Baremetal". See [Provider Choice Baremetal ](#provider-choice-baremetal) below for details.

`gcp` - (Optional) x-displayName: "GCP". See [Provider Choice Gcp ](#provider-choice-gcp) below for details.

`kvm` - (Optional) x-displayName: "KVM". See [Provider Choice Kvm ](#provider-choice-kvm) below for details.

`oci` - (Optional) x-displayName: "OCI". See [Provider Choice Oci ](#provider-choice-oci) below for details.

`rseries` - (Optional) x-displayName: "F5 rSeries". See [Provider Choice Rseries ](#provider-choice-rseries) below for details.

`vmware` - (Optional) x-displayName: "VMWare". See [Provider Choice Vmware ](#provider-choice-vmware) below for details.

`re_select` - (Optional) Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).. See [Re Select ](#re-select) below for details.

###### One of the arguments from this list "dc_cluster_group_sli, no_s2s_connectivity_sli" can be set

`dc_cluster_group_sli` - (Optional) Use a DC Cluster Group to connect to other sites.. See [ref](#ref) below for details.

`no_s2s_connectivity_sli` - (Optional) x-displayName: "Disabled" (`Bool`).

###### One of the arguments from this list "dc_cluster_group_slo, no_s2s_connectivity_slo, site_mesh_group_on_slo" can be set

`dc_cluster_group_slo` - (Optional) Use a DC Cluster Group to connect to other sites.. See [ref](#ref) below for details.

`no_s2s_connectivity_slo` - (Optional) x-displayName: "Disabled" (`Bool`).

`site_mesh_group_on_slo` - (Optional) Use a Site Mesh Group to connect to other sites.. See [S2s Connectivity Slo Choice Site Mesh Group On Slo ](#s2s-connectivity-slo-choice-site-mesh-group-on-slo) below for details.

`software_settings` - (Optional) Select OS and Software version for the site. All nodes in the site will run the same OS and Software version. These settings cannot be changed after the site is created.. See [Software Settings ](#software-settings) below for details.

`tunnel_dead_timeout` - (Optional) When not set (== 0), a default value of 10000 msec will be used. (`Int`).

`tunnel_type` - (Optional) Select the type of tunnel to be used for traffic between the site and REs. By default, IPsec will be preferred with SSL as backup. (`String`).

`upgrade_settings` - (Optional) Specify how a site will be upgraded.. See [Upgrade Settings ](#upgrade-settings) below for details.

### Load Balancing

This section contains settings on the site that relate to Load Balancing functionality..

`vip_vrrp_mode` - (Optional) When VIP is configured on both Site Local Outside (SLO) and Site Local Inside (SLI) Local VRF on the site, it is recommended to turn on VRRP and also configure BGP. (`String`).

### Local Vrf

The Site Local Inside (SLI) local VRF is used to connect LAN side workloads to this site. SLI local VRF is optional..

###### One of the arguments from this list "default_sli_config, sli_config" can be set

`default_sli_config` - (Optional) x-displayName: "Default Configuration" (`Bool`).

`sli_config` - (Optional) Configure properties such as static routes, DNS and common VIP for Load Balancing on the Site Local Inside (SLI) local VRF.. See [Sli Choice Sli Config ](#sli-choice-sli-config) below for details.

###### One of the arguments from this list "default_config, slo_config" must be set

`default_config` - (Optional) x-displayName: "Default Configuration" (`Bool`).

`slo_config` - (Optional) Configure properties such as static routes, DNS and common VIP for Load Balancing on the Site Local Outside (SLO) local VRF.. See [Slo Choice Slo Config ](#slo-choice-slo-config) below for details.

### Offline Survivability Mode

When the mode is toggled, services will restart and traffic disruption will be seen..

###### One of the arguments from this list "enable_offline_survivability_mode, no_offline_survivability_mode" must be set

`enable_offline_survivability_mode` - (Optional) x-displayName: "Enabled" (`Bool`).

`no_offline_survivability_mode` - (Optional) x-displayName: "Disabled" (`Bool`).

### Performance Enhancement Mode

Optimize the site for L3 or L7 traffic processing. By default, the site is optimized for L7 traffic processing..

###### One of the arguments from this list "perf_mode_l3_enhanced, perf_mode_l7_enhanced" must be set

`perf_mode_l3_enhanced` - (Optional) Site optimized for L3 traffic processing. See [Perf Mode Choice Perf Mode L3 Enhanced ](#perf-mode-choice-perf-mode-l3-enhanced) below for details.

`perf_mode_l7_enhanced` - (Optional) Site optimized for L7 traffic processing (`Bool`).

### Re Select

Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s)..

###### One of the arguments from this list "geo_proximity, specific_geography, specific_re" can be set

`geo_proximity` - (Optional) Select REs in closest proximity to the site based on the public IP address of the control nodes of the site. (`Bool`).

`specific_geography` - (Optional) Select a list of specific REs. This is useful when a site needs to deterministically connect to a set of REs. A site will always be connected to 2 REs. If >2 REs are chosen, then 2 REs from these will be selected. (`String`).(Deprecated)

`specific_re` - (Optional) Select specific REs. This is useful when a site needs to deterministically connect to a set of REs. A site will always be connected to 2 REs.. See [Re Selection Choice Specific Re ](#re-selection-choice-specific-re) below for details.

### Software Settings

Select OS and Software version for the site. All nodes in the site will run the same OS and Software version. These settings cannot be changed after the site is created..

`os` - (Optional) Select the Operating System version for the site. By default, latest available Operating System version will be used.. See [Software Settings Os ](#software-settings-os) below for details.

`sw` - (Optional) Refer to release notes to find required released SW versions.. See [Software Settings Sw ](#software-settings-sw) below for details.

### Upgrade Settings

Specify how a site will be upgraded..

`kubernetes_upgrade_drain` - (Optional) Specify how worker nodes within a site will be upgraded.. See [Upgrade Settings Kubernetes Upgrade Drain ](#upgrade-settings-kubernetes-upgrade-drain) below for details.

### Address Choice Dhcp Client

Interface gets it's IP address from an external DHCP server..

### Address Choice Dhcp Server

DHCP Server is configured for this interface, Interface IP is derived from DHCP server configuration..

`dhcp_networks` - (Required) List of networks from which DHCP Server can allocate IPv4 Addresses. See [Dhcp Server Dhcp Networks ](#dhcp-server-dhcp-networks) below for details.

`dhcp_option82_tag` - (Optional) Optional tag that can be given to this configuration (`String`).(Deprecated)

`fixed_ip_map` - (Optional) Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client. (`String`).

###### One of the arguments from this list "automatic_from_end, automatic_from_start, interface_ip_map" must be set

`automatic_from_end` - (Optional) Assign automatically from end of the first network in the DHCP Network list (`Bool`).

`automatic_from_start` - (Optional) Assign automatically from start of the first network in the DHCP Network list (`Bool`).

`interface_ip_map` - (Optional) Statically configure a IPv4 address for every node. See [Interfaces Addressing Choice Interface Ip Map ](#interfaces-addressing-choice-interface-ip-map) below for details.

### Address Choice No Ipv4 Address

Interface does not have an IPv4 Address..

### Address Choice Stateful

works along with Router Advertisement' Managed flag.

`dhcp_networks` - (Required) List of networks from which DHCP server can allocate ip addresses. See [Stateful Dhcp Networks ](#stateful-dhcp-networks) below for details.

`fixed_ip_map` - (Optional) Assign fixed IPv6 addresses based on the MAC Address of the DHCP Client. (`String`).

###### One of the arguments from this list "automatic_from_end, automatic_from_start, interface_ip_map" must be set

`automatic_from_end` - (Optional) Assign automatically from End of the first network in the list (`Bool`).

`automatic_from_start` - (Optional) Assign automatically from start of the first network in the list (`Bool`).

`interface_ip_map` - (Optional) Configured address for every node. See [Interfaces Addressing Choice Interface Ip Map ](#interfaces-addressing-choice-interface-ip-map) below for details.

### Address Choice Static Ip

Interface IP address is configured statically..

`default_gw` - (Optional) IP address of the default gateway. (`String`).

`dns_server` - (Optional) IP address of the DNS server (`String`).(Deprecated)

`ip_address` - (Required) IP address of the interface and prefix length (`String`).

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

It is recommended to disable node local services after the nodes register or after configuration/deugging is complete..

`blocked_sevice` - (Optional) x-displayName: "Disable Node Local Services". See [Blocked Services Blocked Sevice ](#blocked-services-blocked-sevice) below for details.

### Blocked Services Value Type Choice Dns

Matches DNS port 53.

### Blocked Services Value Type Choice Ssh

x-displayName: "SSH".

### Blocked Services Value Type Choice Web User Interface

x-displayName: "Web UI".

### Cluster Static Ip Interface Ip Map

Map of Node to Static ip configuration value, Key:Node, Value:IP Address.

`default_gw` - (Optional) IP address of the default gateway. (`String`).

`dns_server` - (Optional) IP address of the DNS server (`String`).(Deprecated)

`ip_address` - (Required) IP address of the interface and prefix length (`String`).

### Dhcp Networks Pools

List of non overlapping ip address ranges..

`end_ip` - (Optional) In case of address allocator, offset is derived based on network prefix. (`String`).

`exclude` - (Optional) If exclude is true, IP addresses are not assigned from this range. (`Bool`).(Deprecated)

`start_ip` - (Optional) 2001::1 with prefix length of 64, start offset is 5 (`String`).

### Dhcp Networks Pools

List of non overlapping ip address ranges..

`end_ip` - (Optional) 10.1.1.200 with prefix length of 24, end offset is 0.0.0.200 (`String`).

`exclude` - (Optional) If exclude is true, IP addresses are not assigned from this range. (`Bool`).(Deprecated)

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

###### One of the arguments from this list "network_prefix, network_prefix_allocator" must be set

`network_prefix` - (Optional) Set the network prefix for the site. ex: 10.1.1.0/24 (`String`).

`network_prefix_allocator` - (Optional) Prefix length from address allocator scheme is used to calculate offsets. See [ref](#ref) below for details.(Deprecated)

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

Enable Forward Proxy for this site. Traffic will be processed in the order that Forward Proxy Policies are added..

`forward_proxy_policies` - (Required) Ordered List of Forward Proxy Policies active. See [ref](#ref) below for details.

### Gateway Choice First Address

First usable address from the network prefix is chosen as default gateway.

### Gateway Choice Last Address

Last usable address from the network prefix is chosen as default gateway.

### Interface Choice Bond Interface

x-displayName: "Bond Interface".

`devices` - (Required) Ethernet devices that will make up this bond (`String`).

###### One of the arguments from this list "active_backup, lacp" must be set

`active_backup` - (Optional) Configure active/backup based bond device (`Bool`).

`lacp` - (Optional) Configure LACP (802.3ad) based bond device. See [Lacp Choice Lacp ](#lacp-choice-lacp) below for details.

`link_polling_interval` - (Required) Link polling interval in milliseconds (`Int`).

`link_up_delay` - (Required) Milliseconds wait before link is declared up (`Int`).

`name` - (Required) Name for the Bond. Ex 'bond0' (`String`).

### Interface Choice Ethernet Interface

x-displayName: "Ethernet Interface".

`device` - (Required) Once configured, this interface will be part of this sites dataplane and can participate in the networking services configured on this site. (`String`).

`mac` - (Optional) x-example: "01:10:20:0a:bb:1c" (`String`).

### Interface Choice Vlan Interface

x-displayName: "VLAN Interface".

`device` - (Required) Select a parent interface from the dropdown. (`String`).

`vlan_id` - (Optional) Configure the VLAN tag for this interface. (`Int`).

### Interface List Network Option

Global VRFs are configured via Networking > Segments. A site can have multple Network Segments (global VRFs)..

###### One of the arguments from this list "segment_network, site_local_inside_network, site_local_network" can be set

`segment_network` - (Optional) x-displayName: "Segment (Global VRF)". See [ref](#ref) below for details.

`site_local_inside_network` - (Optional) x-displayName: "Site Local Inside (Local VRF)" (`Bool`).

`site_local_network` - (Optional) x-displayName: "Site Local Outside (Local VRF)" (`Bool`).

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

Interface IPv6 address will be configured via Auto Configuration..

###### One of the arguments from this list "host, router" must be set

`host` - (Optional) auto configuration routers. This is similar to a DHCP Client. (`Bool`).

`router` - (Optional) System behaves like auto config Router and provides auto config parameters. This is similar to a DHCP Server.. See [Autoconfig Choice Router ](#autoconfig-choice-router) below for details.

### Ipv6 Address Choice No Ipv6 Address

Interface does not have an IPv6 Address..

### Ipv6 Address Choice Static Ipv6 Address

Interface IPv6 address is configured statically..

###### One of the arguments from this list "cluster_static_ip, fleet_static_ip, node_static_ip" must be set

`cluster_static_ip` - (Optional) Static IP configuration for a specific node. See [Network Prefix Choice Cluster Static Ip ](#network-prefix-choice-cluster-static-ip) below for details.

`fleet_static_ip` - (Optional) Static IP configuration for the fleet. See [Network Prefix Choice Fleet Static Ip ](#network-prefix-choice-fleet-static-ip) below for details.(Deprecated)

`node_static_ip` - (Optional) Static IP configuration for the Node. See [Network Prefix Choice Node Static Ip ](#network-prefix-choice-node-static-ip) below for details.

### Kubernetes Upgrade Drain Enable Choice Disable Upgrade Drain

x-displayName: "Disable Node by Node Upgrade".

### Kubernetes Upgrade Drain Enable Choice Enable Upgrade Drain

x-displayName: "Enable Node by Node Upgrade".

###### One of the arguments from this list "drain_max_unavailable_node_count, drain_max_unavailable_node_percentage" must be set

`drain_max_unavailable_node_count` - (Optional) x-example: "1" (`Int`).

`drain_max_unavailable_node_percentage` - (Optional) Max number of worker nodes to be upgraded in parallel by percentage. Note: 1% would mean batch size of 1 worker node. (`Int`).(Deprecated)

`drain_node_timeout` - (Required) (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is recommended to use the default value). (`Int`).

###### One of the arguments from this list "disable_vega_upgrade_mode, enable_vega_upgrade_mode" must be set

`disable_vega_upgrade_mode` - (Optional) Disable Vega Upgrade Mode (`Bool`).(Deprecated)

`enable_vega_upgrade_mode` - (Optional) When enabled, vega will inform RE to stop traffic to the specific node. (`Bool`).(Deprecated)

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

x-displayName: "Enabled".

### Monitoring Choice Monitor Disabled

x-displayName: "Disabled".

### Network Choice Site Local Inside Network

x-displayName: "Site Local Inside (Local VRF)".

### Network Choice Site Local Network

x-displayName: "Site Local Outside (Local VRF)".

### Network Policy Choice Active Enhanced Firewall Policies

Enable Network Firewall for this site. Traffic will be processed in the order that Network Firewall Policies are added..

`enhanced_firewall_policies` - (Required) Ordered List of Enhanced Firewall Policies active. See [ref](#ref) below for details.

### Network Prefix Choice Cluster Static Ip

Static IP configuration for a specific node.

`interface_ip_map` - (Optional) Map of Node to Static ip configuration value, Key:Node, Value:IP Address. See [Cluster Static Ip Interface Ip Map ](#cluster-static-ip-interface-ip-map) below for details.

### Network Prefix Choice Fleet Static Ip

Static IP configuration for the fleet.

`default_gw` - (Optional) IP address offset of the default gateway, prefix len is used to calculate offset (`String`).

`dns_server` - (Optional) IP address offset of the DNS server, prefix len is used to calculate offset (`String`).

`network_prefix_allocator` - (Optional) Static IP configuration for the fleet. See [ref](#ref) below for details.

### Network Prefix Choice Node Static Ip

Static IP configuration for the Node.

`default_gw` - (Optional) IP address of the default gateway. (`String`).

`dns_server` - (Optional) IP address of the DNS server (`String`).(Deprecated)

`ip_address` - (Required) IP address of the interface and prefix length (`String`).

### Next Hop Choice Default Gateway

Traffic matching the ip prefixes is sent to the default gateway.

### Node List Interface List

Manage interfaces belonging to this node.

###### One of the arguments from this list "dhcp_client, dhcp_server, no_ipv4_address, static_ip" must be set

`dhcp_client` - (Optional) Interface gets it's IP address from an external DHCP server. (`Bool`).

`dhcp_server` - (Optional) DHCP Server is configured for this interface, Interface IP is derived from DHCP server configuration.. See [Address Choice Dhcp Server ](#address-choice-dhcp-server) below for details.

`no_ipv4_address` - (Optional) Interface does not have an IPv4 Address. (`Bool`).

`static_ip` - (Optional) Interface IP address is configured statically.. See [Address Choice Static Ip ](#address-choice-static-ip) below for details.

`description` - (Optional) Description for this Interface (`String`).

###### One of the arguments from this list "bond_interface, ethernet_interface, vlan_interface" must be set

`bond_interface` - (Optional) x-displayName: "Bond Interface". See [Interface Choice Bond Interface ](#interface-choice-bond-interface) below for details.

`ethernet_interface` - (Optional) x-displayName: "Ethernet Interface". See [Interface Choice Ethernet Interface ](#interface-choice-ethernet-interface) below for details.

`vlan_interface` - (Optional) x-displayName: "VLAN Interface". See [Interface Choice Vlan Interface ](#interface-choice-vlan-interface) below for details.

###### One of the arguments from this list "ipv6_auto_config, no_ipv6_address, static_ipv6_address" can be set

`ipv6_auto_config` - (Optional) Interface IPv6 address will be configured via Auto Configuration.. See [Ipv6 Address Choice Ipv6 Auto Config ](#ipv6-address-choice-ipv6-auto-config) below for details.

`no_ipv6_address` - (Optional) Interface does not have an IPv6 Address. (`Bool`).

`static_ipv6_address` - (Optional) Interface IPv6 address is configured statically.. See [Ipv6 Address Choice Static Ipv6 Address ](#ipv6-address-choice-static-ipv6-address) below for details.

`is_management` - (Optional) To be used internally to set an interface as management interface (`Bool`).(Deprecated)

`is_primary` - (Optional) Use for Primary Interface (`Bool`).(Deprecated)

`labels` - (Optional) Add Labels for this Interface, these labels can be used in firewall policy (`String`).

###### One of the arguments from this list "monitor, monitor_disabled" can be set

`monitor` - (Optional) x-displayName: "Enabled". See [Monitoring Choice Monitor ](#monitoring-choice-monitor) below for details.

`monitor_disabled` - (Optional) x-displayName: "Disabled" (`Bool`).

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

`name` - (Optional) Name of this Interface (`String`).

`network_option` - (Required) Global VRFs are configured via Networking > Segments. A site can have multple Network Segments (global VRFs).. See [Interface List Network Option ](#interface-list-network-option) below for details.

`priority` - (Optional) Greater the value, higher the priority (`Int`).

###### One of the arguments from this list "site_to_site_connectivity_interface_disabled, site_to_site_connectivity_interface_enabled" can be set

`site_to_site_connectivity_interface_disabled` - (Optional) Do not use this interface for site to site connectivity. (`Bool`).

`site_to_site_connectivity_interface_enabled` - (Optional) Use this this interface for site to site connectivity. (`Bool`).

### Not Managed Node List

Once a node is created and registers with the site, it will be shown in this section..

`hostname` - (Optional) Hostname for this Node (`String`).

`interface_list` - (Optional) Manage interfaces belonging to this node. See [Node List Interface List ](#node-list-interface-list) below for details.

`public_ip` - (Optional) Public IP for this Node (`String`).

`type` - (Optional) Type for this Node, can be Control or Worker (`String`).

### Offline Survivability Mode Choice Enable Offline Survivability Mode

x-displayName: "Enabled".

### Offline Survivability Mode Choice No Offline Survivability Mode

x-displayName: "Disabled".

### Operating System Version Choice Default Os Version

Will assign latest available OS version.

### Orchestration Choice Not Managed

or by using automation tools such as Terraform..

`node_list` - (Optional) Once a node is created and registers with the site, it will be shown in this section.. See [Not Managed Node List ](#not-managed-node-list) below for details.

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

### Provider Choice Aws

x-displayName: "AWS".

###### One of the arguments from this list "not_managed" can be set

`not_managed` - (Optional) or by using automation tools such as Terraform.. See [Orchestration Choice Not Managed ](#orchestration-choice-not-managed) below for details.

### Provider Choice Azure

x-displayName: "Azure".

###### One of the arguments from this list "not_managed" can be set

`not_managed` - (Optional) or by using automation tools such as Terraform.. See [Orchestration Choice Not Managed ](#orchestration-choice-not-managed) below for details.

### Provider Choice Baremetal

x-displayName: "Baremetal".

###### One of the arguments from this list "not_managed" can be set

`not_managed` - (Optional) or by using automation tools such as Terraform.. See [Orchestration Choice Not Managed ](#orchestration-choice-not-managed) below for details.

### Provider Choice Gcp

x-displayName: "GCP".

###### One of the arguments from this list "not_managed" can be set

`not_managed` - (Optional) or by using automation tools such as Terraform.. See [Orchestration Choice Not Managed ](#orchestration-choice-not-managed) below for details.

### Provider Choice Kvm

x-displayName: "KVM".

###### One of the arguments from this list "not_managed" can be set

`not_managed` - (Optional) or by using automation tools such as Terraform.. See [Orchestration Choice Not Managed ](#orchestration-choice-not-managed) below for details.

### Provider Choice Oci

x-displayName: "OCI".

###### One of the arguments from this list "not_managed" can be set

`not_managed` - (Optional) or by using automation tools such as Terraform.. See [Orchestration Choice Not Managed ](#orchestration-choice-not-managed) below for details.

### Provider Choice Rseries

x-displayName: "F5 rSeries".

###### One of the arguments from this list "not_managed" can be set

`not_managed` - (Optional) or by using automation tools such as Terraform.. See [Orchestration Choice Not Managed ](#orchestration-choice-not-managed) below for details.

### Provider Choice Vmware

x-displayName: "VMWare".

###### One of the arguments from this list "not_managed" can be set

`not_managed` - (Optional) or by using automation tools such as Terraform.. See [Orchestration Choice Not Managed ](#orchestration-choice-not-managed) below for details.

### Re Selection Choice Geo Proximity

Select REs in closest proximity to the site based on the public IP address of the control nodes of the site..

### Re Selection Choice Specific Re

Select specific REs. This is useful when a site needs to deterministically connect to a set of REs. A site will always be connected to 2 REs..

`backup_re` - (Optional) Select backup RE for this site. (`String`).(Deprecated)

`primary_re` - (Optional) Select primary RE for this site. (`String`).

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

### S2s Connectivity Slo Choice Site Mesh Group On Slo

Use a Site Mesh Group to connect to other sites..

###### One of the arguments from this list "no_site_mesh_group, site_mesh_group" must be set

`no_site_mesh_group` - (Optional) This site is not a member of Site Mesh group (`Bool`).(Deprecated)

`site_mesh_group` - (Optional) This site is member of Site Mesh Group via network. See [ref](#ref) below for details.(Deprecated)

###### One of the arguments from this list "sm_connection_public_ip, sm_connection_pvt_ip" must be set

`sm_connection_public_ip` - (Optional) tunnels to other sites which are part of the site mesh group. (`Bool`).

`sm_connection_pvt_ip` - (Optional) If multiple interfaces on a control node belong to the Site Local Outside (SLO) Local VRF, the interface which has 'Use for Site to Site Connectivity' set will be used. (`Bool`).

### Site Mesh Group Choice No Site Mesh Group

This site is not a member of Site Mesh group.

### Site Mesh Group Ip Choice Sm Connection Public Ip

tunnels to other sites which are part of the site mesh group..

### Site Mesh Group Ip Choice Sm Connection Pvt Ip

If multiple interfaces on a control node belong to the Site Local Outside (SLO) Local VRF, the interface which has 'Use for Site to Site Connectivity' set will be used..

### Site To Site Connectivity Interface Choice Site To Site Connectivity Interface Disabled

Do not use this interface for site to site connectivity..

### Site To Site Connectivity Interface Choice Site To Site Connectivity Interface Enabled

Use this this interface for site to site connectivity..

### Sli Choice Default Sli Config

x-displayName: "Default Configuration".

### Sli Choice Sli Config

Configure properties such as static routes, DNS and common VIP for Load Balancing on the Site Local Inside (SLI) local VRF..

`labels` - (Optional) Add Labels for this network, these labels can be used in firewall policy (`String`).

`nameserver` - (Optional) Optional DNS V4 server IP to be used for name resolution (`String`).

`nameserver_v6` - (Optional) Optional DNS V6 server IP to be used for name resolution (`String`).

###### One of the arguments from this list "no_static_routes, static_routes" must be set

`no_static_routes` - (Optional) Static IPv4 Routes disabled for this site local network (VRF). (`Bool`).

`static_routes` - (Optional) Manage IPv4 static routes for this site local network (VRF).. See [Static Route Choice Static Routes ](#static-route-choice-static-routes) below for details.

###### One of the arguments from this list "no_v6_static_routes, static_v6_routes" must be set

`no_v6_static_routes` - (Optional) Static IPv6 Routes disabled for this site local network (VRF). (`Bool`).

`static_v6_routes` - (Optional) Manage IPv6 static routes for this site local network (VRF).. See [Static V6 Route Choice Static V6 Routes ](#static-v6-route-choice-static-v6-routes) below for details.

`vip` - (Optional) Optional common virtual V4 IP across all nodes to be used as automatic VIP. (`String`).

`vip_v6` - (Optional) Optional common virtual V6 IP across all nodes to be used as automatic VIP. (`String`).

### Slo Choice Default Config

x-displayName: "Default Configuration".

### Slo Choice Slo Config

Configure properties such as static routes, DNS and common VIP for Load Balancing on the Site Local Outside (SLO) local VRF..

`labels` - (Optional) Add Labels for this network, these labels can be used in firewall policy (`String`).

`nameserver` - (Optional) Optional DNS V4 server IP to be used for name resolution (`String`).

`nameserver_v6` - (Optional) Optional DNS V6 server IP to be used for name resolution (`String`).

###### One of the arguments from this list "no_static_routes, static_routes" must be set

`no_static_routes` - (Optional) Static IPv4 Routes disabled for this site local network (VRF). (`Bool`).

`static_routes` - (Optional) Manage IPv4 static routes for this site local network (VRF).. See [Static Route Choice Static Routes ](#static-route-choice-static-routes) below for details.

###### One of the arguments from this list "no_v6_static_routes, static_v6_routes" must be set

`no_v6_static_routes` - (Optional) Static IPv6 Routes disabled for this site local network (VRF). (`Bool`).

`static_v6_routes` - (Optional) Manage IPv6 static routes for this site local network (VRF).. See [Static V6 Route Choice Static V6 Routes ](#static-v6-route-choice-static-v6-routes) below for details.

`vip` - (Optional) Optional common virtual V4 IP across all nodes to be used as automatic VIP. (`String`).

`vip_v6` - (Optional) Optional common virtual V6 IP across all nodes to be used as automatic VIP. (`String`).

### Software Settings Os

Select the Operating System version for the site. By default, latest available Operating System version will be used..

###### One of the arguments from this list "default_os_version, operating_system_version" must be set

`default_os_version` - (Optional) Will assign latest available OS version (`Bool`).

`operating_system_version` - (Optional) Specify a OS version to be used e.g. 9.2024.6. (`String`).

### Software Settings Sw

Refer to release notes to find required released SW versions..

###### One of the arguments from this list "default_sw_version, volterra_software_version" must be set

`default_sw_version` - (Optional) Will assign latest available F5XC Software Version (`Bool`).

`volterra_software_version` - (Optional) Specify a F5XC Software Version to be used e.g. crt-20210329-1002. (`String`).

### Stateful Dhcp Networks

List of networks from which DHCP server can allocate ip addresses.

###### One of the arguments from this list "network_prefix, network_prefix_allocator" must be set

`network_prefix` - (Optional) Network Prefix to be used for IPV6 address auto configuration (`String`).

`network_prefix_allocator` - (Optional) Prefix length from address allocator scheme is used to calculate offsets. See [ref](#ref) below for details.(Deprecated)

`pool_settings` - (Required) Controls how DHCPV6 pools are handled (`String`).

`pools` - (Optional) List of non overlapping ip address ranges.. See [Dhcp Networks Pools ](#dhcp-networks-pools) below for details.

### Static Route Choice No Static Routes

Static IPv4 Routes disabled for this site local network (VRF)..

### Static Route Choice Static Routes

Manage IPv4 static routes for this site local network (VRF)..

`static_routes` - (Required) x-required. See [Static Routes Static Routes ](#static-routes-static-routes) below for details.

### Static Routes Static Routes

x-required.

`attrs` - (Optional) List of attributes that control forwarding, dynamic routing and control plane (host) reachability (`List of Strings`).

`ip_prefixes` - (Required) List of route prefixes that have common next hop and attributes (`String`).

###### One of the arguments from this list "default_gateway, interface, ip_address" must be set

`default_gateway` - (Optional) Traffic matching the ip prefixes is sent to the default gateway (`Bool`).

`interface` - (Optional) Traffic matching the ip prefixes is sent to this interface. See [ref](#ref) below for details.

`ip_address` - (Optional) Traffic matching the ip prefixes is sent to this IP Address (`String`).

### Static V6 Route Choice No V6 Static Routes

Static IPv6 Routes disabled for this site local network (VRF)..

### Static V6 Route Choice Static V6 Routes

Manage IPv6 static routes for this site local network (VRF)..

`static_routes` - (Required) List of IPv6 static routes. See [Static V6 Routes Static Routes ](#static-v6-routes-static-routes) below for details.

### Static V6 Routes Static Routes

List of IPv6 static routes.

`attrs` - (Optional) List of attributes that control forwarding, dynamic routing and control plane (host) reachability (`List of Strings`).

`ip_prefixes` - (Required) List of IPv6 route prefixes that have common next hop and attributes (`String`).

###### One of the arguments from this list "default_gateway, interface, ip_address" must be set

`default_gateway` - (Optional) Traffic matching the ip prefixes is sent to the default gateway (`Bool`).

`interface` - (Optional) Traffic matching the ip prefixes is sent to this interface. See [ref](#ref) below for details.

`ip_address` - (Optional) Traffic matching the ip prefixes is sent to this IP Address (`String`).

### Upgrade Settings Kubernetes Upgrade Drain

Specify how worker nodes within a site will be upgraded..

###### One of the arguments from this list "disable_upgrade_drain, enable_upgrade_drain" must be set

`disable_upgrade_drain` - (Optional) x-displayName: "Disable Node by Node Upgrade" (`Bool`).

`enable_upgrade_drain` - (Optional) x-displayName: "Enable Node by Node Upgrade". See [Kubernetes Upgrade Drain Enable Choice Enable Upgrade Drain ](#kubernetes-upgrade-drain-enable-choice-enable-upgrade-drain) below for details.

### Vega Upgrade Mode Toggle Choice Disable Vega Upgrade Mode

Disable Vega Upgrade Mode.

### Vega Upgrade Mode Toggle Choice Enable Vega Upgrade Mode

When enabled, vega will inform RE to stop traffic to the specific node..

### Volterra Sw Version Choice Default Sw Version

Will assign latest available F5XC Software Version.

Attribute Reference
-------------------

-	`id` - This is the id of the configured securemesh_site_v2.
