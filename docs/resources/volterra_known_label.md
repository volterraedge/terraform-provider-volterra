---

page_title: "Volterra: volterra_known_label"
description: "The volterra_known_label allows CRD of namespace resource on Volterra SaaS"

---

Resource: volterra_known_label
==============================

The volterra_known_label resource enables the addition of values to a key within an independent workspace within a tenant. This facilitates the creation of a Custom Resource Definition (CRD) for the known label resource on the Volterra platform.

~> **Note:** Please refer to [Known label api docs](https://docs.cloud.f5.com/docs/api/known-label) to learn more

Example Usage
-------------

```hcl
resource "volterra_known_label" "volterra_known_label" {
  key = "key"
  namespace = "shared"
  value       = "value"
  description = "description of value"
}
```

Argument Reference
------------------

-	`key` - (Required) Key of the label (`String`).
-	`namespace` - (Required) Namespace in which to create the label (`String`).
-	`value` - (Required) Value of the label (`String`).
-	`description` - (Optional) Description of the label to be created (`String`).

Attribute Reference
-------------------

-	`id` - ID of the known label returned
-	`key` - Key of the known label returned
-	`value` - Key of the known label returned
