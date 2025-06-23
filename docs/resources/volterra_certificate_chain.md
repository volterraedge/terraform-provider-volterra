---

page_title: "Volterra: certificate_chain"

description: "The certificate_chain allows CRUD of Certificate Chain resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_certificate_chain
===================================

The Certificate Chain allows CRUD of Certificate Chain resource on Volterra SaaS

~> **Note:** Please refer to [Certificate Chain API docs](https://docs.cloud.f5.com/docs-v2/api/certificate-chain) to learn more

Example Usage
-------------

```hcl
resource "volterra_certificate_chain" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  certificate_url = ["value"]
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

`certificate_url` - (Required) Certificate chain is the list of intermediate certificates in PEM format including the PEM headers. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured certificate_chain.
