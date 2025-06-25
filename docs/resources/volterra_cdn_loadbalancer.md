---

page_title: "Volterra: cdn_loadbalancer"

description: "The cdn_loadbalancer allows CRUD of Cdn Loadbalancer resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_cdn_loadbalancer
==================================

The Cdn Loadbalancer allows CRUD of Cdn Loadbalancer resource on Volterra SaaS

~> **Note:** Please refer to [Cdn Loadbalancer API docs](https://docs.cloud.f5.com/docs-v2/api/views-cdn-loadbalancer) to learn more

Example Usage
-------------

```hcl
resource "volterra_cdn_loadbalancer" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "captcha_challenge enable_challenge js_challenge no_challenge policy_based_challenge" must be set

  no_challenge = true
  domains = ["www.foo.com"]

  // One of the arguments from this list "l7_ddos_action_block l7_ddos_action_default l7_ddos_action_js_challenge" must be set

  l7_ddos_action_default = true

  // One of the arguments from this list "http https https_auto_cert" must be set

  https {
    add_hsts = true

    http_redirect = true

    tls_cert_options {
      // One of the arguments from this list "tls_cert_params tls_inline_params" must be set

      tls_inline_params {
        // One of the arguments from this list "no_mtls use_mtls" must be set

        no_mtls = true

        tls_certificates {
          certificate_url = "value"

          description = "Certificate used in production environment"

          // One of the arguments from this list "custom_hash_algorithms disable_ocsp_stapling use_system_defaults" can be set

          use_system_defaults {}
          private_key {
            // One of the arguments from this list "blindfold_secret_info clear_secret_info" must be set

            blindfold_secret_info {
              decryption_provider = "value"

              location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

              store_provider = "value"
            }
          }
        }

        tls_config {
          // One of the arguments from this list "custom_security default_security low_security medium_security" must be set

          default_security = true
        }
      }
    }
  }

  // One of the arguments from this list "disable_malicious_user_detection enable_malicious_user_detection" must be set

  disable_malicious_user_detection = true
  origin_pool {
    more_origin_options {
      disable_byte_range_request = true

      websocket_proxy = true
    }

    origin_request_timeout = "100s"

    origin_servers {
      // One of the arguments from this list "public_ip public_name" must be set

      public_ip {
        // One of the arguments from this list "ip ipv6" must be set

        ip = "8.8.8.8"
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

  // One of the arguments from this list "api_rate_limit disable_rate_limit rate_limit" must be set

  rate_limit {
    // One of the arguments from this list "custom_ip_allowed_list ip_allowed_list no_ip_allowed_list" must be set

    no_ip_allowed_list = true

    // One of the arguments from this list "no_policies policies" must be set

    policies {
      policies {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }
    }
    rate_limiter {
      // One of the arguments from this list "action_block disabled" can be set

      disabled = true

      burst_multiplier = "1"

      period_multiplier = "1"

      total_number = "1"

      unit = "unit"
    }
  }

  // One of the arguments from this list "active_service_policies no_service_policies service_policies_from_namespace" must be set

  service_policies_from_namespace = true

  // One of the arguments from this list "slow_ddos_mitigation system_default_timeouts" must be set

  system_default_timeouts = true

  // One of the arguments from this list "user_id_client_ip user_identification" must be set

  user_id_client_ip = true

  // One of the arguments from this list "app_firewall disable_waf" must be set

  disable_waf = true
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

`blocked_clients` - (Optional) Define rules to block IP Prefixes or AS numbers.. See [Blocked Clients ](#blocked-clients) below for details.

###### One of the arguments from this list "captcha_challenge, enable_challenge, js_challenge, no_challenge, policy_based_challenge" must be set

`captcha_challenge` - (Optional) Configure Captcha challenge on this load balancer. See [Challenge Type Captcha Challenge ](#challenge-type-captcha-challenge) below for details.

`enable_challenge` - (Optional) Configure auto mitigation i.e risk based challenges for malicious users for this load balancer. See [Challenge Type Enable Challenge ](#challenge-type-enable-challenge) below for details.

`js_challenge` - (Optional) Configure JavaScript challenge on this load balancer. See [Challenge Type Js Challenge ](#challenge-type-js-challenge) below for details.

`no_challenge` - (Optional) No challenge is enabled for this load balancer (`Bool`).

`policy_based_challenge` - (Optional) Specifies the settings for policy rule based challenge. See [Challenge Type Policy Based Challenge ](#challenge-type-policy-based-challenge) below for details.

`cors_policy` - (Optional) resources from a server at a different origin. See [Cors Policy ](#cors-policy) below for details.

`csrf_policy` - (Optional) Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.).. See [Csrf Policy ](#csrf-policy) below for details.

`custom_cache_rule` - (Optional) Caching policies for CDN. . See [Custom Cache Rule ](#custom-cache-rule) below for details.

`data_guard_rules` - (Optional) Note: App Firewall should be enabled, to use Data Guard feature.. See [Data Guard Rules ](#data-guard-rules) below for details.

`ddos_mitigation_rules` - (Optional) Define manual mitigation rules to block L7 DDoS attacks.. See [Ddos Mitigation Rules ](#ddos-mitigation-rules) below for details.

`default_cache_action` - (Optional) Default value for Cache action.. See [Default Cache Action ](#default-cache-action) below for details.

`domains` - (Required) [This can be a domain or a sub-domain](`List of String`).

`graphql_rules` - (Optional) queries and prevent GraphQL tailored attacks.. See [Graphql Rules ](#graphql-rules) below for details.

###### One of the arguments from this list "disable_ip_reputation, enable_ip_reputation" can be set

`disable_ip_reputation` - (Optional) No IP reputation configured this distribution (`Bool`).

`enable_ip_reputation` - (Optional) Enable IP reputation for all requests in this distribution. See [Ip Reputation Choice Enable Ip Reputation ](#ip-reputation-choice-enable-ip-reputation) below for details.

`jwt_validation` - (Optional) tokens or tokens that are not yet valid.. See [Jwt Validation ](#jwt-validation) below for details.

###### One of the arguments from this list "l7_ddos_action_block, l7_ddos_action_default, l7_ddos_action_js_challenge" must be set

`l7_ddos_action_block` - (Optional) Block suspicious sources (`Bool`).

`l7_ddos_action_default` - (Optional) Block suspicious sources (`Bool`).

`l7_ddos_action_js_challenge` - (Optional) Serve JavaScript challenge to suspicious sources. See [L7 Ddos Auto Mitigation Action L7 Ddos Action Js Challenge ](#l7-ddos-auto-mitigation-action-l7-ddos-action-js-challenge) below for details.

###### One of the arguments from this list "http, https, https_auto_cert" must be set

`http` - (Optional) CDN Distribution serving content over HTTP. See [Loadbalancer Type Http ](#loadbalancer-type-http) below for details.

`https` - (Optional) User is responsible for managing DNS.. See [Loadbalancer Type Https ](#loadbalancer-type-https) below for details.

`https_auto_cert` - (Optional) DNS records will be managed by Volterra.. See [Loadbalancer Type Https Auto Cert ](#loadbalancer-type-https-auto-cert) below for details.

###### One of the arguments from this list "disable_malicious_user_detection, enable_malicious_user_detection" must be set

`disable_malicious_user_detection` - (Optional) Disable malicious user detection for this distribution (`Bool`).

`enable_malicious_user_detection` - (Optional) Enable malicious user detection for all requests in this distribution (`Bool`).

`origin_pool` - (Required) x-required. See [Origin Pool ](#origin-pool) below for details.

`other_settings` - (Optional) x-displayName: "Other Settings". See [Other Settings ](#other-settings) below for details.

`protected_cookies` - (Optional) Note: We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.. See [Protected Cookies ](#protected-cookies) below for details.

###### One of the arguments from this list "api_rate_limit, disable_rate_limit, rate_limit" must be set

`api_rate_limit` - (Optional) Define rate limiting for one or more API endpoints. See [Rate Limit Choice Api Rate Limit ](#rate-limit-choice-api-rate-limit) below for details.

`disable_rate_limit` - (Optional) Rate limiting is not currently enabled for this load balancer (`Bool`).

`rate_limit` - (Optional) Define custom rate limiting parameters for this load balancer. See [Rate Limit Choice Rate Limit ](#rate-limit-choice-rate-limit) below for details.

###### One of the arguments from this list "active_service_policies, no_service_policies, service_policies_from_namespace" must be set

`active_service_policies` - (Optional) Apply the specified list of service policies and bypass the namespace service policy set. See [Service Policy Choice Active Service Policies ](#service-policy-choice-active-service-policies) below for details.

`no_service_policies` - (Optional) Do not apply any service policies i.e. bypass the namespace service policy set (`Bool`).

`service_policies_from_namespace` - (Optional) Apply the active service policies configured as part of the namespace service policy set (`Bool`).

###### One of the arguments from this list "slow_ddos_mitigation, system_default_timeouts" must be set

`slow_ddos_mitigation` - (Optional) Custom Settings for Slow DDoS Mitigation. See [Slow Ddos Mitigation Choice Slow Ddos Mitigation ](#slow-ddos-mitigation-choice-slow-ddos-mitigation) below for details.

`system_default_timeouts` - (Optional) Default Settings for Slow DDoS Mitigation (`Bool`).

`trusted_clients` - (Optional) Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients.. See [Trusted Clients ](#trusted-clients) below for details.

###### One of the arguments from this list "user_id_client_ip, user_identification" must be set

`user_id_client_ip` - (Optional) Use the Client IP address as the user identifier. (`Bool`).

`user_identification` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier.. See [ref](#ref) below for details.

###### One of the arguments from this list "app_firewall, disable_waf" must be set

`app_firewall` - (Optional) Enable WAF configuration for all requests in this distribution. See [ref](#ref) below for details.

`disable_waf` - (Optional) No WAF configuration for this load balancer (`Bool`).

`waf_exclusion_rules` - (Optional) When an exclusion rule is matched, then this exclusion rule takes effect and no more rules are evaluated.. See [Waf Exclusion Rules ](#waf-exclusion-rules) below for details.

### Blocked Clients

Define rules to block IP Prefixes or AS numbers..

`actions` - (Optional) Actions that should be taken when client identifier matches the rule (`List of Strings`).

###### One of the arguments from this list "as_number, http_header, ip_prefix, ipv6_prefix, user_identifier" must be set

`as_number` - (Optional) RFC 6793 defined 4-byte AS number (`Int`).

`http_header` - (Optional) Request header name and value pairs. See [Client Source Choice Http Header ](#client-source-choice-http-header) below for details.

`ip_prefix` - (Optional) IPv4 prefix string. (`String`).

`ipv6_prefix` - (Optional) IPv6 prefix string. (`String`).

`user_identifier` - (Optional) Identify user based on user identifier. User identifier value needs to be copied from security event. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Blocked Clients Metadata ](#blocked-clients-metadata) below for details.

### Cors Policy

resources from a server at a different origin.

`allow_credentials` - (Optional) Specifies whether the resource allows credentials (`Bool`).

`allow_headers` - (Optional) Specifies the content for the access-control-allow-headers header (`String`).

`allow_methods` - (Optional) Specifies the content for the access-control-allow-methods header (`String`).

`allow_origin` - (Optional) An origin is allowed if either allow_origin or allow_origin_regex match (`String`).

`allow_origin_regex` - (Optional) An origin is allowed if either allow_origin or allow_origin_regex match (`String`).

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`expose_headers` - (Optional) Specifies the content for the access-control-expose-headers header (`String`).

`maximum_age` - (Optional) Maximum permitted value is 86400 seconds (24 hours) (`Int`).

### Csrf Policy

Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)..

###### One of the arguments from this list "all_load_balancer_domains, custom_domain_list, disabled" must be set

`all_load_balancer_domains` - (Optional) Add All load balancer domains to source origin (allow) list. (`Bool`).

`custom_domain_list` - (Optional) Add one or more domains to source origin (allow) list.. See [Allowed Domains Custom Domain List ](#allowed-domains-custom-domain-list) below for details.

`disabled` - (Optional) Allow all source origin domains. (`Bool`).

### Custom Cache Rule

Caching policies for CDN. .

`cdn_cache_rules` - (Optional) Reference to CDN Cache Rule configuration object. See [ref](#ref) below for details.

### Data Guard Rules

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

### Ddos Mitigation Rules

Define manual mitigation rules to block L7 DDoS attacks..

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Ddos Mitigation Rules Metadata ](#ddos-mitigation-rules-metadata) below for details.

###### One of the arguments from this list "block" must be set

`block` - (Optional) Block user for a duration determined by the expiration time (`Bool`).

###### One of the arguments from this list "ddos_client_source, ip_prefix_list" must be set

`ddos_client_source` - (Optional) Combination of Region, ASN and TLS Fingerprints. See [Mitigation Choice Ddos Client Source ](#mitigation-choice-ddos-client-source) below for details.

`ip_prefix_list` - (Optional) IP prefix string.. See [Mitigation Choice Ip Prefix List ](#mitigation-choice-ip-prefix-list) below for details.

### Default Cache Action

Default value for Cache action..

###### One of the arguments from this list "cache_disabled, cache_ttl_default, cache_ttl_override" can be set

`cache_disabled` - (Optional) Do not cache any content from the origin (`Bool`).

`cache_ttl_default` - (Optional) Use Cache TTL Provided by Origin, and set a contigency TTL value in case one is not provided (`String`).

`cache_ttl_override` - (Optional) Always override the Cahce TTL provided by Origin (`String`).

### Graphql Rules

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

### Jwt Validation

tokens or tokens that are not yet valid..

`action` - (Required) x-required. See [Jwt Validation Action ](#jwt-validation-action) below for details.

###### One of the arguments from this list "jwks_config" must be set

`jwks_config` - (Optional) The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details.. See [Jwks Configuration Jwks Config ](#jwks-configuration-jwks-config) below for details.

`mandatory_claims` - (Optional) If the claim does not exist JWT token validation will fail.. See [Jwt Validation Mandatory Claims ](#jwt-validation-mandatory-claims) below for details.

`reserved_claims` - (Optional) the token validation of these claims should be disabled.. See [Jwt Validation Reserved Claims ](#jwt-validation-reserved-claims) below for details.

`target` - (Required) Define endpoints for which JWT token validation will be performed. See [Jwt Validation Target ](#jwt-validation-target) below for details.

`token_location` - (Required) Define where in the HTTP request the JWT token will be extracted. See [Jwt Validation Token Location ](#jwt-validation-token-location) below for details.

### Origin Pool

x-required.

`more_origin_options` - (Optional) x-displayName: "Advanced Configuration". See [Origin Pool More Origin Options ](#origin-pool-more-origin-options) below for details.

`origin_request_timeout` - (Optional) Configures the time after which a request to the origin will time out waiting for a response (`String`).

`origin_servers` - (Required) List of original servers. See [Origin Pool Origin Servers ](#origin-pool-origin-servers) below for details.

`public_name` - (Optional) The DNS name to be used as the host header for the request to the origin server. See [Origin Pool Public Name ](#origin-pool-public-name) below for details.

###### One of the arguments from this list "no_tls, use_tls" must be set

`no_tls` - (Optional) Origin servers do not use TLS (`Bool`).

`use_tls` - (Optional) Origin servers use TLS. See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

### Other Settings

x-displayName: "Other Settings".

`add_location` - (Optional) Appends header x-volterra-location = <re-site-name> in responses. (`Bool`).

`header_options` - (Optional) Request/Response header related options. See [Other Settings Header Options ](#other-settings-header-options) below for details.

`logging_options` - (Optional) Logging related options. See [Other Settings Logging Options ](#other-settings-logging-options) below for details.

### Protected Cookies

Note: We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection..

###### One of the arguments from this list "disable_tampering_protection, enable_tampering_protection" must be set

`disable_tampering_protection` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_tampering_protection` - (Optional) x-displayName: "Enable" (`Bool`).

###### One of the arguments from this list "add_httponly, ignore_httponly" can be set

`add_httponly` - (Optional) x-displayName: "Add" (`Bool`).

`ignore_httponly` - (Optional) x-displayName: "Ignore" (`Bool`).

`name` - (Required) Name of the Cookie (`String`).

###### One of the arguments from this list "ignore_samesite, samesite_lax, samesite_none, samesite_strict" can be set

`ignore_samesite` - (Optional) Ignore Samesite attribute (`Bool`).

`samesite_lax` - (Optional) Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests (`Bool`).

`samesite_none` - (Optional) Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests (`Bool`).

`samesite_strict` - (Optional) Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests (`Bool`).

###### One of the arguments from this list "add_secure, ignore_secure" can be set

`add_secure` - (Optional) x-displayName: "Add" (`Bool`).

`ignore_secure` - (Optional) x-displayName: "Ignore" (`Bool`).

### Trusted Clients

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients..

`actions` - (Optional) Actions that should be taken when client identifier matches the rule (`List of Strings`).

###### One of the arguments from this list "as_number, http_header, ip_prefix, ipv6_prefix, user_identifier" must be set

`as_number` - (Optional) RFC 6793 defined 4-byte AS number (`Int`).

`http_header` - (Optional) Request header name and value pairs. See [Client Source Choice Http Header ](#client-source-choice-http-header) below for details.

`ip_prefix` - (Optional) IPv4 prefix string. (`String`).

`ipv6_prefix` - (Optional) IPv6 prefix string. (`String`).

`user_identifier` - (Optional) Identify user based on user identifier. User identifier value needs to be copied from security event. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Trusted Clients Metadata ](#trusted-clients-metadata) below for details.

### Waf Exclusion Rules

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

### Action Choice Action Block

Blocks the user for a specified duration of time.

###### One of the arguments from this list "hours, minutes, seconds" can be set

`hours` - (Optional) User block mitigation time in Hours. See [Block Duration Choice Hours ](#block-duration-choice-hours) below for details.

`minutes` - (Optional) User block mitigation time in Minutes. See [Block Duration Choice Minutes ](#block-duration-choice-minutes) below for details.

`seconds` - (Optional) User block mitigation time in Seconds. See [Block Duration Choice Seconds ](#block-duration-choice-seconds) below for details.

### Action Choice Apply Data Guard

x-displayName: "Apply".

### Action Choice Block

Block the request and report the issue.

### Action Choice Disabled

x-displayName: "Disabled".

### Action Choice Report

Allow the request and report the issue.

### Action Choice Skip Data Guard

x-displayName: "Skip".

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

### Api Endpoint Rules Api Endpoint Method

The predicate evaluates to true if the actual HTTP method belongs is present in the list of expected values..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`methods` - (Optional) x-example: "['GET', 'POST', 'DELETE']" (`List of Strings`).

### Api Endpoint Rules Client Matcher

Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc..

###### One of the arguments from this list "any_client, client_selector, ip_threat_category_list" must be set

`any_client` - (Optional) Any Client (`Bool`).

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.

`ip_threat_category_list` - (Optional) IP threat categories to choose from. See [Client Choice Ip Threat Category List ](#client-choice-ip-threat-category-list) below for details.

###### One of the arguments from this list "any_ip, asn_list, asn_matcher, ip_matcher, ip_prefix_list" must be set

`any_ip` - (Optional) Any Source IP (`Bool`).

`asn_list` - (Optional) The predicate evaluates to true if the origin ASN is present in the ASN list.. See [Ip Asn Choice Asn List ](#ip-asn-choice-asn-list) below for details.

`asn_matcher` - (Optional) The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects.. See [Ip Asn Choice Asn Matcher ](#ip-asn-choice-asn-matcher) below for details.

`ip_matcher` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes in the IP Prefix Sets.. See [Ip Asn Choice Ip Matcher ](#ip-asn-choice-ip-matcher) below for details.

`ip_prefix_list` - (Optional) The predicate evaluates to true if the client IP Address is covered by one or more of the IP Prefixes from the list.. See [Ip Asn Choice Ip Prefix List ](#ip-asn-choice-ip-prefix-list) below for details.

`tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints.. See [Client Matcher Tls Fingerprint Matcher ](#client-matcher-tls-fingerprint-matcher) below for details.

### Api Endpoint Rules Request Matcher

Conditions related to the request, such as query parameters, headers, etc..

`cookie_matchers` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Request Matcher Cookie Matchers ](#request-matcher-cookie-matchers) below for details.

`headers` - (Optional) Note that all specified header predicates must evaluate to true.. See [Request Matcher Headers ](#request-matcher-headers) below for details.

`jwt_claims` - (Optional) Note that this feature only works on LBs with JWT Validation feature enabled.. See [Request Matcher Jwt Claims ](#request-matcher-jwt-claims) below for details.

`query_params` - (Optional) Note that all specified query parameter predicates must evaluate to true.. See [Request Matcher Query Params ](#request-matcher-query-params) below for details.

### Api Rate Limit Api Endpoint Rules

For creating rule that contain a whole domain or group of endpoints, please use the server URL rules above..

`api_endpoint_method` - (Optional) The predicate evaluates to true if the actual HTTP method belongs is present in the list of expected values.. See [Api Endpoint Rules Api Endpoint Method ](#api-endpoint-rules-api-endpoint-method) below for details.

`api_endpoint_path` - (Required) The endpoint (path) of the request. (`String`).

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

### Block Duration Choice Hours

User block mitigation time in Hours.

`duration` - (Optional) x-displayName: "Duration" (`Int`).

### Block Duration Choice Minutes

User block mitigation time in Minutes.

`duration` - (Optional) x-displayName: "Duration" (`Int`).

### Block Duration Choice Seconds

User block mitigation time in Seconds.

`duration` - (Optional) x-displayName: "Duration" (`Int`).

### Blocked Clients Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Bypass Rate Limiting Rules Bypass Rate Limiting Rules

This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting..

`client_matcher` - (Optional) Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc.. See [Bypass Rate Limiting Rules Client Matcher ](#bypass-rate-limiting-rules-client-matcher) below for details.

###### One of the arguments from this list "any_url, api_endpoint, api_groups, base_path" must be set

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

###### One of the arguments from this list "any_client, client_selector, ip_threat_category_list" must be set

`any_client` - (Optional) Any Client (`Bool`).

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.

`ip_threat_category_list` - (Optional) IP threat categories to choose from. See [Client Choice Ip Threat Category List ](#client-choice-ip-threat-category-list) below for details.

###### One of the arguments from this list "any_ip, asn_list, asn_matcher, ip_matcher, ip_prefix_list" must be set

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

### Cache Actions Cache Disabled

Do not cache any content from the origin.

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

### Challenge Type Enable Challenge

Configure auto mitigation i.e risk based challenges for malicious users for this load balancer.

###### One of the arguments from this list "captcha_challenge_parameters, default_captcha_challenge_parameters" can be set

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

### Challenge Type Policy Based Challenge

Specifies the settings for policy rule based challenge.

###### One of the arguments from this list "captcha_challenge_parameters, default_captcha_challenge_parameters" can be set

`captcha_challenge_parameters` - (Optional) Configure captcha challenge parameters. See [Captcha Challenge Parameters Choice Captcha Challenge Parameters ](#captcha-challenge-parameters-choice-captcha-challenge-parameters) below for details.

`default_captcha_challenge_parameters` - (Optional) Use default parameters (`Bool`).

###### One of the arguments from this list "always_enable_captcha_challenge, always_enable_js_challenge, no_challenge" must be set

`always_enable_captcha_challenge` - (Optional) Challenge rules can be used to selectively disable Captcha challenge or enable JavaScript challenge for some requests. (`Bool`).

`always_enable_js_challenge` - (Optional) Challenge rules can be used to selectively disable JavaScript challenge or enable Captcha challenge for some requests. (`Bool`).

`no_challenge` - (Optional) Challenge rules can be used to selectively enable JavaScript or Captcha challenge for some requests. (`Bool`).

###### One of the arguments from this list "default_js_challenge_parameters, js_challenge_parameters" can be set

`default_js_challenge_parameters` - (Optional) Use default parameters (`Bool`).

`js_challenge_parameters` - (Optional) Configure JavaScript challenge parameters. See [Js Challenge Parameters Choice Js Challenge Parameters ](#js-challenge-parameters-choice-js-challenge-parameters) below for details.

###### One of the arguments from this list "default_mitigation_settings, malicious_user_mitigation" can be set

`default_mitigation_settings` - (Optional) For high level, users will be temporarily blocked. (`Bool`).

`malicious_user_mitigation` - (Optional) Define the mitigation actions to be taken for different threat levels. See [ref](#ref) below for details.

`rule_list` - (Optional) list challenge rules to be used in policy based challenge. See [Policy Based Challenge Rule List ](#policy-based-challenge-rule-list) below for details.

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

any_client.

### Client Choice Client Selector

client_selector.

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

### Cookie Tampering Disable Tampering Protection

x-displayName: "Disable".

### Cookie Tampering Enable Tampering Protection

x-displayName: "Enable".

### Count By Choice Use Http Lb User Id

Defined in HTTP-LB Security Configuration -> User Identifier..

### Crl Choice No Crl

Client certificate revocation status is not verified.

### Data Guard Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Data Guard Rules Path

URI path matcher..

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Ddos Client Source Asn List

The ASN is obtained by performing a lookup for the source IPv4 Address in a GeoIP DB..

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer. (`Int`).

### Ddos Client Source Ja4 Tls Fingerprint Matcher

The predicate evaluates to true if source JA4 TLS fingerprint matches any of the exact values of JA4 TLS fingerprints..

`exact_values` - (Optional) A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against (`String`).

### Ddos Client Source Tls Fingerprint Matcher

The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints..

`classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`exact_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

`excluded_values` - (Optional) or more known TLS fingerprint classes in the enclosing matcher. (`String`).

### Ddos Mitigation Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

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

Enable Data Guard for any domain.

### Graphql Rules Graphql Settings

GraphQL configuration..

###### One of the arguments from this list "disable_introspection, enable_introspection" must be set

`disable_introspection` - (Optional) Disable introspection queries for the load balancer. (`Bool`).

`enable_introspection` - (Optional) Enable introspection queries for the load balancer. (`Bool`).

`max_batched_queries` - (Required) Specify maximum number of queries in a single batched request. (`Int`).

`max_depth` - (Required) Specify maximum depth for the GraphQL query. (`Int`).

`max_total_length` - (Required) Specify maximum length in bytes for the GraphQL query. (`Int`).

### Graphql Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Header Options Request Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "secret_value, value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Header Options Response Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "secret_value, value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Http Header Headers

List of HTTP header name and value pairs.

`invert_match` - (Optional) Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Required) Name of the header (`String`).

###### One of the arguments from this list "exact, presence, regex" can be set

`exact` - (Optional) Header value to match exactly (`String`).

`presence` - (Optional) If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Httponly Add Httponly

x-displayName: "Add".

### Httponly Ignore Httponly

x-displayName: "Ignore".

### Https Tls Cert Options

TLS Certificate Options.

###### One of the arguments from this list "tls_cert_params, tls_inline_params" must be set

`tls_cert_params` - (Optional) Select/Add one or more TLS Certificate objects to associate with this Load Balancer. See [Tls Certificates Choice Tls Cert Params ](#tls-certificates-choice-tls-cert-params) below for details.

`tls_inline_params` - (Optional) Upload a TLS certificate covering all domain names for this Load Balancer. See [Tls Certificates Choice Tls Inline Params ](#tls-certificates-choice-tls-inline-params) below for details.

### Https Auto Cert Tls Config

TLS Configuration Parameters.

###### One of the arguments from this list "tls_11_plus, tls_12_plus" must be set

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

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

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

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

### Ip Reputation Choice Enable Ip Reputation

Enable IP reputation for all requests in this distribution.

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

###### One of the arguments from this list "audience, audience_disable" must be set

`audience` - (Optional) x-displayName: "Exact Match". See [Audience Validation Audience ](#audience-validation-audience) below for details.

`audience_disable` - (Optional) x-displayName: "Disable" (`Bool`).

###### One of the arguments from this list "issuer, issuer_disable" must be set

`issuer` - (Optional) x-displayName: "Exact Match" (`String`).

`issuer_disable` - (Optional) x-displayName: "Disable" (`Bool`).

###### One of the arguments from this list "validate_period_disable, validate_period_enable" must be set

`validate_period_disable` - (Optional) x-displayName: "Disable" (`Bool`).

`validate_period_enable` - (Optional) x-displayName: "Enable" (`Bool`).

### Jwt Validation Target

Define endpoints for which JWT token validation will be performed.

###### One of the arguments from this list "all_endpoint, api_groups, base_paths" must be set

`all_endpoint` - (Optional) Validation will be performed for all requests on this LB (`Bool`).

`api_groups` - (Optional) Validation will be performed for the endpoints mentioned in the API Groups. See [Target Api Groups ](#target-api-groups) below for details.

`base_paths` - (Optional) Validation will be performed for selected path prefixes. See [Target Base Paths ](#target-base-paths) below for details.

### Jwt Validation Token Location

Define where in the HTTP request the JWT token will be extracted.

###### One of the arguments from this list "bearer_token" must be set

`bearer_token` - (Optional) Token is found in Authorization HTTP header with Bearer authentication scheme (`Bool`).

### L7 Ddos Auto Mitigation Action L7 Ddos Action Js Challenge

Serve JavaScript challenge to suspicious sources.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

`js_script_delay` - (Optional) Delay introduced by Javascript, in milliseconds. (`Int`).

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

`tls_cert_options` - (Optional) TLS Certificate Options. See [Https Tls Cert Options ](#https-tls-cert-options) below for details.

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

### Malicious User Mitigation Choice Default Mitigation Settings

For high level, users will be temporarily blocked..

### Match Check Not Present

Check that the argument is not present..

### Match Check Present

Check that the argument is present..

### Match Item

Criteria for matching the values for the Arg. The match is successful if any of the values in the input satisfies the criteria in the matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Max Session Keys Type Default Session Key Caching

Default session key caching. Only one session key will be cached..

### Max Session Keys Type Disable Session Key Caching

Disable session key caching. This will disable TLS session resumption..

### Method Choice Method Get

x-displayName: "GET".

### Method Choice Method Post

x-displayName: "POST".

### Mitigation Action Block

Block user for a duration determined by the expiration time.

### Mitigation Choice Ddos Client Source

Combination of Region, ASN and TLS Fingerprints.

`asn_list` - (Optional) The ASN is obtained by performing a lookup for the source IPv4 Address in a GeoIP DB.. See [Ddos Client Source Asn List ](#ddos-client-source-asn-list) below for details.

`country_list` - (Optional) Sources that are located in one of the countries in the given list (`List of Strings`).

`ja4_tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if source JA4 TLS fingerprint matches any of the exact values of JA4 TLS fingerprints.. See [Ddos Client Source Ja4 Tls Fingerprint Matcher ](#ddos-client-source-ja4-tls-fingerprint-matcher) below for details.

`tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints.. See [Ddos Client Source Tls Fingerprint Matcher ](#ddos-client-source-tls-fingerprint-matcher) below for details.

### Mitigation Choice Ip Prefix List

IP prefix string..

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

### Mtls Choice No Mtls

x-displayName: "Disable".

### Mtls Choice Use Mtls

x-displayName: "Enable".

`client_certificate_optional` - (Optional) the connection will be accepted. (`Bool`).

###### One of the arguments from this list "crl, no_crl" can be set

`crl` - (Optional) Specify the CRL server information to download the certificate revocation list. See [ref](#ref) below for details.

`no_crl` - (Optional) Client certificate revocation status is not verified (`Bool`).

###### One of the arguments from this list "trusted_ca, trusted_ca_url" must be set

`trusted_ca` - (Optional) Select/Add a Root CA Certificate object to associate with this Load Balancer. See [ref](#ref) below for details.

`trusted_ca_url` - (Optional) Upload a Root CA Certificate specifically for this Load Balancer (`String`).

###### One of the arguments from this list "xfcc_disabled, xfcc_options" can be set

`xfcc_disabled` - (Optional) No X-Forwarded-Client-Cert header will be added (`Bool`).

`xfcc_options` - (Optional) X-Forwarded-Client-Cert header will be added with the configured fields. See [Xfcc Header Xfcc Options ](#xfcc-header-xfcc-options) below for details.

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

### Other Settings Header Options

Request/Response header related options.

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Header Options Request Headers To Add ](#header-options-request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Header Options Response Headers To Add ](#header-options-response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

### Other Settings Logging Options

Logging related options.

`client_log_options` - (Optional) Client request headers to log. See [Logging Options Client Log Options ](#logging-options-client-log-options) below for details.

`origin_log_options` - (Optional) Origin response headers to log. See [Logging Options Origin Log Options ](#logging-options-origin-log-options) below for details.

### Path Choice Any Path

Match all paths.

### Policy Based Challenge Rule List

list challenge rules to be used in policy based challenge.

`rules` - (Optional) these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions. See [Rule List Rules ](#rule-list-rules) below for details.

### Policy Choice No Policies

Do not apply additional rate limiter policies..

### Policy Choice Policies

to the action configured in the rule. If there's no match, the rate limiting configuration for the HTTP load balancer is honored..

`policies` - (Required) Ordered list of rate limiter policies.. See [ref](#ref) below for details.

### Rate Limit Rate Limiter

Requests to the virtual_host are rate limited based on the parameters specified in the rate_limiter..

###### One of the arguments from this list "action_block, disabled" can be set

`action_block` - (Optional) Blocks the user for a specified duration of time. See [Action Choice Action Block ](#action-choice-action-block) below for details.

`disabled` - (Optional) x-displayName: "Disabled" (`Bool`).

`burst_multiplier` - (Optional) The maximum burst of requests to accommodate, expressed as a multiple of the rate. (`Int`).

`period_multiplier` - (Optional) This setting, combined with Per Period units, provides a duration (`Int`).

`total_number` - (Required) The total number of allowed requests per rate-limiting period. (`Int`).

`unit` - (Required) Unit for the period per which the rate limit is applied. (`String`).

### Rate Limit Choice Api Rate Limit

Define rate limiting for one or more API endpoints.

`api_endpoint_rules` - (Optional) For creating rule that contain a whole domain or group of endpoints, please use the server URL rules above.. See [Api Rate Limit Api Endpoint Rules ](#api-rate-limit-api-endpoint-rules) below for details.

###### One of the arguments from this list "bypass_rate_limiting_rules, custom_ip_allowed_list, ip_allowed_list, no_ip_allowed_list" must be set

`bypass_rate_limiting_rules` - (Optional) This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting.. See [Ip Allowed List Choice Bypass Rate Limiting Rules ](#ip-allowed-list-choice-bypass-rate-limiting-rules) below for details.

`custom_ip_allowed_list` - (Optional) IP Allowed list using existing ip_prefix_set objects.. See [Ip Allowed List Choice Custom Ip Allowed List ](#ip-allowed-list-choice-custom-ip-allowed-list) below for details.

`ip_allowed_list` - (Optional) List of IP(s) for which rate limiting will be disabled.. See [Ip Allowed List Choice Ip Allowed List ](#ip-allowed-list-choice-ip-allowed-list) below for details.

`no_ip_allowed_list` - (Optional) There is no ip allowed list for rate limiting, all clients go through rate limiting. (`Bool`).

`server_url_rules` - (Optional) For matching also specific endpoints you can use the API endpoint rules set bellow.. See [Api Rate Limit Server Url Rules ](#api-rate-limit-server-url-rules) below for details.

### Rate Limit Choice Rate Limit

Define custom rate limiting parameters for this load balancer.

###### One of the arguments from this list "custom_ip_allowed_list, ip_allowed_list, no_ip_allowed_list" must be set

`custom_ip_allowed_list` - (Optional) IP Allowed list using existing ip_prefix_set objects.. See [Ip Allowed List Choice Custom Ip Allowed List ](#ip-allowed-list-choice-custom-ip-allowed-list) below for details.

`ip_allowed_list` - (Optional) List of IP(s) for which rate limiting will be disabled.. See [Ip Allowed List Choice Ip Allowed List ](#ip-allowed-list-choice-ip-allowed-list) below for details.

`no_ip_allowed_list` - (Optional) There is no ip allowed list for rate limiting, all clients go through rate limiting. (`Bool`).

###### One of the arguments from this list "no_policies, policies" must be set

`no_policies` - (Optional) Do not apply additional rate limiter policies. (`Bool`).

`policies` - (Optional) to the action configured in the rule. If there's no match, the rate limiting configuration for the HTTP load balancer is honored.. See [Policy Choice Policies ](#policy-choice-policies) below for details.

`rate_limiter` - (Optional) Requests to the virtual_host are rate limited based on the parameters specified in the rate_limiter.. See [Rate Limit Rate Limiter ](#rate-limit-rate-limiter) below for details.

### Rate Limiter Choice Inline Rate Limiter

Specify rate values for the rule..

###### One of the arguments from this list "ref_user_id, use_http_lb_user_id" must be set

`ref_user_id` - (Optional) If traffic cannot be identified by the rules in the user_identification object, by default it will be identified by the HTTP-LB User Identifier.. See [ref](#ref) below for details.

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

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the cookie is not present. (`Bool`).

`check_present` - (Optional) Check that the cookie is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the cookie. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-sensitive cookie name. (`String`).

### Request Matcher Headers

Note that all specified header predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Request Matcher Jwt Claims

Note that this feature only works on LBs with JWT Validation feature enabled..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the JWT Claim is not present. (`Bool`).

`check_present` - (Optional) Check that the JWT Claim is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the JWT Claim. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) JWT claim name. (`String`).

### Request Matcher Query Params

Note that all specified query parameter predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`key` - (Required) A case-sensitive HTTP query parameter name. (`String`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the query parameter is not present. (`Bool`).

`check_present` - (Optional) Check that the query parameter is present. (`Bool`).

`item` - (Optional) criteria in the matcher.. See [Match Item ](#match-item) below for details.

### Request Timeout Choice Disable Request Timeout

x-displayName: "No Timeout".

### Rule List Rules

these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Rules Metadata ](#rules-metadata) below for details.

`spec` - (Required) Specification for the rule including match predicates and actions.. See [Rules Spec ](#rules-spec) below for details.

### Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Rules Spec

Specification for the rule including match predicates and actions..

`arg_matchers` - (Optional)arg_matchers. See [Spec Arg Matchers ](#spec-arg-matchers) below for details.

###### One of the arguments from this list "any_asn, asn_list, asn_matcher" can be set

`any_asn` - (Optional)any_asn (`Bool`).

`asn_list` - (Optional)asn_list. See [Asn Choice Asn List ](#asn-choice-asn-list) below for details.

`asn_matcher` - (Optional)asn_matcher. See [Asn Choice Asn Matcher ](#asn-choice-asn-matcher) below for details.

`body_matcher` - (Optional)body_matcher. See [Spec Body Matcher ](#spec-body-matcher) below for details.

###### One of the arguments from this list "disable_challenge, enable_captcha_challenge, enable_javascript_challenge" must be set

`disable_challenge` - (Optional) Disable the challenge type selected in PolicyBasedChallenge (`Bool`).

`enable_captcha_challenge` - (Optional) Enable captcha challenge (`Bool`).

`enable_javascript_challenge` - (Optional) Enable javascript challenge (`Bool`).

###### One of the arguments from this list "any_client, client_selector" can be set

`any_client` - (Optional)any_client (`Bool`).

`client_selector` - (Optional)client_selector. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.

`cookie_matchers` - (Optional)cookie_matchers. See [Spec Cookie Matchers ](#spec-cookie-matchers) below for details.

`domain_matcher` - (Optional)domain_matcher. See [Spec Domain Matcher ](#spec-domain-matcher) below for details.

`expiration_timestamp` - (Optional)expiration_timestamp (`String`).

`headers` - (Optional)headers. See [Spec Headers ](#spec-headers) below for details.

`http_method` - (Optional)http_method. See [Spec Http Method ](#spec-http-method) below for details.

###### One of the arguments from this list "any_ip, ip_matcher, ip_prefix_list" can be set

`any_ip` - (Optional)any_ip (`Bool`).

`ip_matcher` - (Optional)ip_matcher. See [Ip Choice Ip Matcher ](#ip-choice-ip-matcher) below for details.

`ip_prefix_list` - (Optional)ip_prefix_list. See [Ip Choice Ip Prefix List ](#ip-choice-ip-prefix-list) below for details.

`path` - (Optional)path. See [Spec Path ](#spec-path) below for details.

`query_params` - (Optional)query_params. See [Spec Query Params ](#spec-query-params) below for details.

###### One of the arguments from this list "tls_fingerprint_matcher" can be set

`tls_fingerprint_matcher` - (Optional)tls_fingerprint_matcher. See [Tls Fingerprint Choice Tls Fingerprint Matcher ](#tls-fingerprint-choice-tls-fingerprint-matcher) below for details.

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

### Secure Add Secure

x-displayName: "Add".

### Secure Ignore Secure

x-displayName: "Ignore".

### Server Url Rules Client Matcher

Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc..

###### One of the arguments from this list "any_client, client_selector, ip_threat_category_list" must be set

`any_client` - (Optional) Any Client (`Bool`).

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Choice Client Selector ](#client-choice-client-selector) below for details.

`ip_threat_category_list` - (Optional) IP threat categories to choose from. See [Client Choice Ip Threat Category List ](#client-choice-ip-threat-category-list) below for details.

###### One of the arguments from this list "any_ip, asn_list, asn_matcher, ip_matcher, ip_prefix_list" must be set

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

###### One of the arguments from this list "trusted_ca, trusted_ca_url" must be set

`trusted_ca` - (Optional) Select/Add a Root CA Certificate object to associate with this Origin Pool for verification of server's certificate. See [ref](#ref) below for details.

`trusted_ca_url` - (Optional) Upload a Root CA Certificate specifically for this Origin Pool for verification of server's certificate (`String`).

### Server Validation Choice Volterra Trusted Ca

Perform origin server verification using F5XC Default Root CA Certificate.

### Service Policy Choice Active Service Policies

Apply the specified list of service policies and bypass the namespace service policy set.

`policies` - (Required) If all policies are evaluated and none match, then the request will be denied by default.. See [ref](#ref) below for details.

### Slow Ddos Mitigation Choice Slow Ddos Mitigation

Custom Settings for Slow DDoS Mitigation.

`request_headers_timeout` - (Optional) provides protection against Slowloris attacks. (`Int`).

###### One of the arguments from this list "disable_request_timeout, request_timeout" must be set

`disable_request_timeout` - (Optional) x-displayName: "No Timeout" (`Bool`).

`request_timeout` - (Optional) x-example: "60000" (`Int`).

### Sni Choice Disable Sni

Do not use SNI..

### Sni Choice Use Host Header As Sni

Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied..

### Spec Arg Matchers

arg_matchers.

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the argument is not present. (`Bool`).

`check_present` - (Optional) Check that the argument is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the Arg. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-sensitive JSON path in the HTTP request body. (`String`).

### Spec Body Matcher

body_matcher.

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Spec Cookie Matchers

cookie_matchers.

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the cookie is not present. (`Bool`).

`check_present` - (Optional) Check that the cookie is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the cookie. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-sensitive cookie name. (`String`).

### Spec Domain Matcher

domain_matcher.

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Spec Headers

headers.

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

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

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the query parameter is not present. (`Bool`).

`check_present` - (Optional) Check that the query parameter is present. (`Bool`).

`item` - (Optional) criteria in the matcher.. See [Match Item ](#match-item) below for details.

### Target All Endpoint

Validation will be performed for all requests on this LB.

### Target Api Groups

Validation will be performed for the endpoints mentioned in the API Groups.

`api_groups` - (Required) x-required (`String`).

### Target Base Paths

Validation will be performed for selected path prefixes.

`base_paths` - (Required) x-required (`String`).

### Tls Cert Params Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Tls Certificates Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Tls Certificates Choice Tls Cert Params

Select/Add one or more TLS Certificate objects to associate with this Load Balancer.

`certificates` - (Required) Select one or more certificates with any domain names.. See [ref](#ref) below for details.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Cert Params Tls Config ](#tls-cert-params-tls-config) below for details.

### Tls Certificates Choice Tls Inline Params

Upload a TLS certificate covering all domain names for this Load Balancer.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Inline Params Tls Certificates ](#tls-inline-params-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Inline Params Tls Config ](#tls-inline-params-tls-config) below for details.

### Tls Choice No Tls

Origin servers do not use TLS.

### Tls Choice Use Tls

Origin servers use TLS.

###### One of the arguments from this list "default_session_key_caching, disable_session_key_caching, max_session_keys" must be set

`default_session_key_caching` - (Optional) Default session key caching. Only one session key will be cached. (`Bool`).

`disable_session_key_caching` - (Optional) Disable session key caching. This will disable TLS session resumption. (`Bool`).

`max_session_keys` - (Optional) Number of session keys that are cached. (`Int`).

###### One of the arguments from this list "no_mtls, use_mtls, use_mtls_obj" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Upload a client authentication certificate specifically for this Origin Pool". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`use_mtls_obj` - (Optional) x-displayName: "Select/add a TLS Certificate object for client authentication". See [ref](#ref) below for details.

###### One of the arguments from this list "skip_server_verification, use_server_verification, volterra_trusted_ca" must be set

`skip_server_verification` - (Optional) Skip origin server verification (`Bool`).

`use_server_verification` - (Optional) Perform origin server verification using the provided Root CA Certificate. See [Server Validation Choice Use Server Verification ](#server-validation-choice-use-server-verification) below for details.

`volterra_trusted_ca` - (Optional) Perform origin server verification using F5XC Default Root CA Certificate (`Bool`).

###### One of the arguments from this list "disable_sni, sni, use_host_header_as_sni" must be set

`disable_sni` - (Optional) Do not use SNI. (`Bool`).

`sni` - (Optional) SNI value to be used. (`String`).

`use_host_header_as_sni` - (Optional) Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied. (`Bool`).

`tls_config` - (Required) TLS parameters such as min/max TLS version and ciphers. See [Use Tls Tls Config ](#use-tls-tls-config) below for details.

### Tls Fingerprint Choice Tls Fingerprint Matcher

tls_fingerprint_matcher.

`classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`exact_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

`excluded_values` - (Optional) or more known TLS fingerprint classes in the enclosing matcher. (`String`).

### Tls Inline Params Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Tls Inline Params Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Token Location Bearer Token

Token is found in Authorization HTTP header with Bearer authentication scheme.

### Trusted Clients Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Use Mtls Tls Certificates

mTLS Client Certificate.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Use Tls Tls Config

TLS parameters such as min/max TLS version and ciphers.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Validate Period Validate Period Disable

x-displayName: "Disable".

### Validate Period Validate Period Enable

x-displayName: "Enable".

### Value Choice Secret Value

Secret Value of the HTTP header..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Waf Advanced Configuration App Firewall Detection Control

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria..

`exclude_attack_type_contexts` - (Optional) Attack Types to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Attack Type Contexts ](#app-firewall-detection-control-exclude-attack-type-contexts) below for details.

`exclude_bot_name_contexts` - (Optional) Bot Names to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Bot Name Contexts ](#app-firewall-detection-control-exclude-bot-name-contexts) below for details.

`exclude_signature_contexts` - (Optional) Signature IDs to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Signature Contexts ](#app-firewall-detection-control-exclude-signature-contexts) below for details.

`exclude_violation_contexts` - (Optional) Violations to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Violation Contexts ](#app-firewall-detection-control-exclude-violation-contexts) below for details.

### Waf Advanced Configuration Waf Skip Processing

Skip all App Firewall processing for this request.

### Waf Exclusion Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Xfcc Header Xfcc Disabled

No X-Forwarded-Client-Cert header will be added.

### Xfcc Header Xfcc Options

X-Forwarded-Client-Cert header will be added with the configured fields.

`xfcc_header_elements` - (Required) X-Forwarded-Client-Cert header elements to be added to requests (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured cdn_loadbalancer.
