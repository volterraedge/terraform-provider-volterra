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

  // One of the arguments from this list "no_cluster_wide_apps cluster_wide_app_list" must be set
  no_cluster_wide_apps = true

  // One of the arguments from this list "use_custom_cluster_role_bindings use_default_cluster_role_bindings" must be set
  use_default_cluster_role_bindings = true

  // One of the arguments from this list "use_default_cluster_roles use_custom_cluster_role_list" must be set
  use_default_cluster_roles = true

  // One of the arguments from this list "cluster_scoped_access_deny cluster_scoped_access_permit" must be set
  cluster_scoped_access_deny = true

  // One of the arguments from this list "no_global_access global_access_enable" must be set
  no_global_access = true

  // One of the arguments from this list "no_insecure_registries insecure_registry_list" must be set

  insecure_registry_list {
    insecure_registries = ["example.com:5000"]
  }
  // One of the arguments from this list "no_local_access local_access_config" must be set
  no_local_access = true
  // One of the arguments from this list "use_default_psp use_custom_psp_list" must be set
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

`cluster_wide_app_list` - (Optional) Select cluster wide applications to be deployed. See [Cluster Wide App List ](#cluster-wide-app-list) below for details.

`no_cluster_wide_apps` - (Optional) There are no cluster wide applications to be deployed (bool).

`use_custom_cluster_role_bindings` - (Optional) Select custom K8s cluster role bindings for this K8s cluster. See [Use Custom Cluster Role Bindings ](#use-custom-cluster-role-bindings) below for details.

`use_default_cluster_role_bindings` - (Optional) Select default K8s cluster role bindings for this K8s cluster (bool).

`use_custom_cluster_role_list` - (Optional) Select custom K8s cluster roles for this K8s cluster. See [Use Custom Cluster Role List ](#use-custom-cluster-role-list) below for details.

`use_default_cluster_roles` - (Optional) Select default K8s cluster roles for this K8s cluster (bool).

`cluster_scoped_access_deny` - (Optional) Access to Create, Patch, Replace, Update and Delete for ClusterRoles, ClusterRoleBindings, MutatingWebhookConfiguration and ValidatingWebhookConfiguration will not be allowed through K8s cluster API. It can be managed only through VoltConsole. (bool).

`cluster_scoped_access_permit` - (Optional) Access to Create, Patch, Replace, Update and Delete for ClusterRoles, ClusterRoleBindings, MutatingWebhookConfiguration and ValidatingWebhookConfiguration will be allowed through K8s cluster API. This allows native k8s API operation with ClusterRoles and ClusterRoleBindings. (bool).

`global_access_enable` - (Optional) Access via VoltConsole to site K8s API server is enabled (bool).

`no_global_access` - (Optional) Access via VoltConsole to site K8s API server is not enabled (bool).

`insecure_registry_list` - (Optional) Select Docker insecure registries for this K8s cluster. See [Insecure Registry List ](#insecure-registry-list) below for details.

`no_insecure_registries` - (Optional) There are no Docker insecure registries to be configured (bool).

`local_access_config` - (Optional) Local access to site K8s cluster is enabled. See [Local Access Config ](#local-access-config) below for details.

`no_local_access` - (Optional) Local access to site K8s cluster is not enabled (bool).

`use_custom_psp_list` - (Optional) Select custom pod security policies for this K8s cluster. See [Use Custom Psp List ](#use-custom-psp-list) below for details.

`use_default_psp` - (Optional) Select default pod security policies for this K8s cluster (bool).

### Argo Cd

Deploy Argo Continuous Deployment(CD) application.

`generated_yaml` - (Optional) Generated YAML (`String`).

`local_domain` - (Optional) Local domain to access argocd for example argocd.localdomain. See [Local Domain ](#local-domain) below for details.

### Blindfold Secret Info

Blindfold Secret is used for the secrets managed by Volterra Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Cluster Wide App List

Select cluster wide applications to be deployed.

`cluster_wide_apps` - (Required) List of cluster wide applications. See [Cluster Wide Apps ](#cluster-wide-apps) below for details.

### Cluster Wide Apps

List of cluster wide applications.

`argo_cd` - (Optional) Deploy Argo Continuous Deployment(CD) application. See [Argo Cd ](#argo-cd) below for details.

`dashboard` - (Optional) Deploy Kubernetes Dashboard application. See [Dashboard ](#dashboard) below for details.

`metrics_server` - (Optional) Deploy Kubernetes Metrics Server application. See [Metrics Server ](#metrics-server) below for details.

`prometheus` - (Optional) Prometheus access via k8s api server endpoint. See [Prometheus ](#prometheus) below for details.

### Dashboard

Deploy Kubernetes Dashboard application.

`generated_yaml` - (Optional) Generated YAML (`String`).

### Default Port

Use default port 443 for ArgoCD server..

### Insecure Registry List

Select Docker insecure registries for this K8s cluster.

`insecure_registries` - (Required) List of docker insecure registries in format "example.com:5000" (`String`).

### Local Access Config

Local access to site K8s cluster is enabled.

`local_domain` - (Required) Local K8s API server will be accessible at <site name>.<local domain>. (`String`).

`default_port` - (Optional) Use default port 65443 for K8s API server. (bool).

`port` - (Optional) Available port range is less than 65000 except reserved ports. (`Int`).

### Local Domain

Local domain to access argocd for example argocd.localdomain.

`local_domain` - (Required) ArgoCD will be accessible at <site name>.<local domain>. (`String`).

`password` - (Required) Select blindfold or clear text password for ArgoCD admin.. See [Password ](#password) below for details.

`default_port` - (Optional) Use default port 443 for ArgoCD server. (bool).

`port` - (Optional) Available port range is less than 65000 except reserved ports. (`Int`).

### Metrics Server

Deploy Kubernetes Metrics Server application.

`generated_yaml` - (Optional) Generated YAML (`String`).

### Password

Select blindfold or clear text password for ArgoCD admin..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Prometheus

Prometheus access via k8s api server endpoint.

`generated_yaml` - (Optional) Generated YAML (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Use Custom Cluster Role Bindings

Select custom K8s cluster role bindings for this K8s cluster.

`cluster_role_bindings` - (Required) List of active cluster role binding list for a K8s cluster. See [ref](#ref) below for details.

### Use Custom Cluster Role List

Select custom K8s cluster roles for this K8s cluster.

`cluster_roles` - (Required) List of active cluster role list for a K8s cluster. See [ref](#ref) below for details.

### Use Custom Psp List

Select custom pod security policies for this K8s cluster.

`pod_security_policies` - (Required) List of active Pod security policies for a K8s cluster. See [ref](#ref) below for details.

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Wingman Secret Info

Secret is given as bootstrap secret in Volterra Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured k8s_cluster.
