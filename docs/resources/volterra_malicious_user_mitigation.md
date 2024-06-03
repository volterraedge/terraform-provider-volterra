---

page_title: "Volterra: malicious_user_mitigation"

description: "The malicious_user_mitigation allows CRUD of Malicious User Mitigation resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------------------

Resource volterra_malicious_user_mitigation
===========================================

The Malicious User Mitigation allows CRUD of Malicious User Mitigation resource on Volterra SaaS

~> **Note:** Please refer to [Malicious User Mitigation API docs](https://docs.cloud.f5.com/docs/api/malicious-user-mitigation) to learn more

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

### Mitigation Type

Malicious user mitigation rules specifies the actions to be taken for users to different threat levels.

`rules` - (Required) Define the threat levels and the corresponding mitigation actions to be taken. See [Mitigation Type Rules ](#mitigation-type-rules) below for details.

### Mitigation Action Alert Only

Generate alert while not taking any invasive actions.

### Mitigation Action Block Temporarily

assigned to this mitigation action.

### Mitigation Action Captcha Challenge

configured on the corresponding http load balancer.

### Mitigation Action Javascript Challenge

configured on the corresponding http load balancer.

### Mitigation Action None

No mitigation actions.

### Mitigation Type Rules

Define the threat levels and the corresponding mitigation actions to be taken.

`mitigation_action` - (Required) The action to be taken at the specified threat level. See [Rules Mitigation Action ](#rules-mitigation-action) below for details.

`threat_level` - (Required) The threat level at which mitigation actions will be taken. See [Rules Threat Level ](#rules-threat-level) below for details.

### Rules Mitigation Action

The action to be taken at the specified threat level.

###### One of the arguments from this list "block_temporarily, none, alert_only, javascript_challenge, captcha_challenge" must be set

`alert_only` - (Optional) Generate alert while not taking any invasive actions (`Bool`).(Deprecated)

`block_temporarily` - (Optional) assigned to this mitigation action (`Bool`).

`captcha_challenge` - (Optional) configured on the corresponding http load balancer (`Bool`).

`javascript_challenge` - (Optional) configured on the corresponding http load balancer (`Bool`).

`none` - (Optional) No mitigation actions (`Bool`).(Deprecated)

### Rules Threat Level

The threat level at which mitigation actions will be taken.

###### One of the arguments from this list "low, medium, high" must be set

`high` - (Optional) (`Bool`).

`low` - (Optional) (`Bool`).

`medium` - (Optional) (`Bool`).

### Threat Level High

.

### Threat Level Low

.

### Threat Level Medium

.

Attribute Reference
-------------------

-	`id` - This is the id of the configured malicious_user_mitigation.
