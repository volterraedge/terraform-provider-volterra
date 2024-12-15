---

page_title: "Volterra: volterra_public_ip" 

description: "The volterra_public_ip helps update public IP fields on Volterra SaaS"

---

Resource volterra_public_ip
===========================

=============================

volterra_public_ip helps to update fields of site object. This is not applicable for site created through view style APIs.

~> **Note:** Please refer to [Public IP api docs](https://docs.cloud.f5.com/docs/api/public-ip) to learn more

Example Usage
-------------

---

```hcl
resource "volterra_public_ip" "example" {
  name      = "ip1"
  virtual_sites  {
    "name" = "fleet-car"
    "namespace" = "shared"
  }
}

```

### Argument Reference

---

`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).

`description` - (Optional) Human readable description for the object (`String`).

`disable` - (Optional) A value of true will administratively disable the object (`Bool`).

`labels` - (Optional) by selector expression (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

`namespace` - (Required) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).

`address` - (Optional) Site's geographical address that can be used determine its latitude and longitude. Example: "123 Street, city, country, postal code" (`String`).

`virtual_sites` - (Optional) Reference to virtual sites configuration objects. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String)

Attribute Reference
-------------------

---

-	`id` - This is the id of Volterra site object.
