---

page_title: "Volterra: uztna_application"

description: "The uztna_application allows CRUD of Uztna Application resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_uztna_application
===================================

The Uztna Application allows CRUD of Uztna Application resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Application API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-application) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_application" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  app_tags = ["app_tags"]

  location {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }

  port = ["443"]
  url  = ["test.f5net.com"]

  uztna_domain_ref {
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

`app_tags` - (Required) ZTNA Application tags to associate with Application (`List of String`).

`app_vh` - (Optional) Internal reference to Application VH. See [ref](#ref) below for details.(Deprecated)

`application_vip` - (Optional) for the Application on TMM. See [Application Vip ](#application-vip) below for details.

`extended_app_tags` - (Optional) Application tags updated from the Application group (`List of String`).

`location` - (Required) BIG-IP Next site to deploy the Application onto. See [ref](#ref) below for details.

`origin` - (Optional) SNAT pool is used to configure the SNAT Pool IPs. See [Origin ](#origin) below for details.

`port` - (Required) The Port on which the Application listens on (`Int`).

`protocol` - (Optional) Protocol used by the Application for communication. See [Protocol ](#protocol) below for details.

`url` - (Required) Hostname of the Application to be onboarded (`String`).

`uztna_domain_ref` - (Required) Reference to the UZTNA Domain Non-View Object. See [ref](#ref) below for details.

### Application Vip

for the Application on TMM.

###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Origin

SNAT pool is used to configure the SNAT Pool IPs.

`origin_pool` - (Optional) of the Application. See [ref](#ref) below for details.

`snat_pool` - (Optional) that another address is used from the pool. See [ref](#ref) below for details.

### Protocol

Protocol used by the Application for communication.

###### One of the arguments from this list "HTTP, HTTPS, TCP, UDP" must be set

`HTTP` - (Optional) x-displayName: "HTTP" (`Bool`).

`HTTPS` - (Optional) x-displayName: "HTTPS". See [Protocol Choice HTTPS ](#protocol-choice-HTTPS) below for details.

`TCP` - (Optional) x-displayName: "TCP" (`Bool`).

`UDP` - (Optional) x-displayName: "UDP" (`Bool`).

### Protocol Choice HTTPS

x-displayName: "HTTPS".

`certificate` - (Optional) Select/Add a TLS Certificate object to associate with this Application. See [ref](#ref) below for details.

### Protocol Choice HTTP

x-displayName: "HTTP".

### Protocol Choice TCP

x-displayName: "TCP".

### Protocol Choice UDP

x-displayName: "UDP".

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Ver Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ver Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_application.
