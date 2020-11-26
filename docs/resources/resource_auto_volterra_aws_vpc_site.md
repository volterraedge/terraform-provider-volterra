---

page_title: "Volterra: aws_vpc_site"

description: "The aws_vpc_site allows CRUD of Aws Vpc Site resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_aws_vpc_site
==============================

The Aws Vpc Site allows CRUD of Aws Vpc Site resource on Volterra SaaS

~> **Note:** Please refer to [Aws Vpc Site API docs](https://volterra.io/docs/api/views-aws-vpc-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_aws_vpc_site" "example" {
  name       = "acmecorp-web"
  namespace  = "staging"
  aws_region = ["us-east-1"]

  // One of the arguments from this list "aws_cred assisted" must be set

  aws_cred {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }
  instance_type = ["a1.xlarge"]

  // One of the arguments from this list "ingress_gw ingress_egress_gw voltstack_cluster" must be set

  ingress_gw {
    aws_certified_hw = "aws-byol-voltmesh"

    az_nodes {
      aws_az_name = "us-west-2a"
      disk_size   = "disk_size"

      local_subnet {
        // One of the arguments from this list "subnet_param existing_subnet_id" must be set

        subnet_param {
          ipv4 = "10.1.2.0/24"
          ipv6 = "1234:568:abcd:9100::/64"
        }
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

`aws_region` - (Required) Name for AWS Region. (`String`).

`assisted` - (Optional) In assisted deployment get AWS parameters generated in status of this objects and run volterra provided terraform script. (bool).

`aws_cred` - (Optional) Reference to AWS credentials for automatic deployment. See [ref](#ref) below for details.

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`Int`).

`instance_type` - (Required) m5.4xlarge (16 x vCPU, 64GB RAM) very high performance (`String`).

`nodes_per_az` - (Optional) Auto scale maximum worker nodes limit up to 21, value of zero will disable auto scale (`Int`).

`operating_system_version` - (Optional) Desired Operating System version for this site. (`String`).

`ingress_egress_gw` - (Optional) Two interface site is useful when site is used as ingress/egress gateway to the VPC.. See [Ingress Egress Gw ](#ingress-egress-gw) below for details.

`ingress_gw` - (Optional) One interface site is useful when site is only used as ingress gateway to the VPC.. See [Ingress Gw ](#ingress-gw) below for details.

`voltstack_cluster` - (Optional) Voltstack Cluster using single interface, useful for deploying k8s cluster.. See [Voltstack Cluster ](#voltstack-cluster) below for details.

`ssh_key` - (Optional) Public SSH key for accessing the site. (`String`).

`volterra_software_version` - (Optional) Desired Volterra software version for this site, a string matching released set of software components. (`String`).

`vpc` - (Optional) Choice of using existing VPC or create new VPC. See [Vpc ](#vpc) below for details.

### Az Nodes

Only Single AZ or Three AZ(s) nodes are supported currently..

`aws_az_name` - (Required) Name for AWS availability Zone, should match with AWS region selected. (`String`).

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`String`).

`inside_subnet` - (Optional) Subnets for the inside interface of the node. See [Inside Subnet ](#inside-subnet) below for details.

`outside_subnet` - (Optional) Subnets for the outside interface of the node. See [Outside Subnet ](#outside-subnet) below for details.

### Ingress Egress Gw

Two interface site is useful when site is used as ingress/egress gateway to the VPC..

`aws_certified_hw` - (Required) Name for AWS certified hardware. (`String`).

`az_nodes` - (Optional) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`inside_static_routes` - (Optional) Manage static routes for inside network.. See [Inside Static Routes ](#inside-static-routes) below for details.

`no_inside_static_routes` - (Optional) Static Routes disabled for inside network. (bool).

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy_policy` - (Optional) Disable Forward Proxy for this site (bool).

### Ingress Gw

One interface site is useful when site is only used as ingress gateway to the VPC..

`aws_certified_hw` - (Required) Name for AWS certified hardware. (`String`).

`az_nodes` - (Required) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

### Inside Subnet

Subnets for the inside interface of the node.

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Subnet Param ](#subnet-param) below for details.

### Local Subnet

Subnets for the site local interface of the node.

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Subnet Param ](#subnet-param) below for details.

### Outside Subnet

Subnets for the outside interface of the node.

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Subnet Param ](#subnet-param) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Voltstack Cluster

Voltstack Cluster using single interface, useful for deploying k8s cluster..

`aws_certified_hw` - (Required) Name for AWS certified hardware. (`String`).

`az_nodes` - (Optional) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`no_outside_static_routes` - (Optional) Static Routes disabled for site local network. (bool).

`outside_static_routes` - (Optional) Manage static routes for site local network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy_policy` - (Optional) Disable Forward Proxy for this site (bool).

### Vpc

Choice of using existing VPC or create new VPC.

`new_vpc` - (Optional) Parameters for creating new VPC. See [New Vpc ](#new-vpc) below for details.

`vpc_id` - (Optional) Information about existing VPC ID (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured aws_vpc_site.
