---

page_title: "Volterra: voltstack_site"
description: "The voltstack_site allows CRUD of Voltstack Site resource on Volterra SaaS"

---

Resource volterra_voltstack_site
================================

The Voltstack Site allows CRUD of Voltstack Site resource on Volterra SaaS

~> **Note:** Please refer to [Voltstack Site API docs](https://docs.cloud.f5.com/docs-v2/api/views-voltstack-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_voltstack_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "blocked_services default_blocked_services" must be set

  default_blocked_services = true

  // One of the arguments from this list "bond_device_list no_bond_devices" must be set

  no_bond_devices = true

  // One of the arguments from this list "disable_gpu enable_gpu enable_vgpu" must be set

  disable_gpu = true

  // One of the arguments from this list "k8s_cluster no_k8s_cluster" must be set

  no_k8s_cluster = true

  // One of the arguments from this list "log_receiver logs_streaming_disabled" must be set

  logs_streaming_disabled = true
  master_node_configuration {
    name = "master-0"

    public_ip = "192.168.0.156"
  }

  // One of the arguments from this list "custom_network_config default_network_config" must be set

  default_network_config = true

  // One of the arguments from this list "default_sriov_interface sriov_interfaces" must be set

  sriov_interfaces {
    sriov_interface {
      interface_name = "eth0"

      number_of_vfio_vfs = "2"

      number_of_vfs = "3"
    }
  }

  // One of the arguments from this list "custom_storage_config default_storage_config" must be set

  default_storage_config = true

  // One of the arguments from this list "allow_all_usb deny_all_usb usb_policy" must be set

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

###### One of the arguments from this list "blocked_services, default_blocked_services" must be set

`blocked_services` - (Optional) Use custom blocked services configuration. See [Blocked Services Choice Blocked Services ](#blocked-services-choice-blocked-services) below for details.

`default_blocked_services` - (Optional) Use default behavior, which allows SSH (port 22), HTTPS (port 65500) and ICMP node access in blocked services (`Bool`).

###### One of the arguments from this list "bond_device_list, no_bond_devices" must be set

`bond_device_list` - (Optional) Configure Bond Devices for this App Stack site. See [Bond Choice Bond Device List ](#bond-choice-bond-device-list) below for details.

`no_bond_devices` - (Optional) No Bond Devices configured for this App Stack site (`Bool`).

`coordinates` - (Optional) Coordinates of the site, longitude and latitude. See [Coordinates ](#coordinates) below for details.

`custom_dns` - (Optional) custom dns configure to the CE site. See [Custom Dns ](#custom-dns) below for details.

###### One of the arguments from this list "disable_gpu, enable_gpu, enable_vgpu" must be set

`disable_gpu` - (Optional) GPU is not enabled for this Site (`Bool`).

`enable_gpu` - (Optional) GPU is enabled for this Site (`Bool`).

`enable_vgpu` - (Optional) Enable NVIDIA vGPU hosted on VMware. See [Gpu Choice Enable Vgpu ](#gpu-choice-enable-vgpu) below for details.

###### One of the arguments from this list "k8s_cluster, no_k8s_cluster" must be set

`k8s_cluster` - (Optional) Site Local K8s API access is enabled, using k8s_cluster object. See [ref](#ref) below for details.

`no_k8s_cluster` - (Optional) Site Local K8s API access is disabled (`Bool`).

`kubernetes_upgrade_drain` - (Optional) Enable Kubernetes Drain during OS or SW upgrade. See [Kubernetes Upgrade Drain ](#kubernetes-upgrade-drain) below for details.

###### One of the arguments from this list "local_control_plane, no_local_control_plane" can be set

`local_control_plane` - (Optional) Site Local control plane is enabled. See [Local Control Plane Choice Local Control Plane ](#local-control-plane-choice-local-control-plane) below for details.(Deprecated)

`no_local_control_plane` - (Optional) Site Local control plane is disabled (`Bool`).(Deprecated)

###### One of the arguments from this list "log_receiver, logs_streaming_disabled" must be set

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (`Bool`).

`master_node_configuration` - (Required) Configuration of master nodes. See [Master Node Configuration ](#master-node-configuration) below for details.

`master_nodes` - (Optional) Names of master nodes (`List of String`).(Deprecated)

###### One of the arguments from this list "custom_network_config, default_network_config" must be set

`custom_network_config` - (Optional) Use custom networking configuration. See [Network Cfg Choice Custom Network Config ](#network-cfg-choice-custom-network-config) below for details.

`default_network_config` - (Optional) Use default networking configuration based on certified hardware. (`Bool`).

`offline_survivability_mode` - (Optional) Enable/Disable offline survivability mode. See [Offline Survivability Mode ](#offline-survivability-mode) below for details.

`os` - (Optional) Operating System Details. See [Os ](#os) below for details.

###### One of the arguments from this list "default_sriov_interface, sriov_interfaces" must be set

`default_sriov_interface` - (Optional) Disable Single Root I/O Virtualization interfaces (`Bool`).

`sriov_interfaces` - (Optional) Use custom Single Root I/O Virtualization interfaces. See [Sriov Interface Choice Sriov Interfaces ](#sriov-interface-choice-sriov-interfaces) below for details.

###### One of the arguments from this list "custom_storage_config, default_storage_config" must be set

`custom_storage_config` - (Optional) Use custom storage configuration. See [Storage Cfg Choice Custom Storage Config ](#storage-cfg-choice-custom-storage-config) below for details.

`default_storage_config` - (Optional) Use default storage configuration (`Bool`).

`sw` - (Optional) F5XC Software Details. See [Sw ](#sw) below for details.

###### One of the arguments from this list "allow_all_usb, deny_all_usb, usb_policy" must be set

`allow_all_usb` - (Optional) All USB devices are allowed (`Bool`).

`deny_all_usb` - (Optional) All USB devices are denied (`Bool`).

`usb_policy` - (Optional) Allow only specific USB devices. See [ref](#ref) below for details.

###### One of the arguments from this list "disable_vm, enable_vm" can be set

`disable_vm` - (Optional) VMs support is not enabled for this Site (`Bool`).

`enable_vm` - (Optional) VMs support is enabled for this Site. See [Vm Choice Enable Vm ](#vm-choice-enable-vm) below for details.

`volterra_certified_hw` - (Required) Name for generic server certified hardware to form this App Stack site. (`String`).

`worker_nodes` - (Optional) Names of worker nodes (`List of String`).

### Coordinates

Coordinates of the site, longitude and latitude.

`latitude` - (Optional) Latitude of the site location (`Float`).

`longitude` - (Optional) longitude of site location (`Float`).

### Custom Dns

custom dns configure to the CE site.

`inside_nameserver` - (Optional) Optional DNS server IP to be used for name resolution in inside network (`String`).

`inside_nameserver_v6` - (Optional) Optional DNS server IPv6 to be used for name resolution in inside network (`String`).

`outside_nameserver` - (Optional) Optional DNS server IP to be used for name resolution in outside network (`String`).

`outside_nameserver_v6` - (Optional) Optional DNS server IPv6 to be used for name resolution in outside network (`String`).

### Kubernetes Upgrade Drain

Enable Kubernetes Drain during OS or SW upgrade.

###### One of the arguments from this list "disable_upgrade_drain, enable_upgrade_drain" must be set

`disable_upgrade_drain` - (Optional) x-displayName: "Disable Node by Node Upgrade" (`Bool`).

`enable_upgrade_drain` - (Optional) x-displayName: "Enable Node by Node Upgrade". See [Kubernetes Upgrade Drain Enable Choice Enable Upgrade Drain ](#kubernetes-upgrade-drain-enable-choice-enable-upgrade-drain) below for details.

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

`dhcp_option82_tag` - (Optional) Optional tag that can be given to this configuration (`String`).(Deprecated)

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

###### One of the arguments from this list "cluster_static_ip, fleet_static_ip, node_static_ip" must be set

`cluster_static_ip` - (Optional) Static IP configuration for a specific node. See [Network Prefix Choice Cluster Static Ip ](#network-prefix-choice-cluster-static-ip) below for details.

`fleet_static_ip` - (Optional) Static IP configuration for the fleet. See [Network Prefix Choice Fleet Static Ip ](#network-prefix-choice-fleet-static-ip) below for details.(Deprecated)

`node_static_ip` - (Optional) Static IP configuration for the Node. See [Network Prefix Choice Node Static Ip ](#network-prefix-choice-node-static-ip) below for details.

### Api Token Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Arrays Flash Array

For FlashArrays you must set the "mgmt_endpoint" and "api_token".

`default_fs_opt` - (Optional) Block volume default mkfs options. Not recommended to change! (`String`).

`default_fs_type` - (Required) Block volume default filesystem type. Not recommended to change! (`String`).

`default_mount_opts` - (Optional) Block volume default filesystem mount options. Not recommended to change! (`String`).

`disable_preempt_attachments` - (Optional) Enable/Disable attachment preemption! (`Bool`).

`flash_arrays` - (Required) For FlashArrays you must set the "mgmt_endpoint" and "api_token". See [Flash Array Flash Arrays ](#flash-array-flash-arrays) below for details.

`iscsi_login_timeout` - (Required) iSCSI login timeout in seconds. Not recommended to change! (`Int`).

`san_type` - (Required) Block volume access protocol, either ISCSI or FC (`String`).

### Arrays Flash Blade

Specify what storage flash blades should be managed the plugin.

`enable_snapshot_directory` - (Optional) Enable/Disable FlashBlade snapshots (`Bool`).

`export_rules` - (Optional) NFS Export rules (`String`).

`flash_blades` - (Required) For FlashBlades you must set the "mgmt_endpoint", "api_token" and nfs_endpoint. See [Flash Blade Flash Blades ](#flash-blade-flash-blades) below for details.

### Autoconfig Choice Host

auto configuration routers. This is similar to a DHCP Client..

### Autoconfig Choice Router

System behaves like auto config Router and provides auto config parameters. This is similar to a DHCP Server..

###### One of the arguments from this list "network_prefix, stateful" must be set

`network_prefix` - (Optional) Allowed only /64 prefix length as per RFC 4862 (`String`).

`stateful` - (Optional) works along with Router Advertisement' Managed flag. See [Address Choice Stateful ](#address-choice-stateful) below for details.

`dns_config` - (Optional) Dns information that needs to added in the RouterAdvetisement. See [Router Dns Config ](#router-dns-config) below for details.

### Backend Choice Netapp Backend Ontap Nas

Backend configuration for ONTAP NAS.

`auto_export_cidrs` - (Optional) List of CIDRs to filter Kubernetes’ node IPs against when autoExportPolicy is enabled. See [Netapp Backend Ontap Nas Auto Export Cidrs ](#netapp-backend-ontap-nas-auto-export-cidrs) below for details.

`auto_export_policy` - (Optional) Enable automatic export policy creation and updating (`Bool`).

`backend_name` - (Optional) Configuration of Backend Name. Driver is name + "_" + dataLIF (`String`).

`client_certificate` - (Optional) Please Enter Base64-encoded value of client certificate. Used for certificate-based auth. (`String`).

`client_private_key` - (Optional) Please Enter value of client private key. Used for certificate-based auth.. See [Netapp Backend Ontap Nas Client Private Key ](#netapp-backend-ontap-nas-client-private-key) below for details.

###### One of the arguments from this list "data_lif_dns_name, data_lif_ip" can be set

`data_lif_dns_name` - (Optional) Backend Data LIF IP Address's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

`data_lif_ip` - (Optional) Backend Data LIF IP Address is reachable at the given ip address (`String`).

`labels` - (Optional) List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection. (`String`).

`limit_aggregate_usage` - (Optional) Fail provisioning if usage is above this percentage. Not enforced by default. (`String`).

`limit_volume_size` - (Optional) Fail provisioning if requested volume size is above this value. Not enforced by default. (`String`).

###### One of the arguments from this list "management_lif_dns_name, management_lif_ip" must be set

`management_lif_dns_name` - (Optional) Backend Management LIF IP Address's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

`management_lif_ip` - (Optional) Backend Management LIF IP Address is reachable at the given ip address (`String`).

`nfs_mount_options` - (Optional) Comma-separated list of NFS mount options. Not enforced by default. (`String`).

`password` - (Optional) Please Enter you password. Password to connect to the cluster/SVM. See [Netapp Backend Ontap Nas Password ](#netapp-backend-ontap-nas-password) below for details.

`region` - (Optional) Virtual Pool Region (`String`).

`storage` - (Optional) List of Virtual Storage Pool definitions which are referred back by Storage Class label match selection.. See [Netapp Backend Ontap Nas Storage ](#netapp-backend-ontap-nas-storage) below for details.

`storage_driver_name` - (Required) Configuration of Backend Name (`String`).

`storage_prefix` - (Optional) Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated (`String`).

`svm` - (Optional) Storage virtual machine to use. Derived if an SVM managementLIF is specified (`String`).

`trusted_ca_certificate` - (Optional) Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based auth.. (`String`).

`username` - (Required) Username to connect to the cluster/SVM (`String`).

`volume_defaults` - (Optional) List of QoS volume defaults types. See [Netapp Backend Ontap Nas Volume Defaults ](#netapp-backend-ontap-nas-volume-defaults) below for details.

### Backend Choice Netapp Backend Ontap San

Backend configuration for ONTAP SAN.

###### One of the arguments from this list "no_chap, use_chap" can be set

`no_chap` - (Optional) CHAP disabled (`Bool`).

`use_chap` - (Optional) Device NetApp Backend ONTAP SAN CHAP configuration options for enabled CHAP. See [Chap Choice Use Chap ](#chap-choice-use-chap) below for details.

`client_certificate` - (Optional) Please Enter Base64-encoded value of client certificate. Used for certificate-based auth. (`String`).

`client_private_key` - (Optional) Please Enter value of client private key. Used for certificate-based auth.. See [Netapp Backend Ontap San Client Private Key ](#netapp-backend-ontap-san-client-private-key) below for details.

###### One of the arguments from this list "data_lif_dns_name, data_lif_ip" can be set

`data_lif_dns_name` - (Optional) Backend Data LIF IP Address's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

`data_lif_ip` - (Optional) Backend Data LIF IP Address is reachable at the given ip address (`String`).

`igroup_name` - (Optional) Name of the igroup for SAN volumes to use (`String`).

`labels` - (Optional) List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection. (`String`).

`limit_aggregate_usage` - (Optional) Fail provisioning if usage is above this percentage. Not enforced by default. (`Int`).

`limit_volume_size` - (Optional) Fail provisioning if requested volume size in GBi is above this value. Not enforced by default. (`Int`).

###### One of the arguments from this list "management_lif_dns_name, management_lif_ip" must be set

`management_lif_dns_name` - (Optional) Backend Management LIF IP Address's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

`management_lif_ip` - (Optional) Backend Management LIF IP Address is reachable at the given ip address (`String`).

`password` - (Optional) Please Enter you password. Password to connect to the cluster/SVM. See [Netapp Backend Ontap San Password ](#netapp-backend-ontap-san-password) below for details.

`region` - (Optional) Virtual Pool Region (`String`).

`storage` - (Optional) List of Virtual Storage Pool definitions which are referred back by Storage Class label match selection.. See [Netapp Backend Ontap San Storage ](#netapp-backend-ontap-san-storage) below for details.

`storage_driver_name` - (Required) Configuration of Backend Name (`String`).

`storage_prefix` - (Optional) Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated (`String`).

`svm` - (Optional) Storage virtual machine to use. Derived if an SVM managementLIF is specified (`String`).

`trusted_ca_certificate` - (Optional) Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based auth.. (`String`).

`username` - (Required) Username to connect to the cluster/SVM (`String`).

`volume_defaults` - (Optional) List of QoS volume defaults types. See [Netapp Backend Ontap San Volume Defaults ](#netapp-backend-ontap-san-volume-defaults) below for details.

### Bgp Config Peers

BGP parameters for peer.

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

Configure Bond Devices for this App Stack site.

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

### Chap Choice No Chap

CHAP disabled.

### Chap Choice Use Chap

Device NetApp Backend ONTAP SAN CHAP configuration options for enabled CHAP.

`chap_initiator_secret` - (Optional) CHAP initiator secret. Required if useCHAP=true. See [Use Chap Chap Initiator Secret ](#use-chap-chap-initiator-secret) below for details.

`chap_target_initiator_secret` - (Optional) CHAP target initiator secret. Required if useCHAP=true. See [Use Chap Chap Target Initiator Secret ](#use-chap-chap-target-initiator-secret) below for details.

`chap_target_username` - (Optional) Target username. Required if useCHAP=true (`String`).

`chap_username` - (Optional) Inbound username. Required if useCHAP=true (`String`).

### Chap Initiator Secret Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Chap Target Initiator Secret Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Client Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

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

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Dc Cluster Group Choice No Dc Cluster Group

This site is not a member of dc cluster group.

### Dc Cluster Group Connectivity Interface Choice Dc Cluster Group Connectivity Interface Disabled

Do not use this interface to connect to DC Cluster Group peers. .

### Dc Cluster Group Connectivity Interface Choice Dc Cluster Group Connectivity Interface Enabled

Use this interface to connect to DC Cluster Group peers..

### Device Choice Custom Storage

Device configuration for Custom Storage.

### Device Choice Custom Storage

Storage configuration for Custom Storage.

`yaml` - (Optional) K8s YAML for StorageClass (`String`).

### Device Choice Hpe Storage

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

### Device Choice Hpe Storage

Device configuration for HPE Storage.

`api_server_port` - (Optional) Enter Storage Server Port (`Int`).

`csi_version` - (Optional) CSI Version (`String`).(Deprecated)

`iscsi_chap_password` - (Optional) chap Password to connect to the HPE storage. See [Hpe Storage Iscsi Chap Password ](#hpe-storage-iscsi-chap-password) below for details.

`iscsi_chap_user` - (Optional) chap Username to connect to the HPE storage (`String`).

`log_level` - (Optional) logLevel of CSI (`String`).(Deprecated)

`password` - (Optional) Please Enter you password.. See [Hpe Storage Password ](#hpe-storage-password) below for details.

`storage_server_ip_address` - (Optional) Enter storage server IP address (`String`).

`storage_server_name` - (Optional) Enter storage server Name (`String`).

`username` - (Required) Username to connect to the HPE storage management IP (`String`).

### Device Choice Netapp Trident

Storage class Device configuration for NetApp Trident.

`selector` - (Optional) The volume will have the aspects defined in the chosen virtual pool. (`String`).

`storage_pools` - (Optional) The storagePools parameter is used to further restrict the set of pools that match any specified attributes (`String`).

### Device Choice Netapp Trident

Device configuration for NetApp Trident.

###### One of the arguments from this list "netapp_backend_ontap_nas, netapp_backend_ontap_san" must be set

`netapp_backend_ontap_nas` - (Optional) Backend configuration for ONTAP NAS. See [Backend Choice Netapp Backend Ontap Nas ](#backend-choice-netapp-backend-ontap-nas) below for details.

`netapp_backend_ontap_san` - (Optional) Backend configuration for ONTAP SAN. See [Backend Choice Netapp Backend Ontap San ](#backend-choice-netapp-backend-ontap-san) below for details.

### Device Choice Pure Service Orchestrator

Storage class Device configuration for Pure Service Orchestrator.

`backend` - (Optional) The volume will have the aspects defined in the chosen virtual pool. (`String`).

`bandwidth_limit` - (Optional) Valid unit symbols are K, M, G, representing KiB, MiB, and GiB. (`String`).

`iops_limit` - (Optional) Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not defined. (`Int`).

### Device Choice Pure Service Orchestrator

Device configuration for Pure Storage Service Orchestrator.

`arrays` - (Required) This section configure PSO storage arrays. See [Pure Service Orchestrator Arrays ](#pure-service-orchestrator-arrays) below for details.

`cluster_id` - (Required) characters allowed: alphanumeric and underscores (`String`).

`enable_storage_topology` - (Optional) This option is to enable/disable the csi topology feature for pso-csi (`Bool`).

`enable_strict_topology` - (Optional) This option is to enable/disable the strict csi topology feature for pso-csi (`Bool`).

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

### Enable Disable Choice Disable Interception

Disable Interception.

### Enable Disable Choice Enable Interception

Enable Interception.

### Flash Array Flash Arrays

For FlashArrays you must set the "mgmt_endpoint" and "api_token".

`api_token` - (Optional) Please Enter API TOken. Token to connect to management interface. See [Flash Arrays Api Token ](#flash-arrays-api-token) below for details.

`labels` - (Optional) The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner. (`String`).

###### One of the arguments from this list "mgmt_dns_name, mgmt_ip" must be set

`mgmt_dns_name` - (Optional) Management Endpoint's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

`mgmt_ip` - (Optional) Management Endpoint is reachable at the given ip address (`String`).

### Flash Arrays Api Token

Please Enter API TOken. Token to connect to management interface.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Api Token Blindfold Secret Info Internal ](#api-token-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Flash Blade Flash Blades

For FlashBlades you must set the "mgmt_endpoint", "api_token" and nfs_endpoint.

`api_token` - (Optional) Please Enter API TOken. Token to connect to management interface. See [Flash Blades Api Token ](#flash-blades-api-token) below for details.

`lables` - (Optional) The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner. (`String`).

###### One of the arguments from this list "mgmt_dns_name, mgmt_ip" must be set

`mgmt_dns_name` - (Optional) Management Endpoint's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

`mgmt_ip` - (Optional) Management Endpoint is reachable at the given ip address (`String`).

###### One of the arguments from this list "nfs_endpoint_dns_name, nfs_endpoint_ip" must be set

`nfs_endpoint_dns_name` - (Optional) Endpoint's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

`nfs_endpoint_ip` - (Optional) Endpoint is reachable at the given ip address (`String`).

### Flash Blades Api Token

Please Enter API TOken. Token to connect to management interface.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Api Token Blindfold Secret Info Internal ](#api-token-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Forward Proxy Choice Active Forward Proxy Policies

Enable Forward Proxy for this site and manage policies.

`forward_proxy_policies` - (Required) Ordered List of Forward Proxy Policies active. See [ref](#ref) below for details.

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

### Gpu Choice Enable Vgpu

Enable NVIDIA vGPU hosted on VMware.

`feature_type` - (Required) Set Feature to be enabled (`String`).

`server_address` - (Optional) Set License Server Address (`String`).

`server_port` - (Optional) Set License Server port number (`Int`).

### Hpe Storage Iscsi Chap Password

chap Password to connect to the HPE storage.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Iscsi Chap Password Blindfold Secret Info Internal ](#iscsi-chap-password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Hpe Storage Password

Please Enter you password..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Password Blindfold Secret Info Internal ](#password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Interception Policy Choice Enable For All Domains

Enable interception for all domains.

### Interception Policy Choice Policy

Policy to enable/disable specific domains, with implicit enable all domains.

`interception_rules` - (Required) List of ordered rules to enable or disable for TLS interception. See [Policy Interception Rules ](#policy-interception-rules) below for details.

### Interception Rules Domain Match

Domain value or regular expression to match.

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

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

###### One of the arguments from this list "inside_network, ip_fabric_network, segment_network, site_local_inside_network, site_local_network, srv6_network, storage_network" must be set

`inside_network` - (Optional) Interface belongs to user configured inside network. See [ref](#ref) below for details.(Deprecated)

`ip_fabric_network` - (Optional) Interface belongs to IP Fabric network (`Bool`).(Deprecated)

`segment_network` - (Optional) x-displayName: "Segment". See [ref](#ref) below for details.

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (`Bool`).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (`Bool`).

`srv6_network` - (Optional) Interface belongs to per site srv6 network. See [ref](#ref) below for details.(Deprecated)

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

`interfaces` - (Required) Configure network interfaces for this App Stack site. See [Interface List Interfaces ](#interface-list-interfaces) below for details.

### Interface Choice Loopback Interface

Loopback device..

###### One of the arguments from this list "dhcp_client, dhcp_server, static_ip" must be set

`dhcp_client` - (Optional) Interface gets it IP address from external DHCP server (`Bool`).

`dhcp_server` - (Optional) DHCP Server is configured for this interface. IP for this Interface will be derived from the DHCP Server configuration.. See [Address Choice Dhcp Server ](#address-choice-dhcp-server) below for details.

`static_ip` - (Optional) Interface IP is configured statically. See [Address Choice Static Ip ](#address-choice-static-ip) below for details.

`device` - (Required) Interface configuration for the Loopback Ethernet device (`String`).

###### One of the arguments from this list "no_ipv6_address, static_ipv6_address" can be set

`no_ipv6_address` - (Optional) Interface does not have an IPv6 Address. (`Bool`).

`static_ipv6_address` - (Optional) Interface IP is configured statically. See [Ipv6 Address Choice Static Ipv6 Address ](#ipv6-address-choice-static-ipv6-address) below for details.

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

###### One of the arguments from this list "ip_fabric_network, site_local_inside_network, site_local_network" must be set

`ip_fabric_network` - (Optional) Interface belongs to IP Fabric network (`Bool`).(Deprecated)

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (`Bool`).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (`Bool`).

###### One of the arguments from this list "cluster, node" must be set

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site. (`Bool`).

`node` - (Optional) Configuration will apply to a device on the given node. (`String`).

### Interface Choice Tunnel Interface

Tunnel interface, Ipsec tunnels to other networking devices..

`mtu` - (Optional) When configured, mtu must be between 512 and 16384 (`Int`).

###### One of the arguments from this list "inside_network, site_local_inside_network, site_local_network" must be set

`inside_network` - (Optional) Interface belongs to user configured inside network. See [ref](#ref) below for details.(Deprecated)

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (`Bool`).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (`Bool`).

###### One of the arguments from this list "cluster, node" must be set

`cluster` - (Optional) Configuration will apply to given device on all nodes of the site (`Bool`).(Deprecated)

`node` - (Optional) Configuration will apply to a given device on the given node. (`String`).

`priority` - (Optional) Greater the value, higher the priority (`Int`).

`static_ip` - (Required) Interface IP is configured statically. See [Tunnel Interface Static Ip ](#tunnel-interface-static-ip) below for details.

`tunnel` - (Required) Tunnel Configuration for this Interface. See [ref](#ref) below for details.

### Interface List Interfaces

Configure network interfaces for this App Stack site.

###### One of the arguments from this list "dc_cluster_group_connectivity_interface_disabled, dc_cluster_group_connectivity_interface_enabled" must be set

`dc_cluster_group_connectivity_interface_disabled` - (Optional) Do not use this interface to connect to DC Cluster Group peers. (`Bool`).

`dc_cluster_group_connectivity_interface_enabled` - (Optional) Use this interface to connect to DC Cluster Group peers. (`Bool`).

`description` - (Optional) Description for this Interface (`String`).

###### One of the arguments from this list "dedicated_interface, dedicated_management_interface, ethernet_interface, loopback_interface, tunnel_interface" must be set

`dedicated_interface` - (Optional) Networking configuration for dedicated interface is configured locally on site e.g. (outside/inside)Ethernet. See [Interface Choice Dedicated Interface ](#interface-choice-dedicated-interface) below for details.

`dedicated_management_interface` - (Optional) Fallback management interfaces can be made into dedicated management interface. See [Interface Choice Dedicated Management Interface ](#interface-choice-dedicated-management-interface) below for details.

`ethernet_interface` - (Optional) Ethernet interface configuration.. See [Interface Choice Ethernet Interface ](#interface-choice-ethernet-interface) below for details.

`loopback_interface` - (Optional) Loopback device.. See [Interface Choice Loopback Interface ](#interface-choice-loopback-interface) below for details.(Deprecated)

`tunnel_interface` - (Optional) Tunnel interface, Ipsec tunnels to other networking devices.. See [Interface Choice Tunnel Interface ](#interface-choice-tunnel-interface) below for details.

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

###### One of the arguments from this list "cluster_static_ip, fleet_static_ip, node_static_ip" must be set

`cluster_static_ip` - (Optional) Static IP configuration for a specific node. See [Network Prefix Choice Cluster Static Ip ](#network-prefix-choice-cluster-static-ip) below for details.

`fleet_static_ip` - (Optional) Static IP configuration for the fleet. See [Network Prefix Choice Fleet Static Ip ](#network-prefix-choice-fleet-static-ip) below for details.(Deprecated)

`node_static_ip` - (Optional) Static IP configuration for the Node. See [Network Prefix Choice Node Static Ip ](#network-prefix-choice-node-static-ip) below for details.

### Iscsi Chap Password Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

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

### Local Control Plane Bgp Config

BGP configuration for local control plane.

`asn` - (Required) Autonomous System Number (`Int`).

`peers` - (Optional) BGP parameters for peer. See [Bgp Config Peers ](#bgp-config-peers) below for details.

### Local Control Plane Choice Local Control Plane

Site Local control plane is enabled.

`bgp_config` - (Optional) BGP configuration for local control plane. See [Local Control Plane Bgp Config ](#local-control-plane-bgp-config) below for details.

###### One of the arguments from this list "inside_vn, outside_vn" must be set

`inside_vn` - (Optional) Local control plane will work on inside network (`Bool`).

`outside_vn` - (Optional) Local control plane will work on outside network (`Bool`).

### Local Dns Choice First Address

First usable address from the network prefix is chosen as dns server.

### Local Dns Choice Last Address

Last usable address from the network prefix is chosen as dns server.

### Monitoring Choice Monitor

Link Quality Monitoring parameters. Choosing the option will enable link quality monitoring..

### Monitoring Choice Monitor Disabled

Link quality monitoring disabled on the interface..

### Netapp Backend Ontap Nas Auto Export Cidrs

List of CIDRs to filter Kubernetes’ node IPs against when autoExportPolicy is enabled.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Netapp Backend Ontap Nas Client Private Key

Please Enter value of client private key. Used for certificate-based auth..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Client Private Key Blindfold Secret Info Internal ](#client-private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Netapp Backend Ontap Nas Password

Please Enter you password. Password to connect to the cluster/SVM.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Password Blindfold Secret Info Internal ](#password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Netapp Backend Ontap Nas Storage

List of Virtual Storage Pool definitions which are referred back by Storage Class label match selection..

`labels` - (Optional) List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match selection. (`String`).

`volume_defaults` - (Optional) List of QoS volume default types. See [Storage Volume Defaults ](#storage-volume-defaults) below for details.

`zone` - (Optional) Virtual Storage Pool zone definition. (`String`).

### Netapp Backend Ontap Nas Volume Defaults

List of QoS volume defaults types.

`encryption` - (Optional) Enable NetApp volume encryption. (`Bool`).

`export_policy` - (Optional) Export policy to use. (`String`).

###### One of the arguments from this list "adaptive_qos_policy, no_qos, qos_policy" must be set

`adaptive_qos_policy` - (Optional) Enter Adaptive QoS Policy Name (`String`).

`no_qos` - (Optional) No QoS configured (`Bool`).

`qos_policy` - (Optional) Enter QoS Policy Name (`String`).

`security_style` - (Optional) Security style for new volumes. (`String`).

`snapshot_dir` - (Optional) Access to the .snapshot directory. (`Bool`).

`snapshot_policy` - (Optional) Snapshot policy to use (`String`).

`snapshot_reserve` - (Optional) Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else "" (`String`).

`space_reserve` - (Optional) Space reservation mode; “none” (thin) or “volume” (thick) (`String`).

`split_on_clone` - (Optional) Split a clone from its parent upon creation. (`Bool`).

`tiering_policy` - (Optional) Tiering policy to use. "none" is default. (`String`).

`unix_permissions` - (Optional) Unix permission mode for new volumes. All allowed 777 (`Int`).

### Netapp Backend Ontap San Client Private Key

Please Enter value of client private key. Used for certificate-based auth..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Client Private Key Blindfold Secret Info Internal ](#client-private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Netapp Backend Ontap San Password

Please Enter you password. Password to connect to the cluster/SVM.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Password Blindfold Secret Info Internal ](#password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Netapp Backend Ontap San Storage

List of Virtual Storage Pool definitions which are referred back by Storage Class label match selection..

`labels` - (Optional) List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match selection. (`String`).

`volume_defaults` - (Optional) List of QoS volume default types. See [Storage Volume Defaults ](#storage-volume-defaults) below for details.

`zone` - (Optional) Virtual Storage Pool zone definition. (`String`).

### Netapp Backend Ontap San Volume Defaults

List of QoS volume defaults types.

`encryption` - (Optional) Enable NetApp volume encryption. (`Bool`).

`export_policy` - (Optional) Export policy to use. (`String`).

###### One of the arguments from this list "adaptive_qos_policy, no_qos, qos_policy" must be set

`adaptive_qos_policy` - (Optional) Enter Adaptive QoS Policy Name (`String`).

`no_qos` - (Optional) No QoS configured (`Bool`).

`qos_policy` - (Optional) Enter QoS Policy Name (`String`).

`security_style` - (Optional) Security style for new volumes. (`String`).

`snapshot_dir` - (Optional) Access to the .snapshot directory. (`Bool`).

`snapshot_policy` - (Optional) Snapshot policy to use (`String`).

`snapshot_reserve` - (Optional) Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else "" (`String`).

`space_reserve` - (Optional) Space reservation mode; “none” (thin) or “volume” (thick) (`String`).

`split_on_clone` - (Optional) Split a clone from its parent upon creation. (`Bool`).

`tiering_policy` - (Optional) Tiering policy to use. "none" is default. (`String`).

`unix_permissions` - (Optional) Unix permission mode for new volumes. All allowed 777 (`Int`).

### Network Cfg Choice Custom Network Config

Use custom networking configuration.

`bgp_peer_address` - (Optional) to fetch BGP peer address from site Object. This can be used to change peer address per site in fleet. (`String`).

`bgp_peer_address_v6` - (Optional) to fetch BGP peer IPv6 address from site Object. This can be used to change peer IPv6 address per site in fleet. (`String`).

`bgp_router_id` - (Optional) fetch BGP router ID from site object. (`String`).

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

`outside_nameserver` - (Optional) Optional DNS server V4 IP to be used for name resolution in local network (`String`).

`outside_nameserver_v6` - (Optional) Optional DNS server V6 IP to be used for name resolution in local network (`String`).

`outside_vip` - (Optional) Optional common virtual V4 IP across all nodes to be used as automatic VIP for site local network. (`String`).

`outside_vip_v6` - (Optional) Optional common virtual V6 IP across all nodes to be used as automatic VIP for site local network. (`String`).

###### One of the arguments from this list "site_to_site_tunnel_ip, sm_connection_public_ip, sm_connection_pvt_ip" must be set

`site_to_site_tunnel_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (`String`).

`sm_connection_public_ip` - (Optional) which are part of the site mesh group (`Bool`).

`sm_connection_pvt_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (`Bool`).

###### One of the arguments from this list "default_sli_config, sli_config" can be set

`default_sli_config` - (Optional) Use default configuration for site local network (`Bool`).(Deprecated)

`sli_config` - (Optional) Configuration for site local inside network. See [Sli Choice Sli Config ](#sli-choice-sli-config) below for details.(Deprecated)

###### One of the arguments from this list "default_config, slo_config" must be set

`default_config` - (Optional) Use default configuration for site local network (`Bool`).

`slo_config` - (Optional) Configuration for site local network. See [Slo Choice Slo Config ](#slo-choice-slo-config) below for details.

`tunnel_dead_timeout` - (Optional) When not set (== 0), a default value of 10000 msec will be used. (`Int`).

`vip_vrrp_mode` - (Optional) When Outside VIP / Inside VIP are configured, it is recommended to turn on vrrp and also configure BGP. (`String`).

### Network Choice Inside Vn

Local control plane will work on inside network.

### Network Choice Ip Fabric Network

Interface belongs to IP Fabric network.

### Network Choice Outside Vn

Local control plane will work on outside network.

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

x-displayName: "Enabled".

### Offline Survivability Mode Choice No Offline Survivability Mode

x-displayName: "Disabled".

### Operating System Version Choice Default Os Version

Will assign latest available OS version.

### Password Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

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

### Pure Service Orchestrator Arrays

This section configure PSO storage arrays.

`flash_array` - (Optional) For FlashArrays you must set the "mgmt_endpoint" and "api_token". See [Arrays Flash Array ](#arrays-flash-array) below for details.

`flash_blade` - (Optional) Specify what storage flash blades should be managed the plugin. See [Arrays Flash Blade ](#arrays-flash-blade) below for details.

### Qos Policy Choice No Qos

No QoS configured.

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

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

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

###### One of the arguments from this list "no_static_routes, static_routes" must be set

`no_static_routes` - (Optional) Static Routes disabled for site local inside network. (`Bool`).

`static_routes` - (Optional) Manage static routes for site local inside network.. See [Static Route Choice Static Routes ](#static-route-choice-static-routes) below for details.

###### One of the arguments from this list "no_v6_static_routes, static_v6_routes" must be set

`no_v6_static_routes` - (Optional) Static IPv6 Routes disabled for site local inside network. (`Bool`).

`static_v6_routes` - (Optional) Manage IPv6 static routes for site local inside network.. See [Static V6 Route Choice Static V6 Routes ](#static-v6-route-choice-static-v6-routes) below for details.

### Slo Choice Default Config

Use default configuration for site local network.

### Slo Choice Slo Config

Configuration for site local network.

###### One of the arguments from this list "dc_cluster_group, no_dc_cluster_group" must be set

`dc_cluster_group` - (Optional) This site is member of dc cluster group via network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (`Bool`).

`dc_cluster_group_interface` - (Optional) This App Stack is member of dc cluster group and connected to network over this interface. By default it takes default gateway interface.. See [ref](#ref) below for details.(Deprecated)

`labels` - (Optional) Add Labels for this network, these labels can be used in firewall policy (`String`).

###### One of the arguments from this list "no_static_routes, static_routes" must be set

`no_static_routes` - (Optional) Static Routes disabled for site local network. (`Bool`).

`static_routes` - (Optional) Manage static routes for site local network.. See [Static Route Choice Static Routes ](#static-route-choice-static-routes) below for details.

###### One of the arguments from this list "no_static_v6_routes, static_v6_routes" must be set

`no_static_v6_routes` - (Optional) Static IPv6 Routes disabled for site local network. (`Bool`).

`static_v6_routes` - (Optional) Manage static IPv6 routes for site local network.. See [Static V6 Route Choice Static V6 Routes ](#static-v6-route-choice-static-v6-routes) below for details.

### Sriov Interface Choice Sriov Interfaces

Use custom Single Root I/O Virtualization interfaces.

`sriov_interface` - (Optional) Use custom SR-IOV interfaces Configuration. See [Sriov Interfaces Sriov Interface ](#sriov-interfaces-sriov-interface) below for details.

### Sriov Interfaces Sriov Interface

Use custom SR-IOV interfaces Configuration.

`interface_name` - (Required) Name of SR-IOV physical interface (`String`).

`number_of_vfio_vfs` - (Optional) Number of virtual functions reserved for VNFs and DPDK-based CNFs (`Int`).

`number_of_vfs` - (Required) Total number of virtual functions (`Int`).

### Stateful Dhcp Networks

List of networks from which DHCP server can allocate ip addresses.

###### One of the arguments from this list "network_prefix, network_prefix_allocator" must be set

`network_prefix` - (Optional) Network Prefix to be used for IPV6 address auto configuration (`String`).

`network_prefix_allocator` - (Optional) Prefix length from address allocator scheme is used to calculate offsets. See [ref](#ref) below for details.(Deprecated)

`pool_settings` - (Required) Controls how DHCPV6 pools are handled (`String`).

`pools` - (Optional) List of non overlapping ip address ranges.. See [Dhcp Networks Pools ](#dhcp-networks-pools) below for details.

### Static Route Choice No Static Routes

Static Routes disabled for site local inside network..

### Static Route Choice Static Routes

Manage static routes for site local inside network..

`static_routes` - (Required) List of static routes. See [Static Routes Static Routes ](#static-routes-static-routes) below for details.

### Static Routes Static Routes

List of static routes.

`attrs` - (Optional) List of attributes that control forwarding, dynamic routing and control plane (host) reachability (`List of Strings`).

`ip_prefixes` - (Required) List of route prefixes that have common next hop and attributes (`String`).

###### One of the arguments from this list "default_gateway, interface, ip_address" must be set

`default_gateway` - (Optional) Traffic matching the ip prefixes is sent to the default gateway (`Bool`).

`interface` - (Optional) Traffic matching the ip prefixes is sent to this interface. See [ref](#ref) below for details.

`ip_address` - (Optional) Traffic matching the ip prefixes is sent to this IP Address (`String`).

### Static V6 Route Choice No Static V6 Routes

Static IPv6 Routes disabled for site local network..

### Static V6 Route Choice No V6 Static Routes

Static IPv6 Routes disabled for site local inside network..

### Static V6 Route Choice Static V6 Routes

Manage IPv6 static routes for site local inside network..

`static_routes` - (Required) List of IPv6 static routes. See [Static V6 Routes Static Routes ](#static-v6-routes-static-routes) below for details.

### Static V6 Routes Static Routes

List of IPv6 static routes.

`attrs` - (Optional) List of attributes that control forwarding, dynamic routing and control plane (host) reachability (`List of Strings`).

`ip_prefixes` - (Required) List of IPv6 route prefixes that have common next hop and attributes (`String`).

###### One of the arguments from this list "default_gateway, interface, ip_address" must be set

`default_gateway` - (Optional) Traffic matching the ip prefixes is sent to the default gateway (`Bool`).

`interface` - (Optional) Traffic matching the ip prefixes is sent to this interface. See [ref](#ref) below for details.

`ip_address` - (Optional) Traffic matching the ip prefixes is sent to this IP Address (`String`).

### Storage Volume Defaults

List of QoS volume default types.

`encryption` - (Optional) Enable NetApp volume encryption. (`Bool`).

`export_policy` - (Optional) Export policy to use. (`String`).

###### One of the arguments from this list "adaptive_qos_policy, no_qos, qos_policy" must be set

`adaptive_qos_policy` - (Optional) Enter Adaptive QoS Policy Name (`String`).

`no_qos` - (Optional) No QoS configured (`Bool`).

`qos_policy` - (Optional) Enter QoS Policy Name (`String`).

`security_style` - (Optional) Security style for new volumes. (`String`).

`snapshot_dir` - (Optional) Access to the .snapshot directory. (`Bool`).

`snapshot_policy` - (Optional) Snapshot policy to use (`String`).

`snapshot_reserve` - (Optional) Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else "" (`String`).

`space_reserve` - (Optional) Space reservation mode; “none” (thin) or “volume” (thick) (`String`).

`split_on_clone` - (Optional) Split a clone from its parent upon creation. (`Bool`).

`tiering_policy` - (Optional) Tiering policy to use. "none" is default. (`String`).

`unix_permissions` - (Optional) Unix permission mode for new volumes. All allowed 777 (`Int`).

### Storage Cfg Choice Custom Storage Config

Use custom storage configuration.

###### One of the arguments from this list "no_static_routes, static_routes" must be set

`no_static_routes` - (Optional) Static Routes not required for storage network. (`Bool`).

`static_routes` - (Optional) Manage static routes for storage network.. See [Static Route Choice Static Routes ](#static-route-choice-static-routes) below for details.

###### One of the arguments from this list "default_storage_class, storage_class_list" must be set

`default_storage_class` - (Optional) Use only default storage class in kubernetes (`Bool`).

`storage_class_list` - (Optional) Add additional custom storage classes in kubernetes. See [Storage Class Choice Storage Class List ](#storage-class-choice-storage-class-list) below for details.

###### One of the arguments from this list "no_storage_device, storage_device_list" must be set

`no_storage_device` - (Optional) This site does not have any storage devices (`Bool`).

`storage_device_list` - (Optional) Add all storage devices belonging to this site. See [Storage Device Choice Storage Device List ](#storage-device-choice-storage-device-list) below for details.

###### One of the arguments from this list "no_storage_interfaces, storage_interface_list" must be set

`no_storage_interfaces` - (Optional) This site does not have any storage interfaces (`Bool`).

`storage_interface_list` - (Optional) Add all storage interfaces belonging to this site. See [Storage Interface Choice Storage Interface List ](#storage-interface-choice-storage-interface-list) below for details.

### Storage Class Choice Default Storage Class

Use only default storage class in kubernetes.

### Storage Class Choice Storage Class List

Add additional custom storage classes in kubernetes.

`storage_classes` - (Optional) List of custom storage classes. See [Storage Class List Storage Classes ](#storage-class-list-storage-classes) below for details.

### Storage Class List Storage Classes

List of custom storage classes.

`advanced_storage_parameters` - (Optional) Map of parameter name and string value (`String`).

`allow_volume_expansion` - (Optional) Allow volume expansion. (`Bool`).

`default_storage_class` - (Optional) Make this storage class default storage class for the K8s cluster (`Bool`).

`description` - (Optional) Description for this storage class (`String`).

###### One of the arguments from this list "custom_storage, hpe_storage, netapp_trident, pure_service_orchestrator" must be set

`custom_storage` - (Optional) Storage configuration for Custom Storage. See [Device Choice Custom Storage ](#device-choice-custom-storage) below for details.

`hpe_storage` - (Optional) Storage configuration for HPE Storage. See [Device Choice Hpe Storage ](#device-choice-hpe-storage) below for details.

`netapp_trident` - (Optional) Storage class Device configuration for NetApp Trident. See [Device Choice Netapp Trident ](#device-choice-netapp-trident) below for details.

`pure_service_orchestrator` - (Optional) Storage class Device configuration for Pure Service Orchestrator. See [Device Choice Pure Service Orchestrator ](#device-choice-pure-service-orchestrator) below for details.

`reclaim_policy` - (Optional) Reclaim Policy (`String`).

`storage_class_name` - (Required) Name of the storage class as it will appear in K8s. (`String`).

`storage_device` - (Required) Storage device that this class will use. The Device name defined at previous step. (`String`).

### Storage Device Choice No Storage Device

This site does not have any storage devices.

### Storage Device Choice Storage Device List

Add all storage devices belonging to this site.

`storage_devices` - (Optional) List of custom storage devices. See [Storage Device List Storage Devices ](#storage-device-list-storage-devices) below for details.

### Storage Device List Storage Devices

List of custom storage devices.

`advanced_advanced_parameters` - (Optional) Map of parameter name and string value (`String`).

###### One of the arguments from this list "custom_storage, hpe_storage, netapp_trident, pure_service_orchestrator" must be set

`custom_storage` - (Optional) Device configuration for Custom Storage (`Bool`).

`hpe_storage` - (Optional) Device configuration for HPE Storage. See [Device Choice Hpe Storage ](#device-choice-hpe-storage) below for details.

`netapp_trident` - (Optional) Device configuration for NetApp Trident. See [Device Choice Netapp Trident ](#device-choice-netapp-trident) below for details.

`pure_service_orchestrator` - (Optional) Device configuration for Pure Storage Service Orchestrator. See [Device Choice Pure Service Orchestrator ](#device-choice-pure-service-orchestrator) below for details.

`storage_device` - (Required) Storage device and device unit (`String`).

### Storage Interface Choice No Storage Interfaces

This site does not have any storage interfaces.

### Storage Interface Choice Storage Interface List

Add all storage interfaces belonging to this site.

`storage_interfaces` - (Required) Configure storage interfaces for this App Stack site. See [Storage Interface List Storage Interfaces ](#storage-interface-list-storage-interfaces) below for details.

### Storage Interface List Storage Interfaces

Configure storage interfaces for this App Stack site.

`description` - (Optional) Description for this Interface (`String`).

`labels` - (Optional) Add Labels for this Interface, these labels can be used in firewall policy (`String`).

`storage_interface` - (Required) Configure storage interface for this App Stack site. See [Storage Interfaces Storage Interface ](#storage-interfaces-storage-interface) below for details.

### Storage Interfaces Storage Interface

Configure storage interface for this App Stack site.

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

###### One of the arguments from this list "inside_network, ip_fabric_network, segment_network, site_local_inside_network, site_local_network, srv6_network, storage_network" must be set

`inside_network` - (Optional) Interface belongs to user configured inside network. See [ref](#ref) below for details.(Deprecated)

`ip_fabric_network` - (Optional) Interface belongs to IP Fabric network (`Bool`).(Deprecated)

`segment_network` - (Optional) x-displayName: "Segment". See [ref](#ref) below for details.

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (`Bool`).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (`Bool`).

`srv6_network` - (Optional) Interface belongs to per site srv6 network. See [ref](#ref) below for details.(Deprecated)

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

### Tunnel Interface Static Ip

Interface IP is configured statically.

###### One of the arguments from this list "cluster_static_ip, fleet_static_ip, node_static_ip" must be set

`cluster_static_ip` - (Optional) Static IP configuration for a specific node. See [Network Prefix Choice Cluster Static Ip ](#network-prefix-choice-cluster-static-ip) below for details.

`fleet_static_ip` - (Optional) Static IP configuration for the fleet. See [Network Prefix Choice Fleet Static Ip ](#network-prefix-choice-fleet-static-ip) below for details.(Deprecated)

`node_static_ip` - (Optional) Static IP configuration for the Node. See [Network Prefix Choice Node Static Ip ](#network-prefix-choice-node-static-ip) below for details.

### Use Chap Chap Initiator Secret

CHAP initiator secret. Required if useCHAP=true.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Chap Initiator Secret Blindfold Secret Info Internal ](#chap-initiator-secret-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Use Chap Chap Target Initiator Secret

CHAP target initiator secret. Required if useCHAP=true.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Chap Target Initiator Secret Blindfold Secret Info Internal ](#chap-target-initiator-secret-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Vega Upgrade Mode Toggle Choice Disable Vega Upgrade Mode

Disable Vega Upgrade Mode.

### Vega Upgrade Mode Toggle Choice Enable Vega Upgrade Mode

When enabled, vega will inform RE to stop traffic to the specific node..

### Vlan Choice Untagged

Configure a untagged ethernet interface.

### Vm Choice Enable Vm

VMs support is enabled for this Site.

### Volterra Sw Version Choice Default Sw Version

Will assign latest available F5XC Software Version.

Attribute Reference
-------------------

-	`id` - This is the id of the configured voltstack_site.
