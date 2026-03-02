---

page_title: "Volterra: cdn_cache_rule"

description: "The cdn_cache_rule allows CRUD of Cdn Cache Rule resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_cdn_cache_rule
================================

The Cdn Cache Rule allows CRUD of Cdn Cache Rule resource on Volterra SaaS

~> **Note:** Please refer to [Cdn Cache Rule API docs](https://docs.cloud.f5.com/docs-v2/api/cdn-cache-rule) to learn more

Example Usage
-------------

```hcl
resource "volterra_cdn_cache_rule" "example" {
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

`cache_rules` - (Optional) Rules are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs.. See [Cache Rules ](#cache-rules) below for details.

### Cache Rules

Rules are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs..

###### One of the arguments from this list "cache_bypass, eligible_for_cache" must be set

`cache_bypass` - (Optional) Do not cache content if rule match occurs (`Bool`).

`eligible_for_cache` - (Optional) Cache content if rule match occurs. See [Cache Actions Eligible For Cache ](#cache-actions-eligible-for-cache) below for details.

`rule_expression_list` - (Required) Expressions are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs... See [Cache Rules Rule Expression List ](#cache-rules-rule-expression-list) below for details.

`rule_name` - (Required) Name of the Cache Rule (`String`).

### Cache Actions Cache Bypass

Do not cache content if rule match occurs.

### Cache Actions Eligible For Cache

Cache content if rule match occurs.

###### One of the arguments from this list "scheme_proxy_host_request_uri, scheme_proxy_host_uri" must be set

`scheme_proxy_host_request_uri` - (Optional) x-displayName: "Scheme + Proxy Host + Request URI". See [Eligible For Cache Scheme Proxy Host Request Uri ](#eligible-for-cache-scheme-proxy-host-request-uri) below for details.

`scheme_proxy_host_uri` - (Optional) x-displayName: "Scheme + Proxy Host + URI". See [Eligible For Cache Scheme Proxy Host Uri ](#eligible-for-cache-scheme-proxy-host-uri) below for details.

### Cache Headers Operator

Available operators.

###### One of the arguments from this list "Contains, DoesNotContain, DoesNotEndWith, DoesNotEqual, DoesNotStartWith, Endswith, Equals, MatchRegex, Startswith" must be set

`Contains` - (Optional) The header value must include the specified value as a substring. (`String`).

`DoesNotContain` - (Optional) The header value must not include the specified value as a substring. (`String`).

`DoesNotEndWith` - (Optional) The header value must not end with the specified value. (`String`).

`DoesNotEqual` - (Optional) The header value must not match the specified value. (`String`).

`DoesNotStartWith` - (Optional) The header value must not begin with the specified value. (`String`).

`Endswith` - (Optional) The header value must end with the specified value. (`String`).

`Equals` - (Optional) The header value must exactly match the specified value. (`String`).

`MatchRegex` - (Optional) The header value must match the specified regular expression pattern. (`String`).

`Startswith` - (Optional) The header value must begin with the specified value. (`String`).

### Cache Rule Expression Cache Headers

Configure cache rule headers to match the criteria.

`name` - (Required) Select the name of the header from the list (`String`).

`operator` - (Optional) Available operators. See [Cache Headers Operator ](#cache-headers-operator) below for details.

### Cache Rule Expression Cookie Matcher

Note that all specified cookie matcher predicates must evaluate to true..

`name` - (Required) Enter the name of the cookie to match (`String`).

`operator` - (Optional) x-displayName: "Operator". See [Cookie Matcher Operator ](#cookie-matcher-operator) below for details.

### Cache Rule Expression Path Match

URI path of route.

`operator` - (Optional) A specification of path match. See [Path Match Operator ](#path-match-operator) below for details.

### Cache Rule Expression Query Parameters

List of (key, value) query parameters.

`key` - (Required) The name of the query parameter to match (`String`).

`operator` - (Optional) x-displayName: "Operator". See [Query Parameters Operator ](#query-parameters-operator) below for details.

### Cache Rules Rule Expression List

Expressions are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs...

`cache_rule_expression` - (Required) The Cache Rule Expression Terms that are ANDed. See [Rule Expression List Cache Rule Expression ](#rule-expression-list-cache-rule-expression) below for details.

`expression_name` - (Required) Name of the Expressions items that are ANDed (`String`).

### Cookie Matcher Operator

x-displayName: "Operator".

###### One of the arguments from this list "Contains, DoesNotContain, DoesNotEndWith, DoesNotEqual, DoesNotStartWith, Endswith, Equals, MatchRegex, Startswith" must be set

`Contains` - (Optional) The cookie value must include the specified value as a substring. (`String`).

`DoesNotContain` - (Optional) The cookie value must not include the specified value as a substring. (`String`).

`DoesNotEndWith` - (Optional) The cookie value must not end with the specified value. (`String`).

`DoesNotEqual` - (Optional) The cookie value must not match the specified value. (`String`).

`DoesNotStartWith` - (Optional) The cookie value must not begin with the specified value. (`String`).

`Endswith` - (Optional) The cookie value must end with the specified value. (`String`).

`Equals` - (Optional) The cookie value must exactly match the specified value. (`String`).

`MatchRegex` - (Optional) The cookie value must match the specified regular expression pattern in PCRE format. (`String`).

`Startswith` - (Optional) The cookie value must begin with the specified value. (`String`).

### Eligible For Cache Scheme Proxy Host Request Uri

x-displayName: "Scheme + Proxy Host + Request URI".

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Eligible For Cache Scheme Proxy Host Uri

x-displayName: "Scheme + Proxy Host + URI".

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Path Match Operator

A specification of path match.

###### One of the arguments from this list "Contains, DoesNotContain, DoesNotEndWith, DoesNotEqual, DoesNotStartWith, Endswith, Equals, MatchRegex, Startswith" must be set

`Contains` - (Optional) The path must include the specified value as a substring, up to the filename. (`String`).

`DoesNotContain` - (Optional) The path must not include the specified value as a substring, up to the filename. (`String`).

`DoesNotEndWith` - (Optional) The path must not end with the specified value, up to the filename. (`String`).

`DoesNotEqual` - (Optional) The path must not match the specified value, up to the filename. (`String`).

`DoesNotStartWith` - (Optional) The path must not begin with the specified value, up to the filename. (`String`).

`Endswith` - (Optional) The path must end with the specified value, up to the filename. (`String`).

`Equals` - (Optional) The path must exactly match the specified value, up to the filename. (`String`).

`MatchRegex` - (Optional) The path must match the specified regular expression pattern in PCRE format. (`String`).

`Startswith` - (Optional) The path must begin with the specified value, up to the filename. (`String`).

### Query Parameters Operator

x-displayName: "Operator".

###### One of the arguments from this list "Contains, DoesNotContain, DoesNotEndWith, DoesNotEqual, DoesNotStartWith, Endswith, Equals, MatchRegex, Startswith" must be set

`Contains` - (Optional) The query parameter value must include the specified value as a substring. (`String`).

`DoesNotContain` - (Optional) The query parameter value must not include the specified value as a substring. (`String`).

`DoesNotEndWith` - (Optional) The query parameter value must not end with the specified value. (`String`).

`DoesNotEqual` - (Optional) The query parameter value must not match the specified value. (`String`).

`DoesNotStartWith` - (Optional) The query parameter value must not begin with the specified value. (`String`).

`Endswith` - (Optional) The query parameter value must end with the specified value. (`String`).

`Equals` - (Optional) The query parameter value must exactly match the specified value. (`String`).

`MatchRegex` - (Optional) The query parameter value must match the specified regular expression pattern in PCRE format. (`String`).

`Startswith` - (Optional) The query parameter value must begin with the specified value. (`String`).

### Rule Expression List Cache Rule Expression

The Cache Rule Expression Terms that are ANDed.

`cache_headers` - (Optional) Configure cache rule headers to match the criteria. See [Cache Rule Expression Cache Headers ](#cache-rule-expression-cache-headers) below for details.

`cookie_matcher` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Cache Rule Expression Cookie Matcher ](#cache-rule-expression-cookie-matcher) below for details.

`path_match` - (Optional) URI path of route. See [Cache Rule Expression Path Match ](#cache-rule-expression-path-match) below for details.

`query_parameters` - (Optional) List of (key, value) query parameters. See [Cache Rule Expression Query Parameters ](#cache-rule-expression-query-parameters) below for details.

Attribute Reference
-------------------

*   `id` - This is the id of the configured cdn_cache_rule.
