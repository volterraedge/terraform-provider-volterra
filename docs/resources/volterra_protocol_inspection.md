---

page_title: "Volterra: protocol_inspection"

description: "The protocol_inspection allows CRUD of Protocol Inspection resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------

Resource volterra_protocol_inspection
=====================================

The Protocol Inspection allows CRUD of Protocol Inspection resource on Volterra SaaS

~> **Note:** Please refer to [Protocol Inspection API docs](https://docs.cloud.f5.com/docs/api/protocol-inspection) to learn more

Example Usage
-------------

```hcl
resource "volterra_protocol_inspection" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  enable_disable_compliance_checks {
    // One of the arguments from this list "enable_compliance_checks disable_compliance_checks" must be set

    enable_compliance_checks {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  enable_disable_signatures {
    // One of the arguments from this list "disable_signature enable_signature" must be set

    enable_signature = true
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

`action` - (Optional) Action to take based on inspection (`String`).

`enable_disable_compliance_checks` - (Required) Enable or Disable Compliance Checks. See [Enable Disable Compliance Checks ](#enable-disable-compliance-checks) below for details.

`enable_disable_signatures` - (Required) Confirmation of applying IPS. See [Enable Disable Signatures ](#enable-disable-signatures) below for details.

### Enable Disable Compliance Checks

Enable or Disable Compliance Checks.

###### One of the arguments from this list "enable_compliance_checks, disable_compliance_checks" must be set

`disable_compliance_checks` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_compliance_checks` - (Optional) Enabling compliance checks and selecting a type by default (right now this enables dns_compliance_checks). See [ref](#ref) below for details.

### Enable Disable Signatures

Confirmation of applying IPS.

###### One of the arguments from this list "enable_signature, disable_signature" must be set

`disable_signature` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_signature` - (Optional) x-displayName: "Enable" (`Bool`).

### Compliance Check Choice Disable Compliance Checks

x-displayName: "Disable".

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Signature Choice Disable Signature

x-displayName: "Disable".

### Signature Choice Enable Signature

x-displayName: "Enable".

Attribute Reference
-------------------

-	`id` - This is the id of the configured protocol_inspection.
