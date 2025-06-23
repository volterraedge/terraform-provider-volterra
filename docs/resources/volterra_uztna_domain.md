---

page_title: "Volterra: uztna_domain"

description: "The uztna_domain allows CRUD of Uztna Domain resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_uztna_domain
==============================

The Uztna Domain allows CRUD of Uztna Domain resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Domain API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-domain) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_domain" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  access_fqdn = ["access_fqdn"]

  app_vip_subnet = ["app_vip_subnet"]

  cert {
    certificate {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  gateways {
    bigip_site {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }

    perimeter_re {
      // One of the arguments from this list "all_cloud re_sites" can be set

      all_cloud = true
    }
  }

  lease_pool {
    // One of the arguments from this list "ipv4_leasepool" must be set

    ipv4_leasepool {
      ipv4_leasepool {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }
    }
  }

  policy {
    policy {
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

`access_fqdn` - (Required) FQDN to access the gateways (`String`).

`app_vip_pool` - (Optional) Appliacation VIP Pools . See [App Vip Pool ](#app-vip-pool) below for details.

`app_vip_subnet` - (Required) Application VIP prefix holds array of subnet, under which Global ip allocator allocates ip on this range (`List of String`).

`cert` - (Required) Domain in XC is an established pattern and we would reuse the same.. See [Cert ](#cert) below for details.

`gateways` - (Required) List of all Cloud Gateways and Big-IP Edge Gateways. See [Gateways ](#gateways) below for details.

`lease_pool` - (Required) The Lease Pool List assigned to the Zero Trust Domain. . See [Lease Pool ](#lease-pool) below for details.

`policy` - (Required) This is used to import or create new ZTNA Policy. See [Policy ](#policy) below for details.

### App Vip Pool

Appliacation VIP Pools .

###### One of the arguments from this list "ipv4_app_vip_pool" can be set

`ipv4_app_vip_pool` - (Required) IPv4 App VIP Pools. See [Ipaddress Type Ipv4 App Vip Pool ](#ipaddress-type-ipv4-app-vip-pool) below for details.

### Cert

Domain in XC is an established pattern and we would reuse the same..

`certificate` - (Optional) Select/Add one or more TLS Certificate objects to associate with this ZeroTrust Domain. See [ref](#ref) below for details.

### Gateways

List of all Cloud Gateways and Big-IP Edge Gateways.

`bigip_site` - (Optional) From the available Big-IP Edge Gateway List select Big-IP Edge Gateway.. See [ref](#ref) below for details.

`perimeter_re` - (Optional) List of all Cloud Gateways and Big-IP Edge Gateways. See [Gateways Perimeter Re ](#gateways-perimeter-re) below for details.

### Lease Pool

The Lease Pool List assigned to the Zero Trust Domain. .

###### One of the arguments from this list "ipv4_leasepool" must be set

`ipv4_leasepool` - (Optional) IPv4 Lease Pools. See [Ipaddress Type Ipv4 Leasepool ](#ipaddress-type-ipv4-leasepool) below for details.

### Policy

This is used to import or create new ZTNA Policy.

`policy` - (Required) Select/Add ZTNA Policy to associate with this ZeroTrust Domain. See [ref](#ref) below for details.

### Cloud Gateway Choice All Cloud

This option will allow to advertise on all Cloud Gateways.

### Cloud Gateway Choice Re Sites

This option will allow advertise on selected Cloud Gateways.

`cloud_gateway` - (Optional) Selected Cloud Gateway. See [ref](#ref) below for details.

### Gateways Perimeter Re

List of all Cloud Gateways and Big-IP Edge Gateways.

###### One of the arguments from this list "all_cloud, re_sites" can be set

`all_cloud` - (Optional) This option will allow to advertise on all Cloud Gateways (`Bool`).

`re_sites` - (Optional) This option will allow advertise on selected Cloud Gateways. See [Cloud Gateway Choice Re Sites ](#cloud-gateway-choice-re-sites) below for details.

### Ipaddress Type Ipv4 App Vip Pool

IPv4 App VIP Pools.

`ipv4_app_vip_pool` - (Required) Select or create new IPv4 App VIP Pool. See [ref](#ref) below for details.

### Ipaddress Type Ipv4 Leasepool

IPv4 Lease Pools.

`ipv4_leasepool` - (Required) Select or create new IPv4 Lease Pool. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_domain.
