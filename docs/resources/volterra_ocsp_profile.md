---

page_title: "Volterra: ocsp_profile"

description: "The ocsp_profile allows CRUD of Ocsp Profile resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_ocsp_profile
==============================

The Ocsp Profile allows CRUD of Ocsp Profile resource on Volterra SaaS

~> **Note:** Please refer to [Ocsp Profile API docs](https://docs.cloud.f5.com/docs-v2/api/ocsp-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_ocsp_profile" "example" {
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

`full_path` - (Optional) x-displayName: "Full Path" (`String`).

`max_age` - (Optional) Specifies the value for HTTP Response cache control header max-age. The max-age header sent to client is the lesser of the configured value and validity of OCSP response. Default value is 604800 seconds. (`Int`).

`nonce` - (Optional) Specifies whether OCSP Nonce Request Extension is supported by the OCSP profile. Default value is Enabled. (`String`).

`sub_path` - (Optional) x-displayName: "Sub Path" (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured ocsp_profile.
