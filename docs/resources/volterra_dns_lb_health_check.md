---

page_title: "Volterra: dns_lb_health_check"

description: "The dns_lb_health_check allows CRUD of Dns Lb Health Check resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------

Resource volterra_dns_lb_health_check
=====================================

The Dns Lb Health Check allows CRUD of Dns Lb Health Check resource on Volterra SaaS

~> **Note:** Please refer to [Dns Lb Health Check API docs](https://docs.cloud.f5.com/docs-v2/api/dns-lb-health-check) to learn more

Example Usage
-------------

```hcl
resource "volterra_dns_lb_health_check" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "http_health_check https_health_check icmp_health_check tcp_health_check tcp_hex_health_check udp_health_check" must be set

  tcp_hex_health_check {
    health_check_port = "80"

    health_check_secondary_port = "443"

    receive = "00000034"

    send = "000000FF"
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

###### One of the arguments from this list "http_health_check, https_health_check, icmp_health_check, tcp_health_check, tcp_hex_health_check, udp_health_check" must be set

`http_health_check` - (Optional) HTTP Health Check. See [Health Check Http Health Check ](#health-check-http-health-check) below for details.

`https_health_check` - (Optional) HTTPS Health Check. See [Health Check Https Health Check ](#health-check-https-health-check) below for details.

`icmp_health_check` - (Optional) ICMP Health Check (`Bool`).

`tcp_health_check` - (Optional) TCP Health Check. See [Health Check Tcp Health Check ](#health-check-tcp-health-check) below for details.

`tcp_hex_health_check` - (Optional) TCP Health Check with Hex Encoded Payload. See [Health Check Tcp Hex Health Check ](#health-check-tcp-hex-health-check) below for details.

`udp_health_check` - (Optional) UDP Health Check. See [Health Check Udp Health Check ](#health-check-udp-health-check) below for details.

### Health Check Http Health Check

HTTP Health Check.

`health_check_port` - (Required) x-example: "80" (`Int`).

`health_check_secondary_port` - (Optional) x-example: "443" (`Int`).

`receive` - (Optional) Regular expression used to match against the response to the health check's request. Mark node up upon receipt of a successful regular expression match. Uses re2 regular expression syntax. (`String`).

`send` - (Optional) HTTP payload to send to the target (`String`).

### Health Check Https Health Check

HTTPS Health Check.

`health_check_port` - (Required) x-example: "80" (`Int`).

`health_check_secondary_port` - (Optional) x-example: "443" (`Int`).

`receive` - (Optional) Regular expression used to match against the response to the health check's request. Mark node up upon receipt of a successful regular expression match. Uses re2 regular expression syntax. (`String`).

`send` - (Optional) HTTP payload to send to the target (`String`).

### Health Check Tcp Health Check

TCP Health Check.

`health_check_port` - (Required) x-example: "80" (`Int`).

`health_check_secondary_port` - (Optional) x-example: "443" (`Int`).

`receive` - (Optional) Regular expression used to match against the response to the monitor's request. Mark node up upon receipt of a successful regular expression match. Uses re2 regular expression syntax. (`String`).

`send` - (Optional) Send this string to target (default empty. When send and receive are both empty, monitor just tests 3WHS) (`String`).

### Health Check Tcp Hex Health Check

TCP Health Check with Hex Encoded Payload.

`health_check_port` - (Required) x-example: "80" (`Int`).

`health_check_secondary_port` - (Optional) x-example: "443" (`Int`).

`receive` - (Optional) Hex encoded raw bytes expected in the response. (`String`).

`send` - (Optional) Hex encoded raw bytes sent in the request. Empty payloads imply a connect-only health check. (`String`).

### Health Check Udp Health Check

UDP Health Check.

`health_check_port` - (Required) x-example: "80" (`Int`).

`health_check_secondary_port` - (Optional) x-example: "443" (`Int`).

`receive` - (Required) UDP response to be matched. It can be a regex. (`String`).

`send` - (Required) UDP payload (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured dns_lb_health_check.
