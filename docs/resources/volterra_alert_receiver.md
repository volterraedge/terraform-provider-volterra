---

page_title: "Volterra: alert_receiver"

description: "The alert_receiver allows CRUD of Alert Receiver resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_alert_receiver
================================

The Alert Receiver allows CRUD of Alert Receiver resource on Volterra SaaS

~> **Note:** Please refer to [Alert Receiver API docs](https://docs.cloud.f5.com/docs-v2/api/alert-receiver) to learn more

Example Usage
-------------

```hcl
resource "volterra_alert_receiver" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "email opsgenie pagerduty slack sms webhook" must be set

  email {
    email = "abc@email.com"
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

###### One of the arguments from this list "email, opsgenie, pagerduty, slack, sms, webhook" must be set

`email` - (Optional) Send alert notifications as Email. See [Receiver Email ](#receiver-email) below for details.

`opsgenie` - (Optional) Send alert notifications to OpsGenie. See [Receiver Opsgenie ](#receiver-opsgenie) below for details.

`pagerduty` - (Optional) Send alert notifications to PagerDuty. See [Receiver Pagerduty ](#receiver-pagerduty) below for details.

`slack` - (Optional) Send alert notifications to Slack. See [Receiver Slack ](#receiver-slack) below for details.

`sms` - (Optional) Send alert notifications as SMS. See [Receiver Sms ](#receiver-sms) below for details.

`webhook` - (Optional) Send alert notifications to Webhook. See [Receiver Webhook ](#receiver-webhook) below for details.

### Auth Choice Auth Token

Configure an Access Token for authentication to the HTTP(s) server (such as a Bearer Token).

`token` - (Required) F5XC Secret. URL for token, needs to be fetched from this path. See [Auth Token Token ](#auth-token-token) below for details.

### Auth Choice Basic Auth

Use HTTP Basic Auth for authentication to the HTTP(s) server.

`password` - (Required) HTTP Basic Auth Password. See [Basic Auth Password ](#basic-auth-password) below for details.

`user_name` - (Required) HTTP Basic Auth User Name (`String`).

### Auth Choice Client Cert Obj

Use certificate and key files for client cert authentication to the server..

`use_tls_obj` - (Optional) Reference to client certificate object. See [ref](#ref) below for details.

### Auth Choice No Authorization

Do not use authentication to the HTTP(s) server.

### Auth Token Token

F5XC Secret. URL for token, needs to be fetched from this path.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Basic Auth Password

HTTP Basic Auth Password.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Opsgenie Api Key

API integration key to send alert notifications using REST API to OpsGenie service..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Pagerduty Routing Key

PagerDuty integration key (choose Integration Type: Events API v2).

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Receiver Email

Send alert notifications as Email.

`email` - (Optional) Email id of the user (`String`).

### Receiver Opsgenie

Send alert notifications to OpsGenie.

`api_key` - (Required) API integration key to send alert notifications using REST API to OpsGenie service.. See [Opsgenie Api Key ](#opsgenie-api-key) below for details.

`url` - (Required) URL to send API requests to (`String`).

### Receiver Pagerduty

Send alert notifications to PagerDuty.

`routing_key` - (Required) PagerDuty integration key (choose Integration Type: Events API v2). See [Pagerduty Routing Key ](#pagerduty-routing-key) below for details.

`url` - (Required) URL to send API requests to (`String`).

### Receiver Slack

Send alert notifications to Slack.

`channel` - (Required) Channel or user to send notifications to (`String`).

`url` - (Required) API Key is embedded in the webhook URL.. See [Slack Url ](#slack-url) below for details.

### Receiver Sms

Send alert notifications as SMS.

`contact_number` - (Optional) Contact number of the user in ITU E.164 format [+][country code][subscriber number including area code](`String`).

### Receiver Webhook

Send alert notifications to Webhook.

`http_config` - (Required) Configuration for HTTP endpoint. See [Webhook Http Config ](#webhook-http-config) below for details.

`url` - (Required) Incoming webhook url to send alert notifications.. See [Webhook Url ](#webhook-url) below for details.

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

### Server Validation Choice Use Server Verification

Perform server verification using the provided trusted CA list.

`ca_cert_obj` - (Optional) Trusted CA List for verification of Server's certificate. See [Use Server Verification Ca Cert Obj ](#use-server-verification-ca-cert-obj) below for details.

### Server Validation Choice Volterra Trusted Ca

Perform server verification using F5XC default trusted CA list.

### Slack Url

API Key is embedded in the webhook URL..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Sni Choice Disable Sni

Do not use SNI..

### Tls Choice No Tls

x-displayName: "Disable".

### Tls Choice Use Tls

x-displayName: "Enable".

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

###### One of the arguments from this list "use_server_verification, volterra_trusted_ca" must be set

`use_server_verification` - (Optional) Perform server verification using the provided trusted CA list. See [Server Validation Choice Use Server Verification ](#server-validation-choice-use-server-verification) below for details.

`volterra_trusted_ca` - (Optional) Perform server verification using F5XC default trusted CA list (`Bool`).

###### One of the arguments from this list "disable_sni, sni" must be set

`disable_sni` - (Optional) Do not use SNI. (`Bool`).

`sni` - (Optional) SNI value to be used. (`String`).

### Use Server Verification Ca Cert Obj

Trusted CA List for verification of Server's certificate.

`trusted_ca` - (Optional) Reference to client certificate object. See [ref](#ref) below for details.

### Webhook Http Config

Configuration for HTTP endpoint.

###### One of the arguments from this list "auth_token, basic_auth, client_cert_obj, no_authorization" must be set

`auth_token` - (Optional) Configure an Access Token for authentication to the HTTP(s) server (such as a Bearer Token). See [Auth Choice Auth Token ](#auth-choice-auth-token) below for details.

`basic_auth` - (Optional) Use HTTP Basic Auth for authentication to the HTTP(s) server. See [Auth Choice Basic Auth ](#auth-choice-basic-auth) below for details.

`client_cert_obj` - (Optional) Use certificate and key files for client cert authentication to the server.. See [Auth Choice Client Cert Obj ](#auth-choice-client-cert-obj) below for details.

`no_authorization` - (Optional) Do not use authentication to the HTTP(s) server (`Bool`).

`enable_http2` - (Optional) Configure to use HTTP2 protocol. (`Bool`).

`follow_redirects` - (Optional) Configure whether HTTP requests follow HTTP 3xx redirects. (`Bool`).

###### One of the arguments from this list "no_tls, use_tls" must be set

`no_tls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_tls` - (Optional) x-displayName: "Enable". See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

### Webhook Url

Incoming webhook url to send alert notifications..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured alert_receiver.
