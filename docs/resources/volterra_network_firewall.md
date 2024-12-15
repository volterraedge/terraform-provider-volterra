---

page_title: "Volterra: network_firewall"

description: "The network_firewall allows CRUD of Network Firewall resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_network_firewall
==================================

The Network Firewall allows CRUD of Network Firewall resource on Volterra SaaS

~> **Note:** Please refer to [Network Firewall API docs](https://docs.cloud.f5.com/docs-v2/api/network-firewall) to learn more

Example Usage
-------------

```hcl
resource "volterra_network_firewall" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "active_fast_acls disable_fast_acl fast_acl_set" must be set

  disable_fast_acl = true

  // One of the arguments from this list "active_forward_proxy_policies disable_forward_proxy_policy forward_proxy_policy_set" must be set

  active_forward_proxy_policies {
    forward_proxy_policies {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  // One of the arguments from this list "active_enhanced_firewall_policies active_network_policies disable_network_policy network_policy_set" must be set

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

###### One of the arguments from this list "active_fast_acls, disable_fast_acl, fast_acl_set" must be set

`active_fast_acls` - (Optional) Fast ACL Active for ths network firewall.. See [Fast Acl Choice Active Fast Acls ](#fast-acl-choice-active-fast-acls) below for details.

`disable_fast_acl` - (Optional) Fast ACL is disabled for this network firewall (`Bool`).

`fast_acl_set` - (Optional) The list of Virtual Networks / Interfaces is selected by the Fast ACL set object. See [ref](#ref) below for details.(Deprecated)

###### One of the arguments from this list "active_forward_proxy_policies, disable_forward_proxy_policy, forward_proxy_policy_set" must be set

`active_forward_proxy_policies` - (Optional) L7 firewall for forward proxy.. See [Forward Proxy Policy Choice Active Forward Proxy Policies ](#forward-proxy-policy-choice-active-forward-proxy-policies) below for details.

`disable_forward_proxy_policy` - (Optional) Forward Proxy Policy is disabled for this network firewall (`Bool`).

`forward_proxy_policy_set` - (Optional) L7 firewall for forward proxy. Assign service_policy_set to be used for forward proxies in this firewall.. See [ref](#ref) below for details.(Deprecated)

###### One of the arguments from this list "active_enhanced_firewall_policies, active_network_policies, disable_network_policy, network_policy_set" must be set

`active_enhanced_firewall_policies` - (Optional) with an additional option for service insertion.. See [Network Policy Choice Active Enhanced Firewall Policies ](#network-policy-choice-active-enhanced-firewall-policies) below for details.

`active_network_policies` - (Optional) Active firewall policies for this network firewall(L3/L4 firewall).. See [Network Policy Choice Active Network Policies ](#network-policy-choice-active-network-policies) below for details.

`disable_network_policy` - (Optional) Firewall Policy is disabled for this network firewall (`Bool`).

`network_policy_set` - (Optional) - Site Local. See [ref](#ref) below for details.(Deprecated)

### Fast Acl Choice Active Fast Acls

Fast ACL Active for ths network firewall..

`fast_acls` - (Required) Ordered List of Fast ACL(s) active for this network firewall. See [ref](#ref) below for details.

### Forward Proxy Policy Choice Active Forward Proxy Policies

L7 firewall for forward proxy..

`forward_proxy_policies` - (Required) Ordered List of Forward Proxy Policies active. See [ref](#ref) below for details.

### Network Policy Choice Active Enhanced Firewall Policies

with an additional option for service insertion..

`enhanced_firewall_policies` - (Required) Ordered List of Enhanced Firewall Policies active. See [ref](#ref) below for details.

### Network Policy Choice Active Network Policies

Active firewall policies for this network firewall(L3/L4 firewall)..

`network_policies` - (Required) Ordered List of Firewall Policies active for this network firewall. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_firewall.
