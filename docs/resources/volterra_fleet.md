---

page_title: "Volterra: fleet"

description: "The fleet allows CRUD of Fleet resource on Volterra SaaS"
-----------------------------------------------------------------------

Resource volterra_fleet
=======================

The Fleet allows CRUD of Fleet resource on Volterra SaaS

~> **Note:** Please refer to [Fleet API docs](https://docs.cloud.f5.com/docs/api/fleet) to learn more

Example Usage
-------------

```hcl
resource "volterra_fleet" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "no_bond_devices bond_device_list" must be set
  no_bond_devices = true

  // One of the arguments from this list "no_dc_cluster_group dc_cluster_group dc_cluster_group_inside" must be set
  no_dc_cluster_group = true
  fleet_label         = ["sfo"]

  // One of the arguments from this list "disable_gpu enable_gpu enable_vgpu" must be set

  enable_vgpu {
    feature_type   = "feature_type"
    server_address = "gridlicense1.example.com"
    server_port    = "7070"
  }

  // One of the arguments from this list "default_config device_list interface_list" must be set

  interface_list {
    interfaces {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }
  // One of the arguments from this list "logs_streaming_disabled log_receiver" must be set
  logs_streaming_disabled = true
  // One of the arguments from this list "sriov_interfaces default_sriov_interface" must be set
  default_sriov_interface = true
  // One of the arguments from this list "default_storage_class storage_class_list" must be set
  default_storage_class = true
  // One of the arguments from this list "no_storage_device storage_device_list" must be set
  no_storage_device = true
  // One of the arguments from this list "no_storage_interfaces storage_interface_list" must be set
  no_storage_interfaces = true
  // One of the arguments from this list "storage_static_routes no_storage_static_routes" must be set
  no_storage_static_routes = true
  // One of the arguments from this list "allow_all_usb usb_policy deny_all_usb" must be set
  allow_all_usb = true
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

`blocked_services` - (Optional) Configuration to block the default services allowed by the platform. See [Blocked Services ](#blocked-services) below for details.

`bond_device_list` - (Optional) Configure Bond Devices for this fleet. See [Bond Device List ](#bond-device-list) below for details.

`no_bond_devices` - (Optional) No Bond Devices configured for this Fleet (bool).

`dc_cluster_group` - (Optional) This fleet is member of dc cluster group via site local network. See [ref](#ref) below for details.

`dc_cluster_group_inside` - (Optional) This fleet is member of dc cluster group via site local inside network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This fleet is not a member of a DC cluster group (bool).

`enable_default_fleet_config_download` - (Optional) Enable default fleet config, It must be set for storage config and gpu config (`Bool`).

`fleet_label` - (Required) fleet_label with "sfo" will create a known_label "ves.io/fleet=sfo" in tenant for the fleet (`String`).

`disable_gpu` - (Optional) GPU is not enabled for this fleet (bool).

`enable_gpu` - (Optional) GPU is enabled for this fleet (bool).

`enable_vgpu` - (Optional) Enable NVIDIA vGPU hosted on VMware. See [Enable Vgpu ](#enable-vgpu) below for details.

`inside_virtual_network` - (Optional) Default inside (site local) virtual network for the fleet. See [ref](#ref) below for details.

`default_config` - (Optional) Use default configuration for interfaces belonging to this fleet (bool).

`device_list` - (Optional) Add device for all interfaces belonging to this fleet. See [Device List ](#device-list) below for details.

`interface_list` - (Optional) Add all interfaces belonging to this fleet. See [Interface List ](#interface-list) below for details.

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (bool).

`network_connectors` - (Optional) The network connectors configuration is applied on all sites that are member of the fleet.. See [ref](#ref) below for details.

`network_firewall` - (Optional) The Network Firewall is applied on Virtual Networks of type site local network and site local inside network. See [ref](#ref) below for details.

`operating_system_version` - (Optional) Current Operating System version can be overridden via site config. (`String`).

`outside_virtual_network` - (Optional) Default outside (site local) virtual network for the fleet. See [ref](#ref) below for details.

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode to optimize for L3 or L7 networking. See [Performance Enhancement Mode ](#performance-enhancement-mode) below for details.

`default_sriov_interface` - (Optional) Disable Single Root I/O Virtualization interfaces (bool).

`sriov_interfaces` - (Optional) Use custom Single Root I/O Virtualization interfaces. See [Sriov Interfaces ](#sriov-interfaces) below for details.

`default_storage_class` - (Optional) Use only default storage class in kubernetes (bool).

`storage_class_list` - (Optional) Add additional custom storage classes in kubernetes for this fleet. See [Storage Class List ](#storage-class-list) below for details.

`no_storage_device` - (Optional) This fleet does not have any storage devices (bool).

`storage_device_list` - (Optional) Add all storage devices belonging to this fleet. See [Storage Device List ](#storage-device-list) below for details.

`no_storage_interfaces` - (Optional) This fleet does not have any storage interfaces (bool).

`storage_interface_list` - (Optional) Add all storage interfaces belonging to this fleet. See [Storage Interface List ](#storage-interface-list) below for details.

`no_storage_static_routes` - (Optional) This fleet does not have any storage static routes (bool).

`storage_static_routes` - (Optional) Add all storage storage static routes. See [Storage Static Routes ](#storage-static-routes) below for details.

`allow_all_usb` - (Optional) All USB devices are allowed (bool).

`deny_all_usb` - (Optional) All USB devices are denied (bool).

`usb_policy` - (Optional) Allow only specific USB devices. See [ref](#ref) below for details.

`disable_vm` - (Optional) VMs support is not enabled for this fleet (bool).

`enable_vm` - (Optional) VMs support is enabled for this fleet. See [Enable Vm ](#enable-vm) below for details.

`volterra_software_version` - (Optional) Current software installed can be overridden via site config. (`String`).

### Active Backup

Configure active/backup based bond device.

### Blocked Services

Configuration to block the default services allowed by the platform.

`dns` - (Optional) Matches DNS port 53 (bool).

`ssh` - (Optional) Matches ssh port 22 (bool).

`web_user_interface` - (Optional) Matches the web user interface port (bool).

`network_type` - (Optional) Network type in which these ports get blocked (`String`).

### Bond Device List

Configure Bond Devices for this fleet.

`bond_devices` - (Required) List of bond devices. See [Bond Devices ](#bond-devices) below for details.

### Bond Devices

List of bond devices.

`devices` - (Required) Ethernet devices that will make up this bond (`String`).

`active_backup` - (Optional) Configure active/backup based bond device (bool).

`lacp` - (Optional) Configure LACP (802.3ad) based bond device. See [Lacp ](#lacp) below for details.

`link_polling_interval` - (Required) Link polling interval in millisecond (`Int`).

`link_up_delay` - (Required) Milliseconds wait before link is declared up (`Int`).

`name` - (Required) Bond device name (`String`).

### Custom Storage

Storage configuration for Custom Storage.

`yaml` - (Optional) K8s YAML for StorageClass (`String`).

### Device List

Add device for all interfaces belonging to this fleet.

`devices` - (Optional) device instance specific sections. See [Devices ](#devices) below for details.

### Devices

device instance specific sections.

`network_device` - (Optional) Device instance is a networking device like ethernet. See [Network Device ](#network-device) below for details.

`name` - (Optional) Name of the device including the unit number (e.g. eth0 or disk1). The name must match name of device in host-os of node (`String`).

`owner` - (Required) This option is not yet supported (`String`).

### Dns

Matches DNS port 53.

### Enable Vgpu

Enable NVIDIA vGPU hosted on VMware.

`feature_type` - (Required) Set Feature to be enabled (`String`).

`server_address` - (Optional) Set License Server Address (`String`).

`server_port` - (Optional) Set License Server port number (`Int`).

### Enable Vm

VMs support is enabled for this fleet.

### Hpe Storage

Storage configuration for HPE Storage.

`allow_mutations` - (Optional) mutation can override specified parameters (`String`).

`allow_overrides` - (Optional) PVC can override specified parameters (`String`).

`dedupe_enabled` - (Optional) Indicates that the volume should enable deduplication. (`Bool`).

`description` - (Optional) The SecretName parameter is used to identify name of secret to identify backend storage's auth information (`String`).

`destroy_on_delete` - (Optional) Indicates the backing Nimble volume (including snapshots) should be destroyed when the PVC is deleted (`Bool`).

`encrypted` - (Optional) Indicates that the volume should be encrypted. (`Bool`).

`folder` - (Optional) The name of the folder in which to place the volume. (`String`).

`limit_iops` - (Optional) The IOPS limit of the volume. (`Int`).

`limit_mbps` - (Optional) The IOPS limit of the volume. (`Int`).

`performance_policy` - (Optional) The name of the performance policy to assign to the volume. (`String`).

`pool` - (Optional) The name of the pool in which to place the volume. (`String`).

`protection_template` - (Optional) The name of the performance policy to assign to the volume. (`String`).

`secret_name` - (Optional) The SecretName parameter is used to identify name of secret to identify backend storage's auth information (`String`).

`secret_namespace` - (Optional) The SecretNamespace parameter is used to identify name of namespace where secret resides (`String`).

`sync_on_detach` - (Optional) Indicates that a snapshot of the volume should be synced to the replication partner each time it is detached from a node. (`Bool`).

`thick` - (Optional) Indicates that the volume should be thick provisioned. (`Bool`).

### Interface List

Add all interfaces belonging to this fleet.

`interfaces` - (Required) Add all interfaces belonging to this fleet. See [ref](#ref) below for details.

### Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

### Jumbo

L3 performance mode enhancement to use jumbo frame.

### Lacp

Configure LACP (802.3ad) based bond device.

`rate` - (Optional) Interval in seconds to transmit LACP packets (`Int`).

### Netapp Trident

Storage class Device configuration for NetApp Trident.

`selector` - (Optional) The volume will have the aspects defined in the chosen virtual pool. (`String`).

`storage_pools` - (Optional) The storagePools parameter is used to further restrict the set of pools that match any specified attributes (`String`).

### Network Device

Device instance is a networking device like ethernet.

`interface` - (Required) if use is NETWORK_INTERFACE_USE_INSIDE, the virtual-network must of type VIRTUAL_NETWORK_SITE_LOCAL_INSIDE. See [ref](#ref) below for details.

`use` - (Required) Specifies if the network interface is connected to inside network or outside network. (`String`).

### Nexthop

Nexthop for the route.

`interface` - (Optional) Nexthop is network interface when type is "Network-Interface". See [ref](#ref) below for details.

`nexthop_address` - (Optional) Nexthop address when type is "Use-Configured". See [Nexthop Address ](#nexthop-address) below for details.

`type` - (Optional) Identifies the type of next-hop (`String`).

### Nexthop Address

Nexthop address when type is "Use-Configured".

`ipv4` - (Optional) IPv4 Address. See [Ipv4 ](#ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ipv6 ](#ipv6) below for details.

### No Jumbo

L3 performance mode enhancement without jumbo frame.

### Perf Mode L3 Enhanced

When the mode is toggled to l3 enhanced, traffic disruption will be seen.

`jumbo` - (Optional) L3 performance mode enhancement to use jumbo frame (bool).

`no_jumbo` - (Optional) L3 performance mode enhancement without jumbo frame (bool).

### Perf Mode L7 Enhanced

When the mode is toggled to l7 enhanced, traffic disruption will be seen.

### Performance Enhancement Mode

Performance Enhancement Mode to optimize for L3 or L7 networking.

`perf_mode_l3_enhanced` - (Optional) When the mode is toggled to l3 enhanced, traffic disruption will be seen. See [Perf Mode L3 Enhanced ](#perf-mode-l3-enhanced) below for details.

`perf_mode_l7_enhanced` - (Optional) When the mode is toggled to l7 enhanced, traffic disruption will be seen (bool).

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

### Sriov Interface

Use custom SR-IOV interfaces Configuration.

`interface_name` - (Required) Name of SR-IOV physical interface (`String`).

`number_of_vfio_vfs` - (Optional) Number of virtual functions reserved for VNFs and DPDK-based CNFs (`Int`).

`number_of_vfs` - (Required) Total number of virtual functions (`Int`).

### Sriov Interfaces

Use custom Single Root I/O Virtualization interfaces.

`sriov_interface` - (Optional) Use custom SR-IOV interfaces Configuration. See [Sriov Interface ](#sriov-interface) below for details.

### Ssh

Matches ssh port 22.

### Storage Class List

Add additional custom storage classes in kubernetes for this fleet.

`storage_classes` - (Optional) List of custom storage classes. See [Storage Classes ](#storage-classes) below for details.

### Storage Classes

List of custom storage classes.

`advanced_storage_parameters` - (Optional) Map of parameter name and string value (`String`).

`allow_volume_expansion` - (Optional) Allow volume expansion. (`Bool`).

`default_storage_class` - (Optional) Make this storage class default storage class for the K8s cluster (`Bool`).

`description` - (Optional) Description for this storage class (`String`).

`custom_storage` - (Optional) Storage configuration for Custom Storage. See [Custom Storage ](#custom-storage) below for details.

`hpe_storage` - (Optional) Storage configuration for HPE Storage. See [Hpe Storage ](#hpe-storage) below for details.

`netapp_trident` - (Optional) Storage class Device configuration for NetApp Trident. See [Netapp Trident ](#netapp-trident) below for details.

`pure_service_orchestrator` - (Optional) Storage class Device configuration for Pure Service Orchestrator. See [Pure Service Orchestrator ](#pure-service-orchestrator) below for details.

`reclaim_policy` - (Optional) Reclaim Policy (`String`).

`storage_class_name` - (Required) Name of the storage class as it will appear in K8s. (`String`).

`storage_device` - (Required) Storage device that this class will use. The Device name defined at previous step. (`String`).

### Storage Device List

Add all storage devices belonging to this fleet.

`storage_devices` - (Optional) List of custom storage devices. See [Storage Devices ](#storage-devices) below for details.

### Storage Devices

List of custom storage devices.

`advanced_advanced_parameters` - (Optional) Map of parameter name and string value (`String`).

`custom_storage` - (Optional) Device configuration for Custom Storage (bool).

`hpe_storage` - (Optional) Device configuration for HPE Storage. See [Hpe Storage ](#hpe-storage) below for details.

`netapp_trident` - (Optional) Device configuration for NetApp Trident. See [Netapp Trident ](#netapp-trident) below for details.

`pure_service_orchestrator` - (Optional) Device configuration for Pure Storage Service Orchestrator. See [Pure Service Orchestrator ](#pure-service-orchestrator) below for details.

`storage_device` - (Required) Storage device and device unit (`String`).

### Storage Interface List

Add all storage interfaces belonging to this fleet.

`interfaces` - (Required) Add all interfaces belonging to this fleet. See [ref](#ref) below for details.

### Storage Routes

List of storage static routes.

`attrs` - (Optional) List of route attributes associated with the static route (`List of Strings`).

`labels` - (Optional) Add Labels for this Static Route, these labels can be used in network policy (`String`).

`nexthop` - (Optional) Nexthop for the route. See [Nexthop ](#nexthop) below for details.

`subnets` - (Required) List of route prefixes. See [Subnets ](#subnets) below for details.

### Storage Static Routes

Add all storage storage static routes.

`storage_routes` - (Required) List of storage static routes. See [Storage Routes ](#storage-routes) below for details.

### Subnets

List of route prefixes.

`ipv4` - (Optional) IPv4 Subnet Address. See [Ipv4 ](#ipv4) below for details.

`ipv6` - (Optional) IPv6 Subnet Address. See [Ipv6 ](#ipv6) below for details.

### Web User Interface

Matches the web user interface port.

Attribute Reference
-------------------

-	`id` - This is the id of the configured fleet.
