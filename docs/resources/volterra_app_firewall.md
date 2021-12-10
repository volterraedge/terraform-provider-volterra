---

page_title: "Volterra: app_firewall"

description: "The app_firewall allows CRUD of App Firewall resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_app_firewall
==============================

The App Firewall allows CRUD of App Firewall resource on Volterra SaaS

~> **Note:** Please refer to [App Firewall API docs](https://volterra.io/docs/api/app-firewall) to learn more

Example Usage
-------------

```hcl
resource "volterra_app_firewall" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "allow_all_response_codes allowed_response_codes" must be set
  allow_all_response_codes = true

  // One of the arguments from this list "default_anonymization custom_anonymization disable_anonymization" must be set
  default_anonymization = true

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

`allowed_response_codes` - (Optional) List of HTTP response status codes that are allowed. See [Allowed Response Codes ](#allowed-response-codes) below for details.

`custom_anonymization` - (Optional) WAF performs full case-insensitive match.. See [Custom Anonymization ](#custom-anonymization) below for details.

`default_anonymization` - (Optional) Values of query parameters "card", "pass", "pwd" and "password" will be masked. (bool).

`disable_anonymization` - (Optional) Disable masking of sensitive parameters in logs. (bool).

`blocking_page` - (Optional) Use custom blocking response page. See [Blocking Page ](#blocking-page) below for details.

`use_default_blocking_page` - (Optional) Use default blocking response page (bool).

`bot_protection_setting` - (Optional) Bot protection configuration, mitigation actions for the different types of bots. See [Bot Protection Setting ](#bot-protection-setting) below for details.

`default_bot_setting` - (Optional) Suspicious and Good bots will be reported. (bool).

`default_detection_settings` - (Optional) These settings are recommended for detecting malicious requests with a low false positive rate. (bool).

`detection_settings` - (Optional) Custom detection settings. See [Detection Settings ](#detection-settings) below for details.

`blocking` - (Optional) Override the load balancer setting and set mode to blocking (bool).

`monitoring` - (Optional) Override the load balancer setting and set mode to monitoring (bool).

`use_loadbalancer_setting` - (Optional) Use the mode as specified in the load balancer (bool).

### Allowed Response Codes

List of HTTP response status codes that are allowed.

`response_code` - (Required) response code to send (`Int`).

### Anonymization Config

List of HTTP headers, cookies and query parameters which values will be masked..

`cookie` - (Optional) Anonymize sensitive HTTP Cookie information. See [Cookie ](#cookie) below for details.

`http_header` - (Optional) Anonymize sensitive HTTP Header information. See [Http Header ](#http-header) below for details.

`query_parameter` - (Optional) Anonymize sensitive HTTP query parameters. See [Query Parameter ](#query-parameter) below for details.

### Attack Type Settings

Specifies attack-type settings to be used by WAF.

`disabled_attack_types` - (Required) List of attack types that will be ignored and not trigger a detection (`List of Strings`).

### Blocking Page

Use custom blocking response page.

`blocking_page` - (Optional) Custom blocking page response body (`String`).

### Bot Protection Setting

Bot protection configuration, mitigation actions for the different types of bots.

`good_bot_action` - (Optional) Action to be taken when a good (benign) bot agent is detected (`String`).

`malicious_bot_action` - (Optional) Action to be taken when a malicious bot agent is detected (`String`).

`suspicious_bot_action` - (Optional) Action to be taken when a suspicious bot agent is detected (`String`).

### Cookie

Anonymize sensitive HTTP Cookie information.

`cookie_name` - (Required) Specify name of the sensitive HTTP Cookie (`String`).

### Custom Anonymization

WAF performs full case-insensitive match..

`anonymization_config` - (Required) List of HTTP headers, cookies and query parameters which values will be masked.. See [Anonymization Config ](#anonymization-config) below for details.

### Default Attack Type Settings

Use default settings. This will enable all attack types for detection..

### Default Violation Settings

Use default violation settings. This will enable all violation types for detection..

### Detection Settings

Custom detection settings.

`disable_suppression` - (Optional) Disable Suppression (bool).

`enable_suppression` - (Optional) Enable Suppression (bool).

`signature_selection_setting` - (Optional) Specifies signature settings to be used by WAF. See [Signature Selection Setting ](#signature-selection-setting) below for details.

`disable_threat_campaigns` - (Optional) Disable Threat Campaigns (bool).

`enable_threat_campaigns` - (Optional) Enable Threat Campaigns (bool).

`default_violation_settings` - (Optional) Use default violation settings. This will enable all violation types for detection. (bool).

`violation_settings` - (Required) Specifies violation settings to be used by WAF. See [Violation Settings ](#violation-settings) below for details.

### Disable Suppression

Disable Suppression.

### Disable Threat Campaigns

Disable Threat Campaigns.

### Enable Suppression

Enable Suppression.

### Enable Threat Campaigns

Enable Threat Campaigns.

### High Medium Accuracy Signatures

Enables high and medium accuracy signatures.

### High Medium Low Accuracy Signatures

Enables high, medium and low accuracy signatures.

### Http Header

Anonymize sensitive HTTP Header information.

`header_name` - (Required) Specify name of the sensitive HTTP Header (`String`).

### Only High Accuracy Signatures

Enables only high accuracy signatures.

### Query Parameter

Anonymize sensitive HTTP query parameters.

`query_param_name` - (Required) Specify name of the sensitive HTTP Query Parameter (`String`).

### Signature Selection Setting

Specifies signature settings to be used by WAF.

`attack_type_settings` - (Optional) Specifies attack-type settings to be used by WAF. See [Attack Type Settings ](#attack-type-settings) below for details.

`default_attack_type_settings` - (Optional) Use default settings. This will enable all attack types for detection. (bool).

`high_medium_accuracy_signatures` - (Optional) Enables high and medium accuracy signatures (bool).

`high_medium_low_accuracy_signatures` - (Optional) Enables high, medium and low accuracy signatures (bool).

`only_high_accuracy_signatures` - (Optional) Enables only high accuracy signatures (bool).

### Violation Settings

Specifies violation settings to be used by WAF.

`disabled_violation_types` - (Required) List of violations to be excluded (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured app_firewall.
