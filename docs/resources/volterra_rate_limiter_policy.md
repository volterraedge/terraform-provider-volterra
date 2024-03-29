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

`any_server` - (Optional) Any Server (`Bool`).(Deprecated)

`server_name` - (Optional) The expected name of the server. The actual names for the server are extracted from the HTTP Host header and the name of the virtual_host for the request. (`String`).(Deprecated)

`server_name_matcher` - (Optional) regular expressions.. See [Server Choice Server Name Matcher ](#server-choice-server-name-matcher) below for details.(Deprecated)

`server_selector` - (Optional) true if the expressions in the label selector are true for the server labels.. See [Server Choice Server Selector ](#server-choice-server-selector) below for details.(Deprecated)

### Rules

A list of RateLimiterRules that are evaluated sequentially till a matching rule is identified..

`metadata` - (Required) Common attributes for the rule including name and description.. See [Rules Metadata ](#rules-metadata) below for details.

`spec` - (Required) Specification for the rule including match preicates and actions.. See [Rules Spec ](#rules-spec) below for details.

### Action Choice Apply Rate Limiter

Apply the rate limiter configured on the HTTP loadbalancer..

### Action Choice Bypass Rate Limiter

Bypass the rate limiter configured on the HTTP loadbalancer..

### Asn Choice Any Asn

any_asn.

### Asn Choice Asn List

asn_list.

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. (`Int`).

### Asn Choice Asn Matcher

asn_matcher.

`asn_sets` - (Required) A list of references to bgp_asn_set objects.. See [ref](#ref) below for details.

### Country Choice Any Country

x-displayName: "Any Country".

### Country Choice Country List

x-displayName: "Country List".

`country_codes` - (Required) List of Country Codes (`List of Strings`).

`invert_match` - (Optional) Invert the match result. (`Bool`).

### Ip Choice Any Ip

any_ip.

### Ip Choice Ip Matcher

ip_matcher.

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_sets` - (Required) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Ip Choice Ip Prefix List

ip_prefix_list.

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

### Match Check Not Present

Check that the header is not present..

### Match Check Present

Check that the header is present..

### Match Item

Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Rules Spec

Specification for the rule including match preicates and actions..

###### One of the arguments from this list "bypass_rate_limiter, apply_rate_limiter, custom_rate_limiter" must be set

`apply_rate_limiter` - (Optional) Apply the rate limiter configured on the HTTP loadbalancer. (`Bool`).

`bypass_rate_limiter` - (Optional) Bypass the rate limiter configured on the HTTP loadbalancer. (`Bool`).

`custom_rate_limiter` - (Optional) Apply a custom rate limiter.. See [ref](#ref) below for details.

###### One of the arguments from this list "any_asn, asn_list, asn_matcher" can be set

`any_asn` - (Optional)any_asn (`Bool`).

`asn_list` - (Optional)asn_list. See [Asn Choice Asn List ](#asn-choice-asn-list) below for details.

`asn_matcher` - (Optional)asn_matcher. See [Asn Choice Asn Matcher ](#asn-choice-asn-matcher) below for details.

###### One of the arguments from this list "any_country, country_list" must be set

`any_country` - (Optional) x-displayName: "Any Country" (`Bool`).

`country_list` - (Optional) x-displayName: "Country List". See [Country Choice Country List ](#country-choice-country-list) below for details.

`domain_matcher` - (Optional)domain_matcher. See [Spec Domain Matcher ](#spec-domain-matcher) below for details.

`headers` - (Optional)headers. See [Spec Headers ](#spec-headers) below for details.

`http_method` - (Optional)http_method. See [Spec Http Method ](#spec-http-method) below for details.

###### One of the arguments from this list "any_ip, ip_prefix_list, ip_matcher" can be set

`any_ip` - (Optional)any_ip (`Bool`).

`ip_matcher` - (Optional)ip_matcher. See [Ip Choice Ip Matcher ](#ip-choice-ip-matcher) below for details.

`ip_prefix_list` - (Optional)ip_prefix_list. See [Ip Choice Ip Prefix List ](#ip-choice-ip-prefix-list) below for details.

`path` - (Optional)path. See [Spec Path ](#spec-path) below for details.

### Server Choice Server Name Matcher

regular expressions..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Server Choice Server Selector

true if the expressions in the label selector are true for the server labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Spec Domain Matcher

domain_matcher.

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Spec Headers

headers.

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

###### One of the arguments from this list "presence, check_present, check_not_present, item" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the header is present or absent. (`Bool`).(Deprecated)

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Spec Http Method

http_method.

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`methods` - (Optional) x-example: "['GET', 'POST', 'DELETE']" (`List of Strings`).

### Spec Path

path.

`exact_values` - (Optional) A list of exact path values to match the input HTTP path against. (`String`).

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_values` - (Optional) A list of path prefix values to match the input HTTP path against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input HTTP path against. (`String`).

`suffix_values` - (Optional) A list of path suffix values to match the input HTTP path against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured rate_limiter_policy.
