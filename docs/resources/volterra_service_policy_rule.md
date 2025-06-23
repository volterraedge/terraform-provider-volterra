---

page_title: "Volterra: service_policy_rule"

description: "The service_policy_rule allows CRUD of Service Policy Rule resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------

Resource volterra_service_policy_rule
=====================================

The Service Policy Rule allows CRUD of Service Policy Rule resource on Volterra SaaS

~> **Note:** Please refer to [Service Policy Rule API docs](https://docs.cloud.f5.com/docs-v2/api/service-policy-rule) to learn more

Example Usage
-------------

```hcl
resource "volterra_service_policy_rule" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  action = ["action"]

  // One of the arguments from this list "any_asn asn_list asn_matcher" must be set

  any_asn = true

  // One of the arguments from this list "any_client client_name client_name_matcher client_selector ip_threat_category_list" must be set

  any_client = true

  // One of the arguments from this list "any_ip ip_matcher ip_prefix_list" must be set

  any_ip = true
  waf_action {
    // One of the arguments from this list "app_firewall_detection_control none waf_skip_processing" must be set

    none = true
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

`action` - (Required) Action to be enforced if the input request matches the rule. (`String`).

`api_group_matcher` - (Optional) The predicate evaluates to true if any of the actual API group names for the request is equal to any of the values in the api group matcher.. See [Api Group Matcher ](#api-group-matcher) below for details.

`arg_matchers` - (Optional) Note that all specified arg matcher predicates must evaluate to true.. See [Arg Matchers ](#arg-matchers) below for details.

###### One of the arguments from this list "any_asn, asn_list, asn_matcher" must be set

`any_asn` - (Optional) Any origin ASN. (`Bool`).

`asn_list` - (Optional) The predicate evaluates to true if the origin ASN is present in the ASN list.. See [Asn Choice Asn List ](#asn-choice-asn-list) below for details.

`asn_matcher` - (Optional) The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects.. See [Asn Choice Asn Matcher ](#asn-choice-asn-matcher) below for details.

`body_matcher` - (Optional) The actual request body value is extracted from the request API as a string.. See [Body Matcher ](#body-matcher) below for details.

`bot_action` - (Optional) Bot action to be enforced if the input request matches the rule.. See [Bot Action ](#bot-action) below for details.

###### One of the arguments from this list "any_client, client_name, client_name_matcher, client_selector, ip_threat_category_list" must be set

`any_client` - (Optional) Any Client (`Bool`).

`client_name` - (Optional) The predicate evaluates to true if any of the actual names is the same as the expected client name. (`String`).

`client_name_matcher` - (Optional) The predicate evaluates to true if any of the client's actual names match any of the exact values or regular expressions in the client name matcher.. See [Client Choice Client Name Matcher ](#client-choice-client-name-matcher) below for details.

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.

`ip_threat_category_list` - (Optional) IP threat categories to choose from. See [Client Choice Ip Threat Category List ](#client-choice-ip-threat-category-list) below for details.

`cookie_matchers` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Cookie Matchers ](#cookie-matchers) below for details.

`domain_matcher` - (Optional) matcher.. See [Domain Matcher ](#domain-matcher) below for details.

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`headers` - (Optional) Note that all specified header predicates must evaluate to true.. See [Headers ](#headers) below for details.

`http_method` - (Optional) The predicate evaluates to true if the actual HTTP method belongs is present in the list of expected values.. See [Http Method ](#http-method) below for details.

###### One of the arguments from this list "any_ip, ip_matcher, ip_prefix_list" must be set

`any_ip` - (Optional) Any Source IP (`Bool`).

`ip_matcher` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes in the IP Prefix Sets.. See [Ip Choice Ip Matcher ](#ip-choice-ip-matcher) below for details.

`ip_prefix_list` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes from the list.. See [Ip Choice Ip Prefix List ](#ip-choice-ip-prefix-list) below for details.

`jwt_claims` - (Optional) Note that all specified JWT claim predicates must evaluate to true.. See [Jwt Claims ](#jwt-claims) below for details.

`label_matcher` - (Optional) other labels do not matter.. See [Label Matcher ](#label-matcher) below for details.

`mum_action` - (Optional) Specifies how Malicious User Mitigation is handled. See [Mum Action ](#mum-action) below for details.

`path` - (Optional) The predicate evaluates to true if the actual path value matches any of the exact or prefix values or regular expressions in the path matcher.. See [Path ](#path) below for details.

`port_matcher` - (Optional) The list of port ranges to which the destination port should belong. In case of an HTTP Connect, the port is extracted from the desired destination.. See [Port Matcher ](#port-matcher) below for details.

`query_params` - (Optional) Note that all specified query parameter predicates must evaluate to true.. See [Query Params ](#query-params) below for details.

`request_constraints` - (Optional) Place limits on request based on the request attributes. The request matches if any of the attribute sizes exceed the corresponding maximum value.. See [Request Constraints ](#request-constraints) below for details.

`segment_policy` - (Optional) Skip the configuration or set option as Any to ignore corresponding segment match. See [Segment Policy ](#segment-policy) below for details.

###### One of the arguments from this list "ja4_tls_fingerprint, tls_fingerprint_matcher" can be set

`ja4_tls_fingerprint` - (Optional) SSL/TLS clients and potentially has a different structure and length.. See [Tls Fingerprint Choice Ja4 Tls Fingerprint ](#tls-fingerprint-choice-ja4-tls-fingerprint) below for details.

`tls_fingerprint_matcher` - (Optional) parameters of the Client Hello packet during the handshake.. See [Tls Fingerprint Choice Tls Fingerprint Matcher ](#tls-fingerprint-choice-tls-fingerprint-matcher) below for details.

`waf_action` - (Required) App Firewall action to be enforced if the input request matches the rule.. See [Waf Action ](#waf-action) below for details.

### Api Group Matcher

The predicate evaluates to true if any of the actual API group names for the request is equal to any of the values in the api group matcher..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`match` - (Required) A list of exact values to match the input against. (`String`).

### Arg Matchers

Note that all specified arg matcher predicates must evaluate to true..

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the argument is not present. (`Bool`).

`check_present` - (Optional) Check that the argument is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the Arg. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-sensitive JSON path in the HTTP request body. (`String`).

### Body Matcher

The actual request body value is extracted from the request API as a string..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Bot Action

Bot action to be enforced if the input request matches the rule..

###### One of the arguments from this list "bot_skip_processing, none" must be set

`bot_skip_processing` - (Optional) Skip all Bot processing for this request (`Bool`).

`none` - (Optional) Perform normal Bot processing for this request (`Bool`).

### Cookie Matchers

Note that all specified cookie matcher predicates must evaluate to true..

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the cookie is not present. (`Bool`).

`check_present` - (Optional) Check that the cookie is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the cookie. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-sensitive cookie name. (`String`).

### Domain Matcher

matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Headers

Note that all specified header predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Http Method

The predicate evaluates to true if the actual HTTP method belongs is present in the list of expected values..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`methods` - (Optional) x-example: "['GET', 'POST', 'DELETE']" (`List of Strings`).

### Jwt Claims

Note that all specified JWT claim predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the JWT Claim is not present. (`Bool`).

`check_present` - (Optional) Check that the JWT Claim is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the JWT Claim. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) JWT claim name. (`String`).

### Label Matcher

other labels do not matter..

`keys` - (Optional) The list of label key names that have to match (`String`).

### Mum Action

Specifies how Malicious User Mitigation is handled.

###### One of the arguments from this list "default, skip_processing" must be set

`default` - (Optional) Perform the default enforcement for this request (`Bool`).

`skip_processing` - (Optional) Do not perform enforcement for this request (`Bool`).

### Path

The predicate evaluates to true if the actual path value matches any of the exact or prefix values or regular expressions in the path matcher..

`exact_values` - (Optional) A list of exact path values to match the input HTTP path against. (`String`).

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_values` - (Optional) A list of path prefix values to match the input HTTP path against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input HTTP path against. (`String`).

`suffix_values` - (Optional) A list of path suffix values to match the input HTTP path against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Port Matcher

The list of port ranges to which the destination port should belong. In case of an HTTP Connect, the port is extracted from the desired destination..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`ports` - (Required) to be part of the range. (`String`).

### Query Params

Note that all specified query parameter predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`key` - (Required) A case-sensitive HTTP query parameter name. (`String`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the query parameter is not present. (`Bool`).

`check_present` - (Optional) Check that the query parameter is present. (`Bool`).

`item` - (Optional) criteria in the matcher.. See [Match Item ](#match-item) below for details.

### Request Constraints

Place limits on request based on the request attributes. The request matches if any of the attribute sizes exceed the corresponding maximum value..

###### One of the arguments from this list "max_cookie_count_exceeds, max_cookie_count_none" must be set

`max_cookie_count_exceeds` - (Optional) x-example: "40" (`Int`).

`max_cookie_count_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

###### One of the arguments from this list "max_cookie_key_size_exceeds, max_cookie_key_size_none" must be set

`max_cookie_key_size_exceeds` - (Optional) x-example: "64" (`Int`).

`max_cookie_key_size_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

###### One of the arguments from this list "max_cookie_value_size_exceeds, max_cookie_value_size_none" must be set

`max_cookie_value_size_exceeds` - (Optional) x-example: "4096" (`Int`).

`max_cookie_value_size_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

###### One of the arguments from this list "max_header_count_exceeds, max_header_count_none" must be set

`max_header_count_exceeds` - (Optional) x-example: "20" (`Int`).

`max_header_count_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

###### One of the arguments from this list "max_header_key_size_exceeds, max_header_key_size_none" must be set

`max_header_key_size_exceeds` - (Optional) x-example: "32" (`Int`).

`max_header_key_size_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

###### One of the arguments from this list "max_header_value_size_exceeds, max_header_value_size_none" must be set

`max_header_value_size_exceeds` - (Optional) x-example: "1024" (`Int`).

`max_header_value_size_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

###### One of the arguments from this list "max_parameter_count_exceeds, max_parameter_count_none" must be set

`max_parameter_count_exceeds` - (Optional) x-example: "4" (`Int`).

`max_parameter_count_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

###### One of the arguments from this list "max_parameter_name_size_exceeds, max_parameter_name_size_none" must be set

`max_parameter_name_size_exceeds` - (Optional) x-example: "64" (`Int`).

`max_parameter_name_size_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

###### One of the arguments from this list "max_parameter_value_size_exceeds, max_parameter_value_size_none" must be set

`max_parameter_value_size_exceeds` - (Optional) x-example: "1000" (`Int`).

`max_parameter_value_size_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

###### One of the arguments from this list "max_query_size_exceeds, max_query_size_none" must be set

`max_query_size_exceeds` - (Optional) x-example: "4096" (`Int`).

`max_query_size_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

###### One of the arguments from this list "max_request_line_size_exceeds, max_request_line_size_none" must be set

`max_request_line_size_exceeds` - (Optional) x-example: "4096" (`Int`).

`max_request_line_size_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

###### One of the arguments from this list "max_request_size_exceeds, max_request_size_none" must be set

`max_request_size_exceeds` - (Optional) x-example: "32768" (`Int`).

`max_request_size_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

###### One of the arguments from this list "max_url_size_exceeds, max_url_size_none" must be set

`max_url_size_exceeds` - (Optional) x-example: "4096" (`Int`).

`max_url_size_none` - (Optional) x-displayName: "Not Configured" (`Bool`).

### Segment Policy

Skip the configuration or set option as Any to ignore corresponding segment match.

###### One of the arguments from this list "dst_any, dst_segments, intra_segment" can be set

`dst_any` - (Optional) Traffic is not matched against any segment (`Bool`).

`dst_segments` - (Optional) Traffic is matched against destination segment in selected segments. See [Dst Segment Choice Dst Segments ](#dst-segment-choice-dst-segments) below for details.

`intra_segment` - (Optional) Traffic is matched for source and destination on the same segment (`Bool`).

###### One of the arguments from this list "src_any, src_segments" can be set

`src_any` - (Optional) Traffic is not matched against any segment (`Bool`).

`src_segments` - (Optional) Source traffic is matched against selected segments. See [Src Segment Choice Src Segments ](#src-segment-choice-src-segments) below for details.

### Waf Action

App Firewall action to be enforced if the input request matches the rule..

###### One of the arguments from this list "app_firewall_detection_control, none, waf_skip_processing" must be set

`app_firewall_detection_control` - (Optional) Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.. See [Action Type App Firewall Detection Control ](#action-type-app-firewall-detection-control) below for details.

`none` - (Optional) Perform normal App Firewall processing for this request (`Bool`).

`waf_skip_processing` - (Optional) Skip all App Firewall processing for this request (`Bool`).

### Action Type App Firewall Detection Control

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria..

`exclude_attack_type_contexts` - (Optional) Attack Types to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Attack Type Contexts ](#app-firewall-detection-control-exclude-attack-type-contexts) below for details.

`exclude_bot_name_contexts` - (Optional) Bot Names to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Bot Name Contexts ](#app-firewall-detection-control-exclude-bot-name-contexts) below for details.

`exclude_signature_contexts` - (Optional) Signature IDs to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Signature Contexts ](#app-firewall-detection-control-exclude-signature-contexts) below for details.

`exclude_violation_contexts` - (Optional) Violations to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Violation Contexts ](#app-firewall-detection-control-exclude-violation-contexts) below for details.

### Action Type Bot Skip Processing

Skip all Bot processing for this request.

### Action Type Default

Perform the default enforcement for this request.

### Action Type None

Perform normal Bot processing for this request.

### Action Type Skip Processing

Do not perform enforcement for this request.

### Action Type Waf Skip Processing

Skip all App Firewall processing for this request.

### App Firewall Detection Control Exclude Attack Type Contexts

Attack Types to be excluded for the defined match criteria.

`context` - (Required) x-required (`String`).

`context_name` - (Optional) with an wildcard asterisk (*). (`String`).

`exclude_attack_type` - (Required) x-required (`String`).

### App Firewall Detection Control Exclude Bot Name Contexts

Bot Names to be excluded for the defined match criteria.

`bot_name` - (Required) x-example: "Hydra" (`String`).

### App Firewall Detection Control Exclude Signature Contexts

Signature IDs to be excluded for the defined match criteria.

`context` - (Required) x-required (`String`).

`context_name` - (Optional) with an wildcard asterisk (*). (`String`).

`signature_id` - (Required) 0 implies that all signatures will be excluded for the specified context. (`Int`).

### App Firewall Detection Control Exclude Violation Contexts

Violations to be excluded for the defined match criteria.

`context` - (Required) x-required (`String`).

`context_name` - (Optional) with an wildcard asterisk (*). (`String`).

`exclude_violation` - (Required) x-required (`String`).

### Asn Choice Asn List

The predicate evaluates to true if the origin ASN is present in the ASN list..

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer. (`Int`).

### Asn Choice Asn Matcher

The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects..

`asn_sets` - (Required) A list of references to bgp_asn_set objects.. See [ref](#ref) below for details.

### Client Choice Client Name Matcher

The predicate evaluates to true if any of the client's actual names match any of the exact values or regular expressions in the client name matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Client Choice Client Selector

The predicate evaluates to true if the expressions in the label selector are true for the client labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Client Choice Ip Threat Category List

IP threat categories to choose from.

`ip_threat_categories` - (Required) The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions (`List of Strings`).

### Dst Segment Choice Dst Any

Traffic is not matched against any segment.

### Dst Segment Choice Dst Segments

Traffic is matched against destination segment in selected segments.

`segments` - (Required) Select list of segments. See [ref](#ref) below for details.

### Dst Segment Choice Intra Segment

Traffic is matched for source and destination on the same segment.

### Ip Choice Ip Matcher

The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes in the IP Prefix Sets..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_sets` - (Required) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Ip Choice Ip Prefix List

The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes from the list..

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

### Match Check Not Present

Check that the argument is not present..

### Match Check Present

Check that the argument is present..

### Match Item

Criteria for matching the values for the Arg. The match is successful if any of the values in the input satisfies the criteria in the matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Max Cookie Count Choice Max Cookie Count None

x-displayName: "Not Configured".

### Max Cookie Key Size Choice Max Cookie Key Size None

x-displayName: "Not Configured".

### Max Cookie Value Size Choice Max Cookie Value Size None

x-displayName: "Not Configured".

### Max Header Count Choice Max Header Count None

x-displayName: "Not Configured".

### Max Header Key Size Choice Max Header Key Size None

x-displayName: "Not Configured".

### Max Header Value Size Choice Max Header Value Size None

x-displayName: "Not Configured".

### Max Parameter Count Choice Max Parameter Count None

x-displayName: "Not Configured".

### Max Parameter Name Size Choice Max Parameter Name Size None

x-displayName: "Not Configured".

### Max Parameter Value Size Choice Max Parameter Value Size None

x-displayName: "Not Configured".

### Max Query Size Choice Max Query Size None

x-displayName: "Not Configured".

### Max Request Line Size Choice Max Request Line Size None

x-displayName: "Not Configured".

### Max Request Size Choice Max Request Size None

x-displayName: "Not Configured".

### Max Url Size Choice Max Url Size None

x-displayName: "Not Configured".

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Src Segment Choice Src Any

Traffic is not matched against any segment.

### Src Segment Choice Src Segments

Source traffic is matched against selected segments.

`segments` - (Required) Select list of segments. See [ref](#ref) below for details.

### Tls Fingerprint Choice Ja4 Tls Fingerprint

SSL/TLS clients and potentially has a different structure and length..

`exact_values` - (Optional) A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against (`String`).

### Tls Fingerprint Choice Tls Fingerprint Matcher

parameters of the Client Hello packet during the handshake..

`classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`exact_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

`excluded_values` - (Optional) or more known TLS fingerprint classes in the enclosing matcher. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured service_policy_rule.
