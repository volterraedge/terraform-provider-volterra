











---
page_title: "Volterra: cluster"
description: "The cluster allows CRUD of Cluster  resource on Volterra SaaS"
---
# Resource volterra_cluster

The Cluster  allows CRUD of Cluster  resource on Volterra SaaS

~> **Note:** Please refer to [Cluster  API docs](https://docs.cloud.f5.com/docs-v2/api/cluster) to learn more

## Example Usage

```hcl
resource "volterra_cluster" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
}

```

## Argument Reference

### Metadata Argument Reference
`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).


`description` - (Optional) Human readable description for the object (`String`).


`disable` - (Optional) A value of true will administratively disable the object (`Bool`).


`labels` - (Optional) by selector expression (`String`).


`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).


`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).



### Spec Argument Reference

`circuit_breaker` - (Optional) allows to apply back pressure on downstream quickly.. See [Circuit Breaker ](#circuit-breaker) below for details.








`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2 seconds (`Int`).



`default_subset` - (Optional) which gets used when route specifies no metadata or no subset matching the metadata exists. (`String`).


`endpoint_selection` - (Optional) Policy for selection of endpoints from local site or remote site or both (`String`).



`endpoint_subsets` - (Optional). See [Endpoint Subsets ](#endpoint-subsets) below for details.




`endpoints` - (Optional) List of references to all endpoint objects that belong to this cluster.. See [ref](#ref) below for details.

`fallback_policy` - (Optional) metadata defined as default_set (`String`).



`header_transformation_type` - (Optional) Settings to normalize the headers of upstream requests.. See [Header Transformation Type ](#header-transformation-type) below for details.(Deprecated)




		




		




		




		





`health_checks` - (Optional) List of references to healthcheck object for this cluster.. See [ref](#ref) below for details.


`http_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 5 minutes. (`Int`).




`auto_http_config` - (Optional) and will use whichever protocol is negotiated by ALPN with the upstream. (`Bool`).


`http1_config` - (Optional) Enable HTTP/1.1 for upstream connections. See [Http Protocol Type Http1 Config ](#http-protocol-type-http1-config) below for details.
		


		

















`http2_options` - (Optional) Enable HTTP/2 for upstream connections. See [Http Protocol Type Http2 Options ](#http-protocol-type-http2-options) below for details.
		





`loadbalancer_algorithm` - (Optional) loadbalancer_algorithm to determine which host is selected. (`String`).



`outlier_detection` - (Optional) healthy load balancing set. Outlier detection is a form of passive health checking.. See [Outlier Detection ](#outlier-detection) below for details.









`no_panic_threshold` - (Optional) Disable panic threshold. Only healthy endpoints are considered for loadbalancing. (`Bool`).


`panic_threshold` - (Optional) all endpoints will be considered for loadbalancing ignoring its health status. (`Int`).





`disable_proxy_protocol` - (Optional) Disable Proxy Protocol for upstream connections (`Bool`).


`proxy_protocol_v1` - (Optional) Enable Proxy Protocol V1 for upstream connections (`Bool`).


`proxy_protocol_v2` - (Optional) Enable Proxy Protocol V2 for upstream connections (`Bool`).




`tls_parameters` - (Optional) TLS parameters to access upstream endpoints for this cluster. See [Tls Parameters ](#tls-parameters) below for details.




		




		







		





		






		






		





		










		





		






		





		




		




		


		








		







		






		









		






		














### Circuit Breaker 

 allows to apply back pressure on downstream quickly..

`connection_limit` - (Optional) Remove endpoint out of load balancing decision, if number of connections reach connection limit. (`Int`).

`max_requests` - (Optional) Remove endpoint out of load balancing decision, if requests exceed this count. (`Int`).

`pending_requests` - (Optional) Remove endpoint out of load balancing decision, if pending request reach  pending_request. (`Int`).

`priority` - (Optional) matched with priority of CircuitBreaker to select the CircuitBreaker (`String`).

`retries` - (Optional) Remove endpoint out of load balancing decision, if retries for request exceed this count. (`Int`).



### Endpoint Subsets 

.

`keys` - (Required) List of keys that define a cluster subset class. (`String`).



### Header Transformation Type 

 Settings to normalize the headers of upstream requests..



###### One of the arguments from this list "default_header_transformation, proper_case_header_transformation, preserve_case_header_transformation, legacy_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).


`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).


`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).


`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).




### Outlier Detection 

 healthy load balancing set. Outlier detection is a form of passive health checking..

`base_ejection_time` - (Optional) Defaults to 30000ms or 30s. Specified in milliseconds. (`Int`).

`consecutive_5xx` - (Optional) a consecutive 5xx ejection occurs. Defaults to 5. (`Int`).

`consecutive_gateway_failure` - (Optional) before a consecutive gateway failure ejection occurs. Defaults to 5. (`Int`).

`interval` - (Optional) to 10000ms or 10s. Specified in milliseconds. (`Int`).

`max_ejection_percent` - (Optional) detection. Defaults to 10% but will eject at least one host regardless of the value. (`Int`).



### Tls Parameters 

 TLS parameters to access upstream endpoints for this cluster.



###### One of the arguments from this list "default_session_key_caching, disable_session_key_caching, max_session_keys" must be set

`default_session_key_caching` - (Optional) Default session key caching. Only one session key will be cached. (`Bool`).


`disable_session_key_caching` - (Optional) Disable session key caching. This will disable TLS session resumption. (`Bool`).


`max_session_keys` - (Optional) Number of session keys that are cached. (`Int`).




###### One of the arguments from this list "use_host_header_as_sni, disable_sni, sni" must be set

`disable_sni` - (Optional) Do not use SNI.. See [Sni Choice Disable Sni ](#sni-choice-disable-sni) below for details.


`sni` - (Optional) SNI value to be used. (`String`).


`use_host_header_as_sni` - (Optional) Use the host header as SNI. See [Sni Choice Use Host Header As Sni ](#sni-choice-use-host-header-as-sni) below for details.




###### One of the arguments from this list "common_params, cert_params" must be set

`cert_params` - (Optional) TLS certificate parameters for upstream connections. See [Tls Params Choice Cert Params ](#tls-params-choice-cert-params) below for details.


`common_params` - (Optional) Common TLS parameters used in upstream connections. See [Tls Params Choice Common Params ](#tls-params-choice-common-params) below for details.




### Cert Params Validation Params 

 and list of Subject Alt Names for verification.

`skip_hostname_verification` - (Optional) is not matched to the connecting hostname (`Bool`).



###### One of the arguments from this list "trusted_ca_url, trusted_ca" must be set

`trusted_ca` - (Optional) Root CA Certificate. See [Trusted Ca Choice Trusted Ca ](#trusted-ca-choice-trusted-ca) below for details.


`trusted_ca_url` - (Optional) Inline Root CA Certificate (`String`).


`use_volterra_trusted_ca_url` - (Optional) Use the F5XC default Root CA URL from the global config for hostname verification. (`Bool`).(Deprecated)

`verify_subject_alt_names` - (Optional) the hostname of the peer will be used for matching against SAN/CN of peer's certificate (`String`).



### Common Params Tls Certificates 

 Set of TLS certificates.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).




###### One of the arguments from this list "disable_ocsp_stapling, custom_hash_algorithms, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.


`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.


`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.


`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.



### Common Params Validation Params 

 and list of Subject Alt Names for verification.

`skip_hostname_verification` - (Optional) is not matched to the connecting hostname (`Bool`).



###### One of the arguments from this list "trusted_ca_url, trusted_ca" must be set

`trusted_ca` - (Optional) Root CA Certificate. See [Trusted Ca Choice Trusted Ca ](#trusted-ca-choice-trusted-ca) below for details.


`trusted_ca_url` - (Optional) Inline Root CA Certificate (`String`).


`use_volterra_trusted_ca_url` - (Optional) Use the F5XC default Root CA URL from the global config for hostname verification. (`Bool`).(Deprecated)

`verify_subject_alt_names` - (Optional) the hostname of the peer will be used for matching against SAN/CN of peer's certificate (`String`).



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



###### One of the arguments from this list "legacy_header_transformation, default_header_transformation, proper_case_header_transformation, preserve_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).


`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).


`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).


`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).




### Http Protocol Type Http1 Config 

 Enable HTTP/1.1 for upstream connections.

`header_transformation` - (Optional) the stateful formatter will take effect, and the stateless formatter will be disregarded.. See [Http1 Config Header Transformation ](#http1-config-header-transformation) below for details.



### Http Protocol Type Http2 Options 

 Enable HTTP/2 for upstream connections.

`enabled` - (Optional) Enable/disable HTTP2 Protocol for upstream connections (`Bool`).



### Max Session Keys Type Default Session Key Caching 

 Default session key caching. Only one session key will be cached..



### Max Session Keys Type Disable Session Key Caching 

 Disable session key caching. This will disable TLS session resumption..



### Ocsp Stapling Choice Custom Hash Algorithms 

 Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).



### Ocsp Stapling Choice Disable Ocsp Stapling 

 This is the default behavior if no choice is selected..



### Ocsp Stapling Choice Use System Defaults 

 F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..



### Private Key Blindfold Secret Info Internal 

 Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).



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



### Sni Choice Disable Sni 

 Do not use SNI..



### Sni Choice Use Host Header As Sni 

 Use the host header as SNI.



### Tls Certificates Private Key 

 TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)



###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.


`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.


`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)


`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)




### Tls Params Choice Cert Params 

 TLS certificate parameters for upstream connections.

`certificates` - (Required) Client TLS Certificate required for mTLS authentication. See [ref](#ref) below for details.

`cipher_suites` - (Optional) will be used. (`String`).

`maximum_protocol_version` - (Optional) Maximum TLS protocol version. (`String`).

`minimum_protocol_version` - (Optional) Minimum TLS protocol version. (`String`).

`validation_params` - (Optional) and list of Subject Alt Names for verification. See [Cert Params Validation Params ](#cert-params-validation-params) below for details.



### Tls Params Choice Common Params 

 Common TLS parameters used in upstream connections.

`cipher_suites` - (Optional) will be used. (`String`).

`maximum_protocol_version` - (Optional) Maximum TLS protocol version. (`String`).

`minimum_protocol_version` - (Optional) Minimum TLS protocol version. (`String`).

`tls_certificates` - (Optional) Set of TLS certificates. See [Common Params Tls Certificates ](#common-params-tls-certificates) below for details.

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).(Deprecated)

`validation_params` - (Optional) and list of Subject Alt Names for verification. See [Common Params Validation Params ](#common-params-validation-params) below for details.



### Trusted Ca Choice Trusted Ca 

 Root CA Certificate.

`trusted_ca_list` - (Optional) Reference to Root CA Certificate. See [ref](#ref) below for details.



## Attribute Reference

* `id` - This is the id of the configured cluster.

