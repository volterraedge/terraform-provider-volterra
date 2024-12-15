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

  // One of the arguments from this list "api_specification api_specification_on_cache_miss disable_api_definition" must be set

  disable_api_definition = true

  // One of the arguments from this list "api_discovery_on_cache_miss disable_api_discovery enable_api_discovery" must be set

  api_discovery_on_cache_miss {
    api_crawler {
      // One of the arguments from this list "api_crawler_config disable_api_crawler" must be set

      disable_api_crawler = true
    }

    api_discovery_from_code_scan {
      code_base_integrations {
        // One of the arguments from this list "all_repos selected_repos" must be set

        selected_repos {
          api_code_repo = ["api_code_repo"]
        }

        code_base_integration {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
      }
    }

    // One of the arguments from this list "custom_api_auth_discovery default_api_auth_discovery" must be set

    custom_api_auth_discovery {
      api_discovery_ref {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }
    }
    discovered_api_settings {}

    // One of the arguments from this list "disable_learn_from_redirect_traffic enable_learn_from_redirect_traffic" must be set

    disable_learn_from_redirect_traffic = true
    sensitive_data_detection_rules {}
  }

  // One of the arguments from this list "bot_defense bot_defense_advanced disable_bot_defense" must be set

  disable_bot_defense = true

  // One of the arguments from this list "captcha_challenge challenge_on_cache_miss enable_challenge js_challenge no_challenge policy_based_challenge" must be set

  js_challenge {
    cookie_expiry = "1000"

    custom_page = "string:///PHA+IFBsZWFzZSBXYWl0IDwvcD4="

    js_script_delay = "1000"
  }

  // One of the arguments from this list "client_side_defense disable_client_side_defense" must be set

  disable_client_side_defense = true
  domains = ["www.foo.com"]

  // One of the arguments from this list "l7_ddos_action_block l7_ddos_action_default l7_ddos_action_js_challenge l7_ddos_action_none" must be set

  l7_ddos_action_default = true

  // One of the arguments from this list "http https https_auto_cert" must be set

  http {
    dns_volterra_managed = true

    // One of the arguments from this list "port port_ranges" must be set

    port = "80"
  }

  // One of the arguments from this list "disable_malicious_user_detection enable_malicious_user_detection malicious_user_detection_on_cache_miss" must be set

  disable_malicious_user_detection = true
  origin_pool {
    follow_origin_redirect = true

    more_origin_options {
      disable_byte_range_request = true

      websocket_proxy = true
    }

    origin_request_timeout = "100s"

    origin_servers {
      // One of the arguments from this list "public_ip public_name" must be set

      public_ip {
        // One of the arguments from this list "ip ipv6" must be set

        ipv6 = "2001::1"
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

  disable_rate_limit = true

  // One of the arguments from this list "default_sensitive_data_policy sensitive_data_policy" must be set

  default_sensitive_data_policy = true

  // One of the arguments from this list "active_service_policies no_service_policies service_policies_from_namespace" must be set

  no_service_policies = true

  // One of the arguments from this list "slow_ddos_mitigation system_default_timeouts" must be set

  system_default_timeouts = true

  // One of the arguments from this list "disable_threat_mesh enable_threat_mesh" must be set

  disable_threat_mesh = true

  // One of the arguments from this list "user_id_client_ip user_identification" must be set

  user_id_client_ip = true

  // One of the arguments from this list "app_firewall app_firewall_on_cache_miss disable_waf" must be set

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

`add_location` - (Optional) Appends header x-volterra-location = <re-site-name> in responses. (`Bool`).(Deprecated)

###### One of the arguments from this list "api_specification, api_specification_on_cache_miss, disable_api_definition" must be set

`api_specification` - (Optional) Specify API definition and OpenAPI Validation. See [Api Definition Choice Api Specification ](#api-definition-choice-api-specification) below for details.(Deprecated)

`api_specification_on_cache_miss` - (Optional) Enable API definition and OpenAPI Validation only on cache miss in this distribution. See [Api Definition Choice Api Specification On Cache Miss ](#api-definition-choice-api-specification-on-cache-miss) below for details.(Deprecated)

`disable_api_definition` - (Optional) API Definition is not currently used for this load balancer (`Bool`).(Deprecated)

###### One of the arguments from this list "api_discovery_on_cache_miss, disable_api_discovery, enable_api_discovery" must be set

`api_discovery_on_cache_miss` - (Optional) Enable api discovery only on cache miss in this distribution. See [Api Discovery Choice Api Discovery On Cache Miss ](#api-discovery-choice-api-discovery-on-cache-miss) below for details.(Deprecated)

`disable_api_discovery` - (Optional) Disable api discovery for this distribution (`Bool`).(Deprecated)

`enable_api_discovery` - (Optional) Enable api discovery for all requests in this distribution. See [Api Discovery Choice Enable Api Discovery ](#api-discovery-choice-enable-api-discovery) below for details.(Deprecated)

`api_protection_rules` - (Optional) Rules can also include additional conditions, for example specific clients can access certain API endpoint or API group.. See [Api Protection Rules ](#api-protection-rules) below for details.(Deprecated)

`blocked_clients` - (Optional) Define rules to block IP Prefixes or AS numbers.. See [Blocked Clients ](#blocked-clients) below for details.

###### One of the arguments from this list "bot_defense, bot_defense_advanced, disable_bot_defense" must be set

`bot_defense` - (Optional) Select Bot Defense Standard. See [Bot Defense Choice Bot Defense ](#bot-defense-choice-bot-defense) below for details.(Deprecated)

`bot_defense_advanced` - (Optional) Select Bot Defense Advanced. See [Bot Defense Choice Bot Defense Advanced ](#bot-defense-choice-bot-defense-advanced) below for details.(Deprecated)

`disable_bot_defense` - (Optional) No Bot Defense configuration for this load balancer (`Bool`).(Deprecated)

`cache_rules` - (Optional) Rules are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs.. See [Cache Rules ](#cache-rules) below for details.

###### One of the arguments from this list "captcha_challenge, challenge_on_cache_miss, enable_challenge, js_challenge, no_challenge, policy_based_challenge" must be set

`captcha_challenge` - (Optional) Configure Captcha challenge on this load balancer. See [Challenge Type Captcha Challenge ](#challenge-type-captcha-challenge) below for details.

`challenge_on_cache_miss` - (Optional) Configure auto mitigation i.e risk based challenges for malicious users only on cache miss in this load balancer. See [Challenge Type Challenge On Cache Miss ](#challenge-type-challenge-on-cache-miss) below for details.(Deprecated)

`enable_challenge` - (Optional) Configure auto mitigation i.e risk based challenges for malicious users for this load balancer. See [Challenge Type Enable Challenge ](#challenge-type-enable-challenge) below for details.

`js_challenge` - (Optional) Configure JavaScript challenge on this load balancer. See [Challenge Type Js Challenge ](#challenge-type-js-challenge) below for details.

`no_challenge` - (Optional) No challenge is enabled for this load balancer (`Bool`).

`policy_based_challenge` - (Optional) Specifies the settings for policy rule based challenge. See [Challenge Type Policy Based Challenge ](#challenge-type-policy-based-challenge) below for details.

###### One of the arguments from this list "client_side_defense, disable_client_side_defense" must be set

`client_side_defense` - (Optional) Client-Side Defense configuration for JavaScript insertion. See [Client Side Defense Choice Client Side Defense ](#client-side-defense-choice-client-side-defense) below for details.(Deprecated)

`disable_client_side_defense` - (Optional) No Client-Side Defense configuration for this load balancer (`Bool`).(Deprecated)

`cors_policy` - (Optional) resources from a server at a different origin. See [Cors Policy ](#cors-policy) below for details.

`csrf_policy` - (Optional) Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.).. See [Csrf Policy ](#csrf-policy) below for details.

`data_guard_rules` - (Optional) Note: App Firewall should be enabled, to use Data Guard feature.. See [Data Guard Rules ](#data-guard-rules) below for details.

`ddos_mitigation_rules` - (Optional) Define manual mitigation rules to block L7 DDoS attacks.. See [Ddos Mitigation Rules ](#ddos-mitigation-rules) below for details.

`default_cache_action` - (Optional) Default value for Cache action.. See [Default Cache Action ](#default-cache-action) below for details.

`domains` - (Required) [This can be a domain or a sub-domain](`List of String`).

`graphql_rules` - (Optional) queries and prevent GraphQL tailored attacks.. See [Graphql Rules ](#graphql-rules) below for details.

###### One of the arguments from this list "disable_ip_reputation, enable_ip_reputation, ip_reputation_on_cache_miss" can be set

`disable_ip_reputation` - (Optional) No IP reputation configured this distribution (`Bool`).

`enable_ip_reputation` - (Optional) Enable IP reputation for all requests in this distribution. See [Ip Reputation Choice Enable Ip Reputation ](#ip-reputation-choice-enable-ip-reputation) below for details.

`ip_reputation_on_cache_miss` - (Optional) Enable IP reputation only on cache miss in this distribution. See [Ip Reputation Choice Ip Reputation On Cache Miss ](#ip-reputation-choice-ip-reputation-on-cache-miss) below for details.(Deprecated)

`jwt_validation` - (Optional) tokens or tokens that are not yet valid.. See [Jwt Validation ](#jwt-validation) below for details.

###### One of the arguments from this list "l7_ddos_action_block, l7_ddos_action_default, l7_ddos_action_js_challenge, l7_ddos_action_none" must be set

`l7_ddos_action_block` - (Optional) Block suspicious sources (`Bool`).

`l7_ddos_action_default` - (Optional) Block suspicious sources (`Bool`).

`l7_ddos_action_js_challenge` - (Optional) Serve JavaScript challenge to suspicious sources. See [L7 Ddos Auto Mitigation Action L7 Ddos Action Js Challenge ](#l7-ddos-auto-mitigation-action-l7-ddos-action-js-challenge) below for details.

`l7_ddos_action_none` - (Optional) Disable auto mitigation (`Bool`).(Deprecated)

###### One of the arguments from this list "http, https, https_auto_cert" must be set

`http` - (Optional) CDN Distribution serving content over HTTP. See [Loadbalancer Type Http ](#loadbalancer-type-http) below for details.

`https` - (Optional) User is responsible for managing DNS.. See [Loadbalancer Type Https ](#loadbalancer-type-https) below for details.

`https_auto_cert` - (Optional) DNS records will be managed by Volterra.. See [Loadbalancer Type Https Auto Cert ](#loadbalancer-type-https-auto-cert) below for details.

###### One of the arguments from this list "disable_malicious_user_detection, enable_malicious_user_detection, malicious_user_detection_on_cache_miss" must be set

`disable_malicious_user_detection` - (Optional) Disable malicious user detection for this distribution (`Bool`).

`enable_malicious_user_detection` - (Optional) Enable malicious user detection for all requests in this distribution (`Bool`).

`malicious_user_detection_on_cache_miss` - (Optional) Enable malicious user detection only on cache miss in this distribution (`Bool`).(Deprecated)

`more_option` - (Optional) More options like header manipulation, compression etc.. See [More Option ](#more-option) below for details.(Deprecated)

`origin_pool` - (Required) x-required. See [Origin Pool ](#origin-pool) below for details.

`other_settings` - (Optional) x-displayName: "Other Settings". See [Other Settings ](#other-settings) below for details.

`protected_cookies` - (Optional) Note: We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.. See [Protected Cookies ](#protected-cookies) below for details.

###### One of the arguments from this list "api_rate_limit, disable_rate_limit, rate_limit" must be set

`api_rate_limit` - (Optional) Define rate limiting for one or more API endpoints. See [Rate Limit Choice Api Rate Limit ](#rate-limit-choice-api-rate-limit) below for details.

`disable_rate_limit` - (Optional) Rate limiting is not currently enabled for this load balancer (`Bool`).

`rate_limit` - (Optional) Define custom rate limiting parameters for this load balancer. See [Rate Limit Choice Rate Limit ](#rate-limit-choice-rate-limit) below for details.

###### One of the arguments from this list "default_sensitive_data_policy, sensitive_data_policy" must be set

`default_sensitive_data_policy` - (Optional) Apply system default sensitive data discovery (`Bool`).(Deprecated)

`sensitive_data_policy` - (Optional) Apply custom sensitive data discovery. See [Sensitive Data Policy Choice Sensitive Data Policy ](#sensitive-data-policy-choice-sensitive-data-policy) below for details.(Deprecated)

###### One of the arguments from this list "active_service_policies, no_service_policies, service_policies_from_namespace" must be set

`active_service_policies` - (Optional) Apply the specified list of service policies and bypass the namespace service policy set. See [Service Policy Choice Active Service Policies ](#service-policy-choice-active-service-policies) below for details.

`no_service_policies` - (Optional) Do not apply any service policies i.e. bypass the namespace service policy set (`Bool`).

`service_policies_from_namespace` - (Optional) Apply the active service policies configured as part of the namespace service policy set (`Bool`).

###### One of the arguments from this list "slow_ddos_mitigation, system_default_timeouts" must be set

`slow_ddos_mitigation` - (Optional) Custom Settings for Slow DDoS Mitigation. See [Slow Ddos Mitigation Choice Slow Ddos Mitigation ](#slow-ddos-mitigation-choice-slow-ddos-mitigation) below for details.

`system_default_timeouts` - (Optional) Default Settings for Slow DDoS Mitigation (`Bool`).

###### One of the arguments from this list "disable_threat_mesh, enable_threat_mesh" must be set

`disable_threat_mesh` - (Optional) x-displayName: "Disable" (`Bool`).(Deprecated)

`enable_threat_mesh` - (Optional) x-displayName: "Enable" (`Bool`).(Deprecated)

`trusted_clients` - (Optional) Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients.. See [Trusted Clients ](#trusted-clients) below for details.

###### One of the arguments from this list "user_id_client_ip, user_identification" must be set

`user_id_client_ip` - (Optional) Use the Client IP address as the user identifier. (`Bool`).

`user_identification` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier.. See [ref](#ref) below for details.

###### One of the arguments from this list "app_firewall, app_firewall_on_cache_miss, disable_waf" must be set

`app_firewall` - (Optional) Enable WAF configuration for all requests in this distribution. See [ref](#ref) below for details.

`app_firewall_on_cache_miss` - (Optional) Enable WAF configuration only on cache miss in this distribution. See [ref](#ref) below for details.(Deprecated)

`disable_waf` - (Optional) No WAF configuration for this load balancer (`Bool`).

`waf_exclusion_rules` - (Optional) When an exclusion rule is matched, then this exclusion rule takes effect and no more rules are evaluated.. See [Waf Exclusion Rules ](#waf-exclusion-rules) below for details.

### Api Protection Rules

Rules can also include additional conditions, for example specific clients can access certain API endpoint or API group..

`api_endpoint_rules` - (Optional) If request matches any of these rules, skipping second category rules.. See [Api Protection Rules Api Endpoint Rules ](#api-protection-rules-api-endpoint-rules) below for details.

`api_groups_rules` - (Optional) For API groups, refer to API Definition which includes API groups derived from uploaded swaggers.. See [Api Protection Rules Api Groups Rules ](#api-protection-rules-api-groups-rules) below for details.

### Blocked Clients

Define rules to block IP Prefixes or AS numbers..

###### One of the arguments from this list "bot_skip_processing, skip_processing, waf_skip_processing" can be set

`bot_skip_processing` - (Optional) Skip Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`skip_processing` - (Optional) Skip both WAF and Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (`Bool`).(Deprecated)

`actions` - (Optional) Actions that should be taken when client identifier matches the rule (`List of Strings`).

###### One of the arguments from this list "as_number, http_header, ip_prefix, ipv6_prefix, user_identifier" must be set

`as_number` - (Optional) RFC 6793 defined 4-byte AS number (`Int`).

`http_header` - (Optional) Request header name and value pairs. See [Client Source Choice Http Header ](#client-source-choice-http-header) below for details.

`ip_prefix` - (Optional) IPv4 prefix string. (`String`).

`ipv6_prefix` - (Optional) IPv6 prefix string. (`String`).

`user_identifier` - (Optional) Identify user based on user identifier. User identifier value needs to be copied from security event. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Blocked Clients Metadata ](#blocked-clients-metadata) below for details.

### Cache Rules

Rules are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs..

###### One of the arguments from this list "cache_bypass, eligible_for_cache" must be set

`cache_bypass` - (Optional) Bypass Caching of content from the origin (`Bool`).

`eligible_for_cache` - (Optional) Eligible for caching the content. See [Cache Actions Eligible For Cache ](#cache-actions-eligible-for-cache) below for details.

`rule_expression_list` - (Required) Expressions are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs... See [Cache Rules Rule Expression List ](#cache-rules-rule-expression-list) below for details.

`rule_name` - (Required) Name of the Cache Rule (`String`).

### Cors Policy

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

### Csrf Policy

Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)..

###### One of the arguments from this list "all_load_balancer_domains, custom_domain_list, disabled" must be set

`all_load_balancer_domains` - (Optional) Add All load balancer domains to source origin (allow) list. (`Bool`).

`custom_domain_list` - (Optional) Add one or more domains to source origin (allow) list.. See [Allowed Domains Custom Domain List ](#allowed-domains-custom-domain-list) below for details.

`disabled` - (Optional) Allow all source origin domains. (`Bool`).

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

###### One of the arguments from this list "cache_disabled, cache_ttl_default, cache_ttl_override, eligible_for_cache" can be set

`cache_disabled` - (Optional) Disable Caching of content from the origin (`Bool`).

`cache_ttl_default` - (Optional) Cache TTL value to use when the origin does not provide one (`String`).

`cache_ttl_override` - (Optional) Override the Cache TTL directive in the response from the origin (`String`).

`eligible_for_cache` - (Optional) Eligible for caching the content. See [Cache Actions Eligible For Cache ](#cache-actions-eligible-for-cache) below for details.(Deprecated)

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

###### One of the arguments from this list "auth_server_uri, jwks, jwks_config" must be set

`auth_server_uri` - (Optional) JWKS URI will be will be retrieved from this URI (`String`).(Deprecated)

`jwks` - (Optional) The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details. (`String`).(Deprecated)

`jwks_config` - (Optional) The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details.. See [Jwks Configuration Jwks Config ](#jwks-configuration-jwks-config) below for details.

`mandatory_claims` - (Optional) If the claim does not exist JWT token validation will fail.. See [Jwt Validation Mandatory Claims ](#jwt-validation-mandatory-claims) below for details.

`reserved_claims` - (Optional) the token validation of these claims should be disabled.. See [Jwt Validation Reserved Claims ](#jwt-validation-reserved-claims) below for details.

`target` - (Required) Define endpoints for which JWT token validation will be performed. See [Jwt Validation Target ](#jwt-validation-target) below for details.

`token_location` - (Required) Define where in the HTTP request the JWT token will be extracted. See [Jwt Validation Token Location ](#jwt-validation-token-location) below for details.

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

### Other Settings

x-displayName: "Other Settings".

`add_location` - (Optional) Appends header x-volterra-location = <re-site-name> in responses. (`Bool`).

`geo_filtering` - (Optional) Geo filtering options. See [Other Settings Geo Filtering ](#other-settings-geo-filtering) below for details.(Deprecated)

`header_options` - (Optional) Request/Response header related options. See [Other Settings Header Options ](#other-settings-header-options) below for details.

`ip_filtering` - (Optional) IP filtering options. See [Other Settings Ip Filtering ](#other-settings-ip-filtering) below for details.(Deprecated)

`logging_options` - (Optional) Logging related options. See [Other Settings Logging Options ](#other-settings-logging-options) below for details.

### Protected Cookies

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

###### One of the arguments from this list "bot_skip_processing, skip_processing, waf_skip_processing" can be set

`bot_skip_processing` - (Optional) Skip Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`skip_processing` - (Optional) Skip both WAF and Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (`Bool`).(Deprecated)

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

### Action Type Block

Block bot request and send response with custom content..

`body` - (Optional) E.g. "<p> Your request was blocked </p>". Base64 encoded string for this html is "LzxwPiBZb3VyIHJlcXVlc3Qgd2FzIGJsb2NrZWQgPC9wPg==" (`String`).

`body_hash` - (Optional) Represents the corresponding MD5 Hash for the body message. (`String`).(Deprecated)

`status` - (Optional) HTTP Status code to respond with (`String`).

### Action Type Flag

Flag the request while not taking any invasive actions..

###### One of the arguments from this list "append_headers, no_headers" can be set

`append_headers` - (Optional) Append mitigation headers.. See [Send Headers Choice Append Headers ](#send-headers-choice-append-headers) below for details.

`no_headers` - (Optional) No mitigation headers. (`Bool`).

### Action Type None

No mitigation actions..

### Action Type Redirect

Redirect bot request to a custom URI..

`uri` - (Required) URI location for redirect may be relative or absolute. (`String`).

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

### Api Crawler Api Crawler Config

Select to activate the API Crawling.

`domains` - (Required) Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer.. See [Api Crawler Config Domains ](#api-crawler-config-domains) below for details.

### Api Crawler Disable Api Crawler

Select to turn off the API Crawling. No API Crawling actions will be performed..

### Api Crawler Config Domains

Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer..

`domain` - (Required) Select the domain to execute API Crawling with given credentials. (`String`).

`simple_login` - (Required) Enter the username and password to assign credentials for the selected domain to crawl. See [Domains Simple Login ](#domains-simple-login) below for details.

### Api Definition Choice Api Specification

Specify API definition and OpenAPI Validation.

`api_definition` - (Required) Specify API definition which includes application API paths and methods derived from swagger files.. See [ref](#ref) below for details.

###### One of the arguments from this list "validation_all_spec_endpoints, validation_custom_list, validation_disabled" must be set

`validation_all_spec_endpoints` - (Optional) All other API endpoints would proceed according to "Fall Through Mode". See [Validation Target Choice Validation All Spec Endpoints ](#validation-target-choice-validation-all-spec-endpoints) below for details.

`validation_custom_list` - (Optional) Any other end-points not listed will act according to "Fall Through Mode". See [Validation Target Choice Validation Custom List ](#validation-target-choice-validation-custom-list) below for details.

`validation_disabled` - (Optional) Don't run OpenAPI validation (`Bool`).

### Api Definition Choice Api Specification On Cache Miss

Enable API definition and OpenAPI Validation only on cache miss in this distribution.

`api_definition` - (Required) Specify API definition which includes application API paths and methods derived from swagger files.. See [ref](#ref) below for details.

###### One of the arguments from this list "validation_all_spec_endpoints, validation_custom_list, validation_disabled" must be set

`validation_all_spec_endpoints` - (Optional) All other API endpoints would proceed according to "Fall Through Mode". See [Validation Target Choice Validation All Spec Endpoints ](#validation-target-choice-validation-all-spec-endpoints) below for details.

`validation_custom_list` - (Optional) Any other end-points not listed will act according to "Fall Through Mode". See [Validation Target Choice Validation Custom List ](#validation-target-choice-validation-custom-list) below for details.

`validation_disabled` - (Optional) Don't run OpenAPI validation (`Bool`).

### Api Definition Choice Disable Api Definition

API Definition is not currently used for this load balancer.

### Api Discovery Choice Api Discovery On Cache Miss

Enable api discovery only on cache miss in this distribution.

`api_crawler` - (Optional) Configure Discovered API Settings.. See [Api Discovery On Cache Miss Api Crawler ](#api-discovery-on-cache-miss-api-crawler) below for details.

`api_discovery_from_code_scan` - (Optional) Select API code repositories to the load balancer to use them as a source for API endpoint discovery.. See [Api Discovery On Cache Miss Api Discovery From Code Scan ](#api-discovery-on-cache-miss-api-discovery-from-code-scan) below for details.

###### One of the arguments from this list "custom_api_auth_discovery, default_api_auth_discovery" must be set

`custom_api_auth_discovery` - (Optional) Apply custom API discovery settings. See [Api Discovery Settings Choice Custom Api Auth Discovery ](#api-discovery-settings-choice-custom-api-auth-discovery) below for details.(Deprecated)

`default_api_auth_discovery` - (Optional) Apply system default API discovery settings (`Bool`).(Deprecated)

`discovered_api_settings` - (Optional) Configure Discovered API Settings.. See [Api Discovery On Cache Miss Discovered Api Settings ](#api-discovery-on-cache-miss-discovered-api-settings) below for details.

###### One of the arguments from this list "disable_learn_from_redirect_traffic, enable_learn_from_redirect_traffic" must be set

`disable_learn_from_redirect_traffic` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`enable_learn_from_redirect_traffic` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`sensitive_data_detection_rules` - (Optional) Manage rules to detect sensitive data in requests and/or response sections.. See [Api Discovery On Cache Miss Sensitive Data Detection Rules ](#api-discovery-on-cache-miss-sensitive-data-detection-rules) below for details.(Deprecated)

### Api Discovery Choice Disable Api Discovery

Disable api discovery for this distribution.

### Api Discovery Choice Enable Api Discovery

Enable api discovery for all requests in this distribution.

`api_crawler` - (Optional) Configure Discovered API Settings.. See [Enable Api Discovery Api Crawler ](#enable-api-discovery-api-crawler) below for details.

`api_discovery_from_code_scan` - (Optional) Select API code repositories to the load balancer to use them as a source for API endpoint discovery.. See [Enable Api Discovery Api Discovery From Code Scan ](#enable-api-discovery-api-discovery-from-code-scan) below for details.

###### One of the arguments from this list "custom_api_auth_discovery, default_api_auth_discovery" must be set

`custom_api_auth_discovery` - (Optional) Apply custom API discovery settings. See [Api Discovery Settings Choice Custom Api Auth Discovery ](#api-discovery-settings-choice-custom-api-auth-discovery) below for details.(Deprecated)

`default_api_auth_discovery` - (Optional) Apply system default API discovery settings (`Bool`).(Deprecated)

`discovered_api_settings` - (Optional) Configure Discovered API Settings.. See [Enable Api Discovery Discovered Api Settings ](#enable-api-discovery-discovered-api-settings) below for details.

###### One of the arguments from this list "disable_learn_from_redirect_traffic, enable_learn_from_redirect_traffic" must be set

`disable_learn_from_redirect_traffic` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`enable_learn_from_redirect_traffic` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`sensitive_data_detection_rules` - (Optional) Manage rules to detect sensitive data in requests and/or response sections.. See [Enable Api Discovery Sensitive Data Detection Rules ](#enable-api-discovery-sensitive-data-detection-rules) below for details.(Deprecated)

### Api Discovery From Code Scan Code Base Integrations

x-required.

###### One of the arguments from this list "all_repos, selected_repos" must be set

`all_repos` - (Optional) x-displayName: "All API Repositories" (`Bool`).

`selected_repos` - (Optional) x-displayName: "Selected API Repositories". See [Api Repos Choice Selected Repos ](#api-repos-choice-selected-repos) below for details.

`code_base_integration` - (Required) Select the code base integration for use in code-based API discovery. See [ref](#ref) below for details.

### Api Discovery On Cache Miss Api Crawler

Configure Discovered API Settings..

###### One of the arguments from this list "api_crawler_config, disable_api_crawler" must be set

`api_crawler_config` - (Optional) Select to activate the API Crawling. See [Api Crawler Api Crawler Config ](#api-crawler-api-crawler-config) below for details.

`disable_api_crawler` - (Optional) Select to turn off the API Crawling. No API Crawling actions will be performed. (`Bool`).

### Api Discovery On Cache Miss Api Discovery From Code Scan

Select API code repositories to the load balancer to use them as a source for API endpoint discovery..

`code_base_integrations` - (Required) x-required. See [Api Discovery From Code Scan Code Base Integrations ](#api-discovery-from-code-scan-code-base-integrations) below for details.

### Api Discovery On Cache Miss Discovered Api Settings

Configure Discovered API Settings..

### Api Discovery On Cache Miss Sensitive Data Detection Rules

Manage rules to detect sensitive data in requests and/or response sections..

### Api Discovery Settings Choice Custom Api Auth Discovery

Apply custom API discovery settings.

`api_discovery_ref` - (Required) API Discovery Settings Object. See [ref](#ref) below for details.

### Api Discovery Settings Choice Default Api Auth Discovery

Apply system default API discovery settings.

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

### Api Repos Choice All Repos

x-displayName: "All API Repositories".

### Api Repos Choice Selected Repos

x-displayName: "Selected API Repositories".

`api_code_repo` - (Required) Code repository which contain API endpoints (`String`).

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

### App Traffic Type Choice Mobile

Mobile traffic channel..

### App Traffic Type Choice Mobile Client

Mobile traffic channel..

### App Traffic Type Choice Web

Web traffic channel..

### App Traffic Type Choice Web Client

Web traffic channel..

### App Traffic Type Choice Web Mobile

Web and mobile traffic channel..

`header` - (Optional) Header that is used by mobile traffic.. See [Web Mobile Header ](#web-mobile-header) below for details.(Deprecated)

`headers` - (Optional) Headers that can be used to identify mobile traffic.. See [Web Mobile Headers ](#web-mobile-headers) below for details.(Deprecated)

`mobile_identifier` - (Optional) Mobile identifier type (`String`).

### App Traffic Type Choice Web Mobile Client

Web and mobile traffic channel..

`header` - (Optional) Header that is used by mobile traffic.. See [Web Mobile Client Header ](#web-mobile-client-header) below for details.(Deprecated)

`headers` - (Optional) Headers that can be used to identify mobile traffic.. See [Web Mobile Client Headers ](#web-mobile-client-headers) below for details.(Deprecated)

`mobile_identifier` - (Optional) Mobile identifier type (`String`).

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

###### One of the arguments from this list "bearer_token, cookie, header, query_param" can be set

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

### Bot Defense Policy

Bot Defense Policy..

###### One of the arguments from this list "disable_js_insert, js_insert_all_pages, js_insert_all_pages_except, js_insertion_rules" must be set

`disable_js_insert` - (Optional) Disable JavaScript insertion. (`Bool`).

`js_insert_all_pages` - (Optional) Insert Bot Defense JavaScript in all pages.. See [Java Script Choice Js Insert All Pages ](#java-script-choice-js-insert-all-pages) below for details.

`js_insert_all_pages_except` - (Optional) Insert Bot Defense JavaScript in all pages with the exceptions.. See [Java Script Choice Js Insert All Pages Except ](#java-script-choice-js-insert-all-pages-except) below for details.

`js_insertion_rules` - (Optional) Specify custom JavaScript insertion rules.. See [Java Script Choice Js Insertion Rules ](#java-script-choice-js-insertion-rules) below for details.

`javascript_mode` - (Required) The larger chunk can be loaded asynchronously or synchronously. It can also be cacheable or non-cacheable on the browser. (`String`).

`js_download_path` - (Optional) Customize Bot Defense Client JavaScript path. If not specified, default `/common.js` (`String`).

###### One of the arguments from this list "disable_mobile_sdk, mobile_sdk_config" must be set

`disable_mobile_sdk` - (Optional) Disable Mobile SDK. (`Bool`).

`mobile_sdk_config` - (Optional) Mobile SDK configuration. See [Mobile Sdk Choice Mobile Sdk Config ](#mobile-sdk-choice-mobile-sdk-config) below for details.

`protected_app_endpoints` - (Required) List of protected application endpoints (max 128 items).. See [Policy Protected App Endpoints ](#policy-protected-app-endpoints) below for details.

### Bot Defense Advanced Policy

Bot Defense Advanced Policy..

`js_download_path` - (Required) Customize Bot Defense Web Client JavaScript path (`String`).

###### One of the arguments from this list "disable_mobile_sdk, mobile_sdk_config" must be set

`disable_mobile_sdk` - (Optional) Disable Mobile SDK. (`Bool`).

`mobile_sdk_config` - (Optional) Enable Mobile SDK Configuration. See [Mobile Sdk Choice Mobile Sdk Config ](#mobile-sdk-choice-mobile-sdk-config) below for details.

`protected_app_endpoints` - (Required) List of protected endpoints (max 128 items). See [Policy Protected App Endpoints ](#policy-protected-app-endpoints) below for details.

### Bot Defense Choice Bot Defense

Select Bot Defense Standard.

###### One of the arguments from this list "disable_cors_support, enable_cors_support" must be set

`disable_cors_support` - (Optional) protect against Bot Attacks. (`Bool`).(Deprecated)

`enable_cors_support` - (Optional) Allows Bot Defense to work with your existing CORS policies. (`Bool`).(Deprecated)

`policy` - (Required) Bot Defense Policy.. See [Bot Defense Policy ](#bot-defense-policy) below for details.

`regional_endpoint` - (Required) x-required (`String`).

`timeout` - (Optional) The timeout for the inference check, in milliseconds. (`Int`).

### Bot Defense Choice Bot Defense Advanced

Select Bot Defense Advanced.

`mobile` - (Optional) Select infrastructure for mobile.. See [ref](#ref) below for details.

`policy` - (Required) Bot Defense Advanced Policy.. See [Bot Defense Advanced Policy ](#bot-defense-advanced-policy) below for details.

`web` - (Optional) Select infrastructure for web.. See [ref](#ref) below for details.

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

### Cache Actions Cache Bypass

Bypass Caching of content from the origin.

### Cache Actions Cache Disabled

Disable Caching of content from the origin.

### Cache Actions Eligible For Cache

Eligible for caching the content.

###### One of the arguments from this list "hostname_uri, scheme_hostname_request_uri, scheme_hostname_uri, scheme_hostname_uri_query, scheme_proxy_host_request_uri, scheme_proxy_host_uri" must be set

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

###### One of the arguments from this list "Contains, DoesNotContain, DoesNotEndWith, DoesNotEqual, DoesNotStartWith, Endswith, Equals, MatchRegex, Startswith" can be set

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

###### One of the arguments from this list "cache_disabled, cache_ttl_default, cache_ttl_override, eligible_for_cache" can be set

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

###### One of the arguments from this list "captcha_challenge_parameters, default_captcha_challenge_parameters" can be set

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

### Challenge Type No Challenge

No challenge is enabled for this load balancer.

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

### Client Side Defense Policy

Please ensure that the same domains are configured in the Client-Side Defense configuration..

###### One of the arguments from this list "disable_js_insert, js_insert_all_pages, js_insert_all_pages_except, js_insertion_rules" must be set

`disable_js_insert` - (Optional) Disable JavaScript insertion. (`Bool`).

`js_insert_all_pages` - (Optional) Insert Client-Side Defense JavaScript in all pages. (`Bool`).

`js_insert_all_pages_except` - (Optional) Insert Client-Side Defense JavaScript in all pages with the exceptions.. See [Java Script Choice Js Insert All Pages Except ](#java-script-choice-js-insert-all-pages-except) below for details.

`js_insertion_rules` - (Optional) Specify custom JavaScript insertion rules.. See [Java Script Choice Js Insertion Rules ](#java-script-choice-js-insertion-rules) below for details.

### Client Side Defense Choice Client Side Defense

Client-Side Defense configuration for JavaScript insertion.

`policy` - (Required) Please ensure that the same domains are configured in the Client-Side Defense configuration.. See [Client Side Defense Policy ](#client-side-defense-policy) below for details.

### Client Source Choice Http Header

Request header name and value pairs.

`headers` - (Required) List of HTTP header name and value pairs. See [Http Header Headers ](#http-header-headers) below for details.

### Common Security Controls Blocked Clients

Define rules to block IP Prefixes or AS numbers..

###### One of the arguments from this list "bot_skip_processing, skip_processing, waf_skip_processing" can be set

`bot_skip_processing` - (Optional) Skip Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`skip_processing` - (Optional) Skip both WAF and Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (`Bool`).(Deprecated)

`actions` - (Optional) Actions that should be taken when client identifier matches the rule (`List of Strings`).

###### One of the arguments from this list "as_number, http_header, ip_prefix, ipv6_prefix, user_identifier" must be set

`as_number` - (Optional) RFC 6793 defined 4-byte AS number (`Int`).

`http_header` - (Optional) Request header name and value pairs. See [Client Source Choice Http Header ](#client-source-choice-http-header) below for details.

`ip_prefix` - (Optional) IPv4 prefix string. (`String`).

`ipv6_prefix` - (Optional) IPv6 prefix string. (`String`).

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

###### One of the arguments from this list "bot_skip_processing, skip_processing, waf_skip_processing" can be set

`bot_skip_processing` - (Optional) Skip Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`skip_processing` - (Optional) Skip both WAF and Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (`Bool`).(Deprecated)

`actions` - (Optional) Actions that should be taken when client identifier matches the rule (`List of Strings`).

###### One of the arguments from this list "as_number, http_header, ip_prefix, ipv6_prefix, user_identifier" must be set

`as_number` - (Optional) RFC 6793 defined 4-byte AS number (`Int`).

`http_header` - (Optional) Request header name and value pairs. See [Client Source Choice Http Header ](#client-source-choice-http-header) below for details.

`ip_prefix` - (Optional) IPv4 prefix string. (`String`).

`ipv6_prefix` - (Optional) IPv6 prefix string. (`String`).

`user_identifier` - (Optional) Identify user based on user identifier. User identifier value needs to be copied from security event. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Trusted Clients Metadata ](#trusted-clients-metadata) below for details.

### Condition Type Choice Api Endpoint

The API endpoint (Path + Method) which this validation applies to.

`methods` - (Optional) Methods to be matched (`List of Strings`).

`path` - (Required) Path to be matched (`String`).

### Cookie Matcher Operator

.

###### One of the arguments from this list "Contains, DoesNotContain, DoesNotEndWith, DoesNotEqual, DoesNotStartWith, Endswith, Equals, MatchRegex, Startswith" can be set

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

### Cors Support Choice Disable Cors Support

protect against Bot Attacks..

### Cors Support Choice Enable Cors Support

Allows Bot Defense to work with your existing CORS policies..

### Count By Choice Use Http Lb User Id

Defined in HTTP-LB Security Configuration -> User Identifier..

### Crl Choice No Crl

Client certificate revocation status is not verified.

### Data Guard Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

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

### Ddos Client Source Tls Fingerprint Matcher

The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints..

`classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`exact_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

`excluded_values` - (Optional) or more known TLS fingerprint classes in the enclosing matcher. (`String`).

### Ddos Mitigation Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

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

The rule will apply for all domains..

### Domain Matcher Choice Any Domain

Any Domain..

### Domain Matcher Choice Domain

Domain matcher..

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Domains Simple Login

Enter the username and password to assign credentials for the selected domain to crawl.

`password` - (Required) x-required. See [Simple Login Password ](#simple-login-password) below for details.

`user` - (Required) x-required (`String`).

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

### Enable Api Discovery Api Crawler

Configure Discovered API Settings..

###### One of the arguments from this list "api_crawler_config, disable_api_crawler" must be set

`api_crawler_config` - (Optional) Select to activate the API Crawling. See [Api Crawler Api Crawler Config ](#api-crawler-api-crawler-config) below for details.

`disable_api_crawler` - (Optional) Select to turn off the API Crawling. No API Crawling actions will be performed. (`Bool`).

### Enable Api Discovery Api Discovery From Code Scan

Select API code repositories to the load balancer to use them as a source for API endpoint discovery..

`code_base_integrations` - (Required) x-required. See [Api Discovery From Code Scan Code Base Integrations ](#api-discovery-from-code-scan-code-base-integrations) below for details.

### Enable Api Discovery Discovered Api Settings

Configure Discovered API Settings..

### Enable Api Discovery Sensitive Data Detection Rules

Manage rules to detect sensitive data in requests and/or response sections..

### Exclude List Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Exclude List Path

URI path matcher..

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Fail Configuration Fail Close

Handle the transaction as it failed the OpenAPI specification validation (Block or Report).

### Fail Configuration Fail Open

Continue to process the transaction without enforcing OpenAPI specification (Allow).

### Fall Through Mode Choice Fall Through Mode Allow

Allow any unprotected end point.

### Fall Through Mode Choice Fall Through Mode Custom

Custom rules for any unprotected end point.

`open_api_validation_rules` - (Required) x-displayName: "Custom Fall Through Rule List". See [Fall Through Mode Custom Open Api Validation Rules ](#fall-through-mode-custom-open-api-validation-rules) below for details.

### Fall Through Mode Custom Open Api Validation Rules

x-displayName: "Custom Fall Through Rule List".

###### One of the arguments from this list "action_block, action_report, action_skip" must be set

`action_block` - (Optional) Block the request and issue an API security event (`Bool`).

`action_report` - (Optional) Continue processing the request and issue an API security event (`Bool`).

`action_skip` - (Optional) Continue processing the request (`Bool`).

###### One of the arguments from this list "api_endpoint, api_group, base_path" must be set

`api_endpoint` - (Optional) The API endpoint (Path + Method) which this validation applies to. See [Condition Type Choice Api Endpoint ](#condition-type-choice-api-endpoint) below for details.

`api_group` - (Optional) The API group which this validation applies to (`String`).

`base_path` - (Optional) The base path which this validation applies to (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Open Api Validation Rules Metadata ](#open-api-validation-rules-metadata) below for details.

### Flow Label Choice Account Management

x-displayName: "Account Management".

###### One of the arguments from this list "create, password_reset" must be set

`create` - (Optional) x-displayName: "Account Creation" (`Bool`).

`password_reset` - (Optional) x-displayName: "Password Reset" (`Bool`).

### Flow Label Choice Authentication

x-displayName: "Authentication".

###### One of the arguments from this list "login, login_mfa, login_partner, logout, token_refresh" must be set

`login` - (Optional) x-displayName: "Login". See [Label Choice Login ](#label-choice-login) below for details.

`login_mfa` - (Optional) x-displayName: "Login MFA" (`Bool`).

`login_partner` - (Optional) x-displayName: "Login for a Channel Partner" (`Bool`).

`logout` - (Optional) x-displayName: "Logout" (`Bool`).

`token_refresh` - (Optional) x-displayName: "Token Refresh" (`Bool`).

### Flow Label Choice Financial Services

x-displayName: "Financial Services".

###### One of the arguments from this list "apply, money_transfer" must be set

`apply` - (Optional) x-displayName: "Apply for a Financial Service Account (e.g., credit card, banking, retirement account)" (`Bool`).

`money_transfer` - (Optional) x-displayName: "Money Transfer" (`Bool`).

### Flow Label Choice Flight

x-displayName: "Flight".

###### One of the arguments from this list "checkin" must be set

`checkin` - (Optional) x-displayName: "Check into Flight" (`Bool`).

### Flow Label Choice Flow Label

x-displayName: "Specify Endpoint label category".

###### One of the arguments from this list "account_management, authentication, financial_services, flight, profile_management, search, shopping_gift_cards" must be set

`account_management` - (Optional) x-displayName: "Account Management". See [Flow Label Choice Account Management ](#flow-label-choice-account-management) below for details.

`authentication` - (Optional) x-displayName: "Authentication". See [Flow Label Choice Authentication ](#flow-label-choice-authentication) below for details.

`financial_services` - (Optional) x-displayName: "Financial Services". See [Flow Label Choice Financial Services ](#flow-label-choice-financial-services) below for details.

`flight` - (Optional) x-displayName: "Flight". See [Flow Label Choice Flight ](#flow-label-choice-flight) below for details.

`profile_management` - (Optional) x-displayName: "Profile Management". See [Flow Label Choice Profile Management ](#flow-label-choice-profile-management) below for details.

`search` - (Optional) x-displayName: "Search". See [Flow Label Choice Search ](#flow-label-choice-search) below for details.

`shopping_gift_cards` - (Optional) x-displayName: "Shopping & Gift Cards". See [Flow Label Choice Shopping Gift Cards ](#flow-label-choice-shopping-gift-cards) below for details.

### Flow Label Choice Profile Management

x-displayName: "Profile Management".

###### One of the arguments from this list "create, update, view" must be set

`create` - (Optional) x-displayName: "Profile Creation" (`Bool`).

`update` - (Optional) x-displayName: "Profile Update" (`Bool`).

`view` - (Optional) x-displayName: "Profile View" (`Bool`).

### Flow Label Choice Search

x-displayName: "Search".

###### One of the arguments from this list "flight_search, product_search, reservation_search, room_search" can be set

`flight_search` - (Optional) x-displayName: "Flight Search" (`Bool`).

`product_search` - (Optional) x-displayName: "Product Search" (`Bool`).

`reservation_search` - (Optional) x-displayName: "Reservation Search (e.g., sporting events, concerts)" (`Bool`).

`room_search` - (Optional) x-displayName: "Room Search" (`Bool`).

### Flow Label Choice Shopping Gift Cards

x-displayName: "Shopping & Gift Cards".

###### One of the arguments from this list "gift_card_make_purchase_with_gift_card, gift_card_validation, shop_add_to_cart, shop_checkout, shop_choose_seat, shop_enter_drawing_submission, shop_make_payment, shop_order, shop_price_inquiry, shop_promo_code_validation, shop_purchase_gift_card, shop_update_quantity" can be set

`gift_card_make_purchase_with_gift_card` - (Optional) x-displayName: "Purchase with Gift Card" (`Bool`).

`gift_card_validation` - (Optional) x-displayName: "Gift Card Validation" (`Bool`).

`shop_add_to_cart` - (Optional) x-displayName: "Add to Cart" (`Bool`).

`shop_checkout` - (Optional) x-displayName: "Checkout" (`Bool`).

`shop_choose_seat` - (Optional) x-displayName: "Select Seat(s)" (`Bool`).

`shop_enter_drawing_submission` - (Optional) x-displayName: "Enter Drawing Submission" (`Bool`).

`shop_make_payment` - (Optional) x-displayName: "Payment / Billing" (`Bool`).

`shop_order` - (Optional) x-displayName: "Order Submit" (`Bool`).

`shop_price_inquiry` - (Optional) x-displayName: "Price Inquiry" (`Bool`).

`shop_promo_code_validation` - (Optional) x-displayName: "Promo Code Validation" (`Bool`).

`shop_purchase_gift_card` - (Optional) x-displayName: "Purchase a Gift Card" (`Bool`).

`shop_update_quantity` - (Optional) x-displayName: "Update Quantity" (`Bool`).

### Flow Label Choice Undefined Flow Label

x-displayName: "Undefined".

### Geo Filtering Type Allow List

Allow list of countries.

`country_codes` - (Required) List of Country Codes (`List of Strings`).

`invert_match` - (Optional) Invert the match result. (`Bool`).

### Geo Filtering Type Block List

Block list of countries.

`country_codes` - (Required) List of Country Codes (`List of Strings`).

`invert_match` - (Optional) Invert the match result. (`Bool`).

### Goodbot Choice Allow Good Bots

System flags Good Bot traffic and allow it to continue to the origin.

### Goodbot Choice Mitigate Good Bots

System flags Good Bot Traffic, but mitigation is handled in the same manner as malicious automated traffic defined above.

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

### Https Tls Parameters

TLS parameters for the downstream connections..

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Parameters Tls Certificates ](#tls-parameters-tls-certificates) below for details.

`tls_config` - (Optional) TLS Configuration Parameters. See [Tls Parameters Tls Config ](#tls-parameters-tls-config) below for details.

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

### Ip Filtering Type Allow List

Allow list of ip prefixes.

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

### Ip Filtering Type Block List

Block list of ip prefixes.

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

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

### Java Script Choice Disable Js Insert

Disable JavaScript insertion..

### Java Script Choice Js Insert All Pages

Insert Client-Side Defense JavaScript in all pages..

### Java Script Choice Js Insert All Pages

Insert Bot Defense JavaScript in all pages..

`javascript_location` - (Optional) Defines where to insert Bot Defense JavaScript in HTML page. (`String`).

### Java Script Choice Js Insert All Pages Except

Insert Client-Side Defense JavaScript in all pages with the exceptions..

`exclude_list` - (Optional) Optional JavaScript insertions exclude list of domain and path matchers.. See [Js Insert All Pages Except Exclude List ](#js-insert-all-pages-except-exclude-list) below for details.

### Java Script Choice Js Insert All Pages Except

Insert Bot Defense JavaScript in all pages with the exceptions..

`exclude_list` - (Optional) Optional JavaScript insertions exclude list of domain and path matchers.. See [Js Insert All Pages Except Exclude List ](#js-insert-all-pages-except-exclude-list) below for details.

`javascript_location` - (Optional) Defines where to insert Bot Defense JavaScript in HTML page. (`String`).

### Java Script Choice Js Insertion Rules

Specify custom JavaScript insertion rules..

`exclude_list` - (Optional) Optional JavaScript insertions exclude list of domain and path matchers.. See [Js Insertion Rules Exclude List ](#js-insertion-rules-exclude-list) below for details.

`rules` - (Required) Required list of pages to insert Client-Side Defense client JavaScript.. See [Js Insertion Rules Rules ](#js-insertion-rules-rules) below for details.

### Java Script Choice Js Insertion Rules

Specify custom JavaScript insertion rules..

`exclude_list` - (Optional) Optional JavaScript insertions exclude list of domain and path matchers.. See [Js Insertion Rules Exclude List ](#js-insertion-rules-exclude-list) below for details.

`rules` - (Required) Required list of pages to insert Bot Defense client JavaScript.. See [Js Insertion Rules Rules ](#js-insertion-rules-rules) below for details.

### Js Challenge Parameters Choice Default Js Challenge Parameters

Use default parameters.

### Js Challenge Parameters Choice Js Challenge Parameters

Configure JavaScript challenge parameters.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

`js_script_delay` - (Optional) Delay introduced by Javascript, in milliseconds. (`Int`).

### Js Insert All Pages Except Exclude List

Optional JavaScript insertions exclude list of domain and path matchers..

###### One of the arguments from this list "any_domain, domain" must be set

`any_domain` - (Optional) Any Domain. (`Bool`).

`domain` - (Optional) Domain matcher.. See [Domain Matcher Choice Domain ](#domain-matcher-choice-domain) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Exclude List Metadata ](#exclude-list-metadata) below for details.

`path` - (Required) URI path matcher.. See [Exclude List Path ](#exclude-list-path) below for details.

### Js Insertion Rules Exclude List

Optional JavaScript insertions exclude list of domain and path matchers..

###### One of the arguments from this list "any_domain, domain" must be set

`any_domain` - (Optional) Any Domain. (`Bool`).

`domain` - (Optional) Domain matcher.. See [Domain Matcher Choice Domain ](#domain-matcher-choice-domain) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Exclude List Metadata ](#exclude-list-metadata) below for details.

`path` - (Required) URI path matcher.. See [Exclude List Path ](#exclude-list-path) below for details.

### Js Insertion Rules Rules

Required list of pages to insert Client-Side Defense client JavaScript..

###### One of the arguments from this list "any_domain, domain" must be set

`any_domain` - (Optional) Any Domain. (`Bool`).

`domain` - (Optional) Domain matcher.. See [Domain Matcher Choice Domain ](#domain-matcher-choice-domain) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Rules Metadata ](#rules-metadata) below for details.

`path` - (Required) URI path matcher.. See [Rules Path ](#rules-path) below for details.

### Js Insertion Rules Rules

Required list of pages to insert Bot Defense client JavaScript..

###### One of the arguments from this list "any_domain, domain" must be set

`any_domain` - (Optional) Any Domain. (`Bool`).

`domain` - (Optional) Domain matcher.. See [Domain Matcher Choice Domain ](#domain-matcher-choice-domain) below for details.

`javascript_location` - (Optional) Defines where to insert Bot Defense JavaScript in HTML page. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Rules Metadata ](#rules-metadata) below for details.

`path` - (Required) URI path matcher.. See [Rules Path ](#rules-path) below for details.

### Jwks Configuration Jwks Config

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details..

`cleartext` - (Optional) The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details. (`String`).

### Jwt Backup Key

Backup JWT Key - If specified is also checked in addition to the primary secret key.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Backup Key Blindfold Secret Info Internal ](#backup-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Jwt Secret Key

Secret Key for JWT.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Secret Key Blindfold Secret Info Internal ](#secret-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

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

###### One of the arguments from this list "bearer_token, cookie, header, query_param" must be set

`bearer_token` - (Optional) Token is found in Authorization HTTP header with Bearer authentication scheme (`Bool`).

`cookie` - (Optional) Token is found in the cookie (`String`).(Deprecated)

`header` - (Optional) Token is found in the header (`String`).(Deprecated)

`query_param` - (Optional) Token is found in the query string parameter (`String`).(Deprecated)

### L7 Ddos Auto Mitigation Action L7 Ddos Action Js Challenge

Serve JavaScript challenge to suspicious sources.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

`js_script_delay` - (Optional) Delay introduced by Javascript, in milliseconds. (`Int`).

### Label Choice Apply

x-displayName: "Apply for a Financial Service Account (e.g., credit card, banking, retirement account)".

### Label Choice Checkin

x-displayName: "Check into Flight".

### Label Choice Create

x-displayName: "Account Creation".

### Label Choice Flight Search

x-displayName: "Flight Search".

### Label Choice Gift Card Make Purchase With Gift Card

x-displayName: "Purchase with Gift Card".

### Label Choice Gift Card Validation

x-displayName: "Gift Card Validation".

### Label Choice Login

x-displayName: "Login".

### Label Choice Login Mfa

x-displayName: "Login MFA".

### Label Choice Login Partner

x-displayName: "Login for a Channel Partner".

### Label Choice Logout

x-displayName: "Logout".

### Label Choice Money Transfer

x-displayName: "Money Transfer".

### Label Choice Password Reset

x-displayName: "Password Reset".

### Label Choice Product Search

x-displayName: "Product Search".

### Label Choice Reservation Search

x-displayName: "Reservation Search (e.g., sporting events, concerts)".

### Label Choice Room Search

x-displayName: "Room Search".

### Label Choice Shop Add To Cart

x-displayName: "Add to Cart".

### Label Choice Shop Checkout

x-displayName: "Checkout".

### Label Choice Shop Choose Seat

x-displayName: "Select Seat(s)".

### Label Choice Shop Enter Drawing Submission

x-displayName: "Enter Drawing Submission".

### Label Choice Shop Make Payment

x-displayName: "Payment / Billing".

### Label Choice Shop Order

x-displayName: "Order Submit".

### Label Choice Shop Price Inquiry

x-displayName: "Price Inquiry".

### Label Choice Shop Promo Code Validation

x-displayName: "Promo Code Validation".

### Label Choice Shop Purchase Gift Card

x-displayName: "Purchase a Gift Card".

### Label Choice Shop Update Quantity

x-displayName: "Update Quantity".

### Label Choice Token Refresh

x-displayName: "Token Refresh".

### Label Choice Update

x-displayName: "Profile Update".

### Label Choice View

x-displayName: "Profile View".

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

`tls_cert_options` - (Optional) TLS Certificate Options. See [Https Tls Cert Options ](#https-tls-cert-options) below for details.

`tls_parameters` - (Optional) TLS parameters for the downstream connections.. See [Https Tls Parameters ](#https-tls-parameters) below for details.(Deprecated)

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

`tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints.. See [Ddos Client Source Tls Fingerprint Matcher ](#ddos-client-source-tls-fingerprint-matcher) below for details.

### Mitigation Choice Ip Prefix List

IP prefix string..

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

### Mobile Identifier Headers

Headers that can be used to identify mobile traffic..

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Mobile Sdk Choice Disable Mobile Sdk

Disable Mobile SDK..

### Mobile Sdk Choice Mobile Sdk Config

Enable Mobile SDK Configuration.

`mobile_identifier` - (Optional) Mobile Request Identifier Headers Type.. See [Mobile Sdk Config Mobile Identifier ](#mobile-sdk-config-mobile-identifier) below for details.

### Mobile Sdk Choice Mobile Sdk Config

Mobile SDK configuration.

`mobile_identifier` - (Optional) Mobile traffic identifier type.. See [Mobile Sdk Config Mobile Identifier ](#mobile-sdk-config-mobile-identifier) below for details.

`reload_header_name` - (Optional) Header that is used for SDK configuration sync. (`String`).(Deprecated)

### Mobile Sdk Config Mobile Identifier

Mobile traffic identifier type..

`headers` - (Optional) Headers that can be used to identify mobile traffic.. See [Mobile Identifier Headers ](#mobile-identifier-headers) below for details.

### More Option Cache Options

Cache Options.

`cache_rules` - (Optional) Rules are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs.. See [Cache Options Cache Rules ](#cache-options-cache-rules) below for details.

`default_cache_action` - (Required) Default value for Cache action.. See [Cache Options Default Cache Action ](#cache-options-default-cache-action) below for details.

### More Option Cache Ttl Options

Cache Options.

###### One of the arguments from this list "cache_disabled, cache_ttl_default, cache_ttl_override" can be set

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

###### One of the arguments from this list "skip_validation, validation_mode_active" must be set

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

### Other Settings Geo Filtering

Geo filtering options.

###### One of the arguments from this list "allow_list, block_list" can be set

`allow_list` - (Optional) Allow list of countries. See [Geo Filtering Type Allow List ](#geo-filtering-type-allow-list) below for details.

`block_list` - (Optional) Block list of countries. See [Geo Filtering Type Block List ](#geo-filtering-type-block-list) below for details.

### Other Settings Header Options

Request/Response header related options.

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Header Options Request Headers To Add ](#header-options-request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Header Options Response Headers To Add ](#header-options-response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

### Other Settings Ip Filtering

IP filtering options.

###### One of the arguments from this list "allow_list, block_list" can be set

`allow_list` - (Optional) Allow list of ip prefixes. See [Ip Filtering Type Allow List ](#ip-filtering-type-allow-list) below for details.

`block_list` - (Optional) Block list of ip prefixes. See [Ip Filtering Type Block List ](#ip-filtering-type-block-list) below for details.

### Other Settings Logging Options

Logging related options.

`client_log_options` - (Optional) Client request headers to log. See [Logging Options Client Log Options ](#logging-options-client-log-options) below for details.

`origin_log_options` - (Optional) Origin response headers to log. See [Logging Options Origin Log Options ](#logging-options-origin-log-options) below for details.

### Oversized Body Choice Oversized Body Fail Validation

Apply the request/response action (block or report) when the body length is too long to verify (default 64Kb).

### Oversized Body Choice Oversized Body Skip Validation

Skip body validation when the body length is too long to verify (default 64Kb).

### Password Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Path Choice Any Path

Match all paths.

### Path Match Operator

A specification of path match.

###### One of the arguments from this list "Contains, DoesNotContain, DoesNotEndWith, DoesNotEqual, DoesNotStartWith, Endswith, Equals, MatchRegex, Startswith" can be set

`Contains` - (Optional) Field must contain (`String`).

`DoesNotContain` - (Optional) Field must not contain (`String`).

`DoesNotEndWith` - (Optional) Field must not end with (`String`).

`DoesNotEqual` - (Optional) Field must not equal (`String`).

`DoesNotStartWith` - (Optional) Field must not start with (`String`).

`Endswith` - (Optional) Field must end with (`String`).

`Equals` - (Optional) Field must exactly match (`String`).

`MatchRegex` - (Optional) Field matches regular expression (`String`).

`Startswith` - (Optional) Field must start with (`String`).

### Policy Protected App Endpoints

List of protected application endpoints (max 128 items)..

###### One of the arguments from this list "mobile, web, web_mobile" must be set

`mobile` - (Optional) Mobile traffic channel. (`Bool`).

`web` - (Optional) Web traffic channel. (`Bool`).

`web_mobile` - (Optional) Web and mobile traffic channel.. See [App Traffic Type Choice Web Mobile ](#app-traffic-type-choice-web-mobile) below for details.

###### One of the arguments from this list "any_domain, domain" can be set

`any_domain` - (Optional) Any Domain. (`Bool`).

`domain` - (Optional) Domain matcher.. See [Domain Matcher Choice Domain ](#domain-matcher-choice-domain) below for details.

###### One of the arguments from this list "flow_label, undefined_flow_label" must be set

`flow_label` - (Optional) x-displayName: "Specify Endpoint label category". See [Flow Label Choice Flow Label ](#flow-label-choice-flow-label) below for details.

`undefined_flow_label` - (Optional) x-displayName: "Undefined" (`Bool`).

###### One of the arguments from this list "allow_good_bots, mitigate_good_bots" must be set

`allow_good_bots` - (Optional) System flags Good Bot traffic and allow it to continue to the origin (`Bool`).

`mitigate_good_bots` - (Optional) System flags Good Bot Traffic, but mitigation is handled in the same manner as malicious automated traffic defined above (`Bool`).

`headers` - (Optional) Note that all specified header predicates must evaluate to true.. See [Protected App Endpoints Headers ](#protected-app-endpoints-headers) below for details.

`http_methods` - (Required) List of HTTP methods. (`List of Strings`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Protected App Endpoints Metadata ](#protected-app-endpoints-metadata) below for details.

`mitigation` - (Required) Mitigation action.. See [Protected App Endpoints Mitigation ](#protected-app-endpoints-mitigation) below for details.

`path` - (Required) Matching URI path of the route.. See [Protected App Endpoints Path ](#protected-app-endpoints-path) below for details.

`protocol` - (Optional) Protocol. (`String`).

`query_params` - (Optional) Note that all specified query parameter predicates must evaluate to true.. See [Protected App Endpoints Query Params ](#protected-app-endpoints-query-params) below for details.

### Policy Protected App Endpoints

List of protected endpoints (max 128 items).

###### One of the arguments from this list "mobile_client, web_client, web_mobile_client" must be set

`mobile_client` - (Optional) Mobile traffic channel. (`Bool`).

`web_client` - (Optional) Web traffic channel. (`Bool`).

`web_mobile_client` - (Optional) Web and mobile traffic channel.. See [App Traffic Type Choice Web Mobile Client ](#app-traffic-type-choice-web-mobile-client) below for details.

###### One of the arguments from this list "any_domain, domain" can be set

`any_domain` - (Optional) Any Domain (`Bool`).

`domain` - (Optional) Select Domain matcher. See [Domain Matcher Choice Domain ](#domain-matcher-choice-domain) below for details.

###### One of the arguments from this list "flow_label, undefined_flow_label" can be set

`flow_label` - (Optional) x-displayName: "Specify endpoint label category". See [Flow Label Choice Flow Label ](#flow-label-choice-flow-label) below for details.

`undefined_flow_label` - (Optional) x-displayName: "Undefined" (`Bool`).

`http_methods` - (Required) List of HTTP methods. (`List of Strings`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Protected App Endpoints Metadata ](#protected-app-endpoints-metadata) below for details.

`path` - (Required) Accepts wildcards * to match multiple characters or ? to match a single character. See [Protected App Endpoints Path ](#protected-app-endpoints-path) below for details.

`query` - (Optional) Enter a regular expression or exact value to match your query parameters of interest. See [Protected App Endpoints Query ](#protected-app-endpoints-query) below for details.

`request_body` - (Optional) Request Body. See [Protected App Endpoints Request Body ](#protected-app-endpoints-request-body) below for details.

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

### Protected App Endpoints Headers

Note that all specified header predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

###### One of the arguments from this list "check_not_present, check_present, item, presence" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the header is present or absent. (`Bool`).(Deprecated)

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Protected App Endpoints Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Protected App Endpoints Mitigation

Mitigation action..

###### One of the arguments from this list "block, flag, none, redirect" can be set

`block` - (Optional) Block bot request and send response with custom content.. See [Action Type Block ](#action-type-block) below for details.

`flag` - (Optional) Flag the request while not taking any invasive actions.. See [Action Type Flag ](#action-type-flag) below for details.

`none` - (Optional) No mitigation actions. (`Bool`).(Deprecated)

`redirect` - (Optional) Redirect bot request to a custom URI.. See [Action Type Redirect ](#action-type-redirect) below for details.

### Protected App Endpoints Path

Matching URI path of the route..

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Protected App Endpoints Query

Enter a regular expression or exact value to match your query parameters of interest.

`name` - (Optional) Enter query parameter name (`String`).

###### One of the arguments from this list "check_presence, exact_value, regex_value" must be set

`check_presence` - (Optional) Parameter name taken which is exist in the query parameter (`Bool`).

`exact_value` - (Optional) Exact query value to match (`String`).

`regex_value` - (Optional) Regular expression of query match (e.g. the value .* will match on all query) (`String`).

### Protected App Endpoints Query Params

Note that all specified query parameter predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`key` - (Required) A case-sensitive HTTP query parameter name. (`String`).

###### One of the arguments from this list "check_not_present, check_present, item, presence" must be set

`check_not_present` - (Optional) Check that the query parameter is not present. (`Bool`).

`check_present` - (Optional) Check that the query parameter is present. (`Bool`).

`item` - (Optional) criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the query parameter is present or absent. (`Bool`).(Deprecated)

### Protected App Endpoints Request Body

Request Body.

`name` - (Optional) Enter request body parameter name (`String`).

###### One of the arguments from this list "exact_value, regex_value" must be set

`exact_value` - (Optional) Exact query value to match (`String`).

`regex_value` - (Optional) Regular expression of query match (e.g. the value .* will match on all query) (`String`).

### Query Parameters Operator

.

###### One of the arguments from this list "Contains, DoesNotContain, DoesNotEndWith, DoesNotEqual, DoesNotStartWith, Endswith, Equals, MatchRegex, Startswith" can be set

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

###### One of the arguments from this list "bypass_rate_limiting_rules, custom_ip_allowed_list, ip_allowed_list, no_ip_allowed_list" must be set

`bypass_rate_limiting_rules` - (Optional) This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting.. See [Ip Allowed List Choice Bypass Rate Limiting Rules ](#ip-allowed-list-choice-bypass-rate-limiting-rules) below for details.

`custom_ip_allowed_list` - (Optional) IP Allowed list using existing ip_prefix_set objects.. See [Ip Allowed List Choice Custom Ip Allowed List ](#ip-allowed-list-choice-custom-ip-allowed-list) below for details.

`ip_allowed_list` - (Optional) List of IP(s) for which rate limiting will be disabled.. See [Ip Allowed List Choice Ip Allowed List ](#ip-allowed-list-choice-ip-allowed-list) below for details.

`no_ip_allowed_list` - (Optional) There is no ip allowed list for rate limiting, all clients go through rate limiting. (`Bool`).

`server_url_rules` - (Optional) For matching also specific endpoints you can use the API endpoint rules set bellow.. See [Api Rate Limit Server Url Rules ](#api-rate-limit-server-url-rules) below for details.

### Rate Limit Choice Disable Rate Limit

Rate limiting is not currently enabled for this load balancer.

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

###### One of the arguments from this list "check_not_present, check_present, item, presence" must be set

`check_not_present` - (Optional) Check that the cookie is not present. (`Bool`).

`check_present` - (Optional) Check that the cookie is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the cookie. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the cookie is present or absent. (`Bool`).(Deprecated)

`name` - (Required) A case-sensitive cookie name. (`String`).

### Request Matcher Headers

Note that all specified header predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

###### One of the arguments from this list "check_not_present, check_present, item, presence" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the header is present or absent. (`Bool`).(Deprecated)

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

###### One of the arguments from this list "check_not_present, check_present, item, presence" must be set

`check_not_present` - (Optional) Check that the query parameter is not present. (`Bool`).

`check_present` - (Optional) Check that the query parameter is present. (`Bool`).

`item` - (Optional) criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the query parameter is present or absent. (`Bool`).(Deprecated)

### Request Timeout Choice Disable Request Timeout

x-displayName: "No Timeout".

### Response Validation Mode Choice Response Validation Mode Active

Enforce OpenAPI validation processing for this event.

`response_validation_properties` - (Required) List of properties of the response to validate according to the OpenAPI specification file (a.k.a. swagger) (`List of Strings`).

###### One of the arguments from this list "enforcement_block, enforcement_report" must be set

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

### Rules Path

URI path matcher..

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

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

###### One of the arguments from this list "any_client, client_name, client_name_matcher, client_selector" can be set

`any_client` - (Optional)any_client (`Bool`).

`client_name` - (Optional)client_name (`String`).(Deprecated)

`client_name_matcher` - (Optional)client_name_matcher. See [Client Choice Client Name Matcher ](#client-choice-client-name-matcher) below for details.(Deprecated)

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

###### One of the arguments from this list "ja4_tls_fingerprint, tls_fingerprint_matcher" can be set

`ja4_tls_fingerprint` - (Optional)ja4_tls_fingerprint. See [Tls Fingerprint Choice Ja4 Tls Fingerprint ](#tls-fingerprint-choice-ja4-tls-fingerprint) below for details.(Deprecated)

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

###### One of the arguments from this list "api_specification, api_specification_on_cache_miss, disable_api_definition" must be set

`api_specification` - (Optional) Specify API definition and OpenAPI Validation. See [Api Definition Choice Api Specification ](#api-definition-choice-api-specification) below for details.

`api_specification_on_cache_miss` - (Optional) Enable API definition and OpenAPI Validation only on cache miss in this distribution. See [Api Definition Choice Api Specification On Cache Miss ](#api-definition-choice-api-specification-on-cache-miss) below for details.(Deprecated)

`disable_api_definition` - (Optional) API Definition is not currently used for this load balancer (`Bool`).

###### One of the arguments from this list "api_discovery_on_cache_miss, disable_api_discovery, enable_api_discovery" must be set

`api_discovery_on_cache_miss` - (Optional) Enable api discovery only on cache miss in this distribution. See [Api Discovery Choice Api Discovery On Cache Miss ](#api-discovery-choice-api-discovery-on-cache-miss) below for details.(Deprecated)

`disable_api_discovery` - (Optional) Disable api discovery for this distribution (`Bool`).

`enable_api_discovery` - (Optional) Enable api discovery for all requests in this distribution. See [Api Discovery Choice Enable Api Discovery ](#api-discovery-choice-enable-api-discovery) below for details.

`api_protection_rules` - (Optional) Rules can also include additional conditions, for example specific clients can access certain API endpoint or API group.. See [Api Protection Api Protection Rules ](#api-protection-api-protection-rules) below for details.

`jwt_validation` - (Optional) tokens or tokens that are not yet valid.. See [Api Protection Jwt Validation ](#api-protection-jwt-validation) below for details.

###### One of the arguments from this list "default_sensitive_data_policy, sensitive_data_policy" must be set

`default_sensitive_data_policy` - (Optional) Apply system default sensitive data discovery (`Bool`).

`sensitive_data_policy` - (Optional) Apply custom sensitive data discovery. See [Sensitive Data Policy Choice Sensitive Data Policy ](#sensitive-data-policy-choice-sensitive-data-policy) below for details.

### Security Options Auth Options

Authentication Options.

###### One of the arguments from this list "custom, disable_auth, jwt" can be set

`custom` - (Optional) Enable Custom Authentication. See [Auth Options Custom ](#auth-options-custom) below for details.

`disable_auth` - (Optional) No Authentication (`Bool`).

`jwt` - (Optional) Enable JWT Authentication. See [Auth Options Jwt ](#auth-options-jwt) below for details.

### Security Options Common Security Controls

x-displayName: "Common Security Controls".

`blocked_clients` - (Optional) Define rules to block IP Prefixes or AS numbers.. See [Common Security Controls Blocked Clients ](#common-security-controls-blocked-clients) below for details.

###### One of the arguments from this list "captcha_challenge, challenge_on_cache_miss, enable_challenge, js_challenge, no_challenge, policy_based_challenge" must be set

`captcha_challenge` - (Optional) Configure Captcha challenge on this load balancer. See [Challenge Type Captcha Challenge ](#challenge-type-captcha-challenge) below for details.

`challenge_on_cache_miss` - (Optional) Configure auto mitigation i.e risk based challenges for malicious users only on cache miss in this load balancer. See [Challenge Type Challenge On Cache Miss ](#challenge-type-challenge-on-cache-miss) below for details.(Deprecated)

`enable_challenge` - (Optional) Configure auto mitigation i.e risk based challenges for malicious users for this load balancer. See [Challenge Type Enable Challenge ](#challenge-type-enable-challenge) below for details.

`js_challenge` - (Optional) Configure JavaScript challenge on this load balancer. See [Challenge Type Js Challenge ](#challenge-type-js-challenge) below for details.

`no_challenge` - (Optional) No challenge is enabled for this load balancer (`Bool`).

`policy_based_challenge` - (Optional) Specifies the settings for policy rule based challenge. See [Challenge Type Policy Based Challenge ](#challenge-type-policy-based-challenge) below for details.

`cors_policy` - (Optional) resources from a server at a different origin. See [Common Security Controls Cors Policy ](#common-security-controls-cors-policy) below for details.

###### One of the arguments from this list "disable_ip_reputation, enable_ip_reputation, ip_reputation_on_cache_miss" can be set

`disable_ip_reputation` - (Optional) No IP reputation configured this distribution (`Bool`).

`enable_ip_reputation` - (Optional) Enable IP reputation for all requests in this distribution. See [Ip Reputation Choice Enable Ip Reputation ](#ip-reputation-choice-enable-ip-reputation) below for details.

`ip_reputation_on_cache_miss` - (Optional) Enable IP reputation only on cache miss in this distribution. See [Ip Reputation Choice Ip Reputation On Cache Miss ](#ip-reputation-choice-ip-reputation-on-cache-miss) below for details.(Deprecated)

###### One of the arguments from this list "disable_malicious_user_detection, enable_malicious_user_detection, malicious_user_detection_on_cache_miss" must be set

`disable_malicious_user_detection` - (Optional) Disable malicious user detection for this distribution (`Bool`).

`enable_malicious_user_detection` - (Optional) Enable malicious user detection for all requests in this distribution (`Bool`).

`malicious_user_detection_on_cache_miss` - (Optional) Enable malicious user detection only on cache miss in this distribution (`Bool`).(Deprecated)

###### One of the arguments from this list "api_rate_limit, disable_rate_limit, rate_limit" must be set

`api_rate_limit` - (Optional) Define rate limiting for one or more API endpoints. See [Rate Limit Choice Api Rate Limit ](#rate-limit-choice-api-rate-limit) below for details.

`disable_rate_limit` - (Optional) Rate limiting is not currently enabled for this load balancer (`Bool`).

`rate_limit` - (Optional) Define custom rate limiting parameters for this load balancer. See [Rate Limit Choice Rate Limit ](#rate-limit-choice-rate-limit) below for details.

###### One of the arguments from this list "active_service_policies, no_service_policies, service_policies_from_namespace" must be set

`active_service_policies` - (Optional) Apply the specified list of service policies and bypass the namespace service policy set. See [Service Policy Choice Active Service Policies ](#service-policy-choice-active-service-policies) below for details.

`no_service_policies` - (Optional) Do not apply any service policies i.e. bypass the namespace service policy set (`Bool`).

`service_policies_from_namespace` - (Optional) Apply the active service policies configured as part of the namespace service policy set (`Bool`).

###### One of the arguments from this list "disable_threat_mesh, enable_threat_mesh" must be set

`disable_threat_mesh` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_threat_mesh` - (Optional) x-displayName: "Enable" (`Bool`).

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

###### One of the arguments from this list "app_firewall, app_firewall_on_cache_miss, disable_waf" must be set

`app_firewall` - (Optional) Enable WAF configuration for all requests in this distribution. See [ref](#ref) below for details.

`app_firewall_on_cache_miss` - (Optional) Enable WAF configuration only on cache miss in this distribution. See [ref](#ref) below for details.(Deprecated)

`disable_waf` - (Optional) No WAF configuration for this load balancer (`Bool`).

`waf_exclusion_rules` - (Optional) When an exclusion rule is matched, then this exclusion rule takes effect and no more rules are evaluated.. See [Web App Firewall Waf Exclusion Rules ](#web-app-firewall-waf-exclusion-rules) below for details.

### Send Headers Choice Append Headers

Append mitigation headers..

`auto_type_header_name` - (Required) A case-insensitive HTTP header name. (`String`).

`inference_header_name` - (Required) A case-insensitive HTTP header name. (`String`).

### Send Headers Choice No Headers

No mitigation headers..

### Sensitive Data Policy Choice Default Sensitive Data Policy

Apply system default sensitive data discovery.

### Sensitive Data Policy Choice Sensitive Data Policy

Apply custom sensitive data discovery.

`sensitive_data_policy_ref` - (Required) Specify Sensitive Data Discovery. See [ref](#ref) below for details.

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

### Service Policy Choice No Service Policies

Do not apply any service policies i.e. bypass the namespace service policy set.

### Service Policy Choice Service Policies From Namespace

Apply the active service policies configured as part of the namespace service policy set.

### Simple Login Password

x-required.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Password Blindfold Secret Info Internal ](#password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

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

###### One of the arguments from this list "check_not_present, check_present, item, presence" must be set

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

###### One of the arguments from this list "check_not_present, check_present, item, presence" must be set

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

###### One of the arguments from this list "check_not_present, check_present, item, presence" must be set

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

###### One of the arguments from this list "check_not_present, check_present, item, presence" must be set

`check_not_present` - (Optional) Check that the query parameter is not present. (`Bool`).

`check_present` - (Optional) Check that the query parameter is present. (`Bool`).

`item` - (Optional) criteria in the matcher.. See [Match Item ](#match-item) below for details.

`presence` - (Optional) Check if the query parameter is present or absent. (`Bool`).(Deprecated)

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

### Threat Mesh Choice Disable Threat Mesh

x-displayName: "Disable".

### Threat Mesh Choice Enable Threat Mesh

x-displayName: "Enable".

### Tls Cert Params Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Tls Certificates Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

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

### Tls Fingerprint Choice Ja4 Tls Fingerprint

ja4_tls_fingerprint.

`exact_values` - (Optional) A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against (`String`).

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

### Tls Parameters Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

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

### Transaction Result Failure Conditions

Failure Conditions.

`name` - (Optional) A case-insensitive HTTP header name. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`status` - (Required) HTTP Status code (`String`).

### Transaction Result Success Conditions

Success Conditions.

`name` - (Optional) A case-insensitive HTTP header name. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`status` - (Required) HTTP Status code (`String`).

### Transaction Result Choice Disable Transaction Result

Disable collection of transaction result..

### Transaction Result Choice Transaction Result

Collect transaction result..

`failure_conditions` - (Optional) Failure Conditions. See [Transaction Result Failure Conditions ](#transaction-result-failure-conditions) below for details.

`success_conditions` - (Optional) Success Conditions. See [Transaction Result Success Conditions ](#transaction-result-success-conditions) below for details.

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

###### One of the arguments from this list "fail_close, fail_open" can be set

`fail_close` - (Optional) Handle the transaction as it failed the OpenAPI specification validation (Block or Report) (`Bool`).(Deprecated)

`fail_open` - (Optional) Continue to process the transaction without enforcing OpenAPI specification (Allow) (`Bool`).(Deprecated)

###### One of the arguments from this list "oversized_body_fail_validation, oversized_body_skip_validation" can be set

`oversized_body_fail_validation` - (Optional) Apply the request/response action (block or report) when the body length is too long to verify (default 64Kb) (`Bool`).

`oversized_body_skip_validation` - (Optional) Skip body validation when the body length is too long to verify (default 64Kb) (`Bool`).

###### One of the arguments from this list "property_validation_settings_custom, property_validation_settings_default" can be set

`property_validation_settings_custom` - (Optional) Use custom settings with Open API specification validation. See [Property Validation Settings Choice Property Validation Settings Custom ](#property-validation-settings-choice-property-validation-settings-custom) below for details.

`property_validation_settings_default` - (Optional) Keep the default settings of OpenAPI specification validation (`Bool`).

### Validation All Spec Endpoints Validation Mode

When a validation mismatch occurs on a request to one of the API Inventory endpoints.

###### One of the arguments from this list "response_validation_mode_active, skip_response_validation" must be set

`response_validation_mode_active` - (Optional) Enforce OpenAPI validation processing for this event. See [Response Validation Mode Choice Response Validation Mode Active ](#response-validation-mode-choice-response-validation-mode-active) below for details.

`skip_response_validation` - (Optional) Skip OpenAPI validation processing for this event (`Bool`).

###### One of the arguments from this list "skip_validation, validation_mode_active" must be set

`skip_validation` - (Optional) Skip OpenAPI validation processing for this event (`Bool`).

`validation_mode_active` - (Optional) Enforce OpenAPI validation processing for this event. See [Validation Mode Choice Validation Mode Active ](#validation-mode-choice-validation-mode-active) below for details.

### Validation Custom List Fall Through Mode

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules).

###### One of the arguments from this list "fall_through_mode_allow, fall_through_mode_custom" must be set

`fall_through_mode_allow` - (Optional) Allow any unprotected end point (`Bool`).

`fall_through_mode_custom` - (Optional) Custom rules for any unprotected end point. See [Fall Through Mode Choice Fall Through Mode Custom ](#fall-through-mode-choice-fall-through-mode-custom) below for details.

### Validation Custom List Open Api Validation Rules

x-displayName: "Validation List".

###### One of the arguments from this list "api_endpoint, api_group, base_path" must be set

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

###### One of the arguments from this list "fail_close, fail_open" can be set

`fail_close` - (Optional) Handle the transaction as it failed the OpenAPI specification validation (Block or Report) (`Bool`).(Deprecated)

`fail_open` - (Optional) Continue to process the transaction without enforcing OpenAPI specification (Allow) (`Bool`).(Deprecated)

###### One of the arguments from this list "oversized_body_fail_validation, oversized_body_skip_validation" can be set

`oversized_body_fail_validation` - (Optional) Apply the request/response action (block or report) when the body length is too long to verify (default 64Kb) (`Bool`).

`oversized_body_skip_validation` - (Optional) Skip body validation when the body length is too long to verify (default 64Kb) (`Bool`).

###### One of the arguments from this list "property_validation_settings_custom, property_validation_settings_default" can be set

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

###### One of the arguments from this list "enforcement_block, enforcement_report" must be set

`enforcement_block` - (Optional) Block the request, trigger an API security event (`Bool`).

`enforcement_report` - (Optional) Allow the request, trigger an API security event (`Bool`).

### Validation Target Choice Validation All Spec Endpoints

All other API endpoints would proceed according to "Fall Through Mode".

`fall_through_mode` - (Required) Determine what to do with unprotected endpoints (not part of the API Inventory or doesn't have a specific rule in custom rules). See [Validation All Spec Endpoints Fall Through Mode ](#validation-all-spec-endpoints-fall-through-mode) below for details.

###### One of the arguments from this list "oversized_body_fail_validation, oversized_body_skip_validation" can be set

`oversized_body_fail_validation` - (Optional) Apply the request/response action (block or report) when the body length is too long to verify (default 64Kb) (`Bool`).(Deprecated)

`oversized_body_skip_validation` - (Optional) Skip body validation when the body length is too long to verify (default 64Kb) (`Bool`).(Deprecated)

`settings` - (Optional) OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom list" enforcement. See [Validation All Spec Endpoints Settings ](#validation-all-spec-endpoints-settings) below for details.

`validation_mode` - (Required) When a validation mismatch occurs on a request to one of the API Inventory endpoints. See [Validation All Spec Endpoints Validation Mode ](#validation-all-spec-endpoints-validation-mode) below for details.

### Validation Target Choice Validation Custom List

Any other end-points not listed will act according to "Fall Through Mode".

`fall_through_mode` - (Required) Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules). See [Validation Custom List Fall Through Mode ](#validation-custom-list-fall-through-mode) below for details.

`open_api_validation_rules` - (Required) x-displayName: "Validation List". See [Validation Custom List Open Api Validation Rules ](#validation-custom-list-open-api-validation-rules) below for details.

###### One of the arguments from this list "oversized_body_fail_validation, oversized_body_skip_validation" can be set

`oversized_body_fail_validation` - (Optional) Apply the request/response action (block or report) when the body length is too long to verify (default 64Kb) (`Bool`).(Deprecated)

`oversized_body_skip_validation` - (Optional) Skip body validation when the body length is too long to verify (default 64Kb) (`Bool`).(Deprecated)

`settings` - (Optional) OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom list" enforcement. See [Validation Custom List Settings ](#validation-custom-list-settings) below for details.

### Validation Target Choice Validation Disabled

Don't run OpenAPI validation.

### Value Choice Secret Value

Secret Value of the HTTP header..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Secret Value Blindfold Secret Info Internal ](#secret-value-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Value Type Check Presence

Parameter name taken which is exist in the query parameter.

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

###### One of the arguments from this list "ignore_samesite, samesite_lax, samesite_none, samesite_strict" can be set

`ignore_samesite` - (Optional) Ignore Samesite attribute (`Bool`).

`samesite_lax` - (Optional) Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests (`Bool`).

`samesite_none` - (Optional) Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests (`Bool`).

`samesite_strict` - (Optional) Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests (`Bool`).

###### One of the arguments from this list "add_secure, ignore_secure" can be set

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

### Web Mobile Header

Header that is used by mobile traffic..

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Web Mobile Headers

Headers that can be used to identify mobile traffic..

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Web Mobile Client Header

Header that is used by mobile traffic..

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Web Mobile Client Headers

Headers that can be used to identify mobile traffic..

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Xfcc Header Xfcc Disabled

No X-Forwarded-Client-Cert header will be added.

### Xfcc Header Xfcc Options

X-Forwarded-Client-Cert header will be added with the configured fields.

`xfcc_header_elements` - (Required) X-Forwarded-Client-Cert header elements to be added to requests (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured cdn_loadbalancer.
