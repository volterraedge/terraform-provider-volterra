---

page_title: "Volterra: nginx_service_discovery"

description: "The nginx_service_discovery allows CRUD of Nginx Service Discovery resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------------

Resource volterra_nginx_service_discovery
=========================================

The Nginx Service Discovery allows CRUD of Nginx Service Discovery resource on Volterra SaaS

~> **Note:** Please refer to [Nginx Service Discovery API docs](https://docs.cloud.f5.com/docs-v2/api/nginx-service-discovery) to learn more

Example Usage
-------------

```hcl
resource "volterra_nginx_service_discovery" "example" {
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

`discovery_target` - (Optional) x-displayName: "Discovery Target". See [Discovery Target ](#discovery-target) below for details.

`server_block_filters` - (Optional) x-textBlockPosition: "inside". See [Server Block Filters ](#server-block-filters) below for details.

### Discovery Target

x-displayName: "Discovery Target".

###### One of the arguments from this list "config_sync_group, nginx_instance" must be set

`config_sync_group` - (Optional) x-displayName: "NGINX Config Sync Group". See [Target Config Sync Group ](#target-config-sync-group) below for details.

`nginx_instance` - (Optional) x-displayName: "NGINX Instance". See [Target Nginx Instance ](#target-nginx-instance) below for details.

### Server Block Filters

x-textBlockPosition: "inside".

`name_regex` - (Optional) Regular expression to match the server name or domain that must be discovered (`String`).

`port_ranges` - (Optional) Maximum number of ports allowed is 1024. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Target Config Sync Group

x-displayName: "NGINX Config Sync Group".

`config_sync_group` - (Required) Select new ConfigSyncGroup. See [ref](#ref) below for details.

### Target Nginx Instance

x-displayName: "NGINX Instance".

`nginx_instance` - (Required) Select new NGINX Instance. See [ref](#ref) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured nginx_service_discovery.
