---

page_title: "Volterra: k8s_cluster"

description: "The k8s_cluster allows CRUD of K8s Cluster resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_k8s_cluster
=============================

The K8s Cluster allows CRUD of K8s Cluster resource on Volterra SaaS

~> **Note:** Please refer to [K8s Cluster API docs](https://volterra.io/docs/api/k8s-cluster) to learn more

Example Usage
-------------

```hcl
resource "volterra_k8s_cluster" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "use_default_cluster_role_bindings use_custom_cluster_role_bindings" must be set
  use_default_cluster_role_bindings = true

  // One of the arguments from this list "use_default_cluster_roles use_custom_cluster_role_list" must be set
  use_default_cluster_roles = true

  // One of the arguments from this list "no_global_access global_access_enable" must be set
  no_global_access = true

  // One of the arguments from this list "no_insecure_registries insecure_registry_list" must be set
  no_insecure_registries = true

  // One of the arguments from this list "no_local_access local_access_config" must be set
  no_local_access = true

  // One of the arguments from this list "use_custom_psp_list use_default_psp" must be set
  use_default_psp = true
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

`use_custom_cluster_role_bindings` - (Optional) Select K8s cluster role bindings for this K8s cluster. See [Use Custom Cluster Role Bindings ](#use-custom-cluster-role-bindings) below for details.

`use_default_cluster_role_bindings` - (Optional) Select default K8s cluster role bindings for this K8s cluster (bool).

`use_custom_cluster_role_list` - (Optional) Select custom K8s cluster roles for this K8s cluster. See [Use Custom Cluster Role List ](#use-custom-cluster-role-list) below for details.

`use_default_cluster_roles` - (Optional) Select default K8s cluster roles for this K8s cluster (bool).

`global_access_enable` - (Optional) Access via VoltConsole to site K8s API server is enabled (bool).

`no_global_access` - (Optional) Access via VoltConsole to site K8s API server is not enabled (bool).

`insecure_registry_list` - (Optional) Select Docker insecure registries for this K8s cluster. See [Insecure Registry List ](#insecure-registry-list) below for details.

`no_insecure_registries` - (Optional) There are no Docker insecure registries to be configured (bool).

`local_access_config` - (Optional) Local access to site K8s cluster is enabled. See [Local Access Config ](#local-access-config) below for details.

`no_local_access` - (Optional) Local access to site K8s cluster is not enabled (bool).

`use_custom_psp_list` - (Optional) Select custom pod security policies for this K8s cluster. See [Use Custom Psp List ](#use-custom-psp-list) below for details.

`use_default_psp` - (Optional) Select default pod security policies for this K8s cluster (bool).

### Default Port

Use default port for K8s API server..

### Insecure Registry List

Select Docker insecure registries for this K8s cluster.

`insecure_registries` - (Required) List of docker insecure registries in format "example.com:5000" (`String`).

### Local Access Config

Local access to site K8s cluster is enabled.

`local_domain` - (Required) Local K8s API server will be accessible at <site name>.<local domain>. (`String`).

`default_port` - (Optional) Use default port for K8s API server. (bool).

`port` - (Optional) Available port range is less than 65000 except 6443. (`Int`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Use Custom Cluster Role Bindings

Select K8s cluster role bindings for this K8s cluster.

`cluster_role_bindings` - (Required) List of active cluster role binding list for a K8s cluster. See [ref](#ref) below for details.

### Use Custom Cluster Role List

Select custom K8s cluster roles for this K8s cluster.

`cluster_roles` - (Required) List of active cluster role list for a K8s cluster. See [ref](#ref) below for details.

### Use Custom Psp List

Select custom pod security policies for this K8s cluster.

`pod_security_policies` - (Required) List of active Pod security policies for a K8s cluster. See [ref](#ref) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured k8s_cluster.
