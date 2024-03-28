---

page_title: "Volterra: alert_policy"

description: "The alert_policy allows CRUD of Alert Policy resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_alert_policy
==============================

The Alert Policy allows CRUD of Alert Policy resource on Volterra SaaS

~> **Note:** Please refer to [Alert Policy API docs](https://docs.cloud.f5.com/docs/api/alert-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_alert_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  receivers {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }

  routes {
    // One of the arguments from this list "send dont_send" must be set
    send = true

    // One of the arguments from this list "severity group alertname alertname_regex custom any" must be set

    group {
      groups = ["groups"]
    }
    notification_parameters {
      // One of the arguments from this list "default individual ves_io_group custom" must be set

      custom {
        labels = ["value"]
      }

      group_interval = "1m"

      group_wait = "30s"

      repeat_interval = "4h"
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

`notification_parameters` - (Optional) Notification parameters to decide how and when the alerts should be sent to the receivers.. See [Notification Parameters ](#notification-parameters) below for details.

`receivers` - (Required) list of Alert Receivers where the alerts will be sent. See [ref](#ref) below for details.

`routes` - (Required) The routes are evaluated in the specified order and terminates on the first match.. See [Routes ](#routes) below for details.

### Notification Parameters

Notification parameters to decide how and when the alerts should be sent to the receivers..

###### One of the arguments from this list "ves_io_group, custom, default, individual" must be set

`custom` - (Optional) Specify set of labels for grouping the alerts. See [Group By Custom ](#group-by-custom) below for details.

`default` - (Optional) Group the alerts by severity, group name and alert name (`Bool`).

`individual` - (Optional) This option disables grouping of alerts (`Bool`).

`ves_io_group` - (Optional) Group the alerts by severity, group name and alert name (`Bool`).

`group_interval` - (Optional) If not specified, group_interval defaults to "1m" (`String`).

`group_wait` - (Optional) If not specified, group_wait defaults to "30s" (`String`).

`repeat_interval` - (Optional) If not specified, group_interval defaults to "4h" (`String`).

### Routes

The routes are evaluated in the specified order and terminates on the first match..

###### One of the arguments from this list "send, dont_send" must be set

`dont_send` - (Optional) Do not send the alert (`Bool`).

`send` - (Optional) Send the alert (`Bool`).

###### One of the arguments from this list "any, severity, group, alertname, alertname_regex, custom" can be set

`alertname` - (Optional) Matches the alertname of the alert (`String`).

`alertname_regex` - (Optional) Regular Expression match for the alertname (`String`).

`any` - (Optional) Matches all alerts in the namespace (`Bool`).

`custom` - (Optional) A set of custom equality/regex matchers an alert has to fulfill to match the route.. See [Matcher Custom ](#matcher-custom) below for details.

`group` - (Optional) Matches the group name of the alert. See [Matcher Group ](#matcher-group) below for details.

`severity` - (Optional) Matches the severity level of the alert. See [Matcher Severity ](#matcher-severity) below for details.

`notification_parameters` - (Optional) notification_config defined in the policy.. See [Routes Notification Parameters ](#routes-notification-parameters) below for details.

### Action Dont Send

Do not send the alert.

### Action Send

Send the alert.

### Custom Alertlabel

AlertLabel to configure the alert policy rule.

###### One of the arguments from this list "exact_match, regex_match" must be set

`exact_match` - (Optional) Equality match value for the label (`String`).

`regex_match` - (Optional) Regular expression match value for the label (`String`).

### Custom Alertname

Alertname Matcher.

###### One of the arguments from this list "regex_match, exact_match" must be set

`exact_match` - (Optional) Equality match value for the label (`String`).

`regex_match` - (Optional) Regular expression match value for the label (`String`).

### Custom Group

Group Matcher.

###### One of the arguments from this list "exact_match, regex_match" must be set

`exact_match` - (Optional) Equality match value for the label (`String`).

`regex_match` - (Optional) Regular expression match value for the label (`String`).

### Custom Severity

Severity Matcher.

###### One of the arguments from this list "exact_match, regex_match" must be set

`exact_match` - (Optional) Equality match value for the label (`String`).

`regex_match` - (Optional) Regular expression match value for the label (`String`).

### Group By Custom

Specify set of labels for grouping the alerts.

`labels` - (Optional) Name of labels to group/aggregate the alerts (`String`).

### Group By Default

Group the alerts by severity, group name and alert name.

### Group By Individual

This option disables grouping of alerts.

### Group By Ves Io Group

Group the alerts by severity, group name and alert name.

### Matcher Any

Matches all alerts in the namespace.

### Matcher Custom

A set of custom equality/regex matchers an alert has to fulfill to match the route..

`alertlabel` - (Optional) AlertLabel to configure the alert policy rule. See [Custom Alertlabel ](#custom-alertlabel) below for details.

`alertname` - (Optional) Alertname Matcher. See [Custom Alertname ](#custom-alertname) below for details.

`group` - (Optional) Group Matcher. See [Custom Group ](#custom-group) below for details.

`severity` - (Optional) Severity Matcher. See [Custom Severity ](#custom-severity) below for details.

### Matcher Group

Matches the group name of the alert.

`groups` - (Optional) Name of groups to match the alert (`List of Strings`).

### Matcher Severity

Matches the severity level of the alert.

`severities` - (Optional) List of severity levels (`List of Strings`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Routes Notification Parameters

notification_config defined in the policy..

###### One of the arguments from this list "default, individual, ves_io_group, custom" must be set

`custom` - (Optional) Specify set of labels for grouping the alerts. See [Group By Custom ](#group-by-custom) below for details.

`default` - (Optional) Group the alerts by severity, group name and alert name (`Bool`).

`individual` - (Optional) This option disables grouping of alerts (`Bool`).

`ves_io_group` - (Optional) Group the alerts by severity, group name and alert name (`Bool`).

`group_interval` - (Optional) If not specified, group_interval defaults to "1m" (`String`).

`group_wait` - (Optional) If not specified, group_wait defaults to "30s" (`String`).

`repeat_interval` - (Optional) If not specified, group_interval defaults to "4h" (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured alert_policy.
