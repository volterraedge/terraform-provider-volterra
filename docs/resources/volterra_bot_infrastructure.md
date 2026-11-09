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

`bot_allowlist_policy_metadata` - (Optional) x-displayName: "Bot Allowlist Policy Metadata". See [Bot Allowlist Policy Metadata ](#bot-allowlist-policy-metadata) below for details.

`bot_endpoint_policy_metadata` - (Optional) x-displayName: "Bot Endpoint Policy Metadata". See [Bot Endpoint Policy Metadata ](#bot-endpoint-policy-metadata) below for details.

`bot_network_policy_metadata` - (Optional) x-displayName: "Bot Network Policy Metadata". See [Bot Network Policy Metadata ](#bot-network-policy-metadata) below for details.

`cloned_infra_name` - (Optional) The cluster group from the source infrastructure will be copied to the new infrastructure. (`String`).(Deprecated)

###### One of the arguments from this list "create_cloud_hosted" can be set

`create_cloud_hosted` - (Optional) Associates an F5 Cloud Hosted service node to the Bot Defense infrastructure. See [Create Bot Infra Choice Create Cloud Hosted ](#create-bot-infra-choice-create-cloud-hosted) below for details.

`deployment_mode` - (Optional) By default, the mode will be Reverse Proxy. You need to submit a support ticket to request for API mode. (`String`).(Deprecated)

`traffic_type` - (Optional) The type of traffic that is routed to and processed by this infrastructure (Web or Mobile). (`String`).

### Bot Allowlist Policy Metadata

x-displayName: "Bot Allowlist Policy Metadata".

`name` - (Optional) Policy Name (`String`).

`version` - (Optional) Policy Version (`String`).

### Bot Endpoint Policy Metadata

x-displayName: "Bot Endpoint Policy Metadata".

`name` - (Optional) Policy Name (`String`).

`type` - (Optional) Policy Type (`String`).(Deprecated)

`version` - (Optional) Policy Version (`String`).

### Bot Network Policy Metadata

x-displayName: "Bot Network Policy Metadata".

`name` - (Optional) Policy Name (`String`).

`version` - (Optional) Policy Version (`String`).

### Create Bot Infra Choice Create Cloud Hosted

Associates an F5 Cloud Hosted service node to the Bot Defense infrastructure.

`host_name_list` - (Optional) Access Control List - Host Names. See [Create Cloud Hosted Host Name List ](#create-cloud-hosted-host-name-list) below for details.

`ip_address_list` - (Required) Access Control List - IP Addresses. See [Create Cloud Hosted Ip Address List ](#create-cloud-hosted-ip-address-list) below for details.

`ip_addresses` - (Optional) Only traffic from these IP addresses is allowed to access this Bot Defense infrastructure. (`String`).(Deprecated)

###### One of the arguments from this list "production, production_v2, testing, testing_v2" must be set

`production` - (Required) Production Protect live production applications and services from automated attacks.. See [Type Choice Production ](#type-choice-production) below for details.(Deprecated)

`production_v2` - (Required) Production Protect live production applications and services from automated attacks.. See [Type Choice Production V2 ](#type-choice-production-v2) below for details.

`testing` - (Required) Test Evaluate new features, settings, and system performance.. See [Type Choice Testing ](#type-choice-testing) below for details.(Deprecated)

`testing_v2` - (Required) Test Evaluate new features, settings, and system performance.. See [Type Choice Testing V2 ](#type-choice-testing-v2) below for details.

### Create Cloud Hosted Host Name List

Access Control List - Host Names.

`host_names` - (Optional) x-displayName: "Allowed Host Names" (`String`).

### Create Cloud Hosted Ip Address List

Access Control List - IP Addresses.

`ip_addresses` - (Required) x-displayName: "Ingress IP Address Control List" (`String`).

### Type Choice Production

Production Protect live production applications and services from automated attacks..

`region_1` - (Required) Traffic is routed equally between the two regions. (`String`).

`region_2` - (Required) Traffic is routed equally between the two regions. (`String`).

### Type Choice Production V2

Production Protect live production applications and services from automated attacks..

###### One of the arguments from this list "active_active_region, active_passive_region, single_active_region" can be set

`active_active_region` - (Required) Traffic routed equally to both regions.. See [Type Choice Production Active Active Region ](#type-choice-production-active-active-region) below for details.

`active_passive_region` - (Required) Traffic routed to Active region with failover to Passive on health failure.. See [Type Choice Production Active Passive Region ](#type-choice-production-active-passive-region) below for details.

`single_active_region` - (Required) Only one region serves traffic.. See [Type Choice Production Single Active Region ](#type-choice-production-single-active-region) below for details.

### Type Choice Testing

Test Evaluate new features, settings, and system performance..

`region_1` - (Required) Traffic is routed to a single region. (`String`).

### Type Choice Testing V2

Test Evaluate new features, settings, and system performance..

###### One of the arguments from this list "single_active_region" can be set

`single_active_region` - (Required) This is a single-active infrastructure configuration where traffic is routed to one region.. See [Type Choice Testing Single Active Region ](#type-choice-testing-single-active-region) below for details.

### Type Choice Production Active Active Region

Traffic routed equally to both regions..

`region_1` - (Required) x-required (`String`).

`region_2` - (Required) x-required (`String`).

### Type Choice Production Active Passive Region

Traffic routed to Active region with failover to Passive on health failure..

`region_1` - (Required) x-required (`String`).

`region_2` - (Required) x-required (`String`).

### Type Choice Production Single Active Region

Only one region serves traffic..

`region_1` - (Required) x-required (`String`).

### Type Choice Testing Single Active Region

This is a single-active infrastructure configuration where traffic is routed to one region..

`region_1` - (Required) x-required (`String`).

Attribute Reference
-------------------

*   `id` - This is the id of the configured bot_infrastructure.
