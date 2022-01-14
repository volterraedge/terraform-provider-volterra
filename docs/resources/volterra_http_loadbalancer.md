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

  // One of the arguments from this list "advertise_on_public advertise_custom do_not_advertise advertise_on_public_default_vip" must be set

  advertise_custom {
    advertise_where {
      // One of the arguments from this list "virtual_network site virtual_site vk8s_service" must be set

      site {
        ip      = "ip"
        network = "network"

        site {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
      }

      // One of the arguments from this list "use_default_port port" must be set
      use_default_port = true
    }
  }

  // One of the arguments from this list "no_challenge js_challenge captcha_challenge policy_based_challenge" must be set

  js_challenge {
    cookie_expiry   = "cookie_expiry"
    custom_page     = "string:///PHA+IFBsZWFzZSBXYWl0IDwvcD4="
    js_script_delay = "js_script_delay"
  }
  domains = ["www.foo.com"]
  // One of the arguments from this list "ring_hash round_robin least_active random source_ip_stickiness cookie_stickiness" must be set
  source_ip_stickiness = true

  // One of the arguments from this list "http https_auto_cert https" must be set

  http {
    dns_volterra_managed = true
  }

  // One of the arguments from this list "single_lb_app multi_lb_app" must be set

  single_lb_app {
    // One of the arguments from this list "enable_discovery disable_discovery" must be set

    enable_discovery {
      // One of the arguments from this list "disable_learn_from_redirect_traffic enable_learn_from_redirect_traffic" must be set
      disable_learn_from_redirect_traffic = true
    }

    // One of the arguments from this list "enable_ddos_detection disable_ddos_detection" must be set
    enable_ddos_detection = true

    // One of the arguments from this list "enable_malicious_user_detection disable_malicious_user_detection" must be set
    enable_malicious_user_detection = true
  }
  // One of the arguments from this list "rate_limit disable_rate_limit" must be set
  disable_rate_limit = true
  // One of the arguments from this list "service_policies_from_namespace no_service_policies active_service_policies" must be set
  service_policies_from_namespace = true

  // One of the arguments from this list "user_id_client_ip user_identification" must be set

  user_identification {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }

  // One of the arguments from this list "waf_rule app_firewall disable_waf waf" must be set

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

`advertise_custom` - (Optional) Advertise this loadbalancer on specific sites. See [Advertise Custom ](#advertise-custom) below for details.

`advertise_on_public` - (Optional) Advertise this loadbalancer on public network. See [Advertise On Public ](#advertise-on-public) below for details.

`advertise_on_public_default_vip` - (Optional) Advertise this loadbalancer on public network with default VIP (bool).

`do_not_advertise` - (Optional) Do not advertise this loadbalancer (bool).

`api_definitions` - (Optional) Use API definitions from OpenAPI specification files to specify API operations of an application. See [Api Definitions ](#api-definitions) below for details.

`blocked_clients` - (Optional) Rules that specify the clients to be blocked. See [Blocked Clients ](#blocked-clients) below for details.

`bot_defense` - (Optional) Bot Defense configuration object. See [Bot Defense ](#bot-defense) below for details.

`disable_bot_defense` - (Optional) No Bot Defense configuration for this load balancer (bool).

`captcha_challenge` - (Optional) Configure Captcha challenge on this load balancer. See [Captcha Challenge ](#captcha-challenge) below for details.

`js_challenge` - (Optional) Configure Javascript challenge on this load balancer. See [Js Challenge ](#js-challenge) below for details.

`no_challenge` - (Optional) No challenge is enabled for this load balancer (bool).

`policy_based_challenge` - (Optional) Specifies the settings for policy rule based challenge. See [Policy Based Challenge ](#policy-based-challenge) below for details.

`cors_policy` - (Optional) resources from a server at a different origin. See [Cors Policy ](#cors-policy) below for details.

`ddos_mitigation_rules` - (Optional) Rules that specify the DDoS clients to be blocked. See [Ddos Mitigation Rules ](#ddos-mitigation-rules) below for details.

`default_route_pools` - (Optional) Origin Pools used when no route is specified (default route). See [Default Route Pools ](#default-route-pools) below for details.

`domains` - (Required) Domains also indicate the list of names for which DNS resolution will be done by VER (`List of String`).

`cookie_stickiness` - (Optional) Request are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server. See [Cookie Stickiness ](#cookie-stickiness) below for details.

`least_active` - (Optional) Request are sent to origin server that has least active requests (bool).

`random` - (Optional) Request are sent to all eligible origin servers in random fashion (bool).

`ring_hash` - (Optional) Request are sent to all eligible origin servers using hash of request based on hash policy. Consistent hashing algorithm, ring hash, is used to select origin server. See [Ring Hash ](#ring-hash) below for details.

`round_robin` - (Optional) Request are sent to all eligible origin servers in round robin fashion (bool).

`source_ip_stickiness` - (Optional) Request are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server (bool).

`http` - (Optional) HTTP Load balancer.. See [Http ](#http) below for details.

`https` - (Optional) User is responsible for managing DNS to this Load Balancer.. See [Https ](#https) below for details.

`https_auto_cert` - (Optional) DNS records will be managed by Volterra.. See [Https Auto Cert ](#https-auto-cert) below for details.

`malicious_user_mitigation` - (Optional) The settings defined in malicious user mitigation specify what mitigation actions to take for users determined to be at different threat levels.. See [ref](#ref) below for details.

`multi_lb_app` - (Optional) ML config is shared among multiple HTTP Load Balancers. It should be set externally (bool).

`single_lb_app` - (Optional) ML Config applied on this Load Balancer. See [Single Lb App ](#single-lb-app) below for details.

`more_option` - (Optional) More options like header manipulation, compression etc.. See [More Option ](#more-option) below for details.

`disable_rate_limit` - (Optional) Rate limiting is not currently enabled for this loadbalancer (bool).

`rate_limit` - (Optional) Rate limiting parameters for this loadbalancer. See [Rate Limit ](#rate-limit) below for details.

`routes` - (Optional) Routes for this loadbalancer. See [Routes ](#routes) below for details.

`active_service_policies` - (Optional) Apply the specified list of service policies and bypass the namespace service policy set. See [Active Service Policies ](#active-service-policies) below for details.

`no_service_policies` - (Optional) Do not apply any service policies i.e. bypass the namespace service policy set (bool).

`service_policies_from_namespace` - (Optional) Apply the service policies configured as part of the namespace service policy set (bool).

`trusted_clients` - (Optional) WAF or/and Bot processing can be skipped for trusted clients. See [Trusted Clients ](#trusted-clients) below for details.

`user_id_client_ip` - (Optional) Use the Client IP address as the user identifier. (bool).

`user_identification` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier.. See [ref](#ref) below for details.

`app_firewall` - (Optional) Reference to App Firewall configuration object. See [ref](#ref) below for details.

`disable_waf` - (Optional) No WAF configuration for this load balancer (bool).

`waf` - (Optional) Reference to WAF intent configuration object. See [ref](#ref) below for details.

`waf_rule` - (Optional) Reference to WAF Rules configuration object. See [ref](#ref) below for details.

`waf_exclusion_rules` - (Optional) Rules that specify the match conditions and the corresponding WAF_RULE_IDs which should be excluded from WAF evaluation. See [Waf Exclusion Rules ](#waf-exclusion-rules) below for details.

### Active Service Policies

Apply the specified list of service policies and bypass the namespace service policy set.

`policies` - (Required) An ordered list of references to service_policy objects.. See [ref](#ref) below for details.

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

`common_hash_policy` - (Optional) Use Load balancer hash policy for this route (bool).

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

`disable_waf` - (Optional) No WAF configuration for this load balancer (bool).

`waf` - (Optional) Reference to WAF intent configuration object. See [ref](#ref) below for details.

`waf_rule` - (Optional) Reference to WAF Rules configuration object. See [ref](#ref) below for details.

`disable_web_socket_config` - (Optional) Websocket upgrade is disabled (bool).

`web_socket_config` - (Optional) Upgrade to Websocket for this route. See [Web Socket Config ](#web-socket-config) below for details.

### Advertise Custom

Advertise this loadbalancer on specific sites.

`advertise_where` - (Required) Where should this load balancer be available. See [Advertise Where ](#advertise-where) below for details.

### Advertise On Public

Advertise this loadbalancer on public network.

`public_ip` - (Required) Dedicated public ip are allocated by volterra on request. See [ref](#ref) below for details.

### Advertise Where

Where should this load balancer be available.

`site` - (Optional) Advertise on a customer site and a given network. . See [Site ](#site) below for details.

`virtual_network` - (Optional) Advertise on a virtual network. See [Virtual Network ](#virtual-network) below for details.

`virtual_site` - (Optional) Advertise on a customer virtual site and a given network.. See [Virtual Site ](#virtual-site) below for details.

`vk8s_service` - (Optional) Advertise on vK8s Service Network on RE.. See [Vk8s Service ](#vk8s-service) below for details.

`port` - (Optional) TCP port to Listen. (`Int`).

`use_default_port` - (Optional) For HTTP, default is 80. For HTTPS/SNI, default is 443. (bool).

### Always Enable Captcha Challenge

Challenge rules can be used to selectively disable Captcha challenge or enable Javascript challenge for some requests..

### Always Enable Js Challenge

Challenge rules can be used to selectively disable Javascript challenge or enable Captcha challenge for some requests..

### Any Domain

Any Domain..

### Api Definitions

Use API definitions from OpenAPI specification files to specify API operations of an application.

`api_definitions` - (Required) API Definitions using OpenAPI specification files. See [ref](#ref) below for details.

### App Firewall Detection Control

App Firewall detection changes to be applied for this request.

`exclude_signature_contexts` - (Optional) App Firewall signature contexts to be excluded for this request. See [Exclude Signature Contexts ](#exclude-signature-contexts) below for details.

`exclude_violation_contexts` - (Optional) App Firewall violation contexts to be excluded for this request. See [Exclude Violation Contexts ](#exclude-violation-contexts) below for details.

### Asn List

The ASN is obtained by performing a lookup for the source IPv4 Address in a GeoIP DB..

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. (`Int`).

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

`status` - (Optional) HTTP Status code to respond with (`String`).

### Blocked Clients

Rules that specify the clients to be blocked.

`bot_skip_processing` - (Optional) Skip Bot processing for clients matching this rule. (bool).

`skip_processing` - (Optional) Skip both WAF and Bot processing for clients matching this rule. (bool).

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (bool).

`as_number` - (Required) RFC 6793 defined 4-byte AS number (`Int`).

`ip_prefix` - (Required) IPv4 prefix string. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

### Bot Defense

Bot Defense configuration object.

`policy` - (Required) Bot Defense Policy.. See [Policy ](#policy) below for details.

`regional_endpoint` - (Required) x-required (`String`).

`timeout` - (Optional) The timeout for the inference check, in milliseconds. (`Int`).

### Bot Skip Processing

Skip Bot processing for clients matching this rule..

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

Check that the header is not present..

### Check Present

Check that the header is present..

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Common Buffering

Use common buffering configuration.

### Common Hash Policy

Use Load balancer hash policy for this route.

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

Specifies that the default value of "volt-adc" should be used for Server Header.

### Default Js Challenge Parameters

Use default parameters.

### Default Mitigation Settings

Use default parameters.

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

TLS v1.2+ with PFS ciphers with strong crypto algorithms..

### Default Temporary Blocking Parameters

Use default parameters.

### Default Vip

Use the default VIP, system allocated or configured in the virtual network.

### Direct Response Route

A direct response route matches on path and/or HTTP method and responds directly to the matching traffic.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`path` - (Optional) URI path of route. See [Path ](#path) below for details.

`route_direct_response` - (Optional) Send direct response. See [Route Direct Response ](#route-direct-response) below for details.

### Disable Ddos Detection

Disable DDoS Detection.

### Disable Discovery

Disable API discovery.

### Disable Host Rewrite

Host header is not modified.

### Disable Js Insert

Disable JavaScript insertion..

### Disable Learn From Redirect Traffic

Disable learning API patterns from traffic with redirect response codes 3xx.

### Disable Malicious User Detection

Disable malicious user detection.

### Disable Mirroring

Disable Mirroring of request.

### Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Disable Path Normalize

Path normalization is disabled.

### Disable Prefix Rewrite

Do not rewrite the path prefix..

### Disable Spdy

SPDY upgrade is disabled.

### Disable Waf

No WAF configuration for this load balancer.

### Disable Web Socket Config

Websocket upgrade is disabled.

### Do Not Retract Cluster

configuration..

### Domain

Domain matcher..

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Enable Ddos Detection

Enable DDoS Detection.

### Enable Discovery

Enable API discovery.

`disable_learn_from_redirect_traffic` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (bool).

`enable_learn_from_redirect_traffic` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (bool).

### Enable Learn From Redirect Traffic

Enable learning API patterns from traffic with redirect response codes 3xx.

### Enable Malicious User Detection

Enable malicious user detection.

### Enable Path Normalize

Path normalization is enabled.

### Enable Spdy

SPDY upgrade is enabled.

### Enable Strict Sni Host Header Check

Enable strict SNI and Host header check".

### Exclude List

Optional JavaScript insertions exclude list of domain and path matchers..

`any_domain` - (Optional) Any Domain. (bool).

`domain` - (Optional) Domain matcher.. See [Domain ](#domain) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`path` - (Required) URI path matcher.. See [Path ](#path) below for details.

### Exclude Signature Contexts

App Firewall signature contexts to be excluded for this request.

`signature_id` - (Required) App Firewall signature ID (`Int`).

### Exclude Violation Contexts

App Firewall violation contexts to be excluded for this request.

`exclude_violation` - (Required) App Firewall violation type (`String`).

### Flag

Flag the request while not taking any invasive actions..

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

### Http

HTTP Load balancer..

`dns_volterra_managed` - (Optional) This requires the domain to be delegated to Volterra using the Delegated Domain feature. (`Bool`).

### Https

User is responsible for managing DNS to this Load Balancer..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) Redirect HTTP traffic to corresponding HTTPS (`Bool`).

`disable_path_normalize` - (Optional) Path normalization is disabled (bool).

`enable_path_normalize` - (Optional) Path normalization is enabled (bool).

`append_server_name` - (Optional) If Server Header is already present it is not overwritten. It is just passed. (`String`).

`default_header` - (Optional) Specifies that the default value of "volt-adc" should be used for Server Header (bool).

`pass_through` - (Optional) appended. (bool).

`server_name` - (Optional) This will overwrite existing values if any for Server Header (`String`).

`tls_parameters` - (Optional) TLS parameters for downstream connections.. See [Tls Parameters ](#tls-parameters) below for details.

### Https Auto Cert

DNS records will be managed by Volterra..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) Redirect HTTP traffic to corresponding HTTPS (`Bool`).

`no_mtls` - (Optional) mTLS with clients is not enabled (bool).

`use_mtls` - (Optional) mTLS with clients is enabled. See [Use Mtls ](#use-mtls) below for details.

`disable_path_normalize` - (Optional) Path normalization is disabled (bool).

`enable_path_normalize` - (Optional) Path normalization is enabled (bool).

`append_server_name` - (Optional) If Server Header is already present it is not overwritten. It is just passed. (`String`).

`default_header` - (Optional) Specifies that the default value of "volt-adc" should be used for Server Header (bool).

`pass_through` - (Optional) appended. (bool).

`server_name` - (Optional) This will overwrite existing values if any for Server Header (`String`).

`tls_config` - (Optional) Configuration for TLS parameters such as min/max TLS version and ciphers. See [Tls Config ](#tls-config) below for details.

### Ip Allowed List

List of IP(s) for which rate limiting will be disabled..

`prefixes` - (Required) List of IPv4 prefixes that represent an endpoint (`String`).

### Ip Prefix List

IPv4 prefix string..

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Required) List of IPv4 prefix strings. (`String`).

### Item

Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher..

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

Low Security chooses TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Medium Security

Medium Security chooses TLS v1.0+ with only PFS ciphers and medium strength crypto algorithms..

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

`flag` - (Optional) Flag the request while not taking any invasive actions. (bool).

`none` - (Optional) No mitigation actions. (bool).

`redirect` - (Optional) Redirect bot request to a custom URI.. See [Redirect ](#redirect) below for details.

### Mobile

Mobile application traffic type..

### More Option

More options like header manipulation, compression etc..

`buffer_policy` - (Optional) specify the maximum buffer size and buffer interval with this config.. See [Buffer Policy ](#buffer-policy) below for details.

`compression_params` - (Optional) Only GZIP compression is supported. See [Compression Params ](#compression-params) below for details.

`custom_errors` - (Optional) matches for a request. (`String`).

`disable_default_error_pages` - (Optional) Disable the use of default Volterra error pages. (`Bool`).

`idle_timeout` - (Optional) received, otherwise the stream is reset. (`Int`).

`javascript_info` - (Optional) Custom JavaScript Configuration. Custom JavaScript code can be executed at various stages of request processing.. See [Javascript Info ](#javascript-info) below for details.

`jwt` - (Optional) audiences and issuer. See [ref](#ref) below for details.

`max_request_header_size` - (Optional) such loadbalancers is used for all the loadbalancers in question. (`Int`).

`disable_path_normalize` - (Optional) Path normalization is disabled (bool).

`enable_path_normalize` - (Optional) Path normalization is enabled (bool).

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Request Headers To Add ](#request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Response Headers To Add ](#response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

`additional_domains` - (Optional) Wildcard names are supported in the suffix or prefix form. See [Additional Domains ](#additional-domains) below for details.

`enable_strict_sni_host_header_check` - (Optional) Enable strict SNI and Host header check" (bool).

### No Challenge

Challenge rules can be used to selectively enable Javascript or Captcha challenge for some requests..

### No Crl

Client certificate revocation status is not verified.

### No Ip Allowed List

There is no ip allowed list for rate limiting, all clients go through rate limiting..

### No Mtls

mTLS with clients is not enabled.

### No Policies

Do not apply additional rate limiter policies..

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

appended..

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

to the action configured in the rule. If there's no match, the rate limiting configuration for the HTTP loadbalancer is honored..

`policies` - (Required) Ordered list of rate limiter policies.. See [ref](#ref) below for details.

### Policy

Bot Defense Policy..

`disable_js_insert` - (Optional) Disable JavaScript insertion. (bool).

`js_insert_all_pages` - (Optional) Insert Bot Defense JavaScript in all pages.. See [Js Insert All Pages ](#js-insert-all-pages) below for details.

`js_insert_all_pages_except` - (Optional) Insert Bot Defense JavaScript in all pages with the exceptions.. See [Js Insert All Pages Except ](#js-insert-all-pages-except) below for details.

`js_insertion_rules` - (Optional) Specify custom JavaScript insertion rules.. See [Js Insertion Rules ](#js-insertion-rules) below for details.

`js_download_path` - (Optional) Customize Bot Defense Client JavaScript path. If not specified, default `/common.js` (`String`).

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

`default_mitigation_settings` - (Optional) Use default parameters (bool).

`malicious_user_mitigation` - (Optional) The settings defined in malicious user mitigation specify what mitigation actions to take for users determined to be at different threat levels.. See [ref](#ref) below for details.

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

### Protected App Endpoints

List of protected application endpoints (max 128 items)..

`mobile` - (Optional) Mobile application traffic type. (bool).

`web` - (Optional) Web application traffic type. (bool).

`web_mobile` - (Optional) Web and mobile Application traffic type.. See [Web Mobile ](#web-mobile) below for details.

`any_domain` - (Optional) Any Domain. (bool).

`domain` - (Optional) Domain matcher.. See [Domain ](#domain) below for details.

`http_methods` - (Required) List of HTTP methods. (`List of Strings`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`mitigation` - (Required) Mitigation action.. See [Mitigation ](#mitigation) below for details.

`path` - (Required) Matching URI path of the route.. See [Path ](#path) below for details.

`protocol` - (Optional) Protocol. (`String`).

### Rate Limit

Rate limiting parameters for this loadbalancer.

`custom_ip_allowed_list` - (Optional) IP Allowed list using existing ip_prefix_set objects.. See [Custom Ip Allowed List ](#custom-ip-allowed-list) below for details.

`ip_allowed_list` - (Optional) List of IP(s) for which rate limiting will be disabled.. See [Ip Allowed List ](#ip-allowed-list) below for details.

`no_ip_allowed_list` - (Optional) There is no ip allowed list for rate limiting, all clients go through rate limiting. (bool).

`no_policies` - (Optional) Do not apply additional rate limiter policies. (bool).

`policies` - (Optional) to the action configured in the rule. If there's no match, the rate limiting configuration for the HTTP loadbalancer is honored.. See [Policies ](#policies) below for details.

`rate_limiter` - (Optional) Requests to the virtual_host are rate limited based on the parameters specified in the rate_limiter.. See [Rate Limiter ](#rate-limiter) below for details.

### Rate Limiter

Requests to the virtual_host are rate limited based on the parameters specified in the rate_limiter..

`burst_multiplier` - (Optional) The maximum burst of requests to accommodate, expressed as a multiple of the rate. (`Int`).

`total_number` - (Required) The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period. (`Int`).

`unit` - (Required) Unit for the period per which the rate limit is applied. (`String`).

### Redirect

Redirect bot request to a custom URI..

`uri` - (Optional) URI location for redirect may be relative or absolute. (`String`).

### Redirect Route

A redirect route matches on path and/or HTTP method and redirects the matching traffic to a different URL.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`path` - (Optional) URI path of route. See [Path ](#path) below for details.

`route_redirect` - (Optional) Send redirect response. See [Route Redirect ](#route-redirect) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Remove All Params

Remove all query parameters.

### Request Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Secret Value ](#secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Response Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Secret Value ](#secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Retain All Params

Retain all query parameters.

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

`path_redirect` - (Optional) swap path part of incoming URL in redirect URL (`String`).

`proto_redirect` - (Optional) When incoming-proto option is specified, swapping of protocol is not done. (`String`).

`all_params` - (Optional) be removed. Default value is false, which means query portion of the URL will NOT be removed (`Bool`).

`remove_all_params` - (Optional) Remove all query parameters (bool).

`retain_all_params` - (Optional) Retain all query parameters (bool).

`strip_query_params` - (Optional) Specifies the list of query params to be removed. Not supported. See [Strip Query Params ](#strip-query-params) below for details.

`response_code` - (Optional) code is MOVED_PERMANENTLY (301). (`Int`).

### Routes

Routes for this loadbalancer.

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

### Simple Route

A simple route matches on path and/or HTTP method and forwards the matching traffic to the associated pools.

`advanced_options` - (Optional) Configure Advanced per route options. See [Advanced Options ](#advanced-options) below for details.

`auto_host_rewrite` - (Optional) Host header will be swapped with hostname of upstream host chosen by the cluster (bool).

`disable_host_rewrite` - (Optional) Host header is not modified (bool).

`host_rewrite` - (Optional) Host header will be swapped with this value (`String`).

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`origin_pools` - (Optional) Origin Pools for this route. See [Origin Pools ](#origin-pools) below for details.

`path` - (Optional) URI path of route. See [Path ](#path) below for details.

### Single Lb App

ML Config applied on this Load Balancer.

`disable_discovery` - (Optional) Disable API discovery (bool).

`enable_discovery` - (Optional) Enable API discovery. See [Enable Discovery ](#enable-discovery) below for details.

`disable_ddos_detection` - (Optional) Disable DDoS Detection (bool).

`enable_ddos_detection` - (Optional) Enable DDoS Detection (bool).

`disable_malicious_user_detection` - (Optional) Disable malicious user detection (bool).

`enable_malicious_user_detection` - (Optional) Enable malicious user detection (bool).

### Site

Advertise on a customer site and a given network. .

`ip` - (Optional) Use given IP address as VIP on the site (`String`).

`network` - (Required) By default VIP chosen as ip address of primary network interface in the network (`String`).

`site` - (Required) Reference to site object. See [ref](#ref) below for details.

### Skip Processing

Skip both WAF and Bot processing for clients matching this rule..

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

Configuration for TLS parameters such as min/max TLS version and ciphers.

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Custom Security ](#custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers with strong crypto algorithms. (bool).

`low_security` - (Optional) Low Security chooses TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (bool).

`medium_security` - (Optional) Medium Security chooses TLS v1.0+ with only PFS ciphers and medium strength crypto algorithms. (bool).

### Tls Fingerprint Matcher

The predicate evaluates to true if the TLS fingerprint matches any of the exact values or classes of known TLS fingerprints..

`classes` - (Optional) A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. (`List of Strings`).

`exact_values` - (Optional) A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against. (`String`).

`excluded_values` - (Optional) or more known TLS fingerprint classes in the enclosing matcher. (`String`).

### Tls Parameters

TLS parameters for downstream connections..

`no_mtls` - (Optional) mTLS with clients is not enabled (bool).

`use_mtls` - (Optional) mTLS with clients is enabled. See [Use Mtls ](#use-mtls) below for details.

`tls_certificates` - (Required) Set of TLS certificates. See [Tls Certificates ](#tls-certificates) below for details.

`tls_config` - (Optional) Configuration for TLS parameters such as min/max TLS version and ciphers. See [Tls Config ](#tls-config) below for details.

### Trusted Clients

WAF or/and Bot processing can be skipped for trusted clients.

`bot_skip_processing` - (Optional) Skip Bot processing for clients matching this rule. (bool).

`skip_processing` - (Optional) Skip both WAF and Bot processing for clients matching this rule. (bool).

`waf_skip_processing` - (Optional) Skip WAF processing for clients matching this rule. (bool).

`as_number` - (Required) RFC 6793 defined 4-byte AS number (`Int`).

`ip_prefix` - (Required) IPv4 prefix string. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

### Use Default Port

For HTTP, default is 80. For HTTPS/SNI, default is 443..

### Use Mtls

mTLS with clients is enabled.

`crl` - (Optional) CRL server information to download the certificate revocation list. See [ref](#ref) below for details.

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

Rules that specify the match conditions and the corresponding WAF_RULE_IDs which should be excluded from WAF evaluation.

`app_firewall_detection_control` - (Optional) App Firewall detection changes to be applied for this request. See [App Firewall Detection Control ](#app-firewall-detection-control) below for details.

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

Web application traffic type..

### Web Mobile

Web and mobile Application traffic type..

`header` - (Required) Header that is used by mobile traffic.. See [Header ](#header) below for details.

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
