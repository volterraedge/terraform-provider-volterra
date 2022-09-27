---

page_title: "Volterra: http_loadbalancer"

description: "The http_loadbalancer allows CRUD of Http Loadbalancer resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_http_loadbalancer
===================================

The Http Loadbalancer allows CRUD of Http Loadbalancer resource on Volterra SaaS

~> **Note:** Please refer to [Http Loadbalancer API docs](https://volterra.io/docs/api/views-http-loadbalancer) to learn more

Example Usage
-------------

```hcl
resource "volterra_http_loadbalancer" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "do_not_advertise advertise_on_public_default_vip advertise_on_public advertise_custom" must be set
  do_not_advertise = true

  // One of the arguments from this list "disable_api_definition api_definition api_definitions" must be set
  disable_api_definition = true

  // One of the arguments from this list "enable_api_discovery disable_api_discovery" must be set

  enable_api_discovery {
    // One of the arguments from this list "disable_learn_from_redirect_traffic enable_learn_from_redirect_traffic" must be set
    disable_learn_from_redirect_traffic = true
  }

  // One of the arguments from this list "no_challenge js_challenge captcha_challenge policy_based_challenge" must be set

  policy_based_challenge {
    // One of the arguments from this list "default_captcha_challenge_parameters captcha_challenge_parameters" must be set
    default_captcha_challenge_parameters = true

    // One of the arguments from this list "no_challenge always_enable_js_challenge always_enable_captcha_challenge" must be set
    no_challenge = true

    // One of the arguments from this list "js_challenge_parameters default_js_challenge_parameters" must be set

    js_challenge_parameters {
      cookie_expiry   = "1000"
      custom_page     = "string:///PHA+IFBsZWFzZSBXYWl0IDwvcD4="
      js_script_delay = "1000"
    }
    // One of the arguments from this list "default_mitigation_settings malicious_user_mitigation" must be set
    default_mitigation_settings = true
    rule_list {
      rules {
        metadata {
          description = "Virtual Host for acmecorp website"
          disable     = true
          name        = "acmecorp-web"
        }

        spec {
          arg_matchers {
            invert_matcher = true

            // One of the arguments from this list "check_present check_not_present item presence" must be set
            check_not_present = true
            name              = "name"
          }

          // One of the arguments from this list "any_asn asn_list asn_matcher" must be set
          any_asn = true

          body_matcher {
            exact_values = ["['new york', 'london', 'sydney', 'tokyo', 'cairo']"]

            regex_values = ["['^new .*$', 'san f.*', '.* del .*']"]

            transformers = ["transformers"]
          }

          // One of the arguments from this list "disable_challenge enable_javascript_challenge enable_captcha_challenge" must be set
          disable_challenge = true

          // One of the arguments from this list "client_selector client_name_matcher any_client client_name" must be set
          any_client = true

          cookie_matchers {
            invert_matcher = true

            // One of the arguments from this list "presence check_present check_not_present item" must be set
            presence = true

            name = "Session"
          }

          domain_matcher {
            exact_values = ["['new york', 'london', 'sydney', 'tokyo', 'cairo']"]

            regex_values = ["['^new .*$', 'san f.*', '.* del .*']"]
          }

          expiration_timestamp = "0001-01-01T00:00:00Z"

          headers {
            invert_matcher = true

            // One of the arguments from this list "presence check_present check_not_present item" must be set
            presence = true

            name = "Accept-Encoding"
          }

          http_method {
            invert_matcher = true

            methods = ["['GET', 'POST', 'DELETE']"]
          }

          // One of the arguments from this list "any_ip ip_prefix_list ip_matcher" must be set
          any_ip = true

          path {
            exact_values = ["['/api/web/namespaces/project179/users/user1', '/api/config/namespaces/accounting/bgps', '/api/data/namespaces/project443/virtual_host_101']"]

            prefix_values = ["['/api/web/namespaces/project179/users/', '/api/config/namespaces/', '/api/data/namespaces/']"]

            regex_values = ["['^/api/web/namespaces/abc/users/([a-z]([-a-z0-9]*[a-z0-9])?)$', '/api/data/namespaces/proj404/virtual_hosts/([a-z]([-a-z0-9]*[a-z0-9])?)$']"]

            suffix_values = ["['.exe', '.shtml', '.wmz']"]

            transformers = ["transformers"]
          }

          query_params {
            invert_matcher = true
            key            = "sourceid"

            // One of the arguments from this list "check_not_present item presence check_present" must be set
            presence = true
          }

          tls_fingerprint_matcher {
            classes = ["classes"]

            exact_values = ["['ed6dfd54b01ebe31b7a65b88abfa7297', '16efcf0e00504ddfedde13bfea997952', 'de364c46b0dfc283b5e38c79ceae3f8f']"]

            excluded_values = ["['fb00055a1196aeea8d1bc609885ba953', 'b386946a5a44d1ddcc843bc75336dfce']"]
          }
        }
      }
    }
    // One of the arguments from this list "default_temporary_blocking_parameters temporary_user_blocking" must be set
    default_temporary_blocking_parameters = true
  }
  // One of the arguments from this list "enable_ddos_detection disable_ddos_detection" must be set
  enable_ddos_detection = true
  domains = ["www.foo.com"]
  // One of the arguments from this list "least_active random source_ip_stickiness cookie_stickiness ring_hash round_robin" must be set
  random = true

  // One of the arguments from this list "http https_auto_cert https" must be set

  http {
    dns_volterra_managed = true
    port                 = "80"
  }
  // One of the arguments from this list "enable_malicious_user_detection disable_malicious_user_detection" must be set
  enable_malicious_user_detection = true
  // One of the arguments from this list "rate_limit disable_rate_limit api_rate_limit" must be set
  disable_rate_limit = true

  // One of the arguments from this list "service_policies_from_namespace no_service_policies active_service_policies" must be set

  active_service_policies {
    policies {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }
  // One of the arguments from this list "disable_trust_client_ip_headers enable_trust_client_ip_headers" must be set
  disable_trust_client_ip_headers = true
  // One of the arguments from this list "user_id_client_ip user_identification" must be set
  user_id_client_ip = true

  // One of the arguments from this list "app_firewall disable_waf waf waf_rule" must be set

  waf {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
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

`add_location` - (Optional) is ignored on CE sites. (`Bool`).

`advertise_custom` - (Optional) Advertise this load balancer on specific sites. See [Advertise Custom ](#advertise-custom) below for details.

`advertise_on_public` - (Optional) Advertise this load balancer on public network. See [Advertise On Public ](#advertise-on-public) below for details.

`advertise_on_public_default_vip` - (Optional) Advertise this load balancer on public network with default VIP (bool).

`do_not_advertise` - (Optional) Do not advertise this load balancer (bool).

`api_definition` - (Optional) Specify API definition which includes application API paths and methods derived from swagger files.. See [ref](#ref) below for details.

`api_definitions` - (Optional) DEPRECATED by 'api_definition'. See [Api Definitions ](#api-definitions) below for details.

`disable_api_definition` - (Optional) API Definition is not currently used for this load balancer (bool).

`disable_api_discovery` - (Optional) x-displayName: "Disable" (bool).

`enable_api_discovery` - (Optional) x-displayName: "Enable". See [Enable Api Discovery ](#enable-api-discovery) below for details.

`api_protection_rules` - (Optional) Rules can also include additional conditions, for example specific clients can access certain API endpoint or API group.. See [Api Protection Rules ](#api-protection-rules) below for details.

`blocked_clients` - (Optional) Define rules to block IP Prefixes or AS numbers.. See [Blocked Clients ](#blocked-clients) below for details.

`bot_defense` - (Optional) Bot Defense configuration for Protected App endpoints and JavaScript insertion. See [Bot Defense ](#bot-defense) below for details.

`disable_bot_defense` - (Optional) No Bot Defense configuration for this load balancer (bool).

`captcha_challenge` - (Optional) Configure Captcha challenge on this load balancer. See [Captcha Challenge ](#captcha-challenge) below for details.

`js_challenge` - (Optional) Configure Javascript challenge on this load balancer. See [Js Challenge ](#js-challenge) below for details.

`no_challenge` - (Optional) No challenge is enabled for this load balancer (bool).

`policy_based_challenge` - (Optional) Specifies the settings for policy rule based challenge. See [Policy Based Challenge ](#policy-based-challenge) below for details.

`client_side_defense` - (Optional) Client-Side Defense configuration for JavaScript insertion. See [Client Side Defense ](#client-side-defense) below for details.

`disable_client_side_defense` - (Optional) No Client-Side Defense configuration for this load balancer (bool).

`cors_policy` - (Optional) resources from a server at a different origin. See [Cors Policy ](#cors-policy) below for details.

`data_guard_rules` - (Optional) Note: App Firewall should be enabled, to use Data Guard feature.. See [Data Guard Rules ](#data-guard-rules) below for details.

`disable_ddos_detection` - (Optional) x-displayName: "Disable" (bool).

`enable_ddos_detection` - (Optional) x-displayName: "Enable" (bool).

`ddos_mitigation_rules` - (Optional) Rules that specify the DDoS clients to be blocked. See [Ddos Mitigation Rules ](#ddos-mitigation-rules) below for details.

`default_route_pools` - (Optional) Origin Pools used when no route is specified (default route). See [Default Route Pools ](#default-route-pools) below for details.

`domains` - (Required) Domains also indicate the list of names for which DNS resolution will be done by VER (`List of String`).

`cookie_stickiness` - (Optional) Request are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server. See [Cookie Stickiness ](#cookie-stickiness) below for details.

`least_active` - (Optional) Request are sent to origin server that has least active requests (bool).

`random` - (Optional) Request are sent to all eligible origin servers in random fashion (bool).

`ring_hash` - (Optional) Request are sent to all eligible origin servers using hash of request based on hash policy. Consistent hashing algorithm, ring hash, is used to select origin server. See [Ring Hash ](#ring-hash) below for details.

`round_robin` - (Optional) Request are sent to all eligible origin servers in round robin fashion (bool).

`source_ip_stickiness` - (Optional) Request are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server (bool).

`disable_ip_reputation` - (Optional) x-displayName: "Disable" (bool).

`enable_ip_reputation` - (Optional) x-displayName: "Enable". See [Enable Ip Reputation ](#enable-ip-reputation) below for details.

`http` - (Optional) HTTP Load Balancer.. See [Http ](#http) below for details.

`https` - (Optional) User is responsible for managing DNS to this load balancer.. See [Https ](#https) below for details.

`https_auto_cert` - (Optional) or a DNS CNAME record should be created in your DNS provider's portal.. See [Https Auto Cert ](#https-auto-cert) below for details.

`disable_malicious_user_detection` - (Optional) x-displayName: "Disable" (bool).

`enable_malicious_user_detection` - (Optional) x-displayName: "Enable" (bool).

`malicious_user_mitigation` - (Optional) The settings defined in malicious user mitigation specify what mitigation actions to take for users determined to be at different threat levels.. See [ref](#ref) below for details.

`multi_lb_app` - (Optional) It should be configured externally using app type feature and label should be added to the HTTP load balancer. (bool).

`single_lb_app` - (Optional) ML Config applied on this load balancer. See [Single Lb App ](#single-lb-app) below for details.

`more_option` - (Optional) More options like header manipulation, compression etc.. See [More Option ](#more-option) below for details.

`api_rate_limit` - (Optional) Define rate limiting for one or more API endpoints. See [Api Rate Limit ](#api-rate-limit) below for details.

`disable_rate_limit` - (Optional) Rate limiting is not currently enabled for this load balancer (bool).

`rate_limit` - (Optional) Define custom rate limiting parameters for this load balancer. See [Rate Limit ](#rate-limit) below for details.

`routes` - (Optional) Routes for this load balancer. See [Routes ](#routes) below for details.

`active_service_policies` - (Optional) Apply the specified list of service policies and bypass the namespace service policy set. See [Active Service Policies ](#active-service-policies) below for details.

`no_service_policies` - (Optional) Do not apply any service policies i.e. bypass the namespace service policy set (bool).

`service_policies_from_namespace` - (Optional) Apply the active service policies configured as part of the namespace service policy set (bool).

`disable_trust_client_ip_headers` - (Optional) x-displayName: "Disable" (bool).

`enable_trust_client_ip_headers` - (Optional) x-displayName: "Enable". See [Enable Trust Client Ip Headers ](#enable-trust-client-ip-headers) below for details.

`trusted_clients` - (Optional) Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients.. See [Trusted Clients ](#trusted-clients) below for details.

`user_id_client_ip` - (Optional) Use the Client IP address as the user identifier. (bool).

`user_identification` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier.. See [ref](#ref) below for details.

`app_firewall` - (Optional) Reference to App Firewall configuration object. See [ref](#ref) below for details.

`disable_waf` - (Optional) No WAF configuration for this load balancer (bool).

`waf` - (Optional) Reference to WAF intent configuration object. See [ref](#ref) below for details.

`waf_rule` - (Optional) Reference to WAF Rules configuration object. See [ref](#ref) below for details.

`waf_exclusion_rules` - (Optional) The match criteria include domain, path and method.. See [Waf Exclusion Rules ](#waf-exclusion-rules) below for details.

### Action

The action to take if the input request matches the rule..

`allow` - (Optional) Allow the request to proceed. (bool).

`deny` - (Optional) Deny the request. (bool).

### Active Service Policies

Apply the specified list of service policies and bypass the namespace service policy set.

`policies` - (Required) If all policies are evaluated and none match, then the request will be denied by default.. See [ref](#ref) below for details.

### Additional Domains

Wildcard names are supported in the suffix or prefix form.

`domains` - (Optional) Wildcard names are supported in the suffix or prefix form. (`String`).

### Advanced Options

Configure Advanced per route options.

`buffer_policy` - (Optional) Route level buffer configuration overrides any configuration at VirtualHost level.. See [Buffer Policy ](#buffer-policy) below for details.

`common_buffering` - (Optional) Use common buffering configuration (bool).

`do_not_retract_cluster` - (Optional) configuration. (bool).

`retract_cluster` - (Optional) for route (bool).

`cors_policy` - (Optional) resources from a server at a different origin. See [Cors Policy ](#cors-policy) below for details.

`disable_location_add` - (Optional) virtual-host level. This configuration is ignored on CE sites. (`Bool`).

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

`common_hash_policy` - (Optional) Use load balancer hash policy for this route (bool).

`specific_hash_policy` - (Optional) Configure hash policy specific for this route. See [Specific Hash Policy ](#specific-hash-policy) below for details.

`disable_mirroring` - (Optional) Disable Mirroring of request (bool).

`mirror_policy` - (Optional) useful for logging. For example, *cluster1* becomes *cluster1-shadow*.. See [Mirror Policy ](#mirror-policy) below for details.

`priority` - (Optional) Also, circuit-breaker configuration at destination cluster is chosen based on the route priority. (`String`).

`request_headers_to_add` - (Optional) Headers are key-value pairs to be added to HTTP request being routed towards upstream.. See [Request Headers To Add ](#request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) Headers are key-value pairs to be added to HTTP response being sent towards downstream.. See [Response Headers To Add ](#response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

`default_retry_policy` - (Optional) Use system default retry policy (bool).

`retry_policy` - (Optional) Configure custom retry policy. See [Retry Policy ](#retry-policy) below for details.

`disable_prefix_rewrite` - (Optional) Do not rewrite the path prefix. (bool).

`prefix_rewrite` - (Optional) the query string) will be swapped with this value. (`String`).

`disable_spdy` - (Optional) SPDY upgrade is disabled (bool).

`enable_spdy` - (Optional) SPDY upgrade is enabled (bool).

`timeout` - (Optional) Should be set to a high value or 0 (infinite timeout) for server-side streaming. (`Int`).

`app_firewall` - (Optional) Reference to App Firewall configuration object. See [ref](#ref) below for details.

`inherited_waf` - (Optional) Hence no custom configuration is applied on the route (bool).

`waf` - (Optional) Reference to WAF intent configuration object. See [ref](#ref) below for details.

`waf_rule` - (Optional) Reference to WAF Rules configuration object. See [ref](#ref) below for details.

`disable_web_socket_config` - (Optional) Websocket upgrade is disabled (bool).

`web_socket_config` - (Optional) Upgrade to Websocket for this route. See [Web Socket Config ](#web-socket-config) below for details.

### Advertise Custom

Advertise this load balancer on specific sites.

`advertise_where` - (Required) Where should this load balancer be available. See [Advertise Where ](#advertise-where) below for details.

### Advertise On Public

Advertise this load balancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Advertise Where

Where should this load balancer be available.

`site` - (Optional) Advertise on a customer site and a given network.. See [Site ](#site) below for details.

`virtual_network` - (Optional) Advertise on a virtual network. See [Virtual Network ](#virtual-network) below for details.

`virtual_site` - (Optional) Advertise on a customer virtual site and a given network.. See [Virtual Site ](#virtual-site) below for details.

`vk8s_service` - (Optional) Advertise on vK8s Service Network on RE.. See [Vk8s Service ](#vk8s-service) below for details.

`port` - (Optional) TCP port to Listen. (`Int`).

`use_default_port` - (Optional) For HTTP, default is 80. For HTTPS/SNI, default is 443. (bool).

### Allow

Allow the request to proceed..

### Always Enable Captcha Challenge

Challenge rules can be used to selectively disable Captcha challenge or enable Javascript challenge for some requests..

### Always Enable Js Challenge

Challenge rules can be used to selectively disable Javascript challenge or enable Captcha challenge for some requests..

### Any Client

Any Client.

### Any Domain

The rule will apply for all domains..

### Any Ip

Any Source IP.

### Api Definitions

DEPRECATED by 'api_definition'.

`api_definitions` - (Optional) API Definitions using OpenAPI specification files. See [ref](#ref) below for details.

### Api Endpoint Method

The predicate evaluates to true if the actual HTTP method belongs is present in the list of expected values..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`methods` - (Optional) x-example: "['GET', 'POST', 'DELETE']" (`List of Strings`).

### Api Endpoint Rules

If request matches any of these rules, skipping second category rules..

`action` - (Required) The action to take if the input request matches the rule.. See [Action ](#action) below for details.

`api_endpoint_method` - (Optional) The predicate evaluates to true if the actual HTTP method belongs is present in the list of expected values.. See [Api Endpoint Method ](#api-endpoint-method) below for details.

`api_endpoint_path` - (Required) The endpoint (path) of the request. (`String`).

`client_matcher` - (Optional) Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc.. See [Client Matcher ](#client-matcher) below for details.

`any_domain` - (Required) The rule will apply for all domains. (bool).

`specific_domain` - (Required) For example: api.example.com (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`request_matcher` - (Optional) Conditions related to the request, such as query parameters, headers, etc.. See [Request Matcher ](#request-matcher) below for details.

### Api Groups Rules

For API groups, refer to API Definition which includes API groups derived from uploaded swaggers..

`action` - (Required) The action to take if the input request matches the rule.. See [Action ](#action) below for details.

`api_group` - (Optional) Custom groups can be created if user tags paths or operations with "x-volterra-api-group" extensions inside swaggers. (`String`).

`base_path` - (Required) For example: /v1 (`String`).

`client_matcher` - (Optional) Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc.. See [Client Matcher ](#client-matcher) below for details.

`any_domain` - (Required) The rule will apply for all domains. (bool).

`specific_domain` - (Required) For example: api.example.com (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`request_matcher` - (Optional) Conditions related to the request, such as query parameters, headers, etc.. See [Request Matcher ](#request-matcher) below for details.

### Api Protection Rules

Rules can also include additional conditions, for example specific clients can access certain API endpoint or API group..

`api_endpoint_rules` - (Optional) If request matches any of these rules, skipping second category rules.. See [Api Endpoint Rules ](#api-endpoint-rules) below for details.

`api_groups_rules` - (Optional) For API groups, refer to API Definition which includes API groups derived from uploaded swaggers.. See [Api Groups Rules ](#api-groups-rules) below for details.

### Api Rate Limit

Define rate limiting for one or more API endpoints.

`api_endpoint_rules` - (Optional) For creating rule that contain a whole domain or group of endpoints, please use the server URL rules above.. See [Api Endpoint Rules ](#api-endpoint-rules) below for details.

`custom_ip_allowed_list` - (Optional) IP Allowed list using existing ip_prefix_set objects.. See [Custom Ip Allowed List ](#custom-ip-allowed-list) below for details.

`ip_allowed_list` - (Optional) List of IP(s) for which rate limiting will be disabled.. See [Ip Allowed List ](#ip-allowed-list) below for details.

`no_ip_allowed_list` - (Optional) There is no ip allowed list for rate limiting, all clients go through rate limiting. (bool).

`server_url_rules` - (Optional) For matching also specific endpoints you can use the API endpoint rules set bellow.. See [Server Url Rules ](#server-url-rules) below for details.

### App Firewall Detection Control

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria..

`exclude_attack_type_contexts` - (Optional) App Firewall attack types contexts to be excluded for this request. See [Exclude Attack Type Contexts ](#exclude-attack-type-contexts) below for details.

`exclude_bot_name_contexts` - (Optional) Bot names contexts to be excluded for this request. See [Exclude Bot Name Contexts ](#exclude-bot-name-contexts) below for details.

`exclude_signature_contexts` - (Optional) App Firewall signature contexts to be excluded for this request. See [Exclude Signature Contexts ](#exclude-signature-contexts) below for details.

`exclude_violation_contexts` - (Optional) App Firewall violation contexts to be excluded for this request. See [Exclude Violation Contexts ](#exclude-violation-contexts) below for details.

### Append Headers

Append mitigation headers..

`auto_type_header_name` - (Required) A case-insensitive HTTP header name. (`String`).

`inference_header_name` - (Required) A case-insensitive HTTP header name. (`String`).

### Apply Data Guard

x-displayName: "Apply".

### Asn List

The predicate evaluates to true if the origin ASN is present in the ASN list..

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. (`Int`).

### Asn Matcher

The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects..

`asn_sets` - (Required) A list of references to bgp_asn_set objects.. See [ref](#ref) below for details.

### Auto Host Rewrite

Host header will be swapped with hostname of upstream host chosen by the cluster.

### Back Off

10 times the base interval.

`base_interval` - (Optional) Specifies the base interval between retries in milliseconds (`Int`).

`max_interval` - (Optional) to the base_interval if set. The default is 10 times the base_interval. (`Int`).

### Blindfold Secret Info

Blindfold Secret is used for the secrets managed by Volterra Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Block

Block bot request and send response with custom content..

`body` - (Optional) E.g. "<p> Your request was blocked </p>". Base64 encoded string for this html is "LzxwPiBZb3VyIHJlcXVlc3Qgd2FzIGJsb2NrZWQgPC9wPg==" (`String`).

`body_hash` - (Optional) Represents the corresponding MD5 Hash for the body message. (`String`).

`status` - (Optional) HTTP Status code to respond with (`String`).

### Blocked Clients

Define rules to block IP Prefixes or AS numbers..

`bot_skip_processing` - (Optional) Skip Bot Defense processing for clients matching this rule. (bool).

`skip_processing` - (Optional) Skip both WAF and Bot Defense processing for clients matching this rule. (bool).

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (bool).

`actions` - (Optional) Action that should be taken when client identifier matches the rule (`List of Strings`).

`as_number` - (Required) RFC 6793 defined 4-byte AS number (`Int`).

`http_header` - (Required) Request header name and value pairs. See [Http Header ](#http-header) below for details.

`ip_prefix` - (Required) IPv4 prefix string. (`String`).

`user_identifier` - (Optional) Identify user based on user identifier. User identifier value needs to be copied from security event. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

### Bot Defense

Bot Defense configuration for Protected App endpoints and JavaScript insertion.

`policy` - (Required) Bot Defense Policy.. See [Policy ](#policy) below for details.

`regional_endpoint` - (Required) x-required (`String`).

`timeout` - (Optional) The timeout for the inference check, in milliseconds. (`Int`).

### Bot Skip Processing

Skip Bot Defense processing for clients matching this rule..

### Buffer Policy

specify the maximum buffer size and buffer interval with this config..

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`max_request_bytes` - (Optional) manager will stop buffering and return a RequestEntityTooLarge (413) response. (`Int`).

`max_request_time` - (Optional) request before returning a RequestTimeout (408) response (`Int`).

### Captcha Challenge

Configure Captcha challenge on this load balancer.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

### Captcha Challenge Parameters

Configure captcha challenge parameters.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

### Check Not Present

Check that the cookie is not present..

### Check Present

Check that the cookie is present..

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Client Matcher

Conditions related to the origin of the request, such as client IP, TLS fingerprint, etc..

`any_client` - (Optional) Any Client (bool).

`client_selector` - (Optional) The predicate evaluates to true if the expressions in the label selector are true for the client labels.. See [Client Selector ](#client-selector) below for details.

`ip_threat_category_list` - (Optional) IP threat categories to choose from. See [Ip Threat Category List ](#ip-threat-category-list) below for details.

`any_ip` - (Optional) Any Source IP (bool).

`asn_list` - (Optional) The predicate evaluates to true if the origin ASN is present in the ASN list.. See [Asn List ](#asn-list) below for details.

`asn_matcher` - (Optional) The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects.. See [Asn Matcher ](#asn-matcher) below for details.

`ip_matcher` - (Optional) The predicate evaluates to true if the client IPv4 Address is covered by one or more of the IPv4 Prefixes in the IP Prefix Sets.. See [Ip Matcher ](#ip-matcher) below for details.

`ip_prefix_list` - (Optional) The predicate evaluates to true if the client IPv4 Address is covered by one or more of the IPv4 Prefixes from the list.. See [Ip Prefix List ](#ip-prefix-list) below for details.

`tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints.. See [Tls Fingerprint Matcher ](#tls-fingerprint-matcher) below for details.

### Client Selector

The predicate evaluates to true if the expressions in the label selector are true for the client labels..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Client Side Defense

Client-Side Defense configuration for JavaScript insertion.

`policy` - (Required) Client-Side Defense Policy.. See [Policy ](#policy) below for details.

### Common Buffering

Use common buffering configuration.

### Common Hash Policy

Use load balancer hash policy for this route.

### Compression Params

Only GZIP compression is supported.

`content_length` - (Optional) Minimum response length, in bytes, which will trigger compression. The default value is 30. (`Int`).

`content_type` - (Optional) "text/xml" (`String`).

`disable_on_etag_header` - (Optional) weak etags will be preserved and the ones that require strong validation will be removed. (`Bool`).

`remove_accept_encoding_header` - (Optional) so that responses do not get compressed before reaching the filter. (`Bool`).

### Cookie

Hash based on cookie.

`name` - (Required) produced (`String`).

`path` - (Optional) will be set for the cookie (`String`).

`ttl` - (Optional) be a session cookie. TTL value is in milliseconds (`Int`).

### Cookie Matchers

Note that all specified cookie matcher predicates must evaluate to true..

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

`check_not_present` - (Optional) Check that the cookie is not present. (bool).

`check_present` - (Optional) Check that the cookie is present. (bool).

`item` - (Optional) Criteria for matching the values for the cookie. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the cookie is present or absent. (`Bool`).

`name` - (Required) A case-sensitive cookie name. (`String`).

### Cookie Stickiness

Request are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server.

`name` - (Required) produced (`String`).

`path` - (Optional) will be set for the cookie (`String`).

`ttl` - (Optional) be a session cookie. TTL value is in milliseconds (`Int`).

### Cors Policy

resources from a server at a different origin.

`allow_credentials` - (Optional) Specifies whether the resource allows credentials (`Bool`).

`allow_headers` - (Optional) Specifies the content for the access-control-allow-headers header (`String`).

`allow_methods` - (Optional) Specifies the content for the access-control-allow-methods header (`String`).

`allow_origin` - (Optional) An origin is allowed if either allow_origin or allow_origin_regex match (`String`).

`allow_origin_regex` - (Optional) An origin is allowed if either allow_origin or allow_origin_regex match (`String`).

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`expose_headers` - (Optional) Specifies the content for the access-control-expose-headers header (`String`).

`max_age` - (Optional) Specifies the content for the access-control-max-age header (`String`).

`maximum_age` - (Optional) Maximum permitted value is 86400 seconds (24 hours) (`Int`).

### Custom Hash Algorithms

Use hash algorithms in the custom order. Volterra will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Custom Ip Allowed List

IP Allowed list using existing ip_prefix_set objects..

`rate_limiter_allowed_prefixes` - (Required) Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting.. See [ref](#ref) below for details.

### Custom Route Object

A custom route uses a route object created outside of this view..

`route_ref` - (Optional) Reference to a custom route object. See [ref](#ref) below for details.

### Custom Security

Custom selection of TLS versions and cipher suites.

`cipher_suites` - (Required) The TLS listener will only support the specified cipher list. (`String`).

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

### Data Guard Rules

Note: App Firewall should be enabled, to use Data Guard feature..

`apply_data_guard` - (Optional) x-displayName: "Apply" (bool).

`skip_data_guard` - (Optional) x-displayName: "Skip" (bool).

`any_domain` - (Optional) Enable Data Guard for any domain (bool).

`exact_value` - (Optional) Exact domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`path` - (Required) URI path matcher.. See [Path ](#path) below for details.

### Ddos Client Source

Combination of Region, ASN and TLS Fingerprints.

`asn_list` - (Optional) The ASN is obtained by performing a lookup for the source IPv4 Address in a GeoIP DB.. See [Asn List ](#asn-list) below for details.

`country_list` - (Optional) Sources that are located in one of the countries in the given list (`List of Strings`).

`tls_fingerprint_matcher` - (Optional) The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints.. See [Tls Fingerprint Matcher ](#tls-fingerprint-matcher) below for details.

### Ddos Mitigation Rules

Rules that specify the DDoS clients to be blocked.

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`block` - (Optional) Block user for a duration determined by the expiration time (bool).

`ddos_client_source` - (Required) Combination of Region, ASN and TLS Fingerprints. See [Ddos Client Source ](#ddos-client-source) below for details.

`ip_prefix_list` - (Required) IPv4 prefix string.. See [Ip Prefix List ](#ip-prefix-list) below for details.

### Default Captcha Challenge Parameters

Use default parameters.

### Default Header

Response header name is “server” and value is “volt-adc”.

### Default Header Transformation

Normalize the headers to lower case.

### Default Js Challenge Parameters

Use default parameters.

### Default Loadbalancer

x-displayName: "Yes".

### Default Mitigation Settings

For high level, users will be temporarily blocked..

### Default Retry Policy

Use system default retry policy.

### Default Route Pools

Origin Pools used when no route is specified (default route).

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

`cluster` - (Required) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Required) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`priority` - (Optional) made active as per the increasing priority. (`Int`).

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

### Default Security

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Default Temporary Blocking Parameters

Use default parameters.

### Default Vip

Use the default VIP, system allocated or configured in the virtual network.

### Deny

Deny the request..

### Direct Response Route

A direct response route matches on path and/or HTTP method and responds directly to the matching traffic.

`headers` - (Optional) List of (key, value) headers. See [Headers ](#headers) below for details.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`path` - (Optional) URI path of route. See [Path ](#path) below for details.

`route_direct_response` - (Optional) Send direct response. See [Route Direct Response ](#route-direct-response) below for details.

### Disable Ddos Detection

x-displayName: "Disable".

### Disable Discovery

x-displayName: "Disable".

### Disable Host Rewrite

Host header is not modified.

### Disable Js Insert

Disable JavaScript insertion..

### Disable Learn From Redirect Traffic

Disable learning API patterns from traffic with redirect response codes 3xx.

### Disable Malicious User Detection

x-displayName: "Disable".

### Disable Mirroring

Disable Mirroring of request.

### Disable Mobile Sdk

Disable Mobile SDK..

### Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Disable Path Normalize

x-displayName: "Disable".

### Disable Prefix Rewrite

Do not rewrite the path prefix..

### Disable Spdy

SPDY upgrade is disabled.

### Disable Web Socket Config

Websocket upgrade is disabled.

### Do Not Retract Cluster

configuration..

### Domain

Domain matcher..

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Enable Api Discovery

x-displayName: "Enable".

`disable_learn_from_redirect_traffic` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (bool).

`enable_learn_from_redirect_traffic` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (bool).

### Enable Ddos Detection

x-displayName: "Enable".

### Enable Discovery

x-displayName: "Enable".

`disable_learn_from_redirect_traffic` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (bool).

`enable_learn_from_redirect_traffic` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (bool).

### Enable Ip Reputation

x-displayName: "Enable".

`ip_threat_categories` - (Required) If the source IP matches on atleast one of the enabled IP threat categories, the request will be denied. (`List of Strings`).

### Enable Learn From Redirect Traffic

Enable learning API patterns from traffic with redirect response codes 3xx.

### Enable Malicious User Detection

x-displayName: "Enable".

### Enable Path Normalize

x-displayName: "Enable".

### Enable Spdy

SPDY upgrade is enabled.

### Enable Strict Sni Host Header Check

Enable strict SNI and Host header check.

### Enable Trust Client Ip Headers

x-displayName: "Enable".

`client_ip_headers` - (Optional) Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom, meaning if the first header is not present in the request, the system will proceed to check for the second header, and so on, until one of the listed headers is found. If none of the defined headers exist, or the value is not an IP address, then the system will use the source IP of the packet. If multiple defined headers with different names are present in the request, the value of the first header name in the configuration will be used. If multiple defined headers with the same name are present in the request, values of all those headers will be combined. The system will read the right-most IP address from header, if there are multiple ip addresses in the header value. (`String`).

### Exclude Attack Type Contexts

App Firewall attack types contexts to be excluded for this request.

`exclude_attack_type` - (Required) x-required (`String`).

### Exclude Bot Name Contexts

Bot names contexts to be excluded for this request.

`bot_name` - (Required) x-example: "Hydra" (`String`).

### Exclude List

Optional JavaScript insertions exclude list of domain and path matchers..

`any_domain` - (Optional) Any Domain. (bool).

`domain` - (Optional) Domain matcher.. See [Domain ](#domain) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`path` - (Required) URI path matcher.. See [Path ](#path) below for details.

### Exclude Signature Contexts

App Firewall signature contexts to be excluded for this request.

`signature_id` - (Required) x-required (`Int`).

### Exclude Violation Contexts

App Firewall violation contexts to be excluded for this request.

`exclude_violation` - (Required) x-required (`String`).

### Flag

Flag the request while not taking any invasive actions..

`append_headers` - (Optional) Append mitigation headers.. See [Append Headers ](#append-headers) below for details.

`no_headers` - (Optional) No mitigation headers. (bool).

### Hash Policy

route the request.

`cookie` - (Optional) Hash based on cookie. See [Cookie ](#cookie) below for details.

`header_name` - (Optional) The name or key of the request header that will be used to obtain the hash key (`String`).

`source_ip` - (Optional) Hash based on source IP address (`Bool`).

`terminal` - (Optional) Specify if its a terminal policy (`Bool`).

### Header

Header that is used by mobile traffic..

`check_not_present` - (Optional) Check that the header is not present. (bool).

`check_present` - (Optional) Check that the header is present. (bool).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Item ](#item) below for details.

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Header Transformation Type

Header transformation options for response headers to the client.

`default_header_transformation` - (Optional) Normalize the headers to lower case (bool).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (bool).

### Headers

Note that all specified header predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`check_not_present` - (Optional) Check that the header is not present. (bool).

`check_present` - (Optional) Check that the header is present. (bool).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the header is present or absent. (`Bool`).

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Http

HTTP Load Balancer..

`dns_volterra_managed` - (Optional) or a DNS CNAME record should be created in your DNS provider's portal. (`Bool`).

`port` - (Optional) x-example: "80" (`Int`).

### Http Header

Request header name and value pairs.

`headers` - (Required) List of HTTP header name and value pairs. See [Headers ](#headers) below for details.

### Https

User is responsible for managing DNS to this load balancer..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`default_loadbalancer` - (Optional) x-displayName: "Yes" (bool).

`non_default_loadbalancer` - (Optional) x-displayName: "No" (bool).

`header_transformation_type` - (Optional) Header transformation options for response headers to the client. See [Header Transformation Type ](#header-transformation-type) below for details.

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

`disable_path_normalize` - (Optional) x-displayName: "Disable" (bool).

`enable_path_normalize` - (Optional) x-displayName: "Enable" (bool).

`port` - (Optional) x-example: "443" (`Int`).

`append_server_name` - (Optional) If header value is already present, it is not overwritten and passed as-is. (`String`).

`default_header` - (Optional) Response header name is “server” and value is “volt-adc” (bool).

`pass_through` - (Optional) Pass existing server header as is. If server header is absent, a new header is not appended. (bool).

`server_name` - (Optional) This will overwrite existing values, if any, for the server header. (`String`).

`tls_parameters` - (Optional) TLS parameters for downstream connections.. See [Tls Parameters ](#tls-parameters) below for details.

### Https Auto Cert

or a DNS CNAME record should be created in your DNS provider's portal..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`default_loadbalancer` - (Optional) x-displayName: "Yes" (bool).

`non_default_loadbalancer` - (Optional) x-displayName: "No" (bool).

`header_transformation_type` - (Optional) Header transformation options for response headers to the client. See [Header Transformation Type ](#header-transformation-type) below for details.

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

`no_mtls` - (Optional) x-displayName: "Disable" (bool).

`use_mtls` - (Optional) x-displayName: "Enable". See [Use Mtls ](#use-mtls) below for details.

`disable_path_normalize` - (Optional) x-displayName: "Disable" (bool).

`enable_path_normalize` - (Optional) x-displayName: "Enable" (bool).

`port` - (Optional) x-example: "443" (`Int`).

`append_server_name` - (Optional) If header value is already present, it is not overwritten and passed as-is. (`String`).

`default_header` - (Optional) Response header name is “server” and value is “volt-adc” (bool).

`pass_through` - (Optional) Pass existing server header as is. If server header is absent, a new header is not appended. (bool).

`server_name` - (Optional) This will overwrite existing values, if any, for the server header. (`String`).

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Config ](#tls-config) below for details.

### Inherited Waf

Hence no custom configuration is applied on the route.

### Inline Rate Limiter

Specify rate values for the rule..

`ref_user_id` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier to be rate limited.. See [ref](#ref) below for details.

`use_http_lb_user_id` - (Optional) Defined in HTTP-LB Security Configuration -> User Identifier. (bool).

`threshold` - (Required) The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period. (`Int`).

`unit` - (Required) Unit for the period per which the rate limit is applied. (`String`).

### Ip Allowed List

List of IP(s) for which rate limiting will be disabled..

`prefixes` - (Required) List of IPv4 prefixes that represent an endpoint (`String`).

### Ip Matcher

The predicate evaluates to true if the client IPv4 Address is covered by one or more of the IPv4 Prefixes in the IP Prefix Sets..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_sets` - (Required) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Ip Prefix List

The predicate evaluates to true if the client IPv4 Address is covered by one or more of the IPv4 Prefixes from the list..

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Required) List of IPv4 prefix strings. (`String`).

### Ip Threat Category List

IP threat categories to choose from.

`ip_threat_categories` - (Required) The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions (`List of Strings`).

### Item

Criteria for matching the values for the cookie. The match is successful if any of the values in the input satisfies the criteria in the matcher..

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Javascript Info

Custom JavaScript Configuration. Custom JavaScript code can be executed at various stages of request processing..

`cache_prefix` - (Optional) KeyValue store referred by script. (`String`).

`custom_script_url` - (Optional) URL of JavaScript that gets executed (`String`).

`script_config` - (Optional) Input passed to the script (`String`).

### Js Challenge

Configure Javascript challenge on this load balancer.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

`js_script_delay` - (Optional) Delay introduced by Javascript, in milliseconds. (`Int`).

### Js Challenge Parameters

Configure Javascript challenge parameters.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

`js_script_delay` - (Optional) Delay introduced by Javascript, in milliseconds. (`Int`).

### Js Insert All Pages

Insert Bot Defense JavaScript in all pages..

`javascript_location` - (Optional) Defines where to insert Bot Defense JavaScript in HTML page. (`String`).

### Js Insert All Pages Except

Insert Bot Defense JavaScript in all pages with the exceptions..

`exclude_list` - (Optional) Optional JavaScript insertions exclude list of domain and path matchers.. See [Exclude List ](#exclude-list) below for details.

`javascript_location` - (Optional) Defines where to insert Bot Defense JavaScript in HTML page. (`String`).

### Js Insertion Rules

Specify custom JavaScript insertion rules..

`exclude_list` - (Optional) Optional JavaScript insertions exclude list of domain and path matchers.. See [Exclude List ](#exclude-list) below for details.

`rules` - (Required) Required list of pages to insert Bot Defense client JavaScript.. See [Rules ](#rules) below for details.

### Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Mirror Policy

useful for logging. For example, *cluster1* becomes *cluster1-shadow*..

`origin_pool` - (Required) referred here must be present.. See [ref](#ref) below for details.

`percent` - (Required) Percentage of requests to be mirrored. See [Percent ](#percent) below for details.

### Mitigation

Mitigation action..

`block` - (Optional) Block bot request and send response with custom content.. See [Block ](#block) below for details.

`flag` - (Optional) Flag the request while not taking any invasive actions.. See [Flag ](#flag) below for details.

`none` - (Optional) No mitigation actions. (bool).

`redirect` - (Optional) Redirect bot request to a custom URI.. See [Redirect ](#redirect) below for details.

### Mobile

Mobile traffic channel..

### Mobile Identifier

Mobile traffic identifier type..

`headers` - (Optional) Headers that can be used to identify mobile traffic.. See [Headers ](#headers) below for details.

### Mobile Sdk Config

Mobile SDK configuration.

`mobile_identifier` - (Optional) Mobile traffic identifier type.. See [Mobile Identifier ](#mobile-identifier) below for details.

`reload_header_name` - (Required) Header that is used for SDK configuration sync. (`String`).

### More Option

More options like header manipulation, compression etc..

`buffer_policy` - (Optional) specify the maximum buffer size and buffer interval with this config.. See [Buffer Policy ](#buffer-policy) below for details.

`compression_params` - (Optional) Only GZIP compression is supported. See [Compression Params ](#compression-params) below for details.

`custom_errors` - (Optional) matches for a request. (`String`).

`disable_default_error_pages` - (Optional) Disable the use of default Volterra error pages. (`Bool`).

`idle_timeout` - (Optional) received, otherwise the stream is reset. (`Int`).

`javascript_info` - (Optional) Custom JavaScript Configuration. Custom JavaScript code can be executed at various stages of request processing.. See [Javascript Info ](#javascript-info) below for details.

`jwt` - (Optional) audiences and issuer. See [ref](#ref) below for details.

`max_request_header_size` - (Optional) such load balancers is used for all the load balancers in question. (`Int`).

`disable_path_normalize` - (Optional) x-displayName: "Disable" (bool).

`enable_path_normalize` - (Optional) x-displayName: "Enable" (bool).

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Request Headers To Add ](#request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Response Headers To Add ](#response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

`additional_domains` - (Optional) Wildcard names are supported in the suffix or prefix form. See [Additional Domains ](#additional-domains) below for details.

`enable_strict_sni_host_header_check` - (Optional) Enable strict SNI and Host header check (bool).

### No Challenge

Challenge rules can be used to selectively enable Javascript or Captcha challenge for some requests..

### No Crl

Client certificate revocation status is not verified.

### No Headers

No mitigation headers..

### No Ip Allowed List

There is no ip allowed list for rate limiting, all clients go through rate limiting..

### No Mtls

x-displayName: "Disable".

### No Policies

Do not apply additional rate limiter policies..

### Non Default Loadbalancer

x-displayName: "No".

### None

No mitigation actions..

### Origin Pools

Origin Pools for this route.

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

`cluster` - (Required) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Required) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`priority` - (Optional) made active as per the increasing priority. (`Int`).

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

### Pass Through

Pass existing server header as is. If server header is absent, a new header is not appended..

### Path

URI path matcher..

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (`String`).

`regex` - (Optional) Regular expression of path match (`String`).

### Percent

Percentage of requests to be mirrored.

`denominator` - (Required) Samples per denominator. numerator part per 100 or 10000 ro 1000000 (`String`).

`numerator` - (Required) sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000 (`Int`).

### Policies

to the action configured in the rule. If there's no match, the rate limiting configuration for the HTTP load balancer is honored..

`policies` - (Required) Ordered list of rate limiter policies.. See [ref](#ref) below for details.

### Policy

Bot Defense Policy..

`disable_js_insert` - (Optional) Disable JavaScript insertion. (bool).

`js_insert_all_pages` - (Optional) Insert Bot Defense JavaScript in all pages.. See [Js Insert All Pages ](#js-insert-all-pages) below for details.

`js_insert_all_pages_except` - (Optional) Insert Bot Defense JavaScript in all pages with the exceptions.. See [Js Insert All Pages Except ](#js-insert-all-pages-except) below for details.

`js_insertion_rules` - (Optional) Specify custom JavaScript insertion rules.. See [Js Insertion Rules ](#js-insertion-rules) below for details.

`js_download_path` - (Optional) Customize Bot Defense Client JavaScript path. If not specified, default `/common.js` (`String`).

`disable_mobile_sdk` - (Optional) Disable Mobile SDK. (bool).

`mobile_sdk_config` - (Optional) Mobile SDK configuration. See [Mobile Sdk Config ](#mobile-sdk-config) below for details.

`protected_app_endpoints` - (Required) List of protected application endpoints (max 128 items).. See [Protected App Endpoints ](#protected-app-endpoints) below for details.

### Policy Based Challenge

Specifies the settings for policy rule based challenge.

`captcha_challenge_parameters` - (Optional) Configure captcha challenge parameters. See [Captcha Challenge Parameters ](#captcha-challenge-parameters) below for details.

`default_captcha_challenge_parameters` - (Optional) Use default parameters (bool).

`always_enable_captcha_challenge` - (Optional) Challenge rules can be used to selectively disable Captcha challenge or enable Javascript challenge for some requests. (bool).

`always_enable_js_challenge` - (Optional) Challenge rules can be used to selectively disable Javascript challenge or enable Captcha challenge for some requests. (bool).

`no_challenge` - (Optional) Challenge rules can be used to selectively enable Javascript or Captcha challenge for some requests. (bool).

`default_js_challenge_parameters` - (Optional) Use default parameters (bool).

`js_challenge_parameters` - (Optional) Configure Javascript challenge parameters. See [Js Challenge Parameters ](#js-challenge-parameters) below for details.

`default_mitigation_settings` - (Optional) For high level, users will be temporarily blocked. (bool).

`malicious_user_mitigation` - (Optional) Define the mitigation actions to be taken for different threat levels. See [ref](#ref) below for details.

`rule_list` - (Optional) list challenge rules to be used in policy based challenge. See [Rule List ](#rule-list) below for details.

`default_temporary_blocking_parameters` - (Optional) Use default parameters (bool).

`temporary_user_blocking` - (Optional) Specifies configuration for temporary user blocking resulting from malicious user detection. See [Temporary User Blocking ](#temporary-user-blocking) below for details.

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Proper Case Header Transformation

For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are”.

### Protected App Endpoints

List of protected application endpoints (max 128 items)..

`mobile` - (Optional) Mobile traffic channel. (bool).

`web` - (Optional) Web traffic channel. (bool).

`web_mobile` - (Optional) Web and mobile traffic channel.. See [Web Mobile ](#web-mobile) below for details.

`any_domain` - (Optional) Any Domain. (bool).

`domain` - (Optional) Domain matcher.. See [Domain ](#domain) below for details.

`http_methods` - (Required) List of HTTP methods. (`List of Strings`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`mitigation` - (Required) Mitigation action.. See [Mitigation ](#mitigation) below for details.

`path` - (Required) Matching URI path of the route.. See [Path ](#path) below for details.

`protocol` - (Optional) Protocol. (`String`).

### Query Params

Note that all specified query parameter predicates must evaluate to true..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`key` - (Required) A case-sensitive HTTP query parameter name. (`String`).

`check_not_present` - (Optional) Check that the query parameter is not present. (bool).

`check_present` - (Optional) Check that the query parameter is present. (bool).

`item` - (Optional) criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the query parameter is present or absent. (`Bool`).

### Rate Limit

Define custom rate limiting parameters for this load balancer.

`custom_ip_allowed_list` - (Optional) IP Allowed list using existing ip_prefix_set objects.. See [Custom Ip Allowed List ](#custom-ip-allowed-list) below for details.

`ip_allowed_list` - (Optional) List of IP(s) for which rate limiting will be disabled.. See [Ip Allowed List ](#ip-allowed-list) below for details.

`no_ip_allowed_list` - (Optional) There is no ip allowed list for rate limiting, all clients go through rate limiting. (bool).

`no_policies` - (Optional) Do not apply additional rate limiter policies. (bool).

`policies` - (Optional) to the action configured in the rule. If there's no match, the rate limiting configuration for the HTTP load balancer is honored.. See [Policies ](#policies) below for details.

`rate_limiter` - (Optional) Requests to the virtual_host are rate limited based on the parameters specified in the rate_limiter.. See [Rate Limiter ](#rate-limiter) below for details.

### Rate Limiter

Requests to the virtual_host are rate limited based on the parameters specified in the rate_limiter..

`burst_multiplier` - (Optional) The maximum burst of requests to accommodate, expressed as a multiple of the rate. (`Int`).

`total_number` - (Required) The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period. (`Int`).

`unit` - (Required) Unit for the period per which the rate limit is applied. (`String`).

### Redirect

Redirect bot request to a custom URI..

`uri` - (Required) URI location for redirect may be relative or absolute. (`String`).

### Redirect Route

A redirect route matches on path and/or HTTP method and redirects the matching traffic to a different URL.

`headers` - (Optional) List of (key, value) headers. See [Headers ](#headers) below for details.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`path` - (Optional) URI path of route. See [Path ](#path) below for details.

`route_redirect` - (Optional) Send redirect response. See [Route Redirect ](#route-redirect) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Remove All Params

x-displayName: "Remove All Parameters".

### Request Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Secret Value ](#secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Request Matcher

Conditions related to the request, such as query parameters, headers, etc..

`cookie_matchers` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Cookie Matchers ](#cookie-matchers) below for details.

`headers` - (Optional) Note that all specified header predicates must evaluate to true.. See [Headers ](#headers) below for details.

`query_params` - (Optional) Note that all specified query parameter predicates must evaluate to true.. See [Query Params ](#query-params) below for details.

### Response Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Secret Value ](#secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Retain All Params

x-displayName: "Retain All Parameters".

### Retract Cluster

for route.

### Retry Policy

Configure custom retry policy.

`back_off` - (Optional) 10 times the base interval. See [Back Off ](#back-off) below for details.

`num_retries` - (Optional) is used between each retry (`Int`).

`per_try_timeout` - (Optional) Specifies a non-zero timeout per retry attempt. In milliseconds (`Int`).

`retriable_status_codes` - (Optional) HTTP status codes that should trigger a retry in addition to those specified by retry_on. (`Int`).

`retry_condition` - (Optional) matching one defined in retriable_status_codes field (`String`).

`retry_on` - (Optional) matching one defined in retriable_status_codes field (`String`).

### Ring Hash

Request are sent to all eligible origin servers using hash of request based on hash policy. Consistent hashing algorithm, ring hash, is used to select origin server.

`hash_policy` - (Required) route the request. See [Hash Policy ](#hash-policy) below for details.

### Route Direct Response

Send direct response.

`response_body` - (Optional) response body to send (`String`).

`response_code` - (Optional) response code to send (`Int`).

### Route Redirect

Send redirect response.

`host_redirect` - (Optional) swap host part of incoming URL in redirect URL (`String`).

`proto_redirect` - (Optional) When incoming-proto option is specified, swapping of protocol is not done. (`String`).

`all_params` - (Optional) be removed. Default value is false, which means query portion of the URL will NOT be removed (`Bool`).

`remove_all_params` - (Optional) x-displayName: "Remove All Parameters" (bool).

`replace_params` - (Optional) x-displayName: "Replace All Parameters" (`String`).

`retain_all_params` - (Optional) x-displayName: "Retain All Parameters" (bool).

`strip_query_params` - (Optional) Specifies the list of query params to be removed. Not supported. See [Strip Query Params ](#strip-query-params) below for details.

`path_redirect` - (Optional) swap path part of incoming URL in redirect URL (`String`).

`prefix_rewrite` - (Optional) This option allows redirect URLs be dynamically created based on the request (`String`).

`response_code` - (Optional) The HTTP status code to use in the redirect response. (`Int`).

### Routes

Routes for this load balancer.

`custom_route_object` - (Optional) A custom route uses a route object created outside of this view.. See [Custom Route Object ](#custom-route-object) below for details.

`direct_response_route` - (Optional) A direct response route matches on path and/or HTTP method and responds directly to the matching traffic. See [Direct Response Route ](#direct-response-route) below for details.

`redirect_route` - (Optional) A redirect route matches on path and/or HTTP method and redirects the matching traffic to a different URL. See [Redirect Route ](#redirect-route) below for details.

`simple_route` - (Optional) A simple route matches on path and/or HTTP method and forwards the matching traffic to the associated pools. See [Simple Route ](#simple-route) below for details.

### Rule List

list challenge rules to be used in policy based challenge.

`rules` - (Optional) these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions. See [Rules ](#rules) below for details.

### Rules

Required list of pages to insert Bot Defense client JavaScript..

`any_domain` - (Optional) Any Domain. (bool).

`domain` - (Optional) Domain matcher.. See [Domain ](#domain) below for details.

`javascript_location` - (Optional) Defines where to insert Bot Defense JavaScript in HTML page. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`path` - (Required) URI path matcher.. See [Path ](#path) below for details.

### Secret Value

Secret Value of the HTTP header..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Server Url Rules

For matching also specific endpoints you can use the API endpoint rules set bellow..

`base_path` - (Required) Prefix of the request path. (`String`).

`any_domain` - (Required) The rule will apply for all domains. (bool).

`specific_domain` - (Required) The rule will apply for a specific domain. (`String`).

`inline_rate_limiter` - (Optional) Specify rate values for the rule.. See [Inline Rate Limiter ](#inline-rate-limiter) below for details.

`ref_rate_limiter` - (Optional) Use external rate limiter.. See [ref](#ref) below for details.

### Simple Route

A simple route matches on path and/or HTTP method and forwards the matching traffic to the associated pools.

`advanced_options` - (Optional) Configure Advanced per route options. See [Advanced Options ](#advanced-options) below for details.

`headers` - (Optional) List of (key, value) headers. See [Headers ](#headers) below for details.

`auto_host_rewrite` - (Optional) Host header will be swapped with hostname of upstream host chosen by the cluster (bool).

`disable_host_rewrite` - (Optional) Host header is not modified (bool).

`host_rewrite` - (Optional) Host header will be swapped with this value (`String`).

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`origin_pools` - (Optional) Origin Pools for this route. See [Origin Pools ](#origin-pools) below for details.

`path` - (Optional) URI path of route. See [Path ](#path) below for details.

### Single Lb App

ML Config applied on this load balancer.

`disable_discovery` - (Optional) x-displayName: "Disable" (bool).

`enable_discovery` - (Optional) x-displayName: "Enable". See [Enable Discovery ](#enable-discovery) below for details.

`disable_ddos_detection` - (Optional) x-displayName: "Disable" (bool).

`enable_ddos_detection` - (Optional) x-displayName: "Enable" (bool).

`disable_malicious_user_detection` - (Optional) x-displayName: "Disable" (bool).

`enable_malicious_user_detection` - (Optional) x-displayName: "Enable" (bool).

### Site

Advertise on a customer site and a given network..

`ip` - (Optional) Use given IP address as VIP on the site (`String`).

`network` - (Required) By default VIP chosen as ip address of primary network interface in the network (`String`).

`site` - (Required) Reference to site object. See [ref](#ref) below for details.

### Skip Data Guard

x-displayName: "Skip".

### Skip Processing

Skip both WAF and Bot Defense processing for clients matching this rule..

### Specific Hash Policy

Configure hash policy specific for this route.

`hash_policy` - (Required) route the request. See [Hash Policy ](#hash-policy) below for details.

### Strip Query Params

Specifies the list of query params to be removed. Not supported.

`query_params` - (Optional) Query params keys to strip while manipulating the HTTP request (`String`).

### Temporary User Blocking

Specifies configuration for temporary user blocking resulting from malicious user detection.

`custom_page` - (Optional) E.g. "<p> Blocked </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

### Tls Certificates

Set of TLS certificates.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. Volterra will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Custom Hash Algorithms ](#custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Disable Ocsp Stapling ](#disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) Volterra will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Use System Defaults ](#use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Custom Security ](#custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (bool).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (bool).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (bool).

### Tls Fingerprint Matcher

The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints..

`classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`exact_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

`excluded_values` - (Optional) or more known TLS fingerprint classes in the enclosing matcher. (`String`).

### Tls Parameters

TLS parameters for downstream connections..

`no_mtls` - (Optional) x-displayName: "Disable" (bool).

`use_mtls` - (Optional) x-displayName: "Enable". See [Use Mtls ](#use-mtls) below for details.

`tls_certificates` - (Required) Set of TLS certificates. See [Tls Certificates ](#tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Config ](#tls-config) below for details.

### Trusted Clients

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients..

`bot_skip_processing` - (Optional) Skip Bot Defense processing for clients matching this rule. (bool).

`skip_processing` - (Optional) Skip both WAF and Bot Defense processing for clients matching this rule. (bool).

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (bool).

`actions` - (Optional) Action that should be taken when client identifier matches the rule (`List of Strings`).

`as_number` - (Required) RFC 6793 defined 4-byte AS number (`Int`).

`http_header` - (Required) Request header name and value pairs. See [Http Header ](#http-header) below for details.

`ip_prefix` - (Required) IPv4 prefix string. (`String`).

`user_identifier` - (Optional) Identify user based on user identifier. User identifier value needs to be copied from security event. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

### Use Default Port

For HTTP, default is 80. For HTTPS/SNI, default is 443..

### Use Http Lb User Id

Defined in HTTP-LB Security Configuration -> User Identifier..

### Use Mtls

x-displayName: "Enable".

`crl` - (Optional) Specify the CRL server information to download the certificate revocation list. See [ref](#ref) below for details.

`no_crl` - (Optional) Client certificate revocation status is not verified (bool).

`trusted_ca_url` - (Required) The URL for a trust store (`String`).

### Use System Defaults

Volterra will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Virtual Network

Advertise on a virtual network.

`default_vip` - (Optional) Use the default VIP, system allocated or configured in the virtual network (bool).

`specific_vip` - (Optional) Use given IP address as VIP on VoltADN private Network (`String`).

`virtual_network` - (Required) Select virtual network reference. See [ref](#ref) below for details.

### Virtual Site

Advertise on a customer virtual site and a given network..

`network` - (Required) IP address of primary network interface in the network (`String`).

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Vk8s Service

Advertise on vK8s Service Network on RE..

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Waf Exclusion Rules

The match criteria include domain, path and method..

`app_firewall_detection_control` - (Optional) Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.. See [App Firewall Detection Control ](#app-firewall-detection-control) below for details.

`any_domain` - (Optional) Apply this WAF exclusion rule for any domain (bool).

`exact_value` - (Optional) Exact domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

`exclude_rule_ids` - (Required) WAF Rules to be excluded when match conditions are met (`List of Strings`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`methods` - (Optional) methods to be matched (`List of Strings`).

`path_regex` - (Required) path regex to be matched (`String`).

### Waf Skip Processing

Skip WAF processing for clients matching this rule..

### Web

Web traffic channel..

### Web Mobile

Web and mobile traffic channel..

`header` - (Optional) Header that is used by mobile traffic.. See [Header ](#header) below for details.

`headers` - (Optional) Headers that can be used to identify mobile traffic.. See [Headers ](#headers) below for details.

`mobile_identifier` - (Optional) Mobile identifier type (`String`).

### Web Socket Config

Upgrade to Websocket for this route.

`idle_timeout` - (Optional) Idle Timeout for Websocket in milli seconds. After timeout, connection will be closed (`Int`).

`max_connect_attempts` - (Optional) giving up. Default is 1 (`Int`).

`use_websocket` - (Optional) a WebSocket connection (`Bool`).

### Wingman Secret Info

Secret is given as bootstrap secret in Volterra Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured http_loadbalancer.
