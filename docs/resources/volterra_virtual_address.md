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

  // One of the arguments from this list "public" must be set

  public {
    // One of the arguments from this list "global" can be set

    global {
      // One of the arguments from this list "f5_global_network" can be set

      f5_global_network = true

      // One of the arguments from this list "specific_vip" can be set

      specific_vip {
        public_ip {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
      }
    }
  }
  state {
    // One of the arguments from this list "state_disabled state_enabled" can be set

    state_disabled {}
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

###### One of the arguments from this list "public" must be set

`public` - (Optional) x-displayName: "Public". See [Deployment Type Choice Public ](#deployment-type-choice-public) below for details.

`state` - (Required) Displays the current status of the virtual address.. See [State ](#state) below for details.

### State

Displays the current status of the virtual address..

###### One of the arguments from this list "state_disabled, state_enabled" can be set

`state_disabled` - (Optional) x-displayName: "Disabled". See [State Choice State Disabled ](#state-choice-state-disabled) below for details.

`state_enabled` - (Optional) x-displayName: "Enabled". See [State Choice State Enabled ](#state-choice-state-enabled) below for details.

### Deployment Type Choice Public

x-displayName: "Public".

###### One of the arguments from this list "global" can be set

`global` - (Optional) x-displayName: "Global". See [Where Selection Choice Global ](#where-selection-choice-global) below for details.

### Network Selection Choice F5 Global Network

x-displayName: "F5 Global Network".

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### State Choice State Disabled

x-displayName: "Disabled".

### State Choice State Enabled

x-displayName: "Enabled".

### Vip Choice Specific Vip

Specific VIP.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Where Selection Choice Global

x-displayName: "Global".

###### One of the arguments from this list "f5_global_network" can be set

`f5_global_network` - (Optional) x-displayName: "F5 Global Network" (`Bool`).

###### One of the arguments from this list "specific_vip" can be set

`specific_vip` - (Optional) Specific VIP. See [Vip Choice Specific Vip ](#vip-choice-specific-vip) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured virtual_address.
