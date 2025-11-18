---

page_title: "Volterra: workload_flavor"

description: "The workload_flavor allows CRUD of Workload Flavor resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_workload_flavor
=================================

The Workload Flavor allows CRUD of Workload Flavor resource on Volterra SaaS

~> **Note:** Please refer to [Workload Flavor API docs](https://docs.cloud.f5.com/docs-v2/api/workload-flavor) to learn more

Example Usage
-------------

```hcl
resource "volterra_workload_flavor" "example" {
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

`ephemeral_storage` - (Optional) Ephemeral storage in MiB (mebibyte) allocated for the workload_flavor. (`Int`).

`memory` - (Optional) Memory in MiB (mebibyte) allocated for the workload_flavor. (`Int`).

`vcpus` - (Optional) Number of vCPUs allocated for the workload_flavor. Each vCPU is a thread on a CPU core. (`Float`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured workload_flavor.
