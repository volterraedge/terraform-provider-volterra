---
page_title: "Volterra: service_policy Data Source"

description: "Data source for service_policy resource"
---

# Data Source: volterra_service_policy

Service Policy creates logical independent workspace within a tenant. Data Source reads the service_policy object and gets values like tenant_name, id of the service_policy object.

## Example Usage

```hcl
data volterra_service_policy" "example" {
  name = "example"
}
```

## Argument Reference

* `name` - (Required) Name of the service_policy to be queried


## Attribute Reference

* `id` - ID of the service_policy returned

* `tenant_name` - Tenant name configured in volterra

