---

page_title: "Volterra: uztna_policy_template"

description: "The uztna_policy_template allows CRUD of Uztna Policy Template resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------------

Resource volterra_uztna_policy_template
=======================================

The Uztna Policy Template allows CRUD of Uztna Policy Template resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Policy Template API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-policy-template) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_policy_template" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  on_start_flow {
    // One of the arguments from this list "simple" must be set

    simple {
      flows {
        // One of the arguments from this list "geolocation_match saml_federation" can be set

        saml_federation {
          saml {
            name      = "test1"
            namespace = "staging"
            tenant    = "acmecorp"
          }
        }
      }
    }
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

`on_start_flow` - (Required) throughout the entirety of the session. See [On Start Flow ](#on-start-flow) below for details.

### On Start Flow

throughout the entirety of the session.

###### One of the arguments from this list "simple" must be set

`simple` - (Optional) This template . See [Template Type Simple ](#template-type-simple) below for details.

### Flow Type Geolocation Match

flows of type Geolocation Match.

`geomatch` - (Required) flows of type GEOLocation Match. See [ref](#ref) below for details.

### Flow Type Saml Federation

flows of type SAML Federation.

`saml` - (Required) flows of type SAML Federation. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Simple Flows

Or you can also Empty Flow.

###### One of the arguments from this list "geolocation_match, saml_federation" can be set

`geolocation_match` - (Optional) flows of type Geolocation Match. See [Flow Type Geolocation Match ](#flow-type-geolocation-match) below for details.

`saml_federation` - (Optional) flows of type SAML Federation. See [Flow Type Saml Federation ](#flow-type-saml-federation) below for details.

### Template Type Simple

This template .

`flows` - (Required) Or you can also Empty Flow. See [Simple Flows ](#simple-flows) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_policy_template.
