---

page_title: "Volterra: gcp_vpc_site"

description: "The gcp_vpc_site allows CRUD of Gcp Vpc Site resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_gcp_vpc_site
==============================

The Gcp Vpc Site allows CRUD of Gcp Vpc Site resource on Volterra SaaS

~> **Note:** Please refer to [Gcp Vpc Site API docs](https://volterra.io/docs/api/views-gcp-vpc-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_gcp_vpc_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "cloud_credentials assisted" must be set

  cloud_credentials {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }
  gcp_region    = ["us-west1"]
  instance_type = ["n1-standard-4"]

  // One of the arguments from this list "ingress_gw ingress_egress_gw voltstack_cluster" must be set

  voltstack_cluster {
    az_nodes {
      gcp_zone_name = "us-west1-a"

      local_subnet {
        // One of the arguments from this list "new_subnet existing_subnet" must be set

        new_subnet {
          primary_ipv4 = "10.1.0.0/16"
          subnet_name  = "subnet1-in-network1"
        }
      }
    }

    gcp_certified_hw = "gcp-byol-voltstack-combo"

    // One of the arguments from this list "no_global_network global_network_list" must be set
    no_global_network = true

    // One of the arguments from this list "no_network_policy active_network_policies" must be set
    no_network_policy = true

    // One of the arguments from this list "no_outside_static_routes outside_static_routes" must be set
    no_outside_static_routes = true

    // One of the arguments from this list "no_forward_proxy_policy active_forward_proxy_policies forward_proxy_allow_all" must be set
    no_forward_proxy_policy = true

    site_local_network {
      // One of the arguments from this list "new_network existing_network" must be set

      new_network {
        name = "network1"
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

`assisted` - (Optional) In assisted deployment get GCP parameters generated in status of this objects and run volterra provided terraform script. (bool).

`cloud_credentials` - (Optional) Reference to GCP credentials for automatic deployment. See [ref](#ref) below for details.

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`Int`).

`gcp_region` - (Required) Name for GCP Region. (`String`).

`instance_type` - (Required) n1-standard-16 (16 x vCPU, 64GB RAM) very high performance (`String`).

`nodes_per_az` - (Optional) Auto scale maximum worker nodes limit up to 21, value of zero will disable auto scale (`Int`).

`operating_system_version` - (Optional) Desired Operating System version for this site. (`String`).

`ingress_egress_gw` - (Optional) Two interface site is useful when site is used as ingress/egress gateway to the VPC network.. See [Ingress Egress Gw ](#ingress-egress-gw) below for details.

`ingress_gw` - (Optional) One interface site is useful when site is only used as ingress gateway to the VPC network.. See [Ingress Gw ](#ingress-gw) below for details.

`voltstack_cluster` - (Optional) Voltstack Cluster using single interface, useful for deploying k8s cluster.. See [Voltstack Cluster ](#voltstack-cluster) below for details.

`ssh_key` - (Optional) Public SSH key for accessing the site. (`String`).

`volterra_software_version` - (Optional) Desired Volterra software version for this site, a string matching released set of software components. (`String`).

### Az Nodes

Only Single zone or Three nodes are supported currently..

`gcp_zone_name` - (Required) Name for GCP zone, should match with region selected. (`String`).

`inside_subnet` - (Optional) Subnets for the inside interface of the node, should be in inside network. See [Inside Subnet ](#inside-subnet) below for details.

`outside_subnet` - (Optional) Subnets for the outside interface of the node, should be in outside network. See [Outside Subnet ](#outside-subnet) below for details.

### Ingress Egress Gw

Two interface site is useful when site is used as ingress/egress gateway to the VPC network..

`az_nodes` - (Optional) Only Single zone or Three nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`gcp_certified_hw` - (Required) Name for GCP certified hardware. (`String`).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`inside_network` - (Optional) Network Subnets for the inside interface of the node. See [Inside Network ](#inside-network) below for details.

`inside_static_routes` - (Optional) Manage static routes for inside network.. See [Inside Static Routes ](#inside-static-routes) below for details.

`no_inside_static_routes` - (Optional) Static Routes disabled for inside network. (bool).

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`outside_network` - (Optional) Network Subnets for the outside interface of the node. See [Outside Network ](#outside-network) below for details.

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy_policy` - (Optional) Disable Forward Proxy for this site (bool).

### Ingress Gw

One interface site is useful when site is only used as ingress gateway to the VPC network..

`az_nodes` - (Required) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`gcp_certified_hw` - (Required) Name for GCP certified hardware. (`String`).

`local_network` - (Optional) Network Subnets for the local interface of the node. See [Local Network ](#local-network) below for details.

### Inside Network

Network Subnets for the inside interface of the node.

`existing_network` - (Optional) Information about existing VPC network. See [Existing Network ](#existing-network) below for details.

`new_network` - (Optional) Parameters for creating new VPC network. See [New Network ](#new-network) below for details.

### Inside Subnet

Subnets for the inside interface of the node, should be in inside network.

`existing_subnet` - (Optional) Information about existing subnet. See [Existing Subnet ](#existing-subnet) below for details.

`new_subnet` - (Optional) Parameters for creating new subnet. See [New Subnet ](#new-subnet) below for details.

### Local Network

Network Subnets for the local interface of the node.

`existing_network` - (Optional) Information about existing VPC network. See [Existing Network ](#existing-network) below for details.

`new_network` - (Optional) Parameters for creating new VPC network. See [New Network ](#new-network) below for details.

### Local Subnet

Subnets for the local interface of the node, should be in local network.

`existing_subnet` - (Optional) Information about existing subnet. See [Existing Subnet ](#existing-subnet) below for details.

`new_subnet` - (Optional) Parameters for creating new subnet. See [New Subnet ](#new-subnet) below for details.

### Outside Network

Network Subnets for the outside interface of the node.

`existing_network` - (Optional) Information about existing VPC network. See [Existing Network ](#existing-network) below for details.

`new_network` - (Optional) Parameters for creating new VPC network. See [New Network ](#new-network) below for details.

### Outside Subnet

Subnets for the outside interface of the node, should be in outside network.

`existing_subnet` - (Optional) Information about existing subnet. See [Existing Subnet ](#existing-subnet) below for details.

`new_subnet` - (Optional) Parameters for creating new subnet. See [New Subnet ](#new-subnet) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Site Local Network

Network Subnets for the site local interface of the node.

`existing_network` - (Optional) Information about existing VPC network. See [Existing Network ](#existing-network) below for details.

`new_network` - (Optional) Parameters for creating new VPC network. See [New Network ](#new-network) below for details.

### Voltstack Cluster

Voltstack Cluster using single interface, useful for deploying k8s cluster..

`az_nodes` - (Optional) Only Single zone or Three nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`gcp_certified_hw` - (Required) Name for GCP certified hardware. (`String`).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy_policy` - (Optional) Disable Forward Proxy for this site (bool).

`site_local_network` - (Optional) Network Subnets for the site local interface of the node. See [Site Local Network ](#site-local-network) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured gcp_vpc_site.
