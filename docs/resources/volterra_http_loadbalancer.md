---

page_title: "Volterra: http_loadbalancer"
description: "The http_loadbalancer allows CRUD of Http Loadbalancer resource on Volterra SaaS"

---

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

  do_not_advertise = true

  // One of the arguments from this list "api_definition api_definitions api_specification disable_api_definition" must be set

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

    discovered_api_settings {
      purge_duration_for_inactive_discovered_apis = "2"
    }

    // One of the arguments from this list "disable_learn_from_redirect_traffic enable_learn_from_redirect_traffic" must be set

    disable_learn_from_redirect_traffic = true
    sensitive_data_detection_rules {
      custom_sensitive_data_detection_rules {
        metadata {
          description = "Virtual Host for acmecorp website"

          disable = true

          name = "acmecorp-web"
        }

        sensitive_data_detection_config {
          // One of the arguments from this list "any_domain specific_domain" must be set

          any_domain = true

          // One of the arguments from this list "key_pattern key_value_pattern value_pattern" must be set

          key_pattern {
            // One of the arguments from this list "exact_value regex_value" must be set

            exact_value = "x-volt-header"
          }

          // One of the arguments from this list "all_request_sections all_response_sections all_sections custom_sections" must be set

          all_sections = true

          // One of the arguments from this list "any_target api_endpoint_target api_group base_path" must be set

          api_group = "oas-all-operations"
        }

        sensitive_data_type {
          type = "EMAIL"
        }
      }

      disabled_built_in_rules {
        name = "[EMAIL, CC]"
      }
    }
  }

  // One of the arguments from this list "captcha_challenge enable_challenge js_challenge no_challenge policy_based_challenge" must be set

  no_challenge = true
  domains = ["www.foo.com"]

  // One of the arguments from this list "cookie_stickiness least_active random ring_hash round_robin source_ip_stickiness" must be set

  round_robin = true

  // One of the arguments from this list "l7_ddos_action_block l7_ddos_action_default l7_ddos_action_js_challenge l7_ddos_action_none" must be set

  l7_ddos_action_none = true

  // One of the arguments from this list "http https https_auto_cert" must be set

  http {
    dns_volterra_managed = true

    // One of the arguments from this list "port port_ranges" must be set

    port = "80"
  }

  // One of the arguments from this list "disable_malicious_user_detection enable_malicious_user_detection" must be set

  enable_malicious_user_detection = true

  // One of the arguments from this list "api_rate_limit disable_rate_limit rate_limit" must be set

  disable_rate_limit = true

  // One of the arguments from this list "default_sensitive_data_policy sensitive_data_policy" must be set

  default_sensitive_data_policy = true

  // One of the arguments from this list "active_service_policies no_service_policies service_policies_from_namespace" must be set

  no_service_policies = true

  // One of the arguments from this list "disable_threat_mesh enable_threat_mesh" must be set

  disable_threat_mesh = true

  // One of the arguments from this list "disable_trust_client_ip_headers enable_trust_client_ip_headers" must be set

  disable_trust_client_ip_headers = true

  // One of the arguments from this list "user_id_client_ip user_identification" must be set

  user_identification {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }

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

###### One of the arguments from this list "api_definition, api_definitions, api_specification, disable_api_definition" must be set

