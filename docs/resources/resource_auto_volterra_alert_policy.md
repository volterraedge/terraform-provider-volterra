---

page_title: "Volterra: alert_policy"

description: "The alert_policy allows CRUD of Alert Policy resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_alert_policy
==============================

The Alert Policy allows CRUD of Alert Policy resource on Volterra SaaS

~> **Note:** Please refer to [Alert Policy API docs](https://volterra.io/docs/api/alert-policy) to learn more

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

    // One of the arguments from this list "alertname_regex custom any severity group alertname" must be set
    any = true

    notification_parameters {
      // One of the arguments from this list "default individual ves_io_group custom" must be set
      individual      = true
      group_interval  = "1m"
      group_wait      = "30s"
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

`custom` - (Optional) Specify set of labels for grouping the alerts. See [Custom ](#custom) below for details.

`default` - (Optional) Group the alerts by severity, group name and alert name (bool).

`individual` - (Optional) This option disables grouping of alerts (bool).

`ves_io_group` - (Optional) Group the alerts by severity, group name and alert name (bool).

`group_interval` - (Optional) If not specified, group_interval defaults to "1m" (`String`).

`group_wait` - (Optional) If not specified, group_wait defaults to "30s" (`String`).

`repeat_interval` - (Optional) If not specified, group_interval defaults to "4h" (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Routes

The routes are evaluated in the specified order and terminates on the first match..

`dont_send` - (Optional) Do not send the alert (bool).

`send` - (Optional) Send the alert (bool).

`alertname` - (Optional) Matches the alertname of the alert (`String`).

`alertname_regex` - (Optional) Regular Expression match for the alertname (`String`).

`any` - (Optional) Matches all alerts in the namespace (bool).

`custom` - (Optional) A set of custom equality/regex matchers an alert has to fulfill to match the route.. See [Custom ](#custom) below for details.

`group` - (Optional) Matches the group name of the alert. See [Group ](#group) below for details.

`severity` - (Optional) Matches the severity level of the alert. See [Severity ](#severity) below for details.

`notification_parameters` - (Optional) notification_config defined in the policy.. See [Notification Parameters ](#notification-parameters) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured alert_policy.
