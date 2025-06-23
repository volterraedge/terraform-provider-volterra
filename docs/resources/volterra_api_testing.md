---

page_title: "Volterra: api_testing"

description: "The api_testing allows CRUD of Api Testing resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_api_testing
=============================

The Api Testing allows CRUD of Api Testing resource on Volterra SaaS

~> **Note:** Please refer to [Api Testing API docs](https://docs.cloud.f5.com/docs-v2/api/api-testing) to learn more

Example Usage
-------------

```hcl
resource "volterra_api_testing" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  domains {
    allow_destructive_methods = true

    credentials {
      credential_name = "credential_name"

      // One of the arguments from this list "api_key basic_auth bearer_token login_endpoint" must be set

      bearer_token {
        token {
          // One of the arguments from this list "blindfold_secret_info clear_secret_info" must be set

          blindfold_secret_info {
            decryption_provider = "value"

            location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

            store_provider = "value"
          }
        }
      }

      // One of the arguments from this list "admin standard" must be set

      admin = true
    }

    domain = "domain"
  }

  // One of the arguments from this list "every_day every_month every_week" must be set

  every_week = true
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

`custom_header_value` - (Optional) Add x-f5-api-testing-identifier header value to prevent security flags on API testing traffic (`String`).

`domains` - (Required) Add and configure testing domains and credentials. See [Domains ](#domains) below for details.

###### One of the arguments from this list "every_day, every_month, every_week" must be set

`every_day` - (Optional) x-displayName: "Every Day" (`Bool`).

`every_month` - (Optional) x-displayName: "Every Month" (`Bool`).

`every_week` - (Optional) x-displayName: "Every Week" (`Bool`).

### Domains

Add and configure testing domains and credentials.

`allow_destructive_methods` - (Optional) Enable to allow API test to execute destructive methods. Be cautious as these can alter or delete data (`Bool`).

`credentials` - (Required) Add credentials for API testing to use in the selected environment.. See [Domains Credentials ](#domains-credentials) below for details.

`domain` - (Required) Add your testing environment domain. Be aware that running tests on a production domain can impact live applications, as API testing cannot distinguish between production and testing environments. (`String`).

### Api Key Value

x-displayName: "Value".

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Basic Auth Password

x-displayName: "Password".

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Bearer Token Token

x-displayName: "Token".

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Credentials Choice Api Key

x-displayName: "API Key".

`key` - (Required) x-displayName: "Key" (`String`).

`value` - (Required) x-displayName: "Value". See [Api Key Value ](#api-key-value) below for details.

### Credentials Choice Basic Auth

x-displayName: "Basic Authentication".

`password` - (Required) x-displayName: "Password". See [Basic Auth Password ](#basic-auth-password) below for details.

`user` - (Required) x-displayName: "User" (`String`).

### Credentials Choice Bearer Token

x-displayName: "Bearer Token".

`token` - (Required) x-displayName: "Token". See [Bearer Token Token ](#bearer-token-token) below for details.

### Credentials Choice Login Endpoint

x-displayName: "Login Endpoint".

`json_payload` - (Required) x-displayName: "JSON Payload". See [Login Endpoint Json Payload ](#login-endpoint-json-payload) below for details.

`method` - (Required) x-displayName: "Method" (`String`).

`path` - (Required) x-displayName: "Path" (`String`).

`token_response_key` - (Required) x-displayName: "Token Response Key" (`String`).

### Domains Credentials

Add credentials for API testing to use in the selected environment..

`credential_name` - (Required) Enter a unique name for the credentials used in API testing (`String`).

###### One of the arguments from this list "api_key, basic_auth, bearer_token, login_endpoint" must be set

`api_key` - (Optional) x-displayName: "API Key". See [Credentials Choice Api Key ](#credentials-choice-api-key) below for details.

`basic_auth` - (Optional) x-displayName: "Basic Authentication". See [Credentials Choice Basic Auth ](#credentials-choice-basic-auth) below for details.

`bearer_token` - (Optional) x-displayName: "Bearer Token". See [Credentials Choice Bearer Token ](#credentials-choice-bearer-token) below for details.

`login_endpoint` - (Optional) x-displayName: "Login Endpoint". See [Credentials Choice Login Endpoint ](#credentials-choice-login-endpoint) below for details.

###### One of the arguments from this list "admin, standard" must be set

`admin` - (Optional) x-displayName: "Admin" (`Bool`).

`standard` - (Optional) x-displayName: "Standard" (`Bool`).

### Login Endpoint Json Payload

x-displayName: "JSON Payload".

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Role Choice Admin

x-displayName: "Admin".

### Role Choice Standard

x-displayName: "Standard".

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

-	`id` - This is the id of the configured api_testing.
