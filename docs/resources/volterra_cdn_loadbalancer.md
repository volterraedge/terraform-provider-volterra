---

page_title: "Volterra: cdn_loadbalancer"

description: "The cdn_loadbalancer allows CRUD of Cdn Loadbalancer resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_cdn_loadbalancer
==================================

The Cdn Loadbalancer allows CRUD of Cdn Loadbalancer resource on Volterra SaaS

~> **Note:** Please refer to [Cdn Loadbalancer API docs](https://volterra.io/docs/api/views-cdn-loadbalancer) to learn more

Example Usage
-------------

```hcl
resource "volterra_cdn_loadbalancer" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  domains = ["www.foo.com"]

  // One of the arguments from this list "https_auto_cert https http" must be set

  http {
    dns_volterra_managed = true
    port                 = "80"
  }
  origin_pool {
    follow_origin_redirect = true
    origin_request_timeout = "100s"

    origin_servers {
      // One of the arguments from this list "public_ip public_name" must be set

      public_ip {
        ip = "8.8.8.8"
      }
    }

    public_name {
      dns_name = "value"
    }

    // One of the arguments from this list "no_tls use_tls" must be set
    no_tls = true
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

`add_location` - (Optional) Appends header x-volterra-location = <re-site-name> in responses. (`Bool`).

`domains` - (Required) [This can be a domain or a sub-domain](`List of String`).

`http` - (Optional) CDN Distribution serving content over HTTP. See [Http ](#http) below for details.

`https` - (Optional) User is responsible for managing DNS.. See [Https ](#https) below for details.

`https_auto_cert` - (Optional) DNS records will be managed by Volterra.. See [Https Auto Cert ](#https-auto-cert) below for details.

`more_option` - (Optional) More options like header manipulation, compression etc.. See [More Option ](#more-option) below for details.

`origin_pool` - (Required) x-required. See [Origin Pool ](#origin-pool) below for details.

### Allow List

Allow list of countries.

`country_codes` - (Required) List of Country Codes (`List of Strings`).

`invert_match` - (Optional) Invert the match result. (`Bool`).

### Auth Options

Authentication Options.

`disable_auth` - (Optional) No Authenticaiton (bool).

`jwt` - (Optional) Enable JWT Authenticaiton. See [Jwt ](#jwt) below for details.

### Bearer Token

Token is found in the Bearer-Token.

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

### Block List

Block list of countries.

`country_codes` - (Required) List of Country Codes (`List of Strings`).

`invert_match` - (Optional) Invert the match result. (`Bool`).

### Cache Disabled

Disable Caching of content from the origin.

### Cache Ttl Options

Cache Options.

`cache_disabled` - (Optional) Disable Caching of content from the origin (bool).

`cache_ttl_default` - (Optional) Cache TTL value to use when the origin does not provide one (`String`).

`cache_ttl_override` - (Optional) Override the Cache TTL directive in the response from the origin (`String`).

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Client Log Options

Client request headers to log.

`header_list` - (Optional) List of headers (`String`).

### Cookie

Token is found in the cookie.

`name` - (Required) A case-insensitive cookie name. (`String`).

### Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Disable Auth

No Authenticaiton.

### Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Disable Sni

Do not use SNI..

### Geo Filtering

Geo filtering options.

`allow_list` - (Optional) Allow list of countries. See [Allow List ](#allow-list) below for details.

`block_list` - (Optional) Block list of countries. See [Block List ](#block-list) below for details.

### Header

Token is found in the header.

`name` - (Required) A case-insensitive field header name. (`String`).

### Header Options

Request/Response header related options.

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Request Headers To Add ](#request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Response Headers To Add ](#response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

### Http

CDN Distribution serving content over HTTP.

`dns_volterra_managed` - (Optional) or a DNS CNAME record should be created in your DNS provider's portal. (`Bool`).

`port` - (Optional) x-example: "80" (`Int`).

### Https

User is responsible for managing DNS..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

`tls_parameters` - (Optional) TLS parameters for the downstream connections.. See [Tls Parameters ](#tls-parameters) below for details.

### Https Auto Cert

DNS records will be managed by Volterra..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

`tls_config` - (Optional) TLS Configuration Parameters. See [Tls Config ](#tls-config) below for details.

### Ip Filtering

IP filtering options.

`allow_list` - (Optional) Allow list of ip prefixes. See [Allow List ](#allow-list) below for details.

`block_list` - (Optional) Block list of ip prefixes. See [Block List ](#block-list) below for details.

### Jwt

Enable JWT Authenticaiton.

`secret_key` - (Optional) Secret Key for JWT. See [Secret Key ](#secret-key) below for details.

`bearer_token` - (Optional) Token is found in the Bearer-Token (bool).

`cookie` - (Optional) Token is found in the cookie. See [Cookie ](#cookie) below for details.

`header` - (Optional) Token is found in the header. See [Header ](#header) below for details.

`query_param` - (Optional) Token is found in the Query-Param. See [Query Param ](#query-param) below for details.

### Logging Options

Logging related options.

`client_log_options` - (Optional) Client request headers to log. See [Client Log Options ](#client-log-options) below for details.

`origin_log_options` - (Optional) Origin response headers to log. See [Origin Log Options ](#origin-log-options) below for details.

### More Option

More options like header manipulation, compression etc..

`cache_ttl_options` - (Optional) Cache Options. See [Cache Ttl Options ](#cache-ttl-options) below for details.

`header_options` - (Optional) Request/Response header related options. See [Header Options ](#header-options) below for details.

`logging_options` - (Optional) Logging related options. See [Logging Options ](#logging-options) below for details.

`security_options` - (Optional) Security related options. See [Security Options ](#security-options) below for details.

### No Mtls

x-displayName: "Disable".

### No Tls

Origin servers do not use TLS.

### Origin Log Options

Origin response headers to log.

`header_list` - (Optional) List of headers (`String`).

### Origin Pool

x-required.

`follow_origin_redirect` - (Optional) Instructs the CDN to follow redirects from the origin server(s) (`Bool`).

`origin_request_timeout` - (Optional) Configures the time after which a request to the origin will time out waiting for a response (`String`).

`origin_servers` - (Required) List of original servers. See [Origin Servers ](#origin-servers) below for details.

`public_name` - (Required) The DNS name to be used as the host header for the request to the origin server. See [Public Name ](#public-name) below for details.

`no_tls` - (Optional) Origin servers do not use TLS (bool).

`use_tls` - (Optional) Origin servers use TLS. See [Use Tls ](#use-tls) below for details.

### Origin Servers

List of original servers.

`public_ip` - (Optional) Specify origin server with public IP. See [Public Ip ](#public-ip) below for details.

`public_name` - (Optional) Specify origin server with public DNS name. See [Public Name ](#public-name) below for details.

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Public Ip

Specify origin server with public IP.

`ip` - (Required) Public IP address (`String`).

### Public Name

Specify origin server with public DNS name.

`dns_name` - (Required) DNS Name (`String`).

### Query Param

Token is found in the Query-Param.

`key` - (Required) A case-sensitive HTTP query parameter name. (`String`).

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

### Secret Key

Secret Key for JWT.

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

### Security Options

Security related options.

`auth_options` - (Optional) Authentication Options. See [Auth Options ](#auth-options) below for details.

`geo_filtering` - (Optional) Geo filtering options. See [Geo Filtering ](#geo-filtering) below for details.

`ip_filtering` - (Optional) IP filtering options. See [Ip Filtering ](#ip-filtering) below for details.

### Skip Server Verification

Skip origin server verification.

### Tls 11 Plus

TLS v1.1+ with PFS ciphers and medium strength crypto algorithms..

### Tls 12 Plus

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Custom Hash Algorithms ](#custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Disable Ocsp Stapling ](#disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Use System Defaults ](#use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Tls Config

TLS Configuration Parameters.

`tls_11_plus` - (Optional) TLS v1.1+ with PFS ciphers and medium strength crypto algorithms. (bool).

`tls_12_plus` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (bool).

### Tls Parameters

TLS parameters for the downstream connections..

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Certificates ](#tls-certificates) below for details.

`tls_config` - (Optional) TLS Configuration Parameters. See [Tls Config ](#tls-config) below for details.

### Use Host Header As Sni

Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied..

### Use Mtls

x-displayName: "Enable".

`tls_certificates` - (Required) TLS Certificates. See [Tls Certificates ](#tls-certificates) below for details.

### Use Server Verification

Perform origin server verification using the provided trusted CA list.

`trusted_ca_url` - (Required) Trusted CA certificates for verification of Server's certificate (`String`).

### Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Use Tls

Origin servers use TLS.

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

### Volterra Trusted Ca

Perform origin server verification using F5XC default trusted CA list.

### Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured cdn_loadbalancer.
