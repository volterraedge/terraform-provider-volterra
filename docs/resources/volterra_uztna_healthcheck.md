---

page_title: "Volterra: uztna_healthcheck"

description: "The uztna_healthcheck allows CRUD of Uztna Healthcheck resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_uztna_healthcheck
===================================

The Uztna Healthcheck allows CRUD of Uztna Healthcheck resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Healthcheck API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-healthcheck) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_healthcheck" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "http_health_check https_health_check icmp_health_check tcp_health_check" must be set

  http_health_check {
    host_header {
      expected_response = "HTTP/1"

      expected_status_codes = ["200-250"]

      headers = {
        "key1" = "value1"
      }

      // One of the arguments from this list "host_header use_origin_server_name" must be set

      use_origin_server_name = true
      password {
        blindfold_secret_info_internal {
          decryption_provider = "value"

          location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

          store_provider = "value"
        }

        secret_encoding_type = "secret_encoding_type"

        // One of the arguments from this list "blindfold_secret_info clear_secret_info vault_secret_info wingman_secret_info" must be set

        blindfold_secret_info {
          decryption_provider = "value"

          location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

          store_provider = "value"
        }
      }
      path = "/healthcheck"
      request_headers_to_remove = ["user-agent"]
      send_payload = "HEAD / HTTP/1.0"
      user_name = "user_name"
    }
  }
  interval = ["10"]
  timeout  = ["1"]
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

###### One of the arguments from this list "http_health_check, https_health_check, icmp_health_check, tcp_health_check" must be set

`http_health_check` - (Optional) 4. Request headers to remove. See [Health Check Http Health Check ](#health-check-http-health-check) below for details.

`https_health_check` - (Optional) 4. Request headers to remove. See [Health Check Https Health Check ](#health-check-https-health-check) below for details.

`icmp_health_check` - (Optional) Specifies ICMP HealthCheck (`Bool`).

`tcp_health_check` - (Optional) Specifies send payload and expected response payload. See [Health Check Tcp Health Check ](#health-check-tcp-health-check) below for details.

`interval` - (Required) Time interval in seconds between two healthcheck requests. (`Int`).

`timeout` - (Required) health check attempt will be considered a failure. (`Int`).

### Health Check Http Health Check

1.	Request headers to remove.

`host_header` - (Required) Host Header . See [Http Health Check Host Header ](#http-health-check-host-header) below for details.

### Health Check Https Health Check

1.	Request headers to remove.

`host_header` - (Required) Host Header for HTTPS Header . See [Https Health Check Host Header ](#https-health-check-host-header) below for details.

### Health Check Tcp Health Check

Specifies send payload and expected response payload.

`expected_response` - (Optional) Specifies a regular expression pattern which will be matched against response payload (`String`).

`send_payload` - (Optional) Text string sent in the request (`String`).

### Host Header Password

x-example: "value".

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Password Blindfold Secret Info Internal ](#password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Host Header Choice Use Origin Server Name

Use the origin server name..

### Http Health Check Host Header

Host Header .

`expected_response` - (Required) Regular expression used to match against the response to the health check's request. Mark node up upon receipt of a successful regular expression match. Uses re2 regular expression syntax. (`String`).

`expected_status_codes` - (Optional) of which is single HTTP status code or a range with start and end values separated by "-". (`String`).

`headers` - (Optional) health checked cluster. This is a list of key-value pairs. (`String`).

###### One of the arguments from this list "host_header, use_origin_server_name" must be set

`host_header` - (Optional) The value of the host header. (`String`).

`use_origin_server_name` - (Optional) Use the origin server name. (`Bool`).

`password` - (Required) x-example: "value". See [Host Header Password ](#host-header-password) below for details.

`path` - (Required) Specifies the HTTPS path that will be requested during health checking. (`String`).

`request_headers_to_remove` - (Optional) health checked cluster. This is a list of keys of headers. (`String`).

`send_payload` - (Required) HTTP payload to send to the target (`String`).

`user_name` - (Required) User Name (`String`).

### Https Health Check Host Header

Host Header for HTTPS Header .

`expected_response` - (Required) Regular expression used to match against the response to the health check's request. Mark node up upon receipt of a successful regular expression match. Uses re2 regular expression syntax. (`String`).

`expected_status_codes` - (Optional) of which is single HTTP status code or a range with start and end values separated by "-". (`String`).

`headers` - (Optional) health checked cluster. This is a list of key-value pairs. (`String`).

###### One of the arguments from this list "host_header, use_origin_server_name" must be set

`host_header` - (Optional) The value of the host header. (`String`).

`use_origin_server_name` - (Optional) Use the origin server name. (`Bool`).

`password` - (Required) x-example: "value". See [Host Header Password ](#host-header-password) below for details.

`path` - (Required) Specifies the HTTPS path that will be requested during health checking. (`String`).

`request_headers_to_remove` - (Optional) health checked cluster. This is a list of keys of headers. (`String`).

`send_payload` - (Required) HTTP payload to send to the target (`String`).

`user_name` - (Required) User Name (`String`).

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

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_healthcheck.
