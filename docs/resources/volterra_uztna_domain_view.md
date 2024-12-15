---

page_title: "Volterra: uztna_domain_view"

description: "The uztna_domain_view allows CRUD of Uztna Domain View resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_uztna_domain_view
===================================

The Uztna Domain View allows CRUD of Uztna Domain View resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Domain View API docs](https://docs.cloud.f5.com/docs-v2/api/views-uztna-domain-view) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_domain_view" "example" {
  name       = "acmecorp-web"
  namespace  = "staging"
  access_url = ["access_url"]

  cert {
    certificate {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  gateways {
    perimeter_re {
      // One of the arguments from this list "all_cloud re_sites" can be set

      all_cloud = true
    }

    uztna_gateway {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  lease_pool {
    // One of the arguments from this list "ipv4_ipv6_leasepool ipv4_leasepool ipv6_leasepool" must be set

    ipv4_leasepool {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  policy {
    policy_name {
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

`access_url` - (Required) Url to access the gateways (`String`).

`app_vip_pool` - (Optional) Application VIP Pools . See [App Vip Pool ](#app-vip-pool) below for details.

`cert` - (Required) Domain in XC is an established pattern and we would reuse the same.. See [Cert ](#cert) below for details.

`gateways` - (Required) List of all RE prime and Big Ip edge CE. See [Gateways ](#gateways) below for details.

`lease_pool` - (Required) The Lease Pool assigned to the Zero Trust Domain. . See [Lease Pool ](#lease-pool) below for details.

`policy` - (Required) The name of the ZTNA profile. See [Policy ](#policy) below for details.

### App Vip Pool

Application VIP Pools .

`app_vip_pool` - (Optional) VIP Pools. See [ref](#ref) below for details.

### Cert

Domain in XC is an established pattern and we would reuse the same..

`certificate` - (Optional) Select/Add one or more TLS Certificate objects to associate with this ZeroTrust Domain. See [ref](#ref) below for details.

### Gateways

List of all RE prime and Big Ip edge CE.

`perimeter_re` - (Optional) Cloud Gateways and Big-IP Edge Gateways. See [Gateways Perimeter Re ](#gateways-perimeter-re) below for details.

`uztna_gateway` - (Optional) Select BIG-IP Edge Gateway for Advertisement .. See [ref](#ref) below for details.

### Lease Pool

The Lease Pool assigned to the Zero Trust Domain. .

###### One of the arguments from this list "ipv4_ipv6_leasepool, ipv4_leasepool, ipv6_leasepool" must be set

`ipv4_ipv6_leasepool` - (Optional) Select or create new IPv4 and IPv6 Leasepools. See [Ipaddress Type Ipv4 Ipv6 Leasepool ](#ipaddress-type-ipv4-ipv6-leasepool) below for details.(Deprecated)

`ipv4_leasepool` - (Optional) Select or create new IPv4 Leasepools. See [ref](#ref) below for details.

`ipv6_leasepool` - (Optional) Select or create new IPv4 Leasepools. See [ref](#ref) below for details.(Deprecated)

### Policy

The name of the ZTNA profile.

`policy_name` - (Optional) Select/Add ZTNA Policy to associate with this ZeroTrust Domain. See [ref](#ref) below for details.

### Cloud Gateway Choice All Cloud

Advertise on all Cloud Gateways.

### Cloud Gateway Choice Re Sites

Advertise on selected Cloud Gateways.

`cloud_gateways` - (Optional) Cloud Gateways. See [ref](#ref) below for details.

### Gateways Perimeter Re

Cloud Gateways and Big-IP Edge Gateways.

###### One of the arguments from this list "all_cloud, re_sites" can be set

`all_cloud` - (Optional) Advertise on all Cloud Gateways (`Bool`).

`re_sites` - (Optional) Advertise on selected Cloud Gateways. See [Cloud Gateway Choice Re Sites ](#cloud-gateway-choice-re-sites) below for details.(Deprecated)

### Ipaddress Type Ipv4 Ipv6 Leasepool

Select or create new IPv4 and IPv6 Leasepools.

`ipv4_leasepool` - (Required) Select or create new IPv4 Leasepools. See [ref](#ref) below for details.

`ipv6_leasepool` - (Required) Select or create new IPv4 Lease Pools. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_domain_view.
