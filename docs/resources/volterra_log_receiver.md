---

page_title: "Volterra: log_receiver"

description: "The log_receiver allows CRUD of Log Receiver resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_log_receiver
==============================

The Log Receiver allows CRUD of Log Receiver resource on Volterra SaaS

~> **Note:** Please refer to [Log Receiver API docs](https://docs.cloud.f5.com/docs-v2/api/log-receiver) to learn more

Example Usage
-------------

```hcl
resource "volterra_log_receiver" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "syslog" must be set

  syslog {
    // One of the arguments from this list "syslog_rfc5424" must be set

    syslog_rfc5424 = "500"

    // One of the arguments from this list "tcp_server tls_server udp_server" must be set

    udp_server {
      port = "514"

      server_name = "server.example.com"
    }
  }

  // One of the arguments from this list "site_local" must be set

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

###### One of the arguments from this list "syslog" must be set

`syslog` - (Optional) Stream log to syslog server. See [Log Receiver Choice Syslog ](#log-receiver-choice-syslog) below for details.

###### One of the arguments from this list "site_local" must be set

`site_local` - (Optional) Log receiver is accessible local to the site where it is used. (`Bool`).

### Ca Choice Volterra Ca

Use F5XC default CA.

### Log Receiver Choice Syslog

Stream log to syslog server.

###### One of the arguments from this list "syslog_rfc5424" must be set

`syslog_rfc5424` - (Optional) Select RFC5424 syslog format and maximum message length. (`Int`).

###### One of the arguments from this list "tcp_server, tls_server, udp_server" must be set

`tcp_server` - (Optional) Syslog transport mode is TCP. See [Mode Choice Tcp Server ](#mode-choice-tcp-server) below for details.

`tls_server` - (Optional) Syslog transport mode is TLS. See [Mode Choice Tls Server ](#mode-choice-tls-server) below for details.

`udp_server` - (Optional) Syslog transport mode is UDP. See [Mode Choice Udp Server ](#mode-choice-udp-server) below for details.

### Mode Choice Tcp Server

Syslog transport mode is TCP.

`port` - (Required) Port number used for communication (`Int`).

`server_name` - (Required) Server name is fully qualified domain name or IP address of the server (`String`).

### Mode Choice Tls Server

Syslog transport mode is TLS.

###### One of the arguments from this list "trusted_ca_url, volterra_ca" must be set

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

`volterra_ca` - (Optional) Use F5XC default CA (`Bool`).

###### One of the arguments from this list "mtls_disabled, mtls_enable" must be set

`mtls_disabled` - (Optional) mTLS is disabled (`Bool`).

`mtls_enable` - (Optional) Enable mTLS configuration. See [Mtls Choice Mtls Enable ](#mtls-choice-mtls-enable) below for details.

###### One of the arguments from this list "default_https_port, default_syslog_tls_port, port" must be set

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

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Port Choice Default Https Port

Default port number for HTTPS is 443..

### Port Choice Default Syslog Tls Port

Default port number for syslog TLS is 6514.

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

-	`id` - This is the id of the configured log_receiver.
