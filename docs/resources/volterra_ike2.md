---

page_title: "Volterra: ike2"

description: "The ike2 allows CRUD of Ike2 resource on Volterra SaaS"
---------------------------------------------------------------------

Resource volterra_ike2
======================

The Ike2 allows CRUD of Ike2 resource on Volterra SaaS

~> **Note:** Please refer to [Ike2 API docs](https://docs.cloud.f5.com/docs-v2/api/ike2) to learn more

Example Usage
-------------

```hcl
resource "volterra_ike2" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "ike_keylifetime_hours ike_keylifetime_minutes use_default_keylifetime" must be set

  use_default_keylifetime = true
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

`authentication_algos` - (Optional) Choose one or more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm encryption algorithms. (`List of Strings`).(Deprecated)

`encryption_algos` - (Optional) Choose one or more encryption algorithms. (`List of Strings`).(Deprecated)

###### One of the arguments from this list "ike_keylifetime_hours, ike_keylifetime_minutes, use_default_keylifetime" must be set

`ike_keylifetime_hours` - (Optional) Set IKE Key Lifetime in Hours. See [Ike Key Lifetime Ike Keylifetime Hours ](#ike-key-lifetime-ike-keylifetime-hours) below for details.(Deprecated)

`ike_keylifetime_minutes` - (Optional) Set IKE Key Lifetime in minutes. See [Ike Key Lifetime Ike Keylifetime Minutes ](#ike-key-lifetime-ike-keylifetime-minutes) below for details.(Deprecated)

`use_default_keylifetime` - (Optional) Use default key lifetime (4Hrs) (`Bool`).(Deprecated)

###### One of the arguments from this list "dh_group_set, disable_pfs" can be set

`dh_group_set` - (Optional) Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile.. See [Pfs Mode Dh Group Set ](#pfs-mode-dh-group-set) below for details.(Deprecated)

`disable_pfs` - (Optional) Disable PFS for this tunnel (`Bool`).(Deprecated)

### Ike Key Lifetime Ike Keylifetime Hours

Set IKE Key Lifetime in Hours.

`duration` - (Optional) x-displayName: "Duration" (`Int`).

### Ike Key Lifetime Ike Keylifetime Minutes

Set IKE Key Lifetime in minutes.

`duration` - (Optional) x-displayName: "Duration" (`Int`).

### Pfs Mode Dh Group Set

Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile..

`dh_groups` - (Optional) x-displayName: "Diffie Hellman Groups" (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured ike2.
