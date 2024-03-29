---

page_title: "Volterra: site_mesh_group"

description: "The site_mesh_group allows CRUD of Site Mesh Group resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_site_mesh_group
=================================

The Site Mesh Group allows CRUD of Site Mesh Group resource on Volterra SaaS

~> **Note:** Please refer to [Site Mesh Group API docs](https://docs.cloud.f5.com/docs/api/site-mesh-group) to learn more

Example Usage
-------------

```hcl
resource "volterra_site_mesh_group" "example" {
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

`hub` - (Optional) Hub must be empty when Site Mesh Group type is Hub or Full Mesh. See [ref](#ref) below for details.(Deprecated)

`full_mesh` - (Optional) Full mesh of tunnels are created between all sites. See [Mesh Choice Full Mesh ](#mesh-choice-full-mesh) below for details.

`hub_mesh` - (Optional) Mesh of Type Hub. See [Mesh Choice Hub Mesh ](#mesh-choice-hub-mesh) below for details.

`spoke_mesh` - (Optional) Mesh of Type Spoke. See [Mesh Choice Spoke Mesh ](#mesh-choice-spoke-mesh) below for details.

`type` - (Optional) Type of Site to Site connectivity (Spoke, Hub or Full Mesh) (`String`).(Deprecated)

`virtual_site` - (Optional) If 'Type' is Full Mesh, then it gives set of sites that are connected in full mesh.. See [ref](#ref) below for details.

### Full Mesh Choice Control And Data Plane Mesh

and control plane peering across sites.

### Full Mesh Choice Data Plane Mesh

Full Mesh of data plane tunnels across sites.

### Hub Full Mesh Choice Control And Data Plane Mesh

and control plane peering across sites.

### Hub Full Mesh Choice Data Plane Mesh

Hub Full Mesh of data plane tunnels across sites.

### Mesh Choice Full Mesh

Full mesh of tunnels are created between all sites.

###### One of the arguments from this list "data_plane_mesh, control_and_data_plane_mesh" can be set

`control_and_data_plane_mesh` - (Optional) and control plane peering across sites (`Bool`).

`data_plane_mesh` - (Optional) Full Mesh of data plane tunnels across sites (`Bool`).

### Mesh Choice Hub Mesh

Mesh of Type Hub.

### Mesh Choice Spoke Mesh

Mesh of Type Spoke.

`hub_mesh_group` - (Required) Spoke sites will connect to all the member sites of Hub Site Mesh Group. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Spoke Hub Mesh Choice Control And Data Plane Mesh

and control plane peering with the hub site/s.

### Spoke Hub Mesh Choice Data Plane Mesh

Mesh of data plane tunnels to the hub site/s.

Attribute Reference
-------------------

-	`id` - This is the id of the configured site_mesh_group.
