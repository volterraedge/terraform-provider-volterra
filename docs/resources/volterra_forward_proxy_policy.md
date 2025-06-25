---

page_title: "Volterra: forward_proxy_policy"

description: "The forward_proxy_policy allows CRUD of Forward Proxy Policy resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------

Resource volterra_forward_proxy_policy
======================================

The Forward Proxy Policy allows CRUD of Forward Proxy Policy resource on Volterra SaaS

~> **Note:** Please refer to [Forward Proxy Policy API docs](https://docs.cloud.f5.com/docs-v2/api/views-forward-proxy-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_forward_proxy_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "any_proxy drp_http_connect network_connector proxy_label_selector" must be set

  network_connector {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }

  // One of the arguments from this list "allow_all allow_list deny_list rule_list" must be set

  allow_all = true
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

###### One of the arguments from this list "any_proxy, drp_http_connect, network_connector, proxy_label_selector" must be set

`any_proxy` - (Optional) This policy is applied to all forward proxies on this site, and not drp/http-connect proxies (`Bool`).

`drp_http_connect` - (Optional) This policy is applied to attached DRP/HTTP-Connect Proxy (applicable only in App namespace) (`Bool`).

`network_connector` - (Optional) Proxy for given network connector. See [ref](#ref) below for details.

`proxy_label_selector` - (Optional) Proxy for Network Connector or HTTP connect proxy selected by Label selector. See [Proxy Choice Proxy Label Selector ](#proxy-choice-proxy-label-selector) below for details.

###### One of the arguments from this list "allow_all, allow_list, deny_list, rule_list" must be set

`allow_all` - (Optional) Allow all connections through this forward proxy (`Bool`).

`allow_list` - (Optional) List of allowed connections. See [Rule Choice Allow List ](#rule-choice-allow-list) below for details.

`deny_list` - (Optional) List of denied connections. See [Rule Choice Deny List ](#rule-choice-deny-list) below for details.

`rule_list` - (Optional) List of custom rules. See [Rule Choice Rule List ](#rule-choice-rule-list) below for details.

`segment_policy` - (Optional) Skip the configuration or set option as Any to ignore corresponding segment match. See [Segment Policy ](#segment-policy) below for details.

### Segment Policy

Skip the configuration or set option as Any to ignore corresponding segment match.

###### One of the arguments from this list "dst_any, dst_segments, intra_segment" can be set

`dst_any` - (Optional) Traffic is not matched against any segment (`Bool`).

`dst_segments` - (Optional) Traffic is matched against destination segment in selected segments. See [Dst Segment Choice Dst Segments ](#dst-segment-choice-dst-segments) below for details.

`intra_segment` - (Optional) Traffic is matched for source and destination on the same segment (`Bool`).

###### One of the arguments from this list "src_any, src_segments" can be set

`src_any` - (Optional) Traffic is not matched against any segment (`Bool`).

`src_segments` - (Optional) Source traffic is matched against selected segments. See [Src Segment Choice Src Segments ](#src-segment-choice-src-segments) below for details.

### Allow List Dest List

L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.

`ipv6_prefixes` - (Optional) Destination IPv6 prefixes. (`String`).

`port_ranges` - (Required) Each port range consists of a single port or two ports separated by "-". (`String`).

`prefixes` - (Optional) Destination IPv4 prefixes. (`String`).

### Allow List Http List

URLs for HTTP connections.

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain names e.g "xyz.com" will match "*.xyz.com" (`String`).

###### One of the arguments from this list "any_path, path_exact_value, path_prefix_value, path_regex_value" must be set

`any_path` - (Optional) All paths are considered match (`Bool`).

`path_exact_value` - (Optional) Exact Path to match. (`String`).

`path_prefix_value` - (Optional) Prefix of Path e.g "/abc/xyz" will match "/abc/xyz/.*" (`String`).

`path_regex_value` - (Optional) Regular Expression value for the Path to match (`String`).

### Allow List Tls List

Domains in SNI for TLS connections.

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Default Action Choice Default Action Allow

Allow all connections.

### Default Action Choice Default Action Deny

Deny all connections.

### Default Action Choice Default Action Next Policy

Evaluate the next forward proxy policy in the active list.

### Deny List Dest List

L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.

`ipv6_prefixes` - (Optional) Destination IPv6 prefixes. (`String`).

`port_ranges` - (Required) Each port range consists of a single port or two ports separated by "-". (`String`).

`prefixes` - (Optional) Destination IPv4 prefixes. (`String`).

### Deny List Http List

URLs for HTTP connections.

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain names e.g "xyz.com" will match "*.xyz.com" (`String`).

###### One of the arguments from this list "any_path, path_exact_value, path_prefix_value, path_regex_value" must be set

`any_path` - (Optional) All paths are considered match (`Bool`).

`path_exact_value` - (Optional) Exact Path to match. (`String`).

`path_prefix_value` - (Optional) Prefix of Path e.g "/abc/xyz" will match "/abc/xyz/.*" (`String`).

`path_regex_value` - (Optional) Regular Expression value for the Path to match (`String`).

### Deny List Tls List

Domains in SNI for TLS connections.

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Destination Choice All Destinations

Match on all destinations.

### Destination Choice Dst Asn List

The ASN is obtained by performing a lookup for the destination IPv4 Address in a GeoIP DB..

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer. (`Int`).

### Destination Choice Dst Label Selector

Destination is the set of prefixes determined by the label selector expression.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Destination Choice Dst Prefix List

Addresses that are covered by the given list of IPv4 prefixes.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Destination Choice Http List

URLs for HTTP connections.

`http_list` - (Optional) URLs for HTTP connections. See [Http List Http List ](#http-list-http-list) below for details.

### Destination Choice Tls List

Domains in SNI for TLS connections.

`tls_list` - (Optional) Domains in SNI for TLS connections. See [Tls List Tls List ](#tls-list-tls-list) below for details.

### Destination Choice Url Category List

URL categories to choose, so that the corresponding label selector expressions can be derived from it.

`url_categories` - (Required) List of url categories to be selected (`List of Strings`).

### Dst Segment Choice Dst Any

Traffic is not matched against any segment.

### Dst Segment Choice Dst Segments

Traffic is matched against destination segment in selected segments.

`segments` - (Required) Select list of segments. See [ref](#ref) below for details.

### Dst Segment Choice Intra Segment

Traffic is matched for source and destination on the same segment.

### Http Connect Choice No Http Connect Port

Ignore destination ports for connections.

### Http Connect Choice Port Matcher

In case of an HTTP Connect, the destination port is extracted from the connect destination..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`ports` - (Required) to be part of the range. (`String`).

### Http List Http List

URLs for HTTP connections.

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain names e.g "xyz.com" will match "*.xyz.com" (`String`).

###### One of the arguments from this list "any_path, path_exact_value, path_prefix_value, path_regex_value" must be set

`any_path` - (Optional) All paths are considered match (`Bool`).

`path_exact_value` - (Optional) Exact Path to match. (`String`).

`path_prefix_value` - (Optional) Prefix of Path e.g "/abc/xyz" will match "/abc/xyz/.*" (`String`).

`path_regex_value` - (Optional) Regular Expression value for the Path to match (`String`).

### Path Choice Any Path

All paths are considered match.

### Proxy Choice Proxy Label Selector

Proxy for Network Connector or HTTP connect proxy selected by Label selector.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rule Choice Allow List

List of allowed connections.

###### One of the arguments from this list "default_action_allow, default_action_deny, default_action_next_policy" must be set

`default_action_allow` - (Optional) Allow all connections (`Bool`).

`default_action_deny` - (Optional) Deny all connections (`Bool`).

`default_action_next_policy` - (Optional) Evaluate the next forward proxy policy in the active list (`Bool`).

`dest_list` - (Optional) L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI. See [Allow List Dest List ](#allow-list-dest-list) below for details.

`http_list` - (Optional) URLs for HTTP connections. See [Allow List Http List ](#allow-list-http-list) below for details.

`tls_list` - (Optional) Domains in SNI for TLS connections. See [Allow List Tls List ](#allow-list-tls-list) below for details.

### Rule Choice Deny List

List of denied connections.

###### One of the arguments from this list "default_action_allow, default_action_deny, default_action_next_policy" must be set

`default_action_allow` - (Optional) Allow all connections (`Bool`).

`default_action_deny` - (Optional) Deny all connections (`Bool`).

`default_action_next_policy` - (Optional) Evaluate the next forward proxy policy in the active list (`Bool`).

`dest_list` - (Optional) L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI. See [Deny List Dest List ](#deny-list-dest-list) below for details.

`http_list` - (Optional) URLs for HTTP connections. See [Deny List Http List ](#deny-list-http-list) below for details.

`tls_list` - (Optional) Domains in SNI for TLS connections. See [Deny List Tls List ](#deny-list-tls-list) below for details.

### Rule Choice Rule List

List of custom rules.

`rules` - (Required) List of custom rules. See [Rule List Rules ](#rule-list-rules) below for details.

### Rule List Rules

List of custom rules.

`action` - (Required) Action to be enforced if the input request matches the rule. (`String`).

###### One of the arguments from this list "all_destinations, dst_asn_list, dst_asn_set, dst_ip_prefix_set, dst_label_selector, dst_prefix_list, http_list, tls_list, url_category_list" must be set

`all_destinations` - (Optional) Match on all destinations (`Bool`).

`dst_asn_list` - (Optional) The ASN is obtained by performing a lookup for the destination IPv4 Address in a GeoIP DB.. See [Destination Choice Dst Asn List ](#destination-choice-dst-asn-list) below for details.

`dst_asn_set` - (Optional) The ASN is obtained by performing a lookup for the destination IPv4 Address in a GeoIP DB.. See [ref](#ref) below for details.

`dst_ip_prefix_set` - (Optional) Addresses that are covered by the prefixes in the given ip_prefix_set. See [ref](#ref) below for details.

`dst_label_selector` - (Optional) Destination is the set of prefixes determined by the label selector expression. See [Destination Choice Dst Label Selector ](#destination-choice-dst-label-selector) below for details.

`dst_prefix_list` - (Optional) Addresses that are covered by the given list of IPv4 prefixes. See [Destination Choice Dst Prefix List ](#destination-choice-dst-prefix-list) below for details.

`http_list` - (Optional) URLs for HTTP connections. See [Destination Choice Http List ](#destination-choice-http-list) below for details.

`tls_list` - (Optional) Domains in SNI for TLS connections. See [Destination Choice Tls List ](#destination-choice-tls-list) below for details.

`url_category_list` - (Optional) URL categories to choose, so that the corresponding label selector expressions can be derived from it. See [Destination Choice Url Category List ](#destination-choice-url-category-list) below for details.

###### One of the arguments from this list "no_http_connect_port, port_matcher" can be set

`no_http_connect_port` - (Optional) Ignore destination ports for connections (`Bool`).

`port_matcher` - (Optional) In case of an HTTP Connect, the destination port is extracted from the connect destination.. See [Http Connect Choice Port Matcher ](#http-connect-choice-port-matcher) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Rules Metadata ](#rules-metadata) below for details.

###### One of the arguments from this list "all_sources, ip_prefix_set, label_selector, prefix_list" must be set

`all_sources` - (Optional) Any source that matches 0/0 ip prefix (`Bool`).

`ip_prefix_set` - (Optional) All ip prefixes that are in a given ip prefix set.. See [ref](#ref) below for details.

`label_selector` - (Optional) Sources is set of prefixes determined by label selector expression. See [Source Choice Label Selector ](#source-choice-label-selector) below for details.

`prefix_list` - (Optional) list is a sublist of both V4 and V6 prefix list. See [Source Choice Prefix List ](#source-choice-prefix-list) below for details.

### Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Source Choice All Sources

Any source that matches 0/0 ip prefix.

### Source Choice Label Selector

Sources is set of prefixes determined by label selector expression.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Source Choice Prefix List

list is a sublist of both V4 and V6 prefix list.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Src Segment Choice Src Any

Traffic is not matched against any segment.

### Src Segment Choice Src Segments

Source traffic is matched against selected segments.

`segments` - (Required) Select list of segments. See [ref](#ref) below for details.

### Tls List Tls List

Domains in SNI for TLS connections.

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured forward_proxy_policy.
