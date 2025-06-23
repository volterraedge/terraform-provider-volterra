---

page_title: "Volterra: dns_domain"

description: "The dns_domain allows CRUD of Dns Domain resource on Volterra SaaS"
---------------------------------------------------------------------------------

Resource volterra_dns_domain
============================

The Dns Domain allows CRUD of Dns Domain resource on Volterra SaaS

~> **Note:** Please refer to [Dns Domain API docs](https://docs.cloud.f5.com/docs-v2/api/dns-domain) to learn more

Example Usage
-------------

```hcl
resource "volterra_dns_domain" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "volterra_managed" must be set

  volterra_managed = true
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

`dnssec_mode` - (Optional) Control whether DNSSEC is enabled on the dns domain or not (`String`).

###### One of the arguments from this list "volterra_managed" must be set

`volterra_managed` - (Optional) sub domain (`Bool`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured dns_domain.
-	`txt_record` - This is the txt-record of the configured dns_domain.
