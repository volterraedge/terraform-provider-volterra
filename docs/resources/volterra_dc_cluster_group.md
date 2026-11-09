---

page_title: "Volterra: dc_cluster_group"

description: "The dc_cluster_group allows CRUD of Dc Cluster Group resource on Volterra SaaS"

---

Resource volterra_dc_cluster_group
==================================

The Dc Cluster Group allows CRUD of Dc Cluster Group resource on Volterra SaaS

~> **Note:** Please refer to [Dc Cluster Group API docs](https://docs.cloud.f5.com/docs/api/dc-cluster-group) to learn more

Example Usage
-------------

```hcl
resource "volterra_dc_cluster_group" "example" {
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

`type` - (Optional) DC Cluster Group Mesh Type configuration. See [Type](#type) below for details.

### Type

DC Cluster Group Mesh Type configuration. Choice of data plane, data plane and control plane dc cluster mesh.

###### One of the arguments from this list "data_plane_mesh, control_and_data_plane_mesh" can be set

`data_plane_mesh` - (Optional) Full Mesh of Data plane connectivity across sites (`Bool`).

`control_and_data_plane_mesh` - (Optional) Full Mesh of data plane connectivity across sites and control plane peering across sites (`Bool`).

Attribute Reference
-------------------

-   `id` - This is the id of the configured dc_cluster_group.
