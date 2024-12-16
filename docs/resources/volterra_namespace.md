---

page_title: "Volterra: volterra_namespace" 

description: "The volterra_namespace allows CRD of namespace resource on Volterra SaaS"

---

Resource: volterra_namespace
============================

Namespace creates logical independent workspace within a tenant. This allows CRD of namespace resource on Volterra.

~> **Note:** Please refer to [Namespace api docs](https://docs.cloud.f5.com/docs/api/namespace) to learn more

Example Usage
-------------

```hcl
resource "volterra_namespace" "example" {
  name = "example"
}
```

Argument Reference
------------------

-	`name` - (Required) Name of the namespace to be queried

Attribute Reference
-------------------

-	`id` - ID of the namespace returned

-	`tenant_name` - Tenant name configured in volterra
