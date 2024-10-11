---
page_title: "Volterra: namespace Data Source"

description: "Data source for namespace resource"
---

# Data Source: volterra_namespace

Namespace creates logical independent workspace within a tenant. Data Source reads the namespace object and gets values like tenant_name, id of the namespace object.

## Example Usage

```hcl
data "volterra_namespace" "example" {
  name = "example"
}
```

## Argument Reference

* `name` - (Required) Name of the namespace to be queried


## Attribute Reference

* `id` - ID of the namespace returned

* `tenant_name` - Tenant name configured in volterra

