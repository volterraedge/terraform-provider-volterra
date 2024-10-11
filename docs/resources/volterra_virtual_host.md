---

page_title: "Volterra: virtual_host"
description: "The virtual_host allows CRUD of Virtual Host resource on Volterra SaaS"

---

Resource volterra_virtual_host
==============================

The Virtual Host allows CRUD of Virtual Host resource on Volterra SaaS

~> **Note:** Please refer to [Virtual Host API docs](https://docs.cloud.f5.com/docs-v2/api/virtual-host) to learn more

Example Usage
-------------

```hcl
resource "volterra_virtual_host" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "captcha_challenge js_challenge no_challenge" must be set

  no_challenge = true
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

`advertise_policies` - (Optional) If advertise policy is not specified then no VIP is assigned for this virtual host.. See [ref](#ref) below for details.

`api_spec` - (Optional) OpenAPI specification settings. See [Api Spec ](#api-spec) below for details.(Deprecated)

###### One of the arguments from this list "authentication, no_authentication" can be set

`authentication` - (Optional) Configure authentication details. See [Authentication Choice Authentication ](#authentication-choice-authentication) below for details.

`no_authentication` - (Optional) Disable Authentication (`Bool`).

`buffer_policy` - (Optional) specify the maximum buffer size and buffer interval with this config.. See [Buffer Policy ](#buffer-policy) below for details.

###### One of the arguments from this list "captcha_challenge, js_challenge, no_challenge" must be set

`captcha_challenge` - (Optional) Configure Captcha challenge on Virtual Host. See [Challenge Type Captcha Challenge ](#challenge-type-captcha-challenge) below for details.

`js_challenge` - (Optional) Configure Javascript challenge on Virtual Host. See [Challenge Type Js Challenge ](#challenge-type-js-challenge) below for details.

`no_challenge` - (Optional) No challenge is enabled for this virtual host (`Bool`).

`compression_params` - (Optional) Only GZIP compression is supported. See [Compression Params ](#compression-params) below for details.

`connection_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 2 minutes. (`Int`).

`cookies_to_modify` - (Optional) List of cookies to be modified from the HTTP response being sent towards downstream.. See [Cookies To Modify ](#cookies-to-modify) below for details.(Deprecated)

`cors_policy` - (Optional) resources from a server at a different origin. See [Cors Policy ](#cors-policy) below for details.

`csrf_policy` - (Optional) CSRF is a mechanism that checks if request received at the server is from legitimate user.. See [Csrf Policy ](#csrf-policy) below for details.

`custom_errors` - (Optional) these pages are not editable. User has an option to disable the use of default F5XC error pages (`String`).

###### One of the arguments from this list "default_loadbalancer, non_default_loadbalancer" can be set

`default_loadbalancer` - (Optional) x-displayName: "Yes" (`Bool`).

`non_default_loadbalancer` - (Optional) x-displayName: "No" (`Bool`).

`disable_default_error_pages` - (Optional) An option to specify whether to disable using default F5XC error pages (`Bool`).

`disable_dns_resolve` - (Optional) for domains configured. This configuration is suitable for HTTP CONNECT proxy. (`Bool`).

`dns_proxy_configuration` - (Optional) Advanced DNS Proxy Configurations like DDoS, Cache are mapped to DNSProxyConfiguration for internal use. See [Dns Proxy Configuration ](#dns-proxy-configuration) below for details.(Deprecated)

`domain_cert_map` - (Optional) which contains repeated Certificate refs). See [Domain Cert Map ](#domain-cert-map) below for details.(Deprecated)

`domains` - (Optional) Domains also indicate the list of names for which DNS resolution will be done by VER (`List of String`).

`dynamic_reverse_proxy` - (Optional) request. The DNS response is cached for 60s by default.. See [Dynamic Reverse Proxy ](#dynamic-reverse-proxy) below for details.

`header_transformation_type` - (Optional) Header transformation options for response headers to the client. See [Header Transformation Type ](#header-transformation-type) below for details.(Deprecated)

`http_protocol_options` - (Optional) HTTP protocol configuration options for downstream connections.. See [Http Protocol Options ](#http-protocol-options) below for details.

`idle_timeout` - (Optional) The default if not specified is 1 minute. (`Int`).

`masking_config` - (Optional) Masking configuration settings. See [Masking Config ](#masking-config) below for details.(Deprecated)

`max_direct_response_body_size` - (Optional) on any of the virtual hosts (`Int`).(Deprecated)

`max_request_header_size` - (Optional) on any of the virtual hosts (`Int`).

###### One of the arguments from this list "disable_path_normalize, enable_path_normalize" can be set

`disable_path_normalize` - (Optional) Path normalization is disabled (`Bool`).

`enable_path_normalize` - (Optional) Path normalization is enabled (`Bool`).

`proxy` - (Optional) Indicates whether the type of proxy is UDP/Secret Management Access (`String`).

`rate_limiter_allowed_prefixes` - (Optional) Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting.. See [ref](#ref) below for details.

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Request Headers To Add ](#request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`List of String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Response Headers To Add ](#response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`List of String`).

`retry_policy` - (Optional) Indicates that the virtual_host has a retry policy.. See [Retry Policy ](#retry-policy) below for details.

`routes` - (Optional) VirtualHosts cannot have DirectResponse or Redirect as actions.. See [ref](#ref) below for details.

`sensitive_data_policy` - (Optional) References to sensitive_data_policy objects.. See [ref](#ref) below for details.

###### One of the arguments from this list "append_server_name, default_header, pass_through, server_name" can be set

`append_server_name` - (Optional) If Server Header is already present it is not overwritten. It is just passed. (`String`).

`default_header` - (Optional) Specifies that the default value of "volt-adc" should be used for Server Header (`Bool`).

`pass_through` - (Optional) appended. (`Bool`).

`server_name` - (Optional) This will overwrite existing values if any for Server Header (`String`).

`slow_ddos_mitigation` - (Optional) This configuration helps to mitigate such type of attacks.. See [Slow Ddos Mitigation ](#slow-ddos-mitigation) below for details.

###### One of the arguments from this list "additional_domains, enable_strict_sni_host_header_check" can be set

`additional_domains` - (Optional) Wildcard names are supported in the suffix or prefix form. See [Strict Sni Host Header Check Choice Additional Domains ](#strict-sni-host-header-check-choice-additional-domains) below for details.(Deprecated)

`enable_strict_sni_host_header_check` - (Optional) Enable strict SNI and Host header check" (`Bool`).(Deprecated)

`temporary_user_blocking` - (Optional) Specifies configuration for temporary user blocking resulting from malicious user detection. See [Temporary User Blocking ](#temporary-user-blocking) below for details.(Deprecated)

###### One of the arguments from this list "tls_cert_params, tls_parameters" can be set

`tls_cert_params` - (Optional) in advertise policy. See [Tls Certificates Choice Tls Cert Params ](#tls-certificates-choice-tls-cert-params) below for details.(Deprecated)

`tls_parameters` - (Optional) in advertise policy. See [Tls Certificates Choice Tls Parameters ](#tls-certificates-choice-tls-parameters) below for details.(Deprecated)

`user_identification` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier to be rate limited.. See [ref](#ref) below for details.

`waf_type` - (Optional) waf_type is the App Firewall profile to use.. See [Waf Type ](#waf-type) below for details.

`ztna_proxy_configurations` - (Optional) Advanced Ztna Proxy configuration containing policy and application configuration. See [Ztna Proxy Configurations ](#ztna-proxy-configurations) below for details.(Deprecated)

### Api Spec

OpenAPI specification settings.

`api_definition` - (Required) API definition is set on this vhost for enforcing OpenAPI on requests. See [ref](#ref) below for details.(Deprecated)

###### One of the arguments from this list "disable_open_api_validation, enable_open_api_validation" must be set

`disable_open_api_validation` - (Optional) No OpenApi Validation configuration for this VH (`Bool`).(Deprecated)

`enable_open_api_validation` - (Optional) OpenApi Validation configuration object. See [Open Api Validation Choice Enable Open Api Validation ](#open-api-validation-choice-enable-open-api-validation) below for details.(Deprecated)

### Buffer Policy

specify the maximum buffer size and buffer interval with this config..

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`max_request_bytes` - (Optional) manager will stop buffering and return a RequestEntityTooLarge (413) response. (`Int`).

`max_request_time` - (Optional) request before returning a RequestTimeout (408) response (`Int`).(Deprecated)

### Compression Params

Only GZIP compression is supported.

`content_length` - (Optional) Minimum response length, in bytes, which will trigger compression. The default value is 30. (`Int`).

`content_type` - (Optional) "text/xml" (`String`).

`disable_on_etag_header` - (Optional) weak etags will be preserved and the ones that require strong validation will be removed. (`Bool`).

`remove_accept_encoding_header` - (Optional) so that responses do not get compressed before reaching the filter. (`Bool`).

### Cookies To Modify

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

CSRF is a mechanism that checks if request received at the server is from legitimate user..

###### One of the arguments from this list "all_load_balancer_domains, custom_domain_list, disabled" must be set

`all_load_balancer_domains` - (Optional) Add All load balancer domains to source origin (allow) list. (`Bool`).

`custom_domain_list` - (Optional) Add one or more domains to source origin (allow) list.. See [Allowed Domains Custom Domain List ](#allowed-domains-custom-domain-list) below for details.

`disabled` - (Optional) Allow all source origin domains. (`Bool`).

### Dns Proxy Configuration

Advanced DNS Proxy Configurations like DDoS, Cache are mapped to DNSProxyConfiguration for internal use.

`cache_profile` - (Optional) which caches DNS replies from the origin DNS servers.. See [Dns Proxy Configuration Cache Profile ](#dns-proxy-configuration-cache-profile) below for details.

`ddos_profile` - (Required) to protect the origin DNS servers from external DDoS attacks.. See [Dns Proxy Configuration Ddos Profile ](#dns-proxy-configuration-ddos-profile) below for details.

`irules` - (Optional) Options for attaching iRules to dns proxy. See [ref](#ref) below for details.

`protocol_inspection` - (Optional) Options for enabling and configuring protocol inspection configuration. See [ref](#ref) below for details.

### Domain Cert Map

which contains repeated Certificate refs).

`ecdsa_certificates` - (Optional) the ECDSA certificate for the domain, if any. See [Domain Cert Map Ecdsa Certificates ](#domain-cert-map-ecdsa-certificates) below for details.(Deprecated)

`rsa_certificates` - (Optional) the RSA certificate for the domain, if any. See [Domain Cert Map Rsa Certificates ](#domain-cert-map-rsa-certificates) below for details.(Deprecated)

### Dynamic Reverse Proxy

request. The DNS response is cached for 60s by default..

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2000 (2 seconds) (`Int`).

`resolution_network` - (Optional) VIRTUAL_NETWORK_GLOBAL. It is ignored for all other network types. See [ref](#ref) below for details.

`resolution_network_type` - (Optional) Type of the network to resolve the destination (`String`).

`resolve_endpoint_dynamically` - (Optional) request. The DNS response is cached for 60s by default. (`Bool`).

### Header Transformation Type

Header transformation options for response headers to the client.

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

### Http Protocol Options

HTTP protocol configuration options for downstream connections..

###### One of the arguments from this list "http_protocol_enable_v1_only, http_protocol_enable_v1_v2, http_protocol_enable_v2_only" must be set

`http_protocol_enable_v1_only` - (Optional) Enable HTTP/1.1 for downstream connections. See [Http Protocol Choice Http Protocol Enable V1 Only ](#http-protocol-choice-http-protocol-enable-v1-only) below for details.

`http_protocol_enable_v1_v2` - (Optional) Enable both HTTP/1.1 and HTTP/2 for downstream connections (`Bool`).

`http_protocol_enable_v2_only` - (Optional) Enable HTTP/2 for downstream connections (`Bool`).

### Masking Config

Masking configuration settings.

###### One of the arguments from this list "disable_masking, enable_masking" can be set

`disable_masking` - (Optional) No Masking for this VH (`Bool`).(Deprecated)

`enable_masking` - (Optional) Enable Masking for this VH (`Bool`).(Deprecated)

### Request Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "secret_value, value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Response Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "secret_value, value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Retry Policy

Indicates that the virtual_host has a retry policy..

`back_off` - (Optional) 10 times the base interval. See [Retry Policy Back Off ](#retry-policy-back-off) below for details.

`num_retries` - (Optional) is used between each retry (`Int`).

`per_try_timeout` - (Optional) Specifies a non-zero timeout per retry attempt. In milliseconds (`Int`).

`retriable_status_codes` - (Optional) HTTP status codes that should trigger a retry in addition to those specified by retry_on. (`Int`).

`retry_condition` - (Required) (disconnect/reset/read timeout.) (`String`).

`retry_on` - (Optional) matching one defined in retriable_status_codes field (`String`).(Deprecated)

### Slow Ddos Mitigation

This configuration helps to mitigate such type of attacks..

`request_headers_timeout` - (Optional) provides protection against Slowloris attacks. (`Int`).

###### One of the arguments from this list "disable_request_timeout, request_timeout" must be set

`disable_request_timeout` - (Optional) x-displayName: "No Timeout" (`Bool`).

`request_timeout` - (Optional) x-example: "60000" (`Int`).

### Temporary User Blocking

Specifies configuration for temporary user blocking resulting from malicious user detection.

`custom_page` - (Optional) E.g. "<p> Blocked </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

### Waf Type

waf_type is the App Firewall profile to use..

###### One of the arguments from this list "app_firewall, disable_waf, inherit_waf" can be set

`app_firewall` - (Optional) A direct reference to an Application Firewall configuration object. See [Ref Type App Firewall ](#ref-type-app-firewall) below for details.

`disable_waf` - (Optional) Any Application Firewall configuration will not be enforced (`Bool`).

`inherit_waf` - (Optional) Any Application Firewall configuration that was configured on a higher level will be enforced (`Bool`).

### Ztna Proxy Configurations

Advanced Ztna Proxy configuration containing policy and application configuration.

`ztna_application_config` - (Optional) Advance configuration of Ztna Application Configuration(Connectivity/Session/Message). See [ref](#ref) below for details.

`ztna_policy_config` - (Optional) Advance configuration of Ztna Policy Configuration(Connectivity/Session/Message). See [ref](#ref) below for details.

### Allowed Domains All Load Balancer Domains

Add All load balancer domains to source origin (allow) list..

### Allowed Domains Custom Domain List

Add one or more domains to source origin (allow) list..

`domains` - (Required) Wildcard names are supported in the suffix or prefix form. (`String`).

### Allowed Domains Disabled

Allow all source origin domains..

### Auth Hmac Prim Key

Primary HMAC Key.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Prim Key Blindfold Secret Info Internal ](#prim-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Auth Hmac Sec Key

Secondary HMAC Key.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Sec Key Blindfold Secret Info Internal ](#sec-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Authentication Choice Authentication

Configure authentication details.

`auth_config` - (Required) Reference to Authentication Config Object. See [ref](#ref) below for details.

###### One of the arguments from this list "cookie_params, use_auth_object_config" can be set

`cookie_params` - (Optional) Configure all Cookie params. See [Cookie Params Choice Cookie Params ](#cookie-params-choice-cookie-params) below for details.

`use_auth_object_config` - (Optional) Use the Cookie Params configured in Authentication Object (`Bool`).

###### One of the arguments from this list "redirect_dynamic, redirect_url" must be set

`redirect_dynamic` - (Optional) This URL must match with the redirect URL configured with the OIDC provider (`Bool`).

`redirect_url` - (Optional) must match with the redirect URL configured with the OIDC provider (`String`).

### Cache Profile Choice Disable Cache Profile

x-displayName: "Disable".

### Challenge Type Captcha Challenge

Configure Captcha challenge on Virtual Host.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

### Challenge Type Js Challenge

Configure Javascript challenge on Virtual Host.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

`js_script_delay` - (Optional) Delay introduced by Javascript, in milliseconds. (`Int`).

### Client Certificate Verify Choice Client Certificate Optional

the connection will be accepted..

### Client Certificate Verify Choice Client Certificate Required

certificate..

### Client Certificate Verify Choice No Client Certificate

it will be ignored (not used for verification).

### Common Params Tls Certificates

Set of TLS certificates.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Common Params Validation Params

and list of Subject Alt Names for verification.

`skip_hostname_verification` - (Optional) is not matched to the connecting hostname (`Bool`).

###### One of the arguments from this list "trusted_ca, trusted_ca_url" must be set

`trusted_ca` - (Optional) Root CA Certificate. See [Trusted Ca Choice Trusted Ca ](#trusted-ca-choice-trusted-ca) below for details.

`trusted_ca_url` - (Optional) Inline Root CA Certificate (`String`).

`use_volterra_trusted_ca_url` - (Optional) Use the F5XC default Root CA URL from the global config for hostname verification. (`Bool`).(Deprecated)

`verify_subject_alt_names` - (Optional) the hostname of the peer will be used for matching against SAN/CN of peer's certificate (`String`).

### Cookie Params Choice Cookie Params

Configure all Cookie params.

`cookie_expiry` - (Optional) Default cookie expiry is 3600 seconds (`Int`).

`cookie_refresh_interval` - (Optional) Default refresh interval is 3000 seconds (`Int`).

###### One of the arguments from this list "auth_hmac, kms_key_hmac" can be set

`auth_hmac` - (Optional) HMAC pair provided as primary and secondary key. See [Secret Choice Auth Hmac ](#secret-choice-auth-hmac) below for details.

`kms_key_hmac` - (Optional) HMAC configured using KMS_KEY. See [Secret Choice Kms Key Hmac ](#secret-choice-kms-key-hmac) below for details.

`session_expiry` - (Optional) Default session expiry is 86400 seconds(24 hours). (`Int`).

### Cookie Params Choice Use Auth Object Config

Use the Cookie Params configured in Authentication Object.

### Cookie Tampering Disable Tampering Protection

x-displayName: "Disable".

### Cookie Tampering Enable Tampering Protection

x-displayName: "Enable".

### Ddos Mitigation Choice Disable Ddos Mitigation

x-displayName: "Disable".

### Ddos Mitigation Choice Enable Ddos Mitigation

x-displayName: "Enable".

### Dns Proxy Configuration Cache Profile

which caches DNS replies from the origin DNS servers..

###### One of the arguments from this list "cache_size, disable_cache_profile" must be set

`cache_size` - (Optional) cache size (`Int`).

`disable_cache_profile` - (Optional) x-displayName: "Disable" (`Bool`).

### Dns Proxy Configuration Ddos Profile

to protect the origin DNS servers from external DDoS attacks..

###### One of the arguments from this list "disable_ddos_mitigation, enable_ddos_mitigation" can be set

`disable_ddos_mitigation` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_ddos_mitigation` - (Optional) x-displayName: "Enable" (`Bool`).

### Domain Cert Map Ecdsa Certificates

the ECDSA certificate for the domain, if any.

`kind` - (Optional) then kind will hold the referred object's kind (e.g. "route") (`String`).

`name` - (Optional) then name will hold the referred object's(e.g. route's) name. (`String`).

`namespace` - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (`String`).

`tenant` - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (`String`).

`uid` - (Optional) then uid will hold the referred object's(e.g. route's) uid. (`String`).

### Domain Cert Map Rsa Certificates

the RSA certificate for the domain, if any.

`kind` - (Optional) then kind will hold the referred object's kind (e.g. "route") (`String`).

`name` - (Optional) then name will hold the referred object's(e.g. route's) name. (`String`).

`namespace` - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (`String`).

`tenant` - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (`String`).

`uid` - (Optional) then uid will hold the referred object's(e.g. route's) uid. (`String`).

### Header Transformation Choice Default Header Transformation

Normalize the headers to lower case.

### Header Transformation Choice Legacy Header Transformation

Use old header transformation if configured earlier.

### Header Transformation Choice Preserve Case Header Transformation

Preserves the original case of headers without any modifications..

### Header Transformation Choice Proper Case Header Transformation

For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are”.

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

x-displayName: "Add".

### Httponly Ignore Httponly

x-displayName: "Ignore".

### Masking Choice Disable Masking

No Masking for this VH.

### Masking Choice Enable Masking

Enable Masking for this VH.

### Max Age Ignore Max Age

Ignore max age attribute.

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ocsp Stapling Choice Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Ocsp Stapling Choice Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Open Api Validation Choice Disable Open Api Validation

No OpenApi Validation configuration for this VH.

### Open Api Validation Choice Enable Open Api Validation

OpenApi Validation configuration object.

`allow_only_specified_headers` - (Optional) Set to fail validation on request/response with header that is not specified in the OpenAPI specification (`Bool`).

`allow_only_specified_query_params` - (Optional) Set to fail validation on request with query parameter that is not specified in the OpenAPI specification (`Bool`).

`fail_close` - (Optional) Set to fail validation on internal errors (`Bool`).

`fail_oversized_body_validation` - (Optional) Set to fail validation on request/response with too long body (`Bool`).

### Prim Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Redirect Url Choice Redirect Dynamic

This URL must match with the redirect URL configured with the OIDC provider.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Ref Type App Firewall

A direct reference to an Application Firewall configuration object.

`app_firewall` - (Required) References to an Application Firewall configuration object. See [ref](#ref) below for details.

### Ref Type Disable Waf

Any Application Firewall configuration will not be enforced.

### Ref Type Inherit Waf

Any Application Firewall configuration that was configured on a higher level will be enforced.

### Request Timeout Choice Disable Request Timeout

x-displayName: "No Timeout".

### Retry Policy Back Off

10 times the base interval.

`base_interval` - (Optional) Specifies the base interval between retries in milliseconds (`Int`).

`max_interval` - (Optional) to the base_interval if set. The default is 10 times the base_interval. (`Int`).

### Samesite Ignore Samesite

Ignore Samesite attribute.

### Samesite Samesite Lax

Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests.

### Samesite Samesite None

Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests.

### Samesite Samesite Strict

Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests.

### Sec Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Choice Auth Hmac

HMAC pair provided as primary and secondary key.

`prim_key` - (Required) Primary HMAC Key. See [Auth Hmac Prim Key ](#auth-hmac-prim-key) below for details.

`prim_key_expiry` - (Required) Primary HMAC Key Expiry time (`String`).

`sec_key` - (Required) Secondary HMAC Key. See [Auth Hmac Sec Key ](#auth-hmac-sec-key) below for details.

`sec_key_expiry` - (Required) Secondary HMAC Key Expiry time (`String`).

### Secret Choice Kms Key Hmac

HMAC configured using KMS_KEY.

`auth_hmac_kms` - (Optional) HMAC configured using the KMS_KEY reference. See [ref](#ref) below for details.(Deprecated)

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

### Secure Add Secure

x-displayName: "Add".

### Secure Ignore Secure

x-displayName: "Ignore".

### Strict Sni Host Header Check Choice Additional Domains

Wildcard names are supported in the suffix or prefix form.

`domains` - (Required) Wildcard names are supported in the suffix or prefix form. (`String`).

### Tls Cert Params Validation Params

and list of Subject Alt Names for verification.

`skip_hostname_verification` - (Optional) is not matched to the connecting hostname (`Bool`).

###### One of the arguments from this list "trusted_ca, trusted_ca_url" must be set

`trusted_ca` - (Optional) Root CA Certificate. See [Trusted Ca Choice Trusted Ca ](#trusted-ca-choice-trusted-ca) below for details.

`trusted_ca_url` - (Optional) Inline Root CA Certificate (`String`).

`use_volterra_trusted_ca_url` - (Optional) Use the F5XC default Root CA URL from the global config for hostname verification. (`Bool`).(Deprecated)

`verify_subject_alt_names` - (Optional) the hostname of the peer will be used for matching against SAN/CN of peer's certificate (`String`).

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

in advertise policy.

`certificates` - (Required) Set of certificates. See [ref](#ref) below for details.

`cipher_suites` - (Optional) will be used. (`String`).

###### One of the arguments from this list "client_certificate_optional, client_certificate_required, no_client_certificate" must be set

`client_certificate_optional` - (Optional) the connection will be accepted.. See [Client Certificate Verify Choice Client Certificate Optional ](#client-certificate-verify-choice-client-certificate-optional) below for details.

`client_certificate_required` - (Optional) certificate.. See [Client Certificate Verify Choice Client Certificate Required ](#client-certificate-verify-choice-client-certificate-required) below for details.

`no_client_certificate` - (Optional) it will be ignored (not used for verification). See [Client Certificate Verify Choice No Client Certificate ](#client-certificate-verify-choice-no-client-certificate) below for details.

`crl` - (Optional) Used to ensure that the client presented certificate is not revoked as per the CRL. See [ref](#ref) below for details.(Deprecated)

`maximum_protocol_version` - (Optional) Maximum TLS protocol version. (`String`).

`minimum_protocol_version` - (Optional) Minimum TLS protocol version. (`String`).

`require_client_certificate` - (Optional) certificate. (`Bool`).(Deprecated)

`validation_params` - (Optional) and list of Subject Alt Names for verification. See [Tls Cert Params Validation Params ](#tls-cert-params-validation-params) below for details.

`xfcc_header_elements` - (Optional) If none are defined, the header will not be added. (`List of Strings`).

### Tls Certificates Choice Tls Parameters

in advertise policy.

###### One of the arguments from this list "client_certificate_optional, client_certificate_required, no_client_certificate" must be set

`client_certificate_optional` - (Optional) the connection will be accepted.. See [Client Certificate Verify Choice Client Certificate Optional ](#client-certificate-verify-choice-client-certificate-optional) below for details.

`client_certificate_required` - (Optional) certificate.. See [Client Certificate Verify Choice Client Certificate Required ](#client-certificate-verify-choice-client-certificate-required) below for details.

`no_client_certificate` - (Optional) it will be ignored (not used for verification). See [Client Certificate Verify Choice No Client Certificate ](#client-certificate-verify-choice-no-client-certificate) below for details.

`common_params` - (Optional) Common TLS parameters used in both upstream and downstream connections. See [Tls Parameters Common Params ](#tls-parameters-common-params) below for details.

`crl` - (Optional) Used to ensure that the client presented certificate is not revoked as per the CRL. See [ref](#ref) below for details.(Deprecated)

`require_client_certificate` - (Optional) certificate. (`Bool`).(Deprecated)

`xfcc_header_elements` - (Optional) If none are defined, the header will not be added. (`List of Strings`).

### Tls Parameters Common Params

Common TLS parameters used in both upstream and downstream connections.

`cipher_suites` - (Optional) will be used. (`String`).

`maximum_protocol_version` - (Optional) Maximum TLS protocol version. (`String`).

`minimum_protocol_version` - (Optional) Minimum TLS protocol version. (`String`).

`tls_certificates` - (Optional) Set of TLS certificates. See [Common Params Tls Certificates ](#common-params-tls-certificates) below for details.

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).(Deprecated)

`validation_params` - (Optional) and list of Subject Alt Names for verification. See [Common Params Validation Params ](#common-params-validation-params) below for details.

### Trusted Ca Choice Trusted Ca

Root CA Certificate.

`trusted_ca_list` - (Optional) Reference to Root CA Certificate. See [ref](#ref) below for details.

### Value Choice Secret Value

Secret Value of the HTTP header..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Secret Value Blindfold Secret Info Internal ](#secret-value-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

Attribute Reference
-------------------

-	`id` - This is the id of the configured virtual_host.
