---

page_title: "Volterra: k8s_cluster_role_binding"
description: "The k8s_cluster_role_binding allows CRUD of K8s Cluster Role Binding resource on Volterra SaaS"

---

Resource volterra_k8s_cluster_role_binding
==========================================

The K8s Cluster Role Binding allows CRUD of K8s Cluster Role Binding resource on Volterra SaaS

~> **Note:** Please refer to [K8s Cluster Role Binding API docs](https://docs.cloud.f5.com/docs-v2/api/k8s-cluster-role-binding) to learn more

Example Usage
-------------

```hcl
resource "volterra_k8s_cluster_role_binding" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  k8s_cluster_role {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }

  subjects {
    // One of the arguments from this list "group service_account user" must be set

    user = "user1@example.com"
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

`k8s_cluster_role` - (Required) K8s Cluster Role for which bindings are defined.. See [ref](#ref) below for details.

`subjects` - (Required) List of subjects (user, group or service account) to which this role is bound. See [Subjects ](#subjects) below for details.

### Subjects

List of subjects (user, group or service account) to which this role is bound.

###### One of the arguments from this list "group, service_account, user" must be set

`group` - (Optional) Group id of the user group (`String`).

`service_account` - (Optional) Name and Namespace of the service account. See [Subject Choice Service Account ](#subject-choice-service-account) below for details.

`user` - (Optional) User id of the user (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Subject Choice Service Account

Name and Namespace of the service account.

`name` - (Required) Name of the service account (`String`).

`namespace` - (Required) Namespace of the service account (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured k8s_cluster_role_binding.
