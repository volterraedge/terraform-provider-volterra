---

page_title: "Volterra: fleet"

description: "The fleet allows CRUD of Fleet resource on Volterra SaaS"
-----------------------------------------------------------------------

Resource volterra_fleet
=======================

The Fleet allows CRUD of Fleet resource on Volterra SaaS

~> **Note:** Please refer to [Fleet API docs](https://docs.cloud.f5.com/docs-v2/api/fleet) to learn more

Example Usage
-------------

```hcl
resource "volterra_fleet" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "bond_device_list no_bond_devices" must be set

  bond_device_list {
    bond_devices {
      devices = ["eth0"]

      // One of the arguments from this list "active_backup lacp" must be set

      lacp {
        rate = "30"
      }
      link_polling_interval = "1000"
      link_up_delay = "200"
      name = "bond0"
    }
  }

  // One of the arguments from this list "dc_cluster_group dc_cluster_group_inside no_dc_cluster_group" must be set

  no_dc_cluster_group = true
  fleet_label = ["sfo"]

  // One of the arguments from this list "disable_gpu enable_gpu enable_vgpu" must be set

  disable_gpu = true

  // One of the arguments from this list "default_config device_list interface_list" must be set

  interface_list {
    interfaces {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  // One of the arguments from this list "log_receiver logs_streaming_disabled" must be set

  logs_streaming_disabled = true

  // One of the arguments from this list "default_sriov_interface sriov_interfaces" must be set

  default_sriov_interface = true

  // One of the arguments from this list "default_storage_class storage_class_list" must be set

  default_storage_class = true

  // One of the arguments from this list "no_storage_device storage_device_list" must be set

  no_storage_device = true

  // One of the arguments from this list "no_storage_interfaces storage_interface_list" must be set

  no_storage_interfaces = true

  // One of the arguments from this list "no_storage_static_routes storage_static_routes" must be set

  no_storage_static_routes = true

  // One of the arguments from this list "allow_all_usb deny_all_usb usb_policy" must be set

  deny_all_usb = true
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

`blocked_services` - (Optional) Disable node local services on this site.. See [Blocked Services ](#blocked-services) below for details.

###### One of the arguments from this list "bond_device_list, no_bond_devices" must be set

`bond_device_list` - (Optional) Configure Bond Devices for this fleet. See [Bond Choice Bond Device List ](#bond-choice-bond-device-list) below for details.

`no_bond_devices` - (Optional) No Bond Devices configured for this Fleet (`Bool`).

###### One of the arguments from this list "dc_cluster_group, dc_cluster_group_inside, no_dc_cluster_group" must be set

`dc_cluster_group` - (Optional) This fleet is member of dc cluster group via site local network. See [ref](#ref) below for details.

`dc_cluster_group_inside` - (Optional) This fleet is member of dc cluster group via site local inside network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This fleet is not a member of a DC cluster group (`Bool`).

`enable_default_fleet_config_download` - (Optional) Enable default fleet config, It must be set for storage config and gpu config (`Bool`).

`fleet_label` - (Required) fleet_label with "sfo" will create a known_label "ves.io/fleet=sfo" in tenant for the fleet (`String`).

###### One of the arguments from this list "disable_gpu, enable_gpu, enable_vgpu" must be set

`disable_gpu` - (Optional) GPU is not enabled for this fleet (`Bool`).

`enable_gpu` - (Optional) GPU is enabled for this fleet (`Bool`).

`enable_vgpu` - (Optional) Enable NVIDIA vGPU hosted on VMware. See [Gpu Choice Enable Vgpu ](#gpu-choice-enable-vgpu) below for details.

`inside_virtual_network` - (Optional) Default inside (site local) virtual network for the fleet. See [ref](#ref) below for details.

###### One of the arguments from this list "default_config, device_list, interface_list" must be set

`default_config` - (Optional) Use default configuration for interfaces belonging to this fleet (`Bool`).

`device_list` - (Optional) Add device for all interfaces belonging to this fleet. See [Interface Choice Device List ](#interface-choice-device-list) below for details.

`interface_list` - (Optional) Add all interfaces belonging to this fleet. See [Interface Choice Interface List ](#interface-choice-interface-list) below for details.

`kubernetes_upgrade_drain` - (Optional) Enable Kubernetes Drain during OS or SW upgrade. See [Kubernetes Upgrade Drain ](#kubernetes-upgrade-drain) below for details.

###### One of the arguments from this list "log_receiver, logs_streaming_disabled" must be set

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (`Bool`).

`network_connectors` - (Optional) The network connectors configuration is applied on all sites that are member of the fleet.. See [ref](#ref) below for details.

`network_firewall` - (Optional) The Network Firewall is applied on Virtual Networks of type site local network and site local inside network. See [ref](#ref) below for details.

`operating_system_version` - (Optional) Current Operating System version can be overridden via site config. (`String`).

`outside_virtual_network` - (Optional) Default outside (site local) virtual network for the fleet. See [ref](#ref) below for details.

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode to optimize for L3 or L7 networking. See [Performance Enhancement Mode ](#performance-enhancement-mode) below for details.

###### One of the arguments from this list "default_sriov_interface, sriov_interfaces" must be set

`default_sriov_interface` - (Optional) Disable Single Root I/O Virtualization interfaces (`Bool`).

`sriov_interfaces` - (Optional) Use custom Single Root I/O Virtualization interfaces. See [Sriov Interface Choice Sriov Interfaces ](#sriov-interface-choice-sriov-interfaces) below for details.

###### One of the arguments from this list "default_storage_class, storage_class_list" must be set

`default_storage_class` - (Optional) Use only default storage class in kubernetes (`Bool`).

`storage_class_list` - (Optional) Add additional custom storage classes in kubernetes for this fleet. See [Storage Class Choice Storage Class List ](#storage-class-choice-storage-class-list) below for details.

###### One of the arguments from this list "no_storage_device, storage_device_list" must be set

`no_storage_device` - (Optional) This fleet does not have any storage devices (`Bool`).

`storage_device_list` - (Optional) Add all storage devices belonging to this fleet. See [Storage Device Choice Storage Device List ](#storage-device-choice-storage-device-list) below for details.

###### One of the arguments from this list "no_storage_interfaces, storage_interface_list" must be set

`no_storage_interfaces` - (Optional) This fleet does not have any storage interfaces (`Bool`).

`storage_interface_list` - (Optional) Add all storage interfaces belonging to this fleet. See [Storage Interface Choice Storage Interface List ](#storage-interface-choice-storage-interface-list) below for details.

###### One of the arguments from this list "no_storage_static_routes, storage_static_routes" must be set

`no_storage_static_routes` - (Optional) This fleet does not have any storage static routes (`Bool`).

`storage_static_routes` - (Optional) Add all storage storage static routes. See [Storage Static Routes Choice Storage Static Routes ](#storage-static-routes-choice-storage-static-routes) below for details.

###### One of the arguments from this list "allow_all_usb, deny_all_usb, usb_policy" must be set

`allow_all_usb` - (Optional) All USB devices are allowed (`Bool`).

`deny_all_usb` - (Optional) All USB devices are denied (`Bool`).

`usb_policy` - (Optional) Allow only specific USB devices. See [ref](#ref) below for details.

###### One of the arguments from this list "disable_vm, enable_vm" can be set

`disable_vm` - (Optional) VMs support is not enabled for this fleet (`Bool`).

`enable_vm` - (Optional) VMs support is enabled for this fleet. See [Vm Choice Enable Vm ](#vm-choice-enable-vm) below for details.

`volterra_software_version` - (Optional) Current software installed can be overridden via site config. (`String`).

### Blocked Services

Disable node local services on this site..

###### One of the arguments from this list "dns, ssh, web_user_interface" can be set

`dns` - (Optional) Matches DNS port 53 (`Bool`).

`ssh` - (Optional) x-displayName: "SSH" (`Bool`).

`web_user_interface` - (Optional) x-displayName: "Web UI" (`Bool`).

`network_type` - (Optional) Site Local VRF on which this service will be disabled (`String`).

### Kubernetes Upgrade Drain

Enable Kubernetes Drain during OS or SW upgrade.

###### One of the arguments from this list "disable_upgrade_drain, enable_upgrade_drain" must be set

`disable_upgrade_drain` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_upgrade_drain` - (Optional) x-displayName: "Enable". See [Kubernetes Upgrade Drain Enable Choice Enable Upgrade Drain ](#kubernetes-upgrade-drain-enable-choice-enable-upgrade-drain) below for details.

### Performance Enhancement Mode

Performance Enhancement Mode to optimize for L3 or L7 networking.

###### One of the arguments from this list "perf_mode_l3_enhanced, perf_mode_l7_enhanced" must be set

`perf_mode_l3_enhanced` - (Optional) Site optimized for L3 traffic processing. See [Perf Mode Choice Perf Mode L3 Enhanced ](#perf-mode-choice-perf-mode-l3-enhanced) below for details.

`perf_mode_l7_enhanced` - (Optional) Site optimized for L7 traffic processing (`Bool`).

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

### Blocked Services Value Type Choice Dns

Matches DNS port 53.

### Blocked Services Value Type Choice Ssh

x-displayName: "SSH".

### Blocked Services Value Type Choice Web User Interface

x-displayName: "Web UI".

### Bond Choice Bond Device List

Configure Bond Devices for this fleet.

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

`iscsi_chap_password` - (Optional) chap Password to connect to the HPE storage. See [Hpe Storage Iscsi Chap Password ](#hpe-storage-iscsi-chap-password) below for details.

`iscsi_chap_user` - (Optional) chap Username to connect to the HPE storage (`String`).

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

### Device Instance Network Device

Device instance is a networking device like ethernet.

`interface` - (Required) if use is NETWORK_INTERFACE_USE_INSIDE, the virtual-network must of type VIRTUAL_NETWORK_SITE_LOCAL_INSIDE. See [ref](#ref) below for details.

`use` - (Required) Specifies if the network interface is connected to inside network or outside network. (`String`).

### Device List Devices

device instance specific sections.

###### One of the arguments from this list "network_device" must be set

`network_device` - (Optional) Device instance is a networking device like ethernet. See [Device Instance Network Device ](#device-instance-network-device) below for details.

`name` - (Optional) Name of the device including the unit number (e.g. eth0 or disk1). The name must match name of device in host-os of node (`String`).

`owner` - (Required) This option is not yet supported (`String`).

### Flash Array Flash Arrays

For FlashArrays you must set the "mgmt_endpoint" and "api_token".

`api_token` - (Optional) Please Enter API TOken. Token to connect to management interface. See [Flash Arrays Api Token ](#flash-arrays-api-token) below for details.

`labels` - (Optional) The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner. (`String`).

###### One of the arguments from this list "mgmt_dns_name, mgmt_ip" must be set

`mgmt_dns_name` - (Optional) Management Endpoint's ip address is discovered using DNS name resolution. The name given here is fully qualified domain name. (`String`).

`mgmt_ip` - (Optional) Management Endpoint is reachable at the given ip address (`String`).

### Flash Arrays Api Token

Please Enter API TOken. Token to connect to management interface.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

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

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Gpu Choice Enable Vgpu

Enable NVIDIA vGPU hosted on VMware.

`feature_type` - (Required) Set Feature to be enabled (`String`).

`server_address` - (Optional) Set License Server Address (`String`).

`server_port` - (Optional) Set License Server port number (`Int`).

### Hpe Storage Iscsi Chap Password

chap Password to connect to the HPE storage.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Hpe Storage Password

Please Enter you password..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Interface Choice Device List

Add device for all interfaces belonging to this fleet.

`devices` - (Optional) device instance specific sections. See [Device List Devices ](#device-list-devices) below for details.

### Interface Choice Interface List

Add all interfaces belonging to this fleet.

`interfaces` - (Required) Add all interfaces belonging to this fleet. See [ref](#ref) below for details.

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

### Netapp Backend Ontap Nas Auto Export Cidrs

List of CIDRs to filter Kubernetes’ node IPs against when autoExportPolicy is enabled.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Netapp Backend Ontap Nas Client Private Key

Please Enter value of client private key. Used for certificate-based auth..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Netapp Backend Ontap Nas Password

Please Enter you password. Password to connect to the cluster/SVM.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

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

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Netapp Backend Ontap San Password

Please Enter you password. Password to connect to the cluster/SVM.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

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

### Nexthop Nexthop Address

Nexthop address when type is "Use-Configured".

###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

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

### Secret Info Oneof Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Info Oneof Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Sriov Interface Choice Sriov Interfaces

Use custom Single Root I/O Virtualization interfaces.

`sriov_interface` - (Optional) Use custom SR-IOV interfaces Configuration. See [Sriov Interfaces Sriov Interface ](#sriov-interfaces-sriov-interface) below for details.

### Sriov Interfaces Sriov Interface

Use custom SR-IOV interfaces Configuration.

`interface_name` - (Required) Name of SR-IOV physical interface (`String`).

`number_of_vfio_vfs` - (Optional) Number of virtual functions reserved for VNFs and DPDK-based CNFs (`Int`).

`number_of_vfs` - (Required) Total number of virtual functions (`Int`).

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

### Storage Class Choice Storage Class List

Add additional custom storage classes in kubernetes for this fleet.

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

### Storage Device Choice Storage Device List

Add all storage devices belonging to this fleet.

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

### Storage Interface Choice Storage Interface List

Add all storage interfaces belonging to this fleet.

`interfaces` - (Required) Add all interfaces belonging to this fleet. See [ref](#ref) below for details.

### Storage Routes Nexthop

Nexthop for the route.

`interface` - (Optional) Nexthop is network interface when type is "Network-Interface". See [ref](#ref) below for details.

`nexthop_address` - (Optional) Nexthop address when type is "Use-Configured". See [Nexthop Nexthop Address ](#nexthop-nexthop-address) below for details.

`type` - (Optional) Identifies the type of next-hop (`String`).

### Storage Routes Subnets

List of route prefixes.

###### One of the arguments from this list "ipv4, ipv6" must be set

`ipv4` - (Optional) IPv4 Subnet Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Subnet Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Storage Static Routes Storage Routes

List of storage static routes.

`attrs` - (Optional) List of route attributes associated with the static route (`List of Strings`).

`labels` - (Optional) Add Labels for this Static Route, these labels can be used in network policy (`String`).

`nexthop` - (Optional) Nexthop for the route. See [Storage Routes Nexthop ](#storage-routes-nexthop) below for details.

`subnets` - (Required) List of route prefixes. See [Storage Routes Subnets ](#storage-routes-subnets) below for details.

### Storage Static Routes Choice Storage Static Routes

Add all storage storage static routes.

`storage_routes` - (Required) List of storage static routes. See [Storage Static Routes Storage Routes ](#storage-static-routes-storage-routes) below for details.

### Use Chap Chap Initiator Secret

CHAP initiator secret. Required if useCHAP=true.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Use Chap Chap Target Initiator Secret

CHAP target initiator secret. Required if useCHAP=true.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

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

### Vm Choice Enable Vm

VMs support is enabled for this fleet.

Attribute Reference
-------------------

-	`id` - This is the id of the configured fleet.
