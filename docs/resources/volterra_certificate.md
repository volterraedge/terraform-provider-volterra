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

`certificate_url` - (Optional) Required unless this is an Automatic type certificate (certificate URL will be added once provisioned). (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected. (`Bool`).

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order. (`Bool`).

`private_key` - (Optional) Required unless this is an Automatic type certificate (private key will be generated).. See [Private Key ](#private-key) below for details.

###### One of the arguments from this list "auto_cert, manual" can be set

`auto_cert` - (Optional) Automated certificate issuance and renewal using the ACME protocol.. See [Provider Auto Cert ](#provider-auto-cert) below for details.

`manual` - (Optional) Certificate PEM and private key are uploaded by the user. (`Bool`).

### Private Key

Required unless this is an Automatic type certificate (private key will be generated)..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, none, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`none` - (Optional) No private key will be stored or used. Choose this only when you need a public certificate for verification (e.g., validating SAML/JWT signatures). Not suitable for TLS or any use case that requires presenting a certificate, those will fail without a private key.. See [Secret Info Oneof None ](#secret-info-oneof-none) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Challenge Choice Dns 01

domain types. Domains reported as already validated are skipped..

###### One of the arguments from this list "cname_delegation, delegated" must be set

`cname_delegation` - (Optional) records are created automatically at the CNAME target. (`Bool`).

`delegated` - (Optional) validation. Each zone must have 'Allow Application Managed Records' enabled. (`Bool`).

### Challenge Choice Http 01

non-wildcard domains..

`virtual_host_ref` - (Required) The route is removed after validation completes.. See [ref](#ref) below for details.

### Challenge Choice Pre Validated

renewal. If any domain requires validation, issuance cannot continue..

### Dns Provider Choice Cname Delegation

records are created automatically at the CNAME target..

### Dns Provider Choice Delegated

validation. Each zone must have 'Allow Application Managed Records' enabled..

### Issuer Choice Custom Issuer

Use a custom ACME issuer..

`issuer_ref` - (Required) Select an existing ACME issuer or create a new one.. See [ref](#ref) below for details.

### Issuer Choice Default

Use the built-in Let's Encrypt issuer managed by F5 Distributed Cloud..

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Provider Auto Cert

Automated certificate issuance and renewal using the ACME protocol..

###### One of the arguments from this list "dns_01, http_01, pre_validated" can be set

`dns_01` - (Optional) domain types. Domains reported as already validated are skipped.. See [Challenge Choice Dns 01 ](#challenge-choice-dns-01) below for details.

`http_01` - (Optional) non-wildcard domains.. See [Challenge Choice Http 01 ](#challenge-choice-http-01) below for details.(Deprecated)

`pre_validated` - (Optional) renewal. If any domain requires validation, issuance cannot continue.. See [Challenge Choice Pre Validated ](#challenge-choice-pre-validated) below for details.

`domains` - (Required) in the ACME certificate order. At least one domain is required. (`String`).

###### One of the arguments from this list "custom_issuer, default" can be set

`custom_issuer` - (Optional) Use a custom ACME issuer.. See [Issuer Choice Custom Issuer ](#issuer-choice-custom-issuer) below for details.

`default` - (Optional) Use the built-in Let's Encrypt issuer managed by F5 Distributed Cloud. (`Bool`).

`key_algorithm` - (Optional) Preferred key algorithm for certificate key pair generation. (`String`).

`renew_before_days` - (Optional) will default to two-thirds through its actual lifetime. (`Int`).

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

### Secret Info Oneof None

No private key will be stored or used. Choose this only when you need a public certificate for verification (e.g., validating SAML/JWT signatures). Not suitable for TLS or any use case that requires presenting a certificate, those will fail without a private key..

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

Attribute Reference
-------------------

*   `id` - This is the id of the configured certificate.
