---

page_title: "Volterra: ike_phase1_profile"

description: "The ike_phase1_profile allows CRUD of Ike Phase1 Profile resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------

Resource volterra_ike_phase1_profile
====================================

The Ike Phase1 Profile allows CRUD of Ike Phase1 Profile resource on Volterra SaaS

~> **Note:** Please refer to [Ike Phase1 Profile API docs](https://docs.cloud.f5.com/docs-v2/api/views-ike-phase1-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_ike_phase1_profile" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  authentication_algos = ["authentication_algos"]

  dh_group = ["dh_group"]

  encryption_algos = ["encryption_algos"]

  // One of the arguments from this list "ike_keylifetime_hours ike_keylifetime_minutes use_default_keylifetime" must be set

  use_default_keylifetime = true

  // One of the arguments from this list "reauth_disabled reauth_timeout_days reauth_timeout_hours" must be set

  reauth_timeout_hours {
    duration = "duration"
  }
  prf = ["prf"]
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

`authentication_algos` - (Required) Choose one or more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm encryption algorithms. (`List of Strings`).

`dh_group` - (Required) Choose the acceptable Diffie Hellman (DH) Group or Groups that you are willing to accept as part of this profile. (`List of Strings`).

`encryption_algos` - (Required) Choose one or more encryption algorithms. (`List of Strings`).

###### One of the arguments from this list "ike_keylifetime_hours, ike_keylifetime_minutes, use_default_keylifetime" must be set

`ike_keylifetime_hours` - (Optional) Set IKE Key Lifetime in Hours. See [Ike Key Lifetime Ike Keylifetime Hours ](#ike-key-lifetime-ike-keylifetime-hours) below for details.

`ike_keylifetime_minutes` - (Optional) Set IKE Key Lifetime in minutes. See [Ike Key Lifetime Ike Keylifetime Minutes ](#ike-key-lifetime-ike-keylifetime-minutes) below for details.

`use_default_keylifetime` - (Optional) Use default key lifetime (4Hrs) (`Bool`).

###### One of the arguments from this list "reauth_disabled, reauth_timeout_days, reauth_timeout_hours" must be set

`reauth_disabled` - (Optional) The default is not use IKE reauthentication. (`Bool`).

`reauth_timeout_days` - (Optional) Set Reauthentication timeout in Days. See [Ike Reauth Timeout Reauth Timeout Days ](#ike-reauth-timeout-reauth-timeout-days) below for details.

`reauth_timeout_hours` - (Optional) Set Reauthentication timeout in Hours. See [Ike Reauth Timeout Reauth Timeout Hours ](#ike-reauth-timeout-reauth-timeout-hours) below for details.

`prf` - (Required) Select PseudoRandomFunction for IKE SA. (`List of Strings`).

### Ike Key Lifetime Ike Keylifetime Hours

Set IKE Key Lifetime in Hours.

`duration` - (Optional) x-displayName: "Duration" (`Int`).

### Ike Key Lifetime Ike Keylifetime Minutes

Set IKE Key Lifetime in minutes.

`duration` - (Optional) x-displayName: "Duration" (`Int`).

### Ike Reauth Timeout Reauth Timeout Days

Set Reauthentication timeout in Days.

`duration` - (Optional) x-displayName: "Duration" (`Int`).

### Ike Reauth Timeout Reauth Timeout Hours

Set Reauthentication timeout in Hours.

`duration` - (Optional) x-displayName: "Duration" (`Int`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured ike_phase1_profile.
