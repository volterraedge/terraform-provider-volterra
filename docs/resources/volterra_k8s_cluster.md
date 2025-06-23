---

page_title: "Volterra: k8s_cluster"

description: "The k8s_cluster allows CRUD of K8s Cluster resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_k8s_cluster
=============================

The K8s Cluster allows CRUD of K8s Cluster resource on Volterra SaaS

~> **Note:** Please refer to [K8s Cluster API docs](https://docs.cloud.f5.com/docs-v2/api/k8s-cluster) to learn more

Example Usage
-------------

```hcl
resource "volterra_k8s_cluster" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "cluster_wide_app_list no_cluster_wide_apps" must be set

  no_cluster_wide_apps = true

  // One of the arguments from this list "use_custom_cluster_role_bindings use_default_cluster_role_bindings" must be set

  use_default_cluster_role_bindings = true

  // One of the arguments from this list "use_custom_cluster_role_list use_default_cluster_roles" must be set

  use_default_cluster_roles = true

  // One of the arguments from this list "cluster_scoped_access_deny cluster_scoped_access_permit" must be set

  cluster_scoped_access_deny = true

  // One of the arguments from this list "global_access_enable no_global_access" must be set

  global_access_enable = true

  // One of the arguments from this list "insecure_registry_list no_insecure_registries" must be set

  no_insecure_registries = true

  // One of the arguments from this list "local_access_config no_local_access" must be set

  no_local_access = true

  // One of the arguments from this list "use_custom_pod_security_admission use_default_pod_security_admission" must be set

  use_default_pod_security_admission = true

  // One of the arguments from this list "use_custom_psp_list use_default_psp" must be set

  use_default_psp = true

  // One of the arguments from this list "vk8s_namespace_access_deny vk8s_namespace_access_permit" must be set

  vk8s_namespace_access_deny = true
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

###### One of the arguments from this list "cluster_wide_app_list, no_cluster_wide_apps" must be set

`cluster_wide_app_list` - (Optional) Select cluster wide applications to be deployed. See [Apps Choice Cluster Wide App List ](#apps-choice-cluster-wide-app-list) below for details.

`no_cluster_wide_apps` - (Optional) There are no cluster wide applications to be deployed (`Bool`).

###### One of the arguments from this list "use_custom_cluster_role_bindings, use_default_cluster_role_bindings" must be set

`use_custom_cluster_role_bindings` - (Optional) Select custom K8s cluster role bindings for this K8s cluster. See [Cluster Role Bindings Choice Use Custom Cluster Role Bindings ](#cluster-role-bindings-choice-use-custom-cluster-role-bindings) below for details.

`use_default_cluster_role_bindings` - (Optional) Select default K8s cluster role bindings for this K8s cluster (`Bool`).

###### One of the arguments from this list "use_custom_cluster_role_list, use_default_cluster_roles" must be set

`use_custom_cluster_role_list` - (Optional) Select custom K8s cluster roles for this K8s cluster. See [Cluster Role Choice Use Custom Cluster Role List ](#cluster-role-choice-use-custom-cluster-role-list) below for details.

`use_default_cluster_roles` - (Optional) Select default K8s cluster roles for this K8s cluster (`Bool`).

###### One of the arguments from this list "cluster_scoped_access_deny, cluster_scoped_access_permit" must be set

`cluster_scoped_access_deny` - (Optional) Access to Create, Patch, Replace, Update and Delete for ClusterRoles, ClusterRoleBindings, MutatingWebhookConfiguration and ValidatingWebhookConfiguration will not be allowed through K8s cluster API. It can be managed only through VoltConsole. (`Bool`).

`cluster_scoped_access_permit` - (Optional) Access to Create, Patch, Replace, Update and Delete for ClusterRoles, ClusterRoleBindings, MutatingWebhookConfiguration and ValidatingWebhookConfiguration will be allowed through K8s cluster API. This allows native k8s API operation with ClusterRoles and ClusterRoleBindings. (`Bool`).

###### One of the arguments from this list "global_access_enable, no_global_access" must be set

`global_access_enable` - (Optional) Access via VoltConsole to site K8s API server is enabled (`Bool`).

`no_global_access` - (Optional) Access via VoltConsole to site K8s API server is not enabled (`Bool`).

###### One of the arguments from this list "insecure_registry_list, no_insecure_registries" must be set

`insecure_registry_list` - (Optional) Select Docker insecure registries for this K8s cluster. See [Insecure Registries Choice Insecure Registry List ](#insecure-registries-choice-insecure-registry-list) below for details.

`no_insecure_registries` - (Optional) There are no Docker insecure registries to be configured (`Bool`).

###### One of the arguments from this list "local_access_config, no_local_access" must be set

`local_access_config` - (Optional) Local access to site K8s cluster is enabled. See [Local Access Choice Local Access Config ](#local-access-choice-local-access-config) below for details.

`no_local_access` - (Optional) Local access to site K8s cluster is not enabled (`Bool`).

###### One of the arguments from this list "use_custom_pod_security_admission, use_default_pod_security_admission" must be set

`use_custom_pod_security_admission` - (Optional) Select Custom Pod Security Admission. See [ref](#ref) below for details.

`use_default_pod_security_admission` - (Optional) Select Default Pod Security Admission (`Bool`).

###### One of the arguments from this list "use_custom_psp_list, use_default_psp" must be set

`use_custom_psp_list` - (Optional) Select custom pod security policies for this K8s cluster. See [Pod Security Policy Choice Use Custom Psp List ](#pod-security-policy-choice-use-custom-psp-list) below for details.

`use_default_psp` - (Optional) Select default pod security policies for this K8s cluster (`Bool`).

###### One of the arguments from this list "vk8s_namespace_access_deny, vk8s_namespace_access_permit" must be set

`vk8s_namespace_access_deny` - (Optional) Access to create, modify and delete resources in VK8s namespaces will be prevented for service accounts and Managed K8s clients. Resources in VK8s namespaces can be managed only through VK8s API or UI. (`Bool`).

`vk8s_namespace_access_permit` - (Optional) Access to create, modify and delete resources in VK8s namespaces will be allowed for service accounts and Managed K8s clients. (`Bool`).

### App Choice Argo Cd

Deploy Argo Continuous Deployment(CD) application.

`local_domain` - (Optional) Local domain to access argocd for example argocd.localdomain. See [Argo Cd Local Domain ](#argo-cd-local-domain) below for details.

### App Choice Dashboard

Deploy Kubernetes Dashboard application.

### App Choice Metrics Server

Deploy Kubernetes Metrics Server application.

### App Choice Prometheus

Prometheus access via k8s api server endpoint.

### Apps Choice Cluster Wide App List

Select cluster wide applications to be deployed.

`cluster_wide_apps` - (Required) List of cluster wide applications. See [Cluster Wide App List Cluster Wide Apps ](#cluster-wide-app-list-cluster-wide-apps) below for details.

### Argo Cd Local Domain

Local domain to access argocd for example argocd.localdomain.

`local_domain` - (Required) ArgoCD will be accessible at <site name>.<local domain>. (`String`).

`password` - (Required) Select blindfold or clear text password for ArgoCD admin.. See [Local Domain Password ](#local-domain-password) below for details.

###### One of the arguments from this list "default_port, port" must be set

`default_port` - (Optional) Use default port 443 for ArgoCD server. (`Bool`).

`port` - (Optional) Available port range is less than 65000 except reserved ports. (`Int`).

### Cluster Role Bindings Choice Use Custom Cluster Role Bindings

Select custom K8s cluster role bindings for this K8s cluster.

`cluster_role_bindings` - (Required) List of active cluster role binding list for a K8s cluster. See [ref](#ref) below for details.

### Cluster Role Choice Use Custom Cluster Role List

Select custom K8s cluster roles for this K8s cluster.

`cluster_roles` - (Required) List of active cluster role list for a K8s cluster. See [ref](#ref) below for details.

### Cluster Wide App List Cluster Wide Apps

List of cluster wide applications.

###### One of the arguments from this list "argo_cd, dashboard, metrics_server, prometheus" must be set

`argo_cd` - (Optional) Deploy Argo Continuous Deployment(CD) application. See [App Choice Argo Cd ](#app-choice-argo-cd) below for details.

`dashboard` - (Optional) Deploy Kubernetes Dashboard application. See [App Choice Dashboard ](#app-choice-dashboard) below for details.

`metrics_server` - (Optional) Deploy Kubernetes Metrics Server application. See [App Choice Metrics Server ](#app-choice-metrics-server) below for details.

`prometheus` - (Optional) Prometheus access via k8s api server endpoint. See [App Choice Prometheus ](#app-choice-prometheus) below for details.

### Insecure Registries Choice Insecure Registry List

Select Docker insecure registries for this K8s cluster.

`insecure_registries` - (Required) List of docker insecure registries in format "example.com:5000" (`String`).

### Local Access Choice Local Access Config

Local access to site K8s cluster is enabled.

`local_domain` - (Required) Local K8s API server will be accessible at <site name>.<local domain>. (`String`).

###### One of the arguments from this list "default_port, port" must be set

`default_port` - (Optional) Use default port 65443 for K8s API server. (`Bool`).

`port` - (Optional) Available port range is less than 65000 except reserved ports. (`Int`).

### Local Domain Password

Select blindfold or clear text password for ArgoCD admin..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Pod Security Policy Choice Use Custom Psp List

Select custom pod security policies for this K8s cluster.

`pod_security_policies` - (Required) List of active Pod security policies for a K8s cluster. See [ref](#ref) below for details.

### Port Choice Default Port

Use default port 443 for ArgoCD server..

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Secret Info Oneof Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Info Oneof Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured k8s_cluster.
