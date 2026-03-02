---

page_title: "Volterra: virtual_address"

description: "The virtual_address allows CRUD of Virtual Address resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_virtual_address
=================================

The Virtual Address allows CRUD of Virtual Address resource on Volterra SaaS

~> **Note:** Please refer to [Virtual Address API docs](https://docs.cloud.f5.com/docs-v2/api/views-virtual-address) to learn more

Example Usage
-------------

```hcl
resource "volterra_virtual_address" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "private public" must be set

  public {
    // One of the arguments from this list "global regional vk8s_service" can be set

    global {
      // One of the arguments from this list "f5_global_network" can be set

      f5_global_network = true

      // One of the arguments from this list "specific_vip tenant_public_default_vip" can be set

      tenant_public_default_vip = true
    }
  }
  state {
    // One of the arguments from this list "state_disabled state_enabled" can be set

    state_enabled {}
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

###### One of the arguments from this list "private, public" must be set

`private` - (Optional) x-displayName: "Private". See [Deployment Type Choice Private ](#deployment-type-choice-private) below for details.

`public` - (Optional) x-displayName: "Public". See [Deployment Type Choice Public ](#deployment-type-choice-public) below for details.

`state` - (Required) Displays the current status of the virtual address.. See [State ](#state) below for details.

### State

Displays the current status of the virtual address..

###### One of the arguments from this list "state_disabled, state_enabled" can be set

`state_disabled` - (Optional) x-displayName: "Disabled". See [State Choice State Disabled ](#state-choice-state-disabled) below for details.

`state_enabled` - (Optional) x-displayName: "Enabled". See [State Choice State Enabled ](#state-choice-state-enabled) below for details.

### Deployment Type Choice Private

x-displayName: "Private".

###### One of the arguments from this list "segment, site, virtual_ce_sites" can be set

`segment` - (Optional) x-displayName: "Segment". See [Site Type Choice Segment ](#site-type-choice-segment) below for details.

`site` - (Optional) x-displayName: "Site". See [Site Type Choice Site ](#site-type-choice-site) below for details.

`virtual_ce_sites` - (Optional) x-displayName: "Virtual Sites". See [Site Type Choice Virtual Ce Sites ](#site-type-choice-virtual-ce-sites) below for details.

### Deployment Type Choice Public

x-displayName: "Public".

###### One of the arguments from this list "global, regional, vk8s_service" can be set

`global` - (Optional) x-displayName: "Global". See [Where Selection Choice Global ](#where-selection-choice-global) below for details.

`regional` - (Optional) x-displayName: "Regional". See [Where Selection Choice Regional ](#where-selection-choice-regional) below for details.

`vk8s_service` - (Optional) Advertise on vK8s Service Network on RE.. See [Where Selection Choice Vk8s Service ](#where-selection-choice-vk8s-service) below for details.

### Network Choice Site Local Inside Network

x-displayName: "Site Local Inside (Local VRF)".

### Network Choice Site Local Network

x-displayName: "Site Local Outside (Local VRF)".

### Network Selection Choice F5 Global Network

x-displayName: "F5 Global Network".

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Site Or Virtual Network Type Regional Site

Select Regional Edge.

###### One of the arguments from this list "f5_global_network" can be set

`f5_global_network` - (Optional) x-displayName: "F5 Global Network" (`Bool`).

`regional_edge_site` - (Optional) Reference to site object. See [ref](#ref) below for details.

###### One of the arguments from this list "specific_vip, tenant_public_default_vip" can be set

`specific_vip` - (Optional) Specific VIP. See [Vip Choice Specific Vip ](#vip-choice-specific-vip) below for details.

`tenant_public_default_vip` - (Optional) Tenant VIP (`Bool`).

### Site Type Choice Segment

x-displayName: "Segment".

`segment_network` - (Optional) x-displayName: "Segment". See [ref](#ref) below for details.

###### One of the arguments from this list "custom_vip" can be set

`custom_vip` - (Optional) Tenant VIP. See [Vip Choice Custom Vip ](#vip-choice-custom-vip) below for details.

### Site Type Choice Site

x-displayName: "Site".

###### One of the arguments from this list "segment_network, site_local_inside_network, site_local_network" can be set

`segment_network` - (Optional) x-displayName: "Segment (Global VRF)". See [ref](#ref) below for details.

`site_local_inside_network` - (Optional) x-displayName: "Site Local Inside (Local VRF)" (`Bool`).

`site_local_network` - (Optional) x-displayName: "Site Local Outside (Local VRF)" (`Bool`).

`site_name` - (Optional) x-example: "smsv2-site-name". See [ref](#ref) below for details.

###### One of the arguments from this list "custom_vip" can be set

`custom_vip` - (Optional) Tenant VIP. See [Vip Choice Custom Vip ](#vip-choice-custom-vip) below for details.

### Site Type Choice Virtual Ce Sites

x-displayName: "Virtual Sites".

`ce_virtual_site` - (Optional) x-example: "shared/ce-site". See [ref](#ref) below for details.

###### One of the arguments from this list "segment_network, site_local_inside_network, site_local_network" can be set

`segment_network` - (Optional) x-displayName: "Segment (Global VRF)". See [ref](#ref) below for details.

`site_local_inside_network` - (Optional) x-displayName: "Site Local Inside (Local VRF)" (`Bool`).

`site_local_network` - (Optional) x-displayName: "Site Local Outside (Local VRF)" (`Bool`).

###### One of the arguments from this list "custom_vip" can be set

`custom_vip` - (Optional) Custom VIP. See [Vip Choice Custom Vip ](#vip-choice-custom-vip) below for details.

### State Choice State Disabled

x-displayName: "Disabled".

### State Choice State Enabled

x-displayName: "Enabled".

### Vip Choice Custom Vip

Tenant VIP.

###### One of the arguments from this list "ip_v4, ip_v6" can be set

`ip_v4` - (Optional) V4 IP (`String`).

`ip_v6` - (Optional) Optional common virtual V6 IP across all nodes to be used as automatic VIP. (`String`).

### Vip Choice Specific Vip

Specific VIP.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Vip Choice Tenant Public Default Vip

Tenant VIP.

### Where Selection Choice Global

x-displayName: "Global".

###### One of the arguments from this list "f5_global_network" can be set

`f5_global_network` - (Optional) x-displayName: "F5 Global Network" (`Bool`).

###### One of the arguments from this list "specific_vip, tenant_public_default_vip" can be set

`specific_vip` - (Optional) Specific VIP. See [Vip Choice Specific Vip ](#vip-choice-specific-vip) below for details.

`tenant_public_default_vip` - (Optional) Tenant VIP (`Bool`).

### Where Selection Choice Regional

x-displayName: "Regional".

###### One of the arguments from this list "regional_site" can be set

`regional_site` - (Optional) Select Regional Edge. See [Site Or Virtual Network Type Regional Site ](#site-or-virtual-network-type-regional-site) below for details.

### Where Selection Choice Vk8s Service

Advertise on vK8s Service Network on RE..

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

Attribute Reference
-------------------

*   `id` - This is the id of the configured virtual_address.
