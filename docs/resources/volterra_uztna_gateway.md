---

page_title: "Volterra: uztna_gateway"
description: "The uztna_gateway allows CRUD of Uztna Gateway resource on Volterra SaaS"

---

Resource volterra_uztna_gateway
===============================

The Uztna Gateway allows CRUD of Uztna Gateway resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Gateway API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-gateway) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_gateway" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  big_ip_instance {
    bigip_site {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  listeners {
    // One of the arguments from this list "ipv4 ipv6" can be set

    ipv4 = "ipv4"
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

`big_ip_instance` - (Required) BIG-IP Instance.. See [Big Ip Instance ](#big-ip-instance) below for details.

`listeners` - (Required) BIG-IP Edge Gateway Listener.. See [Listeners ](#listeners) below for details.

### Big Ip Instance

BIG-IP Instance..

`bigip_site` - (Required) Selected BIG-IP Instance. See [ref](#ref) below for details.

### Listeners

BIG-IP Edge Gateway Listener..

###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address (`String`).

`ipv6` - (Optional) IPv6 Address (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_gateway.
