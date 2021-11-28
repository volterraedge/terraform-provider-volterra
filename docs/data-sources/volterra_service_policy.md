---
page_title: "Volterra: service_policy Data Source"

description: "Data source for service_policy resource"
---

# Data Source: volterra_service_policy

A service_policy object consists of an unordered list of predicates and a list of service policy rules. Data Source reads the service_policy object and gets values like name, namespace, id of the service_policy object.

## Example Usage

```hcl
data volterra_service_policy" "example" {
  name = "example"
  namespace = "system"
}
```

## Argument Reference

* `name` - (Required) Name of the service_policy to be queried

* `namespace` - (Required) Namespace where the tunnel is configure. Example `system`


## Attribute Reference

* `id` - ID of the service_policy returned
