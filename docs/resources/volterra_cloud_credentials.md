---

page_title: "Volterra: cloud_credentials"

description: "The cloud_credentials allows CRUD of Cloud Credentials resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_cloud_credentials
===================================

The Cloud Credentials allows CRUD of Cloud Credentials resource on Volterra SaaS

~> **Note:** Please refer to [Cloud Credentials API docs](https://volterra.io/docs/api/cloud-credentials) to learn more

Example Usage
-------------

```hcl
resource "volterra_cloud_credentials" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "gcp_cred_file aws_secret_key azure_pfx_certificate azure_client_secret" must be set

  aws_secret_key {
    access_key = "access_key"

    secret_key {
      blindfold_secret_info_internal {
        decryption_provider = "decryption_provider"
        location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
        store_provider      = "store_provider"
      }

      secret_encoding_type = "secret_encoding_type"

      // One of the arguments from this list "vault_secret_info clear_secret_info wingman_secret_info blindfold_secret_info" must be set

      blindfold_secret_info {
        decryption_provider = "decryption_provider"
        location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
        store_provider      = "store_provider"
      }
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

`aws_secret_key` - (Optional) AWS authentication using access keys. See [Aws Secret Key ](#aws-secret-key) below for details.

`azure_client_secret` - (Optional) Azure authentication using a service principal account with client secret. See [Azure Client Secret ](#azure-client-secret) below for details.

`azure_pfx_certificate` - (Optional) Azure authentication using a service principal account with client certificate. See [Azure Pfx Certificate ](#azure-pfx-certificate) below for details.

`gcp_cred_file` - (Optional) Google authentication using content of Google Credentials File. See [Gcp Cred File ](#gcp-cred-file) below for details.

### Aws Secret Key

AWS authentication using access keys.

`access_key` - (Required) Access key ID for your AWS account (`String`).

`secret_key` - (Required) Secret Access Key for your AWS account. See [Secret Key ](#secret-key) below for details.

### Azure Client Secret

Azure authentication using a service principal account with client secret.

`client_id` - (Required) Client ID for your Azure service principal (`String`).

`client_secret` - (Required) Client Secret (alias password) for your Azure service principal. See [Client Secret ](#client-secret) below for details.

`subscription_id` - (Required) Subscription ID for your Azure service principal (`String`).

`tenant_id` - (Required) Tenant ID for your Azure service principal (`String`).

### Azure Pfx Certificate

Azure authentication using a service principal account with client certificate.

`certificate_url` - (Required) Here <Base64 of certificate> is base64 of '.pfx' or '.p12' binary file (`String`).

`client_id` - (Required) Client ID for your Azure service principal (`String`).

`password` - (Required) Password for your '.p12' or '.pfx' file whose certificate is linked to service principal object. See [Password ](#password) below for details.

`subscription_id` - (Required) Subscription ID for your Azure service principal (`String`).

`tenant_id` - (Required) Tenant ID for your Azure service principal (`String`).

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

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted .

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Client Secret

Client Secret (alias password) for your Azure service principal.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Credential File

Content of Credential File for your GCP account.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Gcp Cred File

Google authentication using content of Google Credentials File.

`credential_file` - (Required) Content of Credential File for your GCP account. See [Credential File ](#credential-file) below for details.

### Password

Password for your '.p12' or '.pfx' file whose certificate is linked to service principal object.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Secret Key

Secret Access Key for your AWS account.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Wingman Secret Info

Secret is given as bootstrap secret in Volterra Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured cloud_credentials.
