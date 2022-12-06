---

page_title: "Volterra: virtual_host"

description: "The virtual_host allows CRUD of Virtual Host resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_virtual_host
==============================

The Virtual Host allows CRUD of Virtual Host resource on Volterra SaaS

~> **Note:** Please refer to [Virtual Host API docs](https://volterra.io/docs/api/virtual-host) to learn more

Example Usage
-------------

```hcl
resource "volterra_virtual_host" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "captcha_challenge no_challenge js_challenge" must be set
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

`authentication` - (Optional) Configure authentication details. See [Authentication ](#authentication) below for details.

`no_authentication` - (Optional) Disable Authentication (bool).

`buffer_policy` - (Optional) specify the maximum buffer size and buffer interval with this config.. See [Buffer Policy ](#buffer-policy) below for details.

`captcha_challenge` - (Optional) Configure Captcha challenge on Virtual Host. See [Captcha Challenge ](#captcha-challenge) below for details.

`js_challenge` - (Optional) Configure Javascript challenge on Virtual Host. See [Js Challenge ](#js-challenge) below for details.

`no_challenge` - (Optional) No challenge is enabled for this virtual host (bool).

`compression_params` - (Optional) Only GZIP compression is supported. See [Compression Params ](#compression-params) below for details.

`cors_policy` - (Optional) resources from a server at a different origin. See [Cors Policy ](#cors-policy) below for details.

`csrf_policy` - (Optional) CSRF is a mechanism that checks if request received at the server is from legitimate user.. See [Csrf Policy ](#csrf-policy) below for details.

`custom_errors` - (Optional) these pages are not editable. User has an option to disable the use of default F5XC error pages (`String`).

`default_loadbalancer` - (Optional) x-displayName: "Yes" (bool).

`non_default_loadbalancer` - (Optional) x-displayName: "No" (bool).

`disable_default_error_pages` - (Optional) An option to specify whether to disable using default F5XC error pages (`Bool`).

`disable_dns_resolve` - (Optional) for domains configured. This configuration is suitable for HTTP CONNECT proxy. (`Bool`).

`dns_proxy_configuration` - (Optional) Advanced DNS Proxy Configurations like DDoS, Cache are mapped to DNSProxyConfiguration for internal use. See [Dns Proxy Configuration ](#dns-proxy-configuration) below for details.

`domains` - (Optional) Domains also indicate the list of names for which DNS resolution will be done by VER (`List of String`).

`dynamic_reverse_proxy` - (Optional) request. The DNS response is cached for 60s by default.. See [Dynamic Reverse Proxy ](#dynamic-reverse-proxy) below for details.

`header_transformation_type` - (Optional) Header transformation options for response headers to the client. See [Header Transformation Type ](#header-transformation-type) below for details.

`idle_timeout` - (Optional) The default if not specified is 1 minute. (`Int`).

`max_request_header_size` - (Optional) on any of the virtual hosts (`Int`).

`disable_path_normalize` - (Optional) Path normalization is disabled (bool).

`enable_path_normalize` - (Optional) Path normalization is enabled (bool).

`proxy` - (Optional) Indicates whether the type of proxy is UDP/Secret Management Access (`String`).

`rate_limiter` - (Optional) Requests to the virtual_host are rate limited based on the parameters specified in the rate_limiter.. See [ref](#ref) below for details.

`rate_limiter_allowed_prefixes` - (Optional) Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting.. See [ref](#ref) below for details.

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Request Headers To Add ](#request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`List of String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Response Headers To Add ](#response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`List of String`).

`retry_policy` - (Optional) Indicates that the virtual_host has a retry policy.. See [Retry Policy ](#retry-policy) below for details.

`routes` - (Optional) VirtualHosts cannot have DirectResponse or Redirect as actions.. See [ref](#ref) below for details.

`append_server_name` - (Optional) If Server Header is already present it is not overwritten. It is just passed. (`String`).

`default_header` - (Optional) Specifies that the default value of "volt-adc" should be used for Server Header (bool).

`pass_through` - (Optional) appended. (bool).

`server_name` - (Optional) This will overwrite existing values if any for Server Header (`String`).

`additional_domains` - (Optional) Wildcard names are supported in the suffix or prefix form. See [Additional Domains ](#additional-domains) below for details.

`enable_strict_sni_host_header_check` - (Optional) Enable strict SNI and Host header check" (bool).

`temporary_user_blocking` - (Optional) Specifies configuration for temporary user blocking resulting from malicious user detection. See [Temporary User Blocking ](#temporary-user-blocking) below for details.

`tls_parameters` - (Optional) in advertise policy. See [Tls Parameters ](#tls-parameters) below for details.

`user_identification` - (Optional) The rules in the user_identification object are evaluated to determine the user identifier to be rate limited.. See [ref](#ref) below for details.

`waf_type` - (Optional) Enable/Disable individual WAF security rules. See [Waf Type ](#waf-type) below for details.

### Additional Domains

Wildcard names are supported in the suffix or prefix form.

`domains` - (Optional) Wildcard names are supported in the suffix or prefix form. (`String`).

### All Load Balancer Domains

Add All load balancer domains to source origin (allow) list..

### App Firewall

A direct reference to an Application Firewall configuration object.

`app_firewall` - (Optional) References to an Application Firewall configuration object. See [ref](#ref) below for details.

### Auth Hmac

HMAC pair provided as primary and secondary key.

`prim_key` - (Required) Primary HMAC Key. See [Prim Key ](#prim-key) below for details.

`prim_key_expiry` - (Required) Primary HMAC Key Expiry time (`String`).

`sec_key` - (Required) Secondary HMAC Key. See [Sec Key ](#sec-key) below for details.

`sec_key_expiry` - (Required) Secondary HMAC Key Expiry time (`String`).

### Authentication

Configure authentication details.

`auth_config` - (Required) Reference to Authentication Config Object. See [ref](#ref) below for details.

`cookie_params` - (Optional) Configure all Cookie params. See [Cookie Params ](#cookie-params) below for details.

`use_auth_object_config` - (Optional) Use the Cookie Params configured in Authentication Object (bool).

`redirect_dynamic` - (Optional) This URL must match with the redirect URL configured with the OIDC provider (bool).

`redirect_url` - (Optional) must match with the redirect URL configured with the OIDC provider (`String`).

### Back Off

10 times the base interval.

`base_interval` - (Optional) Specifies the base interval between retries in milliseconds (`Int`).

`max_interval` - (Optional) to the base_interval if set. The default is 10 times the base_interval. (`Int`).

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

### Buffer Policy

specify the maximum buffer size and buffer interval with this config..

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`max_request_bytes` - (Optional) manager will stop buffering and return a RequestEntityTooLarge (413) response. (`Int`).

`max_request_time` - (Optional) request before returning a RequestTimeout (408) response (`Int`).

### Cache Profile

which caches DNS replies from the origin DNS servers..

`disable_cache_profile` - (Optional) x-displayName: "Disable" (bool).

`enable_cache_profile` - (Optional) x-displayName: "Enable" (bool).

### Captcha Challenge

Configure Captcha challenge on Virtual Host.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Common Params

Common TLS parameters used in both upstream and downstream connections.

`cipher_suites` - (Optional) will be used. (`String`).

`maximum_protocol_version` - (Optional) Maximum TLS protocol version. (`String`).

`minimum_protocol_version` - (Optional) Minimum TLS protocol version. (`String`).

`tls_certificates` - (Optional) Set of TLS certificates. See [Tls Certificates ](#tls-certificates) below for details.

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

`validation_params` - (Optional) and list of Subject Alt Names for verification. See [Validation Params ](#validation-params) below for details.

### Compression Params

Only GZIP compression is supported.

`content_length` - (Optional) Minimum response length, in bytes, which will trigger compression. The default value is 30. (`Int`).

`content_type` - (Optional) "text/xml" (`String`).

`disable_on_etag_header` - (Optional) weak etags will be preserved and the ones that require strong validation will be removed. (`Bool`).

`remove_accept_encoding_header` - (Optional) so that responses do not get compressed before reaching the filter. (`Bool`).

### Cookie Params

Configure all Cookie params.

`cookie_expiry` - (Optional) Default cookie expiry is 3600 seconds (`Int`).

`cookie_refresh_interval` - (Optional) Default refresh interval is 3000 seconds (`Int`).

`auth_hmac` - (Optional) HMAC pair provided as primary and secondary key. See [Auth Hmac ](#auth-hmac) below for details.

`kms_key_hmac` - (Optional) HMAC configured using KMS_KEY. See [Kms Key Hmac ](#kms-key-hmac) below for details.

`session_expiry` - (Optional) Default session expiry is 86400 seconds(24 hours). (`Int`).

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

### Csrf Policy

CSRF is a mechanism that checks if request received at the server is from legitimate user..

`all_load_balancer_domains` - (Optional) Add All load balancer domains to source origin (allow) list. (bool).

`custom_domain_list` - (Optional) Add one or more domains to source origin (allow) list.. See [Custom Domain List ](#custom-domain-list) below for details.

### Custom Domain List

Add one or more domains to source origin (allow) list..

`domains` - (Optional) Wildcard names are supported in the suffix or prefix form. (`String`).

### Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ddos Profile

to protect the origin DNS servers from external DDoS attacks..

`disable_ddos_mitigation` - (Optional) x-displayName: "Disable" (bool).

`enable_ddos_mitigation` - (Optional) x-displayName: "Enable" (bool).

### Default Header Transformation

Normalize the headers to lower case.

### Disable Cache Profile

x-displayName: "Disable".

### Disable Ddos Mitigation

x-displayName: "Disable".

### Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Dns Proxy Configuration

Advanced DNS Proxy Configurations like DDoS, Cache are mapped to DNSProxyConfiguration for internal use.

`cache_profile` - (Optional) which caches DNS replies from the origin DNS servers.. See [Cache Profile ](#cache-profile) below for details.

`ddos_profile` - (Required) to protect the origin DNS servers from external DDoS attacks.. See [Ddos Profile ](#ddos-profile) below for details.

### Dynamic Reverse Proxy

request. The DNS response is cached for 60s by default..

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2000 (2 seconds) (`Int`).

`resolution_network` - (Optional) VIRTUAL_NETWORK_GLOBAL. It is ignored for all other network types. See [ref](#ref) below for details.

`resolution_network_type` - (Optional) Type of the network to resolve the destination (`String`).

`resolve_endpoint_dynamically` - (Optional) request. The DNS response is cached for 60s by default. (`Bool`).

### Enable Cache Profile

x-displayName: "Enable".

### Enable Ddos Mitigation

x-displayName: "Enable".

### Header Transformation Type

Header transformation options for response headers to the client.

`default_header_transformation` - (Optional) Normalize the headers to lower case (bool).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (bool).

### Js Challenge

Configure Javascript challenge on Virtual Host.

`cookie_expiry` - (Optional) An expired cookie causes the loadbalancer to issue a new challenge. (`Int`).

`custom_page` - (Optional) E.g. "<p> Please Wait </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

`js_script_delay` - (Optional) Delay introduced by Javascript, in milliseconds. (`Int`).

### Kms Key Hmac

HMAC configured using KMS_KEY.

`auth_hmac_kms` - (Optional) HMAC configured using the KMS_KEY reference. See [ref](#ref) below for details.

### Prim Key

Primary HMAC Key.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Proper Case Header Transformation

For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are”.

### Redirect Dynamic

This URL must match with the redirect URL configured with the OIDC provider.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

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

### Retry Policy

Indicates that the virtual_host has a retry policy..

`back_off` - (Optional) 10 times the base interval. See [Back Off ](#back-off) below for details.

`num_retries` - (Optional) is used between each retry (`Int`).

`per_try_timeout` - (Optional) Specifies a non-zero timeout per retry attempt. In milliseconds (`Int`).

`retriable_status_codes` - (Optional) HTTP status codes that should trigger a retry in addition to those specified by retry_on. (`Int`).

`retry_condition` - (Optional) matching one defined in retriable_status_codes field (`String`).

`retry_on` - (Optional) matching one defined in retriable_status_codes field (`String`).

### Sec Key

Secondary HMAC Key.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Secret Value

Secret Value of the HTTP header..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Temporary User Blocking

Specifies configuration for temporary user blocking resulting from malicious user detection.

`custom_page` - (Optional) E.g. "<p> Blocked </p>". Base64 encoded string for this html is "PHA+IFBsZWFzZSBXYWl0IDwvcD4=" (`String`).

### Tls Certificates

Set of TLS certificates.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Custom Hash Algorithms ](#custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Disable Ocsp Stapling ](#disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Use System Defaults ](#use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Tls Parameters

in advertise policy.

`common_params` - (Optional) Common TLS parameters used in both upstream and downstream connections. See [Common Params ](#common-params) below for details.

`crl` - (Optional) Used to ensure that the client presented certificate is not revoked as per the CRL. See [ref](#ref) below for details.

`require_client_certificate` - (Optional) certificate. (`Bool`).

### Use Auth Object Config

Use the Cookie Params configured in Authentication Object.

### Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Validation Params

and list of Subject Alt Names for verification.

`skip_hostname_verification` - (Optional) is not matched to the connecting hostname (`Bool`).

`trusted_ca_url` - (Optional) The URL for a trust store (`String`).

`use_volterra_trusted_ca_url` - (Optional) Ignore the trusted CA URL and use the volterra trusted CA URL from the global config for verification. (`Bool`).

`verify_subject_alt_names` - (Optional) the hostname of the peer will be used for matching against SAN/CN of peer's certificate (`String`).

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Waf

A WAF object direct reference.

`waf` - (Optional) A direct reference to web application firewall configuration object. See [ref](#ref) below for details.

### Waf Rules

A set of direct references of WAF Rules objects.

`waf_rules` - (Optional) References to a set of WAF Rules configuration object. See [ref](#ref) below for details.

### Waf Type

```
 Enable/Disable individual WAF security rules.
```

`app_firewall` - (Optional) A direct reference to an Application Firewall configuration object. See [App Firewall ](#app-firewall) below for details.

`waf` - (Optional) A WAF object direct reference. See [Waf ](#waf) below for details.

`waf_rules` - (Optional) A set of direct references of WAF Rules objects. See [Waf Rules ](#waf-rules) below for details.

### Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured virtual_host.
