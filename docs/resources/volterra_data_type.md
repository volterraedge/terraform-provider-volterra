---

page_title: "Volterra: data_type"
description: "The data_type allows CRUD of Data Type resource on Volterra SaaS"

---

Resource volterra_data_type
===========================

The Data Type allows CRUD of Data Type resource on Volterra SaaS

~> **Note:** Please refer to [Data Type API docs](https://docs.cloud.f5.com/docs-v2/api/data-type) to learn more

Example Usage
-------------

```hcl
resource "volterra_data_type" "example" {
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

`compliances` - (Optional) Choose applicable compliance frameworks such as GDPR, PCI/DSS, or CCPA to ensure the platform identifies whether vulnerabilities in API endpoints handling this data type may cause a compliance breach (`List of Strings`).

`is_pii` - (Optional) Select this option to classify the custom data type as personally identifiable information (PII) (`Bool`).

`is_sensitive_data` - (Optional) Select this option to classify the custom data type as sensitive, enabling detection of API vulnerabilities related to this data type. (`Bool`).

`rules` - (Optional) Configure key/value or regex match rules to enable the platform to detect this custom data type in the API request or response. See [Rules ](#rules) below for details.

### Rules

Configure key/value or regex match rules to enable the platform to detect this custom data type in the API request or response.

###### One of the arguments from this list "key_pattern, key_value_pattern, value_pattern" must be set

`key_pattern` - (Optional) Search for pattern across all field names in the specified sections.. See [Pattern Choice Key Pattern ](#pattern-choice-key-pattern) below for details.

`key_value_pattern` - (Optional) Search for specific field and value patterns in the specified sections.. See [Pattern Choice Key Value Pattern ](#pattern-choice-key-value-pattern) below for details.

`value_pattern` - (Optional) Search for pattern across all field values in the specified sections.. See [Pattern Choice Value Pattern ](#pattern-choice-value-pattern) below for details.

### Key Value Pattern Key Pattern

Pattern for key/field..

###### One of the arguments from this list "exact_values, regex_value, substring_value" must be set

`exact_values` - (Optional) Search for values that are exact match to at least one item of the list. See [Type Choice Exact Values ](#type-choice-exact-values) below for details.

`regex_value` - (Optional) Search for values matching this regular expression (`String`).

`substring_value` - (Optional) Search for values that include this substring (`String`).

### Key Value Pattern Value Pattern

Pattern for value..

###### One of the arguments from this list "exact_values, regex_value, substring_value" must be set

`exact_values` - (Optional) Search for values that are exact match to at least one item of the list. See [Type Choice Exact Values ](#type-choice-exact-values) below for details.

`regex_value` - (Optional) Search for values matching this regular expression (`String`).

`substring_value` - (Optional) Search for values that include this substring (`String`).

### Pattern Choice Key Pattern

Search for pattern across all field names in the specified sections..

###### One of the arguments from this list "exact_values, regex_value, substring_value" must be set

`exact_values` - (Optional) Search for values that are exact match to at least one item of the list. See [Type Choice Exact Values ](#type-choice-exact-values) below for details.

`regex_value` - (Optional) Search for values matching this regular expression (`String`).

`substring_value` - (Optional) Search for values that include this substring (`String`).

### Pattern Choice Key Value Pattern

Search for specific field and value patterns in the specified sections..

`key_pattern` - (Required) Pattern for key/field.. See [Key Value Pattern Key Pattern ](#key-value-pattern-key-pattern) below for details.

`value_pattern` - (Required) Pattern for value.. See [Key Value Pattern Value Pattern ](#key-value-pattern-value-pattern) below for details.

### Pattern Choice Value Pattern

Search for pattern across all field values in the specified sections..

###### One of the arguments from this list "exact_values, regex_value, substring_value" must be set

`exact_values` - (Optional) Search for values that are exact match to at least one item of the list. See [Type Choice Exact Values ](#type-choice-exact-values) below for details.

`regex_value` - (Optional) Search for values matching this regular expression (`String`).

`substring_value` - (Optional) Search for values that include this substring (`String`).

### Type Choice Exact Values

Search for values that are exact match to at least one item of the list.

`exact_values` - (Required) List of exact values to match. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured data_type.
