---

page_title: "Volterra: malicious_user_mitigation"

description: "The malicious_user_mitigation allows CRUD of Malicious User Mitigation resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------------------

Resource volterra_malicious_user_mitigation
===========================================

The Malicious User Mitigation allows CRUD of Malicious User Mitigation resource on Volterra SaaS

~> **Note:** Please refer to [Malicious User Mitigation API docs](https://volterra.io/docs/api/malicious-user-mitigation) to learn more

Example Usage
-------------

```hcl
resource "volterra_malicious_user_mitigation" "example" {
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

`mitigation_type` - (Optional) A threat level is calculated for every user identified using config specified in user_identification by analyzing their activity and reputation.. See [Mitigation Type ](#mitigation-type) below for details.

### Alert Only

Generate alert while not taking any invasive actions.

### Block Temporarily

If temporary blocking is not configured for the virtual host, a software default configuration is used.

### Captcha Challenge

If Captcha Challenge is not configured for the virtual host, a software default configuration is used.

### High

User estimated to be high threat.

### Javascript Challenge

If Javascript Challenge is not configured for the virtual host, a software default configuration is used.

### Low

User estimated to be low threat.

### Medium

User estimated to be medium threat.

### Mitigation Action

The action to be taken at the specified threat level.

`alert_only` - (Optional) Generate alert while not taking any invasive actions (bool).

`block_temporarily` - (Optional) If temporary blocking is not configured for the virtual host, a software default configuration is used (bool).

`captcha_challenge` - (Optional) If Captcha Challenge is not configured for the virtual host, a software default configuration is used (bool).

`javascript_challenge` - (Optional) If Javascript Challenge is not configured for the virtual host, a software default configuration is used (bool).

`none` - (Optional) No mitigation actions (bool).

### Mitigation Type

A threat level is calculated for every user identified using config specified in user_identification by analyzing their activity and reputation..

`rules` - (Required) A threat level is calculated for every user identified using config specified in user_identification by analyzing their activity and reputation.. See [Rules ](#rules) below for details.

### None

No mitigation actions.

### Rules

A threat level is calculated for every user identified using config specified in user_identification by analyzing their activity and reputation..

`mitigation_action` - (Required) The action to be taken at the specified threat level. See [Mitigation Action ](#mitigation-action) below for details.

`threat_level` - (Required) The threat level at which mitigation actions will be taken. See [Threat Level ](#threat-level) below for details.

### Threat Level

The threat level at which mitigation actions will be taken.

`high` - (Optional) User estimated to be high threat (bool).

`low` - (Optional) User estimated to be low threat (bool).

`medium` - (Optional) User estimated to be medium threat (bool).

Attribute Reference
-------------------

-	`id` - This is the id of the configured malicious_user_mitigation.
