---

page_title: "Volterra: dns_compliance_checks"
description: "The dns_compliance_checks allows CRUD of Dns Compliance Checks resource on Volterra SaaS"

---

Resource volterra_dns_compliance_checks
=======================================

The Dns Compliance Checks allows CRUD of Dns Compliance Checks resource on Volterra SaaS

~> **Note:** Please refer to [Dns Compliance Checks API docs](https://docs.cloud.f5.com/docs-v2/api/dns-compliance-checks) to learn more

Example Usage
-------------

```hcl
resource "volterra_dns_compliance_checks" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  domain_denylist = ["www.f5.com"]
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

`disallowed_query_type_list` - (Optional) Disallowed Query Type Values (`List of Strings`).

`disallowed_resource_record_type_list` - (Optional) Disallowed Resource Record Type List (`List of Strings`).

`domain_denylist` - (Required) List of domains to be denied by configuration object (`List of String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured dns_compliance_checks.
