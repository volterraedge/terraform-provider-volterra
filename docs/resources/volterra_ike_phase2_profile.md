---

page_title: "Volterra: ike_phase2_profile"

description: "The ike_phase2_profile allows CRUD of Ike Phase2 Profile resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------

Resource volterra_ike_phase2_profile
====================================

The Ike Phase2 Profile allows CRUD of Ike Phase2 Profile resource on Volterra SaaS

~> **Note:** Please refer to [Ike Phase2 Profile API docs](https://docs.cloud.f5.com/docs-v2/api/views-ike-phase2-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_ike_phase2_profile" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  authentication_algos = ["authentication_algos"]

  encryption_algos = ["encryption_algos"]

  // One of the arguments from this list "ike_keylifetime_hours ike_keylifetime_minutes use_default_keylifetime" must be set

  ike_keylifetime_hours {
    duration = "duration"
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

`authentication_algos` - (Required) Choose one or more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm encryption algorithms. (`List of Strings`).

`encryption_algos` - (Required) Choose one or more encryption algorithms. (`List of Strings`).

###### One of the arguments from this list "ike_keylifetime_hours, ike_keylifetime_minutes, use_default_keylifetime" must be set

`ike_keylifetime_hours` - (Optional) Set IKE Key Lifetime in Hours. See [Ike Key Lifetime Ike Keylifetime Hours ](#ike-key-lifetime-ike-keylifetime-hours) below for details.

`ike_keylifetime_minutes` - (Optional) Set IKE Key Lifetime in minutes. See [Ike Key Lifetime Ike Keylifetime Minutes ](#ike-key-lifetime-ike-keylifetime-minutes) below for details.

`use_default_keylifetime` - (Optional) Use default key lifetime (4Hrs) (`Bool`).

###### One of the arguments from this list "dh_group_set, disable_pfs" can be set

`dh_group_set` - (Optional) Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile.. See [Pfs Mode Dh Group Set ](#pfs-mode-dh-group-set) below for details.

`disable_pfs` - (Optional) Disable PFS for this tunnel (`Bool`).

### Ike Key Lifetime Ike Keylifetime Hours

Set IKE Key Lifetime in Hours.

`duration` - (Optional) x-displayName: "Duration" (`Int`).

### Ike Key Lifetime Ike Keylifetime Minutes

Set IKE Key Lifetime in minutes.

`duration` - (Optional) x-displayName: "Duration" (`Int`).

### Pfs Mode Dh Group Set

Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile..

`dh_groups` - (Required) Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile. (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured ike_phase2_profile.
