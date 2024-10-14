---

page_title: "Volterra: bigip_centralmanager_site"
description: "The bigip_centralmanager_site allows CRUD of Bigip Centralmanager Site resource on Volterra SaaS"

---

Resource volterra_bigip_centralmanager_site
===========================================

The Bigip Centralmanager Site allows CRUD of Bigip Centralmanager Site resource on Volterra SaaS

~> **Note:** Please refer to [Bigip Centralmanager Site API docs](https://docs.cloud.f5.com/docs-v2/api/views-bigip-centralmanager-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_bigip_centralmanager_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
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

`volterra_software` - (Optional) Refer to release notes to find required released SW versions.. See [Volterra Software ](#volterra-software) below for details.

### Volterra Software

Refer to release notes to find required released SW versions..

###### One of the arguments from this list "default_sw_version, volterra_software_version" must be set

`default_sw_version` - (Optional) Will assign latest available F5XC Software Version (`Bool`).

`volterra_software_version` - (Optional) Specify a F5XC Software Version to be used e.g. crt-20210329-1002. (`String`).

### Volterra Sw Version Choice Default Sw Version

Will assign latest available F5XC Software Version.

Attribute Reference
-------------------

-	`id` - This is the id of the configured bigip_centralmanager_site.
