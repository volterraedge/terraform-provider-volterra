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
  advertise_on_public_default_vip = true

  // One of the arguments from this list "captcha_challenge no_challenge js_challenge" must be set

  js_challenge {
    cookie_expiry       = "cookie_expiry"
    custom_page         = "string:///PHA+IFBsZWFzZSBXYWl0IDwvcD4="
    enable_js_challenge = true
    js_script_delay     = "js_script_delay"
  }
  domains = ["www.foo.com"]
  // One of the arguments from this list "least_active random source_ip_stickiness cookie_stickiness ring_hash round_robin" must be set
  source_ip_stickiness = true
  // One of the arguments from this list "http https_auto_cert https" must be set
  http = true
  // One of the arguments from this list "disable_rate_limit rate_limit" must be set
  disable_rate_limit = true
  // One of the arguments from this list "service_policies_from_namespace no_service_policies active_service_policies" must be set
  no_service_policies = true
  // One of the arguments from this list "waf waf_rule disable_waf" must be set
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

`advertise_custom` - (Optional) Advertise this loadbalancer on specific sites. See [Advertise Custom ](#advertise-custom) below for details.

`advertise_on_public` - (Optional) Advertise this loadbalancer on public network. See [Advertise On Public ](#advertise-on-public) below for details.

`advertise_on_public_default_vip` - (Optional) Advertise this loadbalancer on public network with default VIP (bool).

`do_not_advertise` - (Optional) Do not advertise this loadbalancer (bool).

`captcha_challenge` - (Optional) Configure Captcha challenge on this load balancer. See [Captcha Challenge ](#captcha-challenge) below for details.

`js_challenge` - (Optional) Configure Javascript challenge on this load balancer. See [Js Challenge ](#js-challenge) below for details.

`no_challenge` - (Optional) No challenge is enabled for this load balancer (bool).

`cors_policy` - (Optional) resources from a server at a different origin. See [Cors Policy ](#cors-policy) below for details.

`default_route_pools` - (Optional) Origin Pools used when no route is specified (default route). See [Default Route Pools ](#default-route-pools) below for details.

`domains` - (Required) Domains also indicate the list of names for which DNS resolution will be done by VER (`List of String`).

`cookie_stickiness` - (Optional) Request are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server. See [Cookie Stickiness ](#cookie-stickiness) below for details.

`least_active` - (Optional) Request are sent to origin server that has least active requests (bool).

`random` - (Optional) Request are sent to all eligible origin servers in random fashion (bool).

`ring_hash` - (Optional) Request are sent to all eligible origin servers using hash of request based on hash policy. Consistent hashing algorithm, ring hash, is used to select origin server. See [Ring Hash ](#ring-hash) below for details.

`round_robin` - (Optional) Request are sent to all eligible origin servers in round robin fashion (bool).

`source_ip_stickiness` - (Optional) Request are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server (bool).

`http` - (Optional) HTTP Load balancer. User is responsible for managing DNS to this Load Balancer. (bool).

`https` - (Optional) User is responsible for managing DNS to this Load Balancer.. See [Https ](#https) below for details.

`https_auto_cert` - (Optional) DNS records will be managed by Volterra.. See [Https Auto Cert ](#https-auto-cert) below for details.

`malicious_user_mitigation` - (Optional) The settings defined in malicious user mitigation specify what mitigation actions to take for users determined to be at different threat levels.. See [ref](#ref) below for details.

`more_option` - (Optional) More options like header manipulation, compression etc.. See [More Option ](#more-option) below for details.

`disable_rate_limit` - (Optional) Rate limiting is not currently enabled for this loadbalancer (bool).

`rate_limit` - (Optional) Rate limiting parameters for this loadbalancer. See [Rate Limit ](#rate-limit) below for details.

`routes` - (Optional) Routes for this loadbalancer. See [Routes ](#routes) below for details.

`active_service_policies` - (Optional) Apply the specified list of service policies and bypass the namespace service policy set. See [Active Service Policies ](#active-service-policies) below for details.

`no_service_policies` - (Optional) Do not apply any service policies i.e. bypass the namespace service policy set (bool).

`service_policies_from_namespace` - (Optional) Apply the service policies configured as part of the namespace service policy set (bool).

`user_identification` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier to be rate limited.. See [ref](#ref) below for details.

`disable_waf` - (Optional) No WAF configuration for this load balancer (bool).

`waf` - (Optional) Reference to WAF intent configuration object. See [ref](#ref) below for details.

`waf_rule` - (Optional) Reference to WAF Rules configuration object. See [ref](#ref) below for details.

`waf_exclusion_rules` - (Optional) Rules that specify the match conditions and the corresponding WAF_RULE_IDs which should be excluded from WAF evaluation . See [Waf Exclusion Rules ](#waf-exclusion-rules) below for details.

### Active Service Policies

Apply the specified list of service policies and bypass the namespace service policy set.

`policies` - (Required) An ordered list of references to service_policy objects.. See [ref](#ref) below for details.

### Advanced Options

Configure Advanced per route options.

`buffer_policy` - (Optional) Route level buffer configuration overrides any configuration at VirtualHost level.. See [Buffer Policy ](#buffer-policy) below for details.

`common_buffering` - (Optional) Use common buffering configuration (bool).

`cors_policy` - (Optional) resources from a server at a different origin. See [Cors Policy ](#cors-policy) below for details.

`disable_location_add` - (Optional) virtual-host level. This configuration is ignored on CE sites. (`Bool`).

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

`common_hash_policy` - (Optional) Use common hash policy for all routes (bool).

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

`timeout` - (Optional) for infinite timeout (`Int`).

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

`virtual_site` - (Optional) Advertise on a customer virtual site and a given network.. See [Virtual Site ](#virtual-site) below for details.

`vk8s_service` - (Optional) Advertise on vK8s Service Network on RE.. See [Vk8s Service ](#vk8s-service) below for details.

`port` - (Optional) TCP port to Listen. (`Int`).

`use_default_port` - (Optional) For HTTP, default is 80. For HTTPS/SNI, default is 443. (bool).

### Any Domain

Apply this waf exclusion rule for any domain.

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

### Buffer Policy

specify the maximum buffer size and buffer interval with this config..

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`max_request_bytes` - (Optional) manager will stop buffering and return a RequestEntityTooLarge (413) response. (`Int`).

`max_request_time` - (Optional) request before returning a RequestTimeout (408) response (`Int`).

### Captcha Challenge

Configure Captcha challenge on this load balancer.

`cookie_expiry` - (Optional) Default cookie expiry is set as 1 hour (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

`enable_captcha_challenge` - (Optional) Turn this configuration knob to enable Captcha Challenge (`Bool`).

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted .

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Common Buffering

Use common buffering configuration.

### Common Hash Policy

Use common hash policy for all routes.

### Compression Params

Only GZIP compression is supported.

`content_length` - (Optional) Minimum response length, in bytes, which will trigger compression. The default value is 30. (`Int`).

`content_type` - (Optional) "text/xml" (`String`).

`disable_on_etag_header` - (Optional) weak etags will be preserved and the ones that require strong validation will be removed. (`Bool`).

`remove_accept_encoding_header` - (Optional) so that responses do not get compressed before reaching the filter. (`Bool`).

### Cookie

Hash based on cookie.

`name` - (Optional) produced (`String`).

`path` - (Optional) will be set for the cookie (`String`).

`ttl` - (Optional) be a session cookie. TTL value is in milliseconds (`Int`).

### Cookie Stickiness

Request are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server.

`name` - (Optional) produced (`String`).

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

### Custom Ip Allowed List

IP Allowed list using existing ip_prefix_set objects.

`rate_limiter_allowed_prefixes` - (Required) Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting.. See [ref](#ref) below for details.

### Custom Route Object

A custom route uses a route object created outside of this view..

`route_ref` - (Optional) Reference to a custom route object. See [ref](#ref) below for details.

### Custom Security

Custom selection of TLS versions and cipher suites.

`cipher_suites` - (Required) The TLS listener will only support the specified cipher list. (`String`).

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

### Default Retry Policy

Use system default retry policy.

### Default Route Pools

Origin Pools used when no route is specified (default route).

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

`cluster` - (Required) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Required) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

### Default Security

TLS v1.2+ with PFS ciphers with strong crypto algorithms..

### Direct Response Route

A direct response route matches on path and/or HTTP method and responds directly to the matching traffic.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`path` - (Optional) URI path of route. See [Path ](#path) below for details.

`route_direct_response` - (Optional) Send direct response. See [Route Direct Response ](#route-direct-response) below for details.

### Disable Host Rewrite

Host header is not modified.

### Disable Mirroring

Disable Mirroring of request.

### Disable Prefix Rewrite

Do not rewrite the path prefix..

### Disable Spdy

SPDY upgrade is disabled.

### Disable Waf

No WAF configuration for this load balancer.

### Disable Web Socket Config

Websocket upgrade is disabled.

### Enable Spdy

SPDY upgrade is enabled.

### Hash Policy

route the request.

`cookie` - (Optional) Hash based on cookie. See [Cookie ](#cookie) below for details.

`header_name` - (Optional) The name or key of the request header that will be used to obtain the hash key (`String`).

`source_ip` - (Optional) Hash based on source IP address (`Bool`).

`terminal` - (Optional) Specify if its a terminal policy (`Bool`).

### Https

User is responsible for managing DNS to this Load Balancer..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) Redirect HTTP traffic to corresponding HTTPS (`Bool`).

`tls_parameters` - (Optional) TLS parameters for downstream connections. . See [Tls Parameters ](#tls-parameters) below for details.

### Https Auto Cert

DNS records will be managed by Volterra..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) Redirect HTTP traffic to corresponding HTTPS (`Bool`).

### Ip Allowed List

List of IP(s) for which rate limiting will be disabled.

`prefixes` - (Required) List of IPv4 prefixes that represent an endpoint (`String`).

### Javascript Info

Custom JavaScript Configuration. Custom JavaScript code can be executed at various stages of request processing..

`cache_prefix` - (Optional) KeyValue store referred by script. (`String`).

`custom_script_url` - (Optional) URL of JavaScript that gets executed (`String`).

`script_config` - (Optional) Input passed to the script (`String`).

### Js Challenge

Configure Javascript challenge on this load balancer.

`cookie_expiry` - (Optional) Default cookie expiry is set as 1 hour (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

`enable_js_challenge` - (Optional) Turn this configuration knob to enable Javascript Challenge (`Bool`).

`js_script_delay` - (Optional) Default delay is 5 seconds (`Int`).

### Low Security

Low Security chooses TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Medium Security

Medium Security chooses TLS v1.0+ with only PFS ciphers and medium strength crypto algorithms..

### Mirror Policy

useful for logging. For example, *cluster1* becomes *cluster1-shadow*..

`origin_pool` - (Required) referred here must be present.. See [ref](#ref) below for details.

`percent` - (Required) Percentage of requests to be mirrored. See [Percent ](#percent) below for details.

### More Option

More options like header manipulation, compression etc..

`buffer_policy` - (Optional) specify the maximum buffer size and buffer interval with this config.. See [Buffer Policy ](#buffer-policy) below for details.

`compression_params` - (Optional) Only GZIP compression is supported. See [Compression Params ](#compression-params) below for details.

`custom_errors` - (Optional) matches for a request. (`String`).

`idle_timeout` - (Optional) The default if not specified is 1 minute. (`Int`).

`javascript_info` - (Optional) Custom JavaScript Configuration. Custom JavaScript code can be executed at various stages of request processing.. See [Javascript Info ](#javascript-info) below for details.

`jwt` - (Optional) audiences and issuer. See [ref](#ref) below for details.

`max_request_header_size` - (Optional) on any of the virtual hosts (`Int`).

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Request Headers To Add ](#request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Response Headers To Add ](#response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

### No Ip Allowed List

There is no ip allowed list for rate limiting, all clients go thru rate limiting.

### No Mtls

MTLS with clients is not enabled.

### Origin Pools

Origin Pools for this route.

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

`cluster` - (Required) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Required) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

### Path

URI path of route.

`path` - (Optional) path : /logout (`String`).

`prefix` - (Optional) prefix : /register/ (`String`).

`regex` - (Optional) regex : /b[io]t (`String`).

### Percent

Percentage of requests to be mirrored.

`denominator` - (Required) Samples per denominator. numerator part per 100 or 10000 ro 1000000 (`String`).

`numerator` - (Required) sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000 (`Int`).

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Rate Limit

Rate limiting parameters for this loadbalancer.

`custom_ip_allowed_list` - (Optional) IP Allowed list using existing ip_prefix_set objects. See [Custom Ip Allowed List ](#custom-ip-allowed-list) below for details.

`ip_allowed_list` - (Optional) List of IP(s) for which rate limiting will be disabled. See [Ip Allowed List ](#ip-allowed-list) below for details.

`no_ip_allowed_list` - (Optional) There is no ip allowed list for rate limiting, all clients go thru rate limiting (bool).

`rate_limiter` - (Optional) Requests to the virtual_host are rate limited based on the parameters specified in the rate_limiter.. See [Rate Limiter ](#rate-limiter) below for details.

### Rate Limiter

Requests to the virtual_host are rate limited based on the parameters specified in the rate_limiter..

`total_number` - (Required) The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period. (`Int`).

`unit` - (Required) Unit for the period per which the rate limit is applied. (`String`).

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

`value` - (Required) Value of the HTTP header. (`String`).

### Response Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

`value` - (Required) Value of the HTTP header. (`String`).

### Retain All Params

Retain all query parameters.

### Retry Policy

Configure custom retry policy.

`back_off` - (Optional) 10 times the base interval. See [Back Off ](#back-off) below for details.

`num_retries` - (Optional) is used between each retry (`Int`).

`per_try_timeout` - (Optional) Specifies a non-zero timeout per retry attempt. In milliseconds (`Int`).

`retriable_status_codes` - (Optional) HTTP status codes that should trigger a retry in addition to those specified by retry_on. (`Int`).

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

`all_params` - (Optional) be removed. Default value is false, which means query portion of the URL will NOT be removed (`Bool`).

`remove_all_params` - (Optional) Remove all query parameters (bool).

`retain_all_params` - (Optional) Retain all query parameters (bool).

`strip_query_params` - (Optional) Specifies the list of query params to be removed. Not supported. See [Strip Query Params ](#strip-query-params) below for details.

`host_redirect` - (Optional) swap host part of incoming URL in redirect URL (`String`).

`path_redirect` - (Optional) swap path part of incoming URL in redirect URL (`String`).

`proto_redirect` - (Optional)swap proto part of incoming URL in redirect URL (`String`).

`response_code` - (Optional) code is MOVED_PERMANENTLY (301). (`Int`).

### Routes

Routes for this loadbalancer.

`custom_route_object` - (Optional) A custom route uses a route object created outside of this view.. See [Custom Route Object ](#custom-route-object) below for details.

`direct_response_route` - (Optional) A direct response route matches on path and/or HTTP method and responds directly to the matching traffic. See [Direct Response Route ](#direct-response-route) below for details.

`redirect_route` - (Optional) A redirect route matches on path and/or HTTP method and redirects the matching traffic to a different URL. See [Redirect Route ](#redirect-route) below for details.

`simple_route` - (Optional) A simple route matches on path and/or HTTP method and forwards the matching traffic to the associated pools. See [Simple Route ](#simple-route) below for details.

### Simple Route

A simple route matches on path and/or HTTP method and forwards the matching traffic to the associated pools.

`auto_host_rewrite` - (Optional) Host header will be swapped with hostname of upstream host chosen by the cluster (bool).

`disable_host_rewrite` - (Optional) Host header is not modified (bool).

`host_rewrite` - (Optional) Host header will be swapped with this value (`String`).

`advanced_options` - (Optional) Configure Advanced per route options. See [Advanced Options ](#advanced-options) below for details.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`origin_pools` - (Optional) Origin Pools for this route. See [Origin Pools ](#origin-pools) below for details.

`path` - (Optional) URI path of route. See [Path ](#path) below for details.

### Site

Advertise on a customer site and a given network. .

`ip` - (Optional) Use given IP address as VIP on the site (`String`).

`network` - (Required) By default VIP chosen as ip address of primary network interface in the network (`String`).

`site` - (Required) Reference to site object. See [ref](#ref) below for details.

### Specific Hash Policy

Configure hash policy specific for this route.

`hash_policy` - (Required) route the request. See [Hash Policy ](#hash-policy) below for details.

### Strip Query Params

Specifies the list of query params to be removed. Not supported.

`query_params` - (Optional) Query params keys to strip while manipulating the HTTP request (`String`).

### Tls Certificates

Set of TLS certificates.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Tls Config

Configuration for TLS parameters such as min/max TLS version and ciphers.

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Custom Security ](#custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers with strong crypto algorithms. (bool).

`low_security` - (Optional) Low Security chooses TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (bool).

`medium_security` - (Optional) Medium Security chooses TLS v1.0+ with only PFS ciphers and medium strength crypto algorithms. (bool).

### Tls Parameters

TLS parameters for downstream connections. .

`no_mtls` - (Optional) MTLS with clients is not enabled (bool).

`use_mtls` - (Optional) MTLS with clients is enabled. See [Use Mtls ](#use-mtls) below for details.

`tls_certificates` - (Required) Set of TLS certificates. See [Tls Certificates ](#tls-certificates) below for details.

`tls_config` - (Optional) Configuration for TLS parameters such as min/max TLS version and ciphers. See [Tls Config ](#tls-config) below for details.

### Use Default Port

For HTTP, default is 80. For HTTPS/SNI, default is 443..

### Use Mtls

MTLS with clients is enabled.

`trusted_ca_url` - (Required) The URL for a trust store (`String`).

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Virtual Site

Advertise on a customer virtual site and a given network..

`network` - (Required) IP address of primary network interface in the network (`String`).

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Vk8s Service

Advertise on vK8s Service Network on RE..

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Waf Exclusion Rules

Rules that specify the match conditions and the corresponding WAF_RULE_IDs which should be excluded from WAF evaluation .

`description` - (Optional) Description (`String`).

`any_domain` - (Optional) Apply this waf exclusion rule for any domain (bool).

`domain_regex` - (Optional) Domain to be matched (`String`).

`exclude_rule_ids` - (Required) WAF Rules to be excluded when match conditions are met (`List of Strings`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`methods` - (Optional) methods to be matched (`List of Strings`).

`name` - (Required) Exclusion rule name (`String`).

`path_regex` - (Required) path regex to be matched (`String`).

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