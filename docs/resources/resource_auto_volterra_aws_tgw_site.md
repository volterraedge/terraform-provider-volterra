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
      name_tag      = "name_tag"
      primary_ipv4  = "10.1.0.0/16"
    }
    ssh_key = "ssh-rsa AAAAB..."

    // One of the arguments from this list "new_tgw existing_tgw" must be set

    new_tgw {
      // One of the arguments from this list "system_generated user_assigned" must be set
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

`aws_parameters` - (Required) Configure AWS TGW, services VPC and site nodes parameters.. See [Aws Parameters ](#aws-parameters) below for details.

`operating_system_version` - (Optional) Desired Operating System version for this site. (`String`).

`tgw_security` - (Optional) Security Configuration for transit gateway. See [Tgw Security ](#tgw-security) below for details.

`vn_config` - (Optional) Virtual Network Configuration for transit gateway. See [Vn Config ](#vn-config) below for details.

`volterra_software_version` - (Optional) Desired Volterra software version for this site, a string matching released set of software components. (`String`).

`vpc_attachments` - (Optional) VPC attachments to transit gateway. See [Vpc Attachments ](#vpc-attachments) below for details.

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

### Inside Subnet

Subnets for the inside interface of the node.

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Subnet Param ](#subnet-param) below for details.

### Outside Subnet

Subnets for the outside interface of the node.

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Subnet Param ](#subnet-param) below for details.

### Tgw Security

Security Configuration for transit gateway.

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy_policy` - (Optional) Disable Forward Proxy for this site (bool).

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
