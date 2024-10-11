---

page_title: "Volterra: dns_zone_record"
description: "The dns_zone_record allows CRUD of Dns Load Balancer resource on Volterra SaaS"

---

Resource volterra_dns_zone_record
=================================

The Dns Zone Record allows CRUD of Dns Load Balancer resource on Volterra SaaS

~> **Note:** Please refer to [Dns Zone Record API docs](https://docs.cloud.f5.com/docs/api/dns-zone-rrset) to learn more

Example Usage
-------------

```hcl
resource "volterra_dns_zone_record" "example" {
  dns_zone_name = "www.dns-zone.com"
  group_name = "default"
  rrset {
    description = "description"
    ttl = "68"
    a_record {
      name = "name"
      values = ["1.2.1.1"]
    }
  }
}

```

Argument Reference
------------------

### Spec Argument Reference

`dns_zone_name` - (Required) dns zone name.

`group_name` - (Required) Name of Group where all records will be created.

`rrset` - (Required) Collection of DNS resource record sets. See [Rr Set](#rr-set) below for details.

### Rr Set

Collection of DNS resource record sets.

`description` - (Optional) x-displayName: "Comment" (`String`).

`ttl` - (Optional) x-example: "3600" (`Int`).

###### One of the arguments from this list "srv_record, lb_record, naptr_record, ds_record, loc_record, sshfp_record, aaaa_record, cname_record, tlsa_record, dlv_record, cert_record, mx_record, txt_record, afsdb_record, eui48_record, alias_record, ns_record, ptr_record, cds_record, eui64_record, a_record, caa_record" must be set

`a_record` - (Optional) x-displayName: "A". See [A Record ](#a-record) below for details.

`aaaa_record` - (Optional) x-displayName: "AAAA". See [Aaaa Record ](#aaaa-record) below for details.

`afsdb_record` - (Optional) x-displayName: "AFSDB". See [Afsdb Record ](#afsdb-record) below for details.

`alias_record` - (Optional) x-displayName: "ALIAS". See [Alias Record ](#alias-record) below for details.

`caa_record` - (Optional) x-displayName: "CAA". See [Caa Record ](#caa-record) below for details.

`cds_record` - (Optional) x-displayName: "CDS". See [Cds Record ](#cds-record) below for details.

`cert_record` - (Optional) x-displayName: "CERT". See [Cert Record ](#cert-record) below for details.

`cname_record` - (Optional) x-displayName: "CNAME". See [Cname Record ](#cname-record) below for details.

`dlv_record` - (Optional) x-displayName: "DLV". See [Dlv Record ](#dlv-record) below for details.(Deprecated)

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

`sshfp_record` - (Optional) x-displayName: "SSHFP". See [Sshfp Record ](#sshfp-record) below for details.

`tlsa_record` - (Optional) x-displayName: "TLSA". See [Tlsa Record ](#tlsa-record) below for details.

`txt_record` - (Optional) x-displayName: "TXT". See [Txt Record ](#txt-record) below for details.

### A Record

x-displayName: "A".

`name` - (Optional) A Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) A valid IPv4 address, for example: 1.1.1.1 (`String`).

### Aaaa Record

x-displayName: "AAAA".

`name` - (Optional) AAAA Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334 (`String`).

### Afsdb Record

x-displayName: "AFSDB".

`name` - (Optional) AFSDB Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Afsdb Record Values ](#afsdb-record-values) below for details.

### Alias Record

x-displayName: "ALIAS".

`name` - (Optional) Alias Record name, please provide only the specific subdomain or record name without the base domain. (`String`).(Deprecated)

`value` - (Optional) A valid domain name, for example: example.com (`String`).

### Caa Record

x-displayName: "CAA".

`name` - (Optional) CAA Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Optional) x-displayName: "CAA Record Value". See [Caa Record Values ](#caa-record-values) below for details.

### Cds Record

x-displayName: "CDS".

`name` - (Optional) CDS Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Cds Record Values ](#cds-record-values) below for details.

### Cert Record

x-displayName: "CERT".

`name` - (Optional) CERT Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Cert Record Values ](#cert-record-values) below for details.

### Cname Record

x-displayName: "CNAME".

`name` - (Required) CName Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`value` - (Optional) x-example: "example.com" (`String`).

### Dlv Record

x-displayName: "DLV".

`name` - (Optional) DLV Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) It uses the same format as the DS record.. See [Dlv Record Values ](#dlv-record-values) below for details.

### Ds Record

x-displayName: "DS".

`name` - (Optional) DS Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Ds Record Values ](#ds-record-values) below for details.

### Eui48 Record

x-displayName: "EUI48".

`name` - (Optional) EUI48 Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`value` - (Required) A valid eui48 identifier, for example: 01-23-45-67-89-ab (`String`).

### Eui64 Record

x-displayName: "EUI64".

`name` - (Optional) EUI64 Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`value` - (Required) A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef (`String`).

### Lb Record

x-displayName: "DNS Load Balancer".

`name` - (Optional) Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name and not a subdomain of a subdomain. (`String`).

`value` - (Optional) x-displayName: "DNS Load Balancer Record". See [ref](#ref) below for details.

### Loc Record

x-displayName: "LOC".

`name` - (Optional) LOC Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Loc Record Values ](#loc-record-values) below for details.

### Mx Record

x-displayName: "MX".

`name` - (Optional) MX Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Mx Record Values ](#mx-record-values) below for details.

### Naptr Record

x-displayName: "NAPTR".

`name` - (Optional) NAPTR Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Naptr Record Values ](#naptr-record-values) below for details.

### Ns Record

x-displayName: "NS".

`name` - (Optional) NS Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required (`String`).

### Ptr Record

x-displayName: "PTR".

`name` - (Optional) PTR Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required (`String`).

### Srv Record

x-displayName: "SRV".

`name` - (Required) SRV Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Srv Record Values ](#srv-record-values) below for details.

### Sshfp Record

x-displayName: "SSHFP".

`name` - (Optional) SSHFP Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Sshfp Record Values ](#sshfp-record-values) below for details.

### Tlsa Record

x-displayName: "TLSA".

`name` - (Optional) TLSA Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required. See [Tlsa Record Values ](#tlsa-record-values) below for details.

### Txt Record

x-displayName: "TXT".

`name` - (Optional) TXT Record name, please provide only the specific subdomain or record name without the base domain. (`String`).

`values` - (Required) x-required (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured dns_zone_record.

Known Limitation
----------------

Adding Resource Record Sets to DNS Zones

`Description:`

Currently, there is a limitation when adding individual resource record sets to a DNS domain through the provider resource. To successfully add a resource record set, an initial setup must be performed through the F5XC console.

`Steps to Reproduce:`

Attempt to add a resource record set directly through the provider resource without prior setup in the F5XC console.

`Observed Behavior:`

The addition of the resource record set fails due to the absence of prior configuration.

`Workaround:`

To add a resource record set to a DNS domain, follow these steps:

-	Navigate to the `Resource Record Sets` tile in the F5XC console.

-	Enable the `Show Advanced Fields` option.

-	In the `Additional Resource Record Sets` section, add the desired resource record set.

Once the initial setup is complete, you can add individual resource records to the DNS zone through the provider resource without encountering this limitation.

`Upcoming Fix:`

We are aware of this limitation and are actively working on a fix. This issue will be resolved in the next release, allowing for the direct addition of individual resource record sets to DNS zones without the need for prior configuration through the F5XC console.
