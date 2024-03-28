---

page_title: "Volterra: cdn_loadbalancer"

description: "The cdn_loadbalancer allows CRUD of Cdn Loadbalancer resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_cdn_loadbalancer
==================================

The Cdn Loadbalancer allows CRUD of Cdn Loadbalancer resource on Volterra SaaS

~> **Note:** Please refer to [Cdn Loadbalancer API docs](https://docs.cloud.f5.com/docs/api/views-cdn-loadbalancer) to learn more

Example Usage
-------------

```hcl
resource "volterra_cdn_loadbalancer" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  domains = ["www.foo.com"]

  // One of the arguments from this list "http https_auto_cert https" must be set

  https_auto_cert {
    add_hsts = true

    http_redirect = true

    tls_config {
      // One of the arguments from this list "tls_12_plus tls_11_plus" must be set
      tls_12_plus = true
    }
  }
  origin_pool {
    follow_origin_redirect = true

    more_origin_options {
      disable_byte_range_request = true

      websocket_proxy = true
    }

    origin_request_timeout = "100s"

    origin_servers {
      // One of the arguments from this list "public_ip public_name" must be set

      public_ip {
        // One of the arguments from this list "ip ipv6" must be set
        ip = "8.8.8.8"
      }

      port = "80"
    }

    public_name {
      dns_name = "value"

      refresh_interval = "20"
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

`http` - (Optional) CDN Distribution serving content over HTTP. See [Loadbalancer Type Http ](#loadbalancer-type-http) below for details.

`https` - (Optional) User is responsible for managing DNS.. See [Loadbalancer Type Https ](#loadbalancer-type-https) below for details.

`https_auto_cert` - (Optional) DNS records will be managed by Volterra.. See [Loadbalancer Type Https Auto Cert ](#loadbalancer-type-https-auto-cert) below for details.

`more_option` - (Optional) More options like header manipulation, compression etc.. See [More Option ](#more-option) below for details.

`origin_pool` - (Required) x-required. See [Origin Pool ](#origin-pool) below for details.

### More Option

More options like header manipulation, compression etc..

`cache_options` - (Optional) Cache Options. See [More Option Cache Options ](#more-option-cache-options) below for details.

`cache_ttl_options` - (Optional) Cache Options. See [More Option Cache Ttl Options ](#more-option-cache-ttl-options) below for details.(Deprecated)

`header_options` - (Optional) Request/Response header related options. See [More Option Header Options ](#more-option-header-options) below for details.

`logging_options` - (Optional) Logging related options. See [More Option Logging Options ](#more-option-logging-options) below for details.

`security_options` - (Optional) Security related options. See [More Option Security Options ](#more-option-security-options) below for details.

### Origin Pool

x-required.

`follow_origin_redirect` - (Optional) Instructs the CDN to follow redirects from the origin server(s) (`Bool`).(Deprecated)

`more_origin_options` - (Optional) x-displayName: "Advanced Configuration". See [Origin Pool More Origin Options ](#origin-pool-more-origin-options) below for details.

`origin_request_timeout` - (Optional) Configures the time after which a request to the origin will time out waiting for a response (`String`).

`origin_servers` - (Required) List of original servers. See [Origin Pool Origin Servers ](#origin-pool-origin-servers) below for details.

`public_name` - (Required) The DNS name to be used as the host header for the request to the origin server. See [Origin Pool Public Name ](#origin-pool-public-name) below for details.

###### One of the arguments from this list "use_tls, no_tls" must be set

`no_tls` - (Optional) Origin servers do not use TLS (`Bool`).

`use_tls` - (Optional) Origin servers use TLS. See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

### Auth Options Custom

Enable Custom Authenticaiton.

`custom_auth_config` - (Optional) This is custom authentication configuration parameters. Please reach out to the support for custom authentication details. (`String`).

### Auth Options Disable Auth

No Authenticaiton.

### Auth Options Jwt

Enable JWT Authenticaiton.

`backup_key` - (Optional) Backup JWT Key - If specified is also checked in addition to the primary secret key. See [Jwt Backup Key ](#jwt-backup-key) below for details.

`secret_key` - (Required) Secret Key for JWT. See [Jwt Secret Key ](#jwt-secret-key) below for details.

###### One of the arguments from this list "header, cookie, query_param, bearer_token" can be set

`bearer_token` - (Optional) Token is found in the Bearer-Token (`Bool`).

`cookie` - (Optional) Token is found in the cookie. See [Token Source Cookie ](#token-source-cookie) below for details.

`header` - (Optional) Token is found in the header. See [Token Source Header ](#token-source-header) below for details.

`query_param` - (Optional) Token is found in the Query-Param. See [Token Source Query Param ](#token-source-query-param) below for details.

### Backup Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Cache Actions Cache Bypass

Bypass Caching of content from the origin.

### Cache Actions Cache Disabled

Disable Caching of content from the origin.

### Cache Actions Eligible For Cache

Eligible for caching the content.

###### One of the arguments from this list "scheme_hostname_uri, scheme_hostname_request_uri, hostname_uri, scheme_hostname_uri_query, scheme_proxy_host_uri, scheme_proxy_host_request_uri" must be set

`hostname_uri` - (Optional) . See [Eligible For Cache Hostname Uri ](#eligible-for-cache-hostname-uri) below for details.(Deprecated)

`scheme_hostname_request_uri` - (Optional) . See [Eligible For Cache Scheme Hostname Request Uri ](#eligible-for-cache-scheme-hostname-request-uri) below for details.(Deprecated)

`scheme_hostname_uri` - (Optional) . See [Eligible For Cache Scheme Hostname Uri ](#eligible-for-cache-scheme-hostname-uri) below for details.(Deprecated)

`scheme_hostname_uri_query` - (Optional) . See [Eligible For Cache Scheme Hostname Uri Query ](#eligible-for-cache-scheme-hostname-uri-query) below for details.(Deprecated)

`scheme_proxy_host_request_uri` - (Optional) . See [Eligible For Cache Scheme Proxy Host Request Uri ](#eligible-for-cache-scheme-proxy-host-request-uri) below for details.

`scheme_proxy_host_uri` - (Optional) . See [Eligible For Cache Scheme Proxy Host Uri ](#eligible-for-cache-scheme-proxy-host-uri) below for details.

### Cache Actions Eligible For Cache

Eligible for caching the content.

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

### Cache Headers Operator

Available operators.

###### One of the arguments from this list "DoesNotEndWith, MatchRegex, DoesNotEqual, DoesNotContain, Startswith, Endswith, Equals, Contains, DoesNotStartWith" can be set

`Contains` - (Optional) Field must contain (`String`).

`DoesNotContain` - (Optional) Field must not contain (`String`).

`DoesNotEndWith` - (Optional) Field must not end with (`String`).

`DoesNotEqual` - (Optional) Field must not equal (`String`).

`DoesNotStartWith` - (Optional) Field must not start with (`String`).

`Endswith` - (Optional) Field must end with (`String`).

`Equals` - (Optional) Field must exactly match (`String`).

`MatchRegex` - (Optional) Field matches regular expression (`String`).

`Startswith` - (Optional) Field must start with (`String`).

### Cache Options Cache Rules

Rules are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs..

###### One of the arguments from this list "cache_bypass, eligible_for_cache" must be set

`cache_bypass` - (Optional) Bypass Caching of content from the origin (`Bool`).

`eligible_for_cache` - (Optional) Eligible for caching the content. See [Cache Actions Eligible For Cache ](#cache-actions-eligible-for-cache) below for details.

`rule_expression_list` - (Required) Expressions are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs... See [Cache Rules Rule Expression List ](#cache-rules-rule-expression-list) below for details.

`rule_name` - (Required) Name of the Cache Rule (`String`).

### Cache Options Default Cache Action

Default value for Cache action..

###### One of the arguments from this list "eligible_for_cache, cache_ttl_default, cache_ttl_override, cache_disabled" can be set

`cache_disabled` - (Optional) Disable Caching of content from the origin (`Bool`).

`cache_ttl_default` - (Optional) Cache TTL value to use when the origin does not provide one (`String`).

`cache_ttl_override` - (Optional) Override the Cache TTL directive in the response from the origin (`String`).

`eligible_for_cache` - (Optional) Eligible for caching the content. See [Cache Actions Eligible For Cache ](#cache-actions-eligible-for-cache) below for details.(Deprecated)

### Cache Rule Expression Cache Headers

Configure cache rule headers to match the criteria.

`name` - (Optional) Name of the header (`String`).

`operator` - (Optional) Available operators. See [Cache Headers Operator ](#cache-headers-operator) below for details.

### Cache Rule Expression Cookie Matcher

Note that all specified cookie matcher predicates must evaluate to true..

`name` - (Required) A case-sensitive cookie name. (`String`).

`operator` - (Optional) . See [Cookie Matcher Operator ](#cookie-matcher-operator) below for details.

### Cache Rule Expression Path Match

URI path of route.

`operator` - (Optional) A specification of path match. See [Path Match Operator ](#path-match-operator) below for details.

### Cache Rule Expression Query Parameters

List of (key, value) query parameters.

`key` - (Required) In the above example, assignee_username is the key (`String`).

`operator` - (Optional) . See [Query Parameters Operator ](#query-parameters-operator) below for details.

### Cache Rules Rule Expression List

Expressions are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs...

`cache_rule_expression` - (Required) The Cache Rule Expression Terms that are ANDed. See [Rule Expression List Cache Rule Expression ](#rule-expression-list-cache-rule-expression) below for details.

`expression_name` - (Required) Name of the Expressions items that are ANDed (`String`).

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

### Choice Public Ip

Specify origin server with public IP.

###### One of the arguments from this list "ip, ipv6" must be set

`ip` - (Optional) Public IPV4 address (`String`).

`ipv6` - (Optional) Public IPV6 address (`String`).

### Choice Public Name

Specify origin server with public DNS name.

`dns_name` - (Required) DNS Name (`String`).

`refresh_interval` - (Optional) Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767 (`Int`).

### Choice Tls 11 Plus

TLS v1.1+ with PFS ciphers and medium strength crypto algorithms..

### Choice Tls 12 Plus

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Cookie Matcher Operator

.

###### One of the arguments from this list "Equals, Contains, DoesNotStartWith, DoesNotEqual, DoesNotContain, Startswith, Endswith, DoesNotEndWith, MatchRegex" can be set

`Contains` - (Optional) Field must contain (`String`).

`DoesNotContain` - (Optional) Field must not contain (`String`).

`DoesNotEndWith` - (Optional) Field must not end with (`String`).

`DoesNotEqual` - (Optional) Field must not equal (`String`).

`DoesNotStartWith` - (Optional) Field must not start with (`String`).

`Endswith` - (Optional) Field must end with (`String`).

`Equals` - (Optional) Field must exactly match (`String`).

`MatchRegex` - (Optional) Field matches regular expression (`String`).

`Startswith` - (Optional) Field must start with (`String`).

### Eligible For Cache Hostname Uri

.

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Eligible For Cache Scheme Hostname Request Uri

.

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Eligible For Cache Scheme Hostname Uri

.

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Eligible For Cache Scheme Hostname Uri Query

.

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Eligible For Cache Scheme Proxy Host Request Uri

.

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Eligible For Cache Scheme Proxy Host Uri

.

`cache_override` - (Optional) Honour Cache Override (`Bool`).

`cache_ttl` - (Required) Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days (`String`).

`ignore_response_cookie` - (Optional) By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present. (`Bool`).

### Geo Filtering Type Allow List

Allow list of countries.

`country_codes` - (Required) List of Country Codes (`List of Strings`).

`invert_match` - (Optional) Invert the match result. (`Bool`).

### Geo Filtering Type Block List

Block list of countries.

`country_codes` - (Required) List of Country Codes (`List of Strings`).

`invert_match` - (Optional) Invert the match result. (`Bool`).

### Header Options Request Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "value, secret_value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Header Options Response Headers To Add

Headers specified at this level are applied after headers from matched Route are applied.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "value, secret_value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Https Tls Parameters

TLS parameters for the downstream connections..

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Parameters Tls Certificates ](#tls-parameters-tls-certificates) below for details.

`tls_config` - (Optional) TLS Configuration Parameters. See [Tls Parameters Tls Config ](#tls-parameters-tls-config) below for details.

### Https Auto Cert Tls Config

TLS Configuration Parameters.

###### One of the arguments from this list "tls_12_plus, tls_11_plus" must be set

`tls_11_plus` - (Optional) TLS v1.1+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

`tls_12_plus` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

### Ip Filtering Type Allow List

Allow list of ip prefixes.

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

### Ip Filtering Type Block List

Block list of ip prefixes.

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).

### Jwt Backup Key

Backup JWT Key - If specified is also checked in addition to the primary secret key.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Backup Key Blindfold Secret Info Internal ](#backup-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "vault_secret_info, clear_secret_info, wingman_secret_info, blindfold_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Jwt Secret Key

Secret Key for JWT.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Secret Key Blindfold Secret Info Internal ](#secret-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Loadbalancer Type Http

CDN Distribution serving content over HTTP.

`dns_volterra_managed` - (Optional) or a DNS CNAME record should be created in your DNS provider's portal. (`Bool`).

###### One of the arguments from this list "port, port_ranges" must be set

`port` - (Optional) HTTP port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

### Loadbalancer Type Https

User is responsible for managing DNS..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

`tls_parameters` - (Optional) TLS parameters for the downstream connections.. See [Https Tls Parameters ](#https-tls-parameters) below for details.

### Loadbalancer Type Https Auto Cert

DNS records will be managed by Volterra..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

`tls_config` - (Optional) TLS Configuration Parameters. See [Https Auto Cert Tls Config ](#https-auto-cert-tls-config) below for details.

### Logging Options Client Log Options

Client request headers to log.

`header_list` - (Optional) List of headers (`String`).

### Logging Options Origin Log Options

Origin response headers to log.

`header_list` - (Optional) List of headers (`String`).

### More Option Cache Options

Cache Options.

`cache_rules` - (Optional) Rules are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs.. See [Cache Options Cache Rules ](#cache-options-cache-rules) below for details.

`default_cache_action` - (Required) Default value for Cache action.. See [Cache Options Default Cache Action ](#cache-options-default-cache-action) below for details.

### More Option Cache Ttl Options

Cache Options.

###### One of the arguments from this list "cache_disabled, cache_ttl_default, cache_ttl_override" can be set

`cache_disabled` - (Optional) Disable Caching of content from the origin (`Bool`).

`cache_ttl_default` - (Optional) Cache TTL value to use when the origin does not provide one (`String`).

`cache_ttl_override` - (Optional) Override the Cache TTL directive in the response from the origin (`String`).

### More Option Header Options

Request/Response header related options.

`request_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Header Options Request Headers To Add ](#header-options-request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) Headers specified at this level are applied after headers from matched Route are applied. See [Header Options Response Headers To Add ](#header-options-response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

### More Option Logging Options

Logging related options.

`client_log_options` - (Optional) Client request headers to log. See [Logging Options Client Log Options ](#logging-options-client-log-options) below for details.

`origin_log_options` - (Optional) Origin response headers to log. See [Logging Options Origin Log Options ](#logging-options-origin-log-options) below for details.

### More Option Security Options

Security related options.

`auth_options` - (Optional) Authentication Options. See [Security Options Auth Options ](#security-options-auth-options) below for details.

`geo_filtering` - (Optional) Geo filtering options. See [Security Options Geo Filtering ](#security-options-geo-filtering) below for details.

`ip_filtering` - (Optional) IP filtering options. See [Security Options Ip Filtering ](#security-options-ip-filtering) below for details.

### Mtls Choice No Mtls

x-displayName: "Disable".

### Mtls Choice Use Mtls

x-displayName: "Upload a client authentication certificate specifically for this Origin Pool".

`tls_certificates` - (Required) mTLS Client Certificate. See [Use Mtls Tls Certificates ](#use-mtls-tls-certificates) below for details.

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ocsp Stapling Choice Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Ocsp Stapling Choice Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Origin Pool More Origin Options

x-displayName: "Advanced Configuration".

`disable_byte_range_request` - (Optional) Choice to enable/disable origin byte range requrests towards origin (`Bool`).

`websocket_proxy` - (Optional) Option to enable proxying of websocket connections to the origin server (`Bool`).

### Origin Pool Origin Servers

List of original servers.

###### One of the arguments from this list "public_name, public_ip" must be set

`public_ip` - (Optional) Specify origin server with public IP. See [Choice Public Ip ](#choice-public-ip) below for details.

`public_name` - (Optional) Specify origin server with public DNS name. See [Choice Public Name ](#choice-public-name) below for details.

`port` - (Optional) Port the workload can be reached on (`Int`).

### Origin Pool Public Name

The DNS name to be used as the host header for the request to the origin server.

`dns_name` - (Required) DNS Name (`String`).

`refresh_interval` - (Optional) Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767 (`Int`).

### Path Match Operator

A specification of path match.

###### One of the arguments from this list "DoesNotEqual, DoesNotContain, Startswith, Endswith, DoesNotEndWith, MatchRegex, Equals, Contains, DoesNotStartWith" can be set

`Contains` - (Optional) Field must contain (`String`).

`DoesNotContain` - (Optional) Field must not contain (`String`).

`DoesNotEndWith` - (Optional) Field must not end with (`String`).

`DoesNotEqual` - (Optional) Field must not equal (`String`).

`DoesNotStartWith` - (Optional) Field must not start with (`String`).

`Endswith` - (Optional) Field must end with (`String`).

`Equals` - (Optional) Field must exactly match (`String`).

`MatchRegex` - (Optional) Field matches regular expression (`String`).

`Startswith` - (Optional) Field must start with (`String`).

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Query Parameters Operator

.

###### One of the arguments from this list "Equals, Contains, DoesNotStartWith, DoesNotEndWith, MatchRegex, DoesNotEqual, DoesNotContain, Startswith, Endswith" can be set

`Contains` - (Optional) Field must contain (`String`).

`DoesNotContain` - (Optional) Field must not contain (`String`).

`DoesNotEndWith` - (Optional) Field must not end with (`String`).

`DoesNotEqual` - (Optional) Field must not equal (`String`).

`DoesNotStartWith` - (Optional) Field must not start with (`String`).

`Endswith` - (Optional) Field must end with (`String`).

`Equals` - (Optional) Field must exactly match (`String`).

`MatchRegex` - (Optional) Field matches regular expression (`String`).

`Startswith` - (Optional) Field must start with (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rule Expression List Cache Rule Expression

The Cache Rule Expression Terms that are ANDed.

`cache_headers` - (Optional) Configure cache rule headers to match the criteria. See [Cache Rule Expression Cache Headers ](#cache-rule-expression-cache-headers) below for details.

`cookie_matcher` - (Optional) Note that all specified cookie matcher predicates must evaluate to true.. See [Cache Rule Expression Cookie Matcher ](#cache-rule-expression-cookie-matcher) below for details.

`path_match` - (Optional) URI path of route. See [Cache Rule Expression Path Match ](#cache-rule-expression-path-match) below for details.

`query_parameters` - (Optional) List of (key, value) query parameters. See [Cache Rule Expression Query Parameters ](#cache-rule-expression-query-parameters) below for details.

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

### Secret Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Value Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Security Options Auth Options

Authentication Options.

###### One of the arguments from this list "disable_auth, jwt, custom" can be set

`custom` - (Optional) Enable Custom Authenticaiton. See [Auth Options Custom ](#auth-options-custom) below for details.

`disable_auth` - (Optional) No Authenticaiton (`Bool`).

`jwt` - (Optional) Enable JWT Authenticaiton. See [Auth Options Jwt ](#auth-options-jwt) below for details.

### Security Options Geo Filtering

Geo filtering options.

###### One of the arguments from this list "allow_list, block_list" can be set

`allow_list` - (Optional) Allow list of countries. See [Geo Filtering Type Allow List ](#geo-filtering-type-allow-list) below for details.

`block_list` - (Optional) Block list of countries. See [Geo Filtering Type Block List ](#geo-filtering-type-block-list) below for details.

### Security Options Ip Filtering

IP filtering options.

###### One of the arguments from this list "block_list, allow_list" can be set

`allow_list` - (Optional) Allow list of ip prefixes. See [Ip Filtering Type Allow List ](#ip-filtering-type-allow-list) below for details.

`block_list` - (Optional) Block list of ip prefixes. See [Ip Filtering Type Block List ](#ip-filtering-type-block-list) below for details.

### Server Validation Choice Skip Server Verification

Skip origin server verification.

### Server Validation Choice Use Server Verification

Perform origin server verification using the provided Root CA Certificate.

###### One of the arguments from this list "trusted_ca_url, trusted_ca" must be set

`trusted_ca` - (Optional) Select/Add a Root CA Certificate object to associate with this Origin Pool for verification of server's certificate. See [ref](#ref) below for details.

`trusted_ca_url` - (Optional) Upload a Root CA Certificate specifically for this Origin Pool for verification of server's certificate (`String`).

### Server Validation Choice Volterra Trusted Ca

Perform origin server verification using F5XC Default Root CA Certificate.

### Sni Choice Disable Sni

Do not use SNI..

### Sni Choice Use Host Header As Sni

Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied..

### Tls Certificates Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Tls Choice No Tls

Origin servers do not use TLS.

### Tls Choice Use Tls

Origin servers use TLS.

###### One of the arguments from this list "no_mtls, use_mtls, use_mtls_obj" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Upload a client authentication certificate specifically for this Origin Pool". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`use_mtls_obj` - (Optional) x-displayName: "Select/add a TLS Certificate object for client authentication". See [ref](#ref) below for details.

###### One of the arguments from this list "use_server_verification, skip_server_verification, volterra_trusted_ca" must be set

`skip_server_verification` - (Optional) Skip origin server verification (`Bool`).

`use_server_verification` - (Optional) Perform origin server verification using the provided Root CA Certificate. See [Server Validation Choice Use Server Verification ](#server-validation-choice-use-server-verification) below for details.

`volterra_trusted_ca` - (Optional) Perform origin server verification using F5XC Default Root CA Certificate (`Bool`).

###### One of the arguments from this list "sni, use_host_header_as_sni, disable_sni" must be set

`disable_sni` - (Optional) Do not use SNI. (`Bool`).

`sni` - (Optional) SNI value to be used. (`String`).

`use_host_header_as_sni` - (Optional) Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied. (`Bool`).

`tls_config` - (Required) TLS parameters such as min/max TLS version and ciphers. See [Use Tls Tls Config ](#use-tls-tls-config) below for details.

### Tls Parameters Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "use_system_defaults, disable_ocsp_stapling, custom_hash_algorithms" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Tls Parameters Tls Config

TLS Configuration Parameters.

###### One of the arguments from this list "tls_12_plus, tls_11_plus" must be set

`tls_11_plus` - (Optional) TLS v1.1+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

`tls_12_plus` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

### Token Source Bearer Token

Token is found in the Bearer-Token.

### Token Source Cookie

Token is found in the cookie.

`name` - (Required) A case-insensitive cookie name. (`String`).

### Token Source Header

Token is found in the header.

`name` - (Required) A case-insensitive field header name. (`String`).

### Token Source Query Param

Token is found in the Query-Param.

`key` - (Required) A case-sensitive HTTP query parameter name. (`String`).

### Ttl Options Cache Disabled

Disable Caching of content from the origin.

### Use Mtls Tls Certificates

mTLS Client Certificate.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "use_system_defaults, disable_ocsp_stapling, custom_hash_algorithms" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Use Tls Tls Config

TLS parameters such as min/max TLS version and ciphers.

###### One of the arguments from this list "default_security, medium_security, low_security, custom_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Value Choice Secret Value

Secret Value of the HTTP header..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Secret Value Blindfold Secret Info Internal ](#secret-value-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "vault_secret_info, clear_secret_info, wingman_secret_info, blindfold_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

Attribute Reference
-------------------

-	`id` - This is the id of the configured cdn_loadbalancer.
