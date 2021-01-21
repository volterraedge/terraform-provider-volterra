---

page_title: "Volterra: network_connector"

description: "The network_connector allows CRUD of Network Connector resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_network_connector
===================================

The Network Connector allows CRUD of Network Connector resource on Volterra SaaS

~> **Note:** Please refer to [Network Connector API docs](https://volterra.io/docs/api/network-connector) to learn more

Example Usage
-------------

```hcl
resource "volterra_network_connector" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "sli_to_global_dr sli_to_global_snat slo_to_global_dr slo_to_global_snat sli_to_slo_snat sli_to_slo_dr" must be set
  sli_to_slo_dr = true

  // One of the arguments from this list "disable_forward_proxy enable_forward_proxy" must be set
  disable_forward_proxy = true
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

`sli_to_global_dr` - (Optional) Site local inside is connected directly to a given global network. See [Sli To Global Dr ](#sli-to-global-dr) below for details.

`sli_to_global_snat` - (Optional) Site local inside is connected to a given global network, using SNAT. See [Sli To Global Snat ](#sli-to-global-snat) below for details.

`sli_to_slo_dr` - (Optional) Site local inside is connected directly to site local outside (bool).

`sli_to_slo_snat` - (Optional) Site local inside is connected to site local outside, using SNAT. See [Sli To Slo Snat ](#sli-to-slo-snat) below for details.

`slo_to_global_dr` - (Optional) Site local outside is connected directly to a given global network. See [Slo To Global Dr ](#slo-to-global-dr) below for details.

`slo_to_global_snat` - (Optional) Site local outside is connected directly to a given global network. See [Slo To Global Snat ](#slo-to-global-snat) below for details.

`disable_forward_proxy` - (Optional) Forward Proxy is disabled for this connector (bool).

`enable_forward_proxy` - (Optional) Forward Proxy is enabled for this connector. See [Enable Forward Proxy ](#enable-forward-proxy) below for details.

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

Clear Secret is used for the secrets that are not encrypted .

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Custom Certificate

Certificates for generating intermediate certificate for TLS interception..

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Default Gw Snat

Default route in inside network to SNATed network.

### Disable Interception

Disable Interception.

### Domain Match

Domain value or regular expression to match.

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Dynamic Routing

Routes are imported in inside network from SNATed network.

### Enable For All Domains

Enable interception for all domains.

### Enable Forward Proxy

Forward Proxy is enabled for this connector.

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2000 (2 seconds) (`Int`).

`max_connect_attempts` - (Optional) Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1. (`Int`).

`no_interception` - (Optional) No TLS interception is enabled for this network connector (bool).

`tls_intercept` - (Optional) Specify TLS interception configuration for the network connector. See [Tls Intercept ](#tls-intercept) below for details.

`white_listed_ports` - (Optional) Example "tmate" server port (`Int`).

`white_listed_prefixes` - (Optional) Example "tmate" server ip (`String`).

### Enable Interception

Enable Interception.

### Interception Rules

List of ordered rules to enable or disable for TLS interception.

`domain_match` - (Required) Domain value or regular expression to match. See [Domain Match ](#domain-match) below for details.

`disable_interception` - (Optional) Disable Interception (bool).

`enable_interception` - (Optional) Enable Interception (bool).

### Interface Ip

Interface IP of the outgoing interface will be used for SNAT.

### No Interception

No TLS interception is enabled for this network connector.

### Policy

Policy to enable/disable specific domains, with implicit enable all domains.

`interception_rules` - (Required) List of ordered rules to enable or disable for TLS interception. See [Interception Rules ](#interception-rules) below for details.

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Sli To Global Dr

Site local inside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Sli To Global Snat

Site local inside is connected to a given global network, using SNAT.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

`snat_config` - (Optional) SNAT configuration to connect to global network. See [Snat Config ](#snat-config) below for details.

### Sli To Slo Snat

Site local inside is connected to site local outside, using SNAT.

`interface_ip` - (Optional) Interface IP of the outgoing interface will be used for SNAT (bool).

`snat_pool` - (Optional) IP from the ip pool prefix will be used for SNAT (`String`).

`snat_pool_allocator` - (Optional) IP from the ip pool prefix will be used for SNAT. See [ref](#ref) below for details.

`default_gw_snat` - (Optional) Default route in inside network to SNATed network (bool).

`dynamic_routing` - (Optional) Routes are imported in inside network from SNATed network (bool).

### Slo To Global Dr

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Slo To Global Snat

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

`snat_config` - (Optional) SNAT configuration to connect to global network. See [Snat Config ](#snat-config) below for details.

### Snat Config

SNAT configuration to connect to global network.

`interface_ip` - (Optional) Interface IP of the outgoing interface will be used for SNAT (bool).

`snat_pool` - (Optional) IP from the ip pool prefix will be used for SNAT (`String`).

`snat_pool_allocator` - (Optional) IP from the ip pool prefix will be used for SNAT. See [ref](#ref) below for details.

`default_gw_snat` - (Optional) Default route in inside network to SNATed network (bool).

`dynamic_routing` - (Optional) Routes are imported in inside network from SNATed network (bool).

### Tls Intercept

Specify TLS interception configuration for the network connector.

`enable_for_all_domains` - (Optional) Enable interception for all domains (bool).

`policy` - (Optional) Policy to enable/disable specific domains, with implicit enable all domains. See [Policy ](#policy) below for details.

`custom_certificate` - (Optional) Certificates for generating intermediate certificate for TLS interception.. See [Custom Certificate ](#custom-certificate) below for details.

`volterra_certificate` - (Optional) Volterra certificates for generating intermediate certificate for TLS interception. (bool).

`trusted_ca_url` - (Optional) Custom trusted CA certificates for validating upstream server certificate (`String`).

`volterra_trusted_ca` - (Optional) Default volterra trusted CA list for validating upstream server certificate (bool).

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Volterra Certificate

Volterra certificates for generating intermediate certificate for TLS interception..

### Volterra Trusted Ca

Default volterra trusted CA list for validating upstream server certificate.

### Wingman Secret Info

Secret is given as bootstrap secret in Volterra Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_connector.
