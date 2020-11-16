---

page_title: "Volterra: fast_acl"

description: "The fast_acl allows CRUD of Fast Acl resource on Volterra SaaS"
-----------------------------------------------------------------------------

Resource volterra_fast_acl
==========================

The Fast Acl allows CRUD of Fast Acl resource on Volterra SaaS

~> **Note:** Please refer to [Fast Acl API docs](https://volterra.io/docs/api/fast-acl) to learn more

Example Usage
-------------

```hcl
resource "volterra_fast_acl" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "site_acl re_acl legacy_acl" must be set

  site_acl {
    fast_acl_rules {
      action {
        // One of the arguments from this list "protocol_policer_action simple_action policer_action" must be set

        protocol_policer_action {
          ref {
            name      = "test1"
            namespace = "staging"
            tenant    = "acmecorp"
          }
        }
      }

      name = "name"

      port {
        // One of the arguments from this list "all user_defined dns" must be set
        all = true
      }

      // One of the arguments from this list "prefix ip_prefix_set" must be set

      prefix {
        prefix = ["[192.168.1.0/24, 192.168.2.0/24]\" or \"[2001:db8::1::/112, 2001::db8::2::/112]"]
      }
    }

    // One of the arguments from this list "inside_network outside_network" must be set
    outside_network = true

    // One of the arguments from this list "interface_services vip_services all_services destination_ip_address shared_vip_services" must be set

    destination_ip_address {
      address {
        // One of the arguments from this list "ipv4 ipv6" must be set

        ipv4 {
          addr = "192.168.1.1"
        }
      }

      ports {
        // One of the arguments from this list "dns all user_defined" must be set
        all = true
      }

      protocol = "TCP"
    }
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

`legacy_acl` - (Optional) ACL may be applied at regional edge sites or customer edge sites. Not recommended. See [Legacy Acl ](#legacy-acl) below for details.

`re_acl` - (Optional) ACL will be applied at regional edge sites. See [Re Acl ](#re-acl) below for details.

`site_acl` - (Optional) ACL will be applied at customer edge sites. See [Site Acl ](#site-acl) below for details.

### Action

Action to be applied if traffic matched rule (pass, deny or rate limit).

`policer_action` - (Optional) Reference to policer object to which traffic would be subjected. See [Policer Action ](#policer-action) below for details.

`protocol_policer_action` - (Optional) Reference to protocol based policer object. See [Protocol Policer Action ](#protocol-policer-action) below for details.

`simple_action` - (Optional) Simple action like dropping or forwarding the traffic (`String`).

### Destination Type

1.	Explicit IP and port.

`all_services` - (Optional) Regional Edge: Applies the configuration to all the VIPs assigned to tenant (bool).

`destination_ip_address` - (Optional) Regional Edge : Tenant can not configure it. See [Destination Ip Address ](#destination-ip-address) below for details.

`interface_services` - (Optional) Regional Edge : tenants CANNOT use this option on RE (bool).

`selected_vip_address` - (Optional) Valid only for RE.. See [Selected Vip Address ](#selected-vip-address) below for details.

`shared_vip_services` - (Optional) Regional Edge: Applies configuration on all shared VIPs used for services (bool).

`vip_services` - (Optional) Regional Edge: Applies the configuration to VIP which has been assigned by default to tenant. Not applicable for shared VIP(s) (bool).

### Fast Acl Rules

Fast ACL rules to match .

`action` - (Required) Action to be applied if traffic matched rule (pass, deny or rate limit). See [Action ](#action) below for details.

`name` - (Required) Name for this rule, will be used to generate metrics (`String`).

`port` - (Optional) L4 port numbers to match. See [Port ](#port) below for details.

`ip_prefix_set` - (Optional) Reference to IP prefix set object. See [Ip Prefix Set ](#ip-prefix-set) below for details.

`prefix` - (Optional) List of IP prefixes. See [Prefix ](#prefix) below for details.

### Legacy Acl

ACL may be applied at regional edge sites or customer edge sites. Not recommended.

`destination_type` - (Required) 4. Explicit IP and port. See [Destination Type ](#destination-type) below for details.

`network_type` - (Required) CE applies Fast ACLs with network type selector as "site_local" and "site_local_inside" only. See [Network Type ](#network-type) below for details.

`source_rules` - (Optional) List of fast ACL rules to be applied to received packets on this site. See [ref](#ref) below for details.

### Network Type

CE applies Fast ACLs with network type selector as "site_local" and "site_local_inside" only.

`public` - (Optional) Indicates use of public network (bool).

`site_local` - (Optional) Indicates use of site local network (bool).

`site_local_inside` - (Optional) Indicates use of site local inside network (bool).

### Port

L4 port numbers to match.

`all` - (Optional) Matches all port (bool).

`dns` - (Optional) Matches dns port 53 (bool).

`user_defined` - (Optional) Matches the user defined port (`Int`).

### Re Acl

ACL will be applied at regional edge sites.

`fast_acl_rules` - (Optional) Fast ACL rules to match . See [Fast Acl Rules ](#fast-acl-rules) below for details.

`all_public_vips` - (Optional) Apply this fast ACL to all public vips (bool).

`default_tenant_vip` - (Optional) Apply this fast ACL to Default(dedicated) Tenant VIP (bool).

`selected_tenant_vip` - (Optional) Apply this fast ACL to List of some selected public VIP(s). See [Selected Tenant Vip ](#selected-tenant-vip) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Site Acl

ACL will be applied at customer edge sites.

`fast_acl_rules` - (Optional) Fast ACL rules to match . See [Fast Acl Rules ](#fast-acl-rules) below for details.

`inside_network` - (Optional) Site local Outside network (bool).

`outside_network` - (Optional) Site local Inside network (bool).

`all_services` - (Optional) Port and protocol is picked up from advertise policies (bool).

`destination_ip_address` - (Optional) Destination is configured (ip, protocol, port). See [Destination Ip Address ](#destination-ip-address) below for details.

`interface_services` - (Optional) Port and protocol is picked up from advertise policies (bool).

`shared_vip_services` - (Optional) Regional Edge: Applies configuration on all shared VIPs used for services (bool).

`vip_services` - (Optional) Port and protocol is picked up from advertise policies (bool).

Attribute Reference
-------------------

-	`id` - This is the id of the configured fast_acl.
