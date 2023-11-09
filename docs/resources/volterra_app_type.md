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

### All Request Sections

x-displayName: "All Request".

### All Response Sections

x-displayName: "All Response".

### All Sections

x-displayName: "All Request & Response".

### Any Domain

The rule will apply for all domains..

### Any Target

The rule will be applied for all requests on this LB..

### Api Endpoint Target

The rule is applied only for the specified api endpoints..

`api_endpoint_path` - (Required) The rule is applied only for the specified api endpoints. (`String`).

`methods` - (Required) x-example: "['GET', 'POST', 'DELETE']" (`List of Strings`).

### Business Logic Markup Setting

Setting specifying how API Discovery will be performed.

`discovered_api_settings` - (Optional) x-displayName: "Discovered API Settings". See [Discovered Api Settings ](#discovered-api-settings) below for details.

`disable` - (Optional) Disable learning API patterns from traffic with redirect response codes 3xx (bool).

`enable` - (Optional) Enable learning API patterns from traffic with redirect response codes 3xx (bool).

`sensitive_data_detection_rules` - (Optional) Rule to detect sensitive data in requests and/or response sections.. See [Sensitive Data Detection Rules ](#sensitive-data-detection-rules) below for details.

### Custom Sections

x-displayName: "Custom Sections".

`custom_sections` - (Required) Request & Response Sections. (`List of Strings`).

### Custom Sensitive Data Detection Rules

Rules to detect custom sensitive data in requests and/or responses sections..

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`sensitive_data_detection_config` - (Required) The custom data detection config specifies targets, scopes & the pattern to be detected.. See [Sensitive Data Detection Config ](#sensitive-data-detection-config) below for details.

`sensitive_data_type` - (Required) If the pattern is detected, the request is labeled with specified sensitive data type.. See [Sensitive Data Type ](#sensitive-data-type) below for details.

### Disable

Disable learning API patterns from traffic with redirect response codes 3xx.

### Disabled Built In Rules

List of disabled built-in sensitive data detection rules..

`name` - (Required) Built-in rule for sensitive data detection. (`String`).

### Discovered Api Settings

x-displayName: "Discovered API Settings".

`purge_duration_for_inactive_discovered_apis` - (Optional) Inactive discovered API will be deleted after configured duration. (`Int`).

### Enable

Enable learning API patterns from traffic with redirect response codes 3xx.

### Features

List of various AI/ML features enabled.

`type` - (Required) Feature type to be enabled (`String`).

### Key Pattern

Search for pattern across all field names in the specified sections..

`exact_value` - (Optional) Search for values with exact match. (`String`).

`regex_value` - (Optional) Search for values matching this regular expression. (`String`).

### Key Value Pattern

Search for specific field and value patterns in the specified sections..

`key_pattern` - (Required) Pattern for key/field.. See [Key Pattern ](#key-pattern) below for details.

`value_pattern` - (Required) Pattern for value.. See [Value Pattern ](#value-pattern) below for details.

### Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Sensitive Data Detection Config

The custom data detection config specifies targets, scopes & the pattern to be detected..

`any_domain` - (Required) The rule will apply for all domains. (bool).

`specific_domain` - (Required) For example: api.example.com (`String`).

`key_pattern` - (Required) Search for pattern across all field names in the specified sections.. See [Key Pattern ](#key-pattern) below for details.

`key_value_pattern` - (Required) Search for specific field and value patterns in the specified sections.. See [Key Value Pattern ](#key-value-pattern) below for details.

`value_pattern` - (Required) Search for pattern across all field values in the specified sections.. See [Value Pattern ](#value-pattern) below for details.

`all_request_sections` - (Optional) x-displayName: "All Request" (bool).

`all_response_sections` - (Optional) x-displayName: "All Response" (bool).

`all_sections` - (Optional) x-displayName: "All Request & Response" (bool).

`custom_sections` - (Optional) x-displayName: "Custom Sections". See [Custom Sections ](#custom-sections) below for details.

`any_target` - (Required) The rule will be applied for all requests on this LB. (bool).

`api_endpoint_target` - (Required) The rule is applied only for the specified api endpoints.. See [Api Endpoint Target ](#api-endpoint-target) below for details.

`api_group` - (Optional) Custom groups can be created if user tags paths or operations with "x-volterra-api-group" extensions inside swaggers. (`String`).

`base_path` - (Required) The rule is applied only for the requests matching the specified base path. (`String`).

### Sensitive Data Detection Rules

Rule to detect sensitive data in requests and/or response sections..

`custom_sensitive_data_detection_rules` - (Optional) Rules to detect custom sensitive data in requests and/or responses sections.. See [Custom Sensitive Data Detection Rules ](#custom-sensitive-data-detection-rules) below for details.

`disabled_built_in_rules` - (Optional) List of disabled built-in sensitive data detection rules.. See [Disabled Built In Rules ](#disabled-built-in-rules) below for details.

### Sensitive Data Type

If the pattern is detected, the request is labeled with specified sensitive data type..

`type` - (Required) The request is labeled as specified sensitive data type. (`String`).

### Value Pattern

Pattern for value..

`exact_value` - (Optional) Pattern value to be detected. (`String`).

`regex_value` - (Optional) Regular expression for this pattern. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured app_type.
