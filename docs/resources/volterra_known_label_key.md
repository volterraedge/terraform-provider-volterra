---

page_title: "Volterra: volterra_known_label_key" 

description: "The volterra_known_label_key allows CRD of namespace resource on Volterra SaaS"

---

Resource: volterra_known_label_key
==================================

The volterra_known_label_key resource facilitates the creation of a known label key within a tenant's shared workspace. This enables the creation of a Custom Resource Definition (CRD) for the known label key resource on the F5XC platform.

~> **Note:** Please refer to [Known label key api docs](https://docs.cloud.f5.com/docs/api/known-label-key) to learn more

Example Usage
-------------

```hcl
resource "volterra_known_label_key" "volterra_known_label_key" {
    key = "key"
    namespace = "shared"  
    description = "description"  
}
```

Argument Reference
------------------

-	`key` - (Required) Key of the label key (`String`).
-	`namespace` - (Required) Namespace in which to create label key (`String`).
-	`description` - (Optional) Description of what this label key means (`String`).

Attribute Reference
-------------------

-	`id` - ID of the known label key returned
-	`key` - Key of the known label key returned
