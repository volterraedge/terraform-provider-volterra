---

page_title: "Volterra: bigip_site"

description: "The bigip_site allows CRUD of Bigip Site resource on Volterra SaaS"
---------------------------------------------------------------------------------

Resource volterra_bigip_site
============================

The Bigip Site allows CRUD of Bigip Site resource on Volterra SaaS

~> **Note:** Please refer to [Bigip Site API docs](https://docs.cloud.f5.com/docs-v2/api/views-bigip-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_bigip_site" "example" {
  name         = "acmecorp-web"
  namespace    = "staging"
  certified_hw = ["rseries-4k"]

  // One of the arguments from this list "log_receiver logs_streaming_disabled" must be set

  logs_streaming_disabled = true
  master_nodes {
    name = "master-0"

    public_ip = "192.168.0.156"
  }
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

`certified_hw` - (Required) x-required (`String`).

`coordinates` - (Optional) Coordinates of the site, longitude and latitude. See [Coordinates ](#coordinates) below for details.

###### One of the arguments from this list "log_receiver, logs_streaming_disabled" must be set

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (`Bool`).

`master_nodes` - (Required) Configuration of master nodes. See [Master Nodes ](#master-nodes) below for details.

`sw` - (Optional) F5XC Software Details. See [Sw ](#sw) below for details.

### Coordinates

Coordinates of the site, longitude and latitude.

`latitude` - (Optional) Latitude of the site location (`Float`).

`longitude` - (Optional) longitude of site location (`Float`).

### Master Nodes

Configuration of master nodes.

`name` - (Required) Names of master node (`String`).

`public_ip` - (Optional) via Site Mesh Group (`String`).

### Sw

F5XC Software Details.

###### One of the arguments from this list "default_sw_version, volterra_software_version" must be set

`default_sw_version` - (Optional) Will assign latest available F5XC Software Version (`Bool`).

`volterra_software_version` - (Optional) Specify a F5XC Software Version to be used e.g. crt-20210329-1002. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Volterra Sw Version Choice Default Sw Version

Will assign latest available F5XC Software Version.

Attribute Reference
-------------------

-	`id` - This is the id of the configured bigip_site.
