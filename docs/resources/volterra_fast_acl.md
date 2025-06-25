---

page_title: "Volterra: fast_acl"

description: "The fast_acl allows CRUD of Fast Acl resource on Volterra SaaS"
-----------------------------------------------------------------------------

Resource volterra_fast_acl
==========================

The Fast Acl allows CRUD of Fast Acl resource on Volterra SaaS

~> **Note:** Please refer to [Fast Acl API docs](https://docs.cloud.f5.com/docs-v2/api/fast-acl) to learn more

Example Usage
-------------

```hcl
resource "volterra_fast_acl" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "re_acl site_acl" must be set

  site_acl {
    fast_acl_rules {
      action {
        // One of the arguments from this list "policer_action protocol_policer_action simple_action" can be set

        policer_action {
          ref {
            name      = "test1"
            namespace = "staging"
            tenant    = "acmecorp"
          }
        }
      }

      metadata {
        description = "Virtual Host for acmecorp website"

        name = "acmecorp-web"
      }

      port {
        // One of the arguments from this list "all dns user_defined" can be set

        user_defined = "user_defined"
      }

      // One of the arguments from this list "ip_prefix_set prefix" must be set

      prefix {
        ipv6_prefix = ["[2001:db8::1::/112, 2001::db8::2::/112]"]

        prefix = ["[192.168.1.0/24, 192.168.2.0/24]\""]
      }
    }

    // One of the arguments from this list "inside_network outside_network" must be set

    outside_network = true

    // One of the arguments from this list "all_services interface_services vip_services" must be set

    interface_services = true
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

`protocol_policer` - (Optional) on that group, giving a semantics of hash limit for source IP. See [ref](#ref) below for details.

###### One of the arguments from this list "re_acl, site_acl" must be set

`re_acl` - (Optional) ACL will be applied at regional edge sites. See [Site Choice Re Acl ](#site-choice-re-acl) below for details.

`site_acl` - (Optional) ACL will be applied at customer edge sites. See [Site Choice Site Acl ](#site-choice-site-acl) below for details.

### Action Policer Action

Reference to policer object to which traffic would be subjected.

`ref` - (Optional) A policer direct reference. See [ref](#ref) below for details.

### Action Protocol Policer Action

Reference to protocol based policer object.

`ref` - (Optional) Reference to protocol policer object. See [ref](#ref) below for details.

### Fast Acl Rules Action

Action to be applied if traffic matched rule (pass, deny or rate limit).

###### One of the arguments from this list "policer_action, protocol_policer_action, simple_action" can be set

`policer_action` - (Optional) Reference to policer object to which traffic would be subjected. See [Action Policer Action ](#action-policer-action) below for details.

`protocol_policer_action` - (Optional) Reference to protocol based policer object. See [Action Protocol Policer Action ](#action-protocol-policer-action) below for details.

`simple_action` - (Optional) Simple action like dropping or forwarding the traffic (`String`).

### Fast Acl Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Fast Acl Rules Port

L4 port numbers to match.

###### One of the arguments from this list "all, dns, user_defined" can be set

`all` - (Optional) Matches all port (`Bool`).

`dns` - (Optional) Matches dns port 53 (`Bool`).

`user_defined` - (Optional) Matches the user defined port (`Int`).

### Network Choice Inside Network

Site Local Inside network.

### Network Choice Outside Network

Site Local Outside network.

### Port Value Type Choice All

Matches all port.

### Port Value Type Choice Dns

Matches dns port 53.

### Re Acl Fast Acl Rules

Fast ACL rules to match.

`action` - (Required) Action to be applied if traffic matched rule (pass, deny or rate limit). See [Fast Acl Rules Action ](#fast-acl-rules-action) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Fast Acl Rules Metadata ](#fast-acl-rules-metadata) below for details.

`port` - (Optional) L4 port numbers to match. See [Fast Acl Rules Port ](#fast-acl-rules-port) below for details.

###### One of the arguments from this list "ip_prefix_set, prefix" must be set

`ip_prefix_set` - (Optional) Reference to IP prefix set object. See [Source Ip Prefix Set ](#source-ip-prefix-set) below for details.

`prefix` - (Optional) List of IP prefixes. See [Source Prefix ](#source-prefix) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Site Acl Fast Acl Rules

Fast ACL rules to match.

`action` - (Required) Action to be applied if traffic matched rule (pass, deny or rate limit). See [Fast Acl Rules Action ](#fast-acl-rules-action) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Fast Acl Rules Metadata ](#fast-acl-rules-metadata) below for details.

`port` - (Optional) L4 port numbers to match. See [Fast Acl Rules Port ](#fast-acl-rules-port) below for details.

###### One of the arguments from this list "ip_prefix_set, prefix" must be set

`ip_prefix_set` - (Optional) Reference to IP prefix set object. See [Source Ip Prefix Set ](#source-ip-prefix-set) below for details.

`prefix` - (Optional) List of IP prefixes. See [Source Prefix ](#source-prefix) below for details.

### Site Choice Re Acl

ACL will be applied at regional edge sites.

`fast_acl_rules` - (Optional) Fast ACL rules to match. See [Re Acl Fast Acl Rules ](#re-acl-fast-acl-rules) below for details.

###### One of the arguments from this list "all_public_vips, default_tenant_vip, selected_tenant_vip" must be set

`all_public_vips` - (Optional) Apply this Fast ACL to all public vips (`Bool`).

`default_tenant_vip` - (Optional) Apply this Fast ACL to Default(dedicated) Tenant VIP (`Bool`).

`selected_tenant_vip` - (Optional) Apply this Fast ACL to List of some selected public VIP(s). See [Vip Choice Selected Tenant Vip ](#vip-choice-selected-tenant-vip) below for details.

### Site Choice Site Acl

ACL will be applied at customer edge sites.

`fast_acl_rules` - (Optional) Fast ACL rules to match. See [Site Acl Fast Acl Rules ](#site-acl-fast-acl-rules) below for details.

###### One of the arguments from this list "inside_network, outside_network" must be set

`inside_network` - (Optional) Site Local Inside network (`Bool`).

`outside_network` - (Optional) Site Local Outside network (`Bool`).

###### One of the arguments from this list "all_services, interface_services, vip_services" must be set

`all_services` - (Optional) Port and protocol is picked up from advertise policies (`Bool`).

`interface_services` - (Optional) Port and protocol is picked up from advertise policies (`Bool`).

`vip_services` - (Optional) Port and protocol is picked up from advertise policies (`Bool`).

### Source Ip Prefix Set

Reference to IP prefix set object.

`ref` - (Optional) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Source Prefix

List of IP prefixes.

`ipv6_prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

### Vip Choice All Public Vips

Apply this Fast ACL to all public vips.

### Vip Choice All Services

Port and protocol is picked up from advertise policies.

### Vip Choice Default Tenant Vip

Apply this Fast ACL to Default(dedicated) Tenant VIP.

### Vip Choice Interface Services

Port and protocol is picked up from advertise policies.

### Vip Choice Selected Tenant Vip

Apply this Fast ACL to List of some selected public VIP(s).

`default_tenant_vip` - (Optional) Include tenant vip in list of specific VIP(s) (`Bool`).

`public_ip_refs` - (Required) Select additional public VIP(s). See [ref](#ref) below for details.

### Vip Choice Vip Services

Port and protocol is picked up from advertise policies.

Attribute Reference
-------------------

-	`id` - This is the id of the configured fast_acl.
