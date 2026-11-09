---

page_title: "Volterra: quic_profile"

description: "The quic_profile allows CRUD of Quic Profile resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_quic_profile
==============================

The Quic Profile allows CRUD of Quic Profile resource on Volterra SaaS

~> **Note:** Please refer to [Quic Profile API docs](https://docs.cloud.f5.com/docs-v2/api/quic-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_quic_profile" "example" {
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

`app_service` - (Optional) The application service to which the object belongs. (`String`).

`bidi_concurrent_streams_per_connection` - (Optional) Specifies how many bidirectional concurrent streams are allowed to be outstanding on a single QUIC connection. (`Int`).

`command` - (Optional) x-displayName: "Command" (`String`).

`spin_bit` - (Optional) Marks the spin bit in the QUIC header to signal path round-trip time to observers. (`String`).

`sub_path` - (Optional) x-displayName: "Sub Path" (`String`).

`uni_concurrent_streams_per_connection` - (Optional) Specifies how many unidirectional concurrent streams are allowed to be outstanding on a single QUIC connection. (`Int`).

`user_spec` - (Optional) User specified properties. (`List of String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured quic_profile.
