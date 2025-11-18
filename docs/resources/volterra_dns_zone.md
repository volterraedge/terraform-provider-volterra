---

page_title: "Volterra: dns_zone"

description: "The dns_zone allows CRUD of Dns Zone resource on Volterra SaaS"
-----------------------------------------------------------------------------

Resource volterra_dns_zone
==========================

The Dns Zone allows CRUD of Dns Zone resource on Volterra SaaS

~> **Note:** Please refer to [Dns Zone API docs](https://docs.cloud.f5.com/docs-v2/api/dns-zone) to learn more

Example Usage
-------------

```hcl
resource "volterra_dns_zone" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "primary secondary" must be set

  primary {
    allow_http_lb_managed_records = true

    default_rr_set_group {
      description = "Comment"

      ttl = "3600"

      // One of the arguments from this list "a_record aaaa_record afsdb_record alias_record caa_record cds_record cert_record cname_record dlv_record ds_record eui48_record eui64_record lb_record loc_record mx_record naptr_record ns_record ptr_record srv_record sshfp_record tlsa_record txt_record" must be set

      ptr_record {
        name = "www or mail or * or corp.web or *.b"

        values = ["my.example.com"]
      }
    }

    dnssec_mode {
      // One of the arguments from this list "disable enable" must be set

      disable = true
    }

    rr_set_group {
      metadata {
        description = "Virtual Host for acmecorp website"

        disable = true

        name = "acmecorp-web"
      }

      rr_set {
        description = "Comment"

        ttl = "3600"

        // One of the arguments from this list "a_record aaaa_record afsdb_record alias_record caa_record cds_record cert_record cname_record dlv_record ds_record eui48_record eui64_record lb_record loc_record mx_record naptr_record ns_record ptr_record srv_record sshfp_record tlsa_record txt_record" must be set

        caa_record {
          name = "www or mail or * or corp.web or *.b"

          values {
            flags = "0"

            tag = "issue"

            value = "value"
          }
        }
      }
    }

    // One of the arguments from this list "default_soa_parameters soa_parameters" can be set

    default_soa_parameters = true
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

###### One of the arguments from this list "primary, secondary" must be set

`primary` - (Optional) Primary DNS. See [Dns Type Primary ](#dns-type-primary) below for details.

`secondary` - (Optional) Secondary DNS. See [Dns Type Secondary ](#dns-type-secondary) below for details.

### Afsdb Record Values

x-required.

`hostname` - (Required) Server name of the AFS cell database server or the DCE name server. (`String`).

`subtype` - (Required) AFSDB Record Subtype. (`String`).

### Caa Record Values

x-displayName: "CAA Record Value".

`flags` - (Optional) This flag should be an integer between 0 and 255. (`Int`).

`tag` - (Optional) x-example: "issue" (`String`).

`value` - (Optional) x-example: "value" (`String`).

### Cds Record Values

x-required.

###### One of the arguments from this list "sha1_digest, sha256_digest, sha384_digest" must be set

`sha1_digest` - (Optional) x-displayName: "SHA1 Digest". See [Digest Choice Sha1 Digest ](#digest-choice-sha1-digest) below for details.

`sha256_digest` - (Optional) x-displayName: "SHA256 Digest". See [Digest Choice Sha256 Digest ](#digest-choice-sha256-digest) below for details.

`sha384_digest` - (Optional) x-displayName: "SHA384 Digest". See [Digest Choice Sha384 Digest ](#digest-choice-sha384-digest) below for details.

`ds_key_algorithm` - (Required) DS key value must be compatible with the specified algorithm. (`String`).

`key_tag` - (Required) A short numeric value which can help quickly identify the referenced DNSKEY-record. (`Int`).

### Cert Record Values

x-required.

`algorithm` - (Required) x-required (`String`).

`cert_key_tag` - (Required) x-required (`Int`).

`cert_type` - (Required) x-required (`String`).

`certificate` - (Required) Certificate in base 64 format. (`String`).

### Digest Choice Sha1 Digest

x-displayName: "SHA1 Digest".

`digest` - (Required) The 'digest' is the DS key and the actual contents of the DS record. (`String`).

### Digest Choice Sha256 Digest

x-displayName: "SHA256 Digest".

`digest` - (Required) The 'digest' is the DS key and the actual contents of the DS record. (`String`).

### Digest Choice Sha384 Digest

x-displayName: "SHA384 Digest".

`digest` - (Required) The 'digest' is the DS key and the actual contents of the DS record. (`String`).

### Dlv Record Values

It uses the same format as the DS record..

###### One of the arguments from this list "sha1_digest, sha256_digest, sha384_digest" must be set

`sha1_digest` - (Optional) x-displayName: "SHA1 Digest". See [Digest Choice Sha1 Digest ](#digest-choice-sha1-digest) below for details.

`sha256_digest` - (Optional) x-displayName: "SHA256 Digest". See [Digest Choice Sha256 Digest ](#digest-choice-sha256-digest) below for details.

`sha384_digest` - (Optional) x-displayName: "SHA384 Digest". See [Digest Choice Sha384 Digest ](#digest-choice-sha384-digest) below for details.

`ds_key_algorithm` - (Required) DS key value must be compatible with the specified algorithm. (`String`).

`key_tag` - (Required) A short numeric value which can help quickly identify the referenced DNSKEY-record. (`Int`).

### Dns Type Primary

Primary DNS.

`allow_http_lb_managed_records` - (Optional)allow_http_lb_managed_records (`Bool`).

`default_rr_set_group` - (Optional)default_rr_set_group. See [Primary Default Rr Set Group ](#primary-default-rr-set-group) below for details.

`dnssec_mode` - (Optional)dnssec_mode. See [Primary Dnssec Mode ](#primary-dnssec-mode) below for details.

`rr_set_group` - (Optional)rr_set_group. See [Primary Rr Set Group ](#primary-rr-set-group) below for details.

###### One of the arguments from this list "default_soa_parameters, soa_parameters" can be set

`default_soa_parameters` - (Optional)default_soa_parameters (`Bool`).

`soa_parameters` - (Optional)soa_parameters. See [Soa Record Parameters Choice Soa Parameters ](#soa-record-parameters-choice-soa-parameters) below for details.

### Dns Type Secondary

Secondary DNS.

`primary_servers` - (Required) x-required (`String`).

`tsig_key_algorithm` - (Optional) TSIG key value must be compatible with the specified algorithm (`String`).

`tsig_key_name` - (Optional) x-example: "value" (`String`).

`tsig_key_value` - (Optional) x-displayName: "TSIG Key Value in Base 64 Format". See [Secondary Tsig Key Value ](#secondary-tsig-key-value) below for details.

### Ds Record Values

x-required.

###### One of the arguments from this list "sha1_digest, sha256_digest, sha384_digest" must be set

`sha1_digest` - (Optional) x-displayName: "SHA1 Digest". See [Digest Choice Sha1 Digest ](#digest-choice-sha1-digest) below for details.

`sha256_digest` - (Optional) x-displayName: "SHA256 Digest". See [Digest Choice Sha256 Digest ](#digest-choice-sha256-digest) below for details.

`sha384_digest` - (Optional) x-displayName: "SHA384 Digest". See [Digest Choice Sha384 Digest ](#digest-choice-sha384-digest) below for details.

`ds_key_algorithm` - (Required) DS key value must be compatible with the specified algorithm. (`String`).

`key_tag` - (Required) A short numeric value which can help quickly identify the referenced DNSKEY-record. (`Int`).

### Fingerprint Type Sha1 Fingerprint

x-displayName: "SHA1 Fingerprint".

`fingerprint` - (Required) The 'fingerprint' is the DS key and the actual contents of the DS record. (`String`).

### Fingerprint Type Sha256 Fingerprint

x-displayName: "SHA256 Fingerprint".

`fingerprint` - (Required) The 'fingerprint' is the DS key and the actual contents of the DS record. (`String`).

### Loc Record Values

x-required.

`altitude` - (Required) Altitude in meters (`Float`).

`horizontal_precision` - (Optional) Horizontal Precision in meters (`Float`).

`latitude_degree` - (Required) Latitude degree, an integer between 0 and 90, including 0 and 90 (`Int`).

`latitude_hemisphere` - (Required) Latitude hemisphere, N or S (`String`).

`latitude_minute` - (Optional) Latitude minute, an integer between 0 and 59, including 0 and 59 (`Int`).

`latitude_second` - (Optional) Latitude second, an decimal between 0 and 59.999, including 0 and 59.999 (`Float`).

`location_diameter` - (Optional) Diameter of a sphere enclosing the described entity, in meters (`Float`).

`longitude_degree` - (Required) Longitude degree, an integer between 0 and 180, including 0 and 180 (`Int`).

`longitude_hemisphere` - (Required) Longitude hemisphere, E or W (`String`).

`longitude_minute` - (Optional) Longitude minute, an integer between 0 and 59, including 0 and 59 (`Int`).

`longitude_second` - (Optional) Longitude second, an decimal between 0 and 59.999, including 0 and 59.999 (`Float`).

`vertical_precision` - (Optional) Vertical Precision in meters (`Float`).

### Mode Disable

DNSSEC disabled.

### Mode Enable

DNSSEC enable.

### Mx Record Values

x-required.

`domain` - (Optional) Mail exchanger domain name, please provide the full hostname, for example: mail.example.com (`String`).

`priority` - (Optional) Mail exchanger priority code (`Int`).

### Naptr Record Values

x-required.

`flags` - (Required) Flag to control aspects of the rewriting and interpretation of the fields in the record. At this time only four flags, S/A/U/P, are defined. (`String`).

`order` - (Required) Order in which the NAPTR records must be processed. A lower number indicates a higher preference. (`Int`).

`preference` - (Required) Preference when records have the same order. A lower number indicates a higher preference. (`Int`).

`regexp` - (Optional) Regular expression to construct the next domain name to lookup. (`String`).

`replacement` - (Optional) The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field. (`String`).

`service` - (Optional) Specifies the service(s) available down this rewrite path. (`String`).

### Primary Default Rr Set Group

default_rr_set_group.

`description` - (Optional) x-displayName: "Comment" (`String`).

`ttl` - (Optional) x-example: "3600" (`Int`).

###### One of the arguments from this list "a_record, aaaa_record, afsdb_record, alias_record, caa_record, cds_record, cert_record, cname_record, dlv_record, ds_record, eui48_record, eui64_record, lb_record, loc_record, mx_record, naptr_record, ns_record, ptr_record, srv_record, sshfp_record, tlsa_record, txt_record" must be set

`a_record` - (Optional) x-displayName: "A". See [Type Record Set A Record ](#type-record-set-a-record) below for details.

`aaaa_record` - (Optional) x-displayName: "AAAA". See [Type Record Set Aaaa Record ](#type-record-set-aaaa-record) below for details.

`afsdb_record` - (Optional) x-displayName: "AFSDB". See [Type Record Set Afsdb Record ](#type-record-set-afsdb-record) below for details.

`alias_record` - (Optional) x-displayName: "ALIAS". See [Type Record Set Alias Record ](#type-record-set-alias-record) below for details.

`caa_record` - (Optional) x-displayName: "CAA". See [Type Record Set Caa Record ](#type-record-set-caa-record) below for details.

`cds_record` - (Optional) x-displayName: "CDS". See [Type Record Set Cds Record ](#type-record-set-cds-record) below for details.

`cert_record` - (Optional) x-displayName: "CERT". See [Type Record Set Cert Record ](#type-record-set-cert-record) below for details.

`cname_record` - (Optional) x-displayName: "CNAME". See [Type Record Set Cname Record ](#type-record-set-cname-record) below for details.

`dlv_record` - (Optional) x-displayName: "DLV". See [Type Record Set Dlv Record ](#type-record-set-dlv-record) below for details.(Deprecated)

`ds_record` - (Optional) x-displayName: "DS". See [Type Record Set Ds Record ](#type-record-set-ds-record) below for details.

`eui48_record` - (Optional) x-displayName: "EUI48". See [Type Record Set Eui48 Record ](#type-record-set-eui48-record) below for details.

`eui64_record` - (Optional) x-displayName: "EUI64". See [Type Record Set Eui64 Record ](#type-record-set-eui64-record) below for details.

`lb_record` - (Optional) x-displayName: "DNS Load Balancer". See [Type Record Set Lb Record ](#type-record-set-lb-record) below for details.

`loc_record` - (Optional) x-displayName: "LOC". See [Type Record Set Loc Record ](#type-record-set-loc-record) below for details.

`mx_record` - (Optional) x-displayName: "MX". See [Type Record Set Mx Record ](#type-record-set-mx-record) below for details.

`naptr_record` - (Optional) x-displayName: "NAPTR". See [Type Record Set Naptr Record ](#type-record-set-naptr-record) below for details.

`ns_record` - (Optional) x-displayName: "NS". See [Type Record Set Ns Record ](#type-record-set-ns-record) below for details.

`ptr_record` - (Optional) x-displayName: "PTR". See [Type Record Set Ptr Record ](#type-record-set-ptr-record) below for details.

`srv_record` - (Optional) x-displayName: "SRV". See [Type Record Set Srv Record ](#type-record-set-srv-record) below for details.

`sshfp_record` - (Optional) x-displayName: "SSHFP". See [Type Record Set Sshfp Record ](#type-record-set-sshfp-record) below for details.

`tlsa_record` - (Optional) x-displayName: "TLSA". See [Type Record Set Tlsa Record ](#type-record-set-tlsa-record) below for details.

`txt_record` - (Optional) x-displayName: "TXT". See [Type Record Set Txt Record ](#type-record-set-txt-record) below for details.

### Primary Dnssec Mode

dnssec_mode.

###### One of the arguments from this list "disable, enable" must be set

`disable` - (Optional) DNSSEC disabled (`Bool`).

`enable` - (Optional) DNSSEC enable. See [Mode Enable ](#mode-enable) below for details.

### Primary Rr Set Group

rr_set_group.

`metadata` - (Required) x-required. See [Rr Set Group Metadata ](#rr-set-group-metadata) below for details.

`rr_set` - (Optional) Collection of DNS resource record sets. See [Rr Set Group Rr Set ](#rr-set-group-rr-set) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Rr Set Group Metadata

x-required.

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Rr Set Group Rr Set

Collection of DNS resource record sets.

`description` - (Optional) x-displayName: "Comment" (`String`).

`ttl` - (Optional) x-example: "3600" (`Int`).

###### One of the arguments from this list "a_record, aaaa_record, afsdb_record, alias_record, caa_record, cds_record, cert_record, cname_record, dlv_record, ds_record, eui48_record, eui64_record, lb_record, loc_record, mx_record, naptr_record, ns_record, ptr_record, srv_record, sshfp_record, tlsa_record, txt_record" must be set

`a_record` - (Optional) x-displayName: "A". See [Type Record Set A Record ](#type-record-set-a-record) below for details.

`aaaa_record` - (Optional) x-displayName: "AAAA". See [Type Record Set Aaaa Record ](#type-record-set-aaaa-record) below for details.

`afsdb_record` - (Optional) x-displayName: "AFSDB". See [Type Record Set Afsdb Record ](#type-record-set-afsdb-record) below for details.

`alias_record` - (Optional) x-displayName: "ALIAS". See [Type Record Set Alias Record ](#type-record-set-alias-record) below for details.

`caa_record` - (Optional) x-displayName: "CAA". See [Type Record Set Caa Record ](#type-record-set-caa-record) below for details.

`cds_record` - (Optional) x-displayName: "CDS". See [Type Record Set Cds Record ](#type-record-set-cds-record) below for details.

`cert_record` - (Optional) x-displayName: "CERT". See [Type Record Set Cert Record ](#type-record-set-cert-record) below for details.

`cname_record` - (Optional) x-displayName: "CNAME". See [Type Record Set Cname Record ](#type-record-set-cname-record) below for details.

`dlv_record` - (Optional) x-displayName: "DLV". See [Type Record Set Dlv Record ](#type-record-set-dlv-record) below for details.(Deprecated)

`ds_record` - (Optional) x-displayName: "DS". See [Type Record Set Ds Record ](#type-record-set-ds-record) below for details.

`eui48_record` - (Optional) x-displayName: "EUI48". See [Type Record Set Eui48 Record ](#type-record-set-eui48-record) below for details.

`eui64_record` - (Optional) x-displayName: "EUI64". See [Type Record Set Eui64 Record ](#type-record-set-eui64-record) below for details.

`lb_record` - (Optional) x-displayName: "DNS Load Balancer". See [Type Record Set Lb Record ](#type-record-set-lb-record) below for details.

`loc_record` - (Optional) x-displayName: "LOC". See [Type Record Set Loc Record ](#type-record-set-loc-record) below for details.

`mx_record` - (Optional) x-displayName: "MX". See [Type Record Set Mx Record ](#type-record-set-mx-record) below for details.

`naptr_record` - (Optional) x-displayName: "NAPTR". See [Type Record Set Naptr Record ](#type-record-set-naptr-record) below for details.

`ns_record` - (Optional) x-displayName: "NS". See [Type Record Set Ns Record ](#type-record-set-ns-record) below for details.

`ptr_record` - (Optional) x-displayName: "PTR". See [Type Record Set Ptr Record ](#type-record-set-ptr-record) below for details.

`srv_record` - (Optional) x-displayName: "SRV". See [Type Record Set Srv Record ](#type-record-set-srv-record) below for details.

`sshfp_record` - (Optional) x-displayName: "SSHFP". See [Type Record Set Sshfp Record ](#type-record-set-sshfp-record) below for details.

`tlsa_record` - (Optional) x-displayName: "TLSA". See [Type Record Set Tlsa Record ](#type-record-set-tlsa-record) below for details.

`txt_record` - (Optional) x-displayName: "TXT". See [Type Record Set Txt Record ](#type-record-set-txt-record) below for details.

### Secondary Tsig Key Value

x-displayName: "TSIG Key Value in Base 64 Format".

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Tsig Key Value Blindfold Secret Info Internal ](#tsig-key-value-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

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

### Soa Record Parameters Choice Default Soa Parameters

default_soa_parameters.

### Soa Record Parameters Choice Soa Parameters

soa_parameters.

`expire` - (Optional) expire value indicates when secondary nameservers should stop answering request for this zone if primary does not respond (`Int`).

`negative_ttl` - (Optional) negative ttl value indicates how long to cache non-existent resource record for this zone (`Int`).

`refresh` - (Optional) refresh value indicates when secondary nameservers should query for the SOA record to detect zone changes (`Int`).

`retry` - (Optional) retry value indicates when secondary nameservers should retry to request the serial number if primary does not respond (`Int`).

`ttl` - (Optional) SOA record time to live (in seconds) (`Int`).

### Srv Record Values

x-required.

`port` - (Optional) Port on which the service can be found (`Int`).

`priority` - (Optional) Priority of the target. A lower number indicates a higher preference. (`Int`).

`target` - (Optional) Hostname of the machine providing the service (`String`).

`weight` - (Optional) Weight of the target. A higher number indicates a higher preference. (`Int`).

### Sshfp Record Values

x-required.

`algorithm` - (Required) Algorithm of the public key. (`String`).

`fingerprint` - (Optional) The hexadecimal representation of the hash result of the SSH key as text. (`String`).(Deprecated)

###### One of the arguments from this list "sha1_fingerprint, sha256_fingerprint" must be set

`sha1_fingerprint` - (Optional) x-displayName: "SHA1 Fingerprint". See [Fingerprint Type Sha1 Fingerprint ](#fingerprint-type-sha1-fingerprint) below for details.

`sha256_fingerprint` - (Optional) x-displayName: "SHA256 Fingerprint". See [Fingerprint Type Sha256 Fingerprint ](#fingerprint-type-sha256-fingerprint) below for details.

`fingerprinttype` - (Required) Algorithm used to calculate the fingerprint of the public key. (`String`).(Deprecated)

### Tlsa Record Values

x-required.

`certificate_association_data` - (Required) The actual data to be matched given the settings of the other fields. (`String`).

`certificate_usage` - (Required) TLSA Record Certificate Usage. (`String`).

`matching_type` - (Required) TLSA Record Matching Type. (`String`).

`selector` - (Required) TLSA Record Selector. (`String`).

### Tsig Key Value Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Type Record Set A Record

x-displayName: "A".

`name` - (Optional) A Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) A valid IPv4 address, for example: 1.1.1.1 (`String`).

### Type Record Set Aaaa Record

x-displayName: "AAAA".

`name` - (Optional) AAAA Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334 (`String`).

### Type Record Set Afsdb Record

x-displayName: "AFSDB".

`name` - (Optional) AFSDB Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Afsdb Record Values ](#afsdb-record-values) below for details.

### Type Record Set Alias Record

x-displayName: "ALIAS".

`name` - (Optional) Alias Record name, please provide only the specific subdomain or record name without the base domain. (`String`).(Deprecated)

`value` - (Optional) A valid domain name, for example: example.com (`String`).

### Type Record Set Caa Record

x-displayName: "CAA".

`name` - (Optional) CAA Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Optional) x-displayName: "CAA Record Value". See [Caa Record Values ](#caa-record-values) below for details.

### Type Record Set Cds Record

x-displayName: "CDS".

`name` - (Optional) CDS Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Cds Record Values ](#cds-record-values) below for details.

### Type Record Set Cert Record

x-displayName: "CERT".

`name` - (Optional) CERT Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Cert Record Values ](#cert-record-values) below for details.

### Type Record Set Cname Record

x-displayName: "CNAME".

`name` - (Required) CName Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`value` - (Optional) x-example: "example.com" (`String`).

### Type Record Set Dlv Record

x-displayName: "DLV".

`name` - (Optional) DLV Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) It uses the same format as the DS record.. See [Dlv Record Values ](#dlv-record-values) below for details.

### Type Record Set Ds Record

x-displayName: "DS".

`name` - (Optional) DS Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Ds Record Values ](#ds-record-values) below for details.

### Type Record Set Eui48 Record

x-displayName: "EUI48".

`name` - (Optional) EUI48 Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`value` - (Required) A valid eui48 identifier, for example: 01-23-45-67-89-ab (`String`).

### Type Record Set Eui64 Record

x-displayName: "EUI64".

`name` - (Optional) EUI64 Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`value` - (Required) A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef (`String`).

### Type Record Set Lb Record

x-displayName: "DNS Load Balancer".

`name` - (Optional) Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name and not a subdomain of a subdomain. (`String`).

`value` - (Optional) x-displayName: "DNS Load Balancer Record". See [ref](#ref) below for details.

### Type Record Set Loc Record

x-displayName: "LOC".

`name` - (Optional) LOC Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Loc Record Values ](#loc-record-values) below for details.

### Type Record Set Mx Record

x-displayName: "MX".

`name` - (Optional) MX Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Mx Record Values ](#mx-record-values) below for details.

### Type Record Set Naptr Record

x-displayName: "NAPTR".

`name` - (Optional) NAPTR Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Naptr Record Values ](#naptr-record-values) below for details.

### Type Record Set Ns Record

x-displayName: "NS".

`name` - (Optional) NS Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required (`String`).

### Type Record Set Ptr Record

x-displayName: "PTR".

`name` - (Optional) PTR Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required (`String`).

### Type Record Set Srv Record

x-displayName: "SRV".

`name` - (Required) SRV Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Srv Record Values ](#srv-record-values) below for details.

### Type Record Set Sshfp Record

x-displayName: "SSHFP".

`name` - (Optional) SSHFP Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Sshfp Record Values ](#sshfp-record-values) below for details.

### Type Record Set Tlsa Record

x-displayName: "TLSA".

`name` - (Optional) TLSA Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Tlsa Record Values ](#tlsa-record-values) below for details.

### Type Record Set Txt Record

x-displayName: "TXT".

`name` - (Optional) TXT Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured dns_zone.
