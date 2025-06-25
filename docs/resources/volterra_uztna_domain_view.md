---

page_title: "Volterra: uztna_domain_view"

description: "The uztna_domain_view allows CRUD of Uztna Domain View resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_uztna_domain_view
===================================

The Uztna Domain View allows CRUD of Uztna Domain View resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Domain View API docs](https://docs.cloud.f5.com/docs-v2/api/views-uztna-domain-view) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_domain_view" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  access_url = ["access_url"]

  cert {
    certificate {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  gateways {
    perimeter_re {
      // One of the arguments from this list "all_cloud" can be set

      all_cloud = true
    }

    private_gateway {
      uztna_gateway {
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

`access_url` - (Required) This URL would resolve to the Cloud or the Private Gateway (`String`).

`app_vip_pool` - (Optional) Configure this option only if the default range is unacceptable.. See [App Vip Pool ](#app-vip-pool) below for details.

`cert` - (Required) and the server on Cloud and Private gateways. See [Cert ](#cert) below for details.

`gateways` - (Required) List of all RE prime and Big Ip edge CE. See [Gateways ](#gateways) below for details.

`lease_pool` - (Optional) range using the leasepool for this assignment. See [Lease Pool ](#lease-pool) below for details.

`policy` - (Optional) The name of the ZTNA profile. See [Policy ](#policy) below for details.

### App Vip Pool

Configure this option only if the default range is unacceptable..

###### One of the arguments from this list "ipv4_app_vip_pool" can be set

`ipv4_app_vip_pool` - (Required) Select or create new IPv4 App VIP pools. See [ref](#ref) below for details.

### Cert

and the server on Cloud and Private gateways.

`certificate` - (Optional) Select/Add one or more TLS Certificate objects to associate with this ZeroTrust Domain. See [ref](#ref) below for details.

### Gateways

List of all RE prime and Big Ip edge CE.

`perimeter_re` - (Optional) Cloud Gateways and Big-IP Edge Gateways. See [Gateways Perimeter Re ](#gateways-perimeter-re) below for details.

`private_gateway` - (Optional) Private gateways are gateways hosted within the customer's data centers and are typically accessed by users connected to their office or corporate network.. See [Gateways Private Gateway ](#gateways-private-gateway) below for details.

### Lease Pool

range using the leasepool for this assignment.

###### One of the arguments from this list "ipv4_leasepool" must be set

`ipv4_leasepool` - (Required) Select or create new IPv4 leasepool. See [ref](#ref) below for details.

### Policy

The name of the ZTNA profile.

`policy_name` - (Optional) The ZeroTrust Domain enforces an Access policy that all end users must comply with to access private and potentially public applications. The ZTNA policy allows the admin to set up the access policy for this ZeroTrust Domain. See [ref](#ref) below for details.

### Cloud Gateway Choice All Cloud

Advertise on all Cloud Gateways.

### Gateways Perimeter Re

Cloud Gateways and Big-IP Edge Gateways.

###### One of the arguments from this list "all_cloud" can be set

`all_cloud` - (Optional) Advertise on all Cloud Gateways (`Bool`).

### Gateways Private Gateway

Private gateways are gateways hosted within the customer's data centers and are typically accessed by users connected to their office or corporate network..

`uztna_gateway` - (Optional) x-example: "system/alon-ge". See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_domain_view.
