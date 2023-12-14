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

  allowed_response_codes {
    response_code = ["[200, 201, 204, 300, 302, 400, 403, 404, 500, 501, 503]"]
  }
  // One of the arguments from this list "disable_anonymization default_anonymization custom_anonymization" must be set
  disable_anonymization = true
  // One of the arguments from this list "use_default_blocking_page blocking_page" must be set
  use_default_blocking_page = true
  // One of the arguments from this list "default_bot_setting bot_protection_setting" must be set
  default_bot_setting = true
  // One of the arguments from this list "default_detection_settings detection_settings" must be set
  default_detection_settings = true
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

`allow_all_response_codes` - (Optional) All HTTP response status codes are allowed (bool).

`allowed_response_codes` - (Optional) Define list of HTTP response status codes that are allowed. See [Allowed Response Codes ](#allowed-response-codes) below for details.

`custom_anonymization` - (Optional) Define HTTP headers, query parameters, or cookies whose values should be masked. See [Custom Anonymization ](#custom-anonymization) below for details.

`default_anonymization` - (Optional) Values of query parameters "card", "pass", "pwd" and "password" will be masked. (bool).

`disable_anonymization` - (Optional) Disable masking of sensitive parameters in logs (bool).

`blocking_page` - (Optional) The system returns a response page with HTML code that you define. See [Blocking Page ](#blocking-page) below for details.

`use_default_blocking_page` - (Optional) The system returns the system-supplied response page in HTML. No further configuration is needed. (bool).

`bot_protection_setting` - (Optional) Define custom Bot Protection settings. See [Bot Protection Setting ](#bot-protection-setting) below for details.

`default_bot_setting` - (Optional) Malicious bots will be blocked, Suspicious and Good bots will be reported. (bool).

`default_detection_settings` - (Optional) All Attack Types, high and medium accuracy signatures, automatic Attack Signatures tuning, Threat Campaigns and all Violations will be enabled. (bool).

`detection_settings` - (Optional) Define Custom Security Policy settings. See [Detection Settings ](#detection-settings) below for details.

`blocking` - (Optional) Log and block threats (bool).

`monitoring` - (Optional) Log threats (bool).

`use_loadbalancer_setting` - (Optional) Use the mode as specified in the load balancer (bool).

### Allowed Response Codes

Define list of HTTP response status codes that are allowed.

`response_code` - (Required) List of HTTP response status codes that are allowed (`Int`).

### Anonymization Config

List of HTTP headers, cookies and query parameters whose values will be masked.

`cookie` - (Optional) x-displayName: "Cookie". See [Cookie ](#cookie) below for details.

`http_header` - (Optional) x-displayName: "HTTP Header". See [Http Header ](#http-header) below for details.

`query_parameter` - (Optional) x-displayName: "Query Parameter". See [Query Parameter ](#query-parameter) below for details.

### Attack Type Settings

Define Attack Types to be disabled for detection.

`disabled_attack_types` - (Required) List of Attack Types that will be ignored and not trigger a detection (`List of Strings`).

### Blocking Page

The system returns a response page with HTML code that you define.

`blocking_page` - (Optional) Response body can't exceed 4 KB in size. (`String`).

`response_code` - (Optional) HTTP status code to be sent for blocked requests (`String`).

### Bot Protection Setting

Define custom Bot Protection settings.

`good_bot_action` - (Optional) A client that exhibits known search engine behaviors and signatures (`String`).

`malicious_bot_action` - (Optional) A client that exhibits malicious intent such as DoS tools, known exploit tools, and vulnerability scanners (`String`).

`suspicious_bot_action` - (Optional) A client that exhibits non-malicious tools such as site crawlers, monitors, spiders, web downloaders and bots behaviors, signatures such as search bots and social media agents (`String`).

### Cookie

x-displayName: "Cookie".

`cookie_name` - (Required) Masks the cookie value. The setting does not mask the cookie name. (`String`).

### Custom Anonymization

Define HTTP headers, query parameters, or cookies whose values should be masked.

`anonymization_config` - (Required) List of HTTP headers, cookies and query parameters whose values will be masked. See [Anonymization Config ](#anonymization-config) below for details.

### Default Attack Type Settings

All Attack Types are enabled for detection.

### Default Violation Settings

All violations are enabled for detection.

### Detection Settings

Define Custom Security Policy settings.

`disable_suppression` - (Optional) x-displayName: "Disable" (bool).

`enable_suppression` - (Optional) x-displayName: "Enable" (bool).

`signature_selection_setting` - (Optional) Attack Signatures are patterns that identify attacks on a web application and its components. See [Signature Selection Setting ](#signature-selection-setting) below for details.

`disable_staging` - (Optional) Enforce new and updated attack signatures (bool).

`stage_new_and_updated_signatures` - (Optional) would not enforce it i.e signature would be in monitoring mode for staging period (instead of blocking mode). See [Stage New And Updated Signatures ](#stage-new-and-updated-signatures) below for details.

`stage_new_signatures` - (Optional) Stage new attack signatures only. Updated signatures will be enforced. See [Stage New Signatures ](#stage-new-signatures) below for details.

`disable_threat_campaigns` - (Optional) x-displayName: "Disable" (bool).

`enable_threat_campaigns` - (Optional) x-displayName: "Enable" (bool).

`default_violation_settings` - (Optional) All violations are enabled for detection (bool).

`violation_settings` - (Optional) Define violations to be disabled for detection. See [Violation Settings ](#violation-settings) below for details.

### Disable Staging

Enforce new and updated attack signatures.

### Disable Suppression

x-displayName: "Disable".

### Disable Threat Campaigns

x-displayName: "Disable".

### Enable Suppression

x-displayName: "Enable".

### Enable Threat Campaigns

x-displayName: "Enable".

### High Medium Accuracy Signatures

Enables high and medium accuracy signatures.

### High Medium Low Accuracy Signatures

Enables high, medium and low accuracy signatures.

### Http Header

x-displayName: "HTTP Header".

`header_name` - (Required) Masks the HTTP header value. The setting does not mask the HTTP header name. (`String`).

### Only High Accuracy Signatures

Enables only high accuracy signatures.

### Query Parameter

x-displayName: "Query Parameter".

`query_param_name` - (Required) Masks the query parameter value. The setting does not mask the query parameter name. (`String`).

### Signature Selection Setting

Attack Signatures are patterns that identify attacks on a web application and its components.

`attack_type_settings` - (Optional) Define Attack Types to be disabled for detection. See [Attack Type Settings ](#attack-type-settings) below for details.

`default_attack_type_settings` - (Optional) All Attack Types are enabled for detection (bool).

`high_medium_accuracy_signatures` - (Optional) Enables high and medium accuracy signatures (bool).

`high_medium_low_accuracy_signatures` - (Optional) Enables high, medium and low accuracy signatures (bool).

`only_high_accuracy_signatures` - (Optional) Enables only high accuracy signatures (bool).

### Stage New And Updated Signatures

would not enforce it i.e signature would be in monitoring mode for staging period (instead of blocking mode).

`staging_period` - (Required) 20 days. (`Int`).

### Stage New Signatures

Stage new attack signatures only. Updated signatures will be enforced.

`staging_period` - (Required) 20 days. (`Int`).

### Violation Settings

Define violations to be disabled for detection.

`disabled_violation_types` - (Required) List of violations to be excluded (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured app_firewall.
