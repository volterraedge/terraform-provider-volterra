











---
page_title: "Volterra: container_registry"
description: "The container_registry allows CRUD of Container Registry  resource on Volterra SaaS"
---
# Resource volterra_container_registry

The Container Registry  allows CRUD of Container Registry  resource on Volterra SaaS

~> **Note:** Please refer to [Container Registry  API docs](https://docs.cloud.f5.com/docs-v2/api/container-registry) to learn more

## Example Usage

```hcl
resource "volterra_container_registry" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  password {
    blindfold_secret_info_internal {
      decryption_provider = "value"

      location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

      store_provider = "value"
    }

    secret_encoding_type = "secret_encoding_type"

    // One of the arguments from this list "wingman_secret_info blindfold_secret_info vault_secret_info clear_secret_info" must be set

    vault_secret_info {
      key = "key_pem"

      location = "v1/data/vhost_key"

      provider = "vault-vh-provider"

      secret_encoding = "secret_encoding"

      version = "1"
    }
  }

  registry  = ["quay.io"]
  user_name = ["joe"]
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

`email` - (Optional) Email used for the registry (`String`).



`password` - (Required) Password for the above username. See [Password ](#password) below for details.


		








		







		






		









		






`registry` - (Required) Fully qualified name of the registry login server (`String`).



`user_name` - (Required) Username used to access the registry (`String`).



### Password 

 Password for the above username.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Password Blindfold Secret Info Internal ](#password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)



###### One of the arguments from this list "vault_secret_info, clear_secret_info, wingman_secret_info, blindfold_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.


`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.


`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)


`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)




### Password Blindfold Secret Info Internal 

 Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).



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



## Attribute Reference

* `id` - This is the id of the configured container_registry.

