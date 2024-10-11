---

page_title: "Volterra: uztna_domain_view"
description: "The uztna_domain_view allows CRUD of Uztna Domain View resource on Volterra SaaS"

---

Resource volterra_uztna_domain_view
===================================

The Uztna Domain View allows CRUD of Uztna Domain View resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Domain View API docs](https://docs.cloud.f5.com/docs-v2/api/views-uztna-domain-view) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_domain_view" "example" {
  name       = "acmecorp-web"
  namespace  = "staging"
  access_url = ["access_url"]

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
      // One of the arguments from this list "all_perimeter re_sites" can be set

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

`cert` - (Required) Domain in XC is an established pattern and we would reuse the same.. See [Cert ](#cert) below for details.

`gateways` - (Required) List of all RE prime and Big Ip edge CE. See [Gateways ](#gateways) below for details.

`lease_pool` - (Required) The Lease Pool assigned to the Zero Trust Domain. . See [Lease Pool ](#lease-pool) below for details.

`profile_name` - (Required) The name of the ZTNA profile (`String`).

### Cert

Domain in XC is an established pattern and we would reuse the same..

`certificate` - (Optional) Select/Add one or more TLS Certificate objects to associate with this ZeroTrust Domain. See [ref](#ref) below for details.

### Gateways

List of all RE prime and Big Ip edge CE.

`bigip_ce` - (Optional) From the available bigip CE List select bigip CE .. See [Gateways Bigip Ce ](#gateways-bigip-ce) below for details.

`perimeter_re` - (Optional) Select the Gateways (either Perimeter RE or BigIP CE).. See [Gateways Perimeter Re ](#gateways-perimeter-re) below for details.

### Lease Pool

The Lease Pool assigned to the Zero Trust Domain. .

`uztna_lpool` - (Optional) Lease Pool for UZTNA Domain View. See [ref](#ref) below for details.

### Gateways Bigip Ce

From the available bigip CE List select bigip CE ..

`uztna_gateway` - (Optional) Selected Big IP CE . See [ref](#ref) below for details.

### Gateways Perimeter Re

Select the Gateways (either Perimeter RE or BigIP CE)..

###### One of the arguments from this list "all_perimeter, re_sites" can be set

`all_perimeter` - (Optional) This option will allow to advertise on all available perimeter RE sites (`Bool`).

`re_sites` - (Optional) This option will allow advertise on specific Perimeter RE sites. See [Perimeter Re Choice Re Sites ](#perimeter-re-choice-re-sites) below for details.

### Perimeter Re Choice All Perimeter

This option will allow to advertise on all available perimeter RE sites.

### Perimeter Re Choice Re Sites

This option will allow advertise on specific Perimeter RE sites.

`perimeter_re_site` - (Optional) Selected Perimeter RE Site.. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_domain_view.
