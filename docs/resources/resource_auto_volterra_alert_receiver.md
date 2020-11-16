---

page_title: "Volterra: alert_receiver"

description: "The alert_receiver allows CRUD of Alert Receiver resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_alert_receiver
================================

The Alert Receiver allows CRUD of Alert Receiver resource on Volterra SaaS

~> **Note:** Please refer to [Alert Receiver API docs](https://volterra.io/docs/api/alert-receiver) to learn more

Example Usage
-------------

```hcl
resource "volterra_alert_receiver" "example" {
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

`email` - (Optional) Send alert notifications as Email. See [Email ](#email) below for details.

`opsgenie` - (Optional) Send alert notifications to OpsGenie. See [Opsgenie ](#opsgenie) below for details.

`pagerduty` - (Optional) Send alert notifications to PagerDuty. See [Pagerduty ](#pagerduty) below for details.

`slack` - (Optional) Send alert notifications to Slack. See [Slack ](#slack) below for details.

`sms` - (Optional) Send alert notifications as SMS. See [Sms ](#sms) below for details.

### Api Key

API integration key to send alert notifications using REST API to OpsGenie service..

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Email

Send alert notifications as Email.

`email` - (Optional) Email id of the user (`String`).

### Opsgenie

Send alert notifications to OpsGenie.

`api_key` - (Required) API integration key to send alert notifications using REST API to OpsGenie service.. See [Api Key ](#api-key) below for details.

`url` - (Required) URL to send API requests to (`String`).

### Pagerduty

Send alert notifications to PagerDuty.

`routing_key` - (Required) PagerDuty integration key (choose Integration Type: Events API v2). See [Routing Key ](#routing-key) below for details.

`url` - (Required) URL to send API requests to (`String`).

### Routing Key

PagerDuty integration key (choose Integration Type: Events API v2).

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Slack

Send alert notifications to Slack.

`channel` - (Required) Channel or user to send notifications to (`String`).

`url` - (Required) API Key is embedded in the webhook URL.. See [Url ](#url) below for details.

### Sms

Send alert notifications as SMS.

`contact_number` - (Optional) Contact number of the user in ITU E.164 format [+][country code][subscriber number including area code](`String`).

### Url

API Key is embedded in the webhook URL..

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured alert_receiver.
