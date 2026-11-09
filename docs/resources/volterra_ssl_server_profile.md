---

page_title: "Volterra: ssl_server_profile"

description: "The ssl_server_profile allows CRUD of Ssl Server Profile resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------

Resource volterra_ssl_server_profile
====================================

The Ssl Server Profile allows CRUD of Ssl Server Profile resource on Volterra SaaS

~> **Note:** Please refer to [Ssl Server Profile API docs](https://docs.cloud.f5.com/docs-v2/api/ssl-server-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_ssl_server_profile" "example" {
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

`alert_timeout` - (Optional) Specifies the alert timeout in seconds. You can also specify indefinite or immediate. (`String`).

`allow_expired_crl` - (Optional) Use the specified CRL file even if it has expired. (`String`).

`authenticate` - (Optional) Specifies frequency of authentication. The default value is once. Note that if this is set to always session cache and session ticket will be disabled. (`String`).

`authenticate_depth` - (Optional) Specifies the client certificate chain maximum traversal depth. (`Int`).

`authenticate_name` - (Optional) Specifies a Common Name (CN) that is embedded in a server certificate. The system authenticates a server based on the specified CN. (`String`).

`cache_size` - (Optional) Specifies the SSL session cache size in number of sessions. For client-side profiles only, you can configure timeout and size values for the SSL session cache. Because each profile maintains a separate SSL session cache, you can configure the values on a per-profile basis. (`Int`).

###### One of the arguments from this list "certificates" can be set

`certificates` - (Optional) Select one or more certificates with any domain names.. See [Certificates Choice Certificates ](#certificates-choice-certificates) below for details.

###### One of the arguments from this list "ciphers" can be set

`ciphers` - (Optional) Specifies a cipher name. (`String`).

`expire_cert_response_control` - (Optional) Specifies the action for the BIG-IP system to take when the server certificate has expired. The default value is drop, which causes the connection to be dropped. Conversely, you can specify ignore to cause the connection to ignore the error and continue or you can specify mask in case of SSL forward proxy to mask server certificate errors and continue with handshake and forge a good certificate on client-side. Note that drop works only if the certificate is trusted. (`String`).

`handshake_timeout` - (Optional) Specifies the handshake timeout in seconds. You can also specify indefinite. (`Int`).

`peer_cert_mode` - (Optional) Specifies the peer certificate mode. (`String`).

`renegotiate_period` - (Optional) Specifies the number of seconds from the initial connect time after which the system renegotiates an SSL session. The default value is indefinite meaning that you do not want the system to renegotiate SSL sessions. Each time the session renegotiation is successful, a new connection is started. Therefore, the system attempts to renegotiate the session again, in the specified amount of time following the successful session renegotiation. For example, setting the Renegotiate Period to 3600 seconds triggers session renegotiation at least once an hour. (`String`).

`renegotiate_size` - (Optional) Specifies a throughput size, in bytes, of SSL renegotiation. This setting forces the traffic management system to renegotiate an SSL session based on the size, in megabytes, of application data that is transmitted over the secure channel. The default value is indefinite specifying that you do not want a throughput size. (`String`).

`renegotiation` - (Optional) Controls mid-stream renegotiation. The default value is enabled. (`String`).

`revoked_cert_status_response_control` - (Optional) Specifies the system action when the server certificate status is revoked. The default value is drop, which causes the connection to be dropped. You can specify ignore to cause the connection to ignore the error and continue handshake. You can specify mask in case of SSL forward proxy to mask server certificate status error and continue handshake. (`String`).

`secure_renegotiation` - (Optional) Controls secure renegotiation. The default value is require-strict. (`String`).

`server_name` - (Optional) Name matched to TLS/1.1 and above client SSL requests that support the Server Name Indication extension. The default value is empty, which disables support for this extension. (`String`).

`session_ticket` - (Optional) Enables or disables session-ticket. The default option is disabled, see RFC5077. (`String`).

`tm_options` - (Optional) Enables options, including some industry-related workarounds. Enter options inside braces, for example, { dont-insert-empty-fragments }. The default value is { dont-insert-empty-fragments no-tlsv1.3 }. (`List of String`).

`trusted_ca_list` - (Optional) the system validates the client certificate against the specified trusted CA list.. See [ref](#ref) below for details.

`unknown_cert_status_response_control` - (Optional) Specifies the system action when the server certificate status is unknown. The default value is ignore, which causes the connection to ignore the error and continue handshake. You can specify drop which causes the connection to be dropped. You can specify mask in case of SSL forward proxy to mask server certificate status error and continue handshake. (`String`).

`untrusted_cert_response_control` - (Optional) Specifies the system action when the server certificate has untrusted CA. The default value is drop, which causes the connection to be dropped. Conversely, you can specify ignore to cause the connection to ignore the error and continue or you can specify mask in case of SSL forward proxy to mask server certificate errors and continue with handshake and forge a good certificate on client-side. (`String`).

### Certificates Choice Certificates

Select one or more certificates with any domain names..

`certificates` - (Optional) List of certificate references. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured ssl_server_profile.
