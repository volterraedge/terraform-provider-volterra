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

`mitigation_type` - (Optional) Malicious user mitigation rules specifies the actions to be taken for users to different threat levels. See [Mitigation Type ](#mitigation-type) below for details.

### Alert Only

Generate alert while not taking any invasive actions.

### Block Temporarily

assigned to this mitigation action.

### Captcha Challenge

configured on the corresponding http load balancer.

### High

.

### Javascript Challenge

configured on the corresponding http load balancer.

### Low

.

### Medium

.

### Mitigation Action

The action to be taken at the specified threat level.

`alert_only` - (Optional) Generate alert while not taking any invasive actions (bool).

`block_temporarily` - (Optional) assigned to this mitigation action (bool).

`captcha_challenge` - (Optional) configured on the corresponding http load balancer (bool).

`javascript_challenge` - (Optional) configured on the corresponding http load balancer (bool).

`none` - (Optional) No mitigation actions (bool).

### Mitigation Type

Malicious user mitigation rules specifies the actions to be taken for users to different threat levels.

`rules` - (Required) Define the threat levels and the corresponding mitigation actions to be taken. See [Rules ](#rules) below for details.

### None

No mitigation actions.

### Rules

Define the threat levels and the corresponding mitigation actions to be taken.

`mitigation_action` - (Required) The action to be taken at the specified threat level. See [Mitigation Action ](#mitigation-action) below for details.

`threat_level` - (Required) The threat level at which mitigation actions will be taken. See [Threat Level ](#threat-level) below for details.

### Threat Level

The threat level at which mitigation actions will be taken.

`high` - (Optional) (bool).

`low` - (Optional) (bool).

`medium` - (Optional) (bool).

Attribute Reference
-------------------

-	`id` - This is the id of the configured malicious_user_mitigation.
