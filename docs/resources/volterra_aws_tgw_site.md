---

page_title: "Volterra: aws_tgw_site"

description: "The aws_tgw_site allows CRUD of Aws Tgw Site resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_aws_tgw_site
==============================

The Aws Tgw Site allows CRUD of Aws Tgw Site resource on Volterra SaaS

~> **Note:** Please refer to [Aws Tgw Site API docs](https://volterra.io/docs/api/views-aws-tgw-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_aws_tgw_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  aws_parameters {
    aws_certified_hw = "aws-byol-multi-nic-voltmesh"
    aws_region       = "us-east-1"

    az_nodes {
      aws_az_name = "us-west-2a"
      disk_size   = "disk_size"

      inside_subnet {
        // One of the arguments from this list "subnet_param existing_subnet_id" must be set

        subnet_param {
          ipv4 = "10.1.2.0/24"
          ipv6 = "1234:568:abcd:9100::/64"
        }
      }

      outside_subnet {
        // One of the arguments from this list "subnet_param existing_subnet_id" must be set

        subnet_param {
          ipv4 = "10.1.2.0/24"
          ipv6 = "1234:568:abcd:9100::/64"
        }
      }
    }

    // One of the arguments from this list "aws_cred assisted" must be set

    aws_cred {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
    disk_size     = "80"
    instance_type = "a1.xlarge"
    nodes_per_az  = "2"

    // One of the arguments from this list "new_vpc vpc_id" must be set

    new_vpc {
      allocate_ipv6 = true

      // One of the arguments from this list "name_tag autogenerate" must be set
      name_tag = "name_tag"

      primary_ipv4 = "10.1.0.0/16"
    }
    ssh_key = "ssh-rsa AAAAB..."

    // One of the arguments from this list "new_tgw existing_tgw" must be set

    new_tgw {
      // One of the arguments from this list "user_assigned system_generated" must be set
      system_generated = true
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

`aws_parameters` - (Required) Configure AWS TGW, services VPC and site nodes parameters.. See [Aws Parameters ](#aws-parameters) below for details.

`coordinates` - (Optional) Site longitude and latitude co-ordinates. See [Coordinates ](#coordinates) below for details.

`operating_system_version` - (Optional) Desired Operating System version for this site. (`String`).

`tgw_security` - (Optional) Security Configuration for transit gateway. See [Tgw Security ](#tgw-security) below for details.

`vn_config` - (Optional) Virtual Network Configuration for transit gateway. See [Vn Config ](#vn-config) below for details.

`volterra_software_version` - (Optional) Desired Volterra software version for this site, a string matching released set of software components. (`String`).

`vpc_attachments` - (Optional) VPC attachments to transit gateway. See [Vpc Attachments ](#vpc-attachments) below for details.

### Active Forward Proxy Policies

Enable Forward Proxy for this site and manage policies.

`forward_proxy_policies` - (Required) Ordered List of Network Policies active for this network firewall. See [ref](#ref) below for details.

### Active Network Policies

Network Policies active for this site..

`network_policies` - (Required) Ordered List of Network Policies active for this network firewall. See [ref](#ref) below for details.

### Assisted

In assisted deployment get AWS parameters generated in status of this objects and run volterra provided terraform script..

### Autogenerate

Autogenerate the VPC Name.

### Aws Parameters

Configure AWS TGW, services VPC and site nodes parameters..

`aws_certified_hw` - (Required) Name for AWS certified hardware. (`String`).

`aws_region` - (Required) Name for AWS Region. (`String`).

`az_nodes` - (Optional) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`assisted` - (Optional) In assisted deployment get AWS parameters generated in status of this objects and run volterra provided terraform script. (bool).

`aws_cred` - (Optional) Reference to AWS credentials for automatic deployment. See [ref](#ref) below for details.

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`Int`).

`instance_type` - (Required) Select Instance size based on performance needed (`String`).

`nodes_per_az` - (Optional) Auto scale maximum worker nodes limit up to 21, value of zero will disable auto scale (`Int`).

`new_vpc` - (Optional) Parameters for creating new VPC. See [New Vpc ](#new-vpc) below for details.

`vpc_id` - (Optional) Information about existing VPC (`String`).

`ssh_key` - (Optional) Public SSH key for accessing nodes of the site. (`String`).

`existing_tgw` - (Optional) Information about existing TGW. See [Existing Tgw ](#existing-tgw) below for details.

`new_tgw` - (Optional) Parameters for creating new TGW. See [New Tgw ](#new-tgw) below for details.

### Az Nodes

Only Single AZ or Three AZ(s) nodes are supported currently..

`aws_az_name` - (Required) Name for AWS availability Zone, should match with AWS region selected. (`String`).

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

### Existing Tgw

Information about existing TGW.

`tgw_asn` - (Optional) TGW ASN. (`Int`).

`tgw_id` - (Optional) Existing TGW ID (`String`).

`volterra_site_asn` - (Optional) Volterra Site ASN. (`Int`).

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

### Inside Static Routes

Manage static routes for inside network..

`static_route_list` - (Required) List of Static routes. See [Static Route List ](#static-route-list) below for details.

### Inside Subnet

Subnets for the inside interface of the node.

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Subnet Param ](#subnet-param) below for details.

### Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

### New Tgw

Parameters for creating new TGW.

`system_generated` - (Optional) Volterra will automatically assign a private ASN for TGW and Volterra Site (bool).

`user_assigned` - (Optional) User is managing the ASN for TGW and Volterra Site.. See [User Assigned ](#user-assigned) below for details.

### New Vpc

Parameters for creating new VPC.

`allocate_ipv6` - (Optional) Allocate IPv6 CIDR block from AWS (`Bool`).

`autogenerate` - (Optional) Autogenerate the VPC Name (bool).

`name_tag` - (Optional) Specify the VPC Name (`String`).

`primary_ipv4` - (Required) The Primary IPv4 block cannot be modified. All subnets prefixes in this VPC must be part of this CIDR block. (`String`).

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

No global network to connect.

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

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Subnet Param ](#subnet-param) below for details.

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

### Subnet Param

Parameters for creating new subnet.

`ipv4` - (Required) IPv4 subnet prefix for this subnet (`String`).

`ipv6` - (Optional) IPv6 subnet prefix for this subnet (`String`).

### Subnets

List of route prefixes.

`ipv4` - (Optional) IPv4 Subnet Address. See [Ipv4 ](#ipv4) below for details.

`ipv6` - (Optional) IPv6 Subnet Address. See [Ipv6 ](#ipv6) below for details.

### System Generated

Volterra will automatically assign a private ASN for TGW and Volterra Site.

### Tgw Security

Security Configuration for transit gateway.

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

### User Assigned

User is managing the ASN for TGW and Volterra Site..

`tgw_asn` - (Optional) TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534. (`Int`).

`volterra_site_asn` - (Optional) Volterra Site ASN. (`Int`).

### Vn Config

Virtual Network Configuration for transit gateway.

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`inside_static_routes` - (Optional) Manage static routes for inside network.. See [Inside Static Routes ](#inside-static-routes) below for details.

`no_inside_static_routes` - (Optional) Static Routes disabled for inside network. (bool).

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

### Vpc Attachments

VPC attachments to transit gateway.

`vpc_list` - (Optional)vpc_list. See [Vpc List ](#vpc-list) below for details.

### Vpc List

vpc_list.

`labels` - (Optional) Add Labels for each of the VPC ID, these labels can be used in network policy (`String`).

`vpc_id` - (Optional) Information about existing VPC (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured aws_tgw_site.
