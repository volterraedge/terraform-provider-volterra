---

page_title: "Volterra: dns_lb_pool"
description: "The dns_lb_pool allows CRUD of Dns Lb Pool resource on Volterra SaaS"

---

Resource volterra_dns_lb_pool
=============================

The Dns Lb Pool allows CRUD of Dns Lb Pool resource on Volterra SaaS

~> **Note:** Please refer to [Dns Lb Pool API docs](https://docs.cloud.f5.com/docs-v2/api/dns-lb-pool) to learn more

Example Usage
-------------

```hcl
resource "volterra_dns_lb_pool" "example" {
  name                = "acmecorp-web"
  namespace           = "staging"
  load_balancing_mode = ["load_balancing_mode"]

  // One of the arguments from this list "a_pool aaaa_pool cname_pool mx_pool srv_pool" must be set

  a_pool {
    // One of the arguments from this list "disable_health_check health_check" must be set

    disable_health_check = true

    max_answers = "1"

    members {
      disable = true

      ip_endpoint = "8.8.8.8"

      name = "web server 1"

      priority = "10"

      ratio = "10"
    }
  }

  // One of the arguments from this list "ttl use_rrset_ttl" must be set

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

###### One of the arguments from this list "a_pool, aaaa_pool, cname_pool, mx_pool, srv_pool" must be set

`a_pool` - (Optional) x-displayName: "A". See [Pool Type Choice A Pool ](#pool-type-choice-a-pool) below for details.

`aaaa_pool` - (Optional) x-displayName: "AAAA". See [Pool Type Choice Aaaa Pool ](#pool-type-choice-aaaa-pool) below for details.

`cname_pool` - (Optional) x-displayName: "CNAME". See [Pool Type Choice Cname Pool ](#pool-type-choice-cname-pool) below for details.

`mx_pool` - (Optional) x-displayName: "MX". See [Pool Type Choice Mx Pool ](#pool-type-choice-mx-pool) below for details.

`srv_pool` - (Optional) x-displayName: "SRV". See [Pool Type Choice Srv Pool ](#pool-type-choice-srv-pool) below for details.

###### One of the arguments from this list "ttl, use_rrset_ttl" must be set

`ttl` - (Optional) Custom TTL in seconds (default 30) for responses from this pool (`Int`).

`use_rrset_ttl` - (Optional) Use TTL specified in the RRSet of the DNS Load Balancer Record which uses the pool (`Bool`).

### A Pool Members

x-required.

`disable` - (Optional) A value of true will disable the pool-member (`Bool`).

`ip_endpoint` - (Required) Public IP address (`String`).

`name` - (Optional) Pool member name (`String`).

`priority` - (Optional) Used if the pool’s load balancing mode is set to Priority (`Int`).

`ratio` - (Optional) Used if the pool’s load balancing mode is set to Ratio-Member (`Int`).

### Aaaa Pool Members

x-required.

`disable` - (Optional) A value of true will disable the pool-member (`Bool`).

`ip_endpoint` - (Required) Public IP address (`String`).

`name` - (Optional) Pool member name (`String`).

`priority` - (Optional) Used if the pool’s load balancing mode is set to Priority (`Int`).

`ratio` - (Optional) Used if the pool’s load balancing mode is set to Ratio-Member (`Int`).

### Cname Pool Members

x-required.

`domain` - (Required) x-required (`String`).

`final_translation` - (Optional) If this flag is true, the CNAME record will not be translated further. (`Bool`).

`name` - (Optional) Pool member name (`String`).

`ratio` - (Optional) Ratio (`Int`).

### Health Check Choice Disable Health Check

When health check is disabled, the pool member is presumed to be always healthy.

### Mx Pool Members

x-required.

`domain` - (Required) x-required (`String`).

`name` - (Optional) Pool member name (`String`).

`priority` - (Optional) MX Record priority (`Int`).

`ratio` - (Optional) Load Balancing Ratio (`Int`).

### Pool Type Choice A Pool

x-displayName: "A".

###### One of the arguments from this list "disable_health_check, health_check" must be set

`disable_health_check` - (Optional) When health check is disabled, the pool member is presumed to be always healthy (`Bool`).

`health_check` - (Optional) Select the health check to be applied to all the pool members. See [ref](#ref) below for details.

`max_answers` - (Required) Limit on number of Resource Records to be included in the response to query (`Int`).

`members` - (Required) x-required. See [A Pool Members ](#a-pool-members) below for details.

### Pool Type Choice Aaaa Pool

x-displayName: "AAAA".

`max_answers` - (Required) Limit on number of Resource Records to be included in the response to query (`Int`).

`members` - (Required) x-required. See [Aaaa Pool Members ](#aaaa-pool-members) below for details.

### Pool Type Choice Cname Pool

x-displayName: "CNAME".

`members` - (Required) x-required. See [Cname Pool Members ](#cname-pool-members) below for details.

### Pool Type Choice Mx Pool

x-displayName: "MX".

`max_answers` - (Required) Limit on number of Resource Records to be included in the response to query (`Int`).

`members` - (Required) x-required. See [Mx Pool Members ](#mx-pool-members) below for details.

### Pool Type Choice Srv Pool

x-displayName: "SRV".

`max_answers` - (Required) Limit on number of Resource Records to be included in the response to query (`Int`).

`members` - (Required) x-required. See [Srv Pool Members ](#srv-pool-members) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Srv Pool Members

x-required.

`final_translation` - (Optional) If this flag is true, the SRV record will not be translated further. (`Bool`).

`name` - (Optional) Pool member name (`String`).

`port` - (Required) x-required (`Int`).

`priority` - (Required) x-required (`Int`).

`ratio` - (Optional) Ratio (`Int`).

`target` - (Required) x-required (`String`).

`weight` - (Required) x-required (`Int`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured dns_lb_pool.
