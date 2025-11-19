---

page_title: "Volterra: network_connector"

description: "The network_connector allows CRUD of Network Connector resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_network_connector
===================================

The Network Connector allows CRUD of Network Connector resource on Volterra SaaS

~> **Note:** Please refer to [Network Connector API docs](https://docs.cloud.f5.com/docs-v2/api/network-connector) to learn more

Example Usage
-------------

```hcl
resource "volterra_network_connector" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "sli_to_global_dr sli_to_global_snat sli_to_slo_dr sli_to_slo_snat slo_to_global_dr slo_to_global_snat" must be set

  slo_to_global_snat {
    global_vn {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }

    snat_config {
      // One of the arguments from this list "interface_ip snat_pool snat_pool_allocator" must be set

      interface_ip = true

      // One of the arguments from this list "default_gw_snat dynamic_routing" must be set

      default_gw_snat = true
    }
  }

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

###### One of the arguments from this list "sli_to_global_dr, sli_to_global_snat, sli_to_slo_dr, sli_to_slo_snat, slo_to_global_dr, slo_to_global_snat" must be set

`sli_to_global_dr` - (Optional) Site local inside is connected directly to a given global network. See [Connector Choice Sli To Global Dr ](#connector-choice-sli-to-global-dr) below for details.

`sli_to_global_snat` - (Optional) Site local inside is connected to a given global network, using SNAT. See [Connector Choice Sli To Global Snat ](#connector-choice-sli-to-global-snat) below for details.(Deprecated)

`sli_to_slo_dr` - (Optional) Site local inside is connected directly to site local outside (`Bool`).(Deprecated)

`sli_to_slo_snat` - (Optional) Site local inside is connected to site local outside, using SNAT. See [Connector Choice Sli To Slo Snat ](#connector-choice-sli-to-slo-snat) below for details.

`slo_to_global_dr` - (Optional) Site local outside is connected directly to a given global network. See [Connector Choice Slo To Global Dr ](#connector-choice-slo-to-global-dr) below for details.

`slo_to_global_snat` - (Optional) Site local outside is connected directly to a given global network. See [Connector Choice Slo To Global Snat ](#connector-choice-slo-to-global-snat) below for details.(Deprecated)

###### One of the arguments from this list "disable_forward_proxy, enable_forward_proxy" must be set

`disable_forward_proxy` - (Optional) Forward Proxy is disabled for this connector (`Bool`).

`enable_forward_proxy` - (Optional) Forward Proxy is enabled for this connector. See [Forward Proxy Choice Enable Forward Proxy ](#forward-proxy-choice-enable-forward-proxy) below for details.

### Connector Choice Sli To Global Dr

Site local inside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Connector Choice Sli To Global Snat

Site local inside is connected to a given global network, using SNAT.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

`snat_config` - (Optional) SNAT configuration to connect to global network. See [Sli To Global Snat Snat Config ](#sli-to-global-snat-snat-config) below for details.

### Connector Choice Sli To Slo Snat

Site local inside is connected to site local outside, using SNAT.

###### One of the arguments from this list "interface_ip, snat_pool, snat_pool_allocator" must be set

`interface_ip` - (Optional) Interface IP of the outgoing interface will be used for SNAT (`Bool`).

`snat_pool` - (Optional) IP from the ip pool prefix will be used for SNAT (`String`).(Deprecated)

`snat_pool_allocator` - (Optional) IP from the ip pool prefix will be used for SNAT. See [ref](#ref) below for details.(Deprecated)

###### One of the arguments from this list "default_gw_snat, dynamic_routing" must be set

`default_gw_snat` - (Optional) Default route in inside network to SNATed network (`Bool`).

`dynamic_routing` - (Optional) Routes are imported in inside network from SNATed network (`Bool`).(Deprecated)

### Connector Choice Slo To Global Dr

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Connector Choice Slo To Global Snat

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

`snat_config` - (Optional) SNAT configuration to connect to global network. See [Slo To Global Snat Snat Config ](#slo-to-global-snat-snat-config) below for details.

### Custom Certificate Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Enable Disable Choice Disable Interception

Disable Interception.

### Enable Disable Choice Enable Interception

Enable Interception.

### Forward Proxy Choice Enable Forward Proxy

Forward Proxy is enabled for this connector.

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2000 (2 seconds) (`Int`).

`max_connect_attempts` - (Optional) Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1. (`Int`).

###### One of the arguments from this list "no_interception, tls_intercept" can be set

`no_interception` - (Optional) No TLS interception is enabled for this network connector (`Bool`).(Deprecated)

`tls_intercept` - (Optional) Specify TLS interception configuration for the network connector. See [Tls Interception Choice Tls Intercept ](#tls-interception-choice-tls-intercept) below for details.(Deprecated)

`white_listed_ports` - (Optional) Example "tmate" server port (`Int`).

`white_listed_prefixes` - (Optional) Example "tmate" server ip (`String`).

### Interception Policy Choice Enable For All Domains

Enable interception for all domains.

### Interception Policy Choice Policy

Policy to enable/disable specific domains, with implicit enable all domains.

`interception_rules` - (Required) List of ordered rules to enable or disable for TLS interception. See [Policy Interception Rules ](#policy-interception-rules) below for details.

### Interception Rules Domain Match

Domain value or regular expression to match.

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ocsp Stapling Choice Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Ocsp Stapling Choice Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Policy Interception Rules

List of ordered rules to enable or disable for TLS interception.

`domain_match` - (Required) Domain value or regular expression to match. See [Interception Rules Domain Match ](#interception-rules-domain-match) below for details.

###### One of the arguments from this list "disable_interception, enable_interception" must be set

`disable_interception` - (Optional) Disable Interception (`Bool`).

`enable_interception` - (Optional) Enable Interception (`Bool`).

### Pool Choice Interface Ip

Interface IP of the outgoing interface will be used for SNAT.

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Routing Choice Default Gw Snat

Default route in inside network to SNATed network.

### Routing Choice Dynamic Routing

Routes are imported in inside network from SNATed network.

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

### Signing Cert Choice Custom Certificate

Certificates for generating intermediate certificate for TLS interception..

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Custom Certificate Private Key ](#custom-certificate-private-key) below for details.

### Signing Cert Choice Volterra Certificate

F5XC certificates for generating intermediate certificate for TLS interception..

### Sli To Global Snat Snat Config

SNAT configuration to connect to global network.

###### One of the arguments from this list "interface_ip, snat_pool, snat_pool_allocator" must be set

`interface_ip` - (Optional) Interface IP of the outgoing interface will be used for SNAT (`Bool`).

`snat_pool` - (Optional) IP from the ip pool prefix will be used for SNAT (`String`).(Deprecated)

`snat_pool_allocator` - (Optional) IP from the ip pool prefix will be used for SNAT. See [ref](#ref) below for details.(Deprecated)

###### One of the arguments from this list "default_gw_snat, dynamic_routing" must be set

`default_gw_snat` - (Optional) Default route in inside network to SNATed network (`Bool`).

`dynamic_routing` - (Optional) Routes are imported in inside network from SNATed network (`Bool`).(Deprecated)

### Slo To Global Snat Snat Config

SNAT configuration to connect to global network.

###### One of the arguments from this list "interface_ip, snat_pool, snat_pool_allocator" must be set

`interface_ip` - (Optional) Interface IP of the outgoing interface will be used for SNAT (`Bool`).

`snat_pool` - (Optional) IP from the ip pool prefix will be used for SNAT (`String`).(Deprecated)

`snat_pool_allocator` - (Optional) IP from the ip pool prefix will be used for SNAT. See [ref](#ref) below for details.(Deprecated)

###### One of the arguments from this list "default_gw_snat, dynamic_routing" must be set

`default_gw_snat` - (Optional) Default route in inside network to SNATed network (`Bool`).

`dynamic_routing` - (Optional) Routes are imported in inside network from SNATed network (`Bool`).(Deprecated)

### Tls Interception Choice No Interception

No TLS interception is enabled for this network connector.

### Tls Interception Choice Tls Intercept

Specify TLS interception configuration for the network connector.

###### One of the arguments from this list "enable_for_all_domains, policy" must be set

`enable_for_all_domains` - (Optional) Enable interception for all domains (`Bool`).

`policy` - (Optional) Policy to enable/disable specific domains, with implicit enable all domains. See [Interception Policy Choice Policy ](#interception-policy-choice-policy) below for details.

###### One of the arguments from this list "custom_certificate, volterra_certificate" must be set

`custom_certificate` - (Optional) Certificates for generating intermediate certificate for TLS interception.. See [Signing Cert Choice Custom Certificate ](#signing-cert-choice-custom-certificate) below for details.

`volterra_certificate` - (Optional) F5XC certificates for generating intermediate certificate for TLS interception. (`Bool`).

###### One of the arguments from this list "trusted_ca_url, volterra_trusted_ca" must be set

`trusted_ca_url` - (Optional) Custom Root CA Certificate for validating upstream server certificate (`String`).

`volterra_trusted_ca` - (Optional) F5XC Root CA Certificate for validating upstream server certificate (`Bool`).

### Trusted Ca Choice Volterra Trusted Ca

F5XC Root CA Certificate for validating upstream server certificate.

Attribute Reference
-------------------

-	`id` - This is the id of the configured network_connector.