`api_definition` - (Optional) DEPRECATED by 'api_specification'. See [ref](#ref) below for details.(Deprecated)

`api_definitions` - (Optional) DEPRECATED by 'api_definition'. See [Api Definition Choice Api Definitions ](#api-definition-choice-api-definitions) below for details.(Deprecated)

`api_specification` - (Optional) Specify API definition and OpenAPI Validation. See [Api Definition Choice Api Specification ](#api-definition-choice-api-specification) below for details.

`disable_api_definition` - (Optional) API Definition is not currently used for this load balancer (`Bool`).

###### One of the arguments from this list "disable_api_discovery, enable_api_discovery" must be set

`disable_api_discovery` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_api_discovery` - (Optional) x-displayName: "Enable". See [Api Discovery Choice Enable Api Discovery ](#api-discovery-choice-enable-api-discovery) below for details.

`api_protection_rules` - (Optional) Rules can also include additional conditions, for example specific clients can access certain API endpoint or API group.. See [Api Protection Rules ](#api-protection-rules) below for details.

`api_rate_limit_legacy` - (Optional) Legacy value only temporary pre-migration. This value will be copied over to api_rate_limit and removed later.. See [Api Rate Limit Legacy ](#api-rate-limit-legacy) below for details.(Deprecated)

`blocked_clients` - (Optional) Define rules to block IP Prefixes or AS numbers.. See [Blocked Clients ](#blocked-clients) below for details.

###### One of the arguments from this list "bot_defense, bot_defense_advanced, disable_bot_defense" can be set

`bot_defense` - (Optional) Select Bot Defense Standard. See [Bot Defense Choice Bot Defense ](#bot-defense-choice-bot-defense) below for details.

`bot_defense_advanced` - (Optional) Select Bot Defense Advanced. See [Bot Defense Choice Bot Defense Advanced ](#bot-defense-choice-bot-defense-advanced) below for details.(Deprecated)

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

###### One of the arguments from this list "l7_ddos_action_block, l7_ddos_action_default, l7_ddos_action_js_challenge, l7_ddos_action_none" must be set

`l7_ddos_action_block` - (Optional) Block suspicious sources (`Bool`).

`l7_ddos_action_default` - (Optional) Block suspicious sources (`Bool`).

`l7_ddos_action_js_challenge` - (Optional) Serve JavaScript challenge to suspicious sources. See [L7 Ddos Auto Mitigation Action L7 Ddos Action Js Challenge ](#l7-ddos-auto-mitigation-action-l7-ddos-action-js-challenge) below for details.

`l7_ddos_action_none` - (Optional) Disable auto mitigation (`Bool`).(Deprecated)

###### One of the arguments from this list "http, https, https_auto_cert" must be set

`http` - (Optional) HTTP Load Balancer.. See [Loadbalancer Type Http ](#loadbalancer-type-http) below for details.

`https` - (Optional) User is responsible for managing DNS to this load balancer.. See [Loadbalancer Type Https ](#loadbalancer-type-https) below for details.

`https_auto_cert` - (Optional) or a DNS CNAME record should be created in your DNS provider's portal(only for Domains not managed by F5 Distributed Cloud).. See [Loadbalancer Type Https Auto Cert ](#loadbalancer-type-https-auto-cert) below for details.

###### One of the arguments from this list "disable_malicious_user_detection, enable_malicious_user_detection" must be set

`disable_malicious_user_detection` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_malicious_user_detection` - (Optional) x-displayName: "Enable" (`Bool`).

`malicious_user_mitigation` - (Optional) The settings defined in malicious user mitigation specify what mitigation actions to take for users determined to be at different threat levels.. See [ref](#ref) below for details.(Deprecated)

###### One of the arguments from this list "multi_lb_app, single_lb_app" can be set

`multi_lb_app` - (Optional) It should be configured externally using app type feature and label should be added to the HTTP load balancer. (`Bool`).(Deprecated)

`single_lb_app` - (Optional) ML Config applied on this load balancer. See [Ml Config Choice Single Lb App ](#ml-config-choice-single-lb-app) below for details.(Deprecated)

`more_option` - (Optional) More options like header manipulation, compression etc.. See [More Option ](#more-option) below for details.

###### One of the arguments from this list "default_pool, default_pool_list" can be set

`default_pool` - (Optional) Single Origin Pool. See [Origin Pool Choice Default Pool ](#origin-pool-choice-default-pool) below for details.(Deprecated)

`default_pool_list` - (Optional) Multiple Origin Pools with weights and priorities. See [Origin Pool Choice Default Pool List ](#origin-pool-choice-default-pool-list) below for details.(Deprecated)

`origin_server_subset_rule_list` - (Optional) When an Origin server subset rule is matched, then this selection rule takes effect and no more rules are evaluated.. See [Origin Server Subset Rule List ](#origin-server-subset-rule-list) below for details.

`protected_cookies` - (Optional) Note: We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.. See [Protected Cookies ](#protected-cookies) below for details.

###### One of the arguments from this list "api_rate_limit, disable_rate_limit, rate_limit" must be set

`api_rate_limit` - (Optional) Define rate limiting for one or more API endpoints.. See [Rate Limit Choice Api Rate Limit ](#rate-limit-choice-api-rate-limit) below for details.

`disable_rate_limit` - (Optional) Rate limiting is not currently enabled for this load balancer (`Bool`).

`rate_limit` - (Optional) Define custom rate limiting parameters for this load balancer. See [Rate Limit Choice Rate Limit ](#rate-limit-choice-rate-limit) below for details.

`routes` - (Optional) to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic. See [Routes ](#routes) below for details.

`sensitive_data_disclosure_rules` - (Optional) Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses. See [Sensitive Data Disclosure Rules ](#sensitive-data-disclosure-rules) below for details.(Deprecated)

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

### Api Rate Limit Legacy

Legacy value only temporary pre-migration. This value will be copied over to api_rate_limit and removed later..

`api_endpoint_rules` - (Optional) For creating rule that contain a whole domain or group of endpoints, please use the server URL rules above.. See [Api Rate Limit Legacy Api Endpoint Rules ](#api-rate-limit-legacy-api-endpoint-rules) below for details.

###### One of the arguments from this list "bypass_rate_limiting_rules, custom_ip_allowed_list, ip_allowed_list, no_ip_allowed_list" must be set

`bypass_rate_limiting_rules` - (Optional) This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting.. See [Ip Allowed List Choice Bypass Rate Limiting Rules ](#ip-allowed-list-choice-bypass-rate-limiting-rules) below for details.

`custom_ip_allowed_list` - (Optional) IP Allowed list using existing ip_prefix_set objects.. See [Ip Allowed List Choice Custom Ip Allowed List ](#ip-allowed-list-choice-custom-ip-allowed-list) below for details.

`ip_allowed_list` - (Optional) List of IP(s) for which rate limiting will be disabled.. See [Ip Allowed List Choice Ip Allowed List ](#ip-allowed-list-choice-ip-allowed-list) below for details.

`no_ip_allowed_list` - (Optional) There is no ip allowed list for rate limiting, all clients go through rate limiting. (`Bool`).

`server_url_rules` - (Optional) For matching also specific endpoints you can use the API endpoint rules set bellow.. See [Api Rate Limit Legacy Server Url Rules ](#api-rate-limit-legacy-server-url-rules) below for details.

### Blocked Clients

Define rules to block IP Prefixes or AS numbers..

###### One of the arguments from this list "bot_skip_processing, skip_processing, waf_skip_processing" can be set

`bot_skip_processing` - (Optional) Skip Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`skip_processing` - (Optional) Skip both WAF and Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (`Bool`).(Deprecated)

`actions` - (Optional) Actions that should be taken when client identifier matches the rule (`List of Strings`).

###### One of the arguments from this list "as_number, http_header, ip_prefix, user_identifier" must be set

`as_number` - (Optional) RFC 6793 defined 4-byte AS number (`Int`).

`http_header` - (Optional) Request header name and value pairs. See [Client Source Choice Http Header ](#client-source-choice-http-header) below for details.

`ip_prefix` - (Optional) IPv4 prefix string. (`String`).

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

`ip_prefix_list` - (Optional) IPv4 prefix string.. See [Mitigation Choice Ip Prefix List ](#mitigation-choice-ip-prefix-list) below for details.

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

`buffer_policy` - (Optional) specify the maximum buffer size and buffer interval with this config.. See [More Option Buffer Policy ](#more-option-buffer-policy) below for details.

`compression_params` - (Optional) Only GZIP compression is supported. See [More Option Compression Params ](#more-option-compression-params) below for details.

`cookies_to_modify` - (Optional) List of cookies to be modified from the HTTP response being sent towards downstream.. See [More Option Cookies To Modify ](#more-option-cookies-to-modify) below for details.(Deprecated)

`custom_errors` - (Optional) matches for a request. (`String`).

`disable_default_error_pages` - (Optional) Disable the use of default F5XC error pages. (`Bool`).

`idle_timeout` - (Optional) received, otherwise the stream is reset. (`Int`).

`javascript_info` - (Optional) Custom JavaScript Configuration. Custom JavaScript code can be executed at various stages of request processing.. See [More Option Javascript Info ](#more-option-javascript-info) below for details.(Deprecated)

`jwt` - (Optional) audiences and issuer. See [ref](#ref) below for details.(Deprecated)

`max_request_header_size` - (Optional) such load balancers is used for all the load balancers in question. (`Int`).

###### One of the arguments from this list "disable_path_normalize, enable_path_normalize" can be set

`disable_path_normalize` - (Optional) x-displayName: "Disable" (`Bool`).(Deprecated)

`enable_path_normalize` - (Optional) x-displayName: "Enable" (`Bool`).(Deprecated)

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [More Option Request Headers To Add ](#more-option-request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [More Option Response Headers To Add ](#more-option-response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

###### One of the arguments from this list "additional_domains, enable_strict_sni_host_header_check" can be set

`additional_domains` - (Optional) Wildcard names are supported in the suffix or prefix form. See [Strict Sni Host Header Check Choice Additional Domains ](#strict-sni-host-header-check-choice-additional-domains) below for details.(Deprecated)

`enable_strict_sni_host_header_check` - (Optional) Enable strict SNI and Host header check (`Bool`).(Deprecated)

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

### Routes

to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic.

###### One of the arguments from this list "custom_route_object, direct_response_route, redirect_route, simple_route" must be set

`custom_route_object` - (Optional) A custom route uses a route object created outside of this view.. See [Choice Custom Route Object ](#choice-custom-route-object) below for details.

`direct_response_route` - (Optional) A direct response route matches on path and/or HTTP method and responds directly to the matching traffic. See [Choice Direct Response Route ](#choice-direct-response-route) below for details.

`redirect_route` - (Optional) A redirect route matches on path and/or HTTP method and redirects the matching traffic to a different URL. See [Choice Redirect Route ](#choice-redirect-route) below for details.

`simple_route` - (Optional) A simple route matches on path and/or HTTP method and forwards the matching traffic to the associated pools. See [Choice Simple Route ](#choice-simple-route) below for details.

### Sensitive Data Disclosure Rules

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses.

`sensitive_data_types_in_response` - (Optional) Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses . See [Sensitive Data Disclosure Rules Sensitive Data Types In Response ](#sensitive-data-disclosure-rules-sensitive-data-types-in-response) below for details.

### Trusted Clients

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients..

###### One of the arguments from this list "bot_skip_processing, skip_processing, waf_skip_processing" can be set

`bot_skip_processing` - (Optional) Skip Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`skip_processing` - (Optional) Skip both WAF and Bot Defense processing for clients matching this rule. (`Bool`).(Deprecated)

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (`Bool`).(Deprecated)

`actions` - (Optional) Actions that should be taken when client identifier matches the rule (`List of Strings`).

###### One of the arguments from this list "as_number, http_header, ip_prefix, user_identifier" must be set

`as_number` - (Optional) RFC 6793 defined 4-byte AS number (`Int`).

`http_header` - (Optional) Request header name and value pairs. See [Client Source Choice Http Header ](#client-source-choice-http-header) below for details.

`ip_prefix` - (Optional) IPv4 prefix string. (`String`).

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

### Advanced Options Cors Policy

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

### Advanced Options Csrf Policy

Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)..

###### One of the arguments from this list "all_load_balancer_domains, custom_domain_list, disabled" must be set

`all_load_balancer_domains` - (Optional) Add All load balancer domains to source origin (allow) list. (`Bool`).

`custom_domain_list` - (Optional) Add one or more domains to source origin (allow) list.. See [Allowed Domains Custom Domain List ](#allowed-domains-custom-domain-list) below for details.

`disabled` - (Optional) Allow all source origin domains. (`Bool`).

### Advanced Options Header Transformation Type

Settings to normalize the headers of upstream requests..

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

### Advanced Options Request Headers To Add

Headers are key-value pairs to be added to HTTP request being routed towards upstream..

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "secret_value, value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

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

###### One of the arguments from this list "advertise_on_public, cloud_edge_segment, segment, site, site_segment, virtual_network, virtual_site, virtual_site_segment, virtual_site_with_vip, vk8s_service" must be set

`advertise_on_public` - (Optional) Advertise this load balancer on public network. See [Choice Advertise On Public ](#choice-advertise-on-public) below for details.

`site` - (Optional) Advertise on a customer site and a given network.. See [Choice Site ](#choice-site) below for details.

`site_segment` - (Optional) Advertise on a segment on a site. See [Choice Site Segment ](#choice-site-segment) below for details.

`virtual_network` - (Optional) Advertise on a virtual network. See [Choice Virtual Network ](#choice-virtual-network) below for details.

`virtual_site` - (Optional) Advertise on a customer virtual site and a given network.. See [Choice Virtual Site ](#choice-virtual-site) below for details.

`virtual_site_with_vip` - (Optional) Advertise on a customer virtual site and a given network and IP.. See [Choice Virtual Site With Vip ](#choice-virtual-site-with-vip) below for details.

`vk8s_service` - (Optional) Advertise on vK8s Service Network on RE.. See [Choice Vk8s Service ](#choice-vk8s-service) below for details.

###### One of the arguments from this list "port, port_ranges, use_default_port" must be set

`port` - (Optional) TCP port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

`use_default_port` - (Optional) For HTTP, default is 80. For HTTPS/SNI, default is 443. (`Bool`).

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

### Api Definition Choice Api Definitions

DEPRECATED by 'api_definition'.

`api_definitions` - (Optional) API Definitions using OpenAPI specification files. See [ref](#ref) below for details.

### Api Definition Choice Api Specification

Specify API definition and OpenAPI Validation.

`api_definition` - (Required) Specify API definition which includes application API paths and methods derived from swagger files.. See [ref](#ref) below for details.

###### One of the arguments from this list "validation_all_spec_endpoints, validation_custom_list, validation_disabled" must be set

`validation_all_spec_endpoints` - (Optional) All other API endpoints would proceed according to "Fall Through Mode". See [Validation Target Choice Validation All Spec Endpoints ](#validation-target-choice-validation-all-spec-endpoints) below for details.

`validation_custom_list` - (Optional) Any other end-points not listed will act according to "Fall Through Mode". See [Validation Target Choice Validation Custom List ](#validation-target-choice-validation-custom-list) below for details.

`validation_disabled` - (Optional) Don't run OpenAPI validation (`Bool`).

### Api Discovery Choice Disable Discovery

x-displayName: "Disable".

### Api Discovery Choice Enable Api Discovery

x-displayName: "Enable".

`api_discovery_from_code_scan` - (Optional) Select API code repositories to the load balancer to use them as a source for API endpoint discovery.. See [Enable Api Discovery Api Discovery From Code Scan ](#enable-api-discovery-api-discovery-from-code-scan) below for details.

`discovered_api_settings` - (Optional) Configure Discovered API Settings.. See [Enable Api Discovery Discovered Api Settings ](#enable-api-discovery-discovered-api-settings) below for details.

###### One of the arguments from this list "disable_learn_from_redirect_traffic, enable_learn_from_redirect_traffic" must be set

`disable_learn_from_redirect_traffic` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`enable_learn_from_redirect_traffic` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`sensitive_data_detection_rules` - (Optional) Manage rules to detect sensitive data in requests and/or response sections.. See [Enable Api Discovery Sensitive Data Detection Rules ](#enable-api-discovery-sensitive-data-detection-rules) below for details.(Deprecated)

### Api Discovery Choice Enable Discovery

x-displayName: "Enable".

`api_discovery_from_code_scan` - (Optional) Select API code repositories to the load balancer to use them as a source for API endpoint discovery.. See [Enable Discovery Api Discovery From Code Scan ](#enable-discovery-api-discovery-from-code-scan) below for details.

`discovered_api_settings` - (Optional) Configure Discovered API Settings.. See [Enable Discovery Discovered Api Settings ](#enable-discovery-discovered-api-settings) below for details.

###### One of the arguments from this list "disable_learn_from_redirect_traffic, enable_learn_from_redirect_traffic" must be set

`disable_learn_from_redirect_traffic` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`enable_learn_from_redirect_traffic` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`sensitive_data_detection_rules` - (Optional) Manage rules to detect sensitive data in requests and/or response sections.. See [Enable Discovery Sensitive Data Detection Rules ](#enable-discovery-sensitive-data-detection-rules) below for details.(Deprecated)

### Api Discovery From Code Scan Code Base Integrations

x-required.

###### One of the arguments from this list "all_repos, selected_repos" must be set

`all_repos` - (Optional) x-displayName: "All API Repositories" (`Bool`).

`selected_repos` - (Optional) x-displayName: "Selected API Repositories". See [Api Repos Choice Selected Repos ](#api-repos-choice-selected-repos) below for details.

`code_base_integration` - (Required) Select the code base integration for use in code-based API discovery. See [ref](#ref) below for details.

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

### Api Rate Limit Legacy Api Endpoint Rules

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

### Api Rate Limit Legacy Server Url Rules

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

`max_request_time` - (Optional) request before returning a RequestTimeout (408) response (`Int`).(Deprecated)

### Buffer Choice Common Buffering

Use common buffering configuration.

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

###### One of the arguments from this list "default_temporary_blocking_parameters, temporary_user_blocking" can be set

`default_temporary_blocking_parameters` - (Optional) Use default parameters (`Bool`).(Deprecated)

`temporary_user_blocking` - (Optional) Specifies configuration for temporary user blocking resulting from malicious user detection. See [Temporary Blocking Parameters Choice Temporary User Blocking ](#temporary-blocking-parameters-choice-temporary-user-blocking) below for details.(Deprecated)

### Choice Advertise On Public

Advertise this load balancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Choice Consul Service

Specify origin server with Hashi Corp Consul service name and site information.

###### One of the arguments from this list "inside_network, outside_network" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`service_name` - (Required) cluster-id is optional. (`String`).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Consul Service Site Locator ](#consul-service-site-locator) below for details.

### Choice Custom Endpoint Object

Specify origin server with a reference to endpoint object.

`endpoint` - (Required) Reference to an endpoint object. See [ref](#ref) below for details.

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

### Choice K8s Service

Specify origin server with K8s service name and site information.

###### One of the arguments from this list "inside_network, outside_network, vk8s_networks" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`vk8s_networks` - (Optional) origin server are on vK8s network on the site (`Bool`).

###### One of the arguments from this list "service_name, service_selector" must be set

`service_name` - (Optional) Both namespace and cluster-id are optional. (`String`).

`service_selector` - (Optional) discovery has to happen. This implicit label is added to service_selector. See [Service Info Service Selector ](#service-info-service-selector) below for details.(Deprecated)

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [K8s Service Site Locator ](#k8s-service-site-locator) below for details.

### Choice Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Choice Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### Choice Private Ip

Specify origin server with private or public IP address and site information.

###### One of the arguments from this list "inside_network, outside_network, segment" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`segment` - (Optional) Segment where this origin server is located. See [ref](#ref) below for details.

###### One of the arguments from this list "ip, ipv6" must be set

`ip` - (Optional) Private IPV4 address (`String`).

`ipv6` - (Optional) Private IPV6 address (`String`).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Private Ip Site Locator ](#private-ip-site-locator) below for details.

### Choice Private Name

Specify origin server with private or public DNS name and site information.

`dns_name` - (Required) DNS Name (`String`).

###### One of the arguments from this list "inside_network, outside_network, segment" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`segment` - (Optional) Segment where this origin server is located. See [ref](#ref) below for details.

`refresh_interval` - (Optional) Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767 (`Int`).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Private Name Site Locator ](#private-name-site-locator) below for details.

### Choice Public Ip

Specify origin server with public IP.

###### One of the arguments from this list "ip, ipv6" must be set

`ip` - (Optional) Public IPV4 address (`String`).

`ipv6` - (Optional) Public IPV6 address (`String`).

### Choice Public Name

Specify origin server with public DNS name.

`dns_name` - (Required) DNS Name (`String`).

`refresh_interval` - (Optional) Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767 (`Int`).

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

### Choice Vn Private Ip

Specify origin server IP address on virtual network other than inside or outside network.

`virtual_network` - (Required) Virtual Network where this IP will be present. See [ref](#ref) below for details.

###### One of the arguments from this list "ip, ipv6" must be set

`ip` - (Optional) IPV4 address (`String`).

`ipv6` - (Optional) IPV6 address (`String`).

### Choice Vn Private Name

Specify origin server name on virtual network other than inside or outside network.

`dns_name` - (Required) DNS Name (`String`).

`private_network` - (Required) Virtual Network where this Name will be present. See [ref](#ref) below for details.

### Circuit Breaker Choice Circuit Breaker

allows to apply back pressure on downstream quickly..

`connection_limit` - (Optional) Remove endpoint out of load balancing decision, if number of connections reach connection limit. (`Int`).

`max_requests` - (Optional) Remove endpoint out of load balancing decision, if requests exceed this count. (`Int`).

`pending_requests` - (Optional) Remove endpoint out of load balancing decision, if pending request reach pending_request. (`Int`).

`priority` - (Optional) matched with priority of CircuitBreaker to select the CircuitBreaker (`String`).

`retries` - (Optional) Remove endpoint out of load balancing decision, if retries for request exceed this count. (`Int`).

### Circuit Breaker Choice Default Circuit Breaker

requests are 1024 and the default value for retries is 3.

### Circuit Breaker Choice Disable Circuit Breaker

Circuit Breaker is disabled.

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

### Cluster Retract Choice Do Not Retract Cluster

configuration..

### Cluster Retract Choice Retract Cluster

for route.

### Condition Type Choice Api Endpoint

The API endpoint (Path + Method) which this validation applies to.

`methods` - (Optional) Methods to be matched (`List of Strings`).

`path` - (Required) Path to be matched (`String`).

### Consul Service Site Locator

Site or Virtual site where this origin server is located.

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

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

### Custom Sensitive Data Detection Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Custom Sensitive Data Detection Rules Sensitive Data Detection Config

The custom data detection config specifies targets, scopes & the pattern to be detected..

###### One of the arguments from this list "any_domain, specific_domain" must be set

`any_domain` - (Optional) The rule will apply for all domains. (`Bool`).(Deprecated)

`specific_domain` - (Optional) For example: api.example.com (`String`).(Deprecated)

###### One of the arguments from this list "key_pattern, key_value_pattern, value_pattern" must be set

`key_pattern` - (Optional) Search for pattern across all field names in the specified sections.. See [Pattern Choice Key Pattern ](#pattern-choice-key-pattern) below for details.

`key_value_pattern` - (Optional) Search for specific field and value patterns in the specified sections.. See [Pattern Choice Key Value Pattern ](#pattern-choice-key-value-pattern) below for details.

`value_pattern` - (Optional) Search for pattern across all field values in the specified sections.. See [Pattern Choice Value Pattern ](#pattern-choice-value-pattern) below for details.

###### One of the arguments from this list "all_request_sections, all_response_sections, all_sections, custom_sections" must be set

`all_request_sections` - (Optional) x-displayName: "All Request" (`Bool`).

`all_response_sections` - (Optional) x-displayName: "All Response" (`Bool`).

`all_sections` - (Optional) x-displayName: "All Request & Response" (`Bool`).

`custom_sections` - (Optional) x-displayName: "Custom Sections". See [Section Choice Custom Sections ](#section-choice-custom-sections) below for details.

###### One of the arguments from this list "any_target, api_endpoint_target, api_group, base_path" must be set

`any_target` - (Optional) The rule will be applied for all requests on this LB. (`Bool`).

`api_endpoint_target` - (Optional) The rule is applied only for the specified api endpoints.. See [Target Choice Api Endpoint Target ](#target-choice-api-endpoint-target) below for details.

`api_group` - (Optional) Custom groups can be created if user tags paths or operations with "x-volterra-api-group" extensions inside swaggers. (`String`).(Deprecated)

`base_path` - (Optional) The rule is applied only for the requests matching the specified base path. (`String`).(Deprecated)

### Custom Sensitive Data Detection Rules Sensitive Data Type

If the pattern is detected, the request is labeled with specified sensitive data type..

`type` - (Required) The request is labeled as specified sensitive data type. (`String`).

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

### Default Lb Choice Default Loadbalancer

x-displayName: "Yes".

### Default Lb Choice Non Default Loadbalancer

x-displayName: "No".

### Default Pool Advanced Options

Advanced options configuration like timeouts, circuit breaker, subset load balancing.

###### One of the arguments from this list "circuit_breaker, default_circuit_breaker, disable_circuit_breaker" must be set

`circuit_breaker` - (Optional) allows to apply back pressure on downstream quickly.. See [Circuit Breaker Choice Circuit Breaker ](#circuit-breaker-choice-circuit-breaker) below for details.

`default_circuit_breaker` - (Optional) requests are 1024 and the default value for retries is 3 (`Bool`).

`disable_circuit_breaker` - (Optional) Circuit Breaker is disabled (`Bool`).

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2 seconds (`Int`).

`header_transformation_type` - (Optional) Settings to normalize the headers of upstream requests.. See [Advanced Options Header Transformation Type ](#advanced-options-header-transformation-type) below for details.(Deprecated)

`http_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 5 minutes. (`Int`).

###### One of the arguments from this list "auto_http_config, http1_config, http2_options" must be set

`auto_http_config` - (Optional) and will use whichever protocol is negotiated by ALPN with the upstream. (`Bool`).

`http1_config` - (Optional) Enable HTTP/1.1 for upstream connections. See [Http Protocol Type Http1 Config ](#http-protocol-type-http1-config) below for details.

`http2_options` - (Optional) Enable HTTP/2 for upstream connections.. See [Http Protocol Type Http2 Options ](#http-protocol-type-http2-options) below for details.

###### One of the arguments from this list "disable_lb_source_ip_persistance, enable_lb_source_ip_persistance" can be set

`disable_lb_source_ip_persistance` - (Optional) Disable LB source IP persistence (`Bool`).

`enable_lb_source_ip_persistance` - (Optional) Enable LB source IP persistence (`Bool`).

###### One of the arguments from this list "disable_outlier_detection, outlier_detection" must be set

`disable_outlier_detection` - (Optional) Outlier detection is disabled (`Bool`).

`outlier_detection` - (Optional) healthy load balancing set. Outlier detection is a form of passive health checking.. See [Outlier Detection Choice Outlier Detection ](#outlier-detection-choice-outlier-detection) below for details.

###### One of the arguments from this list "no_panic_threshold, panic_threshold" must be set

`no_panic_threshold` - (Optional) Disable panic threshold. Only healthy endpoints are considered for load balancing. (`Bool`).

`panic_threshold` - (Optional) all endpoints will be considered for load balancing ignoring its health status. (`Int`).

###### One of the arguments from this list "disable_proxy_protocol, proxy_protocol_v1, proxy_protocol_v2" can be set

`disable_proxy_protocol` - (Optional) Disable Proxy Protocol for upstream connections (`Bool`).

`proxy_protocol_v1` - (Optional) Enable Proxy Protocol Version 1 for upstream connections (`Bool`).

`proxy_protocol_v2` - (Optional) Enable Proxy Protocol Version 2 for upstream connections (`Bool`).

###### One of the arguments from this list "disable_subsets, enable_subsets" must be set

`disable_subsets` - (Optional) Subset load balancing is disabled. All eligible origin servers will be considered for load balancing. (`Bool`).

`enable_subsets` - (Optional) Subset load balancing is enabled. Based on route, subset of origin servers will be considered for load balancing.. See [Subset Choice Enable Subsets ](#subset-choice-enable-subsets) below for details.

### Default Pool Origin Servers

List of origin servers in this pool.

###### One of the arguments from this list "consul_service, custom_endpoint_object, k8s_service, private_ip, private_name, public_ip, public_name, vn_private_ip, vn_private_name" must be set

`consul_service` - (Optional) Specify origin server with Hashi Corp Consul service name and site information. See [Choice Consul Service ](#choice-consul-service) below for details.

`custom_endpoint_object` - (Optional) Specify origin server with a reference to endpoint object. See [Choice Custom Endpoint Object ](#choice-custom-endpoint-object) below for details.

`k8s_service` - (Optional) Specify origin server with K8s service name and site information. See [Choice K8s Service ](#choice-k8s-service) below for details.

`private_ip` - (Optional) Specify origin server with private or public IP address and site information. See [Choice Private Ip ](#choice-private-ip) below for details.

`private_name` - (Optional) Specify origin server with private or public DNS name and site information. See [Choice Private Name ](#choice-private-name) below for details.

`public_ip` - (Optional) Specify origin server with public IP. See [Choice Public Ip ](#choice-public-ip) below for details.

`public_name` - (Optional) Specify origin server with public DNS name. See [Choice Public Name ](#choice-public-name) below for details.

`vn_private_ip` - (Optional) Specify origin server IP address on virtual network other than inside or outside network. See [Choice Vn Private Ip ](#choice-vn-private-ip) below for details.

`vn_private_name` - (Optional) Specify origin server name on virtual network other than inside or outside network. See [Choice Vn Private Name ](#choice-vn-private-name) below for details.

`labels` - (Optional) Add Labels for this origin server, these labels can be used to form subset. (`String`).

### Default Pool List Pools

List of Origin Pools.

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

###### One of the arguments from this list "cluster, pool" must be set

`cluster` - (Optional) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Optional) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`priority` - (Optional) made active as per the increasing priority. (`Int`).

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

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

### Domain Matcher Choice Any Domain

Any Domain..

### Domain Matcher Choice Domain

Domain matcher..

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Enable Api Discovery Api Discovery From Code Scan

Select API code repositories to the load balancer to use them as a source for API endpoint discovery..

`code_base_integrations` - (Required) x-required. See [Api Discovery From Code Scan Code Base Integrations ](#api-discovery-from-code-scan-code-base-integrations) below for details.

### Enable Api Discovery Discovered Api Settings

Configure Discovered API Settings..

`purge_duration_for_inactive_discovered_apis` - (Optional) Inactive discovered API will be deleted after configured duration. (`Int`).

### Enable Api Discovery Sensitive Data Detection Rules

Manage rules to detect sensitive data in requests and/or response sections..

`custom_sensitive_data_detection_rules` - (Optional) Rules to detect custom sensitive data in requests and/or responses sections.. See [Sensitive Data Detection Rules Custom Sensitive Data Detection Rules ](#sensitive-data-detection-rules-custom-sensitive-data-detection-rules) below for details.

`disabled_built_in_rules` - (Optional) List of disabled built-in sensitive data detection rules.. See [Sensitive Data Detection Rules Disabled Built In Rules ](#sensitive-data-detection-rules-disabled-built-in-rules) below for details.

### Enable Discovery Api Discovery From Code Scan

Select API code repositories to the load balancer to use them as a source for API endpoint discovery..

`code_base_integrations` - (Required) x-required. See [Api Discovery From Code Scan Code Base Integrations ](#api-discovery-from-code-scan-code-base-integrations) below for details.

### Enable Discovery Discovered Api Settings

Configure Discovered API Settings..

`purge_duration_for_inactive_discovered_apis` - (Optional) Inactive discovered API will be deleted after configured duration. (`Int`).

### Enable Discovery Sensitive Data Detection Rules

Manage rules to detect sensitive data in requests and/or response sections..

`custom_sensitive_data_detection_rules` - (Optional) Rules to detect custom sensitive data in requests and/or responses sections.. See [Sensitive Data Detection Rules Custom Sensitive Data Detection Rules ](#sensitive-data-detection-rules-custom-sensitive-data-detection-rules) below for details.

`disabled_built_in_rules` - (Optional) List of disabled built-in sensitive data detection rules.. See [Sensitive Data Detection Rules Disabled Built In Rules ](#sensitive-data-detection-rules-disabled-built-in-rules) below for details.

### Enable Subsets Endpoint Subsets

List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class..

`keys` - (Required) List of keys that define a cluster subset class. (`String`).

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

### Fallback Policy Choice Any Endpoint

Select any origin server from available healthy origin servers in this pool.

### Fallback Policy Choice Default Subset

Use the default subset provided here. Select endpoints matching default subset..

`default_subset` - (Optional) which gets used when route specifies no metadata or no subset matching the metadata exists. (`String`).

### Fallback Policy Choice Fail Request

Request will be failed and error returned, as if cluster has no origin servers..

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

### Health Check Port Choice Same As Endpoint Port

Health check is performed on endpoint port itself.

### Host Rewrite Params Auto Host Rewrite

Host header will be swapped with hostname of upstream host chosen by the cluster.

### Host Rewrite Params Disable Host Rewrite

Host header is not modified.

### Http1 Config Header Transformation

the stateful formatter will take effect, and the stateless formatter will be disregarded..

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

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

### Http Protocol Type Auto Http Config

and will use whichever protocol is negotiated by ALPN with the upstream..

### Http Protocol Type Http1 Config

Enable HTTP/1.1 for upstream connections.

`header_transformation` - (Optional) the stateful formatter will take effect, and the stateless formatter will be disregarded.. See [Http1 Config Header Transformation ](#http1-config-header-transformation) below for details.

### Http Protocol Type Http2 Options

Enable HTTP/2 for upstream connections..

`enabled` - (Optional) Enable/disable HTTP2 Protocol for upstream connections (`Bool`).

### Httponly Add Httponly

Add httponly attribute.

### Httponly Ignore Httponly

Ignore httponly attribute.

### Https Header Transformation Type

Header transformation options for response headers to the client.

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

### Https Http Protocol Options

HTTP protocol configuration options for downstream connections..

###### One of the arguments from this list "http_protocol_enable_v1_only, http_protocol_enable_v1_v2, http_protocol_enable_v2_only" must be set

`http_protocol_enable_v1_only` - (Optional) Enable HTTP/1.1 for downstream connections. See [Http Protocol Choice Http Protocol Enable V1 Only ](#http-protocol-choice-http-protocol-enable-v1-only) below for details.

`http_protocol_enable_v1_v2` - (Optional) Enable both HTTP/1.1 and HTTP/2 for downstream connections (`Bool`).

`http_protocol_enable_v2_only` - (Optional) Enable HTTP/2 for downstream connections (`Bool`).

### Https Auto Cert Header Transformation Type

Header transformation options for response headers to the client.

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

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

###### One of the arguments from this list "bearer_token, cookie, header, query_param" must be set

`bearer_token` - (Optional) Token is found in Authorization HTTP header with Bearer authentication scheme (`Bool`).

`cookie` - (Optional) Token is found in the cookie (`String`).(Deprecated)

`header` - (Optional) Token is found in the header (`String`).(Deprecated)

`query_param` - (Optional) Token is found in the query string parameter (`String`).(Deprecated)

### K8s Service Site Locator

Site or Virtual site where this origin server is located.

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Key Value Pattern Key Pattern

Pattern for key/field..

###### One of the arguments from this list "exact_value, regex_value" must be set

`exact_value` - (Optional) Search for values with exact match. (`String`).

`regex_value` - (Optional) Search for values matching this regular expression. (`String`).

### Key Value Pattern Value Pattern

Pattern for value..

###### One of the arguments from this list "exact_value, regex_value" must be set

`exact_value` - (Optional) Pattern value to be detected. (`String`).

`regex_value` - (Optional) Regular expression for this pattern. (`String`).

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

### Lb Source Ip Persistance Choice Disable Lb Source Ip Persistance

Disable LB source IP persistence.

### Lb Source Ip Persistance Choice Enable Lb Source Ip Persistance

Enable LB source IP persistence.

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

`connection_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 2 minutes. (`Int`).

###### One of the arguments from this list "default_loadbalancer, non_default_loadbalancer" can be set

`default_loadbalancer` - (Optional) x-displayName: "Yes" (`Bool`).

`non_default_loadbalancer` - (Optional) x-displayName: "No" (`Bool`).

`header_transformation_type` - (Optional) Header transformation options for response headers to the client. See [Https Header Transformation Type ](#https-header-transformation-type) below for details.(Deprecated)

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

`connection_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 2 minutes. (`Int`).

###### One of the arguments from this list "default_loadbalancer, non_default_loadbalancer" can be set

`default_loadbalancer` - (Optional) For traffic terminating at this load balancer, the certificate associated with the first configured domain will be used for TLS termination. (`Bool`).

`non_default_loadbalancer` - (Optional) x-displayName: "No" (`Bool`).

`header_transformation_type` - (Optional) Header transformation options for response headers to the client. See [Https Auto Cert Header Transformation Type ](#https-auto-cert-header-transformation-type) below for details.(Deprecated)

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

### Malicious User Detection Choice Disable Malicious User Detection

x-displayName: "Disable".

### Malicious User Detection Choice Enable Malicious User Detection

x-displayName: "Enable".

### Malicious User Mitigation Choice Default Mitigation Settings

For high level, users will be temporarily blocked..

### Masking Mode Choice Mask

x-displayName: "Mask Sensitive Data".

### Masking Mode Choice Report

x-displayName: "Report Sensitive Data".

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

### Mitigation Choice Ddos Client Source

Combination of Region, ASN and TLS Fingerprints.

`asn_list` - (Optional) The ASN is obtained by performing a lookup for the source IPv4 Address in a GeoIP DB.. See [Ddos Client Source Asn List ](#ddos-client-source-asn-list) below for details.

`country_list` - (Optional) Sources that are located in one of the countries in the given list (`List of Strings`).

`tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints.. See [Ddos Client Source Tls Fingerprint Matcher ](#ddos-client-source-tls-fingerprint-matcher) below for details.

### Mitigation Choice Ip Prefix List

IPv4 prefix string..

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

### Ml Config Choice Single Lb App

ML Config applied on this load balancer.

###### One of the arguments from this list "disable_discovery, enable_discovery" must be set

`disable_discovery` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_discovery` - (Optional) x-displayName: "Enable". See [Api Discovery Choice Enable Discovery ](#api-discovery-choice-enable-discovery) below for details.

###### One of the arguments from this list "disable_malicious_user_detection, enable_malicious_user_detection" must be set

`disable_malicious_user_detection` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_malicious_user_detection` - (Optional) x-displayName: "Enable" (`Bool`).

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

### More Option Buffer Policy

specify the maximum buffer size and buffer interval with this config..

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`max_request_bytes` - (Optional) manager will stop buffering and return a RequestEntityTooLarge (413) response. (`Int`).

`max_request_time` - (Optional) request before returning a RequestTimeout (408) response (`Int`).(Deprecated)

### More Option Compression Params

Only GZIP compression is supported.

`content_length` - (Optional) Minimum response length, in bytes, which will trigger compression. The default value is 30. (`Int`).

`content_type` - (Optional) "text/xml" (`String`).

`disable_on_etag_header` - (Optional) weak etags will be preserved and the ones that require strong validation will be removed. (`Bool`).

`remove_accept_encoding_header` - (Optional) so that responses do not get compressed before reaching the filter. (`Bool`).

### More Option Cookies To Modify

List of cookies to be modified from the HTTP response being sent towards downstream..

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

### More Option Javascript Info

Custom JavaScript Configuration. Custom JavaScript code can be executed at various stages of request processing..

`cache_prefix` - (Optional) KeyValue store referred by script. (`String`).

`custom_script_url` - (Optional) URL of JavaScript that gets executed (`String`).

`script_config` - (Optional) Input passed to the script (`String`).

### More Option Request Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "secret_value, value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

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

### Mtls Choice Use Mtls

x-displayName: "Upload a client authentication certificate specifically for this Origin Pool".

`tls_certificates` - (Required) mTLS Client Certificate. See [Use Mtls Tls Certificates ](#use-mtls-tls-certificates) below for details.

### Network Choice Inside Network

Inside network on the site.

### Network Choice Outside Network

Outside network on the site.

### Network Choice Vk8s Networks

origin server are on vK8s network on the site.

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

### Origin Pool Choice Default Pool

Single Origin Pool.

`advanced_options` - (Optional) Advanced options configuration like timeouts, circuit breaker, subset load balancing. See [Default Pool Advanced Options ](#default-pool-advanced-options) below for details.

`endpoint_selection` - (Required) Policy for selection of endpoints from local site or remote site or both (`String`).

###### One of the arguments from this list "health_check_port, same_as_endpoint_port" can be set

`health_check_port` - (Optional) Port used for performing health check (`Int`).

`same_as_endpoint_port` - (Optional) Health check is performed on endpoint port itself (`Bool`).

`healthcheck` - (Optional) Reference to healthcheck configuration objects. See [ref](#ref) below for details.

`loadbalancer_algorithm` - (Required) loadbalancer_algorithm to determine which host is selected. (`String`).

`origin_servers` - (Required) List of origin servers in this pool. See [Default Pool Origin Servers ](#default-pool-origin-servers) below for details.

###### One of the arguments from this list "automatic_port, lb_port, port" must be set

`automatic_port` - (Optional) For other origin server types, port will be automatically set as 443 if TLS is enabled at Origin Pool and 80 if TLS is disabled (`Bool`).

`lb_port` - (Optional) Endpoint port is selected based on loadbalancer port (`Bool`).

`port` - (Optional) Endpoint service is available on this port (`Int`).

###### One of the arguments from this list "no_tls, use_tls" must be set

`no_tls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_tls` - (Optional) x-displayName: "Enable". See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

`view_internal` - (Optional) Reference to view internal object. See [ref](#ref) below for details.

### Origin Pool Choice Default Pool List

Multiple Origin Pools with weights and priorities.

`pools` - (Optional) List of Origin Pools. See [Default Pool List Pools ](#default-pool-list-pools) below for details.

### Origin Server Subset Rule List Origin Server Subset Rules

When an Origin server subset rule is matched, then this selection rule takes effect and no more rules are evaluated..

###### One of the arguments from this list "any_asn, asn_list, asn_matcher" must be set

`any_asn` - (Optional) Any origin ASN. (`Bool`).

`asn_list` - (Optional) The predicate evaluates to true if the origin ASN is present in the ASN list.. See [Asn Choice Asn List ](#asn-choice-asn-list) below for details.

`asn_matcher` - (Optional) The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects.. See [Asn Choice Asn Matcher ](#asn-choice-asn-matcher) below for details.

`body_matcher` - (Optional) The actual request body value is extracted from the request API as a string.. See [Origin Server Subset Rules Body Matcher ](#origin-server-subset-rules-body-matcher) below for details.(Deprecated)

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

### Origin Server Subset Rules Body Matcher

The actual request body value is extracted from the request API as a string..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Origin Server Subset Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Outlier Detection Choice Disable Outlier Detection

Outlier detection is disabled.

### Outlier Detection Choice Outlier Detection

healthy load balancing set. Outlier detection is a form of passive health checking..

`base_ejection_time` - (Optional) Defaults to 30000ms or 30s. Specified in milliseconds. (`Int`).

`consecutive_5xx` - (Optional) a consecutive 5xx ejection occurs. Defaults to 5. (`Int`).

`consecutive_gateway_failure` - (Optional) before a consecutive gateway failure ejection occurs. Defaults to 5. (`Int`).

`interval` - (Optional) to 10000ms or 10s. Specified in milliseconds. (`Int`).

`max_ejection_percent` - (Optional) detection. Defaults to 10% but will eject at least one host regardless of the value. (`Int`).

### Oversized Body Choice Oversized Body Fail Validation

Apply the request/response action (block or report) when the body length is too long to verify (default 64Kb).

### Oversized Body Choice Oversized Body Skip Validation

Skip body validation when the body length is too long to verify (default 64Kb).

### Panic Threshold Type No Panic Threshold

Disable panic threshold. Only healthy endpoints are considered for load balancing..

### Path Choice Any Path

Match all paths.

### Path Normalize Choice Disable Path Normalize

x-displayName: "Disable".

### Path Normalize Choice Enable Path Normalize

x-displayName: "Enable".

### Pattern Choice Key Pattern

Search for pattern across all field names in the specified sections..

###### One of the arguments from this list "exact_value, regex_value" must be set

`exact_value` - (Optional) Search for values with exact match. (`String`).

`regex_value` - (Optional) Search for values matching this regular expression. (`String`).

### Pattern Choice Key Value Pattern

Search for specific field and value patterns in the specified sections..

`key_pattern` - (Required) Pattern for key/field.. See [Key Value Pattern Key Pattern ](#key-value-pattern-key-pattern) below for details.

`value_pattern` - (Required) Pattern for value.. See [Key Value Pattern Value Pattern ](#key-value-pattern-value-pattern) below for details.

### Pattern Choice Value Pattern

Search for pattern across all field values in the specified sections..

###### One of the arguments from this list "exact_value, regex_value" must be set

`exact_value` - (Optional) Pattern value to be detected. (`String`).

`regex_value` - (Optional) Regular expression for this pattern. (`String`).

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

`http_methods` - (Required) List of HTTP methods. (`List of Strings`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Protected App Endpoints Metadata ](#protected-app-endpoints-metadata) below for details.

`mitigation` - (Required) Mitigation action.. See [Protected App Endpoints Mitigation ](#protected-app-endpoints-mitigation) below for details.

`path` - (Required) Matching URI path of the route.. See [Protected App Endpoints Path ](#protected-app-endpoints-path) below for details.

`protocol` - (Optional) Protocol. (`String`).

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

### Port Choice Automatic Port

For other origin server types, port will be automatically set as 443 if TLS is enabled at Origin Pool and 80 if TLS is disabled.

### Port Choice Lb Port

Endpoint port is selected based on loadbalancer port.

### Port Choice Use Default Port

For HTTP, default is 80. For HTTPS/SNI, default is 443..

### Port Match No Port Match

Disable matching of ports.

### Private Ip Site Locator

Site or Virtual site where this origin server is located.

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Private Name Site Locator

Site or Virtual site where this origin server is located.

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

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

### Protected App Endpoints Request Body

Request Body.

`name` - (Optional) Enter request body parameter name (`String`).

###### One of the arguments from this list "exact_value, regex_value" must be set

`exact_value` - (Optional) Exact query value to match (`String`).

`regex_value` - (Optional) Regular expression of query match (e.g. the value .* will match on all query) (`String`).

### Proxy Protocol Choice Disable Proxy Protocol

Disable Proxy Protocol for upstream connections.

### Proxy Protocol Choice Proxy Protocol V1

Enable Proxy Protocol Version 1 for upstream connections.

### Proxy Protocol Choice Proxy Protocol V2

Enable Proxy Protocol Version 2 for upstream connections.

### Query Params Remove All Params

x-displayName: "Remove All Parameters".

### Query Params Retain All Params

x-displayName: "Retain All Parameters".

### Query Params Strip Query Params

Specifies the list of query params to be removed. Not supported.

`query_params` - (Optional) Query params keys to strip while manipulating the HTTP request (`String`).

### Rate Limit Rate Limiter

Requests to the virtual_host are rate limited based on the parameters specified in the rate_limiter..

`burst_multiplier` - (Optional) The maximum burst of requests to accommodate, expressed as a multiple of the rate. (`Int`).

`total_number` - (Required) The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period. (`Int`).

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

`ref_user_id` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier to be rate limited.. See [ref](#ref) below for details.

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

`port_redirect` - (Optional) Specify the port value to redirect to a URL with non default port(443) (`Int`).(Deprecated)

`proto_redirect` - (Optional) When incoming-proto option is specified, swapping of protocol is not done. (`String`).

###### One of the arguments from this list "all_params, remove_all_params, replace_params, retain_all_params, strip_query_params" can be set

`all_params` - (Optional) be removed. Default value is false, which means query portion of the URL will NOT be removed (`Bool`).(Deprecated)

`remove_all_params` - (Optional) x-displayName: "Remove All Parameters" (`Bool`).

`replace_params` - (Optional) x-displayName: "Replace All Parameters" (`String`).

`retain_all_params` - (Optional) x-displayName: "Retain All Parameters" (`Bool`).

`strip_query_params` - (Optional) Specifies the list of query params to be removed. Not supported. See [Query Params Strip Query Params ](#query-params-strip-query-params) below for details.(Deprecated)

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

`retry_on` - (Optional) matching one defined in retriable_status_codes field (`String`).(Deprecated)

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

### Secret Value Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Section Choice All Request Sections

x-displayName: "All Request".

### Section Choice All Response Sections

x-displayName: "All Response".

### Section Choice All Sections

x-displayName: "All Request & Response".

### Section Choice Custom Sections

x-displayName: "Custom Sections".

`custom_sections` - (Required) Request & Response Sections. (`List of Strings`).

### Secure Add Secure

Add secure attribute.

### Secure Ignore Secure

Ignore secure attribute.

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

### Sensitive Data Detection Rules Custom Sensitive Data Detection Rules

Rules to detect custom sensitive data in requests and/or responses sections..

`metadata` - (Required) Common attributes for the rule including name and description.. See [Custom Sensitive Data Detection Rules Metadata ](#custom-sensitive-data-detection-rules-metadata) below for details.

`sensitive_data_detection_config` - (Required) The custom data detection config specifies targets, scopes & the pattern to be detected.. See [Custom Sensitive Data Detection Rules Sensitive Data Detection Config ](#custom-sensitive-data-detection-rules-sensitive-data-detection-config) below for details.

`sensitive_data_type` - (Required) If the pattern is detected, the request is labeled with specified sensitive data type.. See [Custom Sensitive Data Detection Rules Sensitive Data Type ](#custom-sensitive-data-detection-rules-sensitive-data-type) below for details.

### Sensitive Data Detection Rules Disabled Built In Rules

List of disabled built-in sensitive data detection rules..

`name` - (Required) Built-in rule for sensitive data detection. (`String`).

### Sensitive Data Disclosure Rules Sensitive Data Types In Response

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses .

`body` - (Optional) x-displayName: "JSON Path". See [Sensitive Data Types In Response Body ](#sensitive-data-types-in-response-body) below for details.

###### One of the arguments from this list "mask, report" can be set

`mask` - (Optional) x-displayName: "Mask Sensitive Data" (`Bool`).(Deprecated)

`report` - (Optional) x-displayName: "Report Sensitive Data" (`Bool`).(Deprecated)

###### One of the arguments from this list "api_endpoint, api_group, base_path" must be set

`api_endpoint` - (Optional) The API endpoint (Path + Method) which this validation applies to. See [Type Condition Type Choice Api Endpoint ](#type-condition-type-choice-api-endpoint) below for details.

`api_group` - (Optional) The API group which this validation applies to (`String`).(Deprecated)

`base_path` - (Optional) The base path which this validation applies to (`String`).(Deprecated)

### Sensitive Data Policy Choice Sensitive Data Policy

Apply custom sensitive data discovery.

`sensitive_data_policy_ref` - (Required) Specify Sensitive Data Discovery. See [ref](#ref) below for details.

### Sensitive Data Types In Response Body

x-displayName: "JSON Path".

`fields` - (Required) List of JSON Path field values. Use square brackets with an underscore \[*] to indicate array elements, e.g., person.emails\[*]. (`String`).

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

### Server Validation Choice Skip Server Verification

Skip origin server verification.

### Server Validation Choice Use Server Verification

Perform origin server verification using the provided Root CA Certificate.

###### One of the arguments from this list "trusted_ca, trusted_ca_url" must be set

`trusted_ca` - (Optional) Select/Add a Root CA Certificate object to associate with this Origin Pool for verification of server's certificate. See [ref](#ref) below for details.

`trusted_ca_url` - (Optional) Upload a Root CA Certificate specifically for this Origin Pool for verification of server's certificate (`String`).

### Server Validation Choice Volterra Trusted Ca

Perform origin server verification using F5XC Default Root CA Certificate.

### Service Info Service Selector

discovery has to happen. This implicit label is added to service_selector.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Service Policy Choice Active Service Policies

Apply the specified list of service policies and bypass the namespace service policy set.

`policies` - (Required) If all policies are evaluated and none match, then the request will be denied by default.. See [ref](#ref) below for details.

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

`request_headers_to_add` - (Optional) Headers are key-value pairs to be added to HTTP request being routed towards upstream.. See [Advanced Options Request Headers To Add ](#advanced-options-request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

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

### Spdy Choice Disable Spdy

SPDY upgrade is disabled.

### Spdy Choice Enable Spdy

SPDY upgrade is enabled.

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

### Specific Hash Policy Hash Policy

route the request.

###### One of the arguments from this list "cookie, header_name, source_ip" must be set

`cookie` - (Optional) Hash based on cookie. See [Policy Specifier Cookie ](#policy-specifier-cookie) below for details.

`header_name` - (Optional) The name or key of the request header that will be used to obtain the hash key (`String`).

`source_ip` - (Optional) Hash based on source IP address (`Bool`).

`terminal` - (Optional) Specify if its a terminal policy (`Bool`).

### Strict Sni Host Header Check Choice Additional Domains

Wildcard names are supported in the suffix or prefix form.

`domains` - (Required) Wildcard names are supported in the suffix or prefix form. (`String`).

### Strict Sni Host Header Check Choice Enable Strict Sni Host Header Check

Enable strict SNI and Host header check.

### Subset Choice Disable Subsets

Subset load balancing is disabled. All eligible origin servers will be considered for load balancing..

### Subset Choice Enable Subsets

Subset load balancing is enabled. Based on route, subset of origin servers will be considered for load balancing..

`endpoint_subsets` - (Required) List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class.. See [Enable Subsets Endpoint Subsets ](#enable-subsets-endpoint-subsets) below for details.

###### One of the arguments from this list "any_endpoint, default_subset, fail_request" must be set

`any_endpoint` - (Optional) Select any origin server from available healthy origin servers in this pool (`Bool`).

`default_subset` - (Optional) Use the default subset provided here. Select endpoints matching default subset.. See [Fallback Policy Choice Default Subset ](#fallback-policy-choice-default-subset) below for details.

`fail_request` - (Optional) Request will be failed and error returned, as if cluster has no origin servers. (`Bool`).

### Target All Endpoint

Validation will be performed for all requests on this LB.

### Target Api Groups

Validation will be performed for the endpoints mentioned in the API Groups.

`api_groups` - (Required) x-required (`String`).

### Target Base Paths

Validation will be performed for selected path prefixes.

`base_paths` - (Required) x-required (`String`).

### Target Choice Any Target

The rule will be applied for all requests on this LB..

### Target Choice Api Endpoint Target

The rule is applied only for the specified api endpoints..

`api_endpoint_path` - (Required) The rule is applied only for the specified api endpoints. (`String`).

`methods` - (Required) x-example: "['GET', 'POST', 'DELETE']" (`List of Strings`).

### Temporary Blocking Parameters Choice Default Temporary Blocking Parameters

Use default parameters.

### Temporary Blocking Parameters Choice Temporary User Blocking

Specifies configuration for temporary user blocking resulting from malicious user detection.

`custom_page` - (Optional) E.g. "<p> Blocked </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

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

### Tls Certificates Choice Tls Parameters

Upload a TLS certificate covering all domain names for this Load Balancer.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Parameters Tls Certificates ](#tls-parameters-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Parameters Tls Config ](#tls-parameters-tls-config) below for details.

### Tls Choice No Tls

x-displayName: "Disable".

### Tls Choice Use Tls

x-displayName: "Enable".

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

### Trust Client Ip Headers Choice Enable Trust Client Ip Headers

x-displayName: "Enable".

`client_ip_headers` - (Required) For X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client ip (`String`).

### Trusted Clients Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Type Condition Type Choice Api Endpoint

The API endpoint (Path + Method) which this validation applies to.

`methods` - (Optional) Methods to be matched (`List of Strings`).

`path` - (Required) Path to be matched (`String`).

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

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

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
