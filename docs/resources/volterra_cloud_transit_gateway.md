---

page_title: "Volterra: cloud_transit_gateway"

description: "The cloud_transit_gateway allows CRUD of Cloud Transit Gateway resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------------

Resource volterra_cloud_transit_gateway
=======================================

The Cloud Transit Gateway allows CRUD of Cloud Transit Gateway resource on Volterra SaaS

~> **Note:** Please refer to [Cloud Transit Gateway API docs](https://docs.cloud.f5.com/docs-v2/api/cloud-transit-gateway) to learn more

Example Usage
-------------

```hcl
resource "volterra_cloud_transit_gateway" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "aws" must be set

  aws {
    aws_region = "us-east-1"

    cloud_credentials {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }

    primary_secure_mesh_site {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }

    service_vpc {
      subnet_ids = ["subnet-12345678901234567"]

      vpc_id = "vpc-12345678901234567"
    }

    tgw_asn = "64512"

    // One of the arguments from this list "existing_tgw new_tgw" must be set

    new_tgw {
      // One of the arguments from this list "system_generated user_assigned" must be set

      system_generated = true
    }

    // One of the arguments from this list "reserved_tgw_cidr tgw_cidr" must be set

    reserved_tgw_cidr = true
    tgw_cidr_ipv4 = "tgw_cidr_ipv4"
    volterra_site_asn = "64500"
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

###### One of the arguments from this list "aws" must be set

`aws` - (Optional) AWS Backend. We use AWS Transit Gateway to attach VPC via Cloud Connect object.. See [Cloud Choice Aws ](#cloud-choice-aws) below for details.

### Asn Choice System Generated

F5XC will automatically assign a private ASN for TGW and F5XC Site.

### Asn Choice User Assigned

User is managing the ASN for TGW and F5XC Site..

`tgw_asn` - (Optional) TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534. (`Int`).

`volterra_site_asn` - (Optional) F5XC Site ASN. (`Int`).

### Aws Service Vpc

Service VPC info TGW attached to.

`subnet_ids` - (Required) List of subnet IDs (`String`).

`vpc_id` - (Required) Enter the VPC ID where secure mesh site exists (`String`).

### Cloud Choice Aws

AWS Backend. We use AWS Transit Gateway to attach VPC via Cloud Connect object..

`aws_region` - (Required) AWS Region of your transit gateway present. (`String`).

`cloud_credentials` - (Optional) Reference to AWS credentials.. See [ref](#ref) below for details.

`primary_secure_mesh_site` - (Optional) Reference to Secure Mesh Site. See [ref](#ref) below for details.

`service_vpc` - (Required) Service VPC info TGW attached to. See [Aws Service Vpc ](#aws-service-vpc) below for details.

`tgw_asn` - (Optional) TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534. (`Int`).(Deprecated)

###### One of the arguments from this list "existing_tgw, new_tgw" must be set

`existing_tgw` - (Optional) Information about existing TGW. See [Tgw Choice Existing Tgw ](#tgw-choice-existing-tgw) below for details.

`new_tgw` - (Optional) Details needed to create new TGW. See [Tgw Choice New Tgw ](#tgw-choice-new-tgw) below for details.

###### One of the arguments from this list "reserved_tgw_cidr, tgw_cidr" must be set

`reserved_tgw_cidr` - (Optional) Autogenerate and reserve a TGW CIDR Block from the Primary CIDR (`Bool`).

`tgw_cidr` - (Optional) Specify TGW CIDR block. See [Tgw Cidr Choice Tgw Cidr ](#tgw-cidr-choice-tgw-cidr) below for details.

`tgw_cidr_ipv4` - (Optional) Allocated TGW CIDR by config service (`String`).(Deprecated)

`volterra_site_asn` - (Optional) F5XC Site ASN. (`Int`).(Deprecated)

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Tgw Choice Existing Tgw

Information about existing TGW.

`tgw_asn` - (Optional) TGW ASN. (`Int`).

`tgw_id` - (Optional) Existing TGW ID (`String`).

`volterra_site_asn` - (Optional) F5XC Site ASN. (`Int`).

### Tgw Choice New Tgw

Details needed to create new TGW.

###### One of the arguments from this list "system_generated, user_assigned" must be set

`system_generated` - (Optional) F5XC will automatically assign a private ASN for TGW and F5XC Site (`Bool`).

`user_assigned` - (Optional) User is managing the ASN for TGW and F5XC Site.. See [Asn Choice User Assigned ](#asn-choice-user-assigned) below for details.

### Tgw Cidr Choice Reserved Tgw Cidr

Autogenerate and reserve a TGW CIDR Block from the Primary CIDR.

### Tgw Cidr Choice Tgw Cidr

Specify TGW CIDR block.

`ipv4` - (Required) IPv4 subnet prefix for this subnet (`String`).

`ipv6` - (Optional) IPv6 subnet prefix for this subnet (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured cloud_transit_gateway.
