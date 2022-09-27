---

page_title: "Volterra: origin_pool"

description: "The origin_pool allows CRUD of Origin Pool resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_origin_pool
=============================

The Origin Pool allows CRUD of Origin Pool resource on Volterra SaaS

~> **Note:** Please refer to [Origin Pool API docs](https://volterra.io/docs/api/views-origin-pool) to learn more

Example Usage
-------------

```hcl
resource "volterra_origin_pool" "example" {
  name                   = "acmecorp-web"
  namespace              = "staging"
  endpoint_selection     = ["endpoint_selection"]
  loadbalancer_algorithm = ["loadbalancer_algorithm"]

  origin_servers {
    // One of the arguments from this list "custom_endpoint_object private_name k8s_service private_ip consul_service vn_private_ip vn_private_name public_ip public_name" must be set

    private_name {
      dns_name = "value"

      // One of the arguments from this list "inside_network outside_network" must be set
      inside_network = true

      site_locator {
        // One of the arguments from this list "site virtual_site" must be set

        site {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
      }
    }

    labels = {
      "key1" = "value1"
    }
  }

  // One of the arguments from this list "port automatic_port" must be set
  port = "9080"

  // One of the arguments from this list "no_tls use_tls" must be set
  no_tls = true
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

### Advanced Options

Advanced options configuration like timeouts, circuit breaker, subset load balancing.

`circuit_breaker` - (Optional) allows to apply back pressure on downstream quickly.. See [Circuit Breaker ](#circuit-breaker) below for details.

`default_circuit_breaker` - (Optional) requests are 1024 and the default value for retries is 3 (bool).

`disable_circuit_breaker` - (Optional) Circuit Breaker is disabled (bool).

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2 seconds (`Int`).

`header_transformation_type` - (Optional) Header transformation options for upstream request headers. See [Header Transformation Type ](#header-transformation-type) below for details.

`http2_options` - (Optional) Http2 Protocol options for upstream connections. See [Http2 Options ](#http2-options) below for details.

`http_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 5 minutes. (`Int`).

`disable_outlier_detection` - (Optional) Outlier detection is disabled (bool).

`outlier_detection` - (Optional) healthy load balancing set. Outlier detection is a form of passive health checking.. See [Outlier Detection ](#outlier-detection) below for details.

`no_panic_threshold` - (Optional) Disable panic threshold. Only healthy endpoints are considered for load balancing. (bool).

`panic_threshold` - (Optional) all endpoints will be considered for load balancing ignoring its health status. (`Int`).

`disable_subsets` - (Optional) Subset load balancing is disabled. All eligible origin servers will be considered for load balancing. (bool).

`enable_subsets` - (Optional) Subset load balancing is enabled. Based on route, subset of origin servers will be considered for load balancing.. See [Enable Subsets ](#enable-subsets) below for details.

### Any Endpoint

Select any origin server from available healthy origin servers in this pool.

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

### Circuit Breaker

allows to apply back pressure on downstream quickly..

`connection_limit` - (Optional) Remove endpoint out of load balancing decision, if number of connections reach connection limit. (`Int`).

`max_requests` - (Optional) Remove endpoint out of load balancing decision, if requests exceed this count. (`Int`).

`pending_requests` - (Optional) Remove endpoint out of load balancing decision, if pending request reach pending_request. (`Int`).

`priority` - (Optional) matched with priority of CircuitBreaker to select the CircuitBreaker (`String`).

`retries` - (Optional) Remove endpoint out of load balancing decision, if retries for request exceed this count. (`Int`).

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Consul Service

Specify origin server with Hashi Corp Consul service name and site information.

`inside_network` - (Optional) Inside network on the site (bool).

`outside_network` - (Optional) Outside network on the site (bool).

`service_name` - (Required) cluster-id is optional. (`String`).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Site Locator ](#site-locator) below for details.

### Custom Endpoint Object

Specify origin server with a reference to endpoint object.

`endpoint` - (Required) Reference to an endpoint object. See [ref](#ref) below for details.

### Custom Hash Algorithms

Use hash algorithms in the custom order. Volterra will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Custom Security

Custom selection of TLS versions and cipher suites.

`cipher_suites` - (Required) The TLS listener will only support the specified cipher list. (`String`).

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

### Default Circuit Breaker

requests are 1024 and the default value for retries is 3.

### Default Header Transformation

Normalize the headers to lower case.

### Default Security

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Default Subset

Use the default subset provided here. Select endpoints matching default subset..

`default_subset` - (Optional) which gets used when route specifies no metadata or no subset matching the metadata exists. (`String`).

### Disable Circuit Breaker

Circuit Breaker is disabled.

### Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Disable Outlier Detection

Outlier detection is disabled.

### Disable Sni

Do not use SNI..

### Disable Subsets

Subset load balancing is disabled. All eligible origin servers will be considered for load balancing..

### Enable Subsets

Subset load balancing is enabled. Based on route, subset of origin servers will be considered for load balancing..

`endpoint_subsets` - (Optional) List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class.. See [Endpoint Subsets ](#endpoint-subsets) below for details.

`any_endpoint` - (Optional) Select any origin server from available healthy origin servers in this pool (bool).

`default_subset` - (Optional) Use the default subset provided here. Select endpoints matching default subset.. See [Default Subset ](#default-subset) below for details.

`fail_request` - (Optional) Request will be failed and error returned, as if cluster has no origin servers. (bool).

### Endpoint Subsets

List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class..

`keys` - (Optional) List of keys that define a cluster subset class. (`String`).

### Fail Request

Request will be failed and error returned, as if cluster has no origin servers..

### Header Transformation Type

Header transformation options for upstream request headers.

`default_header_transformation` - (Optional) Normalize the headers to lower case (bool).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (bool).

### Http2 Options

Http2 Protocol options for upstream connections.

`enabled` - (Optional) Enable/disable Http2 Protocol for upstream connections. It is disabled by default. (`Bool`).

### Inside Network

Inside network on the site.

### K8s Service

Specify origin server with K8s service name and site information.

`inside_network` - (Optional) Inside network on the site (bool).

`outside_network` - (Optional) Outside network on the site (bool).

`vk8s_networks` - (Optional) origin server are on vK8s network on the site (bool).

`service_name` - (Required) Both namespace and cluster-id are optional. (`String`).

`service_selector` - (Required) discovery has to happen. This implicit label is added to service_selector. See [Service Selector ](#service-selector) below for details.

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Site Locator ](#site-locator) below for details.

### Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### No Mtls

x-displayName: "Disable".

### No Panic Threshold

Disable panic threshold. Only healthy endpoints are considered for load balancing..

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

`base_ejection_time` - (Optional) Defaults to 30000ms or 30s. Specified in milliseconds. (`Int`).

`consecutive_5xx` - (Optional) a consecutive 5xx ejection occurs. Defaults to 5. (`Int`).

`consecutive_gateway_failure` - (Optional) before a consecutive gateway failure ejection occurs. Defaults to 5. (`Int`).

`interval` - (Optional) to 10000ms or 10s. Specified in milliseconds. (`Int`).

`max_ejection_percent` - (Optional) detection. Defaults to 10% but will eject at least one host regardless of the value. (`Int`).

### Outside Network

Outside network on the site.

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

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Private Name

Specify origin server with private or public DNS name and site information.

`dns_name` - (Required) DNS Name (`String`).

`inside_network` - (Optional) Inside network on the site (bool).

`outside_network` - (Optional) Outside network on the site (bool).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Site Locator ](#site-locator) below for details.

### Proper Case Header Transformation

For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are”.

### Public Ip

Specify origin server with public IP.

`ip` - (Required) Public IP address (`String`).

### Public Name

Specify origin server with public DNS name.

`dns_name` - (Required) DNS Name (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Service Selector

discovery has to happen. This implicit label is added to service_selector.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Site Locator

Site or Virtual site where this origin server is located.

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Skip Server Verification

Skip origin server verification.

### Tls Certificates

TLS Certificates.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. Volterra will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Custom Hash Algorithms ](#custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Disable Ocsp Stapling ](#disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) Volterra will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Use System Defaults ](#use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Tls Config

TLS parameters such as min/max TLS version and ciphers.

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Custom Security ](#custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (bool).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (bool).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (bool).

### Use Host Header As Sni

Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied..

### Use Mtls

x-displayName: "Enable".

`tls_certificates` - (Required) TLS Certificates. See [Tls Certificates ](#tls-certificates) below for details.

### Use Server Verification

Perform origin server verification using the provided trusted CA list.

`trusted_ca_url` - (Required) Trusted CA certificates for verification of Server's certificate (`String`).

### Use System Defaults

Volterra will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Use Tls

x-displayName: "Enable".

`no_mtls` - (Optional) x-displayName: "Disable" (bool).

`use_mtls` - (Optional) x-displayName: "Enable". See [Use Mtls ](#use-mtls) below for details.

`skip_server_verification` - (Optional) Skip origin server verification (bool).

`use_server_verification` - (Optional) Perform origin server verification using the provided trusted CA list. See [Use Server Verification ](#use-server-verification) below for details.

`volterra_trusted_ca` - (Optional) Perform origin server verification using Volterra default trusted CA list (bool).

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

### Vk8s Networks

origin server are on vK8s network on the site.

### Vn Private Ip

Specify origin server IP address on virtual network other than inside or outside network.

`ip` - (Required) IP address (`String`).

`virtual_network` - (Required) Virtual Network where this IP will be present. See [ref](#ref) below for details.

### Vn Private Name

Specify origin server name on virtual network other than inside or outside network.

`dns_name` - (Required) DNS Name (`String`).

`private_network` - (Required) Virtual Network where this Name will be present. See [ref](#ref) below for details.

### Volterra Trusted Ca

Perform origin server verification using Volterra default trusted CA list.

### Wingman Secret Info

Secret is given as bootstrap secret in Volterra Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured origin_pool.
