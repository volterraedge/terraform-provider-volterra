---

page_title: "Volterra: advertise_policy"

description: "The advertise_policy allows CRUD of Advertise Policy resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_advertise_policy
==================================

The Advertise Policy allows CRUD of Advertise Policy resource on Volterra SaaS

~> **Note:** Please refer to [Advertise Policy API docs](https://docs.cloud.f5.com/docs/api/advertise-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_advertise_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "port port_ranges" must be set
  port = "80"
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

`address` - (Optional) If inside_vip/outside_vip is not configured in the site object, system use interface ip in the respected networks. (`String`).

`port` - (Optional) Port to advertise. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

`protocol` - (Optional) Protocol to advertise. (`String`).

`public_ip` - (Optional) This field is mutually exclusive with where and address fields.. See [ref](#ref) below for details.

`skip_xff_append` - (Optional) If set, the loadbalancer will not append the remote address to the x-forwarded-for HTTP header. (`Bool`).

`tls_parameters` - (Optional) Optional. TLS parameters to use. If not specified, will take from Virtual Host configuration.. See [Tls Parameters ](#tls-parameters) below for details.

`where` - (Optional) * site Advertised on site local network in case of customer sites and Public network in case of regional sites. See [Where ](#where) below for details.

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

### Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Disable Internet Vip

Do not enable advertise on external internet vip..

### Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Enable Internet Vip

Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site..

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Segment

Reference to Segment.

`ref` - (Required) A segment reference. See [ref](#ref) below for details.

### Segment Site

Reference to Segment object.

`segment` - (Required) Segment in the site. See [ref](#ref) below for details.

`site` - (Required) Reference to a site. See [ref](#ref) below for details.

### Segment Vsite

Reference to Segment in a virtual site.

`segment` - (Required) Segment in the virtual site. See [ref](#ref) below for details.

`vsite` - (Required) Reference to a virtual site. See [ref](#ref) below for details.

### Site

Direct reference to site object.

`disable_internet_vip` - (Optional) Do not enable advertise on external internet vip. (bool).

`enable_internet_vip` - (Optional) Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site. (bool).

`network_type` - (Optional) The type of network on the referred site (`String`).

`ref` - (Required) A site direct reference. See [ref](#ref) below for details.

### Tls Certificates

Set of TLS certificates.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Custom Hash Algorithms ](#custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Disable Ocsp Stapling ](#disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Use System Defaults ](#use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Tls Parameters

Optional. TLS parameters to use. If not specified, will take from Virtual Host configuration..

`common_params` - (Optional) Common TLS parameters used in both upstream and downstream connections. See [Common Params ](#common-params) below for details.

`crl` - (Optional) Used to ensure that the client presented certificate is not revoked as per the CRL. See [ref](#ref) below for details.

`require_client_certificate` - (Optional) certificate. (`Bool`).

`xfcc_header_elements` - (Optional) If none are defined, the header will not be added. (`List of Strings`).

### Trusted Ca

Trusted CA List.

`trusted_ca_list` - (Optional) Reference to Trusted CA List. See [ref](#ref) below for details.

### Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Validation Params

and list of Subject Alt Names for verification.

`skip_hostname_verification` - (Optional) is not matched to the connecting hostname (`Bool`).

`trusted_ca` - (Optional) Trusted CA List. See [Trusted Ca ](#trusted-ca) below for details.

`trusted_ca_url` - (Optional) Inline Trusted CA List (`String`).

`use_volterra_trusted_ca_url` - (Optional) Ignore the trusted CA URL and use the volterra trusted CA URL from the global config for verification. (`Bool`).

`verify_subject_alt_names` - (Optional) the hostname of the peer will be used for matching against SAN/CN of peer's certificate (`String`).

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Virtual Network

Direct reference to virtual network object.

`ref` - (Required) A virtual network direct reference. See [ref](#ref) below for details.

### Virtual Site

Direct reference to virtual site object.

`disable_internet_vip` - (Optional) Do not enable advertise on external internet vip. (bool).

`enable_internet_vip` - (Optional) Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site. (bool).

`network_type` - (Optional) The type of network on the referred virtual_site (`String`).

`ref` - (Required) A virtual_site direct reference. See [ref](#ref) below for details.

### Where

-	site Advertised on site local network in case of customer sites and Public network in case of regional sites.

`segment` - (Optional) Reference to Segment. See [Segment ](#segment) below for details.

`segment_site` - (Optional) Reference to Segment object. See [Segment Site ](#segment-site) below for details.

`segment_vsite` - (Optional) Reference to Segment in a virtual site. See [Segment Vsite ](#segment-vsite) below for details.

`site` - (Optional) Direct reference to site object. See [Site ](#site) below for details.

`virtual_network` - (Optional) Direct reference to virtual network object. See [Virtual Network ](#virtual-network) below for details.

`virtual_site` - (Optional) Direct reference to virtual site object. See [Virtual Site ](#virtual-site) below for details.

### Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured advertise_policy.
