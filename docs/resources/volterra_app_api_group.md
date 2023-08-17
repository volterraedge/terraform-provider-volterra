---

page_title: "Volterra: app_api_group"

description: "The app_api_group allows CRUD of App Api Group resource on Volterra SaaS"
---------------------------------------------------------------------------------------

Resource volterra_app_api_group
===============================

The App Api Group allows CRUD of App Api Group resource on Volterra SaaS

~> **Note:** Please refer to [App Api Group API docs](https://volterra.io/docs/api/views-app-api-group) to learn more

Example Usage
-------------

```hcl
resource "volterra_app_api_group" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  elements {
    methods    = ["['PUT', 'POST', 'DELETE']"]
    path_regex = "/api/config/.*/path[123]/$"
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

`elements` - (Required) List of API group elements with methods and path regex for matching requests.. See [Elements ](#elements) below for details.

### Elements

List of API group elements with methods and path regex for matching requests..

`methods` - (Required) The match is considered to succeed if the input request API method is a member of the list. (`List of Strings`).

`path_regex` - (Required) The match is considered to succeed if the input request API path matches the specified path regex. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured app_api_group.
