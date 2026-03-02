---

page_title: "Volterra: v1_dns_monitor"

description: "The v1_dns_monitor allows CRUD of V1 Dns Monitor resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_v1_dns_monitor
================================

The V1 Dns Monitor allows CRUD of V1 Dns Monitor resource on Volterra SaaS

~> **Note:** Please refer to [V1 Dns Monitor API docs](https://docs.cloud.f5.com/docs-v2/api/v1-dns-monitor) to learn more

Example Usage
-------------

```hcl
resource "volterra_v1_dns_monitor" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
  domain    = "www.example.com"

  external_sources {
    // One of the arguments from this list "aws f5xc" must be set

    aws {
      regions = ["us-east-1"]
    }
  }

  // One of the arguments from this list "interval_12_hours interval_15_mins interval_1_day interval_1_hour interval_1_min interval_30_mins interval_30_secs interval_5_mins interval_6_hours" must be set

  interval_1_hour  = true
  lookup_timeout   = "5000"
  on_failure_count = "2"

  // One of the arguments from this list "on_failure_to_all on_failure_to_any" must be set

  on_failure_to_any         = true
  protocol                  = "UDP"
  record_type               = "A"
  source_critical_threshold = "1"
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

`domain` - (Required) x-example: "www.example.com" (`String`).

`external_sources` - (Required) Internet facing monitor sources. See [External Sources ](#external-sources) below for details.

`health_policy` - (Optional) x-displayName: "Health Policy". See [Health Policy ](#health-policy) below for details.

###### One of the arguments from this list "interval_12_hours, interval_15_mins, interval_1_day, interval_1_hour, interval_1_min, interval_30_mins, interval_30_secs, interval_5_mins, interval_6_hours" must be set

`interval_12_hours` - (Optional) x-displayName: "12 hours" (`Bool`).

`interval_15_mins` - (Optional) x-displayName: "15 minutes" (`Bool`).

`interval_1_day` - (Optional) x-displayName: "1 day" (`Bool`).

`interval_1_hour` - (Optional) x-displayName: "1 hour" (`Bool`).

`interval_1_min` - (Optional) x-displayName: "1 minute" (`Bool`).

`interval_30_mins` - (Optional) x-displayName: "30 minutes" (`Bool`).

`interval_30_secs` - (Optional) x-displayName: "30 seconds" (`Bool`).

`interval_5_mins` - (Optional) x-displayName: "5 minutes" (`Bool`).

`interval_6_hours` - (Optional) x-displayName: "6 hours" (`Bool`).

`lookup_timeout` - (Required) The amount of time in milliseconds before the monitor considers a pending request to be a failure (`Int`).

`name_servers` - (Optional) x-example: "". See [Name Servers ](#name-servers) below for details.

`on_failure_count` - (Required) The number of times a monitor must fail before the global monitor changes health (`Int`).

###### One of the arguments from this list "on_failure_to_all, on_failure_to_any" must be set

`on_failure_to_all` - (Optional) x-displayName: "ALL" (`Bool`).

`on_failure_to_any` - (Optional) x-displayName: "ANY" (`Bool`).

`protocol` - (Required) x-example: "UDP" (`String`).

`receive` - (Optional) The regex pattern that the monitor looks for in the returned resource. If the text string is returned, the monitor is healthy. (`String`).

`record_type` - (Required) x-example: "A" (`String`).

`source_critical_threshold` - (Required) The number of provider-regions a monitor must fail from before the global monitor changes health (`Int`).

### External Sources

Internet facing monitor sources.

###### One of the arguments from this list "aws, f5xc" must be set

`aws` - (Optional) Amazon Web Services. See [Source Choice Aws ](#source-choice-aws) below for details.

`f5xc` - (Optional) F5 Distributed Cloud. See [Source Choice F5xc ](#source-choice-f5xc) below for details.

### Health Policy

x-displayName: "Health Policy".

###### One of the arguments from this list "dynamic_threshold, dynamic_threshold_disabled" must be set

`dynamic_threshold` - (Optional) x-displayName: "Enable". See [Dynamic Threshold Policy Dynamic Threshold ](#dynamic-threshold-policy-dynamic-threshold) below for details.

`dynamic_threshold_disabled` - (Optional) x-displayName: "Disable" (`Bool`).

###### One of the arguments from this list "static_max_threshold, static_max_threshold_disabled" must be set

`static_max_threshold` - (Optional) x-displayName: "Enable". See [Static Max Threshold Policy Static Max Threshold ](#static-max-threshold-policy-static-max-threshold) below for details.

`static_max_threshold_disabled` - (Optional) x-displayName: "Disable" (`Bool`).

###### One of the arguments from this list "static_min_threshold, static_min_threshold_disabled" must be set

`static_min_threshold` - (Optional) x-displayName: "Enable". See [Static Min Threshold Policy Static Min Threshold ](#static-min-threshold-policy-static-min-threshold) below for details.

`static_min_threshold_disabled` - (Optional) x-displayName: "Disable" (`Bool`).

### Name Servers

x-example: "".

`name_server` - (Required) IP address of the nameserver to execute the monitor against (`String`).

`port` - (Optional) Port of the nameserver to execute the monitor against (`Int`).

### Dynamic Threshold Policy Dynamic Threshold

x-displayName: "Enable".

`eval_period` - (Required) The duration that the response time value must exceed the set threshold before the monitor turns critical. (`String`).

`std_dev_val` - (Required) The number of standard deviations response time value must exceed before the monitor turns critical. (`String`).

### Dynamic Threshold Policy Dynamic Threshold Disabled

x-displayName: "Disable".

### Source Choice Aws

Amazon Web Services.

`regions` - (Required) A specific source location within AWS (`String`).

### Source Choice F5xc

F5 Distributed Cloud.

`regions` - (Required) A specific source location within F5 Distributed Cloud (`String`).

### Static Max Threshold Policy Static Max Threshold

x-displayName: "Enable".

`eval_period` - (Required) The duration that the response time value must exceed the set threshold before the monitor turns critical. (`String`).

`max_response_time` - (Required) The maximum response time value that must be exceeded before the monitor turns critical. (`Int`).

### Static Max Threshold Policy Static Max Threshold Disabled

x-displayName: "Disable".

### Static Min Threshold Policy Static Min Threshold

x-displayName: "Enable".

`eval_period` - (Required) The duration that the response time value must fall below the set threshold before the monitor turns critical. (`String`).

`min_response_time` - (Required) The minimum value that response time must fall below before the monitor turns critical. (`Int`).

### Static Min Threshold Policy Static Min Threshold Disabled

x-displayName: "Disable".

Attribute Reference
-------------------

*   `id` - This is the id of the configured v1_dns_monitor.
