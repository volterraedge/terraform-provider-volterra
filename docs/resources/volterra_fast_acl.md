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

  legacy_acl {
    destination_type {
      // One of the arguments from this list "shared_vip_services interface_services vip_services all_services destination_ip_address selected_vip_address" must be set
      shared_vip_services = true
    }

    network_type {
      // One of the arguments from this list "site_local site_local_inside public" must be set
      site_local = true
    }

    source_rules {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
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

### Address

List of IP addresses to match with destination.

`ipv4` - (Optional) IPv4 Address. See [Ipv4 ](#ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ipv6 ](#ipv6) below for details.

### All

Matches all port.

### All Public Vips

Apply this Fast ACL to all public vips.

### All Services

Regional Edge: Applies the configuration to all the VIPs assigned to tenant.

### Default Tenant Vip

Apply this Fast ACL to Default(dedicated) Tenant VIP.

### Destination Ip Address

Regional Edge : Tenant can not configure it.

`address` - (Optional) List of IP addresses to match with destination. See [Address ](#address) below for details.

`ports` - (Optional) Special value "0" means all valid ports on the VIPs. See [Ports ](#ports) below for details.

`protocol` - (Optional) Protocol to match in the traffic (`String`).

### Destination Type

1.	Explicit IP and port.

`all_services` - (Optional) Regional Edge: Applies the configuration to all the VIPs assigned to tenant (bool).

`destination_ip_address` - (Optional) Regional Edge : Tenant can not configure it. See [Destination Ip Address ](#destination-ip-address) below for details.

`interface_services` - (Optional) Regional Edge : tenants CANNOT use this option on RE (bool).

`selected_vip_address` - (Optional) Valid only for RE.. See [Selected Vip Address ](#selected-vip-address) below for details.

`shared_vip_services` - (Optional) Regional Edge: Applies configuration on all shared VIPs used for services (bool).

`vip_services` - (Optional) Regional Edge: Applies the configuration to VIP which has been assigned by default to tenant. Not applicable for shared VIP(s) (bool).

### Dns

Matches dns port 53.

### Fast Acl Rules

Fast ACL rules to match .

`action` - (Required) Action to be applied if traffic matched rule (pass, deny or rate limit). See [Action ](#action) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Metadata ](#metadata) below for details.

`name` - (Optional) Name for this rule, will be used to generate metrics (`String`).

`port` - (Optional) L4 port numbers to match. See [Port ](#port) below for details.

`ip_prefix_set` - (Optional) Reference to IP prefix set object. See [Ip Prefix Set ](#ip-prefix-set) below for details.

`prefix` - (Optional) List of IP prefixes. See [Prefix ](#prefix) below for details.

### Inside Network

Site Local Inside network.

### Interface Services

Regional Edge : tenants CANNOT use this option on RE.

### Ip Prefix Set

Reference to IP prefix set object.

`ref` - (Optional) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.

### Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

### Legacy Acl

ACL may be applied at regional edge sites or customer edge sites. Not recommended.

`destination_type` - (Required) 4. Explicit IP and port. See [Destination Type ](#destination-type) below for details.

`network_type` - (Required) CE applies Fast ACLs with network type selector as "site_local" and "site_local_inside" only. See [Network Type ](#network-type) below for details.

`source_rules` - (Optional) List of Fast ACL rules to be applied to received packets on this site. See [ref](#ref) below for details.

### Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Network Type

CE applies Fast ACLs with network type selector as "site_local" and "site_local_inside" only.

`public` - (Optional) Indicates use of public network (bool).

`site_local` - (Optional) Indicates use of site local network (bool).

`site_local_inside` - (Optional) Indicates use of site local inside network (bool).

### Outside Network

Site Local Outside network.

### Policer Action

Reference to policer object to which traffic would be subjected.

`ref` - (Optional) A policer direct reference. See [ref](#ref) below for details.

### Port

L4 port numbers to match.

`all` - (Optional) Matches all port (bool).

`dns` - (Optional) Matches dns port 53 (bool).

`user_defined` - (Optional) Matches the user defined port (`Int`).

### Ports

Special value "0" means all valid ports on the VIPs.

`all` - (Optional) Matches all port (bool).

`dns` - (Optional) Matches dns port 53 (bool).

`user_defined` - (Optional) Matches the user defined port (`Int`).

### Prefix

List of IP prefixes.

`prefix` - (Optional) IP Address prefix in string format. String must contain both prefix and prefix-length (`String`).

### Protocol Policer Action

Reference to protocol based policer object.

`ref` - (Optional) Reference to protocol policer object. See [ref](#ref) below for details.

### Public

Indicates use of public network.

### Re Acl

ACL will be applied at regional edge sites.

`fast_acl_rules` - (Optional) Fast ACL rules to match . See [Fast Acl Rules ](#fast-acl-rules) below for details.

`all_public_vips` - (Optional) Apply this Fast ACL to all public vips (bool).

`default_tenant_vip` - (Optional) Apply this Fast ACL to Default(dedicated) Tenant VIP (bool).

`selected_tenant_vip` - (Optional) Apply this Fast ACL to List of some selected public VIP(s). See [Selected Tenant Vip ](#selected-tenant-vip) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Selected Tenant Vip

Apply this Fast ACL to List of some selected public VIP(s).

`default_tenant_vip` - (Optional) Include tenant vip in list of specific VIP(s) (`Bool`).

`public_ip_refs` - (Required) Select additional public VIP(s) . See [ref](#ref) below for details.

### Selected Vip Address

Valid only for RE..

`address` - (Optional) List of IP addresses to match with destination. See [Address ](#address) below for details.

### Shared Vip Services

Regional Edge: Applies configuration on all shared VIPs used for services.

### Site Acl

ACL will be applied at customer edge sites.

`fast_acl_rules` - (Optional) Fast ACL rules to match . See [Fast Acl Rules ](#fast-acl-rules) below for details.

`inside_network` - (Optional) Site Local Inside network (bool).

`outside_network` - (Optional) Site Local Outside network (bool).

`all_services` - (Optional) Port and protocol is picked up from advertise policies (bool).

`interface_services` - (Optional) Port and protocol is picked up from advertise policies (bool).

`vip_services` - (Optional) Port and protocol is picked up from advertise policies (bool).

### Site Local

Indicates use of site local network.

### Site Local Inside

Indicates use of site local inside network.

### Vip Services

Regional Edge: Applies the configuration to VIP which has been assigned by default to tenant. Not applicable for shared VIP(s).

Attribute Reference
-------------------

-	`id` - This is the id of the configured fast_acl.
