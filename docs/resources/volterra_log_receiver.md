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

  syslog {
    // One of the arguments from this list "syslog_rfc3164 syslog_rfc5424" must be set

    syslog_rfc5424 = "500"

    // One of the arguments from this list "udp_server tcp_server tls_server" must be set

    tcp_server {
      port = "601"

      server_name = "server.example.com"
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

`data_dog` - (Optional) Stream log to Datadog receiver. See [Log Receiver Choice Data Dog ](#log-receiver-choice-data-dog) below for details.(Deprecated)

`splunk` - (Optional) Stream log to Splunk HEC Receiver. See [Log Receiver Choice Splunk ](#log-receiver-choice-splunk) below for details.(Deprecated)

`syslog` - (Optional) Stream log to syslog server. See [Log Receiver Choice Syslog ](#log-receiver-choice-syslog) below for details.

`log_receiver_sites` - (Optional) Log receiver is accessible on a specific site.. See [Where Choice Log Receiver Sites ](#where-choice-log-receiver-sites) below for details.(Deprecated)

`site_local` - (Optional) Log receiver is accessible local to the site where it is used. (`Bool`).

### Ca Choice Volterra Ca

Use F5XC default CA.

### Compression Choice Compression Disabled

Disable compression of log messages.

### Compression Choice Compression Enabled

Enable compression of log messages.

### Data Dog Datadog Api Key

Secret API key to access datadog servers.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Datadog Api Key Blindfold Secret Info Internal ](#datadog-api-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Datadog Api Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Key Url Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Log Receiver Choice Data Dog

Stream log to Datadog receiver.

###### One of the arguments from this list "volterra_ca, trusted_ca_url" must be set

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

`volterra_ca` - (Optional) Use F5XC default CA (`Bool`).

###### One of the arguments from this list "compression_enabled, compression_disabled" must be set

`compression_disabled` - (Optional) Disable compression of log messages (`Bool`).

`compression_enabled` - (Optional) Enable compression of log messages (`Bool`).

`datadog_api_key` - (Optional) Secret API key to access datadog servers. See [Data Dog Datadog Api Key ](#data-dog-datadog-api-key) below for details.

###### One of the arguments from this list "default_port, port" must be set

`default_port` - (Optional) Default port number https is 443. (`Bool`).

`port` - (Optional) Custom port number used for communication (`Int`).

###### One of the arguments from this list "server_name, datadog_default_server" must be set

`datadog_default_server` - (Optional) Default Datadog server name (`Bool`).

`server_name` - (Optional) Custom fully qualified server name (`String`).

### Log Receiver Choice Splunk

Stream log to Splunk HEC Receiver.

###### One of the arguments from this list "splunk_server_name, splunk_server_ip, splunk_server_tls" must be set

`splunk_server_ip` - (Optional) Splunk HEC server ip address and port number. See [Server Name Choice Splunk Server Ip ](#server-name-choice-splunk-server-ip) below for details.

`splunk_server_name` - (Optional) Fully qualified splunk HEC server name and port number. See [Server Name Choice Splunk Server Name ](#server-name-choice-splunk-server-name) below for details.

`splunk_server_tls` - (Optional) Splunk TLS Server Parameters. See [Server Name Choice Splunk Server Tls ](#server-name-choice-splunk-server-tls) below for details.

`splunk_hec_token` - (Optional) Secret splunk HEC token. See [Splunk Splunk Hec Token ](#splunk-splunk-hec-token) below for details.

### Log Receiver Choice Syslog

Stream log to syslog server.

###### One of the arguments from this list "syslog_rfc5424, syslog_rfc3164" must be set

`syslog_rfc3164` - (Optional) Select RFC3164 syslog format and maximum message length. (`Int`).(Deprecated)

`syslog_rfc5424` - (Optional) Select RFC5424 syslog format and maximum message length. (`Int`).

###### One of the arguments from this list "udp_server, tcp_server, tls_server" must be set

`tcp_server` - (Optional) Syslog transport mode is TCP. See [Mode Choice Tcp Server ](#mode-choice-tcp-server) below for details.

`tls_server` - (Optional) Syslog transport mode is TLS. See [Mode Choice Tls Server ](#mode-choice-tls-server) below for details.

`udp_server` - (Optional) Syslog transport mode is UDP. See [Mode Choice Udp Server ](#mode-choice-udp-server) below for details.

### Mode Choice Tcp Server

Syslog transport mode is TCP.

`port` - (Required) Port number used for communication (`Int`).

`server_name` - (Required) Server name is fully qualified domain name or IP address of the server (`String`).

### Mode Choice Tls Server

Syslog transport mode is TLS.

###### One of the arguments from this list "volterra_ca, trusted_ca_url" must be set

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

`volterra_ca` - (Optional) Use F5XC default CA (`Bool`).

###### One of the arguments from this list "mtls_disabled, mtls_enable" must be set

`mtls_disabled` - (Optional) mTLS is disabled (`Bool`).

`mtls_enable` - (Optional) Enable mTLS configuration. See [Mtls Choice Mtls Enable ](#mtls-choice-mtls-enable) below for details.

###### One of the arguments from this list "default_https_port, port, default_syslog_tls_port" must be set

`default_https_port` - (Optional) Default port number for HTTPS is 443. (`Bool`).

`default_syslog_tls_port` - (Optional) Default port number for syslog TLS is 6514 (`Bool`).

`port` - (Optional) Custom port number used for communication (`Int`).

`server_name` - (Required) certificates against. (`String`).

### Mode Choice Udp Server

Syslog transport mode is UDP.

`port` - (Required) Port number used for communication (`Int`).

`server_name` - (Required) Server name is fully qualified domain name or IP address of the server (`String`).

### Mtls Choice Mtls Disabled

mTLS is disabled.

### Mtls Choice Mtls Enable

Enable mTLS configuration.

`certificate` - (Optional) Client certificate is PEM-encoded certificate or certificate-chain. (`String`).

`key_url` - (Optional) The data may be optionally secured using BlindFold.. See [Mtls Enable Key Url ](#mtls-enable-key-url) below for details.

### Mtls Enable Key Url

The data may be optionally secured using BlindFold..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Key Url Blindfold Secret Info Internal ](#key-url-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "wingman_secret_info, blindfold_secret_info, vault_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Port Choice Default Https Port

Default port number for HTTPS is 443..

### Port Choice Default Port

Default port number https is 443..

### Port Choice Default Syslog Tls Port

Default port number for syslog TLS is 6514.

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

### Server Name Choice Datadog Default Server

Default Datadog server name.

### Server Name Choice Splunk Server Ip

Splunk HEC server ip address and port number.

`ip_address` - (Required) Ip address of the log receiver server (`String`).

`port` - (Required) Port number used for communication (`Int`).

### Server Name Choice Splunk Server Name

Fully qualified splunk HEC server name and port number.

`port` - (Required) Port number used for communication (`Int`).

`server_name` - (Required) Server name is fully qualified domain name of the server (`String`).

### Server Name Choice Splunk Server Tls

Splunk TLS Server Parameters.

###### One of the arguments from this list "volterra_ca, trusted_ca_url" must be set

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

`volterra_ca` - (Optional) Use F5XC default CA (`Bool`).

###### One of the arguments from this list "mtls_disabled, mtls_enable" must be set

`mtls_disabled` - (Optional) mTLS is disabled (`Bool`).

`mtls_enable` - (Optional) Enable mTLS configuration. See [Mtls Choice Mtls Enable ](#mtls-choice-mtls-enable) below for details.

###### One of the arguments from this list "default_syslog_tls_port, default_https_port, port" must be set

`default_https_port` - (Optional) Default port number for HTTPS is 443. (`Bool`).

`default_syslog_tls_port` - (Optional) Default port number for syslog TLS is 6514 (`Bool`).

`port` - (Optional) Custom port number used for communication (`Int`).

`server_name` - (Required) certificates against. (`String`).

### Splunk Splunk Hec Token

Secret splunk HEC token.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Splunk Hec Token Blindfold Secret Info Internal ](#splunk-hec-token-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Splunk Hec Token Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Where Choice Log Receiver Sites

Log receiver is accessible on a specific site..

`network` - (Required) Select Network through which log receiver is accessible (`String`).

`site` - (Required) Reference to CE sites. See [ref](#ref) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured log_receiver.
