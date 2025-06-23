---

page_title: "Volterra: uztna_app_vip_pool"

description: "The uztna_app_vip_pool allows CRUD of Uztna App Vip Pool resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------

Resource volterra_uztna_app_vip_pool
====================================

The Uztna App Vip Pool allows CRUD of Uztna App Vip Pool resource on Volterra SaaS

~> **Note:** Please refer to [Uztna App Vip Pool API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-app-vip-pool) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_app_vip_pool" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  ip_version {
    // One of the arguments from this list "ipv4_vip" can be set

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

`ip_version` - (Required) IP address ranges for the Application VIP Pool Range . See [Ip Version ](#ip-version) below for details.

### Ip Version

IP address ranges for the Application VIP Pool Range .

###### One of the arguments from this list "ipv4_vip" can be set

`ipv4_vip` - (Optional) IPV4 Only Appllication VIP Pool. See [Ip Vip Ipv4 Vip ](#ip-vip-ipv4-vip) below for details.

### Ip Vip Ipv4 Vip

IPV4 Only Appllication VIP Pool.

`prefix` - (Optional) IPV4 Application IP Pool Network (`String`).

`vip4_range` - (Optional) IPV4 Application VIP Pool Range. See [Ipv4 Vip Vip4 Range ](#ipv4-vip-vip4-range) below for details.

### Ipv4 Vip Vip4 Range

IPV4 Application VIP Pool Range.

`end_address` - (Optional) IPV4 End Address. See [Vip4 Range End Address ](#vip4-range-end-address) below for details.

`start_address` - (Optional) IPV4 Start Address. See [Vip4 Range Start Address ](#vip4-range-start-address) below for details.

### Vip4 Range End Address

IPV4 End Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Vip4 Range Start Address

IPV4 Start Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_app_vip_pool.
