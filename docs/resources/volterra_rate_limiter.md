---

page_title: "Volterra: rate_limiter"

description: "The rate_limiter allows CRUD of Rate Limiter resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_rate_limiter
==============================

The Rate Limiter allows CRUD of Rate Limiter resource on Volterra SaaS

~> **Note:** Please refer to [Rate Limiter API docs](https://docs.cloud.f5.com/docs/api/rate-limiter) to learn more

Example Usage
-------------

```hcl
resource "volterra_rate_limiter" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  limits {
    burst_multiplier = "1"
    total_number     = "1"
    unit             = "unit"
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

`limits` - (Required) A list of RateLimitValues that specifies the total number of allowed requests for each specified period.. See [Limits ](#limits) below for details.

`user_identification` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier to be rate limited.. See [ref](#ref) below for details.

### Limits

A list of RateLimitValues that specifies the total number of allowed requests for each specified period..

`burst_multiplier` - (Optional) The maximum burst of requests to accommodate, expressed as a multiple of the rate. (`Int`).

`total_number` - (Required) The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period. (`Int`).

`unit` - (Required) Unit for the period per which the rate limit is applied. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured rate_limiter.
