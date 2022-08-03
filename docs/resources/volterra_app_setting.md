---

page_title: "Volterra: app_setting"

description: "The app_setting allows CRUD of App Setting resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_app_setting
=============================

The App Setting allows CRUD of App Setting resource on Volterra SaaS

~> **Note:** Please refer to [App Setting API docs](https://volterra.io/docs/api/app-setting) to learn more

Example Usage
-------------

```hcl
resource "volterra_app_setting" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  app_type_settings {
    app_type_ref {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }

    business_logic_markup_setting {
      // One of the arguments from this list "disable enable" must be set
      enable = true
    }

    timeseries_analyses_setting {
      metric_selectors {
        metric         = ["metric"]
        metrics_source = "metrics_source"
      }
    }

    user_behavior_analysis_setting {
      // One of the arguments from this list "enable_learning disable_learning" must be set
      enable_learning = true

      // One of the arguments from this list "enable_detection disable_detection" must be set

      enable_detection {
        // One of the arguments from this list "cooling_off_period" must be set
        cooling_off_period = "cooling_off_period"

        // One of the arguments from this list "include_failed_login_activity exclude_failed_login_activity" must be set

        include_failed_login_activity {
          login_failures_threshold = "10"
        }

        // One of the arguments from this list "include_forbidden_activity exclude_forbidden_activity" must be set

        include_forbidden_activity {
          forbidden_requests_threshold = "10"
        }
        // One of the arguments from this list "include_ip_reputation exclude_ip_reputation" must be set
        include_ip_reputation = true
        // One of the arguments from this list "include_waf_activity exclude_waf_activity" must be set
        exclude_waf_activity = true
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

`anomaly_types` - (Optional) List of Anomaly algorithms that need to be enabled (`List of Strings`).

`app_type_refs` - (Optional) List of references to app_type for which monitoring needs to enabled. See [ref](#ref) below for details.

`app_type_settings` - (Required) List of settings to enable for each AppType, given instance of AppType Exist in this Namespace. See [App Type Settings ](#app-type-settings) below for details.

### App Type Settings

List of settings to enable for each AppType, given instance of AppType Exist in this Namespace.

`app_type_ref` - (Required) Associating an AppType reference, will enable analysis on this instance's generated data. See [ref](#ref) below for details.

`business_logic_markup_setting` - (Optional) Setting specifying how Business Logic Markup will be performed.. See [Business Logic Markup Setting ](#business-logic-markup-setting) below for details.

`timeseries_analyses_setting` - (Optional) Configuration for timeseries analyses. See [Timeseries Analyses Setting ](#timeseries-analyses-setting) below for details.

`user_behavior_analysis_setting` - (Optional) Configuration for user behavior analysis in this namespace. See [User Behavior Analysis Setting ](#user-behavior-analysis-setting) below for details.

### Business Logic Markup Setting

Setting specifying how Business Logic Markup will be performed..

`disable` - (Optional) Disable learning API patterns from this namespace (bool).

`enable` - (Optional) Enable learning API patterns from this namespace (bool).

### Disable

Disable learning API patterns from this namespace.

### Disable Detection

Disable malicious user detection.

### Disable Learning

Disable learning user behavior patterns from this namespace.

### Enable

Enable learning API patterns from this namespace.

### Enable Detection

Enable AI based malicious user detection.

`cooling_off_period` - (Required) a high to medium or medium to low or low to none. (`Int`).

`exclude_failed_login_activity` - (Optional) Exclude persistent login failures activity (401 response code) in malicious user detection (bool).

`include_failed_login_activity` - (Optional) Include persistent login failures activity (401 response code) in malicious user detection. See [Include Failed Login Activity ](#include-failed-login-activity) below for details.

`exclude_forbidden_activity` - (Optional) Exclude forbidden activity by policy in malicious user detection (bool).

`include_forbidden_activity` - (Optional) Include forbidden activity by policy in malicious user detection. See [Include Forbidden Activity ](#include-forbidden-activity) below for details.

`exclude_ip_reputation` - (Optional) Exclude IP Reputation by policy in malicious user detection (bool).

`include_ip_reputation` - (Optional) Include IP Reputation by policy in malicious user detection (bool).

`exclude_waf_activity` - (Optional) Exclude WAF activity in malicious user detection (bool).

`include_waf_activity` - (Optional) Include WAF activity in malicious user detection (bool).

### Enable Learning

Enable learning user behavior patterns from this namespace.

### Exclude Failed Login Activity

Exclude persistent login failures activity (401 response code) in malicious user detection.

### Exclude Forbidden Activity

Exclude forbidden activity by policy in malicious user detection.

### Exclude Ip Reputation

Exclude IP Reputation by policy in malicious user detection.

### Exclude Waf Activity

Exclude WAF activity in malicious user detection.

### Include Failed Login Activity

Include persistent login failures activity (401 response code) in malicious user detection.

`login_failures_threshold` - (Required) The number of failed logins beyond which the system will flag this user as malicious (`Int`).

### Include Forbidden Activity

Include forbidden activity by policy in malicious user detection.

`forbidden_requests_threshold` - (Required) The number of forbidden requests beyond which the system will flag this user as malicious (`Int`).

### Include Ip Reputation

Include IP Reputation by policy in malicious user detection.

### Include Waf Activity

Include WAF activity in malicious user detection.

### Metric Selectors

List of metric selectors where each selector specifies which metrics are selected to be analyzed.

`metric` - (Optional) Metrics enabled to be analyzed (`List of Strings`).

`metrics_source` - (Optional) Source from where this metric is gathered for analyses (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Timeseries Analyses Setting

Configuration for timeseries analyses.

`metric_selectors` - (Optional) List of metric selectors where each selector specifies which metrics are selected to be analyzed. See [Metric Selectors ](#metric-selectors) below for details.

### User Behavior Analysis Setting

Configuration for user behavior analysis in this namespace.

`disable_learning` - (Optional) Disable learning user behavior patterns from this namespace (bool).

`enable_learning` - (Optional) Enable learning user behavior patterns from this namespace (bool).

`disable_detection` - (Optional) Disable malicious user detection (bool).

`enable_detection` - (Optional) Enable AI based malicious user detection. See [Enable Detection ](#enable-detection) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured app_setting.
