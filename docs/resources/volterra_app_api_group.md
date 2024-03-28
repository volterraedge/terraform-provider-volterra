---

page_title: "Volterra: app_api_group"

description: "The app_api_group allows CRUD of App Api Group resource on Volterra SaaS"
---------------------------------------------------------------------------------------

Resource volterra_app_api_group
===============================

The App Api Group allows CRUD of App Api Group resource on Volterra SaaS

~> **Note:** Please refer to [App Api Group API docs](https://docs.cloud.f5.com/docs/api/views-app-api-group) to learn more

Example Usage
-------------

```hcl
resource "volterra_app_api_group" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  elements {
    methods = ["['PUT', 'POST', 'DELETE']"]

    path_regex = "/api/config/.*/path[123]/$"
  }

  // One of the arguments from this list "generic http_loadbalancer api_definition" must be set

  http_loadbalancer {
    http_loadbalancer {
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

`api_group_builder` - (Optional) API Group builder defines how to create API group from a list of endpoints. See [Api Group Builder ](#api-group-builder) below for details.(Deprecated)

`elements` - (Required) List of API group elements with methods and path regex for matching requests.. See [Elements ](#elements) below for details.

`api_definition` - (Optional) Set scope to an API Definition object to define the API endpoints list for API groups management. See [Scope Choice Api Definition ](#scope-choice-api-definition) below for details.(Deprecated)

`generic` - (Optional) The generic API Group is simply a list of API endpoints (`Bool`).(Deprecated)

`http_loadbalancer` - (Optional) Set scope to an HTTP Loadbalancer object to define the API endpoints list for API groups management. See [Scope Choice Http Loadbalancer ](#scope-choice-http-loadbalancer) below for details.

### Api Group Builder

API Group builder defines how to create API group from a list of endpoints.

`excluded_operations` - (Optional) The paths appear here with parameters as defined in OpenAPI spec file.. See [Api Group Builder Excluded Operations ](#api-group-builder-excluded-operations) below for details.

`included_operations` - (Optional) The paths appear here with parameters as defined in OpenAPI spec file.. See [Api Group Builder Included Operations ](#api-group-builder-included-operations) below for details.

`label_filter` - (Optional) In the current context a label is a property of an OpenAPI operation or path.. See [Api Group Builder Label Filter ](#api-group-builder-label-filter) below for details.

`metadata` - (Optional) Common attributes for the rule including name and description.. See [Api Group Builder Metadata ](#api-group-builder-metadata) below for details.(Deprecated)

`path_filter` - (Optional) The match is considered to succeed if the input request API path matches the specified path regex. (`String`).

### Elements

List of API group elements with methods and path regex for matching requests..

`methods` - (Required) The match is considered to succeed if the input request API method is a member of the list. (`List of Strings`).

`path_regex` - (Required) The match is considered to succeed if the input request API path matches the specified path regex. (`String`).

### Api Group Builder Excluded Operations

The paths appear here with parameters as defined in OpenAPI spec file..

`method` - (Required) Method to match the input request API method against. (`String`).

`path` - (Required) The path should comply with RFC 3986 and may have parameters according to OpenAPI specification (`String`).

### Api Group Builder Included Operations

The paths appear here with parameters as defined in OpenAPI spec file..

`method` - (Required) Method to match the input request API method against. (`String`).

`path` - (Required) The path should comply with RFC 3986 and may have parameters according to OpenAPI specification (`String`).

### Api Group Builder Label Filter

In the current context a label is a property of an OpenAPI operation or path..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Api Group Builder Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Scope Choice Api Definition

Set scope to an API Definition object to define the API endpoints list for API groups management.

`api_definition` - (Required) Reference to an API Definition object which defines a superset of API Endpoints for the API Group. See [ref](#ref) below for details.

### Scope Choice Http Loadbalancer

Set scope to an HTTP Loadbalancer object to define the API endpoints list for API groups management.

`http_loadbalancer` - (Required) Reference to an HTTP Loadbalancer object which defines a superset of API Endpoints for the API Group. See [ref](#ref) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured app_api_group.
