---

page_title: "Volterra: waf_rules"

description: "The waf_rules allows CRUD of Waf Rules resource on Volterra SaaS"
-------------------------------------------------------------------------------

Resource volterra_waf_rules
===========================

The Waf Rules allows CRUD of Waf Rules resource on Volterra SaaS

~> **Note:** Please refer to [Waf Rules API docs](https://volterra.io/docs/api/waf-rules) to learn more

Example Usage
-------------

```hcl
resource "volterra_waf_rules" "example" {
  name                    = "acmecorp-web"
  namespace               = "staging"
  anomaly_score_threshold = ["4"]
  mode                    = ["mode"]
  paranoia_level          = ["2"]
  rule_list_type          = ["rule_list_type"]
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

`anomaly_score_threshold` - (Required) on non_blocking_mode configuration) (`Int`).

`mode` - (Required) WAF Mode is blocking or Alert (`String`).

`paranoia_level` - (Required) Paranoia level (`Int`).

`rule_ids` - (Optional) rule IDs to be included or excluded in this WAF instance (`List of Strings`).

`rule_list_type` - (Required) Specify whether the defined rule_ids list is to be included or excluded from associated WAF instance (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured waf_rules.
