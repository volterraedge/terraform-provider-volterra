---

page_title: "Volterra: cloud_elastic_ip"

description: "The cloud_elastic_ip allows CRUD of Cloud Elastic Ip resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_cloud_elastic_ip
==================================

The Cloud Elastic Ip allows CRUD of Cloud Elastic Ip resource on Volterra SaaS

~> **Note:** Please refer to [Cloud Elastic Ip API docs](https://docs.cloud.f5.com/docs-v2/api/cloud-elastic-ip) to learn more

Example Usage
-------------

```hcl
resource "volterra_cloud_elastic_ip" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  site_ref {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
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

`count` - (Optional) number of Elastic Ips / Public Ips associated with this object per Node (`Int`).

`site_ref` - (Required) Site to which this cloud elastic ip object is attached. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured cloud_elastic_ip.
