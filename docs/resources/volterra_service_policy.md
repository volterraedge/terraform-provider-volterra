---

page_title: "Volterra: service_policy"

description: "The service_policy allows CRUD of Service Policy resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_service_policy
================================

The Service Policy allows CRUD of Service Policy resource on Volterra SaaS

~> **Note:** Please refer to [Service Policy API docs](https://volterra.io/docs/api/service-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_service_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
  algo      = ["algo"]

  // One of the arguments from this list "legacy_rule_list allow_all_requests deny_all_requests internally_generated allow_list deny_list rule_list" must be set
  allow_all_requests = true

  // One of the arguments from this list "server_name server_selector server_name_matcher any_server" must be set
  server_name = "database.production.customer.volterra.us"
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

`algo` - (Required) - DENY_OVERRIDES Rules with a DENY action are evaluated prior to rules with an ALLOW action (`String`).

`port_matcher` - (Optional) The list of port ranges to which the destination port should belong. In case of an HTTP Connect, the port is extracted from the desired destination.. See [Port Matcher ](#port-matcher) below for details.

`allow_all_requests` - (Optional) Allow all requests (bool).

`allow_list` - (Optional) List of allowed sources for requests. See [Allow List ](#allow-list) below for details.

`deny_all_requests` - (Optional) Deny all requests (bool).

`deny_list` - (Optional) List of denied sources for requests. See [Deny List ](#deny-list) below for details.

`internally_generated` - (Optional) Placeholder that's used for internally generated service_policy objects to satisy the validation check that rule_choice is non-nil. (bool).

`legacy_rule_list` - (Optional) List of references to service_policy_rule objects. See [Legacy Rule List ](#legacy-rule-list) below for details.

`rule_list` - (Optional) List of custom rules. See [Rule List ](#rule-list) below for details.

`any_server` - (Optional) Any Server (bool).

`server_name` - (Optional) The predicate evaluates to true if any of the actual names is the same as the expected server name. (`String`).

`server_name_matcher` - (Optional) The predicate evaluates to true if any of the server's actual names match any of the exact values or regular expressions in the server name matcher.. See [Server Name Matcher ](#server-name-matcher) below for details.

`server_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the server labels.. See [Server Selector ](#server-selector) below for details.

### Allow List

List of allowed sources for requests.

`asn_list` - (Optional) The ASN is obtained by performing a lookup for the source IPv4 Address in a GeoIP DB.. See [Asn List ](#asn-list) below for details.

`asn_set` - (Optional) The ASN is obtained by performing a lookup for the source IPv4 Address in a GeoIP DB.. See [ref](#ref) below for details.

`country_list` - (Optional) The country is obtained by performing a lookup for the source IPv4 Address in a GeoIP DB. (`List of Strings`).

`default_action_allow` - (Optional) Allow all requests (bool).

`default_action_deny` - (Optional) Deny all requests (bool).

`default_action_next_policy` - (Optional) Evaluate the next service policy (bool).

`ip_prefix_set` - (Optional) Addresses that are covered by the prefixes in the given ip_prefix_set. See [ref](#ref) below for details.

`prefix_list` - (Optional) Addresses that are covered by the given list of IPv4 prefixes. See [Prefix List ](#prefix-list) below for details.

`tls_fingerprint_classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`tls_fingerprint_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

### Any Asn

Any origin ASN..

### Any Client

Any Client.

### Any Dst Asn

Any origin ASN..

### Any Dst Ip

Any Destination IP.

### Any Ip

Any Source IP.

### Api Group Matcher

The predicate evaluates to true if any of the actual API group names for the request is equal to any of the values in the api group matcher..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`match` - (Required) A list of exact values to match the input against. (`String`).

### Arg Matchers

Note that all specified arg matcher predicates must evaluate to true..

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

`check_not_present` - (Optional) Check that the argument is not present. (bool).

`check_present` - (Optional) Check that the argument is present. (bool).

`item` - (Optional) Criteria for matching the values for the Arg. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the arg is present or absent. (`Bool`).

`name` - (Required) A case-sensitive JSON path in the HTTP request body. (`String`).

### Asn List

The ASN is obtained by performing a lookup for the source IPv4 Address in a GeoIP DB..

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. (`Int`).

### Asn Matcher

The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects..

`asn_sets` - (Required) A list of references to bgp_asn_set objects.. See [ref](#ref) below for details.

### Body Matcher

The actual request body value is extracted from the request API as a string..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Check Not Present

Check that the argument is not present..

### Check Present

Check that the argument is present..

### Client Name Matcher

The predicate evaluates to true if any of the client's actual names match any of the exact values or regular expressions in the client name matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Client Role

The predicate evaluates to true if any of the client's roles match the value(s) specified in client role..

`match` - (Required) Value of the expected role. (`String`).

### Client Selector

The predicate evaluates to true if the expressions in the label selector are true for the client labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Cookie Matchers

Note that all specified cookie matcher predicates must evaluate to true..

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

`check_not_present` - (Optional) Check that the cookie is not present. (bool).

`check_present` - (Optional) Check that the cookie is present. (bool).

`item` - (Optional) Criteria for matching the values for the cookie. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the cookie is present or absent. (`Bool`).

`name` - (Required) A case-sensitive cookie name. (`String`).

### Default Action Allow

Allow all requests.

### Default Action Deny

Deny all requests.

### Default Action Next Policy

Evaluate the next service policy.

### Deny List

List of denied sources for requests.

`asn_list` - (Optional) The ASN is obtained by performing a lookup for the source IPv4 Address in a GeoIP DB.. See [Asn List ](#asn-list) below for details.

`asn_set` - (Optional) The ASN is obtained by performing a lookup for the source IPv4 Address in a GeoIP DB.. See [ref](#ref) below for details.

`country_list` - (Optional) The country is obtained by performing a lookup for the source IPv4 Address in a GeoIP DB. (`List of Strings`).

`default_action_allow` - (Optional) Allow all requests (bool).

`default_action_deny` - (Optional) Deny all requests (bool).

`default_action_next_policy` - (Optional) Evaluate the next service policy (bool).

`ip_prefix_set` - (Optional) Addresses that are covered by the prefixes in the given ip_prefix_set. See [ref](#ref) below for details.

`prefix_list` - (Optional) Addresses that are covered by the given list of IPv4 prefixes. See [Prefix List ](#prefix-list) below for details.

`tls_fingerprint_classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`tls_fingerprint_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

### Domain Matcher

matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Dst Asn List

The predicate evaluates to true if the destination ASN is present in the ASN list..

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. (`Int`).

### Dst Asn Matcher

The predicate evaluates to true if the destination ASN is present in one of the BGP ASN Set objects..

`asn_sets` - (Required) A list of references to bgp_asn_set objects.. See [ref](#ref) below for details.

### Dst Ip Matcher

The predicate evaluates to true if the client IPv4 Address is covered by one or more of the IPv4 Prefixes in the IP Prefix Sets..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_sets` - (Required) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Dst Ip Prefix List

The predicate evaluates to true if the destination address is covered by one or more of the IPv4 Prefixes from the list..

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Required) List of IPv4 prefix strings. (`String`).

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

The predicate evaluates to true if the client IPv4 Address is covered by one or more of the IPv4 Prefixes in the IP Prefix Sets..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_sets` - (Required) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Ip Prefix List

The predicate evaluates to true if the client IPv4 Address is covered by one or more of the IPv4 Prefixes from the list..

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Required) List of IPv4 prefix strings. (`String`).

### Item

Criteria for matching the values for the Arg. The match is successful if any of the values in the input satisfies the criteria in the matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### L4 Dest Matcher

IP matches one of the prefixes and the destination port belongs to the port range..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`l4_dests` - (Required) A list of L4 destinations used as match criteria. The match is considered successful if the destination IP and path match any of the L4 destinations.. See [L4 Dests ](#l4-dests) below for details.

### L4 Dests

A list of L4 destinations used as match criteria. The match is considered successful if the destination IP and path match any of the L4 destinations..

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

`prefixes` - (Required) Destination IPv4 prefixes. (`String`).

### Label Matcher

other labels do not matter..

`keys` - (Optional) The list of label key names that have to match (`String`).

### Legacy Rule List

List of references to service_policy_rule objects.

`rules` - (Optional) The order of evaluation of the rules depends on the rule combining algorithm.. See [ref](#ref) below for details.

### Malicious User Mitigation Bypass

the appropriate match conditions in the enclosing policy rule and setting malicious user mitigation bypass flag..

### Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### None

Perform normal App Firewall processing for this request.

### Path

The predicate evaluates to true if the actual path value matches any of the exact or prefix values or regular expressions in the path matcher..

`exact_values` - (Optional) A list of exact path values to match the input HTTP path against. (`String`).

`prefix_values` - (Optional) A list of path prefix values to match the input HTTP path against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input HTTP path against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Port Matcher

The list of port ranges to which the destination port should belong. In case of an HTTP Connect, the port is extracted from the desired destination..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`ports` - (Required) to be part of the range. (`String`).

### Prefix List

Addresses that are covered by the given list of IPv4 prefixes.

`prefixes` - (Required) List of IPv4 prefixes that represent an endpoint (`String`).

### Query Params

Note that all specified query parameter predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`key` - (Required) A case-sensitive HTTP query parameter name. (`String`).

`check_not_present` - (Optional) Check that the query parameter is not present. (bool).

`check_present` - (Optional) Check that the query parameter is present. (bool).

`item` - (Optional) criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the query parameter is present or absent. (`Bool`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rule List

List of custom rules.

`rules` - (Optional) The order of evaluation of the rules depends on the rule combining algorithm.. See [Rules ](#rules) below for details.

### Rules

The order of evaluation of the rules depends on the rule combining algorithm..

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`spec` - (Required) Specification for the rule including match predicates and actions.. See [Spec ](#spec) below for details.

### Server Name Matcher

The predicate evaluates to true if any of the server's actual names match any of the exact values or regular expressions in the server name matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Server Selector

The predicate evaluates to true if the expressions in the label selector are true for the server labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Spec

Specification for the rule including match predicates and actions..

`action` - (Required) Action to be enforced if the input request matches the rule. (`String`).

`api_group_matcher` - (Optional) The predicate evaluates to true if any of the actual API group names for the request is equal to any of the values in the api group matcher.. See [Api Group Matcher ](#api-group-matcher) below for details.

`arg_matchers` - (Optional) Note that all specified arg matcher predicates must evaluate to true.. See [Arg Matchers ](#arg-matchers) below for details.

`any_asn` - (Optional) Any origin ASN. (bool).

`asn_list` - (Optional) The predicate evaluates to true if the origin ASN is present in the ASN list.. See [Asn List ](#asn-list) below for details.

`asn_matcher` - (Optional) The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects.. See [Asn Matcher ](#asn-matcher) below for details.

`body_matcher` - (Optional) The actual request body value is extracted from the request API as a string.. See [Body Matcher ](#body-matcher) below for details.

`challenge_action` - (Required) Select challenge action, enable javascript/captcha challenge or disable challenge (`String`).

`any_client` - (Optional) Any Client (bool).

`client_name` - (Optional) The predicate evaluates to true if any of the actual names is the same as the expected client name. (`String`).

`client_name_matcher` - (Optional) The predicate evaluates to true if any of the client's actual names match any of the exact values or regular expressions in the client name matcher.. See [Client Name Matcher ](#client-name-matcher) below for details.

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Selector ](#client-selector) below for details.

`client_role` - (Optional) The predicate evaluates to true if any of the client's roles match the value(s) specified in client role.. See [Client Role ](#client-role) below for details.

`cookie_matchers` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Cookie Matchers ](#cookie-matchers) below for details.

`domain_matcher` - (Optional) matcher.. See [Domain Matcher ](#domain-matcher) below for details.

`any_dst_asn` - (Optional) Any origin ASN. (bool).

`dst_asn_list` - (Optional) The predicate evaluates to true if the destination ASN is present in the ASN list.. See [Dst Asn List ](#dst-asn-list) below for details.

`dst_asn_matcher` - (Optional) The predicate evaluates to true if the destination ASN is present in one of the BGP ASN Set objects.. See [Dst Asn Matcher ](#dst-asn-matcher) below for details.

`any_dst_ip` - (Optional) Any Destination IP (bool).

`dst_ip_matcher` - (Optional) The predicate evaluates to true if the client IPv4 Address is covered by one or more of the IPv4 Prefixes in the IP Prefix Sets.. See [Dst Ip Matcher ](#dst-ip-matcher) below for details.

`dst_ip_prefix_list` - (Optional) The predicate evaluates to true if the destination address is covered by one or more of the IPv4 Prefixes from the list.. See [Dst Ip Prefix List ](#dst-ip-prefix-list) below for details.

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`forwarding_class` - (Optional) Action valid only when the policy is used PBR. See [ref](#ref) below for details.

`goto_policy` - (Optional) The target policy must be part of the current policy set and must be after the current policy in the policy set.. See [ref](#ref) below for details.

`headers` - (Optional) Note that all specified header predicates must evaluate to true.. See [Headers ](#headers) below for details.

`http_method` - (Optional) The predicate evaluates to true if the actual HTTP method belongs is present in the list of expected values.. See [Http Method ](#http-method) below for details.

`any_ip` - (Optional) Any Source IP (bool).

`ip_matcher` - (Optional) The predicate evaluates to true if the client IPv4 Address is covered by one or more of the IPv4 Prefixes in the IP Prefix Sets.. See [Ip Matcher ](#ip-matcher) below for details.

`ip_prefix_list` - (Optional) The predicate evaluates to true if the client IPv4 Address is covered by one or more of the IPv4 Prefixes from the list.. See [Ip Prefix List ](#ip-prefix-list) below for details.

`l4_dest_matcher` - (Optional) IP matches one of the prefixes and the destination port belongs to the port range.. See [L4 Dest Matcher ](#l4-dest-matcher) below for details.

`label_matcher` - (Optional) other labels do not matter.. See [Label Matcher ](#label-matcher) below for details.

`malicious_user_mitigation_bypass` - (Optional) the appropriate match conditions in the enclosing policy rule and setting malicious user mitigation bypass flag. (bool).

`path` - (Optional) The predicate evaluates to true if the actual path value matches any of the exact or prefix values or regular expressions in the path matcher.. See [Path ](#path) below for details.

`port_matcher` - (Optional) The list of port ranges to which the destination port should belong. In case of an HTTP Connect, the port is extracted from the desired destination.. See [Port Matcher ](#port-matcher) below for details.

`query_params` - (Optional) Note that all specified query parameter predicates must evaluate to true.. See [Query Params ](#query-params) below for details.

`rate_limiter` - (Optional) Requests matching this the enclosing rule are subjected to the specified rate_limiter.. See [ref](#ref) below for details.

`scheme` - (Optional) The scheme in the request. (`String`).

`server_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the server labels.. See [Server Selector ](#server-selector) below for details.

`tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints.. See [Tls Fingerprint Matcher ](#tls-fingerprint-matcher) below for details.

`url_matcher` - (Optional) A URL matcher specifies a list of URL items as match criteria. The match is considered successful if the domain and path match any of the URL items.. See [Url Matcher ](#url-matcher) below for details.

`virtual_host_matcher` - (Optional) Hidden because this will be used only in system generated rate limiting service_policy_sets.. See [Virtual Host Matcher ](#virtual-host-matcher) below for details.

`waf_action` - (Required) App Firewall action to be enforced if the input request matches the rule.. See [Waf Action ](#waf-action) below for details.

### Tls Fingerprint Matcher

The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints..

`classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`exact_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

`excluded_values` - (Optional) or more known TLS fingerprint classes in the enclosing matcher. (`String`).

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

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Waf Action

App Firewall action to be enforced if the input request matches the rule..

`none` - (Optional) Perform normal App Firewall processing for this request (bool).

`waf_in_monitoring_mode` - (Optional) App Firewall will run in monitoring mode without blocking the request (bool).

`waf_inline_rule_control` - (Optional) App Firewall rule changes to be applied for this request. See [Waf Inline Rule Control ](#waf-inline-rule-control) below for details.

`waf_rule_control` - (Optional) App Firewall rule changes to be applied for this request. See [Waf Rule Control ](#waf-rule-control) below for details.

`waf_skip_processing` - (Optional) Skip all App Firewall processing for this request (bool).

### Waf In Monitoring Mode

App Firewall will run in monitoring mode without blocking the request.

### Waf Inline Rule Control

App Firewall rule changes to be applied for this request.

`exclude_rule_ids` - (Optional) App Firewall Rule IDs to be excluded for this request (`List of Strings`).

`monitoring_mode` - (Optional) App Firewall will run in monitoring mode without blocking the request (`Bool`).

### Waf Rule Control

App Firewall rule changes to be applied for this request.

`exclude_rule_ids` - (Optional) App Firewall Rule List specifying the rule IDs to be excluded for this request. See [ref](#ref) below for details.

`monitoring_mode` - (Optional) App Firewall will run in monitoring mode without blocking the request (`Bool`).

### Waf Skip Processing

Skip all App Firewall processing for this request.

Attribute Reference
-------------------

-	`id` - This is the id of the configured service_policy.
