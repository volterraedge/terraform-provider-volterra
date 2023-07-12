---

page_title: "Volterra: dns_zone"

description: "The dns_zone allows CRUD of Dns Zone resource on Volterra SaaS"
-----------------------------------------------------------------------------

Resource volterra_dns_zone
==========================

The Dns Zone allows CRUD of Dns Zone resource on Volterra SaaS

~> **Note:** Please refer to [Dns Zone API docs](https://volterra.io/docs/api/dns-zone) to learn more

Example Usage
-------------

```hcl
resource "volterra_dns_zone" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "primary secondary" must be set

  secondary {
    primary_servers    = ["8.8.8.8"]
    tsig_key_algorithm = "tsig_key_algorithm"
    tsig_key_name      = "value"

    tsig_key_value {
      blindfold_secret_info_internal {
        decryption_provider = "value"
        location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
        store_provider      = "value"
      }

      secret_encoding_type = "secret_encoding_type"

      // One of the arguments from this list "blindfold_secret_info vault_secret_info clear_secret_info wingman_secret_info" must be set

      vault_secret_info {
        key             = "key_pem"
        location        = "v1/data/vhost_key"
        provider        = "vault-vh-provider"
        secret_encoding = "secret_encoding"
        version         = "1"
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

`primary` - (Optional) Primary DNS. See [Primary ](#primary) below for details.

`secondary` - (Optional) Secondary DNS. See [Secondary ](#secondary) below for details.

### A Record

x-displayName: "A".

`name` - (Optional) x-example: "www or mail or * or corp.web or *.b" (`String`).

`values` - (Required) x-required (`String`).

### Aaaa Record

x-displayName: "AAAA".

`name` - (Optional) x-example: "www or mail or * or corp.web or *.b" (`String`).

`values` - (Required) x-required (`String`).

### Afsdb Record

x-displayName: "AFSDB".

`name` - (Optional) x-example: "www or mail or * or ww* or *ab" (`String`).

`values` - (Required) x-required. See [Values ](#values) below for details.

### Alias Record

x-displayName: "ALIAS".

`name` - (Optional) x-example: "" (`String`).

`value` - (Optional) x-example: "example.com" (`String`).

### Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Caa Record

x-displayName: "CAA".

`name` - (Optional) x-example: "www or mail or * or corp.web or *.b" (`String`).

`values` - (Optional) x-displayName: "CAA Record Value". See [Values ](#values) below for details.

### Cds Record

x-displayName: "CDS".

`name` - (Optional) x-example: "www or mail or * or ww* or *ab" (`String`).

`values` - (Required) x-required. See [Values ](#values) below for details.

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Cname Record

x-displayName: "CNAME".

`name` - (Required) x-required (`String`).

`value` - (Optional) x-example: "example.com" (`String`).

### Default Rr Set Group

default_rr_set_group.

`ttl` - (Optional) x-example: "3600" (`Int`).

`a_record` - (Optional) x-displayName: "A". See [A Record ](#a-record) below for details.

`aaaa_record` - (Optional) x-displayName: "AAAA". See [Aaaa Record ](#aaaa-record) below for details.

`afsdb_record` - (Optional) x-displayName: "AFSDB". See [Afsdb Record ](#afsdb-record) below for details.

`alias_record` - (Optional) x-displayName: "ALIAS". See [Alias Record ](#alias-record) below for details.

`caa_record` - (Optional) x-displayName: "CAA". See [Caa Record ](#caa-record) below for details.

`cds_record` - (Optional) x-displayName: "CDS". See [Cds Record ](#cds-record) below for details.

`cname_record` - (Optional) x-displayName: "CNAME". See [Cname Record ](#cname-record) below for details.

`ds_record` - (Optional) x-displayName: "DS". See [Ds Record ](#ds-record) below for details.

`eui48_record` - (Optional) x-displayName: "EUI48". See [Eui48 Record ](#eui48-record) below for details.

`eui64_record` - (Optional) x-displayName: "EUI64". See [Eui64 Record ](#eui64-record) below for details.

`lb_record` - (Optional) x-displayName: "DNS Load Balancer". See [Lb Record ](#lb-record) below for details.

`loc_record` - (Optional) x-displayName: "LOC". See [Loc Record ](#loc-record) below for details.

`mx_record` - (Optional) x-displayName: "MX". See [Mx Record ](#mx-record) below for details.

`naptr_record` - (Optional) x-displayName: "NAPTR". See [Naptr Record ](#naptr-record) below for details.

`ns_record` - (Optional) x-displayName: "NS". See [Ns Record ](#ns-record) below for details.

`ptr_record` - (Optional) x-displayName: "PTR". See [Ptr Record ](#ptr-record) below for details.

`srv_record` - (Optional) x-displayName: "SRV". See [Srv Record ](#srv-record) below for details.

`txt_record` - (Optional) x-displayName: "TXT". See [Txt Record ](#txt-record) below for details.

### Default Soa Parameters

default_soa_parameters.

### Disable

DNSSEC disabled.

### Dnssec Mode

dnssec_mode.

`disable` - (Optional) DNSSEC disabled (bool).

`enable` - (Optional) DNSSEC enable. See [Enable ](#enable) below for details.

### Ds Record

x-displayName: "DS".

`name` - (Optional) x-example: "www or mail or * or ww* or *ab" (`String`).

`values` - (Required) x-required. See [Values ](#values) below for details.

### Enable

DNSSEC enable.

### Eui48 Record

x-displayName: "EUI48".

`name` - (Optional) x-example: "www or mail or * or ww* or *ab" (`String`).

`value` - (Required) x-required (`String`).

### Eui64 Record

x-displayName: "EUI64".

`name` - (Optional) x-example: "www or mail or * or ww* or *ab" (`String`).

`value` - (Required) x-required (`String`).

### Lb Record

x-displayName: "DNS Load Balancer".

`name` - (Optional) x-example: "www or mail or * or ww* or *ab" (`String`).

`value` - (Optional) x-displayName: "DNS Load Balancer Record". See [ref](#ref) below for details.

### Loc Record

x-displayName: "LOC".

`name` - (Optional) x-example: "www or mail or * or ww* or *ab" (`String`).

`values` - (Required) x-required. See [Values ](#values) below for details.

### Metadata

x-required.

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Mx Record

x-displayName: "MX".

`name` - (Optional) x-example: "www or mail or * or corp.web or *.b" (`String`).

`values` - (Required) x-required. See [Values ](#values) below for details.

### Naptr Record

x-displayName: "NAPTR".

`name` - (Optional) x-example: "www or mail or * or ww* or *ab" (`String`).

`values` - (Required) x-required. See [Values ](#values) below for details.

### Ns Record

x-displayName: "NS".

`name` - (Optional) x-example: "www or mail or * or corp.we" or *.b" (`String`).

`values` - (Required) x-required (`String`).

### Primary

Primary DNS.

`allow_http_lb_managed_records` - (Optional)allow_http_lb_managed_records (`Bool`).

`default_rr_set_group` - (Optional)default_rr_set_group. See [Default Rr Set Group ](#default-rr-set-group) below for details.

`dnssec_mode` - (Optional)dnssec_mode. See [Dnssec Mode ](#dnssec-mode) below for details.

`rr_set_group` - (Optional)rr_set_group. See [Rr Set Group ](#rr-set-group) below for details.

`default_soa_parameters` - (Optional)default_soa_parameters (bool).

`soa_parameters` - (Optional)soa_parameters. See [Soa Parameters ](#soa-parameters) below for details.

### Ptr Record

x-displayName: "PTR".

`name` - (Optional) x-example: "www or mail or * or corp.web or *.b" (`String`).

`values` - (Required) x-required (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rr Set

Collection of DNS resource record sets.

`ttl` - (Optional) x-example: "3600" (`Int`).

`a_record` - (Optional) x-displayName: "A". See [A Record ](#a-record) below for details.

`aaaa_record` - (Optional) x-displayName: "AAAA". See [Aaaa Record ](#aaaa-record) below for details.

`afsdb_record` - (Optional) x-displayName: "AFSDB". See [Afsdb Record ](#afsdb-record) below for details.

`alias_record` - (Optional) x-displayName: "ALIAS". See [Alias Record ](#alias-record) below for details.

`caa_record` - (Optional) x-displayName: "CAA". See [Caa Record ](#caa-record) below for details.

`cds_record` - (Optional) x-displayName: "CDS". See [Cds Record ](#cds-record) below for details.

`cname_record` - (Optional) x-displayName: "CNAME". See [Cname Record ](#cname-record) below for details.

`ds_record` - (Optional) x-displayName: "DS". See [Ds Record ](#ds-record) below for details.

`eui48_record` - (Optional) x-displayName: "EUI48". See [Eui48 Record ](#eui48-record) below for details.

`eui64_record` - (Optional) x-displayName: "EUI64". See [Eui64 Record ](#eui64-record) below for details.

`lb_record` - (Optional) x-displayName: "DNS Load Balancer". See [Lb Record ](#lb-record) below for details.

`loc_record` - (Optional) x-displayName: "LOC". See [Loc Record ](#loc-record) below for details.

`mx_record` - (Optional) x-displayName: "MX". See [Mx Record ](#mx-record) below for details.

`naptr_record` - (Optional) x-displayName: "NAPTR". See [Naptr Record ](#naptr-record) below for details.

`ns_record` - (Optional) x-displayName: "NS". See [Ns Record ](#ns-record) below for details.

`ptr_record` - (Optional) x-displayName: "PTR". See [Ptr Record ](#ptr-record) below for details.

`srv_record` - (Optional) x-displayName: "SRV". See [Srv Record ](#srv-record) below for details.

`txt_record` - (Optional) x-displayName: "TXT". See [Txt Record ](#txt-record) below for details.

### Rr Set Group

rr_set_group.

`metadata` - (Required) x-required. See [Metadata ](#metadata) below for details.

`rr_set` - (Optional) Collection of DNS resource record sets. See [Rr Set ](#rr-set) below for details.

### Secondary

Secondary DNS.

`primary_servers` - (Required) x-required (`String`).

`tsig_key_algorithm` - (Optional) TSIG key value must be compatible with the specified algorithm (`String`).

`tsig_key_name` - (Optional) x-example: "value" (`String`).

`tsig_key_value` - (Optional) x-displayName: "TSIG key value in base 64 format". See [Tsig Key Value ](#tsig-key-value) below for details.

### Soa Parameters

soa_parameters.

`expire` - (Optional) expire value indicates when secondary nameservers should stop answering request for this zone if primary does not respond (`Int`).

`negative_ttl` - (Optional) negative ttl value indicates how long to cache non-existent resource record for this zone (`Int`).

`refresh` - (Optional) refresh value indicates when secondary nameservers should query for the SOA record to detect zone changes (`Int`).

`retry` - (Optional) retry value indicates when secondary nameservers should retry to request the serial number if primary does not respond (`Int`).

`ttl` - (Optional) SOA record time to live (in seconds) (`Int`).

### Srv Record

x-displayName: "SRV".

`name` - (Required) x-required (`String`).

`values` - (Required) x-required. See [Values ](#values) below for details.

### Tsig Key Value

x-displayName: "TSIG key value in base 64 format".

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Txt Record

x-displayName: "TXT".

`name` - (Optional) x-example: "www or mail or * or corp.web or *.b" (`String`).

`values` - (Required) x-required (`String`).

### Values

x-required.

`hostname` - (Required) Server name of the AFS cell database server or the DCE name server. (`String`).

`subtype` - (Required) AFSDB Record Subtype. (`String`).

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured dns_zone.
