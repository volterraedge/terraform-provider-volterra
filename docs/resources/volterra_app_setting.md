---

page_title: "Volterra: app_setting"

description: "The app_setting allows CRUD of App Setting resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_app_setting
=============================

The App Setting allows CRUD of App Setting resource on Volterra SaaS

~> **Note:** Please refer to [App Setting API docs](https://docs.cloud.f5.com/docs-v2/api/app-setting) to learn more

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
      // One of the arguments from this list "disable enable" can be set

      enable = true
    }

    timeseries_analyses_setting {
      metric_selectors {
        metric = ["metric"]

        metrics_source = "metrics_source"
      }
    }

    user_behavior_analysis_setting {
      // One of the arguments from this list "disable_learning enable_learning" must be set

      enable_learning = true

      // One of the arguments from this list "disable_detection enable_detection" must be set

      enable_detection {
        // One of the arguments from this list "bola_detection_automatic bola_detection_manual exclude_bola_detection" can be set

        exclude_bola_detection = true

        // One of the arguments from this list "exclude_bot_defense_activity include_bot_defense_activity" must be set

        include_bot_defense_activity = true

        // One of the arguments from this list "cooling_off_period" must be set

        cooling_off_period = "cooling_off_period"

        // One of the arguments from this list "exclude_failed_login_activity include_failed_login_activity" must be set

        include_failed_login_activity {
          login_failures_threshold = "10"
        }

        // One of the arguments from this list "exclude_forbidden_activity include_forbidden_activity" must be set

        include_forbidden_activity {
          forbidden_requests_threshold = "10"
        }

        // One of the arguments from this list "exclude_ip_reputation include_ip_reputation" must be set

        include_ip_reputation = true

        // One of the arguments from this list "exclude_non_existent_url_activity include_non_existent_url_activity_automatic include_non_existent_url_activity_custom" can be set

        exclude_non_existent_url_activity = true

        // One of the arguments from this list "exclude_rate_limit include_rate_limit" must be set

        include_rate_limit = true

        // One of the arguments from this list "exclude_waf_activity include_waf_activity" must be set

        include_waf_activity = true
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

`anomaly_types` - (Optional) List of Anomaly algorithms that need to be enabled (`List of Strings`).(Deprecated)

`app_type_refs` - (Optional) List of references to app_type for which monitoring needs to enabled. See [ref](#ref) below for details.(Deprecated)

`app_type_settings` - (Required) List of settings to enable for each AppType, given instance of AppType Exist in this Namespace. See [App Type Settings ](#app-type-settings) below for details.

### App Type Settings

List of settings to enable for each AppType, given instance of AppType Exist in this Namespace.

`app_type_ref` - (Required) Associating an AppType reference, will enable analysis on this instance's generated data. See [ref](#ref) below for details.

`business_logic_markup_setting` - (Optional) Setting specifying how API Discovery will be performed.. See [App Type Settings Business Logic Markup Setting ](#app-type-settings-business-logic-markup-setting) below for details.

`timeseries_analyses_setting` - (Optional) The clients are flagged if anomalies are observed. See [App Type Settings Timeseries Analyses Setting ](#app-type-settings-timeseries-analyses-setting) below for details.

`user_behavior_analysis_setting` - (Optional) The risk score of the user decays over time, if no further suspicious activity is noted. See [App Type Settings User Behavior Analysis Setting ](#app-type-settings-user-behavior-analysis-setting) below for details.

### App Type Settings Business Logic Markup Setting

Setting specifying how API Discovery will be performed..

###### One of the arguments from this list "disable, enable" can be set

`disable` - (Optional) API Endpoints are not discovered in this namespace (`Bool`).

`enable` - (Optional) API Endpoints are discovered in this namespace (`Bool`).

### App Type Settings Timeseries Analyses Setting

The clients are flagged if anomalies are observed.

`metric_selectors` - (Optional) be included in the detection logic. See [Timeseries Analyses Setting Metric Selectors ](#timeseries-analyses-setting-metric-selectors) below for details.

### App Type Settings User Behavior Analysis Setting

The risk score of the user decays over time, if no further suspicious activity is noted.

###### One of the arguments from this list "disable_learning, enable_learning" must be set

`disable_learning` - (Optional) Disable learning user behavior patterns from this namespace (`Bool`).

`enable_learning` - (Optional) Enable learning user behavior patterns from this namespace (`Bool`).

###### One of the arguments from this list "disable_detection, enable_detection" must be set

`disable_detection` - (Optional) Disable malicious user detection (`Bool`).

`enable_detection` - (Optional) Enable AI based malicious user detection. See [Malicious User Detection Enable Detection ](#malicious-user-detection-enable-detection) below for details.

### Bola Activity Choice Bola Detection Automatic

Detect Enumeration attack automatically..

### Bola Activity Choice Bola Detection Manual

Detect Enumeration attack using user defined threshold..

###### One of the arguments from this list "threshold_level_1, threshold_level_2, threshold_level_3, threshold_level_4, threshold_level_5, threshold_level_6" must be set

`threshold_level_1` - (Optional) Detected in range: 10 - 150 (`Bool`).

`threshold_level_2` - (Optional) Detected in range: 25 - 400 (`Bool`).

`threshold_level_3` - (Optional) Detected in range: 50 - 800 (`Bool`).

`threshold_level_4` - (Optional) Detected in range: 100 - 1500 (`Bool`).

`threshold_level_5` - (Optional) Detected in range: 200 - 3000 (`Bool`).

`threshold_level_6` - (Optional) Detected in range: 500 - 8000 (`Bool`).

### Bola Activity Choice Exclude Bola Detection

Disable Enumeration attack detection.

### Bot Defense Activity Choice Exclude Bot Defense Activity

Exclude Bot Defense activity in malicious user detection.

### Bot Defense Activity Choice Include Bot Defense Activity

Include Bot Defense activity in malicious user detection.

### Failed Login Activity Choice Exclude Failed Login Activity

Exclude persistent login failures activity (401 response code) in malicious user detection.

### Failed Login Activity Choice Include Failed Login Activity

Include persistent login failures activity (401 response code) in malicious user detection.

`login_failures_threshold` - (Required) The number of failed logins beyond which the system will flag this user as malicious (`Int`).

### Forbidden Activity Choice Exclude Forbidden Activity

Exclude forbidden activity by policy in malicious user detection.

### Forbidden Activity Choice Include Forbidden Activity

Include forbidden activity by policy in malicious user detection.

`forbidden_requests_threshold` - (Required) The number of forbidden requests beyond which the system will flag this user as malicious (`Int`).

### Ip Reputation Choice Exclude Ip Reputation

Exclude IP Reputation in malicious user detection.

### Ip Reputation Choice Include Ip Reputation

Include IP Reputation in malicious user detection.

### Learn From Namespace Disable

API Endpoints are not discovered in this namespace.

### Learn From Namespace Disable Learning

Disable learning user behavior patterns from this namespace.

### Learn From Namespace Enable

API Endpoints are discovered in this namespace.

### Learn From Namespace Enable Learning

Enable learning user behavior patterns from this namespace.

### Malicious User Detection Disable Detection

Disable malicious user detection.

### Malicious User Detection Enable Detection

Enable AI based malicious user detection.

###### One of the arguments from this list "bola_detection_automatic, bola_detection_manual, exclude_bola_detection" can be set

`bola_detection_automatic` - (Optional) Detect Enumeration attack automatically. (`Bool`).(Deprecated)

`bola_detection_manual` - (Optional) Detect Enumeration attack using user defined threshold.. See [Bola Activity Choice Bola Detection Manual ](#bola-activity-choice-bola-detection-manual) below for details.(Deprecated)

`exclude_bola_detection` - (Optional) Disable Enumeration attack detection (`Bool`).(Deprecated)

###### One of the arguments from this list "exclude_bot_defense_activity, include_bot_defense_activity" must be set

`exclude_bot_defense_activity` - (Optional) Exclude Bot Defense activity in malicious user detection (`Bool`).

`include_bot_defense_activity` - (Optional) Include Bot Defense activity in malicious user detection (`Bool`).

###### One of the arguments from this list "cooling_off_period" must be set

`cooling_off_period` - (Optional) a high to medium or medium to low or low to none. (`Int`).

###### One of the arguments from this list "exclude_failed_login_activity, include_failed_login_activity" must be set

`exclude_failed_login_activity` - (Optional) Exclude persistent login failures activity (401 response code) in malicious user detection (`Bool`).

`include_failed_login_activity` - (Optional) Include persistent login failures activity (401 response code) in malicious user detection. See [Failed Login Activity Choice Include Failed Login Activity ](#failed-login-activity-choice-include-failed-login-activity) below for details.

###### One of the arguments from this list "exclude_forbidden_activity, include_forbidden_activity" must be set

`exclude_forbidden_activity` - (Optional) Exclude forbidden activity by policy in malicious user detection (`Bool`).

`include_forbidden_activity` - (Optional) Include forbidden activity by policy in malicious user detection. See [Forbidden Activity Choice Include Forbidden Activity ](#forbidden-activity-choice-include-forbidden-activity) below for details.

###### One of the arguments from this list "exclude_ip_reputation, include_ip_reputation" must be set

`exclude_ip_reputation` - (Optional) Exclude IP Reputation in malicious user detection (`Bool`).

`include_ip_reputation` - (Optional) Include IP Reputation in malicious user detection (`Bool`).

###### One of the arguments from this list "exclude_non_existent_url_activity, include_non_existent_url_activity_automatic, include_non_existent_url_activity_custom" can be set

`exclude_non_existent_url_activity` - (Optional) Exclude Non-Existent URL activity in malicious user detection (`Bool`).(Deprecated)

`include_non_existent_url_activity_automatic` - (Optional) Include Non-Existent URL Activity using automatic threshold in malicious user detection. See [Non Existent Url Activity Choice Include Non Existent Url Activity Automatic ](#non-existent-url-activity-choice-include-non-existent-url-activity-automatic) below for details.(Deprecated)

`include_non_existent_url_activity_custom` - (Optional) Include Non-Existent URL Activity using custom threshold in malicious user detection. See [Non Existent Url Activity Choice Include Non Existent Url Activity Custom ](#non-existent-url-activity-choice-include-non-existent-url-activity-custom) below for details.(Deprecated)

###### One of the arguments from this list "exclude_rate_limit, include_rate_limit" must be set

`exclude_rate_limit` - (Optional) Exclude Rate Limiting in malicious user detection (`Bool`).

`include_rate_limit` - (Optional) Include Rate Limiting in malicious user detection (`Bool`).

###### One of the arguments from this list "exclude_waf_activity, include_waf_activity" must be set

`exclude_waf_activity` - (Optional) Exclude WAF activity in malicious user detection (`Bool`).

`include_waf_activity` - (Optional) Include WAF activity in malicious user detection (`Bool`).

### Non Existent Url Activity Choice Exclude Non Existent Url Activity

Exclude Non-Existent URL activity in malicious user detection.

### Non Existent Url Activity Choice Include Non Existent Url Activity Automatic

Include Non-Existent URL Activity using automatic threshold in malicious user detection.

###### One of the arguments from this list "high, low, medium" must be set

`high` - (Optional) Use auto-calculated threshold decreased by margin for more sensitive detection (`Bool`).

`low` - (Optional) Use auto-calculated threshold with margin for less sensitive detection (`Bool`).

`medium` - (Optional) Use auto-calculated threshold learnt from statistics per given application (`Bool`).

### Non Existent Url Activity Choice Include Non Existent Url Activity Custom

Include Non-Existent URL Activity using custom threshold in malicious user detection.

`nonexistent_requests_threshold` - (Required) The percentage of non-existent requests beyond which the system will flag this user as malicious (`Int`).

### Rate Limit Choice Exclude Rate Limit

Exclude Rate Limiting in malicious user detection.

### Rate Limit Choice Include Rate Limit

Include Rate Limiting in malicious user detection.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Sensitivity High

Use auto-calculated threshold decreased by margin for more sensitive detection.

### Sensitivity Low

Use auto-calculated threshold with margin for less sensitive detection.

### Sensitivity Medium

Use auto-calculated threshold learnt from statistics per given application.

### Threshold Levels Threshold Level 1

Detected in range: 10 - 150.

### Threshold Levels Threshold Level 2

Detected in range: 25 - 400.

### Threshold Levels Threshold Level 3

Detected in range: 50 - 800.

### Threshold Levels Threshold Level 4

Detected in range: 100 - 1500.

### Threshold Levels Threshold Level 5

Detected in range: 200 - 3000.

### Threshold Levels Threshold Level 6

Detected in range: 500 - 8000.

### Timeseries Analyses Setting Metric Selectors

be included in the detection logic.

`metric` - (Optional) Choose one or more metrics to be included in the detection logic (`List of Strings`).

`metrics_source` - (Optional) Choose the source for the metrics to be included in the detection logic (`String`).

### Waf Activity Choice Exclude Waf Activity

Exclude WAF activity in malicious user detection.

### Waf Activity Choice Include Waf Activity

Include WAF activity in malicious user detection.

Attribute Reference
-------------------

-	`id` - This is the id of the configured app_setting.
