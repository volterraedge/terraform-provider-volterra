---

page_title: "Volterra: dns_load_balancer"

description: "The dns_load_balancer allows CRUD of Dns Load Balancer resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_dns_load_balancer
===================================

The Dns Load Balancer allows CRUD of Dns Load Balancer resource on Volterra SaaS

~> **Note:** Please refer to [Dns Load Balancer API docs](https://volterra.io/docs/api/dns-load-balancer) to learn more

Example Usage
-------------

```hcl
resource "volterra_dns_load_balancer" "example" {
  name        = "acmecorp-web"
  namespace   = "staging"
  record_type = ["record_type"]

  rule_list {
    rules {
      // One of the arguments from this list "pool nxdomain" must be set

      pool {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }

      // One of the arguments from this list "geo_location_label_selector geo_location_set" must be set

      geo_location_label_selector {
        expressions = ["region in (us-west1, us-west2),tier in (staging)"]
      }
      score = "50"
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

`record_type` - (Required) x-required (`String`).

`response_cache` - (Optional) Response Cache Parameters. See [Response Cache ](#response-cache) below for details.

`rule_list` - (Required) Load Balancing Rules. See [Rule List ](#rule-list) below for details.

### Default Response Cache Parameters

Default Parameters for caching the DNS responses.

### Disable

When Response Cache is disabled, responses will be computed for each request.

### Geo Location Label Selector

with the translated geo locations derived from incoming EDNS-S0 client-subnet in the DNS request..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Nxdomain

Do not perform any load-balancing. Instead return NXDOMAIN.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Response Cache

Response Cache Parameters.

`default_response_cache_parameters` - (Optional) Default Parameters for caching the DNS responses (bool).

`disable` - (Optional) When Response Cache is disabled, responses will be computed for each request (bool).

`response_cache_parameters` - (Optional) Customize the parameters for Response cache. See [Response Cache Parameters ](#response-cache-parameters) below for details.

### Response Cache Parameters

Customize the parameters for Response cache.

`cache_cidr_ipv4` - (Optional) Length of CIDR masks used to group IPv4 clients (`Int`).

`cache_cidr_ipv6` - (Optional) Length of CIDR masks used to group IPv6 clients (`Int`).

`cache_ttl` - (Optional) TTL for response cache (`Int`).

### Rule List

Load Balancing Rules.

`rules` - (Required) Rules to perform load balancing. See [Rules ](#rules) below for details.

### Rules

Rules to perform load balancing.

`nxdomain` - (Optional) Do not perform any load-balancing. Instead return NXDOMAIN (bool).

`pool` - (Optional) Use this pool for the Load Balancing.. See [ref](#ref) below for details.

`geo_location_label_selector` - (Optional) with the translated geo locations derived from incoming EDNS-S0 client-subnet in the DNS request.. See [Geo Location Label Selector ](#geo-location-label-selector) below for details.

`geo_location_set` - (Optional) with the translated geo locations derived from incoming EDNS-S0 client-subnet in the DNS request.. See [ref](#ref) below for details.

`score` - (Optional) When multiple load balancing rules match a query, the one with the highest score is chosen (`Int`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured dns_load_balancer.
