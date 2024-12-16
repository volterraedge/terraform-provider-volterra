---

page_title: "Volterra: uztna_leasepool"

description: "The uztna_leasepool allows CRUD of Uztna Leasepool resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_uztna_leasepool
=================================

The Uztna Leasepool allows CRUD of Uztna Leasepool resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Leasepool API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-leasepool) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_leasepool" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  ip_version {
    // One of the arguments from this list "ipv4_vip ipv6_vip" can be set

    ipv4_vip {
      prefix = ["['10.2.1.0/24', '192.168.8.0/29', '10.7.64.160/27']"]

      vip4_range {
        end_address {
          addr = "192.168.1.1"
        }

        start_address {
          addr = "192.168.1.1"
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

`ip_version` - (Required) IP address ranges for the Lease Pools. See [Ip Version ](#ip-version) below for details.

### Ip Version

IP address ranges for the Lease Pools.

###### One of the arguments from this list "ipv4_vip, ipv6_vip" can be set

`ipv4_vip` - (Optional) IPV4 Only Lease Pool. See [Ip Vip Ipv4 Vip ](#ip-vip-ipv4-vip) below for details.

`ipv6_vip` - (Optional) IPV6 Only Lease Pool. See [Ip Vip Ipv6 Vip ](#ip-vip-ipv6-vip) below for details.(Deprecated)

### Ip Vip Ipv4 Vip

IPV4 Only Lease Pool.

`prefix` - (Required) IPV4 Lease Pool Network (`String`).

`vip4_range` - (Optional) IPV4 Lease Pool Range. See [Ipv4 Vip Vip4 Range ](#ipv4-vip-vip4-range) below for details.

### Ip Vip Ipv6 Vip

IPV6 Only Lease Pool.

`ipv6_prefix` - (Optional) IPV6 Lease Pool Network (`String`).

`vip6_range` - (Required) IPV6 Lease Pool Range. See [Ipv6 Vip Vip6 Range ](#ipv6-vip-vip6-range) below for details.

### Ipv4 Vip Vip4 Range

IPV4 Lease Pool Range.

`end_address` - (Optional) IPV4 End Address. See [Vip4 Range End Address ](#vip4-range-end-address) below for details.

`start_address` - (Optional) IPV4 Start Address. See [Vip4 Range Start Address ](#vip4-range-start-address) below for details.

### Ipv6 Vip Vip6 Range

IPV6 Lease Pool Range.

`end_address` - (Optional) IPV6 End Address. See [Vip6 Range End Address ](#vip6-range-end-address) below for details.

`start_address` - (Optional) IPV6 Start Address. See [Vip6 Range Start Address ](#vip6-range-start-address) below for details.

### Vip4 Range End Address

IPV4 End Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Vip4 Range Start Address

IPV4 Start Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Vip6 Range End Address

IPV6 End Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

### Vip6 Range Start Address

IPV6 Start Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_leasepool.
