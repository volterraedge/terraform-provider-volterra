---

page_title: "Volterra: bot_infrastructure"

description: "The bot_infrastructure allows CRUD of Bot Infrastructure resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------

Resource volterra_bot_infrastructure
====================================

The Bot Infrastructure allows CRUD of Bot Infrastructure resource on Volterra SaaS

~> **Note:** Please refer to [Bot Infrastructure API docs](https://docs.cloud.f5.com/docs-v2/api/bot-infrastructure) to learn more

Example Usage
-------------

```hcl
resource "volterra_bot_infrastructure" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
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

###### One of the arguments from this list "create_cloud_hosted" can be set

`create_cloud_hosted` - (Optional) Associates an F5 Cloud Hosted service node to the Bot Defense infrastructure. See [Create Bot Infra Choice Create Cloud Hosted ](#create-bot-infra-choice-create-cloud-hosted) below for details.

`traffic_type` - (Optional) The type of traffic that is routed to and processed by this infrastructure (Web or Mobile). (`String`).

### Create Bot Infra Choice Create Cloud Hosted

Associates an F5 Cloud Hosted service node to the Bot Defense infrastructure.

`ip_addresses` - (Optional) Only traffic from these IP addresses is allowed to access this Bot Defense infrastructure. (`String`).

###### One of the arguments from this list "production, testing" must be set

`production` - (Required) Production Protect live production applications and services from automated attacks.. See [Type Choice Production ](#type-choice-production) below for details.

`testing` - (Required) Test Evaluate new features, settings, and system performance.. See [Type Choice Testing ](#type-choice-testing) below for details.

### Type Choice Production

Production Protect live production applications and services from automated attacks..

`region_1` - (Required) This is an Active-Active Infrastructure configuration where traffic is routed equally between the two regions. (`String`).

`region_2` - (Required) This is an Active-Active Infrastructure configuration where traffic is routed equally between the two regions. (`String`).

### Type Choice Testing

Test Evaluate new features, settings, and system performance..

`region_1` - (Required) This is an Active-Passive Infrastructure configuration where traffic is routed to a single region. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured bot_infrastructure.
