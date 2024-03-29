---

page_title: "Volterra: securemesh_site"

description: "The securemesh_site allows CRUD of Securemesh Site resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_securemesh_site
=================================

The Securemesh Site allows CRUD of Securemesh Site resource on Volterra SaaS

~> **Note:** Please refer to [Securemesh Site API docs](https://docs.cloud.f5.com/docs/api/views-securemesh-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_securemesh_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "default_blocked_services blocked_services" must be set
  default_blocked_services = true

  // One of the arguments from this list "no_bond_devices bond_device_list" must be set
  no_bond_devices = true

  // One of the arguments from this list "logs_streaming_disabled log_receiver" must be set
  logs_streaming_disabled = true

  master_node_configuration {
    name = "master-0"

    public_ip = "192.168.0.156"
  }

  // One of the arguments from this list "default_network_config custom_network_config" must be set

  custom_network_config {
    bgp_peer_address = "10.1.1.1"

    bgp_peer_address_v6 = "3c0f:7554:352a:a2dc:333f:67c5:c2b5:7326"

    bgp_router_id = "10.1.1.1"

    // One of the arguments from this list "forward_proxy_allow_all no_forward_proxy active_forward_proxy_policies" must be set
    no_forward_proxy = true

    // One of the arguments from this list "no_global_network global_network_list" must be set

    global_network_list {
      global_network_connections {
        // One of the arguments from this list "sli_to_global_dr slo_to_global_dr" must be set

        sli_to_global_dr {
          global_vn {
            name      = "test1"
            namespace = "staging"
            tenant    = "acmecorp"
          }
        }

        // One of the arguments from this list "disable_forward_proxy enable_forward_proxy" must be set
        disable_forward_proxy = true
      }
    }
    // One of the arguments from this list "default_interface_config interface_list" must be set
    default_interface_config = true
    // One of the arguments from this list "active_network_policies active_enhanced_firewall_policies no_network_policy" must be set
    no_network_policy = true
    // One of the arguments from this list "sm_connection_public_ip sm_connection_pvt_ip" must be set
    sm_connection_public_ip = true
    // One of the arguments from this list "default_sli_config sli_config" must be set
    default_sli_config = true
    // One of the arguments from this list "default_config slo_config" must be set
    default_config = true
    tunnel_dead_timeout = "0"
    vip_vrrp_mode = "vip_vrrp_mode"
  }
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

`blocked_services` - (Optional) Use custom blocked services configuration. See [Blocked Services Choice Blocked Services ](#blocked-services-choice-blocked-services) below for details.

`default_blocked_services` - (Optional) Use default behavior of allowing ports mentioned in blocked services (`Bool`).

`bond_device_list` - (Optional) Configure Bond Devices for this Secure Mesh site. See [Bond Choice Bond Device List ](#bond-choice-bond-device-list) below for details.

`no_bond_devices` - (Optional) No Bond Devices configured for this Secure Mesh site (`Bool`).

`coordinates` - (Optional) Coordinates of the site, longitude and latitude. See [Coordinates ](#coordinates) below for details.

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (`Bool`).

`master_node_configuration` - (Required) Configuration of master nodes. See [Master Node Configuration ](#master-node-configuration) below for details.

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

### Master Node Configuration

Configuration of master nodes.

`name` - (Required) Names of master node (`String`).

`public_ip` - (Optional) via Site Mesh Group (`String`).

### Offline Survivability Mode

Enable/Disable offline survivability mode.

###### One of the arguments from this list "no_offline_survivability_mode, enable_offline_survivability_mode" must be set

`enable_offline_survivability_mode` - (Optional) When this feature is enabled on an existing site, the pods/services on this site will be restarted. (`Bool`).

`no_offline_survivability_mode` - (Optional) When this feature is disabled on an existing site, the pods/services on this site will be restarted. (`Bool`).

### Os

Operating System Details.

###### One of the arguments from this list "default_os_version, operating_system_version" must be set

`default_os_version` - (Optional) Will assign latest available OS version (`Bool`).

`operating_system_version` - (Optional) Operating System Version is optional parameter, which allows to specify target OS version for particular site e.g. 7.2009.10. (`String`).

### Performance Enhancement Mode

Performance Enhancement Mode to optimize for L3 or L7 networking.

###### One of the arguments from this list "perf_mode_l3_enhanced, perf_mode_l7_enhanced" must be set

`perf_mode_l3_enhanced` - (Optional) When the mode is toggled to l3 enhanced, traffic disruption will be seen. See [Perf Mode Choice Perf Mode L3 Enhanced ](#perf-mode-choice-perf-mode-l3-enhanced) below for details.

`perf_mode_l7_enhanced` - (Optional) When the mode is toggled to l7 enhanced, traffic disruption will be seen (`Bool`).

### Sw

F5XC Software Details.

###### One of the arguments from this list "default_sw_version, volterra_software_version" must be set

`default_sw_version` - (Optional) Will assign latest available SW version (`Bool`).

`volterra_software_version` - (Optional) F5XC Software Version is optional parameter, which allows to specify target SW version for particular site e.g. crt-20210329-1002. (`String`).

### Address Choice Dhcp Client

Interface gets it IP address from external DHCP server.

### Address Choice Dhcp Server

DHCP Server is configured for this interface, Interface IP from DHCP server configuration..

`dhcp_networks` - (Required) List of networks from which DHCP server can allocate ip addresses. See [Dhcp Server Dhcp Networks ](#dhcp-server-dhcp-networks) below for details.

`dhcp_option82_tag` - (Optional) Optional tag that can be given to this configuration (`String`).(Deprecated)

`fixed_ip_map` - (Optional) Fixed MAC address to ip assignments, Key: Mac address, Value: IP Address (`String`).

###### One of the arguments from this list "interface_ip_map, automatic_from_start, automatic_from_end" must be set

`automatic_from_end` - (Optional) Assign automatically from End of the first network in the list (`Bool`).

`automatic_from_start` - (Optional) Assign automatically from start of the first network in the list (`Bool`).

`interface_ip_map` - (Optional) Configured address for every node. See [Interfaces Addressing Choice Interface Ip Map ](#interfaces-addressing-choice-interface-ip-map) below for details.

### Address Choice Stateful

works along with Router Advertisement' Managed flag.

`dhcp_networks` - (Required) List of networks from which DHCP server can allocate ip addresses. See [Stateful Dhcp Networks ](#stateful-dhcp-networks) below for details.

`fixed_ip_map` - (Optional) Fixed MAC address to ipv6 assignments, Key: Mac address, Value: IPV6 Address (`String`).

###### One of the arguments from this list "automatic_from_start, automatic_from_end, interface_ip_map" must be set

`automatic_from_end` - (Optional) Assign automatically from End of the first network in the list (`Bool`).

`automatic_from_start` - (Optional) Assign automatically from start of the first network in the list (`Bool`).

`interface_ip_map` - (Optional) Configured address for every node. See [Interfaces Addressing Choice Interface Ip Map ](#interfaces-addressing-choice-interface-ip-map) below for details.

### Address Choice Static Ip

Interface IP is configured statically.

###### One of the arguments from this list "node_static_ip, cluster_static_ip, fleet_static_ip" must be set

`cluster_static_ip` - (Optional) Static IP configuration for a specific node. See [Network Prefix Choice Cluster Static Ip ](#network-prefix-choice-cluster-static-ip) below for details.

`fleet_static_ip` - (Optional) Static IP configuration for the fleet. See [Network Prefix Choice Fleet Static Ip ](#network-prefix-choice-fleet-static-ip) below for details.(Deprecated)

`node_static_ip` - (Optional) Static IP configuration for the Node. See [Network Prefix Choice Node Static Ip ](#network-prefix-choice-node-static-ip) below for details.

### Autoconfig Choice Host

auto configuration routers.

### Autoconfig Choice Router

System behaves like Auto config Router and provides auto config parameters.

###### One of the arguments from this list "network_prefix, stateful" must be set

`network_prefix` - (Optional) Nework prefix that is used as Prefix information (`String`).

`stateful` - (Optional) works along with Router Advertisement' Managed flag. See [Address Choice Stateful ](#address-choice-stateful) below for details.

`dns_config` - (Optional) Dns information that needs to added in the RouterAdvetisement. See [Router Dns Config ](#router-dns-config) below for details.

### Blocked Services Blocked Sevice

Use custom blocked services configuration.

###### One of the arguments from this list "ssh, web_user_interface, dns" can be set

`dns` - (Optional) Matches DNS port 53 (`Bool`).

`ssh` - (Optional) Matches ssh port 22 (`Bool`).

`web_user_interface` - (Optional) Matches the web user interface port (`Bool`).

`network_type` - (Optional) Network type in which these ports get blocked (`String`).

### Blocked Services Choice Blocked Services

Use custom blocked services configuration.

`blocked_sevice` - (Optional) Use custom blocked services configuration. See [Blocked Services Blocked Sevice ](#blocked-services-blocked-sevice) below for details.

### Blocked Services Value Type Choice Dns

Matches DNS port 53.

### Blocked Services Value Type Choice Ssh

Matches ssh port 22.

### Blocked Services Value Type Choice Web User Interface

Matches the web user interface port.

### Bond Choice Bond Device List

Configure Bond Devices for this Secure Mesh site.

`bond_devices` - (Required) List of bond devices. See [Bond Device List Bond Devices ](#bond-device-list-bond-devices) below for details.

### Bond Device List Bond Devices

List of bond devices.

`devices` - (Required) Ethernet devices that will make up this bond (`String`).

###### One of the arguments from this list "lacp, active_backup" must be set

`active_backup` - (Optional) Configure active/backup based bond device (`Bool`).

`lacp` - (Optional) Configure LACP (802.3ad) based bond device. See [Lacp Choice Lacp ](#lacp-choice-lacp) below for details.

`link_polling_interval` - (Required) Link polling interval in millisecond (`Int`).

`link_up_delay` - (Required) Milliseconds wait before link is declared up (`Int`).

`name` - (Required) Bond device name (`String`).

### Cluster Static Ip Interface Ip Map

Map of Node to Static ip configuration value, Key:Node, Value:IP Address.

`default_gw` - (Optional) IP address of the default gateway. (`String`).

`dns_server` - (Optional) IP address of the DNS server (`String`).(Deprecated)

`ip_address` - (Required) IP address of the interface and prefix length (`String`).

### Connection Choice Sli To Global Dr

Site local inside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Connection Choice Slo To Global Dr

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Custom Certificate Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Dc Cluster Group Choice No Dc Cluster Group

This site is not a member of dc cluster group.

### Dc Cluster Group Connectivity Interface Choice Dc Cluster Group Connectivity Interface Disabled

Do not use this interface to connect to DC cluster group peers. .

### Dc Cluster Group Connectivity Interface Choice Dc Cluster Group Connectivity Interface Enabled

Use this interface to connect to DC cluster group peers..

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

List of networks from which DHCP server can allocate ip addresses.

###### One of the arguments from this list "same_as_dgw, dns_address" must be set

`dns_address` - (Optional) Configured address is chosen as DNS server address in DHCP response. (`String`).

`same_as_dgw` - (Optional) DNS server address is same as default gateway address (`Bool`).

###### One of the arguments from this list "last_address, dgw_address, first_address" must be set

`dgw_address` - (Optional) Configured address from the network prefix is chosen as default gateway. (`String`).

`first_address` - (Optional) First usable address from the network prefix is chosen as default gateway (`Bool`).

`last_address` - (Optional) Last usable address from the network prefix is chosen as default gateway (`Bool`).

###### One of the arguments from this list "network_prefix, network_prefix_allocator" must be set

`network_prefix` - (Optional) Network Prefix for a single site. (`String`).

`network_prefix_allocator` - (Optional) Prefix length from address allocator scheme is used to calculate offsets. See [ref](#ref) below for details.(Deprecated)

`pool_settings` - (Required) Controls how DHCP pools are handled (`String`).

`pools` - (Optional) List of non overlapping ip address ranges.. See [Dhcp Networks Pools ](#dhcp-networks-pools) below for details.

### Dns Choice Configured List

Configured address outside network range - external dns server.

`dns_list` - (Required) List of IPV6 Addresses acting as Dns servers (`String`).

### Dns Choice Local Dns

Choose the address from the network prefix range as dns server.

###### One of the arguments from this list "first_address, last_address, configured_address" must be set

`configured_address` - (Optional) Configured address from the network prefix is chosen as dns server (`String`).

`first_address` - (Optional) First usable address from the network prefix is chosen as dns server (`Bool`).

`last_address` - (Optional) Last usable address from the network prefix is chosen as dns server (`Bool`).

### Dns Choice Same As Dgw

DNS server address is same as default gateway address.

### Enable Disable Choice Disable Interception

Disable Interception.

### Enable Disable Choice Enable Interception

Enable Interception.

### Forward Proxy Choice Active Forward Proxy Policies

Enable Forward Proxy for this site and manage policies.

`forward_proxy_policies` - (Required) List of Forward Proxy Policies. See [ref](#ref) below for details.

### Forward Proxy Choice Disable Forward Proxy

Forward Proxy is disabled for this connector.

### Forward Proxy Choice Enable Forward Proxy

Forward Proxy is enabled for this connector.

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2000 (2 seconds) (`Int`).

`max_connect_attempts` - (Optional) Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1. (`Int`).

###### One of the arguments from this list "no_interception, tls_intercept" can be set

`no_interception` - (Optional) No TLS interception is enabled for this network connector (`Bool`).(Deprecated)

`tls_intercept` - (Optional) Specify TLS interception configuration for the network connector. See [Tls Interception Choice Tls Intercept ](#tls-interception-choice-tls-intercept) below for details.(Deprecated)

`white_listed_ports` - (Optional) Example "tmate" server port (`Int`).

`white_listed_prefixes` - (Optional) Example "tmate" server ip (`String`).

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

###### One of the arguments from this list "disable_forward_proxy, enable_forward_proxy" can be set

`disable_forward_proxy` - (Optional) Forward Proxy is disabled for this connector (`Bool`).(Deprecated)

`enable_forward_proxy` - (Optional) Forward Proxy is enabled for this connector. See [Forward Proxy Choice Enable Forward Proxy ](#forward-proxy-choice-enable-forward-proxy) below for details.(Deprecated)

### Interception Policy Choice Enable For All Domains

Enable interception for all domains.

### Interception Policy Choice Policy

Policy to enable/disable specific domains, with implicit enable all domains.

`interception_rules` - (Required) List of ordered rules to enable or disable for TLS interception. See [Policy Interception Rules ](#policy-interception-rules) below for details.

### Interception Rules Domain Match

Domain value or regular expression to match.

###### One of the arguments from this list "exact_value, suffix_value, regex_value" must be set

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

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

###### One of the arguments from this list "not_primary, is_primary" must be set

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

`dhcp_client` - (Optional) Interface gets it IP address from external DHCP server (`Bool`).

`dhcp_server` - (Optional) DHCP Server is configured for this interface, Interface IP from DHCP server configuration.. See [Address Choice Dhcp Server ](#address-choice-dhcp-server) below for details.

`static_ip` - (Optional) Interface IP is configured statically. See [Address Choice Static Ip ](#address-choice-static-ip) below for details.

`device` - (Required) Interface configuration for the ethernet device (`String`).

###### One of the arguments from this list "ipv6_auto_config, no_ipv6_address, static_ipv6_address" can be set

`ipv6_auto_config` - (Optional) Configuration corresponding to IPV6 auto configuration. See [Ipv6 Address Choice Ipv6 Auto Config ](#ipv6-address-choice-ipv6-auto-config) below for details.

`no_ipv6_address` - (Optional) Interface does not have an IPv6 Address. (`Bool`).

`static_ipv6_address` - (Optional) Interface IP is configured statically. See [Ipv6 Address Choice Static Ipv6 Address ](#ipv6-address-choice-static-ipv6-address) below for details.

###### One of the arguments from this list "monitor_disabled, monitor" can be set

`monitor` - (Optional) Link Quality Monitoring parameters. Choosing the option will enable link quality monitoring.. See [Monitoring Choice Monitor ](#monitoring-choice-monitor) below for details.

`monitor_disabled` - (Optional) Link quality monitoring disabled on the interface. (`Bool`).

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

###### One of the arguments from this list "site_local_network, site_local_inside_network, inside_network, storage_network, srv6_network, ip_fabric_network, segment_network" must be set

`inside_network` - (Optional) Interface belongs to user configured inside network. See [ref](#ref) below for details.(Deprecated)

`ip_fabric_network` - (Optional) Interface belongs to IP Fabric network (`Bool`).(Deprecated)

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (`Bool`).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (`Bool`).

`srv6_network` - (Optional) Interface belongs to per site srv6 network. See [ref](#ref) below for details.(Deprecated)

`storage_network` - (Optional) Interface belongs to site local network inside (`Bool`).

###### One of the arguments from this list "cluster, node" must be set

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site. (`Bool`).

`node` - (Optional) Configuration will apply to a device on the given node. (`String`).

###### One of the arguments from this list "not_primary, is_primary" must be set

`is_primary` - (Optional) This interface is primary (`Bool`).

`not_primary` - (Optional) This interface is not primary (`Bool`).

`priority` - (Optional) Greater the value, higher the priority (`Int`).

###### One of the arguments from this list "untagged, vlan_id" must be set

`untagged` - (Optional) Configure a untagged ethernet interface (`Bool`).

`vlan_id` - (Optional) Configure a VLAN tagged ethernet interface (`Int`).

### Interface Choice Interface List

Add all interfaces belonging to this site.

`interfaces` - (Required) Configure network interfaces for this Secure Mesh site. See [Interface List Interfaces ](#interface-list-interfaces) below for details.

### Interface Choice Loopback Interface

Loopback device..

###### One of the arguments from this list "dhcp_client, dhcp_server, static_ip" must be set

`dhcp_client` - (Optional) Interface gets it IP address from external DHCP server (`Bool`).

`dhcp_server` - (Optional) DHCP Server is configured for this interface, Interface IP from DHCP server configuration.. See [Address Choice Dhcp Server ](#address-choice-dhcp-server) below for details.

`static_ip` - (Optional) Interface IP is configured statically. See [Address Choice Static Ip ](#address-choice-static-ip) below for details.

`device` - (Required) Interface configuration for the Loopback Ethernet device (`String`).

###### One of the arguments from this list "no_ipv6_address, static_ipv6_address" can be set

`no_ipv6_address` - (Optional) Interface does not have an IPv6 Address. (`Bool`).

`static_ipv6_address` - (Optional) Interface IP is configured statically. See [Ipv6 Address Choice Static Ipv6 Address ](#ipv6-address-choice-static-ipv6-address) below for details.

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

###### One of the arguments from this list "site_local_network, site_local_inside_network, ip_fabric_network" must be set

`ip_fabric_network` - (Optional) Interface belongs to IP Fabric network (`Bool`).(Deprecated)

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (`Bool`).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (`Bool`).

###### One of the arguments from this list "cluster, node" must be set

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site. (`Bool`).

`node` - (Optional) Configuration will apply to a device on the given node. (`String`).

### Interface List Interfaces

Configure network interfaces for this Secure Mesh site.

###### One of the arguments from this list "dc_cluster_group_connectivity_interface_disabled, dc_cluster_group_connectivity_interface_enabled" must be set

`dc_cluster_group_connectivity_interface_disabled` - (Optional) Do not use this interface to connect to DC cluster group peers. (`Bool`).

`dc_cluster_group_connectivity_interface_enabled` - (Optional) Use this interface to connect to DC cluster group peers. (`Bool`).

`description` - (Optional) Description for this Interface (`String`).

###### One of the arguments from this list "ethernet_interface, dedicated_interface, dedicated_management_interface, loopback_interface" must be set

`dedicated_interface` - (Optional) Networking configuration for dedicated interface is configured locally on site e.g. (outside/inside)Ethernet. See [Interface Choice Dedicated Interface ](#interface-choice-dedicated-interface) below for details.

`dedicated_management_interface` - (Optional) Fallback management interfaces can be made into dedicated management interface. See [Interface Choice Dedicated Management Interface ](#interface-choice-dedicated-management-interface) below for details.

`ethernet_interface` - (Optional) Ethernet interface configuration.. See [Interface Choice Ethernet Interface ](#interface-choice-ethernet-interface) below for details.

`loopback_interface` - (Optional) Loopback device.. See [Interface Choice Loopback Interface ](#interface-choice-loopback-interface) below for details.(Deprecated)

`labels` - (Optional) Add Labels for this Interface, these labels can be used in firewall policy (`String`).

### Interfaces Addressing Choice Automatic From End

Assign automatically from End of the first network in the list.

### Interfaces Addressing Choice Automatic From Start

Assign automatically from start of the first network in the list.

### Interfaces Addressing Choice Interface Ip Map

Configured address for every node.

`interface_ip_map` - (Optional) Map of Site:Node to IP address. (`String`).

### Interfaces Addressing Choice Interface Ip Map

Configured address for every node.

`interface_ip_map` - (Optional) Map of Site:Node to IPV6 address. (`String`).

### Ipv6 Address Choice Ipv6 Auto Config

Configuration corresponding to IPV6 auto configuration.

###### One of the arguments from this list "router, host" must be set

`host` - (Optional) auto configuration routers (`Bool`).

`router` - (Optional) System behaves like Auto config Router and provides auto config parameters. See [Autoconfig Choice Router ](#autoconfig-choice-router) below for details.

### Ipv6 Address Choice No Ipv6 Address

Interface does not have an IPv6 Address..

### Ipv6 Address Choice Static Ipv6 Address

Interface IP is configured statically.

###### One of the arguments from this list "node_static_ip, cluster_static_ip, fleet_static_ip" must be set

`cluster_static_ip` - (Optional) Static IP configuration for a specific node. See [Network Prefix Choice Cluster Static Ip ](#network-prefix-choice-cluster-static-ip) below for details.

`fleet_static_ip` - (Optional) Static IP configuration for the fleet. See [Network Prefix Choice Fleet Static Ip ](#network-prefix-choice-fleet-static-ip) below for details.(Deprecated)

`node_static_ip` - (Optional) Static IP configuration for the Node. See [Network Prefix Choice Node Static Ip ](#network-prefix-choice-node-static-ip) below for details.

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

`bgp_peer_address` - (Optional) to fetch BGP peer address from site Object. This can be used to change peer address per site in fleet. (`String`).(Deprecated)

`bgp_peer_address_v6` - (Optional) to fetch BGP peer IPv6 address from site Object. This can be used to change peer IPv6 address per site in fleet. (`String`).(Deprecated)

`bgp_router_id` - (Optional) fetch BGP router ID from site object. (`String`).(Deprecated)

###### One of the arguments from this list "no_forward_proxy, active_forward_proxy_policies, forward_proxy_allow_all" must be set

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Forward Proxy Choice Active Forward Proxy Policies ](#forward-proxy-choice-active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (`Bool`).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (`Bool`).

###### One of the arguments from this list "no_global_network, global_network_list" must be set

`global_network_list` - (Optional) List of global network connections. See [Global Network Choice Global Network List ](#global-network-choice-global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (`Bool`).

###### One of the arguments from this list "default_interface_config, interface_list" must be set

`default_interface_config` - (Optional) Interface configuration is done based on certified hardware for this site (`Bool`).

`interface_list` - (Optional) Add all interfaces belonging to this site. See [Interface Choice Interface List ](#interface-choice-interface-list) below for details.

###### One of the arguments from this list "active_enhanced_firewall_policies, no_network_policy, active_network_policies" must be set

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

### Network Choice Ip Fabric Network

Interface belongs to IP Fabric network.

### Network Choice Site Local Inside Network

Interface belongs to site local network inside.

### Network Choice Site Local Network

Interface belongs to site local network (outside).

### Network Choice Storage Network

Interface belongs to site local network inside.

### Network Policy Choice Active Enhanced Firewall Policies

with an additional option for service insertion..

`enhanced_firewall_policies` - (Required) Ordered List of Enhaned Firewall Policy active for this network firewall. See [ref](#ref) below for details.

### Network Policy Choice Active Network Policies

Firewall Policies active for this site..

`network_policies` - (Required) Ordered List of Firewall Policies active for this network firewall. See [ref](#ref) below for details.

### Network Policy Choice No Network Policy

Firewall Policy is disabled for this site..

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

Traffic matching the ip prefixes is sent to default gateway.

### Node Choice Cluster

Configuration will apply to given device on all nodes of the site..

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ocsp Stapling Choice Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Ocsp Stapling Choice Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Offline Survivability Mode Choice Enable Offline Survivability Mode

When this feature is enabled on an existing site, the pods/services on this site will be restarted..

### Offline Survivability Mode Choice No Offline Survivability Mode

When this feature is disabled on an existing site, the pods/services on this site will be restarted..

### Operating System Version Choice Default Os Version

Will assign latest available OS version.

### Perf Mode Choice Jumbo

L3 performance mode enhancement to use jumbo frame.

### Perf Mode Choice No Jumbo

L3 performance mode enhancement without jumbo frame.

### Perf Mode Choice Perf Mode L3 Enhanced

When the mode is toggled to l3 enhanced, traffic disruption will be seen.

###### One of the arguments from this list "no_jumbo, jumbo" must be set

`jumbo` - (Optional) L3 performance mode enhancement to use jumbo frame (`Bool`).

`no_jumbo` - (Optional) L3 performance mode enhancement without jumbo frame (`Bool`).

### Perf Mode Choice Perf Mode L7 Enhanced

When the mode is toggled to l7 enhanced, traffic disruption will be seen.

### Policy Interception Rules

List of ordered rules to enable or disable for TLS interception.

`domain_match` - (Required) Domain value or regular expression to match. See [Interception Rules Domain Match ](#interception-rules-domain-match) below for details.

###### One of the arguments from this list "disable_interception, enable_interception" must be set

`disable_interception` - (Optional) Disable Interception (`Bool`).

`enable_interception` - (Optional) Enable Interception (`Bool`).

### Primary Choice Is Primary

This interface is primary.

### Primary Choice Not Primary

This interface is not primary.

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

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

### Secret Info Oneof Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Info Oneof Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Secret Info Oneof Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Secret Info Oneof Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

### Signing Cert Choice Custom Certificate

Certificates for generating intermediate certificate for TLS interception..

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "use_system_defaults, disable_ocsp_stapling, custom_hash_algorithms" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Custom Certificate Private Key ](#custom-certificate-private-key) below for details.

### Signing Cert Choice Volterra Certificate

F5XC certificates for generating intermediate certificate for TLS interception..

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

`dc_cluster_group_interface` - (Optional) This Secure Mesh is member of dc cluster group and connected to network over this interface. By default it takes default gateway interface.. See [ref](#ref) below for details.(Deprecated)

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

###### One of the arguments from this list "no_dc_cluster_group, dc_cluster_group" must be set

`dc_cluster_group` - (Optional) This site is member of dc cluster group via network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (`Bool`).

`dc_cluster_group_interface` - (Optional) This Secure Mesh is member of dc cluster group and connected to network over this interface. By default it takes default gateway interface.. See [ref](#ref) below for details.(Deprecated)

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

###### One of the arguments from this list "network_prefix, network_prefix_allocator" must be set

`network_prefix` - (Optional) Network Prefix to be used for IPV6 address auto configuration (`String`).

`network_prefix_allocator` - (Optional) Prefix length from address allocator scheme is used to calculate offsets. See [ref](#ref) below for details.(Deprecated)

`pools` - (Optional) List of non overlapping ip address ranges.. See [Dhcp Networks Pools ](#dhcp-networks-pools) below for details.

### Static Route Choice No Static Routes

Static Routes disabled for site local network..

### Static Route Choice Static Routes

Manage static routes for site local network..

`static_routes` - (Required) List of static routes. See [Static Routes Static Routes ](#static-routes-static-routes) below for details.

### Static Routes Static Routes

List of static routes.

`attrs` - (Optional) List of attributes that control forwarding, dynamic routing and control plane(host) reachability (`List of Strings`).

`ip_prefixes` - (Required) List of route prefixes that have common next hop and attributes (`String`).

###### One of the arguments from this list "ip_address, interface, default_gateway" must be set

`default_gateway` - (Optional) Traffic matching the ip prefixes is sent to default gateway (`Bool`).

`interface` - (Optional) Traffic matching the ip prefixes is sent to the interface. See [ref](#ref) below for details.

`ip_address` - (Optional) Traffic matching the ip prefixes is sent to IP Address (`String`).

### Static V6 Route Choice No V6 Static Routes

Static IPv6 Routes disabled for site local network..

### Static V6 Route Choice Static V6 Routes

Manage IPv6 static routes for site local network..

`static_routes` - (Required) List of IPv6 static routes. See [Static V6 Routes Static Routes ](#static-v6-routes-static-routes) below for details.

### Static V6 Routes Static Routes

List of IPv6 static routes.

`attrs` - (Optional) List of attributes that control forwarding, dynamic routing and control plane(host) reachability (`List of Strings`).

`ip_prefixes` - (Required) List of IPv6 route prefixes that have common next hop and attributes (`String`).

###### One of the arguments from this list "ip_address, interface, default_gateway" must be set

`default_gateway` - (Optional) Traffic matching the ip prefixes is sent to default gateway (`Bool`).

`interface` - (Optional) Traffic matching the ip prefixes is sent to the interface. See [ref](#ref) below for details.

`ip_address` - (Optional) Traffic matching the ip prefixes is sent to IP Address (`String`).

### Tls Interception Choice No Interception

No TLS interception is enabled for this network connector.

### Tls Interception Choice Tls Intercept

Specify TLS interception configuration for the network connector.

###### One of the arguments from this list "enable_for_all_domains, policy" must be set

`enable_for_all_domains` - (Optional) Enable interception for all domains (`Bool`).

`policy` - (Optional) Policy to enable/disable specific domains, with implicit enable all domains. See [Interception Policy Choice Policy ](#interception-policy-choice-policy) below for details.

###### One of the arguments from this list "custom_certificate, volterra_certificate" must be set

`custom_certificate` - (Optional) Certificates for generating intermediate certificate for TLS interception.. See [Signing Cert Choice Custom Certificate ](#signing-cert-choice-custom-certificate) below for details.

`volterra_certificate` - (Optional) F5XC certificates for generating intermediate certificate for TLS interception. (`Bool`).

###### One of the arguments from this list "trusted_ca_url, volterra_trusted_ca" must be set

`trusted_ca_url` - (Optional) Custom Root CA Certificate for validating upstream server certificate (`String`).

`volterra_trusted_ca` - (Optional) F5XC Root CA Certificate for validating upstream server certificate (`Bool`).

### Trusted Ca Choice Volterra Trusted Ca

F5XC Root CA Certificate for validating upstream server certificate.

### Vlan Choice Untagged

Configure a untagged ethernet interface.

### Volterra Sw Version Choice Default Sw Version

Will assign latest available SW version.

Attribute Reference
-------------------

-	`id` - This is the id of the configured securemesh_site.
