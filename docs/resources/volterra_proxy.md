---

page_title: "Volterra: proxy"

description: "The proxy allows CRUD of Proxy resource on Volterra SaaS"
-----------------------------------------------------------------------

Resource volterra_proxy
=======================

The Proxy allows CRUD of Proxy resource on Volterra SaaS

~> **Note:** Please refer to [Proxy API docs](https://docs.cloud.f5.com/docs-v2/api/views-proxy) to learn more

Example Usage
-------------

```hcl


resource "volterra_proxy" "example" {
	name = "acmecorp-web"
	namespace = "staging"
// One of the arguments from this list "site_local_inside_network site_local_network" must be set

	site_local_network = true
// One of the arguments from this list "dynamic_proxy http_proxy" must be set


http_proxy {
	
// One of the arguments from this list "enable_http enable_https" must be set

	enable_http = true
	

more_option {
	

buffer_policy {
	
			disabled = true
	
			max_request_bytes = "2048"
	
			max_request_time = "30"
	
}
	

compression_params {
	
			content_length = "100"
	
content_type = ["application/json"]
	
			disable_on_etag_header = true
	
			remove_accept_encoding_header = true
	
}
	

cookies_to_modify {
	
// One of the arguments from this list "disable_tampering_protection enable_tampering_protection" must be set

	disable_tampering_protection = true
	

// One of the arguments from this list "add_httponly ignore_httponly" can be set

	ignore_httponly = true
	

// One of the arguments from this list "ignore_max_age max_age_value" can be set

	ignore_max_age = true
	
			name = "value"
	

// One of the arguments from this list "ignore_samesite samesite_lax samesite_none samesite_strict" can be set

	samesite_none = true
	

// One of the arguments from this list "add_secure ignore_secure" can be set

	ignore_secure = true
	
}
	
			custom_errors = {
				"key1" = "value1"
			}
	
			disable_default_error_pages = true
	
// One of the arguments from this list "disable_http10 enable_http10" must be set

	enable_http10 = true
	
			idle_timeout = "2000"
	

javascript_info {
	
			cache_prefix = "value"
	
			custom_script_url = "value"
	
			script_config = "{\"struct\": \"json\"}"
	
}
	
jwt {
	name = "test1"
	namespace = "staging"
	tenant = "acmecorp"
}
		
	
			max_request_header_size = "60"
	
// One of the arguments from this list "max_requests_per_connection no_request_limit_per_connection" must be set

	no_request_limit_per_connection = true
	

// One of the arguments from this list "disable_path_normalize enable_path_normalize" can be set

	enable_path_normalize = true
	

request_cookies_to_add {
	
			name = "value"
	
			overwrite = true
	
// One of the arguments from this list "secret_value value" must be set

	value = "value"

	
}
	
request_cookies_to_remove = ["request_cookies_to_remove"]
	

request_headers_to_add {
	
			append = true
	
			name = "value"
	
// One of the arguments from this list "secret_value value" must be set

	value = "value"

	
}
	
request_headers_to_remove = ["host"]
	

response_cookies_to_add {
	

// One of the arguments from this list "add_domain ignore_domain" can be set

	ignore_domain = true
	

// One of the arguments from this list "add_expiry ignore_expiry" can be set

	ignore_expiry = true
	

// One of the arguments from this list "add_httponly ignore_httponly" can be set

	ignore_httponly = true
	

// One of the arguments from this list "ignore_max_age max_age_value" can be set

	ignore_max_age = true
	
			name = "value"
	
			overwrite = true
	

// One of the arguments from this list "add_partitioned ignore_partitioned" can be set

	ignore_partitioned = true
	

// One of the arguments from this list "add_path ignore_path" can be set

	ignore_path = true
	

// One of the arguments from this list "ignore_samesite samesite_lax samesite_none samesite_strict" can be set

	ignore_samesite = true
	

// One of the arguments from this list "add_secure ignore_secure" can be set

	ignore_secure = true
	

// One of the arguments from this list "ignore_value secret_value value" can be set

	ignore_value = true
	
}
	
response_cookies_to_remove = ["response_cookies_to_remove"]
	

response_headers_to_add {
	
			append = true
	
			name = "value"
	
// One of the arguments from this list "secret_value value" must be set

	value = "value"

	
}
	
response_headers_to_remove = ["host"]
	

// One of the arguments from this list "additional_domains enable_strict_sni_host_header_check" can be set

	enable_strict_sni_host_header_check = true
	
}
	
}
// One of the arguments from this list "active_forward_proxy_policies no_forward_proxy_policy" must be set

	no_forward_proxy_policy = true
// One of the arguments from this list "do_not_advertise site_virtual_sites" must be set


site_virtual_sites {
	

advertise_where {
	
// One of the arguments from this list "site virtual_site" must be set


site {
	
			ip = "8.8.8.8"
	
			ipv6 = "2001::1"
	
			network = "network"
	
site {
	name = "test1"
	namespace = "staging"
	tenant = "acmecorp"
}
		
	
}
	

// One of the arguments from this list "port use_default_port" can be set

	use_default_port = true
	
}
	
}
// One of the arguments from this list "no_interception tls_intercept" must be set

	no_interception = true
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

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2000 (2 seconds) (`Int`).

###### One of the arguments from this list "site_local_inside_network, site_local_network" must be set

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (`Bool`).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (`Bool`).

###### One of the arguments from this list "dynamic_proxy, http_proxy" must be set

`dynamic_proxy` - (Optional) This is dynamic reverse proxy, Upstream connection will be determined based on Host header or SNI header. See [Proxy Choice Dynamic Proxy ](#proxy-choice-dynamic-proxy) below for details.

`http_proxy` - (Optional) This is HTTP connect Proxy, Upstream connection will be determined based on HTTP connect protocol. See [Proxy Choice Http Proxy ](#proxy-choice-http-proxy) below for details.

###### One of the arguments from this list "active_forward_proxy_policies, no_forward_proxy_policy" must be set

`active_forward_proxy_policies` - (Optional) Forward proxy policies active for this proxy. See [Proxy Policy Choice Active Forward Proxy Policies ](#proxy-policy-choice-active-forward-proxy-policies) below for details.

`no_forward_proxy_policy` - (Optional) Proxy Policy is disabled for this proxy. (`Bool`).

###### One of the arguments from this list "do_not_advertise, site_virtual_sites" must be set

`do_not_advertise` - (Optional) Do not Instantiate this Proxy (`Bool`).

`site_virtual_sites` - (Optional) Instantiate this proxy on specific sites and/or virtual sites. See [Site Choice Site Virtual Sites ](#site-choice-site-virtual-sites) below for details.

###### One of the arguments from this list "no_interception, tls_intercept" must be set

`no_interception` - (Optional) TLS interception is not enabled (`Bool`).

`tls_intercept` - (Optional) Specify TLS interception configuration. See [Tls Interception Choice Tls Intercept ](#tls-interception-choice-tls-intercept) below for details.

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

### Choice Site

Advertise on a customer site and a given network..

`ip` - (Optional) Use given IP address as VIP on the site (`String`).

`ipv6` - (Optional) Use given IPv6 address as VIP on the site (`String`).

`network` - (Required) By default VIP chosen as ip address of primary network interface in the network. Accepted values: `SITE_NETWORK_INSIDE_AND_OUTSIDE` (default), `SITE_NETWORK_INSIDE`, `SITE_NETWORK_OUTSIDE`, `SITE_NETWORK_SERVICE`, `SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_IP_FABRIC`. Note: an invalid/unrecognized value may currently fall back to the default `SITE_NETWORK_INSIDE_AND_OUTSIDE` instead of returning a validation error. (`String`).

`site` - (Required) Reference to site object. See [ref](#ref) below for details.

### Choice Virtual Site

Advertise on a customer virtual site and a given network..

`network` - (Required) IP address of primary network interface in the network. Accepted values: `SITE_NETWORK_INSIDE_AND_OUTSIDE` (default), `SITE_NETWORK_INSIDE`, `SITE_NETWORK_OUTSIDE`, `SITE_NETWORK_SERVICE`, `SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_IP_FABRIC`. Note: an invalid/unrecognized value may currently fall back to the default `SITE_NETWORK_INSIDE_AND_OUTSIDE` instead of returning a validation error. (`String`).

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Cookie Tampering Disable Tampering Protection

x-displayName: "Disable".

### Cookie Tampering Enable Tampering Protection

x-displayName: "Enable".

### Crl Choice No Crl

Client certificate revocation status is not verified.

### Custom Certificate Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, none, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`none` - (Optional) No private key will be stored or used. Choose this only when you need a public certificate for verification (e.g., validating SAML/JWT signatures). Not suitable for TLS or any use case that requires presenting a certificate, those will fail without a private key.. See [Secret Info Oneof None ](#secret-info-oneof-none) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Dns Choice Disable Dns Masquerade

DNS queries for proxy domains will not be resolved to proxy VIP..

### Dns Choice Enable Dns Masquerade

DNS queries for proxy domains will be resolved to proxy VIP..

### Domain Choice Ignore Domain

Ignore max age attribute.

### Enable Disable Choice Disable Interception

Disable Interception.

### Enable Disable Choice Enable Interception

Enable Interception.

### Enable Https Tls Params

Downstream TLS Parameters like certificate, private key, etc.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Params Tls Certificates ](#tls-params-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Params Tls Config ](#tls-params-tls-config) below for details.

### Expiry Choice Ignore Expiry

Ignore expiry attribute.

### Http10 Choice Disable Http10

Reject HTTP/1.0 and HTTP/0.9 traffic on this LB..

### Http10 Choice Enable Http10

Allow HTTP/1.0 and HTTP/0.9 traffic on this LB..

### Http Https Choice Enable Http

HTTP connect transaction is in cleartext(unencrypted).

### Http Https Choice Enable Https

HTTP connect transaction is in HTTPS.

`proxy_name` - (Required) Fully qualified domain name(FQDN) of the proxy (`String`).

`tls_params` - (Optional) Downstream TLS Parameters like certificate, private key, etc. See [Enable Https Tls Params ](#enable-https-tls-params) below for details.

### Http Proxy More Option

Advanced More options like header manipulation, compression etc..

`buffer_policy` - (Optional) specify the maximum buffer size and buffer interval with this config.. See [More Option Buffer Policy ](#more-option-buffer-policy) below for details.

`compression_params` - (Optional) Only GZIP compression is supported. See [More Option Compression Params ](#more-option-compression-params) below for details.

`cookies_to_modify` - (Optional) List of cookies to be modified from the HTTP response being sent towards downstream.. See [More Option Cookies To Modify ](#more-option-cookies-to-modify) below for details.(Deprecated)

`custom_errors` - (Optional) matches for a request. (`String`).

`disable_default_error_pages` - (Optional) Disable the use of default F5XC error pages. (`Bool`).

###### One of the arguments from this list "disable_http10, enable_http10" must be set

`disable_http10` - (Optional) Reject HTTP/1.0 and HTTP/0.9 traffic on this LB. (`Bool`).

`enable_http10` - (Optional) Allow HTTP/1.0 and HTTP/0.9 traffic on this LB. (`Bool`).

`idle_timeout` - (Optional) received, otherwise the stream is reset. (`Int`).

`javascript_info` - (Optional) Custom JavaScript Configuration. Custom JavaScript code can be executed at various stages of request processing.. See [More Option Javascript Info ](#more-option-javascript-info) below for details.(Deprecated)

`jwt` - (Optional) audiences and issuer. See [ref](#ref) below for details.(Deprecated)

`max_request_header_size` - (Optional) such load balancers is used for all the load balancers in question. (`Int`).

###### One of the arguments from this list "max_requests_per_connection, no_request_limit_per_connection" must be set

`max_requests_per_connection` - (Optional) Enter a value >=1 to define the request limit per connection. (`Int`).

`no_request_limit_per_connection` - (Optional) When selected, no limit is enforced, and connections can handle unlimited requests. (`Bool`).

###### One of the arguments from this list "disable_path_normalize, enable_path_normalize" can be set

`disable_path_normalize` - (Optional) x-displayName: "Disable" (`Bool`).(Deprecated)

`enable_path_normalize` - (Optional) x-displayName: "Enable" (`Bool`).(Deprecated)

`request_cookies_to_add` - (Optional) Cookies specified at this level are applied after cookies from matched Route are applied. See [More Option Request Cookies To Add ](#more-option-request-cookies-to-add) below for details.

`request_cookies_to_remove` - (Optional) List of keys of Cookies to be removed from the HTTP request being sent towards upstream. (`String`).

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [More Option Request Headers To Add ](#more-option-request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_cookies_to_add` - (Optional) Cookies specified at this level are applied after cookies from matched Route are applied. See [More Option Response Cookies To Add ](#more-option-response-cookies-to-add) below for details.

`response_cookies_to_remove` - (Optional) List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [More Option Response Headers To Add ](#more-option-response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

###### One of the arguments from this list "additional_domains, enable_strict_sni_host_header_check" can be set

`additional_domains` - (Optional) Wildcard names are supported in the suffix or prefix form. See [Strict Sni Host Header Check Choice Additional Domains ](#strict-sni-host-header-check-choice-additional-domains) below for details.(Deprecated)

`enable_strict_sni_host_header_check` - (Optional) Enable strict SNI and Host header check (`Bool`).(Deprecated)

### Httponly Add Httponly

x-displayName: "Add".

### Httponly Ignore Httponly

x-displayName: "Ignore".

### Httponly Choice Add Httponly

x-displayName: "Add".

### Httponly Choice Ignore Httponly

x-displayName: "Ignore".

### Https Proxy More Option

Advanced More options like header manipulation, compression etc..

`buffer_policy` - (Optional) specify the maximum buffer size and buffer interval with this config.. See [More Option Buffer Policy ](#more-option-buffer-policy) below for details.

`compression_params` - (Optional) Only GZIP compression is supported. See [More Option Compression Params ](#more-option-compression-params) below for details.

`cookies_to_modify` - (Optional) List of cookies to be modified from the HTTP response being sent towards downstream.. See [More Option Cookies To Modify ](#more-option-cookies-to-modify) below for details.(Deprecated)

`custom_errors` - (Optional) matches for a request. (`String`).

`disable_default_error_pages` - (Optional) Disable the use of default F5XC error pages. (`Bool`).

###### One of the arguments from this list "disable_http10, enable_http10" must be set

`disable_http10` - (Optional) Reject HTTP/1.0 and HTTP/0.9 traffic on this LB. (`Bool`).

`enable_http10` - (Optional) Allow HTTP/1.0 and HTTP/0.9 traffic on this LB. (`Bool`).

`idle_timeout` - (Optional) received, otherwise the stream is reset. (`Int`).

`javascript_info` - (Optional) Custom JavaScript Configuration. Custom JavaScript code can be executed at various stages of request processing.. See [More Option Javascript Info ](#more-option-javascript-info) below for details.(Deprecated)

`jwt` - (Optional) audiences and issuer. See [ref](#ref) below for details.(Deprecated)

`max_request_header_size` - (Optional) such load balancers is used for all the load balancers in question. (`Int`).

###### One of the arguments from this list "max_requests_per_connection, no_request_limit_per_connection" must be set

`max_requests_per_connection` - (Optional) Enter a value >=1 to define the request limit per connection. (`Int`).

`no_request_limit_per_connection` - (Optional) When selected, no limit is enforced, and connections can handle unlimited requests. (`Bool`).

###### One of the arguments from this list "disable_path_normalize, enable_path_normalize" can be set

`disable_path_normalize` - (Optional) x-displayName: "Disable" (`Bool`).(Deprecated)

`enable_path_normalize` - (Optional) x-displayName: "Enable" (`Bool`).(Deprecated)

`request_cookies_to_add` - (Optional) Cookies specified at this level are applied after cookies from matched Route are applied. See [More Option Request Cookies To Add ](#more-option-request-cookies-to-add) below for details.

`request_cookies_to_remove` - (Optional) List of keys of Cookies to be removed from the HTTP request being sent towards upstream. (`String`).

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [More Option Request Headers To Add ](#more-option-request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_cookies_to_add` - (Optional) Cookies specified at this level are applied after cookies from matched Route are applied. See [More Option Response Cookies To Add ](#more-option-response-cookies-to-add) below for details.

`response_cookies_to_remove` - (Optional) List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [More Option Response Headers To Add ](#more-option-response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

###### One of the arguments from this list "additional_domains, enable_strict_sni_host_header_check" can be set

`additional_domains` - (Optional) Wildcard names are supported in the suffix or prefix form. See [Strict Sni Host Header Check Choice Additional Domains ](#strict-sni-host-header-check-choice-additional-domains) below for details.(Deprecated)

`enable_strict_sni_host_header_check` - (Optional) Enable strict SNI and Host header check (`Bool`).(Deprecated)

### Https Proxy Tls Params

Fake down stream TLS Parameters like certificate, private key, etc for all destinations of proxy.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Params Tls Certificates ](#tls-params-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Params Tls Config ](#tls-params-tls-config) below for details.

### Interception Policy Choice Enable For All Domains

Enable interception for all domains.

### Interception Policy Choice Policy

Policy to enable/disable specific domains, with implicit enable all domains.

`interception_rules` - (Required) List of ordered rules to enable or disable for TLS interception. See [Policy Interception Rules ](#policy-interception-rules) below for details.

### Interception Rules Domain Match

Domain value or regular expression to match.

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Max Age Ignore Max Age

Ignore max age attribute.

### Max Age Choice Ignore Max Age

Ignore max age attribute.

### Max Requests Per Connection Choice No Request Limit Per Connection

When selected, no limit is enforced, and connections can handle unlimited requests..

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

### Partitioned Choice Add Partitioned

x-displayName: "Add".

### Partitioned Choice Ignore Partitioned

x-displayName: "Ignore".

### Path Choice Ignore Path

Ignore path attribute.

### Path Normalize Choice Disable Path Normalize

x-displayName: "Disable".

### Path Normalize Choice Enable Path Normalize

x-displayName: "Enable".

### Policy Interception Rules

List of ordered rules to enable or disable for TLS interception.

`domain_match` - (Required) Domain value or regular expression to match. See [Interception Rules Domain Match ](#interception-rules-domain-match) below for details.

###### One of the arguments from this list "disable_interception, enable_interception" must be set

`disable_interception` - (Optional) Disable Interception (`Bool`).

`enable_interception` - (Optional) Enable Interception (`Bool`).

### Port Choice Use Default Port

For HTTP, default is 80. For HTTPS/SNI-Proxy, default is 443..

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Proxy Choice Dynamic Proxy

This is dynamic reverse proxy, Upstream connection will be determined based on Host header or SNI header.

###### One of the arguments from this list "disable_dns_masquerade, enable_dns_masquerade" must be set

`disable_dns_masquerade` - (Optional) DNS queries for proxy domains will not be resolved to proxy VIP. (`Bool`).

`enable_dns_masquerade` - (Optional) DNS queries for proxy domains will be resolved to proxy VIP. (`Bool`).

`domains` - (Required) (`String`).

###### One of the arguments from this list "http_proxy, https_proxy, sni_proxy" must be set

`http_proxy` - (Optional) Destination is determined based on Host header in HTTP connections. See [Proxy Choice Http Proxy ](#proxy-choice-http-proxy) below for details.

`https_proxy` - (Optional) Connections are selected based on SNI and destination is determined based on Host header in HTTP connections. See [Proxy Choice Https Proxy ](#proxy-choice-https-proxy) below for details.

`sni_proxy` - (Optional) Destination is determined based on SNI in TLS connections. See [Proxy Choice Sni Proxy ](#proxy-choice-sni-proxy) below for details.

### Proxy Choice Http Proxy

Destination is determined based on Host header in HTTP connections.

`more_option` - (Optional) Advanced More options like header manipulation, compression etc.. See [Http Proxy More Option ](#http-proxy-more-option) below for details.

### Proxy Choice Http Proxy

This is HTTP connect Proxy, Upstream connection will be determined based on HTTP connect protocol.

###### One of the arguments from this list "enable_http, enable_https" must be set

`enable_http` - (Optional) HTTP connect transaction is in cleartext(unencrypted) (`Bool`).

`enable_https` - (Optional) HTTP connect transaction is in HTTPS. See [Http Https Choice Enable Https ](#http-https-choice-enable-https) below for details.(Deprecated)

`more_option` - (Optional) Advanced More options like header manipulation, compression etc.. See [Http Proxy More Option ](#http-proxy-more-option) below for details.

### Proxy Choice Https Proxy

Connections are selected based on SNI and destination is determined based on Host header in HTTP connections.

`more_option` - (Optional) Advanced More options like header manipulation, compression etc.. See [Https Proxy More Option ](#https-proxy-more-option) below for details.

`tls_params` - (Optional) Fake down stream TLS Parameters like certificate, private key, etc for all destinations of proxy. See [Https Proxy Tls Params ](#https-proxy-tls-params) below for details.

### Proxy Choice Sni Proxy

Destination is determined based on SNI in TLS connections.

`idle_timeout` - (Optional) The amount of time that a stream can exist without upstream or downstream activity, in milliseconds. (`Int`).

### Proxy Policy Choice Active Forward Proxy Policies

Forward proxy policies active for this proxy.

`forward_proxy_policies` - (Required) Ordered List of Forward Proxy Policies active. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

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

### Secret Info Oneof None

No private key will be stored or used. Choose this only when you need a public certificate for verification (e.g., validating SAML/JWT signatures). Not suitable for TLS or any use case that requires presenting a certificate, those will fail without a private key..

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

### Secure Add Secure

x-displayName: "Add".

### Secure Ignore Secure

x-displayName: "Ignore".

### Secure Choice Add Secure

x-displayName: "Add".

### Secure Choice Ignore Secure

x-displayName: "Ignore".

### Signing Cert Choice Custom Certificate

Certificates for generating intermediate certificate for TLS interception..

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Custom Certificate Private Key ](#custom-certificate-private-key) below for details.

### Signing Cert Choice Volterra Certificate

F5XC certificates for generating intermediate certificate for TLS interception..

### Site Choice Site Virtual Sites

Instantiate this proxy on specific sites and/or virtual sites.

`advertise_where` - (Required) Where should this load balancer be available. See [Site Virtual Sites Advertise Where ](#site-virtual-sites-advertise-where) below for details.

### Site Virtual Sites Advertise Where

Where should this load balancer be available.

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Advertise on a customer site and a given network.. See [Choice Site ](#choice-site) below for details.

`virtual_site` - (Optional) Advertise on a customer virtual site and a given network.. See [Choice Virtual Site ](#choice-virtual-site) below for details.

###### One of the arguments from this list "port, use_default_port" can be set

`port` - (Optional) TCP port to Listen. (`Int`).

`use_default_port` - (Optional) For HTTP, default is 80. For HTTPS/SNI-Proxy, default is 443. (`Bool`).

### Strict Sni Host Header Check Choice Additional Domains

Wildcard names are supported in the suffix or prefix form.

`domains` - (Required) Wildcard names are supported in the suffix or prefix form. (`String`).

### Strict Sni Host Header Check Choice Enable Strict Sni Host Header Check

Enable strict SNI and Host header check.

### Tls Certificates Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, none, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`none` - (Optional) No private key will be stored or used. Choose this only when you need a public certificate for verification (e.g., validating SAML/JWT signatures). Not suitable for TLS or any use case that requires presenting a certificate, those will fail without a private key.. See [Secret Info Oneof None ](#secret-info-oneof-none) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Tls Interception Choice Tls Intercept

Specify TLS interception configuration.

###### One of the arguments from this list "enable_for_all_domains, policy" must be set

`enable_for_all_domains` - (Optional) Enable interception for all domains (`Bool`).

`policy` - (Optional) Policy to enable/disable specific domains, with implicit enable all domains. See [Interception Policy Choice Policy ](#interception-policy-choice-policy) below for details.

###### One of the arguments from this list "custom_certificate, volterra_certificate" must be set

`custom_certificate` - (Optional) Certificates for generating intermediate certificate for TLS interception.. See [Signing Cert Choice Custom Certificate ](#signing-cert-choice-custom-certificate) below for details.

`volterra_certificate` - (Optional) F5XC certificates for generating intermediate certificate for TLS interception. (`Bool`).

###### One of the arguments from this list "trusted_ca_url, volterra_trusted_ca" must be set

`trusted_ca_url` - (Optional) Custom Root CA Certificate for validating upstream server certificate (`String`).

`volterra_trusted_ca` - (Optional) F5XC Root CA Certificate for validating upstream server certificate (`Bool`).

### Tls Params Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Tls Params Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Trusted Ca Choice Volterra Trusted Ca

F5XC Root CA Certificate for validating upstream server certificate.

### Value Choice Ignore Value

Ignore value of cookie.

### Value Choice Secret Value

Secret Value of the Cookie header.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Secret Value Blindfold Secret Info Internal ](#secret-value-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, none, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`none` - (Optional) No private key will be stored or used. Choose this only when you need a public certificate for verification (e.g., validating SAML/JWT signatures). Not suitable for TLS or any use case that requires presenting a certificate, those will fail without a private key.. See [Secret Info Oneof None ](#secret-info-oneof-none) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Xfcc Header Xfcc Disabled

No X-Forwarded-Client-Cert header will be added.

### Xfcc Header Xfcc Options

X-Forwarded-Client-Cert header will be added with the configured fields.

`xfcc_header_elements` - (Required) X-Forwarded-Client-Cert header elements to be added to requests (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured proxy.
