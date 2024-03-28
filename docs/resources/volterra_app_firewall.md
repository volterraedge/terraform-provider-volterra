---

page_title: "Volterra: app_firewall"

description: "The app_firewall allows CRUD of App Firewall resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_app_firewall
==============================

The App Firewall allows CRUD of App Firewall resource on Volterra SaaS

~> **Note:** Please refer to [App Firewall API docs](https://docs.cloud.f5.com/docs/api/app-firewall) to learn more

Example Usage
-------------

```hcl
resource "volterra_app_firewall" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "allow_all_response_codes allowed_response_codes" must be set
  allow_all_response_codes = true

  // One of the arguments from this list "disable_anonymization default_anonymization custom_anonymization" must be set
  disable_anonymization = true

  // One of the arguments from this list "use_default_blocking_page blocking_page" must be set
  use_default_blocking_page = true

  // One of the arguments from this list "default_bot_setting bot_protection_setting" must be set
  default_bot_setting = true

  // One of the arguments from this list "default_detection_settings detection_settings" must be set

  detection_settings {
    // One of the arguments from this list "enable_suppression disable_suppression" must be set
    enable_suppression = true

    signature_selection_setting {
      // One of the arguments from this list "default_attack_type_settings attack_type_settings" must be set
      default_attack_type_settings = true

      // One of the arguments from this list "only_high_accuracy_signatures high_medium_accuracy_signatures high_medium_low_accuracy_signatures" must be set
      only_high_accuracy_signatures = true
    }

    // One of the arguments from this list "disable_staging stage_new_signatures stage_new_and_updated_signatures" must be set

    stage_new_and_updated_signatures {
      staging_period = "7"
    }
    // One of the arguments from this list "enable_threat_campaigns disable_threat_campaigns" must be set
    disable_threat_campaigns = true
    // One of the arguments from this list "default_violation_settings violation_settings" must be set
    default_violation_settings = true
  }
  // One of the arguments from this list "use_loadbalancer_setting blocking monitoring" must be set
  use_loadbalancer_setting = true
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

`allow_all_response_codes` - (Optional) All HTTP response status codes are allowed (`Bool`).

`allowed_response_codes` - (Optional) Define list of HTTP response status codes that are allowed. See [Allowed Response Codes Choice Allowed Response Codes ](#allowed-response-codes-choice-allowed-response-codes) below for details.

`custom_anonymization` - (Optional) Define HTTP headers, query parameters, or cookies whose values should be masked. See [Anonymization Setting Custom Anonymization ](#anonymization-setting-custom-anonymization) below for details.

`default_anonymization` - (Optional) Values of query parameters "card", "pass", "pwd" and "password" will be masked. (`Bool`).

`disable_anonymization` - (Optional) Disable masking of sensitive parameters in logs (`Bool`).

`blocking_page` - (Optional) The system returns a response page with HTML code that you define. See [Blocking Page Choice Blocking Page ](#blocking-page-choice-blocking-page) below for details.

`use_default_blocking_page` - (Optional) The system returns the system-supplied response page in HTML. No further configuration is needed. (`Bool`).

`bot_protection_setting` - (Optional) Define custom Bot Protection settings. See [Bot Protection Choice Bot Protection Setting ](#bot-protection-choice-bot-protection-setting) below for details.

`default_bot_setting` - (Optional) Malicious bots will be blocked, Suspicious and Good bots will be reported. (`Bool`).

`default_detection_settings` - (Optional) All Attack Types, high and medium accuracy signatures, automatic Attack Signatures tuning, Threat Campaigns and all Violations will be enabled. (`Bool`).

`detection_settings` - (Optional) Define Custom Security Policy settings. See [Detection Setting Choice Detection Settings ](#detection-setting-choice-detection-settings) below for details.

`blocking` - (Optional) Log and block threats (`Bool`).

`monitoring` - (Optional) Log threats (`Bool`).

`use_loadbalancer_setting` - (Optional) Use the mode as specified in the load balancer (`Bool`).(Deprecated)

### Allowed Response Codes Choice Allowed Response Codes

Define list of HTTP response status codes that are allowed.

`response_code` - (Required) List of HTTP response status codes that are allowed (`Int`).

### Anonymization Choice Cookie

x-displayName: "Cookie".

`cookie_name` - (Required) Masks the cookie value. The setting does not mask the cookie name. (`String`).

### Anonymization Choice Http Header

x-displayName: "HTTP Header".

`header_name` - (Required) Masks the HTTP header value. The setting does not mask the HTTP header name. (`String`).

### Anonymization Choice Query Parameter

x-displayName: "Query Parameter".

`query_param_name` - (Required) Masks the query parameter value. The setting does not mask the query parameter name. (`String`).

### Anonymization Setting Custom Anonymization

Define HTTP headers, query parameters, or cookies whose values should be masked.

`anonymization_config` - (Required) List of HTTP headers, cookies and query parameters whose values will be masked. See [Custom Anonymization Anonymization Config ](#custom-anonymization-anonymization-config) below for details.

### Attack Type Setting Attack Type Settings

Define Attack Types to be disabled for detection.

`disabled_attack_types` - (Required) List of Attack Types that will be ignored and not trigger a detection (`List of Strings`).

### Attack Type Setting Default Attack Type Settings

All Attack Types are enabled for detection.

### Blocking Page Choice Blocking Page

The system returns a response page with HTML code that you define.

`blocking_page` - (Optional) which would be about 3070 bytes in plain text. (`String`).

`response_code` - (Optional) HTTP status code to be sent for blocked requests (`String`).

### Bot Protection Choice Bot Protection Setting

Define custom Bot Protection settings.

`good_bot_action` - (Optional) A client that exhibits known search engine behaviors and signatures (`String`).

`malicious_bot_action` - (Optional) A client that exhibits malicious intent such as DoS tools, known exploit tools, and vulnerability scanners (`String`).

`suspicious_bot_action` - (Optional) A client that exhibits non-malicious tools such as site crawlers, monitors, spiders, web downloaders and bots behaviors, signatures such as search bots and social media agents (`String`).

### Custom Anonymization Anonymization Config

List of HTTP headers, cookies and query parameters whose values will be masked.

###### One of the arguments from this list "http_header, query_parameter, cookie" must be set

`cookie` - (Optional) x-displayName: "Cookie". See [Anonymization Choice Cookie ](#anonymization-choice-cookie) below for details.

`http_header` - (Optional) x-displayName: "HTTP Header". See [Anonymization Choice Http Header ](#anonymization-choice-http-header) below for details.

`query_parameter` - (Optional) x-displayName: "Query Parameter". See [Anonymization Choice Query Parameter ](#anonymization-choice-query-parameter) below for details.

### Detection Setting Choice Detection Settings

Define Custom Security Policy settings.

###### One of the arguments from this list "enable_suppression, disable_suppression" must be set

`disable_suppression` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_suppression` - (Optional) x-displayName: "Enable" (`Bool`).

`signature_selection_setting` - (Optional) Attack Signatures are patterns that identify attacks on a web application and its components. See [Detection Settings Signature Selection Setting ](#detection-settings-signature-selection-setting) below for details.

###### One of the arguments from this list "disable_staging, stage_new_signatures, stage_new_and_updated_signatures" can be set

`disable_staging` - (Optional) Enforce new and updated attack signatures (`Bool`).

`stage_new_and_updated_signatures` - (Optional) would not enforce it i.e signature would be in monitoring mode for staging period (instead of blocking mode). See [Signatures Staging Settings Stage New And Updated Signatures ](#signatures-staging-settings-stage-new-and-updated-signatures) below for details.

`stage_new_signatures` - (Optional) Stage new attack signatures only. Updated signatures will be enforced. See [Signatures Staging Settings Stage New Signatures ](#signatures-staging-settings-stage-new-signatures) below for details.

###### One of the arguments from this list "enable_threat_campaigns, disable_threat_campaigns" must be set

`disable_threat_campaigns` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_threat_campaigns` - (Optional) x-displayName: "Enable" (`Bool`).

###### One of the arguments from this list "default_violation_settings, violation_settings" must be set

`default_violation_settings` - (Optional) All violations are enabled for detection (`Bool`).

`violation_settings` - (Optional) Define violations to be disabled for detection. See [Violation Detection Setting Violation Settings ](#violation-detection-setting-violation-settings) below for details.

### Detection Settings Signature Selection Setting

Attack Signatures are patterns that identify attacks on a web application and its components.

###### One of the arguments from this list "default_attack_type_settings, attack_type_settings" must be set

`attack_type_settings` - (Optional) Define Attack Types to be disabled for detection. See [Attack Type Setting Attack Type Settings ](#attack-type-setting-attack-type-settings) below for details.

`default_attack_type_settings` - (Optional) All Attack Types are enabled for detection (`Bool`).

###### One of the arguments from this list "only_high_accuracy_signatures, high_medium_accuracy_signatures, high_medium_low_accuracy_signatures" must be set

`high_medium_accuracy_signatures` - (Optional) Enables high and medium accuracy signatures (`Bool`).

`high_medium_low_accuracy_signatures` - (Optional) Enables high, medium and low accuracy signatures (`Bool`).

`only_high_accuracy_signatures` - (Optional) Enables only high accuracy signatures (`Bool`).

### False Positive Suppression Disable Suppression

x-displayName: "Disable".

### False Positive Suppression Enable Suppression

x-displayName: "Enable".

### Signature Selection By Accuracy High Medium Accuracy Signatures

Enables high and medium accuracy signatures.

### Signature Selection By Accuracy High Medium Low Accuracy Signatures

Enables high, medium and low accuracy signatures.

### Signature Selection By Accuracy Only High Accuracy Signatures

Enables only high accuracy signatures.

### Signatures Staging Settings Disable Staging

Enforce new and updated attack signatures.

### Signatures Staging Settings Stage New And Updated Signatures

would not enforce it i.e signature would be in monitoring mode for staging period (instead of blocking mode).

`staging_period` - (Required) 20 days. (`Int`).

### Signatures Staging Settings Stage New Signatures

Stage new attack signatures only. Updated signatures will be enforced.

`staging_period` - (Required) 20 days. (`Int`).

### Threat Campaign Choice Disable Threat Campaigns

x-displayName: "Disable".

### Threat Campaign Choice Enable Threat Campaigns

x-displayName: "Enable".

### Violation Detection Setting Default Violation Settings

All violations are enabled for detection.

### Violation Detection Setting Violation Settings

Define violations to be disabled for detection.

`disabled_violation_types` - (Required) List of violations to be excluded (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured app_firewall.
