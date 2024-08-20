











---
page_title: "Volterra: k8s_cluster"
description: "The k8s_cluster allows CRUD of K8s Cluster  resource on Volterra SaaS"
---
# Resource volterra_k8s_cluster

The K8s Cluster  allows CRUD of K8s Cluster  resource on Volterra SaaS

~> **Note:** Please refer to [K8s Cluster  API docs](https://docs.cloud.f5.com/docs-v2/api/k8s-cluster) to learn more

## Example Usage

```hcl
resource "volterra_k8s_cluster" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "no_cluster_wide_apps cluster_wide_app_list" must be set

  no_cluster_wide_apps = true

  // One of the arguments from this list "use_default_cluster_role_bindings use_custom_cluster_role_bindings" must be set

  use_default_cluster_role_bindings = true

  // One of the arguments from this list "use_default_cluster_roles use_custom_cluster_role_list" must be set

  use_default_cluster_roles = true

  // One of the arguments from this list "cluster_scoped_access_deny cluster_scoped_access_permit" must be set

  cluster_scoped_access_deny = true

  // One of the arguments from this list "no_global_access global_access_enable" must be set

  no_global_access = true

  // One of the arguments from this list "no_insecure_registries insecure_registry_list" must be set

  no_insecure_registries = true

  // One of the arguments from this list "no_local_access local_access_config" must be set

  local_access_config {
    local_domain = "example.com"

    // One of the arguments from this list "default_port port" must be set

    default_port = true
  }

  // One of the arguments from this list "use_default_pod_security_admission use_custom_pod_security_admission" must be set

  use_default_pod_security_admission = true

  // One of the arguments from this list "use_default_psp use_custom_psp_list" must be set

  use_default_psp = true

  // One of the arguments from this list "vk8s_namespace_access_deny vk8s_namespace_access_permit" must be set

  vk8s_namespace_access_deny = true
}

```

## Argument Reference

### Metadata Argument Reference
`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).


`description` - (Optional) Human readable description for the object (`String`).


`disable` - (Optional) A value of true will administratively disable the object (`Bool`).


`labels` - (Optional) by selector expression (`String`).


`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).


`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).



### Spec Argument Reference


`cluster_wide_app_list` - (Optional) Select cluster wide applications to be deployed. See [Apps Choice Cluster Wide App List ](#apps-choice-cluster-wide-app-list) below for details.
		


		




		



		



		


		








		







		






		









		







		







		





		





		






`no_cluster_wide_apps` - (Optional) There are no cluster wide applications to be deployed (`Bool`).





`use_custom_cluster_role_bindings` - (Optional) Select custom K8s cluster role bindings for this K8s cluster. See [Cluster Role Bindings Choice Use Custom Cluster Role Bindings ](#cluster-role-bindings-choice-use-custom-cluster-role-bindings) below for details.
		




`use_default_cluster_role_bindings` - (Optional) Select default K8s cluster role bindings for this K8s cluster (`Bool`).





`use_custom_cluster_role_list` - (Optional) Select custom K8s cluster roles for this K8s cluster. See [Cluster Role Choice Use Custom Cluster Role List ](#cluster-role-choice-use-custom-cluster-role-list) below for details.
		




`use_default_cluster_roles` - (Optional) Select default K8s cluster roles for this K8s cluster (`Bool`).





`cluster_scoped_access_deny` - (Optional) Access to Create, Patch, Replace, Update and Delete for ClusterRoles, ClusterRoleBindings, MutatingWebhookConfiguration and ValidatingWebhookConfiguration will not be allowed through K8s cluster API. It can be managed only through VoltConsole. (`Bool`).


`cluster_scoped_access_permit` - (Optional) Access to Create, Patch, Replace, Update and Delete for ClusterRoles, ClusterRoleBindings, MutatingWebhookConfiguration and ValidatingWebhookConfiguration will be allowed through K8s cluster API. This allows native k8s API operation with ClusterRoles and ClusterRoleBindings. (`Bool`).





`global_access_enable` - (Optional) Access via VoltConsole to site K8s API server is enabled (`Bool`).


`no_global_access` - (Optional) Access via VoltConsole to site K8s API server is not enabled (`Bool`).





`insecure_registry_list` - (Optional) Select Docker insecure registries for this K8s cluster. See [Insecure Registries Choice Insecure Registry List ](#insecure-registries-choice-insecure-registry-list) below for details.
		




`no_insecure_registries` - (Optional) There are no Docker insecure registries to be configured (`Bool`).





`local_access_config` - (Optional) Local access to site K8s cluster is enabled. See [Local Access Choice Local Access Config ](#local-access-choice-local-access-config) below for details.
		










`no_local_access` - (Optional) Local access to site K8s cluster is not enabled (`Bool`).





`use_custom_pod_security_admission` - (Optional) Select Custom Pod Security Admission. See [ref](#ref) below for details.


`use_default_pod_security_admission` - (Optional) Select Default Pod Security Admission (`Bool`).





`use_custom_psp_list` - (Optional) Select custom pod security policies for this K8s cluster. See [Pod Security Policy Choice Use Custom Psp List ](#pod-security-policy-choice-use-custom-psp-list) below for details.
		




`use_default_psp` - (Optional) Select default pod security policies for this K8s cluster (`Bool`).





`vk8s_namespace_access_deny` - (Optional) Access to create, modify and delete resources in VK8s namespaces will be prevented for service accounts and Managed K8s clients. Resources in VK8s namespaces can be managed only through VK8s API or UI. (`Bool`).


`vk8s_namespace_access_permit` - (Optional) Access to create, modify and delete resources in VK8s namespaces will be allowed for service accounts and Managed K8s clients. (`Bool`).




### App Choice Argo Cd 

 Deploy Argo Continuous Deployment(CD) application.

`generated_yaml` - (Optional) Generated YAML (`String`).(Deprecated)

`local_domain` - (Optional) Local domain to access argocd for example argocd.localdomain. See [Argo Cd Local Domain ](#argo-cd-local-domain) below for details.



### App Choice Dashboard 

 Deploy Kubernetes Dashboard application.

`generated_yaml` - (Optional) Generated YAML (`String`).(Deprecated)



### App Choice Metrics Server 

 Deploy Kubernetes Metrics Server application.

`generated_yaml` - (Optional) Generated YAML (`String`).(Deprecated)



### App Choice Prometheus 

 Prometheus access via k8s api server endpoint.

`generated_yaml` - (Optional) Generated YAML (`String`).(Deprecated)



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



###### One of the arguments from this list "dashboard, metrics_server, prometheus, argo_cd" must be set

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

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Password Blindfold Secret Info Internal ](#password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)



###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.


`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.


`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)


`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)




### Password Blindfold Secret Info Internal 

 Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).



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



### Secret Info Oneof Vault Secret Info 

 Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).



### Secret Info Oneof Wingman Secret Info 

 Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).



## Attribute Reference

* `id` - This is the id of the configured k8s_cluster.

