---

page_title: "Volterra: crl"
description: "The crl allows CRUD of Crl resource on Volterra SaaS"

---

Resource volterra_crl
=====================

The Crl allows CRUD of Crl resource on Volterra SaaS

~> **Note:** Please refer to [Crl API docs](https://docs.cloud.f5.com/docs-v2/api/crl) to learn more

Example Usage
-------------

```hcl
resource "volterra_crl" "example" {
  name           = "acmecorp-web"
  namespace      = "staging"
  server_address = ["mycrl.it.com"]
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

###### One of the arguments from this list "http_access" can be set

`http_access` - (Optional) http://<server_address>:<server_port>/<path>. See [Access Info Choice Http Access ](#access-info-choice-http-access) below for details.

`refresh_interval` - (Optional) CRL refresh interval, in hours. (`Int`).

`server_address` - (Required) CRL server address or hostname (`String`).

`server_port` - (Optional) Set CRL Server port number (`Int`).

`timeout` - (Optional) CRL download wait time, in seconds (`Int`).

`verify_all_certs_with_crl` - (Optional) from that chain. (`Bool`).(Deprecated)

### Access Info Choice Http Access

http://<server_address>:<server_port>/<path>.

`path` - (Optional) CRL file location (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured crl.
