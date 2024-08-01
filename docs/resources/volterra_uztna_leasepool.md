











---
page_title: "Volterra: uztna_leasepool"
description: "The uztna_leasepool allows CRUD of Uztna Leasepool  resource on Volterra SaaS"
---
# Resource volterra_uztna_leasepool

The Uztna Leasepool  allows CRUD of Uztna Leasepool  resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Leasepool  API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-leasepool) to learn more

## Example Usage

```hcl
resource "volterra_uztna_leasepool" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  lease_pool {
    end_address {
      // One of the arguments from this list "ipv6 ipv4" can be set

      ipv4 {
        addr = "192.168.1.1"
      }
    }

    start_address {
      // One of the arguments from this list "ipv6 ipv4" can be set

      ipv4 {
        addr = "192.168.1.1"
      }
    }
  }

  network {
    // One of the arguments from this list "ipv4 ipv6" must be set

    ipv4 {
      plen = "24"

      prefix = "192.168.1.0"
    }
  }
}

```

## Argument Reference

### Metadata Argument Reference
`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).


`description` - (Optional) Human readable description for the object (`String`).


`disable` - (Optional) A value of true will administratively disable the object (`Bool`).


`labels` - (Optional) by selector expression (`String`).


`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).


`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).



### Spec Argument Reference

`lease_pool` - (Required) IP address ranges for the lease pool . See [Lease Pool ](#lease-pool) below for details.


		




		





		





		











`network` - (Required) IPv4 or IPv6 network addresses. See [Network ](#network) below for details.




		






		







### Lease Pool 

 IP address ranges for the lease pool .

`end_address` - (Required) End Address of Lease Pool member. See [Lease Pool End Address ](#lease-pool-end-address) below for details.

`start_address` - (Required) startAddress of Lease Pool member. See [Lease Pool Start Address ](#lease-pool-start-address) below for details.



### Network 

 IPv4 or IPv6 network addresses.



###### One of the arguments from this list "ipv4, ipv6" must be set

`ipv4` - (Optional) IPv4 Subnet Address. See [Ver Ipv4 ](#ver-ipv4) below for details.


`ipv6` - (Optional) IPv6 Subnet Address. See [Ver Ipv6 ](#ver-ipv6) below for details.




### Lease Pool End Address 

 End Address of Lease Pool member.




###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.


`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.




### Lease Pool Start Address 

 startAddress of Lease Pool member.




###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.


`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.




### Ver Ipv4 

 IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).



### Ver Ipv4 

 IPv4 Subnet Address.

`plen` - (Optional) Prefix-length of the IPv4 subnet. Must be <= 32 (`Int`).

`prefix` - (Optional) Prefix part of the IPv4 subnet in string form with dot-decimal notation (`String`).



### Ver Ipv6 

 IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).



### Ver Ipv6 

 IPv6 Subnet Address.

`plen` - (Optional) Prefix length of the IPv6 subnet. Must be <= 128 (`Int`).

`prefix` - (Optional) e.g. "2001:db8::2::" (`String`).



## Attribute Reference

* `id` - This is the id of the configured uztna_leasepool.

