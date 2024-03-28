---

page_title: "Volterra: origin_pool"

description: "The origin_pool allows CRUD of Origin Pool resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_origin_pool
=============================

The Origin Pool allows CRUD of Origin Pool resource on Volterra SaaS

~> **Note:** Please refer to [Origin Pool API docs](https://docs.cloud.f5.com/docs/api/views-origin-pool) to learn more

Example Usage
-------------

```hcl
resource "volterra_origin_pool" "example" {
  name                   = "acmecorp-web"
  namespace              = "staging"
  endpoint_selection     = ["endpoint_selection"]
  loadbalancer_algorithm = ["loadbalancer_algorithm"]

  origin_servers {
    // One of the arguments from this list "segment_ip private_name consul_service custom_endpoint_object vn_private_ip vn_private_name public_ip public_name private_ip k8s_service segment_name" must be set

    vn_private_name {
      dns_name = "value"

      private_network {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }
    }

    labels = {
      "key1" = "value1"
    }
  }

  // One of the arguments from this list "lb_port port automatic_port" must be set
  lb_port = true

  // One of the arguments from this list "use_tls no_tls" must be set
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

`same_as_endpoint_port` - (Optional) Health check is performed on endpoint port itself (`Bool`).

`healthcheck` - (Optional) Reference to healthcheck configuration objects. See [ref](#ref) below for details.

`loadbalancer_algorithm` - (Required) loadbalancer_algorithm to determine which host is selected. (`String`).

`origin_servers` - (Required) List of origin servers in this pool. See [Origin Servers ](#origin-servers) below for details.

`automatic_port` - (Optional) For other origin server types, port will be automatically set as 443 if TLS is enabled at Origin Pool and 80 if TLS is disabled (`Bool`).

`lb_port` - (Optional) Endpoint port is selected based on loadbalancer port (`Bool`).

`port` - (Optional) Endpoint service is available on this port (`Int`).

`no_tls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_tls` - (Optional) x-displayName: "Enable". See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

### Advanced Options

Advanced options configuration like timeouts, circuit breaker, subset load balancing.

###### One of the arguments from this list "default_circuit_breaker, disable_circuit_breaker, circuit_breaker" must be set

`circuit_breaker` - (Optional) allows to apply back pressure on downstream quickly.. See [Circuit Breaker Choice Circuit Breaker ](#circuit-breaker-choice-circuit-breaker) below for details.

`default_circuit_breaker` - (Optional) requests are 1024 and the default value for retries is 3 (`Bool`).

`disable_circuit_breaker` - (Optional) Circuit Breaker is disabled (`Bool`).

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2 seconds (`Int`).

`header_transformation_type` - (Optional) Settings to normalize the headers of upstream requests.. See [Advanced Options Header Transformation Type ](#advanced-options-header-transformation-type) below for details.

`http_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 5 minutes. (`Int`).

###### One of the arguments from this list "auto_http_config, http1_config, http2_options" must be set

`auto_http_config` - (Optional) and will use whichever protocol is negotiated by ALPN with the upstream. (`Bool`).

`http1_config` - (Optional) Enable HTTP/1.1 for upstream connections (`Bool`).

`http2_options` - (Optional) Enable HTTP/2 for upstream connections.. See [Http Protocol Type Http2 Options ](#http-protocol-type-http2-options) below for details.

###### One of the arguments from this list "enable_lb_source_ip_persistance, disable_lb_source_ip_persistance" can be set

`disable_lb_source_ip_persistance` - (Optional) Disable LB source IP persistance (`Bool`).(Deprecated)

`enable_lb_source_ip_persistance` - (Optional) Enable LB source IP persistance (`Bool`).(Deprecated)

###### One of the arguments from this list "disable_outlier_detection, outlier_detection" must be set

`disable_outlier_detection` - (Optional) Outlier detection is disabled (`Bool`).

`outlier_detection` - (Optional) healthy load balancing set. Outlier detection is a form of passive health checking.. See [Outlier Detection Choice Outlier Detection ](#outlier-detection-choice-outlier-detection) below for details.

###### One of the arguments from this list "no_panic_threshold, panic_threshold" must be set

`no_panic_threshold` - (Optional) Disable panic threshold. Only healthy endpoints are considered for load balancing. (`Bool`).

`panic_threshold` - (Optional) all endpoints will be considered for load balancing ignoring its health status. (`Int`).

###### One of the arguments from this list "disable_subsets, enable_subsets" must be set

`disable_subsets` - (Optional) Subset load balancing is disabled. All eligible origin servers will be considered for load balancing. (`Bool`).

`enable_subsets` - (Optional) Subset load balancing is enabled. Based on route, subset of origin servers will be considered for load balancing.. See [Subset Choice Enable Subsets ](#subset-choice-enable-subsets) below for details.

### Origin Servers

List of origin servers in this pool.

###### One of the arguments from this list "public_name, private_ip, k8s_service, segment_name, public_ip, consul_service, custom_endpoint_object, vn_private_ip, vn_private_name, segment_ip, private_name" must be set

`consul_service` - (Optional) Specify origin server with Hashi Corp Consul service name and site information. See [Choice Consul Service ](#choice-consul-service) below for details.

`custom_endpoint_object` - (Optional) Specify origin server with a reference to endpoint object. See [Choice Custom Endpoint Object ](#choice-custom-endpoint-object) below for details.

`k8s_service` - (Optional) Specify origin server with K8s service name and site information. See [Choice K8s Service ](#choice-k8s-service) below for details.

`private_ip` - (Optional) Specify origin server with private or public IP address and site information. See [Choice Private Ip ](#choice-private-ip) below for details.

`private_name` - (Optional) Specify origin server with private or public DNS name and site information. See [Choice Private Name ](#choice-private-name) below for details.

`public_ip` - (Optional) Specify origin server with public IP. See [Choice Public Ip ](#choice-public-ip) below for details.

`public_name` - (Optional) Specify origin server with public DNS name. See [Choice Public Name ](#choice-public-name) below for details.

`vn_private_ip` - (Optional) Specify origin server IP address on virtual network other than inside or outside network. See [Choice Vn Private Ip ](#choice-vn-private-ip) below for details.

`vn_private_name` - (Optional) Specify origin server name on virtual network other than inside or outside network. See [Choice Vn Private Name ](#choice-vn-private-name) below for details.

`labels` - (Optional) Add Labels for this origin server, these labels can be used to form subset. (`String`).

### Advanced Options Header Transformation Type

Settings to normalize the headers of upstream requests..

###### One of the arguments from this list "default_header_transformation, proper_case_header_transformation, preserve_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

### Choice Consul Service

Specify origin server with Hashi Corp Consul service name and site information.

###### One of the arguments from this list "inside_network, outside_network" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`service_name` - (Required) cluster-id is optional. (`String`).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Consul Service Site Locator ](#consul-service-site-locator) below for details.

### Choice Custom Endpoint Object

Specify origin server with a reference to endpoint object.

`endpoint` - (Required) Reference to an endpoint object. See [ref](#ref) below for details.

### Choice Custom Security

Custom selection of TLS versions and cipher suites.

`cipher_suites` - (Required) The TLS listener will only support the specified cipher list. (`String`).

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

### Choice Default Security

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Choice K8s Service

Specify origin server with K8s service name and site information.

###### One of the arguments from this list "outside_network, vk8s_networks, inside_network" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`vk8s_networks` - (Optional) origin server are on vK8s network on the site (`Bool`).

###### One of the arguments from this list "service_selector, service_name" must be set

`service_name` - (Optional) Both namespace and cluster-id are optional. (`String`).

`service_selector` - (Optional) discovery has to happen. This implicit label is added to service_selector. See [Service Info Service Selector ](#service-info-service-selector) below for details.(Deprecated)

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [K8s Service Site Locator ](#k8s-service-site-locator) below for details.

### Choice Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Choice Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### Choice Private Ip

Specify origin server with private or public IP address and site information.

###### One of the arguments from this list "inside_network, outside_network" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

###### One of the arguments from this list "ip, ipv6" must be set

`ip` - (Optional) Private IPV4 address (`String`).

`ipv6` - (Optional) Private IPV6 address (`String`).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Private Ip Site Locator ](#private-ip-site-locator) below for details.

### Choice Private Name

Specify origin server with private or public DNS name and site information.

`dns_name` - (Required) DNS Name (`String`).

###### One of the arguments from this list "inside_network, outside_network" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`refresh_interval` - (Optional) Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767 (`Int`).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Private Name Site Locator ](#private-name-site-locator) below for details.

### Choice Public Ip

Specify origin server with public IP.

###### One of the arguments from this list "ip, ipv6" must be set

`ip` - (Optional) Public IPV4 address (`String`).

`ipv6` - (Optional) Public IPV6 address (`String`).

### Choice Public Name

Specify origin server with public DNS name.

`dns_name` - (Required) DNS Name (`String`).

`refresh_interval` - (Optional) Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767 (`Int`).

### Choice Vn Private Ip

Specify origin server IP address on virtual network other than inside or outside network.

`virtual_network` - (Required) Virtual Network where this IP will be present. See [ref](#ref) below for details.

###### One of the arguments from this list "ip, ipv6" must be set

`ip` - (Optional) IPV4 address (`String`).

`ipv6` - (Optional) IPV6 address (`String`).

### Choice Vn Private Name

Specify origin server name on virtual network other than inside or outside network.

`dns_name` - (Required) DNS Name (`String`).

`private_network` - (Required) Virtual Network where this Name will be present. See [ref](#ref) below for details.

### Circuit Breaker Choice Circuit Breaker

allows to apply back pressure on downstream quickly..

`connection_limit` - (Optional) Remove endpoint out of load balancing decision, if number of connections reach connection limit. (`Int`).

`max_requests` - (Optional) Remove endpoint out of load balancing decision, if requests exceed this count. (`Int`).

`pending_requests` - (Optional) Remove endpoint out of load balancing decision, if pending request reach pending_request. (`Int`).

`priority` - (Optional) matched with priority of CircuitBreaker to select the CircuitBreaker (`String`).

`retries` - (Optional) Remove endpoint out of load balancing decision, if retries for request exceed this count. (`Int`).

### Circuit Breaker Choice Default Circuit Breaker

requests are 1024 and the default value for retries is 3.

### Circuit Breaker Choice Disable Circuit Breaker

Circuit Breaker is disabled.

### Consul Service Site Locator

Site or Virtual site where this origin server is located.

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Enable Subsets Endpoint Subsets

List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class..

`keys` - (Required) List of keys that define a cluster subset class. (`String`).

### Fallback Policy Choice Any Endpoint

Select any origin server from available healthy origin servers in this pool.

### Fallback Policy Choice Default Subset

Use the default subset provided here. Select endpoints matching default subset..

`default_subset` - (Optional) which gets used when route specifies no metadata or no subset matching the metadata exists. (`String`).

### Fallback Policy Choice Fail Request

Request will be failed and error returned, as if cluster has no origin servers..

### Header Transformation Choice Default Header Transformation

Normalize the headers to lower case.

### Header Transformation Choice Preserve Case Header Transformation

Preserves the original case of headers without any modifications..

### Header Transformation Choice Proper Case Header Transformation

For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are”.

### Http Protocol Type Auto Http Config

and will use whichever protocol is negotiated by ALPN with the upstream..

### Http Protocol Type Http1 Config

Enable HTTP/1.1 for upstream connections.

### Http Protocol Type Http2 Options

Enable HTTP/2 for upstream connections..

`enabled` - (Optional) Enable/disable HTTP2 Protocol for upstream connections (`Bool`).

### K8s Service Site Locator

Site or Virtual site where this origin server is located.

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Lb Source Ip Persistance Choice Disable Lb Source Ip Persistance

Disable LB source IP persistance.

### Lb Source Ip Persistance Choice Enable Lb Source Ip Persistance

Enable LB source IP persistance.

### Mtls Choice No Mtls

x-displayName: "Disable".

### Mtls Choice Use Mtls

x-displayName: "Upload a client authentication certificate specifically for this Origin Pool".

`tls_certificates` - (Required) mTLS Client Certificate. See [Use Mtls Tls Certificates ](#use-mtls-tls-certificates) below for details.

### Network Choice Inside Network

Inside network on the site.

### Network Choice Outside Network

Outside network on the site.

### Network Choice Vk8s Networks

origin server are on vK8s network on the site.

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ocsp Stapling Choice Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Ocsp Stapling Choice Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Outlier Detection Choice Disable Outlier Detection

Outlier detection is disabled.

### Outlier Detection Choice Outlier Detection

healthy load balancing set. Outlier detection is a form of passive health checking..

`base_ejection_time` - (Optional) Defaults to 30000ms or 30s. Specified in milliseconds. (`Int`).

`consecutive_5xx` - (Optional) a consecutive 5xx ejection occurs. Defaults to 5. (`Int`).

`consecutive_gateway_failure` - (Optional) before a consecutive gateway failure ejection occurs. Defaults to 5. (`Int`).

`interval` - (Optional) to 10000ms or 10s. Specified in milliseconds. (`Int`).

`max_ejection_percent` - (Optional) detection. Defaults to 10% but will eject at least one host regardless of the value. (`Int`).

### Panic Threshold Type No Panic Threshold

Disable panic threshold. Only healthy endpoints are considered for load balancing..

### Private Ip Site Locator

Site or Virtual site where this origin server is located.

###### One of the arguments from this list "virtual_site, site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Private Name Site Locator

Site or Virtual site where this origin server is located.

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

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

### Server Validation Choice Skip Server Verification

Skip origin server verification.

### Server Validation Choice Use Server Verification

Perform origin server verification using the provided Root CA Certificate.

###### One of the arguments from this list "trusted_ca_url, trusted_ca" must be set

`trusted_ca` - (Optional) Select/Add a Root CA Certificate object to associate with this Origin Pool for verification of server's certificate. See [ref](#ref) below for details.

`trusted_ca_url` - (Optional) Upload a Root CA Certificate specifically for this Origin Pool for verification of server's certificate (`String`).

### Server Validation Choice Volterra Trusted Ca

Perform origin server verification using F5XC Default Root CA Certificate.

### Service Info Service Selector

discovery has to happen. This implicit label is added to service_selector.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Sni Choice Disable Sni

Do not use SNI..

### Sni Choice Use Host Header As Sni

Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied..

### Subset Choice Disable Subsets

Subset load balancing is disabled. All eligible origin servers will be considered for load balancing..

### Subset Choice Enable Subsets

Subset load balancing is enabled. Based on route, subset of origin servers will be considered for load balancing..

`endpoint_subsets` - (Required) List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class.. See [Enable Subsets Endpoint Subsets ](#enable-subsets-endpoint-subsets) below for details.

###### One of the arguments from this list "any_endpoint, default_subset, fail_request" must be set

`any_endpoint` - (Optional) Select any origin server from available healthy origin servers in this pool (`Bool`).

`default_subset` - (Optional) Use the default subset provided here. Select endpoints matching default subset.. See [Fallback Policy Choice Default Subset ](#fallback-policy-choice-default-subset) below for details.

`fail_request` - (Optional) Request will be failed and error returned, as if cluster has no origin servers. (`Bool`).

### Tls Certificates Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Tls Choice Use Tls

x-displayName: "Enable".

###### One of the arguments from this list "use_mtls_obj, no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Upload a client authentication certificate specifically for this Origin Pool". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`use_mtls_obj` - (Optional) x-displayName: "Select/add a TLS Certificate object for client authentication". See [ref](#ref) below for details.

###### One of the arguments from this list "volterra_trusted_ca, use_server_verification, skip_server_verification" must be set

`skip_server_verification` - (Optional) Skip origin server verification (`Bool`).

`use_server_verification` - (Optional) Perform origin server verification using the provided Root CA Certificate. See [Server Validation Choice Use Server Verification ](#server-validation-choice-use-server-verification) below for details.

`volterra_trusted_ca` - (Optional) Perform origin server verification using F5XC Default Root CA Certificate (`Bool`).

###### One of the arguments from this list "sni, use_host_header_as_sni, disable_sni" must be set

`disable_sni` - (Optional) Do not use SNI. (`Bool`).

`sni` - (Optional) SNI value to be used. (`String`).

`use_host_header_as_sni` - (Optional) Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied. (`Bool`).

`tls_config` - (Required) TLS parameters such as min/max TLS version and ciphers. See [Use Tls Tls Config ](#use-tls-tls-config) below for details.

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

Attribute Reference
-------------------

-	`id` - This is the id of the configured origin_pool.
