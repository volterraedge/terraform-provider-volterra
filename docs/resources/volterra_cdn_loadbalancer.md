---

page_title: "Volterra: cdn_loadbalancer"

description: "The cdn_loadbalancer allows CRUD of Cdn Loadbalancer resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_cdn_loadbalancer
==================================

The Cdn Loadbalancer allows CRUD of Cdn Loadbalancer resource on Volterra SaaS

~> **Note:** Please refer to [Cdn Loadbalancer API docs](https://docs.cloud.f5.com/docs/api/views-cdn-loadbalancer) to learn more

Example Usage
-------------

```hcl
resource "volterra_cdn_loadbalancer" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  domains = ["www.foo.com"]

  // One of the arguments from this list "https_auto_cert https http" must be set

  http {
    dns_volterra_managed = true

    // One of the arguments from this list "port port_ranges" must be set

    port = "80"
  }
  origin_pool {
    follow_origin_redirect = true

    more_origin_options {
      disable_byte_range_request = true

      websocket_proxy = true
    }

    origin_request_timeout = "100s"

    origin_servers {
      // One of the arguments from this list "public_ip public_name" must be set

      public_name {
        dns_name = "value"

        refresh_interval = "20"
      }

      port = "80"
    }

    public_name {
      dns_name = "value"

      refresh_interval = "20"
    }

    // One of the arguments from this list "no_tls use_tls" must be set

    no_tls = true
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

`add_location` - (Optional) Appends header x-volterra-location = <re-site-name> in responses. (`Bool`).

`domains` - (Required) [This can be a domain or a sub-domain](`List of String`).

`http` - (Optional) CDN Distribution serving content over HTTP. See [Loadbalancer Type Http ](#loadbalancer-type-http) below for details.

`https` - (Optional) User is responsible for managing DNS.. See [Loadbalancer Type Https ](#loadbalancer-type-https) below for details.

`https_auto_cert` - (Optional) DNS records will be managed by Volterra.. See [Loadbalancer Type Https Auto Cert ](#loadbalancer-type-https-auto-cert) below for details.

`more_option` - (Optional) More options like header manipulation, compression etc.. See [More Option ](#more-option) below for details.

`origin_pool` - (Required) x-required. See [Origin Pool ](#origin-pool) below for details.

### More Option

More options like header manipulation, compression etc..

`cache_options` - (Optional) Cache Options. See [More Option Cache Options ](#more-option-cache-options) below for details.

`cache_ttl_options` - (Optional) Cache Options. See [More Option Cache Ttl Options ](#more-option-cache-ttl-options) below for details.(Deprecated)

`header_options` - (Optional) Request/Response header related options. See [More Option Header Options ](#more-option-header-options) below for details.

`logging_options` - (Optional) Logging related options. See [More Option Logging Options ](#more-option-logging-options) below for details.

`security_options` - (Optional) Security related options. See [More Option Security Options ](#more-option-security-options) below for details.

### Origin Pool

x-required.

`follow_origin_redirect` - (Optional) Instructs the CDN to follow redirects from the origin server(s) (`Bool`).(Deprecated)

`more_origin_options` - (Optional) x-displayName: "Advanced Configuration". See [Origin Pool More Origin Options ](#origin-pool-more-origin-options) below for details.

`origin_request_timeout` - (Optional) Configures the time after which a request to the origin will time out waiting for a response (`String`).

`origin_servers` - (Required) List of original servers. See [Origin Pool Origin Servers ](#origin-pool-origin-servers) below for details.

`public_name` - (Required) The DNS name to be used as the host header for the request to the origin server. See [Origin Pool Public Name ](#origin-pool-public-name) below for details.

###### One of the arguments from this list "no_tls, use_tls" must be set

`no_tls` - (Optional) Origin servers do not use TLS (`Bool`).

`use_tls` - (Optional) Origin servers use TLS. See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

### Action Allow

Allow the request to proceed..

### Action Deny

Deny the request..

### Action Choice Action Block

Block the request and issue an API security event.

### Action Choice Action Report

Continue processing the request and issue an API security event.

### Action Choice Action Skip

Continue processing the request.

### Action Choice Apply Data Guard

x-displayName: "Apply".

### Action Choice Block

Block the request and report the issue.

### Action Choice Bot Skip Processing

Skip Bot Defense processing for clients matching this rule..

### Action Choice Report

Allow the request and report the issue.

### Action Choice Skip Data Guard

x-displayName: "Skip".

### Action Choice Skip Processing

Skip both WAF and Bot Defense processing for clients matching this rule..

### Action Choice Waf Skip Processing

Skip WAF processing for clients matching this rule..

### Additional Headers Choice Allow Additional Headers

Allow extra headers (on top of what specified in the OAS documentation).

### Additional Headers Choice Disallow Additional Headers

Disallow extra headers (on top of what specified in the OAS documentation).

### Additional Parameters Choice Allow Additional Parameters

Allow extra query parameters (on top of what specified in the OAS documentation).

### Additional Parameters Choice Disallow Additional Parameters

Disallow extra query parameters (on top of what specified in the OAS documentation).

### Allow Introspection Queries Choice Disable Introspection

Disable introspection queries for the load balancer..

### Allow Introspection Queries Choice Enable Introspection

Enable introspection queries for the load balancer..

### Allowed Domains All Load Balancer Domains

Add All load balancer domains to source origin (allow) list..

### Allowed Domains Custom Domain List

Add one or more domains to source origin (allow) list..

`domains` - (Required) Wildcard names are supported in the suffix or prefix form. (`String`).

### Allowed Domains Disabled

Allow all source origin domains..

### Api Definition Choice Api Specification

Specify API definition and OpenAPI Validation.

`api_definition` - (Required) Specify API definition which includes application API paths and methods derived from swagger files.. See [ref](#ref) below for details.

###### One of the arguments from this list "validation_disabled, validation_all_spec_endpoints, validation_custom_list" must be set

`validation_all_spec_endpoints` - (Optional) All other API endpoints would proceed according to "Fall Through Mode". See [Validation Target Choice Validation All Spec Endpoints ](#validation-target-choice-validation-all-spec-endpoints) below for details.

`validation_custom_list` - (Optional) Any other end-points not listed will act according to "Fall Through Mode". See [Validation Target Choice Validation Custom List ](#validation-target-choice-validation-custom-list) below for details.

`validation_disabled` - (Optional) Don't run OpenAPI validation (`Bool`).

### Api Definition Choice Api Specification On Cache Miss

Enable API definition and OpenAPI Validation only on cache miss in this distribution.

`api_definition` - (Required) Specify API definition which includes application API paths and methods derived from swagger files.. See [ref](#ref) below for details.

###### One of the arguments from this list "validation_disabled, validation_all_spec_endpoints, validation_custom_list" must be set

`validation_all_spec_endpoints` - (Optional) All other API endpoints would proceed according to "Fall Through Mode". See [Validation Target Choice Validation All Spec Endpoints ](#validation-target-choice-validation-all-spec-endpoints) below for details.

`validation_custom_list` - (Optional) Any other end-points not listed will act according to "Fall Through Mode". See [Validation Target Choice Validation Custom List ](#validation-target-choice-validation-custom-list) below for details.

`validation_disabled` - (Optional) Don't run OpenAPI validation (`Bool`).

### Api Definition Choice Disable Api Definition

API Definition is not currently used for this load balancer.

### Api Discovery Choice Api Discovery On Cache Miss

Enable api discovery only on cache miss in this distribution.

`discovered_api_settings` - (Optional) Configure Discovered API Settings.. See [Api Discovery On Cache Miss Discovered Api Settings ](#api-discovery-on-cache-miss-discovered-api-settings) below for details.

###### One of the arguments from this list "disable_learn_from_redirect_traffic, enable_learn_from_redirect_traffic" must be set

`disable_learn_from_redirect_traffic` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`enable_learn_from_redirect_traffic` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`sensitive_data_detection_rules` - (Optional) Manage rules to detect sensitive data in requests and/or response sections.. See [Api Discovery On Cache Miss Sensitive Data Detection Rules ](#api-discovery-on-cache-miss-sensitive-data-detection-rules) below for details.

### Api Discovery Choice Disable Api Discovery

Disable api discovery for this distribution.

### Api Discovery Choice Enable Api Discovery

Enable api discovery for all requests in this distribution.

`discovered_api_settings` - (Optional) Configure Discovered API Settings.. See [Enable Api Discovery Discovered Api Settings ](#enable-api-discovery-discovered-api-settings) below for details.

###### One of the arguments from this list "enable_learn_from_redirect_traffic, disable_learn_from_redirect_traffic" must be set

`disable_learn_from_redirect_traffic` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`enable_learn_from_redirect_traffic` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`sensitive_data_detection_rules` - (Optional) Manage rules to detect sensitive data in requests and/or response sections.. See [Enable Api Discovery Sensitive Data Detection Rules ](#enable-api-discovery-sensitive-data-detection-rules) below for details.

### Api Discovery On Cache Miss Discovered Api Settings

Configure Discovered API Settings..

### Api Discovery On Cache Miss Sensitive Data Detection Rules

Manage rules to detect sensitive data in requests and/or response sections..

### Api Endpoint Rules Action

The action to take if the input request matches the rule..

###### One of the arguments from this list "allow, deny" must be set

`allow` - (Optional) Allow the request to proceed. (`Bool`).

`deny` - (Optional) Deny the request. (`Bool`).

### Api Endpoint Rules Api Endpoint Method

The predicate evaluates to true if the actual HTTP method belongs is present in the list of expected values..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`methods` - (Optional) x-example: "['GET', 'POST', 'DELETE']" (`List of Strings`).

### Api Endpoint Rules Client Matcher

Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc..

###### One of the arguments from this list "ip_threat_category_list, client_selector, any_client" must be set

`any_client` - (Optional) Any Client (`Bool`).

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.

`ip_threat_category_list` - (Optional) IP threat categories to choose from. See [Client Choice Ip Threat Category List ](#client-choice-ip-threat-category-list) below for details.

###### One of the arguments from this list "any_ip, ip_prefix_list, ip_matcher, asn_list, asn_matcher" must be set

`any_ip` - (Optional) Any Source IP (`Bool`).

`asn_list` - (Optional) The predicate evaluates to true if the origin ASN is present in the ASN list.. See [Ip Asn Choice Asn List ](#ip-asn-choice-asn-list) below for details.

`asn_matcher` - (Optional) The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects.. See [Ip Asn Choice Asn Matcher ](#ip-asn-choice-asn-matcher) below for details.

`ip_matcher` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes in the IP Prefix Sets.. See [Ip Asn Choice Ip Matcher ](#ip-asn-choice-ip-matcher) below for details.

`ip_prefix_list` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes from the list.. See [Ip Asn Choice Ip Prefix List ](#ip-asn-choice-ip-prefix-list) below for details.

`tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints.. See [Client Matcher Tls Fingerprint Matcher ](#client-matcher-tls-fingerprint-matcher) below for details.

### Api Endpoint Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Api Endpoint Rules Request Matcher

Conditions related to the request, such as query parameters, headers, etc..

`cookie_matchers` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Request Matcher Cookie Matchers ](#request-matcher-cookie-matchers) below for details.

`headers` - (Optional) Note that all specified header predicates must evaluate to true.. See [Request Matcher Headers ](#request-matcher-headers) below for details.

`jwt_claims` - (Optional) Note that this feature only works on LBs with JWT Validation feature enabled.. See [Request Matcher Jwt Claims ](#request-matcher-jwt-claims) below for details.

`query_params` - (Optional) Note that all specified query parameter predicates must evaluate to true.. See [Request Matcher Query Params ](#request-matcher-query-params) below for details.

### Api Groups Rules Action

The action to take if the input request matches the rule..

###### One of the arguments from this list "allow, deny" must be set

`allow` - (Optional) Allow the request to proceed. (`Bool`).

`deny` - (Optional) Deny the request. (`Bool`).

### Api Groups Rules Client Matcher

Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc..

###### One of the arguments from this list "any_client, ip_threat_category_list, client_selector" must be set

`any_client` - (Optional) Any Client (`Bool`).

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.

`ip_threat_category_list` - (Optional) IP threat categories to choose from. See [Client Choice Ip Threat Category List ](#client-choice-ip-threat-category-list) below for details.

###### One of the arguments from this list "any_ip, ip_prefix_list, ip_matcher, asn_list, asn_matcher" must be set

`any_ip` - (Optional) Any Source IP (`Bool`).

`asn_list` - (Optional) The predicate evaluates to true if the origin ASN is present in the ASN list.. See [Ip Asn Choice Asn List ](#ip-asn-choice-asn-list) below for details.

`asn_matcher` - (Optional) The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects.. See [Ip Asn Choice Asn Matcher ](#ip-asn-choice-asn-matcher) below for details.

`ip_matcher` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes in the IP Prefix Sets.. See [Ip Asn Choice Ip Matcher ](#ip-asn-choice-ip-matcher) below for details.

`ip_prefix_list` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes from the list.. See [Ip Asn Choice Ip Prefix List ](#ip-asn-choice-ip-prefix-list) below for details.

`tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints.. See [Client Matcher Tls Fingerprint Matcher ](#client-matcher-tls-fingerprint-matcher) below for details.

### Api Groups Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Api Groups Rules Request Matcher

Conditions related to the request, such as query parameters, headers, etc..

`cookie_matchers` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Request Matcher Cookie Matchers ](#request-matcher-cookie-matchers) below for details.

`headers` - (Optional) Note that all specified header predicates must evaluate to true.. See [Request Matcher Headers ](#request-matcher-headers) below for details.

`jwt_claims` - (Optional) Note that this feature only works on LBs with JWT Validation feature enabled.. See [Request Matcher Jwt Claims ](#request-matcher-jwt-claims) below for details.

`query_params` - (Optional) Note that all specified query parameter predicates must evaluate to true.. See [Request Matcher Query Params ](#request-matcher-query-params) below for details.

### Api Protection Api Protection Rules

Rules can also include additional conditions, for example specific clients can access certain API endpoint or API group..

`api_endpoint_rules` - (Optional) If request matches any of these rules, skipping second category rules.. See [Api Protection Rules Api Endpoint Rules ](#api-protection-rules-api-endpoint-rules) below for details.

`api_groups_rules` - (Optional) For API groups, refer to API Definition which includes API groups derived from uploaded swaggers.. See [Api Protection Rules Api Groups Rules ](#api-protection-rules-api-groups-rules) below for details.

### Api Protection Jwt Validation

tokens or tokens that are not yet valid..

`action` - (Required) x-required. See [Jwt Validation Action ](#jwt-validation-action) below for details.

###### One of the arguments from this list "auth_server_uri, jwks, jwks_config" must be set

`auth_server_uri` - (Optional) JWKS URI will be will be retrieved from this URI (`String`).(Deprecated)

`jwks` - (Optional) The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details. (`String`).(Deprecated)

`jwks_config` - (Optional) The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details.. See [Jwks Configuration Jwks Config ](#jwks-configuration-jwks-config) below for details.

`mandatory_claims` - (Optional) If the claim does not exist JWT token validation will fail.. See [Jwt Validation Mandatory Claims ](#jwt-validation-mandatory-claims) below for details.

`reserved_claims` - (Optional) the token validation of these claims should be disabled.. See [Jwt Validation Reserved Claims ](#jwt-validation-reserved-claims) below for details.

`target` - (Required) Define endpoints for which JWT token validation will be performed. See [Jwt Validation Target ](#jwt-validation-target) below for details.

`token_location` - (Required) Define where in the HTTP request the JWT token will be extracted. See [Jwt Validation Token Location ](#jwt-validation-token-location) below for details.

### Api Protection Rules Api Endpoint Rules

If request matches any of these rules, skipping second category rules..

`action` - (Required) The action to take if the input request matches the rule.. See [Api Endpoint Rules Action ](#api-endpoint-rules-action) below for details.

`api_endpoint_method` - (Optional) The predicate evaluates to true if the actual HTTP method belongs is present in the list of expected values.. See [Api Endpoint Rules Api Endpoint Method ](#api-endpoint-rules-api-endpoint-method) below for details.

`api_endpoint_path` - (Required) The endpoint (path) of the request. (`String`).

`client_matcher` - (Optional) Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc.. See [Api Endpoint Rules Client Matcher ](#api-endpoint-rules-client-matcher) below for details.

###### One of the arguments from this list "any_domain, specific_domain" must be set

`any_domain` - (Optional) The rule will apply for all domains. (`Bool`).

`specific_domain` - (Optional) For example: api.example.com (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Api Endpoint Rules Metadata ](#api-endpoint-rules-metadata) below for details.

`request_matcher` - (Optional) Conditions related to the request, such as query parameters, headers, etc.. See [Api Endpoint Rules Request Matcher ](#api-endpoint-rules-request-matcher) below for details.

### Api Protection Rules Api Groups Rules

For API groups, refer to API Definition which includes API groups derived from uploaded swaggers..

`action` - (Required) The action to take if the input request matches the rule.. See [Api Groups Rules Action ](#api-groups-rules-action) below for details.

`api_group` - (Optional) Custom groups can be created if user tags paths or operations with "x-volterra-api-group" extensions inside swaggers. (`String`).

`base_path` - (Required) For example: /v1 (`String`).

`client_matcher` - (Optional) Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc.. See [Api Groups Rules Client Matcher ](#api-groups-rules-client-matcher) below for details.

###### One of the arguments from this list "any_domain, specific_domain" must be set

`any_domain` - (Optional) The rule will apply for all domains. (`Bool`).

`specific_domain` - (Optional) For example: api.example.com (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Api Groups Rules Metadata ](#api-groups-rules-metadata) below for details.

`request_matcher` - (Optional) Conditions related to the request, such as query parameters, headers, etc.. See [Api Groups Rules Request Matcher ](#api-groups-rules-request-matcher) below for details.

### Api Rate Limit Api Endpoint Rules

For creating rule that contain a whole domain or group of endpoints, please use the server URL rules above..

`api_endpoint_method` - (Optional) The predicate evaluates to true if the actual HTTP method belongs is present in the list of expected values.. See [Api Endpoint Rules Api Endpoint Method ](#api-endpoint-rules-api-endpoint-method) below for details.

`api_endpoint_path` - (Required) The endpoint (path) of the request. (`String`).

`base_path` - (Optional) The request base path. (`String`).(Deprecated)

`client_matcher` - (Optional) Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc.. See [Api Endpoint Rules Client Matcher ](#api-endpoint-rules-client-matcher) below for details.

###### One of the arguments from this list "any_domain, specific_domain" must be set

`any_domain` - (Optional) The rule will apply for all domains. (`Bool`).

`specific_domain` - (Optional) The rule will apply for a specific domain. (`String`).

###### One of the arguments from this list "inline_rate_limiter, ref_rate_limiter" must be set

`inline_rate_limiter` - (Optional) Specify rate values for the rule.. See [Rate Limiter Choice Inline Rate Limiter ](#rate-limiter-choice-inline-rate-limiter) below for details.

`ref_rate_limiter` - (Optional) Select external rate limiter.. See [ref](#ref) below for details.

`request_matcher` - (Optional) Conditions related to the request, such as query parameters, headers, etc.. See [Api Endpoint Rules Request Matcher ](#api-endpoint-rules-request-matcher) below for details.

### Api Rate Limit Server Url Rules

For matching also specific endpoints you can use the API endpoint rules set bellow..

`api_group` - (Optional) Custom groups can be created if user tags paths or operations with "x-volterra-api-group" extensions inside swaggers. (`String`).

`base_path` - (Required) Prefix of the request path. (`String`).

`client_matcher` - (Optional) Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc.. See [Server Url Rules Client Matcher ](#server-url-rules-client-matcher) below for details.

###### One of the arguments from this list "any_domain, specific_domain" must be set

`any_domain` - (Optional) The rule will apply for all domains. (`Bool`).

`specific_domain` - (Optional) The rule will apply for a specific domain. (`String`).

###### One of the arguments from this list "inline_rate_limiter, ref_rate_limiter" must be set

`inline_rate_limiter` - (Optional) Specify rate values for the rule.. See [Rate Limiter Choice Inline Rate Limiter ](#rate-limiter-choice-inline-rate-limiter) below for details.

`ref_rate_limiter` - (Optional) Use external rate limiter.. See [ref](#ref) below for details.

`request_matcher` - (Optional) Conditions related to the request, such as query parameters, headers, etc.. See [Server Url Rules Request Matcher ](#server-url-rules-request-matcher) below for details.

### App Firewall Detection Control Exclude Attack Type Contexts

Attack Types to be excluded for the defined match criteria.

`context` - (Required) x-required (`String`).

`context_name` - (Optional) Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. (`String`).

`exclude_attack_type` - (Required) x-required (`String`).

### App Firewall Detection Control Exclude Bot Name Contexts

Bot Names to be excluded for the defined match criteria.

`bot_name` - (Required) x-example: "Hydra" (`String`).

### App Firewall Detection Control Exclude Signature Contexts

Signature IDs to be excluded for the defined match criteria.

`context` - (Required) x-required (`String`).

`context_name` - (Optional) Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. (`String`).

`signature_id` - (Required) 0 implies that all signatures will be excluded for the specified context. (`Int`).

### App Firewall Detection Control Exclude Violation Contexts

Violations to be excluded for the defined match criteria.

`context` - (Required) x-required (`String`).

`context_name` - (Optional) Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. (`String`).

`exclude_violation` - (Required) x-required (`String`).

### Asn Choice Any Asn

any_asn.

### Asn Choice Asn List

asn_list.

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer. (`Int`).

### Asn Choice Asn Matcher

asn_matcher.

`asn_sets` - (Required) A list of references to bgp_asn_set objects.. See [ref](#ref) below for details.

### Audience Validation Audience

x-displayName: "Exact Match".

`audiences` - (Required) x-required (`String`).

### Audience Validation Audience Disable

x-displayName: "Disable".

### Auth Options Custom

Enable Custom Authentication.

`custom_auth_config` - (Optional) This is custom authentication configuration parameters. Please reach out to the support for custom authentication details. (`String`).

### Auth Options Disable Auth

No Authentication.

### Auth Options Jwt

Enable JWT Authentication.

`backup_key` - (Optional) Backup JWT Key - If specified is also checked in addition to the primary secret key. See [Jwt Backup Key ](#jwt-backup-key) below for details.

`secret_key` - (Required) Secret Key for JWT. See [Jwt Secret Key ](#jwt-secret-key) below for details.

###### One of the arguments from this list "header, cookie, query_param, bearer_token" can be set

`bearer_token` - (Optional) Token is found in the Bearer-Token (`Bool`).

`cookie` - (Optional) Token is found in the cookie. See [Token Source Cookie ](#token-source-cookie) below for details.

`header` - (Optional) Token is found in the header. See [Token Source Header ](#token-source-header) below for details.

`query_param` - (Optional) Token is found in the Query-Param. See [Token Source Query Param ](#token-source-query-param) below for details.

### Backup Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Blocked Clients Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Bypass Rate Limiting Rules Bypass Rate Limiting Rules

This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting..

`client_matcher` - (Optional) Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc.. See [Bypass Rate Limiting Rules Client Matcher ](#bypass-rate-limiting-rules-client-matcher) below for details.

###### One of the arguments from this list "any_url, base_path, api_endpoint, api_groups" must be set

`any_url` - (Optional) Any URL (`Bool`).

`api_endpoint` - (Required) The endpoint (path) of the request.. See [Destination Type Api Endpoint ](#destination-type-api-endpoint) below for details.

`api_groups` - (Optional) Validation will be performed for the endpoints mentioned in the API Groups. See [Destination Type Api Groups ](#destination-type-api-groups) below for details.

`base_path` - (Optional) The base path which this validation applies to (`String`).

###### One of the arguments from this list "any_domain, specific_domain" must be set

`any_domain` - (Optional) The rule will apply for all domains. (`Bool`).

`specific_domain` - (Optional) For example: api.example.com (`String`).

`request_matcher` - (Optional) Conditions related to the request, such as query parameters, headers, etc.. See [Bypass Rate Limiting Rules Request Matcher ](#bypass-rate-limiting-rules-request-matcher) below for details.

### Bypass Rate Limiting Rules Client Matcher

Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc..

###### One of the arguments from this list "any_client, ip_threat_category_list, client_selector" must be set

`any_client` - (Optional) Any Client (`Bool`).

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.

`ip_threat_category_list` - (Optional) IP threat categories to choose from. See [Client Choice Ip Threat Category List ](#client-choice-ip-threat-category-list) below for details.

###### One of the arguments from this list "any_ip, ip_prefix_list, ip_matcher, asn_list, asn_matcher" must be set

`any_ip` - (Optional) Any Source IP (`Bool`).

`asn_list` - (Optional) The predicate evaluates to true if the origin ASN is present in the ASN list.. See [Ip Asn Choice Asn List ](#ip-asn-choice-asn-list) below for details.

`asn_matcher` - (Optional) The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects.. See [Ip Asn Choice Asn Matcher ](#ip-asn-choice-asn-matcher) below for details.

`ip_matcher` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes in the IP Prefix Sets.. See [Ip Asn Choice Ip Matcher ](#ip-asn-choice-ip-matcher) below for details.

`ip_prefix_list` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes from the list.. See [Ip Asn Choice Ip Prefix List ](#ip-asn-choice-ip-prefix-list) below for details.

`tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints.. See [Client Matcher Tls Fingerprint Matcher ](#client-matcher-tls-fingerprint-matcher) below for details.

### Bypass Rate Limiting Rules Request Matcher

Conditions related to the request, such as query parameters, headers, etc..

`cookie_matchers` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Request Matcher Cookie Matchers ](#request-matcher-cookie-matchers) below for details.

`headers` - (Optional) Note that all specified header predicates must evaluate to true.. See [Request Matcher Headers ](#request-matcher-headers) below for details.

`jwt_claims` - (Optional) Note that this feature only works on LBs with JWT Validation feature enabled.. See [Request Matcher Jwt Claims ](#request-matcher-jwt-claims) below for details.

`query_params` - (Optional) Note that all specified query parameter predicates must evaluate to true.. See [Request Matcher Query Params ](#request-matcher-query-params) below for details.

### Cache Actions Cache Bypass

Bypass Caching of content from the origin.

### Cache Actions Cache Disabled

Disable Caching of content from the origin.

### Cache Actions Eligible For Cache

Eligible for caching the content.

###### One of the arguments from this list "scheme_hostname_request_uri, hostname_uri, scheme_hostname_uri_query, scheme_proxy_host_uri, scheme_proxy_host_request_uri, scheme_hostname_uri" must be set

`hostname_uri` - (Optional) . See [Eligible For Cache Hostname Uri ](#eligible-for-cache-hostname-uri) below for details.(Deprecated)

`scheme_hostname_request_uri` - (Optional) . See [Eligible For Cache Scheme Hostname Request Uri ](#eligible-for-cache-scheme-hostname-request-uri) below for details.(Deprecated)

`scheme_hostname_uri` - (Optional) . See [Eligible For Cache Scheme Hostname Uri ](#eligible-for-cache-scheme-hostname-uri) below for details.(Deprecated)

`scheme_hostname_uri_query` - (Optional) . See [Eligible For Cache Scheme Hostname Uri Query ](#eligible-for-cache-scheme-hostname-uri-query) below for details.(Deprecated)

`scheme_proxy_host_request_uri` - (Optional) . See [Eligible For Cache Scheme Proxy Host Request Uri ](#eligible-for-cache-scheme-proxy-host-request-uri) below for details.

`scheme_proxy_host_uri` - (Optional) . See [Eligible For Cache Scheme Proxy Host Uri ](#eligible-for-cache-scheme-proxy-host-uri) below for details.

### Cache Actions Eligible For Cache

Eligible for caching the content.

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

### Cache Headers Operator

Available operators.

###### One of the arguments from this list "Equals, DoesNotEqual, Endswith, DoesNotEndWith, MatchRegex, Contains, DoesNotContain, Startswith, DoesNotStartWith" can be set

`Contains` - (Optional) Field must contain (`String`).

`DoesNotContain` - (Optional) Field must not contain (`String`).

`DoesNotEndWith` - (Optional) Field must not end with (`String`).

`DoesNotEqual` - (Optional) Field must not equal (`String`).

`DoesNotStartWith` - (Optional) Field must not start with (`String`).

`Endswith` - (Optional) Field must end with (`String`).

`Equals` - (Optional) Field must exactly match (`String`).

`MatchRegex` - (Optional) Field matches regular expression (`String`).

`Startswith` - (Optional) Field must start with (`String`).

### Cache Options Cache Rules

Rules are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs..

###### One of the arguments from this list "cache_bypass, eligible_for_cache" must be set

`cache_bypass` - (Optional) Bypass Caching of content from the origin (`Bool`).

`eligible_for_cache` - (Optional) Eligible for caching the content. See [Cache Actions Eligible For Cache ](#cache-actions-eligible-for-cache) below for details.

`rule_expression_list` - (Required) Expressions are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs... See [Cache Rules Rule Expression List ](#cache-rules-rule-expression-list) below for details.

`rule_name` - (Required) Name of the Cache Rule (`String`).

### Cache Options Default Cache Action

Default value for Cache action..

###### One of the arguments from this list "cache_ttl_override, cache_disabled, eligible_for_cache, cache_ttl_default" can be set

`cache_disabled` - (Optional) Disable Caching of content from the origin (`Bool`).

`cache_ttl_default` - (Optional) Cache TTL value to use when the origin does not provide one (`String`).

`cache_ttl_override` - (Optional) Override the Cache TTL directive in the response from the origin (`String`).

`eligible_for_cache` - (Optional) Eligible for caching the content. See [Cache Actions Eligible For Cache ](#cache-actions-eligible-for-cache) below for details.(Deprecated)

### Cache Rule Expression Cache Headers

Configure cache rule headers to match the criteria.

`name` - (Optional) Name of the header (`String`).

`operator` - (Optional) Available operators. See [Cache Headers Operator ](#cache-headers-operator) below for details.

### Cache Rule Expression Cookie Matcher

Note that all specified cookie matcher predicates must evaluate to true..

`name` - (Required) A case-sensitive cookie name. (`String`).

`operator` - (Optional) . See [Cookie Matcher Operator ](#cookie-matcher-operator) below for details.

### Cache Rule Expression Path Match

URI path of route.

`operator` - (Optional) A specification of path match. See [Path Match Operator ](#path-match-operator) below for details.

### Cache Rule Expression Query Parameters

List of (key, value) query parameters.

`key` - (Required) In the above example, assignee_username is the key (`String`).

`operator` - (Optional) . See [Query Parameters Operator ](#query-parameters-operator) below for details.

### Cache Rules Rule Expression List

Expressions are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs...

`cache_rule_expression` - (Required) The Cache Rule Expression Terms that are ANDed. See [Rule Expression List Cache Rule Expression ](#rule-expression-list-cache-rule-expression) below for details.

`expression_name` - (Required) Name of the Expressions items that are ANDed (`String`).

### Captcha Challenge Parameters Choice Captcha Challenge Parameters

Configure captcha challenge parameters.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

### Captcha Challenge Parameters Choice Default Captcha Challenge Parameters

Use default parameters.

### Challenge Action Disable Challenge

Disable the challenge type selected in PolicyBasedChallenge.

### Challenge Action Enable Captcha Challenge

Enable captcha challenge.

### Challenge Action Enable Javascript Challenge

Enable javascript challenge.

### Challenge Choice Always Enable Captcha Challenge

Challenge rules can be used to selectively disable Captcha challenge or enable JavaScript challenge for some requests..

### Challenge Choice Always Enable Js Challenge

Challenge rules can be used to selectively disable JavaScript challenge or enable Captcha challenge for some requests..

### Challenge Choice No Challenge

Challenge rules can be used to selectively enable JavaScript or Captcha challenge for some requests..

### Challenge Type Captcha Challenge

Configure Captcha challenge on this load balancer.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

### Challenge Type Challenge On Cache Miss

Configure auto mitigation i.e risk based challenges for malicious users only on cache miss in this load balancer.

###### One of the arguments from this list "default_captcha_challenge_parameters, captcha_challenge_parameters" can be set

`captcha_challenge_parameters` - (Optional) Configure captcha challenge parameters. See [Captcha Challenge Parameters Choice Captcha Challenge Parameters ](#captcha-challenge-parameters-choice-captcha-challenge-parameters) below for details.

`default_captcha_challenge_parameters` - (Optional) Use default parameters (`Bool`).

###### One of the arguments from this list "default_js_challenge_parameters, js_challenge_parameters" can be set

`default_js_challenge_parameters` - (Optional) Use default parameters (`Bool`).

`js_challenge_parameters` - (Optional) Configure JavaScript challenge parameters. See [Js Challenge Parameters Choice Js Challenge Parameters ](#js-challenge-parameters-choice-js-challenge-parameters) below for details.

###### One of the arguments from this list "default_mitigation_settings, malicious_user_mitigation" can be set

`default_mitigation_settings` - (Optional) For high level, users will be temporarily blocked. (`Bool`).

`malicious_user_mitigation` - (Optional) Define the mitigation actions to be taken for different threat levels. See [ref](#ref) below for details.

### Challenge Type Enable Challenge

Configure auto mitigation i.e risk based challenges for malicious users for this load balancer.

###### One of the arguments from this list "default_captcha_challenge_parameters, captcha_challenge_parameters" can be set

`captcha_challenge_parameters` - (Optional) Configure captcha challenge parameters. See [Captcha Challenge Parameters Choice Captcha Challenge Parameters ](#captcha-challenge-parameters-choice-captcha-challenge-parameters) below for details.

`default_captcha_challenge_parameters` - (Optional) Use default parameters (`Bool`).

###### One of the arguments from this list "default_js_challenge_parameters, js_challenge_parameters" can be set

`default_js_challenge_parameters` - (Optional) Use default parameters (`Bool`).

`js_challenge_parameters` - (Optional) Configure JavaScript challenge parameters. See [Js Challenge Parameters Choice Js Challenge Parameters ](#js-challenge-parameters-choice-js-challenge-parameters) below for details.

###### One of the arguments from this list "default_mitigation_settings, malicious_user_mitigation" can be set

`default_mitigation_settings` - (Optional) For high level, users will be temporarily blocked. (`Bool`).

`malicious_user_mitigation` - (Optional) Define the mitigation actions to be taken for different threat levels. See [ref](#ref) below for details.

### Challenge Type Js Challenge

Configure JavaScript challenge on this load balancer.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

`js_script_delay` - (Optional) Delay introduced by Javascript, in milliseconds. (`Int`).

### Challenge Type No Challenge

No challenge is enabled for this load balancer.

### Challenge Type Policy Based Challenge

Specifies the settings for policy rule based challenge.

###### One of the arguments from this list "default_captcha_challenge_parameters, captcha_challenge_parameters" can be set

`captcha_challenge_parameters` - (Optional) Configure captcha challenge parameters. See [Captcha Challenge Parameters Choice Captcha Challenge Parameters ](#captcha-challenge-parameters-choice-captcha-challenge-parameters) below for details.

`default_captcha_challenge_parameters` - (Optional) Use default parameters (`Bool`).

###### One of the arguments from this list "no_challenge, always_enable_js_challenge, always_enable_captcha_challenge" must be set

`always_enable_captcha_challenge` - (Optional) Challenge rules can be used to selectively disable Captcha challenge or enable JavaScript challenge for some requests. (`Bool`).

`always_enable_js_challenge` - (Optional) Challenge rules can be used to selectively disable JavaScript challenge or enable Captcha challenge for some requests. (`Bool`).

`no_challenge` - (Optional) Challenge rules can be used to selectively enable JavaScript or Captcha challenge for some requests. (`Bool`).

###### One of the arguments from this list "js_challenge_parameters, default_js_challenge_parameters" can be set

`default_js_challenge_parameters` - (Optional) Use default parameters (`Bool`).

`js_challenge_parameters` - (Optional) Configure JavaScript challenge parameters. See [Js Challenge Parameters Choice Js Challenge Parameters ](#js-challenge-parameters-choice-js-challenge-parameters) below for details.

###### One of the arguments from this list "default_mitigation_settings, malicious_user_mitigation" can be set

`default_mitigation_settings` - (Optional) For high level, users will be temporarily blocked. (`Bool`).

`malicious_user_mitigation` - (Optional) Define the mitigation actions to be taken for different threat levels. See [ref](#ref) below for details.

`rule_list` - (Optional) list challenge rules to be used in policy based challenge. See [Policy Based Challenge Rule List ](#policy-based-challenge-rule-list) below for details.

###### One of the arguments from this list "default_temporary_blocking_parameters, temporary_user_blocking" can be set

`default_temporary_blocking_parameters` - (Optional) Use default parameters (`Bool`).(Deprecated)

`temporary_user_blocking` - (Optional) Specifies configuration for temporary user blocking resulting from malicious user detection. See [Temporary Blocking Parameters Choice Temporary User Blocking ](#temporary-blocking-parameters-choice-temporary-user-blocking) below for details.(Deprecated)

### Choice Custom Security

Custom selection of TLS versions and cipher suites.

`cipher_suites` - (Required) The TLS listener will only support the specified cipher list. (`String`).

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

### Choice Default Security

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Choice Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Choice Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### Choice Public Ip

Specify origin server with public IP.

###### One of the arguments from this list "ip, ipv6" must be set

`ip` - (Optional) Public IPV4 address (`String`).

`ipv6` - (Optional) Public IPV6 address (`String`).

### Choice Public Name

Specify origin server with public DNS name.

`dns_name` - (Required) DNS Name (`String`).

`refresh_interval` - (Optional) Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767 (`Int`).

### Choice Tls 11 Plus

TLS v1.1+ with PFS ciphers and medium strength crypto algorithms..

### Choice Tls 12 Plus

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Client Choice Any Client

Any Client.

### Client Choice Client Name Matcher

client_name_matcher.

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Client Choice Client Selector

The predicate evaluates to true if the expressions in the label selector are true for the client labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Client Choice Ip Threat Category List

IP threat categories to choose from.

`ip_threat_categories` - (Required) The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions (`List of Strings`).

### Client Matcher Tls Fingerprint Matcher

The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints..

`classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`exact_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

`excluded_values` - (Optional) or more known TLS fingerprint classes in the enclosing matcher. (`String`).

### Client Source Choice Http Header

Request header name and value pairs.

`headers` - (Required) List of HTTP header name and value pairs. See [Http Header Headers ](#http-header-headers) below for details.

### Common Security Controls Blocked Clients

Define rules to block IP Prefixes or AS numbers..

###### One of the arguments from this list "skip_processing, waf_skip_processing, bot_skip_processing" can be set

`bot_skip_processing` - (Optional) Skip Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`skip_processing` - (Optional) Skip both WAF and Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (`Bool`).(Deprecated)

`actions` - (Optional) Actions that should be taken when client identifier matches the rule (`List of Strings`).

###### One of the arguments from this list "as_number, http_header, user_identifier, ip_prefix" must be set

`as_number` - (Optional) RFC 6793 defined 4-byte AS number (`Int`).

`http_header` - (Optional) Request header name and value pairs. See [Client Source Choice Http Header ](#client-source-choice-http-header) below for details.

`ip_prefix` - (Optional) IPv4 prefix string. (`String`).

`user_identifier` - (Optional) Identify user based on user identifier. User identifier value needs to be copied from security event. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Blocked Clients Metadata ](#blocked-clients-metadata) below for details.

### Common Security Controls Cors Policy

resources from a server at a different origin.

`allow_credentials` - (Optional) Specifies whether the resource allows credentials (`Bool`).

`allow_headers` - (Optional) Specifies the content for the access-control-allow-headers header (`String`).

`allow_methods` - (Optional) Specifies the content for the access-control-allow-methods header (`String`).

`allow_origin` - (Optional) An origin is allowed if either allow_origin or allow_origin_regex match (`String`).

`allow_origin_regex` - (Optional) An origin is allowed if either allow_origin or allow_origin_regex match (`String`).

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`expose_headers` - (Optional) Specifies the content for the access-control-expose-headers header (`String`).

`max_age` - (Optional) Specifies the content for the access-control-max-age header (`String`).(Deprecated)

`maximum_age` - (Optional) Maximum permitted value is 86400 seconds (24 hours) (`Int`).

### Common Security Controls Trusted Clients

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients..

###### One of the arguments from this list "skip_processing, waf_skip_processing, bot_skip_processing" can be set

`bot_skip_processing` - (Optional) Skip Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`skip_processing` - (Optional) Skip both WAF and Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (`Bool`).(Deprecated)

`actions` - (Optional) Actions that should be taken when client identifier matches the rule (`List of Strings`).

###### One of the arguments from this list "user_identifier, ip_prefix, as_number, http_header" must be set

`as_number` - (Optional) RFC 6793 defined 4-byte AS number (`Int`).

`http_header` - (Optional) Request header name and value pairs. See [Client Source Choice Http Header ](#client-source-choice-http-header) below for details.

`ip_prefix` - (Optional) IPv4 prefix string. (`String`).

`user_identifier` - (Optional) Identify user based on user identifier. User identifier value needs to be copied from security event. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Trusted Clients Metadata ](#trusted-clients-metadata) below for details.

### Condition Type Choice Api Endpoint

The API endpoint (Path + Method) which this validation applies to.

`methods` - (Optional) Methods to be matched (`List of Strings`).

`path` - (Required) Path to be matched (`String`).

### Cookie Matcher Operator

.

###### One of the arguments from this list "Startswith, DoesNotStartWith, Contains, DoesNotContain, Endswith, DoesNotEndWith, MatchRegex, Equals, DoesNotEqual" can be set

`Contains` - (Optional) Field must contain (`String`).

`DoesNotContain` - (Optional) Field must not contain (`String`).

`DoesNotEndWith` - (Optional) Field must not end with (`String`).

`DoesNotEqual` - (Optional) Field must not equal (`String`).

`DoesNotStartWith` - (Optional) Field must not start with (`String`).

`Endswith` - (Optional) Field must end with (`String`).

`Equals` - (Optional) Field must exactly match (`String`).

`MatchRegex` - (Optional) Field matches regular expression (`String`).

`Startswith` - (Optional) Field must start with (`String`).

### Cookie Tampering Disable Tampering Protection

x-displayName: "Disable".

### Cookie Tampering Enable Tampering Protection

x-displayName: "Enable".

### Count By Choice Use Http Lb User Id

Defined in HTTP-LB Security Configuration -> User Identifier..

### Data Guard Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Data Guard Rules Path

URI path matcher..

###### One of the arguments from this list "prefix, path, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Destination Type Any Url

Any URL .

### Destination Type Api Endpoint

The endpoint (path) of the request..

`methods` - (Optional) Methods to be matched (`List of Strings`).

`path` - (Required) Path to be matched (`String`).

### Destination Type Api Groups

Validation will be performed for the endpoints mentioned in the API Groups.

`api_groups` - (Required) x-required (`String`).

### Domain Choice Any Domain

The rule will apply for all domains..

### Eligible For Cache Hostname Uri

.

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Eligible For Cache Scheme Hostname Request Uri

.

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Eligible For Cache Scheme Hostname Uri

.

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Eligible For Cache Scheme Hostname Uri Query

.

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Eligible For Cache Scheme Proxy Host Request Uri

.

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Eligible For Cache Scheme Proxy Host Uri

.

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Enable Api Discovery Discovered Api Settings

Configure Discovered API Settings..

### Enable Api Discovery Sensitive Data Detection Rules

Manage rules to detect sensitive data in requests and/or response sections..

### Fall Through Mode Choice Fall Through Mode Allow

Allow any unprotected end point.

### Fall Through Mode Choice Fall Through Mode Custom

Custom rules for any unprotected end point.

`open_api_validation_rules` - (Required) x-displayName: "Custom Fall Through Rule List". See [Fall Through Mode Custom Open Api Validation Rules ](#fall-through-mode-custom-open-api-validation-rules) below for details.

### Fall Through Mode Custom Open Api Validation Rules

x-displayName: "Custom Fall Through Rule List".

###### One of the arguments from this list "action_skip, action_report, action_block" must be set

`action_block` - (Optional) Block the request and issue an API security event (`Bool`).

`action_report` - (Optional) Continue processing the request and issue an API security event (`Bool`).

`action_skip` - (Optional) Continue processing the request (`Bool`).

###### One of the arguments from this list "api_endpoint, base_path, api_group" must be set

`api_endpoint` - (Optional) The API endpoint (Path + Method) which this validation applies to. See [Condition Type Choice Api Endpoint ](#condition-type-choice-api-endpoint) below for details.

`api_group` - (Optional) The API group which this validation applies to (`String`).

`base_path` - (Optional) The base path which this validation applies to (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Open Api Validation Rules Metadata ](#open-api-validation-rules-metadata) below for details.

### Geo Filtering Type Allow List

Allow list of countries.

`country_codes` - (Required) List of Country Codes (`List of Strings`).

`invert_match` - (Optional) Invert the match result. (`Bool`).

### Geo Filtering Type Block List

Block list of countries.

`country_codes` - (Required) List of Country Codes (`List of Strings`).

`invert_match` - (Optional) Invert the match result. (`Bool`).

### Graphql Rules Graphql Settings

GraphQL configuration..

###### One of the arguments from this list "disable_introspection, enable_introspection" must be set

`disable_introspection` - (Optional) Disable introspection queries for the load balancer. (`Bool`).

`enable_introspection` - (Optional) Enable introspection queries for the load balancer. (`Bool`).

`max_batched_queries` - (Required) Specify maximum number of queries in a single batched request. (`Int`).

`max_depth` - (Required) Specify maximum depth for the GraphQL query. (`Int`).

`max_total_length` - (Required) Specify maximum length in bytes for the GraphQL query. (`Int`).

`max_value_length` - (Required) Specify maximum value length in bytes for the GraphQL query. (`Int`).(Deprecated)

`policy_name` - (Optional) Sets the BD Policy to use (`String`).(Deprecated)

### Graphql Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Header Options Request Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "value, secret_value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Header Options Response Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "value, secret_value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Http Header Headers

List of HTTP header name and value pairs.

`invert_match` - (Optional) Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Required) Name of the header (`String`).

###### One of the arguments from this list "exact, regex, presence" can be set

`exact` - (Optional) Header value to match exactly (`String`).

`presence` - (Optional) If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Httponly Add Httponly

x-displayName: "Add".

### Httponly Ignore Httponly

x-displayName: "Ignore".

### Https Tls Parameters

TLS parameters for the downstream connections..

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Parameters Tls Certificates ](#tls-parameters-tls-certificates) below for details.

`tls_config` - (Optional) TLS Configuration Parameters. See [Tls Parameters Tls Config ](#tls-parameters-tls-config) below for details.

### Https Auto Cert Tls Config

TLS Configuration Parameters.

###### One of the arguments from this list "tls_12_plus, tls_11_plus" must be set

`tls_11_plus` - (Optional) TLS v1.1+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

`tls_12_plus` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

### Ip Allowed List Choice Bypass Rate Limiting Rules

This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting..

`bypass_rate_limiting_rules` - (Optional) This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting.. See [Bypass Rate Limiting Rules Bypass Rate Limiting Rules ](#bypass-rate-limiting-rules-bypass-rate-limiting-rules) below for details.

### Ip Allowed List Choice Custom Ip Allowed List

IP Allowed list using existing ip_prefix_set objects..

`rate_limiter_allowed_prefixes` - (Required) Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting.. See [ref](#ref) below for details.

### Ip Allowed List Choice Ip Allowed List

List of IP(s) for which rate limiting will be disabled..

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Ip Allowed List Choice No Ip Allowed List

There is no ip allowed list for rate limiting, all clients go through rate limiting..

### Ip Asn Choice Any Ip

Any Source IP.

### Ip Asn Choice Asn List

The predicate evaluates to true if the origin ASN is present in the ASN list..

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer. (`Int`).

### Ip Asn Choice Asn Matcher

The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects..

`asn_sets` - (Required) A list of references to bgp_asn_set objects.. See [ref](#ref) below for details.

### Ip Asn Choice Ip Matcher

The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes in the IP Prefix Sets..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_sets` - (Required) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Ip Asn Choice Ip Prefix List

The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes from the list..

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

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

### Ip Filtering Type Allow List

Allow list of ip prefixes.

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

### Ip Filtering Type Block List

Block list of ip prefixes.

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

### Ip Reputation Choice Disable Ip Reputation

No IP reputation configured this distribution.

### Ip Reputation Choice Enable Ip Reputation

Enable IP reputation for all requests in this distribution.

`ip_threat_categories` - (Required) If the source IP matches on atleast one of the enabled IP threat categories, the request will be denied. (`List of Strings`).

### Ip Reputation Choice Ip Reputation On Cache Miss

Enable IP reputation only on cache miss in this distribution.

`ip_threat_categories` - (Required) If the source IP matches on atleast one of the enabled IP threat categories, the request will be denied. (`List of Strings`).

### Issuer Validation Issuer Disable

x-displayName: "Disable".

### Js Challenge Parameters Choice Default Js Challenge Parameters

Use default parameters.

### Js Challenge Parameters Choice Js Challenge Parameters

Configure JavaScript challenge parameters.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

`js_script_delay` - (Optional) Delay introduced by Javascript, in milliseconds. (`Int`).

### Jwks Configuration Jwks Config

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details..

`cleartext` - (Optional) The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details. (`String`).

### Jwt Backup Key

Backup JWT Key - If specified is also checked in addition to the primary secret key.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Backup Key Blindfold Secret Info Internal ](#backup-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "vault_secret_info, clear_secret_info, wingman_secret_info, blindfold_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Jwt Secret Key

Secret Key for JWT.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Secret Key Blindfold Secret Info Internal ](#secret-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Jwt Validation Action

x-required.

###### One of the arguments from this list "block, report" must be set

`block` - (Optional) Block the request and report the issue (`Bool`).

`report` - (Optional) Allow the request and report the issue (`Bool`).

### Jwt Validation Mandatory Claims

If the claim does not exist JWT token validation will fail..

`claim_names` - (Optional) x-displayName: "Claim Names" (`String`).

### Jwt Validation Reserved Claims

the token validation of these claims should be disabled..

###### One of the arguments from this list "audience_disable, audience" must be set

`audience` - (Optional) x-displayName: "Exact Match". See [Audience Validation Audience ](#audience-validation-audience) below for details.

`audience_disable` - (Optional) x-displayName: "Disable" (`Bool`).

###### One of the arguments from this list "issuer_disable, issuer" must be set

`issuer` - (Optional) x-displayName: "Exact Match" (`String`).

`issuer_disable` - (Optional) x-displayName: "Disable" (`Bool`).

###### One of the arguments from this list "validate_period_disable, validate_period_enable" must be set

`validate_period_disable` - (Optional) x-displayName: "Disable" (`Bool`).

`validate_period_enable` - (Optional) x-displayName: "Enable" (`Bool`).

### Jwt Validation Target

Define endpoints for which JWT token validation will be performed.

###### One of the arguments from this list "api_groups, base_paths, all_endpoint" must be set

`all_endpoint` - (Optional) Validation will be performed for all requests on this LB (`Bool`).

`api_groups` - (Optional) Validation will be performed for the endpoints mentioned in the API Groups. See [Target Api Groups ](#target-api-groups) below for details.

`base_paths` - (Optional) Validation will be performed for selected path prefixes. See [Target Base Paths ](#target-base-paths) below for details.

### Jwt Validation Token Location

Define where in the HTTP request the JWT token will be extracted.

###### One of the arguments from this list "query_param, bearer_token, cookie, header" must be set

`bearer_token` - (Optional) Token is found in Authorization HTTP header with Bearer authentication scheme (`Bool`).

`cookie` - (Optional) Token is found in the cookie (`String`).(Deprecated)

`header` - (Optional) Token is found in the header (`String`).(Deprecated)

`query_param` - (Optional) Token is found in the query string parameter (`String`).(Deprecated)

### Learn From Redirect Traffic Disable Learn From Redirect Traffic

Disable learning API patterns from traffic with redirect response codes 3xx.

### Learn From Redirect Traffic Enable Learn From Redirect Traffic

Enable learning API patterns from traffic with redirect response codes 3xx.

### Loadbalancer Type Http

CDN Distribution serving content over HTTP.

`dns_volterra_managed` - (Optional) or a DNS CNAME record should be created in your DNS provider's portal. (`Bool`).

###### One of the arguments from this list "port, port_ranges" must be set

`port` - (Optional) HTTP port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

### Loadbalancer Type Https

User is responsible for managing DNS..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

`tls_parameters` - (Optional) TLS parameters for the downstream connections.. See [Https Tls Parameters ](#https-tls-parameters) below for details.

### Loadbalancer Type Https Auto Cert

DNS records will be managed by Volterra..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

`tls_config` - (Optional) TLS Configuration Parameters. See [Https Auto Cert Tls Config ](#https-auto-cert-tls-config) below for details.

### Logging Options Client Log Options

Client request headers to log.

`header_list` - (Optional) List of headers (`String`).

### Logging Options Origin Log Options

Origin response headers to log.

`header_list` - (Optional) List of headers (`String`).

### Malicious User Detection Choice Disable Malicious User Detection

Disable malicious user detection for this distribution.

### Malicious User Detection Choice Enable Malicious User Detection

Enable malicious user detection for all requests in this distribution.

### Malicious User Detection Choice Malicious User Detection On Cache Miss

Enable malicious user detection only on cache miss in this distribution.

### Malicious User Mitigation Choice Default Mitigation Settings

For high level, users will be temporarily blocked..

### Match Check Not Present

Check that the cookie is not present..

### Match Check Present

Check that the cookie is present..

### Match Item

Criteria for matching the values for the cookie. The match is successful if any of the values in the input satisfies the criteria in the matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Max Age Ignore Max Age

Ignore max age attribute.

### Method Choice Method Get

x-displayName: "GET".

### Method Choice Method Post

x-displayName: "POST".

### More Option Cache Options

Cache Options.

`cache_rules` - (Optional) Rules are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs.. See [Cache Options Cache Rules ](#cache-options-cache-rules) below for details.

`default_cache_action` - (Required) Default value for Cache action.. See [Cache Options Default Cache Action ](#cache-options-default-cache-action) below for details.

### More Option Cache Ttl Options

Cache Options.

###### One of the arguments from this list "cache_ttl_default, cache_ttl_override, cache_disabled" can be set

`cache_disabled` - (Optional) Disable Caching of content from the origin (`Bool`).

`cache_ttl_default` - (Optional) Cache TTL value to use when the origin does not provide one (`String`).

`cache_ttl_override` - (Optional) Override the Cache TTL directive in the response from the origin (`String`).

### More Option Header Options

Request/Response header related options.

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Header Options Request Headers To Add ](#header-options-request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Header Options Response Headers To Add ](#header-options-response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

### More Option Logging Options

Logging related options.

`client_log_options` - (Optional) Client request headers to log. See [Logging Options Client Log Options ](#logging-options-client-log-options) below for details.

`origin_log_options` - (Optional) Origin response headers to log. See [Logging Options Origin Log Options ](#logging-options-origin-log-options) below for details.

### More Option Security Options

Security related options.

`api_protection` - (Optional) x-displayName: "API Protection". See [Security Options Api Protection ](#security-options-api-protection) below for details.

`auth_options` - (Optional) Authentication Options. See [Security Options Auth Options ](#security-options-auth-options) below for details.

`common_security_controls` - (Optional) x-displayName: "Common Security Controls". See [Security Options Common Security Controls ](#security-options-common-security-controls) below for details.

`geo_filtering` - (Optional) Geo filtering options. See [Security Options Geo Filtering ](#security-options-geo-filtering) below for details.

`ip_filtering` - (Optional) IP filtering options. See [Security Options Ip Filtering ](#security-options-ip-filtering) below for details.

`web_app_firewall` - (Optional) Web Application Firewall. See [Security Options Web App Firewall ](#security-options-web-app-firewall) below for details.

### Mtls Choice No Mtls

x-displayName: "Disable".

### Mtls Choice Use Mtls

x-displayName: "Upload a client authentication certificate specifically for this Origin Pool".

`tls_certificates` - (Required) mTLS Client Certificate. See [Use Mtls Tls Certificates ](#use-mtls-tls-certificates) below for details.

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ocsp Stapling Choice Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Ocsp Stapling Choice Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Open Api Validation Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Open Api Validation Rules Validation Mode

When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger).

###### One of the arguments from this list "response_validation_mode_active, skip_response_validation" must be set

`response_validation_mode_active` - (Optional) Enforce OpenAPI validation processing for this event. See [Response Validation Mode Choice Response Validation Mode Active ](#response-validation-mode-choice-response-validation-mode-active) below for details.

`skip_response_validation` - (Optional) Skip OpenAPI validation processing for this event (`Bool`).

###### One of the arguments from this list "validation_mode_active, skip_validation" must be set

`skip_validation` - (Optional) Skip OpenAPI validation processing for this event (`Bool`).

`validation_mode_active` - (Optional) Enforce OpenAPI validation processing for this event. See [Validation Mode Choice Validation Mode Active ](#validation-mode-choice-validation-mode-active) below for details.

### Origin Pool More Origin Options

x-displayName: "Advanced Configuration".

`disable_byte_range_request` - (Optional) Choice to enable/disable origin byte range requrests towards origin (`Bool`).

`websocket_proxy` - (Optional) Option to enable proxying of websocket connections to the origin server (`Bool`).

### Origin Pool Origin Servers

List of original servers.

###### One of the arguments from this list "public_ip, public_name" must be set

`public_ip` - (Optional) Specify origin server with public IP. See [Choice Public Ip ](#choice-public-ip) below for details.

`public_name` - (Optional) Specify origin server with public DNS name. See [Choice Public Name ](#choice-public-name) below for details.

`port` - (Optional) Port the workload can be reached on (`Int`).

### Origin Pool Public Name

The DNS name to be used as the host header for the request to the origin server.

`dns_name` - (Required) DNS Name (`String`).

`refresh_interval` - (Optional) Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767 (`Int`).

### Oversized Body Choice Oversized Body Fail Validation

Apply the request/response action (block or report) when the body length is too long to verify (default 64Kb).

### Oversized Body Choice Oversized Body Skip Validation

Skip body validation when the body length is too long to verify (default 64Kb).

### Path Choice Any Path

Match all paths.

### Path Match Operator

A specification of path match.

###### One of the arguments from this list "Equals, DoesNotEqual, Endswith, DoesNotEndWith, MatchRegex, Contains, DoesNotContain, Startswith, DoesNotStartWith" can be set

`Contains` - (Optional) Field must contain (`String`).

`DoesNotContain` - (Optional) Field must not contain (`String`).

`DoesNotEndWith` - (Optional) Field must not end with (`String`).

`DoesNotEqual` - (Optional) Field must not equal (`String`).

`DoesNotStartWith` - (Optional) Field must not start with (`String`).

`Endswith` - (Optional) Field must end with (`String`).

`Equals` - (Optional) Field must exactly match (`String`).

`MatchRegex` - (Optional) Field matches regular expression (`String`).

`Startswith` - (Optional) Field must start with (`String`).

### Policy Based Challenge Rule List

list challenge rules to be used in policy based challenge.

`rules` - (Optional) these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions. See [Rule List Rules ](#rule-list-rules) below for details.

### Policy Choice No Policies

Do not apply additional rate limiter policies..

### Policy Choice Policies

to the action configured in the rule. If there's no match, the rate limiting configuration for the HTTP load balancer is honored..

`policies` - (Required) Ordered list of rate limiter policies.. See [ref](#ref) below for details.

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Property Validation Settings Choice Property Validation Settings Custom

Use custom settings with Open API specification validation.

`headers` - (Optional) Custom settings for headers validation. See [Property Validation Settings Custom Headers ](#property-validation-settings-custom-headers) below for details.(Deprecated)

`queryParameters` - (Optional) Custom settings for query parameters validation. See [Property Validation Settings Custom QueryParameters ](#property-validation-settings-custom-queryParameters) below for details.

### Property Validation Settings Choice Property Validation Settings Default

Keep the default settings of OpenAPI specification validation.

### Property Validation Settings Custom Headers

Custom settings for headers validation.

###### One of the arguments from this list "allow_additional_headers, disallow_additional_headers" must be set

`allow_additional_headers` - (Optional) Allow extra headers (on top of what specified in the OAS documentation) (`Bool`).

`disallow_additional_headers` - (Optional) Disallow extra headers (on top of what specified in the OAS documentation) (`Bool`).

### Property Validation Settings Custom QueryParameters

Custom settings for query parameters validation.

###### One of the arguments from this list "allow_additional_parameters, disallow_additional_parameters" must be set

`allow_additional_parameters` - (Optional) Allow extra query parameters (on top of what specified in the OAS documentation) (`Bool`).

`disallow_additional_parameters` - (Optional) Disallow extra query parameters (on top of what specified in the OAS documentation) (`Bool`).

### Query Parameters Operator

.

###### One of the arguments from this list "Equals, DoesNotEqual, Endswith, DoesNotEndWith, MatchRegex, Contains, DoesNotContain, Startswith, DoesNotStartWith" can be set

`Contains` - (Optional) Field must contain (`String`).

`DoesNotContain` - (Optional) Field must not contain (`String`).

`DoesNotEndWith` - (Optional) Field must not end with (`String`).

`DoesNotEqual` - (Optional) Field must not equal (`String`).

`DoesNotStartWith` - (Optional) Field must not start with (`String`).

`Endswith` - (Optional) Field must end with (`String`).

`Equals` - (Optional) Field must exactly match (`String`).

`MatchRegex` - (Optional) Field matches regular expression (`String`).

`Startswith` - (Optional) Field must start with (`String`).

### Rate Limit Rate Limiter

Requests to the virtual_host are rate limited based on the parameters specified in the rate_limiter..

`burst_multiplier` - (Optional) The maximum burst of requests to accommodate, expressed as a multiple of the rate. (`Int`).

`total_number` - (Required) The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period. (`Int`).

`unit` - (Required) Unit for the period per which the rate limit is applied. (`String`).

### Rate Limit Choice Api Rate Limit

Define rate limiting for one or more API endpoints.

`api_endpoint_rules` - (Optional) For creating rule that contain a whole domain or group of endpoints, please use the server URL rules above.. See [Api Rate Limit Api Endpoint Rules ](#api-rate-limit-api-endpoint-rules) below for details.

###### One of the arguments from this list "custom_ip_allowed_list, bypass_rate_limiting_rules, no_ip_allowed_list, ip_allowed_list" must be set

`bypass_rate_limiting_rules` - (Optional) This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting.. See [Ip Allowed List Choice Bypass Rate Limiting Rules ](#ip-allowed-list-choice-bypass-rate-limiting-rules) below for details.

`custom_ip_allowed_list` - (Optional) IP Allowed list using existing ip_prefix_set objects.. See [Ip Allowed List Choice Custom Ip Allowed List ](#ip-allowed-list-choice-custom-ip-allowed-list) below for details.

`ip_allowed_list` - (Optional) List of IP(s) for which rate limiting will be disabled.. See [Ip Allowed List Choice Ip Allowed List ](#ip-allowed-list-choice-ip-allowed-list) below for details.

`no_ip_allowed_list` - (Optional) There is no ip allowed list for rate limiting, all clients go through rate limiting. (`Bool`).

`server_url_rules` - (Optional) For matching also specific endpoints you can use the API endpoint rules set bellow.. See [Api Rate Limit Server Url Rules ](#api-rate-limit-server-url-rules) below for details.

### Rate Limit Choice Disable Rate Limit

Rate limiting is not currently enabled for this load balancer.

### Rate Limit Choice Rate Limit

Define custom rate limiting parameters for this load balancer.

###### One of the arguments from this list "no_ip_allowed_list, ip_allowed_list, custom_ip_allowed_list" must be set

`custom_ip_allowed_list` - (Optional) IP Allowed list using existing ip_prefix_set objects.. See [Ip Allowed List Choice Custom Ip Allowed List ](#ip-allowed-list-choice-custom-ip-allowed-list) below for details.

`ip_allowed_list` - (Optional) List of IP(s) for which rate limiting will be disabled.. See [Ip Allowed List Choice Ip Allowed List ](#ip-allowed-list-choice-ip-allowed-list) below for details.

`no_ip_allowed_list` - (Optional) There is no ip allowed list for rate limiting, all clients go through rate limiting. (`Bool`).

###### One of the arguments from this list "no_policies, policies" must be set

`no_policies` - (Optional) Do not apply additional rate limiter policies. (`Bool`).

`policies` - (Optional) to the action configured in the rule. If there's no match, the rate limiting configuration for the HTTP load balancer is honored.. See [Policy Choice Policies ](#policy-choice-policies) below for details.

`rate_limiter` - (Optional) Requests to the virtual_host are rate limited based on the parameters specified in the rate_limiter.. See [Rate Limit Rate Limiter ](#rate-limit-rate-limiter) below for details.

### Rate Limiter Choice Inline Rate Limiter

Specify rate values for the rule..

###### One of the arguments from this list "use_http_lb_user_id, ref_user_id" must be set

`ref_user_id` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier to be rate limited.. See [ref](#ref) below for details.

`use_http_lb_user_id` - (Optional) Defined in HTTP-LB Security Configuration -> User Identifier. (`Bool`).

`threshold` - (Required) The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period. (`Int`).

`unit` - (Required) Unit for the period per which the rate limit is applied. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Request Matcher Cookie Matchers

Note that all specified cookie matcher predicates must evaluate to true..

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

###### One of the arguments from this list "presence, check_present, check_not_present, item" must be set

`check_not_present` - (Optional) Check that the cookie is not present. (`Bool`).

`check_present` - (Optional) Check that the cookie is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the cookie. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the cookie is present or absent. (`Bool`).(Deprecated)

`name` - (Required) A case-sensitive cookie name. (`String`).

### Request Matcher Headers

Note that all specified header predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

###### One of the arguments from this list "presence, check_present, check_not_present, item" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the header is present or absent. (`Bool`).(Deprecated)

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Request Matcher Jwt Claims

Note that this feature only works on LBs with JWT Validation feature enabled..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

###### One of the arguments from this list "item, check_present, check_not_present" must be set

`check_not_present` - (Optional) Check that the JWT Claim is not present. (`Bool`).

`check_present` - (Optional) Check that the JWT Claim is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the JWT Claim. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) JWT claim name. (`String`).

### Request Matcher Query Params

Note that all specified query parameter predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`key` - (Required) A case-sensitive HTTP query parameter name. (`String`).

###### One of the arguments from this list "item, presence, check_present, check_not_present" must be set

`check_not_present` - (Optional) Check that the query parameter is not present. (`Bool`).

`check_present` - (Optional) Check that the query parameter is present. (`Bool`).

`item` - (Optional) criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the query parameter is present or absent. (`Bool`).(Deprecated)

### Response Validation Mode Choice Response Validation Mode Active

Enforce OpenAPI validation processing for this event.

`response_validation_properties` - (Required) List of properties of the response to validate according to the OpenAPI specification file (a.k.a. swagger) (`List of Strings`).

###### One of the arguments from this list "enforcement_report, enforcement_block" must be set

`enforcement_block` - (Optional) Block the response, trigger an API security event (`Bool`).

`enforcement_report` - (Optional) Allow the response, trigger an API security event (`Bool`).

### Response Validation Mode Choice Skip Response Validation

Skip OpenAPI validation processing for this event.

### Rule Expression List Cache Rule Expression

The Cache Rule Expression Terms that are ANDed.

`cache_headers` - (Optional) Configure cache rule headers to match the criteria. See [Cache Rule Expression Cache Headers ](#cache-rule-expression-cache-headers) below for details.

`cookie_matcher` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Cache Rule Expression Cookie Matcher ](#cache-rule-expression-cookie-matcher) below for details.

`path_match` - (Optional) URI path of route. See [Cache Rule Expression Path Match ](#cache-rule-expression-path-match) below for details.

`query_parameters` - (Optional) List of (key, value) query parameters. See [Cache Rule Expression Query Parameters ](#cache-rule-expression-query-parameters) below for details.

### Rule List Rules

these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Rules Metadata ](#rules-metadata) below for details.

`spec` - (Required) Specification for the rule including match predicates and actions.. See [Rules Spec ](#rules-spec) below for details.

### Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Rules Spec

Specification for the rule including match predicates and actions..

`arg_matchers` - (Optional)arg_matchers. See [Spec Arg Matchers ](#spec-arg-matchers) below for details.

###### One of the arguments from this list "asn_list, asn_matcher, any_asn" can be set

`any_asn` - (Optional)any_asn (`Bool`).

`asn_list` - (Optional)asn_list. See [Asn Choice Asn List ](#asn-choice-asn-list) below for details.

`asn_matcher` - (Optional)asn_matcher. See [Asn Choice Asn Matcher ](#asn-choice-asn-matcher) below for details.

`body_matcher` - (Optional)body_matcher. See [Spec Body Matcher ](#spec-body-matcher) below for details.

###### One of the arguments from this list "disable_challenge, enable_javascript_challenge, enable_captcha_challenge" must be set

`disable_challenge` - (Optional) Disable the challenge type selected in PolicyBasedChallenge (`Bool`).

`enable_captcha_challenge` - (Optional) Enable captcha challenge (`Bool`).

`enable_javascript_challenge` - (Optional) Enable javascript challenge (`Bool`).

###### One of the arguments from this list "client_name, client_selector, client_name_matcher, any_client" can be set

`any_client` - (Optional)any_client (`Bool`).

`client_name` - (Optional)client_name (`String`).(Deprecated)

`client_name_matcher` - (Optional)client_name_matcher. See [Client Choice Client Name Matcher ](#client-choice-client-name-matcher) below for details.(Deprecated)

`client_selector` - (Optional)client_selector. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.

`cookie_matchers` - (Optional)cookie_matchers. See [Spec Cookie Matchers ](#spec-cookie-matchers) below for details.

`domain_matcher` - (Optional)domain_matcher. See [Spec Domain Matcher ](#spec-domain-matcher) below for details.

`expiration_timestamp` - (Optional)expiration_timestamp (`String`).

`headers` - (Optional)headers. See [Spec Headers ](#spec-headers) below for details.

`http_method` - (Optional)http_method. See [Spec Http Method ](#spec-http-method) below for details.

###### One of the arguments from this list "any_ip, ip_prefix_list, ip_matcher" can be set

`any_ip` - (Optional)any_ip (`Bool`).

`ip_matcher` - (Optional)ip_matcher. See [Ip Choice Ip Matcher ](#ip-choice-ip-matcher) below for details.

`ip_prefix_list` - (Optional)ip_prefix_list. See [Ip Choice Ip Prefix List ](#ip-choice-ip-prefix-list) below for details.

`path` - (Optional)path. See [Spec Path ](#spec-path) below for details.

`query_params` - (Optional)query_params. See [Spec Query Params ](#spec-query-params) below for details.

`tls_fingerprint_matcher` - (Optional)tls_fingerprint_matcher. See [Spec Tls Fingerprint Matcher ](#spec-tls-fingerprint-matcher) below for details.

### Samesite Ignore Samesite

Ignore Samesite attribute.

### Samesite Samesite Lax

Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests.

### Samesite Samesite None

Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests.

### Samesite Samesite Strict

Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests.

### Secret Info Oneof Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Info Oneof Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Secret Info Oneof Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Secret Info Oneof Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

### Secret Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Value Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secure Add Secure

x-displayName: "Add".

### Secure Ignore Secure

x-displayName: "Ignore".

### Security Options Api Protection

x-displayName: "API Protection".

###### One of the arguments from this list "disable_api_definition, api_specification, api_specification_on_cache_miss" must be set

`api_specification` - (Optional) Specify API definition and OpenAPI Validation. See [Api Definition Choice Api Specification ](#api-definition-choice-api-specification) below for details.

`api_specification_on_cache_miss` - (Optional) Enable API definition and OpenAPI Validation only on cache miss in this distribution. See [Api Definition Choice Api Specification On Cache Miss ](#api-definition-choice-api-specification-on-cache-miss) below for details.

`disable_api_definition` - (Optional) API Definition is not currently used for this load balancer (`Bool`).

###### One of the arguments from this list "enable_api_discovery, disable_api_discovery, api_discovery_on_cache_miss" must be set

`api_discovery_on_cache_miss` - (Optional) Enable api discovery only on cache miss in this distribution. See [Api Discovery Choice Api Discovery On Cache Miss ](#api-discovery-choice-api-discovery-on-cache-miss) below for details.

`disable_api_discovery` - (Optional) Disable api discovery for this distribution (`Bool`).

`enable_api_discovery` - (Optional) Enable api discovery for all requests in this distribution. See [Api Discovery Choice Enable Api Discovery ](#api-discovery-choice-enable-api-discovery) below for details.

`api_protection_rules` - (Optional) Rules can also include additional conditions, for example specific clients can access certain API endpoint or API group.. See [Api Protection Api Protection Rules ](#api-protection-api-protection-rules) below for details.

`jwt_validation` - (Optional) tokens or tokens that are not yet valid.. See [Api Protection Jwt Validation ](#api-protection-jwt-validation) below for details.

### Security Options Auth Options

Authentication Options.

###### One of the arguments from this list "disable_auth, jwt, custom" can be set

`custom` - (Optional) Enable Custom Authentication. See [Auth Options Custom ](#auth-options-custom) below for details.

`disable_auth` - (Optional) No Authentication (`Bool`).

`jwt` - (Optional) Enable JWT Authentication. See [Auth Options Jwt ](#auth-options-jwt) below for details.

### Security Options Common Security Controls

x-displayName: "Common Security Controls".

`blocked_clients` - (Optional) Define rules to block IP Prefixes or AS numbers.. See [Common Security Controls Blocked Clients ](#common-security-controls-blocked-clients) below for details.

###### One of the arguments from this list "policy_based_challenge, no_challenge, enable_challenge, challenge_on_cache_miss, js_challenge, captcha_challenge" must be set

`captcha_challenge` - (Optional) Configure Captcha challenge on this load balancer. See [Challenge Type Captcha Challenge ](#challenge-type-captcha-challenge) below for details.

`challenge_on_cache_miss` - (Optional) Configure auto mitigation i.e risk based challenges for malicious users only on cache miss in this load balancer. See [Challenge Type Challenge On Cache Miss ](#challenge-type-challenge-on-cache-miss) below for details.

`enable_challenge` - (Optional) Configure auto mitigation i.e risk based challenges for malicious users for this load balancer. See [Challenge Type Enable Challenge ](#challenge-type-enable-challenge) below for details.

`js_challenge` - (Optional) Configure JavaScript challenge on this load balancer. See [Challenge Type Js Challenge ](#challenge-type-js-challenge) below for details.

`no_challenge` - (Optional) No challenge is enabled for this load balancer (`Bool`).

`policy_based_challenge` - (Optional) Specifies the settings for policy rule based challenge. See [Challenge Type Policy Based Challenge ](#challenge-type-policy-based-challenge) below for details.

`cors_policy` - (Optional) resources from a server at a different origin. See [Common Security Controls Cors Policy ](#common-security-controls-cors-policy) below for details.

###### One of the arguments from this list "disable_ip_reputation, enable_ip_reputation, ip_reputation_on_cache_miss" can be set

`disable_ip_reputation` - (Optional) No IP reputation configured this distribution (`Bool`).

`enable_ip_reputation` - (Optional) Enable IP reputation for all requests in this distribution. See [Ip Reputation Choice Enable Ip Reputation ](#ip-reputation-choice-enable-ip-reputation) below for details.

`ip_reputation_on_cache_miss` - (Optional) Enable IP reputation only on cache miss in this distribution. See [Ip Reputation Choice Ip Reputation On Cache Miss ](#ip-reputation-choice-ip-reputation-on-cache-miss) below for details.

###### One of the arguments from this list "disable_malicious_user_detection, enable_malicious_user_detection, malicious_user_detection_on_cache_miss" must be set

`disable_malicious_user_detection` - (Optional) Disable malicious user detection for this distribution (`Bool`).

`enable_malicious_user_detection` - (Optional) Enable malicious user detection for all requests in this distribution (`Bool`).

`malicious_user_detection_on_cache_miss` - (Optional) Enable malicious user detection only on cache miss in this distribution (`Bool`).

###### One of the arguments from this list "api_rate_limit, rate_limit, disable_rate_limit" must be set

`api_rate_limit` - (Optional) Define rate limiting for one or more API endpoints. See [Rate Limit Choice Api Rate Limit ](#rate-limit-choice-api-rate-limit) below for details.

`disable_rate_limit` - (Optional) Rate limiting is not currently enabled for this load balancer (`Bool`).

`rate_limit` - (Optional) Define custom rate limiting parameters for this load balancer. See [Rate Limit Choice Rate Limit ](#rate-limit-choice-rate-limit) below for details.

###### One of the arguments from this list "active_service_policies, service_policies_from_namespace, no_service_policies" must be set

`active_service_policies` - (Optional) Apply the specified list of service policies and bypass the namespace service policy set. See [Service Policy Choice Active Service Policies ](#service-policy-choice-active-service-policies) below for details.

`no_service_policies` - (Optional) Do not apply any service policies i.e. bypass the namespace service policy set (`Bool`).

`service_policies_from_namespace` - (Optional) Apply the active service policies configured as part of the namespace service policy set (`Bool`).

`trusted_clients` - (Optional) Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients.. See [Common Security Controls Trusted Clients ](#common-security-controls-trusted-clients) below for details.

###### One of the arguments from this list "user_id_client_ip, user_identification" must be set

`user_id_client_ip` - (Optional) Use the Client IP address as the user identifier. (`Bool`).

`user_identification` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier.. See [ref](#ref) below for details.

### Security Options Geo Filtering

Geo filtering options.

###### One of the arguments from this list "allow_list, block_list" can be set

`allow_list` - (Optional) Allow list of countries. See [Geo Filtering Type Allow List ](#geo-filtering-type-allow-list) below for details.

`block_list` - (Optional) Block list of countries. See [Geo Filtering Type Block List ](#geo-filtering-type-block-list) below for details.

### Security Options Ip Filtering

IP filtering options.

###### One of the arguments from this list "allow_list, block_list" can be set

`allow_list` - (Optional) Allow list of ip prefixes. See [Ip Filtering Type Allow List ](#ip-filtering-type-allow-list) below for details.

`block_list` - (Optional) Block list of ip prefixes. See [Ip Filtering Type Block List ](#ip-filtering-type-block-list) below for details.

### Security Options Web App Firewall

Web Application Firewall.

`csrf_policy` - (Optional) Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.).. See [Web App Firewall Csrf Policy ](#web-app-firewall-csrf-policy) below for details.

`data_guard_rules` - (Optional) Note: App Firewall should be enabled, to use Data Guard feature.. See [Web App Firewall Data Guard Rules ](#web-app-firewall-data-guard-rules) below for details.

`graphql_rules` - (Optional) queries and prevent GraphQL tailored attacks.. See [Web App Firewall Graphql Rules ](#web-app-firewall-graphql-rules) below for details.

`protected_cookies` - (Optional) Note: We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.. See [Web App Firewall Protected Cookies ](#web-app-firewall-protected-cookies) below for details.

###### One of the arguments from this list "disable_waf, app_firewall, app_firewall_on_cache_miss" must be set

`app_firewall` - (Optional) Enable WAF configuration for all requests in this distribution. See [ref](#ref) below for details.

`app_firewall_on_cache_miss` - (Optional) Enable WAF configuration only on cache miss in this distribution. See [ref](#ref) below for details.

`disable_waf` - (Optional) No WAF configuration for this load balancer (`Bool`).

`waf_exclusion_rules` - (Optional) When an exclusion rule is matched, then this exclusion rule takes effect and no more rules are evaluated.. See [Web App Firewall Waf Exclusion Rules ](#web-app-firewall-waf-exclusion-rules) below for details.

### Server Url Rules Client Matcher

Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc..

###### One of the arguments from this list "ip_threat_category_list, client_selector, any_client" must be set

`any_client` - (Optional) Any Client (`Bool`).

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.

`ip_threat_category_list` - (Optional) IP threat categories to choose from. See [Client Choice Ip Threat Category List ](#client-choice-ip-threat-category-list) below for details.

###### One of the arguments from this list "asn_matcher, any_ip, ip_prefix_list, ip_matcher, asn_list" must be set

`any_ip` - (Optional) Any Source IP (`Bool`).

`asn_list` - (Optional) The predicate evaluates to true if the origin ASN is present in the ASN list.. See [Ip Asn Choice Asn List ](#ip-asn-choice-asn-list) below for details.

`asn_matcher` - (Optional) The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects.. See [Ip Asn Choice Asn Matcher ](#ip-asn-choice-asn-matcher) below for details.

`ip_matcher` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes in the IP Prefix Sets.. See [Ip Asn Choice Ip Matcher ](#ip-asn-choice-ip-matcher) below for details.

`ip_prefix_list` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes from the list.. See [Ip Asn Choice Ip Prefix List ](#ip-asn-choice-ip-prefix-list) below for details.

`tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints.. See [Client Matcher Tls Fingerprint Matcher ](#client-matcher-tls-fingerprint-matcher) below for details.

### Server Url Rules Request Matcher

Conditions related to the request, such as query parameters, headers, etc..

`cookie_matchers` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Request Matcher Cookie Matchers ](#request-matcher-cookie-matchers) below for details.

`headers` - (Optional) Note that all specified header predicates must evaluate to true.. See [Request Matcher Headers ](#request-matcher-headers) below for details.

`jwt_claims` - (Optional) Note that this feature only works on LBs with JWT Validation feature enabled.. See [Request Matcher Jwt Claims ](#request-matcher-jwt-claims) below for details.

`query_params` - (Optional) Note that all specified query parameter predicates must evaluate to true.. See [Request Matcher Query Params ](#request-matcher-query-params) below for details.

### Server Validation Choice Skip Server Verification

Skip origin server verification.

### Server Validation Choice Use Server Verification

Perform origin server verification using the provided Root CA Certificate.

###### One of the arguments from this list "trusted_ca_url, trusted_ca" must be set

`trusted_ca` - (Optional) Select/Add a Root CA Certificate object to associate with this Origin Pool for verification of server's certificate. See [ref](#ref) below for details.

`trusted_ca_url` - (Optional) Upload a Root CA Certificate specifically for this Origin Pool for verification of server's certificate (`String`).

### Server Validation Choice Volterra Trusted Ca

Perform origin server verification using F5XC Default Root CA Certificate.

### Service Policy Choice Active Service Policies

Apply the specified list of service policies and bypass the namespace service policy set.

`policies` - (Required) If all policies are evaluated and none match, then the request will be denied by default.. See [ref](#ref) below for details.

### Service Policy Choice No Service Policies

Do not apply any service policies i.e. bypass the namespace service policy set.

### Service Policy Choice Service Policies From Namespace

Apply the active service policies configured as part of the namespace service policy set.

### Sni Choice Disable Sni

Do not use SNI..

### Sni Choice Use Host Header As Sni

Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied..

### Spec Arg Matchers

arg_matchers.

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

###### One of the arguments from this list "item, presence, check_present, check_not_present" must be set

`check_not_present` - (Optional) Check that the argument is not present. (`Bool`).

`check_present` - (Optional) Check that the argument is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the Arg. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the arg is present or absent. (`Bool`).(Deprecated)

`name` - (Required) A case-sensitive JSON path in the HTTP request body. (`String`).

### Spec Body Matcher

body_matcher.

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Spec Cookie Matchers

cookie_matchers.

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

###### One of the arguments from this list "presence, check_present, check_not_present, item" must be set

`check_not_present` - (Optional) Check that the cookie is not present. (`Bool`).

`check_present` - (Optional) Check that the cookie is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the cookie. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the cookie is present or absent. (`Bool`).(Deprecated)

`name` - (Required) A case-sensitive cookie name. (`String`).

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

### Spec Query Params

query_params.

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`key` - (Required) A case-sensitive HTTP query parameter name. (`String`).

###### One of the arguments from this list "presence, check_present, check_not_present, item" must be set

`check_not_present` - (Optional) Check that the query parameter is not present. (`Bool`).

`check_present` - (Optional) Check that the query parameter is present. (`Bool`).

`item` - (Optional) criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the query parameter is present or absent. (`Bool`).(Deprecated)

### Spec Tls Fingerprint Matcher

tls_fingerprint_matcher.

`classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`exact_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

`excluded_values` - (Optional) or more known TLS fingerprint classes in the enclosing matcher. (`String`).

### Target All Endpoint

Validation will be performed for all requests on this LB.

### Target Api Groups

Validation will be performed for the endpoints mentioned in the API Groups.

`api_groups` - (Required) x-required (`String`).

### Target Base Paths

Validation will be performed for selected path prefixes.

`base_paths` - (Required) x-required (`String`).

### Temporary Blocking Parameters Choice Default Temporary Blocking Parameters

Use default parameters.

### Temporary Blocking Parameters Choice Temporary User Blocking

Specifies configuration for temporary user blocking resulting from malicious user detection.

`custom_page` - (Optional) E.g. "<p> Blocked </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

### Tls Certificates Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "wingman_secret_info, blindfold_secret_info, vault_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Tls Choice No Tls

Origin servers do not use TLS.

### Tls Choice Use Tls

Origin servers use TLS.

###### One of the arguments from this list "no_mtls, use_mtls, use_mtls_obj" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Upload a client authentication certificate specifically for this Origin Pool". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`use_mtls_obj` - (Optional) x-displayName: "Select/add a TLS Certificate object for client authentication". See [ref](#ref) below for details.

###### One of the arguments from this list "use_server_verification, skip_server_verification, volterra_trusted_ca" must be set

`skip_server_verification` - (Optional) Skip origin server verification (`Bool`).

`use_server_verification` - (Optional) Perform origin server verification using the provided Root CA Certificate. See [Server Validation Choice Use Server Verification ](#server-validation-choice-use-server-verification) below for details.

`volterra_trusted_ca` - (Optional) Perform origin server verification using F5XC Default Root CA Certificate (`Bool`).

###### One of the arguments from this list "disable_sni, sni, use_host_header_as_sni" must be set

`disable_sni` - (Optional) Do not use SNI. (`Bool`).

`sni` - (Optional) SNI value to be used. (`String`).

`use_host_header_as_sni` - (Optional) Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied. (`Bool`).

`tls_config` - (Required) TLS parameters such as min/max TLS version and ciphers. See [Use Tls Tls Config ](#use-tls-tls-config) below for details.

### Tls Parameters Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "use_system_defaults, disable_ocsp_stapling, custom_hash_algorithms" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Tls Parameters Tls Config

TLS Configuration Parameters.

###### One of the arguments from this list "tls_11_plus, tls_12_plus" must be set

`tls_11_plus` - (Optional) TLS v1.1+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

`tls_12_plus` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

### Token Location Bearer Token

Token is found in Authorization HTTP header with Bearer authentication scheme.

### Token Source Bearer Token

Token is found in the Bearer-Token.

### Token Source Cookie

Token is found in the cookie.

`name` - (Required) A case-insensitive cookie name. (`String`).

### Token Source Header

Token is found in the header.

`name` - (Required) A case-insensitive field header name. (`String`).

### Token Source Query Param

Token is found in the Query-Param.

`key` - (Required) A case-sensitive HTTP query parameter name. (`String`).

### Trusted Clients Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Ttl Options Cache Disabled

Disable Caching of content from the origin.

### Use Mtls Tls Certificates

mTLS Client Certificate.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "use_system_defaults, disable_ocsp_stapling, custom_hash_algorithms" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Use Tls Tls Config

TLS parameters such as min/max TLS version and ciphers.

###### One of the arguments from this list "custom_security, default_security, medium_security, low_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### User Id Choice User Id Client Ip

Use the Client IP address as the user identifier..

### Validate Period Validate Period Disable

x-displayName: "Disable".

### Validate Period Validate Period Enable

x-displayName: "Enable".

### Validation All Spec Endpoints Fall Through Mode

Determine what to do with unprotected endpoints (not part of the API Inventory or doesn't have a specific rule in custom rules).

###### One of the arguments from this list "fall_through_mode_allow, fall_through_mode_custom" must be set

`fall_through_mode_allow` - (Optional) Allow any unprotected end point (`Bool`).

`fall_through_mode_custom` - (Optional) Custom rules for any unprotected end point. See [Fall Through Mode Choice Fall Through Mode Custom ](#fall-through-mode-choice-fall-through-mode-custom) below for details.

### Validation All Spec Endpoints Settings

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom list" enforcement.

###### One of the arguments from this list "oversized_body_skip_validation, oversized_body_fail_validation" can be set

`oversized_body_fail_validation` - (Optional) Apply the request/response action (block or report) when the body length is too long to verify (default 64Kb) (`Bool`).

`oversized_body_skip_validation` - (Optional) Skip body validation when the body length is too long to verify (default 64Kb) (`Bool`).

###### One of the arguments from this list "property_validation_settings_default, property_validation_settings_custom" can be set

`property_validation_settings_custom` - (Optional) Use custom settings with Open API specification validation. See [Property Validation Settings Choice Property Validation Settings Custom ](#property-validation-settings-choice-property-validation-settings-custom) below for details.

`property_validation_settings_default` - (Optional) Keep the default settings of OpenAPI specification validation (`Bool`).

### Validation All Spec Endpoints Validation Mode

When a validation mismatch occurs on a request to one of the API Inventory endpoints.

###### One of the arguments from this list "skip_response_validation, response_validation_mode_active" must be set

`response_validation_mode_active` - (Optional) Enforce OpenAPI validation processing for this event. See [Response Validation Mode Choice Response Validation Mode Active ](#response-validation-mode-choice-response-validation-mode-active) below for details.

`skip_response_validation` - (Optional) Skip OpenAPI validation processing for this event (`Bool`).

###### One of the arguments from this list "validation_mode_active, skip_validation" must be set

`skip_validation` - (Optional) Skip OpenAPI validation processing for this event (`Bool`).

`validation_mode_active` - (Optional) Enforce OpenAPI validation processing for this event. See [Validation Mode Choice Validation Mode Active ](#validation-mode-choice-validation-mode-active) below for details.

### Validation Custom List Fall Through Mode

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules).

###### One of the arguments from this list "fall_through_mode_allow, fall_through_mode_custom" must be set

`fall_through_mode_allow` - (Optional) Allow any unprotected end point (`Bool`).

`fall_through_mode_custom` - (Optional) Custom rules for any unprotected end point. See [Fall Through Mode Choice Fall Through Mode Custom ](#fall-through-mode-choice-fall-through-mode-custom) below for details.

### Validation Custom List Open Api Validation Rules

x-displayName: "Validation List".

###### One of the arguments from this list "api_endpoint, base_path, api_group" must be set

`api_endpoint` - (Optional) The API endpoint (Path + Method) which this validation applies to. See [Condition Type Choice Api Endpoint ](#condition-type-choice-api-endpoint) below for details.

`api_group` - (Optional) The API group which this validation applies to (`String`).

`base_path` - (Optional) The base path which this validation applies to (`String`).

###### One of the arguments from this list "any_domain, specific_domain" must be set

`any_domain` - (Optional) The rule will apply for all domains. (`Bool`).

`specific_domain` - (Optional) The rule will apply for a specific domain. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Open Api Validation Rules Metadata ](#open-api-validation-rules-metadata) below for details.

`validation_mode` - (Required) When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger). See [Open Api Validation Rules Validation Mode ](#open-api-validation-rules-validation-mode) below for details.

### Validation Custom List Settings

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom list" enforcement.

###### One of the arguments from this list "oversized_body_fail_validation, oversized_body_skip_validation" can be set

`oversized_body_fail_validation` - (Optional) Apply the request/response action (block or report) when the body length is too long to verify (default 64Kb) (`Bool`).

`oversized_body_skip_validation` - (Optional) Skip body validation when the body length is too long to verify (default 64Kb) (`Bool`).

###### One of the arguments from this list "property_validation_settings_default, property_validation_settings_custom" can be set

`property_validation_settings_custom` - (Optional) Use custom settings with Open API specification validation. See [Property Validation Settings Choice Property Validation Settings Custom ](#property-validation-settings-choice-property-validation-settings-custom) below for details.

`property_validation_settings_default` - (Optional) Keep the default settings of OpenAPI specification validation (`Bool`).

### Validation Enforcement Type Enforcement Block

Block the response, trigger an API security event.

### Validation Enforcement Type Enforcement Report

Allow the response, trigger an API security event.

### Validation Mode Choice Skip Validation

Skip OpenAPI validation processing for this event.

### Validation Mode Choice Validation Mode Active

Enforce OpenAPI validation processing for this event.

`request_validation_properties` - (Required) List of properties of the request to validate according to the OpenAPI specification file (a.k.a. swagger) (`List of Strings`).

###### One of the arguments from this list "enforcement_report, enforcement_block" must be set

`enforcement_block` - (Optional) Block the request, trigger an API security event (`Bool`).

`enforcement_report` - (Optional) Allow the request, trigger an API security event (`Bool`).

### Validation Target Choice Validation All Spec Endpoints

All other API endpoints would proceed according to "Fall Through Mode".

`fall_through_mode` - (Required) Determine what to do with unprotected endpoints (not part of the API Inventory or doesn't have a specific rule in custom rules). See [Validation All Spec Endpoints Fall Through Mode ](#validation-all-spec-endpoints-fall-through-mode) below for details.

###### One of the arguments from this list "oversized_body_skip_validation, oversized_body_fail_validation" can be set

`oversized_body_fail_validation` - (Optional) Apply the request/response action (block or report) when the body length is too long to verify (default 64Kb) (`Bool`).(Deprecated)

`oversized_body_skip_validation` - (Optional) Skip body validation when the body length is too long to verify (default 64Kb) (`Bool`).(Deprecated)

`settings` - (Optional) OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom list" enforcement. See [Validation All Spec Endpoints Settings ](#validation-all-spec-endpoints-settings) below for details.

`validation_mode` - (Required) When a validation mismatch occurs on a request to one of the API Inventory endpoints. See [Validation All Spec Endpoints Validation Mode ](#validation-all-spec-endpoints-validation-mode) below for details.

### Validation Target Choice Validation Custom List

Any other end-points not listed will act according to "Fall Through Mode".

`fall_through_mode` - (Required) Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules). See [Validation Custom List Fall Through Mode ](#validation-custom-list-fall-through-mode) below for details.

`open_api_validation_rules` - (Required) x-displayName: "Validation List". See [Validation Custom List Open Api Validation Rules ](#validation-custom-list-open-api-validation-rules) below for details.

###### One of the arguments from this list "oversized_body_skip_validation, oversized_body_fail_validation" can be set

`oversized_body_fail_validation` - (Optional) Apply the request/response action (block or report) when the body length is too long to verify (default 64Kb) (`Bool`).(Deprecated)

`oversized_body_skip_validation` - (Optional) Skip body validation when the body length is too long to verify (default 64Kb) (`Bool`).(Deprecated)

`settings` - (Optional) OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom list" enforcement. See [Validation Custom List Settings ](#validation-custom-list-settings) below for details.

### Validation Target Choice Validation Disabled

Don't run OpenAPI validation.

### Value Choice Secret Value

Secret Value of the HTTP header..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Secret Value Blindfold Secret Info Internal ](#secret-value-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "clear_secret_info, wingman_secret_info, blindfold_secret_info, vault_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Waf Advanced Configuration App Firewall Detection Control

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria..

`exclude_attack_type_contexts` - (Optional) Attack Types to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Attack Type Contexts ](#app-firewall-detection-control-exclude-attack-type-contexts) below for details.

`exclude_bot_name_contexts` - (Optional) Bot Names to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Bot Name Contexts ](#app-firewall-detection-control-exclude-bot-name-contexts) below for details.

`exclude_signature_contexts` - (Optional) Signature IDs to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Signature Contexts ](#app-firewall-detection-control-exclude-signature-contexts) below for details.

`exclude_violation_contexts` - (Optional) Violations to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Violation Contexts ](#app-firewall-detection-control-exclude-violation-contexts) below for details.

### Waf Advanced Configuration Waf Skip Processing

Skip all App Firewall processing for this request.

### Waf Choice Disable Waf

No WAF configuration for this load balancer.

### Waf Exclusion Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Web App Firewall Csrf Policy

Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)..

###### One of the arguments from this list "all_load_balancer_domains, custom_domain_list, disabled" must be set

`all_load_balancer_domains` - (Optional) Add All load balancer domains to source origin (allow) list. (`Bool`).

`custom_domain_list` - (Optional) Add one or more domains to source origin (allow) list.. See [Allowed Domains Custom Domain List ](#allowed-domains-custom-domain-list) below for details.

`disabled` - (Optional) Allow all source origin domains. (`Bool`).

### Web App Firewall Data Guard Rules

Note: App Firewall should be enabled, to use Data Guard feature..

###### One of the arguments from this list "apply_data_guard, skip_data_guard" must be set

`apply_data_guard` - (Optional) x-displayName: "Apply" (`Bool`).

`skip_data_guard` - (Optional) x-displayName: "Skip" (`Bool`).

###### One of the arguments from this list "any_domain, exact_value, suffix_value" must be set

`any_domain` - (Optional) Enable Data Guard for any domain (`Bool`).

`exact_value` - (Optional) Exact domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Data Guard Rules Metadata ](#data-guard-rules-metadata) below for details.

`path` - (Required) URI path matcher.. See [Data Guard Rules Path ](#data-guard-rules-path) below for details.

### Web App Firewall Graphql Rules

queries and prevent GraphQL tailored attacks..

###### One of the arguments from this list "any_domain, exact_value, suffix_value" must be set

`any_domain` - (Optional) Enable GraphQL inspection for any domain (`Bool`).

`exact_value` - (Optional) Exact domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

`exact_path` - (Required) Specifies the exact path to GraphQL endpoint. Default value is /graphql. (`String`).

`graphql_settings` - (Optional) GraphQL configuration.. See [Graphql Rules Graphql Settings ](#graphql-rules-graphql-settings) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Graphql Rules Metadata ](#graphql-rules-metadata) below for details.

###### One of the arguments from this list "method_get, method_post" must be set

`method_get` - (Optional) x-displayName: "GET" (`Bool`).

`method_post` - (Optional) x-displayName: "POST" (`Bool`).

### Web App Firewall Protected Cookies

Note: We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection..

###### One of the arguments from this list "disable_tampering_protection, enable_tampering_protection" must be set

`disable_tampering_protection` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_tampering_protection` - (Optional) x-displayName: "Enable" (`Bool`).

###### One of the arguments from this list "add_httponly, ignore_httponly" can be set

`add_httponly` - (Optional) x-displayName: "Add" (`Bool`).

`ignore_httponly` - (Optional) x-displayName: "Ignore" (`Bool`).

###### One of the arguments from this list "ignore_max_age, max_age_value" can be set

`ignore_max_age` - (Optional) Ignore max age attribute (`Bool`).(Deprecated)

`max_age_value` - (Optional) Add max age attribute (`Int`).(Deprecated)

`name` - (Required) Name of the Cookie (`String`).

###### One of the arguments from this list "ignore_samesite, samesite_strict, samesite_lax, samesite_none" can be set

`ignore_samesite` - (Optional) Ignore Samesite attribute (`Bool`).

`samesite_lax` - (Optional) Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests (`Bool`).

`samesite_none` - (Optional) Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests (`Bool`).

`samesite_strict` - (Optional) Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests (`Bool`).

###### One of the arguments from this list "ignore_secure, add_secure" can be set

`add_secure` - (Optional) x-displayName: "Add" (`Bool`).

`ignore_secure` - (Optional) x-displayName: "Ignore" (`Bool`).

### Web App Firewall Waf Exclusion Rules

When an exclusion rule is matched, then this exclusion rule takes effect and no more rules are evaluated..

###### One of the arguments from this list "any_domain, exact_value, suffix_value" must be set

`any_domain` - (Optional) Apply this WAF exclusion rule for any domain (`Bool`).

`exact_value` - (Optional) Exact domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Waf Exclusion Rules Metadata ](#waf-exclusion-rules-metadata) below for details.

`methods` - (Optional) methods to be matched (`List of Strings`).

###### One of the arguments from this list "any_path, path_prefix, path_regex" must be set

`any_path` - (Optional) Match all paths (`Bool`).

`path_prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`path_regex` - (Optional) Define the regex for the path. For example, the regex ^/.*$ will match on all paths (`String`).

###### One of the arguments from this list "app_firewall_detection_control, waf_skip_processing" can be set

`app_firewall_detection_control` - (Optional) Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.. See [Waf Advanced Configuration App Firewall Detection Control ](#waf-advanced-configuration-app-firewall-detection-control) below for details.

`waf_skip_processing` - (Optional) Skip all App Firewall processing for this request (`Bool`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured cdn_loadbalancer.
