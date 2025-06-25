---

page_title: "Volterra: bgp_routing_policy"

description: "The bgp_routing_policy allows CRUD of Bgp Routing Policy resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------

Resource volterra_bgp_routing_policy
====================================

The Bgp Routing Policy allows CRUD of Bgp Routing Policy resource on Volterra SaaS

~> **Note:** Please refer to [Bgp Routing Policy API docs](https://docs.cloud.f5.com/docs-v2/api/bgp-routing-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_bgp_routing_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  rules {
    action {
      // One of the arguments from this list "aggregate allow as_path community deny local_preference metric" can be set

      as_path = "as_path"
    }

    match {
      // One of the arguments from this list "as_path community ip_prefixes" must be set

      ip_prefixes {
        prefixes {
          ip_prefixes = "ip_prefixes"

          // One of the arguments from this list "equal_or_longer_than exact_match longer_than" can be set

          exact_match = true
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

`rules` - (Required) of rules is critical as rules are applied top to bottom.. See [Rules ](#rules) below for details.

### Rules

of rules is critical as rules are applied top to bottom..

`action` - (Optional) Actions to be applied if route matches the rule. See [Rules Action ](#rules-action) below for details.

`match` - (Optional) Predicates which have to match information in route for action to be applied. See [Rules Match ](#rules-match) below for details.

### Action Type Aggregate

without sharing the more specific prefix.

### Action Type Allow

Allow routes being shared.

### Action Type Community

policies.

`community` - (Required) An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being value (`String`).

### Action Type Deny

Generally used for route filtering purposes.

### Ip Prefixes Prefixes

List of IP prefix.

`ip_prefixes` - (Optional) IP prefix to match on BGP route (`String`).

###### One of the arguments from this list "equal_or_longer_than, exact_match, longer_than" can be set

`equal_or_longer_than` - (Optional) match the 10.1.1.0/20. (`Bool`).

`exact_match` - (Optional) Match the exact prefix length specified in IP prefix (`Bool`).

`longer_than` - (Optional) 10.1.2.0/24, 10.1.3.0/24. This wouldn’t match the 10.1.1.0/20. (`Bool`).

### Prefix Length Match Equal Or Longer Than

match the 10.1.1.0/20..

### Prefix Length Match Exact Match

Match the exact prefix length specified in IP prefix.

### Prefix Length Match Longer Than

10.1.2.0/24, 10.1.3.0/24. This wouldn’t match the 10.1.1.0/20..

### Rules Action

Actions to be applied if route matches the rule.

###### One of the arguments from this list "aggregate, allow, as_path, community, deny, local_preference, metric" can be set

`aggregate` - (Optional) without sharing the more specific prefix (`Bool`).

`allow` - (Optional) Allow routes being shared (`Bool`).

`as_path` - (Optional) AS-Path Prepending is generally used to influence incoming traffic (`String`).

`community` - (Optional) policies. See [Action Type Community ](#action-type-community) below for details.

`deny` - (Optional) Generally used for route filtering purposes (`Bool`).

`local_preference` - (Optional) BGP Local Preference is generally used to influence outgoing traffic (`Int`).

`metric` - (Optional) The Multi-Exit Discriminator metric to indicate the preferred path to AS (`Int`).

### Rules Match

Predicates which have to match information in route for action to be applied.

###### One of the arguments from this list "as_path, community, ip_prefixes" must be set

`as_path` - (Optional) information (`String`).

`community` - (Optional) List of BGP communities to match in route information. See [Type Of Match Community ](#type-of-match-community) below for details.

`ip_prefixes` - (Optional) Select a prefix or group of prefix. See [Type Of Match Ip Prefixes ](#type-of-match-ip-prefixes) below for details.

### Type Of Match Community

List of BGP communities to match in route information.

`community` - (Required) An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being value (`String`).

### Type Of Match Ip Prefixes

Select a prefix or group of prefix.

`prefixes` - (Required) List of IP prefix. See [Ip Prefixes Prefixes ](#ip-prefixes-prefixes) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured bgp_routing_policy.
