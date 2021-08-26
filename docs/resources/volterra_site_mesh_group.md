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
  name        = "acmecorp-web"
  namespace   = "staging"
  tunnel_type = ["SITE_TO_SITE_TUNNEL_IPSEC"]
  type        = ["mesh-1"]
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

`tunnel_type` - (Required) Only tunnels of type IPSec are support currently. (`String`).

`type` - (Required) Type of Site to Site connectivity (Spoke, Hub or Full Mesh) (`String`).

`virtual_site` - (Optional) If 'Type' is Full Mesh, then it gives set of sites that are connected in full mesh.. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured site_mesh_group.
