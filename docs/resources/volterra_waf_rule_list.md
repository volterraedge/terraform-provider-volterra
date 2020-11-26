---

page_title: "Volterra: waf_rule_list"

description: "The waf_rule_list allows CRUD of Waf Rule List resource on Volterra SaaS"
---------------------------------------------------------------------------------------

Resource volterra_waf_rule_list
===============================

The Waf Rule List allows CRUD of Waf Rule List resource on Volterra SaaS

~> **Note:** Please refer to [Waf Rule List API docs](https://volterra.io/docs/api/waf-rule-list) to learn more

Example Usage
-------------

```hcl
resource "volterra_waf_rule_list" "example" {
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

`rule_ids` - (Optional) List of rule ids (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured waf_rule_list.
