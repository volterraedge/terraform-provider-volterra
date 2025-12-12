---

page_title: "Volterra: origin_pool"

description: "The origin_pool allows CRUD of Origin Pool resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_origin_pool
=============================

The Origin Pool allows CRUD of Origin Pool resource on Volterra SaaS

~> **Note:** Please refer to [Origin Pool API docs](https://docs.cloud.f5.com/docs-v2/api/views-origin-pool) to learn more

Example Usage
-------------

```hcl
resource "volterra_origin_pool" "example" {
  name                   = "acmecorp-web"
  namespace              = "staging"
  endpoint_selection     = ["endpoint_selection"]
  loadbalancer_algorithm = ["loadbalancer_algorithm"]

  origin_servers {
    // One of the arguments from this list "cbip_service consul_service custom_endpoint_object k8s_service private_ip private_name public_ip public_name vn_private_ip vn_private_name" must be set

    private_name {
      dns_name = "value"

      // One of the arguments from this list "inside_network outside_network segment" must be set

      inside_network = true
      refresh_interval = "20"
      site_locator {
        // One of the arguments from this list "site virtual_site" must be set

        site {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
      }
      snat_pool {
        // One of the arguments from this list "no_snat_pool snat_pool" can be set

        no_snat_pool = true
      }
    }

    labels = {
      "key1" = "value1"
    }
  }

  // One of the arguments from this list "automatic_port lb_port port" must be set

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

###### One of the arguments from this list "health_check_port, same_as_endpoint_port" can be set

`health_check_port` - (Optional) Port used for performing health check (`Int`).

`same_as_endpoint_port` - (Optional) Health check is performed on endpoint port itself (`Bool`).

`healthcheck` - (Optional) Reference to healthcheck configuration objects. See [ref](#ref) below for details.

`loadbalancer_algorithm` - (Required) loadbalancer_algorithm to determine which host is selected. (`String`).

`origin_servers` - (Required) List of origin servers in this pool. See [Origin Servers ](#origin-servers) below for details.

###### One of the arguments from this list "automatic_port, lb_port, port" must be set

`automatic_port` - (Optional) For other origin server types, port will be automatically set as 443 if TLS is enabled at Origin Pool and 80 if TLS is disabled (`Bool`).

`lb_port` - (Optional) Endpoint port is selected based on loadbalancer port (`Bool`).

`port` - (Optional) Endpoint service is available on this port (`Int`).

###### One of the arguments from this list "no_tls, use_tls" must be set

`no_tls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_tls` - (Optional) x-displayName: "Enable". See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

`upstream_conn_pool_reuse_type` - (Optional) This configuration choice is for HTTP(S) LB only.. See [Upstream Conn Pool Reuse Type ](#upstream-conn-pool-reuse-type) below for details.

### Advanced Options

Advanced options configuration like timeouts, circuit breaker, subset load balancing.

###### One of the arguments from this list "circuit_breaker, default_circuit_breaker, disable_circuit_breaker" must be set

`circuit_breaker` - (Optional) allows to apply back pressure on downstream quickly.. See [Circuit Breaker Choice Circuit Breaker ](#circuit-breaker-choice-circuit-breaker) below for details.

`default_circuit_breaker` - (Optional) requests are 1024 and the default value for retries is 3 (`Bool`).

`disable_circuit_breaker` - (Optional) Circuit Breaker is disabled (`Bool`).

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2 seconds (`Int`).

`header_transformation_type` - (Optional) Settings to normalize the headers of upstream requests.. See [Advanced Options Header Transformation Type ](#advanced-options-header-transformation-type) below for details.(Deprecated)

`http_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 5 minutes. (`Int`).

###### One of the arguments from this list "auto_http_config, http1_config, http2_options" must be set

`auto_http_config` - (Optional) and will use whichever protocol is negotiated by ALPN with the upstream. (`Bool`).

`http1_config` - (Optional) Enable HTTP/1.1 for upstream connections. See [Http Protocol Type Http1 Config ](#http-protocol-type-http1-config) below for details.

`http2_options` - (Optional) Enable HTTP/2 for upstream connections.. See [Http Protocol Type Http2 Options ](#http-protocol-type-http2-options) below for details.

###### One of the arguments from this list "disable_lb_source_ip_persistance, enable_lb_source_ip_persistance" can be set

`disable_lb_source_ip_persistance` - (Optional) Disable LB source IP persistence (`Bool`).

`enable_lb_source_ip_persistance` - (Optional) Enable LB source IP persistence (`Bool`).

###### One of the arguments from this list "disable_outlier_detection, outlier_detection" must be set

`disable_outlier_detection` - (Optional) Outlier detection is disabled (`Bool`).

`outlier_detection` - (Optional) healthy load balancing set. Outlier detection is a form of passive health checking.. See [Outlier Detection Choice Outlier Detection ](#outlier-detection-choice-outlier-detection) below for details.

###### One of the arguments from this list "no_panic_threshold, panic_threshold" must be set

`no_panic_threshold` - (Optional) Disable panic threshold. Only healthy endpoints are considered for load balancing. (`Bool`).

`panic_threshold` - (Optional) all endpoints will be considered for load balancing ignoring its health status. (`Int`).

###### One of the arguments from this list "disable_proxy_protocol, proxy_protocol_v1, proxy_protocol_v2" can be set

`disable_proxy_protocol` - (Optional) Disable Proxy Protocol for upstream connections (`Bool`).

`proxy_protocol_v1` - (Optional) Enable Proxy Protocol Version 1 for upstream connections (`Bool`).

`proxy_protocol_v2` - (Optional) Enable Proxy Protocol Version 2 for upstream connections (`Bool`).

###### One of the arguments from this list "disable_subsets, enable_subsets" must be set

`disable_subsets` - (Optional) Subset load balancing is disabled. All eligible origin servers will be considered for load balancing. (`Bool`).

`enable_subsets` - (Optional) Subset load balancing is enabled. Based on route, subset of origin servers will be considered for load balancing.. See [Subset Choice Enable Subsets ](#subset-choice-enable-subsets) below for details.

### Origin Servers

List of origin servers in this pool.

###### One of the arguments from this list "cbip_service, consul_service, custom_endpoint_object, k8s_service, private_ip, private_name, public_ip, public_name, vn_private_ip, vn_private_name" must be set

`cbip_service` - (Optional) Specify origin server with cBIP service name. See [Choice Cbip Service ](#choice-cbip-service) below for details.

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

### Upstream Conn Pool Reuse Type

This configuration choice is for HTTP(S) LB only..

###### One of the arguments from this list "disable_conn_pool_reuse, enable_conn_pool_reuse" must be set

`disable_conn_pool_reuse` - (Optional) Open new upstream connection pool for every new downstream connection (`Bool`).

`enable_conn_pool_reuse` - (Optional) Reuse upstream connection pool for multiple downstream connections (`Bool`).

### Advanced Options Header Transformation Type

Settings to normalize the headers of upstream requests..

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

### Choice Cbip Service

Specify origin server with cBIP service name.

`service_name` - (Required) Name of the discovered Classic BIG-IP virtual server to be used as origin. (`String`).

### Choice Consul Service

Specify origin server with Hashi Corp Consul service name and site information.

###### One of the arguments from this list "inside_network, outside_network" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`service_name` - (Required) The format is servicename:cluster-id. (`String`).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Consul Service Site Locator ](#consul-service-site-locator) below for details.

`snat_pool` - (Optional) x-displayName: "SNAT Pool Configuration". See [Consul Service Snat Pool ](#consul-service-snat-pool) below for details.

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

###### One of the arguments from this list "inside_network, outside_network, vk8s_networks" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`vk8s_networks` - (Optional) origin server are on vK8s network on the site (`Bool`).

`protocol` - (Optional) Protocol to be used in the discovery. (`String`).

###### One of the arguments from this list "service_name, service_selector" must be set

`service_name` - (Optional) Both namespace and cluster-id are optional. (`String`).

`service_selector` - (Optional) discovery has to happen. This implicit label is added to service_selector. See [Service Info Service Selector ](#service-info-service-selector) below for details.(Deprecated)

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [K8s Service Site Locator ](#k8s-service-site-locator) below for details.

`snat_pool` - (Optional) x-displayName: "SNAT Pool Configuration". See [K8s Service Snat Pool ](#k8s-service-snat-pool) below for details.

### Choice Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Choice Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### Choice Private Ip

Specify origin server with private or public IP address and site information.

###### One of the arguments from this list "inside_network, outside_network, segment" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`segment` - (Optional) Segment where this origin server is located. See [ref](#ref) below for details.

###### One of the arguments from this list "ip, ipv6" must be set

`ip` - (Optional) Private IPV4 address (`String`).

`ipv6` - (Optional) Private IPV6 address (`String`).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Private Ip Site Locator ](#private-ip-site-locator) below for details.

`snat_pool` - (Optional) x-displayName: "SNAT Pool Configuration". See [Private Ip Snat Pool ](#private-ip-snat-pool) below for details.

### Choice Private Name

Specify origin server with private or public DNS name and site information.

`dns_name` - (Required) DNS Name (`String`).

###### One of the arguments from this list "inside_network, outside_network, segment" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`segment` - (Optional) Segment where this origin server is located. See [ref](#ref) below for details.

`refresh_interval` - (Optional) Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767 (`Int`).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Private Name Site Locator ](#private-name-site-locator) below for details.

`snat_pool` - (Optional) x-displayName: "SNAT Pool Configuration". See [Private Name Snat Pool ](#private-name-snat-pool) below for details.

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

### Consul Service Snat Pool

x-displayName: "SNAT Pool Configuration".

###### One of the arguments from this list "no_snat_pool, snat_pool" can be set

`no_snat_pool` - (Optional) No configured SNAT Pool to reach Origin Server (`Bool`).

`snat_pool` - (Optional) Configure SNAT Pool to reach Origin Server. See [Snat Pool Choice Snat Pool ](#snat-pool-choice-snat-pool) below for details.

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

### Header Transformation Choice Legacy Header Transformation

Use old header transformation if configured earlier.

### Header Transformation Choice Preserve Case Header Transformation

Preserves the original case of headers without any modifications..

### Header Transformation Choice Proper Case Header Transformation

For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are”.

### Http1 Config Header Transformation

the stateful formatter will take effect, and the stateless formatter will be disregarded..

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

### Http Protocol Type Auto Http Config

and will use whichever protocol is negotiated by ALPN with the upstream..

### Http Protocol Type Http1 Config

Enable HTTP/1.1 for upstream connections.

`header_transformation` - (Optional) the stateful formatter will take effect, and the stateless formatter will be disregarded.. See [Http1 Config Header Transformation ](#http1-config-header-transformation) below for details.

### Http Protocol Type Http2 Options

Enable HTTP/2 for upstream connections..

`enabled` - (Optional) Enable/disable HTTP2 Protocol for upstream connections (`Bool`).

### K8s Service Site Locator

Site or Virtual site where this origin server is located.

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### K8s Service Snat Pool

x-displayName: "SNAT Pool Configuration".

###### One of the arguments from this list "no_snat_pool, snat_pool" can be set

`no_snat_pool` - (Optional) No configured SNAT Pool to reach Origin Server (`Bool`).

`snat_pool` - (Optional) Configure SNAT Pool to reach Origin Server. See [Snat Pool Choice Snat Pool ](#snat-pool-choice-snat-pool) below for details.

### Lb Source Ip Persistance Choice Disable Lb Source Ip Persistance

Disable LB source IP persistence.

### Lb Source Ip Persistance Choice Enable Lb Source Ip Persistance

Enable LB source IP persistence.

### Map Downstream To Upstream Conn Pool Type Disable Conn Pool Reuse

Open new upstream connection pool for every new downstream connection.

### Map Downstream To Upstream Conn Pool Type Enable Conn Pool Reuse

Reuse upstream connection pool for multiple downstream connections.

### Max Session Keys Type Default Session Key Caching

Default session key caching. Only one session key will be cached..

### Max Session Keys Type Disable Session Key Caching

Disable session key caching. This will disable TLS session resumption..

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

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Private Ip Snat Pool

x-displayName: "SNAT Pool Configuration".

###### One of the arguments from this list "no_snat_pool, snat_pool" can be set

`no_snat_pool` - (Optional) No configured SNAT Pool to reach Origin Server (`Bool`).

`snat_pool` - (Optional) Configure SNAT Pool to reach Origin Server. See [Snat Pool Choice Snat Pool ](#snat-pool-choice-snat-pool) below for details.

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

### Private Name Snat Pool

x-displayName: "SNAT Pool Configuration".

###### One of the arguments from this list "no_snat_pool, snat_pool" can be set

`no_snat_pool` - (Optional) No configured SNAT Pool to reach Origin Server (`Bool`).

`snat_pool` - (Optional) Configure SNAT Pool to reach Origin Server. See [Snat Pool Choice Snat Pool ](#snat-pool-choice-snat-pool) below for details.

### Proxy Protocol Choice Disable Proxy Protocol

Disable Proxy Protocol for upstream connections.

### Proxy Protocol Choice Proxy Protocol V1

Enable Proxy Protocol Version 1 for upstream connections.

### Proxy Protocol Choice Proxy Protocol V2

Enable Proxy Protocol Version 2 for upstream connections.

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

###### One of the arguments from this list "trusted_ca, trusted_ca_url" must be set

`trusted_ca` - (Optional) Select/Add a Root CA Certificate object to associate with this Origin Pool for verification of server's certificate. See [ref](#ref) below for details.

`trusted_ca_url` - (Optional) Upload a Root CA Certificate specifically for this Origin Pool for verification of server's certificate (`String`).

### Server Validation Choice Volterra Trusted Ca

Perform origin server verification using F5XC Default Root CA Certificate.

### Service Info Service Selector

discovery has to happen. This implicit label is added to service_selector.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Snat Pool Choice No Snat Pool

No configured SNAT Pool to reach Origin Server.

### Snat Pool Choice Snat Pool

Configure SNAT Pool to reach Origin Server.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

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

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Tls Choice Use Tls

x-displayName: "Enable".

###### One of the arguments from this list "default_session_key_caching, disable_session_key_caching, max_session_keys" must be set

`default_session_key_caching` - (Optional) Default session key caching. Only one session key will be cached. (`Bool`).

`disable_session_key_caching` - (Optional) Disable session key caching. This will disable TLS session resumption. (`Bool`).

`max_session_keys` - (Optional) Number of session keys that are cached. (`Int`).

###### One of the arguments from this list "no_mtls, use_mtls, use_mtls_obj" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Upload a client authentication certificate specifically for this Origin Pool". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`use_mtls_obj` - (Optional) x-displayName: "Select/add a TLS Certificate object for client authentication". See [ref](#ref) below for details.

###### One of the arguments from this list "skip_server_verification, use_server_verification, volterra_trusted_ca" must be set

`skip_server_verification` - (Optional) Skip origin server verification (`Bool`).

`use_server_verification` - (Optional) Perform origin server verification using the provided Root CA Certificate. See [Server Validation Choice Use Server Verification ](#server-validation-choice-use-server-verification) below for details.

`volterra_trusted_ca` - (Optional) Perform origin server verification using F5XC Default Root CA Certificate (`Bool`).

###### One of the arguments from this list "disable_sni, sni, use_host_header_as_sni" must be set

`disable_sni` - (Optional) Do not use SNI. (`Bool`).

`sni` - (Optional) SNI value to be used. (`String`).

`use_host_header_as_sni` - (Optional) Use the host header as SNI. The host header value is extracted after any configured rewrites have been applied. (`Bool`).

`tls_config` - (Required) TLS parameters such as min/max TLS version and ciphers. See [Use Tls Tls Config ](#use-tls-tls-config) below for details.

### Use Mtls Tls Certificates

mTLS Client Certificate.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Use Tls Tls Config

TLS parameters such as min/max TLS version and ciphers.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured origin_pool.
