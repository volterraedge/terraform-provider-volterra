---

page_title: "Volterra: app_type"

description: "The app_type allows CRUD of App Type resource on Volterra SaaS"
-----------------------------------------------------------------------------

Resource volterra_app_type
==========================

The App Type allows CRUD of App Type resource on Volterra SaaS

~> **Note:** Please refer to [App Type API docs](https://docs.cloud.f5.com/docs/api/app-type) to learn more

Example Usage
-------------

```hcl
resource "volterra_app_type" "example" {
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

`business_logic_markup_setting` - (Optional) Setting specifying how API Discovery will be performed. See [Business Logic Markup Setting ](#business-logic-markup-setting) below for details.

`features` - (Optional) List of various AI/ML features enabled. See [Features ](#features) below for details.

### Business Logic Markup Setting

Setting specifying how API Discovery will be performed.

`discovered_api_settings` - (Optional) x-displayName: "Discovered API Settings". See [Business Logic Markup Setting Discovered Api Settings ](#business-logic-markup-setting-discovered-api-settings) below for details.

###### One of the arguments from this list "disable, enable" can be set

`disable` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`enable` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (`Bool`).

`sensitive_data_detection_rules` - (Optional) Rule to detect sensitive data in requests and/or response sections.. See [Business Logic Markup Setting Sensitive Data Detection Rules ](#business-logic-markup-setting-sensitive-data-detection-rules) below for details.(Deprecated)

### Features

List of various AI/ML features enabled.

`type` - (Required) Feature type to be enabled (`String`).

### Business Logic Markup Setting Discovered Api Settings

x-displayName: "Discovered API Settings".

`purge_duration_for_inactive_discovered_apis` - (Optional) Inactive discovered API will be deleted after configured duration. (`Int`).

### Business Logic Markup Setting Sensitive Data Detection Rules

Rule to detect sensitive data in requests and/or response sections..

`custom_sensitive_data_detection_rules` - (Optional) Rules to detect custom sensitive data in requests and/or responses sections.. See [Sensitive Data Detection Rules Custom Sensitive Data Detection Rules ](#sensitive-data-detection-rules-custom-sensitive-data-detection-rules) below for details.

`disabled_built_in_rules` - (Optional) List of disabled built-in sensitive data detection rules.. See [Sensitive Data Detection Rules Disabled Built In Rules ](#sensitive-data-detection-rules-disabled-built-in-rules) below for details.

### Custom Sensitive Data Detection Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Custom Sensitive Data Detection Rules Sensitive Data Detection Config

The custom data detection config specifies targets, scopes & the pattern to be detected..

###### One of the arguments from this list "any_domain, specific_domain" must be set

`any_domain` - (Optional) The rule will apply for all domains. (`Bool`).(Deprecated)

`specific_domain` - (Optional) For example: api.example.com (`String`).(Deprecated)

###### One of the arguments from this list "key_pattern, value_pattern, key_value_pattern" must be set

`key_pattern` - (Optional) Search for pattern across all field names in the specified sections.. See [Pattern Choice Key Pattern ](#pattern-choice-key-pattern) below for details.

`key_value_pattern` - (Optional) Search for specific field and value patterns in the specified sections.. See [Pattern Choice Key Value Pattern ](#pattern-choice-key-value-pattern) below for details.

`value_pattern` - (Optional) Search for pattern across all field values in the specified sections.. See [Pattern Choice Value Pattern ](#pattern-choice-value-pattern) below for details.

###### One of the arguments from this list "all_sections, all_request_sections, all_response_sections, custom_sections" must be set

`all_request_sections` - (Optional) x-displayName: "All Request" (`Bool`).

`all_response_sections` - (Optional) x-displayName: "All Response" (`Bool`).

`all_sections` - (Optional) x-displayName: "All Request & Response" (`Bool`).

`custom_sections` - (Optional) x-displayName: "Custom Sections". See [Section Choice Custom Sections ](#section-choice-custom-sections) below for details.

###### One of the arguments from this list "api_group, any_target, api_endpoint_target, base_path" must be set

`any_target` - (Optional) The rule will be applied for all requests on this LB. (`Bool`).

`api_endpoint_target` - (Optional) The rule is applied only for the specified api endpoints.. See [Target Choice Api Endpoint Target ](#target-choice-api-endpoint-target) below for details.

`api_group` - (Optional) Custom groups can be created if user tags paths or operations with "x-volterra-api-group" extensions inside swaggers. (`String`).(Deprecated)

`base_path` - (Optional) The rule is applied only for the requests matching the specified base path. (`String`).(Deprecated)

### Custom Sensitive Data Detection Rules Sensitive Data Type

If the pattern is detected, the request is labeled with specified sensitive data type..

`type` - (Required) The request is labeled as specified sensitive data type. (`String`).

### Domain Choice Any Domain

The rule will apply for all domains..

### Key Value Pattern Key Pattern

Pattern for key/field..

###### One of the arguments from this list "exact_value, regex_value" must be set

`exact_value` - (Optional) Search for values with exact match. (`String`).

`regex_value` - (Optional) Search for values matching this regular expression. (`String`).

### Key Value Pattern Value Pattern

Pattern for value..

###### One of the arguments from this list "regex_value, exact_value" must be set

`exact_value` - (Optional) Pattern value to be detected. (`String`).

`regex_value` - (Optional) Regular expression for this pattern. (`String`).

### Learn From Redirect Traffic Disable

Disable learning API patterns from traffic with redirect response codes 3xx.

### Learn From Redirect Traffic Enable

Enable learning API patterns from traffic with redirect response codes 3xx.

### Pattern Choice Key Pattern

Search for pattern across all field names in the specified sections..

###### One of the arguments from this list "exact_value, regex_value" must be set

`exact_value` - (Optional) Search for values with exact match. (`String`).

`regex_value` - (Optional) Search for values matching this regular expression. (`String`).

### Pattern Choice Key Value Pattern

Search for specific field and value patterns in the specified sections..

`key_pattern` - (Required) Pattern for key/field.. See [Key Value Pattern Key Pattern ](#key-value-pattern-key-pattern) below for details.

`value_pattern` - (Required) Pattern for value.. See [Key Value Pattern Value Pattern ](#key-value-pattern-value-pattern) below for details.

### Pattern Choice Value Pattern

Search for pattern across all field values in the specified sections..

###### One of the arguments from this list "exact_value, regex_value" must be set

`exact_value` - (Optional) Pattern value to be detected. (`String`).

`regex_value` - (Optional) Regular expression for this pattern. (`String`).

### Section Choice All Request Sections

x-displayName: "All Request".

### Section Choice All Response Sections

x-displayName: "All Response".

### Section Choice All Sections

x-displayName: "All Request & Response".

### Section Choice Custom Sections

x-displayName: "Custom Sections".

`custom_sections` - (Required) Request & Response Sections. (`List of Strings`).

### Sensitive Data Detection Rules Custom Sensitive Data Detection Rules

Rules to detect custom sensitive data in requests and/or responses sections..

`metadata` - (Required) Common attributes for the rule including name and description.. See [Custom Sensitive Data Detection Rules Metadata ](#custom-sensitive-data-detection-rules-metadata) below for details.

`sensitive_data_detection_config` - (Required) The custom data detection config specifies targets, scopes & the pattern to be detected.. See [Custom Sensitive Data Detection Rules Sensitive Data Detection Config ](#custom-sensitive-data-detection-rules-sensitive-data-detection-config) below for details.

`sensitive_data_type` - (Required) If the pattern is detected, the request is labeled with specified sensitive data type.. See [Custom Sensitive Data Detection Rules Sensitive Data Type ](#custom-sensitive-data-detection-rules-sensitive-data-type) below for details.

### Sensitive Data Detection Rules Disabled Built In Rules

List of disabled built-in sensitive data detection rules..

`name` - (Required) Built-in rule for sensitive data detection. (`String`).

### Target Choice Any Target

The rule will be applied for all requests on this LB..

### Target Choice Api Endpoint Target

The rule is applied only for the specified api endpoints..

`api_endpoint_path` - (Required) The rule is applied only for the specified api endpoints. (`String`).

`methods` - (Required) x-example: "['GET', 'POST', 'DELETE']" (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured app_type.
