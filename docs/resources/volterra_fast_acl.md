---

page_title: "Volterra: fast_acl"

description: "The fast_acl allows CRUD of Fast Acl resource on Volterra SaaS"
-----------------------------------------------------------------------------

Resource volterra_fast_acl
==========================

The Fast Acl allows CRUD of Fast Acl resource on Volterra SaaS

~> **Note:** Please refer to [Fast Acl API docs](https://docs.cloud.f5.com/docs/api/fast-acl) to learn more

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
        // One of the arguments from this list "simple_action policer_action protocol_policer_action" must be set
        simple_action = "simple_action"
      }

      metadata {
        description = "Virtual Host for acmecorp website"

        disable = true

        name = "acmecorp-web"
      }

      name = "value"

      port {
        // One of the arguments from this list "all user_defined dns" must be set
        all = true
      }

      // One of the arguments from this list "prefix ip_prefix_set" must be set

      prefix {
        ipv6_prefix = ["[2001:db8::1::/112, 2001::db8::2::/112]"]

        prefix = ["[192.168.1.0/24, 192.168.2.0/24]\""]
      }
    }

    // One of the arguments from this list "outside_network inside_network" must be set
    inside_network = true

    // One of the arguments from this list "interface_services vip_services all_services" must be set
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

`legacy_acl` - (Optional) ACL may be applied at regional edge sites or customer edge sites. Not recommended. See [Site Choice Legacy Acl ](#site-choice-legacy-acl) below for details.(Deprecated)

`re_acl` - (Optional) ACL will be applied at regional edge sites. See [Site Choice Re Acl ](#site-choice-re-acl) below for details.

`site_acl` - (Optional) ACL will be applied at customer edge sites. See [Site Choice Site Acl ](#site-choice-site-acl) below for details.

### Action Policer Action

Reference to policer object to which traffic would be subjected.

`ref` - (Optional) A policer direct reference. See [ref](#ref) below for details.

### Action Protocol Policer Action

Reference to protocol based policer object.

`ref` - (Optional) Reference to protocol policer object. See [ref](#ref) below for details.

### Destination Ip Address Address

List of IP addresses to match with destination.

###### One of the arguments from this list "ipv6, ipv4" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Destination Ip Address Ports

Special value "0" means all valid ports on the VIPs.

###### One of the arguments from this list "all, user_defined, dns" can be set

`all` - (Optional) Matches all port (`Bool`).

`dns` - (Optional) Matches dns port 53 (`Bool`).

`user_defined` - (Optional) Matches the user defined port (`Int`).

### Destination Type Choice All Services

Regional Edge: Applies the configuration to all the VIPs assigned to tenant.

### Destination Type Choice Destination Ip Address

Regional Edge : Tenant can not configure it.

`address` - (Optional) List of IP addresses to match with destination. See [Destination Ip Address Address ](#destination-ip-address-address) below for details.

`ports` - (Required) Special value "0" means all valid ports on the VIPs. See [Destination Ip Address Ports ](#destination-ip-address-ports) below for details.

`protocol` - (Optional) Protocol to match in the traffic (`String`).

### Destination Type Choice Interface Services

Regional Edge : tenants CANNOT use this option on RE.

### Destination Type Choice Selected Vip Address

Valid only for RE..

`address` - (Optional) List of IP addresses to match with destination. See [Selected Vip Address Address ](#selected-vip-address-address) below for details.

### Destination Type Choice Shared Vip Services

Regional Edge: Applies configuration on all shared VIPs used for services.

### Destination Type Choice Vhost

Regional Edge: Allowed.

### Destination Type Choice Vip Services

Regional Edge: Applies the configuration to VIP which has been assigned by default to tenant. Not applicable for shared VIP(s).

### Fast Acl Rules Action

Action to be applied if traffic matched rule (pass, deny or rate limit).

###### One of the arguments from this list "simple_action, policer_action, protocol_policer_action" can be set

`policer_action` - (Optional) Reference to policer object to which traffic would be subjected. See [Action Policer Action ](#action-policer-action) below for details.

`protocol_policer_action` - (Optional) Reference to protocol based policer object. See [Action Protocol Policer Action ](#action-protocol-policer-action) below for details.

`simple_action` - (Optional) Simple action like dropping or forwarding the traffic (`String`).

### Fast Acl Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Fast Acl Rules Port

L4 port numbers to match.

###### One of the arguments from this list "all, user_defined, dns" can be set

`all` - (Optional) Matches all port (`Bool`).

`dns` - (Optional) Matches dns port 53 (`Bool`).

`user_defined` - (Optional) Matches the user defined port (`Int`).

### Legacy Acl Destination Type

1.	Explicit IP and port.

###### One of the arguments from this list "destination_ip_address, selected_vip_address, shared_vip_services, vhost, interface_services, vip_services, all_services" can be set

`all_services` - (Optional) Regional Edge: Applies the configuration to all the VIPs assigned to tenant (`Bool`).

`destination_ip_address` - (Optional) Regional Edge : Tenant can not configure it. See [Destination Type Choice Destination Ip Address ](#destination-type-choice-destination-ip-address) below for details.

`interface_services` - (Optional) Regional Edge : tenants CANNOT use this option on RE (`Bool`).

`selected_vip_address` - (Optional) Valid only for RE.. See [Destination Type Choice Selected Vip Address ](#destination-type-choice-selected-vip-address) below for details.

`shared_vip_services` - (Optional) Regional Edge: Applies configuration on all shared VIPs used for services (`Bool`).

`vhost` - (Optional) Regional Edge: Allowed (`Bool`).(Deprecated)

`vip_services` - (Optional) Regional Edge: Applies the configuration to VIP which has been assigned by default to tenant. Not applicable for shared VIP(s) (`Bool`).

### Legacy Acl Network Type

CE applies Fast ACLs with network type selector as "site_local" and "site_local_inside" only.

###### One of the arguments from this list "public, site_local, site_local_inside" can be set

`public` - (Optional) Indicates use of public network (`Bool`).

`site_local` - (Optional) Indicates use of site local network (`Bool`).

`site_local_inside` - (Optional) Indicates use of site local inside network (`Bool`).

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

`name` - (Optional) Name for this rule, will be used to generate metrics (`String`).(Deprecated)

`port` - (Optional) L4 port numbers to match. See [Fast Acl Rules Port ](#fast-acl-rules-port) below for details.

###### One of the arguments from this list "prefix, ip_prefix_set" must be set

`ip_prefix_set` - (Optional) Reference to IP prefix set object. See [Source Ip Prefix Set ](#source-ip-prefix-set) below for details.

`prefix` - (Optional) List of IP prefixes. See [Source Prefix ](#source-prefix) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Selected Vip Address Address

List of IP addresses to match with destination.

###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Site Acl Fast Acl Rules

Fast ACL rules to match.

`action` - (Required) Action to be applied if traffic matched rule (pass, deny or rate limit). See [Fast Acl Rules Action ](#fast-acl-rules-action) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Fast Acl Rules Metadata ](#fast-acl-rules-metadata) below for details.

`name` - (Optional) Name for this rule, will be used to generate metrics (`String`).(Deprecated)

`port` - (Optional) L4 port numbers to match. See [Fast Acl Rules Port ](#fast-acl-rules-port) below for details.

###### One of the arguments from this list "ip_prefix_set, prefix" must be set

`ip_prefix_set` - (Optional) Reference to IP prefix set object. See [Source Ip Prefix Set ](#source-ip-prefix-set) below for details.

`prefix` - (Optional) List of IP prefixes. See [Source Prefix ](#source-prefix) below for details.

### Site Choice Legacy Acl

ACL may be applied at regional edge sites or customer edge sites. Not recommended.

`destination_type` - (Required) 4. Explicit IP and port. See [Legacy Acl Destination Type ](#legacy-acl-destination-type) below for details.

`network_type` - (Required) CE applies Fast ACLs with network type selector as "site_local" and "site_local_inside" only. See [Legacy Acl Network Type ](#legacy-acl-network-type) below for details.

`source_rules` - (Optional) List of Fast ACL rules to be applied to received packets on this site. See [ref](#ref) below for details.

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

###### One of the arguments from this list "outside_network, inside_network" must be set

`inside_network` - (Optional) Site Local Inside network (`Bool`).

`outside_network` - (Optional) Site Local Outside network (`Bool`).

###### One of the arguments from this list "interface_services, vip_services, all_services" must be set

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

### Ver Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ver Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

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

### Vn Type Choice Public

Indicates use of public network.

### Vn Type Choice Site Local

Indicates use of site local network.

### Vn Type Choice Site Local Inside

Indicates use of site local inside network.

Attribute Reference
-------------------

-	`id` - This is the id of the configured fast_acl.
