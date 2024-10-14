---

page_title: "Volterra: uztna_domain"
description: "The uztna_domain allows CRUD of Uztna Domain resource on Volterra SaaS"

---

Resource volterra_uztna_domain
==============================

The Uztna Domain allows CRUD of Uztna Domain resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Domain API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-domain) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_domain" "example" {
  name       = "acmecorp-web"
  namespace  = "staging"
  access_url = ["https://example.com/landing/page"]

  cert {
    certificate {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  gateways {
    bigip_ce {
      uztna_gateway {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }
    }

    perimeter_re {
      // One of the arguments from this list "all_perimeter domain_re_sites" can be set

      all_perimeter = true
    }
  }

  lease_pool {
    uztna_lpool {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  profile_name = ["profile_name"]
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

`access_url` - (Required) Url to access the gateways (`String`).

`cdn_ce_vh_api_gw` - (Optional) Internal reference to cdn ce api gateway VH. See [ref](#ref) below for details.(Deprecated)

`cert` - (Required) Domain in XC is an establis. See [Cert ](#cert) below for details.

`gateways` - (Required) List of all RE prime and Big Ip edge CE. See [Gateways ](#gateways) below for details.

`lease_pool` - (Required) This is List of LeasePool Object created.From the List we can assign one LeasePool to one ZeroTrust Domain. . See [Lease Pool ](#lease-pool) below for details.

`profile_name` - (Required) Ztna Profile Name used to refer the ZTNA profile (`String`).

### Cert

Domain in XC is an establis.

`certificate` - (Optional) Select/Add one or more TLS Certificate objects to associate with this ZeroTrust Domain. See [ref](#ref) below for details.

### Gateways

List of all RE prime and Big Ip edge CE.

`bigip_ce` - (Optional) From the available bigip CE List select bigip CE .. See [Gateways Bigip Ce ](#gateways-bigip-ce) below for details.

`perimeter_re` - (Optional) Select the Gateways (either Perimeter RE or BigIP CE).. See [Gateways Perimeter Re ](#gateways-perimeter-re) below for details.

### Lease Pool

This is List of LeasePool Object created.From the List we can assign one LeasePool to one ZeroTrust Domain. .

`uztna_lpool` - (Optional) Lease Pool for UZTNA Domain. See [ref](#ref) below for details.

### Gateways Bigip Ce

From the available bigip CE List select bigip CE ..

`uztna_gateway` - (Optional) List of All avaliable Big IP CE List. See [ref](#ref) below for details.

### Gateways Perimeter Re

Select the Gateways (either Perimeter RE or BigIP CE)..

###### One of the arguments from this list "all_perimeter, domain_re_sites" can be set

`all_perimeter` - (Optional) This option will allow to advertise on all available perimeter RE sites (`Bool`).

`domain_re_sites` - (Optional) This option will allow advertise on specific Perimeter RE sites. See [Perimter Re Choice Domain Re Sites ](#perimter-re-choice-domain-re-sites) below for details.

### Perimter Re Choice All Perimeter

This option will allow to advertise on all available perimeter RE sites.

### Perimter Re Choice Domain Re Sites

This option will allow advertise on specific Perimeter RE sites.

`perimeter_re_site` - (Optional) Selected Perimeter RE Site.. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_domain.
