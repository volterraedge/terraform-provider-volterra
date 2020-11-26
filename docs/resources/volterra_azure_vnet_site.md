---

page_title: "Volterra: azure_vnet_site"

description: "The azure_vnet_site allows CRUD of Azure Vnet Site resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_azure_vnet_site
=================================

The Azure Vnet Site allows CRUD of Azure Vnet Site resource on Volterra SaaS

~> **Note:** Please refer to [Azure Vnet Site API docs](https://volterra.io/docs/api/views-azure-vnet-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_azure_vnet_site" "example" {
  name         = "acmecorp-web"
  namespace    = "staging"
  azure_region = ["East US"]

  // One of the arguments from this list "assisted azure_cred" must be set

  azure_cred {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }
  resource_group = ["my-resources"]

  // One of the arguments from this list "ingress_egress_gw voltstack_cluster ingress_gw" must be set

  ingress_egress_gw {
    az_nodes {
      azure_az  = "1"
      disk_size = "disk_size"

      inside_subnet {
        // One of the arguments from this list "subnet_param subnet" must be set

        subnet_param {
          ipv4 = "10.1.2.0/24"
          ipv6 = "1234:568:abcd:9100::/64"
        }
      }

      outside_subnet {
        // One of the arguments from this list "subnet subnet_param" must be set

        subnet {
          // One of the arguments from this list "subnet_resource_grp vnet_resource_group" must be set
          subnet_resource_grp = "subnet_resource_grp"

          subnet_name = "MySubnet"
        }
      }
    }

    azure_certified_hw = "azure-byol-multi-nic-voltmesh"

    // One of the arguments from this list "no_forward_proxy active_forward_proxy_policies forward_proxy_allow_all" must be set
    no_forward_proxy = true

    // One of the arguments from this list "no_global_network global_network_list" must be set
    no_global_network = true

    // One of the arguments from this list "no_inside_static_routes inside_static_routes" must be set
    no_inside_static_routes = true

    // One of the arguments from this list "no_network_policy active_network_policies" must be set
    no_network_policy = true

    // One of the arguments from this list "outside_static_routes no_outside_static_routes" must be set
    no_outside_static_routes = true
  }
  vnet {
    // One of the arguments from this list "new_vnet existing_vnet" must be set

    new_vnet {
      // One of the arguments from this list "name autogenerate" must be set
      name = "name"

      primary_ipv4 = "10.1.0.0/16"
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

`address` - (Optional) Site's geographical address that can be used determine its latitude and longitude. (`String`).

`azure_region` - (Required) name for azure region in which this site will be launched. (`String`).

`coordinates` - (Optional) Site longitude and latitude co-ordinates. See [Coordinates ](#coordinates) below for details.

`assisted` - (Optional) In assisted deployment get Azure parameters generated in status of this objects and run volterra provided terraform script. (bool).

`azure_cred` - (Optional) Reference to Azure credentials for automatic deployment. See [ref](#ref) below for details.

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`Int`).

`machine_type` - (Optional) Standard_D5_v2 (16 x vCPU, 56GB RAM) very high performance (`String`).

`nodes_per_az` - (Optional) Auto scale maximum worker nodes limit up to 21, value of zero will disable auto scale (`Int`).

`operating_system_version` - (Optional) Desired Operating System version for this site. (`String`).

`resource_group` - (Required) Azure resource group for resources that will be created (`String`).

`ingress_egress_gw` - (Optional) Two interface site is useful when site is used as ingress/egress gateway to the Vnet.. See [Ingress Egress Gw ](#ingress-egress-gw) below for details.

`ingress_gw` - (Optional) One interface site is useful when site is only used as ingress gateway to the Vnet.. See [Ingress Gw ](#ingress-gw) below for details.

`voltstack_cluster` - (Optional) Voltstack Cluster using single interface, useful for deploying K8s cluster.. See [Voltstack Cluster ](#voltstack-cluster) below for details.

`ssh_key` - (Optional) Public SSH key for accessing the site. (`String`).

`vnet` - (Required) Choice of using existing Vnet or create new Vnet. See [Vnet ](#vnet) below for details.

`volterra_software_version` - (Optional) Desired Volterra software version for this site, a string matching released set of software components. (`String`).

### Active Forward Proxy Policies

Enable Forward Proxy for this site and manage policies.

`forward_proxy_policies` - (Required) Ordered List of Network Policies active for this network firewall. See [ref](#ref) below for details.

### Active Network Policies

Network Policies active for this site..

`network_policies` - (Required) Ordered List of Network Policies active for this network firewall. See [ref](#ref) below for details.

### Autogenerate

Autogenerate the Vnet Name.

### Az Nodes

Only Single AZ or Three AZ(s) nodes are supported currently..

`azure_az` - (Required) Name for AWS availability Zone. (`String`).

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`String`).

`inside_subnet` - (Optional) Subnets for the inside interface of the node. See [Inside Subnet ](#inside-subnet) below for details.

`outside_subnet` - (Optional) Subnets for the outside interface of the node. See [Outside Subnet ](#outside-subnet) below for details.

### Coordinates

Site longitude and latitude co-ordinates.

`latitude` - (Optional) Latitude of the site location (`Float`).

`longitude` - (Optional) longitude of site location (`Float`).

### Custom Static Route

Use Custom static route to configure all advanced options.

`attrs` - (Optional) List of route attributes associated with the static route (`List of Strings`).

`labels` - (Optional) Add Labels for this Static Route, these labels can be used in network policy (`String`).

`nexthop` - (Optional) Nexthop for the route. See [Nexthop ](#nexthop) below for details.

`subnets` - (Optional) List of route prefixes. See [Subnets ](#subnets) below for details.

### Disable Forward Proxy

Forward Proxy is disabled for this connector.

### Enable Forward Proxy

Forward Proxy is enabled for this connector.

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2 seconds (`Int`).

`max_connect_attempts` - (Optional) Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1. (`Int`).

`white_listed_ports` - (Optional) Example "tmate" server port (`Int`).

`white_listed_prefixes` - (Optional) Example "tmate" server ip (`String`).

### Existing Vnet

Information about existing Vnet.

`resource_group` - (Required) Resource group of existing Vnet (`String`).

`vnet_name` - (Required) Name of existing Vnet (`String`).

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

### Ingress Egress Gw

Two interface site is useful when site is used as ingress/egress gateway to the Vnet..

`az_nodes` - (Optional) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`inside_static_routes` - (Optional) Manage static routes for inside network.. See [Inside Static Routes ](#inside-static-routes) below for details.

`no_inside_static_routes` - (Optional) Static Routes disabled for inside network. (bool).

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

### Ingress Gw

One interface site is useful when site is only used as ingress gateway to the Vnet..

`az_nodes` - (Optional) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

### Inside Static Routes

Manage static routes for inside network..

`static_route_list` - (Required) List of Static routes. See [Static Route List ](#static-route-list) below for details.

### Inside Subnet

Subnets for the inside interface of the node.

`subnet` - (Optional) Information about existing subnet.. See [Subnet ](#subnet) below for details.

`subnet_param` - (Optional) Parameters for creating new subnet. . See [Subnet Param ](#subnet-param) below for details.

### Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

### New Vnet

Parameters for creating new Vnet.

`autogenerate` - (Optional) Autogenerate the Vnet Name (bool).

`name` - (Optional) Specify the Vnet Name (`String`).

`primary_ipv4` - (Required) IPv4 CIDR block for this Vnet. It has to be private address space. (`String`).

### Nexthop

Nexthop for the route.

`interface` - (Optional) Nexthop is network interface when type is "Network-Interface". See [ref](#ref) below for details.

`nexthop_address` - (Optional) Nexthop address when type is "Use-Configured". See [Nexthop Address ](#nexthop-address) below for details.

`type` - (Optional) Identifies the type of next-hop (`String`).

### Nexthop Address

Nexthop address when type is "Use-Configured".

`ipv4` - (Optional) IPv4 Address. See [Ipv4 ](#ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ipv6 ](#ipv6) below for details.

### No Forward Proxy

Disable Forward Proxy for this site.

### No Global Network

No global network to connect .

### No Inside Static Routes

Static Routes disabled for inside network..

### No Network Policy

Network Policy is disabled for this site..

### No Outside Static Routes

Static Routes disabled for outside network..

### Outside Static Routes

Manage static routes for outside network..

`static_route_list` - (Required) List of Static routes. See [Static Route List ](#static-route-list) below for details.

### Outside Subnet

Subnets for the outside interface of the node.

`subnet` - (Optional) Information about existing subnet.. See [Subnet ](#subnet) below for details.

`subnet_param` - (Optional) Parameters for creating new subnet. . See [Subnet Param ](#subnet-param) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Sli To Global Dr

Site local inside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Slo To Global Dr

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Static Route List

List of Static routes.

`custom_static_route` - (Optional) Use Custom static route to configure all advanced options. See [Custom Static Route ](#custom-static-route) below for details.

`simple_static_route` - (Optional) Use simple static route for prefix pointing to single interface in the network (`String`).

### Subnet

Information about existing subnet..

`subnet_resource_grp` - (Optional) Specify name of Resource Group (`String`).

`vnet_resource_group` - (Optional) Use the same Resource Group as the Vnet (bool).

`subnet_name` - (Required) Name of existing subnet. (`String`).

### Subnet Param

Parameters for creating new subnet. .

`ipv4` - (Required) IPv4 subnet prefix for this subnet (`String`).

`ipv6` - (Optional) IPv6 subnet prefix for this subnet (`String`).

### Subnets

List of route prefixes.

`ipv4` - (Optional) IPv4 Subnet Address. See [Ipv4 ](#ipv4) below for details.

`ipv6` - (Optional) IPv6 Subnet Address. See [Ipv6 ](#ipv6) below for details.

### Vnet

Choice of using existing Vnet or create new Vnet.

`existing_vnet` - (Optional) Information about existing Vnet. See [Existing Vnet ](#existing-vnet) below for details.

`new_vnet` - (Optional) Parameters for creating new Vnet. See [New Vnet ](#new-vnet) below for details.

### Vnet Resource Group

Use the same Resource Group as the Vnet.

### Voltstack Cluster

Voltstack Cluster using single interface, useful for deploying K8s cluster..

`az_nodes` - (Optional) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outisde network.. See [Outside Static Routes ](#outside-static-routes) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured azure_vnet_site.
