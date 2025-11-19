---

page_title: "Volterra: filter_set"

description: "The filter_set allows CRUD of Filter Set resource on Volterra SaaS"
---------------------------------------------------------------------------------

Resource volterra_filter_set
============================

The Filter Set allows CRUD of Filter Set resource on Volterra SaaS

~> **Note:** Please refer to [Filter Set API docs](https://docs.cloud.f5.com/docs-v2/api/filter-set) to learn more

Example Usage
-------------

```hcl
resource "volterra_filter_set" "example" {
  name        = "acmecorp-web"
  namespace   = "staging"
  context_key = ["context_key"]

  filter_fields {
    field_id = "field_id"

    // One of the arguments from this list "date_field filter_expression_field label_selector_field string_field" must be set

    string_field {
      field_values = ["field_values"]
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

`context_key` - (Required) indexable context key that identifies a page or page type for which the FilterSet is applicable (`String`).

`filter_fields` - (Required) list of fields and their values selected by the user. See [Filter Fields ](#filter-fields) below for details.

### Filter Fields

list of fields and their values selected by the user.

`field_id` - (Required) an identifier for the field that maps to some UI filter component (`String`).

###### One of the arguments from this list "date_field, filter_expression_field, label_selector_field, string_field" must be set

`date_field` - (Optional) x-displayName: "Date/Time Range". See [Field Value Date Field ](#field-value-date-field) below for details.

`filter_expression_field` - (Optional) x-displayName: "Key/Value Selector Expression". See [Field Value Filter Expression Field ](#field-value-filter-expression-field) below for details.

`label_selector_field` - (Optional) x-displayName: "Kubernetes-style Label Selector Expression (deprecated)". See [Field Value Label Selector Field ](#field-value-label-selector-field) below for details.(Deprecated)

`string_field` - (Optional) x-displayName: "String". See [Field Value String Field ](#field-value-string-field) below for details.

### Field Value Date Field

x-displayName: "Date/Time Range".

###### One of the arguments from this list "absolute, relative" must be set

`absolute` - (Optional) absolute start and end timestamps. See [Range Type Absolute ](#range-type-absolute) below for details.

`relative` - (Optional) relative time duration (`String`).

### Field Value Filter Expression Field

x-displayName: "Key/Value Selector Expression".

`expression` - (Required) differs in that it allows special characters in the keys and values (`String`).

### Field Value Label Selector Field

x-displayName: "Kubernetes-style Label Selector Expression (deprecated)".

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Field Value String Field

x-displayName: "String".

`field_values` - (Required) x-required (`String`).

### Range Type Absolute

absolute start and end timestamps.

`end_date` - (Required) Contains end date (`String`).

`start_date` - (Required) Contains start date (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured filter_set.
