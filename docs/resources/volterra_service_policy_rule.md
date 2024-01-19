---

page_title: "Volterra: service_policy_rule"

description: "The service_policy_rule allows CRUD of Service Policy Rule resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------

Resource volterra_service_policy_rule
=====================================

The Service Policy Rule allows CRUD of Service Policy Rule resource on Volterra SaaS

~> **Note:** Please refer to [Service Policy Rule API docs](https://docs.cloud.f5.com/docs/api/service-policy-rule) to learn more

Example Usage
-------------

```hcl
resource "volterra_service_policy_rule" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
  action    = ["action"]

  // One of the arguments from this list "asn_list asn_matcher any_asn" must be set
  any_asn          = true
  challenge_action = ["challenge_action"]

  // One of the arguments from this list "any_client client_name ip_threat_category_list client_selector client_name_matcher" must be set
  any_client = true

  // One of the arguments from this list "any_ip ip_prefix_list ip_matcher" must be set
  any_ip = true

  waf_action {
    // One of the arguments from this list "waf_skip_processing waf_in_monitoring_mode app_firewall_detection_control data_guard_control jwt_validation jwt_claims_validation none" must be set
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

`any_asn` - (Optional) Any origin ASN. (bool).

`asn_list` - (Optional) The predicate evaluates to true if the origin ASN is present in the ASN list.. See [Asn List ](#asn-list) below for details.

`asn_matcher` - (Optional) The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects.. See [Asn Matcher ](#asn-matcher) below for details.

`body_matcher` - (Optional) The actual request body value is extracted from the request API as a string.. See [Body Matcher ](#body-matcher) below for details.

`bot_action` - (Optional) Bot action to be enforced if the input request matches the rule.. See [Bot Action ](#bot-action) below for details.

`challenge_action` - (Required) Select challenge action, enable javascript/captcha challenge or disable challenge (`String`).

`any_client` - (Optional) Any Client (bool).

`client_name` - (Optional) The predicate evaluates to true if any of the actual names is the same as the expected client name. (`String`).

`client_name_matcher` - (Optional) The predicate evaluates to true if any of the client's actual names match any of the exact values or regular expressions in the client name matcher.. See [Client Name Matcher ](#client-name-matcher) below for details.

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Selector ](#client-selector) below for details.

`ip_threat_category_list` - (Optional) IP threat categories to choose from. See [Ip Threat Category List ](#ip-threat-category-list) below for details.

`client_role` - (Optional) The predicate evaluates to true if any of the client's roles match the value(s) specified in client role.. See [Client Role ](#client-role) below for details.

`content_rewrite_action` - (Optional) Rewrite HTML response action to insert HTML content such as Javascript <script> tags into the HTML document. See [Content Rewrite Action ](#content-rewrite-action) below for details.

`cookie_matchers` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Cookie Matchers ](#cookie-matchers) below for details.

`domain_matcher` - (Optional) matcher.. See [Domain Matcher ](#domain-matcher) below for details.

`any_dst_asn` - (Optional) Any origin ASN. (bool).

`dst_asn_list` - (Optional) The predicate evaluates to true if the destination ASN is present in the ASN list.. See [Dst Asn List ](#dst-asn-list) below for details.

`dst_asn_matcher` - (Optional) The predicate evaluates to true if the destination ASN is present in one of the BGP ASN Set objects.. See [Dst Asn Matcher ](#dst-asn-matcher) below for details.

`any_dst_ip` - (Optional) Any Destination IP (bool).

`dst_ip_matcher` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes in the IP Prefix Sets.. See [Dst Ip Matcher ](#dst-ip-matcher) below for details.

`dst_ip_prefix_list` - (Optional) The predicate evaluates to true if the destination address is covered by one or more of the IP Prefixes from the list.. See [Dst Ip Prefix List ](#dst-ip-prefix-list) below for details.

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`goto_policy` - (Optional) The target policy must be part of the current policy set and must be after the current policy in the policy set.. See [ref](#ref) below for details.

`headers` - (Optional) Note that all specified header predicates must evaluate to true.. See [Headers ](#headers) below for details.

`http_method` - (Optional) The predicate evaluates to true if the actual HTTP method belongs is present in the list of expected values.. See [Http Method ](#http-method) below for details.

`any_ip` - (Optional) Any Source IP (bool).

`ip_matcher` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes in the IP Prefix Sets.. See [Ip Matcher ](#ip-matcher) below for details.

`ip_prefix_list` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes from the list.. See [Ip Prefix List ](#ip-prefix-list) below for details.

`ip_reputation_action` - (Optional) Specifies how IP Reputation is handled. See [Ip Reputation Action ](#ip-reputation-action) below for details.

`l4_dest_matcher` - (Optional) IP matches one of the prefixes and the destination port belongs to the port range.. See [L4 Dest Matcher ](#l4-dest-matcher) below for details.

`label_matcher` - (Optional) other labels do not matter.. See [Label Matcher ](#label-matcher) below for details.

`mum_action` - (Optional) Specifies how Malicious User Mitigation is handled. See [Mum Action ](#mum-action) below for details.

`origin_server_subsets_action` - (Optional) Add Labels for this origin server, these labels can be used to form subset. (`String`).

`path` - (Optional) The predicate evaluates to true if the actual path value matches any of the exact or prefix values or regular expressions in the path matcher.. See [Path ](#path) below for details.

`port_matcher` - (Optional) The list of port ranges to which the destination port should belong. In case of an HTTP Connect, the port is extracted from the desired destination.. See [Port Matcher ](#port-matcher) below for details.

`query_params` - (Optional) Note that all specified query parameter predicates must evaluate to true.. See [Query Params ](#query-params) below for details.

`rate_limiter` - (Optional) Requests matching this the enclosing rule are subjected to the specified rate_limiter.. See [ref](#ref) below for details.

`request_constraints` - (Optional) Place limits on request based on the request attributes. The request matches if any of the attribute sizes exceed the corresponding maximum value.. See [Request Constraints ](#request-constraints) below for details.

`scheme` - (Optional) The scheme in the request. (`List of String`).

`server_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the server labels.. See [Server Selector ](#server-selector) below for details.

`shape_protected_endpoint_action` - (Optional) Shape Protected Endpoint Action that include application traffic type and mitigation. See [Shape Protected Endpoint Action ](#shape-protected-endpoint-action) below for details.

`tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints.. See [Tls Fingerprint Matcher ](#tls-fingerprint-matcher) below for details.

`url_matcher` - (Optional) A URL matcher specifies a list of URL items as match criteria. The match is considered successful if the domain and path match any of the URL items.. See [Url Matcher ](#url-matcher) below for details.

`virtual_host_matcher` - (Optional) Hidden because this will be used only in system generated rate limiting service_policy_sets.. See [Virtual Host Matcher ](#virtual-host-matcher) below for details.

`waf_action` - (Required) App Firewall action to be enforced if the input request matches the rule.. See [Waf Action ](#waf-action) below for details.

### Api Group Matcher

The predicate evaluates to true if any of the actual API group names for the request is equal to any of the values in the api group matcher..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`match` - (Required) A list of exact values to match the input against. (`String`).

### App Firewall Detection Control

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria..

`exclude_attack_type_contexts` - (Optional) Attack Types to be excluded for the defined match criteria. See [Exclude Attack Type Contexts ](#exclude-attack-type-contexts) below for details.

`exclude_bot_name_contexts` - (Optional) Bot Names to be excluded for the defined match criteria. See [Exclude Bot Name Contexts ](#exclude-bot-name-contexts) below for details.

`exclude_signature_contexts` - (Optional) Signature IDs to be excluded for the defined match criteria. See [Exclude Signature Contexts ](#exclude-signature-contexts) below for details.

`exclude_violation_contexts` - (Optional) Violations to be excluded for the defined match criteria. See [Exclude Violation Contexts ](#exclude-violation-contexts) below for details.

### Append Headers

Append mitigation headers..

`auto_type_header_name` - (Required) A case-insensitive HTTP header name. (`String`).

`inference_header_name` - (Required) A case-insensitive HTTP header name. (`String`).

### Arg Matchers

Note that all specified arg matcher predicates must evaluate to true..

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

`check_not_present` - (Optional) Check that the argument is not present. (bool).

`check_present` - (Optional) Check that the argument is present. (bool).

`item` - (Optional) Criteria for matching the values for the Arg. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the arg is present or absent. (`Bool`).

`name` - (Required) A case-sensitive JSON path in the HTTP request body. (`String`).

### Asn List

The predicate evaluates to true if the origin ASN is present in the ASN list..

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. (`Int`).

### Asn Matcher

The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects..

`asn_sets` - (Required) A list of references to bgp_asn_set objects.. See [ref](#ref) below for details.

### Block

Block bot request and send response with custom content..

`body` - (Optional) E.g. "<p> Your request was blocked </p>". Base64 encoded string for this html is "LzxwPiBZb3VyIHJlcXVlc3Qgd2FzIGJsb2NrZWQgPC9wPg==" (`String`).

`body_hash` - (Optional) Represents the corresponding MD5 Hash for the body message. (`String`).

`status` - (Optional) HTTP Status code to respond with (`String`).

### Body Matcher

The actual request body value is extracted from the request API as a string..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Bot Action

Bot action to be enforced if the input request matches the rule..

`bot_skip_processing` - (Optional) Skip all Bot processing for this request (bool).

`none` - (Optional) Perform normal Bot processing for this request (bool).

### Bot Skip Processing

Skip all Bot processing for this request.

### Check Not Present

Check that the argument is not present..

### Check Present

Check that the argument is present..

### Client Name Matcher

The predicate evaluates to true if any of the client's actual names match any of the exact values or regular expressions in the client name matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Client Role

The predicate evaluates to true if any of the client's roles match the value(s) specified in client role..

`match` - (Required) Value of the expected role. (`String`).

### Client Selector

The predicate evaluates to true if the expressions in the label selector are true for the client labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Content Rewrite Action

Rewrite HTML response action to insert HTML content such as Javascript <script> tags into the HTML document.

`element_selector` - (Required) Element selector to insert into. (`String`).

`insert_content` - (Optional) HTML content to insert. (`String`).

`inserted_types` - (Optional) Inserted types of security configuration like Bot Defense, Client Side Defense. (`Bool`).

`position` - (Optional) Position of HTML content to be inserted within HTML tag. (`String`).

### Cookie Matchers

Note that all specified cookie matcher predicates must evaluate to true..

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

`check_not_present` - (Optional) Check that the cookie is not present. (bool).

`check_present` - (Optional) Check that the cookie is present. (bool).

`item` - (Optional) Criteria for matching the values for the cookie. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the cookie is present or absent. (`Bool`).

`name` - (Required) A case-sensitive cookie name. (`String`).

### Data Guard Control

Data Guard changes to be applied for this request.

`policy_name` - (Optional) Sets the BD Policy to use (`String`).

### Default

Perform the default enforcement for this request.

### Domain Matcher

matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Dst Asn List

The predicate evaluates to true if the destination ASN is present in the ASN list..

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. (`Int`).

### Dst Asn Matcher

The predicate evaluates to true if the destination ASN is present in one of the BGP ASN Set objects..

`asn_sets` - (Required) A list of references to bgp_asn_set objects.. See [ref](#ref) below for details.

### Dst Ip Matcher

The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes in the IP Prefix Sets..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_sets` - (Required) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Dst Ip Prefix List

The predicate evaluates to true if the destination address is covered by one or more of the IP Prefixes from the list..

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

### Exclude Attack Type Contexts

Attack Types to be excluded for the defined match criteria.

`exclude_attack_type` - (Required) x-required (`String`).

### Exclude Bot Name Contexts

Bot Names to be excluded for the defined match criteria.

`bot_name` - (Required) x-example: "Hydra" (`String`).

### Exclude Signature Contexts

Signature IDs to be excluded for the defined match criteria.

`context` - (Required) x-required (`String`).

`context_name` - (Optional) Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. (`String`).

`signature_id` - (Required) x-required (`Int`).

### Exclude Violation Contexts

Violations to be excluded for the defined match criteria.

`context` - (Required) x-required (`String`).

`context_name` - (Optional) Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. (`String`).

`exclude_violation` - (Required) x-required (`String`).

### Failure Conditions

Failure Conditions.

`name` - (Optional) A case-insensitive HTTP header name. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`status` - (Required) HTTP Status code (`String`).

### Flag

Flag the request while not taking any invasive actions..

`append_headers` - (Optional) Append mitigation headers.. See [Append Headers ](#append-headers) below for details.

`no_headers` - (Optional) No mitigation headers. (bool).

### Headers

Note that all specified header predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`check_not_present` - (Optional) Check that the header is not present. (bool).

`check_present` - (Optional) Check that the header is present. (bool).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the header is present or absent. (`Bool`).

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Http Method

The predicate evaluates to true if the actual HTTP method belongs is present in the list of expected values..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`methods` - (Optional) x-example: "['GET', 'POST', 'DELETE']" (`List of Strings`).

### Ip Matcher

The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes in the IP Prefix Sets..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_sets` - (Required) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Ip Prefix List

The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes from the list..

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

### Ip Reputation Action

Specifies how IP Reputation is handled.

`default` - (Optional) Perform the default enforcement for this request (bool).

`skip_processing` - (Optional) Do not perform enforcement for this request (bool).

### Ip Threat Category List

IP threat categories to choose from.

`ip_threat_categories` - (Required) The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions (`List of Strings`).

### Item

Criteria for matching the values for the Arg. The match is successful if any of the values in the input satisfies the criteria in the matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Jwt Claims Validation

Validate JWT Claims for this request.

### Jwt Validation

Validate JWT for this request.

### L4 Dest Matcher

IP matches one of the prefixes and the destination port belongs to the port range..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`l4_dests` - (Required) A list of L4 destinations used as match criteria. The match is considered successful if the destination IP and path match any of the L4 destinations.. See [L4 Dests ](#l4-dests) below for details.

### L4 Dests

A list of L4 destinations used as match criteria. The match is considered successful if the destination IP and path match any of the L4 destinations..

`port_ranges` - (Required) Each port range consists of a single port or two ports separated by "-". (`String`).

`prefixes` - (Required) Destination IPv4 prefixes. (`String`).

### Label Matcher

other labels do not matter..

`keys` - (Optional) The list of label key names that have to match (`String`).

### Max Cookie Count None

x-displayName: "Not Configured".

### Max Cookie Key Size None

x-displayName: "Not Configured".

### Max Cookie Value Size None

x-displayName: "Not Configured".

### Max Header Count None

x-displayName: "Not Configured".

### Max Header Key Size None

x-displayName: "Not Configured".

### Max Header Value Size None

x-displayName: "Not Configured".

### Max Parameter Count None

x-displayName: "Not Configured".

### Max Parameter Name Size None

x-displayName: "Not Configured".

### Max Parameter Value Size None

x-displayName: "Not Configured".

### Max Query Size None

x-displayName: "Not Configured".

### Max Request Line Size None

x-displayName: "Not Configured".

### Max Request Size None

x-displayName: "Not Configured".

### Max Url Size None

x-displayName: "Not Configured".

### Mitigation

Mitigation action for protected endpoint.

`block` - (Optional) Block bot request and send response with custom content.. See [Block ](#block) below for details.

`flag` - (Optional) Flag the request while not taking any invasive actions.. See [Flag ](#flag) below for details.

`none` - (Optional) No mitigation actions. (bool).

`redirect` - (Optional) Redirect bot request to a custom URI.. See [Redirect ](#redirect) below for details.

### Mum Action

Specifies how Malicious User Mitigation is handled.

`default` - (Optional) Perform the default enforcement for this request (bool).

`skip_processing` - (Optional) Do not perform enforcement for this request (bool).

### No Headers

No mitigation headers..

### None

Perform normal Bot processing for this request.

### Path

The predicate evaluates to true if the actual path value matches any of the exact or prefix values or regular expressions in the path matcher..

`exact_values` - (Optional) A list of exact path values to match the input HTTP path against. (`String`).

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

`check_not_present` - (Optional) Check that the query parameter is not present. (bool).

`check_present` - (Optional) Check that the query parameter is present. (bool).

`item` - (Optional) criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the query parameter is present or absent. (`Bool`).

### Redirect

Redirect bot request to a custom URI..

`uri` - (Required) URI location for redirect may be relative or absolute. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Request Constraints

Place limits on request based on the request attributes. The request matches if any of the attribute sizes exceed the corresponding maximum value..

`max_cookie_count_exceeds` - (Optional) x-example: "40" (`Int`).

`max_cookie_count_none` - (Optional) x-displayName: "Not Configured" (bool).

`max_cookie_key_size_exceeds` - (Optional) x-example: "64" (`Int`).

`max_cookie_key_size_none` - (Optional) x-displayName: "Not Configured" (bool).

`max_cookie_value_size_exceeds` - (Optional) x-example: "4096" (`Int`).

`max_cookie_value_size_none` - (Optional) x-displayName: "Not Configured" (bool).

`max_header_count_exceeds` - (Optional) x-example: "20" (`Int`).

`max_header_count_none` - (Optional) x-displayName: "Not Configured" (bool).

`max_header_key_size_exceeds` - (Optional) x-example: "32" (`Int`).

`max_header_key_size_none` - (Optional) x-displayName: "Not Configured" (bool).

`max_header_value_size_exceeds` - (Optional) x-example: "1024" (`Int`).

`max_header_value_size_none` - (Optional) x-displayName: "Not Configured" (bool).

`max_parameter_count_exceeds` - (Optional) x-example: "4" (`Int`).

`max_parameter_count_none` - (Optional) x-displayName: "Not Configured" (bool).

`max_parameter_name_size_exceeds` - (Optional) x-example: "64" (`Int`).

`max_parameter_name_size_none` - (Optional) x-displayName: "Not Configured" (bool).

`max_parameter_value_size_exceeds` - (Optional) x-example: "1000" (`Int`).

`max_parameter_value_size_none` - (Optional) x-displayName: "Not Configured" (bool).

`max_query_size_exceeds` - (Optional) x-example: "4096" (`Int`).

`max_query_size_none` - (Optional) x-displayName: "Not Configured" (bool).

`max_request_line_size_exceeds` - (Optional) x-example: "4096" (`Int`).

`max_request_line_size_none` - (Optional) x-displayName: "Not Configured" (bool).

`max_request_size_exceeds` - (Optional) x-example: "32768" (`Int`).

`max_request_size_none` - (Optional) x-displayName: "Not Configured" (bool).

`max_url_size_exceeds` - (Optional) x-example: "4096" (`Int`).

`max_url_size_none` - (Optional) x-displayName: "Not Configured" (bool).

### Server Selector

The predicate evaluates to true if the expressions in the label selector are true for the server labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Shape Protected Endpoint Action

Shape Protected Endpoint Action that include application traffic type and mitigation.

`allow_goodbot` - (Required) Good bot (`Bool`).

`app_traffic_type` - (Required) Traffic type (`String`).

`flow_label` - (Required) Flow label (`String`).

`mitigation` - (Required) Mitigation action for protected endpoint. See [Mitigation ](#mitigation) below for details.

`transaction_result` - (Optional) Success/failure Criteria for transaction result. See [Transaction Result ](#transaction-result) below for details.

`web_scraping` - (Required) Web scraping protection enabled for protected endpoint (`Bool`).

### Skip Processing

Do not perform enforcement for this request.

### Success Conditions

Success Conditions.

`name` - (Optional) A case-insensitive HTTP header name. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`status` - (Required) HTTP Status code (`String`).

### Tls Fingerprint Matcher

The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints..

`classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`exact_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

`excluded_values` - (Optional) or more known TLS fingerprint classes in the enclosing matcher. (`String`).

### Transaction Result

Success/failure Criteria for transaction result.

`failure_conditions` - (Optional) Failure Conditions. See [Failure Conditions ](#failure-conditions) below for details.

`success_conditions` - (Optional) Success Conditions. See [Success Conditions ](#success-conditions) below for details.

### Url Items

A list of URL items used as match criteria. The match is considered successful if the domain and path match any of the URL items..

`domain_regex` - (Optional) A regular expression to match the domain against. (`String`).

`domain_value` - (Optional) An exact value to match the domain against. (`String`).

`path_prefix` - (Optional) An prefix value to match the path against. (`String`).

`path_regex` - (Optional) A regular expression to match the path against. (`String`).

`path_value` - (Optional) An exact value to match the path against. (`String`).

### Url Matcher

A URL matcher specifies a list of URL items as match criteria. The match is considered successful if the domain and path match any of the URL items..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`url_items` - (Required) A list of URL items used as match criteria. The match is considered successful if the domain and path match any of the URL items.. See [Url Items ](#url-items) below for details.

### Virtual Host Matcher

Hidden because this will be used only in system generated rate limiting service_policy_sets..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Waf Action

App Firewall action to be enforced if the input request matches the rule..

`app_firewall_detection_control` - (Optional) Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.. See [App Firewall Detection Control ](#app-firewall-detection-control) below for details.

`data_guard_control` - (Optional) Data Guard changes to be applied for this request. See [Data Guard Control ](#data-guard-control) below for details.

`jwt_claims_validation` - (Optional) Validate JWT Claims for this request (bool).

`jwt_validation` - (Optional) Validate JWT for this request (bool).

`none` - (Optional) Perform normal App Firewall processing for this request (bool).

`waf_in_monitoring_mode` - (Optional) App Firewall will run in monitoring mode without blocking the request (bool).

`waf_skip_processing` - (Optional) Skip all App Firewall processing for this request (bool).

### Waf In Monitoring Mode

App Firewall will run in monitoring mode without blocking the request.

### Waf Skip Processing

Skip all App Firewall processing for this request.

Attribute Reference
-------------------

-	`id` - This is the id of the configured service_policy_rule.
