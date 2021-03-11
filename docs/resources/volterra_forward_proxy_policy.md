---

page_title: "Volterra: forward_proxy_policy"

description: "The forward_proxy_policy allows CRUD of Forward Proxy Policy resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------

Resource volterra_forward_proxy_policy
======================================

The Forward Proxy Policy allows CRUD of Forward Proxy Policy resource on Volterra SaaS

~> **Note:** Please refer to [Forward Proxy Policy API docs](https://volterra.io/docs/api/views-forward-proxy-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_forward_proxy_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "any_proxy network_connector proxy_label_selector drp_http_connect" must be set
  any_proxy = true

  // One of the arguments from this list "deny_list rule_list allow_all allow_list" must be set

  deny_list {
    // One of the arguments from this list "default_action_next_policy default_action_deny default_action_allow" must be set
    default_action_next_policy = true

    dest_list {
      port_ranges = "80,443,8080-8191,9080"

      prefixes = ["prefixes"]
    }

    http_list {
      // One of the arguments from this list "regex_value exact_value suffix_value" must be set
      suffix_value = "xyz.com"

      // One of the arguments from this list "path_exact_value path_prefix_value path_regex_value any_path" must be set
      path_prefix_value = "/abc/xyz/"
    }

    metadata {
      description = "Virtual Host for acmecorp website"
      disable     = true
      name        = "acmecorp-web"
    }

    rule_description = "Rule to block example.com"
    rule_name        = "my-policy-allow-github.com"

    tls_list {
      // One of the arguments from this list "exact_value suffix_value regex_value" must be set
      exact_value = "abc.zyz.com"
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

`any_proxy` - (Optional) This policy is applied to all forward proxies on this site, and not drp/http-connect proxies (bool).

`drp_http_connect` - (Optional) This policy is applied to attached DRP/HTTP-Connect Proxy (applicable only in App namespace) (bool).

`network_connector` - (Optional) Proxy for given network connector. See [ref](#ref) below for details.

`proxy_label_selector` - (Optional) Proxy for Network Connector or HTTP connect proxy selected by Label selector. See [Proxy Label Selector ](#proxy-label-selector) below for details.

`allow_all` - (Optional) Allow all connections through this forward proxy (bool).

`allow_list` - (Optional) List of allowed connections. See [Allow List ](#allow-list) below for details.

`deny_list` - (Optional) List of denied connections. See [Deny List ](#deny-list) below for details.

`rule_list` - (Optional) List of custom rules. See [Rule List ](#rule-list) below for details.

### All Destinations

Match on all destinations.

### All Sources

Any source that matches 0/0 ip prefix.

### Allow List

List of allowed connections.

`default_action_allow` - (Optional) Allow all connections (bool).

`default_action_deny` - (Optional) Deny all connections (bool).

`default_action_next_policy` - (Optional) Evaluate the next forward proxy policy in the active list (bool).

`dest_list` - (Optional) L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI. See [Dest List ](#dest-list) below for details.

`http_list` - (Optional) URLs for HTTP connections. See [Http List ](#http-list) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`rule_description` - (Optional) Human readable description for the rule (`String`).

`rule_name` - (Optional) Rule Name that will be used to query metrics for this rule. (`String`).

`tls_list` - (Optional) Domains in SNI for TLS connections. See [Tls List ](#tls-list) below for details.

### Any Path

All paths are considered match.

### Default Action Allow

Allow all connections.

### Default Action Deny

Deny all connections.

### Default Action Next Policy

Evaluate the next forward proxy policy in the active list.

### Deny List

List of denied connections.

`default_action_allow` - (Optional) Allow all connections (bool).

`default_action_deny` - (Optional) Deny all connections (bool).

`default_action_next_policy` - (Optional) Evaluate the next forward proxy policy in the active list (bool).

`dest_list` - (Optional) L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI. See [Dest List ](#dest-list) below for details.

`http_list` - (Optional) URLs for HTTP connections. See [Http List ](#http-list) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`rule_description` - (Optional) Human readable description for the rule (`String`).

`rule_name` - (Optional) Rule Name that will be used to query metrics for this rule. (`String`).

`tls_list` - (Optional) Domains in SNI for TLS connections. See [Tls List ](#tls-list) below for details.

### Dest List

L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

`prefixes` - (Required) Destination IPv4 prefixes. (`String`).

### Dst Asn List

The ASN is obtained by performing a lookup for the destination IPv4 Address in a GeoIP DB..

`as_numbers` - (Optional) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. (`Int`).

### Dst Label Selector

Destination is the set of prefixes determined by the label selector expression.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Dst Prefix List

Addresses that are covered by the given list of IPv4 prefixes.

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Http List

URLs for HTTP connections.

`exact_value` - (Optional) Exact domain name (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain names e.g "xyz.com" will match "*.xyz.com" (`String`).

`any_path` - (Optional) All paths are considered match (bool).

`path_exact_value` - (Optional) Exact Path to match. (`String`).

`path_prefix_value` - (Optional) Prefix of Path e.g "/abc/xyz" will match "/abc/xyz/.*" (`String`).

`path_regex_value` - (Optional) Regular Expression value for the Path to match (`String`).

### Inside Sources

All ip prefixes that are reachable via inside interfaces are chosen as Endpoints.

### Label Selector

Sources is set of prefixes determined by label selector expression.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### No Http Connect Port

Ignore destination ports for connections.

### Port Matcher

In case of an HTTP Connect, the destination port is extracted from the connect destination..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`ports` - (Required) to be part of the range. (`String`).

### Prefix List

list of ip prefixes that are representing source of traffic seen by proxy.

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Proxy Label Selector

Proxy for Network Connector or HTTP connect proxy selected by Label selector.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rule List

List of custom rules.

`rules` - (Optional) List of custom rules. See [Rules ](#rules) below for details.

### Rules

List of custom rules.

`action` - (Required) Action to be enforced if the input request matches the rule. (`String`).

`all_destinations` - (Optional) Match on all destinations (bool).

`dst_asn_list` - (Optional) The ASN is obtained by performing a lookup for the destination IPv4 Address in a GeoIP DB.. See [Dst Asn List ](#dst-asn-list) below for details.

`dst_asn_set` - (Optional) The ASN is obtained by performing a lookup for the destination IPv4 Address in a GeoIP DB.. See [ref](#ref) below for details.

`dst_ip_prefix_set` - (Optional) Addresses that are covered by the prefixes in the given ip_prefix_set. See [ref](#ref) below for details.

`dst_label_selector` - (Optional) Destination is the set of prefixes determined by the label selector expression. See [Dst Label Selector ](#dst-label-selector) below for details.

`dst_prefix_list` - (Optional) Addresses that are covered by the given list of IPv4 prefixes. See [Dst Prefix List ](#dst-prefix-list) below for details.

`http_list` - (Optional) URLs for HTTP connections. See [Http List ](#http-list) below for details.

`tls_list` - (Optional) Domains in SNI for TLS connections. See [Tls List ](#tls-list) below for details.

`no_http_connect_port` - (Optional) Ignore destination ports for connections (bool).

`port_matcher` - (Optional) In case of an HTTP Connect, the destination port is extracted from the connect destination.. See [Port Matcher ](#port-matcher) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`rule_description` - (Optional) Human readable description for the rule (`String`).

`rule_name` - (Optional) Rule Name that will be used to query metrics for this rule. (`String`).

`all_sources` - (Optional) Any source that matches 0/0 ip prefix (bool).

`inside_sources` - (Optional) All ip prefixes that are reachable via inside interfaces are chosen as Endpoints (bool).

`interface` - (Optional) All ip prefixes that are reachable via an interfaces are chosen as Endpoints. See [ref](#ref) below for details.

`ip_prefix_set` - (Optional) All ip prefixes that are in a given ip prefix set.. See [ref](#ref) below for details.

`label_selector` - (Optional) Sources is set of prefixes determined by label selector expression. See [Label Selector ](#label-selector) below for details.

`namespace` - (Optional) All ip prefixes that are of a namespace are chosen as Endpoints (`String`).

`prefix_list` - (Optional) list of ip prefixes that are representing source of traffic seen by proxy. See [Prefix List ](#prefix-list) below for details.

### Tls List

Domains in SNI for TLS connections.

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured forward_proxy_policy.
