---

page_title: "Volterra: route"

description: "The route allows CRUD of Route resource on Volterra SaaS"
-----------------------------------------------------------------------

Resource volterra_route
=======================

The Route allows CRUD of Route resource on Volterra SaaS

~> **Note:** Please refer to [Route API docs](https://volterra.io/docs/api/route) to learn more

Example Usage
-------------

```hcl
resource "volterra_route" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  routes {
    // One of the arguments from this list "route_destination route_redirect route_direct_response" must be set

    route_destination {
      // One of the arguments from this list "host_rewrite auto_host_rewrite" must be set
      host_rewrite = "one.volterra.com"

      buffer_policy {
        disabled          = true
        max_request_bytes = "2048"
        max_request_time  = "30"
      }

      cors_policy {
        allow_credentials = true
        allow_headers     = "allow_headers"
        allow_methods     = "allow_methods"

        allow_origin = ["allow_origin"]

        allow_origin_regex = ["allow_origin_regex"]
        disabled           = true
        expose_headers     = "expose_headers"
        max_age            = "max_age"
        maximum_age        = "maximum_age"
      }

      destinations {
        cluster {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }

        endpoint_subsets = {
          "key1" = "value1"
        }

        weight = "weight"
      }

      endpoint_subsets = {
        "key1" = "value1"
      }

      hash_policy {
        // One of the arguments from this list "header_name cookie source_ip" must be set
        header_name = "host"

        terminal = true
      }

      mirror_policy {
        cluster {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }

        percent {
          denominator = "denominator"
          numerator   = "5"
        }
      }

      prefix_rewrite = "/"
      priority       = "priority"

      retry_policy {
        back_off {
          base_interval = "base_interval"
          max_interval  = "max_interval"
        }

        num_retries     = "num_retries"
        per_try_timeout = "per_try_timeout"

        retriable_status_codes = ["retriable_status_codes"]
        retry_on               = "5xx"
      }

      spdy_config {
        use_spdy = true
      }

      timeout = "timeout"

      web_socket_config {
        idle_timeout         = "idle_timeout"
        max_connect_attempts = "max_connect_attempts"
        use_websocket        = true
      }
    }

    disable_custom_script = true
    disable_location_add  = true

    match {
      headers {
        invert_match = true
        name         = "Content-Type"

        // One of the arguments from this list "exact regex presence" must be set
        exact = "application/json"
      }

      http_method = "http_method"

      path {
        // One of the arguments from this list "prefix path regex" must be set
        path = "path"
      }

      query_params {
        key = "assignee_username"

        // One of the arguments from this list "exact regex" must be set
        exact = "exact"
      }
    }

    request_headers_to_add {
      append = true
      name   = "name"
      value  = "value"
    }

    request_headers_to_remove = ["host"]

    response_headers_to_add {
      append = true
      name   = "name"
      value  = "value"
    }

    response_headers_to_remove = ["host"]

    service_policy {
      disable = true
    }

    waf_type {
      // One of the arguments from this list "waf waf_rules" must be set

      waf_rules {
        waf_rules {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
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

`routes` - (Required) List of routes to match for incoming request. See [Routes ](#routes) below for details.

### Headers

List of (key, value) headers.

`invert_match` - (Optional) Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Optional) Name of the header (`String`).

`exact` - (Optional) Header value to match exactly (`String`).

`presence` - (Optional) If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Match

route match condition.

`headers` - (Optional) List of (key, value) headers. See [Headers ](#headers) below for details.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`path` - (Optional) URI path of route. See [Path ](#path) below for details.

`query_params` - (Optional) List of (key, value) query parameters. See [Query Params ](#query-params) below for details.

### Path

URI path of route.

`path` - (Optional) path : /logout (`String`).

`prefix` - (Optional) prefix : /register/ (`String`).

`regex` - (Optional) regex : /b[io]t (`String`).

### Query Params

List of (key, value) query parameters.

`key` - (Optional) In the above example, assignee_username is the key (`String`).

`exact` - (Optional) Exact match value for the query parameter key (`String`).

`regex` - (Optional) Regex match value for the query parameter key (`String`).

### Request Headers To Add

enclosing VirtualHost object level.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

`value` - (Required) Value of the HTTP header. (`String`).

### Response Headers To Add

enclosing VirtualHost object level.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

`value` - (Required) Value of the HTTP header. (`String`).

### Routes

List of routes to match for incoming request.

`route_destination` - (Optional) Send request to one of the destination from list of destinations. See [Route Destination ](#route-destination) below for details.

`route_direct_response` - (Optional) Send direct response. See [Route Direct Response ](#route-direct-response) below for details.

`route_redirect` - (Optional) Send redirect response. See [Route Redirect ](#route-redirect) below for details.

`disable_custom_script` - (Optional) disable execution of Javascript at route level, if it is configured at virtual-host level (`Bool`).

`disable_location_add` - (Optional) virtual-host level. This configuration is ignored on CE sites. (`Bool`).

`match` - (Optional) route match condition. See [Match ](#match) below for details.

`request_headers_to_add` - (Optional) enclosing VirtualHost object level. See [Request Headers To Add ](#request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) enclosing VirtualHost object level. See [Response Headers To Add ](#response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

`service_policy` - (Optional) service policy configuration at route level which overrides configuration at virtual-host level. See [Service Policy ](#service-policy) below for details.

`waf_type` - (Optional) waf_type specified at route level overrides waf configuration at VirtualHost level. See [Waf Type ](#waf-type) below for details.

### Service Policy

service policy configuration at route level which overrides configuration at virtual-host level.

`disable` - (Optional) disable service policy at route level, if it is configured at virtual-host level (`Bool`).

### Waf Type

waf_type specified at route level overrides waf configuration at VirtualHost level.

`waf` - (Optional) A WAF object direct reference. See [Waf ](#waf) below for details.

`waf_rules` - (Optional) A set of direct references of WAF Rules objects. See [Waf Rules ](#waf-rules) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured route.
