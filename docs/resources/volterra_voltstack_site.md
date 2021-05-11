---

page_title: "Volterra: voltstack_site"

description: "The voltstack_site allows CRUD of Voltstack Site resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_voltstack_site
================================

The Voltstack Site allows CRUD of Voltstack Site resource on Volterra SaaS

~> **Note:** Please refer to [Voltstack Site API docs](https://volterra.io/docs/api/views-voltstack-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_voltstack_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "no_bond_devices bond_device_list" must be set
  no_bond_devices = true

  // One of the arguments from this list "disable_gpu enable_gpu" must be set
  disable_gpu = true

  // One of the arguments from this list "no_k8s_cluster k8s_cluster" must be set
  no_k8s_cluster = true

  // One of the arguments from this list "logs_streaming_disabled log_receiver" must be set
  logs_streaming_disabled = true

  master_nodes = ["master-0"]

  // One of the arguments from this list "custom_network_config default_network_config" must be set
  default_network_config = true

  // One of the arguments from this list "default_storage_config custom_storage_config" must be set
  default_storage_config = true

  // One of the arguments from this list "deny_all_usb allow_all_usb usb_policy" must be set
  deny_all_usb          = true
  volterra_certified_hw = ["isv-8000-series-voltstack-combo"]
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

`bond_device_list` - (Optional) Configure Bond Devices for this voltstack site. See [Bond Device List ](#bond-device-list) below for details.

`no_bond_devices` - (Optional) No Bond Devices configured for this voltstack site (bool).

`coordinates` - (Optional) Coordinates of the site, longitude and latitude. See [Coordinates ](#coordinates) below for details.

`disable_gpu` - (Optional) GPU is not enabled for this Site (bool).

`enable_gpu` - (Optional) GPU is enabled for this Site (bool).

`k8s_cluster` - (Optional) Site Local K8s API access is enabled, using k8s_cluster object. See [ref](#ref) below for details.

`no_k8s_cluster` - (Optional) Site Local K8s API access is disabled (bool).

`local_control_plane` - (Optional) Site Local control plane is enabled. See [Local Control Plane ](#local-control-plane) below for details.

`no_local_control_plane` - (Optional) Site Local control plane is disabled (bool).

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (bool).

`master_nodes` - (Required) Names of master nodes (`List of String`).

`custom_network_config` - (Optional) Use custom networking configuration. See [Custom Network Config ](#custom-network-config) below for details.

`default_network_config` - (Optional) Use default networking configuration based on certified hardware. (bool).

`os` - (Optional) Operating System Details. See [Os ](#os) below for details.

`custom_storage_config` - (Optional) Use custom storage configuration. See [Custom Storage Config ](#custom-storage-config) below for details.

`default_storage_config` - (Optional) Use default storage configuration (bool).

`sw` - (Optional) Volterra Software Details. See [Sw ](#sw) below for details.

`allow_all_usb` - (Optional) All USB devices are allowed (bool).

`deny_all_usb` - (Optional) All USB devices are denied (bool).

`usb_policy` - (Optional) Allow only specific USB devices. See [ref](#ref) below for details.

`volterra_certified_hw` - (Required) Name for generic server certified hardware to form this voltstack site. (`String`).

`worker_nodes` - (Optional) Names of worker nodes (`List of String`).

### Active Backup

Configure active/backup based bond device.

### Active Forward Proxy Policies

Enable Forward Proxy for this site and manage policies.

`forward_proxy_policies` - (Required) List of Forward Proxy Policies. See [ref](#ref) below for details.

### Active Network Policies

Network Policies active for this site..

`network_policies` - (Required) Ordered List of Network Policies active for this network firewall. See [ref](#ref) below for details.

### Automatic From End

Assign automatically from End of the first network in the list.

### Automatic From Start

Assign automatically from start of the first network in the list.

### Bgp Config

BGP configuration for local control plane.

`asn` - (Required) Autonomous System Number (`Int`).

`peers` - (Optional) BGP parameters for peer. See [Peers ](#peers) below for details.

### Blindfold Secret Info

Blindfold Secret is used for the secrets managed by Volterra Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Bond Device List

Configure Bond Devices for this voltstack site.

`bond_devices` - (Required) List of bond devices for this fleet. See [Bond Devices ](#bond-devices) below for details.

### Bond Devices

List of bond devices for this fleet.

`devices` - (Required) Ethernet devices that will make up this bond (`String`).

`active_backup` - (Optional) Configure active/backup based bond device (bool).

`lacp` - (Optional) Configure LACP (802.3ad) based bond device. See [Lacp ](#lacp) below for details.

`link_polling_interval` - (Required) Link polling interval in millisecond (`Int`).

`link_up_delay` - (Required) Milliseconds wait before link is declared up (`Int`).

`name` - (Required) Bond device name (`String`).

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted .

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Cluster

Configuration will apply to given device on all nodes of the site..

### Cluster Static Ip

Static IP configuration for a specific node.

`interface_ip_map` - (Optional) Map of Node to Static ip configuration value, Key:Node, Value:IP Address. See [Interface Ip Map ](#interface-ip-map) below for details.

### Coordinates

Coordinates of the site, longitude and latitude.

`latitude` - (Optional) Latitude of the site location (`Float`).

`longitude` - (Optional) longitude of site location (`Float`).

### Custom Certificate

Certificates for generating intermediate certificate for TLS interception..

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Custom Network Config

Use custom networking configuration.

`bgp_peer_address` - (Optional) to fetch BGP peer address from site Object. This can be used to change peer address per site in fleet. (`String`).

`bgp_router_id` - (Optional) fetch BGP router ID from site object. (`String`).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`default_interface_config` - (Optional) Interface configuration is done based on certified hardware for this site (bool).

`interface_list` - (Optional) Add all interfaces belonging to this site. See [Interface List ](#interface-list) below for details.

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`outside_nameserver` - (Optional) Optional DNS server IP to be used for name resolution in local network (`String`).

`outside_vip` - (Optional) Optional common virtual IP across all nodes to be used as automatic VIP for site local network. (`String`).

`site_to_site_tunnel_ip` - (Optional) Optional, VIP in the site_to_site_network_type configured above used for terminating IPSec/SSL tunnels created with SiteMeshGroup. (`String`).

`default_config` - (Optional) Use default configuration for site local network (bool).

`slo_config` - (Optional) Configuration for site local network. See [Slo Config ](#slo-config) below for details.

`tunnel_dead_timeout` - (Optional) When not set (== 0), a default value of 10000 msec will be used. (`Int`).

`vip_vrrp_mode` - (Optional) When Outside VIP / Inside VIP are configured, it is recommended to turn on vrrp and also configure BGP. (`String`).

### Custom Storage Config

Use custom storage configuration.

`no_static_routes` - (Optional) Static Routes not required for storage network. (bool).

`static_routes` - (Optional) Manage static routes for storage network.. See [Static Routes ](#static-routes) below for details.

`default_storage_class` - (Optional) Use only default storage class in kubernetes (bool).

`storage_class_list` - (Optional) Add additional custom storage classes in kubernetes. See [Storage Class List ](#storage-class-list) below for details.

`no_storage_device` - (Optional) This site does not have any storage devices (bool).

`storage_device_list` - (Optional) Add all storage devices belonging to this site. See [Storage Device List ](#storage-device-list) below for details.

`no_storage_interfaces` - (Optional) This site does not have any storage interfaces (bool).

`storage_interface_list` - (Optional) Add all storage interfaces belonging to this site. See [Storage Interface List ](#storage-interface-list) below for details.

### Dedicated Interface

Networking configuration for dedicated interface is configured locally on site e.g. (outside/inside)Ethernet.

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

Fallback management interfaces can be made into dedicated management interface.

`device` - (Required) Name of the device for which interface is configured (`String`).

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site. (bool).

`node` - (Optional) Configuration will apply to a device on the given node of the site. (`String`).

### Default Config

Use default configuration for site local network.

### Default Interface Config

Interface configuration is done based on certified hardware for this site.

### Default Os Version

Will assign latest available OS version.

### Default Storage Class

Use only default storage class in kubernetes.

### Default Sw Version

Will assign latest available SW version.

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

### Disable Forward Proxy

Forward Proxy is disabled for this connector.

### Disable Interception

Disable Interception.

### Domain Match

Domain value or regular expression to match.

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Enable For All Domains

Enable interception for all domains.

### Enable Forward Proxy

Forward Proxy is enabled for this connector.

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2000 (2 seconds) (`Int`).

`max_connect_attempts` - (Optional) Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1. (`Int`).

`no_interception` - (Optional) No TLS interception is enabled for this network connector (bool).

`tls_intercept` - (Optional) Specify TLS interception configuration for the network connector. See [Tls Intercept ](#tls-intercept) below for details.

`white_listed_ports` - (Optional) Example "tmate" server port (`Int`).

`white_listed_prefixes` - (Optional) Example "tmate" server ip (`String`).

### Enable Interception

Enable Interception.

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

### Forward Proxy Allow All

Enable Forward Proxy for this site and allow all requests..

### Global Network Connections

Global network connections.

`sli_to_global_dr` - (Optional) Site local inside is connected directly to a given global network. See [Sli To Global Dr ](#sli-to-global-dr) below for details.

`slo_to_global_dr` - (Optional) Site local outside is connected directly to a given global network. See [Slo To Global Dr ](#slo-to-global-dr) below for details.

`disable_forward_proxy` - (Optional) Forward Proxy is disabled for this connector (bool).

`enable_forward_proxy` - (Optional) Forward Proxy is enabled for this connector. See [Enable Forward Proxy ](#enable-forward-proxy) below for details.

### Global Network List

List of global network connections.

`global_network_connections` - (Required) Global network connections. See [Global Network Connections ](#global-network-connections) below for details.

### Inside Vn

Local control plane will work on inside network.

### Interception Rules

List of ordered rules to enable or disable for TLS interception.

`domain_match` - (Required) Domain value or regular expression to match. See [Domain Match ](#domain-match) below for details.

`disable_interception` - (Optional) Disable Interception (bool).

`enable_interception` - (Optional) Enable Interception (bool).

### Interface Ip Map

Configured address for every node.

`interface_ip_map` - (Optional) Map of Site:Node to IP address. (`String`).

### Interface List

Add all interfaces belonging to this site.

`interfaces` - (Required) Configure network interfaces for this voltstack. See [Interfaces ](#interfaces) below for details.

### Interfaces

Configure network interfaces for this voltstack.

`description` - (Optional) Description for this Interface (`String`).

`dedicated_interface` - (Optional) Networking configuration for dedicated interface is configured locally on site e.g. (outside/inside)Ethernet. See [Dedicated Interface ](#dedicated-interface) below for details.

`dedicated_management_interface` - (Optional) Fallback management interfaces can be made into dedicated management interface. See [Dedicated Management Interface ](#dedicated-management-interface) below for details.

`ethernet_interface` - (Optional) Ethernet interface configuration.. See [Ethernet Interface ](#ethernet-interface) below for details.

`tunnel_interface` - (Optional) Tunnel interface, Ipsec tunnels to other networking devices.. See [Tunnel Interface ](#tunnel-interface) below for details.

`labels` - (Optional) Add Labels for this Interface, these labels can be used in network policy (`String`).

### Is Primary

This interface is primary.

### Lacp

Configure LACP (802.3ad) based bond device.

`rate` - (Optional) Interval in seconds to transmit LACP packets (`Int`).

### Last Address

Last usable address from the network prefix is chosen as default gateway.

### Local Control Plane

Site Local control plane is enabled.

`bgp_config` - (Optional) BGP configuration for local control plane. See [Bgp Config ](#bgp-config) below for details.

`inside_vn` - (Optional) Local control plane will work on inside network (bool).

`outside_vn` - (Optional) Local control plane will work on outside network (bool).

### Monitor

Link Quality Monitoring parameters. Choosing the option will enable link quality monitoring..

### Monitor Disabled

Link quality monitoring disabled on the interface..

### Netapp Trident

Storage class Device configuration for NetApp Trident.

`selector` - (Optional) The volume will have the aspects defined in the chosen virtual pool. (`String`).

`storage_pools` - (Optional) The storagePools parameter is used to further restrict the set of pools that match any specified attributes (`String`).

### No Dc Cluster Group

This site is not a member of dc cluster group.

### No Forward Proxy

Disable Forward Proxy for this site.

### No Global Network

No global network to connect.

### No Interception

No TLS interception is enabled for this network connector.

### No Ipv6 Address

Interface does not have an IPv6 Address..

### No Network Policy

Network Policy is disabled for this site..

### No Static Routes

Static Routes disabled for site local network..

### No Storage Device

This site does not have any storage devices.

### No Storage Interfaces

This site does not have any storage interfaces.

### Node Static Ip

Static IP configuration for the Node.

`default_gw` - (Optional) IP address of the default gateway. (`String`).

`dns_server` - (Optional) IP address of the DNS server (`String`).

`ip_address` - (Required) IP address of the interface and prefix length (`String`).

### Not Primary

This interface is not primary.

### Openebs Enterprise

Storage class Device configuration for OpenEBS Enterprise.

`protocol` - (Optional) Defines type of transport protocol used to mount the PV to the worker node hosting the associated application pod (NVMe-oF) (`String`).

`replication` - (Optional) Replication sets the replication factor of the PV, i.e. the number of data replicas to be maintained for it such as 1 or 3. (`Int`).

### Os

Operating System Details.

`default_os_version` - (Optional) Will assign latest available OS version (bool).

`operating_system_version` - (Optional) Operating System Version is optional parameter, which allows to specify target OS version for particular site e.g. 7.2009.10. (`String`).

### Outside Vn

Local control plane will work on outside network.

### Peers

BGP parameters for peer.

### Policy

Policy to enable/disable specific domains, with implicit enable all domains.

`interception_rules` - (Required) List of ordered rules to enable or disable for TLS interception. See [Interception Rules ](#interception-rules) below for details.

### Pools

List of non overlapping ip address ranges..

`end_ip` - (Optional) 10.1.1.200 with prefix length of 24, end offset is 0.0.0.200 (`String`).

`exclude` - (Optional) If exclude is true, IP addresses are not assigned from this range. (`Bool`).

`start_ip` - (Optional) 10.1.1.5 with prefix length of 24, start offset is 0.0.0.5 (`String`).

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Pure Service Orchestrator

Storage class Device configuration for Pure Service Orchestrator.

`backend` - (Optional) The volume will have the aspects defined in the chosen virtual pool. (`String`).

`bandwidth_limit` - (Optional) Valid unit symbols are K, M, G, representing KiB, MiB, and GiB. (`String`).

`iops_limit` - (Optional) Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not defined. (`Int`).

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

### Sli To Global Dr

Site local inside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Slo Config

Configuration for site local network.

`dc_cluster_group` - (Optional) This site is member of dc cluster group via network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (bool).

`labels` - (Optional) Add Labels for this network, these labels can be used in network policy (`String`).

`no_static_routes` - (Optional) Static Routes disabled for site local network. (bool).

`static_routes` - (Optional) Manage static routes for site local network.. See [Static Routes ](#static-routes) below for details.

### Slo To Global Dr

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

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

### Static Routes

Manage static routes for site local network..

`static_routes` - (Required) List of static routes. See [Static Routes ](#static-routes) below for details.

### Storage Class List

Add additional custom storage classes in kubernetes.

`storage_classes` - (Optional) List of custom storage classes. See [Storage Classes ](#storage-classes) below for details.

### Storage Classes

List of custom storage classes.

`advanced_storage_parameters` - (Optional) Map of parameter name and string value (`String`).

`allow_volume_expansion` - (Optional) Allow volume expansion. (`Bool`).

`default_storage_class` - (Optional) Make this storage class default storage class for the K8s cluster (`Bool`).

`description` - (Optional) Description for this storage class (`String`).

`netapp_trident` - (Optional) Storage class Device configuration for NetApp Trident. See [Netapp Trident ](#netapp-trident) below for details.

`openebs_enterprise` - (Optional) Storage class Device configuration for OpenEBS Enterprise. See [Openebs Enterprise ](#openebs-enterprise) below for details.

`pure_service_orchestrator` - (Optional) Storage class Device configuration for Pure Service Orchestrator. See [Pure Service Orchestrator ](#pure-service-orchestrator) below for details.

`reclaim_policy` - (Optional) Reclaim Policy (`String`).

`storage_class_name` - (Required) Name of the storage class as it will appear in K8s. (`String`).

`storage_device` - (Required) Storage device that this class will use. The Device name defined at previous step. (`String`).

### Storage Device List

Add all storage devices belonging to this site.

`storage_devices` - (Optional) List of custom storage devices. See [Storage Devices ](#storage-devices) below for details.

### Storage Devices

List of custom storage devices.

`advanced_advanced_parameters` - (Optional) Map of parameter name and string value (`String`).

`netapp_trident` - (Optional) Device configuration for NetApp Trident. See [Netapp Trident ](#netapp-trident) below for details.

`openebs_enterprise` - (Optional) Device configuration for Pure Storage Service Orchestrator. See [Openebs Enterprise ](#openebs-enterprise) below for details.

`pure_service_orchestrator` - (Optional) Device configuration for Pure Storage Service Orchestrator. See [Pure Service Orchestrator ](#pure-service-orchestrator) below for details.

`storage_device` - (Required) Storage device and device unit (`String`).

### Storage Interface

Configure storage interface for this voltstack.

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

### Storage Interface List

Add all storage interfaces belonging to this site.

`storage_interfaces` - (Required) Configure storage interfaces for this voltstack. See [Storage Interfaces ](#storage-interfaces) below for details.

### Storage Interfaces

Configure storage interfaces for this voltstack.

`description` - (Optional) Description for this Interface (`String`).

`labels` - (Optional) Add Labels for this Interface, these labels can be used in network policy (`String`).

`storage_interface` - (Required) Configure storage interface for this voltstack. See [Storage Interface ](#storage-interface) below for details.

### Storage Network

Interface belongs to site local network inside.

### Sw

Volterra Software Details.

`default_sw_version` - (Optional) Will assign latest available SW version (bool).

`volterra_software_version` - (Optional) Volterra Software Version is optional parameter, which allows to specify target SW version for particular site e.g. crt-20210329-1002. (`String`).

### Tls Intercept

Specify TLS interception configuration for the network connector.

`enable_for_all_domains` - (Optional) Enable interception for all domains (bool).

`policy` - (Optional) Policy to enable/disable specific domains, with implicit enable all domains. See [Policy ](#policy) below for details.

`custom_certificate` - (Optional) Certificates for generating intermediate certificate for TLS interception.. See [Custom Certificate ](#custom-certificate) below for details.

`volterra_certificate` - (Optional) Volterra certificates for generating intermediate certificate for TLS interception. (bool).

`trusted_ca_url` - (Optional) Custom trusted CA certificates for validating upstream server certificate (`String`).

`volterra_trusted_ca` - (Optional) Default volterra trusted CA list for validating upstream server certificate (bool).

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

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Volterra Certificate

Volterra certificates for generating intermediate certificate for TLS interception..

### Volterra Trusted Ca

Default volterra trusted CA list for validating upstream server certificate.

### Wingman Secret Info

Secret is given as bootstrap secret in Volterra Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured voltstack_site.
