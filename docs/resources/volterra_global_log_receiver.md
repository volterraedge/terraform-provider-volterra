---

page_title: "Volterra: global_log_receiver"

description: "The global_log_receiver allows CRUD of Global Log Receiver resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------

Resource volterra_global_log_receiver
=====================================

The Global Log Receiver allows CRUD of Global Log Receiver resource on Volterra SaaS

~> **Note:** Please refer to [Global Log Receiver API docs](https://volterra.io/docs/api/global-log-receiver) to learn more

Example Usage
-------------

```hcl
resource "volterra_global_log_receiver" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "ns_current ns_all ns_list ns_system" must be set

  ns_list {
    namespaces = ["default"]
  }

  // One of the arguments from this list "s3_receiver" must be set

  s3_receiver {
    aws_cred {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }

    aws_region = "us-east-1"
    bucket     = "S3 Buket Name"
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

`ns_all` - (Optional) Select logs from all namespaces (bool).

`ns_current` - (Optional) Select logs from current namespace (bool).

`ns_list` - (Optional) Select logs in specific namespaces. See [Ns List ](#ns-list) below for details.

`ns_system` - (Optional) Select logs from System namespace (bool).

`s3_receiver` - (Optional) S3 Receiver. See [S3 Receiver ](#s3-receiver) below for details.

### Ns List

Select logs in specific namespaces.

`namespaces` - (Required) List of namespaces to stream logs for (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### S3 Receiver

S3 Receiver.

`aws_cred` - (Required) Reference to AWS credentials for access to S3 bucket. See [ref](#ref) below for details.

`aws_region` - (Required) Name for AWS Region. (`String`).

`bucket` - (Required) S3 Bucket Name (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured global_log_receiver.
