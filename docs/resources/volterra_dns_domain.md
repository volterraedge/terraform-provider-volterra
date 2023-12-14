---

page_title: "Volterra: dns_domain"

description: "The dns_domain allows CRUD of Dns Domain resource on Volterra SaaS"
---------------------------------------------------------------------------------

Resource volterra_dns_domain
============================

The Dns Domain allows CRUD of Dns Domain resource on Volterra SaaS

~> **Note:** Please refer to [Dns Domain API docs](https://docs.cloud.f5.com/docs/api/dns-domain) to learn more

Example Usage
-------------

```hcl
resource "volterra_dns_domain" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "verification_only volterra_managed route53" must be set
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

`route53` - (Optional) sub domain in Amazon Route 53 zone owned by users. See [Route53 ](#route53) below for details.

`verification_only` - (Optional) F5XC will verify this domain, but will not manage it. (bool).

`volterra_managed` - (Optional) sub domain (bool).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Route53

sub domain in Amazon Route 53 zone owned by users.

`creds` - (Optional) Reference to AWS credentials to program route53. See [ref](#ref) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured dns_domain.
