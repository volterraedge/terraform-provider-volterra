---

page_title: "Volterra: discovery"

description: "The discovery allows CRUD of Discovery resource on Volterra SaaS"
-------------------------------------------------------------------------------

Resource volterra_discovery
===========================

The Discovery allows CRUD of Discovery resource on Volterra SaaS

~> **Note:** Please refer to [Discovery API docs](https://volterra.io/docs/api/discovery) to learn more

Example Usage
-------------

```hcl
resource "volterra_discovery" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "discovery_k8s discovery_consul" must be set

  discovery_k8s {
    access_info {
      // One of the arguments from this list "connection_info in_cluster kubeconfig_url" must be set

      kubeconfig_url {
        secret_encoding_type = "secret_encoding_type"

        // One of the arguments from this list "clear_secret_info wingman_secret_info blindfold_secret_info vault_secret_info" must be set

        blindfold_secret_info {
          decryption_provider = "decryption_provider"
          location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
          store_provider      = "store_provider"
        }
      }

      // One of the arguments from this list "isolated reachable" must be set
      isolated = true
    }

    publish_info {
      // One of the arguments from this list "disable publish publish_fqdns dns_delegation" must be set
      disable = true
    }
  }
  where {
    // One of the arguments from this list "virtual_network site virtual_site" must be set

    virtual_network {
      ref {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }
    }
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

`discovery_consul` - (Optional) Discovery configuration for Hashicorp Consul. See [Discovery Consul ](#discovery-consul) below for details.

`discovery_k8s` - (Optional) Discovery configuration for K8s.. See [Discovery K8s ](#discovery-k8s) below for details.

`where` - (Required) All the sites where this discovery config is valid.. See [Where ](#where) below for details.

### Access Info

Credentials to access Hashicorp Consul service discovery.

`connection_info` - (Optional) Configuration details to access Hashicorp Consul API service using REST.. See [Connection Info ](#connection-info) below for details.

`http_basic_auth_info` - (Optional) Username and password used for HTTP/HTTPS access. See [Http Basic Auth Info ](#http-basic-auth-info) below for details.

`scheme` - (Optional) scheme (`String`).

### Ca Certificate Url

Volterra Secret. URL to fetch the server CA certificate file.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Certificate Url

Volterra Secret. URL to fetch the client certificate file.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Connection Info

Configuration details to access Hashicorp Consul API service using REST..

`api_server` - (Required) API server must be a fully qualified domain string and port specified as host:port pair (`String`).

`tls_info` - (Optional) TLS settings to enable transport layer security. See [Tls Info ](#tls-info) below for details.

### Discovery Consul

Discovery configuration for Hashicorp Consul.

`access_info` - (Required) Credentials to access Hashicorp Consul service discovery. See [Access Info ](#access-info) below for details.

`publish_info` - (Required) Configuration to publish VIPs. See [Publish Info ](#publish-info) below for details.

### Discovery K8s

Discovery configuration for K8s..

`access_info` - (Required) Credentials can be kubeconfig file or mutual TLS using PKI certificates. See [Access Info ](#access-info) below for details.

`publish_info` - (Required) Configuration to publish VIPs. See [Publish Info ](#publish-info) below for details.

### Http Basic Auth Info

Username and password used for HTTP/HTTPS access.

`passwd_url` - (Optional) Volterra Secret. URL for password, needs to be fetched from this path. See [Passwd Url ](#passwd-url) below for details.

`user_name` - (Optional) username in consul (`String`).

### Key Url

The data may be optionally secured using BlindFold..

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Passwd Url

Volterra Secret. URL for password, needs to be fetched from this path.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Publish Info

Configuration to publish VIPs.

`disable` - (Optional) Disable VIP Publishing (bool).

`publish` - (Optional) Publish domain to VIP mapping. (bool).

### Tls Info

TLS settings to enable transport layer security.

`ca_certificate_url` - (Optional) Volterra Secret. URL to fetch the server CA certificate file. See [Ca Certificate Url ](#ca-certificate-url) below for details.

`certificate` - (Optional) Client certificate is PEM-encoded certificate or certificate-chain. (`String`).

`certificate_url` - (Optional) Volterra Secret. URL to fetch the client certificate file. See [Certificate Url ](#certificate-url) below for details.

`key_url` - (Optional) The data may be optionally secured using BlindFold.. See [Key Url ](#key-url) below for details.

`server_name` - (Optional) server is used (`String`).

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

### Where

All the sites where this discovery config is valid..

`site` - (Optional) Direct reference to site object. See [Site ](#site) below for details.

`virtual_network` - (Optional) Direct reference to virtual network object. See [Virtual Network ](#virtual-network) below for details.

`virtual_site` - (Optional) Direct reference to virtual site object. See [Virtual Site ](#virtual-site) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured discovery.
