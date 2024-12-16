---

page_title: "Volterra: volterra_service_credential" 

description: "The volterra_service_credential allows creation of service_credential object on Volterra SaaS"

---

Resource volterra_service_credential
====================================

================================

Credentials created via volterra_service_credential will have a separte service user associated and RBAC can be managed differently per namespace than that of the original user creating it.

For each these credentials, user can request credential in 3 different formats as required by their usecase:

API certificate - a password protected P12 certificate bundle virtual k8s kubeconfig - a kubeconfig to access a cluster with embedded user access with client certificate. API token - an easy to use secret that can be send part of HTTP request header Authorization: APIToken

~> **Note:** Please refer to [API Credential docs](https://docs.cloud.f5.com/docs/api/api-credential) to learn more

Example Usage
-------------

---

```hcl
resource "volterra_service_credential" "example" {
  created_at = timestamp()
  name      = "service-cred-example"
  service_credential_type = "SITE_GLOBAL_KUBE_CONFIG"
  service_site_kubeconfig {
    site = "test-site"
  }
}

```

Argument Reference
------------------

---

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

`created_at` - (Required) This field has been introduced to deal with renewing an expired credential. Valid values are `timestamp()` or a UTC timestamp string in RFC 3339 format. (`String`).

`namespace_roles` - (Optional) List of roles per namespace to be assigned to service credentials. See [Namespace Roles ](#namespace-roles) below for details.

`service_credential_type` - (Required) Types of API credential given. Valid values are `SERVICE_API_TOKEN`, `SERVICE_API_CERTIFICATE`, `SERVICE_KUBE_CONFIG`, `SITE_GLOBAL_KUBE_CONFIG`. (`String`).

`service_api_certificate` - (Optional) Applicable when `service_credential_type` is `SERVICE_API_CERTIFICATE`. See [Service API Certificate](#service-api-certificate) below for details.

`site_kubeconfig` - (Optional) Applicable when `service_credential_type` is `SITE_GLOBAL_KUBE_CONFIG`. See [Site Kube Config](#site-cube-config) below for details.

`vk8s_kubeconfig` - (Optional) Applicable when `service_credential_type` is `SERVICE_KUBE_CONFIG`. See [vk8s Kube Config](#vk8s-cube-config) below for details.

`expiration_days` - (Optional) Number of days after which the certificate will be expired. (`Int`).

`user_group_names` - (Optional) List of user_groups assigned to this service credential. (`List of Strings`).

`automatic_approval_api_token` - (Optional) A value of true will renew an expired user's credential. Renewal is only supported for the credential type `SERVICE_API_TOKEN`.(`Bool`).

### Namespace Roles

`namespace` - (Required) Namespace the role applies to. '\*' value implies all namespaces. (`String`).

`role` - (Required) Users role for this namespace. (`String`).

### Service API Certificate

`password` - (Required) Password is used for generating an API certificate P12 bundle user can use to protect access to it. this password will not be saved/persisted anywhere in the system. Applicable when `service_credential_type` is `SERVICE_API_CERTIFICATE`. Users have to use this password when they use the certificate, e.g. in curl or while adding to key chain. (`String`).

### Site Kube Config

`site` - (Required) Name of the site for which kubeconfig is being requested. Applicable when `service_credential_type` is `SERVICE_API_CERTIFICATE`. (`String`).

### vk8s Kube Config

`virtual_k8s_namespace` - (Required) Namespace of virtual_k8s cluster. Applicable when `api_credential_type` is `KUBE_CONFIG`. (`String`).

`virtual_k8s_name` - (Required) Name of virtual_k8s cluster. Applicable when `api_credential_type` is `KUBE_CONFIG`. (`String`).

Attribute Reference
-------------------

---

-	`id` - This is the id of service credential.

-	`data` - Data is the response format based on the Service credential type. In case of API_CERTIFICATES, the response is the base64 encoded value of certificate in PKCS12 format. In case of KUBE_CONFIG, the response is the base64 encoded value of the k8s kubeconfig file with contents as requested - cluster,namespace and base64 encoded certificate, key and CA (`string`).
