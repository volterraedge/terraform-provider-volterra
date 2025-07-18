---

page_title: "Volterra: certificate"

description: "The certificate allows CRUD of Certificate resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_certificate
=============================

The Certificate allows CRUD of Certificate resource on Volterra SaaS

~> **Note:** Please refer to [Certificate API docs](https://docs.cloud.f5.com/docs-v2/api/certificate) to learn more

Example Usage
-------------

```hcl
resource "volterra_certificate" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  certificate_url = ["value"]

  private_key {
    // One of the arguments from this list "blindfold_secret_info clear_secret_info" must be set

    blindfold_secret_info {
      decryption_provider = "value"

      location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

      store_provider = "value"
    }
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

`certificate_chain` - (Optional) Intermediate Certificate Chain is used to establish the chain of trust. Intermediate Certificate Chain contains the list of intermediate certificates, excluding the end-entity certificate.. See [ref](#ref) below for details.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected. (`Bool`).

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order. (`Bool`).

`private_key` - (Required) Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. Key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Private Key

Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. Key has to match the accompanying certificate..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

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

Attribute Reference
-------------------

-	`id` - This is the id of the configured certificate.
