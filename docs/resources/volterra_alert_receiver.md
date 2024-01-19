---

page_title: "Volterra: alert_receiver"

description: "The alert_receiver allows CRUD of Alert Receiver resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_alert_receiver
================================

The Alert Receiver allows CRUD of Alert Receiver resource on Volterra SaaS

~> **Note:** Please refer to [Alert Receiver API docs](https://docs.cloud.f5.com/docs/api/alert-receiver) to learn more

Example Usage
-------------

```hcl
resource "volterra_alert_receiver" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "slack pagerduty opsgenie email sms webhook" must be set

  pagerduty {
    routing_key {
      blindfold_secret_info_internal {
        decryption_provider = "value"
        location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
        store_provider      = "value"
      }

      secret_encoding_type = "secret_encoding_type"

      // One of the arguments from this list "vault_secret_info clear_secret_info wingman_secret_info blindfold_secret_info" must be set

      blindfold_secret_info {
        decryption_provider = "value"
        location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
        store_provider      = "value"
      }
    }

    url = "value"
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

`email` - (Optional) Send alert notifications as Email. See [Email ](#email) below for details.

`opsgenie` - (Optional) Send alert notifications to OpsGenie. See [Opsgenie ](#opsgenie) below for details.

`pagerduty` - (Optional) Send alert notifications to PagerDuty. See [Pagerduty ](#pagerduty) below for details.

`slack` - (Optional) Send alert notifications to Slack. See [Slack ](#slack) below for details.

`sms` - (Optional) Send alert notifications as SMS. See [Sms ](#sms) below for details.

`webhook` - (Optional) Send alert notifications to Webhook. See [Webhook ](#webhook) below for details.

### Api Key

API integration key to send alert notifications using REST API to OpsGenie service..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Auth Token

Configure an Access Token for authentication to the HTTP(s) server (such as a Bearer Token).

`token` - (Required) F5XC Secret. URL for token, needs to be fetched from this path. See [Token ](#token) below for details.

### Basic Auth

Use HTTP Basic Auth for authentication to the HTTP(s) server.

`password` - (Required) HTTP Basic Auth Password. See [Password ](#password) below for details.

`user_name` - (Required) HTTP Basic Auth User Name (`String`).

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

### Ca Cert Obj

Trusted CA List for verification of Server's certificate.

`trusted_ca` - (Optional) Reference to client certificate object. See [ref](#ref) below for details.

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Client Cert Obj

Use certificate and key files for client cert authentication to the server..

`use_tls_obj` - (Optional) Reference to client certificate object. See [ref](#ref) below for details.

### Disable Sni

Do not use SNI..

### Email

Send alert notifications as Email.

`email` - (Optional) Email id of the user (`String`).

### Http Config

Configuration for HTTP endpoint.

`auth_token` - (Optional) Configure an Access Token for authentication to the HTTP(s) server (such as a Bearer Token). See [Auth Token ](#auth-token) below for details.

`basic_auth` - (Optional) Use HTTP Basic Auth for authentication to the HTTP(s) server. See [Basic Auth ](#basic-auth) below for details.

`client_cert_obj` - (Optional) Use certificate and key files for client cert authentication to the server.. See [Client Cert Obj ](#client-cert-obj) below for details.

`no_authorization` - (Optional) Do not use authentication to the HTTP(s) server (bool).

`enable_http2` - (Optional) Configure to use HTTP2 protocol. (`Bool`).

`follow_redirects` - (Optional) Configure whether HTTP requests follow HTTP 3xx redirects. (`Bool`).

`no_tls` - (Optional) x-displayName: "Disable" (bool).

`use_tls` - (Optional) x-displayName: "Enable". See [Use Tls ](#use-tls) below for details.

### No Authorization

Do not use authentication to the HTTP(s) server.

### No Tls

x-displayName: "Disable".

### Opsgenie

Send alert notifications to OpsGenie.

`api_key` - (Required) API integration key to send alert notifications using REST API to OpsGenie service.. See [Api Key ](#api-key) below for details.

`url` - (Required) URL to send API requests to (`String`).

### Pagerduty

Send alert notifications to PagerDuty.

`routing_key` - (Required) PagerDuty integration key (choose Integration Type: Events API v2). See [Routing Key ](#routing-key) below for details.

`url` - (Required) URL to send API requests to (`String`).

### Password

HTTP Basic Auth Password.

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

### Routing Key

PagerDuty integration key (choose Integration Type: Events API v2).

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Slack

Send alert notifications to Slack.

`channel` - (Required) Channel or user to send notifications to (`String`).

`url` - (Required) API Key is embedded in the webhook URL.. See [Url ](#url) below for details.

### Sms

Send alert notifications as SMS.

`contact_number` - (Optional) Contact number of the user in ITU E.164 format [+][country code][subscriber number including area code](`String`).

### Token

F5XC Secret. URL for token, needs to be fetched from this path.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Url

API Key is embedded in the webhook URL..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Use Server Verification

Perform server verification using the provided trusted CA list.

`ca_cert_obj` - (Optional) Trusted CA List for verification of Server's certificate. See [Ca Cert Obj ](#ca-cert-obj) below for details.

### Use Tls

x-displayName: "Enable".

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

`use_server_verification` - (Optional) Perform server verification using the provided trusted CA list. See [Use Server Verification ](#use-server-verification) below for details.

`volterra_trusted_ca` - (Optional) Perform server verification using F5XC default trusted CA list (bool).

`disable_sni` - (Optional) Do not use SNI. (bool).

`sni` - (Optional) SNI value to be used. (`String`).

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Volterra Trusted Ca

Perform server verification using F5XC default trusted CA list.

### Webhook

Send alert notifications to Webhook.

`http_config` - (Required) Configuration for HTTP endpoint. See [Http Config ](#http-config) below for details.

`url` - (Required) Incoming webhook url to send alert notifications.. See [Url ](#url) below for details.

### Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured alert_receiver.
