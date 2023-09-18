---

page_title: "Volterra: log_receiver"

description: "The log_receiver allows CRUD of Log Receiver resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_log_receiver
==============================

The Log Receiver allows CRUD of Log Receiver resource on Volterra SaaS

~> **Note:** Please refer to [Log Receiver API docs](https://docs.cloud.f5.com/docs/api/log-receiver) to learn more

Example Usage
-------------

```hcl
resource "volterra_log_receiver" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "syslog data_dog splunk" must be set

  splunk {
    // One of the arguments from this list "splunk_server_name splunk_server_ip splunk_server_tls" must be set

    splunk_server_name {
      port        = "3000"
      server_name = "server.example.com"
    }

    splunk_hec_token {
      blindfold_secret_info_internal {
        decryption_provider = "value"
        location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
        store_provider      = "value"
      }

      secret_encoding_type = "secret_encoding_type"

      // One of the arguments from this list "blindfold_secret_info vault_secret_info clear_secret_info wingman_secret_info" must be set

      blindfold_secret_info {
        decryption_provider = "value"
        location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
        store_provider      = "value"
      }
    }
  }
  // One of the arguments from this list "site_local log_receiver_sites" must be set
  site_local = true
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

`data_dog` - (Optional) Stream log to Datadog receiver. See [Data Dog ](#data-dog) below for details.

`splunk` - (Optional) Stream log to Splunk HEC Receiver. See [Splunk ](#splunk) below for details.

`syslog` - (Optional) Stream log to syslog server. See [Syslog ](#syslog) below for details.

`log_receiver_sites` - (Optional) Log receiver is accessible on a specific site.. See [Log Receiver Sites ](#log-receiver-sites) below for details.

`site_local` - (Optional) Log receiver is accessible local to the site where it is used. (bool).

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

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Compression Disabled

Disable compression of log messages.

### Compression Enabled

Enable compression of log messages.

### Data Dog

Stream log to Datadog receiver.

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

`volterra_ca` - (Optional) Use Volterra default CA (bool).

`compression_disabled` - (Optional) Disable compression of log messages (bool).

`compression_enabled` - (Optional) Enable compression of log messages (bool).

`datadog_api_key` - (Optional) Secret API key to access datadog servers. See [Datadog Api Key ](#datadog-api-key) below for details.

`default_port` - (Optional) Default port number https is 443. (bool).

`port` - (Required) Custom port number used for communication (`Int`).

`datadog_default_server` - (Optional) Default Datadog server name (bool).

`server_name` - (Optional) Custom fully qualified server name (`String`).

### Datadog Api Key

Secret API key to access datadog servers.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Datadog Default Server

Default Datadog server name.

### Default Https Port

Default port number for HTTPS is 443..

### Default Port

Default port number https is 443..

### Default Syslog Tls Port

Default port number for syslog TLS is 6514.

### Key Url

The data may be optionally secured using BlindFold..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Log Receiver Sites

Log receiver is accessible on a specific site..

`network` - (Required) Select Network through which log receiver is accessible (`String`).

`site` - (Required) Reference to CE sites. See [ref](#ref) below for details.

### Mtls Disabled

mTLS is disabled.

### Mtls Enable

Enable mTLS configuration.

`certificate` - (Optional) Client certificate is PEM-encoded certificate or certificate-chain. (`String`).

`key_url` - (Optional) The data may be optionally secured using BlindFold.. See [Key Url ](#key-url) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Splunk

Stream log to Splunk HEC Receiver.

`splunk_server_ip` - (Optional) Splunk HEC server ip address and port number. See [Splunk Server Ip ](#splunk-server-ip) below for details.

`splunk_server_name` - (Optional) Fully qualified splunk HEC server name and port number. See [Splunk Server Name ](#splunk-server-name) below for details.

`splunk_server_tls` - (Optional) Splunk TLS Server Parameters. See [Splunk Server Tls ](#splunk-server-tls) below for details.

`splunk_hec_token` - (Optional) Secret splunk HEC token. See [Splunk Hec Token ](#splunk-hec-token) below for details.

### Splunk Hec Token

Secret splunk HEC token.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Splunk Server Ip

Splunk HEC server ip address and port number.

`ip_address` - (Required) Ip address of the log receiver server (`String`).

`port` - (Required) Port number used for communication (`Int`).

### Splunk Server Name

Fully qualified splunk HEC server name and port number.

`port` - (Required) Port number used for communication (`Int`).

`server_name` - (Required) Server name is fully qualified domain name of the server (`String`).

### Splunk Server Tls

Splunk TLS Server Parameters.

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

`volterra_ca` - (Optional) Use Volterra default CA (bool).

`mtls_disabled` - (Optional) mTLS is disabled (bool).

`mtls_enable` - (Optional) Enable mTLS configuration. See [Mtls Enable ](#mtls-enable) below for details.

`default_https_port` - (Optional) Default port number for HTTPS is 443. (bool).

`default_syslog_tls_port` - (Optional) Default port number for syslog TLS is 6514 (bool).

`port` - (Required) Custom port number used for communication (`Int`).

`server_name` - (Required) certificates against. (`String`).

### Syslog

Stream log to syslog server.

`syslog_rfc3164` - (Optional) Select RFC3164 syslog format and maximum message length. (`Int`).

`syslog_rfc5424` - (Optional) Select RFC5424 syslog format and maximum message length. (`Int`).

`tcp_server` - (Optional) Syslog transport mode is TCP. See [Tcp Server ](#tcp-server) below for details.

`tls_server` - (Optional) Syslog transport mode is TLS. See [Tls Server ](#tls-server) below for details.

`udp_server` - (Optional) Syslog transport mode is UDP. See [Udp Server ](#udp-server) below for details.

### Tcp Server

Syslog transport mode is TCP.

`port` - (Required) Port number used for communication (`Int`).

`server_name` - (Required) Server name is fully qualified domain name or IP address of the server (`String`).

### Tls Server

Syslog transport mode is TLS.

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

`volterra_ca` - (Optional) Use Volterra default CA (bool).

`mtls_disabled` - (Optional) mTLS is disabled (bool).

`mtls_enable` - (Optional) Enable mTLS configuration. See [Mtls Enable ](#mtls-enable) below for details.

`default_https_port` - (Optional) Default port number for HTTPS is 443. (bool).

`default_syslog_tls_port` - (Optional) Default port number for syslog TLS is 6514 (bool).

`port` - (Required) Custom port number used for communication (`Int`).

`server_name` - (Required) certificates against. (`String`).

### Udp Server

Syslog transport mode is UDP.

`port` - (Required) Port number used for communication (`Int`).

`server_name` - (Required) Server name is fully qualified domain name or IP address of the server (`String`).

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Volterra Ca

Use Volterra default CA.

### Wingman Secret Info

Secret is given as bootstrap secret in Volterra Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured log_receiver.
