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

  // One of the arguments from this list "dns_proxy_icmp_health_check http_health_check tcp_health_check dns_proxy_tcp_health_check dns_proxy_udp_health_check dns_health_check" must be set

  http_health_check {
    expected_status_codes = ["200-250"]

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

`dns_health_check` - (Optional) 4. Expected IP Address. See [Dns Health Check ](#dns-health-check) below for details.

`dns_proxy_icmp_health_check` - (Optional) Specifies ICMP HealthCheck (bool).

`dns_proxy_tcp_health_check` - (Optional) Specifies send string and expected response payload pattern for TCP health Check. See [Dns Proxy Tcp Health Check ](#dns-proxy-tcp-health-check) below for details.

`dns_proxy_udp_health_check` - (Optional) Specifies send string and expected response payload pattern for UDP health Check. See [Dns Proxy Udp Health Check ](#dns-proxy-udp-health-check) below for details.

`http_health_check` - (Optional) 4. Request headers to remove. See [Http Health Check ](#http-health-check) below for details.

`tcp_health_check` - (Optional) Specifies send payload and expected response payload. See [Tcp Health Check ](#tcp-health-check) below for details.

`healthy_threshold` - (Required) required to mark a host healthy. (`Int`).

`interval` - (Required) Time interval in seconds between two healthcheck requests. (`Int`).

`jitter_percent` - (Optional) Add a random amount of time as a percent value to the interval between successive healthcheck requests. (`Int`).

`timeout` - (Required) health check attempt will be considered a failure. (`Int`).

`unhealthy_threshold` - (Required) this threshold is ignored and the host is considered unhealthy immediately. (`Int`).

### Dns Health Check

1.	Expected IP Address.

`expected_rcode` - (Required) Specifies an expected Rcode in the answer section of DNS Response, option [no-error, any](`String`).

`expected_record_type` - (Required) options: [REQUESTED_QUERY_TYPE, RECORD_TYPE_ANY] when REQUESTED_QUERY_TYPE is set, health monitor expects record type same as requested query type (`String`).

`expected_response` - (Required) Specifies an IPv4 or IPv6 address in the answer section of DNS Response (`String`).

`query_name` - (Required) The query name that the monitor sends a DNS query for. (`String`).

`query_type` - (Required) The DNS query type that the monitor sends. Supported types are: [a, aaaa] default: a (`String`).

`reverse` - (Optional) string match marks the monitored object down instead of up. (`Bool`).

### Dns Proxy Tcp Health Check

Specifies send string and expected response payload pattern for TCP health Check.

`expected_response` - (Required) Specifies a regular expression pattern which will be matched against response payload (`String`).

`send_payload` - (Required) Text string sent in the request (`String`).

### Dns Proxy Udp Health Check

Specifies send string and expected response payload pattern for UDP health Check.

`expected_response` - (Required) Specifies a regular expression pattern which will be matched against response payload (`String`).

`send_payload` - (Required) Text string sent in the request (`String`).

### Http Health Check

1.	Request headers to remove.

`expected_status_codes` - (Optional) of which is single HTTP status code or a range with start and end values separated by "-". (`String`).

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
