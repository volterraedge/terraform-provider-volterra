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

  // One of the arguments from this list "js_challenge captcha_challenge policy_based_challenge no_challenge" must be set
  no_challenge = true

  domains = ["www.foo.com"]

  // One of the arguments from this list "cookie_stickiness ring_hash round_robin least_active random source_ip_stickiness" must be set
  round_robin = true

  // One of the arguments from this list "http https_auto_cert https" must be set

  http {
    dns_volterra_managed = true
  }
  // One of the arguments from this list "rate_limit disable_rate_limit" must be set
  disable_rate_limit = true
  // One of the arguments from this list "service_policies_from_namespace no_service_policies active_service_policies" must be set
  no_service_policies = true

  // One of the arguments from this list "waf waf_rule disable_waf" must be set

  waf_rule {
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

`blocked_clients` - (Optional) Rules that specify the clients to be blocked. See [Blocked Clients ](#blocked-clients) below for details.

`captcha_challenge` - (Optional) Configure Captcha challenge on this load balancer. See [Captcha Challenge ](#captcha-challenge) below for details.

`js_challenge` - (Optional) Configure Javascript challenge on this load balancer. See [Js Challenge ](#js-challenge) below for details.

`no_challenge` - (Optional) No challenge is enabled for this load balancer (bool).

`policy_based_challenge` - (Optional) Specifies the settings for policy rule based challenge. See [Policy Based Challenge ](#policy-based-challenge) below for details.

`cors_policy` - (Optional) resources from a server at a different origin. See [Cors Policy ](#cors-policy) below for details.

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

`more_option` - (Optional) More options like header manipulation, compression etc.. See [More Option ](#more-option) below for details.

`disable_rate_limit` - (Optional) Rate limiting is not currently enabled for this loadbalancer (bool).

`rate_limit` - (Optional) Rate limiting parameters for this loadbalancer. See [Rate Limit ](#rate-limit) below for details.

`routes` - (Optional) Routes for this loadbalancer. See [Routes ](#routes) below for details.

`active_service_policies` - (Optional) Apply the specified list of service policies and bypass the namespace service policy set. See [Active Service Policies ](#active-service-policies) below for details.

`no_service_policies` - (Optional) Do not apply any service policies i.e. bypass the namespace service policy set (bool).

`service_policies_from_namespace` - (Optional) Apply the service policies configured as part of the namespace service policy set (bool).

`trusted_clients` - (Optional) WAF processing is skipped for trusted clients. See [Trusted Clients ](#trusted-clients) below for details.

`user_identification` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier to be rate limited.. See [ref](#ref) below for details.

`disable_waf` - (Optional) No WAF configuration for this load balancer (bool).

`waf` - (Optional) Reference to WAF intent configuration object. See [ref](#ref) below for details.

`waf_rule` - (Optional) Reference to WAF Rules configuration object. See [ref](#ref) below for details.

`waf_exclusion_rules` - (Optional) Rules that specify the match conditions and the corresponding WAF_RULE_IDs which should be excluded from WAF evaluation. See [Waf Exclusion Rules ](#waf-exclusion-rules) below for details.

### Active Service Policies

Apply the specified list of service policies and bypass the namespace service policy set.

`policies` - (Required) An ordered list of references to service_policy objects.. See [ref](#ref) below for details.

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

`private_network` - (Optional) Advertise on a VoltADN private network. See [Private Network ](#private-network) below for details.

`site` - (Optional) Advertise on a customer site and a given network. . See [Site ](#site) below for details.

`srv6_network` - (Optional) Advertise on a Per site srv6 network. See [Srv6 Network ](#srv6-network) below for details.

`virtual_site` - (Optional) Advertise on a customer virtual site and a given network.. See [Virtual Site ](#virtual-site) below for details.

`vk8s_service` - (Optional) Advertise on vK8s Service Network on RE.. See [Vk8s Service ](#vk8s-service) below for details.

`port` - (Optional) TCP port to Listen. (`Int`).

`use_default_port` - (Optional) For HTTP, default is 80. For HTTPS/SNI, default is 443. (bool).

### Always Enable Captcha Challenge

Challenge rules can be used to selectively disable Captcha challenge or enable Javascript challenge for some requests..

### Always Enable Js Challenge

Challenge rules can be used to selectively disable Javascript challenge or enable Captcha challenge for some requests..

### Any Asn

any_asn.

### Any Client

any_client.

### Any Domain

Apply this waf exclusion rule for any domain.

### Any Ip

any_ip.

### Arg Matchers

arg_matchers.

`invert_matcher` - (Optional) Invert Match of the expression defined (`Bool`).

`check_not_present` - (Optional) Check that the argument is not present. (bool).

`check_present` - (Optional) Check that the argument is present. (bool).

`item` - (Optional) Criteria for matching the values for the Arg. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the arg is present or absent. (`Bool`).

`name` - (Required) A case-sensitive JSON path in the HTTP request body. (`String`).

### Asn List

asn_list.

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. (`Int`).

### Asn Matcher

asn_matcher.

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

### Blocked Clients

Rules that specify the clients to be blocked.

`as_number` - (Required) RFC 6793 defined 4-byte AS number (`Int`).

`ip_prefix` - (Required) IPv4 prefix string. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

### Body Matcher

body_matcher.

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

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

Check that the argument is not present..

### Check Present

Check that the argument is present..

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted .

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Client Name Matcher

client_name_matcher.

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Client Selector

client_selector.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

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

### Cookie Matchers

cookie_matchers.

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

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

### Default Security

TLS v1.2+ with PFS ciphers with strong crypto algorithms..

### Default Temporary Blocking Parameters

Use default parameters.

### Default Vip

Use the default VIP, system allocated or configured in the VoltADN Private Network.

### Direct Response Route

A direct response route matches on path and/or HTTP method and responds directly to the matching traffic.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`path` - (Optional) URI path of route. See [Path ](#path) below for details.

`route_direct_response` - (Optional) Send direct response. See [Route Direct Response ](#route-direct-response) below for details.

### Disable Challenge

Disable the challenge type selected in PolicyBasedChallenge.

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

### Do Not Retract Cluster

configuration..

### Domain Matcher

domain_matcher.

`exact_values` - (Optional) A list of exact values to match the input against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

### Enable Captcha Challenge

Enable captcha challenge.

### Enable Javascript Challenge

Enable javascript challenge.

### Enable Spdy

SPDY upgrade is enabled.

### Hash Policy

route the request.

`cookie` - (Optional) Hash based on cookie. See [Cookie ](#cookie) below for details.

`header_name` - (Optional) The name or key of the request header that will be used to obtain the hash key (`String`).

`source_ip` - (Optional) Hash based on source IP address (`Bool`).

`terminal` - (Optional) Specify if its a terminal policy (`Bool`).

### Headers

headers.

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`check_not_present` - (Optional) Check that the header is not present. (bool).

`check_present` - (Optional) Check that the header is present. (bool).

`item` - (Optional) Criteria for matching the values for the header. The match is successful if any of the values in the input satisfies the criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the header is present or absent. (`Bool`).

`name` - (Required) A case-insensitive HTTP header name. (`String`).

### Http

HTTP Load balancer..

`dns_volterra_managed` - (Optional) This requires the domain to be delegated to Volterra using the Delegated Domain feature. (`Bool`).

### Http Method

http_method.

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`methods` - (Optional) x-example: "['GET', 'POST', 'DELETE']" (`List of Strings`).

### Https

User is responsible for managing DNS to this Load Balancer..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) Redirect HTTP traffic to corresponding HTTPS (`Bool`).

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

`append_server_name` - (Optional) If Server Header is already present it is not overwritten. It is just passed. (`String`).

`default_header` - (Optional) Specifies that the default value of "volt-adc" should be used for Server Header (bool).

`pass_through` - (Optional) appended. (bool).

`server_name` - (Optional) This will overwrite existing values if any for Server Header (`String`).

`tls_config` - (Optional) Configuration for TLS parameters such as min/max TLS version and ciphers. See [Tls Config ](#tls-config) below for details.

### Ip Allowed List

List of IP(s) for which rate limiting will be disabled..

`prefixes` - (Required) List of IPv4 prefixes that represent an endpoint (`String`).

### Ip Matcher

ip_matcher.

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_sets` - (Required) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Ip Prefix List

ip_prefix_list.

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Required) List of IPv4 prefix strings. (`String`).

### Item

Criteria for matching the values for the Arg. The match is successful if any of the values in the input satisfies the criteria in the matcher..

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

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Request Headers To Add ](#request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Response Headers To Add ](#response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

### No Challenge

Challenge rules can be used to selectively enable Javascript or Captcha challenge for some requests..

### No Ip Allowed List

There is no ip allowed list for rate limiting, all clients go through rate limiting..

### No Mtls

mTLS with clients is not enabled.

### No Policies

Do not apply additional rate limiter policies..

### Origin Pools

Origin Pools for this route.

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

`cluster` - (Required) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Required) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

### Pass Through

appended..

### Path

path.

`exact_values` - (Optional) A list of exact path values to match the input HTTP path against. (`String`).

`prefix_values` - (Optional) A list of path prefix values to match the input HTTP path against. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input HTTP path against. (`String`).

`transformers` - (Optional) An ordered list of transformers (starting from index 0) to be applied to the path before matching. (`List of Strings`).

### Percent

Percentage of requests to be mirrored.

`denominator` - (Required) Samples per denominator. numerator part per 100 or 10000 ro 1000000 (`String`).

`numerator` - (Required) sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000 (`Int`).

### Policies

to the action configured in the rule. If there's no match, the rate limiting configuration for the HTTP loadbalancer is honored..

`policies` - (Required) Ordered list of rate limiter policies.. See [ref](#ref) below for details.

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

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Private Network

Advertise on a VoltADN private network.

`private_network` - (Required) Select VoltADN private network reference. See [ref](#ref) below for details.

`default_vip` - (Optional) Use the default VIP, system allocated or configured in the VoltADN Private Network (bool).

`specific_vip` - (Optional) Use given IP address as VIP on VoltADN private Network (`String`).

### Query Params

query_params.

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`key` - (Required) A case-sensitive HTTP query parameter name. (`String`).

`check_not_present` - (Optional) Check that the query parameter is not present. (bool).

`check_present` - (Optional) Check that the query parameter is present. (bool).

`item` - (Optional) criteria in the matcher.. See [Item ](#item) below for details.

`presence` - (Optional) Check if the query parameter is present or absent. (`Bool`).

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

### Retract Cluster

for route.

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

`host_redirect` - (Optional) swap host part of incoming URL in redirect URL (`String`).

`path_redirect` - (Optional) swap path part of incoming URL in redirect URL (`String`).

`proto_redirect` - (Optional)swap proto part of incoming URL in redirect URL (`String`).

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

these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`spec` - (Required) Specification for the rule including match predicates and actions.. See [Spec ](#spec) below for details.

### Simple Route

A simple route matches on path and/or HTTP method and forwards the matching traffic to the associated pools.

`advanced_options` - (Optional) Configure Advanced per route options. See [Advanced Options ](#advanced-options) below for details.

`auto_host_rewrite` - (Optional) Host header will be swapped with hostname of upstream host chosen by the cluster (bool).

`disable_host_rewrite` - (Optional) Host header is not modified (bool).

`host_rewrite` - (Optional) Host header will be swapped with this value (`String`).

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`origin_pools` - (Optional) Origin Pools for this route. See [Origin Pools ](#origin-pools) below for details.

`path` - (Optional) URI path of route. See [Path ](#path) below for details.

### Site

Advertise on a customer site and a given network. .

`ip` - (Optional) Use given IP address as VIP on the site (`String`).

`network` - (Required) By default VIP chosen as ip address of primary network interface in the network (`String`).

`site` - (Required) Reference to site object. See [ref](#ref) below for details.

### Spec

Specification for the rule including match predicates and actions..

`arg_matchers` - (Optional)arg_matchers. See [Arg Matchers ](#arg-matchers) below for details.

`any_asn` - (Optional)any_asn (bool).

`asn_list` - (Optional)asn_list. See [Asn List ](#asn-list) below for details.

`asn_matcher` - (Optional)asn_matcher. See [Asn Matcher ](#asn-matcher) below for details.

`body_matcher` - (Optional)body_matcher. See [Body Matcher ](#body-matcher) below for details.

`disable_challenge` - (Optional) Disable the challenge type selected in PolicyBasedChallenge (bool).

`enable_captcha_challenge` - (Optional) Enable captcha challenge (bool).

`enable_javascript_challenge` - (Optional) Enable javascript challenge (bool).

`any_client` - (Optional)any_client (bool).

`client_name` - (Optional)client_name (`String`).

`client_name_matcher` - (Optional)client_name_matcher. See [Client Name Matcher ](#client-name-matcher) below for details.

`client_selector` - (Optional)client_selector. See [Client Selector ](#client-selector) below for details.

`cookie_matchers` - (Optional)cookie_matchers. See [Cookie Matchers ](#cookie-matchers) below for details.

`domain_matcher` - (Optional)domain_matcher. See [Domain Matcher ](#domain-matcher) below for details.

`expiration_timestamp` - (Optional)expiration_timestamp (`String`).

`headers` - (Optional)headers. See [Headers ](#headers) below for details.

`http_method` - (Optional)http_method. See [Http Method ](#http-method) below for details.

`any_ip` - (Optional)any_ip (bool).

`ip_matcher` - (Optional)ip_matcher. See [Ip Matcher ](#ip-matcher) below for details.

`ip_prefix_list` - (Optional)ip_prefix_list. See [Ip Prefix List ](#ip-prefix-list) below for details.

`path` - (Optional)path. See [Path ](#path) below for details.

`query_params` - (Optional)query_params. See [Query Params ](#query-params) below for details.

`tls_fingerprint_matcher` - (Optional)tls_fingerprint_matcher. See [Tls Fingerprint Matcher ](#tls-fingerprint-matcher) below for details.

### Specific Hash Policy

Configure hash policy specific for this route.

`hash_policy` - (Required) route the request. See [Hash Policy ](#hash-policy) below for details.

### Srv6 Network

Advertise on a Per site srv6 network.

`private_network` - (Required) Select per site srv6 network. See [ref](#ref) below for details.

`default_vip` - (Optional) Use the default VIP, system allocated or configured in the VoltADN Private Network (bool).

`specific_vip` - (Optional) Use given IP address as VIP on VoltADN private Network (`String`).

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

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Tls Config

Configuration for TLS parameters such as min/max TLS version and ciphers.

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Custom Security ](#custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers with strong crypto algorithms. (bool).

`low_security` - (Optional) Low Security chooses TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (bool).

`medium_security` - (Optional) Medium Security chooses TLS v1.0+ with only PFS ciphers and medium strength crypto algorithms. (bool).

### Tls Fingerprint Matcher

tls_fingerprint_matcher.

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

WAF processing is skipped for trusted clients.

`as_number` - (Required) RFC 6793 defined 4-byte AS number (`Int`).

`ip_prefix` - (Required) IPv4 prefix string. (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

### Use Default Port

For HTTP, default is 80. For HTTPS/SNI, default is 443..

### Use Mtls

mTLS with clients is enabled.

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

Rules that specify the match conditions and the corresponding WAF_RULE_IDs which should be excluded from WAF evaluation.

`any_domain` - (Optional) Apply this waf exclusion rule for any domain (bool).

`domain_regex` - (Optional) Domain to be matched (`String`).

`exclude_rule_ids` - (Required) WAF Rules to be excluded when match conditions are met (`List of Strings`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`methods` - (Optional) methods to be matched (`List of Strings`).

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
