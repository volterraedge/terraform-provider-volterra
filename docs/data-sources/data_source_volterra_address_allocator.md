---
page_title: "Volterra: volterra_address_allocator Data Source"

description: "Data source for volterra_address_allocator resource"
---

# Data Source: volterra_address_allocator

volterra_address_allocator is used to get mode of the allocator object and allocation map. Mode determines if the allocation is limited to VERs within the local site / cluster (local) or if the allocation is across different sites (global).


## Example Usage

```hcl
data "volterra_address_allocator" "example" {
  name = "aalloc-1"
  namespace = "system"
}
```

## Argument Reference

* `name` - (Required) Name of the tunnel

* `namespace` - (Required) Namespace where the tunnel is configure. Example `system`

## Attribute Reference

* `mode` - Mode of the address allocator. Expected values are `LOCAL` `GLOBAL_PER_SITE_NODE` (`string`)

* `allocation_map` - Map of per site node allocation (`map(string)`)