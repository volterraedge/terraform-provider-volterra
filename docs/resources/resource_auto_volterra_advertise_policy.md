---
page_title: "Volterra: advertise_policy"

description: "The advertise_policy allows CRUD of Advertise Policy resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_advertise_policy
==================================

The Advertise Policy allows CRUD of Advertise Policy resource on Volterra SaaS

~> **Note:** Please refer to [Advertise Policy API docs](https://volterra.io/docs/api/advertise-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_advertise_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
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

`port` - (Optional) Optional. port to advertise. If not specified, system will assume 80. (`Int`).

`protocol` - (Optional) Remove Hidden when more protocols are supported (`String`).

`public_ip` - (Optional) This field is mutually exclusive with where and address fields.. See [ref](#ref) below for details.

`skip_xff_append` - (Optional) If set, the loadbalancer will not append the remote address to the x-forwarded-for HTTP header. (`Bool`).

`tls_parameters` - (Optional) Optional. TLS parameters to use. If not specified, will take from Virtual Host configuration.. See [Tls Parameters ](#tls-parameters) below for details.

`where` - (Optional) * site Advertised on site local network in case of customer sites and Public network in case of regional sites. See [Where ](#where) below for details.

### Common Params

Common TLS parameters used in both upstream and downstream connections.

`cipher_suites` - (Optional)cipher_suites (`String`).

`maximum_protocol_version` - (Optional) Maximum TLS protocol version. (`String`).

`minimum_protocol_version` - (Optional) Minimum TLS protocol version. (`String`).

`tls_certificates` - (Optional) Set of TLS certificates. See [Tls Certificates ](#tls-certificates) below for details.

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

`validation_params` - (Optional) and list of Subject Alt Names for verification. See [Validation Params ](#validation-params) below for details.

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Tls Certificates

Set of TLS certificates.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Tls Parameters

Optional. TLS parameters to use. If not specified, will take from Virtual Host configuration..

`common_params` - (Optional) Common TLS parameters used in both upstream and downstream connections. See [Common Params ](#common-params) below for details.

`require_client_certificate` - (Optional) certificate. (`Bool`).

### Validation Params

and list of Subject Alt Names for verification.

`skip_hostname_verification` - (Optional) is not matched to the connecting hostname (`Bool`).

`trusted_ca_url` - (Optional) The URL for a trust store (`String`).

`use_volterra_trusted_ca_url` - (Optional) Ignore the trusted CA URL and use the volterra trusted CA URL from the global config for verification. (`Bool`).

`verify_subject_alt_names` - (Optional) the hostname of the peer will be used for matching against SAN/CN of peer's certificate (`String`).

### Where

-	site Advertised on site local network in case of customer sites and Public network in case of regional sites.

`site` - (Optional) Direct reference to site object. See [Site ](#site) below for details.

`virtual_network` - (Optional) Direct reference to virtual network object. See [Virtual Network ](#virtual-network) below for details.

`virtual_site` - (Optional) Direct reference to virtual site object. See [Virtual Site ](#virtual-site) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured advertise_policy.
