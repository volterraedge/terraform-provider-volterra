---

page_title: "Volterra: volterra_service_credential"

description: "The volterra_service_credential allows creation of service_credential object on Volterra SaaS"

---

Resource volterra_service_credential
====================================

================================

Credentials created via volterra_service_credential will have a separte service user associated and RBAC can be managed differently per namespace than that of the original user creating it.

For each these credentials, user can request credential in different formats as required by their usecase:

API certificate - a password protected P12 certificate bundle. Virtual k8s kubeconfig - a kubeconfig to access a cluster with embedded user access with client certificate. Site kubeconfig - a kubeconfig to access a site. API token - an easy to use secret that can be send part of HTTP request header Authorization: APIToken.

~> **Note:** Please refer to [API Credential docs](https://docs.cloud.f5.com/docs/api/api-credential) to learn more

Example Usage
-------------

---

```hcl
resource "volterra_service_credential" "api_token" {
  created_at = timestamp()
  name       = "service-api-token-example"

  service_api_token = true

  namespace_roles {
    namespace = "*"
    role      = "ves-io-admin-role"
  }

  user_group_names = ["service-users"]

  expiration_days              = 10
  automatic_approval_api_token = true
}
```

```hcl
resource "volterra_service_credential" "api_certificate" {
  created_at = timestamp()
  name       = "service-api-cert-example"

  service_api_certificate {
    password = "ChangeMe123!"
  }

  expiration_days = 10
}
```

Argument Reference
------------------

---

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

`created_at` - (Required) This field has been introduced to deal with renewing an expired credential. Valid values are `timestamp()` or a UTC timestamp string in RFC 3339 format. (`String`).

`namespace_roles` - (Optional) List of roles per namespace to be assigned to service credentials. See [Namespace Roles ](#namespace-roles) below for details.

`service_api_token` - (Optional) Selects service API token credential generation. (`Bool`).

`service_api_certificate` - (Optional) Selects service API certificate credential generation. See [Service API Certificate](#service-api-certificate) below for details.

`site_kubeconfig` - (Optional) Selects site kubeconfig credential generation. See [Site Kube Config](#site-kube-config) below for details.

`vk8s_kubeconfig` - (Optional) Selects virtual k8s kubeconfig credential generation. See [vk8s Kube Config](#vk8s-kube-config) below for details.

`expiration_days` - (Optional) Number of days after which the certificate will be expired. (`Int`).

`user_group_names` - (Optional) List of user_groups assigned to this service credential. (`List of Strings`).

`automatic_approval_api_token` - (Optional) A value of true will renew an expired user's credential. Renewal is only supported for the credential type `SERVICE_API_TOKEN`.(`Bool`).

Only one of `service_api_token`, `service_api_certificate`, `site_kubeconfig`, or `vk8s_kubeconfig` should be configured.

### Namespace Roles

`namespace` - (Required) Namespace the role applies to. '\*' value implies all namespaces. (`String`).

`role` - (Required) Users role for this namespace. (`String`).

### Service API Certificate

`password` - (Required) Password is used for generating an API certificate P12 bundle user can use to protect access to it. this password will not be saved/persisted anywhere in the system. Users have to use this password when they use the certificate, e.g. in curl or while adding to key chain. (`String`).

### Site Kube Config

`site` - (Required) Name of the site for which kubeconfig is being requested. (`String`).

### vk8s Kube Config

`vk8s_namespace` - (Optional) Namespace of virtual k8s cluster. (`String`).

`vk8s_cluster_name` - (Optional) Name of virtual k8s cluster. (`String`).

Attribute Reference
-------------------

---

-	`id` - This is the id of service credential.

-	`data` - Data is the response format based on the Service credential type. In case of API_CERTIFICATES, the response is the base64 encoded value of certificate in PKCS12 format. In case of KUBE_CONFIG, the response is the base64 encoded value of the k8s kubeconfig file with contents as requested - cluster,namespace and base64 encoded certificate, key and CA (`string`).
