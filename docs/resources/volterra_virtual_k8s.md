---

page_title: "Volterra: virtual_k8s"

description: "The virtual_k8s allows CRUD of Virtual K8s resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_virtual_k8s
=============================

The Virtual K8s allows CRUD of Virtual K8s resource on Volterra SaaS

~> **Note:** Please refer to [Virtual K8s API docs](https://volterra.io/docs/api/virtual-k8s) to learn more

Example Usage
-------------

```hcl
resource "volterra_virtual_k8s" "example" {
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

`default_flavor_ref` - (Optional) Default workfload flavor for all workloads launched in this Virtual K8s. See [ref](#ref) below for details.

`disabled` - (Optional) setting the ves.io/serviceIsolation annotation to true (bool).

`isolated` - (Optional) overridden by the K8s service via setting the ves.io/serviceIsolation annotation to false (bool).

`vsite_refs` - (Optional) Kubernetes API resource object. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured virtual_k8s.
