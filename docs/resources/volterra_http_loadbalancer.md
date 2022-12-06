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

  advertise_custom {
    advertise_where {
      // One of the arguments from this list "site virtual_site vk8s_service virtual_network" must be set

      virtual_network {
        // One of the arguments from this list "default_vip specific_vip" must be set
        default_vip = true

        virtual_network {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
      }

      // One of the arguments from this list "use_default_port port" must be set
      use_default_port = true
    }
  }
  // One of the arguments from this list "disable_api_definition api_definition api_definitions" must be set
  disable_api_definition = true

  // One of the arguments from this list "enable_api_discovery disable_api_discovery" must be set

  enable_api_discovery {
    // One of the arguments from this list "disable_learn_from_redirect_traffic enable_learn_from_redirect_traffic" must be set
    disable_learn_from_redirect_traffic = true
  }

  // One of the arguments from this list "captcha_challenge policy_based_challenge no_challenge js_challenge" must be set

  js_challenge {
    cookie_expiry   = "1000"
    custom_page     = "string:///PHA+IFBsZWFzZSBXYWl0IDwvcD4="
    js_script_delay = "1000"
  }
  // One of the arguments from this list "enable_ddos_detection disable_ddos_detection" must be set
  disable_ddos_detection = true
  domains = ["www.foo.com"]
  // One of the arguments from this list "ring_hash round_robin least_active random source_ip_stickiness cookie_stickiness" must be set
  random = true

  // One of the arguments from this list "https http https_auto_cert" must be set

  http {
    dns_volterra_managed = true
    port                 = "80"
  }
  // One of the arguments from this list "enable_malicious_user_detection disable_malicious_user_detection" must be set
  enable_malicious_user_detection = true
  // One of the arguments from this list "disable_rate_limit api_rate_limit rate_limit" must be set
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

  user_identification {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }
  // One of the arguments from this list "disable_waf waf waf_rule app_firewall" must be set
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

`csrf_policy` - (Optional) Policy configuration to protect against CSRF attacks.. See [Csrf Policy ](#csrf-policy) below for details.

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

`default_pool` - (Optional) Single Origin Pool. See [Default Pool ](#default-pool) below for details.

`default_pool_list` - (Optional) Multiple Origin Pools with weights and priorities. See [Default Pool List ](#default-pool-list) below for details.

`api_rate_limit` - (Optional) Define rate limiting for one or more API endpoints. See [Api Rate Limit ](#api-rate-limit) below for details.

`disable_rate_limit` - (Optional) Rate limiting is not currently enabled for this load balancer (bool).

`rate_limit` - (Optional) Define custom rate limiting parameters for this load balancer. See [Rate Limit ](#rate-limit) below for details.

`routes` - (Optional) to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic. See [Routes ](#routes) below for details.

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

### Account Management

x-displayName: "Account Management".

`create` - (Optional) x-displayName: "Account Creation" (bool).

`password_reset` - (Optional) x-displayName: "Password Reset" (bool).

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

Advanced options configuration like timeouts, circuit breaker, subset load balancing.

`circuit_breaker` - (Optional) allows to apply back pressure on downstream quickly.. See [Circuit Breaker ](#circuit-breaker) below for details.

`default_circuit_breaker` - (Optional) requests are 1024 and the default value for retries is 3 (bool).

`disable_circuit_breaker` - (Optional) Circuit Breaker is disabled (bool).

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2 seconds (`Int`).

`header_transformation_type` - (Optional) Settings to normalize the headers of upstream requests.. See [Header Transformation Type ](#header-transformation-type) below for details.

`http2_options` - (Optional) Http2 Protocol options for upstream connections. See [Http2 Options ](#http2-options) below for details.

`http_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 5 minutes. (`Int`).

`disable_outlier_detection` - (Optional) Outlier detection is disabled (bool).

`outlier_detection` - (Optional) healthy load balancing set. Outlier detection is a form of passive health checking.. See [Outlier Detection ](#outlier-detection) below for details.

`no_panic_threshold` - (Optional) Disable panic threshold. Only healthy endpoints are considered for load balancing. (bool).

`panic_threshold` - (Optional) all endpoints will be considered for load balancing ignoring its health status. (`Int`).

`disable_subsets` - (Optional) Subset load balancing is disabled. All eligible origin servers will be considered for load balancing. (bool).

`enable_subsets` - (Optional) Subset load balancing is enabled. Based on route, subset of origin servers will be considered for load balancing.. See [Enable Subsets ](#enable-subsets) below for details.

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

### All Load Balancer Domains

Add All load balancer domains to source origin (allow) list..

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

### Any Endpoint

Select any origin server from available healthy origin servers in this pool.

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

### Apply

x-displayName: "Apply for a Financial Service Account (e.g., credit card, banking, retirement account)".

### Apply Data Guard

x-displayName: "Apply".

### Asn List

The predicate evaluates to true if the origin ASN is present in the ASN list..

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. (`Int`).

### Asn Matcher

The predicate evaluates to true if the origin ASN is present in one of the BGP ASN Set objects..

`asn_sets` - (Required) A list of references to bgp_asn_set objects.. See [ref](#ref) below for details.

### Authentication

x-displayName: "Authentication".

`login` - (Optional) x-displayName: "Login" (bool).

`login_mfa` - (Optional) x-displayName: "Login MFA" (bool).

`login_partner` - (Optional) x-displayName: "Login for a Channel Partner" (bool).

`logout` - (Optional) x-displayName: "Logout" (bool).

`token_refresh` - (Optional) x-displayName: "Token Refresh" (bool).

### Auto Host Rewrite

Host header will be swapped with hostname of upstream host chosen by the cluster.

### Automatic Port

For other origin server types, port will be automatically set as 443 if TLS is enabled at Origin Pool and 80 if TLS is disabled.

### Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

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

`actions` - (Optional) Actions that should be taken when client identifier matches the rule (`List of Strings`).

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

### Checkin

x-displayName: "Check into Flight".

### Circuit Breaker

allows to apply back pressure on downstream quickly..

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

`policy` - (Required) Please ensure that the same domains are configured in the Client-Side Defense configuration.. See [Policy ](#policy) below for details.

### Compression Params

Only GZIP compression is supported.

`content_length` - (Optional) Minimum response length, in bytes, which will trigger compression. The default value is 30. (`Int`).

`content_type` - (Optional) "text/xml" (`String`).

`disable_on_etag_header` - (Optional) weak etags will be preserved and the ones that require strong validation will be removed. (`Bool`).

`remove_accept_encoding_header` - (Optional) so that responses do not get compressed before reaching the filter. (`Bool`).

### Consul Service

Specify origin server with Hashi Corp Consul service name and site information.

`inside_network` - (Optional) Inside network on the site (bool).

`outside_network` - (Optional) Outside network on the site (bool).

`service_name` - (Required) cluster-id is optional. (`String`).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Site Locator ](#site-locator) below for details.

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

### Create

x-displayName: "Account Creation".

### Csrf Policy

Policy configuration to protect against CSRF attacks..

`all_load_balancer_domains` - (Optional) Add All load balancer domains to source origin (allow) list. (bool).

`custom_domain_list` - (Optional) Add one or more domains to source origin (allow) list.. See [Custom Domain List ](#custom-domain-list) below for details.

### Custom Domain List

Add one or more domains to source origin (allow) list..

`domains` - (Optional) Wildcard names are supported in the suffix or prefix form. (`String`).

### Custom Endpoint Object

Specify origin server with a reference to endpoint object.

`endpoint` - (Required) Reference to an endpoint object. See [ref](#ref) below for details.

### Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

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

### Default Circuit Breaker

requests are 1024 and the default value for retries is 3.

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

### Default Pool

Single Origin Pool.

`advanced_options` - (Optional) Advanced options configuration like timeouts, circuit breaker, subset load balancing. See [Advanced Options ](#advanced-options) below for details.

`endpoint_selection` - (Required) Policy for selection of endpoints from local site or remote site or both (`String`).

`health_check_port` - (Optional) Port used for performing health check (`Int`).

`same_as_endpoint_port` - (Optional) Health check is performed on endpoint port itself (bool).

`healthcheck` - (Optional) Reference to healthcheck configuration objects. See [ref](#ref) below for details.

`loadbalancer_algorithm` - (Required) loadbalancer_algorithm to determine which host is selected. (`String`).

`origin_servers` - (Required) List of origin servers in this pool. See [Origin Servers ](#origin-servers) below for details.

`automatic_port` - (Optional) For other origin server types, port will be automatically set as 443 if TLS is enabled at Origin Pool and 80 if TLS is disabled (bool).

`port` - (Optional) Endpoint service is available on this port (`Int`).

`no_tls` - (Optional) x-displayName: "Disable" (bool).

`use_tls` - (Optional) x-displayName: "Enable". See [Use Tls ](#use-tls) below for details.

`view_internal` - (Optional) Reference to view internal object. See [ref](#ref) below for details.

### Default Pool List

Multiple Origin Pools with weights and priorities.

`pools` - (Optional) List of Origin Pools. See [Pools ](#pools) below for details.

### Default Route Pools

Origin Pools used when no route is specified (default route).

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

`cluster` - (Required) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Required) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`priority` - (Optional) made active as per the increasing priority. (`Int`).

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

### Default Security

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Default Subset

Use the default subset provided here. Select endpoints matching default subset..

`default_subset` - (Optional) which gets used when route specifies no metadata or no subset matching the metadata exists. (`String`).

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

### Disable Circuit Breaker

Circuit Breaker is disabled.

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

### Disable Mobile Sdk

Disable Mobile SDK..

### Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Disable Outlier Detection

Outlier detection is disabled.

### Disable Path Normalize

x-displayName: "Disable".

### Disable Sni

Do not use SNI..

### Disable Subsets

Subset load balancing is disabled. All eligible origin servers will be considered for load balancing..

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

### Enable Strict Sni Host Header Check

Enable strict SNI and Host header check.

### Enable Subsets

Subset load balancing is enabled. Based on route, subset of origin servers will be considered for load balancing..

`endpoint_subsets` - (Optional) List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class.. See [Endpoint Subsets ](#endpoint-subsets) below for details.

`any_endpoint` - (Optional) Select any origin server from available healthy origin servers in this pool (bool).

`default_subset` - (Optional) Use the default subset provided here. Select endpoints matching default subset.. See [Default Subset ](#default-subset) below for details.

`fail_request` - (Optional) Request will be failed and error returned, as if cluster has no origin servers. (bool).

### Enable Trust Client Ip Headers

x-displayName: "Enable".

`client_ip_headers` - (Optional) For X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client ip (`String`).

### Endpoint Subsets

List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class..

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

### Fail Request

Request will be failed and error returned, as if cluster has no origin servers..

### Financial Services

x-displayName: "Financial Services".

`apply` - (Optional) x-displayName: "Apply for a Financial Service Account (e.g., credit card, banking, retirement account)" (bool).

`money_transfer` - (Optional) x-displayName: "Money Transfer" (bool).

### Flag

Flag the request while not taking any invasive actions..

`append_headers` - (Optional) Append mitigation headers.. See [Append Headers ](#append-headers) below for details.

`no_headers` - (Optional) No mitigation headers. (bool).

### Flight

x-displayName: "Flight".

`checkin` - (Optional) x-displayName: "Check into Flight" (bool).

### Flight Search

x-displayName: "Flight Search".

### Flow Label

x-displayName: "Specify flow label category".

`account_management` - (Optional) x-displayName: "Account Management". See [Account Management ](#account-management) below for details.

`authentication` - (Optional) x-displayName: "Authentication". See [Authentication ](#authentication) below for details.

`financial_services` - (Optional) x-displayName: "Financial Services". See [Financial Services ](#financial-services) below for details.

`flight` - (Optional) x-displayName: "Flight". See [Flight ](#flight) below for details.

`profile_management` - (Optional) x-displayName: "Profile Management". See [Profile Management ](#profile-management) below for details.

`search` - (Optional) x-displayName: "Search". See [Search ](#search) below for details.

`shopping_gift_cards` - (Optional) x-displayName: "Shopping & Gift Cards". See [Shopping Gift Cards ](#shopping-gift-cards) below for details.

### Gift Card Make Purchase With Gift Card

x-displayName: "Purchase with Gift Card".

### Gift Card Validation

x-displayName: "Gift Card Validation".

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

### Http2 Options

Http2 Protocol options for upstream connections.

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

### Inline Rate Limiter

Specify rate values for the rule..

`ref_user_id` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier to be rate limited.. See [ref](#ref) below for details.

`use_http_lb_user_id` - (Optional) Defined in HTTP-LB Security Configuration -> User Identifier. (bool).

`threshold` - (Required) The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period. (`Int`).

`unit` - (Required) Unit for the period per which the rate limit is applied. (`String`).

### Inside Network

Inside network on the site.

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

### K8s Service

Specify origin server with K8s service name and site information.

`inside_network` - (Optional) Inside network on the site (bool).

`outside_network` - (Optional) Outside network on the site (bool).

`vk8s_networks` - (Optional) origin server are on vK8s network on the site (bool).

`service_name` - (Required) Both namespace and cluster-id are optional. (`String`).

`service_selector` - (Required) discovery has to happen. This implicit label is added to service_selector. See [Service Selector ](#service-selector) below for details.

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Site Locator ](#site-locator) below for details.

### Login

x-displayName: "Login".

### Login Mfa

x-displayName: "Login MFA".

### Login Partner

x-displayName: "Login for a Channel Partner".

### Logout

x-displayName: "Logout".

### Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

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

### Money Transfer

x-displayName: "Money Transfer".

### More Option

More options like header manipulation, compression etc..

`buffer_policy` - (Optional) specify the maximum buffer size and buffer interval with this config.. See [Buffer Policy ](#buffer-policy) below for details.

`compression_params` - (Optional) Only GZIP compression is supported. See [Compression Params ](#compression-params) below for details.

`custom_errors` - (Optional) matches for a request. (`String`).

`disable_default_error_pages` - (Optional) Disable the use of default F5XC error pages. (`Bool`).

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

### No Panic Threshold

Disable panic threshold. Only healthy endpoints are considered for load balancing..

### No Policies

Do not apply additional rate limiter policies..

### No Tls

x-displayName: "Disable".

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

### Origin Servers

List of origin servers in this pool.

`consul_service` - (Optional) Specify origin server with Hashi Corp Consul service name and site information. See [Consul Service ](#consul-service) below for details.

`custom_endpoint_object` - (Optional) Specify origin server with a reference to endpoint object. See [Custom Endpoint Object ](#custom-endpoint-object) below for details.

`k8s_service` - (Optional) Specify origin server with K8s service name and site information. See [K8s Service ](#k8s-service) below for details.

`private_ip` - (Optional) Specify origin server with private or public IP address and site information. See [Private Ip ](#private-ip) below for details.

`private_name` - (Optional) Specify origin server with private or public DNS name and site information. See [Private Name ](#private-name) below for details.

`public_ip` - (Optional) Specify origin server with public IP. See [Public Ip ](#public-ip) below for details.

`public_name` - (Optional) Specify origin server with public DNS name. See [Public Name ](#public-name) below for details.

`vn_private_ip` - (Optional) Specify origin server IP address on virtual network other than inside or outside network. See [Vn Private Ip ](#vn-private-ip) below for details.

`vn_private_name` - (Optional) Specify origin server name on virtual network other than inside or outside network. See [Vn Private Name ](#vn-private-name) below for details.

`labels` - (Optional) Add Labels for this origin server, these labels can be used to form subset. (`String`).

### Outlier Detection

healthy load balancing set. Outlier detection is a form of passive health checking..

### Outside Network

Outside network on the site.

### Pass Through

Pass existing server header as is. If server header is absent, a new header is not appended..

### Password Reset

x-displayName: "Password Reset".

### Path

URI path matcher..

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (`String`).

`regex` - (Optional) Regular expression of path match (`String`).

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

### Pools

List of Origin Pools.

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

`cluster` - (Required) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Required) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`priority` - (Optional) made active as per the increasing priority. (`Int`).

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

### Private Ip

Specify origin server with private or public IP address and site information.

`ip` - (Required) IP address (`String`).

`inside_network` - (Optional) Inside network on the site (bool).

`outside_network` - (Optional) Outside network on the site (bool).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Site Locator ](#site-locator) below for details.

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Private Name

Specify origin server with private or public DNS name and site information.

`dns_name` - (Required) DNS Name (`String`).

`inside_network` - (Optional) Inside network on the site (bool).

`outside_network` - (Optional) Outside network on the site (bool).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Site Locator ](#site-locator) below for details.

### Product Search

x-displayName: "Product Search".

### Profile Management

x-displayName: "Profile Management".

`create` - (Optional) x-displayName: "Profile Creation" (bool).

`update` - (Optional) x-displayName: "Profile Update" (bool).

`view` - (Optional) x-displayName: "Profile View" (bool).

### Proper Case Header Transformation

For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are”.

### Protected App Endpoints

List of protected application endpoints (max 128 items)..

`mobile` - (Optional) Mobile traffic channel. (bool).

`web` - (Optional) Web traffic channel. (bool).

`web_mobile` - (Optional) Web and mobile traffic channel.. See [Web Mobile ](#web-mobile) below for details.

`any_domain` - (Optional) Any Domain. (bool).

`domain` - (Optional) Domain matcher.. See [Domain ](#domain) below for details.

`flow_label` - (Optional) x-displayName: "Specify flow label category". See [Flow Label ](#flow-label) below for details.

`undefined_flow_label` - (Optional) x-displayName: "Undefined" (bool).

`http_methods` - (Required) List of HTTP methods. (`List of Strings`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`mitigation` - (Required) Mitigation action.. See [Mitigation ](#mitigation) below for details.

`path` - (Required) Matching URI path of the route.. See [Path ](#path) below for details.

`protocol` - (Optional) Protocol. (`String`).

### Public Ip

Specify origin server with public IP.

`ip` - (Required) Public IP address (`String`).

### Public Name

Specify origin server with public DNS name.

`dns_name` - (Required) DNS Name (`String`).

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

### Reservation Search

x-displayName: "Reservation Search (e.g., sporting events, concerts)".

### Response Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Secret Value ](#secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Retain All Params

x-displayName: "Retain All Parameters".

### Ring Hash

Request are sent to all eligible origin servers using hash of request based on hash policy. Consistent hashing algorithm, ring hash, is used to select origin server.

`hash_policy` - (Required) route the request. See [Hash Policy ](#hash-policy) below for details.

### Room Search

x-displayName: "Room Search".

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

to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic.

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

### Same As Endpoint Port

Health check is performed on endpoint port itself.

### Search

x-displayName: "Search".

`flight_search` - (Optional) x-displayName: "Flight Search" (bool).

`product_search` - (Optional) x-displayName: "Product Search" (bool).

`reservation_search` - (Optional) x-displayName: "Reservation Search (e.g., sporting events, concerts)" (bool).

`room_search` - (Optional) x-displayName: "Room Search" (bool).

### Secret Value

Secret Value of the HTTP header..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Server Url Rules

For matching also specific endpoints you can use the API endpoint rules set bellow..

`base_path` - (Required) Prefix of the request path. (`String`).

`any_domain` - (Required) The rule will apply for all domains. (bool).

`specific_domain` - (Required) The rule will apply for a specific domain. (`String`).

`inline_rate_limiter` - (Optional) Specify rate values for the rule.. See [Inline Rate Limiter ](#inline-rate-limiter) below for details.

`ref_rate_limiter` - (Optional) Use external rate limiter.. See [ref](#ref) below for details.

### Service Selector

discovery has to happen. This implicit label is added to service_selector.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Shop Add To Cart

x-displayName: "Add to Cart".

### Shop Checkout

x-displayName: "Checkout".

### Shop Choose Seat

x-displayName: "Select Seat(s)".

### Shop Enter Drawing Submission

x-displayName: "Enter Drawing Submission".

### Shop Make Payment

x-displayName: "Payment / Billing".

### Shop Order

x-displayName: "Order Submit".

### Shop Price Inquiry

x-displayName: "Price Inquiry".

### Shop Promo Code Validation

x-displayName: "Promo Code Validation".

### Shop Purchase Gift Card

x-displayName: "Purchase a Gift Card".

### Shop Update Quantity

x-displayName: "Update Quantity".

### Shopping Gift Cards

x-displayName: "Shopping & Gift Cards".

`gift_card_make_purchase_with_gift_card` - (Optional) x-displayName: "Purchase with Gift Card" (bool).

`gift_card_validation` - (Optional) x-displayName: "Gift Card Validation" (bool).

`shop_add_to_cart` - (Optional) x-displayName: "Add to Cart" (bool).

`shop_checkout` - (Optional) x-displayName: "Checkout" (bool).

`shop_choose_seat` - (Optional) x-displayName: "Select Seat(s)" (bool).

`shop_enter_drawing_submission` - (Optional) x-displayName: "Enter Drawing Submission" (bool).

`shop_make_payment` - (Optional) x-displayName: "Payment / Billing" (bool).

`shop_order` - (Optional) x-displayName: "Order Submit" (bool).

`shop_price_inquiry` - (Optional) x-displayName: "Price Inquiry" (bool).

`shop_promo_code_validation` - (Optional) x-displayName: "Promo Code Validation" (bool).

`shop_purchase_gift_card` - (Optional) x-displayName: "Purchase a Gift Card" (bool).

`shop_update_quantity` - (Optional) x-displayName: "Update Quantity" (bool).

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

### Site Locator

Site or Virtual site where this origin server is located.

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Skip Data Guard

x-displayName: "Skip".

### Skip Processing

Skip both WAF and Bot Defense processing for clients matching this rule..

### Skip Server Verification

Skip origin server verification.

### Strip Query Params

Specifies the list of query params to be removed. Not supported.

`query_params` - (Optional) Query params keys to strip while manipulating the HTTP request (`String`).

### Temporary User Blocking

Specifies configuration for temporary user blocking resulting from malicious user detection.

`custom_page` - (Optional) E.g. "<p> Blocked </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

### Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Custom Hash Algorithms ](#custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Disable Ocsp Stapling ](#disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Use System Defaults ](#use-system-defaults) below for details.

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

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Certificates ](#tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Config ](#tls-config) below for details.

### Token Refresh

x-displayName: "Token Refresh".

### Trusted Clients

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients..

`bot_skip_processing` - (Optional) Skip Bot Defense processing for clients matching this rule. (bool).

`skip_processing` - (Optional) Skip both WAF and Bot Defense processing for clients matching this rule. (bool).

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (bool).

`actions` - (Optional) Actions that should be taken when client identifier matches the rule (`List of Strings`).

`as_number` - (Required) RFC 6793 defined 4-byte AS number (`Int`).

`http_header` - (Required) Request header name and value pairs. See [Http Header ](#http-header) below for details.

`ip_prefix` - (Required) IPv4 prefix string. (`String`).

`user_identifier` - (Optional) Identify user based on user identifier. User identifier value needs to be copied from security event. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

### Undefined Flow Label

x-displayName: "Undefined".

### Update

x-displayName: "Profile Update".

### Use Default Port

For HTTP, default is 80. For HTTPS/SNI, default is 443..

### Use Host Header As Sni

Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied..

### Use Http Lb User Id

Defined in HTTP-LB Security Configuration -> User Identifier..

### Use Mtls

x-displayName: "Enable".

`crl` - (Optional) Specify the CRL server information to download the certificate revocation list. See [ref](#ref) below for details.

`no_crl` - (Optional) Client certificate revocation status is not verified (bool).

`trusted_ca_url` - (Required) The URL for a trust store (`String`).

### Use Server Verification

Perform origin server verification using the provided trusted CA list.

`trusted_ca_url` - (Required) Trusted CA certificates for verification of Server's certificate (`String`).

### Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Use Tls

x-displayName: "Enable".

`no_mtls` - (Optional) x-displayName: "Disable" (bool).

`use_mtls` - (Optional) x-displayName: "Enable". See [Use Mtls ](#use-mtls) below for details.

`skip_server_verification` - (Optional) Skip origin server verification (bool).

`use_server_verification` - (Optional) Perform origin server verification using the provided trusted CA list. See [Use Server Verification ](#use-server-verification) below for details.

`volterra_trusted_ca` - (Optional) Perform origin server verification using F5XC default trusted CA list (bool).

`disable_sni` - (Optional) Do not use SNI. (bool).

`sni` - (Optional) SNI value to be used. (`String`).

`use_host_header_as_sni` - (Optional) Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied. (bool).

`tls_config` - (Required) TLS parameters such as min/max TLS version and ciphers. See [Tls Config ](#tls-config) below for details.

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### View

x-displayName: "Profile View".

### Virtual Network

Advertise on a virtual network.

`default_vip` - (Optional) Use the default VIP, system allocated or configured in the virtual network (bool).

`specific_vip` - (Optional) Use given IP address as VIP on VoltADN private Network (`String`).

`virtual_network` - (Required) Select virtual network reference. See [ref](#ref) below for details.

### Virtual Site

Advertise on a customer virtual site and a given network..

`network` - (Required) IP address of primary network interface in the network (`String`).

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Vk8s Networks

origin server are on vK8s network on the site.

### Vk8s Service

Advertise on vK8s Service Network on RE..

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Vn Private Ip

Specify origin server IP address on virtual network other than inside or outside network.

`ip` - (Required) IP address (`String`).

`virtual_network` - (Required) Virtual Network where this IP will be present. See [ref](#ref) below for details.

### Vn Private Name

Specify origin server name on virtual network other than inside or outside network.

`dns_name` - (Required) DNS Name (`String`).

`private_network` - (Required) Virtual Network where this Name will be present. See [ref](#ref) below for details.

### Volterra Trusted Ca

Perform origin server verification using F5XC default trusted CA list.

### Waf Exclusion Rules

The match criteria include domain, path and method..

`any_domain` - (Optional) Apply this WAF exclusion rule for any domain (bool).

`exact_value` - (Optional) Exact domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

`exclude_rule_ids` - (Required) WAF Rules to be excluded when match conditions are met (`List of Strings`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`methods` - (Optional) methods to be matched (`List of Strings`).

`path_regex` - (Required) path regex to be matched (`String`).

`app_firewall_detection_control` - (Optional) Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.. See [App Firewall Detection Control ](#app-firewall-detection-control) below for details.

`waf_skip_processing` - (Optional) Skip all App Firewall processing for this request (bool).

### Waf Skip Processing

Skip WAF processing for clients matching this rule..

### Web

Web traffic channel..

### Web Mobile

Web and mobile traffic channel..

`header` - (Optional) Header that is used by mobile traffic.. See [Header ](#header) below for details.

`headers` - (Optional) Headers that can be used to identify mobile traffic.. See [Headers ](#headers) below for details.

`mobile_identifier` - (Optional) Mobile identifier type (`String`).

### Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured http_loadbalancer.
