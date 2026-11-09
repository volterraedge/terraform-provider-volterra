---

page_title: "Volterra: http_profile"

description: "The http_profile allows CRUD of Http Profile resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_http_profile
==============================

The Http Profile allows CRUD of Http Profile resource on Volterra SaaS

~> **Note:** Please refer to [Http Profile API docs](https://docs.cloud.f5.com/docs-v2/api/http-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_http_profile" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
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

`accept_xff` - (Optional) Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request's XFF (X-forwarded-for) headers, if they exist. (`String`).

`app_service` - (Optional) The application service to which the object belongs. (`String`).

`basic_auth_realm` - (Optional) Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none. (`String`).

`command` - (Optional) x-displayName: "Command" (`String`).

`defaults_from` - (Optional) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. (`String`).

`encrypt_cookie_secret` - (Optional) Specifies a passphrase for the cookie encryption. (`String`).

`encrypt_cookies` - (Optional) Encrypts specified cookies that the BIG-IP system sends to a client system. (`List of String`).

`enforcement` - (Optional) x-displayName: "Enforcement". See [Enforcement ](#enforcement) below for details.

`expiration_micros` - (Optional) x-displayName: "Expiration Micros" (`Int`).

`explicit_proxy` - (Optional) x-displayName: "Explicit Proxy". See [Explicit Proxy ](#explicit-proxy) below for details.

`fallback_host` - (Optional) Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number, or URI path. For example, if all members of the targeted pool are unavailable (that is, the members are disabled, marked as down, or have exceeded their connection limit), the system can redirect the HTTP request to the fallback host, with the HTTP reply Status Code 302 Found. (`String`).

`fallback_status_codes` - (Optional) Specifies one or more three-digit status codes that can be returned by an HTTP server. (`List of String`).

`full_path` - (Optional) x-displayName: "Full Path" (`String`).

`generation` - (Optional) x-displayName: "Generation" (`Int`).

`header_erase` - (Optional) Specifies the header string that you want to erase from an HTTP request. You can also specify none. (`String`).

`header_insert` - (Optional) Specifies a quoted header string that you want to insert into an HTTP request. You can also specify none. The HTTP header being inserted can include a client IP address. Including a client IP address in an HTTP header is useful when a connection goes through a secure network address translation (SNAT) and you need to preserve the original client IP address. When you assign the configured HTTP profile to a virtual server, the system then inserts the header specified by the profile into any HTTP request that the system sends to a pool or pool member. (`String`).

`hsts` - (Optional) x-displayName: "HSTS". See [Hsts ](#hsts) below for details.

`insert_xforwarded_for` - (Optional) When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address. (`String`).

`kind` - (Optional) x-displayName: "Kind" (`String`).

`last_update_micros` - (Optional) x-displayName: "Last Update Micros" (`Int`).

`lws_separator` - (Optional) Specifies the linear white space separator that the system should use between HTTP headers when a header exceeds the maximum width specified by the lws width setting. (`String`).

`lws_width` - (Optional) Specifies the maximum number of columns allowed for a header that is inserted into an HTTP request. (`Int`).

`oneconnect_status_reuse` - (Optional) x-displayName: "OneConnect Status Reuse" (`String`).

`oneconnect_transformations` - (Optional) Enables the system to perform HTTP header transformations for the purpose of keeping server-side connections open. This feature requires configuration of a OneConnect profile. (`String`).(Deprecated)

`proxy_type_choice` - (Optional) Specifies the type of HTTP proxy. (`String`).

`redirect_rewrite_choice` - (Optional) Specifies which of the application HTTP redirects the system rewrites to HTTPS. Use this feature when the application is generating HTTP redirects that send the client to HTTP (a non-secure channel) when you want the client to continue accessing the application using HTTPS (a secure channel). This is a common occurrence when using client-side SSL processing on a BIG-IP system. (`String`).

`request_chunking_choice` - (Optional) Specifies how to handle chunked and unchunked requests. (`String`).

`response_chunking_choice` - (Optional) Specifies how to handle chunked and unchunked responses. (`String`).

`response_headers_permitted` - (Optional) Specifies headers that the BIG-IP system allows in an HTTP response. (`List of String`).

`server_agent_name` - (Optional) Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is "BigIP". If no string is specified, then no Server header will be added to such responses. (`String`).

`sflow` - (Optional) x-displayName: "Sflow". See [Sflow ](#sflow) below for details.

`sub_path` - (Optional) x-displayName: "Sub Path" (`String`).

`user_spec` - (Optional) User specified properties. (`List of String`).

`via_host_name` - (Optional) Specifies the hostname to include into Via header. (`String`).

`via_request_choice` - (Optional) Specifies whether to append, remove, or preserve a Via header in an HTTP request. (`String`).

`via_response_choice` - (Optional) Specifies whether to append, remove, or preserve a Via header in an HTTP response. (`String`).

`xff_alternative_names` - (Optional) Specifies alternative XFF headers instead of the default X-forwarded-for header. (`List of String`).

### Enforcement

x-displayName: "Enforcement".

`allow_ws_header_name` - (Optional) Specifies whether to allow or reject white space(s) between a header name and a colon. (`String`).

`excess_client_headers` - (Optional) Specifies the behavior when too many client headers are received. If enabled, will switch to pass through mode instead of rejecting the connection. (`String`).

`excess_server_headers` - (Optional) Specifies the behavior when too many server headers are received. If enabled, will switch to pass through mode instead of rejecting the connection. (`String`).

`known_methods` - (Optional) Specifies which HTTP methods count as being known. Removing RFC-defined methods from this list will cause the HTTP filter to not recognize them. (`String`).

`max_header_count` - (Optional) Specifies the maximum number of headers allowed in HTTP request/response. The default is 64 headers. (`Int`).

`max_header_size` - (Optional) Specifies the maximum header size. (`Int`).

`max_requests` - (Optional) Specifies the number of requests that the system accepts on a per-connection basis. The default value is 0 (zero), which means the system does not limit the number of requests per connection. (`Int`).

`oversize_client_headers` - (Optional) Specifies the behavior when too-large client headers are received. If enabled, will switch to pass through mode instead of rejecting the connection. (`String`).

`oversize_server_headers` - (Optional) Specifies the behavior when too-large server headers are received. If enabled, will switch to pass through mode instead of rejecting the connection. (`String`).

`pipeline_choice` - (Optional) Enables HTTP/1.1 pipelining. This allows clients to make requests even when prior requests have not received a response. In order for this to succeed, however, destination servers must include support for pipelining. If set to pass-through, pipelined data will cause the BigIP to immediately switch to pass-through mode and disable the HTTP filter. (`String`).

`rfc_compliance` - (Optional) Specifies whether to allow or reject non HTTP rfc compliant traffic. (`String`).

`truncated_redirects` - (Optional) Specifies what happens if a truncated redirect is seen from a server. If enabled, the redirect will be forwarded to the client, otherwise the malformed HTTP will be silently ignored. (`String`).

`unknown_method_choice` - (Optional) x-displayName: "Unknown Method Choice" (`String`).

### Explicit Proxy

x-displayName: "Explicit Proxy".

`bad_request_message` - (Optional) Specifies the error message that will be returned to the browser when a proxy request can't be completed because the request was malformed. (`String`).

`bad_response_message` - (Optional) Specifies the error message that will be returned to the browser when a proxy request can't be completed because the response was malformed. (`String`).

`connect_error_message` - (Optional) Specifies the error message that will be returned to the browser when a proxy request can't be completed because of a failure to establish the outbound connection. (`String`).

`default_connect_handling` - (Optional) Specifies the behavior of the proxy service for CONNECT requests. If set to 'deny', CONNECT requests will only be honored if there is another virtual server listening for the requested outbound connection. If set to 'allow' outbound connections will be made regardless of other virtual servers. (`String`).

`dns_error_message` - (Optional) Specifies the error message that will be returned to the browser when a proxy request can't be completed because of a failure to resolve the hostname in the request. (`String`).

`dns_resolver` - (Optional) Specifies the dns-resolver object that will be used to resolve hostnames in proxy requests. (`String`).

`host_names` - (Optional) Specifies the which host names are to be treated as local. Proxy requests made for those hosts will be treated as regular HTTP requests and will be sent to the configured default pool. (`String`).

`ipv6` - (Optional) Specifies that URIs will attempted to be resolved as IPv6 addresses before trying as IPv4. (`String`).

`route_domain` - (Optional) Specifies the route-domain that will be used for outbound proxy requests. (`String`).

`tunnel_name` - (Optional) Specifies the tunnel that will be used for outbound proxy requests. This enables other virtual servers to receive connections initiated by the proxy service. (`String`).

`tunnel_on_any_request` - (Optional) Specifies that the tunnel will be used for non-CONNECT requests. If set to 'yes', virtual servers listening on a tunnel will be able to receive any requests and 'default-connect-handling' option effect will be extended to all outbound proxy requests. (`String`).

### Hsts

x-displayName: "HSTS".

`include_subdomains` - (Optional) Specifies whether to include the includeSubdomains directive in the HSTS header. The default is enabled. (`String`).

`maximum_age` - (Optional) Specifies the maximum age to assume the connection should remain secure. The default is 16070400 seconds. (`Int`).

`mode` - (Optional) Specifies whether to include the HSTS response header. The default is Disabled (`String`).

`preload` - (Optional) Specifies whether to include the preload directive in the HSTS header. The default is Disabled. (`String`).

### Sflow

x-displayName: "Sflow".

`poll_interval` - (Optional) Specifies the maximum interval in seconds between two pollings. To enable this setting, you must also set the poll-interval-global setting to no. (`Int`).

`poll_interval_global` - (Optional) Specifies whether the global HTTP poll-interval setting overrides the object-level poll-interval setting. (`String`).

`sampling_rate` - (Optional) Specifies the ratio of packets observed to the samples generated. For example, a sampling rate of 2000 specifies that 1 sample will be randomly generated for every 2000 packets observed. To enable this setting, you must also set the sampling-rate-global setting to no. (`Int`).

`sampling_rate_global` - (Optional) Specifies whether the global HTTP sampling-rate setting overrides the object-level sampling-rate setting. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured http_profile.
