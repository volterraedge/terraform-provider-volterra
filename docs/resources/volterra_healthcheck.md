---

page_title: "Volterra: healthcheck"

description: "The healthcheck allows CRUD of Healthcheck resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_healthcheck
=============================

The Healthcheck allows CRUD of Healthcheck resource on Volterra SaaS

~> **Note:** Please refer to [Healthcheck API docs](https://volterra.io/docs/api/healthcheck) to learn more

Example Usage
-------------

```hcl
resource "volterra_healthcheck" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "http_health_check tcp_health_check" must be set

  http_health_check {
    headers = {
      "key1" = "value1"
    }

    // One of the arguments from this list "use_origin_server_name host_header" must be set
    use_origin_server_name = true
    path                   = "/healthcheck"

    request_headers_to_remove = ["user-agent"]
    use_http2                 = true
  }
  healthy_threshold   = ["2"]
  interval            = ["10"]
  timeout             = ["1"]
  unhealthy_threshold = ["5"]
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

`http_health_check` - (Optional) 4. Request headers to remove. See [Http Health Check ](#http-health-check) below for details.

`tcp_health_check` - (Optional) Specifies send payload and expected response payload. See [Tcp Health Check ](#tcp-health-check) below for details.

`healthy_threshold` - (Required) required to mark a host healthy. (`Int`).

`interval` - (Required) Time interval in seconds between two healthcheck requests. (`Int`).

`jitter_percent` - (Optional) Add a random amount of time as a percent value to the interval between successive healthcheck requests. (`Int`).

`timeout` - (Required) health check attempt will be considered a failure. (`Int`).

`unhealthy_threshold` - (Required) this threshold is ignored and the host is considered unhealthy immediately. (`Int`).

### Http Health Check

1.	Request headers to remove.

`headers` - (Optional) health checked cluster. This is a list of key-value pairs. (`String`).

`host_header` - (Optional) The value of the host header. (`String`).

`use_origin_server_name` - (Optional) Use the origin server name. (bool).

`path` - (Required) Specifies the HTTP path that will be requested during health checking. (`String`).

`request_headers_to_remove` - (Optional) health checked cluster. This is a list of keys of headers. (`String`).

`use_http2` - (Optional) If set, health checks will be made using http/2. (`Bool`).

### Tcp Health Check

Specifies send payload and expected response payload.

`expected_response` - (Optional) Hex encoded payload. (`String`).

`send_payload` - (Optional) Hex encoded payload. (`String`).

### Use Origin Server Name

Use the origin server name..

Attribute Reference
-------------------

-	`id` - This is the id of the configured healthcheck.
