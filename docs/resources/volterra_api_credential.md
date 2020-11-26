---

page_title: "Volterra: volterra_api_credential"

description: "The volterra_api_credential allows creation of api_credential object on Volterra SaaS"
----------------------------------------------------------------------------------------------------

Resource volterra_api_credential
================================

================================

Credentials created via volterra_api_credential inherits same user that of the creator and it gets same RBAC applied to the user. So when this credential is used in APIs, it will have same response/audit as creator of the credential is.

For each these credentials, user can request credential in 3 different formats as required by their usecase:

API certificate - a password protected P12 certificate bundle virtual k8s kubeconfig - a kubeconfig to access a cluster with embedded user access with client certificate. API token - an easy to use secret that can be send part of HTTP request header Authorization: APIToken

~> **Note:** Please refer to [API Credential docs](https://volterra.io/docs/api/api-credential) to learn more

Example Usage
-------------

---

```hcl
resource "volterra_api_credential" "example" {
  name      = "api-cred-example"
  api_credential_type = "KUBE_CONFIG"
  virtual_k8s_namespace = "ns1"
  virtual_k8s_name = "vk8s1"
}

```

Argument Reference
------------------

---

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

`api_credential_type` - (Required) Types of API credential given. Valid values are `API_CERTIFICATE`, `KUBE_CONFIG`, `API_TOKEN` (`String`).

`virtual_k8s_namespace` - (Optional) Namespace of virtual_k8s cluster. Applicable when `api_credential_type` is `KUBE_CONFIG` (`String`).

`virtual_k8s_name` - (Optional) Name of virtual_k8s cluster. Applicable when `api_credential_type` is `KUBE_CONFIG` (`String`).

`api_credential_password` - (Optional) Password is used for generating an API certificate P12 bundle user can use to protect access to it. this password will not be saved/persisted anywhere in the system. Applicable when `api_credential_type` is `API_CERTIFICATE`. Users have to use this password when they use the certificate, e.g. in curl or while adding to key chain (`String`).

`expiry_days` - (Optional) Number of days after which the certificate will be expired (`Int`).

Attribute Reference
-------------------

---

-	`id` - This is the id of api credential.

-	`data` - Data is the response format based on the API credential type. In case of API_CERTIFICATES, the response is the base64 encoded value of certificate in PKCS12 format. In case of KUBE_CONFIG, the response is the base64 encoded value of the k8s kubeconfig file with contents as requested - cluster,namespace and base64 encoded certificate, key and CA (`string`).
