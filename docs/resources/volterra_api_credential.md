---

page_title: "Volterra: volterra_api_credential"

description: "The volterra_api_credential allows creation of api_credential object on Volterra SaaS"

---

Resource volterra_api_credential
================================

================================

Credentials created via volterra_api_credential inherits same user that of the creator and it gets same RBAC applied to the user. So when this credential is used in APIs, it will have same response/audit as creator of the credential is.

For each these credentials, user can request credential in different formats as required by their usecase:

API certificate - a password protected P12 certificate bundle. API token - an easy to use secret that can be send part of HTTP request header Authorization: APIToken. Legacy virtual k8s kubeconfig is supported through deprecated arguments for backward compatibility.

~> **Note:** Please refer to [API Credential docs](https://docs.cloud.f5.com/docs/api/api-credential) to learn more

Example Usage
-------------

---

```hcl
resource "volterra_api_credential" "api_token" {
  created_at = timestamp()
  name       = "api-token-example"

  api_token = true

  expiry_days                  = 10
  automatic_approval_api_token = true
}
```

```hcl
resource "volterra_api_credential" "api_certificate" {
  created_at = timestamp()
  name       = "api-cert-example"

  api_certificate {
    password = "ChangeMe123!"
  }

  expiry_days = 10
}
```

Argument Reference
------------------

---

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

`created_at` - (Required) This field has been introduced to deal with renewing an expired credential. Valid values are `timestamp()` or a UTC timestamp string in RFC 3339 format. (`String`).

`api_token` - (Optional) Selects API token credential generation. Conflicts with `api_certificate`. (`Bool`).

`api_certificate` - (Optional) Selects API certificate credential generation. Conflicts with `api_token`. See [API Certificate](#api-certificate) below for details.

`api_credential_type` - (Optional, Deprecated) Legacy type selector for API credential generation. Valid values are `API_CERTIFICATE`, `KUBE_CONFIG`, `API_TOKEN` (`String`). Use `api_token` or `api_certificate` for new configurations.

`virtual_k8s_namespace` - (Optional, Deprecated) Namespace of virtual_k8s cluster. Applicable when `api_credential_type` is `KUBE_CONFIG` (`String`).

`virtual_k8s_name` - (Optional, Deprecated) Name of virtual_k8s cluster. Applicable when `api_credential_type` is `KUBE_CONFIG` (`String`).

`api_credential_password` - (Optional, Deprecated) Password is used for generating an API certificate P12 bundle user can use to protect access to it. this password will not be saved/persisted anywhere in the system. Applicable when `api_credential_type` is `API_CERTIFICATE`. Use `password` inside `api_certificate` for new configurations. (`String`).

`expiry_days` - (Optional) Number of days after which the certificate will be expired (`Int`).

`automatic_approval_api_token` - (Optional) A value of true will renew an expired user's credential. Renewal is only supported for the credential type API_TOKEN.(`Bool`\)

### API Certificate

`password` - (Required) Password is used for generating an API certificate P12 bundle user can use to protect access to it. this password will not be saved/persisted anywhere in the system. Users have to use this password when they use the certificate, e.g. in curl or while adding to key chain. (`String`).

Attribute Reference
-------------------

---

-   `id` - This is the id of api credential.

-   `data` - Data is the response format based on the API credential type. In case of API_CERTIFICATES, the response is the base64 encoded value of certificate in PKCS12 format. In case of KUBE_CONFIG, the response is the base64 encoded value of the k8s kubeconfig file with contents as requested - cluster,namespace and base64 encoded certificate, key and CA (`string`).

