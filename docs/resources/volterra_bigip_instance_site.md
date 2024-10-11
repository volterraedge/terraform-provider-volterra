---

page_title: "Volterra: bigip_instance_site"
description: "The bigip_instance_site allows CRUD of Bigip Instance Site resource on Volterra SaaS"

---

Resource volterra_bigip_instance_site
=====================================

The Bigip Instance Site allows CRUD of Bigip Instance Site resource on Volterra SaaS

~> **Note:** Please refer to [Bigip Instance Site API docs](https://docs.cloud.f5.com/docs-v2/api/views-bigip-instance-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_bigip_instance_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  central_manager {
    central_manager_site {
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

`central_manager` - (Required) BIG-IP Central manager site managing this BIG-IP instance. See [Central Manager ](#central-manager) below for details.

`node_list` - (Optional) Once a node is created and registers with the site, it will be shown in this section.. See [Node List ](#node-list) below for details.

`volterra_software` - (Optional) Refer to release notes to find required released SW versions.. See [Volterra Software ](#volterra-software) below for details.

### Central Manager

BIG-IP Central manager site managing this BIG-IP instance.

`central_manager_site` - (Optional) BIG-IP Central manager site managing this BIG-IP instance. See [ref](#ref) below for details.

### Node List

Once a node is created and registers with the site, it will be shown in this section..

`hostname` - (Optional) Hostname for this Node (`String`).

`interface_list` - (Optional) Manage interfaces belonging to this node. See [Node List Interface List ](#node-list-interface-list) below for details.

`public_ip` - (Optional) Public IP for this Node (`String`).

`type` - (Optional) Type for this Node, can be Control or Worker (`String`).

### Volterra Software

Refer to release notes to find required released SW versions..

###### One of the arguments from this list "default_sw_version, volterra_software_version" must be set

`default_sw_version` - (Optional) Will assign latest available F5XC Software Version (`Bool`).

`volterra_software_version` - (Optional) Specify a F5XC Software Version to be used e.g. crt-20210329-1002. (`String`).

### Address Choice Dhcp Client

Interface gets it's IP address from an external DHCP server..

### Address Choice Dhcp Server

DHCP Server is configured for this interface, Interface IP is derived from DHCP server configuration..

### Address Choice No Ipv4 Address

Interface does not have an IPv4 Address..

### Address Choice Static Ip

Interface IP address is configured statically..

### Interface Choice Bond Interface

x-displayName: "Bond Interface".

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

### Ipv6 Address Choice Ipv6 Auto Config

Interface IPv6 address will be configured via Auto Configuration..

### Ipv6 Address Choice No Ipv6 Address

Interface does not have an IPv6 Address..

### Ipv6 Address Choice Static Ipv6 Address

Interface IPv6 address is configured statically..

### Monitoring Choice Monitor

x-displayName: "Enabled".

### Monitoring Choice Monitor Disabled

x-displayName: "Disabled".

### Network Choice Site Local Inside Network

x-displayName: "Site Local Inside (Local VRF)".

### Network Choice Site Local Network

x-displayName: "Site Local Outside (Local VRF)".

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

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Site To Site Connectivity Interface Choice Site To Site Connectivity Interface Disabled

Do not use this interface for site to site connectivity..

### Site To Site Connectivity Interface Choice Site To Site Connectivity Interface Enabled

Use this this interface for site to site connectivity..

### Volterra Sw Version Choice Default Sw Version

Will assign latest available F5XC Software Version.

Attribute Reference
-------------------

-	`id` - This is the id of the configured bigip_instance_site.
