---

page_title: "Volterra: infraprotect_firewall_rule_group"

description: "The infraprotect_firewall_rule_group allows CRUD of Infraprotect Firewall Rule Group resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------------------------------

Resource volterra_infraprotect_firewall_rule_group
==================================================

The Infraprotect Firewall Rule Group allows CRUD of Infraprotect Firewall Rule Group resource on Volterra SaaS

~> **Note:** Please refer to [Infraprotect Firewall Rule Group API docs](https://docs.cloud.f5.com/docs-v2/api/infraprotect-firewall-rule-group) to learn more

Example Usage
-------------

```hcl
resource "volterra_infraprotect_firewall_rule_group" "example" {
  name                     = "acmecorp-web"
  namespace                = "staging"
  firewall_rule_group_name = ["Group Name"]
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

`firewall_rule_group_name` - (Required) Firewall Rule Group Name (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured infraprotect_firewall_rule_group.
