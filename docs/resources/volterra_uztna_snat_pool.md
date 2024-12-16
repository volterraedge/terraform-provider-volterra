---

page_title: "Volterra: uztna_snat_pool"

description: "The uztna_snat_pool allows CRUD of Uztna Snat Pool resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_uztna_snat_pool
=================================

The Uztna Snat Pool allows CRUD of Uztna Snat Pool resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Snat Pool API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-snat-pool) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_snat_pool" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  ip_addresses {
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

`ip_addresses` - (Required) Combination of IPV4 and IPV6 can exist in the list.. See [Ip Addresses ](#ip-addresses) below for details.

### Ip Addresses

Combination of IPV4 and IPV6 can exist in the list..

###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Ver Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ver Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_snat_pool.
