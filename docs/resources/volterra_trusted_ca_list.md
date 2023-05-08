---

page_title: "Volterra: trusted_ca_list"

description: "The trusted_ca_list allows CRUD of Trusted Ca List resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_trusted_ca_list
=================================

The Trusted Ca List allows CRUD of Trusted Ca List resource on Volterra SaaS

~> **Note:** Please refer to [Trusted Ca List API docs](https://volterra.io/docs/api/trusted-ca-list) to learn more

Example Usage
-------------

```hcl
resource "volterra_trusted_ca_list" "example" {
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

`trusted_ca_url` - (Optional) Trusted CA certificates for validating certificates (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured trusted_ca_list.
