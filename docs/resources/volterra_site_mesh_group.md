---

page_title: "Volterra: site_mesh_group"

description: "The site_mesh_group allows CRUD of Site Mesh Group resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_site_mesh_group
=================================

The Site Mesh Group allows CRUD of Site Mesh Group resource on Volterra SaaS

~> **Note:** Please refer to [Site Mesh Group API docs](https://volterra.io/docs/api/site-mesh-group) to learn more

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

`hub` - (Optional) Hub must be empty when Site Mesh Group type is Hub or Full Mesh. See [ref](#ref) below for details.

`full_mesh` - (Optional) Full mesh of tunnels are created between all sites. See [Full Mesh ](#full-mesh) below for details.

`hub_mesh` - (Optional) Mesh of Type Hub (bool).

`spoke_mesh` - (Optional) Mesh of Type Spoke. See [Spoke Mesh ](#spoke-mesh) below for details.

`tunnel_type` - (Optional) Only tunnels of type IPSec are support currently. (`String`).

`type` - (Optional) Type of Site to Site connectivity (Spoke, Hub or Full Mesh) (`String`).

`virtual_site` - (Optional) If 'Type' is Full Mesh, then it gives set of sites that are connected in full mesh.. See [ref](#ref) below for details.

### Control And Data Plane Mesh

and control plane peering across sites.

### Data Plane Mesh

Full Mesh of data plane tunnels across sites.

### Full Mesh

Full mesh of tunnels are created between all sites.

`control_and_data_plane_mesh` - (Optional) and control plane peering across sites (bool).

`data_plane_mesh` - (Optional) Full Mesh of data plane tunnels across sites (bool).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Spoke Mesh

Mesh of Type Spoke.

`hub_mesh_group` - (Required) Spoke sites will connect to all the member sites of Hub Site Mesh Group. See [ref](#ref) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured site_mesh_group.
