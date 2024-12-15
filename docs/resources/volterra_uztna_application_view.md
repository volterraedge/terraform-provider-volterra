---

page_title: "Volterra: uztna_application_view"

description: "The uztna_application_view allows CRUD of Uztna Application View resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------------

Resource volterra_uztna_application_view
========================================

The Uztna Application View allows CRUD of Uztna Application View resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Application View API docs](https://docs.cloud.f5.com/docs-v2/api/views-uztna-application-view) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_application_view" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  app_tags {
    app_tags {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  fqdn = ["test.f5net.com"]

  location {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }

  port = ["443"]
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

`app_tags` - (Required) List of Tags to be associated with the Application. See [App Tags ](#app-tags) below for details.

`fqdn` - (Required) Hostname of the application to be onboarded. (`String`).

`location` - (Required) BIG-IP Next sites to onboard the Application. See [ref](#ref) below for details.

`origin` - (Optional) SNAT pool is an advanced field which is used to configure the SNAT pool IPs. See [Origin ](#origin) below for details.

`port` - (Required) The port on which the application listens . (`Int`).

`protocol` - (Optional) Protocol used by the Application for communication. See [Protocol ](#protocol) below for details.

### App Tags

List of Tags to be associated with the Application.

`app_tags` - (Required) List of Tags to be associated with the Application.. See [ref](#ref) below for details.

### Origin

SNAT pool is an advanced field which is used to configure the SNAT pool IPs.

`origin_pool` - (Optional) to which the BIG-IP Next will distribute the application traffic. See [ref](#ref) below for details.

`snat_pool` - (Optional) connections with unique ports.. See [ref](#ref) below for details.

### Protocol

Protocol used by the Application for communication.

###### One of the arguments from this list "HTTP, HTTPS, TCP, UDP" must be set

`HTTP` - (Optional) x-displayName: "HTTP" (`Bool`).

`HTTPS` - (Optional) Automatically from XC or User provided certificate.. See [Protocol Choice HTTPS ](#protocol-choice-HTTPS) below for details.

`TCP` - (Optional) x-displayName: "TCP" (`Bool`).

`UDP` - (Optional) x-displayName: "UDP" (`Bool`).

### Protocol Choice HTTPS

Automatically from XC or User provided certificate..

`certificate` - (Optional) Select/Add TLS Certificate objects to associate with this Application. See [ref](#ref) below for details.

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

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_application_view.
