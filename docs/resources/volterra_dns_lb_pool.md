---

page_title: "Volterra: dns_lb_pool"

description: "The dns_lb_pool allows CRUD of Dns Lb Pool resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_dns_lb_pool
=============================

The Dns Lb Pool allows CRUD of Dns Lb Pool resource on Volterra SaaS

~> **Note:** Please refer to [Dns Lb Pool API docs](https://docs.cloud.f5.com/docs/api/dns-lb-pool) to learn more

Example Usage
-------------

```hcl
resource "volterra_dns_lb_pool" "example" {
  name                = "acmecorp-web"
  namespace           = "staging"
  load_balancing_mode = ["load_balancing_mode"]

  // One of the arguments from this list "aaaa_pool cname_pool mx_pool a_pool" must be set

  a_pool {
    // One of the arguments from this list "disable_health_check health_check" must be set
    disable_health_check = true
    max_answers          = "1"

    members {
      ip_endpoint = "8.8.8.8"
      priority    = "10"
      ratio       = "10"
    }
  }
  // One of the arguments from this list "use_rrset_ttl ttl" must be set
  use_rrset_ttl = true
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

`load_balancing_mode` - (Required) x-required (`String`).

`a_pool` - (Optional) x-displayName: "A". See [A Pool ](#a-pool) below for details.

`aaaa_pool` - (Optional) x-displayName: "AAAA". See [Aaaa Pool ](#aaaa-pool) below for details.

`cname_pool` - (Optional) x-displayName: "CNAME". See [Cname Pool ](#cname-pool) below for details.

`mx_pool` - (Optional) x-displayName: "MX". See [Mx Pool ](#mx-pool) below for details.

`ttl` - (Optional) Custom TTL in seconds (default 30) for responses from this pool (`Int`).

`use_rrset_ttl` - (Optional) Use TTL specified in the RRSet of the DNS Load Balancer Record which uses the pool (bool).

### A Pool

x-displayName: "A".

`disable_health_check` - (Optional) When health check is disabled, the pool member is presumed to be always healthy (bool).

`health_check` - (Optional) Select the health check to be applied to all the pool members. See [ref](#ref) below for details.

`max_answers` - (Required) Limit on number of Resource Records to be included in the response to query (`Int`).

`members` - (Required) x-required. See [Members ](#members) below for details.

### Aaaa Pool

x-displayName: "AAAA".

`max_answers` - (Required) Limit on number of Resource Records to be included in the response to query (`Int`).

`members` - (Required) x-required. See [Members ](#members) below for details.

### Cname Pool

x-displayName: "CNAME".

`members` - (Required) x-required. See [Members ](#members) below for details.

### Disable Health Check

When health check is disabled, the pool member is presumed to be always healthy.

### Members

x-required.

`ip_endpoint` - (Required) Public IP address (`String`).

`priority` - (Optional) Used if the pool’s load balancing mode is set to Priority (`Int`).

`ratio` - (Optional) Used if the pool’s load balancing mode is set to Ratio-Member (`Int`).

### Mx Pool

x-displayName: "MX".

`max_answers` - (Required) Limit on number of Resource Records to be included in the response to query (`Int`).

`members` - (Required) x-required. See [Members ](#members) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured dns_lb_pool.
