











---
page_title: "Volterra: sensitive_data_policy"
description: "The sensitive_data_policy allows CRUD of Sensitive Data Policy  resource on Volterra SaaS"
---
# Resource volterra_sensitive_data_policy

The Sensitive Data Policy  allows CRUD of Sensitive Data Policy  resource on Volterra SaaS

~> **Note:** Please refer to [Sensitive Data Policy  API docs](https://docs.cloud.f5.com/docs-v2/api/sensitive-data-policy) to learn more

## Example Usage

```hcl
resource "volterra_sensitive_data_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
}

```

## Argument Reference

### Metadata Argument Reference
`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).


`description` - (Optional) Human readable description for the object (`String`).


`disable` - (Optional) A value of true will administratively disable the object (`Bool`).


`labels` - (Optional) by selector expression (`String`).


`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).


`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).



### Spec Argument Reference
`compliances` - (Optional) Select relevant compliance frameworks, such as GDPR, HIPAA, or PCI-DSS, to ensure monitoring under your sensitive data discovery. (`List of Strings`).(Deprecated)



`custom_data_types` - (Optional) Select your custom data types to be monitored in the API discovery. See [Custom Data Types ](#custom-data-types) below for details.




`disabled_predefined_data_types` - (Optional) Select which pre-configured data types to disable, disabled data types will not be shown as sensitive in the API discovery (`List of String`).



### Custom Data Types 

 Select your custom data types to be monitored in the API discovery.

`custom_data_type_ref` - (Optional) List of custom data types to monitor. See [ref](#ref) below for details.



### Ref 


Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).



## Attribute Reference

* `id` - This is the id of the configured sensitive_data_policy.

