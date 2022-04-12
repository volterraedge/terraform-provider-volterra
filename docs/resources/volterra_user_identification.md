---

page_title: "Volterra: user_identification"

description: "The user_identification allows CRUD of User Identification resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------

Resource volterra_user_identification
=====================================

The User Identification allows CRUD of User Identification resource on Volterra SaaS

~> **Note:** Please refer to [User Identification API docs](https://volterra.io/docs/api/user-identification) to learn more

Example Usage
-------------

```hcl
resource "volterra_user_identification" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  rules {
    // One of the arguments from this list "tls_fingerprint ip_and_http_header_name none client_asn ip_and_tls_fingerprint client_ip query_param_key http_header_name cookie_name" must be set
    cookie_name = "Session"
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

`rules` - (Required) Evaluation of the rules is terminated once a user identifier has been extracted. . See [Rules ](#rules) below for details.

### Client Asn

The client ASN is obtained by performing a lookup for the client IP Address in a GeoIP DB..

### Client Ip

Use client IP address as user identifier..

### Ip And Tls Fingerprint

Use the combination of Client IP and TLS Fingerprint as user identifier..

### None

Do not use any user identifier..

### Rules

Evaluation of the rules is terminated once a user identifier has been extracted. .

`client_asn` - (Optional) The client ASN is obtained by performing a lookup for the client IP Address in a GeoIP DB. (bool).

`client_ip` - (Optional) Use client IP address as user identifier. (bool).

`cookie_name` - (Optional) Use the HTTP cookie value for the given name as user identifier. (`String`).

`http_header_name` - (Optional) Use the HTTP header value for the given name as user identifier. (`String`).

`ip_and_http_header_name` - (Optional) Use the combination of Client IP and HTTP header value for the given name as user identifier. (`String`).

`ip_and_tls_fingerprint` - (Optional) Use the combination of Client IP and TLS Fingerprint as user identifier. (bool).

`none` - (Optional) Do not use any user identifier. (bool).

`query_param_key` - (Optional) Use the query parameter value for the given key as user identifier. (`String`).

`tls_fingerprint` - (Optional) Use TLS Fingerprint as user identifier. (bool).

### Tls Fingerprint

Use TLS Fingerprint as user identifier..

Attribute Reference
-------------------

-	`id` - This is the id of the configured user_identification.
