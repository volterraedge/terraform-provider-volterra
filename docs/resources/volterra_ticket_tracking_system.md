---

page_title: "Volterra: ticket_tracking_system"

description: "The ticket_tracking_system allows CRUD of Ticket Tracking System resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------------

Resource volterra_ticket_tracking_system
========================================

The Ticket Tracking System allows CRUD of Ticket Tracking System resource on Volterra SaaS

~> **Note:** Please refer to [Ticket Tracking System API docs](https://docs.cloud.f5.com/docs-v2/api/ticket-tracking-system) to learn more

Example Usage
-------------

```hcl
resource "volterra_ticket_tracking_system" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "jira_config" must be set

  jira_config {
    // One of the arguments from this list "adhoc_rest_api" can be set

    adhoc_rest_api {
      account_email = "account_email"

      api_token = "api_token"

      organization_domain = "sample-org.atlassian.net"
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

###### One of the arguments from this list "jira_config" must be set

`jira_config` - (Optional) Configuration when JIRA is the ticket provider system. See [Provider Config Jira Config ](#provider-config-jira-config) below for details.

### Config Adhoc Rest Api

x-displayName: "Ad-hoc REST API".

`account_email` - (Required) Username (email) for the Atlassian account (`String`).

`api_token` - (Required) Since this field is confidential, the Get/List public APIs will return this value as an empty string. (`String`).

`organization_domain` - (Required) A valid hostname for the Atlassian organization, as defined by RFC 1034 (`String`).

### Provider Config Jira Config

Configuration when JIRA is the ticket provider system.

###### One of the arguments from this list "adhoc_rest_api" can be set

`adhoc_rest_api` - (Optional) x-displayName: "Ad-hoc REST API". See [Config Adhoc Rest Api ](#config-adhoc-rest-api) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured ticket_tracking_system.
