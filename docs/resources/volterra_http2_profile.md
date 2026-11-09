---

page_title: "Volterra: http2_profile"

description: "The http2_profile allows CRUD of Http2 Profile resource on Volterra SaaS"
---------------------------------------------------------------------------------------

Resource volterra_http2_profile
===============================

The Http2 Profile allows CRUD of Http2 Profile resource on Volterra SaaS

~> **Note:** Please refer to [Http2 Profile API docs](https://docs.cloud.f5.com/docs-v2/api/http2-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_http2_profile" "example" {
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

`activation_modes` - (Optional) Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default value alpn specifies that the TLS application-layer-protocol-negotiation extension will be used. (`List of String`).

`app_service` - (Optional) The application service to which the object belongs. (`String`).

`command` - (Optional) x-displayName: "Command" (`String`).

`concurrent_streams_per_connection` - (Optional) Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection. (`Int`).

`connection_idle_timeout` - (Optional) Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. (`Int`).

`enforce_tls_requirements` - (Optional) x-displayName: "Enforce Tls Requirements" (`String`).

`frame_size` - (Optional) The size in bytes of the data frames that will be produced by HTTP/2. (`Int`).

`header_table_size` - (Optional) Specifies what table size will be used for the compression of headers. (`Int`).

`include_content_length` - (Optional) x-displayName: "Include Content Length" (`String`).

`insert_header` - (Optional) Specifies whether an HTTP header should be added to the HTTP request to show the request was received via HTTP/2. (`String`).

`insert_header_name` - (Optional) Specifies the name of the header that is added to the HTTP request when insert-header is enabled. (`String`).

`receive_window` - (Optional) Specifies in KB the size of the receive window for HTTP/2 flow-control. (`Int`).

`user_spec` - (Optional) User specified properties. (`List of String`).

`write_size` - (Optional) The size in bytes of the SSL records that will be produced by HTTP/2.handled. (`Int`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured http2_profile.
