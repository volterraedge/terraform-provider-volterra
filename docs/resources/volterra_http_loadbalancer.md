---

page_title: "Volterra: http_loadbalancer"

description: "The http_loadbalancer allows CRUD of Http Loadbalancer resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_http_loadbalancer
===================================

The Http Loadbalancer allows CRUD of Http Loadbalancer resource on Volterra SaaS

~> **Note:** Please refer to [Http Loadbalancer API docs](https://docs.cloud.f5.com/docs-v2/api/views-http-loadbalancer) to learn more

Example Usage
-------------

```hcl
resource "volterra_http_loadbalancer" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "advertise_custom advertise_on_public advertise_on_public_default_vip do_not_advertise" must be set

  advertise_custom {
    advertise_where {
      site {
        site {
          name   = "site1"
          namespace = "system"
          tenant    = "acmecorp"
        }
         network = "SITE_NETWORK_INSIDE_AND_OUTSIDE"
      }
    }
    advertise_where {
      site {
        site {
          name   = "site2"
          namespace = "system"
          tenant    = "acmecorp"
        }
        network = "SITE_NETWORK_INSIDE_AND_OUTSIDE"
      }
    }
  }

  // One of the arguments from this list "api_specification disable_api_definition" must be set

  api_specification {
    api_definition {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }

    // One of the arguments from this list "validation_all_spec_endpoints validation_custom_list validation_disabled" must be set

    validation_disabled = true
  }

  // One of the arguments from this list "disable_api_discovery enable_api_discovery" must be set

  enable_api_discovery {
    api_crawler {
      // One of the arguments from this list "api_crawler_config disable_api_crawler" must be set

      disable_api_crawler = true
    }

    api_discovery_from_code_scan {
      code_base_integrations {
        // One of the arguments from this list "all_repos selected_repos" must be set

        all_repos = true

        code_base_integration {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
      }
    }

    // One of the arguments from this list "custom_api_auth_discovery default_api_auth_discovery" must be set

    default_api_auth_discovery = true
    discovered_api_settings {
      purge_duration_for_inactive_discovered_apis = "2"
    }

    // One of the arguments from this list "disable_learn_from_redirect_traffic enable_learn_from_redirect_traffic" must be set

    disable_learn_from_redirect_traffic = true
  }

  // One of the arguments from this list "api_testing disable_api_testing" must be set

  disable_api_testing = true

  // One of the arguments from this list "captcha_challenge enable_challenge js_challenge no_challenge policy_based_challenge" must be set

  captcha_challenge {
    cookie_expiry = "1000"

    custom_page = "string:///PHA+IFBsZWFzZSBXYWl0IDwvcD4="
  }
  domains = ["www.foo.com"]

  // One of the arguments from this list "cookie_stickiness least_active random ring_hash round_robin source_ip_stickiness" must be set

  random = true

  // One of the arguments from this list "http https https_auto_cert" must be set

  https_auto_cert {
    add_hsts = true

    coalescing_options {
      // One of the arguments from this list "default_coalescing strict_coalescing" must be set

      default_coalescing = true
    }

    connection_idle_timeout = "60000"

    // One of the arguments from this list "default_loadbalancer non_default_loadbalancer" can be set

    non_default_loadbalancer = true
    http_protocol_options {
      // One of the arguments from this list "http_protocol_enable_v1_only http_protocol_enable_v1_v2 http_protocol_enable_v2_only" must be set

      http_protocol_enable_v1_only {
        header_transformation {
          // One of the arguments from this list "default_header_transformation legacy_header_transformation preserve_case_header_transformation proper_case_header_transformation" must be set

          legacy_header_transformation = true
        }
      }
    }
    http_redirect = true

    // One of the arguments from this list "no_mtls use_mtls" must be set

    no_mtls = true

    // One of the arguments from this list "disable_path_normalize enable_path_normalize" must be set

    enable_path_normalize = true

    // One of the arguments from this list "port port_ranges" can be set

    port = "443"

    // One of the arguments from this list "append_server_name default_header pass_through server_name" can be set

    server_name = "server_name"
    tls_config {
      // One of the arguments from this list "custom_security default_security low_security medium_security" must be set

        default_security = true
    }
  }

  // One of the arguments from this list "disable_malicious_user_detection enable_malicious_user_detection" must be set

  enable_malicious_user_detection = true

  // One of the arguments from this list "disable_malware_protection malware_protection_settings" must be set

  disable_malware_protection = true

  // One of the arguments from this list "api_rate_limit disable_rate_limit rate_limit" must be set

  disable_rate_limit = true

  // One of the arguments from this list "default_sensitive_data_policy sensitive_data_policy" must be set

  default_sensitive_data_policy = true

  // One of the arguments from this list "active_service_policies no_service_policies service_policies_from_namespace" must be set

  service_policies_from_namespace = true

  // One of the arguments from this list "disable_threat_mesh enable_threat_mesh" must be set

  disable_threat_mesh = true

  // One of the arguments from this list "disable_trust_client_ip_headers enable_trust_client_ip_headers" must be set

  disable_trust_client_ip_headers = true

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

`add_location` - (Optional) is ignored on CE sites. (`Bool`).

###### One of the arguments from this list "advertise_custom, advertise_on_public, advertise_on_public_default_vip, do_not_advertise" must be set

`advertise_custom` - (Optional) Advertise this load balancer on specific sites. See [Advertise Choice Advertise Custom ](#advertise-choice-advertise-custom) below for details.

`advertise_on_public` - (Optional) Advertise this load balancer on public network. See [Advertise Choice Advertise On Public ](#advertise-choice-advertise-on-public) below for details.

`advertise_on_public_default_vip` - (Optional) Advertise this load balancer on public network with default VIP (`Bool`).

`do_not_advertise` - (Optional) Do not advertise this load balancer (`Bool`).

###### One of the arguments from this list "api_specification, disable_api_definition" must be set

`api_specification` - (Optional) Specify API definition and OpenAPI Validation. See [Api Definition Choice Api Specification ](#api-definition-choice-api-specification) below for details.

`disable_api_definition` - (Optional) API Definition is not currently used for this load balancer (`Bool`).

###### One of the arguments from this list "disable_api_discovery, enable_api_discovery" must be set

`disable_api_discovery` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_api_discovery` - (Optional) x-displayName: "Enable". See [Api Discovery Choice Enable Api Discovery ](#api-discovery-choice-enable-api-discovery) below for details.

`api_protection_rules` - (Optional) Rules can also include additional conditions, for example specific clients can access certain API endpoint or API group.. See [Api Protection Rules ](#api-protection-rules) below for details.

###### One of the arguments from this list "api_testing, disable_api_testing" must be set

`api_testing` - (Optional) x-displayName: "Enable". See [Api Testing Choice Api Testing ](#api-testing-choice-api-testing) below for details.

`disable_api_testing` - (Optional) x-displayName: "Disable" (`Bool`).

`blocked_clients` - (Optional) Define rules to block IP Prefixes or AS numbers.. See [Blocked Clients ](#blocked-clients) below for details.

###### One of the arguments from this list "bot_defense, disable_bot_defense" can be set

`bot_defense` - (Optional) Select Bot Defense Standard. See [Bot Defense Choice Bot Defense ](#bot-defense-choice-bot-defense) below for details.

`disable_bot_defense` - (Optional) No Bot Defense configuration for this load balancer (`Bool`).

###### One of the arguments from this list "captcha_challenge, enable_challenge, js_challenge, no_challenge, policy_based_challenge" must be set

`captcha_challenge` - (Optional) Configure Captcha challenge on this load balancer. See [Challenge Type Captcha Challenge ](#challenge-type-captcha-challenge) below for details.

`enable_challenge` - (Optional) Configure auto mitigation i.e risk based challenges for malicious users. See [Challenge Type Enable Challenge ](#challenge-type-enable-challenge) below for details.

`js_challenge` - (Optional) Configure JavaScript challenge on this load balancer. See [Challenge Type Js Challenge ](#challenge-type-js-challenge) below for details.

`no_challenge` - (Optional) No challenge is enabled for this load balancer (`Bool`).

`policy_based_challenge` - (Optional) Specifies the settings for policy rule based challenge. See [Challenge Type Policy Based Challenge ](#challenge-type-policy-based-challenge) below for details.

###### One of the arguments from this list "client_side_defense, disable_client_side_defense" can be set

`client_side_defense` - (Optional) Client-Side Defense configuration for JavaScript insertion. See [Client Side Defense Choice Client Side Defense ](#client-side-defense-choice-client-side-defense) below for details.

`disable_client_side_defense` - (Optional) No Client-Side Defense configuration for this load balancer (`Bool`).

`cors_policy` - (Optional) resources from a server at a different origin. See [Cors Policy ](#cors-policy) below for details.

`csrf_policy` - (Optional) Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.).. See [Csrf Policy ](#csrf-policy) below for details.

`data_guard_rules` - (Optional) Note: App Firewall should be enabled, to use Data Guard feature.. See [Data Guard Rules ](#data-guard-rules) below for details.

`ddos_mitigation_rules` - (Optional) Define manual mitigation rules to block L7 DDoS attacks.. See [Ddos Mitigation Rules ](#ddos-mitigation-rules) below for details.

`default_route_pools` - (Optional) Origin Pools used when no route is specified (default route). See [Default Route Pools ](#default-route-pools) below for details.

`domains` - (Required) Domains also indicate the list of names for which DNS resolution will be done by VER (`List of String`).

`graphql_rules` - (Optional) queries and prevent GraphQL tailored attacks.. See [Graphql Rules ](#graphql-rules) below for details.

###### One of the arguments from this list "cookie_stickiness, least_active, random, ring_hash, round_robin, source_ip_stickiness" must be set

`cookie_stickiness` - (Optional) Consistent hashing algorithm, ring hash, is used to select origin server. See [Hash Policy Choice Cookie Stickiness ](#hash-policy-choice-cookie-stickiness) below for details.

`least_active` - (Optional) Request are sent to origin server that has least active requests (`Bool`).

`random` - (Optional) Request are sent to all eligible origin servers in random fashion (`Bool`).

`ring_hash` - (Optional) Request are sent to all eligible origin servers using hash of request based on hash policy. Consistent hashing algorithm, ring hash, is used to select origin server. See [Hash Policy Choice Ring Hash ](#hash-policy-choice-ring-hash) below for details.

`round_robin` - (Optional) Request are sent to all eligible origin servers in round robin fashion (`Bool`).

`source_ip_stickiness` - (Optional) Request are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server (`Bool`).

###### One of the arguments from this list "disable_ip_reputation, enable_ip_reputation" can be set

`disable_ip_reputation` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_ip_reputation` - (Optional) x-displayName: "Enable". See [Ip Reputation Choice Enable Ip Reputation ](#ip-reputation-choice-enable-ip-reputation) below for details.

`jwt_validation` - (Optional) tokens or tokens that are not yet valid.. See [Jwt Validation ](#jwt-validation) below for details.

`l7_ddos_protection` - (Optional) L7 DDoS attack. See [L7 Ddos Protection ](#l7-ddos-protection) below for details.

###### One of the arguments from this list "http, https, https_auto_cert" must be set

`http` - (Optional) HTTP Load Balancer.. See [Loadbalancer Type Http ](#loadbalancer-type-http) below for details.

`https` - (Optional) User is responsible for managing DNS to this load balancer.. See [Loadbalancer Type Https ](#loadbalancer-type-https) below for details.

`https_auto_cert` - (Optional) or a DNS CNAME record should be created in your DNS provider's portal(only for Domains not managed by F5 Distributed Cloud).. See [Loadbalancer Type Https Auto Cert ](#loadbalancer-type-https-auto-cert) below for details.

###### One of the arguments from this list "disable_malicious_user_detection, enable_malicious_user_detection" must be set

`disable_malicious_user_detection` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_malicious_user_detection` - (Optional) x-displayName: "Enable" (`Bool`).

###### One of the arguments from this list "disable_malware_protection, malware_protection_settings" must be set

`disable_malware_protection` - (Optional) x-displayName: "Disable" (`Bool`).

`malware_protection_settings` - (Optional) x-displayName: "Enable". See [Malware Protection Malware Protection Settings ](#malware-protection-malware-protection-settings) below for details.

`more_option` - (Optional) More options like header manipulation, compression etc.. See [More Option ](#more-option) below for details.

`origin_server_subset_rule_list` - (Optional) When an Origin server subset rule is matched, then this selection rule takes effect and no more rules are evaluated.. See [Origin Server Subset Rule List ](#origin-server-subset-rule-list) below for details.

`protected_cookies` - (Optional) Note: We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.. See [Protected Cookies ](#protected-cookies) below for details.

###### One of the arguments from this list "api_rate_limit, disable_rate_limit, rate_limit" must be set

`api_rate_limit` - (Optional) Define rate limiting for one or more API endpoints.. See [Rate Limit Choice Api Rate Limit ](#rate-limit-choice-api-rate-limit) below for details.

`disable_rate_limit` - (Optional) Rate limiting is not currently enabled for this load balancer (`Bool`).

`rate_limit` - (Optional) Define custom rate limiting parameters for this load balancer. See [Rate Limit Choice Rate Limit ](#rate-limit-choice-rate-limit) below for details.

`routes` - (Optional) to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic. See [Routes ](#routes) below for details.

`sensitive_data_disclosure_rules` - (Optional) Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses. See [Sensitive Data Disclosure Rules ](#sensitive-data-disclosure-rules) below for details.

###### One of the arguments from this list "default_sensitive_data_policy, sensitive_data_policy" must be set

`default_sensitive_data_policy` - (Optional) Apply system default sensitive data discovery (`Bool`).

`sensitive_data_policy` - (Optional) Apply custom sensitive data discovery. See [Sensitive Data Policy Choice Sensitive Data Policy ](#sensitive-data-policy-choice-sensitive-data-policy) below for details.

###### One of the arguments from this list "active_service_policies, no_service_policies, service_policies_from_namespace" must be set

`active_service_policies` - (Optional) Apply the specified list of service policies and bypass the namespace service policy set. See [Service Policy Choice Active Service Policies ](#service-policy-choice-active-service-policies) below for details.

`no_service_policies` - (Optional) Do not apply any service policies i.e. bypass the namespace service policy set (`Bool`).

`service_policies_from_namespace` - (Optional) Apply the active service policies configured as part of the namespace service policy set (`Bool`).

###### One of the arguments from this list "slow_ddos_mitigation, system_default_timeouts" can be set

`slow_ddos_mitigation` - (Optional) Custom Settings for Slow DDoS Mitigation. See [Slow Ddos Mitigation Choice Slow Ddos Mitigation ](#slow-ddos-mitigation-choice-slow-ddos-mitigation) below for details.

`system_default_timeouts` - (Optional) Default Settings for Slow DDoS Mitigation (`Bool`).

###### One of the arguments from this list "disable_threat_mesh, enable_threat_mesh" must be set

`disable_threat_mesh` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_threat_mesh` - (Optional) x-displayName: "Enable" (`Bool`).

###### One of the arguments from this list "disable_trust_client_ip_headers, enable_trust_client_ip_headers" must be set

`disable_trust_client_ip_headers` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_trust_client_ip_headers` - (Optional) x-displayName: "Enable". See [Trust Client Ip Headers Choice Enable Trust Client Ip Headers ](#trust-client-ip-headers-choice-enable-trust-client-ip-headers) below for details.

`trusted_clients` - (Optional) Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients.. See [Trusted Clients ](#trusted-clients) below for details.

###### One of the arguments from this list "user_id_client_ip, user_identification" must be set

`user_id_client_ip` - (Optional) Use the Client IP address as the user identifier. (`Bool`).

`user_identification` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier.. See [ref](#ref) below for details.

###### One of the arguments from this list "app_firewall, disable_waf" must be set

`app_firewall` - (Optional) Reference to App Firewall configuration object. See [ref](#ref) below for details.

`disable_waf` - (Optional) No WAF configuration for this load balancer (`Bool`).

`waf_exclusion_rules` - (Optional) When an exclusion rule is matched, then this exclusion rule takes effect and no more rules are evaluated.. See [Waf Exclusion Rules ](#waf-exclusion-rules) below for details.

### Api Protection Rules

Rules can also include additional conditions, for example specific clients can access certain API endpoint or API group..

`api_endpoint_rules` - (Optional) If request matches any of these rules, skipping second category rules.. See [Api Protection Rules Api Endpoint Rules ](#api-protection-rules-api-endpoint-rules) below for details.

`api_groups_rules` - (Optional) For API groups, refer to API Definition which includes API groups derived from uploaded swaggers.. See [Api Protection Rules Api Groups Rules ](#api-protection-rules-api-groups-rules) below for details.

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

### Default Route Pools

Origin Pools used when no route is specified (default route).

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

###### One of the arguments from this list "cluster, pool" must be set

`cluster` - (Optional) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Optional) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`priority` - (Optional) made active as per the increasing priority. (`Int`).

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

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

### L7 Ddos Protection

L7 DDoS attack.

###### One of the arguments from this list "ddos_policy_custom, ddos_policy_none" must be set

`ddos_policy_custom` - (Optional) Apply a custom service policy during an ongoing DDoS attack. See [ref](#ref) below for details.

`ddos_policy_none` - (Optional) Do not apply additional service policy during an ongoing DDoS attack (`Bool`).

###### One of the arguments from this list "mitigation_block, mitigation_js_challenge" must be set

`mitigation_block` - (Optional) Block suspicious sources (`Bool`).

`mitigation_js_challenge` - (Optional) Serve JavaScript challenge to suspicious sources. See [Mitigation Action Choice Mitigation Js Challenge ](#mitigation-action-choice-mitigation-js-challenge) below for details.

### More Option

More options like header manipulation, compression etc..

`buffer_policy` - (Optional) specify the maximum buffer size and buffer interval with this config.. See [More Option Buffer Policy ](#more-option-buffer-policy) below for details.

`compression_params` - (Optional) Only GZIP compression is supported. See [More Option Compression Params ](#more-option-compression-params) below for details.

`cookies_to_modify` - (Optional) List of cookies to be modified from the HTTP response being sent towards downstream.. See [More Option Cookies To Modify ](#more-option-cookies-to-modify) below for details.(Deprecated)

`custom_errors` - (Optional) Map of integer error codes as keys and string values that can be used to provide custom http pages for each error code. Key of the map can be either response code class or HTTP Error code. Response code classes for key is configured as follows 3 -- for 3xx response code class, 4 -- for 4xx response code class, 5 -- for 5xx response code class. Value of the map is string which represents custom HTTP responses. Specific response code takes preference when both response code and response code class matches for a request. (`map(string)`).

`disable_default_error_pages` - (Optional) Disable the use of default F5XC error pages. (`Bool`).

`idle_timeout` - (Optional) received, otherwise the stream is reset. (`Int`).

`max_request_header_size` - (Optional) such load balancers is used for all the load balancers in question. (`Int`).

`request_cookies_to_add` - (Optional) Cookies specified at this level are applied after cookies from matched Route are applied. See [More Option Request Cookies To Add ](#more-option-request-cookies-to-add) below for details.

`request_cookies_to_remove` - (Optional) List of keys of Cookies to be removed from the HTTP request being sent towards upstream. (`String`).

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [More Option Request Headers To Add ](#more-option-request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_cookies_to_add` - (Optional) Cookies specified at this level are applied after cookies from matched Route are applied. See [More Option Response Cookies To Add ](#more-option-response-cookies-to-add) below for details.

`response_cookies_to_remove` - (Optional) List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [More Option Response Headers To Add ](#more-option-response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

### Origin Server Subset Rule List

When an Origin server subset rule is matched, then this selection rule takes effect and no more rules are evaluated..

`origin_server_subset_rules` - (Optional) When an Origin server subset rule is matched, then this selection rule takes effect and no more rules are evaluated.. See [Origin Server Subset Rule List Origin Server Subset Rules ](#origin-server-subset-rule-list-origin-server-subset-rules) below for details.

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

### Routes

to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic.

###### One of the arguments from this list "custom_route_object, direct_response_route, redirect_route, simple_route" must be set

`custom_route_object` - (Optional) A custom route uses a route object created outside of this view.. See [Choice Custom Route Object ](#choice-custom-route-object) below for details.

`direct_response_route` - (Optional) A direct response route matches on path and/or HTTP method and responds directly to the matching traffic. See [Choice Direct Response Route ](#choice-direct-response-route) below for details.

`redirect_route` - (Optional) A redirect route matches on path and/or HTTP method and redirects the matching traffic to a different URL. See [Choice Redirect Route ](#choice-redirect-route) below for details.

`simple_route` - (Optional) A simple route matches on path and/or HTTP method and forwards the matching traffic to the associated pools. See [Choice Simple Route ](#choice-simple-route) below for details.

### Sensitive Data Disclosure Rules

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses.

`sensitive_data_types_in_response` - (Optional) Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses. See [Sensitive Data Disclosure Rules Sensitive Data Types In Response ](#sensitive-data-disclosure-rules-sensitive-data-types-in-response) below for details.

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

### Action Allow

Allow the request to proceed..

### Action Deny

Deny the request..

### Action Choice Action Block

Block the request and issue an API security event.

### Action Choice Action Block

Blocks the user for a specified duration of time.

###### One of the arguments from this list "hours, minutes, seconds" can be set

`hours` - (Optional) User block mitigation time in Hours. See [Block Duration Choice Hours ](#block-duration-choice-hours) below for details.

`minutes` - (Optional) User block mitigation time in Minutes. See [Block Duration Choice Minutes ](#block-duration-choice-minutes) below for details.

`seconds` - (Optional) User block mitigation time in Seconds. See [Block Duration Choice Seconds ](#block-duration-choice-seconds) below for details.

### Action Choice Action Report

Continue processing the request and issue an API security event.

### Action Choice Action Skip

Continue processing the request.

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

### Action Type Block

Block bot request and send response with custom content..

`body` - (Optional) E.g. "<p> Your request was blocked </p>". Base64 encoded string for this html is "LzxwPiBZb3VyIHJlcXVlc3Qgd2FzIGJsb2NrZWQgPC9wPg==" (`String`).

`status` - (Optional) HTTP Status code to respond with (`String`).

### Action Type Flag

Flag the request while not taking any invasive actions..

###### One of the arguments from this list "append_headers, no_headers" can be set

`append_headers` - (Optional) Append mitigation headers.. See [Send Headers Choice Append Headers ](#send-headers-choice-append-headers) below for details.

`no_headers` - (Optional) No mitigation headers. (`Bool`).

### Action Type Redirect

Redirect bot request to a custom URI..

`uri` - (Required) URI location for redirect may be relative or absolute. (`String`).

### Additional Parameters Choice Allow Additional Parameters

Allow extra query parameters (on top of what specified in the OAS documentation).

### Additional Parameters Choice Disallow Additional Parameters

Disallow extra query parameters (on top of what specified in the OAS documentation).

### Advanced Options Cors Policy

resources from a server at a different origin.

`allow_credentials` - (Optional) Specifies whether the resource allows credentials (`Bool`).

`allow_headers` - (Optional) Specifies the content for the access-control-allow-headers header (`String`).

`allow_methods` - (Optional) Specifies the content for the access-control-allow-methods header (`String`).

`allow_origin` - (Optional) An origin is allowed if either allow_origin or allow_origin_regex match (`String`).

`allow_origin_regex` - (Optional) An origin is allowed if either allow_origin or allow_origin_regex match (`String`).

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`expose_headers` - (Optional) Specifies the content for the access-control-expose-headers header (`String`).

`maximum_age` - (Optional) Maximum permitted value is 86400 seconds (24 hours) (`Int`).

### Advanced Options Csrf Policy

Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)..

###### One of the arguments from this list "all_load_balancer_domains, custom_domain_list, disabled" must be set

`all_load_balancer_domains` - (Optional) Add All load balancer domains to source origin (allow) list. (`Bool`).

`custom_domain_list` - (Optional) Add one or more domains to source origin (allow) list.. See [Allowed Domains Custom Domain List ](#allowed-domains-custom-domain-list) below for details.

`disabled` - (Optional) Allow all source origin domains. (`Bool`).

### Advanced Options Request Cookies To Add

Cookies specified at this level are applied after cookies from matched Route are applied.

`name` - (Required) Name of the cookie in Cookie header. (`String`).

`overwrite` - (Optional) Default value is do not overwrite (`Bool`).

###### One of the arguments from this list "secret_value, value" must be set

`secret_value` - (Optional) Secret Value of the Cookie header. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the Cookie header. (`String`).

### Advanced Options Request Headers To Add

Headers are key-value pairs to be added to HTTP request being routed towards upstream..

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "secret_value, value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Advanced Options Response Cookies To Add

Cookies specified at this level are applied after cookies from matched Route are applied.

###### One of the arguments from this list "add_domain, ignore_domain" can be set

`add_domain` - (Optional) Add domain attribute (`String`).

`ignore_domain` - (Optional) Ignore max age attribute (`Bool`).

###### One of the arguments from this list "add_expiry, ignore_expiry" can be set

`add_expiry` - (Optional) Add expiry attribute (`String`).

`ignore_expiry` - (Optional) Ignore expiry attribute (`Bool`).

###### One of the arguments from this list "add_httponly, ignore_httponly" can be set

`add_httponly` - (Optional) x-displayName: "Add" (`Bool`).

`ignore_httponly` - (Optional) x-displayName: "Ignore" (`Bool`).

###### One of the arguments from this list "ignore_max_age, max_age_value" can be set

`ignore_max_age` - (Optional) Ignore max age attribute (`Bool`).

`max_age_value` - (Optional) Add max age attribute (`Int`).

`name` - (Required) Name of the cookie in Cookie header. (`String`).

`overwrite` - (Optional) Default value is do not overwrite (`Bool`).

###### One of the arguments from this list "add_partitioned, ignore_partitioned" can be set

`add_partitioned` - (Optional) x-displayName: "Add" (`Bool`).

`ignore_partitioned` - (Optional) x-displayName: "Ignore" (`Bool`).

###### One of the arguments from this list "add_path, ignore_path" can be set

`add_path` - (Optional) Add path attribute (`String`).

`ignore_path` - (Optional) Ignore path attribute (`Bool`).

###### One of the arguments from this list "ignore_samesite, samesite_lax, samesite_none, samesite_strict" can be set

`ignore_samesite` - (Optional) Ignore Samesite attribute (`Bool`).

`samesite_lax` - (Optional) Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests (`Bool`).

`samesite_none` - (Optional) Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests (`Bool`).

`samesite_strict` - (Optional) Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests (`Bool`).

###### One of the arguments from this list "add_secure, ignore_secure" can be set

`add_secure` - (Optional) x-displayName: "Add" (`Bool`).

`ignore_secure` - (Optional) x-displayName: "Ignore" (`Bool`).

###### One of the arguments from this list "ignore_value, secret_value, value" can be set

`ignore_value` - (Optional) Ignore value of cookie (`Bool`).

`secret_value` - (Optional) Secret Value of the Cookie header. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the Cookie header. (`String`).

### Advanced Options Response Headers To Add

Headers are key-value pairs to be added to HTTP response being sent towards downstream..

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "secret_value, value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Advertise Choice Advertise Custom

Advertise this load balancer on specific sites.

`advertise_where` - (Required) Where should this load balancer be available. See [Advertise Custom Advertise Where ](#advertise-custom-advertise-where) below for details.

### Advertise Choice Advertise On Public

Advertise this load balancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Advertise Custom Advertise Where

Where should this load balancer be available.

###### One of the arguments from this list "advertise_on_public, site, site_segment, virtual_network, virtual_site, virtual_site_segment, virtual_site_with_vip, vk8s_service" must be set

`advertise_on_public` - (Optional) Advertise this load balancer on public network. See [Choice Advertise On Public ](#choice-advertise-on-public) below for details.

`site` - (Optional) Advertise on a customer site and a given network.. See [Choice Site ](#choice-site) below for details.

`site_segment` - (Optional) Advertise on a segment on a site. See [Choice Site Segment ](#choice-site-segment) below for details.

`virtual_network` - (Optional) Advertise on a virtual network. See [Choice Virtual Network ](#choice-virtual-network) below for details.

`virtual_site` - (Optional) Advertise on a customer virtual site and a given network.. See [Choice Virtual Site ](#choice-virtual-site) below for details.

`virtual_site_segment` - (Optional) Advertise on a segment on a virtual site. See [Choice Virtual Site Segment ](#choice-virtual-site-segment) below for details.

`virtual_site_with_vip` - (Optional) Advertise on a customer virtual site and a given network and IP.. See [Choice Virtual Site With Vip ](#choice-virtual-site-with-vip) below for details.

`vk8s_service` - (Optional) Advertise on vK8s Service Network on RE.. See [Choice Vk8s Service ](#choice-vk8s-service) below for details.

###### One of the arguments from this list "port, port_ranges, use_default_port" must be set

`port` - (Optional) Port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

`use_default_port` - (Optional) Inherit the Load Balancer's Listen Port. (`Bool`).

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

`simple_login` - (Optional) Enter the username and password to assign credentials for the selected domain to crawl. See [Domains Simple Login ](#domains-simple-login) below for details.

### Api Definition Choice Api Specification

Specify API definition and OpenAPI Validation.

`api_definition` - (Required) Specify API definition which includes application API paths and methods derived from swagger files.. See [ref](#ref) below for details.

###### One of the arguments from this list "validation_all_spec_endpoints, validation_custom_list, validation_disabled" must be set

`validation_all_spec_endpoints` - (Optional) All other API endpoints would proceed according to "Fall Through Mode". See [Validation Target Choice Validation All Spec Endpoints ](#validation-target-choice-validation-all-spec-endpoints) below for details.

`validation_custom_list` - (Optional) Any other end-points not listed will act according to "Fall Through Mode". See [Validation Target Choice Validation Custom List ](#validation-target-choice-validation-custom-list) below for details.

`validation_disabled` - (Optional) Don't run OpenAPI validation (`Bool`).

### Api Discovery Choice Enable Api Discovery

x-displayName: "Enable".

`api_crawler` - (Optional) Configure Discovered API Settings.. See [Enable Api Discovery Api Crawler ](#enable-api-discovery-api-crawler) below for details.

`api_discovery_from_code_scan` - (Optional) Select API code repositories to the load balancer to use them as a source for API endpoint discovery.. See [Enable Api Discovery Api Discovery From Code Scan ](#enable-api-discovery-api-discovery-from-code-scan) below for details.

###### One of the arguments from this list "custom_api_auth_discovery, default_api_auth_discovery" must be set

`custom_api_auth_discovery` - (Optional) Apply custom API discovery settings. See [Api Discovery Settings Choice Custom Api Auth Discovery ](#api-discovery-settings-choice-custom-api-auth-discovery) below for details.

`default_api_auth_discovery` - (Optional) Apply system default API discovery settings (`Bool`).

`discovered_api_settings` - (Optional) Configure Discovered API Settings.. See [Enable Api Discovery Discovered Api Settings ](#enable-api-discovery-discovered-api-settings) below for details.

###### One of the arguments from this list "disable_learn_from_redirect_traffic, enable_learn_from_redirect_traffic" must be set

`disable_learn_from_redirect_traffic` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`enable_learn_from_redirect_traffic` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

### Api Discovery From Code Scan Code Base Integrations

x-required.

###### One of the arguments from this list "all_repos, selected_repos" must be set

`all_repos` - (Optional) x-displayName: "All API Repositories" (`Bool`).

`selected_repos` - (Optional) x-displayName: "Selected API Repositories". See [Api Repos Choice Selected Repos ](#api-repos-choice-selected-repos) below for details.

`code_base_integration` - (Required) Select the code base integration for use in code-based API discovery. See [ref](#ref) below for details.

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

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Api Groups Rules Request Matcher

Conditions related to the request, such as query parameters, headers, etc..

`cookie_matchers` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Request Matcher Cookie Matchers ](#request-matcher-cookie-matchers) below for details.

`headers` - (Optional) Note that all specified header predicates must evaluate to true.. See [Request Matcher Headers ](#request-matcher-headers) below for details.

`jwt_claims` - (Optional) Note that this feature only works on LBs with JWT Validation feature enabled.. See [Request Matcher Jwt Claims ](#request-matcher-jwt-claims) below for details.

`query_params` - (Optional) Note that all specified query parameter predicates must evaluate to true.. See [Request Matcher Query Params ](#request-matcher-query-params) below for details.

### Api Key Value

x-displayName: "Value".

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

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

### Api Testing Domains

Add and configure testing domains and credentials.

`allow_destructive_methods` - (Optional) Enable to allow API test to execute destructive methods. Be cautious as these can alter or delete data. (`Bool`).

`credentials` - (Required) Add credentials for API testing to use in the selected environment.. See [Domains Credentials ](#domains-credentials) below for details.

`domain` - (Required) Add your testing environment domain. Be aware that running tests on a production domain can impact live applications, as API testing cannot distinguish between production and testing environments. (`String`).

### Api Testing Choice Api Testing

x-displayName: "Enable".

`custom_header_value` - (Optional) Add x-f5-api-testing-identifier header value to prevent security flags on API testing traffic (`String`).

`domains` - (Required) Add and configure testing domains and credentials. See [Api Testing Domains ](#api-testing-domains) below for details.

###### One of the arguments from this list "every_day, every_month, every_week" must be set

`every_day` - (Optional) x-displayName: "Every Day" (`Bool`).

`every_month` - (Optional) x-displayName: "Every Month" (`Bool`).

`every_week` - (Optional) x-displayName: "Every Week" (`Bool`).

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

### App Traffic Type Choice Web

Web traffic channel..

### App Traffic Type Choice Web Mobile

Web and mobile traffic channel..

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

### Basic Auth Password

x-displayName: "Password".

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Bearer Token Token

x-displayName: "Token".

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

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

`protected_app_endpoints` - (Required) List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32 endpoints per LB' after 4 LBs.. See [Policy Protected App Endpoints ](#policy-protected-app-endpoints) below for details.

### Bot Defense Choice Bot Defense

Select Bot Defense Standard.

`policy` - (Required) Bot Defense Policy.. See [Bot Defense Policy ](#bot-defense-policy) below for details.

`regional_endpoint` - (Required) x-required (`String`).

`timeout` - (Optional) The timeout for the inference check, in milliseconds. (`Int`).

### Bot Defense Javascript Injection Javascript Tags

Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first..

`javascript_url` - (Required) Please enter the full URL (include domain and path), or relative path. (`String`).

`tag_attributes` - (Optional) Add the tag attributes you want to include in your Javascript tag.. See [Javascript Tags Tag Attributes ](#javascript-tags-tag-attributes) below for details.

### Bot Defense Javascript Injection Choice Bot Defense Javascript Injection

Configuration for Bot Defense JavaScript Injection.

`javascript_location` - (Optional) Select the location where you would like to insert the Javascript tag(s). (`String`).

`javascript_tags` - (Required) Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first.. See [Bot Defense Javascript Injection Javascript Tags ](#bot-defense-javascript-injection-javascript-tags) below for details.

### Bot Defense Javascript Injection Choice Inherited Bot Defense Javascript Injection

Hence no custom configuration is applied on the route.

### Buffer Choice Buffer Policy

Route level buffer configuration overrides any configuration at VirtualHost level..

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`max_request_bytes` - (Optional) manager will stop buffering and return a RequestEntityTooLarge (413) response. (`Int`).

### Buffer Choice Common Buffering

Use common buffering configuration.

### Bypass Rate Limiting Rules Bypass Rate Limiting Rules

This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting..

`client_matcher` - (Optional) Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc.. See [Bypass Rate Limiting Rules Client Matcher ](#bypass-rate-limiting-rules-client-matcher) below for details.

###### One of the arguments from this list "any_url, api_endpoint, api_groups, base_path" must be set

`any_url` - (Optional) Any URL (`Bool`).

`api_endpoint` - (Optional) The endpoint (path) of the request.. See [Destination Type Api Endpoint ](#destination-type-api-endpoint) below for details.

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

Configure auto mitigation i.e risk based challenges for malicious users.

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

### Choice Advertise On Public

Advertise this load balancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Choice Custom Route Object

A custom route uses a route object created outside of this view..

`route_ref` - (Optional) Reference to a custom route object. See [ref](#ref) below for details.

### Choice Custom Security

Custom selection of TLS versions and cipher suites.

`cipher_suites` - (Required) The TLS listener will only support the specified cipher list. (`String`).

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

### Choice Default Security

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Choice Direct Response Route

A direct response route matches on path and/or HTTP method and responds directly to the matching traffic.

`headers` - (Optional) List of (key, value) headers. See [Direct Response Route Headers ](#direct-response-route-headers) below for details.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`incoming_port` - (Optional) The port on which the request is received. See [Direct Response Route Incoming Port ](#direct-response-route-incoming-port) below for details.

`path` - (Required) URI path of route. See [Direct Response Route Path ](#direct-response-route-path) below for details.

`route_direct_response` - (Optional) Send direct response. See [Direct Response Route Route Direct Response ](#direct-response-route-route-direct-response) below for details.

### Choice Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Choice Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### Choice Redirect Route

A redirect route matches on path and/or HTTP method and redirects the matching traffic to a different URL.

`headers` - (Optional) List of (key, value) headers. See [Redirect Route Headers ](#redirect-route-headers) below for details.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`incoming_port` - (Optional) The port on which the request is received. See [Redirect Route Incoming Port ](#redirect-route-incoming-port) below for details.

`path` - (Required) URI path of route. See [Redirect Route Path ](#redirect-route-path) below for details.

`route_redirect` - (Optional) Send redirect response. See [Redirect Route Route Redirect ](#redirect-route-route-redirect) below for details.

### Choice Simple Route

A simple route matches on path and/or HTTP method and forwards the matching traffic to the associated pools.

`advanced_options` - (Optional) Configure Advanced per route options. See [Simple Route Advanced Options ](#simple-route-advanced-options) below for details.

`headers` - (Optional) List of (key, value) headers. See [Simple Route Headers ](#simple-route-headers) below for details.

###### One of the arguments from this list "auto_host_rewrite, disable_host_rewrite, host_rewrite" must be set

`auto_host_rewrite` - (Optional) Host header will be swapped with hostname of upstream host chosen by the cluster (`Bool`).

`disable_host_rewrite` - (Optional) Host header is not modified (`Bool`).

`host_rewrite` - (Optional) Host header will be swapped with this value (`String`).

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`incoming_port` - (Optional) The port on which the request is received. See [Simple Route Incoming Port ](#simple-route-incoming-port) below for details.

`origin_pools` - (Required) Origin Pools for this route. See [Simple Route Origin Pools ](#simple-route-origin-pools) below for details.

`path` - (Required) URI path of route. See [Simple Route Path ](#simple-route-path) below for details.

`query_params` - (Optional) Handling of incoming query parameters in simple route.. See [Simple Route Query Params ](#simple-route-query-params) below for details.

### Choice Site

Advertise on a customer site and a given network..

`ip` - (Optional) Use given IP address as VIP on the site (`String`).

`ipv6` - (Optional) Use given IPv6 address as VIP on the site (`String`).

`network` - (Required) By default VIP chosen as ip address of primary network interface in the network (`String`).

`site` - (Required) Reference to site object. See [ref](#ref) below for details.

### Choice Site Segment

Advertise on a segment on a site.

`ip` - (Required) Use given IP address as VIP on the site (`String`).

`ipv6` - (Optional) Use given IPv6 address as VIP on the site (`String`).

`segment` - (Required) x-required. See [ref](#ref) below for details.

`site` - (Required) x-required. See [ref](#ref) below for details.

### Choice Virtual Network

Advertise on a virtual network.

###### One of the arguments from this list "default_v6_vip, specific_v6_vip" can be set

`default_v6_vip` - (Optional) Use the default VIP, system allocated or configured in the virtual network (`Bool`).

`specific_v6_vip` - (Optional) Use given IPV6 address as VIP on virtual Network (`String`).

###### One of the arguments from this list "default_vip, specific_vip" can be set

`default_vip` - (Optional) Use the default VIP, system allocated or configured in the virtual network (`Bool`).

`specific_vip` - (Optional) Use given IPV4 address as VIP on virtual Network (`String`).

`virtual_network` - (Required) Select network reference. See [ref](#ref) below for details.

### Choice Virtual Site

Advertise on a customer virtual site and a given network..

`network` - (Required) IP address of primary network interface in the network (`String`).

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Choice Virtual Site Segment

Advertise on a segment on a virtual site.

`ip` - (Required) Use given IP address as VIP on the site (`String`).

`ipv6` - (Optional) Use given IPv6 address as VIP on the site (`String`).

`segment` - (Required) x-required. See [ref](#ref) below for details.

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Choice Virtual Site With Vip

Advertise on a customer virtual site and a given network and IP..

`ip` - (Optional) Use given IP address as VIP on the site (`String`).

`ipv6` - (Optional) Use given IPv6 address as VIP on the site (`String`).

`network` - (Required) IP address of primary network interface in the network (`String`).

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Choice Vk8s Service

Advertise on vK8s Service Network on RE..

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Client Choice Any Client

Any Client.

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

### Cluster Retract Choice Do Not Retract Cluster

configuration..

### Cluster Retract Choice Retract Cluster

for route.

### Coalescing Choice Default Coalescing

or Cipher suite configuration.

### Coalescing Choice Strict Coalescing

and/or Cipher suite configuration.

### Condition Type Choice Api Endpoint

The API endpoint (Path + Method) which this validation applies to.

`methods` - (Optional) Methods to be matched (`List of Strings`).

`path` - (Required) Path to be matched (`String`).

### Cookie Tampering Disable Tampering Protection

x-displayName: "Disable".

### Cookie Tampering Enable Tampering Protection

x-displayName: "Enable".

### Count By Choice Use Http Lb User Id

Defined in HTTP-LB Security Configuration -> User Identifier..

### Credentials Choice Api Key

x-displayName: "API Key".

`key` - (Required) x-displayName: "Key" (`String`).

`value` - (Required) x-displayName: "Value". See [Api Key Value ](#api-key-value) below for details.

### Credentials Choice Basic Auth

x-displayName: "Basic Authentication".

`password` - (Required) x-displayName: "Password". See [Basic Auth Password ](#basic-auth-password) below for details.

`user` - (Required) x-displayName: "User" (`String`).

### Credentials Choice Bearer Token

x-displayName: "Bearer Token".

`token` - (Required) x-displayName: "Token". See [Bearer Token Token ](#bearer-token-token) below for details.

### Credentials Choice Login Endpoint

x-displayName: "Login Endpoint".

`json_payload` - (Required) Defines the structure of the API request payload, including payload structure, fields and values.. See [Login Endpoint Json Payload ](#login-endpoint-json-payload) below for details.

`method` - (Required) x-displayName: "Method" (`String`).

`path` - (Required) x-displayName: "Path" (`String`).

`token_response_key` - (Required) Specifies how to handle the API response, extracting authentication tokens. (`String`).

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

### Ddos Policy Choice Ddos Policy None

Do not apply additional service policy during an ongoing DDoS attack.

### Default Lb Choice Default Loadbalancer

x-displayName: "Yes".

### Default Lb Choice Non Default Loadbalancer

x-displayName: "No".

### Destination Type Any Url

Any URL .

### Destination Type Api Endpoint

The endpoint (path) of the request..

`methods` - (Optional) Methods to be matched (`List of Strings`).

`path` - (Required) Path to be matched (`String`).

### Destination Type Api Groups

Validation will be performed for the endpoints mentioned in the API Groups.

`api_groups` - (Required) x-required (`String`).

### Direct Response Route Headers

List of (key, value) headers.

`invert_match` - (Optional) Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Required) Name of the header (`String`).

###### One of the arguments from this list "exact, presence, regex" can be set

`exact` - (Optional) Header value to match exactly (`String`).

`presence` - (Optional) If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Direct Response Route Incoming Port

The port on which the request is received.

###### One of the arguments from this list "no_port_match, port, port_ranges" can be set

`no_port_match` - (Optional) Disable matching of ports (`Bool`).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Direct Response Route Path

URI path of route.

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Direct Response Route Route Direct Response

Send direct response.

`response_body` - (Optional) response body to send (`String`).

`response_code` - (Optional) response code to send (`Int`).

### Domain Choice Any Domain

The rule will apply for all domains..

### Domain Choice Ignore Domain

Ignore max age attribute.

### Domain Matcher Any Domain

x-displayName: "Any Domain".

### Domain Matcher Domain

x-displayName: "Domain".

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Domain Matcher Choice Any Domain

Any Domain..

### Domain Matcher Choice Domain

Domain matcher..

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Domains Credentials

Add credentials for API testing to use in the selected environment..

`credential_name` - (Required) Enter a unique name for the credentials used in API testing (`String`).

###### One of the arguments from this list "api_key, basic_auth, bearer_token, login_endpoint" must be set

`api_key` - (Optional) x-displayName: "API Key". See [Credentials Choice Api Key ](#credentials-choice-api-key) below for details.

`basic_auth` - (Optional) x-displayName: "Basic Authentication". See [Credentials Choice Basic Auth ](#credentials-choice-basic-auth) below for details.

`bearer_token` - (Optional) x-displayName: "Bearer Token". See [Credentials Choice Bearer Token ](#credentials-choice-bearer-token) below for details.

`login_endpoint` - (Optional) x-displayName: "Login Endpoint". See [Credentials Choice Login Endpoint ](#credentials-choice-login-endpoint) below for details.

###### One of the arguments from this list "admin, standard" must be set

`admin` - (Optional) x-displayName: "Admin" (`Bool`).

`standard` - (Optional) x-displayName: "Standard" (`Bool`).

### Domains Simple Login

Enter the username and password to assign credentials for the selected domain to crawl.

`password` - (Optional) Enter the password to assign credentials for the selected domain to crawl. See [Simple Login Password ](#simple-login-password) below for details.

`user` - (Optional) Enter the username to assign credentials for the selected domain to crawl (`String`).

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

`purge_duration_for_inactive_discovered_apis` - (Optional) Inactive discovered API will be deleted after configured duration. (`Int`).

### Exclude List Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Exclude List Path

URI path matcher..

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Expiry Choice Ignore Expiry

Ignore expiry attribute.

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

### Frequency Choice Every Day

x-displayName: "Every Day".

### Frequency Choice Every Month

x-displayName: "Every Month".

### Frequency Choice Every Week

x-displayName: "Every Week".

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

### Graphql Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Hash Policy Choice Common Hash Policy

Use load balancer hash policy for this route.

### Hash Policy Choice Cookie Stickiness

Consistent hashing algorithm, ring hash, is used to select origin server.

###### One of the arguments from this list "add_httponly, ignore_httponly" can be set

`add_httponly` - (Optional) Add httponly attribute (`Bool`).

`ignore_httponly` - (Optional) Ignore httponly attribute (`Bool`).

`name` - (Required) produced (`String`).

`path` - (Optional) will be set for the cookie (`String`).

###### One of the arguments from this list "ignore_samesite, samesite_lax, samesite_none, samesite_strict" can be set

`ignore_samesite` - (Optional) Ignore Samesite attribute (`Bool`).

`samesite_lax` - (Optional) Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests (`Bool`).

`samesite_none` - (Optional) Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests (`Bool`).

`samesite_strict` - (Optional) Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests (`Bool`).

###### One of the arguments from this list "add_secure, ignore_secure" can be set

`add_secure` - (Optional) Add secure attribute (`Bool`).

`ignore_secure` - (Optional) Ignore secure attribute (`Bool`).

`ttl` - (Optional) be a session cookie. TTL value is in milliseconds (`Int`).

### Hash Policy Choice Ring Hash

Request are sent to all eligible origin servers using hash of request based on hash policy. Consistent hashing algorithm, ring hash, is used to select origin server.

`hash_policy` - (Required) route the request. See [Ring Hash Hash Policy ](#ring-hash-hash-policy) below for details.

### Hash Policy Choice Specific Hash Policy

Configure hash policy specific for this route.

`hash_policy` - (Required) route the request. See [Specific Hash Policy Hash Policy ](#specific-hash-policy-hash-policy) below for details.

### Header Transformation Choice Default Header Transformation

Normalize the headers to lower case.

### Header Transformation Choice Legacy Header Transformation

Use old header transformation if configured earlier.

### Header Transformation Choice Preserve Case Header Transformation

Preserves the original case of headers without any modifications..

### Header Transformation Choice Proper Case Header Transformation

For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are”.

### Host Rewrite Params Auto Host Rewrite

Host header will be swapped with hostname of upstream host chosen by the cluster.

### Host Rewrite Params Disable Host Rewrite

Host header is not modified.

### Http Header Headers

List of HTTP header name and value pairs.

`invert_match` - (Optional) Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Required) Name of the header (`String`).

###### One of the arguments from this list "exact, presence, regex" can be set

`exact` - (Optional) Header value to match exactly (`String`).

`presence` - (Optional) If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Http Protocol Choice Http Protocol Enable V1 Only

Enable HTTP/1.1 for downstream connections.

`header_transformation` - (Optional) the stateful formatter will take effect, and the stateless formatter will be disregarded.. See [Http Protocol Enable V1 Only Header Transformation ](#http-protocol-enable-v1-only-header-transformation) below for details.

### Http Protocol Choice Http Protocol Enable V1 V2

Enable both HTTP/1.1 and HTTP/2 for downstream connections.

### Http Protocol Choice Http Protocol Enable V2 Only

Enable HTTP/2 for downstream connections.

### Http Protocol Enable V1 Only Header Transformation

the stateful formatter will take effect, and the stateless formatter will be disregarded..

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

### Httponly Add Httponly

Add httponly attribute.

### Httponly Ignore Httponly

Ignore httponly attribute.

### Httponly Choice Add Httponly

x-displayName: "Add".

### Httponly Choice Ignore Httponly

x-displayName: "Ignore".

### Https Coalescing Options

Options for coalescing TLS for multiple HTTPS Load Balancers.

###### One of the arguments from this list "default_coalescing, strict_coalescing" must be set

`default_coalescing` - (Optional) or Cipher suite configuration (`Bool`).

`strict_coalescing` - (Optional) and/or Cipher suite configuration (`Bool`).

### Https Http Protocol Options

HTTP protocol configuration options for downstream connections..

###### One of the arguments from this list "http_protocol_enable_v1_only, http_protocol_enable_v1_v2, http_protocol_enable_v2_only" must be set

`http_protocol_enable_v1_only` - (Optional) Enable HTTP/1.1 for downstream connections. See [Http Protocol Choice Http Protocol Enable V1 Only ](#http-protocol-choice-http-protocol-enable-v1-only) below for details.

`http_protocol_enable_v1_v2` - (Optional) Enable both HTTP/1.1 and HTTP/2 for downstream connections (`Bool`).

`http_protocol_enable_v2_only` - (Optional) Enable HTTP/2 for downstream connections (`Bool`).

### Https Auto Cert Coalescing Options

Options for coalescing TLS for multiple HTTPS Load Balancers.

###### One of the arguments from this list "default_coalescing, strict_coalescing" must be set

`default_coalescing` - (Optional) or Cipher suite configuration (`Bool`).

`strict_coalescing` - (Optional) and/or Cipher suite configuration (`Bool`).

### Https Auto Cert Http Protocol Options

HTTP protocol configuration options for downstream connections..

###### One of the arguments from this list "http_protocol_enable_v1_only, http_protocol_enable_v1_v2, http_protocol_enable_v2_only" must be set

`http_protocol_enable_v1_only` - (Optional) Enable HTTP/1.1 for downstream connections. See [Http Protocol Choice Http Protocol Enable V1 Only ](#http-protocol-choice-http-protocol-enable-v1-only) below for details.

`http_protocol_enable_v1_v2` - (Optional) Enable both HTTP/1.1 and HTTP/2 for downstream connections (`Bool`).

`http_protocol_enable_v2_only` - (Optional) Enable HTTP/2 for downstream connections (`Bool`).

### Https Auto Cert Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

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

x-displayName: "Enable".

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

### Javascript Tags Tag Attributes

Add the tag attributes you want to include in your Javascript tag..

`javascript_tag` - (Optional) Select from one of the predefined tag attibutes. (`String`).

`tag_value` - (Optional) Add the tag attribute value. (`String`).

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

HTTP Load Balancer..

`dns_volterra_managed` - (Optional) or a DNS CNAME record should be created in your DNS provider's portal. (`Bool`).

###### One of the arguments from this list "port, port_ranges" must be set

`port` - (Optional) HTTP port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

### Loadbalancer Type Https

User is responsible for managing DNS to this load balancer..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`coalescing_options` - (Optional) Options for coalescing TLS for multiple HTTPS Load Balancers. See [Https Coalescing Options ](#https-coalescing-options) below for details.

`connection_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 2 minutes. (`Int`).

###### One of the arguments from this list "default_loadbalancer, non_default_loadbalancer" can be set

`default_loadbalancer` - (Optional) x-displayName: "Yes" (`Bool`).

`non_default_loadbalancer` - (Optional) x-displayName: "No" (`Bool`).

`http_protocol_options` - (Optional) HTTP protocol configuration options for downstream connections.. See [Https Http Protocol Options ](#https-http-protocol-options) below for details.

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

###### One of the arguments from this list "disable_path_normalize, enable_path_normalize" must be set

`disable_path_normalize` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_path_normalize` - (Optional) x-displayName: "Enable" (`Bool`).

###### One of the arguments from this list "port, port_ranges" must be set

`port` - (Optional) HTTPS port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

###### One of the arguments from this list "append_server_name, default_header, pass_through, server_name" can be set

`append_server_name` - (Optional) If header value is already present, it is not overwritten and passed as-is. (`String`).

`default_header` - (Optional) Response header name is “server” and value is “volt-adc” (`Bool`).

`pass_through` - (Optional) Pass existing server header as is. If server header is absent, a new header is not appended. (`Bool`).

`server_name` - (Optional) This will overwrite existing values, if any, for the server header. (`String`).

###### One of the arguments from this list "tls_cert_params, tls_parameters" must be set

`tls_cert_params` - (Optional) Select/Add one or more TLS Certificate objects to associate with this Load Balancer. See [Tls Certificates Choice Tls Cert Params ](#tls-certificates-choice-tls-cert-params) below for details.

`tls_parameters` - (Optional) Upload a TLS certificate covering all domain names for this Load Balancer. See [Tls Certificates Choice Tls Parameters ](#tls-certificates-choice-tls-parameters) below for details.

### Loadbalancer Type Https Auto Cert

or a DNS CNAME record should be created in your DNS provider's portal(only for Domains not managed by F5 Distributed Cloud)..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`coalescing_options` - (Optional) Options for coalescing TLS for multiple HTTPS Load Balancers. See [Https Auto Cert Coalescing Options ](#https-auto-cert-coalescing-options) below for details.

`connection_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 2 minutes. (`Int`).

###### One of the arguments from this list "default_loadbalancer, non_default_loadbalancer" can be set

`default_loadbalancer` - (Optional) For traffic terminating at this load balancer, the certificate associated with the first configured domain will be used for TLS termination. (`Bool`).

`non_default_loadbalancer` - (Optional) x-displayName: "No" (`Bool`).

`http_protocol_options` - (Optional) HTTP protocol configuration options for downstream connections.. See [Https Auto Cert Http Protocol Options ](#https-auto-cert-http-protocol-options) below for details.

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

###### One of the arguments from this list "disable_path_normalize, enable_path_normalize" must be set

`disable_path_normalize` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_path_normalize` - (Optional) x-displayName: "Enable" (`Bool`).

###### One of the arguments from this list "port, port_ranges" can be set

`port` - (Optional) HTTPS port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

###### One of the arguments from this list "append_server_name, default_header, pass_through, server_name" can be set

`append_server_name` - (Optional) If header value is already present, it is not overwritten and passed as-is. (`String`).

`default_header` - (Optional) Response header name is “server” and value is “volt-adc” (`Bool`).

`pass_through` - (Optional) Pass existing server header as is. If server header is absent, a new header is not appended. (`Bool`).

`server_name` - (Optional) This will overwrite existing values, if any, for the server header. (`String`).

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Https Auto Cert Tls Config ](#https-auto-cert-tls-config) below for details.

### Login Endpoint Json Payload

Defines the structure of the API request payload, including payload structure, fields and values..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Malicious User Mitigation Choice Default Mitigation Settings

For high level, users will be temporarily blocked..

### Malware Protection Malware Protection Settings

x-displayName: "Enable".

`malware_protection_rules` - (Required) Configure the match criteria to trigger Malware Protection Scan. See [Malware Protection Settings Malware Protection Rules ](#malware-protection-settings-malware-protection-rules) below for details.

### Malware Protection Rules Action

Report will identify and log threats, whereas Block will both log and block threats.

###### One of the arguments from this list "block, report" must be set

`block` - (Optional) Block the request and report the issue (`Bool`).

`report` - (Optional) Allow the request and report the issue (`Bool`).

### Malware Protection Rules Domain

Domain to be matched.

###### One of the arguments from this list "any_domain, domain" must be set

`any_domain` - (Optional) x-displayName: "Any Domain". See [Domain Matcher Any Domain ](#domain-matcher-any-domain) below for details.

`domain` - (Optional) x-displayName: "Domain". See [Domain Matcher Domain ](#domain-matcher-domain) below for details.

### Malware Protection Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Malware Protection Rules Path

Path to be matched.

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Malware Protection Settings Malware Protection Rules

Configure the match criteria to trigger Malware Protection Scan.

`action` - (Optional) Report will identify and log threats, whereas Block will both log and block threats. See [Malware Protection Rules Action ](#malware-protection-rules-action) below for details.

`domain` - (Optional) Domain to be matched. See [Malware Protection Rules Domain ](#malware-protection-rules-domain) below for details.

`http_methods` - (Optional) Methods to be matched (`List of Strings`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Malware Protection Rules Metadata ](#malware-protection-rules-metadata) below for details.

`path` - (Optional) Path to be matched. See [Malware Protection Rules Path ](#malware-protection-rules-path) below for details.

### Match Check Not Present

Check that the cookie is not present..

### Match Check Present

Check that the cookie is present..

### Match Item

Criteria for matching the values for the cookie. The match is successful if any of the values in the input satisfies the criteria in the matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Max Age Choice Ignore Max Age

Ignore max age attribute.

### Method Choice Method Get

x-displayName: "GET".

### Method Choice Method Post

x-displayName: "POST".

### Mirror Policy Percent

Percentage of requests to be mirrored.

`denominator` - (Required) Samples per denominator. numerator part per 100 or 10000 ro 1000000 (`String`).

`numerator` - (Required) sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000 (`Int`).

### Mirroring Choice Disable Mirroring

Disable Mirroring of request.

### Mirroring Choice Mirror Policy

useful for logging. For example, *cluster1* becomes *cluster1-shadow*..

`origin_pool` - (Required) referred here must be present.. See [ref](#ref) below for details.

`percent` - (Required) Percentage of requests to be mirrored. See [Mirror Policy Percent ](#mirror-policy-percent) below for details.

### Mitigation Action Block

Block user for a duration determined by the expiration time.

### Mitigation Action Choice Mitigation Block

Block suspicious sources.

### Mitigation Action Choice Mitigation Js Challenge

Serve JavaScript challenge to suspicious sources.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

`js_script_delay` - (Optional) Delay introduced by Javascript, in milliseconds. (`Int`).

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

Mobile SDK configuration.

`mobile_identifier` - (Optional) Mobile traffic identifier type.. See [Mobile Sdk Config Mobile Identifier ](#mobile-sdk-config-mobile-identifier) below for details.

### Mobile Sdk Config Mobile Identifier

Mobile traffic identifier type..

`headers` - (Optional) Headers that can be used to identify mobile traffic.. See [Mobile Identifier Headers ](#mobile-identifier-headers) below for details.

### More Option Buffer Policy

specify the maximum buffer size and buffer interval with this config..

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`max_request_bytes` - (Optional) manager will stop buffering and return a RequestEntityTooLarge (413) response. (`Int`).

### More Option Compression Params

Only GZIP compression is supported.

`content_length` - (Optional) Minimum response length, in bytes, which will trigger compression. The default value is 30. (`Int`).

`content_type` - (Optional) "text/xml" (`String`).

`disable_on_etag_header` - (Optional) weak etags will be preserved and the ones that require strong validation will be removed. (`Bool`).

`remove_accept_encoding_header` - (Optional) so that responses do not get compressed before reaching the filter. (`Bool`).

### More Option Request Cookies To Add

Cookies specified at this level are applied after cookies from matched Route are applied.

`name` - (Required) Name of the cookie in Cookie header. (`String`).

`overwrite` - (Optional) Default value is do not overwrite (`Bool`).

###### One of the arguments from this list "secret_value, value" must be set

`secret_value` - (Optional) Secret Value of the Cookie header. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the Cookie header. (`String`).

### More Option Request Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "secret_value, value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### More Option Response Cookies To Add

Cookies specified at this level are applied after cookies from matched Route are applied.

###### One of the arguments from this list "add_domain, ignore_domain" can be set

`add_domain` - (Optional) Add domain attribute (`String`).

`ignore_domain` - (Optional) Ignore max age attribute (`Bool`).

###### One of the arguments from this list "add_expiry, ignore_expiry" can be set

`add_expiry` - (Optional) Add expiry attribute (`String`).

`ignore_expiry` - (Optional) Ignore expiry attribute (`Bool`).

###### One of the arguments from this list "add_httponly, ignore_httponly" can be set

`add_httponly` - (Optional) x-displayName: "Add" (`Bool`).

`ignore_httponly` - (Optional) x-displayName: "Ignore" (`Bool`).

###### One of the arguments from this list "ignore_max_age, max_age_value" can be set

`ignore_max_age` - (Optional) Ignore max age attribute (`Bool`).

`max_age_value` - (Optional) Add max age attribute (`Int`).

`name` - (Required) Name of the cookie in Cookie header. (`String`).

`overwrite` - (Optional) Default value is do not overwrite (`Bool`).

###### One of the arguments from this list "add_partitioned, ignore_partitioned" can be set

`add_partitioned` - (Optional) x-displayName: "Add" (`Bool`).

`ignore_partitioned` - (Optional) x-displayName: "Ignore" (`Bool`).

###### One of the arguments from this list "add_path, ignore_path" can be set

`add_path` - (Optional) Add path attribute (`String`).

`ignore_path` - (Optional) Ignore path attribute (`Bool`).

###### One of the arguments from this list "ignore_samesite, samesite_lax, samesite_none, samesite_strict" can be set

`ignore_samesite` - (Optional) Ignore Samesite attribute (`Bool`).

`samesite_lax` - (Optional) Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests (`Bool`).

`samesite_none` - (Optional) Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests (`Bool`).

`samesite_strict` - (Optional) Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests (`Bool`).

###### One of the arguments from this list "add_secure, ignore_secure" can be set

`add_secure` - (Optional) x-displayName: "Add" (`Bool`).

`ignore_secure` - (Optional) x-displayName: "Ignore" (`Bool`).

###### One of the arguments from this list "ignore_value, secret_value, value" can be set

`ignore_value` - (Optional) Ignore value of cookie (`Bool`).

`secret_value` - (Optional) Secret Value of the Cookie header. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the Cookie header. (`String`).

### More Option Response Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "secret_value, value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

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

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Open Api Validation Rules Validation Mode

When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger).

###### One of the arguments from this list "response_validation_mode_active, skip_response_validation" must be set

`response_validation_mode_active` - (Optional) Enforce OpenAPI validation processing for this event. See [Response Validation Mode Choice Response Validation Mode Active ](#response-validation-mode-choice-response-validation-mode-active) below for details.

`skip_response_validation` - (Optional) Skip OpenAPI validation processing for this event (`Bool`).

###### One of the arguments from this list "skip_validation, validation_mode_active" must be set

`skip_validation` - (Optional) Skip OpenAPI validation processing for this event (`Bool`).

`validation_mode_active` - (Optional) Enforce OpenAPI validation processing for this event. See [Validation Mode Choice Validation Mode Active ](#validation-mode-choice-validation-mode-active) below for details.

### Origin Server Subset Rule List Origin Server Subset Rules

When an Origin server subset rule is matched, then this selection rule takes effect and no more rules are evaluated..

###### One of the arguments from this list "any_asn, asn_list, asn_matcher" must be set

`any_asn` - (Optional) Any origin ASN. (`Bool`).

`asn_list` - (Optional) The predicate evaluates to true if the origin ASN is present in the ASN list.. See [Asn Choice Asn List ](#asn-choice-asn-list) below for details.

`asn_matcher` - (Optional) The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects.. See [Asn Choice Asn Matcher ](#asn-choice-asn-matcher) below for details.

`country_codes` - (Optional) List of Country Codes (`List of Strings`).

###### One of the arguments from this list "any_ip, ip_matcher, ip_prefix_list" must be set

`any_ip` - (Optional) Any Source IP (`Bool`).

`ip_matcher` - (Optional) The predicate evaluates to true if the client IPv4 Address is covered by one or more of the IPv4 Prefixes in the IP Prefix Sets.. See [Ip Choice Ip Matcher ](#ip-choice-ip-matcher) below for details.

`ip_prefix_list` - (Optional) The predicate evaluates to true if the client IPv4 Address is covered by one or more of the IPv4 Prefixes from the list.. See [Ip Choice Ip Prefix List ](#ip-choice-ip-prefix-list) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Origin Server Subset Rules Metadata ](#origin-server-subset-rules-metadata) below for details.

`origin_server_subsets_action` - (Required) 2. Enable subset load balancing in the Origin Server Subsets section and configure keys in origin server subsets classes (`String`).

`re_name_list` - (Optional) List of RE names for match (`String`).

###### One of the arguments from this list "client_selector, none" must be set

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Selector Choice Client Selector ](#selector-choice-client-selector) below for details.

`none` - (Optional) No Label Selector (`Bool`).

### Origin Server Subset Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Oversized Body Choice Oversized Body Fail Validation

Apply the request/response action (block or report) when the body length is too long to verify (default 64Kb).

### Oversized Body Choice Oversized Body Skip Validation

Skip body validation when the body length is too long to verify (default 64Kb).

### Partitioned Choice Add Partitioned

x-displayName: "Add".

### Partitioned Choice Ignore Partitioned

x-displayName: "Ignore".

### Path Choice Any Path

Match all paths.

### Path Choice Ignore Path

Ignore path attribute.

### Path Normalize Choice Disable Path Normalize

x-displayName: "Disable".

### Path Normalize Choice Enable Path Normalize

x-displayName: "Enable".

### Policy Protected App Endpoints

List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32 endpoints per LB' after 4 LBs..

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

### Policy Based Challenge Rule List

list challenge rules to be used in policy based challenge.

`rules` - (Optional) these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions. See [Rule List Rules ](#rule-list-rules) below for details.

### Policy Choice No Policies

Do not apply additional rate limiter policies..

### Policy Choice Policies

to the action configured in the rule. If there's no match, the rate limiting configuration for the HTTP load balancer is honored..

`policies` - (Required) Ordered list of rate limiter policies.. See [ref](#ref) below for details.

### Policy Specifier Cookie

Hash based on cookie.

###### One of the arguments from this list "add_httponly, ignore_httponly" can be set

`add_httponly` - (Optional) Add httponly attribute (`Bool`).

`ignore_httponly` - (Optional) Ignore httponly attribute (`Bool`).

`name` - (Required) produced (`String`).

`path` - (Optional) will be set for the cookie (`String`).

###### One of the arguments from this list "ignore_samesite, samesite_lax, samesite_none, samesite_strict" can be set

`ignore_samesite` - (Optional) Ignore Samesite attribute (`Bool`).

`samesite_lax` - (Optional) Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests (`Bool`).

`samesite_none` - (Optional) Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests (`Bool`).

`samesite_strict` - (Optional) Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests (`Bool`).

###### One of the arguments from this list "add_secure, ignore_secure" can be set

`add_secure` - (Optional) Add secure attribute (`Bool`).

`ignore_secure` - (Optional) Ignore secure attribute (`Bool`).

`ttl` - (Optional) be a session cookie. TTL value is in milliseconds (`Int`).

### Port Choice Use Default Port

Inherit the Load Balancer's Listen Port..

### Port Match No Port Match

Disable matching of ports.

### Property Validation Settings Choice Property Validation Settings Custom

Use custom settings with Open API specification validation.

`queryParameters` - (Optional) Custom settings for query parameters validation. See [Property Validation Settings Custom QueryParameters ](#property-validation-settings-custom-queryParameters) below for details.

### Property Validation Settings Choice Property Validation Settings Default

Keep the default settings of OpenAPI specification validation.

### Property Validation Settings Custom QueryParameters

Custom settings for query parameters validation.

###### One of the arguments from this list "allow_additional_parameters, disallow_additional_parameters" must be set

`allow_additional_parameters` - (Optional) Allow extra query parameters (on top of what specified in the OAS documentation) (`Bool`).

`disallow_additional_parameters` - (Optional) Disallow extra query parameters (on top of what specified in the OAS documentation) (`Bool`).

### Protected App Endpoints Headers

Note that all specified header predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the header is not present. (`Bool`).

`check_present` - (Optional) Check that the header is present. (`Bool`).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Match Item ](#match-item) below for details.

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Protected App Endpoints Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Protected App Endpoints Mitigation

Mitigation action..

###### One of the arguments from this list "block, flag, redirect" can be set

`block` - (Optional) Block bot request and send response with custom content.. See [Action Type Block ](#action-type-block) below for details.

`flag` - (Optional) Flag the request while not taking any invasive actions.. See [Action Type Flag ](#action-type-flag) below for details.

`redirect` - (Optional) Redirect bot request to a custom URI.. See [Action Type Redirect ](#action-type-redirect) below for details.

### Protected App Endpoints Path

Matching URI path of the route..

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Protected App Endpoints Query Params

Note that all specified query parameter predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`key` - (Required) A case-sensitive HTTP query parameter name. (`String`).

###### One of the arguments from this list "check_not_present, check_present, item" must be set

`check_not_present` - (Optional) Check that the query parameter is not present. (`Bool`).

`check_present` - (Optional) Check that the query parameter is present. (`Bool`).

`item` - (Optional) criteria in the matcher.. See [Match Item ](#match-item) below for details.

### Query Params Remove All Params

x-displayName: "Remove All Parameters".

### Query Params Retain All Params

x-displayName: "Retain All Parameters".

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

Define rate limiting for one or more API endpoints..

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

### Redirect Route Headers

List of (key, value) headers.

`invert_match` - (Optional) Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Required) Name of the header (`String`).

###### One of the arguments from this list "exact, presence, regex" can be set

`exact` - (Optional) Header value to match exactly (`String`).

`presence` - (Optional) If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Redirect Route Incoming Port

The port on which the request is received.

###### One of the arguments from this list "no_port_match, port, port_ranges" can be set

`no_port_match` - (Optional) Disable matching of ports (`Bool`).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Redirect Route Path

URI path of route.

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Redirect Route Route Redirect

Send redirect response.

`host_redirect` - (Optional) swap host part of incoming URL in redirect URL (`String`).

`proto_redirect` - (Optional) When incoming-proto option is specified, swapping of protocol is not done. (`String`).

###### One of the arguments from this list "remove_all_params, replace_params, retain_all_params" can be set

`remove_all_params` - (Optional) x-displayName: "Remove All Parameters" (`Bool`).

`replace_params` - (Optional) x-displayName: "Replace All Parameters" (`String`).

`retain_all_params` - (Optional) x-displayName: "Retain All Parameters" (`Bool`).

###### One of the arguments from this list "path_redirect, prefix_rewrite" can be set

`path_redirect` - (Optional) swap path part of incoming URL in redirect URL (`String`).

`prefix_rewrite` - (Optional) This option allows redirect URLs be dynamically created based on the request (`String`).

`response_code` - (Optional) The HTTP status code to use in the redirect response. (`Int`).

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

### Response Validation Mode Choice Response Validation Mode Active

Enforce OpenAPI validation processing for this event.

`response_validation_properties` - (Required) List of properties of the response to validate according to the OpenAPI specification file (a.k.a. swagger) (`List of Strings`).

###### One of the arguments from this list "enforcement_block, enforcement_report" must be set

`enforcement_block` - (Optional) Block the response, trigger an API security event (`Bool`).

`enforcement_report` - (Optional) Allow the response, trigger an API security event (`Bool`).

### Response Validation Mode Choice Skip Response Validation

Skip OpenAPI validation processing for this event.

### Retry Policy Back Off

10 times the base interval.

`base_interval` - (Optional) Specifies the base interval between retries in milliseconds (`Int`).

`max_interval` - (Optional) to the base_interval if set. The default is 10 times the base_interval. (`Int`).

### Retry Policy Choice Default Retry Policy

Use system default retry policy.

### Retry Policy Choice No Retry Policy

Do not configure retry policy.

### Retry Policy Choice Retry Policy

Configure custom retry policy.

`back_off` - (Optional) 10 times the base interval. See [Retry Policy Back Off ](#retry-policy-back-off) below for details.

`num_retries` - (Optional) is used between each retry (`Int`).

`per_try_timeout` - (Optional) Specifies a non-zero timeout per retry attempt. In milliseconds (`Int`).

`retriable_status_codes` - (Optional) HTTP status codes that should trigger a retry in addition to those specified by retry_on. (`Int`).

`retry_condition` - (Required) (disconnect/reset/read timeout.) (`String`).

### Rewrite Choice Disable Prefix Rewrite

Do not rewrite any path portion..

### Rewrite Choice Regex Rewrite

with the substitution value..

`pattern` - (Optional) The regular expression used to find portions of a string that should be replaced. (`String`).

`substitution` - (Optional) substitution operation to produce a new string. (`String`).

### Ring Hash Hash Policy

route the request.

###### One of the arguments from this list "cookie, header_name, source_ip" must be set

`cookie` - (Optional) Hash based on cookie. See [Policy Specifier Cookie ](#policy-specifier-cookie) below for details.

`header_name` - (Optional) The name or key of the request header that will be used to obtain the hash key (`String`).

`source_ip` - (Optional) Hash based on source IP address (`Bool`).

`terminal` - (Optional) Specify if its a terminal policy (`Bool`).

### Role Choice Admin

x-displayName: "Admin".

### Role Choice Standard

x-displayName: "Standard".

### Rule List Rules

these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Rules Metadata ](#rules-metadata) below for details.

`spec` - (Required) Specification for the rule including match predicates and actions.. See [Rules Spec ](#rules-spec) below for details.

### Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

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

### Samesite Choice Ignore Samesite

Ignore Samesite attribute.

### Samesite Choice Samesite Lax

Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests.

### Samesite Choice Samesite None

Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests.

### Samesite Choice Samesite Strict

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

Add secure attribute.

### Secure Ignore Secure

Ignore secure attribute.

### Secure Choice Add Secure

x-displayName: "Add".

### Secure Choice Ignore Secure

x-displayName: "Ignore".

### Selector Choice Client Selector

The predicate evaluates to true if the expressions in the label selector are true for the client labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Selector Choice None

No Label Selector.

### Send Headers Choice Append Headers

Append mitigation headers..

`auto_type_header_name` - (Required) A case-insensitive HTTP header name. (`String`).

`inference_header_name` - (Required) A case-insensitive HTTP header name. (`String`).

### Send Headers Choice No Headers

No mitigation headers..

### Sensitive Data Disclosure Rules Sensitive Data Types In Response

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses.

`body` - (Optional) x-displayName: "JSON Path". See [Sensitive Data Types In Response Body ](#sensitive-data-types-in-response-body) below for details.

###### One of the arguments from this list "api_endpoint" must be set

`api_endpoint` - (Optional) The API endpoint (Path + Method) which this validation applies to. See [Type Condition Type Choice Api Endpoint ](#type-condition-type-choice-api-endpoint) below for details.

### Sensitive Data Policy Choice Sensitive Data Policy

Apply custom sensitive data discovery.

`sensitive_data_policy_ref` - (Required) Specify Sensitive Data Discovery. See [ref](#ref) below for details.

### Sensitive Data Types In Response Body

x-displayName: "JSON Path".

`fields` - (Required) List of JSON Path field values. Use square brackets with an underscore \[*] to indicate array elements (e.g., person.emails\[*]). To reference JSON keys that contain spaces, enclose the entire path in double quotes. For example: "person.first name". (`String`).

### Server Header Choice Default Header

Response header name is “server” and value is “volt-adc”.

### Server Header Choice Pass Through

Pass existing server header as is. If server header is absent, a new header is not appended..

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

### Service Policy Choice Active Service Policies

Apply the specified list of service policies and bypass the namespace service policy set.

`policies` - (Required) If all policies are evaluated and none match, then the request will be denied by default.. See [ref](#ref) below for details.

### Simple Login Password

Enter the password to assign credentials for the selected domain to crawl.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Simple Route Advanced Options

Configure Advanced per route options.

###### One of the arguments from this list "bot_defense_javascript_injection, inherited_bot_defense_javascript_injection" can be set

`bot_defense_javascript_injection` - (Optional) Configuration for Bot Defense JavaScript Injection. See [Bot Defense Javascript Injection Choice Bot Defense Javascript Injection ](#bot-defense-javascript-injection-choice-bot-defense-javascript-injection) below for details.

`inherited_bot_defense_javascript_injection` - (Optional) Hence no custom configuration is applied on the route (`Bool`).

###### One of the arguments from this list "buffer_policy, common_buffering" must be set

`buffer_policy` - (Optional) Route level buffer configuration overrides any configuration at VirtualHost level.. See [Buffer Choice Buffer Policy ](#buffer-choice-buffer-policy) below for details.

`common_buffering` - (Optional) Use common buffering configuration (`Bool`).

###### One of the arguments from this list "do_not_retract_cluster, retract_cluster" must be set

`do_not_retract_cluster` - (Optional) configuration. (`Bool`).

`retract_cluster` - (Optional) for route (`Bool`).

`cors_policy` - (Optional) resources from a server at a different origin. See [Advanced Options Cors Policy ](#advanced-options-cors-policy) below for details.

`csrf_policy` - (Optional) Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.).. See [Advanced Options Csrf Policy ](#advanced-options-csrf-policy) below for details.

`disable_location_add` - (Optional) virtual-host level. This configuration is ignored on CE sites. (`Bool`).

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

###### One of the arguments from this list "common_hash_policy, specific_hash_policy" must be set

`common_hash_policy` - (Optional) Use load balancer hash policy for this route (`Bool`).

`specific_hash_policy` - (Optional) Configure hash policy specific for this route. See [Hash Policy Choice Specific Hash Policy ](#hash-policy-choice-specific-hash-policy) below for details.

###### One of the arguments from this list "disable_mirroring, mirror_policy" must be set

`disable_mirroring` - (Optional) Disable Mirroring of request (`Bool`).

`mirror_policy` - (Optional) useful for logging. For example, *cluster1* becomes *cluster1-shadow*.. See [Mirroring Choice Mirror Policy ](#mirroring-choice-mirror-policy) below for details.

`priority` - (Optional) Also, circuit-breaker configuration at destination cluster is chosen based on the route priority. (`String`).

`request_cookies_to_add` - (Optional) Cookies specified at this level are applied after cookies from matched Route are applied. See [Advanced Options Request Cookies To Add ](#advanced-options-request-cookies-to-add) below for details.

`request_cookies_to_remove` - (Optional) List of keys of Cookies to be removed from the HTTP request being sent towards upstream. (`String`).

`request_headers_to_add` - (Optional) Headers are key-value pairs to be added to HTTP request being routed towards upstream.. See [Advanced Options Request Headers To Add ](#advanced-options-request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_cookies_to_add` - (Optional) Cookies specified at this level are applied after cookies from matched Route are applied. See [Advanced Options Response Cookies To Add ](#advanced-options-response-cookies-to-add) below for details.

`response_cookies_to_remove` - (Optional) List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed (`String`).

`response_headers_to_add` - (Optional) Headers are key-value pairs to be added to HTTP response being sent towards downstream.. See [Advanced Options Response Headers To Add ](#advanced-options-response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

###### One of the arguments from this list "default_retry_policy, no_retry_policy, retry_policy" must be set

`default_retry_policy` - (Optional) Use system default retry policy (`Bool`).

`no_retry_policy` - (Optional) Do not configure retry policy (`Bool`).

`retry_policy` - (Optional) Configure custom retry policy. See [Retry Policy Choice Retry Policy ](#retry-policy-choice-retry-policy) below for details.

###### One of the arguments from this list "disable_prefix_rewrite, prefix_rewrite, regex_rewrite" must be set

`disable_prefix_rewrite` - (Optional) Do not rewrite any path portion. (`Bool`).

`prefix_rewrite` - (Optional) the query string) will be swapped with this value. (`String`).

`regex_rewrite` - (Optional) with the substitution value.. See [Rewrite Choice Regex Rewrite ](#rewrite-choice-regex-rewrite) below for details.

###### One of the arguments from this list "disable_spdy, enable_spdy" must be set

`disable_spdy` - (Optional) SPDY upgrade is disabled (`Bool`).

`enable_spdy` - (Optional) SPDY upgrade is enabled (`Bool`).

`timeout` - (Optional) Should be set to a high value or 0 (infinite timeout) for server-side streaming. (`Int`).

###### One of the arguments from this list "app_firewall, disable_waf, inherited_waf" can be set

`app_firewall` - (Optional) Reference to App Firewall configuration object. See [ref](#ref) below for details.

`disable_waf` - (Optional) App Firewall configuration that is configured in the Load Balancer will not be enforced on this route (`Bool`).

`inherited_waf` - (Optional) Hence no custom configuration is applied on the route (`Bool`).

###### One of the arguments from this list "disable_web_socket_config, web_socket_config" must be set

`disable_web_socket_config` - (Optional) Websocket upgrade is disabled (`Bool`).

`web_socket_config` - (Optional) Upgrade to Websocket for this route. See [Websocket Choice Web Socket Config ](#websocket-choice-web-socket-config) below for details.

### Simple Route Headers

List of (key, value) headers.

`invert_match` - (Optional) Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Required) Name of the header (`String`).

###### One of the arguments from this list "exact, presence, regex" can be set

`exact` - (Optional) Header value to match exactly (`String`).

`presence` - (Optional) If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Simple Route Incoming Port

The port on which the request is received.

###### One of the arguments from this list "no_port_match, port, port_ranges" can be set

`no_port_match` - (Optional) Disable matching of ports (`Bool`).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Simple Route Origin Pools

Origin Pools for this route.

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

###### One of the arguments from this list "cluster, pool" must be set

`cluster` - (Optional) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Optional) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`priority` - (Optional) made active as per the increasing priority. (`Int`).

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

### Simple Route Path

URI path of route.

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Simple Route Query Params

Handling of incoming query parameters in simple route..

###### One of the arguments from this list "remove_all_params, replace_params, retain_all_params" must be set

`remove_all_params` - (Optional) x-displayName: "Remove All Parameters" (`Bool`).

`replace_params` - (Optional) x-displayName: "Replace All Parameters" (`String`).

`retain_all_params` - (Optional) x-displayName: "Retain All Parameters" (`Bool`).

### Slow Ddos Mitigation Choice Slow Ddos Mitigation

Custom Settings for Slow DDoS Mitigation.

`request_headers_timeout` - (Optional) provides protection against Slowloris attacks. (`Int`).

###### One of the arguments from this list "disable_request_timeout, request_timeout" must be set

`disable_request_timeout` - (Optional) x-displayName: "No Timeout" (`Bool`).

`request_timeout` - (Optional) x-example: "60000" (`Int`).

### Spdy Choice Disable Spdy

SPDY upgrade is disabled.

### Spdy Choice Enable Spdy

SPDY upgrade is enabled.

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

### Specific Hash Policy Hash Policy

route the request.

###### One of the arguments from this list "cookie, header_name, source_ip" must be set

`cookie` - (Optional) Hash based on cookie. See [Policy Specifier Cookie ](#policy-specifier-cookie) below for details.

`header_name` - (Optional) The name or key of the request header that will be used to obtain the hash key (`String`).

`source_ip` - (Optional) Hash based on source IP address (`Bool`).

`terminal` - (Optional) Specify if its a terminal policy (`Bool`).

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

### Tls Certificates Choice Tls Parameters

Upload a TLS certificate covering all domain names for this Load Balancer.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Parameters Tls Certificates ](#tls-parameters-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Parameters Tls Config ](#tls-parameters-tls-config) below for details.

### Tls Fingerprint Choice Tls Fingerprint Matcher

tls_fingerprint_matcher.

`classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`exact_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

`excluded_values` - (Optional) or more known TLS fingerprint classes in the enclosing matcher. (`String`).

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

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Token Location Bearer Token

Token is found in Authorization HTTP header with Bearer authentication scheme.

### Trust Client Ip Headers Choice Enable Trust Client Ip Headers

x-displayName: "Enable".

`client_ip_headers` - (Required) For X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client ip (`String`).

### Trusted Clients Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Type Condition Type Choice Api Endpoint

The API endpoint (Path + Method) which this validation applies to.

`methods` - (Optional) Methods to be matched (`List of Strings`).

`path` - (Required) Path to be matched (`String`).

### V6 Vip Choice Default V6 Vip

Use the default VIP, system allocated or configured in the virtual network.

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

`settings` - (Optional) OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom list" enforcement. See [Validation All Spec Endpoints Settings ](#validation-all-spec-endpoints-settings) below for details.

`validation_mode` - (Required) When a validation mismatch occurs on a request to one of the API Inventory endpoints. See [Validation All Spec Endpoints Validation Mode ](#validation-all-spec-endpoints-validation-mode) below for details.

### Validation Target Choice Validation Custom List

Any other end-points not listed will act according to "Fall Through Mode".

`fall_through_mode` - (Required) Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules). See [Validation Custom List Fall Through Mode ](#validation-custom-list-fall-through-mode) below for details.

`open_api_validation_rules` - (Required) x-displayName: "Validation List". See [Validation Custom List Open Api Validation Rules ](#validation-custom-list-open-api-validation-rules) below for details.

`settings` - (Optional) OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom list" enforcement. See [Validation Custom List Settings ](#validation-custom-list-settings) below for details.

### Validation Target Choice Validation Disabled

Don't run OpenAPI validation.

### Value Choice Ignore Value

Ignore value of cookie.

### Value Choice Secret Value

Secret Value of the Cookie header.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Vip Choice Default Vip

Use the default VIP, system allocated or configured in the virtual network.

### Waf Advanced Configuration App Firewall Detection Control

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria..

`exclude_attack_type_contexts` - (Optional) Attack Types to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Attack Type Contexts ](#app-firewall-detection-control-exclude-attack-type-contexts) below for details.

`exclude_bot_name_contexts` - (Optional) Bot Names to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Bot Name Contexts ](#app-firewall-detection-control-exclude-bot-name-contexts) below for details.

`exclude_signature_contexts` - (Optional) Signature IDs to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Signature Contexts ](#app-firewall-detection-control-exclude-signature-contexts) below for details.

`exclude_violation_contexts` - (Optional) Violations to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Violation Contexts ](#app-firewall-detection-control-exclude-violation-contexts) below for details.

### Waf Advanced Configuration Waf Skip Processing

Skip all App Firewall processing for this request.

### Waf Choice Disable Waf

App Firewall configuration that is configured in the Load Balancer will not be enforced on this route.

### Waf Choice Inherited Waf

Hence no custom configuration is applied on the route.

### Waf Exclusion Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Websocket Choice Disable Web Socket Config

Websocket upgrade is disabled.

### Websocket Choice Web Socket Config

Upgrade to Websocket for this route.

`use_websocket` - (Optional) a WebSocket connection (`Bool`).

### Xfcc Header Xfcc Disabled

No X-Forwarded-Client-Cert header will be added.

### Xfcc Header Xfcc Options

X-Forwarded-Client-Cert header will be added with the configured fields.

`xfcc_header_elements` - (Required) X-Forwarded-Client-Cert header elements to be added to requests (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured http_loadbalancer.
-	`cname` - This is the hostname of the configured http_loadbalancer.
