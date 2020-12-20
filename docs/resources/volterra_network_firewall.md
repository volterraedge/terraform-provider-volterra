---

page_title: "Volterra: network_firewall"

description: "The network_firewall allows CRUD of Network Firewall resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_network_firewall
==================================

The Network Firewall allows CRUD of Network Firewall resource on Volterra SaaS

~> **Note:** Please refer to [Network Firewall API docs](https://volterra.io/docs/api/network-firewall) to learn more

Example Usage
-------------

```hcl
resource "volterra_network_firewall" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "disable_fast_acl active_fast_acls fast_acl_set" must be set
  disable_fast_acl = true

  // One of the arguments from this list "disable_forward_proxy_policy active_forward_proxy_policies forward_proxy_policy_set" must be set
  disable_forward_proxy_policy = true

  // One of the arguments from this list "disable_network_policy active_network_policies network_policy_set" must be set
  disable_network_policy = true
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

`active_fast_acls` - (Optional) Fast ACL Active for ths network firewall.. See [Active Fast Acls ](#active-fast-acls) below for details.

`disable_fast_acl` - (Optional) Fast ACL is disabled for this network firewall (bool).

`fast_acl_set` - (Optional) The list of Virtual Networks / Interfaces is selected by the Fast ACL set object. See [ref](#ref) below for details.

`active_forward_proxy_policies` - (Optional) L7 firewall for forward proxy. . See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`disable_forward_proxy_policy` - (Optional) Forward Proxy Policy is disabled for this network firewall (bool).

`forward_proxy_policy_set` - (Optional) L7 firewall for forward proxy. Assign service_policy_set to be used for forward proxies in this firewall.. See [ref](#ref) below for details.

`active_network_policies` - (Optional) Active network policies for this network firewall(L3/L4 firewall).. See [Active Network Policies ](#active-network-policies) below for details.

`disable_network_policy` - (Optional) Network Policy is disabled for this network firewall (bool).

`network_policy_set` - (Optional) - Site Local. See [ref](#ref) below for details.

### Active Fast Acls

Fast ACL Active for ths network firewall..

`fast_acls` - (Required) Ordered List of Fast ACL(s) active for this network firewall. See [ref](#ref) below for details.

### Active Forward Proxy Policies

L7 firewall for forward proxy. .

`forward_proxy_policies` - (Required) Ordered List of Network Policies active for this network firewall. See [ref](#ref) below for details.

### Active Network Policies

Active network policies for this network firewall(L3/L4 firewall)..

`network_policies` - (Required) Ordered List of Network Policies active for this network firewall. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_firewall.
