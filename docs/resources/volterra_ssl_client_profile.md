---

page_title: "Volterra: ssl_client_profile"

description: "The ssl_client_profile allows CRUD of Ssl Client Profile resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------

Resource volterra_ssl_client_profile
====================================

The Ssl Client Profile allows CRUD of Ssl Client Profile resource on Volterra SaaS

~> **Note:** Please refer to [Ssl Client Profile API docs](https://docs.cloud.f5.com/docs-v2/api/ssl-client-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_ssl_client_profile" "example" {
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

`authenticate_depth` - (Optional) Specifies the authenticate depth. This is the client certificate chain maximum traversal depth. (`Int`).

`cache_size` - (Optional) Specifies the SSL session cache size. For client-side profiles only, you can configure timeout and size values for the SSL session cache. Because each profile maintains a separate SSL session cache, you can configure the values on a per-profile basis. (`Int`).

###### One of the arguments from this list "certificates" can be set

`certificates` - (Optional) Select one or more certificates with any domain names.. See [Certificates Choice Certificates ](#certificates-choice-certificates) below for details.

###### One of the arguments from this list "ciphers" can be set

`ciphers` - (Optional) Specifies a cipher name. (`String`).

`crl` - (Optional) x-displayName: "CRL" (`String`).(Deprecated)

`crl_file` - (Optional) Specifies the certificate revocation list file name. (`String`).(Deprecated)

`handshake_timeout` - (Optional) Specifies the handshake timeout in seconds. You can also specify indefinite. (`Int`).

`peer_cert_mode` - (Optional) Specifies the peer certificate mode. (`String`).

`proxy_ca_passphrase` - (Optional) Specifies the passphrase of the Certification Authority key for SSL Forward Proxy. (This option is deprecated beginning in v14.0.0. Instead, you can use the cert-key-chain option with the usage argument to add an FWDP CA key/cert.) (`String`).(Deprecated)

`renegotiate_max_record_delay` - (Optional) Specifies the maximum number of SSL records that the traffic management system can receive before it renegotiates an SSL session. After the system receives this number of SSL records, it closes the connection. This setting applies to client-side profiles only. The default value is indefinite. (`String`).

`renegotiate_period` - (Optional) Specifies the number of seconds required to renegotiate an SSL session. The default value is indefinite. (`String`).

`renegotiate_size` - (Optional) Specifies the size of the application data, in megabytes, that is transmitted over the secure channel above which the traffic management system must renegotiate the SSL session. The default value is indefinite. (`String`).

`renegotiation` - (Optional) Controls mid-stream renegotiation. The default value is enabled. (`String`).

`secure_renegotiation` - (Optional) Controls secure renegotiation. The default value is require. (`String`).

`server_name` - (Optional) Name matched to TLS/1.1 and above client SSL requests that support the Server Name Indication extension. The default value is empty, which disables support for this extension. (`String`).

`session_ticket` - (Optional) Enables or disables session-ticket. The default option is disabled, see RFC5077. (`String`).

`sni_default` - (Optional) When true, this profile is the default SSL profile when a client connection does not specify a known server name, or does not specify any server name at all. The default value is false. (`String`).

`tm_options` - (Optional) Enables options, including some industry-related workarounds. Enter options inside braces, for example, { dont-insert-empty-fragments }. The default value is { dont-insert-empty-fragments no-tlsv1.3 }. (`List of String`).

`trusted_ca_list` - (Optional) the system validates the client certificate against the specified trusted CA list.. See [ref](#ref) below for details.

### Certificates Choice Certificates

Select one or more certificates with any domain names..

`certificates` - (Required) List of certificate references. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured ssl_client_profile.
