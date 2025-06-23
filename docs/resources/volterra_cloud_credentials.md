---

page_title: "Volterra: cloud_credentials"

description: "The cloud_credentials allows CRUD of Cloud Credentials resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_cloud_credentials
===================================

The Cloud Credentials allows CRUD of Cloud Credentials resource on Volterra SaaS

~> **Note:** Please refer to [Cloud Credentials API docs](https://docs.cloud.f5.com/docs-v2/api/cloud-credentials) to learn more

Example Usage
-------------

```hcl
resource "volterra_cloud_credentials" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "aws_assume_role aws_secret_key azure_client_secret azure_pfx_certificate gcp_cred_file" must be set

  azure_pfx_certificate {
    certificate_url = "value"

    client_id = "value"

    password {
      // One of the arguments from this list "blindfold_secret_info clear_secret_info" must be set

      blindfold_secret_info {
        decryption_provider = "value"

        location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

        store_provider = "value"
      }
    }

    subscription_id = "value"

    tenant_id = "value"
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

###### One of the arguments from this list "aws_assume_role, aws_secret_key, azure_client_secret, azure_pfx_certificate, gcp_cred_file" must be set

`aws_assume_role` - (Optional) F5XC will assume role designated by customer. See [Cloud Aws Assume Role ](#cloud-aws-assume-role) below for details.

`aws_secret_key` - (Optional) AWS authentication using access keys. See [Cloud Aws Secret Key ](#cloud-aws-secret-key) below for details.

`azure_client_secret` - (Optional) Azure authentication using a service principal account with client secret. See [Cloud Azure Client Secret ](#cloud-azure-client-secret) below for details.

`azure_pfx_certificate` - (Optional) Azure authentication using a service principal account with client certificate. See [Cloud Azure Pfx Certificate ](#cloud-azure-pfx-certificate) below for details.

`gcp_cred_file` - (Optional) Google authentication using content of Google Credentials File. See [Cloud Gcp Cred File ](#cloud-gcp-cred-file) below for details.

### Aws Secret Key Secret Key

Secret Access Key for your AWS account.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Azure Client Secret Client Secret

Client Secret (alias password) for your Azure service principal.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Azure Pfx Certificate Password

Password for your '.p12' or '.pfx' file whose certificate is linked to service principal object.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Cloud Aws Assume Role

F5XC will assume role designated by customer.

`duration_seconds` - (Optional) The duration, in seconds of the role session. (`Int`).

`role_arn` - (Required) IAM Role ARN to assume the role (`String`).

`session_name` - (Required) be used for deploy, monitor from F5XC console (`String`).

`session_tags` - (Optional) Session tags are key-value pair attributes that you pass when you assume an IAM role (`String`).

### Cloud Aws Secret Key

AWS authentication using access keys.

`access_key` - (Required) Access key ID for your AWS account (`String`).

`secret_key` - (Required) Secret Access Key for your AWS account. See [Aws Secret Key Secret Key ](#aws-secret-key-secret-key) below for details.

### Cloud Azure Client Secret

Azure authentication using a service principal account with client secret.

`client_id` - (Required) Client ID for your Azure service principal (`String`).

`client_secret` - (Required) Client Secret (alias password) for your Azure service principal. See [Azure Client Secret Client Secret ](#azure-client-secret-client-secret) below for details.

`subscription_id` - (Required) Subscription ID for your Azure service principal (`String`).

`tenant_id` - (Required) Tenant ID for your Azure service principal (`String`).

### Cloud Azure Pfx Certificate

Azure authentication using a service principal account with client certificate.

`certificate_url` - (Required) Here <Base64 of certificate> is base64 of '.pfx' or '.p12' binary file (`String`).

`client_id` - (Required) Client ID for your Azure service principal (`String`).

`password` - (Required) Password for your '.p12' or '.pfx' file whose certificate is linked to service principal object. See [Azure Pfx Certificate Password ](#azure-pfx-certificate-password) below for details.

`subscription_id` - (Required) Subscription ID for your Azure service principal (`String`).

`tenant_id` - (Required) Tenant ID for your Azure service principal (`String`).

### Cloud Gcp Cred File

Google authentication using content of Google Credentials File.

`credential_file` - (Required) Content of Credential File for your GCP account. See [Gcp Cred File Credential File ](#gcp-cred-file-credential-file) below for details.

### Gcp Cred File Credential File

Content of Credential File for your GCP account.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

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

-	`id` - This is the id of the configured cloud_credentials.
