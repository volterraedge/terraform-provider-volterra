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
  name        = "acmecorp-web"
  namespace   = "staging"
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
    // One of the arguments from this list "ipv4_ipv6_leasepool ipv4_leasepool ipv6_leasepool" must be set

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

  vip_dns_proxy {
    // One of the arguments from this list "ipv4 ipv6" can be set

    ipv4 {
      addr = "192.168.1.1"
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

`cdn_ce_vh_api_gw` - (Optional) Internal reference to cdn ce api gateway VH. See [ref](#ref) below for details.(Deprecated)

`cert` - (Required) Domain in XC is an established pattern and we would reuse the same.. See [Cert ](#cert) below for details.

`gateways` - (Required) List of all Cloud Gateways and Big-IP Edge Gateways. See [Gateways ](#gateways) below for details.

`lease_pool` - (Required) The Lease Pool List assigned to the Zero Trust Domain. . See [Lease Pool ](#lease-pool) below for details.

`policy` - (Required) This is used to import or create new ZTNA Policy. See [Policy ](#policy) below for details.

`vip_dns_proxy` - (Required) VIP address for internal DNS proxy. See [Vip Dns Proxy ](#vip-dns-proxy) below for details.(Deprecated)

### App Vip Pool

Appliacation VIP Pools .

`app_vip_pool` - (Optional) Selected VIP Pools. See [ref](#ref) below for details.

### Cert

Domain in XC is an established pattern and we would reuse the same..

`certificate` - (Optional) Select/Add one or more TLS Certificate objects to associate with this ZeroTrust Domain. See [ref](#ref) below for details.

### Gateways

List of all Cloud Gateways and Big-IP Edge Gateways.

`bigip_site` - (Optional) From the available Big-IP Edge Gateway List select Big-IP Edge Gateway.. See [ref](#ref) below for details.

`perimeter_re` - (Optional) List of all Cloud Gateways and Big-IP Edge Gateways. See [Gateways Perimeter Re ](#gateways-perimeter-re) below for details.

### Lease Pool

The Lease Pool List assigned to the Zero Trust Domain. .

###### One of the arguments from this list "ipv4_ipv6_leasepool, ipv4_leasepool, ipv6_leasepool" must be set

`ipv4_ipv6_leasepool` - (Optional) IPv4 and IPv6 Lease Pools. See [Ipaddress Type Ipv4 Ipv6 Leasepool ](#ipaddress-type-ipv4-ipv6-leasepool) below for details.(Deprecated)

`ipv4_leasepool` - (Required) IPv4 Lease Pools. See [Ipaddress Type Ipv4 Leasepool ](#ipaddress-type-ipv4-leasepool) below for details.

`ipv6_leasepool` - (Required) IPv4 Lease Pools. See [Ipaddress Type Ipv6 Leasepool ](#ipaddress-type-ipv6-leasepool) below for details.(Deprecated)

### Policy

This is used to import or create new ZTNA Policy.

`policy` - (Optional) Select/Add ZTNA Policy to associate with this ZeroTrust Domain. See [ref](#ref) below for details.

### Vip Dns Proxy

VIP address for internal DNS proxy.

###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

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

### Ipaddress Type Ipv4 Ipv6 Leasepool

IPv4 and IPv6 Lease Pools.

`ipv4_leasepool` - (Required) Select or create new IPv4 Lease Pools. See [Ipv4 Ipv6 Leasepool Ipv4 Leasepool ](#ipv4-ipv6-leasepool-ipv4-leasepool) below for details.

`ipv6_leasepool` - (Required) Select or create new IPv4 Lease Pools. See [Ipv4 Ipv6 Leasepool Ipv6 Leasepool ](#ipv4-ipv6-leasepool-ipv6-leasepool) below for details.

### Ipaddress Type Ipv4 Leasepool

IPv4 Lease Pools.

`ipv4_leasepool` - (Required) Select or create new IPv4 Lease Pool. See [ref](#ref) below for details.

### Ipaddress Type Ipv6 Leasepool

IPv4 Lease Pools.

`ipv6_leasepool` - (Required) Select or create new IPv6 Lease Pools. See [ref](#ref) below for details.

### Ipv4 Ipv6 Leasepool Ipv4 Leasepool

Select or create new IPv4 Lease Pools.

`ipv4_leasepool` - (Required) Select or create new IPv4 Lease Pool. See [ref](#ref) below for details.

### Ipv4 Ipv6 Leasepool Ipv6 Leasepool

Select or create new IPv4 Lease Pools.

`ipv6_leasepool` - (Required) Select or create new IPv6 Lease Pools. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Ver Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ver Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_domain.
