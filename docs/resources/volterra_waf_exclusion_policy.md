---

page_title: "Volterra: waf_exclusion_policy"

description: "The waf_exclusion_policy allows CRUD of Waf Exclusion Policy resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------

Resource volterra_waf_exclusion_policy
======================================

The Waf Exclusion Policy allows CRUD of Waf Exclusion Policy resource on Volterra SaaS

~> **Note:** Please refer to [Waf Exclusion Policy API docs](https://docs.cloud.f5.com/docs-v2/api/waf-exclusion-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_waf_exclusion_policy" "example" {
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

`waf_exclusion_rules` - (Optional) An ordered list of rules.. See [Waf Exclusion Rules ](#waf-exclusion-rules) below for details.

### Waf Exclusion Rules

An ordered list of rules..

###### One of the arguments from this list "any_domain, exact_value, suffix_value" must be set

`any_domain` - (Optional) Apply this WAF exclusion rule for any domain (`Bool`).

`exact_value` - (Optional) Exact domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

`expiration_timestamp` - (Optional) the configuration but is not applied anymore. (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Waf Exclusion Rules Metadata ](#waf-exclusion-rules-metadata) below for details.

`methods` - (Optional) methods to be matched (`List of Strings`).

###### One of the arguments from this list "any_path, path_prefix, path_regex" must be set

`any_path` - (Optional) Match all paths (`Bool`).

`path_prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`path_regex` - (Optional) Define the regex for the path. For example, the regex ^/.*$ will match on all paths (`String`).

###### One of the arguments from this list "app_firewall_detection_control, waf_skip_processing" can be set

`app_firewall_detection_control` - (Optional) Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.. See [Waf Advanced Configuration App Firewall Detection Control ](#waf-advanced-configuration-app-firewall-detection-control) below for details.

`waf_skip_processing` - (Optional) Skip all App Firewall processing for this request (`Bool`).

### App Firewall Detection Control Exclude Attack Type Contexts

Attack Types to be excluded for the defined match criteria.

`context` - (Required) x-required (`String`).

`context_name` - (Optional) with an wildcard asterisk (*). (`String`).

`exclude_attack_type` - (Required) x-required (`String`).

### App Firewall Detection Control Exclude Bot Name Contexts

Bot Names to be excluded for the defined match criteria.

`bot_name` - (Required) x-example: "Hydra" (`String`).

### App Firewall Detection Control Exclude Signature Contexts

Signature IDs to be excluded for the defined match criteria.

`context` - (Required) x-required (`String`).

`context_name` - (Optional) with an wildcard asterisk (*). (`String`).

`signature_id` - (Required) 0 implies that all signatures will be excluded for the specified context. (`Int`).

### App Firewall Detection Control Exclude Violation Contexts

Violations to be excluded for the defined match criteria.

`context` - (Required) x-required (`String`).

`context_name` - (Optional) with an wildcard asterisk (*). (`String`).

`exclude_violation` - (Required) x-required (`String`).

### Domain Choice Any Domain

Apply this WAF exclusion rule for any domain.

### Path Choice Any Path

Match all paths.

### Waf Advanced Configuration App Firewall Detection Control

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria..

`exclude_attack_type_contexts` - (Optional) Attack Types to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Attack Type Contexts ](#app-firewall-detection-control-exclude-attack-type-contexts) below for details.

`exclude_bot_name_contexts` - (Optional) Bot Names to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Bot Name Contexts ](#app-firewall-detection-control-exclude-bot-name-contexts) below for details.

`exclude_signature_contexts` - (Optional) Signature IDs to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Signature Contexts ](#app-firewall-detection-control-exclude-signature-contexts) below for details.

`exclude_violation_contexts` - (Optional) Violations to be excluded for the defined match criteria. See [App Firewall Detection Control Exclude Violation Contexts ](#app-firewall-detection-control-exclude-violation-contexts) below for details.

### Waf Advanced Configuration Waf Skip Processing

Skip all App Firewall processing for this request.

### Waf Exclusion Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured waf_exclusion_policy.
