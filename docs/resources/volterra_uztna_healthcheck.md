---

page_title: "Volterra: uztna_healthcheck"

description: "The uztna_healthcheck allows CRUD of Uztna Healthcheck resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------

Resource volterra_uztna_healthcheck
===================================

The Uztna Healthcheck allows CRUD of Uztna Healthcheck resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Healthcheck API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-healthcheck) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_healthcheck" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "tcp_health_check" must be set

  tcp_health_check {
    expected_response = ".*"

    send_payload = "send_payload"
  }
  interval = ["10"]
  timeout = ["1"]
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

###### One of the arguments from this list "tcp_health_check" must be set

`tcp_health_check` - (Optional) Specifies send payload and expected response payload. See [Health Check Tcp Health Check ](#health-check-tcp-health-check) below for details.

`interval` - (Required) Time interval in seconds between two healthcheck requests. (`Int`).

`timeout` - (Required) health check attempt will be considered a failure. (`Int`).

### Health Check Tcp Health Check

Specifies send payload and expected response payload.

`expected_response` - (Optional) Specifies a regular expression pattern which will be matched against response payload (`String`).

`send_payload` - (Optional) Text string sent in the request (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_healthcheck.
