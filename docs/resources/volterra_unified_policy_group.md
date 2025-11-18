---

page_title: "Volterra: unified_policy_group"

description: "The unified_policy_group allows CRUD of Unified Policy Group resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------

Resource volterra_unified_policy_group
======================================

The Unified Policy Group allows CRUD of Unified Policy Group resource on Volterra SaaS

~> **Note:** Please refer to [Unified Policy Group API docs](https://docs.cloud.f5.com/docs-v2/api/unified-policy-group) to learn more

Example Usage
-------------

```hcl
resource "volterra_unified_policy_group" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
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

`east_west_policy` - (Optional) Manage policies for inter and intra segment traffic. See [East West Policy ](#east-west-policy) below for details.

`internet_policy` - (Optional) Manage policies internet traffic. See [Internet Policy ](#internet-policy) below for details.

### East West Policy

Manage policies for inter and intra segment traffic.

`policies` - (Optional) x-displayName: "Policies". See [East West Policy Policies ](#east-west-policy-policies) below for details.

### Internet Policy

Manage policies internet traffic.

`policies` - (Optional) x-displayName: "Policies". See [Internet Policy Policies ](#internet-policy-policies) below for details.

### East West Policy Policies

x-displayName: "Policies".

###### One of the arguments from this list "active_enhanced_firewall_policies, no_firewall_policy" must be set

`active_enhanced_firewall_policies` - (Optional) x-displayName: "Enable Enhanced Firewall Policies". See [Firewall Policy Choice Active Enhanced Firewall Policies ](#firewall-policy-choice-active-enhanced-firewall-policies) below for details.

`no_firewall_policy` - (Optional) x-displayName: "Disable Enhanced Firewall Policies" (`Bool`).

###### One of the arguments from this list "active_forward_proxy_policies, active_service_policies, no_forward_proxy" must be set

`active_forward_proxy_policies` - (Optional) x-displayName: "Enable Forward Proxy Policies". See [Forward Proxy Choice Active Forward Proxy Policies ](#forward-proxy-choice-active-forward-proxy-policies) below for details.

`active_service_policies` - (Optional) For advanced outbound control, add Service Policies. See [Forward Proxy Choice Active Service Policies ](#forward-proxy-choice-active-service-policies) below for details.

`no_forward_proxy` - (Optional) x-displayName: "Disable Forward Proxy Policies" (`Bool`).

### Firewall Policy Choice Active Enhanced Firewall Policies

x-displayName: "Enable Enhanced Firewall Policies".

`enhanced_firewall_policies` - (Required) Ordered List of Enhaned Firewall Policies. See [ref](#ref) below for details.

### Firewall Policy Choice No Firewall Policy

x-displayName: "Disable Enhanced Firewall Policies".

### Forward Proxy Choice Active Forward Proxy Policies

x-displayName: "Enable Forward Proxy Policies".

`forward_proxy_policies` - (Required) List of Forward Proxy Policies. See [ref](#ref) below for details.

### Forward Proxy Choice Active Service Policies

For advanced outbound control, add Service Policies.

`service_policies` - (Optional) A list of references to service_policy objects.. See [ref](#ref) below for details.

### Forward Proxy Choice No Forward Proxy

x-displayName: "Disable Forward Proxy Policies".

### Internet Policy Policies

x-displayName: "Policies".

###### One of the arguments from this list "active_enhanced_firewall_policies, no_firewall_policy" must be set

`active_enhanced_firewall_policies` - (Optional) x-displayName: "Enable Enhanced Firewall Policies". See [Firewall Policy Choice Active Enhanced Firewall Policies ](#firewall-policy-choice-active-enhanced-firewall-policies) below for details.

`no_firewall_policy` - (Optional) x-displayName: "Disable Enhanced Firewall Policies" (`Bool`).

###### One of the arguments from this list "active_forward_proxy_policies, active_service_policies, no_forward_proxy" must be set

`active_forward_proxy_policies` - (Optional) x-displayName: "Enable Forward Proxy Policies". See [Forward Proxy Choice Active Forward Proxy Policies ](#forward-proxy-choice-active-forward-proxy-policies) below for details.

`active_service_policies` - (Optional) For advanced outbound control, add Service Policies. See [Forward Proxy Choice Active Service Policies ](#forward-proxy-choice-active-service-policies) below for details.

`no_forward_proxy` - (Optional) x-displayName: "Disable Forward Proxy Policies" (`Bool`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured unified_policy_group.
