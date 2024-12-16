---

page_title: "Volterra: uztna_application_group"

description: "The uztna_application_group allows CRUD of Uztna Application Group resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------------

Resource volterra_uztna_application_group
=========================================

The Uztna Application Group allows CRUD of Uztna Application Group resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Application Group API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-application-group) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_application_group" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  app_tags {
    apptag_list {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  apps {
    app_list {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
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

`app_tags` - (Required) List of Valid Application Tags to be tagged for the list of Applications. See [App Tags ](#app-tags) below for details.

`apps` - (Required) List of Applications to be Grouped. . See [Apps ](#apps) below for details.

### App Tags

List of Valid Application Tags to be tagged for the list of Applications.

`apptag_list` - (Required) Application Tag for Application Group. See [ref](#ref) below for details.

### Apps

List of Applications to be Grouped. .

`app_list` - (Required) Select/Add one or more Application objects to associate with this Application Group. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_application_group.
