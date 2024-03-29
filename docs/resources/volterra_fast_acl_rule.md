---

page_title: "Volterra: fast_acl_rule"

description: "The fast_acl_rule allows CRUD of Fast Acl Rule resource on Volterra SaaS"
---------------------------------------------------------------------------------------

Resource volterra_fast_acl_rule
===============================

The Fast Acl Rule allows CRUD of Fast Acl Rule resource on Volterra SaaS

~> **Note:** Please refer to [Fast Acl Rule API docs](https://docs.cloud.f5.com/docs/api/fast-acl-rule) to learn more

Example Usage
-------------

```hcl
resource "volterra_fast_acl_rule" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  action {
    // One of the arguments from this list "simple_action policer_action protocol_policer_action" must be set
    simple_action = "simple_action"
  }

  port {
    // One of the arguments from this list "all user_defined dns" must be set
    user_defined = "user_defined"
  }

  // One of the arguments from this list "prefix ip_prefix_set" must be set

  prefix {
    ipv6_prefix = ["[2001:db8::1::/112, 2001::db8::2::/112]"]

    prefix = ["[192.168.1.0/24, 192.168.2.0/24]\""]
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

`action` - (Required) Action to be applied if traffic matched rule (allow, deny or police). See [Action ](#action) below for details.

`port` - (Required) L4 port numbers to match. See [Port ](#port) below for details.

`ip_prefix_set` - (Optional) Reference to IP prefix set object. See [Source Ip Prefix Set ](#source-ip-prefix-set) below for details.

`prefix` - (Optional) List of IP prefixes. See [Source Prefix ](#source-prefix) below for details.

### Action

Action to be applied if traffic matched rule (allow, deny or police).

###### One of the arguments from this list "simple_action, policer_action, protocol_policer_action" can be set

`policer_action` - (Optional) Reference to policer object to which traffic would be subjected. See [Action Policer Action ](#action-policer-action) below for details.

`protocol_policer_action` - (Optional) Reference to protocol based policer object. See [Action Protocol Policer Action ](#action-protocol-policer-action) below for details.

`simple_action` - (Optional) Simple action like dropping or forwarding the traffic (`String`).

### Port

L4 port numbers to match.

###### One of the arguments from this list "all, user_defined, dns" can be set

`all` - (Optional) Matches all port (`Bool`).

`dns` - (Optional) Matches dns port 53 (`Bool`).

`user_defined` - (Optional) Matches the user defined port (`Int`).

### Action Policer Action

Reference to policer object to which traffic would be subjected.

`ref` - (Optional) A policer direct reference. See [ref](#ref) below for details.

### Action Protocol Policer Action

Reference to protocol based policer object.

`ref` - (Optional) Reference to protocol policer object. See [ref](#ref) below for details.

### Port Value Type Choice All

Matches all port.

### Port Value Type Choice Dns

Matches dns port 53.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Source Ip Prefix Set

Reference to IP prefix set object.

`ref` - (Optional) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Source Prefix

List of IP prefixes.

`ipv6_prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured fast_acl_rule.
