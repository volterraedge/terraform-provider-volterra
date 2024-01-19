---

page_title: "Volterra: rate_limiter_policy"

description: "The rate_limiter_policy allows CRUD of Rate Limiter Policy resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------

Resource volterra_rate_limiter_policy
=====================================

The Rate Limiter Policy allows CRUD of Rate Limiter Policy resource on Volterra SaaS

~> **Note:** Please refer to [Rate Limiter Policy API docs](https://docs.cloud.f5.com/docs/api/views-rate-limiter-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_rate_limiter_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
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

`rules` - (Optional) A list of RateLimiterRules that are evaluated sequentially till a matching rule is identified.. See [Rules ](#rules) below for details.

`any_server` - (Optional) Any Server (bool).

`server_name` - (Optional) The expected name of the server. The actual names for the server are extracted from the HTTP Host header and the name of the virtual_host for the request. (`String`).

`server_name_matcher` - (Optional) regular expressions.. See [Server Name Matcher ](#server-name-matcher) below for details.

`server_selector` - (Optional) true if the expressions in the label selector are true for the server labels.. See [Server Selector ](#server-selector) below for details.

### Any Asn

any_asn.

### Any Country

x-displayName: "Any Country".

### Any Ip

any_ip.

### Apply Rate Limiter

Apply the rate limiter configured on the HTTP loadbalancer..

### Asn List

asn_list.

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. (`Int`).

### Asn Matcher

asn_matcher.

`asn_sets` - (Required) A list of references to bgp_asn_set objects.. See [ref](#ref) below for details.

### Bypass Rate Limiter

Bypass the rate limiter configured on the HTTP loadbalancer..

### Check Not Present

Check that the header is not present..

### Check Present

Check that the header is present..

### Country List

x-displayName: "Country List".

`country_codes` - (Required) List of Country Codes (`List of Strings`).

`invert_match` - (Optional) Invert the match result. (`Bool`).

### Domain Matcher

domain_matcher.

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Headers

headers.

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`check_not_present` - (Optional) Check that the header is not present. (bool).

`check_present` - (Optional) Check that the header is present. (bool).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the header is present or absent. (`Bool`).

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Http Method

http_method.

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`methods` - (Optional) x-example: "['GET', 'POST', 'DELETE']" (`List of Strings`).

### Ip Matcher

ip_matcher.

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_sets` - (Required) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Ip Prefix List

ip_prefix_list.

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

### Item

Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Path

path.

`exact_values` - (Optional) A list of exact path values to match the input HTTP path against. (`String`).

`prefix_values` - (Optional) A list of path prefix values to match the input HTTP path against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input HTTP path against. (`String`).

`suffix_values` - (Optional) A list of path suffix values to match the input HTTP path against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rules

A list of RateLimiterRules that are evaluated sequentially till a matching rule is identified..

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`spec` - (Required) Specification for the rule including match preicates and actions.. See [Spec ](#spec) below for details.

### Server Name Matcher

regular expressions..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Server Selector

true if the expressions in the label selector are true for the server labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Spec

Specification for the rule including match preicates and actions..

`apply_rate_limiter` - (Optional) Apply the rate limiter configured on the HTTP loadbalancer. (bool).

`bypass_rate_limiter` - (Optional) Bypass the rate limiter configured on the HTTP loadbalancer. (bool).

`custom_rate_limiter` - (Optional) Apply a custom rate limiter.. See [ref](#ref) below for details.

`any_asn` - (Optional)any_asn (bool).

`asn_list` - (Optional)asn_list. See [Asn List ](#asn-list) below for details.

`asn_matcher` - (Optional)asn_matcher. See [Asn Matcher ](#asn-matcher) below for details.

`any_country` - (Optional) x-displayName: "Any Country" (bool).

`country_list` - (Optional) x-displayName: "Country List". See [Country List ](#country-list) below for details.

`domain_matcher` - (Optional)domain_matcher. See [Domain Matcher ](#domain-matcher) below for details.

`headers` - (Optional)headers. See [Headers ](#headers) below for details.

`http_method` - (Optional)http_method. See [Http Method ](#http-method) below for details.

`any_ip` - (Optional)any_ip (bool).

`ip_matcher` - (Optional)ip_matcher. See [Ip Matcher ](#ip-matcher) below for details.

`ip_prefix_list` - (Optional)ip_prefix_list. See [Ip Prefix List ](#ip-prefix-list) below for details.

`path` - (Optional)path. See [Path ](#path) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured rate_limiter_policy.
